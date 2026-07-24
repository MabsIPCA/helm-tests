# Verdicts on `thesis-todos.md`

One row per doubt. Two of the four "contradictions" are false alarms and need no run; the
biggest problem was not on the list at all.

| Item | Verdict |
|---|---|
| **D1** — "the 873 does not exist" | ❌ **False alarm.** 873 is real and sourced. Do not change it |
| **D2** — "Ch. 3 table does not add up" | ❌ **False alarm.** Reconciles exactly; three denominators |
| **A3** — re-aggregate taxonomy | ⛔ **Blocked, but not by D2** — the basis corpus is wrong |
| **Group C** — "REAL, cross-check then unwrap" | ⚠️ **Reclassify.** Real, but measured on the wrong corpus |
| **D3** — corpus funnel unexplained | ✅ **Answered.** 92.5% of charts are in dep-failed archive monorepos |
| **D4** — prevalence vs recoverability | ⚠️ **Two populations, not a contradiction.** Resolve at A3 |
| **Ch. 3.2 corpus figures** | ✅ **Verified.** All reproduce — `scripts/verify_corpus.py` |
| *(not on the list)* | 🔴 **Taxonomy computed on a superseded corpus.** See `02-taxonomy.md` |
| *(not on the list)* | 🔴 **3 prose errors in Ch. 3.2.** See `01-corpus.md` §Corrections |

## D1 — the 873 **does** exist. Do not "fix" it

> *"§4.5.1 claims 873 corpus render failures. This number appears nowhere else in the thesis."*
> *"Likely fix: … The 873 looks like a drafting error for 870."*

**It appears at `chap3-2:508`:** *"Across the 873 classified failures, nil dereference was the
dominant template error with 340 occurrences…"* It is the analyzer's `classified_errors` total.

**Do not apply the suggested fix.** 873, 870 and 2,446 are three different quantities, none a
typo for another:

| Value | Is | Definition |
|---:|---|---|
| **2,446** | corpus template failures | raw failing runs, not deduped (`chap3-2:503`) |
| **873** | classified errors | distinct `(repo, error)` pairs that a rule matched = 869 template + 4 dependency |
| **870** | fix attempts | 869 template + 1 unclassified that was attempted |

Rewriting 873 → 870 would replace a correct number with a wrong one.

**The one real defect:** 873 is called *"corpus render failures"*, but 4 of the 873 are
dependency-resolution failures, not renders. Call it "classified failures" (as `chap3-2:508`
correctly does), or say "869 classified render failures".

**Caveat:** 873 is right *as an interpretation*, but it was measured on the superseded corpus,
so its value will change. The label is what needed fixing, not the arithmetic.

## D2 — the table adds up. It mixes three denominators

> *"Failures: 596 + 273 + 105 = 974 vs Total row 870."*

Not an arithmetic error. The columns count different things:

- **Sub-rows** count *distinct classified errors* → template sub-kinds sum to **869**.
- **Total row** counts *fix attempts* → **870**.
- The **105 unclassified** row is the culprit: only **1** of those 105 was ever fix-attempted
  (the rest are generic `exit status 1` dependency errors with nothing to parse). It
  contributes **1** to the total, not 105.

```
869 template sub-rows + 1 attempted unclassified = 870 ✓
```

The TODO's observation that "excluding the 105, the classified rows sum to 869 … still one
short of 870" is exactly right — **that missing one is the attempted unclassified run.**

The remaining 481-vs-480 wobble is the non-reproduced split, not arithmetic:

```
870 attempted = 479 resolved + 389 unresolved + 2 non-reproduced ✓
```

**Fix the presentation, not the numbers.** Either print the unclassified row as `1` attempted
(and note 104 were never attempted), or drop it from the table and state it in prose. Also add
a non-reproduced column or footnote — see `02-taxonomy.md`.

### Why the raw failure count is so much larger

The analyzer dedups by `(repo_name, error_message)`: the fetcher renders up to 100 value-file
combinations per chart, so one broken chart yields many identical errors. 1,658 raw
occurrences collapse to 978 distinct (**680 collapsed**). Worth one sentence in Ch. 3.3 — the
gap between 2,446 corpus failures and 873 classified is otherwise unexplained, and invites the
same "these don't add up" reading that produced D2.

## A3 — blocked by the corpus, not by D2

The TODO calls A3 "the cheapest win — no new run needed". **It needs a run.** D2 is resolved and
was never the blocker; the sub-kind counts it would aggregate are themselves computed on the
superseded corpus. Re-map onto `tab:ch4-failure-class-map` only after the corrected run.

## Group C — reclassify from "REAL" to "real, wrong corpus"

`thesis-todos.md` says these are *"measured and already published in Chapter 3 with a full table
behind them. The `\myworries{}` around them is over-caution — cross-check each, then unwrap."*

**Do not unwrap yet.** They are measured, but on the superseded corpus, and all are expected to
move:

| Value | Location | Status |
|---|---|---|
| 870 failing renders | chap4:82, 276 | wrong corpus + it's *fix attempts*, not "failing renders" |
| 481 resolved | chap4:82, 269, 276 | wrong corpus **and** stale: really 479 + 2 non-reproduced |
| 55.3% overall | chap4:82, 269, 276; chap5:9 | wrong corpus; 479/870 = 55.1% on current data |
| 480 of 596 | chap4:277 | wrong corpus; 479/596 on current data |
| 80.5% injection-applicable | chap4:82, 277; chap5:9 | wrong corpus; 80.4% on current data |

The group C figures that **are** safe to unwrap are the Ch. 3.2 corpus ones (973 / 72,225 /
10,110 / 75.8%) — verified in `01-corpus.md`, subject to Correction 3 (973 → 963 distinct).

