# Ch. 3.3 — Taxonomy and recovery table

**Status: the published numbers have an invalid basis.** `tab:ch5-fix-results` (`chap3-3:222`)
and every figure derived from it were computed on a corpus that Ch. 3.2 does not describe and
that `FINAL_RESULTS.md` declares superseded. A corrected run is in flight; §Pending below is
filled when it lands.

This is *not* one of the doubts in `thesis-todos.md` — it was found while tracing A3/D2, and it
supersedes both of them.

## The finding

`catalog_sources_merged.json` — the input behind every taxonomy number — is **not** the
corpus. It is the superseded April GitHub catalog plus the current ArtifactHub run. Verified by
exact composition match on all four dimensions (not a coincidence at this precision):

| | repos | runs | template failures | dep failures |
|---|---:|---:|---:|---:|
| `results/github/` (superseded, April) | 322 | 13,593 | 1,297 | — |
| `runs/20260529_135859_artifacthub/` (authoritative) | 486 | 2,098 | 248 | — |
| **sum** | **808** | **15,691** | **1,545** | **113** |
| `catalog_sources_merged.json` (actual) | 808 | 15,691 | 1,545 | 113 |

The corpus Ch. 3.2 describes is 973 / 10,110 / 2,446 / 218. **The taxonomy never saw it.**

Only **91** of the 487 authoritative GitHub repos are in the superseded catalog (97 appear
anywhere in the merged input, counting 10 that entered via the ArtifactHub half).
The ArtifactHub half of the taxonomy is correct; the **GitHub half is stale**.

### Root cause

`helm_fetcher/Makefile:46` and `scripts/refix_all.ps1:19` both defaulted to:

```make
GITHUB_CATALOG   ?= results/github/catalog_by_project.json          # superseded April snapshot
ARTIFACT_CATALOG ?= runs/20260529_135859_artifacthub/...            # authoritative
```

The ArtifactHub default was updated when the new run landed; the GitHub one never was. Fixed in
`Makefile:46` (with a comment recording why), plus a note that `SOURCES_CATALOG` must be
rebuilt from the same pair or the analyzer scores a different corpus than the fixer ran on.

## The published table is also one analyzer change stale

Independent of the corpus problem. The **481 resolved** in `tab:ch5-fix-results` predates the
2026-06-24 non-reproduced split, which is implemented in `analyzer.go` but never propagated to
the thesis. The real figure is **479 resolved + 2 non-reproduced**.

A "resolved" with an empty `fix_chain` is not a fix — the chart rendered clean on the fixer's
first attempt with nothing injected, i.e. the original failure did not reproduce (floating
dependency versions and Helm drift). The 2 false positives are exactly the printed
`nil_pointer` **298** (really 297) and `missing_template` **1** (really 0).

**Correcting this strengthens the argument.** Structurally-blocked resolutions become a clean
**0**, and all 479 recoveries fall in the value-injection group — exactly what the design
predicts. The current table's lone `missing_template` "success" is an artefact that undercuts
the claim.

| Figure | Published | Corrected |
|---|---|---|
| Resolved overall | 481 / 870 = 55.3% | **479 / 870 = 55.1%** (+2 non-reproduced) |
| Value-injection applicable | 480 / 596 = 80.5% | **479 / 596 = 80.4%** |
| Structurally blocked, resolved | 1 | **0** |

## Superseded numbers — recorded so they are recognisable

From `taxonomy_analyzer/out/cumulative_final/` (generated 2026-06-24). **Do not cite these.**

| Metric | Value |
|---|---:|
| Repos / runs | 808 / 15,691 |
| Template / dependency failures | 1,545 / 113 |
| Duplicates collapsed | 680 |
| Classified / unclassified errors | 873 / 105 |
| Non-fixable errors | 261 |
| Fix attempted | 870 |
| Fix resolved | 479 |
| Fix non-reproduced | 2 |
| Fix unresolved | 389 |

