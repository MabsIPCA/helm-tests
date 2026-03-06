package exporter

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/MabsIPCA/helm-tests/helm_fetcher/model"
	"github.com/rs/zerolog/log"
)

// FlushJSON writes the repository results to a JSON file with indentation for readability.
func FlushJSON(filename string, repos []model.RepoResult) {
	data, _ := json.MarshalIndent(repos, "", "  ")
	_ = os.WriteFile(filename, data, 0o644)
}

// WriteMarkdownTable generates a Markdown report summarizing the results for each repository, including a detailed section for failures.
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
	fmt.Fprintln(f, "| # | Repository | Charts | Runs | ✅ Success | ❌ Failures | Success Rate | Status |")
	fmt.Fprintln(f, "|---|------------|--------|------|------------|-------------|--------------|--------|")

	totalCharts := 0
	totalRuns := 0
	totalSuccess := 0
	totalFailures := 0

	for i, repo := range repos {
		rate := float64(0)
		if repo.TotalRuns > 0 {
			rate = float64(repo.TotalSuccess) / float64(repo.TotalRuns) * 100
		}
		status := "Removed"
		if repo.Kept {
			status = "Kept"
		}

		fmt.Fprintf(f, "| %d | [%s](%s) | %d | %d | %d | %d | %.1f%% | %s |\n",
			i+1,
			repo.RepoName,
			repo.RepoURL,
			repo.TotalCharts,
			repo.TotalRuns,
			repo.TotalSuccess,
			repo.TotalFailures,
			rate,
			status,
		)

		totalCharts += repo.TotalCharts
		totalRuns += repo.TotalRuns
		totalSuccess += repo.TotalSuccess
		totalFailures += repo.TotalFailures
	}

	// Totals row
	totalRate := float64(0)
	if totalRuns > 0 {
		totalRate = float64(totalSuccess) / float64(totalRuns) * 100
	}
	fmt.Fprintln(f, "|---|------------|--------|------|------------|-------------|--------------|--------|")
	fmt.Fprintf(f, "| **Total** | **%d repos** | **%d** | **%d** | **%d** | **%d** | **%.1f%%** | - |\n",
		len(repos), totalCharts, totalRuns, totalSuccess, totalFailures, totalRate)

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

// WriteFailuresCSV generates a CSV file listing all failed runs with details for easier analysis in spreadsheet software.
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

// RemoveDir attempts to delete the specified directory and logs the outcome, including any errors encountered.
func RemoveDir(dir string) {
	if err := os.RemoveAll(dir); err != nil {
		log.Warn().Err(err).Str("dir", dir).Msg("Could not remove directory")
	} else {
		log.Info().Str("dir", dir).Msg("Removed (no errors – not needed)")
	}
}
