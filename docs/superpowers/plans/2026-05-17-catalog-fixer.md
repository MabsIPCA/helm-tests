# Catalog Fixer Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Add a `fixer` mode to `helm_fetcher` that reads `catalog_kept.json`, clones missing repos, runs a retry-fix loop on every failing `helm template` run, and writes `catalog_fixed.json` + `catalog_fixer_report.md`.

**Architecture:** The fixer reuses `git.CloneRepo`, `helm.FindCharts`, and `helm.RunHelmTemplate` from the existing pipeline, adds `helm.RunHelmTemplateWithSets` for patched re-runs, and implements the fix loop (same algorithm as kics-render-mock/fixer.go) in a new top-level file `fixer_mode.go`. New model types extend `model.go` for the fixer output JSON schema.

**Tech Stack:** Go 1.21, zerolog, standard library (`encoding/json`, `regexp`, `os/exec`, `strings`), existing `helm_fetcher` packages.

---

## File Map

| Action | File |
|--------|------|
| Modify | `helm_fetcher/model/model.go` |
| Create | `helm_fetcher/model/model_test.go` |
| Modify | `helm_fetcher/helm/helm.go` |
| Create | `helm_fetcher/helm/helm_test.go` |
| Modify | `helm_fetcher/config/config.go` |
| Create | `helm_fetcher/fixer_mode.go` |
| Create | `helm_fetcher/fixer_mode_test.go` |
| Modify | `helm_fetcher/main.go` |
| Modify | `helm_fetcher/Makefile` |

---

### Task 1: Add Fixer Model Types

**Files:**
- Modify: `helm_fetcher/model/model.go`
- Create: `helm_fetcher/model/model_test.go`

- [ ] **Step 1: Write failing test**

Create `helm_fetcher/model/model_test.go`:

```go
package model_test

import (
	"encoding/json"
	"testing"

	"github.com/MabsIPCA/helm-tests/helm_fetcher/model"
)

func TestFixStep_JSONRoundtrip(t *testing.T) {
	step := model.FixStep{
		Attempt:       1,
		ErrorSeen:     "Error: nil pointer evaluating interface {}.foo",
		Kind:          "nil_pointer",
		ValuePath:     "service.type",
		ValueInjected: "",
	}
	data, err := json.Marshal(step)
	if err != nil {
		t.Fatal(err)
	}
	var got model.FixStep
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}
	if got != step {
		t.Errorf("roundtrip: got %+v, want %+v", got, step)
	}
}

func TestFixedRunResult_JSONKeys(t *testing.T) {
	r := model.FixedRunResult{
		Resolved:     true,
		StopReason:   "",
		FixChain:     []model.FixStep{},
		FinalCommand: "helm template test /chart --set foo=bar",
	}
	data, err := json.Marshal(r)
	if err != nil {
		t.Fatal(err)
	}
	m := map[string]interface{}{}
	json.Unmarshal(data, &m)
	for _, key := range []string{"resolved", "stop_reason", "fix_chain", "final_command"} {
		if _, ok := m[key]; !ok {
			t.Errorf("missing JSON key %q", key)
		}
	}
}

func TestRunResultWithFix_EmbedPromotesFields(t *testing.T) {
	orig := model.RunResult{
		ChartPath:   "/some/chart",
		HelmCommand: "helm template test /some/chart",
		Success:     false,
		ErrorMessage: "Error: nil pointer",
	}
	withFix := model.RunResultWithFix{RunResult: orig}
	data, err := json.Marshal(withFix)
	if err != nil {
		t.Fatal(err)
	}
	m := map[string]interface{}{}
	json.Unmarshal(data, &m)
	for _, key := range []string{"chart_path", "helm_command", "success", "error_message"} {
		if _, ok := m[key]; !ok {
			t.Errorf("embedded RunResult field %q not promoted to JSON", key)
		}
	}
	if _, ok := m["fixed"]; ok {
		t.Error("fixed field should be omitted when nil")
	}
}
```

- [ ] **Step 2: Run test to verify it fails**

```
cd helm_fetcher
go test ./model/... -v -run TestFix
```

