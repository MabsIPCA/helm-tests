package git

import (
	"bytes"
	"encoding/json"
	"fmt"
	neturl "net/url"
	"regexp"
	"strings"
	"time"

	"github.com/MabsIPCA/helm-tests/helm_fetcher/model"
	"github.com/rs/zerolog/log"
)

// artifactHubFetchAPI can be overridden in tests.
var artifactHubFetchAPI = curlAPI

const artifactHubHelmKind = "helm"

const (
	artifactHubBatchRetryAttempts        = 5
	artifactHubDetailsRetryAttempts      = 4
	artifactHubRetryBackoff              = 2 * time.Second
	artifactHubRetryBackoffMax           = 20 * time.Second
	artifactHubRateLimitCooldown         = 8 * time.Second
	artifactHubDetailsMinInterval        = 500 * time.Millisecond
	artifactHubMaxConsecutiveBatchErrors = 5
)

var artifactHubSleep = time.Sleep

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

var artifactHubGitHubURLPattern = regexp.MustCompile(`(?i)(?:https?://|ssh://git@|git@)?(?:www\.)?github\.com[:/][^"'\s<>()]+`)

func extractRepoURLFromString(value string) string {
	if normalized := normalizeGitHubRepoURL(value); normalized != "" {
		return normalized
	}

	cleaned := strings.ReplaceAll(value, `\/`, "/")
	for _, candidate := range artifactHubGitHubURLPattern.FindAllString(cleaned, -1) {
		if normalized := normalizeGitHubRepoURL(candidate); normalized != "" {
			return normalized
		}
	}

	return ""
}

func extractRepoURLFromAny(raw any) string {
	switch v := raw.(type) {
	case string:
		return extractRepoURLFromString(v)
	case []any:
		for _, item := range v {
			if normalized := extractRepoURLFromAny(item); normalized != "" {
				return normalized
			}
		}
	case map[string]any:
		for _, item := range v {
			if normalized := extractRepoURLFromAny(item); normalized != "" {
				return normalized
			}
		}
	}

	return ""
}

func extractRepoURLFromRawArtifactPackage(rawPkg json.RawMessage) string {
	var payload any
	if err := json.Unmarshal(rawPkg, &payload); err == nil {
		if normalized := extractRepoURLFromAny(payload); normalized != "" {
			return normalized
		}
	}

	if normalized := extractRepoURLFromString(string(rawPkg)); normalized != "" {
		return normalized
	}

	return ""
}

func buildArtifactHubPackageDetailsURL(pkg model.ArtifactHubPackage) string {
	repoName := strings.TrimSpace(pkg.Repository.Name)
	packageName := strings.TrimSpace(pkg.NormalizedName)
	if packageName == "" {
		packageName = strings.TrimSpace(pkg.Name)
	}
	if repoName == "" || packageName == "" {
		return ""
	}

	return fmt.Sprintf(
		"https://artifacthub.io/api/v1/packages/%s/%s/%s",
		artifactHubHelmKind,
		neturl.PathEscape(repoName),
		neturl.PathEscape(packageName),
	)
}

func extractRepoURLFromPackageDetails(pkg model.ArtifactHubPackage) (string, error) {
	requestURL := buildArtifactHubPackageDetailsURL(pkg)
	if requestURL == "" {
		return "", nil
	}

	var body []byte
	var err error
	for attempt := 1; attempt <= artifactHubDetailsRetryAttempts; attempt++ {
		body, err = artifactHubFetchAPI(requestURL)
		if err == nil && len(bytes.TrimSpace(body)) > 0 {
			break
		}
		if err == nil {
			err = fmt.Errorf("empty details response")
		}
		if !isArtifactHubRetryableError(err) || attempt == artifactHubDetailsRetryAttempts {
			return "", err
		}
		artifactHubSleep(backoffForAttempt(attempt))
	}

	repoURL := extractRepoURLFromRawArtifactPackage(body)
	if repoURL != "" {
		return repoURL, nil
	}

	var detailedPkg model.ArtifactHubPackage
	if err := json.Unmarshal(body, &detailedPkg); err == nil {
		repoURL = extractRepoURL(detailedPkg)
	}

	return repoURL, nil
}

func backoffForAttempt(attempt int) time.Duration {
	if attempt <= 0 {
		return artifactHubRetryBackoff
	}
	backoff := artifactHubRetryBackoff
	for i := 1; i < attempt; i++ {
		backoff *= 2
	}
	if backoff > artifactHubRetryBackoffMax {
		return artifactHubRetryBackoffMax
	}
	return backoff
}

func isArtifactHubRetryableError(err error) bool {
	if err == nil {
		return false
	}
	text := strings.ToLower(err.Error())
	if strings.Contains(text, "429") || strings.Contains(text, "too many requests") || strings.Contains(text, "rate limit") {
		return true
	}
	if strings.Contains(text, "timeout") || strings.Contains(text, "tempor") || strings.Contains(text, "connection reset") {
		return true
	}
	if strings.Contains(text, "unexpected end of json input") || strings.Contains(text, "empty response body") || strings.Contains(text, "empty details response") {
		return true
	}
	return false
}

