# taxonomy_analyzer Fix-Outcome Enrichment — Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Add an optional `--fixed` flag to `taxonomy_analyzer` that loads `catalog_fixed.json`, joins its per-run fix outcomes to the primary catalog by `helm_command`, and adds fix stats (attempted / resolved / unresolved / by_stop_reason) to every taxonomy bucket in all outputs.

**Architecture:** The loader gains `LoadFixedIndex` which streams `catalog_fixed.json` and returns a `map[string]*FixedResult` keyed by `helm_command`. The analyzer holds this map and, for each failed run, looks up the matching fix outcome and attaches it to the `ErrorOccurrence`. The exporter renders fix columns in markdown and CSV when `FixAttempted > 0`.

**Tech Stack:** Go 1.25, standard library only (`encoding/json`, `os`, `fmt`, `flag`, `strings`). Tests use the built-in `testing` package.

---

## File Map

| File | Change |
|---|---|
| `model/model.go` | Add `FixAttempt`, `FixedResult`, `FixOutcome`; extend `RunResult`, `ErrorOccurrence`, `TaxonomyBucket`, `ReportTotals`, `TaxonomyReport` |
| `loader/loader.go` | Add `LoadFixedIndex` function |
| `loader/loader_test.go` | New test file |
| `analyzer/analyzer.go` | Update `Analyzer` struct, `New`, `ConsumeRepo`, `consumeOccurrence`; add `bumpFixOutcome` |
| `analyzer/analyzer_test.go` | New test file |
| `exporter/exporter.go` | Update `writeMarkdown`, `writeCSV` |
| `exporter/exporter_test.go` | New test file |
| `main.go` | Add `--fixed` flag and wiring |

All commands run from `taxonomy_analyzer/`.

---

## Task 1: Extend model types

**Files:**
- Modify: `model/model.go`

- [ ] **Step 1: Replace model.go with the extended version**

```go
package model

import "time"

type FixAttempt struct {
	Attempt       int    `json:"attempt"`
	ErrorSeen     string `json:"error_seen"`
	Kind          string `json:"kind"`
	ValuePath     string `json:"value_path"`
	ValueInjected string `json:"value_injected"`
}

type FixedResult struct {
	Resolved   bool         `json:"resolved"`
	StopReason string       `json:"stop_reason"`
	FixChain   []FixAttempt `json:"fix_chain"`
}

type FixOutcome struct {
	Attempted    int            `json:"attempted"`
	Resolved     int            `json:"resolved"`
	Unresolved   int            `json:"unresolved"`
	ByStopReason map[string]int `json:"by_stop_reason"`
}

type RunResult struct {
	ChartPath    string       `json:"chart_path"`
	ValuesFiles  []string     `json:"values_files,omitempty"`
	HelmCommand  string       `json:"helm_command"`
	Success      bool         `json:"success"`
	ErrorMessage string       `json:"error_message,omitempty"`
	Fixed        *FixedResult `json:"fixed,omitempty"`
}

type ChartSummary struct {
	ChartPath       string      `json:"chart_path"`
	TotalRuns       int         `json:"total_runs"`
	Successes       int         `json:"successes"`
	Failures        int         `json:"failures"`
	DepBuildFailure bool        `json:"dep_build_failure,omitempty"`
	DepBuildError   string      `json:"dep_build_error,omitempty"`
	Runs            []RunResult `json:"runs"`
}

type RepoResult struct {
	RepoURL          string         `json:"repo_url"`
	RepoName         string         `json:"repo_name"`
	ClonedDir        string         `json:"cloned_dir"`
	TotalCharts      int            `json:"total_charts"`
	TotalRuns        int            `json:"total_runs"`
	TotalSuccess     int            `json:"total_successes"`
	TotalFailures    int            `json:"total_failures"`
	TotalDepFailures int            `json:"total_dep_failures"`
	Charts           []ChartSummary `json:"charts"`
	Kept             bool           `json:"kept"`
	DepFailed        bool           `json:"dep_failed"`
}

type ErrorOccurrence struct {
	RepoURL      string       `json:"repo_url"`
	RepoName     string       `json:"repo_name"`
	ChartPath    string       `json:"chart_path"`
	ValuesFiles  []string     `json:"values_files,omitempty"`
	HelmCommand  string       `json:"helm_command,omitempty"`
	ErrorSource  string       `json:"error_source"`
	ErrorKind    string       `json:"error_kind"`
	ErrorSubKind string       `json:"error_sub_kind"`
	ErrorMessage string       `json:"error_message"`
	Fixed        *FixedResult `json:"fixed,omitempty"`
}

type TaxonomyBucket struct {
	Count      int               `json:"count"`
	Examples   []ErrorOccurrence `json:"examples"`
	FixOutcome FixOutcome        `json:"fix_outcome"`
}

type ReportTotals struct {
	Repos             int `json:"repos"`
	Runs              int `json:"runs"`
	TemplateFailures  int `json:"template_failures"`
	DependencyErrors  int `json:"dependency_failures"`
	ClassifiedErrors  int `json:"classified_errors"`
	UnclassifiedError int `json:"unclassified_errors"`
	FixAttempted      int `json:"fix_attempted,omitempty"`
	FixResolved       int `json:"fix_resolved,omitempty"`
	FixUnresolved     int `json:"fix_unresolved,omitempty"`
}

type TaxonomyReport struct {
	GeneratedAt   time.Time                 `json:"generated_at"`
	SourceCatalog string                    `json:"source_catalog"`
	FixedCatalog  string                    `json:"fixed_catalog,omitempty"`
	Totals        ReportTotals              `json:"totals"`
	ByKind        map[string]TaxonomyBucket `json:"by_kind"`
	BySubKind     map[string]TaxonomyBucket `json:"by_sub_kind"`
	ByRepo        map[string]map[string]int `json:"by_repo"`
	Unclassified  []ErrorOccurrence         `json:"unclassified"`
	AllClassified []ErrorOccurrence         `json:"all_classified"`
}
```

