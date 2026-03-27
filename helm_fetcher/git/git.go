package git

import (
	"fmt"
	neturl "net/url"
	"os"
	"os/exec"
	"regexp"
	"strings"

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
