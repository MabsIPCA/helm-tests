# Plot Helpers

Auxiliary scripts to generate 2D SVG plots from `helm_fetcher` outputs.

## Inputs

- `catalog_by_project.json` for problems per project
- `github_search.json` for repository popularity by stars

## Scripts

- `plot_problems_per_project.py`
  - Uses `total_failures + total_dep_failures` per repo
  - Default output: `plots/problems_per_project.svg`
- `plot_repo_popularity.py`
  - Uses `repos[].stars` from `github_search.json`
  - Default output: `plots/repo_popularity_stars.svg`
- `plot_problems_per_popularity.py`
  - Joins repos by `repo_url` between both JSON files
  - Plots stars (X) vs problems (Y)
  - Filters out projects where `problems == 0`
  - Default X-axis scale is `log1p` to separate low-star repositories
  - Default output: `plots/problems_per_popularity.svg`

## Run Directly

From `helm_fetcher/`:

```bash
python scripts/plot_problems_per_project.py --catalog-in catalog_by_project.json --output plots/problems_per_project.svg --top 40
python scripts/plot_repo_popularity.py --search-in github_search.json --output plots/repo_popularity_stars.svg --top 50
python scripts/plot_problems_per_popularity.py --catalog-in catalog_by_project.json --search-in github_search.json --output plots/problems_per_popularity.svg --top 200 --x-scale log1p
```

## Run via Makefile

From `helm_fetcher/`:

```bash
make plot-problems
make plot-popularity
make plot-problems-popularity
make plot-all
```




