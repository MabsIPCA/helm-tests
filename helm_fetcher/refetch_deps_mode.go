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

// runRefetchDepsMode re-processes ONLY the repos in catalogIn that previously
// had a dependency-build failure, using the build+update dependency fetch
// (helm.RunHelmDependencyFetch) instead of the original build-only step. Repos
// without a dep failure are left untouched. The updated catalog is written back
// to catalogIn in place, after a one-time .bak backup.
//
// The point: a chart whose subcharts could not be resolved by "helm dependency
// build" alone may now resolve via "helm dependency update" (URL/OCI deps), so
// its real template errors — not a dependency wall — land in the catalog the
// taxonomy reads.
func runRefetchDepsMode(catalogIn, cloneDir string) {
	data, err := os.ReadFile(catalogIn)
	if err != nil {
		log.Fatal().Err(err).Str("path", catalogIn).Msg("Cannot read catalog")
	}
	var repos []model.RepoResult
	if err := json.Unmarshal(data, &repos); err != nil {
		log.Fatal().Err(err).Msg("Cannot parse catalog JSON")
	}

	// One-time backup so the original snapshot is never lost.
	backup := catalogIn + ".bak"
	if _, statErr := os.Stat(backup); os.IsNotExist(statErr) {
		if writeErr := os.WriteFile(backup, data, 0o644); writeErr != nil {
			log.Warn().Err(writeErr).Str("backup", backup).Msg("Could not write .bak (continuing)")
		} else {
			log.Info().Str("backup", backup).Msg("Backed up original catalog")
		}
	}

	var targets int
	for _, r := range repos {
		if repoHadDepFailure(r) {
			targets++
		}
	}
	log.Info().Int("total_repos", len(repos)).Int("dep_failed_repos", targets).Str("clone_dir", cloneDir).Msg("Re-fetching dependencies for dep-failed repos")

	reprocessed, stillFailed, nowTemplating := 0, 0, 0
	for i := range repos {
		if !repoHadDepFailure(repos[i]) {
			continue
		}
		repo := repos[i]
		destDir := resolveDestDir(repo.RepoURL, cloneDir)

		reprocessed++
		log.Info().Int("n", reprocessed).Int("of", targets).Str("repo", repo.RepoName).Msg("Re-processing dep-failed repo")

		if cloneErr := git.CloneRepo(repo.RepoURL, destDir); cloneErr != nil {
			log.Warn().Err(cloneErr).Str("repo", repo.RepoName).Msg("Clone failed – leaving existing result unchanged")
			continue
		}

		updated, hadDep := reprocessRepoWithDepFetch(repo.RepoURL, repo.RepoName, destDir)
		repos[i] = updated
		if hadDep {
			stillFailed++
		} else {
			nowTemplating++
		}
	}

	out, err := json.MarshalIndent(repos, "", "  ")
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot marshal updated catalog")
	}
	if err := os.WriteFile(catalogIn, out, 0o644); err != nil {
		log.Fatal().Err(err).Str("path", catalogIn).Msg("Cannot write updated catalog")
	}

	fmt.Printf("\nRefetch-deps complete.\n")
	fmt.Printf("  Dep-failed repos re-processed: %d\n", reprocessed)
	fmt.Printf("  Now past dependency build:     %d\n", nowTemplating)
	fmt.Printf("  Still dependency-failed:       %d\n", stillFailed)
	fmt.Printf("  Updated catalog: %s (backup: %s)\n", catalogIn, backup)
}

// repoHadDepFailure reports whether a repo recorded any dependency-build failure.
func repoHadDepFailure(r model.RepoResult) bool {
	if r.DepFailed || r.TotalDepFailures > 0 {
		return true
	}
	for _, c := range r.Charts {
		if c.DepBuildFailure {
			return true
		}
	}
	return false
}

// reprocessRepoWithDepFetch rebuilds a RepoResult from scratch for destDir,
// fetching each chart's dependencies via build+update before templating. It
// mirrors the full-mode per-chart flow but swaps the build-only dependency step
// for helm.RunHelmDependencyFetch. The returned bool is true when at least one
// chart still failed its dependency fetch.
func reprocessRepoWithDepFetch(repoURL, repoName, destDir string) (model.RepoResult, bool) {
	repoResult := model.RepoResult{
		RepoURL:   repoURL,
		RepoName:  repoName,
		ClonedDir: destDir,
	}

	charts := helm.FindCharts(destDir)
	repoResult.TotalCharts = len(charts)

	for _, chartDir := range charts {
		// build+update so URL/OCI deps resolve even without a pre-existing lock.
		if _, depErr := helm.RunHelmDependencyFetch(chartDir); depErr != nil {
			log.Warn().Str("chart", chartDir).Msg("Dependency fetch failure – skipping remaining charts in repo")
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
	return repoResult, repoResult.TotalDepFailures > 0
}
