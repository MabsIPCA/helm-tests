package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rs/zerolog/log"

	"github.com/MabsIPCA/helm-tests/helm_fetcher/git"
	"github.com/MabsIPCA/helm-tests/helm_fetcher/helm"
	"github.com/MabsIPCA/helm-tests/helm_fetcher/model"
)

// runRefilterValuesMode re-templates ONLY the repos whose results the YAML
// prefilter actually changes — i.e. repos with at least one overlay value file
// that the prefilter now rejects (git-crypt encrypted or otherwise unparseable).
// Every other repo would re-template to an identical result, so it is left
// untouched. This makes a targeted pass equivalent to a full fresh re-run (same
// end-state catalog) at a fraction of the cost: no dep-build backoff churn on
// the ~95% of repos that have no unparseable value files.
//
// The updated catalog is written back to catalogIn in place, after a one-time
// .prefilter.bak backup.
func runRefilterValuesMode(catalogIn, cloneDir string) {
	data, err := os.ReadFile(catalogIn)
	if err != nil {
		log.Fatal().Err(err).Str("path", catalogIn).Msg("Cannot read catalog")
	}
	var repos []model.RepoResult
	if err := json.Unmarshal(data, &repos); err != nil {
		log.Fatal().Err(err).Msg("Cannot parse catalog JSON")
	}

	// One-time backup so the pre-prefilter snapshot is never lost.
	backup := catalogIn + ".prefilter.bak"
	if _, statErr := os.Stat(backup); os.IsNotExist(statErr) {
		if writeErr := os.WriteFile(backup, data, 0o644); writeErr != nil {
			log.Warn().Err(writeErr).Str("backup", backup).Msg("Could not write .prefilter.bak (continuing)")
		} else {
			log.Info().Str("backup", backup).Msg("Backed up original catalog")
		}
	}

	// Identify the affected repos up front so we can report progress meaningfully.
	type target struct {
		idx     int
		destDir string
		dropped int
	}
	var targets []target
	for i := range repos {
		destDir := resolveDestDir(repos[i].RepoURL, cloneDir)
		if n := repoDroppedValueFiles(destDir); n > 0 {
			targets = append(targets, target{idx: i, destDir: destDir, dropped: n})
		}
	}
	log.Info().Int("total_repos", len(repos)).Int("affected_repos", len(targets)).Str("clone_dir", cloneDir).Msg("Re-templating repos with prefiltered value files")

	reprocessed := 0
	for _, t := range targets {
		repo := repos[t.idx]
		reprocessed++
		log.Info().Int("n", reprocessed).Int("of", len(targets)).Int("dropped_files", t.dropped).Str("repo", repo.RepoName).Msg("Re-templating affected repo")

		// Clones are reused (CloneRepo skips an existing dir); clone only if the
		// directory is missing so the repo can still be re-templated.
		if cloneErr := git.CloneRepo(repo.RepoURL, t.destDir); cloneErr != nil {
			log.Warn().Err(cloneErr).Str("repo", repo.RepoName).Msg("Clone failed – leaving existing result unchanged")
			continue
		}
		repos[t.idx] = retemplateRepo(repo.RepoURL, repo.RepoName, t.destDir)
	}

	out, err := json.MarshalIndent(repos, "", "  ")
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot marshal updated catalog")
	}
	if err := os.WriteFile(catalogIn, out, 0o644); err != nil {
		log.Fatal().Err(err).Str("path", catalogIn).Msg("Cannot write updated catalog")
	}

	fmt.Printf("\nRefilter-values complete.\n")
	fmt.Printf("  Repos with prefiltered value files re-templated: %d / %d\n", reprocessed, len(repos))
	fmt.Printf("  Updated catalog: %s (backup: %s)\n", catalogIn, backup)
}

// repoDroppedValueFiles counts the overlay value files across all charts in
// destDir that the YAML prefilter rejects. A repo with zero is unaffected by the
// prefilter and is skipped.
func repoDroppedValueFiles(destDir string) int {
	if info, err := os.Stat(destDir); err != nil || !info.IsDir() {
		return 0
	}
	n := 0
	for _, chartDir := range helm.FindCharts(destDir) {
		n += len(helm.DroppedValuesFiles(chartDir))
	}
	return n
}

// retemplateRepo rebuilds a RepoResult from scratch for destDir, mirroring full
// mode's per-chart flow (helm dependency build, then template the default run
// plus every value-file combination). Because FindValuesFiles now applies the
// prefilter, the dropped overlays simply never enter the combination set.
func retemplateRepo(repoURL, repoName, destDir string) model.RepoResult {
	repoResult := model.RepoResult{
		RepoURL:   repoURL,
		RepoName:  repoName,
		ClonedDir: destDir,
	}

	charts := helm.FindCharts(destDir)
	repoResult.TotalCharts = len(charts)

	for _, chartDir := range charts {
		if depErr := helm.RunHelmDependencyBuild(chartDir); depErr != nil {
			log.Warn().Str("chart", chartDir).Msg("Dependency build failure – skipping remaining charts in repo")
			repoResult.Charts = append(repoResult.Charts, model.ChartSummary{
				ChartPath:       chartDir,
				DepBuildFailure: true,
				DepBuildError:   depErr.Error(),
			})
			repoResult.TotalDepFailures++
			break
		}

		chartSummary := model.ChartSummary{ChartPath: chartDir}
		valuesFiles := helm.FindValuesFiles(chartDir)

		cmdStr, output, runErr := helm.RunHelmTemplate(chartDir, nil)
		run := model.RunResult{ChartPath: chartDir, HelmCommand: cmdStr, Success: runErr == nil}
		if runErr != nil {
			run.ErrorMessage = output
			chartSummary.Failures++
		} else {
			chartSummary.Successes++
		}
		chartSummary.Runs = append(chartSummary.Runs, run)
		chartSummary.TotalRuns++

		if len(valuesFiles) > 0 {
			if len(valuesFiles) > 10 {
				valuesFiles = valuesFiles[:10]
			}
			combos := helm.Combinations(valuesFiles)
			if len(combos) > 100 {
				combos = combos[:100]
			}
			for _, combo := range combos {
				cmdStr, output, runErr := helm.RunHelmTemplate(chartDir, combo)
				r := model.RunResult{ChartPath: chartDir, ValuesFiles: combo, HelmCommand: cmdStr, Success: runErr == nil}
				if runErr != nil {
					r.ErrorMessage = output
					chartSummary.Failures++
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

	if repoResult.TotalDepFailures > 0 {
		repoResult.DepFailed = true
		repoResult.Kept = false
	} else if repoResult.TotalFailures > 0 {
		repoResult.Kept = true
	}
	return repoResult
}
