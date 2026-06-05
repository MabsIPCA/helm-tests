// Command unresolved_summary groups unresolved runs in a fixed catalog by the
// first fix-chain kind and a normalized final_error, so you can see WHY a class
// of failures is still failing.
//
// Usage:
//
//	go run ./scripts/unresolved_summary <catalog_fixed.json> [firstKindFilter]
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/MabsIPCA/helm-tests/helm_fetcher/model"
)

var (
	envelopeRe = regexp.MustCompile(`execution error at \([^)]*\):\s*`)
	digitsRe   = regexp.MustCompile(`\d+`)
	upperRe    = regexp.MustCompile(`[A-Z][A-Z0-9_]{2,}`)
)

func normalize(err string) string {
	s := strings.TrimPrefix(strings.TrimSpace(err), "Error: ")
	s = envelopeRe.ReplaceAllString(s, "")
	if i := strings.Index(s, "\n"); i >= 0 {
		s = s[:i]
	}
	s = upperRe.ReplaceAllString(s, "<KEY>")
	s = digitsRe.ReplaceAllString(s, "N")
	if len(s) > 130 {
		s = s[:130]
	}
	return strings.TrimSpace(s)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: unresolved_summary <catalog_fixed.json> [firstKindFilter]")
		os.Exit(2)
	}
	filter := ""
	if len(os.Args) > 2 {
		filter = os.Args[2]
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "read: %v\n", err)
		os.Exit(1)
	}
	var repos []model.RepoResultFixed
	if err := json.Unmarshal(data, &repos); err != nil {
		fmt.Fprintf(os.Stderr, "decode: %v\n", err)
		os.Exit(1)
	}

	type key struct{ kind, norm string }
	counts := map[key]int{}
	examples := map[key]string{}
	total := 0

	for _, r := range repos {
		for _, c := range r.Charts {
			for _, run := range c.Runs {
				f := run.Fixed
				if f == nil || f.Resolved {
					continue
				}
				kind := "(immediate)"
				if len(f.FixChain) > 0 {
					kind = f.FixChain[0].Kind
				}
				if filter != "" && kind != filter {
					continue
				}
				total++
				k := key{kind, normalize(f.FinalError)}
				counts[k]++
				if examples[k] == "" {
					examples[k] = c.ChartPath
				}
			}
		}
	}

	type row struct {
		k key
		n int
	}
	rows := make([]row, 0, len(counts))
	for k, n := range counts {
		rows = append(rows, row{k, n})
	}
	sort.Slice(rows, func(i, j int) bool { return rows[i].n > rows[j].n })

	fmt.Printf("Unresolved runs%s: %d\n\n", filterLabel(filter), total)
	for _, r := range rows {
		fmt.Printf("%4d  [%s] %s\n", r.n, r.k.kind, r.k.norm)
		fmt.Printf("        e.g. %s\n", examples[r.k])
	}
}

func filterLabel(f string) string {
	if f == "" {
		return ""
	}
	return " (first kind = " + f + ")"
}
