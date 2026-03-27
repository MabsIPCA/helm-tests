package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	appconfig "github.com/MabsIPCA/helm-tests/helm_fetcher/config"
	"github.com/MabsIPCA/helm-tests/helm_fetcher/exporter"
	"github.com/MabsIPCA/helm-tests/helm_fetcher/git"
	"github.com/MabsIPCA/helm-tests/helm_fetcher/helm"
	"github.com/MabsIPCA/helm-tests/helm_fetcher/model"
)

type githubSearchOutput struct {
	Source string                 `json:"source"`
	Top    int                    `json:"top"`
	Order  string                 `json:"order"`
	Repos  []git.GitHubRankedRepo `json:"repos"`
}

func loadDotEnvIfPresent() {
	if err := godotenv.Load(); err != nil {
		if !os.IsNotExist(err) {
			log.Warn().Err(err).Msg("Failed to load .env")
		}
		return
	}

	log.Info().Msg("Loaded environment variables from .env")
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.Kitchen})
	loadDotEnvIfPresent()

	cfg := appconfig.Parse()

	selectedMode := strings.ToLower(strings.TrimSpace(cfg.Mode))
	if selectedMode == "github-search-json" {
		searchOnlyErr := runGitHubSearchJSONMode(cfg.PageSize, cfg.SearchTop, cfg.SearchOrder, cfg.SearchOut)
		if searchOnlyErr != nil {
			log.Fatal().Err(searchOnlyErr).Msg("github-search-json mode failed")
		}
		return
	}
	if selectedMode != "full" {
		log.Fatal().Str("mode", cfg.Mode).Msg("Invalid mode. Use 'full' or 'github-search-json'")
	}

	selectedSource := strings.ToLower(strings.TrimSpace(cfg.Source))
	var (
		repos []string
		err   error
	)

	switch selectedSource {
	case "artifacthub":
		repos, err = git.SearchTopArtifactHubRepos(cfg.Top, cfg.PageSize)
	case "github":
		log.Info().
			Str("search_json", cfg.SearchIn).
			Int("top", cfg.Top).
			Msg("GitHub source selected: loading ranked repos from JSON")
		repos, err = loadGitHubReposFromJSON(cfg.SearchIn, cfg.Top)
	default:
		log.Fatal().Str("source", cfg.Source).Msg("Invalid source. Use 'artifacthub' or 'github'")
	}

	if err != nil {
		log.Error().Err(err).Msg("Search encountered an error (continuing with partial results)")
	}
	if len(repos) == 0 {
		log.Fatal().Msg("No repos found – nothing to do")
	}

	// Prepare output files and clone base dir
	cloneBase := filepath.Join(".", "cloned")
	_ = os.MkdirAll(cloneBase, 0o755)
	// Remove old output files
	for _, f := range []string{
		"catalog_by_project.json", "catalog_kept.json", "catalog_removed.json", "catalog_dep_failures.json",
		"catalog_results.md", "catalog_kept.md", "catalog_removed.md", "catalog_dep_failures.md",
		"catalog_failures.csv", "catalog_kept_failures.csv", "catalog_dep_failures.csv",
	} {
		_ = os.Remove(f)
	}

	var allRepos []model.RepoResult
	var allDepFailures []model.DepFailureEntry

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
			depErr := helm.RunHelmDependencyBuild(chartDir)
			if depErr != nil {
				log.Warn().Str("chart", chartDir).Msg("Dependency build failure – skipping remaining charts in repo")
				depErrOutput := depErr.Error()
				chartSummary := model.ChartSummary{
					ChartPath:       chartDir,
					DepBuildFailure: true,
					DepBuildError:   depErrOutput,
				}
				repoResult.Charts = append(repoResult.Charts, chartSummary)
				repoResult.TotalDepFailures++
				allDepFailures = append(allDepFailures, model.DepFailureEntry{
					RepoURL:   repoURL,
					RepoName:  repoName,
					ChartPath: chartDir,
					Error:     depErrOutput,
				})
				break
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
				if len(combos) > 100 {
					log.Warn().Int("n", len(combos)).Msg("Capping combinations to 100")
					combos = combos[:100]
				}
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

		// Decide whether to keep or remove the cloned repo based on failures / dep failures
		if repoResult.TotalDepFailures > 0 {
			repoResult.DepFailed = true
			repoResult.Kept = false
			log.Info().Str("repo", repoName).Int("dep_failures", repoResult.TotalDepFailures).Msg("Marking repo as dep-failed")
		} else if repoResult.TotalFailures > 0 {
			repoResult.Kept = true
			log.Info().Str("repo", repoName).Msg("Keeping cloned repo (has template failures)")
		} else {
			repoResult.Kept = false
			exporter.RemoveDir(destDir)
		}

		allRepos = append(allRepos, repoResult)

		// write continuous output after each repo to avoid losing data on crashes
		exporter.FlushAll(allRepos, allDepFailures)

		log.Info().
			Int("total_runs", repoResult.TotalRuns).
			Int("failures", repoResult.TotalFailures).
			Str("repo", repoName).
			Msg("Repo processing complete – output flushed")
	}

	// export final results
	exporter.FlushAll(allRepos, allDepFailures)

	totalRuns := 0
	totalFailures := 0
	totalDepFails := 0
	for _, r := range allRepos {
		totalRuns += r.TotalRuns
		totalFailures += r.TotalFailures
		totalDepFails += r.TotalDepFailures
	}

	log.Info().
		Int("repos", len(allRepos)).
		Int("total_runs", totalRuns).
		Int("failures", totalFailures).
		Int("dep_failures", totalDepFails).
		Int("successes", totalRuns-totalFailures).
		Msg("Done")

	fmt.Println()
	fmt.Printf("✅  %d repos processed, %d total runs, %d template failures, %d dep-build failures.\n",
		len(allRepos), totalRuns, totalFailures, totalDepFails)
	fmt.Println()
	fmt.Println("Output files:")
	fmt.Println("  catalog_by_project.json      — all results")
	fmt.Println("  catalog_kept.json            — repos with template failures (kept)")
	fmt.Println("  catalog_removed.json         — repos with no failures (removed)")
	fmt.Println("  catalog_dep_failures.json    — repos with dep-build failures")
	fmt.Println("  catalog_results.md           — full markdown summary")
	fmt.Println("  catalog_kept.md              — kept repos markdown")
	fmt.Println("  catalog_removed.md           — removed repos markdown")
	fmt.Println("  catalog_dep_failures.md      — dep-failure details markdown")
	fmt.Println("  catalog_failures.csv         — all template failures (CSV)")
	fmt.Println("  catalog_kept_failures.csv    — kept repos template failures (CSV)")
	fmt.Println("  catalog_dep_failures.csv     — dep-build failures (CSV)")
}