Expected: compile error — `model.FixStep`, `model.FixedRunResult`, `model.RunResultWithFix` undefined.

- [ ] **Step 3: Add types to model.go**

Append to the end of `helm_fetcher/model/model.go`:

```go

// FixStep records one iteration of the fix loop.
type FixStep struct {
	Attempt       int    `json:"attempt"`
	ErrorSeen     string `json:"error_seen"`
	Kind          string `json:"kind"`           // "nil_pointer" | "required_value"
	ValuePath     string `json:"value_path"`
	ValueInjected string `json:"value_injected"`
}

// FixedRunResult holds the fixer outcome for a single helm template run.
type FixedRunResult struct {
	Resolved     bool      `json:"resolved"`
	StopReason   string    `json:"stop_reason"` // "" | "loop_detected" | "unfixable_error_kind" | "max_iterations"
	FixChain     []FixStep `json:"fix_chain"`
	FinalCommand string    `json:"final_command,omitempty"`
}

// RunResultWithFix embeds RunResult and adds the fixer output.
type RunResultWithFix struct {
	RunResult
	Fixed *FixedRunResult `json:"fixed,omitempty"`
}

// ChartSummaryFixed is ChartSummary with RunResultWithFix entries.
type ChartSummaryFixed struct {
	ChartPath       string             `json:"chart_path"`
	TotalRuns       int                `json:"total_runs"`
	Successes       int                `json:"successes"`
	Failures        int                `json:"failures"`
	DepBuildFailure bool               `json:"dep_build_failure,omitempty"`
	DepBuildError   string             `json:"dep_build_error,omitempty"`
	Runs            []RunResultWithFix `json:"runs"`
}

// RepoResultFixed is RepoResult with ChartSummaryFixed entries.
type RepoResultFixed struct {
	RepoURL          string              `json:"repo_url"`
	RepoName         string              `json:"repo_name"`
	ClonedDir        string              `json:"cloned_dir"`
	TotalCharts      int                 `json:"total_charts"`
	TotalRuns        int                 `json:"total_runs"`
	TotalSuccesses   int                 `json:"total_successes"`
	TotalFailures    int                 `json:"total_failures"`
	TotalDepFailures int                 `json:"total_dep_failures"`
	Charts           []ChartSummaryFixed `json:"charts"`
	Kept             bool                `json:"kept"`
	DepFailed        bool                `json:"dep_failed"`
}
```

- [ ] **Step 4: Run tests to verify they pass**

```
cd helm_fetcher
go test ./model/... -v -run TestFix
```

Expected: `PASS` for all three tests.

- [ ] **Step 5: Verify existing code still compiles**

```
cd helm_fetcher
go build ./...
```

Expected: no errors.

- [ ] **Step 6: Commit**

```
git add helm_fetcher/model/model.go helm_fetcher/model/model_test.go
git commit -m "feat(helm_fetcher): add fixer model types (FixStep, FixedRunResult, RunResultWithFix)"
```

---

### Task 2: Add RunHelmTemplateWithSets

**Files:**
- Modify: `helm_fetcher/helm/helm.go`
- Create: `helm_fetcher/helm/helm_test.go`

- [ ] **Step 1: Write failing test**

Create `helm_fetcher/helm/helm_test.go`:

```go
package helm_test

import (
	"strings"
	"testing"

	"github.com/MabsIPCA/helm-tests/helm_fetcher/helm"
)

func TestRunHelmTemplateWithSets_CmdStr(t *testing.T) {
	// cmdStr is always built before exec; helm binary need not be present for this test.
	cmdStr, _, _ := helm.RunHelmTemplateWithSets(
		"/chart",
		[]string{"/values.yaml"},
		[]string{"foo=bar", "baz=qux"},
	)
	if !strings.HasPrefix(cmdStr, "helm template test /chart") {
		t.Errorf("unexpected cmdStr prefix: %q", cmdStr)
	}
	if !strings.Contains(cmdStr, "-f /values.yaml") {
		t.Errorf("cmdStr missing -f flag: %q", cmdStr)
	}
	if !strings.Contains(cmdStr, "--set foo=bar") {
		t.Errorf("cmdStr missing --set foo=bar: %q", cmdStr)
	}
	if !strings.Contains(cmdStr, "--set baz=qux") {
		t.Errorf("cmdStr missing --set baz=qux: %q", cmdStr)
	}
}

func TestRunHelmTemplateWithSets_NoExtras(t *testing.T) {
	cmdStr, _, _ := helm.RunHelmTemplateWithSets("/chart", nil, nil)
	want := "helm template test /chart"
	if cmdStr != want {
		t.Errorf("got %q, want %q", cmdStr, want)
	}
}
```