- [ ] **Step 2: Verify compilation**

```
go build ./...
```

Expected: no output (compiles cleanly).

- [ ] **Step 3: Commit**

```bash
git add model/model.go
git commit -m "feat(taxonomy_analyzer): add FixedResult and FixOutcome model types"
```

---

## Task 2: Add LoadFixedIndex to loader

**Files:**
- Modify: `loader/loader.go`
- Create: `loader/loader_test.go`

- [ ] **Step 1: Write the failing test**

Create `loader/loader_test.go`:

```go
package loader_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/MabsIPCA/helm-tests/taxonomy_analyzer/loader"
)

const fixedSample = `[
  {
    "repo_url": "https://github.com/test/repo-a",
    "repo_name": "test/repo-a",
    "cloned_dir": "/tmp/repo-a",
    "total_charts": 1, "total_runs": 2, "total_successes": 0, "total_failures": 2, "total_dep_failures": 0,
    "charts": [
      {
        "chart_path": "/tmp/repo-a/chart",
        "total_runs": 2, "successes": 0, "failures": 2,
        "runs": [
          {
            "chart_path": "cloned/repo-a/chart",
            "helm_command": "helm template test cloned/repo-a/chart",
            "success": false,
            "error_message": "nil pointer evaluating interface {}.type",
            "fixed": {
              "resolved": true,
              "stop_reason": "",
              "fix_chain": [{"attempt": 1, "error_seen": "nil pointer evaluating interface {}.type", "kind": "nil_pointer", "value_path": "service.type", "value_injected": "ClusterIP"}]
            }
          },
          {
            "chart_path": "cloned/repo-a/chart",
            "helm_command": "helm template test cloned/repo-a/chart -f values.yaml",
            "success": false,
            "error_message": "nil pointer evaluating interface {}.repository",
            "fixed": {"resolved": false, "stop_reason": "unfixable_error_kind", "fix_chain": []}
          }
        ]
      }
    ],
    "kept": true, "dep_failed": false
  }
]`

func TestLoadFixedIndex_IndexesByHelmCommand(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "catalog_fixed.json")
	if err := os.WriteFile(path, []byte(fixedSample), 0o644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}

	idx, err := loader.LoadFixedIndex(path)
	if err != nil {
		t.Fatalf("LoadFixedIndex: %v", err)
	}

	if len(idx) != 2 {
		t.Errorf("index length: got %d, want 2", len(idx))
	}

	cmd1 := "helm template test cloned/repo-a/chart"
	if idx[cmd1] == nil {
		t.Fatalf("missing key %q", cmd1)
	}
	if !idx[cmd1].Resolved {
		t.Errorf("%q: Resolved got false, want true", cmd1)
	}
	if len(idx[cmd1].FixChain) != 1 {
		t.Errorf("%q: FixChain len got %d, want 1", cmd1, len(idx[cmd1].FixChain))
	}
	if idx[cmd1].FixChain[0].Kind != "nil_pointer" {
		t.Errorf("%q: FixChain[0].Kind got %q, want %q", cmd1, idx[cmd1].FixChain[0].Kind, "nil_pointer")
	}

	cmd2 := "helm template test cloned/repo-a/chart -f values.yaml"
	if idx[cmd2] == nil {
		t.Fatalf("missing key %q", cmd2)
	}
	if idx[cmd2].Resolved {
		t.Errorf("%q: Resolved got true, want false", cmd2)
	}
	if idx[cmd2].StopReason != "unfixable_error_kind" {
		t.Errorf("%q: StopReason got %q, want %q", cmd2, idx[cmd2].StopReason, "unfixable_error_kind")
	}
}

func TestLoadFixedIndex_MissingFile_ReturnsError(t *testing.T) {
	_, err := loader.LoadFixedIndex("/nonexistent/path.json")
	if err == nil {
		t.Error("expected error for missing file, got nil")
	}
}
```

- [ ] **Step 2: Run test to verify it fails**

```
go test ./loader/...
```

Expected: FAIL — `loader.LoadFixedIndex undefined`

- [ ] **Step 3: Add LoadFixedIndex to loader.go**

Append to `loader/loader.go` (after the existing `StreamCatalog` function):

