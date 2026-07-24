// Command listcharts enumerates chart source directories from the helm_fetcher
// clones using the exact same discovery as a fetcher run (helm.FindCharts, which
// skips vendored subcharts and file:// umbrella components). It reads the set of
// clone directories from a catalog produced by the fetcher, converts the stored
// Windows paths to WSL paths, and prints one discovered chart directory per line.
//
//	go run ./cmd/listcharts -catalog catalog_cumulative.json > charts.txt
//	go run ./cmd/listcharts -catalog catalog_cumulative.json -count
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	"strings"

	"github.com/MabsIPCA/helm-tests/helm_fetcher/helm"
)

type repo struct {
	RepoURL   string `json:"repo_url"`
	RepoName  string `json:"repo_name"`
	ClonedDir string `json:"cloned_dir"`
	Kept      bool   `json:"kept"`
	DepFailed bool   `json:"dep_failed"`
}

// loadRepoSet reads a newline-delimited file of repo URLs to keep (blank/`#` lines skipped).
func loadRepoSet(path string) map[string]bool {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "listcharts: -repos:", err)
		os.Exit(1)
	}
	set := map[string]bool{}
	for _, ln := range strings.Split(string(data), "\n") {
		ln = strings.TrimSpace(ln)
		if ln != "" && !strings.HasPrefix(ln, "#") {
			set[ln] = true
		}
	}
	return set
}

var winDrive = regexp.MustCompile(`^([A-Za-z]):[\\/](.*)$`)

// toWSL converts `D:\a\b` -> `/mnt/d/a/b`; leaves already-POSIX paths untouched.
// Output is always emitted in this form so the WSL scanners can consume it,
// regardless of whether discovery ran natively on Windows or inside WSL.
func toWSL(p string) string {
	if m := winDrive.FindStringSubmatch(p); m != nil {
		return "/mnt/" + strings.ToLower(m[1]) + "/" + strings.ReplaceAll(m[2], `\`, "/")
	}
	return p
}

// walkBase is the path to hand filepath.Walk: native `D:\…` when running as a
// Windows binary (fast, no 9p), or the `/mnt/d/…` form when running inside WSL.
func walkBase(clonedDir string) string {
	if runtime.GOOS == "windows" {
		return clonedDir
	}
	return toWSL(clonedDir)
}

func main() {
	catalog := flag.String("catalog", "catalog_cumulative.json", "fetcher catalog json")
	countOnly := flag.Bool("count", false, "print only summary counts")
	only := flag.String("only", "", "subset filter: kept | dep_failed | (empty=all)")
	reposFile := flag.String("repos", "", "file of repo URLs to restrict discovery to (one per line)")
	flag.Parse()

	var repoSet map[string]bool
	if *reposFile != "" {
		repoSet = loadRepoSet(*reposFile)
	}

	data, err := os.ReadFile(*catalog)
	if err != nil {
		fmt.Fprintln(os.Stderr, "listcharts:", err)
		os.Exit(1)
	}
	var repos []repo
	if err := json.Unmarshal(data, &repos); err != nil {
		fmt.Fprintln(os.Stderr, "listcharts: parse:", err)
		os.Exit(1)
	}

	// dedupe clone dirs (a repo can appear once, but be defensive)
	seenClone := map[string]bool{}
	var clones []repo
	for _, r := range repos {
		if repoSet != nil && !repoSet[r.RepoURL] {
			continue
		}
		switch *only {
		case "kept":
			if !r.Kept {
				continue
			}
		case "dep_failed":
			if !r.DepFailed {
				continue
			}
		}
		cd := walkBase(r.ClonedDir)
		if cd == "" || seenClone[cd] {
			continue
		}
		seenClone[cd] = true
		r.ClonedDir = cd
		clones = append(clones, r)
	}

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var totalCharts, presentClones, missingClones int
	seenChart := map[string]bool{}
	for _, r := range clones {
		if fi, err := os.Stat(r.ClonedDir); err != nil || !fi.IsDir() {
			missingClones++
			continue
		}
		presentClones++
		for _, chartDir := range helm.FindCharts(r.ClonedDir) {
			c := filepath.Clean(chartDir)
			if seenChart[c] {
				continue
			}
			seenChart[c] = true
			totalCharts++
			if !*countOnly {
				fmt.Fprintln(out, toWSL(c)) // always emit /mnt/d form for WSL scanners
			}
		}
	}

	_ = sort.Strings
	fmt.Fprintf(os.Stderr,
		"listcharts: clones present=%d missing=%d | discovered charts=%d (filter=%q)\n",
		presentClones, missingClones, totalCharts, *only)
}
