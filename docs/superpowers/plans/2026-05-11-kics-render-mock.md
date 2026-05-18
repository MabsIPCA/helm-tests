# KICS Helm Render Mock — Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Build a standalone Go tool (`kics-render-mock`) that mimics the KICS internal Helm rendering pipeline, drives every config-defined toggle/case combination through it, and writes per-test-dir `kics-render-output.json` files capturing both a standard and a debug render run.

**Architecture:** Single `package main` with four source files. `pipeline.go` replicates the KICS Helm Go library pipeline (`newClient`, `setID`, `addID`, `splitManifestYAML`, `runOnce`). `runner.go` reads `config.json`, detects suite format, builds invocations, and runs each one. `output.go` owns all JSON types and the file writer. `main.go` discovers both suites and coordinates. TDD throughout for all pure functions.

**Tech Stack:** Go 1.25.3, `helm.sh/helm/v3` (same library KICS uses), `encoding/json`, standard `testing` package.

---

## File Map

| File | Responsibility |
|------|---------------|
| `kics-render-mock/go.mod` | Module definition, single dep: `helm.sh/helm/v3` |
| `kics-render-mock/output.go` | JSON types (`RenderOutput`, `RenderEntry`, `RunResult`, `DebugRunResult`, `SplitManifestEntry`, `ValuesOptionsSummary`) + `writeOutput` |
| `kics-render-mock/output_test.go` | JSON marshaling round-trip test |
| `kics-render-mock/pipeline.go` | `addID`, `setID`, `updateName`, `splitManifestYAML`, `runOnce`, internal `renderResult` |
| `kics-render-mock/pipeline_test.go` | Unit tests for `addID` and `splitManifestYAML` |
| `kics-render-mock/runner.go` | `testConfig`, `testFunction`, `testCase`, `invocation`, `loadConfig`, `buildInvocations`, `toStandardResult`, `toDebugResult`, `buildRenderEntry`, `runTestDir` |
| `kics-render-mock/runner_test.go` | Unit tests for `loadConfig` and `buildInvocations` |
| `kics-render-mock/main.go` | Suite dir lists, `main()` with `-root` flag, summary output |

Output files land inside each existing test directory:
```
render_problems/test-02-nil-pointer/kics-render-output.json
override_tests_suite/test-01-basic-data-types/kics-render-output.json
```

---

## Task 1: Scaffold the module

**Files:**
- Create: `kics-render-mock/go.mod`
- Create: `kics-render-mock/output.go` (stub)
- Create: `kics-render-mock/pipeline.go` (stub)
- Create: `kics-render-mock/runner.go` (stub)
- Create: `kics-render-mock/main.go` (stub)

- [ ] **Step 1: Create the directory and go.mod**

```
helm-tests/kics-render-mock/
```

`go.mod`:
```
module kics-render-mock

go 1.25.3
```

- [ ] **Step 2: Create stub source files**

`output.go`:
```go
package main
```

`pipeline.go`:
```go
package main
```

`runner.go`:
```go
package main
```

`main.go`:
```go
package main

func main() {}
```

- [ ] **Step 3: Fetch the Helm library**

Run from inside `kics-render-mock/`:
```
go get helm.sh/helm/v3@latest
go mod tidy
```

Expected: `go.sum` created, `go.mod` updated with `helm.sh/helm/v3 vX.Y.Z` and its indirect deps.

- [ ] **Step 4: Verify the module compiles**

```
go build ./...
```

Expected: no output, exit 0.

- [ ] **Step 5: Commit**

```bash
git add kics-render-mock/
git commit -m "feat: scaffold kics-render-mock module"
```

---

## Task 2: Define JSON output types and file writer (`output.go`)

**Files:**
- Write: `kics-render-mock/output.go`
- Create: `kics-render-mock/output_test.go`

- [ ] **Step 1: Write the failing test**

