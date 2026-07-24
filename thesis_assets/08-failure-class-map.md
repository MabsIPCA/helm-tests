# A3 — `tab:ch4-taxonomy` per-class rows (chap4:274-292)

Answers the last open item in `README.md`'s "Still open" line: the six-class split of
`tab:ch4-taxonomy` in `chap4_experimentalEvaluation.tex:274-292`. The Total row (1,211 / 723 /
59.7%) was already set and guarded; only the six `\myworries{?}` class rows were outstanding.

## The task, precisely

The table's own caption (`chap4:290`) says how to fill it in:

> map the sub-kinds of Table `tab:ch5-fix-results` onto these six classes via Table
> `tab:ch4-failure-class-map` and aggregate.

Two tables, two different taxonomies:

- **`tab:ch5-fix-results`** (`chap3-3`, reproduced in `02-taxonomy.md`) — the 16 real sub-kinds
  `taxonomy_analyzer` assigns to production-corpus failure *messages*, by regex pattern
  (`taxonomy_analyzer/classifier/classifier.go`). Grounded in `out/cumulative_v2/taxonomy_report.json`.
- **`tab:ch4-failure-class-map`** (`chap3-2:203-261`) — the 21 *synthetic* `render_problems` test
  groups, each tagged with one of the six classes and a Helm/sprig function family (e.g. "Logic and
  default functions → Type-system mismatch").

These don't share a key. The class-map is organized by *which sprig function family* was
exercised in an isolated test chart; the real corpus's sub-kinds are organized by *which regex
matched the error text* (author guardrails, schema validation, dependency pre-checks, chart
metadata, etc. — none of which the synthetic suite tests as a function family). So there is no
1:1 name lookup; each of the 16 sub-kinds had to be placed under the single best-fit class using
the class-map's entries plus the class definitions in `chap3-2` §5 (`sec:ch4-parse-time`,
`sec:ch4-nil-failures`).

## Source data

```sh
cd <helm-tests>/taxonomy_analyzer/out/cumulative_v2
# per-subkind count / resolved / unresolved
python3 -c "import json; d=json.load(open('taxonomy_report.json')); [print(k,v['count'],(v.get('fix_outcome') or {}).get('resolved',0)) for k,v in d['by_sub_kind'].items() if k.startswith('template.')]"
```

Same basis as `02-taxonomy.md`'s corrected `tab:ch5-fix-results`: `catalog_sources_v2.json` +
`catalog_fixed_cumulative_v2.json`, analyzed by `taxonomy_analyzer -out out/cumulative_v2`.

## Per-sub-kind assignment

| Sub-kind | Count | Resolved | → Class | Why |
|---|---:|---:|---|---|
| `nil_pointer` | 587 | 540 | Nil dereference | Direct match — `chap3-2:297` defines nil dereference as accessing a field/index/result on a nil value; the classifier rule is literally `"nil pointer evaluating"`. |
| `unsupported_builtin` | 8 | 0 | Nil dereference | Same error signature (`"nil pointer"`), just on a Helm built-in (`.Release.Time`, `.Capabilities`, …) instead of a user value — non-fixable by injection, but the underlying language-level fault is the same nil dereference. |
| `required_value` | 263 | 159 | Type-system mismatch | Class-map row 01: "Logic and default functions (`and`, `or`, `default`, `coalesce`) → Type-system mismatch." Helm's `required` builtin is the same function family — a logic/validation call that rejects an absent/invalid value, structurally identical to `default`'s handling of a missing value. |
| `values_schema_validation` | 92 | 0 | Type-system mismatch | A `values.schema.json` violation is, by definition, a type/shape check — the class's own name. Ties directly to `chap4:300`'s explanation of the recovery gap: "supplying a placeholder of the wrong type does not satisfy the template's expectations." |
| `custom_validation` | 20 | 0 | Type-system mismatch | Same logic/validation-guard family as `required_value` (row 01) — a chart author's own hand-written guard clause. |
| `author_assertion` | 21 | 0 | Type-system mismatch | Same family again — an explicit `fail`/`required` guardrail, matching row 01. |
| `type_mismatch` | 15 | 0 | Type-system mismatch | Named match — classifier rule is `"wrong type for value"` / `"expected string"`. |
| `invalid_value` | 2 | 0 | Type-system mismatch | Classifier's own description: "value rejected as structurally invalid" — a shape/type rejection. |
| `runtime_eval` | 21 | 0 | Type-system mismatch | Inspected all 21 raw messages (`taxonomy_occurrences.csv`): the majority (~17/21) are field-access/logic type errors during execution — `"can't evaluate field X in type interface{}"`, boolean-pipeline errors on `and`/`or` — the same row-01 family. A minority (4/21) are `{{template "x"}}`-not-defined indirection misses that fell through `missing_template`'s narrower pattern; documented here as the exception, not split out (see "Judgment calls" below). |
| `missing_template` | 110 | 0 | Structural, parse-time | Class-map row 04, verbatim: "Missing template reference (`include`, `template`) → Structural / parse-time." |
| `malformed_yaml` | 23 | 0 | Structural, parse-time | Class-map row 03 family — the chart fails to load/parse at all. |
| `parse_error` | 9 | 0 | Structural, parse-time | Same family — literal Go-template parse errors, row 03's "Go template syntax." |
| `chart_metadata` | 4 | 0 | Structural, parse-time | Invalid `Chart.yaml` metadata (`apiVersion`/`version`) — a packaging defect caught before rendering starts, same "detectable through parsing, without executing any rendering logic" bucket (`chap3-2:269`). |
| `library_chart_not_installable` | 11 | 0 | Structural, parse-time | A chart-type precondition (`type: library`) checked before rendering ever begins — same pre-render gate family. |
| `dependency_check_failed` | 1 | 0 | Structural, parse-time | A dependency pre-check gate inside the template, blocking before value evaluation — matches `chap3-2:507-508`'s framing of dependency-lock drift as "a pure file-comparison operation available offline." |
| `kube_version_incompatible` | 24 | 24 | Structural, parse-time | See "Judgment calls" — placed here, not Type-system mismatch. |