```go
// LoadFixedIndex streams catalog_fixed.json and returns a map from
// helm_command -> *FixedResult for correlation with primary catalog runs.
func LoadFixedIndex(path string) (map[string]*model.FixedResult, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open fixed catalog: %w", err)
	}
	defer f.Close()

	decoder := json.NewDecoder(f)

	tok, err := decoder.Token()
	if err != nil {
		return nil, fmt.Errorf("read opening token: %w", err)
	}
	delim, ok := tok.(json.Delim)
	if !ok || delim != '[' {
		return nil, fmt.Errorf("invalid fixed catalog format: expected JSON array")
	}

	index := make(map[string]*model.FixedResult)
	for decoder.More() {
		var repo model.RepoResult
		if err := decoder.Decode(&repo); err != nil {
			return nil, fmt.Errorf("decode fixed repository entry: %w", err)
		}
		for _, chart := range repo.Charts {
			for _, run := range chart.Runs {
				if run.Fixed != nil && run.HelmCommand != "" {
					fixed := run.Fixed
					index[run.HelmCommand] = fixed
				}
			}
		}
	}

	if _, err = decoder.Token(); err != nil {
		return nil, fmt.Errorf("read closing token: %w", err)
	}

	return index, nil
}
```

- [ ] **Step 4: Run test to verify it passes**

```
go test ./loader/...
```

Expected: `ok  github.com/MabsIPCA/helm-tests/taxonomy_analyzer/loader`

- [ ] **Step 5: Commit**

```bash
git add loader/loader.go loader/loader_test.go
git commit -m "feat(taxonomy_analyzer): add LoadFixedIndex to loader"
```

---

## Task 3: Update Analyzer for fix enrichment

**Files:**
- Modify: `analyzer/analyzer.go`
- Create: `analyzer/analyzer_test.go`

- [ ] **Step 1: Write the failing tests**

Create `analyzer/analyzer_test.go`:

```go
package analyzer_test

import (
	"testing"

	"github.com/MabsIPCA/helm-tests/taxonomy_analyzer/analyzer"
	"github.com/MabsIPCA/helm-tests/taxonomy_analyzer/model"
)

func repoWithTwoNilPointerRuns() model.RepoResult {
	return model.RepoResult{
		RepoURL:       "https://github.com/test/repo-a",
		RepoName:      "test/repo-a",
		TotalRuns:     2,
		TotalFailures: 2,
		Charts: []model.ChartSummary{
			{
				ChartPath: "/tmp/repo-a/chart",
				TotalRuns: 2,
				Failures:  2,
				Runs: []model.RunResult{
					{
						ChartPath:    "cloned/repo-a/chart",
						HelmCommand:  "helm template test cloned/repo-a/chart",
						Success:      false,
						ErrorMessage: "nil pointer evaluating interface {}.type",
					},
					{
						ChartPath:    "cloned/repo-a/chart",
						HelmCommand:  "helm template test cloned/repo-a/chart -f values.yaml",
						Success:      false,
						ErrorMessage: "nil pointer evaluating interface {}.repository",
					},
				},
			},
		},
	}
}

func TestConsumeRepo_WithFixedIndex_AccumulatesFixStats(t *testing.T) {
	fixedIndex := map[string]*model.FixedResult{
		"helm template test cloned/repo-a/chart": {
			Resolved:   true,
			StopReason: "",
			FixChain:   []model.FixAttempt{{Attempt: 1, Kind: "nil_pointer", ValuePath: "service.type", ValueInjected: "ClusterIP"}},
		},
		"helm template test cloned/repo-a/chart -f values.yaml": {
			Resolved:   false,
			StopReason: "unfixable_error_kind",
			FixChain:   []model.FixAttempt{},
		},
	}

	a := analyzer.New("catalog.json", "catalog_fixed.json", fixedIndex)
	a.ConsumeRepo(repoWithTwoNilPointerRuns())
	report := a.Report()

	if report.FixedCatalog != "catalog_fixed.json" {
		t.Errorf("FixedCatalog: got %q, want %q", report.FixedCatalog, "catalog_fixed.json")
	}
	if report.Totals.FixAttempted != 2 {
		t.Errorf("FixAttempted: got %d, want 2", report.Totals.FixAttempted)
	}
	if report.Totals.FixResolved != 1 {
		t.Errorf("FixResolved: got %d, want 1", report.Totals.FixResolved)
	}
	if report.Totals.FixUnresolved != 1 {
		t.Errorf("FixUnresolved: got %d, want 1", report.Totals.FixUnresolved)
	}

	subBucket := report.BySubKind["template.nil_pointer"]
	if subBucket.FixOutcome.Attempted != 2 {
		t.Errorf("nil_pointer FixOutcome.Attempted: got %d, want 2", subBucket.FixOutcome.Attempted)
	}
	if subBucket.FixOutcome.Resolved != 1 {
		t.Errorf("nil_pointer FixOutcome.Resolved: got %d, want 1", subBucket.FixOutcome.Resolved)
	}
	if subBucket.FixOutcome.Unresolved != 1 {
		t.Errorf("nil_pointer FixOutcome.Unresolved: got %d, want 1", subBucket.FixOutcome.Unresolved)
	}
	if subBucket.FixOutcome.ByStopReason["unfixable_error_kind"] != 1 {
		t.Errorf("nil_pointer ByStopReason[unfixable_error_kind]: got %d, want 1", subBucket.FixOutcome.ByStopReason["unfixable_error_kind"])
	}

	kindBucket := report.ByKind["template"]
	if kindBucket.FixOutcome.Attempted != 2 {
		t.Errorf("template FixOutcome.Attempted: got %d, want 2", kindBucket.FixOutcome.Attempted)
	}
}

func TestConsumeRepo_WithoutFixedIndex_NoFixStats(t *testing.T) {
	a := analyzer.New("catalog.json", "", nil)
	a.ConsumeRepo(repoWithTwoNilPointerRuns())
	report := a.Report()

	if report.Totals.FixAttempted != 0 {
		t.Errorf("FixAttempted: got %d, want 0", report.Totals.FixAttempted)
	}
	if report.BySubKind["template.nil_pointer"].FixOutcome.Attempted != 0 {
		t.Errorf("bucket FixOutcome.Attempted: got %d, want 0", report.BySubKind["template.nil_pointer"].FixOutcome.Attempted)
	}
}

func TestConsumeRepo_RunNotInFixedIndex_NoFixStats(t *testing.T) {
	fixedIndex := map[string]*model.FixedResult{} // empty — no matching entry

	a := analyzer.New("catalog.json", "catalog_fixed.json", fixedIndex)
	a.ConsumeRepo(repoWithTwoNilPointerRuns())
	report := a.Report()

	if report.Totals.FixAttempted != 0 {
		t.Errorf("FixAttempted: got %d, want 0", report.Totals.FixAttempted)
	}
}
```