`output_test.go`:
```go
package main

import (
	"encoding/json"
	"testing"
)

func TestRenderOutputMarshal(t *testing.T) {
	errStr := "template error"
	manifest := "---\napiVersion: v1"
	out := RenderOutput{
		Suite:      "render_problems",
		TestNumber: 2,
		TestName:   "Nil Pointer",
		ChartPath:  "/some/path",
		Renders: []RenderEntry{
			{
				Toggle:   "nil-map-access",
				DataType: "nilMap",
				ValuesOptions: ValuesOptionsSummary{
					Values:     []string{"testNilMapAccess=true", "testDataType=nilMap"},
					ValueFiles: []string{},
				},
				Standard: RunResult{
					Error:          &errStr,
					RawManifest:    nil,
					PartialManifest: nil,
					SplitManifests: []SplitManifestEntry{},
				},
				Debug: DebugRunResult{
					Error:          &errStr,
					DebugLogs:      []string{"[debug] chart path: /some/path"},
					RawManifest:    nil,
					PartialManifest: &manifest,
					SplitManifests: []SplitManifestEntry{},
				},
			},
		},
	}

	data, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var roundTrip RenderOutput
	if err := json.Unmarshal(data, &roundTrip); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if roundTrip.Suite != "render_problems" {
		t.Errorf("Suite: got %q", roundTrip.Suite)
	}
	if len(roundTrip.Renders) != 1 {
		t.Fatalf("Renders len: got %d", len(roundTrip.Renders))
	}
	r := roundTrip.Renders[0]
	if r.Standard.Error == nil || *r.Standard.Error != errStr {
		t.Errorf("Standard.Error: got %v", r.Standard.Error)
	}
	if r.Debug.PartialManifest == nil || *r.Debug.PartialManifest != manifest {
		t.Errorf("Debug.PartialManifest: got %v", r.Debug.PartialManifest)
	}
	if len(r.Debug.DebugLogs) != 1 {
		t.Errorf("Debug.DebugLogs len: got %d", len(r.Debug.DebugLogs))
	}
}
```

- [ ] **Step 2: Run to verify it fails (types not defined)**

```
go test ./... 2>&1 | head -20
```

Expected: compilation error — `RenderOutput undefined`.

- [ ] **Step 3: Write the output types and file writer**

`output.go`:
```go
package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type ValuesOptionsSummary struct {
	Values       []string `json:"values"`
	ValueFiles   []string `json:"valueFiles"`
	StringValues []string `json:"stringValues,omitempty"`
}

type SplitManifestEntry struct {
	Path            string `json:"path"`
	Content         string `json:"content"`
	OriginalContent string `json:"originalContent"`
	SplitID         string `json:"splitID,omitempty"`
}

type RunResult struct {
	Error           *string              `json:"error"`
	RawManifest     *string              `json:"rawManifest"`
	PartialManifest *string              `json:"partialManifest"`
	SplitManifests  []SplitManifestEntry `json:"splitManifests"`
}

type DebugRunResult struct {
	Error           *string              `json:"error"`
	DebugLogs       []string             `json:"debugLogs"`
	RawManifest     *string              `json:"rawManifest"`
	PartialManifest *string              `json:"partialManifest"`
	SplitManifests  []SplitManifestEntry `json:"splitManifests"`
}

type RenderEntry struct {
	Toggle        string               `json:"toggle,omitempty"`
	DataType      string               `json:"dataType,omitempty"`
	CaseName      string               `json:"caseName,omitempty"`
	ValuesOptions ValuesOptionsSummary `json:"valuesOptions"`
	Standard      RunResult            `json:"standard"`
	Debug         DebugRunResult       `json:"debug"`
}

type RenderOutput struct {
	Suite      string        `json:"suite"`
	TestNumber int           `json:"testNumber"`
	TestName   string        `json:"testName"`
	ChartPath  string        `json:"chartPath"`
	Renders    []RenderEntry `json:"renders"`
}

func writeOutput(testDir string, out RenderOutput) error {
	data, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(testDir, "kics-render-output.json"), data, 0644)
}
```

- [ ] **Step 4: Run the test to verify it passes**

```
go test ./... -run TestRenderOutputMarshal -v
```

Expected:
```
--- PASS: TestRenderOutputMarshal (0.00s)
PASS
```

- [ ] **Step 5: Commit**

```bash
git add kics-render-mock/output.go kics-render-mock/output_test.go
git commit -m "feat: define JSON output types and file writer"
```

---

## Task 3: Implement ID injection (`pipeline.go`)

**Files:**
- Write: `kics-render-mock/pipeline.go` (partial — `addID`, `setID`, `updateName`, constants)
- Create: `kics-render-mock/pipeline_test.go`

- [ ] **Step 1: Write the failing test**

