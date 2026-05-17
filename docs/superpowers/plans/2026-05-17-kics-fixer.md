# KICS Fixer Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Add a retry-fix loop to `kics-render-mock` that satisfies `nil_pointer` and `required_value` render errors by injecting missing values, writing `kics-fixed-output.json` per test dir with before/after results.

**Architecture:** One new file (`fixer.go`) in the existing `kics-render-mock` package handles error parsing, value patching, and the retry loop. New types in `output.go` describe the fixed output shape. `main.go` calls the fixer automatically after each base run and prints a before/after summary line.

**Tech Stack:** Go 1.21+, `helm.sh/helm/v3` (already vendored), `regexp`, `encoding/json`

---

### Task 1: Add fixed output types and writeFixedOutput

**Files:**
- Modify: `kics-render-mock/output.go` — append new types and writer
- Modify: `kics-render-mock/output_test.go` — append two new tests

- [ ] **Step 1: Write the failing tests**

Append to `kics-render-mock/output_test.go` (after the existing `TestRenderOutputMarshal`):

```go
func TestFixedRenderEntryMarshal_EmbeddedFieldsInlined(t *testing.T) {
	errStr := "nil pointer"
	entry := FixedRenderEntry{
		RenderEntry: RenderEntry{
			Toggle:   "nil-test",
			DataType: "nilMap",
			ValuesOptions: ValuesOptionsSummary{
				Values:     []string{"testNilMapAccess=true"},
				ValueFiles: []string{},
			},
			Standard: RunResult{Error: &errStr, SplitManifests: []SplitManifestEntry{}},
			Debug:    DebugRunResult{Error: &errStr, DebugLogs: []string{}, SplitManifests: []SplitManifestEntry{}},
		},
		Resolved:   false,
		StopReason: "unfixable_error_kind",
		FixChain:   []FixStep{},
	}
	data, err := json.Marshal(entry)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"toggle", "dataType", "standard", "debug", "resolved", "stopReason", "fixChain"} {
		if _, ok := m[key]; !ok {
			t.Errorf("key %q missing from JSON", key)
		}
	}
}

func TestWriteFixedOutput_CreatesFile(t *testing.T) {
	dir := t.TempDir()
	out := FixedRenderOutput{
		Suite:      "render_problems",
		TestNumber: 1,
		TestName:   "Test",
		ChartPath:  dir,
		Renders:    []FixedRenderEntry{},
	}
	if err := writeFixedOutput(dir, out); err != nil {
		t.Fatalf("writeFixedOutput: %v", err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "kics-fixed-output.json"))
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	var got FixedRenderOutput
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if got.Suite != "render_problems" {
		t.Errorf("Suite: got %q", got.Suite)
	}
	if got.Renders == nil {
		t.Error("Renders must not be nil")
	}
}
```

Add the missing imports to `output_test.go` — it already has `encoding/json` and `testing`; add `"os"` and `"path/filepath"`.

- [ ] **Step 2: Run tests to verify they fail**

```
cd kics-render-mock
go test -run "TestFixedRenderEntry|TestWriteFixedOutput" -v
```

Expected: `FAIL` — `FixedRenderEntry`, `FixStep`, `FixedRenderOutput`, `writeFixedOutput` are not defined yet.

- [ ] **Step 3: Add types and writer to output.go**

Append to the bottom of `kics-render-mock/output.go`:

```go
type FixStep struct {
	Attempt       int    `json:"attempt"`
	ErrorSeen     string `json:"errorSeen"`
	Kind          string `json:"kind"`
	ValuePath     string `json:"valuePath"`
	ValueInjected string `json:"valueInjected"`
}

type FixedRenderEntry struct {
	RenderEntry
	Resolved    bool       `json:"resolved"`
	StopReason  string     `json:"stopReason"`
	FixChain    []FixStep  `json:"fixChain"`
	FixedResult *RunResult `json:"fixedResult,omitempty"`
}

type FixedRenderOutput struct {
	Suite      string             `json:"suite"`
	TestNumber int                `json:"testNumber"`
	TestName   string             `json:"testName"`
	ChartPath  string             `json:"chartPath"`
	Renders    []FixedRenderEntry `json:"renders"`
}

func writeFixedOutput(testDir string, out FixedRenderOutput) error {
	data, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(testDir, "kics-fixed-output.json"), data, 0644)
}
```

- [ ] **Step 4: Run tests to verify they pass**

