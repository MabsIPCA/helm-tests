package git

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/MabsIPCA/helm-tests/helm_fetcher/model"
	"github.com/rs/zerolog/log"
)

const (
	GitHubMaxCandidateRepos = 50000
	GitHubTopReposToReturn  = 1000
	OrderAsc                = "asc"
	OrderDesc               = "desc"
)

type GitHubRankedRepo struct {
	RepoURL string `json:"repo_url"`
	Stars   int    `json:"stars"`
}

type scoredRepo struct {
	url   string
	stars int
}

func githubAuthHeader() string {
	token := strings.TrimSpace(os.Getenv("GITHUB_TOKEN"))
	if token == "" {
		return ""
	}
	return "Authorization: Bearer " + token
}

func fetchRepoStars(repoURL string) (int, error) {
	repoPath := strings.TrimPrefix(repoURL, "https://github.com/")
	parts := strings.Split(repoPath, "/")
	if len(parts) < 2 || parts[0] == "" || parts[1] == "" {
		return 0, fmt.Errorf("invalid repo URL: %s", repoURL)
	}

	requestURL := fmt.Sprintf("https://api.github.com/repos/%s/%s", parts[0], parts[1])
	headers := []string{"Accept: application/vnd.github+json"}
	if auth := githubAuthHeader(); auth != "" {
		headers = append(headers, auth)
	}

	body, err := curlAPI(requestURL, headers...)
	if err != nil {
		return 0, err
	}

	var repoInfo struct {
		StargazersCount int    `json:"stargazers_count"`
		Message         string `json:"message"`
	}
	if err := json.Unmarshal(body, &repoInfo); err != nil {
		return 0, fmt.Errorf("decode repo stars failed: %w", err)
	}
	if repoInfo.Message != "" && repoInfo.StargazersCount == 0 {
		return 0, fmt.Errorf("github API error: %s", repoInfo.Message)
	}

	return repoInfo.StargazersCount, nil
}

func rankAndTrimReposByStars(candidates []scoredRepo, topN int, order string) []scoredRepo {
	descending := order != OrderAsc

	sort.Slice(candidates, func(i, j int) bool {
		if candidates[i].stars == candidates[j].stars {
			return candidates[i].url < candidates[j].url
		}
		if descending {
			return candidates[i].stars > candidates[j].stars
		}
		return candidates[i].stars < candidates[j].stars
	})

	if topN > len(candidates) {
		topN = len(candidates)
	}

	ranked := make([]scoredRepo, 0, topN)
	for i := 0; i < topN; i++ {
		ranked = append(ranked, candidates[i])
	}
	return ranked
}

// SearchGitHubRepos searches GitHub code for Chart.yaml files, collects up to 50k unique repos,
// ranks them by stars, and returns the requested amount in the selected order.
func SearchGitHubRepos(pageSize, limit int, order string) ([]GitHubRankedRepo, error) {
	if pageSize <= 0 || pageSize > 100 {
		return nil, fmt.Errorf("pageSize must be between 1 and 100")
	}
	if limit <= 0 {
		return nil, fmt.Errorf("limit must be > 0")
	}
	if order != OrderAsc && order != OrderDesc {
		return nil, fmt.Errorf("order must be '%s' or '%s'", OrderAsc, OrderDesc)
	}

	seen := map[string]bool{}
	var candidates []scoredRepo
	starLookupFailures := 0
	pagesFetched := 0

	headers := []string{"Accept: application/vnd.github+json"}
	if auth := githubAuthHeader(); auth != "" {
		headers = append(headers, auth)
	} else {
		log.Warn().Msg("GITHUB_TOKEN not set; GitHub API rate limits may reduce coverage")
	}

	for page := 1; len(candidates) < GitHubMaxCandidateRepos; page++ {
		pagesFetched++

		requestURL := fmt.Sprintf(
			"https://api.github.com/search/code?q=filename:Chart.yaml&per_page=%d&page=%d",
			pageSize, page,
		)
		log.Info().
			Str("url", requestURL).
			Int("page", page).
			Int("page_size", pageSize).
			Int("unique_repos_so_far", len(candidates)).
			Msg("Searching GitHub via curl")

		body, err := curlAPI(requestURL, headers...)
		if err != nil {
			ranked := rankAndTrimReposByStars(candidates, limit, order)
			return toGitHubRankedRepos(ranked), fmt.Errorf("page %d: %w", page, err)
		}

		var result model.GitHubSearchResult
		if err := json.Unmarshal(body, &result); err != nil {
			ranked := rankAndTrimReposByStars(candidates, limit, order)
			return toGitHubRankedRepos(ranked), fmt.Errorf("JSON decode page %d: %w – body snippet: %.300s", page, err, string(body))
		}
		if len(result.Items) == 0 {
			log.Info().Int("page", page).Msg("GitHub returned no items; stopping pagination")
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
				stars, starErr := fetchRepoStars(repoURL)
				if starErr != nil {
					starLookupFailures++
					log.Warn().Err(starErr).Str("repo", repoURL).Msg("Failed to fetch repo stars; using 0")
					stars = 0
				}

				candidates = append(candidates, scoredRepo{url: repoURL, stars: stars})
				newReposInPage++
				if len(candidates) >= GitHubMaxCandidateRepos {
					break
				}
			}
		}

		log.Info().
			Int("new_repos_in_page", newReposInPage).
			Int("unique_repos_so_far", len(candidates)).
			Int("star_lookup_failures", starLookupFailures).
			Msg("GitHub page complete")
	}

	ranked := rankAndTrimReposByStars(candidates, limit, order)
	repos := toGitHubRankedRepos(ranked)
	log.Info().
		Int("pages_fetched", pagesFetched).
		Int("candidate_repos", len(candidates)).
		Int("returned_repos", len(repos)).
		Str("order", order).
		Int("star_lookup_failures", starLookupFailures).
		Msg("GitHub search finished")

	return repos, nil
}

func toGitHubRankedRepos(ranked []scoredRepo) []GitHubRankedRepo {
	out := make([]GitHubRankedRepo, 0, len(ranked))
	for _, repo := range ranked {
		out = append(out, GitHubRankedRepo{RepoURL: repo.url, Stars: repo.stars})
	}
	return out
}

// SearchTopGitHubRepos preserves the previous behavior used by full catalog mode integrations.
func SearchTopGitHubRepos(pageSize int) ([]string, error) {
	rankedRepos, err := SearchGitHubRepos(pageSize, GitHubTopReposToReturn, OrderDesc)
	if err != nil {
		partial := make([]string, 0, len(rankedRepos))
		for _, repo := range rankedRepos {
			partial = append(partial, repo.RepoURL)
		}
		return partial, err
	}

	repos := make([]string, 0, len(rankedRepos))
	for _, repo := range rankedRepos {
		repos = append(repos, repo.RepoURL)
	}
	return repos, nil
}