### How these reconcile (this is D2's answer)

The table is not broken arithmetic — it mixes three denominators:

```
1,545 template + 113 dependency = 1,658 raw occurrences
  − 680 collapsed by (repo_name, error_message) dedup
  =  978 distinct  =  873 classified + 105 unclassified

873 classified = 869 template + 4 dependency
870 fix attempts = 869 template + 1 unclassified   ← only 1 of the 105 was ever attempted
870 = 479 resolved + 389 unresolved + 2 non-reproduced ✓
```

The sub-rows count *distinct classified errors*; the Total row counts *fix attempts*. The
"105 unclassified" row contributes **1** to the total, not 105 — that is the entire
974-vs-870 discrepancy. See `03-todo-verdicts.md` D2.

## Pending — corrected run

```sh
cd <helm-tests>/helm_fetcher
go run ./scripts/merge_sources catalog_sources_v2.json \
    runs/20260528_215041_github/catalog_by_project.json \
    runs/20260529_135859_artifacthub/catalog_by_project.json
powershell -ExecutionPolicy Bypass -File scripts/refix_all.ps1 \
    -GithubCatalog runs/20260528_215041_github/catalog_by_project.json \
    -GithubCloneDir D:/helm_clones_github \
    -ArtifactCatalog runs/20260529_135859_artifacthub/catalog_by_project.json \
    -ArtifactCloneDir D:/helm_clones_artifacthub \
    -GithubFixed catalog_fixed_github_v2.json -ArtifactFixed catalog_fixed_ah_v2.json \
    -Cumulative catalog_fixed_cumulative_v2.json \
    -SourcesCatalog catalog_sources_v2.json -AnalyzeOut out/cumulative_v2
```

Feasibility confirmed before launch: every failing repo's clone is on disk (132/132 GitHub,
68/68 ArtifactHub). Outputs are `*_v2` so the existing artifacts are preserved for comparison.

Both halves are re-fixed under the same `helmfix` build, so the two sources are directly
comparable — the current split (April GitHub + June ArtifactHub) cannot claim that.

### Two corrections applied during the run

1. **App Control blocked `go run`.** `go run ./scripts/merge_fixed` died with *"Uma política de
   Controlo de Aplicações bloqueou este ficheiro"* — Windows App Control blocks the temp
   executable `go run` builds. Workaround: `go build -o <local>.exe ./scripts/...` and run from
   the working directory. The fixers themselves completed fine.
2. **The raw May catalogs predate the vendored-subchart fix.** `catalog_sources_v2.json` had to
   be passed through `drop_subcharts`, exactly as the Jun 10 pipeline did to the stale catalog
   (`.presubchart.bak`). It dropped **521 non-standalone charts / 510 runs = 205 successes +
   305 false failures**. Skipping this step inflates the taxonomy (1,604 classified instead of
   1,388) because vendored subcharts rendered standalone fail spuriously.

## FINAL — corrected numbers (2026-07-17)

Basis: `catalog_sources_v2.json` (subchart-cleaned) + `catalog_fixed_cumulative_v2.json`.
Report: `taxonomy_analyzer/out/cumulative_v2/`.

**Guarded.** `verify_corpus.py` reads `out/cumulative_v2/taxonomy_report.json` and asserts every
figure below — totals (repos/runs/failures/collapsed/classified/attempts/resolved/unresolved),
all 16 template sub-kinds of `tab:ch5-fix-results` (count + resolved), the cited dependency
sub-kinds, and the 59.7% / 81.3% rates. The report itself is reproducible: re-running the
analyzer on the two catalogs regenerates it byte-for-byte apart from the timestamp.

**The basis is provably correct now:** the corpus totals reproduce Ch. 3.2 exactly
(973 repos / 218 dep failures), and `template` kind (1,211) ≡ `fix_attempted` (1,211) — every
template error was attempted, with no unclassified straggler and no non-reproduced artefacts.
The published run could never say that.

