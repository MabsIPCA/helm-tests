# Helm Error Taxonomy Report

Generated at: `2026-05-17 23:36:24 UTC`

Source catalog: `C:\Users\miabs\GolandProjects\helm-tests\helm_fetcher\backup\run_002_274\catalog_by_project.json`

Fixed catalog: `C:\Users\miabs\GolandProjects\helm-tests\helm_fetcher\catalog_fixed.json`

## Totals

| Metric | Value |
|---|---:|
| Repositories | 274 |
| Helm runs | 3325 |
| Template failures | 473 |
| Dependency failures | 106 |
| Classified errors | 579 |
| Unclassified errors | 0 |
| Fix attempts | 373 |
| Fix resolved | 225 |
| Fix unresolved | 148 |

## Taxonomy by Kind

| Kind | Count |
|---|---:|
| `template` | 473 |
| `dependency` | 106 |

## Taxonomy by SubKind

| SubKind | Count | Fix Resolved | Fix Unresolved |
|---|---:|---:|---:|
| `template.nil_pointer` | 328 | 225 | 38 |
| `template.required_value` | 39 | 0 | 34 |
| `dependency.missing_repository` | 37 | 0 | 0 |
| `template.runtime_eval` | 27 | 0 | 26 |
| `dependency.missing_subchart` | 23 | 0 | 0 |
| `dependency.rate_limit` | 18 | 0 | 0 |
| `template.invalid_value` | 16 | 0 | 0 |
| `template.type_mismatch` | 16 | 0 | 10 |
| `dependency.network_dns` | 14 | 0 | 0 |
| `template.values_schema_validation` | 12 | 0 | 12 |
| `template.missing_template` | 11 | 0 | 8 |
| `template.yaml_render` | 7 | 0 | 7 |
| `dependency.lock_file_out_of_sync` | 6 | 0 | 0 |
| `template.dependency_check_failed` | 6 | 0 | 6 |
| `template.library_chart_not_installable` | 6 | 0 | 5 |
| `dependency.repo_update` | 4 | 0 | 0 |
| `dependency.chart_validation` | 3 | 0 | 0 |
| `template.parse_error` | 3 | 0 | 1 |
| `dependency.unpack_error` | 1 | 0 | 0 |
| `template.kube_version_incompatible` | 1 | 0 | 0 |
| `template.values_merge_error` | 1 | 0 | 1 |