- [ ] **Step 2: Run test to verify it fails**

```
cd helm_fetcher
go test ./helm/... -v -run TestRunHelmTemplateWithSets
```

Expected: compile error — `helm.RunHelmTemplateWithSets` undefined.

- [ ] **Step 3: Implement RunHelmTemplateWithSets in helm/helm.go**

Append to the end of `helm_fetcher/helm/helm.go`:

```go

// RunHelmTemplateWithSets runs "helm template test <chartPath> [-f vf...] [--set k=v...]"
// and returns the command string, combined output, and any exec error.
func RunHelmTemplateWithSets(chartPath string, valFiles, setFlags []string) (cmdStr, output string, err error) {
	args := []string{"template", "test", chartPath}
	for _, vf := range valFiles {
		args = append(args, "-f", vf)
	}
	for _, sf := range setFlags {
		args = append(args, "--set", sf)
	}
	cmdStr = "helm " + strings.Join(args, " ")

	cmd := exec.Command("helm", args...)
	out, runErr := cmd.CombinedOutput()
	return cmdStr, string(out), runErr
}
```

- [ ] **Step 4: Run tests to verify they pass**

```
cd helm_fetcher
go test ./helm/... -v -run TestRunHelmTemplateWithSets
```

Expected: `PASS` for both tests.

- [ ] **Step 5: Commit**

```
git add helm_fetcher/helm/helm.go helm_fetcher/helm/helm_test.go
git commit -m "feat(helm_fetcher/helm): add RunHelmTemplateWithSets"
```

---

### Task 3: Add CatalogIn Config Flag

**Files:**
- Modify: `helm_fetcher/config/config.go`

No new test: config parsing is integration-level; it is exercised by the E2E step in Task 6.

- [ ] **Step 1: Add constant, struct field, and flag registration**

In `helm_fetcher/config/config.go`, add after the `cloneDirDefault` block (line 33) and into `Config` and `Parse()`:

Replace the constants block to add `catalogInFlag` / `catalogInDefault`:

```go
const (
	modeFlag           = "mode"
	modeUsage          = "Execution mode: full, github-search-json, artifacthub-search-json, or fixer"
	modeDefault        = "full"
	sourceFlag         = "source"
	sourceUsage        = "Discovery source: artifacthub or github (full mode loads repos from -search-in JSON)"
	sourceDefault      = "artifacthub"
	topFlag            = "top"
	topUsage           = "Top N repositories to inspect from -search-in JSON in full mode"
	topDefault         = 400
	pageSizeFlag       = "page-size"
	pageSizeUsage      = "Results per API request/page (artifacthub max 60, github max 100)"
	pageSizeDefault    = 60
	searchTopFlag      = "search-top"
	searchTopUsage     = "In *-search-json modes, number of repos to include in output"
	searchTopDefault   = 1000
	searchOrderFlag    = "search-order"
	searchOrderUsage   = "In github-search-json mode, star ordering: desc or asc"
	searchOrderDefault = "desc"
	searchInFlag       = "search-in"
	searchInUsage      = "In full mode, input JSON file generated by github-search-json or artifacthub-search-json mode"
	searchInDefault    = "github_search.json"
	searchOutFlag      = "search-out"
	searchOutUsage     = "In *-search-json modes, output JSON file path"
	searchOutDefault   = "github_search.json"
	cloneDirFlag       = "clone-dir"
	cloneDirUsage      = "Directory where repositories are cloned"
	cloneDirDefault    = "cloned"
	catalogInFlag      = "catalog-in"
	catalogInUsage     = "In fixer mode, input catalog JSON (e.g. catalog_kept.json)"
	catalogInDefault   = "catalog_kept.json"
)
```