- [ ] **Step 2: Run test to verify it fails**

```
go test ./analyzer/...
```

Expected: FAIL — `analyzer.New` does not accept 3 arguments

- [ ] **Step 3: Replace analyzer.go with the updated version**

```go
package analyzer

import (
	"fmt"
	"sort"
	"time"

	"github.com/MabsIPCA/helm-tests/taxonomy_analyzer/classifier"
	"github.com/MabsIPCA/helm-tests/taxonomy_analyzer/model"
)

const maxExamplesPerBucket = 5

// Analyzer keeps running aggregation state while catalog entries are streamed.
type Analyzer struct {
	report     model.TaxonomyReport
	fixedIndex map[string]*model.FixedResult
}

// New creates a new analyzer report collector.
// fixedCatalog and fixedIndex may be empty/nil when --fixed is not provided.
func New(sourceCatalog, fixedCatalog string, fixedIndex map[string]*model.FixedResult) *Analyzer {
	r := model.TaxonomyReport{
		GeneratedAt:   time.Now().UTC(),
		SourceCatalog: sourceCatalog,
		ByKind:        map[string]model.TaxonomyBucket{},
		BySubKind:     map[string]model.TaxonomyBucket{},
		ByRepo:        map[string]map[string]int{},
	}
	if fixedCatalog != "" {
		r.FixedCatalog = fixedCatalog
	}
	return &Analyzer{report: r, fixedIndex: fixedIndex}
}

// ConsumeRepo feeds one repo result into the taxonomy analyzer.
func (a *Analyzer) ConsumeRepo(repo model.RepoResult) {
	a.report.Totals.Repos++
	a.report.Totals.Runs += repo.TotalRuns
	a.report.Totals.TemplateFailures += repo.TotalFailures
	a.report.Totals.DependencyErrors += repo.TotalDepFailures

	for _, chart := range repo.Charts {
		if chart.DepBuildFailure && chart.DepBuildError != "" {
			a.consumeOccurrence(model.ErrorOccurrence{
				RepoURL:      repo.RepoURL,
				RepoName:     repo.RepoName,
				ChartPath:    chart.ChartPath,
				ErrorSource:  "dependency",
				ErrorMessage: chart.DepBuildError,
			})
		}

		for _, run := range chart.Runs {
			if run.Success || run.ErrorMessage == "" {
				continue
			}
			occ := model.ErrorOccurrence{
				RepoURL:      repo.RepoURL,
				RepoName:     repo.RepoName,
				ChartPath:    run.ChartPath,
				ValuesFiles:  run.ValuesFiles,
				HelmCommand:  run.HelmCommand,
				ErrorSource:  "template",
				ErrorMessage: run.ErrorMessage,
			}
			if a.fixedIndex != nil {
				occ.Fixed = a.fixedIndex[run.HelmCommand]
			}
			a.consumeOccurrence(occ)
		}
	}
}

func (a *Analyzer) consumeOccurrence(occ model.ErrorOccurrence) {
	res := classifier.Classify(occ.ErrorMessage, occ.ErrorSource)
	occ.ErrorKind = res.Kind
	occ.ErrorSubKind = res.SubKind

	if !res.Classified {
		a.report.Totals.UnclassifiedError++
		a.report.Unclassified = append(a.report.Unclassified, occ)
	} else {
		a.report.Totals.ClassifiedErrors++
	}

	kindKey := occ.ErrorKind
	subKindKey := fmt.Sprintf("%s.%s", occ.ErrorKind, occ.ErrorSubKind)
	a.bumpBucket(a.report.ByKind, kindKey, occ)
	a.bumpBucket(a.report.BySubKind, subKindKey, occ)

	if occ.Fixed != nil {
		a.report.Totals.FixAttempted++
		if occ.Fixed.Resolved {
			a.report.Totals.FixResolved++
		} else {
			a.report.Totals.FixUnresolved++
		}
		a.bumpFixOutcome(a.report.ByKind, kindKey, occ.Fixed)
		a.bumpFixOutcome(a.report.BySubKind, subKindKey, occ.Fixed)
	}

	if _, ok := a.report.ByRepo[occ.RepoName]; !ok {
		a.report.ByRepo[occ.RepoName] = map[string]int{}
	}
	a.report.ByRepo[occ.RepoName][subKindKey]++
	a.report.AllClassified = append(a.report.AllClassified, occ)
}

func (a *Analyzer) bumpBucket(buckets map[string]model.TaxonomyBucket, key string, occ model.ErrorOccurrence) {
	bucket := buckets[key]
	bucket.Count++
	if len(bucket.Examples) < maxExamplesPerBucket {
		bucket.Examples = append(bucket.Examples, occ)
	}
	buckets[key] = bucket
}

func (a *Analyzer) bumpFixOutcome(buckets map[string]model.TaxonomyBucket, key string, fixed *model.FixedResult) {
	bucket := buckets[key]
	fo := bucket.FixOutcome
	fo.Attempted++
	if fixed.Resolved {
		fo.Resolved++
	} else {
		fo.Unresolved++
		if fixed.StopReason != "" {
			if fo.ByStopReason == nil {
				fo.ByStopReason = map[string]int{}
			}
			fo.ByStopReason[fixed.StopReason]++
		}
	}
	bucket.FixOutcome = fo
	buckets[key] = bucket
}

// Report returns a normalized copy of the report for exporting.
func (a *Analyzer) Report() model.TaxonomyReport {
	report := a.report
	report.ByKind = sortBucketsByCount(report.ByKind)
	report.BySubKind = sortBucketsByCount(report.BySubKind)
	report.ByRepo = sortRepoKinds(report.ByRepo)
	return report
}

func sortBucketsByCount(src map[string]model.TaxonomyBucket) map[string]model.TaxonomyBucket {
	type pair struct {
		key   string
		count int
	}
	pairs := make([]pair, 0, len(src))
	for k, v := range src {
		pairs = append(pairs, pair{key: k, count: v.Count})
	}
	sort.SliceStable(pairs, func(i, j int) bool {
		if pairs[i].count == pairs[j].count {
			return pairs[i].key < pairs[j].key
		}
		return pairs[i].count > pairs[j].count
	})

	out := make(map[string]model.TaxonomyBucket, len(src))
	for _, p := range pairs {
		out[p.key] = src[p.key]
	}
	return out
}

func sortRepoKinds(src map[string]map[string]int) map[string]map[string]int {
	repoNames := make([]string, 0, len(src))
	for repo := range src {
		repoNames = append(repoNames, repo)
	}
	sort.Strings(repoNames)

	out := make(map[string]map[string]int, len(src))
	for _, repo := range repoNames {
		kinds := src[repo]
		keys := make([]string, 0, len(kinds))
		for k := range kinds {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		normalized := make(map[string]int, len(kinds))
		for _, k := range keys {
			normalized[k] = kinds[k]
		}
		out[repo] = normalized
	}
	return out
}
```