**Reconciliation**: 595 (nil deref) + 434 (type-system) + 182 (structural) = 1,211. Resolved:
540 + 159 + 24 = 723. Both match the guarded Total row exactly.

## Judgment calls (the two sub-kinds without an obvious class-map row)

**`kube_version_incompatible` → Structural, parse-time, not Type-system mismatch.** It isn't a
value-type error at all: it's Helm comparing the chart's declared `kubeVersion` SemVer constraint
against the target cluster version, checked during chart *loading* — before the values merge or
any template executes. That is the same pre-render, packaging-level gate as `chart_metadata`,
`library_chart_not_installable`, and `dependency_check_failed`, so it was grouped with them rather
than with the value-type-mismatch family. This is also why it recovers cleanly (24/24, injected via
`--kube-version` rather than `--set`) while every other sub-kind in Structural/parse-time is 0% —
worth calling out explicitly if this table is discussed in prose, since it makes that row's 13.2%
entirely carried by one sub-kind.

**`runtime_eval` → Type-system mismatch, not split by message.** The caption's instruction is to
map *sub-kinds* onto classes, i.e. keep the aggregation at the granularity `tab:ch5-fix-results`
already uses — not re-classify individual error messages, which would be a different (and
undocumented-elsewhere) methodology. Inspected anyway for transparency (see table above): of 21
raw messages, ~17 are type/logic execution errors and 4 are template-indirection misses. Kept whole
under the majority class rather than fabricating a fractional split.

## Why three classes are zero

Searched all 1,211 classified `template.*` messages in `taxonomy_occurrences.csv` for each class's
defining symptom:

```sh
python3 - <<'PY'
import csv, re
rows = [r for r in csv.DictReader(open("taxonomy_occurrences.csv", encoding="utf-8")) if r["error_kind"]=="template"]
def scan(pat): return sum(1 for r in rows if re.search(pat, r["error_message"], re.I))
print("divide/division by zero:", scan(r"divide by zero|division by zero"))
print("genuine `tpl` invocation:", scan(r"at <\(?tpl[\s(]"))
print("uuid/rand/now function errors:", scan(r"uuidv4|randalpha|randnumeric|genprivatekey"))
PY
```

- **Division by zero: 0.** No message in the corpus mentions a divide/modulo-by-zero panic.
`runtime_eval` — the only sub-kind that could plausibly hide one — was inspected message-by-message
above and contains none.
- **Indirection propagation: 0** (as its own bucket). Only 3 messages anywhere in the corpus are a
genuine dynamic `tpl` function call failing, and all 3 already carry a more specific sub-kind
(`nil_pointer` ×1, `required_value` ×2) rather than existing as an isolated indirection failure —
there is no sub-kind in `tab:ch5-fix-results` dedicated to `tpl`/`include` propagation distinct from
`missing_template` (which is itself classified as Structural/parse-time per the class-map, not
Indirection propagation).
- **Non-deterministic output: 0.** By construction this can't appear in a *failure* taxonomy:
`uuidv4`, `randAlphaNum`, `now` and similar functions don't error, they render successfully with
different output each time — so a render that fails is never classified into this bucket. Zero
occurrences here is expected, not a gap in the classifier.

This means the real corpus's 1,211 failures land in only three of the six classes. `chap4:271`
currently guesses ("the remaining four classes forming the tail") that all four non-dominant
classes have some presence — this data shows three of them are exactly zero, not merely small.
That line should be corrected alongside filling the table.

## Result — ready to paste into `tab:ch4-taxonomy` (chap4:280-285)

| Failure class | Count | Recovered | Recovery rate |
|---|---:|---:|---:|
| Type-system mismatch | 434 | 159 | 36.6% |
| Nil dereference | 595 | 540 | 90.8% |
| Structural, parse-time | 182 | 24 | 13.2% |
| Division by zero | 0 | 0 | n/a |
| Indirection propagation | 0 | 0 | n/a |
| Non-deterministic output | 0 | 0 | n/a |
| **Total** | **1,211** | **723** | **59.7%** |

```latex
Type-system mismatch & 434 & 159 & 36.6\% \\
Nil dereference & 595 & 540 & 90.8\% \\
Structural, parse-time & 182 & 24 & 13.2\% \\
Division by zero & 0 & 0 & n/a \\
Indirection propagation & 0 & 0 & n/a \\
Non-deterministic output & 0 & 0 & n/a \\
```

The existing prose already gets this right by luck: `chap4:298-301` says type-system mismatch is
"common yet resists recovery" — true here (434 count, only 36.6%, the worst rate of the three
non-zero classes) — and that prevalence and recoverability are otherwise aligned — also true:
Nil dereference is both the largest class (595) and the best-recovered (90.8%). Only the
"remaining four classes forming the tail" aside (`chap4:271`) needs correcting, per above.

**Not applied to the `.tex` file** — this documents the derivation and the fill-in values only;
the thesis table itself (`chap4_experimentalEvaluation.tex:280-285`) still has to be edited by hand
or on request.