`pipeline_test.go`:
```go
package main

import (
	"strings"
	"testing"

	"helm.sh/helm/v3/pkg/chart"
)

func TestAddID_InjectsMarkerBeforeAPIVersion(t *testing.T) {
	f := &chart.File{
		Name: "templates/test.yaml",
		Data: []byte("apiVersion: v1\nkind: ConfigMap\nmetadata:\n  name: test"),
	}
	result := addID(f)
	lines := strings.Split(string(result.Data), "\n")

	found := false
	for i, l := range lines {
		if strings.Contains(l, "# KICS_HELM_ID_0:") {
			found = true
			if i+1 >= len(lines) || !strings.Contains(lines[i+1], "apiVersion:") {
				t.Errorf("KICS_HELM_ID_0: must immediately precede apiVersion:, got next line: %q", lines[i+1])
			}
			break
		}
	}
	if !found {
		t.Errorf("KICS_HELM_ID_0: marker not found in:\n%s", string(result.Data))
	}
}

func TestAddID_MultipleAPIVersions(t *testing.T) {
	f := &chart.File{
		Name: "templates/multi.yaml",
		Data: []byte("apiVersion: v1\nkind: ConfigMap\n---\napiVersion: apps/v1\nkind: Deployment"),
	}
	result := addID(f)
	content := string(result.Data)
	if !strings.Contains(content, "# KICS_HELM_ID_0:") {
		t.Error("missing KICS_HELM_ID_0:")
	}
	// Second apiVersion: shifts by 2 due to first insertion
	if !strings.Contains(content, kicsHelmID) {
		t.Error("missing additional KICS_HELM_ID marker")
	}
}
```

- [ ] **Step 2: Run to verify it fails**

```
go test ./... -run TestAddID 2>&1 | head -10
```

Expected: compilation error — `addID undefined`, `kicsHelmID undefined`.

- [ ] **Step 3: Implement `addID`, `setID`, `updateName`**

`pipeline.go` (replace stub):
```go
package main

import (
	"fmt"
	"io"
	"log"
	"path/filepath"
	"strings"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/chartutil"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/cli/values"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/release"
)

const kicsHelmID = "# KICS_HELM_ID_"

var helmSettings = cli.New()

// renderResult is the raw output of one runOnce call.
type renderResult struct {
	rel  *release.Release
	err  error
	logs []string
}

// addID mirrors KICS's addID: injects a # KICS_HELM_ID_<line>: comment before
// each "apiVersion:" line in a template file. Mutates file.Data in place.
func addID(file *chart.File) *chart.File {
	split := strings.Split(string(file.Data), "\n")
	for i := 0; i < len(split); i++ {
		if strings.Contains(split[i], "apiVersion:") {
			split = append(split, "")
			copy(split[i+1:], split[i:])
			split[i] = fmt.Sprintf("%s%d:", kicsHelmID, i)
			i++
		}
	}
	file.Data = []byte(strings.Join(split, "\n"))
	return file
}

// setID mirrors KICS's setID: applies addID to all templates in the chart and
// its dependencies.
func setID(c *chart.Chart) *chart.Chart {
	for _, t := range c.Templates {
		addID(t)
	}
	for _, dep := range c.Dependencies() {
		setID(dep)
	}
	return c
}

// updateName mirrors KICS's updateName: prefixes all template file names with
// the chart name hierarchy and collects them into a flat slice.
func updateName(files []*chart.File, c *chart.Chart, name string) []*chart.File {
	if name != c.Name() {
		name = filepath.Join(name, c.Name())
	}
	for _, t := range c.Templates {
		t.Name = filepath.Join(name, t.Name)
	}
	files = append(files, c.Templates...)
	for _, dep := range c.Dependencies() {
		files = updateName(files, dep, filepath.Join(name, "charts"))
	}
	return files
}
```

Note: `runOnce` and `splitManifestYAML` are added in Tasks 4 and 5.

- [ ] **Step 4: Run the tests**

```
go test ./... -run TestAddID -v
```

Expected:
```
--- PASS: TestAddID_InjectsMarkerBeforeAPIVersion (0.00s)
--- PASS: TestAddID_MultipleAPIVersions (0.00s)
PASS
```

- [ ] **Step 5: Commit**

```bash
git add kics-render-mock/pipeline.go kics-render-mock/pipeline_test.go
git commit -m "feat: implement KICS-faithful ID injection pipeline"
```

---

## Task 4: Implement the render pipeline — `runOnce` (`pipeline.go`)

**Files:**
- Modify: `kics-render-mock/pipeline.go` (add `runOnce`)

No unit test possible here — requires a real chart on disk. Verified in Task 9.

- [ ] **Step 1: Add `runOnce` to `pipeline.go`**