```
cd kics-render-mock
go test -run "TestFixedRenderEntry|TestWriteFixedOutput" -v
```

Expected: both `PASS`.

- [ ] **Step 5: Run full test suite to check for regressions**

```
cd kics-render-mock
go test ./...
```

Expected: all existing tests still `PASS`.

- [ ] **Step 6: Commit**

```
git add kics-render-mock/output.go kics-render-mock/output_test.go
git commit -m "feat(fixer): add FixedRenderOutput types and writeFixedOutput"
```

---

### Task 2: Add error parser and value patcher

**Files:**
- Create: `kics-render-mock/fixer.go` — `parseError`, `applyPatch`, constants, regexes
- Create: `kics-render-mock/fixer_test.go` — 7 unit tests

- [ ] **Step 1: Write the failing tests**

Create `kics-render-mock/fixer_test.go`:

```go
package main

import (
	"testing"

	"helm.sh/helm/v3/pkg/cli/values"
)

func TestParseError_NilPointer(t *testing.T) {
	errStr := `template: mychart/templates/deploy.yaml:12:8: executing "mychart/templates/deploy.yaml" at <.Values.ingress.hosts>: nil pointer evaluating interface {}.hosts`
	fix, ok := parseError(errStr)
	if !ok {
		t.Fatal("expected parseable, got false")
	}
	if fix.kind != kindNilPointer {
		t.Errorf("kind: got %q, want %q", fix.kind, kindNilPointer)
	}
	if fix.path != "ingress.hosts" {
		t.Errorf("path: got %q, want %q", fix.path, "ingress.hosts")
	}
	if fix.value != "" {
		t.Errorf("value: got %q, want empty string", fix.value)
	}
}

func TestParseError_RequiredValue(t *testing.T) {
	errStr := `template: mychart/templates/deploy.yaml:5:15: executing "mychart/templates/deploy.yaml" at <required .Values.db.host "db.host is required">: error calling required: db.host is required`
	fix, ok := parseError(errStr)
	if !ok {
		t.Fatal("expected parseable, got false")
	}
	if fix.kind != kindRequiredValue {
		t.Errorf("kind: got %q, want %q", fix.kind, kindRequiredValue)
	}
	if fix.path != "db.host" {
		t.Errorf("path: got %q, want %q", fix.path, "db.host")
	}
	if fix.value != "kics-placeholder" {
		t.Errorf("value: got %q, want %q", fix.value, "kics-placeholder")
	}
}

func TestParseError_Unfixable(t *testing.T) {
	errStr := `template: mychart/templates/deploy.yaml:1:1: executing "..." at <.Values.foo>: error calling fail: chart cannot be installed directly`
	_, ok := parseError(errStr)
	if ok {
		t.Fatal("expected unfixable, got parseable")
	}
}

func TestParseError_NilPointerWithoutAtClause(t *testing.T) {
	errStr := `template: mychart/templates/deploy.yaml:12: nil pointer evaluating something that has no at clause`
	_, ok := parseError(errStr)
	if ok {
		t.Fatal("expected unfixable without at <.Values.x> clause, got parseable")
	}
}

func TestApplyPatch_AppendsToOriginalValues(t *testing.T) {
	orig := &values.Options{
		Values:     []string{"foo=bar"},
		ValueFiles: []string{"/some/values.yaml"},
	}
	patch := map[string]string{"db.host": "kics-placeholder"}
	patched := applyPatch(orig, patch)

	if len(patched.ValueFiles) != 1 || patched.ValueFiles[0] != "/some/values.yaml" {
		t.Error("ValueFiles not preserved")
	}
	if len(patched.Values) != 2 {
		t.Fatalf("Values len: got %d, want 2", len(patched.Values))
	}
	if patched.Values[0] != "foo=bar" {
		t.Errorf("original value not at index 0: got %q", patched.Values[0])
	}
	if patched.Values[1] != "db.host=kics-placeholder" {
		t.Errorf("patch not appended correctly: got %q", patched.Values[1])
	}
}

func TestApplyPatch_DoesNotMutateOriginal(t *testing.T) {
	orig := &values.Options{Values: []string{"foo=bar"}}
	_ = applyPatch(orig, map[string]string{"extra": "val"})
	if len(orig.Values) != 1 {
		t.Errorf("original Values was mutated: len=%d", len(orig.Values))
	}
}

func TestApplyPatch_EmptyPatch(t *testing.T) {
	orig := &values.Options{Values: []string{"foo=bar"}}
	patched := applyPatch(orig, map[string]string{})
	if len(patched.Values) != 1 {
		t.Errorf("Values len: got %d, want 1", len(patched.Values))
	}
}
```

