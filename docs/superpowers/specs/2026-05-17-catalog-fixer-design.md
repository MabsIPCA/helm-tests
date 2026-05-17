# Catalog Fixer — Design Spec

**Date:** 2026-05-17
**Status:** Approved

## Goal

Add a `fixer` mode to `helm_fetcher` that reads `catalog_kept.json`, clones any missing repos, and runs a retry-fix loop on every failing `helm template` run. For each fixable error (`nil_pointer`, `required_value`), the loop injects `--set` flags and re-runs until success, a cascading loop is detected, or a non-fixable error is hit. Results are written to `catalog_fixed.json` and `catalog_fixer_report.md`, giving a before/after view of the taxonomy for the thesis.

Targeted error kinds (top 2 from taxonomy):
1. `template.nil_pointer` — 328 occurrences
2. `template.required_value` — 39 occurrences

---

## Location

No new module or binary. Everything extends the existing `helm_fetcher/` package.

```
helm_fetcher/
├── main.go          — MODIFY: add "fixer" case to mode switch
├── Makefile         — MODIFY: add run-fixer target
├── helm/helm.go     — MODIFY: add RunHelmTemplateWithSets
├── model/model.go   — MODIFY: add FixStep, FixedRunResult, FixedRepoResult
└── fixer_mode.go    — NEW: runFixerMode, fix loop, parseHelmCLIError, report writers
```

---

## Invocation

```
make run-fixer CATALOG_IN=backup/run_002_274/catalog_kept.json CLONE_DIR=cloned
```

New Makefile target:

```makefile
CATALOG_IN ?= catalog_kept.json

run-fixer:
	go run . \
		-mode=fixer \
		-catalog-in=$(CATALOG_IN) \
		-clone-dir=$(CLONE_DIR)
```

---

## Helm Extension (`helm/helm.go`)

New function alongside existing `RunHelmTemplate`:

```go
func RunHelmTemplateWithSets(chartPath string, valFiles, setFlags []string) (cmdStr, output string, err error)
```

Builds: `helm template test <chartPath> [-f vf...] [--set k=v...]`  
Returns the same `(cmdStr, output, error)` triple as `RunHelmTemplate`.

---

## New Model Types (`model/model.go`)

```go
type FixStep struct {
    Attempt       int    `json:"attempt"`
    ErrorSeen     string `json:"error_seen"`
    Kind          string `json:"kind"`          // "nil_pointer" | "required_value"
    ValuePath     string `json:"value_path"`
    ValueInjected string `json:"value_injected"`
}

type FixedRunResult struct {
    Resolved     bool      `json:"resolved"`
    StopReason   string    `json:"stop_reason"` // "" | "loop_detected" | "unfixable_error_kind" | "max_iterations"
    FixChain     []FixStep `json:"fix_chain"`
    FinalCommand string    `json:"final_command,omitempty"` // helm command that succeeded
}

// RunResultWithFix extends RunResult with the fixer output.
// Used only in catalog_fixed.json — does not replace RunResult in existing files.
type RunResultWithFix struct {
    RunResult
    Fixed *FixedRunResult `json:"fixed,omitempty"`
}

type ChartSummaryFixed struct {
    ChartPath       string             `json:"chart_path"`
    TotalRuns       int                `json:"total_runs"`
    Successes       int                `json:"successes"`
    Failures        int                `json:"failures"`
    DepBuildFailure bool               `json:"dep_build_failure,omitempty"`
    DepBuildError   string             `json:"dep_build_error,omitempty"`
    Runs            []RunResultWithFix `json:"runs"`
}

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

---

## Fix Loop (`fixer_mode.go`)

### `parseHelmCLIError(errMsg string) (kind, valuePath, injectedValue string, ok bool)`

Strips the `"Error: "` CLI prefix, then applies the same two regexes from `kics-render-mock`:

- `nilPtrRe  = regexp.MustCompile(`at <(\.Values\.[^ >|]+)>`)` — matched when error contains `"nil pointer"`
  - Injects `value_injected: ""` (empty string unblocks nil map access)
- `requiredRe = regexp.MustCompile(`required (\.Values\.[^ "]+)`)` — matched when error contains `"error calling required"`
  - Injects `value_injected: "kics-placeholder"` (non-empty passes Helm's `required` check)
- Any other error → `ok = false` (unfixable)

### `fixRun(chartPath string, orig RunResult) FixedRunResult`

```
patch     := map[string]string{}
seenErrs  := map[string]bool{}