Append to `pipeline.go` (after the existing functions):
```go
// runOnce runs the KICS-faithful Helm pipeline once for the given chart path
// and values options. When debugMode is true, cfg.Log is wired to collect log
// lines and helmSettings.Debug is set to true for the duration of the call.
//
// Even when err != nil, rel may be non-nil (partial render before failure).
func runOnce(chartPath string, valOpts *values.Options, debugMode bool) renderResult {
	var logs []string

	cfg := new(action.Configuration)
	if debugMode {
		cfg.Log = func(format string, v ...interface{}) {
			logs = append(logs, fmt.Sprintf(format, v...))
		}
		helmSettings.Debug = true
		defer func() { helmSettings.Debug = false }()
	}

	client := action.NewInstall(cfg)
	client.DryRun = true
	client.ReleaseName = "kics-helm"
	client.Replace = true
	client.ClientOnly = true
	client.APIVersions = chartutil.VersionSet([]string{})
	client.IncludeCRDs = false
	client.Namespace = "kics-namespace"

	// Suppress Helm's internal std logger output.
	log.SetOutput(io.Discard)
	defer log.SetOutput(nil)

	p := getter.All(helmSettings)
	vals, err := valOpts.MergeValues(p)
	if err != nil {
		return renderResult{err: err, logs: logs}
	}

	chartReq, err := loader.Load(chartPath)
	if err != nil {
		return renderResult{err: err, logs: logs}
	}

	chartReq = setID(chartReq)

	rel, runErr := client.Run(chartReq, vals)
	return renderResult{rel: rel, err: runErr, logs: logs}
}
```

- [ ] **Step 2: Verify the module still compiles**

```
go build ./...
```

Expected: no output, exit 0.

- [ ] **Step 3: Commit**

```bash
git add kics-render-mock/pipeline.go
git commit -m "feat: add runOnce render pipeline (KICS-faithful)"
```

---

## Task 5: Implement `splitManifestYAML` (`pipeline.go`)

**Files:**
- Modify: `kics-render-mock/pipeline.go` (add `splitManifestYAML`)
- Modify: `kics-render-mock/pipeline_test.go` (add `TestSplitManifestYAML`)

- [ ] **Step 1: Write the failing test**

Add to `pipeline_test.go`:
```go
func TestSplitManifestYAML_SingleTemplate(t *testing.T) {
	templateSrc := "# KICS_HELM_ID_0:\napiVersion: v1\nkind: ConfigMap\nmetadata:\n  name: test"
	rel := &release.Release{
		Manifest: "---\n# Source: mychart/templates/test.yaml\n# KICS_HELM_ID_0:\napiVersion: v1\nkind: ConfigMap\nmetadata:\n  name: test\n",
		Chart: &chart.Chart{
			Metadata: &chart.Metadata{Name: "mychart"},
			Templates: []*chart.File{
				{Name: "templates/test.yaml", Data: []byte(templateSrc)},
			},
		},
	}

	entries := splitManifestYAML(rel)

	if len(entries) != 1 {
		t.Fatalf("expected 1 entry, got %d", len(entries))
	}
	e := entries[0]
	if e.Path != "mychart/templates/test.yaml" {
		t.Errorf("Path: got %q", e.Path)
	}
	if e.SplitID != "# KICS_HELM_ID_0:" {
		t.Errorf("SplitID: got %q", e.SplitID)
	}
	if !strings.Contains(e.OriginalContent, "KICS_HELM_ID_0:") {
		t.Errorf("OriginalContent should contain ID marker, got: %q", e.OriginalContent)
	}
}

func TestSplitManifestYAML_EmptyManifest(t *testing.T) {
	rel := &release.Release{
		Manifest: "",
		Chart: &chart.Chart{
			Metadata:  &chart.Metadata{Name: "mychart"},
			Templates: []*chart.File{},
		},
	}
	entries := splitManifestYAML(rel)
	if len(entries) != 0 {
		t.Errorf("expected 0 entries, got %d", len(entries))
	}
}
```

Add the missing import to `pipeline_test.go` imports block:
```go
import (
	"strings"
	"testing"

	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/release"
)
```

- [ ] **Step 2: Run to verify it fails**

```
go test ./... -run TestSplitManifestYAML 2>&1 | head -10
```

Expected: compilation error — `splitManifestYAML undefined`.

- [ ] **Step 3: Implement `splitManifestYAML`**

