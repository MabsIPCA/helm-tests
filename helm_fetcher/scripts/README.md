# Helm Fetcher helpers

Run these commands from `helm_fetcher/`. See the [module README](../README.md)
for setup and the [final results](../FINAL_RESULTS.md) for dataset provenance.

## Catalog maintenance

The Go helpers use positional arguments. Replace angle-bracket placeholders
with your own paths.

| Helper | Invocation | Effect |
|---|---|---|
| Merge source catalogs | `go run ./scripts/merge_sources <out.json> <in1.json> <in2.json> [<in3.json> ...]` | Writes a merged source catalog |
| Merge fixed catalogs | `go run ./scripts/merge_fixed <out.json> <in1.json> <in2.json> [<in3.json> ...]` | Writes a merged fixed catalog |
| Select repositories | `go run ./scripts/pick_repos <in.json> <out.json> <substring> [<substring> ...]` | Writes repos whose names contain a supplied substring |
| Collapse subcharts | `go run ./scripts/drop_subcharts <catalog.json> [<catalog.json> ...]` | Rewrites catalogs in place and creates `.bak` backups |
| Audit subcharts | `go run ./scripts/subchart_audit <catalog_fixed.json>` | Prints subchart accounting |
| Summarize unresolved failures | `go run ./scripts/unresolved_summary <catalog_fixed.json> [firstKindFilter]` | Groups failures by fix kind and normalized final error |

Other maintenance modes (`merge`, `refetch-deps`, `refilter-values`, and
`render-md`) live in the main CLI; see the [command reference](../README.md#commands).

## PowerShell automation

- `get_token.ps1`: authenticates through GitHub CLI and saves a token in the
  ignored `.env` file. Run `make get-token`.
- `refix_all.ps1`: runs the fixer for both sources, merges fixed catalogs, and
  regenerates taxonomy outputs. Run `make refix-all`, overriding catalog and clone
  paths as needed; `SKIP_ANALYZE=1` stops after merging. The direct script's GitHub
  default uses the historical `results/github/` snapshot, while the Make target
  supplies the May run. Specify inputs explicitly when reproducing a particular
  corpus, and use matching source and fixed catalogs for analysis.
- [`../run_both_sources.ps1`](../run_both_sources.ps1): orchestrates both sources
  with separate clone directories and logs. Run `make run-both` or invoke the
  script with `-GithubCloneDir` and `-ArtifactCloneDir` for your machine.

## SVG plots

The Python helpers use the standard library and the local `svg_bar_plot.py` /
`svg_scatter_plot.py` renderers; no plotting package installation is needed.

| Script | Input and measure | Default output |
|---|---|---|
| `plot_problems_per_project.py` | Catalog: `total_failures + total_dep_failures` per repo | `plots/problems_per_project.svg` |
| `plot_repo_popularity.py` | GitHub seed: `repos[].stars` | `plots/repo_popularity_stars.svg` |
| `plot_problems_per_popularity.py` | Joins catalog and GitHub seed by `repo_url`; stars vs failures | `plots/problems_per_popularity.svg` |

The scatter plot excludes repositories with zero problems and uses a `log1p`
X-axis by default.

```powershell
python scripts/plot_problems_per_project.py --catalog-in catalog_sources_v2.json --output plots/problems_per_project.svg --top 40
python scripts/plot_repo_popularity.py --search-in github_search.json --output plots/repo_popularity_stars.svg --top 50
python scripts/plot_problems_per_popularity.py --catalog-in catalog_sources_v2.json --search-in github_search.json --output plots/problems_per_popularity.svg --top 200 --x-scale log1p
```

Equivalent Make targets are `plot-problems`, `plot-popularity`,
`plot-problems-popularity`, and `plot-all`. Make uses `catalog_cumulative.json`
by default. For the v2 data:

```powershell
make plot-all CATALOG_CUMULATIVE=catalog_sources_v2.json
```
