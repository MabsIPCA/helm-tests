// Command merge_sources concatenates two or more SOURCE catalog JSON files
// (each a []model.RepoResult array) into a single output file. It is the
// source-catalog counterpart of merge_fixed (which merges []RepoResultFixed),
// used to rebuild catalog_sources_merged.json from the per-source catalogs.
//
// Usage:
//
//	go run ./scripts/merge_sources <out.json> <in1.json> <in2.json> [<in3.json> ...]
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/MabsIPCA/helm-tests/helm_fetcher/model"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Fprintln(os.Stderr, "usage: merge_sources <out.json> <in1.json> <in2.json> [<in3.json> ...]")
		os.Exit(2)
	}
	out := os.Args[1]
	ins := os.Args[2:]

	var merged []model.RepoResult
	for _, in := range ins {
		data, err := os.ReadFile(in)
		if err != nil {
			fmt.Fprintf(os.Stderr, "read %s: %v\n", in, err)
			os.Exit(1)
		}
		var repos []model.RepoResult
		if err := json.Unmarshal(data, &repos); err != nil {
			fmt.Fprintf(os.Stderr, "decode %s: %v\n", in, err)
			os.Exit(1)
		}
		merged = append(merged, repos...)
		fmt.Printf("  + %-50s %d repos\n", in, len(repos))
	}

	encoded, err := json.MarshalIndent(merged, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "encode: %v\n", err)
		os.Exit(1)
	}
	if err := os.WriteFile(out, encoded, 0o644); err != nil {
		fmt.Fprintf(os.Stderr, "write %s: %v\n", out, err)
		os.Exit(1)
	}
	fmt.Printf("merged %d repos -> %s\n", len(merged), out)
}
