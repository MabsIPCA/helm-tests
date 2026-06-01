package exporter

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/MabsIPCA/helm-tests/helm_fetcher/model"
	"github.com/rs/zerolog/log"
)

// ---- helpers ---------------------------------------------------------------

// FlushJSON writes the repository results to a JSON file with indentation for readability.
func FlushJSON(filename string, repos []model.RepoResult) {
	data, _ := json.MarshalIndent(repos, "", "  ")
	_ = os.WriteFile(filename, data, 0o644)
}

// FlushDepFailuresJSON writes the dep-failure entries to a JSON file.
func FlushDepFailuresJSON(filename string, entries []model.DepFailureEntry) {
	data, _ := json.MarshalIndent(entries, "", "  ")
	_ = os.WriteFile(filename, data, 0o644)
}

// filter helpers
func keptRepos(repos []model.RepoResult) []model.RepoResult {
	var out []model.RepoResult
	for _, r := range repos {
		if r.Kept {
			out = append(out, r)
		}
	}
	return out
}

func removedRepos(repos []model.RepoResult) []model.RepoResult {
	var out []model.RepoResult
	for _, r := range repos {
		if !r.Kept && !r.DepFailed {
			out = append(out, r)
		}
	}
	return out
}

func depFailedRepos(repos []model.RepoResult) []model.RepoResult {
	var out []model.RepoResult
	for _, r := range repos {
		if r.DepFailed {
			out = append(out, r)
		}
	}
	return out
}

// ---- per-type flush helpers ------------------------------------------------

// FlushAll writes all catalog files into dir (the run-specific subdirectory).
func FlushAll(allRepos []model.RepoResult, depFailures []model.DepFailureEntry, dir string) {
	p := func(name string) string { return filepath.Join(dir, name) }

	FlushJSON(p("catalog_by_project.json"), allRepos)
	FlushJSON(p("catalog_kept.json"), keptRepos(allRepos))
	FlushJSON(p("catalog_removed.json"), removedRepos(allRepos))
	FlushDepFailuresJSON(p("catalog_dep_failures.json"), depFailures)

	_ = WriteMarkdownTable(p("catalog_results.md"), allRepos)
	_ = WriteMarkdownTable(p("catalog_kept.md"), keptRepos(allRepos))
	_ = WriteMarkdownTable(p("catalog_removed.md"), removedRepos(allRepos))
	_ = WriteDepFailuresMarkdown(p("catalog_dep_failures.md"), depFailures)

	_ = WriteFailuresCSV(p("catalog_failures.csv"), allRepos)
	_ = WriteFailuresCSV(p("catalog_kept_failures.csv"), keptRepos(allRepos))
	_ = WriteDepFailuresCSV(p("catalog_dep_failures.csv"), depFailures)
}

// MergeCumulative scans all runs/<*>/catalog_by_project.json files under runsDir,
// deduplicates entries by RepoURL (first-seen wins), and writes catalog_cumulative.json
// and catalog_cumulative_dep_failures.json in the current directory.
func MergeCumulative(runsDir string) {
	// --- repos ---
	repoMatches, _ := filepath.Glob(filepath.Join(runsDir, "*", "catalog_by_project.json"))
	seenURL := make(map[string]struct{})
	var repos []model.RepoResult
	for _, path := range repoMatches {
		data, err := os.ReadFile(path)
		if err != nil {
			log.Warn().Err(err).Str("file", path).Msg("MergeCumulative: skipping unreadable file")
			continue
		}
		var batch []model.RepoResult
		if err := json.Unmarshal(data, &batch); err != nil {
			log.Warn().Err(err).Str("file", path).Msg("MergeCumulative: skipping unparseable file")
			continue
		}
		for _, r := range batch {
			if _, ok := seenURL[r.RepoURL]; !ok {
				seenURL[r.RepoURL] = struct{}{}
				repos = append(repos, r)
			}
		}
	}
	FlushJSON("catalog_cumulative.json", repos)

	// --- dep failures ---
	depMatches, _ := filepath.Glob(filepath.Join(runsDir, "*", "catalog_dep_failures.json"))
	type depKey struct{ repoURL, chartPath string }
	seenDep := make(map[depKey]struct{})
	var deps []model.DepFailureEntry
	for _, path := range depMatches {
		data, err := os.ReadFile(path)
		if err != nil {
			continue
		}
		var batch []model.DepFailureEntry
		if err := json.Unmarshal(data, &batch); err != nil {
			continue
		}
		for _, e := range batch {
			k := depKey{e.RepoURL, e.ChartPath}
			if _, ok := seenDep[k]; !ok {
				seenDep[k] = struct{}{}
				deps = append(deps, e)
			}
		}
	}
	FlushDepFailuresJSON("catalog_cumulative_dep_failures.json", deps)

	log.Info().
		Str("runs_dir", runsDir).
		Int("repos", len(repos)).
		Int("dep_failures", len(deps)).
		Msg("Cumulative catalog rebuilt")
}

// ---- markdown writers ------------------------------------------------------

