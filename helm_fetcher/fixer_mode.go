package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/MabsIPCA/helm-tests/helm_fetcher/git"
	"github.com/MabsIPCA/helm-tests/helm_fetcher/helm"
	"github.com/MabsIPCA/helm-tests/helm_fetcher/model"
)

const maxFixIterations = 10

var (
	nilPtrRe   = regexp.MustCompile(`at <(\.Values\.[^ >|]+)>`)
	requiredRe = regexp.MustCompile(`required (\.Values\.[^ "]+)`)
)

// parseHelmCLIError parses a helm CLI error string and returns fix instructions.
// ok is false for errors that cannot be fixed by value injection.
func parseHelmCLIError(errMsg string) (kind, valuePath, injectedValue string, ok bool) {
	s := strings.TrimPrefix(errMsg, "Error: ")

	if strings.Contains(s, "nil pointer") {
		if m := nilPtrRe.FindStringSubmatch(s); len(m) == 2 {
			path := strings.TrimPrefix(m[1], ".Values.")
			return "nil_pointer", path, "", true
		}
	}
	if strings.Contains(s, "error calling required") {
		if m := requiredRe.FindStringSubmatch(s); len(m) == 2 {
			path := strings.TrimPrefix(m[1], ".Values.")
			return "required_value", path, "kics-placeholder", true
		}
	}
	return "", "", "", false
}

// fixRun runs the fix loop for a single failing RunResult.
// If orig.Success is true, it returns immediately with Resolved=true and an empty chain.
func fixRun(chartPath string, orig model.RunResult) model.FixedRunResult {
	if orig.Success {
		return model.FixedRunResult{Resolved: true, FixChain: []model.FixStep{}}
	}

	patch := map[string]string{}
	seenErrs := map[string]bool{}
	chain := []model.FixStep{}

	for attempt := 1; attempt <= maxFixIterations; attempt++ {
		setFlags := make([]string, 0, len(patch))
		for k, v := range patch {
			setFlags = append(setFlags, k+"="+v)
		}

		cmdStr, output, err := helm.RunHelmTemplateWithSets(chartPath, orig.ValuesFiles, setFlags)

		if err == nil {
			return model.FixedRunResult{
				Resolved:     true,
				FixChain:     chain,
				FinalCommand: cmdStr,
			}
		}

		errStr := output
		if seenErrs[errStr] {
			return model.FixedRunResult{
				StopReason: "loop_detected",
				FixChain:   chain,
			}
		}
		seenErrs[errStr] = true

		kind, path, val, ok := parseHelmCLIError(errStr)
		if !ok {
			return model.FixedRunResult{
				StopReason: "unfixable_error_kind",
				FixChain:   chain,
			}
		}

		chain = append(chain, model.FixStep{
			Attempt:       attempt,
			ErrorSeen:     errStr,
			Kind:          kind,
			ValuePath:     path,
			ValueInjected: val,
		})
		patch[path] = val
	}

	return model.FixedRunResult{
		StopReason: "max_iterations",
		FixChain:   chain,
	}
}

// resolveDestDir derives the clone destination directory from repoURL and cloneDir,
// using the same algorithm as main.go so paths stay consistent.
func resolveDestDir(repoURL, cloneDir string) string {
	repoName := strings.TrimPrefix(repoURL, "https://github.com/")
	safeName := strings.ReplaceAll(repoName, "/", "__")
	return filepath.Join(cloneDir, safeName)
}