- [ ] **Step 4: Run tests to verify they pass**

```
go test ./analyzer/...
```

Expected: `ok  github.com/MabsIPCA/helm-tests/taxonomy_analyzer/analyzer`

- [ ] **Step 5: Run all tests to verify no regressions**

```
go test ./...
```

Expected: all packages pass.

- [ ] **Step 6: Commit**

```bash
git add analyzer/analyzer.go analyzer/analyzer_test.go
git commit -m "feat(taxonomy_analyzer): enrich analyzer with fix outcomes per bucket"
```

---

## Task 4: Update Exporter for fix stats output

**Files:**
- Modify: `exporter/exporter.go`
- Create: `exporter/exporter_test.go`

- [ ] **Step 1: Write the failing tests**

Create `exporter/exporter_test.go`:

```go
package exporter_test

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/MabsIPCA/helm-tests/taxonomy_analyzer/exporter"
	"github.com/MabsIPCA/helm-tests/taxonomy_analyzer/model"
)

func buildReportWithFix() model.TaxonomyReport {
	occ := model.ErrorOccurrence{
		RepoURL:      "https://github.com/test/repo",
		RepoName:     "test/repo",
		ChartPath:    "chart",
		HelmCommand:  "helm template test chart",
		ErrorSource:  "template",
		ErrorKind:    "template",
		ErrorSubKind: "nil_pointer",
		ErrorMessage: "nil pointer evaluating interface {}.type",
		Fixed: &model.FixedResult{
			Resolved:   false,
			StopReason: "unfixable_error_kind",
		},
	}
	bucket := model.TaxonomyBucket{
		Count:    1,
		Examples: []model.ErrorOccurrence{occ},
		FixOutcome: model.FixOutcome{
			Attempted:    1,
			Unresolved:   1,
			ByStopReason: map[string]int{"unfixable_error_kind": 1},
		},
	}
	return model.TaxonomyReport{
		GeneratedAt:   time.Now().UTC(),
		SourceCatalog: "catalog.json",
		FixedCatalog:  "catalog_fixed.json",
		Totals: model.ReportTotals{
			Repos:             1,
			Runs:              1,
			TemplateFailures:  1,
			ClassifiedErrors:  1,
			FixAttempted:      1,
			FixResolved:       0,
			FixUnresolved:     1,
		},
		ByKind:        map[string]model.TaxonomyBucket{"template": bucket},
		BySubKind:     map[string]model.TaxonomyBucket{"template.nil_pointer": bucket},
		ByRepo:        map[string]map[string]int{"test/repo": {"template.nil_pointer": 1}},
		AllClassified: []model.ErrorOccurrence{occ},
	}
}

func buildReportWithoutFix() model.TaxonomyReport {
	occ := model.ErrorOccurrence{
		RepoURL:      "https://github.com/test/repo",
		RepoName:     "test/repo",
		ChartPath:    "chart",
		HelmCommand:  "helm template test chart",
		ErrorSource:  "template",
		ErrorKind:    "template",
		ErrorSubKind: "nil_pointer",
		ErrorMessage: "nil pointer evaluating interface {}.type",
	}
	return model.TaxonomyReport{
		GeneratedAt:   time.Now().UTC(),
		SourceCatalog: "catalog.json",
		Totals: model.ReportTotals{
			Repos: 1, Runs: 1, TemplateFailures: 1, ClassifiedErrors: 1,
		},
		ByKind:        map[string]model.TaxonomyBucket{"template": {Count: 1, Examples: []model.ErrorOccurrence{occ}}},
		BySubKind:     map[string]model.TaxonomyBucket{"template.nil_pointer": {Count: 1, Examples: []model.ErrorOccurrence{occ}}},
		ByRepo:        map[string]map[string]int{"test/repo": {"template.nil_pointer": 1}},
		AllClassified: []model.ErrorOccurrence{occ},
	}
}

func TestWriteMarkdown_IncludesFixStats_WhenFixDataPresent(t *testing.T) {
	dir := t.TempDir()
	if err := exporter.WriteAll(dir, buildReportWithFix()); err != nil {
		t.Fatalf("WriteAll: %v", err)
	}

	data, err := os.ReadFile(filepath.Join(dir, "taxonomy_summary.md"))
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	md := string(data)

	for _, want := range []string{"Fix attempts", "Fix resolved", "Fix unresolved", "Fix Resolved", "Fix Unresolved"} {
		if !strings.Contains(md, want) {
			t.Errorf("markdown missing %q", want)
		}
	}
}

func TestWriteMarkdown_NoFixStats_WhenNoFixData(t *testing.T) {
	dir := t.TempDir()
	if err := exporter.WriteAll(dir, buildReportWithoutFix()); err != nil {
		t.Fatalf("WriteAll: %v", err)
	}

	data, err := os.ReadFile(filepath.Join(dir, "taxonomy_summary.md"))
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	md := string(data)

	if strings.Contains(md, "Fix attempts") {
		t.Error("markdown should not contain 'Fix attempts' when no fix data")
	}
}

func TestWriteCSV_IncludesFixColumns(t *testing.T) {
	dir := t.TempDir()
	if err := exporter.WriteAll(dir, buildReportWithFix()); err != nil {
		t.Fatalf("WriteAll: %v", err)
	}

	f, err := os.Open(filepath.Join(dir, "taxonomy_occurrences.csv"))
	if err != nil {
		t.Fatalf("open csv: %v", err)
	}
	defer f.Close()

	rows, err := csv.NewReader(f).ReadAll()
	if err != nil {
		t.Fatalf("read csv: %v", err)
	}
	if len(rows) < 2 {
		t.Fatalf("expected header + 1 data row, got %d rows", len(rows))
	}

	header := rows[0]
	if header[len(header)-2] != "fix_resolved" {
		t.Errorf("second-to-last column: got %q, want \"fix_resolved\"", header[len(header)-2])
	}
	if header[len(header)-1] != "fix_stop_reason" {
		t.Errorf("last column: got %q, want \"fix_stop_reason\"", header[len(header)-1])
	}

	row := rows[1]
	if row[len(row)-2] != "false" {
		t.Errorf("fix_resolved value: got %q, want \"false\"", row[len(row)-2])
	}
	if row[len(row)-1] != "unfixable_error_kind" {
		t.Errorf("fix_stop_reason value: got %q, want \"unfixable_error_kind\"", row[len(row)-1])
	}
}

func TestWriteCSV_EmptyFixColumns_WhenNoFixData(t *testing.T) {
	dir := t.TempDir()
	if err := exporter.WriteAll(dir, buildReportWithoutFix()); err != nil {
		t.Fatalf("WriteAll: %v", err)
	}

	f, err := os.Open(filepath.Join(dir, "taxonomy_occurrences.csv"))
	if err != nil {
		t.Fatalf("open csv: %v", err)
	}
	defer f.Close()

	rows, err := csv.NewReader(f).ReadAll()
	if err != nil {
		t.Fatalf("read csv: %v", err)
	}
	if len(rows) < 2 {
		t.Fatalf("expected header + 1 data row, got %d rows", len(rows))
	}

	row := rows[1]
	if row[len(row)-2] != "" {
		t.Errorf("fix_resolved (no fix data): got %q, want \"\"", row[len(row)-2])
	}
	if row[len(row)-1] != "" {
		t.Errorf("fix_stop_reason (no fix data): got %q, want \"\"", row[len(row)-1])
	}
}
```