func fetchArtifactHubSearchBatch(requestURL string, offset int) (model.ArtifactHubSearchResponse, []json.RawMessage, error) {
	var lastErr error

	for attempt := 1; attempt <= artifactHubBatchRetryAttempts; attempt++ {
		body, err := artifactHubFetchAPI(requestURL)
		if err != nil {
			lastErr = fmt.Errorf("offset %d attempt %d/%d fetch failed: %w", offset, attempt, artifactHubBatchRetryAttempts, err)
		} else if len(bytes.TrimSpace(body)) == 0 {
			lastErr = fmt.Errorf("offset %d attempt %d/%d empty response body", offset, attempt, artifactHubBatchRetryAttempts)
		} else {
			var result model.ArtifactHubSearchResponse
			if err := json.Unmarshal(body, &result); err != nil {
				lastErr = fmt.Errorf("offset %d attempt %d/%d JSON decode failed: %w – body snippet: %.300s", offset, attempt, artifactHubBatchRetryAttempts, err, string(body))
			} else {
				var rawResult struct {
					Packages []json.RawMessage `json:"packages"`
				}
				if err := json.Unmarshal(body, &rawResult); err != nil {
					lastErr = fmt.Errorf("offset %d attempt %d/%d raw JSON decode failed: %w – body snippet: %.300s", offset, attempt, artifactHubBatchRetryAttempts, err, string(body))
				} else {
					return result, rawResult.Packages, nil
				}
			}
		}

		if attempt < artifactHubBatchRetryAttempts {
			log.Debug().Err(lastErr).Int("offset", offset).Int("attempt", attempt).Int("max_attempts", artifactHubBatchRetryAttempts).Msg("Artifact Hub batch request failed (will retry)")
		} else {
			log.Warn().Err(lastErr).Int("offset", offset).Int("attempt", attempt).Int("max_attempts", artifactHubBatchRetryAttempts).Msg("Artifact Hub batch request failed")
		}
		if attempt < artifactHubBatchRetryAttempts {
			artifactHubSleep(backoffForAttempt(attempt))
		}
	}

	return model.ArtifactHubSearchResponse{}, nil, lastErr
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
	packageRepoCache := map[string]string{}
	var repos []string
	var skippedNoRepo int
	var detailLookups int
	var detailLookupFailures int
	var detailResolved int
	consecutiveBatchErrors := 0
	detailsDisabledUntil := time.Time{}
	lastDetailsLookupAt := time.Time{}

	for offset := 0; len(repos) < limit; offset += pageSize {
		currentLimit := pageSize

		requestURL := fmt.Sprintf(
			"https://artifacthub.io/api/v1/packages/search?kind=0&sort=stars&limit=%d&offset=%d",
			currentLimit, offset,
		)
		log.Info().Str("url", requestURL).Int("offset", offset).Int("batch_limit", currentLimit).Msg("Searching Artifact Hub via curl")

		result, rawPackages, err := fetchArtifactHubSearchBatch(requestURL, offset)
		if err != nil {
			consecutiveBatchErrors++
			log.Warn().Err(err).Int("offset", offset).Int("consecutive_batch_errors", consecutiveBatchErrors).Msg("Skipping Artifact Hub batch after retries")
			if isArtifactHubRetryableError(err) {
				detailsDisabledUntil = time.Now().Add(artifactHubRateLimitCooldown)
			}
			if consecutiveBatchErrors >= artifactHubMaxConsecutiveBatchErrors {
				log.Warn().Int("consecutive_batch_errors", consecutiveBatchErrors).Int("unique_repos_so_far", len(repos)).Msg("Stopping Artifact Hub search due to repeated batch failures")
				break
			}
			continue
		}
		consecutiveBatchErrors = 0
		if len(result.Packages) == 0 {
			break
		}

		batchNewRepos := 0
		for i, pkg := range result.Packages {
			repoURL := extractRepoURL(pkg)
			if repoURL == "" && i < len(rawPackages) {
				repoURL = extractRepoURLFromRawArtifactPackage(rawPackages[i])
			}
			if repoURL == "" {
				cacheKey := strings.TrimSpace(pkg.PackageID)
				if cacheKey != "" {
					repoURL = packageRepoCache[cacheKey]
				}
				if repoURL == "" {
					if !detailsDisabledUntil.IsZero() && time.Now().Before(detailsDisabledUntil) {
						skippedNoRepo++
						continue
					}
					if !lastDetailsLookupAt.IsZero() {
						since := time.Since(lastDetailsLookupAt)
						if since < artifactHubDetailsMinInterval {
							artifactHubSleep(artifactHubDetailsMinInterval - since)
						}
					}
					lastDetailsLookupAt = time.Now()
					detailLookups++
					resolved, err := extractRepoURLFromPackageDetails(pkg)
					if err != nil {
						detailLookupFailures++
						if isArtifactHubRetryableError(err) {
							detailsDisabledUntil = time.Now().Add(artifactHubRateLimitCooldown)
						}
						log.Debug().Err(err).Str("package", pkg.Name).Str("repo", pkg.Repository.Name).Msg("Artifact Hub details lookup failed")
					} else {
						repoURL = resolved
						if repoURL != "" {
							detailResolved++
						}
					}
					if cacheKey != "" {
						packageRepoCache[cacheKey] = repoURL
					}
				}
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
			Int("details_lookups", detailLookups).
			Int("details_resolved", detailResolved).
			Int("details_failures", detailLookupFailures).
			Msg("Artifact Hub batch complete")
	}

	log.Info().
		Int("unique_repos", len(repos)).
		Int("skipped_without_repo", skippedNoRepo).
		Int("details_lookups", detailLookups).
		Int("details_resolved", detailResolved).
		Int("details_failures", detailLookupFailures).
		Msg("Artifact Hub search finished")
	return repos, nil
}
