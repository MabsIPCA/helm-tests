// Command pick_repos writes a sub-catalog containing only the repos whose
// RepoName contains one of the given substrings. Useful for running the fixer
// against a focused slice of a large catalog.
//
// Usage:
//
//	go run ./scripts/pick_repos <in.json> <out.json> <substr> [<substr> ...]
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/MabsIPCA/helm-tests/helm_fetcher/model"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Fprintln(os.Stderr, "usage: pick_repos <in.json> <out.json> <substr> [<substr> ...]")
		os.Exit(2)
	}
	in, out := os.Args[1], os.Args[2]
	subs := os.Args[3:]

	data, err := os.ReadFile(in)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read: %v\n", err)
		os.Exit(1)
	}
	var repos []model.RepoResult
	if err := json.Unmarshal(data, &repos); err != nil {
		fmt.Fprintf(os.Stderr, "decode: %v\n", err)
		os.Exit(1)
	}

	var picked []model.RepoResult
	for _, r := range repos {
		for _, s := range subs {
			if strings.Contains(r.RepoName, s) {
				picked = append(picked, r)
				break
			}
		}
	}

	encoded, err := json.MarshalIndent(picked, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "encode: %v\n", err)
		os.Exit(1)
	}
	if err := os.WriteFile(out, encoded, 0o644); err != nil {
		fmt.Fprintf(os.Stderr, "write: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("picked %d repos -> %s\n", len(picked), out)
}