Replace the `Config` struct:

```go
// Config groups all runtime settings parsed from CLI flags.
type Config struct {
	Mode        string
	Source      string
	Top         int
	PageSize    int
	SearchTop   int
	SearchOrder string
	SearchIn    string
	SearchOut   string
	CloneDir    string
	CatalogIn   string
}
```

Replace the `Parse()` function:

```go
func Parse() Config {
	mode := flag.String(modeFlag, modeDefault, modeUsage)
	source := flag.String(sourceFlag, sourceDefault, sourceUsage)
	top := flag.Int(topFlag, topDefault, topUsage)
	pageSize := flag.Int(pageSizeFlag, pageSizeDefault, pageSizeUsage)
	searchTop := flag.Int(searchTopFlag, searchTopDefault, searchTopUsage)
	searchOrder := flag.String(searchOrderFlag, searchOrderDefault, searchOrderUsage)
	searchIn := flag.String(searchInFlag, searchInDefault, searchInUsage)
	searchOut := flag.String(searchOutFlag, searchOutDefault, searchOutUsage)
	cloneDir := flag.String(cloneDirFlag, cloneDirDefault, cloneDirUsage)
	catalogIn := flag.String(catalogInFlag, catalogInDefault, catalogInUsage)
	flag.Parse()

	return Config{
		Mode:        *mode,
		Source:      *source,
		Top:         *top,
		PageSize:    *pageSize,
		SearchTop:   *searchTop,
		SearchOrder: *searchOrder,
		SearchIn:    *searchIn,
		SearchOut:   *searchOut,
		CloneDir:    *cloneDir,
		CatalogIn:   *catalogIn,
	}
}
```

- [ ] **Step 2: Verify compilation**

```
cd helm_fetcher
go build ./...
```

Expected: no errors.

- [ ] **Step 3: Commit**

```
git add helm_fetcher/config/config.go
git commit -m "feat(helm_fetcher/config): add -catalog-in flag for fixer mode"
```

---

### Task 4: Implement parseHelmCLIError and fixRun

**Files:**
- Create: `helm_fetcher/fixer_mode.go` (partial — just the fix loop, no orchestration yet)
- Create: `helm_fetcher/fixer_mode_test.go`

- [ ] **Step 1: Write failing tests**

Create `helm_fetcher/fixer_mode_test.go`:

