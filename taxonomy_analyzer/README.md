# taxonomy_analyzer

`taxonomy_analyzer` is a Go CLI that reads Helm fetcher catalog JSON files and builds a taxonomy of template and dependency failures.

## Default input

The default input file is:

`../helm_fetcher/catalog_cumulative.json`

You can override it with `-input`.

## Outputs

The analyzer writes these files to the output folder (`out` by default):

- `taxonomy_report.json` (full machine-readable report)
- `taxonomy_summary.md` (human-readable summary)
- `taxonomy_occurrences.csv` (one row per failure occurrence)
- `taxonomy_errors_by_bucket.md` (complete error messages grouped by taxonomy bucket)

## Run

```sh
go run .
```

```sh
# Fix-correlation report. Use a matched input+fixed pair (here the merged
# github+artifacthub catalogs) so every failure correlates by helm_command and
# Fix attempts == Template failures.
go run . -input ../helm_fetcher/catalog_sources_merged.json --fixed ../helm_fetcher/catalog_fixed_cumulative.json -out out/cumulative_final
```

## Make targets

- `make run`
- `make test`
- `make analyze-cumulative`
- `make analyze-github`
- `make analyze-artifacthub`

## Package structure

- `model/` data contracts for input and taxonomy report
- `loader/` streaming JSON decoder for large catalogs
- `classifier/` rule-based error classifier
- `analyzer/` aggregation and counting logic
- `exporter/` JSON, Markdown, CSV exporters

## Current taxonomy examples

Template subkinds:

- `template.nil_pointer`
- `template.parse_error`
- `template.missing_template`
- `template.type_mismatch`

Dependency subkinds:

- `dependency.missing_repository`
- `dependency.lock_file_out_of_sync`
- `dependency.rate_limit`
- `dependency.version_resolution`


