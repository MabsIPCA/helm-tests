package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/MabsIPCA/helm-tests/helm_fetcher/exporter"
	"github.com/MabsIPCA/helm-tests/helm_fetcher/git"
	"github.com/MabsIPCA/helm-tests/helm_fetcher/helm"
	"github.com/MabsIPCA/helm-tests/helm_fetcher/model"
)

const (
	perPageFlag     = "per-page"
	perPageUsage    = "Results per GitHub search page (max 100)"
	perPageDefault  = 30
	maxPagesFlag    = "max-pages"
	maxPagesUsage   = "Number of GitHub search pages to fetch (max 10)"
	maxPagesDefault = 3
)

const (
	jsonFile = "catalog_by_project.json"
	mdFile   = "catalog_results.md"
	csvFile  = "catalog_failures.csv"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.Kitchen})

	perPage := flag.Int(perPageFlag, perPageDefault, perPageUsage)
	maxPages := flag.Int(maxPagesFlag, maxPagesDefault, maxPagesUsage)
	flag.Parse()

	if err := godotenv.Load(); err != nil {
		log.Warn().Msg(".env not found – using environment variables")
	}
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatal().Msg("GITHUB_TOKEN is required (set in .env or environment)")
	}

	// search for repos with Chart.yaml files
	repos, err := git.SearchGitHubCharts(token, *perPage, *maxPages)
	if err != nil {
		log.Error().Err(err).Msg("Search encountered an error (continuing with partial results)")
	}
	if len(repos) == 0 {
		log.Fatal().Msg("No repos found – nothing to do")
	}

	// Prepare output files and clone base dir
	cloneBase := filepath.Join(".", "cloned")
	_ = os.MkdirAll(cloneBase, 0o755)
	_ = os.Remove(jsonFile)
	_ = os.Remove(mdFile)
	_ = os.Remove(csvFile)

	var allRepos []model.RepoResult

	// iterate over repos
	for i, repoURL := range repos {
		repoName := strings.TrimPrefix(repoURL, "https://github.com/")
		safeName := strings.ReplaceAll(repoName, "/", "__")
		destDir := filepath.Join(cloneBase, safeName)

		log.Info().
			Int("repo_index", i+1).
			Int("total_repos", len(repos)).
			Str("repo", repoName).
			Msg("Processing repo")

		repoResult := model.RepoResult{
			RepoURL:   repoURL,
			RepoName:  repoName,
			ClonedDir: destDir,
		}

		// clone the repo (if not already cloned)
		if err := git.CloneRepo(repoURL, destDir); err != nil {
			log.Error().Err(err).Str("repo", repoName).Msg("Clone failed – skipping")
			continue
		}

		// search for charts (directories with Chart.yaml)
		charts := helm.FindCharts(destDir)
		repoResult.TotalCharts = len(charts)
		log.Info().Int("charts", len(charts)).Msg("Charts found in repo")

		// iterate over charts
		for _, chartDir := range charts {
			err := helm.RunHelmDependencyBuild(chartDir)
			if err != nil {
				log.Warn().Str("chart", chartDir).Msg("Skipping helm template runs due to dependency build failure")
				continue
			}

			chartSummary := model.ChartSummary{ChartPath: chartDir}

			valuesFiles := helm.FindValuesFiles(chartDir)
			log.Info().Str("chart", chartDir).Int("values_files", len(valuesFiles)).Msg("Scanning chart")

			// default run (no extra values files)
			cmdStr, output, runErr := helm.RunHelmTemplate(chartDir, nil)
			run := model.RunResult{
				ChartPath:   chartDir,
				HelmCommand: cmdStr,
				Success:     runErr == nil,
			}
			if runErr != nil {
				run.ErrorMessage = output
				chartSummary.Failures++
				log.Warn().Str("cmd", cmdStr).Msg("FAIL (default)")
			} else {
				chartSummary.Successes++
			}
			chartSummary.Runs = append(chartSummary.Runs, run)
			chartSummary.TotalRuns++

			// runs with values file combinations
			if len(valuesFiles) > 0 {
				if len(valuesFiles) > 10 {
					log.Warn().Int("n", len(valuesFiles)).Msg("Capping values files to 10")
					valuesFiles = valuesFiles[:10]
				}
				combos := helm.Combinations(valuesFiles)
				log.Info().Int("combinations", len(combos)).Msg("Running value-file combinations")

				for _, combo := range combos {
					cmdStr, output, runErr := helm.RunHelmTemplate(chartDir, combo)
					r := model.RunResult{
						ChartPath:   chartDir,
						ValuesFiles: combo,
						HelmCommand: cmdStr,
						Success:     runErr == nil,
					}
					if runErr != nil {
						r.ErrorMessage = output
						chartSummary.Failures++
						log.Warn().Str("cmd", cmdStr).Msg("FAIL")
					} else {
						chartSummary.Successes++
					}
					chartSummary.Runs = append(chartSummary.Runs, r)
					chartSummary.TotalRuns++
				}
			}

			repoResult.Charts = append(repoResult.Charts, chartSummary)
			repoResult.TotalRuns += chartSummary.TotalRuns
			repoResult.TotalSuccess += chartSummary.Successes
			repoResult.TotalFailures += chartSummary.Failures
		}

		// Decide whether to keep or remove the cloned repo based on failures
		if repoResult.TotalFailures > 0 {
			repoResult.Kept = true
			log.Info().Str("repo", repoName).Msg("Keeping cloned repo (has failures)")
		} else {
			repoResult.Kept = false
			exporter.RemoveDir(destDir)
		}

		allRepos = append(allRepos, repoResult)

		// write continuous output after each repo to avoid losing data on crashes
		exporter.FlushJSON(jsonFile, allRepos)
		_ = exporter.WriteMarkdownTable(mdFile, allRepos)

		log.Info().
			Int("total_runs", repoResult.TotalRuns).
			Int("failures", repoResult.TotalFailures).
			Str("repo", repoName).
			Msg("Repo processing complete – output flushed")
	}

	// export final results
	exporter.FlushJSON(jsonFile, allRepos)
	_ = exporter.WriteMarkdownTable(mdFile, allRepos)
	_ = exporter.WriteFailuresCSV(csvFile, allRepos)

	totalRuns := 0
	totalFailures := 0
	for _, r := range allRepos {
		totalRuns += r.TotalRuns
		totalFailures += r.TotalFailures
	}

	log.Info().
		Int("repos", len(allRepos)).
		Int("total_runs", totalRuns).
		Int("failures", totalFailures).
		Int("successes", totalRuns-totalFailures).
		Msg("Done")

	fmt.Println()
	fmt.Printf("✅  %d repos processed, %d total runs, %d failures.\n", len(allRepos), totalRuns, totalFailures)
	fmt.Println("   → catalog_by_project.json  (all results by project)")
	fmt.Println("   → catalog_results.md       (markdown summary table)")
	fmt.Println("   → catalog_failures.csv     (failures only)")
}