```go
package main

import (
	"testing"

	"github.com/MabsIPCA/helm-tests/helm_fetcher/model"
)

func TestParseHelmCLIError_NilPointer(t *testing.T) {
	errMsg := `Error: template: mychart/templates/svc.yaml:12:8: executing "mychart/templates/svc.yaml" at <.Values.service.type>: nil pointer evaluating interface {}.type`
	kind, path, val, ok := parseHelmCLIError(errMsg)
	if !ok {
		t.Fatal("expected ok=true")
	}
	if kind != "nil_pointer" {
		t.Errorf("kind: got %q, want %q", kind, "nil_pointer")
	}
	if path != "service.type" {
		t.Errorf("path: got %q, want %q", path, "service.type")
	}
	if val != "" {
		t.Errorf("val: got %q, want empty string", val)
	}
}

func TestParseHelmCLIError_RequiredValue(t *testing.T) {
	errMsg := `Error: template: mychart/templates/dep.yaml:5:15: executing "mychart/templates/dep.yaml" at <required .Values.db.host "db.host is required">: error calling required: db.host is required`
	kind, path, val, ok := parseHelmCLIError(errMsg)
	if !ok {
		t.Fatal("expected ok=true")
	}
	if kind != "required_value" {
		t.Errorf("kind: got %q, want %q", kind, "required_value")
	}
	if path != "db.host" {
		t.Errorf("path: got %q, want %q", path, "db.host")
	}
	if val != "kics-placeholder" {
		t.Errorf("val: got %q, want %q", val, "kics-placeholder")
	}
}

func TestParseHelmCLIError_Unfixable(t *testing.T) {
	errMsg := `Error: parse error at (mychart/templates/bad.yaml:3): unexpected "}" in command`
	_, _, _, ok := parseHelmCLIError(errMsg)
	if ok {
		t.Error("expected ok=false for parse error")
	}
}

func TestParseHelmCLIError_StripsCLIPrefix(t *testing.T) {
	// "Error: " prefix must be stripped before matching
	withPrefix := `Error: template: c/t.yaml:1:1: executing "" at <.Values.foo>: nil pointer evaluating interface {}.foo`
	withoutPrefix := `template: c/t.yaml:1:1: executing "" at <.Values.foo>: nil pointer evaluating interface {}.foo`
	_, path1, _, ok1 := parseHelmCLIError(withPrefix)
	_, path2, _, ok2 := parseHelmCLIError(withoutPrefix)
	if !ok1 || !ok2 {
		t.Fatal("both forms must parse ok")
	}
	if path1 != path2 {
		t.Errorf("path mismatch: %q vs %q", path1, path2)
	}
}

func TestParseHelmCLIError_PipeFilter(t *testing.T) {
	// Pipe-filter expressions like <.Values.foo | quote> must NOT match nilPtrRe
	errMsg := `Error: template: c/t.yaml:3:5: executing "" at <.Values.foo | quote>: nil pointer evaluating interface {}.foo`
	_, _, _, ok := parseHelmCLIError(errMsg)
	// The pipe-filter form doesn't match nilPtrRe (stops at space/pipe) so falls through to unfixable.
	// It may or may not match depending on the exact error. This test just verifies no panic.
	_ = ok
}

func TestFixRun_AlreadySucceeded(t *testing.T) {
	orig := model.RunResult{Success: true}
	result := fixRun("/nonexistent", orig)
	if !result.Resolved {
		t.Error("already-succeeded run must return Resolved=true")
	}
	if len(result.FixChain) != 0 {
		t.Errorf("FixChain must be empty for already-succeeded run, got %d steps", len(result.FixChain))
	}
	if result.StopReason != "" {
		t.Errorf("StopReason must be empty, got %q", result.StopReason)
	}
}

func TestFixRun_UnfixableImmediately(t *testing.T) {
	// A chart path that doesn't exist → helm returns an exec/path error which is unfixable.
	orig := model.RunResult{
		Success:      false,
		ErrorMessage: "Error: stat /nonexistent: no such file or directory",
	}
	result := fixRun("/nonexistent", orig)
	// Should stop immediately: unfixable_error_kind or loop_detected (depending on helm output).
	// Either way it must not resolve.
	if result.Resolved {
		t.Error("expected Resolved=false for non-existent chart")
	}
}
```

- [ ] **Step 2: Run tests to verify they fail**

```
cd helm_fetcher
go test . -v -run "TestParseHelmCLIError|TestFixRun"
```

Expected: compile error — `parseHelmCLIError` and `fixRun` undefined.

- [ ] **Step 3: Create fixer_mode.go with parseHelmCLIError and fixRun**

Create `helm_fetcher/fixer_mode.go`:

```go
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
```

- [ ] **Step 4: Run tests to verify they pass**

```
cd helm_fetcher
go test . -v -run "TestParseHelmCLIError|TestFixRun"
```

Expected: `PASS` for all 7 tests (`TestFixRun_UnfixableImmediately` may vary by environment but must not panic or resolve).

- [ ] **Step 5: Commit**

```
git add helm_fetcher/fixer_mode.go helm_fetcher/fixer_mode_test.go
git commit -m "feat(helm_fetcher): add parseHelmCLIError and fixRun"
```

---

### Task 5: Implement runFixerMode, writeFixedJSON, writeFixerReport

**Files:**
- Modify: `helm_fetcher/fixer_mode.go` (replace placeholder bodies)

- [ ] **Step 1: Replace runFixerMode, writeFixedJSON, writeFixerReport**