## D3 — corpus funnel: measured, and the cause is concentrated

> *"72,225 charts discovered but only 10,110 executions, about 14%."*

The TODO's proposed explanation (dependency resolution) is **correct, and stronger than it
looks** — but for a specific reason worth stating, because "14%" implies a broad, uniform loss
and the reality is a handful of outlier repos.

Charts and runs by repo status:

| Status | Repos | Charts | Runs | Charts/repo |
|---|---:|---:|---:|---:|
| **Dep-failed** | 218 | **66,837 (92.5%)** | 2,907 | 307 |
| Kept (has failures) | 153 | 3,603 | 5,221 | 24 |
| Removed (clean) | 602 | 1,785 | 1,982 | 3 |
| **Total** | **973** | **72,225** | **10,110** | 74 |

**92.5% of all discovered charts sit in the 218 dep-failed repos.** They are chart *archives* —
mirrors and catalogs holding thousands of versioned charts — averaging 307 charts/repo against
24 for kept repos and 3 for clean ones. Discovery is dominated by a few of them:

| Repo | Charts | Runs | Status |
|---|---:|---:|---|
| `rancher/partner-charts` | 8,606 | 32 | dep-failed |
| `hey101/scale-catalog` | 6,209 | **0** | dep-failed |
| `adstanley/archive` | 4,619 | **0** | dep-failed |
| `Makhuta/truecharts-archive-scale-catalog` | 4,593 | **0** | dep-failed |
| `Yuzu815/TruenasScaleArchive` | 4,593 | **0** | dep-failed |

**8 repos hold 56% of all discovered charts; the top 50 hold 88.6%.** 93 repos have charts but
produced **zero** runs, stranding 33,553 charts on their own.

Two consequences for the thesis:

1. **The funnel is not a 14% success rate.** Excluding dep-failed repos, 5,388 charts produced
   7,203 runs — *more* than one run per chart, because each chart renders once per value-file
   combination. There is no unexplained 86% loss; there is one dominant cause.
2. **"72,225 charts discovered" overstates corpus breadth** and should not be read as scale.
   It is mostly TrueNAS/TrueCharts archive mirrors and rancher catalogs. The meaningful
   denominator is the 5,388 charts in renderable repos, or the 10,110 executions.

### The mechanism — a `break`, and it is a limitation to disclose

The above is *what*; this is *why*, and it is sharper than "dependency resolution failed".
`main.go:205-223`:

```go
for _, chartDir := range charts {
    depErr := helm.RunHelmDependencyBuild(chartDir)
    if depErr != nil {
        log.Warn().Msg("Dependency build failure – skipping remaining charts in repo")
        repoResult.Charts = append(repoResult.Charts, chartSummary)
        repoResult.TotalDepFailures++
        break                    // abandons EVERY remaining chart in this repo
    }
    ...
}
```

**One chart failing `helm dependency build` abandons the whole repo, mid-iteration.** Charts are
walked in `filepath.Walk` (lexical) order, so everything after the first failure is dropped:

| Repo | Discovered | Chart summaries stored | Runs | Stranded |
|---|---:|---:|---:|---:|
| `hey101/scale-catalog` | 6,209 | **1** | 0 | 6,208 |
| `adstanley/archive` | 4,619 | **1** | 0 | 4,618 |
| `Jaebytes/TrueCharts` | 4,494 | **1** | 0 | 4,493 |
| `rancher/partner-charts` | 8,606 | 33 | 32 | 8,573 |

**GitHub: 69,034 charts discovered, only 6,188 ever stored — 62,846 (91.0%) stranded by the
`break`.** `hey101/scale-catalog`'s *first* chart failed, so 6,208 others were never attempted.

Three consequences, in order of severity:

1. **"72,225 charts discovered" cannot stand as a measure of scale.** 91% were never attempted,
   and not because of anything about those charts — a control-flow decision dropped them. The
   defensible figure is charts *attempted* (~7,922 with summaries; 7,401 after subchart
   cleaning), or the 9,600 executions.
2. **The sample is order-biased, not random.** Within an archive repo, the alphabetically-early
   charts are rendered and the rest are dropped. This should be stated as a threat to validity —
   the corpus is not a uniform sample of the discovered population.
3. **It explains Correction 2 exactly.** Dep-failed repos still show template failures (46 repos,
   332 failures) because charts *before* the break rendered normally. The prose's "excluded from
   template rendering" describes neither the intent nor the implementation.

**Fix:** in §4.2/§3.2, report attempted charts alongside discovered, state the abort-on-first-
dep-failure behaviour, and disclose the ordering bias. If the `break` were a `continue`, the
corpus would be far larger — that is future work, not a re-run to attempt now.

> Note: vendored subcharts are **not** the explanation for the funnel — `helm/helm.go:44`
> (`FindCharts`) already excludes `IsVendoredSubchart` dirs and `file://` components at
> discovery, so they are never in the 72,225. `main.go:201` sets
> `TotalCharts = len(FindCharts(...))`. They *are* a separate correctness issue: see
> `02-taxonomy.md` — the raw May catalogs predate that fix and carry 305 false failures.

## D4 — resolve after A3

The `type_mismatch` tension is real: rare as a sub-kind (**9** occurrences) yet `chap4:62` calls
type-system mismatch "the most common class across the function library". Those are different
populations — controlled test-suite function coverage vs. production corpus frequency. Say so
explicitly when A3 maps sub-kinds onto classes; the two views are not in conflict once the
populations are named.

## E3 — seed date

See `01-corpus.md` §Open reproducibility gap. Filesystem evidence points to **2026-05-28**, but
it must be confirmed and recorded at `chap3-2:502` and `chap4:112`.
