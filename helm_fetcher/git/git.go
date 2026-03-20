package git

import (
	"encoding/json"
	"fmt"
	neturl "net/url"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/MabsIPCA/helm-tests/helm_fetcher/model"
	"github.com/rs/zerolog/log"
)

// curlAPI performs a GET request to the specified URL using curl.
func curlAPI(url string, headers ...string) ([]byte, error) {
	args := []string{
		"-s", "-S",
		"--max-time", "30",
		"-H", "User-Agent: helm-fetcher/1.0",
		"-H", "Accept: application/json",
	}
	for _, header := range headers {
		args = append(args, "-H", header)
	}
	args = append(args, url)
	cmd := exec.Command("curl", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("curl failed: %w – output: %s", err, string(out))
	}
	return out, nil
}

func sanitizeGitHubURLCandidate(raw string) string {
	trimmed := strings.TrimSpace(raw)
	// Remove common trailing punctuation/noise from free-text URLs.
	return strings.TrimRight(trimmed, ")]}>,.;:'\"`|\\")
}

var githubNamePattern = regexp.MustCompile(`^[A-Za-z0-9_.-]+$`)

func normalizeGitHubRepoURL(raw string) string {
	raw = sanitizeGitHubURLCandidate(raw)
	if raw == "" {
		return ""
	}
	u, err := neturl.Parse(raw)
	if err != nil || u.Host == "" {
		return ""
	}
	host := strings.ToLower(u.Hostname())
	if host != "github.com" && host != "www.github.com" {
		return ""
	}
	parts := strings.Split(strings.Trim(u.Path, "/"), "/")
	if len(parts) < 2 {
		return ""
	}
	owner := parts[0]
	repo := strings.TrimSuffix(parts[1], ".git")
	repo = strings.TrimSuffix(repo, "/")
	repo = strings.TrimSuffix(repo, ".")
	if owner == "" || repo == "" {
		return ""
	}
	if !githubNamePattern.MatchString(owner) || !githubNamePattern.MatchString(repo) {
		return ""
	}
	return fmt.Sprintf("https://github.com/%s/%s", owner, repo)
}

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

// SearchTopGitHubRepos searches GitHub code for Chart.yaml files and returns unique repositories.
func SearchTopGitHubRepos(limit, pageSize int) ([]string, error) {
	if limit <= 0 {
		return nil, fmt.Errorf("limit must be > 0")
	}
	if pageSize <= 0 || pageSize > 100 {
		return nil, fmt.Errorf("pageSize must be between 1 and 100")
	}

	seen := map[string]bool{}
	var repos []string
	maxPages := (limit + pageSize - 1) / pageSize

	for page := 1; page <= maxPages && len(repos) < limit; page++ {
		remaining := limit - len(repos)
		currentPageSize := pageSize
		if remaining < currentPageSize {
			currentPageSize = remaining
		}

		requestURL := fmt.Sprintf(
			"https://api.github.com/search/code?q=filename:Chart.yaml&per_page=%d&page=%d",
			currentPageSize, page,
		)
		log.Info().Str("url", requestURL).Int("page", page).Int("page_size", currentPageSize).Msg("Searching GitHub via curl")

		body, err := curlAPI(requestURL, "Accept: application/vnd.github+json")
		if err != nil {
			return repos, fmt.Errorf("page %d: %w", page, err)
		}

		var result model.GitHubSearchResult
		if err := json.Unmarshal(body, &result); err != nil {
			return repos, fmt.Errorf("JSON decode page %d: %w – body snippet: %.300s", page, err, string(body))
		}
		if len(result.Items) == 0 {
			break
		}

		newReposInPage := 0
		for _, item := range result.Items {
			repoURL := normalizeGitHubRepoURL(item.Repository.HTMLURL)
			if repoURL == "" && item.Repository.FullName != "" {
				repoURL = normalizeGitHubRepoURL("https://github.com/" + item.Repository.FullName)
			}
			if repoURL == "" {
				continue
			}
			if !seen[repoURL] {
				seen[repoURL] = true
				repos = append(repos, repoURL)
				newReposInPage++
				if len(repos) >= limit {
					break
				}
			}
		}

		log.Info().
			Int("new_repos_in_page", newReposInPage).
			Int("unique_repos_so_far", len(repos)).
			Msg("GitHub page complete")
	}

	log.Info().Int("unique_repos", len(repos)).Msg("GitHub search finished")
	return repos, nil
}

// CloneRepo clones the given repo URL into the destination directory using "git clone --depth 1".
func CloneRepo(repoURL, destDir string) error {
	if info, err := os.Stat(destDir); err == nil && info.IsDir() {
		log.Info().Str("dir", destDir).Msg("Already cloned – skipping")
		return nil
	}
	log.Info().Str("repo", repoURL).Str("dest", destDir).Msg("Cloning (depth=1)")
	cmd := exec.Command("git", "clone", "--depth", "1", repoURL+".git", destDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
