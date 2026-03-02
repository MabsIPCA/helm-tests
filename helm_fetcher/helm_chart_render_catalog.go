package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// ---------------------------------------------------------------------------
// Types
// ---------------------------------------------------------------------------

// ChartResult holds the result for a chart + values combination.
type ChartResult struct {
	RepoURL      string   `json:"repo_url"`
	RepoName     string   `json:"repo_name"`
	ChartPath    string   `json:"chart_path"`
	ValuesFiles  []string `json:"values_files"`
	HelmCommand  string   `json:"helm_command"`
	Success      bool     `json:"success"`
	ErrorMessage string   `json:"error_message,omitempty"`
}

// gitHubSearchResult models the JSON returned by the GitHub code-search API.
type gitHubSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []struct {
		Name       string `json:"name"`
		Path       string `json:"path"`
		Repository struct {
			FullName string `json:"full_name"`
			HTMLURL  string `json:"html_url"`
		} `json:"repository"`
	} `json:"items"`
}

// ---------------------------------------------------------------------------
// GitHub search  – uses curl to avoid the Go HTTP hang on Windows
// ---------------------------------------------------------------------------

// curlGitHubAPI shells out to curl.exe so we sidestep every Go-on-Windows TLS
// / proxy / HTTP2 quirk that causes the request to hang forever.
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

// searchGitHubCharts searches for repos that contain a Chart.yaml using the
// GitHub v3 Code Search API via curl.
func searchGitHubCharts(token string, perPage, maxPages int) ([]string, error) {
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

		var result gitHubSearchResult
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

		// GitHub code-search: 10 req/min for authenticated users.
		if page < maxPages {
			log.Info().Msg("Sleeping 7 s to respect rate-limit …")
			time.Sleep(7 * time.Second)
		}
	}

	log.Info().Int("unique_repos", len(repos)).Msg("GitHub search finished")
	return repos, nil
}

// ---------------------------------------------------------------------------
// Git helpers
// ---------------------------------------------------------------------------

func cloneRepo(repoURL, destDir string) error {
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

// ---------------------------------------------------------------------------
// Chart / values helpers
// ---------------------------------------------------------------------------

func findCharts(baseDir string) []string {
	var charts []string
	_ = filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.Name() == "Chart.yaml" {
			charts = append(charts, filepath.Dir(path))
		}
		return nil
	})
	return charts
}

func findValuesFiles(chartDir string) []string {
	var all []string
	for _, pat := range []string{"values*.yaml", "values*.yml"} {
		matches, _ := filepath.Glob(filepath.Join(chartDir, pat))
		all = append(all, matches...)
	}
	return all
}

// combinations returns every non-empty subset of items.
func combinations(items []string) [][]string {
	n := len(items)
	if n == 0 {
		return nil
	}
	total := 1 << n
	var out [][]string
	for mask := 1; mask < total; mask++ {
		var combo []string
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				combo = append(combo, items[i])
			}
		}
		out = append(out, combo)
	}
	return out
}

// ---------------------------------------------------------------------------
// helm dependency build
// ---------------------------------------------------------------------------

func runHelmDependencyBuild(chartPath string) error {
	cmd := exec.Command("helm", "dependency", "build", chartPath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Warn().Str("chart", chartPath).Str("output", string(out)).Msg("helm dependency build failed (continuing)")
		return err
	}
	log.Info().Str("chart", chartPath).Msg("helm dependency build succeeded")
	return nil
}

// ---------------------------------------------------------------------------
// helm template runner
// ---------------------------------------------------------------------------

func runHelmTemplate(chartPath string, valuesFiles []string) (cmdStr, output string, err error) {
	args := []string{"template", "test", chartPath}
	for _, vf := range valuesFiles {
		args = append(args, "-f", vf)
	}
	cmdStr = "helm " + strings.Join(args, " ")

	cmd := exec.Command("helm", args...)
	out, runErr := cmd.CombinedOutput()
	return cmdStr, string(out), runErr
}

// ---------------------------------------------------------------------------
// CSV writer
// ---------------------------------------------------------------------------

func writeCSV(filename string, results []ChartResult) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	header := []string{"repo_url", "repo_name", "chart_path", "values_files", "helm_command", "success", "error_message"}
	_ = w.Write(header)

	for _, r := range results {
		_ = w.Write([]string{
			r.RepoURL,
			r.RepoName,
			r.ChartPath,
			strings.Join(r.ValuesFiles, " | "),
			r.HelmCommand,
			fmt.Sprintf("%t", r.Success),
			r.ErrorMessage,
		})
	}
	return nil
}

// appendCSVRow appends a single result row to the CSV file (creates with header if missing).
func appendCSVRow(filename string, r ChartResult) {
	needsHeader := false
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		needsHeader = true
	}
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		log.Error().Err(err).Msg("Cannot open CSV for append")
		return
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	if needsHeader {
		_ = w.Write([]string{"repo_url", "repo_name", "chart_path", "values_files", "helm_command", "success", "error_message"})
	}
	_ = w.Write([]string{
		r.RepoURL,
		r.RepoName,
		r.ChartPath,
		strings.Join(r.ValuesFiles, " | "),
		r.HelmCommand,
		fmt.Sprintf("%t", r.Success),
		r.ErrorMessage,
	})
}

// flushAllResultsJSON overwrites the JSON file with the current full slice.
func flushAllResultsJSON(filename string, results []ChartResult) {
	data, _ := json.MarshalIndent(results, "", "  ")
	_ = os.WriteFile(filename, data, 0o644)
}