- [ ] **Step 2: Run tests to verify they fail**

```
cd kics-render-mock
go test -run "TestParseError|TestApplyPatch" -v
```

Expected: `FAIL` — `parseError`, `applyPatch`, `kindNilPointer`, `kindRequiredValue` not defined yet.

- [ ] **Step 3: Create fixer.go with parseError and applyPatch**

Create `kics-render-mock/fixer.go`:

```go
package main

import (
	"fmt"
	"regexp"
	"strings"

	"helm.sh/helm/v3/pkg/cli/values"
)

const (
	kindNilPointer    = "nil_pointer"
	kindRequiredValue = "required_value"
	maxFixIterations  = 10
)

var (
	nilPtrRe   = regexp.MustCompile(`at <(\.Values\.[^>]+)>`)
	requiredRe = regexp.MustCompile(`required (\.Values\.[^ "]+)`)
)

type parsedFix struct {
	kind  string
	path  string
	value string
}

func parseError(errStr string) (parsedFix, bool) {
	if strings.Contains(errStr, "nil pointer") {
		if m := nilPtrRe.FindStringSubmatch(errStr); m != nil {
			path := strings.TrimPrefix(m[1], ".Values.")
			return parsedFix{kind: kindNilPointer, path: path, value: ""}, true
		}
	}
	if strings.Contains(errStr, "error calling required") {
		if m := requiredRe.FindStringSubmatch(errStr); m != nil {
			path := strings.TrimPrefix(m[1], ".Values.")
			return parsedFix{kind: kindRequiredValue, path: path, value: "kics-placeholder"}, true
		}
	}
	return parsedFix{}, false
}

func applyPatch(orig *values.Options, patch map[string]string) *values.Options {
	extra := make([]string, 0, len(patch))
	for k, v := range patch {
		extra = append(extra, fmt.Sprintf("%s=%s", k, v))
	}
	return &values.Options{
		ValueFiles:   orig.ValueFiles,
		Values:       append(append([]string{}, orig.Values...), extra...),
		StringValues: orig.StringValues,
		FileValues:   orig.FileValues,
	}
}
```

- [ ] **Step 4: Run tests to verify they pass**

```
cd kics-render-mock
go test -run "TestParseError|TestApplyPatch" -v
```

Expected: all 7 `PASS`.

- [ ] **Step 5: Run full test suite**

```
cd kics-render-mock
go test ./...
```

Expected: all `PASS`.

- [ ] **Step 6: Commit**

```
git add kics-render-mock/fixer.go kics-render-mock/fixer_test.go
git commit -m "feat(fixer): add parseError and applyPatch with unit tests"
```

---

### Task 3: Add fix loop (fixInvocation + runFixDir)

**Files:**
- Modify: `kics-render-mock/fixer.go` — append `fixInvocation` and `runFixDir`

No unit tests for these functions: they call `runOnce` which requires a real chart on disk. Correctness is verified end-to-end in Task 4.

- [ ] **Step 1: Append fixInvocation and runFixDir to fixer.go**

Append to the bottom of `kics-render-mock/fixer.go`:

```go
func fixInvocation(chartPath string, inv invocation) FixedRenderEntry {
	entry := FixedRenderEntry{FixChain: []FixStep{}}

	patch := map[string]string{}
	seenErrs := map[string]bool{}

	for attempt := 1; attempt <= maxFixIterations; attempt++ {
		patchedOpts := applyPatch(inv.valOpts, patch)
		res := runOnce(chartPath, patchedOpts, false)

		if res.err == nil {
			entry.Resolved = true
			fixed := toStandardResult(res)
			entry.FixedResult = &fixed
			return entry
		}

		errStr := res.err.Error()
		if seenErrs[errStr] {
			entry.StopReason = "loop_detected"
			return entry
		}
		seenErrs[errStr] = true

		fix, ok := parseError(errStr)
		if !ok {
			entry.StopReason = "unfixable_error_kind"
			return entry
		}

		entry.FixChain = append(entry.FixChain, FixStep{
			Attempt:       attempt,
			ErrorSeen:     errStr,
			Kind:          fix.kind,
			ValuePath:     fix.path,
			ValueInjected: fix.value,
		})
		patch[fix.path] = fix.value
	}

	entry.StopReason = "max_iterations"
	return entry
}

func runFixDir(testDir, suite string, base RenderOutput) (FixedRenderOutput, error) {
	cfg, err := loadConfig(testDir)
	if err != nil {
		return FixedRenderOutput{}, err
	}
	invocations := buildInvocations(cfg, testDir)

	out := FixedRenderOutput{
		Suite:      suite,
		TestNumber: base.TestNumber,
		TestName:   base.TestName,
		ChartPath:  base.ChartPath,
		Renders:    make([]FixedRenderEntry, 0, len(base.Renders)),
	}

	for i, inv := range invocations {
		baseEntry := base.Renders[i]
		fixed := FixedRenderEntry{
			RenderEntry: baseEntry,
			FixChain:    []FixStep{},
		}

		if baseEntry.Standard.Error == nil {
			fixed.Resolved = true
			out.Renders = append(out.Renders, fixed)
			continue
		}

		loopResult := fixInvocation(testDir, inv)
		fixed.Resolved = loopResult.Resolved
		fixed.StopReason = loopResult.StopReason
		fixed.FixChain = loopResult.FixChain
		fixed.FixedResult = loopResult.FixedResult
		out.Renders = append(out.Renders, fixed)
	}

	return out, nil
}
```

- [ ] **Step 2: Verify it compiles**

```
cd kics-render-mock
go build ./...
```

Expected: no errors.

- [ ] **Step 3: Run full test suite**

```
cd kics-render-mock
go test ./...
```

Expected: all `PASS`.

- [ ] **Step 4: Commit**

```
git add kics-render-mock/fixer.go
git commit -m "feat(fixer): add fixInvocation retry loop and runFixDir"
```

---

### Task 4: Wire fixer into main.go and update summary

**Files:**
- Modify: `kics-render-mock/main.go` — call `runFixDir`, track resolved count, update summary

- [ ] **Step 1: Replace the suite loop body in main.go**

The current loop body (lines 86–114) handles `runTestDir` and `writeOutput`. Replace it with the version below that adds the fixer call and tracks `totalResolved`. The full updated `main()` function:

```go
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
		totalResolved := 0

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

			fixedOut, err := runFixDir(testDir, suite.name, out)
			if err != nil {
				fmt.Fprintf(os.Stderr, "[%s/%s] fixer error: %v\n", suite.name, dirName, err)
			} else {
				for _, r := range fixedOut.Renders {
					if r.Resolved && len(r.FixChain) > 0 {
						totalResolved++
					}
				}
				if err := writeFixedOutput(testDir, fixedOut); err != nil {
					fmt.Fprintf(os.Stderr, "[%s/%s] fixed write error: %v\n", suite.name, dirName, err)
				}
			}
		}

		resolved := ""
		if totalErrors > 0 {
			pct := totalResolved * 100 / totalErrors
			resolved = fmt.Sprintf(" → %d resolved (%d%%)", totalResolved, pct)
		}
		fmt.Printf("%-25s %2d dirs   %4d invocations   %4d errors%s\n",
			suite.name, len(suite.dirs), totalInvocations, totalErrors, resolved)
	}

	fmt.Println("JSON + fixed files written to each test directory.")
}
```

- [ ] **Step 2: Build to verify no compile errors**

```
cd kics-render-mock
go build ./...
```

Expected: no errors.

- [ ] **Step 3: Run unit tests**

```
cd kics-render-mock
go test ./...
```

Expected: all `PASS`.

- [ ] **Step 4: Do a real end-to-end run against the test suites**

From inside `kics-render-mock/`:

```
go run . -root ..
```

Expected output shape (numbers will vary):

```
=== KICS Render Mock ===
render_problems           21 dirs    126 invocations     89 errors → N resolved (X%)
override_tests_suite      10 dirs     42 invocations      0 errors
JSON + fixed files written to each test directory.
```

Verify that `kics-fixed-output.json` was created in at least one test directory:

```
dir ..\render_problems\test-02-nil-pointer\kics-fixed-output.json
```

Open one `kics-fixed-output.json` and confirm:
- Entries that had `Standard.Error != null` in `kics-render-output.json` now have a `fixChain` array (possibly non-empty)
- `resolved: true` entries that were fixed have a non-null `fixedResult` with `error: null`
- Entries that were already passing have `resolved: true`, `fixChain: []`, `fixedResult: null`

- [ ] **Step 5: Commit**

```
git add kics-render-mock/main.go
git commit -m "feat(fixer): wire runFixDir into main loop with before/after summary"
```