Append to `pipeline.go`:
```go
// splitManifestYAML mirrors KICS's splitManifestYAML: splits rel.Manifest by
// "---" separators, maps each section back to its source template file, and
// returns structured entries. originalContent is the post-setID template source.
func splitManifestYAML(rel *release.Release) []SplitManifestEntry {
	// Build path→originalContent map using the same updateName logic as KICS.
	sources := updateName(nil, rel.Chart, rel.Chart.Name())
	origData := make(map[string][]byte, len(sources))
	for _, f := range sources {
		key := filepath.ToSlash(f.Name) // manifest uses forward slashes
		origData[key] = []byte(strings.ReplaceAll(string(f.Data), "\r", ""))
	}

	var entries []SplitManifestEntry
	for _, section := range strings.Split(rel.Manifest, "---") {
		section = strings.ReplaceAll(section, "\r", "")
		if strings.TrimSpace(section) == "" {
			continue
		}

		lines := strings.Split(section, "\n")

		// Extract # Source: path
		path := ""
		for _, l := range lines {
			if strings.HasPrefix(l, "# Source: ") {
				path = strings.TrimSpace(strings.TrimPrefix(l, "# Source: "))
				break
			}
		}
		if path == "" {
			continue
		}

		orig, ok := origData[path]
		if !ok {
			continue
		}

		// Extract KICS split ID if present
		splitID := ""
		for _, l := range lines {
			if strings.Contains(l, kicsHelmID) {
				splitID = strings.TrimSpace(l)
				break
			}
		}

		entries = append(entries, SplitManifestEntry{
			Path:            path,
			Content:         section,
			OriginalContent: string(orig),
			SplitID:         splitID,
		})
	}
	return entries
}
```

- [ ] **Step 4: Run the tests**

```
go test ./... -run TestSplitManifestYAML -v
```

Expected:
```
--- PASS: TestSplitManifestYAML_SingleTemplate (0.00s)
--- PASS: TestSplitManifestYAML_EmptyManifest (0.00s)
PASS
```

- [ ] **Step 5: Commit**

```bash
git add kics-render-mock/pipeline.go kics-render-mock/pipeline_test.go
git commit -m "feat: implement splitManifestYAML"
```

---

## Task 6: Config loading and invocation building (`runner.go`)

**Files:**
- Write: `kics-render-mock/runner.go` (partial — types, `loadConfig`, `buildInvocations`)
- Create: `kics-render-mock/runner_test.go`

- [ ] **Step 1: Write the failing tests**

