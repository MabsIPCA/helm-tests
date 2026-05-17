# taxonomy_analyzer: fix-outcome enrichment

**Date:** 2026-05-17
**Status:** approved

## Goal

Extend `taxonomy_analyzer` to optionally load `catalog_fixed.json` alongside the primary catalog and report, per taxonomy bucket, how many failures were attempted by the fixer, how many were resolved, how many remained unresolved, and the breakdown of stop reasons.

## Chosen approach

Load `catalog_fixed.json` via a new `--fixed` flag. Build a lookup map keyed by `helm_command` → `*FixedResult`. When processing each run from the primary catalog, look up its fix outcome and attach it to the `ErrorOccurrence`. Aggregate fix stats per taxonomy bucket.

The `--fixed` flag is entirely optional. When omitted, the tool behaves identically to today.

## Model changes (`model/model.go`)

### New types

```go
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
```

### Extended types

- `RunResult` gains `Fixed *FixedResult \`json:"fixed,omitempty"\``
- `ErrorOccurrence` gains `Fixed *FixedResult \`json:"fixed,omitempty"\``
- `TaxonomyBucket` gains `FixOutcome FixOutcome \`json:"fix_outcome"\``
- `ReportTotals` gains `FixAttempted int`, `FixResolved int`, `FixUnresolved int`
- `TaxonomyReport` gains `FixedCatalog string \`json:"fixed_catalog,omitempty"\``

## Loader changes (`loader/loader.go`)

New function, existing `StreamCatalog` unchanged:

```go
func LoadFixedIndex(path string) (map[string]*model.FixedResult, error)
```

Streams `catalog_fixed.json` using the same decoder pattern. For each run with a non-nil `fixed` field, inserts `run.HelmCommand → &run.Fixed` into the map. Join key is `helm_command` (a fully-qualified CLI string, unique per run).

## Analyzer changes (`analyzer/analyzer.go`)

```go
type Analyzer struct {
    report     model.TaxonomyReport
    fixedIndex map[string]*model.FixedResult
}

func New(sourceCatalog, fixedCatalog string, fixedIndex map[string]*model.FixedResult) *Analyzer
```

- `New` sets `report.FixedCatalog` when `fixedCatalog != ""`
- In `ConsumeRepo`, for each run: look up `fixedIndex[run.HelmCommand]`, attach to `occ.Fixed`
- In `consumeOccurrence`, after bucket bumps, if `occ.Fixed != nil`:
  - Increment `report.Totals.FixAttempted`; `FixResolved` or `FixUnresolved` by `occ.Fixed.Resolved`
  - Update `bucket.FixOutcome` for both `ByKind` and `BySubKind` buckets (Attempted, Resolved/Unresolved, `ByStopReason[stopReason]++`)

## Exporter changes (`exporter/exporter.go`)

### `taxonomy_summary.md`

- Totals table: add `Fix attempts`, `Fix resolved`, `Fix unresolved` rows when `FixAttempted > 0`
- SubKind table: add `Fix Resolved` and `Fix Unresolved` columns when fix data is present

### `taxonomy_occurrences.csv`

- Two new trailing columns: `fix_resolved` (`true`/`false`/`""`) and `fix_stop_reason` (string or `""`)

### `taxonomy_report.json`

No extra work — `FixOutcome` marshals automatically as part of each bucket.

### `taxonomy_errors_by_bucket.md`

Unchanged.

## `main.go` changes

```go
fixedPath := flag.String("fixed", "", "Optional path to catalog_fixed.json")
```

After flag parsing:
1. If `*fixedPath != ""`: resolve abs path, call `loader.LoadFixedIndex`, pass result to `analyzer.New`
2. If `*fixedPath == ""`: pass `""` and `nil` to `analyzer.New`

## Data flow

```
catalog_by_project.json  →  StreamCatalog  →  ConsumeRepo  ─┐
                                                              ├→  consumeOccurrence  →  TaxonomyBucket.FixOutcome
catalog_fixed.json       →  LoadFixedIndex  →  fixedIndex   ─┘
                             (helm_command key)
```

## Error handling

- Missing or unreadable `--fixed` file: fatal error before analysis starts
- A run in the primary catalog with no matching entry in `fixedIndex`: `occ.Fixed` stays `nil`, no fix stats accumulated for that run — treated as "not yet fixed"
