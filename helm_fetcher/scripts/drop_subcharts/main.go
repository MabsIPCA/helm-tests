// Command drop_subcharts removes non-standalone charts from one or more catalog
// JSON files — both vendored subcharts (helm.IsVendoredSubchart) and file://
// relative components of a sibling umbrella (helm.RelativeComponentDirs) — then
// recomputes repo/chart totals and rewrites each file in place after backing it
// up to <file>.bak. Rendering either standalone is a false failure.
//
// It decodes every file as []model.RepoResultFixed, which is a superset of the
// plain []model.RepoResult shape, so both input and fixed catalogs round-trip
// cleanly (fix fields are omitempty and stay absent when not present).
//
// Usage:
//
//	go run ./scripts/drop_subcharts <catalog.json> [<catalog.json> ...]
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/MabsIPCA/helm-tests/helm_fetcher/helm"
	"github.com/MabsIPCA/helm-tests/helm_fetcher/model"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: drop_subcharts <catalog.json> [<catalog.json> ...]")
		os.Exit(2)
	}
	for _, path := range os.Args[1:] {
		if err := process(path); err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", path, err)
			os.Exit(1)
		}
	}
}

func process(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read: %w", err)
	}
	var repos []model.RepoResultFixed
	if err := json.Unmarshal(data, &repos); err != nil {
		return fmt.Errorf("decode: %w", err)
	}

	var chartsDropped, runsDropped int

	// Drop non-standalone charts (vendored subcharts + file:// relative
	// components) — rendering them alone is a false failure. Repos are always
	// preserved (even if they end up chart-less) so we never remove a
	// legitimately-cataloged repo — that would make a future run re-process it.
	out := repos[:0]
	for _, r := range repos {
		// Components are repo-scoped: scan this repo's clone for file:// deps.
		var components map[string]bool
		if r.ClonedDir != "" {
			components = helm.RelativeComponentDirs(r.ClonedDir)
		}
		keptCharts := r.Charts[:0]
		for _, c := range r.Charts {
			if helm.IsVendoredSubchart(c.ChartPath) || components[filepath.Clean(c.ChartPath)] {
				chartsDropped++
				runsDropped += len(c.Runs)
				continue
			}
			keptCharts = append(keptCharts, c)
		}
		r.Charts = keptCharts
		recompute(&r)
		out = append(out, r)
	}

	if chartsDropped == 0 {
		fmt.Printf("%-46s no non-standalone charts found (unchanged)\n", path)
		return nil
	}

	// Back up the original once (never clobber an existing .bak).
	bak := path + ".bak"
	if _, statErr := os.Stat(bak); os.IsNotExist(statErr) {
		if err := os.WriteFile(bak, data, 0o644); err != nil {
			return fmt.Errorf("write backup: %w", err)
		}
	}

	encoded, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		return fmt.Errorf("encode: %w", err)
	}
	if err := os.WriteFile(path, encoded, 0o644); err != nil {
		return fmt.Errorf("write: %w", err)
	}

	fmt.Printf("%-46s %d repos kept, dropped %d non-standalone charts / %d runs (backup: %s)\n",
		path, len(out), chartsDropped, runsDropped, bak)
	return nil
}

// recompute refreshes the repo-level rollup counts after charts are removed.
func recompute(r *model.RepoResultFixed) {
	r.TotalCharts = len(r.Charts)
	r.TotalRuns, r.TotalSuccesses, r.TotalFailures, r.TotalDepFailures = 0, 0, 0, 0
	for _, c := range r.Charts {
		r.TotalRuns += c.TotalRuns
		r.TotalSuccesses += c.Successes
		r.TotalFailures += c.Failures
		if c.DepBuildFailure {
			r.TotalDepFailures++
		}
	}
	r.DepFailed = r.TotalDepFailures > 0
}