// runFixerMode reads catalogIn, clones missing repos, runs the fix loop on every
// failing helm template run, and writes catalog_fixed.json + catalog_fixer_report.md.
func runFixerMode(catalogIn, cloneDir string) {
	data, err := os.ReadFile(catalogIn)
	if err != nil {
		log.Fatal().Err(err).Str("path", catalogIn).Msg("Cannot read catalog")
	}
	var repos []model.RepoResult
	if err := json.Unmarshal(data, &repos); err != nil {
		log.Fatal().Err(err).Msg("Cannot parse catalog JSON")
	}

	var fixedRepos []model.RepoResultFixed
	var skipped int

	for i, repo := range repos {
		destDir := resolveDestDir(repo.RepoURL, cloneDir)

		log.Info().
			Int("repo_index", i+1).
			Int("total", len(repos)).
			Str("repo", repo.RepoName).
			Msg("Processing repo")

		if cloneErr := git.CloneRepo(repo.RepoURL, destDir); cloneErr != nil {
			log.Warn().Err(cloneErr).Str("repo", repo.RepoName).Msg("Clone failed – skipping")
			skipped++
			continue
		}

		freshCharts := helm.FindCharts(destDir)
		if len(freshCharts) == 0 {
			log.Warn().Str("repo", repo.RepoName).Msg("No charts found – skipping")
			skipped++
			continue
		}

		// Index original chart summaries by their path relative to the original clonedDir.
		// This lets us match even when cloneDir has changed between the original run and now.
		origByRel := map[string]model.ChartSummary{}
		for _, cs := range repo.Charts {
			rel := strings.TrimPrefix(cs.ChartPath, repo.ClonedDir)
			rel = strings.TrimPrefix(rel, string(filepath.Separator))
			origByRel[rel] = cs
		}

		fixedRepo := model.RepoResultFixed{
			RepoURL:   repo.RepoURL,
			RepoName:  repo.RepoName,
			ClonedDir: destDir,
			Kept:      repo.Kept,
			DepFailed: repo.DepFailed,
		}

		for _, chartDir := range freshCharts {
			rel := strings.TrimPrefix(chartDir, destDir)
			rel = strings.TrimPrefix(rel, string(filepath.Separator))

			origChart, found := origByRel[rel]
			if !found {
				log.Warn().Str("chart", chartDir).Msg("Chart not in catalog – skipping")
				continue
			}

			fixedChart := model.ChartSummaryFixed{
				ChartPath:       chartDir,
				DepBuildFailure: origChart.DepBuildFailure,
				DepBuildError:   origChart.DepBuildError,
			}

			for _, run := range origChart.Runs {
				// Remap values files to the current destDir in case cloneDir changed.
				remappedRun := run
				if len(run.ValuesFiles) > 0 {
					remapped := make([]string, len(run.ValuesFiles))
					for j, vf := range run.ValuesFiles {
						vfRel := strings.TrimPrefix(vf, repo.ClonedDir)
						vfRel = strings.TrimPrefix(vfRel, string(filepath.Separator))
						remapped[j] = filepath.Join(destDir, vfRel)
					}
					remappedRun.ValuesFiles = remapped
				}

				withFix := model.RunResultWithFix{RunResult: remappedRun}
				if !run.Success {
					result := fixRun(chartDir, remappedRun)
					withFix.Fixed = &result
				}

				fixedChart.Runs = append(fixedChart.Runs, withFix)
				fixedChart.TotalRuns++
				if run.Success {
					fixedChart.Successes++
				} else {
					fixedChart.Failures++
				}
			}

			fixedRepo.Charts = append(fixedRepo.Charts, fixedChart)
			fixedRepo.TotalCharts++
			fixedRepo.TotalRuns += fixedChart.TotalRuns
			fixedRepo.TotalSuccesses += fixedChart.Successes
			fixedRepo.TotalFailures += fixedChart.Failures
			if fixedChart.DepBuildFailure {
				fixedRepo.TotalDepFailures++
			}
		}

		fixedRepos = append(fixedRepos, fixedRepo)
	}

	if err := writeFixedJSON(fixedRepos); err != nil {
		log.Fatal().Err(err).Msg("Failed to write catalog_fixed.json")
	}
	log.Info().Msg("catalog_fixed.json written")

	if err := writeFixerReport(catalogIn, fixedRepos); err != nil {
		log.Error().Err(err).Msg("Failed to write catalog_fixer_report.md (continuing)")
	} else {
		log.Info().Msg("catalog_fixer_report.md written")
	}

	// Print summary
	var totalFailing, totalResolved int
	for _, r := range fixedRepos {
		for _, c := range r.Charts {
			for _, run := range c.Runs {
				if run.Fixed != nil {
					totalFailing++
					if run.Fixed.Resolved {
						totalResolved++
					}
				}
			}
		}
	}
	pct := 0
	if totalFailing > 0 {
		pct = totalResolved * 100 / totalFailing
	}
	fmt.Printf("\nCatalog Fixer complete.\n")
	fmt.Printf("  Repos processed: %d\n", len(fixedRepos))
	if skipped > 0 {
		fmt.Printf("  Repos skipped:   %d (clone failed or no charts)\n", skipped)
	}
	fmt.Printf("  Failing runs:    %d\n", totalFailing)
	fmt.Printf("  Resolved:        %d\n", totalResolved)
	fmt.Printf("  Resolution rate: %d%%\n", pct)
	fmt.Printf("\nOutput files:\n")
	fmt.Printf("  catalog_fixed.json\n")
	fmt.Printf("  catalog_fixer_report.md\n")
}

