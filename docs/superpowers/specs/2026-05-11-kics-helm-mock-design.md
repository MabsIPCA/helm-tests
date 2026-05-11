# KICS Helm Render Mock — Design Spec

**Date:** 2026-05-11
**Status:** Approved

## Goal

Build a standalone Go tool (`kics-render-mock`) that mimics the KICS engine's internal Helm rendering pipeline and runs it against every toggle/case combination in both test suites (`render_problems` and `override_tests_suite`). Each test directory receives a `kics-render-output.json` file capturing both a standard and a debug render run, enabling offline analysis of what errors KICS would encounter per chart and per toggle.

---

## Location

```
helm-tests/
└── kics-render-mock/
    ├── main.go       — entry point: discovers both suites, prints summary
    ├── pipeline.go   — KICS-faithful Helm render pipeline
    ├── runner.go     — iterates test dirs, drives pipeline per toggle/case
    ├── output.go     — JSON structs + file writer
    └── go.mod        — depends on helm.sh/helm/v3 only
```

Output files land inside each existing test directory (next to `TEST-RESULTS.md`):

```
render_problems/test-02-nil-pointer/kics-render-output.json
override_tests_suite/test-01-basic-data-types/kics-render-output.json
...
```

---

## Pipeline (`pipeline.go`)

Replicates the KICS `pkg/resolver/helm` pipeline exactly, using the same Helm Go library (`helm.sh/helm/v3`):

1. `newClient()` — creates `action.Install` with `DryRun=true`, `ClientOnly=true`, `Replace=true`, `IncludeCRDs=false`, `Namespace="kics-namespace"`
2. `loader.Load(chartPath)` — loads the chart from disk
3. `setID(chart)` — injects `# KICS_HELM_ID_<n>:` markers before each `apiVersion:` line (identical to KICS)
4. `client.Run(chart, vals)` — renders the chart; **always inspect both the return value and the error** — a non-nil `*release.Release` alongside a non-nil `error` contains the partial manifest rendered before the failure
5. `splitManifestYAML(release)` — splits the rendered manifest by `---` and maps each section back to its source template file

Each invocation runs **twice**:
- **Standard run:** `settings.Debug = false`, `cfg.Log = nil`
- **Debug run:** `settings.Debug = true`, `cfg.Log` wired to a `[]string` collector

The two clients are constructed independently per run to prevent state bleed.

---

## Toggle Iteration (`runner.go`)

Two config formats are supported, detected by which top-level key is present in `config.json`:

### `render_problems` (has `testFunctions` + `dataTypes`)

For each `(testFunction, dataType)` pair, one render invocation:
```go
values.Options{
    Values: []string{
        fmt.Sprintf("%s=true", tf.Name),
        fmt.Sprintf("testDataType=%s", dataType),
    },
}
```

### `override_tests_suite` (has `testCases`)

For each `testCase`, one render invocation:
```go
values.Options{
    ValueFiles:   absoluteValueFiles, // tc.ValuesFiles joined with testDir to make absolute paths
    Values:       tc.SetFlags,
    StringValues: tc.SetStringFlags,
}
```

---

## JSON Schema

One `kics-render-output.json` per test directory:

```json
{
  "suite":       "render_problems",
  "testNumber":  2,
  "testName":    "Nil Pointer Dereference Scenarios",
  "chartPath":   "C:\\...\\test-02-nil-pointer",
  "renders": [
    {
      "toggle":    "testNilMapAccess",
      "dataType":  "nilMap",
      "caseName":  "",
      "valuesOptions": {
        "values":     ["testNilMapAccess=true", "testDataType=nilMap"],
        "valueFiles": []
      },
      "standard": {
        "error":          "template: ...: nil pointer evaluating interface {}.someKey",
        "rawManifest":    null,
        "partialManifest": null,
        "splitManifests": []
      },
      "debug": {
        "error":          "template: ...: nil pointer evaluating interface {}.someKey",
        "debugLogs":      ["[debug] CHART PATH: ...", "[debug] rendering manifest failed: ..."],
        "rawManifest":    null,
        "partialManifest": "---\n# Source: .../nil-pointer-test.yaml\napiVersion: v1\n...",
        "splitManifests": []
      }
    }
  ]
}
```

**Field rules:**
- `toggle` / `dataType` — non-empty for `render_problems` runs
- `caseName` — non-empty for `override_tests_suite` runs
- `error` — string on failure, `null` on success
- `rawManifest` — full manifest string on success, `null` on failure
- `partialManifest` — non-null only when `error != null` AND Helm still returned a partial `release.Manifest`
- `splitManifests` — populated on success only; each entry: `{ path, content, originalContent, splitID }`
- `debugLogs` — present in `debug` block only; empty slice if Helm logged nothing

---

## SplitManifest entry

```json
{
  "path":            "test-nil-pointer/templates/nil-pointer-test.yaml",
  "content":         "---\n# Source: ...\napiVersion: v1\nkind: ConfigMap\n...",
  "originalContent": "apiVersion: v1\nkind: ConfigMap\n...",
  "splitID":         "# KICS_HELM_ID_0:"
}
```

`originalContent` is the pre-render template source with KICS ID markers injected (same as what KICS stores in `splitManifest.original`).

---

## Error Handling

- If `config.json` is missing or malformed for a test dir, skip it and log a warning to stdout; do not abort the whole run.
- If both `standard` and `debug` runs error, still write the JSON — the error strings are the data.
- If writing `kics-render-output.json` fails, print to stderr and continue to the next test.

---

## Summary Output (stdout)

```
=== KICS Render Mock ===
render_problems         21 dirs   126 invocations   89 errors
override_tests_suite    10 dirs    42 invocations    0 errors
JSON files written to each test directory.
```

---

## Out of Scope

- The experimental multi-values-file combination logic from `renderHelm` (power-set combinations) — not replicated; each test case is driven by `config.json` only.
- Comparing results against expected values — this tool is data collection only.
- CI integration — run manually.