`runner_test.go`:
```go
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfig_RenderProblems(t *testing.T) {
	dir := t.TempDir()
	raw := `{
		"testNumber": 2,
		"testName": "Nil Pointer",
		"chartName": "test-nil-pointer",
		"testFunctions": [
			{"name": "testNilMapAccess", "func": "nil-map-access"},
			{"name": "testNilNestedMap", "func": "nil-nested-map"}
		],
		"dataTypes": ["nilMap", "nilNested"]
	}`
	if err := os.WriteFile(filepath.Join(dir, "config.json"), []byte(raw), 0644); err != nil {
		t.Fatal(err)
	}

	cfg, err := loadConfig(dir)
	if err != nil {
		t.Fatalf("loadConfig: %v", err)
	}
	if cfg.TestNumber != 2 {
		t.Errorf("TestNumber: got %d", cfg.TestNumber)
	}
	if len(cfg.TestFunctions) != 2 {
		t.Errorf("TestFunctions len: got %d", len(cfg.TestFunctions))
	}
	if cfg.TestFunctions[0].Name != "testNilMapAccess" {
		t.Errorf("TestFunctions[0].Name: got %q", cfg.TestFunctions[0].Name)
	}
	if len(cfg.DataTypes) != 2 {
		t.Errorf("DataTypes len: got %d", len(cfg.DataTypes))
	}
}

func TestLoadConfig_OverrideTests(t *testing.T) {
	dir := t.TempDir()
	raw := `{
		"testNumber": 1,
		"testName": "Basic Data Types",
		"chartName": "test-basic",
		"testCases": [
			{"name": "default-values", "valuesFiles": [], "expected": {}},
			{"name": "string-override", "valuesFiles": ["values-string.yaml"], "expected": {}}
		]
	}`
	if err := os.WriteFile(filepath.Join(dir, "config.json"), []byte(raw), 0644); err != nil {
		t.Fatal(err)
	}

	cfg, err := loadConfig(dir)
	if err != nil {
		t.Fatalf("loadConfig: %v", err)
	}
	if len(cfg.TestCases) != 2 {
		t.Errorf("TestCases len: got %d", len(cfg.TestCases))
	}
	if cfg.TestCases[1].ValuesFiles[0] != "values-string.yaml" {
		t.Errorf("TestCases[1].ValuesFiles[0]: got %q", cfg.TestCases[1].ValuesFiles[0])
	}
}

func TestBuildInvocations_RenderProblems(t *testing.T) {
	cfg := &testConfig{
		TestFunctions: []testFunction{
			{Name: "testNilMapAccess", Func: "nil-map-access"},
		},
		DataTypes: []string{"nilMap", "nilNested"},
	}
	invs := buildInvocations(cfg, "/test/dir")
	if len(invs) != 2 {
		t.Fatalf("expected 2 invocations, got %d", len(invs))
	}
	if invs[0].toggle != "nil-map-access" {
		t.Errorf("toggle: got %q", invs[0].toggle)
	}
	if invs[0].dataType != "nilMap" {
		t.Errorf("dataType: got %q", invs[0].dataType)
	}
	if len(invs[0].valOpts.Values) != 2 {
		t.Fatalf("Values len: got %d", len(invs[0].valOpts.Values))
	}
	if invs[0].valOpts.Values[0] != "testNilMapAccess=true" {
		t.Errorf("Values[0]: got %q", invs[0].valOpts.Values[0])
	}
	if invs[0].valOpts.Values[1] != "testDataType=nilMap" {
		t.Errorf("Values[1]: got %q", invs[0].valOpts.Values[1])
	}
}

func TestBuildInvocations_OverrideTests(t *testing.T) {
	cfg := &testConfig{
		TestCases: []testCase{
			{Name: "string-override", ValuesFiles: []string{"values-string.yaml"}},
		},
	}
	invs := buildInvocations(cfg, "/test/dir")
	if len(invs) != 1 {
		t.Fatalf("expected 1 invocation, got %d", len(invs))
	}
	if invs[0].caseName != "string-override" {
		t.Errorf("caseName: got %q", invs[0].caseName)
	}
	if len(invs[0].valOpts.ValueFiles) != 1 {
		t.Fatalf("ValueFiles len: got %d", len(invs[0].valOpts.ValueFiles))
	}
	// Value file must be absolute (joined with testDir)
	expected := filepath.Join("/test/dir", "values-string.yaml")
	if invs[0].valOpts.ValueFiles[0] != expected {
		t.Errorf("ValueFiles[0]: got %q, want %q", invs[0].valOpts.ValueFiles[0], expected)
	}
}

func TestBuildInvocations_DefaultDataType(t *testing.T) {
	cfg := &testConfig{
		TestFunctions: []testFunction{
			{Name: "testFoo", Func: "foo"},
		},
		DataTypes: []string{}, // empty → default
	}
	invs := buildInvocations(cfg, "/test/dir")
	if len(invs) != 1 {
		t.Fatalf("expected 1 invocation, got %d", len(invs))
	}
	// When dataType is "default", testDataType is NOT added to Values
	if len(invs[0].valOpts.Values) != 1 {
		t.Errorf("Values len for default dataType: got %d, want 1", len(invs[0].valOpts.Values))
	}
}
```

- [ ] **Step 2: Run to verify it fails**

```
go test ./... -run TestLoadConfig -run TestBuildInvocations 2>&1 | head -10
```

Expected: compilation error — types undefined.

- [ ] **Step 3: Implement config types, `loadConfig`, and `buildInvocations`**