func writeFixedJSON(repos []model.RepoResultFixed) error {
	data, err := json.MarshalIndent(repos, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}
	return os.WriteFile("catalog_fixed.json", data, 0o644)
}

func writeFixerReport(catalogIn string, repos []model.RepoResultFixed) error {
	type kindStats struct{ before, resolved int }
	stats := map[string]*kindStats{
		"nil_pointer":    {},
		"required_value": {},
		"other":          {},
	}
	stopReasons := map[string]int{}
	var totalFailing, totalResolved int

	for _, r := range repos {
		for _, c := range r.Charts {
			for _, run := range c.Runs {
				if run.Fixed == nil {
					continue
				}
				totalFailing++
				if run.Fixed.Resolved {
					totalResolved++
					stopReasons["resolved"]++
				} else {
					stopReasons[run.Fixed.StopReason]++
				}

				// Classify by first fix step kind; if no steps, it was immediately unfixable.
				if len(run.Fixed.FixChain) == 0 {
					stats["other"].before++
					if run.Fixed.Resolved {
						stats["other"].resolved++
					}
				} else {
					k := run.Fixed.FixChain[0].Kind
					if _, known := stats[k]; !known {
						stats["other"].before++
						if run.Fixed.Resolved {
							stats["other"].resolved++
						}
					} else {
						stats[k].before++
						if run.Fixed.Resolved {
							stats[k].resolved++
						}
					}
				}
			}
		}
	}

	pct := 0
	if totalFailing > 0 {
		pct = totalResolved * 100 / totalFailing
	}

	date := time.Now().Format("2006-01-02")
	var sb strings.Builder

	fmt.Fprintf(&sb, "# Catalog Fixer Report\n\n")
	fmt.Fprintf(&sb, "**Source:** %s\n", catalogIn)
	fmt.Fprintf(&sb, "**Date:** %s\n\n", date)

	fmt.Fprintf(&sb, "## Summary\n\n")
	fmt.Fprintf(&sb, "| Metric | Count |\n|---|---:|\n")
	fmt.Fprintf(&sb, "| Repos processed | %d |\n", len(repos))
	fmt.Fprintf(&sb, "| Failing runs (before) | %d |\n", totalFailing)
	fmt.Fprintf(&sb, "| Resolved | %d |\n", totalResolved)
	fmt.Fprintf(&sb, "| Resolution rate | %d%% |\n\n", pct)

	fmt.Fprintf(&sb, "## By Error Kind\n\n")
	fmt.Fprintf(&sb, "| Kind | Before | Resolved | Still failing |\n|---|---:|---:|---:|\n")
	for _, k := range []string{"nil_pointer", "required_value", "other"} {
		s := stats[k]
		fmt.Fprintf(&sb, "| %s | %d | %d | %d |\n", k, s.before, s.resolved, s.before-s.resolved)
	}

	fmt.Fprintf(&sb, "\n## Stop Reasons\n\n")
	fmt.Fprintf(&sb, "| Reason | Count |\n|---|---:|\n")
	for _, reason := range []string{"resolved", "unfixable_error_kind", "loop_detected", "max_iterations"} {
		fmt.Fprintf(&sb, "| %s | %d |\n", reason, stopReasons[reason])
	}

	return os.WriteFile("catalog_fixer_report.md", []byte(sb.String()), 0o644)
}
