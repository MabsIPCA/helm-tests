# Thesis Assets — where every number comes from

Companion to `thesis-todos.md`. That file lists the doubts; this folder answers them and
records the provenance of every figure, so no number in the thesis is unsourced.

**Rule:** a number is only allowed in the thesis if it appears in this folder with (a) the
artifact it was read from and (b) the command that regenerates it.

## Contents

| File | Answers |
|---|---|
| `01-corpus.md` | Ch. 3.2 corpus figures — 973/72,225/10,110/2,446/75.8% — and 3 prose errors found |
| `02-taxonomy.md` | Ch. 3.3 taxonomy + recovery table. **The published numbers are on the wrong corpus** |
| `03-todo-verdicts.md` | Verdict on each `thesis-todos.md` item (D1, D2, group C). Two are false alarms |
| `04-helmgoat-misconfig.md` | §4.4.1 HELM GOAT baseline-vs-enhanced KICS (A1/A2), native builds only: 6→322 findings, 1→9 categories, +316 gain |
| `05-helmgoat-image-cve.md` | §4.4.2 HELM GOAT image CVEs, full `hcs` run (A4): 14 images, 4,155 CVEs, 322 misconfigs, 4,477 unified, 15 runs. Resolves the old "189" TEMP value |
| `06-engine-crosswalk.md` | §4.4.3(?) / A5 capability crosswalk — Kubescape vs Trivy vs kube-linter: misconfig + single-command CVE scanning, render-failure handling, and dep-up on `render-problems/madgoat-render`. Draft `tab:ch4-crosswalk` content. **Not yet applied to the thesis text** |
| `07-render-problems-engines.md` | §4.4.x(?) / A5b — five engines (KICS default + KICS enhanced + Kubescape + Trivy + kube-linter) on `render-problems/madgoat-render` at committed defaults, **forcing dep-up** (unvendored `file://` subchart + render errors on). Headline: **only KICS-enhanced clears both hurdles** — its fixer repairs the render (value injection) *and* vendors the missing dependency, reaching **322** (139 parent + 183 subchart) and mutating disk; Kubescape dep-ups but wastes it behind the unfixed render (1 resource, score 94); Trivy/kube-linter produce 0. Confirms `04`'s 322 is achievable without pre-vendoring. Also carries an **OWASP Kubernetes Top 10 coverage** table per engine (via `scripts/scan_coverage.py`): on `madgoat-render` only KICS-enhanced detects any category (K01/K02/K05/K09); on the vendored `helm/madgoat` the union reaches K01–K03/K05/K09. **Not yet applied to the thesis text** |
| `08-failure-class-map.md` | A3 — `tab:ch4-taxonomy` (chap4:274-292) per-class split. Maps the 16 real sub-kinds of `tab:ch5-fix-results` onto the six classes of `tab:ch4-failure-class-map` (chap3-2). Result: Type-system mismatch 434/159/36.6%, Nil dereference 595/540/90.8%, Structural-parse-time 182/24/13.2%, and **Division by zero / Indirection propagation / Non-deterministic output are all 0** in the real corpus (verified by keyword search, not assumed). Ties out exactly to the guarded 1,211/723/59.7% total. **Not yet applied to the thesis text** |
| `09-kics-corpus-baseline-enhanced.md` | §4.4.1(?) / A7 — Baseline vs Enhanced KICS at corpus scale: every chart in the **top-100 GitHub + top-100 ArtifactHub** repos (9,861 charts, 154 repos). **Charts analysed 9,861 / 9,861 · findings 204,086 → 508,248 · net gain +304,162 (+149%, ≈2.5×)**; charts-with-findings 5,838 → 7,337. Native-Windows KICS builds (not poolable with WSL numbers); ~49% is rancher versioned near-duplicates (not deduped). Scales `04`'s single-chart 6→322. **Not yet applied to the thesis text** |
| `scripts/scan_coverage.py` | Runs the 5 engines over one source, counts each run's findings, and maps fired rule ids → OWASP Kubernetes Top 10 (`K01`–`K10`) using `scripts/vulnerabilities.yaml` (synced to `tab:scanner-coverage`). Emits text, `--latex` (`tab:scanner-owasp-coverage`), or `--json`; non-destructive |
| `scripts/batch_scan.py` | Batch-runs the engines over a corpus (chart-list from `helm_fetcher/cmd/listcharts`), resumable, one JSONL record per chart; `agg` builds the engine-comparison + OWASP tables. Used for `09` and the `scan_out/corpus-batch/` 5-engine run |
| `scripts/verify_corpus.py` | Recomputes every Ch. 3.2 figure from the catalogs and asserts it |