func runGitHubSearchJSONMode(pageSize, top int, order, outputPath string) error {
	normalizedOrder := strings.ToLower(strings.TrimSpace(order))
	repos, err := git.SearchGitHubRepos(pageSize, top, normalizedOrder)
	if err != nil {
		log.Error().Err(err).Msg("GitHub search encountered an error (writing partial results)")
	}
	if len(repos) == 0 {
		return fmt.Errorf("no repos found")
	}

	payload := githubSearchOutput{
		Source: "github",
		Top:    top,
		Order:  normalizedOrder,
		Repos:  repos,
	}

	data, marshalErr := json.MarshalIndent(payload, "", "  ")
	if marshalErr != nil {
		return fmt.Errorf("marshal github search output: %w", marshalErr)
	}
	if writeErr := os.WriteFile(outputPath, data, 0o644); writeErr != nil {
		return fmt.Errorf("write github search output: %w", writeErr)
	}

	log.Info().Str("output", outputPath).Int("repos", len(repos)).Str("order", normalizedOrder).Msg("GitHub search JSON generated")
	return err
}

func loadGitHubReposFromJSON(path string, top int) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read search JSON (%s): %w", path, err)
	}

	var payload githubSearchOutput
	if err := json.Unmarshal(data, &payload); err != nil {
		return nil, fmt.Errorf("decode search JSON (%s): %w", path, err)
	}
	if len(payload.Repos) == 0 {
		return nil, fmt.Errorf("search JSON has no repos: %s", path)
	}

	limit := len(payload.Repos)
	if top > 0 && top < limit {
		limit = top
	}

	repos := make([]string, 0, limit)
	for i := 0; i < limit; i++ {
		repoURL := strings.TrimSpace(payload.Repos[i].RepoURL)
		if repoURL == "" {
			continue
		}
		repos = append(repos, repoURL)
	}
	if len(repos) == 0 {
		return nil, fmt.Errorf("search JSON contains only empty repo URLs: %s", path)
	}

	return repos, nil
}