- [ ] **Step 2: Run test to verify it fails**

```
go test ./exporter/...
```

Expected: FAIL — markdown and CSV assertions fail (fix columns/rows not yet present)

- [ ] **Step 3: Update writeMarkdown in exporter.go**

Replace the existing `writeMarkdown` function body's totals table section and SubKind table section. The full updated `writeMarkdown` function:

```go
func writeMarkdown(path string, report model.TaxonomyReport) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create markdown report: %w", err)
	}
	defer f.Close()

	fmt.Fprintln(f, "# Helm Error Taxonomy Report")
	fmt.Fprintln(f)
	fmt.Fprintf(f, "Generated at: `%s`\n\n", report.GeneratedAt.Format("2006-01-02 15:04:05 MST"))
	fmt.Fprintf(f, "Source catalog: `%s`\n\n", report.SourceCatalog)
	if report.FixedCatalog != "" {
		fmt.Fprintf(f, "Fixed catalog: `%s`\n\n", report.FixedCatalog)
	}

	fmt.Fprintln(f, "## Totals")
	fmt.Fprintln(f)
	fmt.Fprintln(f, "| Metric | Value |")
	fmt.Fprintln(f, "|---|---:|")
	fmt.Fprintf(f, "| Repositories | %d |\n", report.Totals.Repos)
	fmt.Fprintf(f, "| Helm runs | %d |\n", report.Totals.Runs)
	fmt.Fprintf(f, "| Template failures | %d |\n", report.Totals.TemplateFailures)
	fmt.Fprintf(f, "| Dependency failures | %d |\n", report.Totals.DependencyErrors)
	fmt.Fprintf(f, "| Classified errors | %d |\n", report.Totals.ClassifiedErrors)
	fmt.Fprintf(f, "| Unclassified errors | %d |\n", report.Totals.UnclassifiedError)
	if report.Totals.FixAttempted > 0 {
		fmt.Fprintf(f, "| Fix attempts | %d |\n", report.Totals.FixAttempted)
		fmt.Fprintf(f, "| Fix resolved | %d |\n", report.Totals.FixResolved)
		fmt.Fprintf(f, "| Fix unresolved | %d |\n", report.Totals.FixUnresolved)
	}
	fmt.Fprintln(f)

	fmt.Fprintln(f, "## Taxonomy by Kind")
	fmt.Fprintln(f)
	fmt.Fprintln(f, "| Kind | Count |")
	fmt.Fprintln(f, "|---|---:|")
	for _, key := range sortedKeysByCount(report.ByKind) {
		fmt.Fprintf(f, "| `%s` | %d |\n", key, report.ByKind[key].Count)
	}
	fmt.Fprintln(f)

	fmt.Fprintln(f, "## Taxonomy by SubKind")
	fmt.Fprintln(f)
	if report.Totals.FixAttempted > 0 {
		fmt.Fprintln(f, "| SubKind | Count | Fix Resolved | Fix Unresolved |")
		fmt.Fprintln(f, "|---|---:|---:|---:|")
		for _, key := range sortedKeysByCount(report.BySubKind) {
			b := report.BySubKind[key]
			fmt.Fprintf(f, "| `%s` | %d | %d | %d |\n", key, b.Count, b.FixOutcome.Resolved, b.FixOutcome.Unresolved)
		}
	} else {
		fmt.Fprintln(f, "| SubKind | Count |")
		fmt.Fprintln(f, "|---|---:|")
		for _, key := range sortedKeysByCount(report.BySubKind) {
			fmt.Fprintf(f, "| `%s` | %d |\n", key, report.BySubKind[key].Count)
		}
	}
	fmt.Fprintln(f)

	if len(report.Unclassified) > 0 {
		fmt.Fprintln(f, "## Unclassified Samples")
		fmt.Fprintln(f)
		limit := 10
		if len(report.Unclassified) < limit {
			limit = len(report.Unclassified)
		}
		for i := 0; i < limit; i++ {
			u := report.Unclassified[i]
			msg := strings.TrimSpace(u.ErrorMessage)
			msg = strings.ReplaceAll(msg, "\n", " ")
			if len(msg) > 180 {
				msg = msg[:180] + "..."
			}
			fmt.Fprintf(f, "- `%s` `%s` `%s`: %s\n", u.RepoName, u.ChartPath, u.ErrorSource, msg)
		}
	}

	return nil
}
```

