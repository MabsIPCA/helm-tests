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
	var chain []model.FixStep

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

// runFixerMode is implemented in Task 5.
func runFixerMode(catalogIn, cloneDir string) {
	panic("not yet implemented")
}

// writeFixedJSON is implemented in Task 5.
func writeFixedJSON(repos []model.RepoResultFixed) error {
	panic("not yet implemented")
}

// writeFixerReport is implemented in Task 5.
func writeFixerReport(catalogIn string, repos []model.RepoResultFixed) error {
	panic("not yet implemented")
}

// suppress unused import errors during incremental development
var (
	_ = json.Marshal
	_ = fmt.Sprintf
	_ = os.ReadFile
	_ = time.Now
	_ = log.Info
	_ = git.CloneRepo
	_ = helm.FindCharts
)