Replace the three stub functions and the unused-import block at the bottom of `helm_fetcher/fixer_mode.go` with the full implementations below. Also remove the `_ = ...` suppression block since all imports will be used.

The complete replacement for everything from `// runFixerMode is implemented in Task 5.` to end of file:

```go
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

	for i, repo := range repos {
		destDir := resolveDestDir(repo.RepoURL, cloneDir)

		log.Info().
			Int("repo_index", i+1).
			Int("total", len(repos)).
			Str("repo", repo.RepoName).
			Msg("Processing repo")

		if cloneErr := git.CloneRepo(repo.RepoURL, destDir); cloneErr != nil {
			log.Warn().Err(cloneErr).Str("repo", repo.RepoName).Msg("Clone failed – skipping")
			continue
		}

		freshCharts := helm.FindCharts(destDir)
		if len(freshCharts) == 0 {
			log.Warn().Str("repo", repo.RepoName).Msg("No charts found – skipping")
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
				} else {
					k := run.Fixed.FixChain[0].Kind
					if _, known := stats[k]; !known {
						stats["other"].before++
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
```

- [ ] **Step 2: Verify the file compiles cleanly**

```
cd helm_fetcher
go build ./...
```

Expected: no errors.

- [ ] **Step 3: Run all tests**

```
cd helm_fetcher
go test ./... -v -run "TestParseHelmCLIError|TestFixRun|TestFix|TestRunHelmTemplateWithSets"
```

Expected: all tests pass.

- [ ] **Step 4: Commit**

```
git add helm_fetcher/fixer_mode.go
git commit -m "feat(helm_fetcher): implement runFixerMode, writeFixedJSON, writeFixerReport"
```

---

### Task 6: Wire into main.go and Makefile, E2E Verify

**Files:**
- Modify: `helm_fetcher/main.go`
- Modify: `helm_fetcher/Makefile`

- [ ] **Step 1: Add fixer case to main.go**

In `helm_fetcher/main.go`, add the fixer mode case before the `if selectedMode != "full"` check (after line 70). Insert:

```go
	if selectedMode == "fixer" {
		runFixerMode(cfg.CatalogIn, cfg.CloneDir)
		return
	}
```

The block between the `artifacthub-search-json` check and the `full` validation should now look like:

```go
	if selectedMode == "artifacthub-search-json" {
		searchOnlyErr := runArtifactHubSearchJSONMode(cfg.PageSize, cfg.SearchTop, cfg.SearchOut)
		if searchOnlyErr != nil {
			log.Fatal().Err(searchOnlyErr).Msg("artifacthub-search-json mode failed")
		}
		return
	}
	if selectedMode == "fixer" {
		runFixerMode(cfg.CatalogIn, cfg.CloneDir)
		return
	}
	if selectedMode != "full" {
		log.Fatal().Str("mode", cfg.Mode).Msg("Invalid mode. Use 'full', 'github-search-json', 'artifacthub-search-json', or 'fixer'")
	}
```

Also update the error message on the `!= "full"` check to include `fixer`:

```go
	if selectedMode != "full" {
		log.Fatal().Str("mode", cfg.Mode).Msg("Invalid mode. Use 'full', 'github-search-json', 'artifacthub-search-json', or 'fixer'")
	}
```

- [ ] **Step 2: Add run-fixer target to Makefile**

In `helm_fetcher/Makefile`, add after the `CLONE_DIR` variable definition and after the `.PHONY` line, add `run-fixer` to `.PHONY` and insert the target.

Add `CATALOG_IN` variable after line 35 (`CLONE_DIR ?= ...`):

```makefile
# In fixer mode, the catalog JSON to read (produced by a previous 'run' invocation)
CATALOG_IN ?= catalog_kept.json
```

Add `run-fixer` to the `.PHONY` line so it reads:

```makefile
.PHONY: run run-github-search-json run-artifacthub-search-json run-fixer plot-problems plot-popularity plot-problems-popularity plot-all clean-cloned clean-output clean help
```

Add the target after the `run-artifacthub-search-json` block:

