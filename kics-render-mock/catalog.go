package main

import (
	"encoding/json"
	"fmt"
	"os"

	"helm.sh/helm/v3/pkg/cli/values"
)

// catalogRun mirrors helm_fetcher's model.RunResult for JSON parsing.
type catalogRun struct {
	ChartPath   string   `json:"chart_path"`
	ValuesFiles []string `json:"values_files"`
	Success     bool     `json:"success"`
	ErrorMsg    string   `json:"error_message"`
}

type catalogChart struct {
	ChartPath string       `json:"chart_path"`
	Runs      []catalogRun `json:"runs"`
}

type catalogRepo struct {
	RepoURL  string         `json:"repo_url"`
	RepoName string         `json:"repo_name"`
	Charts   []catalogChart `json:"charts"`
}

// runCatalogMode reads a helm_fetcher catalog_kept.json and processes each
// chart through the KICS-faithful Helm SDK pipeline, writing
// kics-render-output.json and kics-fixed-output.json into each chart directory.
func runCatalogMode(catalogPath string) {
	data, err := os.ReadFile(catalogPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read catalog %q: %v\n", catalogPath, err)
		os.Exit(1)
	}

	var repos []catalogRepo
	if err := json.Unmarshal(data, &repos); err != nil {
		fmt.Fprintf(os.Stderr, "parse catalog: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("=== KICS Render Mock — Catalog Mode ===")

	totalInvocations := 0
	totalErrors := 0
	totalResolved := 0

	for _, repo := range repos {
		for _, chart := range repo.Charts {
			chartPath := chart.ChartPath

			if _, statErr := os.Stat(chartPath); os.IsNotExist(statErr) {
				fmt.Fprintf(os.Stderr, "[%s] chart path not found, skipping\n", chartPath)
				continue
			}

			invocations := buildCatalogInvocations(chart.Runs)

			out := RenderOutput{
				Suite:     "catalog",
				ChartPath: chartPath,
			}
			for _, inv := range invocations {
				stdRes := runOnce(chartPath, inv.valOpts, false)
				dbgRes := runOnce(chartPath, inv.valOpts, true)
				out.Renders = append(out.Renders, buildRenderEntry(inv, stdRes, dbgRes))
				totalInvocations++
				if stdRes.err != nil {
					totalErrors++
				}
			}

			if err := writeOutput(chartPath, out); err != nil {
				fmt.Fprintf(os.Stderr, "[%s] write error: %v\n", chartPath, err)
			}

			fixedOut := fixFromInvocations(chartPath, out, invocations)
			if err := writeFixedOutput(chartPath, fixedOut); err != nil {
				fmt.Fprintf(os.Stderr, "[%s] fixed write error: %v\n", chartPath, err)
			}
			for _, r := range fixedOut.Renders {
				if r.Resolved && len(r.FixChain) > 0 {
					totalResolved++
				}
			}
		}
	}

	resolved := ""
	if totalErrors > 0 {
		pct := totalResolved * 100 / totalErrors
		resolved = fmt.Sprintf(" → %d resolved (%d%%)", totalResolved, pct)
	}
	fmt.Printf("catalog   %4d invocations   %4d errors%s\n", totalInvocations, totalErrors, resolved)
	fmt.Println("kics-render-output.json + kics-fixed-output.json written to each chart directory.")
}

func buildCatalogInvocations(runs []catalogRun) []invocation {
	invs := make([]invocation, 0, len(runs))
	for _, run := range runs {
		vf := run.ValuesFiles
		if vf == nil {
			vf = []string{}
		}
		invs = append(invs, invocation{
			valOpts: &values.Options{ValueFiles: vf},
		})
	}
	return invs
}