`runner.go`:
```go
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"helm.sh/helm/v3/pkg/cli/values"
)

// testConfig is a unified shape that covers both suite config.json formats.
// Presence of TestFunctions → render_problems suite.
// Presence of TestCases    → override_tests_suite.
type testConfig struct {
	TestNumber    int            `json:"testNumber"`
	TestName      string         `json:"testName"`
	ChartName     string         `json:"chartName"`
	TestFunctions []testFunction `json:"testFunctions,omitempty"`
	DataTypes     []string       `json:"dataTypes,omitempty"`
	TestCases     []testCase     `json:"testCases,omitempty"`
}

type testFunction struct {
	Name string `json:"name"` // --set flag name, e.g. "testNilMapAccess"
	Func string `json:"func"` // display key,    e.g. "nil-map-access"
}

type testCase struct {
	Name           string   `json:"name"`
	ValuesFiles    []string `json:"valuesFiles"`
	SetFlags       []string `json:"setFlags,omitempty"`
	SetStringFlags []string `json:"setStringFlags,omitempty"`
}

// invocation is one render call to make.
type invocation struct {
	toggle   string // render_problems: display key (Func field)
	dataType string // render_problems: data type name
	caseName string // override_tests_suite: test case name
	valOpts  *values.Options
}

func loadConfig(testDir string) (*testConfig, error) {
	data, err := os.ReadFile(filepath.Join(testDir, "config.json"))
	if err != nil {
		return nil, fmt.Errorf("read config.json: %w", err)
	}
	var cfg testConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parse config.json: %w", err)
	}
	return &cfg, nil
}

func buildInvocations(cfg *testConfig, testDir string) []invocation {
	if len(cfg.TestFunctions) > 0 {
		return buildRenderProblemsInvocations(cfg)
	}
	return buildOverrideTestsInvocations(cfg, testDir)
}

func buildRenderProblemsInvocations(cfg *testConfig) []invocation {
	dataTypes := cfg.DataTypes
	if len(dataTypes) == 0 {
		dataTypes = []string{"default"}
	}
	var invs []invocation
	for _, tf := range cfg.TestFunctions {
		for _, dt := range dataTypes {
			vals := []string{fmt.Sprintf("%s=true", tf.Name)}
			if dt != "default" {
				vals = append(vals, fmt.Sprintf("testDataType=%s", dt))
			}
			invs = append(invs, invocation{
				toggle:   tf.Func,
				dataType: dt,
				valOpts:  &values.Options{Values: vals},
			})
		}
	}
	return invs
}

func buildOverrideTestsInvocations(cfg *testConfig, testDir string) []invocation {
	var invs []invocation
	for _, tc := range cfg.TestCases {
		absFiles := make([]string, len(tc.ValuesFiles))
		for i, vf := range tc.ValuesFiles {
			absFiles[i] = filepath.Join(testDir, vf)
		}
		invs = append(invs, invocation{
			caseName: tc.Name,
			valOpts: &values.Options{
				ValueFiles:   absFiles,
				Values:       tc.SetFlags,
				StringValues: tc.SetStringFlags,
			},
		})
	}
	return invs
}
```

- [ ] **Step 4: Run the tests**

```
go test ./... -run "TestLoadConfig|TestBuildInvocations" -v
```

Expected: all 5 tests PASS.

- [ ] **Step 5: Commit**

```bash
git add kics-render-mock/runner.go kics-render-mock/runner_test.go
git commit -m "feat: implement config loading and invocation builder"
```

---

## Task 7: Implement `runTestDir` and result assembly (`runner.go`)

**Files:**
- Modify: `kics-render-mock/runner.go` (add `toStandardResult`, `toDebugResult`, `buildRenderEntry`, `runTestDir`)

- [ ] **Step 1: Append the result assembly and runner functions to `runner.go`**

```go
func toStandardResult(res renderResult) RunResult {
	r := RunResult{SplitManifests: []SplitManifestEntry{}}
	if res.err != nil {
		errStr := res.err.Error()
		r.Error = &errStr
		if res.rel != nil && res.rel.Manifest != "" {
			r.PartialManifest = &res.rel.Manifest
		}
		return r
	}
	if res.rel != nil {
		r.RawManifest = &res.rel.Manifest
		r.SplitManifests = splitManifestYAML(res.rel)
	}
	return r
}

func toDebugResult(res renderResult) DebugRunResult {
	logs := res.logs
	if logs == nil {
		logs = []string{}
	}
	r := DebugRunResult{DebugLogs: logs, SplitManifests: []SplitManifestEntry{}}
	if res.err != nil {
		errStr := res.err.Error()
		r.Error = &errStr
		if res.rel != nil && res.rel.Manifest != "" {
			r.PartialManifest = &res.rel.Manifest
		}
		return r
	}
	if res.rel != nil {
		r.RawManifest = &res.rel.Manifest
		r.SplitManifests = splitManifestYAML(res.rel)
	}
	return r
}

func buildRenderEntry(inv invocation, stdRes, dbgRes renderResult) RenderEntry {
	vf := inv.valOpts.ValueFiles
	if vf == nil {
		vf = []string{}
	}
	sv := inv.valOpts.StringValues
	vals := inv.valOpts.Values
	if vals == nil {
		vals = []string{}
	}
	return RenderEntry{
		Toggle:   inv.toggle,
		DataType: inv.dataType,
		CaseName: inv.caseName,
		ValuesOptions: ValuesOptionsSummary{
			Values:       vals,
			ValueFiles:   vf,
			StringValues: sv,
		},
		Standard: toStandardResult(stdRes),
		Debug:    toDebugResult(dbgRes),
	}
}

func runTestDir(testDir, suite string) (RenderOutput, error) {
	cfg, err := loadConfig(testDir)
	if err != nil {
		return RenderOutput{}, err
	}

	invocations := buildInvocations(cfg, testDir)

	out := RenderOutput{
		Suite:      suite,
		TestNumber: cfg.TestNumber,
		TestName:   cfg.TestName,
		ChartPath:  testDir,
		Renders:    make([]RenderEntry, 0, len(invocations)),
	}

	for _, inv := range invocations {
		stdRes := runOnce(testDir, inv.valOpts, false)
		dbgRes := runOnce(testDir, inv.valOpts, true)
		out.Renders = append(out.Renders, buildRenderEntry(inv, stdRes, dbgRes))
	}

	return out, nil
}
```