## Status

| Group | State |
|---|---|
| Ch. 3.2 corpus | ✅ **Set in the thesis** on the subchart-cleaned basis (Correction 4 applied) |
| Ch. 3.2 prose | ✅ Corrections 1–4 applied; the taxonomy sentences carry a new `\myworries` |
| Ch. 3.3 taxonomy | ✅ **Re-run complete and set.** `tab:ch5-fix-results` rewritten |
| D1 / D2 | ✅ **Resolved, both false alarms** — and now moot: the totals coincide on the correct basis |
| D3 | ✅ **Answered and written up** as a sampling limitation, with a `\myworries` |
| A5 crosswalk | 🟡 **Measured, not applied** — `06-engine-crosswalk.md` has the numbers and draft table; still needs a `.tex` home and a confirmed section number |
| A3 class map | 🟡 **Measured, not applied** — `08-failure-class-map.md` has the six-class split and the full sub-kind mapping; `chap4:280-285` still has the `\myworries{?}` placeholders |
| Group C "REAL" | ✅ **All five values replaced and unwrapped** |

### Applied to the thesis 2026-07-17

| File | Change |
|---|---|
| `chap3-2:499-502` | Per-chart dep resolution + abort-on-first-failure; new `\myworries` on the sampling bias |
| `chap3-2:504-507` | 963 distinct repos · 9,600 executions · 77.7% · 2,141 failures across 178 repos |
| `chap3-2:515-518` | Taxonomy sentences left intact but flagged — superseded corpus, 4 → 177 dependency |
| `chap3-3:218-255` | `tab:ch5-fix-results` rewritten: 889/723/166 · 322/0/322 · **1,211/723/488** |
| `chap3-3:260-273` | Prose reset to 59.7% / 81.3%; unclassified row removed entirely |
| `chap4:124-128` | `tab:ch4-corpus` synced; **Charts processed (7,401)** row added |
| `chap4:115` | Funnel `AIREV` resolved — the `break`, not the 100-combination cap |
| `chap4:256` | 873-vs-corpus `AIREV` resolved — dedup: 2,347 raw → 1,417 distinct |
| `chap4:275-284` | Total row set; the "873 vs 870 differ by three" worry retired (they coincide) |
| `chap4:82`, `chap5:9` | Group C figures unwrapped to 59.7% / 81.3% |

**Not compiled** — no TeX toolchain on this machine. Edits verified brace-neutral against `HEAD`
(chap3-2 carries a pre-existing 2-brace counter artefact, unrelated). Build before trusting.

**Still open:** applying `08-failure-class-map.md`'s numbers to `chap4:280-285` (A3 is measured,
not yet written into the `.tex`), and E3, the seed capture date.

## Headline

**The corrected recovery rate is 59.7% (723/1,211), not the published 55.3%** — the claim gets
*stronger*. Injection-applicable recovery is **81.3% (723/889)**, up from 80.5%.

Two figures that must not be quoted again:
- **72,225 charts** — 91% were never attempted (`main.go:207` aborts a whole repo on the first
  dep-build failure). Report attempted charts, not discovered.
- **2,446 template failures** — 305 are spurious vendored-subchart renders. True count **2,141**.

## Verifying

```sh
cd <helm-tests>/helm_fetcher
py ../../thesis/Detection_of_Misconfiguration_and_Vulnerabilities_in_Helm_Charts/thesis_assets/scripts/verify_corpus.py
```

Exits non-zero if any thesis figure stops matching the catalogs.

## The two corpora — do not confuse them

The single most important distinction in this folder.

| | Authoritative | Superseded |
|---|---|---|
| GitHub catalog | `runs/20260528_215041_github/` | `results/github/` (April) |
| Repos | 487 | 322 |
| Runs | 8,012 | 13,593 |
| Template failures | 2,198 | 1,297 |

`FINAL_RESULTS.md:55` declares `results/github/` superseded. The two catalogs share only **91**
repos — independently confirmed by that same line's "~91/322 overlap" note.

Ch. 3.2 describes the authoritative corpus. Ch. 3.3's taxonomy was computed on the superseded
one, because `Makefile:46` still defaulted to it (now fixed). See `02-taxonomy.md`.

Two overlap figures, easily confused — quote the right one:

- **91** of the 487 authoritative GitHub repos are in the superseded *GitHub catalog*.
- **97** of the 487 appear *anywhere* in `catalog_sources_merged.json` (91 via the stale
  GitHub half + 10 that entered via ArtifactHub − 4 present in both).
