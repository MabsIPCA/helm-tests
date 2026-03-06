package git

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/MabsIPCA/helm-tests/helm_fetcher/model"
	"github.com/rs/zerolog/log"
)

// curlGitHubAPI performs a GET request to the specified GitHub API URL using curl with the provided token for authentication.
func curlGitHubAPI(url, token string) ([]byte, error) {
	args := []string{
		"-s", "-S",
		"--max-time", "30",
		"-H", "User-Agent: helm-fetcher/1.0",
		"-H", "Accept: application/vnd.github.v3+json",
		"-H", "Authorization: Bearer " + token,
		url,
	}
	cmd := exec.Command("curl", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("curl failed: %w – output: %s", err, string(out))
	}
	return out, nil
}

// SearchGitHubCharts searches GitHub for code files named "Chart.yaml" using the GitHub Search API,
// returning a list of unique repository URLs that contain such files. It handles pagination and respects
// rate limits by sleeping between requests.
func SearchGitHubCharts(token string, perPage, maxPages int) ([]string, error) {
	seen := map[string]bool{}
	var repos []string

	for page := 1; page <= maxPages; page++ {
		url := fmt.Sprintf(
			"https://api.github.com/search/code?q=filename:Chart.yaml&per_page=%d&page=%d",
			perPage, page,
		)
		log.Info().Str("url", url).Int("page", page).Msg("Searching GitHub via curl")

		body, err := curlGitHubAPI(url, token)
		if err != nil {
			return repos, fmt.Errorf("page %d: %w", page, err)
		}

		var result model.GitHubSearchResult
		if err := json.Unmarshal(body, &result); err != nil {
			return repos, fmt.Errorf("JSON decode page %d: %w – body snippet: %.300s", page, err, string(body))
		}

		for _, item := range result.Items {
			repoURL := item.Repository.HTMLURL
			if !seen[repoURL] {
				seen[repoURL] = true
				repos = append(repos, repoURL)
			}
		}

		log.Info().
			Int("total_count", result.TotalCount).
			Int("new_repos_this_page", len(result.Items)).
			Int("unique_repos_so_far", len(repos)).
			Msg("Page complete")

		if len(result.Items) == 0 {
			break
		}

		if page < maxPages {
			log.Info().Msg("Sleeping 7 s to respect rate-limit …")
			time.Sleep(7 * time.Second)
		}
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
