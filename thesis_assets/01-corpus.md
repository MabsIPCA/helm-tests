# Ch. 3.2 — Corpus figures and provenance

**Status: verified 2026-07-17.** Every figure below reproduces exactly from the catalogs via
`scripts/verify_corpus.py`. These are the numbers `thesis-todos.md` group C calls "already
unwrapped and correct in `tab:ch4-corpus`" — that assessment holds.

## Source of truth

The corpus is the concatenation of the two authoritative per-source runs:

```sh
cd <helm-tests>/helm_fetcher
go run ./scripts/merge_sources catalog_sources_v2.json \
    runs/20260528_215041_github/catalog_by_project.json \
    runs/20260529_135859_artifacthub/catalog_by_project.json
```

Rendered summary: `helm_fetcher/FINAL_RESULTS.md` (generated 2026-06-16).

> **Not** `catalog_sources_merged.json` — that file is the *superseded* April GitHub catalog
> plus ArtifactHub. See `02-taxonomy.md`.

## Headline figures — as cited in the thesis

Cited at `chap3-2:501` and `tab:ch4-corpus` (`chap4:128-131`).

| Figure | Value | Provenance |
|---|---:|---|
| Repositories (seed entries) | 973 | 487 GitHub + 486 ArtifactHub |
| Repositories (**distinct**) | **963** | 10 repos appear in both seeds — see Correction 3 |
| Charts discovered | 72,225 | sum `total_charts` |
| `helm template` executions | 10,110 | sum `total_runs` |
| Successes | 7,664 | sum `total_successes` |
| Template failures | 2,446 | sum `total_failures` |
| Dependency-build failures | 218 | sum `total_dep_failures` |
| Baseline success rate | 75.8% | 7,664 / 10,110 = 75.806% |

## By source

| Metric | GitHub | ArtifactHub | Combined |
|---|---:|---:|---:|
| Repos | 487 | 486 | 973 |
| Charts discovered | 69,034 | 3,191 | 72,225 |
| Template runs | 8,012 | 2,098 | 10,110 |
| Successes | 5,814 | 1,850 | 7,664 |
| Template failures | 2,198 | 248 | 2,446 |
| Dep-build failures | 194 | 24 | 218 |
| **Success rate** | **72.6%** | **88.2%** | **75.8%** |

## Repo status breakdown

`kept` = has template failures **and** not dep-failed. These three are mutually exclusive and
sum to the corpus.

| Status | GitHub | ArtifactHub | Combined |
|---|---:|---:|---:|
| Kept (has template failures) | 86 | 67 | 153 |
| Removed (clean, no failures) | 207 | 395 | 602 |
| Dep-failed | 194 | 24 | 218 |
| **Total** | **487** | **486** | **973** |

## Corrections required in Ch. 3.2 prose

Three errors found while tracing. All verified; none needs a re-run.

### Correction 1 — "153 repositories" pairs mismatched populations (`chap3-2:503`)

> *"Of these, 2,446 rendered with template failures across 153 repositories"*

2,446 is corpus-wide; 153 is the *filtered* (`kept`) count. They don't describe the same set.

| Population | Repos | Failures |
|---|---:|---:|
| All repos with ≥1 template failure | **200** | **2,446** |
| — of which `kept` | 153 | 2,102 |
| — of which dep-failed but still rendered | 47 | 338 |
| — neither (AH edge case) | 2 | 6 |

**Fix:** either *"2,446 template failures across 200 repositories"*, or *"2,102 template
failures across the 153 repositories retained for analysis"*. Not the current mix.

### Correction 2 — the dep-failure exclusion is per-chart, not per-repo (`chap3-2:500`)

> *"Repositories for which dependency resolution fails are recorded separately and excluded
> from template rendering, as a completed dependency step is a prerequisite."*

This describes a whole-repo exclusion that does not occur. `helm dependency build` fails **per
chart**: 46 GitHub repos are marked dep-failed yet produced **332** template failures from
their other charts. A monorepo can have one chart fail dep-build while its siblings render.

**Fix:** state that dependency resolution is evaluated per chart, and that a repo is labelled
dep-failed while its resolvable charts still render.

### Correction 3 — 973 is entries, not repositories (`chap3-2:501`)

