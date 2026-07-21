# Helm Fetcher — Final Catalog Results

_Generated 2026-07-17 (v2 basis). Top-500 seeds per source, value-prefiltered
and subchart-collapsed._

## Corpus provenance

The corpus was cloned once and **never re-pulled** (each repo's git reflog is a
single clone entry), so chart content is frozen at clone time.

| Source | Seed file (`top: 500`) | Cloned to | Clone performed | Clone dirs on disk | Cataloged repos |
|---|---|---|---|---:|---:|
| GitHub | `github_search.json` (2026-05-28) | `D:\helm_clones_github` | **2026-05-28** | 733 | 492 |
| ArtifactHub | `artifacthub_search.json` (2026-05-28) | `D:\helm_clones_artifacthub` | **2026-05-29** | 496 | 481 |

The clone directories on disk (733 / 496) are a **superset** of the cataloged
repos (492 / 481): they include repos cloned by earlier/larger runs that are
outside the current top-500 seed, and seeds that yielded no chart. The in-repo
`helm_fetcher/cloned/` holds only the 7 large monorepos as a kept subset; the
full corpus lives on `D:\`.

## Processing model — and the dependency-failure early-stop

`-mode=full` runs, per repo: clone → discover charts (`FindCharts`, with vendored
subcharts collapsed) → for each chart, `helm dependency build`, then
`helm template` on the default values plus every value-file combination.

**Important:** within a repo the chart loop **stops at the first
`helm dependency build` failure** (`main.go:223`, `break`): that chart is recorded
as a dependency failure, the repo is marked *dep-failed*, and **every remaining
chart in the repo is skipped**. One broken dependency therefore suppresses all of
its sibling charts. This is why "charts recorded" is far below the number of
`Chart.yaml` files on disk — the gap is concentrated in dep-failed **archive /
mirror monorepos** (e.g. `rancher/partner-charts` = 8,606 charts on disk but only
33 processed; TrueCharts/TrueNAS archives ≈ 4,600–6,200 each), where an early
failure skips thousands of siblings.

| Chart accounting | GitHub | ArtifactHub | Combined |
|---|---:|---:|---:|
| `Chart.yaml` files on disk (raw) | 75,597 | 3,823 | 79,420 |
| Charts discovered by `FindCharts` (June, pre–early-stop) | 69,034 | 3,191 | 72,225 |
| Charts **recorded** in v2 (after early-stop + subchart collapse) | 5,913 | 1,488 | 7,401 |

## Results by source (v2)

| Metric | GitHub | ArtifactHub | Combined |
|---|---:|---:|---:|
| Repos | 492 | 481 | 973 |
| Charts recorded | 5,913 | 1,488 | 7,401 |
| Helm template runs | 7,749 | 1,851 | 9,600 |
| ✅ Successes | 5,787 | 1,672 | 7,459 |
| ❌ Template failures | 1,962 | 179 | 2,141 |
| 🔧 Dep-build failures (charts) | 182 | 24 | 206 |
| **Success rate** | **74.7%** | **90.3%** | **77.7%** |

> "Charts recorded" is the number of charts written to the catalog (one
> `ChartSummary` each). Of the 7,401, **206 are dep-failed** (the chart that
> triggered the early-stop, 0 template runs); the remaining ~7,195 produced the
> 9,600 template runs. Success rate is over template runs.

### Repo status breakdown (v2)

| Status | GitHub | ArtifactHub | Combined |
|---|---:|---:|---:|
| Kept (has template failures) | 81 | 61 | 142 |
| Removed (clean, no failures) | 229 | 396 | 625 |
| Dep-failed (dependency build failed) | 182 | 24 | 206 |
| **Total** | **492** | **481** | **973** |

Because of the early-stop, each dep-failed repo contributes exactly **one**
dep-build failure, so dep-failed repos (206) = dep-build failures (206).

## Notes

- **Authoritative data (v2)** is `catalog_sources_v2.json` (merged source
  catalog) paired with `catalog_fixed_cumulative_v2.json` (merged fixed catalog).
  The taxonomy report in `taxonomy_analyzer/out/cumulative_v2/` is computed from
  exactly this pair and reproduces these totals (9,600 runs / 2,141 template
  failures / 206 dep failures).
- **v2 is derived from the June run, not a fresh corpus scan.** It was produced
  from the June catalogs via refetch/refilter passes, so it inherits June's
  post–early-stop chart set: v2 knows only about the ~7,401 charts June recorded
  and does **not** re-discover the ~72k charts skipped by the early-stop. A fresh
  `-mode=full` scan would re-discover them (dominated by the mirror-archive
  monorepos above).
- The June basis (`catalog_sources_merged.json` +
  `catalog_fixed_cumulative.json`, 10,110 runs) and the April top-500 GitHub
  data in `results/github/` + `backup/GHTop500/` are **superseded**.
- Per-source *fixed* catalogs are split out (`catalog_fixed_github_v2.json`,
  `catalog_fixed_ah_v2.json`); a per-source *source* catalog split for v2 is not
  materialised — the merged `catalog_sources_v2.json` is the basis.

## Source files (v2)

```
catalog_sources_v2.json            # merged source catalog (GitHub + ArtifactHub), authoritative
catalog_fixed_cumulative_v2.json   # merged fixed catalog (paired with the above)
catalog_fixed_github_v2.json       # GitHub fixed catalog
catalog_fixed_ah_v2.json           # ArtifactHub fixed catalog
github_search.json / artifacthub_search.json   # top-500 seeds (2026-05-28)
```