- [ ] **Step 4: Update writeCSV in exporter.go**

Replace the existing `writeCSV` function:

```go
func writeCSV(path string, occurrences []model.ErrorOccurrence) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create csv report: %w", err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	header := []string{"repo_url", "repo_name", "chart_path", "values_files", "helm_command", "error_source", "error_kind", "error_sub_kind", "error_message", "fix_resolved", "fix_stop_reason"}
	if err := w.Write(header); err != nil {
		return fmt.Errorf("write csv header: %w", err)
	}

	for _, occ := range occurrences {
		fixResolved := ""
		fixStopReason := ""
		if occ.Fixed != nil {
			if occ.Fixed.Resolved {
				fixResolved = "true"
			} else {
				fixResolved = "false"
			}
			fixStopReason = occ.Fixed.StopReason
		}
		if err := w.Write([]string{
			occ.RepoURL,
			occ.RepoName,
			occ.ChartPath,
			strings.Join(occ.ValuesFiles, " | "),
			occ.HelmCommand,
			occ.ErrorSource,
			occ.ErrorKind,
			occ.ErrorSubKind,
			occ.ErrorMessage,
			fixResolved,
			fixStopReason,
		}); err != nil {
			return fmt.Errorf("write csv row: %w", err)
		}
	}

	return nil
}
```

