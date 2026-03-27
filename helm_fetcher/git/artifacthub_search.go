package git

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/MabsIPCA/helm-tests/helm_fetcher/model"
	"github.com/rs/zerolog/log"
)

func extractRepoURL(pkg model.ArtifactHubPackage) string {
	for _, l := range pkg.Links {
		if strings.EqualFold(l.Name, "source") {
			if normalized := normalizeGitHubRepoURL(l.URL); normalized != "" {
				return normalized
			}
		}
	}
	for _, l := range pkg.Links {
		if normalized := normalizeGitHubRepoURL(l.URL); normalized != "" {
			return normalized
		}
	}
	return normalizeGitHubRepoURL(pkg.Repository.URL)
}

var artifactHubGitHubURLPattern = regexp.MustCompile(`https?://(?:www\.)?github\.com/[^"\s]+`)

func extractRepoURLFromRawArtifactPackage(rawPkg json.RawMessage) string {
	matches := artifactHubGitHubURLPattern.FindAllString(string(rawPkg), -1)
	for _, candidate := range matches {
		if normalized := normalizeGitHubRepoURL(candidate); normalized != "" {
			return normalized
		}
	}
	return ""
}

// SearchTopArtifactHubRepos fetches top packages from Artifact Hub sorted by stars
// and returns unique GitHub repository URLs.
func SearchTopArtifactHubRepos(limit, pageSize int) ([]string, error) {
	if limit <= 0 {
		return nil, fmt.Errorf("limit must be > 0")
	}
	if pageSize <= 0 || pageSize > 60 {
		return nil, fmt.Errorf("pageSize must be between 1 and 60 for artifacthub")
	}

	seen := map[string]bool{}
	var repos []string
	var skippedNoRepo int

	for offset := 0; len(repos) < limit; offset += pageSize {
		remaining := limit - len(repos)
		currentLimit := pageSize
		if remaining < currentLimit {
			currentLimit = remaining
		}

		requestURL := fmt.Sprintf(
			"https://artifacthub.io/api/v1/packages/search?kind=0&sort=stars&limit=%d&offset=%d",
			currentLimit, offset,
		)
		log.Info().Str("url", requestURL).Int("offset", offset).Int("batch_limit", currentLimit).Msg("Searching Artifact Hub via curl")

		body, err := curlAPI(requestURL)
		if err != nil {
			return repos, fmt.Errorf("offset %d: %w", offset, err)
		}

		var result model.ArtifactHubSearchResponse
		if err := json.Unmarshal(body, &result); err != nil {
			return repos, fmt.Errorf("JSON decode offset %d: %w – body snippet: %.300s", offset, err, string(body))
		}
		var rawResult struct {
			Packages []json.RawMessage `json:"packages"`
		}
		if err := json.Unmarshal(body, &rawResult); err != nil {
			return repos, fmt.Errorf("raw JSON decode offset %d: %w – body snippet: %.300s", offset, err, string(body))
		}
		if len(result.Packages) == 0 {
			break
		}

		batchNewRepos := 0
		for i, pkg := range result.Packages {
			repoURL := extractRepoURL(pkg)
			if repoURL == "" && i < len(rawResult.Packages) {
				repoURL = extractRepoURLFromRawArtifactPackage(rawResult.Packages[i])
			}
			if repoURL == "" {
				skippedNoRepo++
				continue
			}
			if !seen[repoURL] {
				seen[repoURL] = true
				repos = append(repos, repoURL)
				batchNewRepos++
				if len(repos) >= limit {
					break
				}
			}
		}

		log.Info().
			Int("packages_in_batch", len(result.Packages)).
			Int("new_repos_in_batch", batchNewRepos).
			Int("unique_repos_so_far", len(repos)).
			Int("skipped_without_repo", skippedNoRepo).
			Msg("Artifact Hub batch complete")
	}

	log.Info().Int("unique_repos", len(repos)).Int("skipped_without_repo", skippedNoRepo).Msg("Artifact Hub search finished")
	return repos, nil
}