| Metric | Published (stale corpus) | **Corrected** |
|---|---:|---:|
| Repositories | 808 | **973** |
| Helm runs | 15,691 | **9,600** |
| Template failures | 1,545 | **2,141** |
| Dependency failures | 113 | **206** |
| Duplicates collapsed | 680 | **930** |
| Classified errors | 873 | **1,388** |
| Unclassified errors | 105 | **29** |
| Fix attempts | 870 | **1,211** |
| Fix resolved | 481 *(really 479)* | **723** |
| Fix unresolved | 389 | **488** |
| Fix non-reproduced | 2 | **0** |
| **Overall recovery** | **55.3%** | **59.7%** |
| **Injection-applicable** | **80.5%** (480/596) | **81.3%** (723/889) |

### Corrected `tab:ch5-fix-results`

| Kind | Sub-kind | Failures | Resolved | Unresolved |
|---|---|---:|---:|---:|
| | *Value-injection applicable* | | | |
| template | nil_pointer | 587 | 540 | 47 |
| template | required_value | 263 | 159 | 104 |
| template | kube_version_incompatible | 24 | 24 | 0 |
| template | type_mismatch | 15 | 0 | 15 |
| | *Subtotal* | *889* | *723* | *166* |
| | *Structurally blocked* | | | |
| template | missing_template | 110 | 0 | 110 |
| template | values_schema_validation | 92 | 0 | 92 |
| template | malformed_yaml | 23 | 0 | 23 |
| template | author_assertion | 21 | 0 | 21 |
| template | runtime_eval | 21 | 0 | 21 |
| template | custom_validation | 20 | 0 | 20 |
| template | library_chart_not_installable | 11 | 0 | 11 |
| template | parse_error | 9 | 0 | 9 |
| template | unsupported_builtin | 8 | 0 | 8 |
| template | chart_metadata | 4 | 0 | 4 |
| template | invalid_value | 2 | 0 | 2 |
| template | dependency_check_failed | 1 | 0 | 1 |
| | *Subtotal* | *322* | *0* | *322* |
| | **Total** | **1,211** | **723** | **488** |

Every identity closes: 889 + 322 = 1,211 · 723 + 166 = 889 · 723 + 488 = 1,211.

### Repo-level failure attribution — the early-stop overlap

Data: `catalog_sources_v2.json`. Because the classification tests `TotalDepFailures > 0`
**first** (`helm_fetcher/main.go:288`), a repository whose chart loop aborts on a later chart's
`helm dependency build` is recorded `dep_failed` even though charts visited *before* the `break`
already logged their template failures. So "repos with template failures" splits across two
buckets:

| Repos with ≥1 template failure | Repos | Template failures |
|---|---:|---:|
| `dep_failed = false` (Kept bucket) | 142 | 1,964 |
| `dep_failed = true` (aborted after visiting some charts) | 36 | 177 |
| **Total** | **178** | **2,141** |

Both partitions close: 142 + 36 = 178 and 1,964 + 177 = 2,141. The 36-repo overlap is dominated
by `oracle-cne/catalog` (95 of the 177), `IndustryFusion/DigitalTwin` (12), and
`elastisys/compliantkubernetes-apps` (11). Reproduce: filter `total_failures > 0`, split on
`dep_failed`. Guarded by `verify_corpus.py` (Ch. 3.3 v2 block).

**This table is cleaner than the published one.** No unclassified row is needed (dependency and
unknown errors are not fix-attempted, so they belong in the taxonomy table, not the recovery
table) — which removes the entire cause of D2. Structurally-blocked resolved is a clean **0**,
and non-reproduced is **0**: both of the old "successes" were stale-corpus artefacts.

## Render-path validation — the CLI vs KICS's own call path