for attempt := 1..10:
    setFlags := patch entries formatted as "key=value"
    _, output, err = RunHelmTemplateWithSets(chartPath, orig.ValuesFiles, setFlags)

    if err == nil:
        return FixedRunResult{Resolved: true, FinalCommand: cmdStr, FixChain: chain}

    errStr = output (combined stderr+stdout from CLI)
    if seenErrs[errStr]:
        return FixedRunResult{StopReason: "loop_detected", FixChain: chain}
    seenErrs[errStr] = true

    kind, path, val, ok = parseHelmCLIError(errStr)
    if !ok:
        return FixedRunResult{StopReason: "unfixable_error_kind", FixChain: chain}

    chain = append(chain, FixStep{attempt, errStr, kind, path, val})
    patch[path] = val

return FixedRunResult{StopReason: "max_iterations", FixChain: chain}
```

Already-successful runs (`orig.Success == true`) are returned immediately as `FixedRunResult{Resolved: true, FixChain: []}` — no fix loop runs.

### `runFixerMode(catalogIn, cloneDir string)`

1. Read and decode `catalogIn` → `[]model.RepoResult`
2. For each repo: call `git.CloneRepo(repoURL, clonedDir)` if dir absent
3. For each chart and each run: call `fixRun`, build `RunResultWithFix`
4. Write `catalog_fixed.json`
5. Write `catalog_fixer_report.md`
6. Print summary to stdout

---

## Output Files

### `catalog_fixed.json`

Array of `RepoResultFixed`. Each run entry:

```json
{
  "chart_path": "cloned/ayarinesrine__pfe-deployment/helm-charts/backend/authentification",
  "helm_command": "helm template test ...",
  "success": false,
  "error_message": "Error: template: ... nil pointer evaluating interface {}.type",
  "fixed": {
    "resolved": true,
    "stop_reason": "",
    "fix_chain": [
      {
        "attempt": 1,
        "kind": "nil_pointer",
        "value_path": "service.type",
        "value_injected": "",
        "error_seen": "Error: template: ..."
      }
    ],
    "final_command": "helm template test ... --set service.type="
  }
}
```

Already-passing runs have `"fixed": null` (omitted via `omitempty`).

### `catalog_fixer_report.md`

```markdown
# Catalog Fixer Report

**Source:** backup/run_002_274/catalog_kept.json
**Date:** YYYY-MM-DD

## Summary

| Metric | Count |
|---|---:|
| Repos processed | 274 |
| Failing runs (before) | N |
| Resolved | N |
| Resolution rate | N% |

## By Error Kind

| Kind | Before | Resolved | Still failing |
|---|---:|---:|---:|
| nil_pointer | 328 | N | N |
| required_value | 39 | N | N |
| other | N | 0 | N |

## Stop Reasons

| Reason | Count |
|---|---:|
| resolved | N |
| unfixable_error_kind | N |
| loop_detected | N |
| max_iterations | N |
```

---

## Error Handling

- Clone failure → log warning, skip repo, continue
- Chart not found after clone → skip repo
- `RunHelmTemplateWithSets` error unrelated to template rendering (e.g. binary not found) → log error, record as unfixable
- `catalog_fixed.json` write failure → fatal (no partial write)
- `catalog_fixer_report.md` write failure → log error, continue (JSON is more important)

---

## Config additions (`config/config.go`)

New flag `-catalog-in string` (default `"catalog_kept.json"`) added to `Config` struct and registered in `Parse()`. All flags are always registered (Go's `flag` package); `-catalog-in` is only *used* in `fixer` mode. `-clone-dir` already exists and is reused.

## Clone path resolution

The `cloned_dir` field in catalog entries is a snapshot from the original run and may be stale. The fixer derives the destination directory from `repo_url` the same way `main.go` does:

```go
repoName := strings.TrimPrefix(repoURL, "https://github.com/")
safeName := strings.ReplaceAll(repoName, "/", "__")
destDir  := filepath.Join(cloneDir, safeName)
```

`git.CloneRepo(repoURL, destDir)` is called; it is a no-op if `destDir` already exists. Chart paths inside each repo are discovered fresh via `helm.FindCharts(destDir)` rather than trusting the catalog's `chart_path` values.

---

## Out of Scope

- Re-running the original helm_fetcher pipeline (only reads existing catalog)
- Fixing `dependency.*` errors — requires network; not attempted
- Parallel cloning — sequential, same as existing helm_fetcher behavior
- Modifying chart templates — only `--set` flags are injected
