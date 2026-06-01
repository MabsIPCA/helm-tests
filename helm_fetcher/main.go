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

type artifactHubSearchRepo struct {
	RepoURL string `json:"repo_url"`
}

type artifactHubSearchOutput struct {
	Source string                  `json:"source"`
	Top    int                     `json:"top"`
	Repos  []artifactHubSearchRepo `json:"repos"`
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
	if selectedMode == "artifacthub-search-json" {
		searchOnlyErr := runArtifactHubSearchJSONMode(cfg.PageSize, cfg.SearchTop, cfg.SearchOut)
		if searchOnlyErr != nil {
			log.Fatal().Err(searchOnlyErr).Msg("artifacthub-search-json mode failed")
		}
		return
	}
	if selectedMode == "fixer" {
		runFixerMode(cfg.CatalogIn, cfg.CloneDir)
		return
	}
	if selectedMode == "merge" {
		exporter.MergeCumulative("runs")
		return
	}
	if selectedMode != "full" {
		log.Fatal().Str("mode", cfg.Mode).Msg("Invalid mode. Use 'full', 'github-search-json', 'artifacthub-search-json', 'fixer', or 'merge'")
	}

	selectedSource := strings.ToLower(strings.TrimSpace(cfg.Source))
	var (
		repos []string
		err   error
	)

	switch selectedSource {
	case "artifacthub":
		log.Info().
			Str("search_json", cfg.SearchIn).
			Int("top", cfg.Top).
			Msg("Artifact Hub source selected: loading repos from JSON")
		repos, err = loadArtifactHubReposFromJSON(cfg.SearchIn, cfg.Top)
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

	// Prepare clone base dir
	cloneBase := strings.TrimSpace(cfg.CloneDir)
	if cloneBase == "" {
		cloneBase = "cloned"
	}
	cloneBase = filepath.Clean(cloneBase)
	_ = os.MkdirAll(cloneBase, 0o755)

	// Determine run output directory; auto-generate when not supplied.
	runDir := strings.TrimSpace(cfg.RunDir)
	if runDir == "" {
		ts := time.Now().Format("20060102_150405")
		runDir = filepath.Join("runs", ts+"_"+selectedSource)
	}
	_ = os.MkdirAll(runDir, 0o755)
	log.Info().Str("run_dir", runDir).Msg("Run outputs directory")

	// Snapshot the search input into the run folder so each run is self-contained.
	if raw, readErr := os.ReadFile(cfg.SearchIn); readErr == nil {
		dest := filepath.Join(runDir, filepath.Base(cfg.SearchIn))
		if writeErr := os.WriteFile(dest, raw, 0o644); writeErr != nil {
			log.Warn().Err(writeErr).Str("dest", dest).Msg("Could not snapshot search input into run dir")
		} else {
			log.Info().Str("src", cfg.SearchIn).Str("dest", dest).Msg("Search input snapshotted")
		}
	}

	// Cross-reference against the cumulative catalog so repos processed in any
	// previous run are skipped immediately without re-cloning or re-templating.
	existingCatalog := loadExistingCatalog("catalog_cumulative.json")
	existingDepFails := loadExistingDepFailures("catalog_cumulative_dep_failures.json")

	var allRepos []model.RepoResult
	var allDepFailures []model.DepFailureEntry
	newlyProcessed := 0

	// iterate over repos
	for i, repoURL := range repos {
		repoName := strings.TrimPrefix(repoURL, "https://github.com/")
		safeName := strings.ReplaceAll(repoName, "/", "__")
		destDir := filepath.Join(cloneBase, safeName)

		// Reuse existing result when the repo was already cataloged.
		if existing, ok := existingCatalog[repoURL]; ok {
			log.Info().
				Int("repo_index", i+1).
				Int("total_repos", len(repos)).
				Str("repo", repoName).
				Msg("Already cataloged – reusing existing result")
			allRepos = append(allRepos, existing)
			if deps, ok := existingDepFails[repoURL]; ok {
				allDepFailures = append(allDepFailures, deps...)
			}
			continue
		}

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
			if !cfg.KeepClones {
				exporter.RemoveDir(destDir)
			} else {
				log.Info().Str("repo", repoName).Msg("No failures – keeping clone for reuse (-keep-clones)")
			}
		}

		allRepos = append(allRepos, repoResult)
		newlyProcessed++

		// write continuous output after each repo to avoid losing data on crashes
		exporter.FlushAll(allRepos, allDepFailures, runDir)

		log.Info().
			Int("total_runs", repoResult.TotalRuns).
			Int("failures", repoResult.TotalFailures).
			Str("repo", repoName).
			Msg("Repo processing complete – output flushed")
	}

	// All repos were already in the cumulative catalog – nothing new to write.
	if newlyProcessed == 0 {
		_ = os.RemoveAll(runDir)
		log.Info().
			Int("reused", len(allRepos)).
			Str("run_dir", runDir).
			Msg("All repos already cataloged – run folder suppressed")
		fmt.Println()
		fmt.Printf("ℹ️  All %d repos already in cumulative catalog – no new run folder created.\n", len(allRepos))
		fmt.Println("   Use 'make plot-all' to generate plots from the existing cumulative data.")
		return
	}

	// export final results for this run
	exporter.FlushAll(allRepos, allDepFailures, runDir)

	// rebuild cumulative catalog from all run subfolders
	exporter.MergeCumulative("runs")

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
		Int("newly_processed", newlyProcessed).
		Int("reused_from_cumulative", len(allRepos)-newlyProcessed).
		Int("total_runs", totalRuns).
		Int("failures", totalFailures).
		Int("dep_failures", totalDepFails).
		Int("successes", totalRuns-totalFailures).
		Msg("Done")

	fmt.Println()
	fmt.Printf("✅  %d repos total (%d new, %d from cumulative), %d helm runs, %d failures, %d dep-build failures.\n",
		len(allRepos), newlyProcessed, len(allRepos)-newlyProcessed, totalRuns, totalFailures, totalDepFails)
	fmt.Println()
	fmt.Printf("Run output:  %s/\n", runDir)
	fmt.Printf("  %-36s — top-500 list used for this run\n", filepath.Base(cfg.SearchIn))
	fmt.Println("  catalog_by_project.json              — this run, all results")
	fmt.Println("  catalog_kept.json                    — this run, repos with template failures")
	fmt.Println("  catalog_removed.json                 — this run, repos with no failures")
	fmt.Println("  catalog_dep_failures.json            — this run, dep-build failures")
	fmt.Println()
	fmt.Println("Cumulative (all runs merged):")
	fmt.Println("  catalog_cumulative.json              — all repos ever processed")
	fmt.Println("  catalog_cumulative_dep_failures.json — all dep-build failures ever")
}