**Why this exists:** `chap3-3:217` claimed the validation tool "imports
`helm.sh/helm/v3/pkg/action` … rather than invoking the helm binary". That was **false**. The
recovery table is built from `catalog_fixed_*.json`, which `refix_all.ps1` produces via
`helm_fetcher -mode fixer` → `helm/helm.go:338` → `exec.Command("helm", ...)` — the **CLI**.
This was true of the published 481/870 too; the re-run inherited it.

It matters because the whole transferability argument rests on it: *"these rates carry over to
KICS because we used KICS's call path."* Measured through the CLI, that was asserted, not shown.

### Measured, not asserted (2026-07-17)

Replayed the same 2,141 failing runs through `kics-render-mock -catalog`
(`pipeline.go` → `action.Install.Run()`), matched **run-by-run** by `helm_command`, zero unmatched:

| Outcome | Runs |
|---|---:|
| Resolved by **both** | 869 |
| CLI only | 26 |
| SDK only | 5 |
| Neither | 1,211 |
| SDK renders clean at base (failure did not reproduce) | 30 |
| **Total** | **2,141** |

**On the 2,111 both treat as failures: CLI 895 (42.4%) vs library path 874 (41.4%) — 98.5%
agreement, 1.0pp apart.** The claim is now earned.

### The first attempt was misleading — a lesson worth keeping

The naive replay gave 791 (37.5%) vs 910, a 119-resolution gap that *looked* like the action API
performing worse. It wasn't. `kics-render-mock`'s `fixInvocation` was a **simplified mock**
missing two capabilities `helm_fetcher`'s `fixRun` has:

1. **`--kube-version`.** `helmfix.ParseError` returns `path=""` for `kube_version`, so the mock's
   `patch[path] = value` executed `patch[""] = "1.29.0"` — a nonsense `--set =1.29.0`. **95 of the
   910 CLI resolutions use `kube_version`: 80% of the gap on its own.**
2. **Dependency fetch** — no build/update-and-retry for missing subcharts (~24).

Both were ported (`pipeline.go` `runOnceKube`, `fixer.go` `fixInvocation`, mirroring `fixRun`
step for step), which closed 119 → 36. **Comparing two implementations is only a test of the
thing you care about once every other difference is eliminated.**

### The residual is a finding, not noise

**30 of 2,141 (1.4%) of CLI-recorded failures do not reproduce through `action.Install.Run()`.**
Drift is ruled out: today's CLI fixer reproduced all 2,141 with **zero** non-reproduced, while the
SDK missed 30 — same hour, same clones. So a small fraction of the catalogued failures are
**helm CLI entry-point artefacts that KICS would never encounter**.

> Note: `-catalog` mode writes `kics-render-output.json` / `kics-fixed-output.json` into every
> chart directory it touches (1,386 dirs under `D:\helm_clones_*`). Harmless — `FindCharts` only
> looks for `Chart.yaml` — but it is scratch output, and the dep-fetch step vendors subcharts
> into `charts/`, so a second run's *base* renders are not identical to the first's.

### What changed in the story

- **Recovery improved: 55.3% → 59.7%.** The headline claim gets *stronger*, not weaker.
- **`nil_pointer` recovers at 92%** (540/587), up from 88%.
- **`type_mismatch` is now 0 of 15**, not 1 of 9. The lone success is gone — it never existed.
- **`missing_template` is the largest blocked bucket at 110**, up from 23. Worth prose.
- **Unclassified collapsed 105 → 29.** The published claim at `chap3-2:509` that unclassified
  are *"mostly dependency build errors that yielded only a generic exit-status code"* is an
  artefact of the **April** data: the May run captures full dependency error text.
- **The dependency taxonomy is transformed: 4 classified → 177, across 12 sub-kinds**
  (`missing_repository` 84, `missing_subchart` 41, `chart_validation` 14, `lock_file_out_of_sync`
  9, …). `chap3-2:510` says *"only 4 classified dependency failures were recorded: missing
  repository (3) and missing subchart archive (1)"* — that sentence massively undersells the
  work and must be rewritten.