```makefile
## run-fixer: read catalog_kept.json, clone repos, run fix loop, write catalog_fixed.json
run-fixer:
	go run . \
		-mode=fixer \
		-catalog-in=$(CATALOG_IN) \
		-clone-dir=$(CLONE_DIR)
```

Also add an entry to the `help` target's echo block:

```makefile
	@echo     run-fixer              - read catalog and apply fix loop (writes catalog_fixed.json)
```

And add `CATALOG_IN` to the Variables section:

```makefile
	@echo     CATALOG_IN   (default: catalog_kept.json) input catalog for run-fixer mode
```

- [ ] **Step 3: Build to confirm wiring**

```
cd helm_fetcher
go build ./...
```

Expected: no errors.

- [ ] **Step 4: Smoke test with --help**

```
cd helm_fetcher
go run . -help 2>&1 | head -30
```

Expected: `-catalog-in` flag is listed in the output.

- [ ] **Step 5: E2E dry-run against catalog_kept.json**

Run against the real catalog with the existing cloned directory:

```
cd helm_fetcher
go run . -mode=fixer -catalog-in=backup/run_002_274/catalog_kept.json -clone-dir=cloned 2>&1 | tail -20
```

Expected output ends with something like:

```
Catalog Fixer complete.
  Repos processed: N
  Failing runs:    N
  Resolved:        N
  Resolution rate: N%

Output files:
  catalog_fixed.json
  catalog_fixer_report.md
```

Verify the output files exist:

```
dir catalog_fixed.json catalog_fixer_report.md
```

Expected: both files present with non-zero size.

Spot-check one resolved entry in catalog_fixed.json:

```
findstr /C:"\"resolved\": true" catalog_fixed.json | find "" /c
```

Expected: a non-zero count.

- [ ] **Step 6: Run full test suite**

```
cd helm_fetcher
go test ./... -count=1
```

Expected: `ok` for all packages, no failures.

- [ ] **Step 7: Commit**

```
git add helm_fetcher/main.go helm_fetcher/Makefile
git commit -m "feat(helm_fetcher): wire fixer mode into main.go and Makefile"
```

---

## Self-Review

**Spec coverage check:**

| Spec requirement | Task |
|---|---|
| `-catalog-in` flag, default `catalog_kept.json` | Task 3 |
| `RunHelmTemplateWithSets` function | Task 2 |
| `FixStep`, `FixedRunResult`, `RunResultWithFix`, `ChartSummaryFixed`, `RepoResultFixed` model types | Task 1 |
| `parseHelmCLIError` with nilPtrRe / requiredRe | Task 4 |
| Fix loop: max 10 iterations, seenErrs guard, stop reasons | Task 4 |
| Already-successful runs return immediately with `Resolved: true, FixChain: []` | Task 4 |
| `runFixerMode`: read catalog → clone → FindCharts fresh → fixRun → write JSON → write report | Task 5 |
| `catalog_fixed.json` written; fatal on failure | Task 5 |
| `catalog_fixer_report.md` written; log error and continue on failure | Task 5 |
| Markdown report: Summary table, By Error Kind table, Stop Reasons table | Task 5 |
| `run-fixer` Makefile target with `CATALOG_IN` variable | Task 6 |
| `fixer` case wired into `main.go` mode switch | Task 6 |
| Clone path resolution via `resolveDestDir` (same algorithm as main.go) | Task 4/5 |
| Values files remapped if cloneDir changed | Task 5 |

**Placeholder scan:** No TBDs. All code blocks are complete.

**Type consistency check:**
- `model.FixStep` used in `fixer_mode.go` chain ✓
- `model.FixedRunResult` returned by `fixRun`, stored in `RunResultWithFix.Fixed` ✓
- `model.RunResultWithFix` collected in `ChartSummaryFixed.Runs` ✓
- `model.RepoResultFixed` collected in `fixedRepos` slice, passed to `writeFixedJSON`/`writeFixerReport` ✓
- `helm.RunHelmTemplateWithSets` called in `fixRun` ✓
- `git.CloneRepo`, `helm.FindCharts` called in `runFixerMode` ✓