- [ ] **Step 2: Verify the module still compiles**

```
go build ./...
```

Expected: no output, exit 0.

- [ ] **Step 3: Run all existing tests to ensure nothing broke**

```
go test ./... -v
```

Expected: all previously passing tests still PASS.

- [ ] **Step 4: Commit**

```bash
git add kics-render-mock/runner.go
git commit -m "feat: implement runTestDir and render entry assembly"
```

---

## Task 8: Implement main entry point (`main.go`)

**Files:**
- Write: `kics-render-mock/main.go`

- [ ] **Step 1: Write `main.go`**

```go
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
```

- [ ] **Step 2: Build the binary**

Run from inside `kics-render-mock/`:
```
go build -o kics-render-mock .
```

Expected: `kics-render-mock` (or `kics-render-mock.exe` on Windows) binary created, exit 0.

- [ ] **Step 3: Commit**

```bash
git add kics-render-mock/main.go
git commit -m "feat: implement main entry point with suite discovery"
```

---

## Task 9: Integration run and verification

**Files:** none — read-only verification

- [ ] **Step 1: Run the tool from inside `kics-render-mock/`**

```
go run . 
```

Expected stdout (numbers will vary):
```
=== KICS Render Mock ===
render_problems           21 dirs    NNN invocations    NNN errors
override_tests_suite      10 dirs    NNN invocations      0 errors
JSON files written to each test directory.
```

- [ ] **Step 2: Verify a JSON file was created for a render_problems test**

```
cat ..\render_problems\test-02-nil-pointer\kics-render-output.json | head -60
```

Expected: valid JSON with `"suite": "render_problems"`, `"renders"` array, each entry with `"standard"` and `"debug"` objects.

- [ ] **Step 3: Verify a debug run captured a partial manifest**

```powershell
$json = Get-Content "..\render_problems\test-02-nil-pointer\kics-render-output.json" | ConvertFrom-Json
$errorRenders = $json.renders | Where-Object { $_.standard.error -ne $null }
$withPartial  = $errorRenders | Where-Object { $_.debug.partialManifest -ne $null }
Write-Host "Error renders: $($errorRenders.Count)"
Write-Host "With partial manifest in debug: $($withPartial.Count)"
```

Expected: `Error renders` > 0; `With partial manifest in debug` ideally > 0 (shows the debug run recovered more than the standard run).

- [ ] **Step 4: Verify an override_tests_suite JSON file**

```
cat ..\override_tests_suite\test-01-basic-data-types\kics-render-output.json | head -40
```

Expected: `"suite": "override_tests_suite"`, renders keyed by `caseName`, `standard.error` is `null` for successful renders.

- [ ] **Step 5: Run all unit tests one final time**

```
go test ./... -v
```

Expected: all tests PASS.

- [ ] **Step 6: Final commit**

```bash
git add kics-render-mock/
git commit -m "feat: complete kics-render-mock tool with integration verified"
```

---

## Self-Review Notes

- `addID` and `setID` mutate chart template data in place — this is intentional and matches KICS exactly. Each `runOnce` loads the chart fresh from disk, so there is no cross-run contamination.
- `updateName` also mutates template `Name` fields. Called only inside `splitManifestYAML` after the render is complete — safe.
- `helmSettings.Debug` is a package-level global. Standard and debug runs execute sequentially per invocation so there is no race condition.
- `values.Options.MergeValues` only reads files; passing the same `*values.Options` to both runs is safe.
- `log.SetOutput(io.Discard)` suppresses Helm's internal `log` package noise during render. `defer log.SetOutput(nil)` resets to nil (no output) rather than `os.Stderr` to avoid log noise leaking to test output — adjust to `os.Stderr` if needed for debugging.