// WriteMarkdownTable generates a Markdown report summarizing the results for each repository.
func WriteMarkdownTable(filename string, repos []model.RepoResult) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	// Header
	fmt.Fprintln(f, "# Helm Chart Catalog Results")
	fmt.Fprintln(f)
	fmt.Fprintln(f, "## Summary by Repository")
	fmt.Fprintln(f)
	fmt.Fprintln(f, "| # | Repository | Charts | Runs | ✅ Success | ❌ Failures | 🔧 Dep Failures | Success Rate | Status |")
	fmt.Fprintln(f, "|---|------------|--------|------|------------|-------------|-----------------|--------------|--------|")

	totalCharts := 0
	totalRuns := 0
	totalSuccess := 0
	totalFailures := 0
	totalDepFail := 0

	for i, repo := range repos {
		rate := float64(0)
		if repo.TotalRuns > 0 {
			rate = float64(repo.TotalSuccess) / float64(repo.TotalRuns) * 100
		}
		status := "Removed"
		if repo.DepFailed {
			status = "Dep-Failed"
		} else if repo.Kept {
			status = "Kept"
		}

		fmt.Fprintf(f, "| %d | [%s](%s) | %d | %d | %d | %d | %d | %.1f%% | %s |\n",
			i+1,
			repo.RepoName,
			repo.RepoURL,
			repo.TotalCharts,
			repo.TotalRuns,
			repo.TotalSuccess,
			repo.TotalFailures,
			repo.TotalDepFailures,
			rate,
			status,
		)

		totalCharts += repo.TotalCharts
		totalRuns += repo.TotalRuns
		totalSuccess += repo.TotalSuccess
		totalFailures += repo.TotalFailures
		totalDepFail += repo.TotalDepFailures
	}

	// Totals row
	totalRate := float64(0)
	if totalRuns > 0 {
		totalRate = float64(totalSuccess) / float64(totalRuns) * 100
	}
	fmt.Fprintln(f, "|---|------------|--------|------|------------|-------------|-----------------|--------------|--------|")
	fmt.Fprintf(f, "| **Total** | **%d repos** | **%d** | **%d** | **%d** | **%d** | **%d** | **%.1f%%** | - |\n",
		len(repos), totalCharts, totalRuns, totalSuccess, totalFailures, totalDepFail, totalRate)

	// Detailed failures section
	fmt.Fprintln(f)
	fmt.Fprintln(f, "## Failure Details")
	fmt.Fprintln(f)

	for _, repo := range repos {
		if repo.TotalFailures == 0 {
			continue
		}

		fmt.Fprintf(f, "### %s\n\n", repo.RepoName)

		for _, chart := range repo.Charts {
			if chart.Failures == 0 {
				continue
			}

			fmt.Fprintf(f, "#### `%s`\n\n", chart.ChartPath)
			fmt.Fprintln(f, "| Values Files | Command | Error |")
			fmt.Fprintln(f, "|--------------|---------|-------|")

			for _, run := range chart.Runs {
				if run.Success {
					continue
				}
				values := "(default)"
				if len(run.ValuesFiles) > 0 {
					values = strings.Join(run.ValuesFiles, ", ")
				}
				errMsg := strings.ReplaceAll(run.ErrorMessage, "\n", " ")
				errMsg = strings.ReplaceAll(errMsg, "|", "\\|")
				fmt.Fprintf(f, "| %s | `%s` | %s |\n", values, run.HelmCommand, errMsg)
			}
			fmt.Fprintln(f)
		}
	}

	return nil
}

// WriteDepFailuresMarkdown generates a Markdown report for dependency build failures.
func WriteDepFailuresMarkdown(filename string, entries []model.DepFailureEntry) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Fprintln(f, "# Helm Dependency Build Failures")
	fmt.Fprintln(f)
	fmt.Fprintf(f, "Total dependency failures: **%d**\n\n", len(entries))
	fmt.Fprintln(f, "| # | Repository | Chart Path | Error |")
	fmt.Fprintln(f, "|---|------------|------------|-------|")

	for i, e := range entries {
		errMsg := strings.ReplaceAll(e.Error, "\n", " ")
		errMsg = strings.ReplaceAll(errMsg, "|", "\\|")
		fmt.Fprintf(f, "| %d | [%s](%s) | `%s` | %s |\n",
			i+1, e.RepoName, e.RepoURL, e.ChartPath, errMsg)
	}

	// Detailed errors
	if len(entries) > 0 {
		fmt.Fprintln(f)
		fmt.Fprintln(f, "## Full Error Details")
		fmt.Fprintln(f)
		for _, e := range entries {
			fmt.Fprintf(f, "### `%s` — %s\n\n", e.ChartPath, e.RepoName)
			fmt.Fprintf(f, "```\n%s\n```\n\n", strings.TrimSpace(e.Error))
		}
	}

	return nil
}

// ---- CSV writers -----------------------------------------------------------

// WriteFailuresCSV generates a CSV file listing all failed template runs.
func WriteFailuresCSV(filename string, repos []model.RepoResult) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	header := []string{"repo_url", "repo_name", "chart_path", "values_files", "helm_command", "error_message"}
	_ = w.Write(header)

	for _, repo := range repos {
		for _, chart := range repo.Charts {
			for _, run := range chart.Runs {
				if run.Success {
					continue
				}
				_ = w.Write([]string{
					repo.RepoURL,
					repo.RepoName,
					run.ChartPath,
					strings.Join(run.ValuesFiles, " | "),
					run.HelmCommand,
					run.ErrorMessage,
				})
			}
		}
	}
	return nil
}

// WriteDepFailuresCSV generates a CSV file listing all dependency build failures.
func WriteDepFailuresCSV(filename string, entries []model.DepFailureEntry) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	_ = w.Write([]string{"repo_url", "repo_name", "chart_path", "error"})
	for _, e := range entries {
		_ = w.Write([]string{e.RepoURL, e.RepoName, e.ChartPath, e.Error})
	}
	return nil
}

// ---- misc ------------------------------------------------------------------

// RemoveDir attempts to delete the specified directory and logs the outcome.
func RemoveDir(dir string) {
	if err := os.RemoveAll(dir); err != nil {
		log.Warn().Err(err).Str("dir", dir).Msg("Could not remove directory")
	} else {
		log.Info().Str("dir", dir).Msg("Removed (no errors – not needed)")
	}
}
