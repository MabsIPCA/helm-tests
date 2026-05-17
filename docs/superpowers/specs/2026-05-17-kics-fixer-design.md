# KICS Fixer — Design Spec

**Date:** 2026-05-17
**Status:** Approved

## Goal

Extend `kics-render-mock` with an automatic retry-fix loop that, for each failed invocation, parses the render error, injects the minimum values needed to satisfy it, and re-renders — repeating until the chart renders successfully, a cascading loop is detected, or a non-fixable error is hit. Each test directory receives a `kics-fixed-output.json` alongside the existing `kics-render-output.json`, enabling before/after comparison of error resolution rates.

Targeted error kinds (from taxonomy):
1. `template.nil_pointer` — 328 occurrences
2. `template.required_value` — 39 occurrences

---

## Location

No new binary. One new file in the existing `kics-render-mock` package:

```
kics-render-mock/
├── main.go          — updated: prints before/after summary line
├── pipeline.go      — unchanged
├── runner.go        — unchanged
├── output.go        — unchanged
├── fixer.go         — NEW: retry loop, error parsing, default injection
```

Output files land next to the existing output:

```
render_problems/test-02-nil-pointer/kics-render-output.json   ← unchanged
render_problems/test-02-nil-pointer/kics-fixed-output.json    ← new
override_tests_suite/test-01-basic-data-types/kics-render-output.json
override_tests_suite/test-01-basic-data-types/kics-fixed-output.json
```

---

## Invocation

`main.go` calls the fixer automatically after each test dir completes — no flags required. The KICS-faithful render path is not touched.

---

## Retry Loop (`fixer.go`)

```
patch     := map[string]string{}   // grows each iteration; keys are Helm --set paths
seenErrs  := map[string]bool{}     // loop guard: same error string twice = stop
maxIter   := 10

for attempt := 1..maxIter:
    build patchedValOpts = original valOpts + patch (appended as extra Values entries)
    result = runOnce(chartPath, patchedValOpts, debugMode=false)

    if result.err == nil:
        resolved = true; break

    errStr = result.err.Error()
    if seenErrs[errStr]:
        stopReason = "loop_detected"; break
    seenErrs[errStr] = true

    fix, ok = parseError(errStr)
    if !ok:
        stopReason = "unfixable_error_kind"; break

    record FixStep{attempt, errStr, fix.kind, fix.path, fix.value}
    patch[fix.path] = fix.value
```

The patch is merged via `values.Options{Values: []string{"a.b=val", "x[0].y=val"}}` appended to the original invocation's `Values`. Helm's own precedence rules apply — the patch wins over chart defaults and values files, same as `--set`.

---

## Error Parsing

### `nil_pointer`

**Error pattern:**
```
nil pointer evaluating interface {}.someKey
```

Full error from Helm looks like:
```
template: chartname/templates/foo.yaml:12:8: executing "..." at <.Values.foo.bar>: nil pointer evaluating interface {}.bar
```

**Extraction:**
1. Regex match `at <(\.Values\.[^>]+)>` to get the full dotted path
2. Strip leading `.Values.` → Helm `--set` path
3. Inject `path=""` (empty string satisfies nil map access)

### `required_value`

**Error pattern:**
```
error calling required: <user message>
```

Full error from Helm looks like:
```
template: chartname/templates/foo.yaml:5:15: executing "..." at <required .Values.db.host "db.host is required">: error calling required: db.host is required
```

**Extraction:**
1. Regex match `required (\.Values\.[^ ]+)` to get the dotted path
2. Strip leading `.Values.` → Helm `--set` path
3. Inject a type-inferred default (see table below)

### Default Value Table

| Error kind | Value injected |
|---|---|
| `nil_pointer` | `""` (empty string — satisfies nil map access) |
| `required_value` | `"kics-placeholder"` (non-empty — Helm's `required` rejects `""`) |

Helm's `required` function fails on both nil *and* empty string, so `required_value` fixes must inject a non-empty string. `nil_pointer` fixes only need a non-nil value, so `""` is sufficient. No richer type inference is attempted — the goal is a valid partial render, not a semantically correct one.

### Unfixable — stop immediately

Any error that does not match either pattern above. Examples: parse errors, schema validation failures, dependency errors. Recorded as `stopReason: "unfixable_error_kind"`.

---

## JSON Schema — `kics-fixed-output.json`

Same top-level shape as `kics-render-output.json`. Each render entry gains three extra fields:

```json
{
  "suite":      "render_problems",
  "testNumber": 2,
  "testName":   "Nil Pointer Dereference Scenarios",
  "chartPath":  "C:\\...\\test-02-nil-pointer",
  "renders": [
    {
      "toggle":    "testNilMapAccess",
      "dataType":  "nilMap",
      "caseName":  "",
      "valuesOptions": { ... },
      "standard":  { ... },
      "debug":     { ... },

      "resolved":   true,
      "stopReason": "",
      "fixChain": [
        {
          "attempt":       1,
          "errorSeen":     "nil pointer evaluating interface {}.hosts",
          "kind":          "nil_pointer",
          "valuePath":     "ingress.hosts",
          "valueInjected": ""
        },
        {
          "attempt":       2,
          "errorSeen":     "error calling required: db.host is required",
          "kind":          "required_value",
          "valuePath":     "db.host",
          "valueInjected": "kics-placeholder"
        }
      ]
    }
  ]
}
```

**Field rules:**
- `resolved` — `true` if the final retry attempt returned no error
- `stopReason` — `""` on success; `"loop_detected"` or `"unfixable_error_kind"` or `"max_iterations"` on failure
- `fixChain` — ordered list of fix steps applied; empty slice `[]` if the original render already succeeded (no fix needed)
- Invocations that were already passing in `kics-render-output.json` are included with `resolved: true`, `fixChain: []`

---

## Before/After Summary (stdout)

Appended to the existing summary block:

```
=== KICS Render Mock ===
render_problems      21 dirs  126 invocations  89 errors → 34 resolved (38%)
override_tests_suite 10 dirs   42 invocations   0 errors → 0 resolved
JSON + fixed files written to each test directory.
```

The "resolved" count includes only invocations that were failing in the base run and are now passing after the fix loop.

---

## Error Handling

- If `fixer.go` fails to write `kics-fixed-output.json`, print to stderr and continue — same policy as the base runner.
- If an invocation was already passing in the base run, copy its result into the fixed output with `resolved: true`, `fixChain: []` — no retry loop is run.
- The fix loop always runs in non-debug mode (`runOnce(..., false)`) to keep error messages clean and consistent.

---

## Out of Scope

- Fixing `dependency.*` errors — these require network access or local chart updates, not values injection.
- Richer type inference (e.g. detecting that a value should be an integer and injecting `0`) — empty string is sufficient for the research goal.
- Modifying chart templates — the fixer only touches values, never template source files.
- Persisting the patch as an actual values file on disk — the patch lives in memory as `--set` flags.