The two top-500 seeds overlap: **10 repos are in both**, so they are cloned, rendered and
counted twice.

| Repo | GitHub runs / fails | ArtifactHub runs / fails |
|---|---|---|
| `nicholaswilde/helm-charts` | 63 / **63** | 63 / **63** |
| `EugenMayer/helm-charts` | 22 / 5 | 22 / 5 |
| `secureCodeBox/secureCodeBox` | 42 / 0 | 42 / 0 |
| `bitnami/charts` | 1 / 0 | 118 / 1 |
| `prometheus-community/helm-charts` | 7 / 1 | 44 / 2 |
| `wiremind/wiremind-helm-charts` | 24 / 0 | 80 / 1 |
| `DataDog/helm-charts` | 1 / 0 | 9 / 0 |
| `grafana/helm-charts` | 1 / 0 | 2 / 0 |
| `k8s-home-lab/helm-charts` | 0 / 0 | 46 / 0 |
| `openebs/openebs` | 2 / 0 | 2 / 0 |

`nicholaswilde/helm-charts` shows identical 63 runs / 63 failures on both sides — the same
clone rendered twice, contributing 126 of the 2,446 failures where it should contribute 63.
~70 failures are double-counted; 591 runs are attributable to double-counted repos.

**Fix:** report **963 distinct repositories (973 seed entries; 10 repos appear in both
top-500 lists)**, or dedup before counting.

**Does not affect the taxonomy.** `analyzer.go` dedups occurrences by
`(repo_name, error_message)` — keyed on repo *name*, not (source, repo) — so duplicate entries
collapse automatically. This is why the 808 merged entries reduce to 799 distinct.

### Correction 4 — the headline figures predate the vendored-subchart fix

**Found 2026-07-17, and it moves the success rate.** `FINAL_RESULTS.md` (2026-06-16) is on the
*value-prefiltered* basis, but the vendored-subchart fix
(`FindCharts`/`IsVendoredSubchart`/`RelativeComponentDirs`) landed **after** the May runs. The
authoritative catalogs therefore still render vendored subcharts and `file://` components
standalone — which always fail, spuriously.

Passing the corpus through `scripts/drop_subcharts` (the same step the Jun 10 pipeline applied
to the stale catalog, hence `catalog_sources_merged.json.presubchart.bak`) drops **521
non-standalone charts / 510 runs**:

| Basis | Runs | Successes | Template failures | Success rate |
|---|---:|---:|---:|---:|
| As published (raw May) | 10,110 | 7,664 | 2,446 | **75.8%** |
| Subchart-cleaned | 9,600 | 7,459 | **2,141** | **77.7%** |
| *delta* | −510 | −205 | **−305** | +1.9pp |

**305 of the published 2,446 template failures are false** — subcharts that were never meant to
render standalone. The corrected figures are **9,600 executions, 2,141 template failures,
77.7% success**.

**Decision needed.** Either (a) report the cleaned basis and note the correction, or (b) keep
the published basis and disclose that it includes ~305 known-spurious failures. (a) is
recommended: the taxonomy in `02-taxonomy.md` is already computed on the cleaned basis, so (b)
leaves Ch. 3.2 and Ch. 3.3 on different bases — the exact class of error this folder exists to
prevent.

> **Tooling gotcha:** `drop_subcharts` *recomputes* `total_charts` as `len(Charts)`, collapsing
> it from 72,225 to 7,401 — because the catalog only stores summaries for charts actually
> processed (see `03-todo-verdicts.md` D3). This is a **lossy side effect, not a real finding**:
> the 72,225 discovery count remains valid and is read from the pristine run catalogs by
> `verify_corpus.py`. Never quote `total_charts` from a `drop_subcharts`-processed catalog.

## Open reproducibility gap

`thesis-todos.md` E3 stands: **the top-500 seed capture date is unrecorded** (`chap3-2:502`,
`chap4:112`). `github_search.json` / `artifacthub_search.json` are dated 2026-05-28 on disk,
and the run directories are `20260528_215041` / `20260529_135859` — consistent with a
**2026-05-28** capture, but that is filesystem inference, not a recorded fact. Confirm and
state it in both places.
