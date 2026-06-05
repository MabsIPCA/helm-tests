// Command subchart_audit quantifies how many template failures in a fixed
// catalog come from rendering a vendored subchart standalone (a false failure
// the fetcher now avoids — see helm.FindCharts / helm.IsVendoredSubchart).
//
// Usage:
//
//	go run ./scripts/subchart_audit <catalog_fixed.json>
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/MabsIPCA/helm-tests/helm_fetcher/helm"
	"github.com/MabsIPCA/helm-tests/helm_fetcher/model"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: subchart_audit <catalog_fixed.json>")
		os.Exit(2)
	}
	path := os.Args[1]

	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read %s: %v\n", path, err)
		os.Exit(1)
	}
	var repos []model.RepoResultFixed
	if err := json.Unmarshal(data, &repos); err != nil {
		fmt.Fprintf(os.Stderr, "decode %s: %v\n", path, err)
		os.Exit(1)
	}

	var totalFail, failSub int
	var unresolved, unresolvedSub int
	subByRepo := map[string]int{}

	for _, r := range repos {
		for _, c := range r.Charts {
			for _, run := range c.Runs {
				if run.Success || run.ErrorMessage == "" {
					continue
				}
				totalFail++
				isSub := helm.IsVendoredSubchart(run.ChartPath)
				if isSub {
					failSub++
					subByRepo[r.RepoName]++
				}
				if run.Fixed != nil && !run.Fixed.Resolved {
					unresolved++
					if isSub {
						unresolvedSub++
					}
				}
			}
		}
	}

	pct := func(n, d int) float64 {
		if d == 0 {
			return 0
		}
		return 100 * float64(n) / float64(d)
	}

	fmt.Printf("Catalog: %s\n\n", path)
	fmt.Printf("Template failures          : %d\n", totalFail)
	fmt.Printf("  from vendored subcharts  : %d (%.1f%%)\n", failSub, pct(failSub, totalFail))
	fmt.Printf("Unresolved failures        : %d\n", unresolved)
	fmt.Printf("  from vendored subcharts  : %d (%.1f%%)\n", unresolvedSub, pct(unresolvedSub, unresolved))
	fmt.Printf("\nTop repos by subchart failures:\n")

	type kv struct {
		repo string
		n    int
	}
	rows := make([]kv, 0, len(subByRepo))
	for k, v := range subByRepo {
		rows = append(rows, kv{k, v})
	}
	sort.Slice(rows, func(i, j int) bool {
		if rows[i].n != rows[j].n {
			return rows[i].n > rows[j].n
		}
		return rows[i].repo < rows[j].repo
	})
	for i, row := range rows {
		if i >= 15 {
			fmt.Printf("  ... and %d more repos\n", len(rows)-15)
			break
		}
		fmt.Printf("  %4d  %s\n", row.n, row.repo)
	}
}