// removeDir removes a directory tree and logs it.
func removeDir(dir string) {
	if err := os.RemoveAll(dir); err != nil {
		log.Warn().Err(err).Str("dir", dir).Msg("Could not remove directory")
	} else {
		log.Info().Str("dir", dir).Msg("Removed (no errors – not needed)")
	}
}

// filterFailures returns only the results that have Success == false.
func filterFailures(results []ChartResult) []ChartResult {
	var out []ChartResult
	for _, r := range results {
		if !r.Success {
			out = append(out, r)
		}
	}
	return out
}

// ---------------------------------------------------------------------------
// main
// ---------------------------------------------------------------------------

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.Kitchen})

	perPage := flag.Int("per-page", 30, "Results per GitHub search page (max 100)")
	maxPages := flag.Int("max-pages", 3, "Number of pages to fetch (≈ per-page × max-pages unique repos)")
	flag.Parse()

	if err := godotenv.Load(); err != nil {
		log.Warn().Msg(".env not found – using environment variables")
	}

	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatal().Msg("GITHUB_TOKEN is required (set in .env or environment)")
	}

	// ---- 1. Search GitHub ----
	repos, err := searchGitHubCharts(token, *perPage, *maxPages)
	if err != nil {
		log.Error().Err(err).Msg("Search encountered an error (continuing with partial results)")
	}
	if len(repos) == 0 {
		log.Fatal().Msg("No repos found – nothing to do")
	}

	cloneBase := filepath.Join(".", "cloned")
	_ = os.MkdirAll(cloneBase, 0o755)

	const allJSONFile = "catalog_all_results.json"
	const failCSVFile = "catalog_failures.csv"

	// Remove stale output from previous runs.
	_ = os.Remove(allJSONFile)
	_ = os.Remove(failCSVFile)

	var allResults []ChartResult
	totalFailures := 0

	for i, repoURL := range repos {
		repoName := strings.TrimPrefix(repoURL, "https://github.com/")
		safeName := strings.ReplaceAll(repoName, "/", "__")
		destDir := filepath.Join(cloneBase, safeName)

		log.Info().
			Int("repo_index", i+1).
			Int("total_repos", len(repos)).
			Str("repo", repoName).
			Msg("Processing repo")

		// ---- 2. Clone & checkout ----
		if err := cloneRepo(repoURL, destDir); err != nil {
			log.Error().Err(err).Str("repo", repoName).Msg("Clone failed – skipping")
			continue
		}

		// ---- 3. Find charts ----
		charts := findCharts(destDir)
		log.Info().Int("charts", len(charts)).Msg("Charts found in repo")

		repoHasFailures := false

		for _, chartDir := range charts {
			// ---- 3b. Build dependencies first ----
			runHelmDependencyBuild(chartDir)

			valuesFiles := findValuesFiles(chartDir)
			log.Info().
				Str("chart", chartDir).
				Int("values_files", len(valuesFiles)).
				Msg("Scanning chart")

			// ---- 4a. Default run (no -f flags) ----
			cmdStr, output, runErr := runHelmTemplate(chartDir, nil)
			res := ChartResult{
				RepoURL:     repoURL,
				RepoName:    repoName,
				ChartPath:   chartDir,
				HelmCommand: cmdStr,
				Success:     runErr == nil,
			}
			if runErr != nil {
				res.ErrorMessage = output
				appendCSVRow(failCSVFile, res)
				repoHasFailures = true
				totalFailures++
				log.Warn().Str("cmd", cmdStr).Msg("FAIL (default)")
			}
			allResults = append(allResults, res)

			// ---- 4b. Every non-empty subset of values files ----
			if len(valuesFiles) > 0 {
				if len(valuesFiles) > 10 {
					log.Warn().Int("n", len(valuesFiles)).Msg("Capping values files to 10")
					valuesFiles = valuesFiles[:10]
				}
				combos := combinations(valuesFiles)
				log.Info().Int("combinations", len(combos)).Msg("Running value-file combinations")

				for _, combo := range combos {
					cmdStr, output, runErr := runHelmTemplate(chartDir, combo)
					r := ChartResult{
						RepoURL:     repoURL,
						RepoName:    repoName,
						ChartPath:   chartDir,
						ValuesFiles: combo,
						HelmCommand: cmdStr,
						Success:     runErr == nil,
					}
					if runErr != nil {
						r.ErrorMessage = output
						appendCSVRow(failCSVFile, r)
						repoHasFailures = true
						totalFailures++
						log.Warn().Str("cmd", cmdStr).Msg("FAIL")
					}
					allResults = append(allResults, r)
				}
			}
		}

		// ---- Incremental: flush JSON after each repo ----
		flushAllResultsJSON(allJSONFile, allResults)

		// ---- Cleanup: remove repos that produced zero errors ----
		if !repoHasFailures {
			removeDir(destDir)
		} else {
			log.Info().Str("repo", repoName).Msg("Keeping cloned repo (has failures)")
		}

		log.Info().
			Int("total_runs_so_far", len(allResults)).
			Int("failures_so_far", totalFailures).
			Str("repo", repoName).
			Msg("Repo processing complete – output flushed")
	}

	// ---- 5. Final summary write ----
	flushAllResultsJSON(allJSONFile, allResults)
	if err := writeCSV(failCSVFile, filterFailures(allResults)); err != nil {
		log.Error().Err(err).Msg("CSV write failed")
	}

	log.Info().
		Int("total_runs", len(allResults)).
		Int("failures", totalFailures).
		Int("successes", len(allResults)-totalFailures).
		Msg("Done")

	fmt.Println()
	fmt.Printf("✅  %d total runs, %d failures cataloged.\n", len(allResults), totalFailures)
	fmt.Println("   → catalog_all_results.json  (all results)")
	fmt.Println("   → catalog_failures.csv       (failures only)")
}