// loadExistingCatalog reads catalog_by_project.json and returns a map of repoURL → RepoResult.
// Returns nil (no-op) when the file is absent or unparseable.
func loadExistingCatalog(path string) map[string]model.RepoResult {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	var repos []model.RepoResult
	if err := json.Unmarshal(data, &repos); err != nil {
		log.Warn().Err(err).Str("file", path).Msg("Could not parse existing catalog – starting fresh")
		return nil
	}
	m := make(map[string]model.RepoResult, len(repos))
	for _, r := range repos {
		m[r.RepoURL] = r
	}
	log.Info().Str("file", path).Int("repos", len(m)).Msg("Loaded existing catalog for cross-reference")
	return m
}

// loadExistingDepFailures reads catalog_dep_failures.json and returns a map of repoURL → []DepFailureEntry.
// Returns nil (no-op) when the file is absent or unparseable.
func loadExistingDepFailures(path string) map[string][]model.DepFailureEntry {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	var entries []model.DepFailureEntry
	if err := json.Unmarshal(data, &entries); err != nil {
		log.Warn().Err(err).Str("file", path).Msg("Could not parse existing dep failures – starting fresh")
		return nil
	}
	m := make(map[string][]model.DepFailureEntry)
	for _, e := range entries {
		m[e.RepoURL] = append(m[e.RepoURL], e)
	}
	log.Info().Str("file", path).Int("entries", len(entries)).Msg("Loaded existing dep failures for cross-reference")
	return m
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

func runArtifactHubSearchJSONMode(pageSize, top int, outputPath string) error {
	repos, err := git.SearchTopArtifactHubRepos(top, pageSize)
	if err != nil {
		log.Error().Err(err).Msg("Artifact Hub search encountered an error (writing partial results)")
	}
	if len(repos) == 0 {
		return fmt.Errorf("no repos found")
	}

	entries := make([]artifactHubSearchRepo, 0, len(repos))
	for _, repoURL := range repos {
		repoURL = strings.TrimSpace(repoURL)
		if repoURL == "" {
			continue
		}
		entries = append(entries, artifactHubSearchRepo{RepoURL: repoURL})
	}
	if len(entries) == 0 {
		return fmt.Errorf("artifacthub search output contains only empty repo URLs")
	}

	payload := artifactHubSearchOutput{
		Source: "artifacthub",
		Top:    top,
		Repos:  entries,
	}

	data, marshalErr := json.MarshalIndent(payload, "", "  ")
	if marshalErr != nil {
		return fmt.Errorf("marshal artifacthub search output: %w", marshalErr)
	}
	if writeErr := os.WriteFile(outputPath, data, 0o644); writeErr != nil {
		return fmt.Errorf("write artifacthub search output: %w", writeErr)
	}

	log.Info().
		Str("output", outputPath).
		Int("repos", len(entries)).
		Int("requested", top).
		Msg("Artifact Hub search JSON generated")
	if top >= 500 && len(entries) < 500 {
		log.Warn().Int("requested", top).Int("returned", len(entries)).Msg("Artifact Hub JSON contains fewer than 500 repos; continuing with partial list")
	}

	// Partial results are acceptable as long as we have a non-empty list.
	return nil
}

func loadArtifactHubReposFromJSON(path string, top int) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read artifacthub search JSON (%s): %w", path, err)
	}

	var payload artifactHubSearchOutput
	if err := json.Unmarshal(data, &payload); err != nil {
		return nil, fmt.Errorf("decode artifacthub search JSON (%s): %w", path, err)
	}
	if len(payload.Repos) == 0 {
		return nil, fmt.Errorf("artifacthub search JSON has no repos: %s", path)
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
		return nil, fmt.Errorf("artifacthub search JSON contains only empty repo URLs: %s", path)
	}

	return repos, nil
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
