# Helm Fetcher

Discover Helm chart repositories from GitHub or Artifact Hub, build dependencies,
render charts with combinations of values files, and record failures for repair
and taxonomy analysis. This module uses the shared [helmfix library](../helmfix/).

## Start here

- [Repository overview](../README.md): test suites and the complete workflow.
- [Final results and corpus provenance](FINAL_RESULTS.md): authoritative datasets and limitations.
- [Helper scripts](scripts/README.md): catalog maintenance, diagnostics, and plots.
- [Taxonomy analyzer](../taxonomy_analyzer/README.md): failure classification and fix chains.

## Setup

Run commands below from `helm_fetcher/`. Install Go 1.25.3 or newer, Git, and
Helm 3 on `PATH`. Python 3 is needed for plots; PowerShell and GNU Make are
needed for the corresponding scripts and Make targets.

Keep the sibling `helmfix/` directory: `go.mod` references it with a local
`replace`. Local `vendor/` directories are ignored. If you use vendoring,
refresh it with `go mod vendor` after changing dependencies or `helmfix`.

For authenticated GitHub discovery, set `GITHUB_TOKEN` in your environment or
in an ignored `.env` file here. `make get-token` can populate it using GitHub CLI.

## Discover and scan

Capture a small GitHub seed, then scan it:

```powershell
go run . -mode=github-search-json -search-top=10 -search-out=github_sample.json
go run . -mode=full -source=github -search-in=github_sample.json -top=10 -clone-dir=cloned/github -keep-clones=true
```

For Artifact Hub:

```powershell
go run . -mode=artifacthub-search-json -search-top=10 -search-out=artifacthub_sample.json
go run . -mode=full -source=artifacthub -search-in=artifacthub_sample.json -top=10 -clone-dir=cloned/artifacthub -keep-clones=true
```

Full mode creates `runs/<timestamp>_<source>/`, snapshots the seed there, and
writes JSON catalogs plus Markdown and CSV reports. Use `-run-dir` to choose
the output directory. It also updates the cumulative catalog and skips
repositories already recorded there; a rerun is not automatically a fresh scan.

Specify both `-source` and a matching `-search-in`: the CLI defaults to
`artifacthub`, but its default input filename is `github_search.json`.
Without `-keep-clones=true`, clean repositories may be deleted after processing.

## Repair and analyze

Pass a catalog from the run you want to repair and its matching clone directory
(replace `<run>` with the actual directory name):

```powershell
go run . -mode=fixer -catalog-in=runs/<run>/catalog_kept.json -clone-dir=cloned/github
```

The fixer writes `catalog_fixed.json` and `catalog_fixer_report.md`. It works
against cloned charts and can modify their files. Preserve each fixed catalog
under a distinct name before processing another source.

To analyze the checked-in v2 research snapshot, starting at the repository root:

```powershell
cd taxonomy_analyzer
go run . -input ../helm_fetcher/catalog_sources_v2.json -fixed ../helm_fetcher/catalog_fixed_cumulative_v2.json -out out/cumulative_v2
```

Open `out/cumulative_v2/viewer.html` from that directory. Source and fixed
catalogs must describe the same corpus for meaningful fix correlation.

## Commands

Use `go run . -h` for flag defaults.

| Mode | Purpose | Main inputs / effects |
|---|---|---|
| `github-search-json` | Capture ranked GitHub repositories | `-search-top`, `-search-order`, `-search-out` |
| `artifacthub-search-json` | Capture Artifact Hub repositories | `-search-top`, `-search-out` |
| `full` | Clone, discover charts, build dependencies, and render | `-source`, `-search-in`, `-clone-dir`, `-run-dir` |
| `fixer` | Attempt repairs for recorded failures | `-catalog-in`, `-clone-dir`; writes fixed catalog and report |
| `merge` | Rebuild cumulative outputs from existing runs | Reads `runs/` |
| `refetch-deps` | Retry dependency failures against local clones | Rewrites `-catalog-in`; creates a `.bak` backup |
| `refilter-values` | Apply values-file filtering | Rewrites `-catalog-in`; creates a `.prefilter.bak` backup |
| `render-md` | Regenerate Markdown and CSV sidecars | `-catalog-in` accepts a catalog file or run directory |

`make help` lists convenience targets. Make defaults differ from the CLI:
`SOURCE=github`, `TOP=500`, `KEEP_CLONES=true`, and
`CLONE_DIR=D:/helm_clones_artifacthub`. Override the clone directory for your
machine. `make refix-all` reruns both sources and the analyzer using older June
inputs by default; it does not reproduce the v2 snapshot.

## Code and data layout

| Location | Contents |
|---|---|
| `main.go`, `*_mode.go` | Pipeline and mode entry points |
| `config/` | CLI flags |
| `git/` | Repository discovery and cloning |
| `helm/` | Chart discovery, values combinations, dependencies, rendering |
| `model/`, `exporter/` | Catalog types and JSON / CSV / Markdown export |
| `cmd/listcharts/` | Chart inventory for downstream scans |
| `scripts/` | Maintenance, diagnostics, authentication, and plotting helpers |
| `github_search.json`, `artifacthub_search.json` | Research seed snapshots |
| `catalog_sources_v2.json`, `catalog_fixed_cumulative_v2.json` | Authoritative paired v2 catalogs |
| `runs/` | Timestamped run catalogs and seed snapshots |
| `results/`, `backup/` | Historical results retained for provenance |
| `plots/` | SVG figures |
| `cloned/` | Ignored local repository clones |

Keep historical catalogs separate from the authoritative v2 pair when
interpreting results. [FINAL_RESULTS.md](FINAL_RESULTS.md) explains superseded
snapshots and the dependency-failure early stop that limits chart coverage.
Catalog paths remain stable because scripts and thesis records use them.

## Development and cleanup

```powershell
go test ./...
go build -o helm_fetcher.exe .
```

Binaries, logs, automatic `.bak` files, and local assistant settings are ignored.
Research catalogs, reports, seeds, and plots remain versioned. `make clean-cloned`
deletes the configured clone directory; `make clean-output` deletes selected
root-level outputs. Check `make help` and your paths before using these targets.
