package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var renderProblemsDirs = []string{
	"test-01-logic-default-functions",
	"test-02-nil-pointer",
	"test-03-template-syntax-error",
	"test-04-missing-template",
	"test-05-division-by-zero",
	"test-06-invalid-regex",
	"test-07-file-read-error",
	"test-08-invalid-types",
	"test-09-list-operation-error",
	"test-10-string-operation-error",
	"test-11-map-dict-error",
	"test-12-date-time-error",
	"test-13-crypto-error",
	"test-14-helper-template",
	"test-15-tpl-function",
	"test-16-math-operations",
	"test-17-uuid-functions",
	"test-18-url-functions",
	"test-19-encoding-functions",
	"test-20-reflection-functions",
	"test-21-semver-functions",
}

var overrideTestsDirs = []string{
	"test-01-basic-data-types",
	"test-02-multiple-values-files",
	"test-03-values-with-defaults",
	"test-04-type-changes",
	"test-05-list-override",
	"test-06-deep-nested-override",
	"test-07-null-value-override",
	"test-08-subchart-override",
	"test-09-global-values-override",
	"test-10-set-flag-override",
}

type suiteSpec struct {
	name string
	path string
	dirs []string
}

func main() {
	root := flag.String("root", "", "path to helm-tests repo root (default: parent of working directory)")
	flag.Parse()

	repoRoot := *root
	if repoRoot == "" {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Fprintf(os.Stderr, "getwd: %v\n", err)
			os.Exit(1)
		}
		repoRoot = filepath.Dir(wd)
	}

	suites := []suiteSpec{
		{
			name: "render_problems",
			path: filepath.Join(repoRoot, "render_problems"),
			dirs: renderProblemsDirs,
		},
		{
			name: "override_tests_suite",
			path: filepath.Join(repoRoot, "override_tests_suite"),
			dirs: overrideTestsDirs,
		},
	}

	fmt.Println("=== KICS Render Mock ===")

	for _, suite := range suites {
		totalInvocations := 0
		totalErrors := 0

		for _, dirName := range suite.dirs {
			testDir := filepath.Join(suite.path, dirName)

			if _, err := os.Stat(testDir); os.IsNotExist(err) {
				fmt.Fprintf(os.Stderr, "[%s/%s] directory not found, skipping\n", suite.name, dirName)
				continue
			}

			out, err := runTestDir(testDir, suite.name)
			if err != nil {
				fmt.Fprintf(os.Stderr, "[%s/%s] config error: %v\n", suite.name, dirName, err)
				continue
			}

			for _, r := range out.Renders {
				totalInvocations++
				if r.Standard.Error != nil {
					totalErrors++
				}
			}

			if err := writeOutput(testDir, out); err != nil {
				fmt.Fprintf(os.Stderr, "[%s/%s] write error: %v\n", suite.name, dirName, err)
			}
		}

		fmt.Printf("%-25s %2d dirs   %4d invocations   %4d errors\n",
			suite.name, len(suite.dirs), totalInvocations, totalErrors)
	}

	fmt.Println("JSON files written to each test directory.")
}