- [ ] **Step 5: Run tests to verify they pass**

```
go test ./exporter/...
```

Expected: `ok  github.com/MabsIPCA/helm-tests/taxonomy_analyzer/exporter`

- [ ] **Step 6: Run all tests**

```
go test ./...
```

Expected: all packages pass.

- [ ] **Step 7: Commit**

```bash
git add exporter/exporter.go exporter/exporter_test.go
git commit -m "feat(taxonomy_analyzer): render fix stats in markdown and CSV outputs"
```

---

## Task 5: Wire --fixed flag in main.go

**Files:**
- Modify: `main.go`

- [ ] **Step 1: Replace main.go**

```go
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/MabsIPCA/helm-tests/taxonomy_analyzer/analyzer"
	"github.com/MabsIPCA/helm-tests/taxonomy_analyzer/exporter"
	"github.com/MabsIPCA/helm-tests/taxonomy_analyzer/loader"
	"github.com/MabsIPCA/helm-tests/taxonomy_analyzer/model"
)

func main() {
	defaultCatalog := filepath.Join("..", "helm_fetcher", "backup", "run_002_274", "catalog_by_project.json")

	catalogPath := flag.String("input", defaultCatalog, "Path to catalog_by_project.json")
	fixedPath := flag.String("fixed", "", "Optional path to catalog_fixed.json")
	outputDir := flag.String("out", "out", "Output directory for taxonomy reports")
	flag.Parse()

	absCatalog, err := filepath.Abs(*catalogPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to resolve input path: %v\n", err)
		os.Exit(1)
	}

	var absFixed string
	var fixedIndex map[string]*model.FixedResult
	if *fixedPath != "" {
		absFixed, err = filepath.Abs(*fixedPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to resolve fixed path: %v\n", err)
			os.Exit(1)
		}
		fixedIndex, err = loader.LoadFixedIndex(absFixed)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to load fixed catalog: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Loaded fixed catalog: %d runs indexed\n", len(fixedIndex))
	}

	collector := analyzer.New(absCatalog, absFixed, fixedIndex)
	err = loader.StreamCatalog(absCatalog, func(repo model.RepoResult) error {
		collector.ConsumeRepo(repo)
		return nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to scan catalog: %v\n", err)
		os.Exit(1)
	}

	report := collector.Report()
	if err := exporter.WriteAll(*outputDir, report); err != nil {
		fmt.Fprintf(os.Stderr, "failed to export report: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Taxonomy analysis complete. Output directory: %s\n", *outputDir)
	fmt.Printf("Classified errors: %d | Unclassified errors: %d\n", report.Totals.ClassifiedErrors, report.Totals.UnclassifiedError)
	if report.Totals.FixAttempted > 0 {
		fmt.Printf("Fix attempts: %d | Resolved: %d | Unresolved: %d\n", report.Totals.FixAttempted, report.Totals.FixResolved, report.Totals.FixUnresolved)
	}
}
```

- [ ] **Step 2: Build to verify compilation**

```
go build ./...
```

Expected: no output.

- [ ] **Step 3: Run all tests**

```
go test ./...
```

Expected: all packages pass.

- [ ] **Step 4: Smoke test without --fixed (existing behavior)**

```
go run . -input ../helm_fetcher/catalog_fixed.json -out out/smoke-no-fixed
```

Expected: runs cleanly, `out/smoke-no-fixed/taxonomy_summary.md` contains no "Fix attempts" row.

- [ ] **Step 5: Smoke test with --fixed**

```
go run . -input ../helm_fetcher/catalog_fixed.json -fixed ../helm_fetcher/catalog_fixed.json -out out/smoke-with-fixed
```

Expected: prints `Loaded fixed catalog: N runs indexed` and `Fix attempts: N | Resolved: N | Unresolved: N`. `out/smoke-with-fixed/taxonomy_summary.md` contains "Fix attempts", "Fix Resolved", "Fix Unresolved" columns. `taxonomy_occurrences.csv` has `fix_resolved` and `fix_stop_reason` columns with values.

- [ ] **Step 6: Commit**

```bash
git add main.go
git commit -m "feat(taxonomy_analyzer): add --fixed flag to load catalog_fixed.json"
```
