# Helm Error Taxonomy Report

Generated at: `2026-03-13 00:39:39 UTC`

Source catalog: `C:\Users\miabs\GolandProjects\helm-tests\helm_fetcher\backup\run_002_274\catalog_by_project.json`

## Totals

| Metric | Value |
|---|---:|
| Repositories | 274 |
| Helm runs | 3325 |
| Template failures | 473 |
| Dependency failures | 106 |
| Classified errors | 579 |
| Unclassified errors | 0 |

## Taxonomy by Kind

| Kind | Count |
|---|---:|
| `template` | 473 |
| `dependency` | 106 |

## Taxonomy by SubKind

| SubKind | Count |
|---|---:|
| `template.nil_pointer` | 328 |
| `template.required_value` | 39 |
| `dependency.missing_repository` | 37 |
| `template.runtime_eval` | 27 |
| `dependency.missing_subchart` | 23 |
| `dependency.rate_limit` | 18 |
| `template.invalid_value` | 16 |
| `template.type_mismatch` | 16 |
| `dependency.network_dns` | 14 |
| `template.values_schema_validation` | 12 |
| `template.missing_template` | 11 |
| `template.yaml_render` | 7 |
| `dependency.lock_file_out_of_sync` | 6 |
| `template.dependency_check_failed` | 6 |
| `template.library_chart_not_installable` | 6 |
| `dependency.repo_update` | 4 |
| `dependency.chart_validation` | 3 |
| `template.parse_error` | 3 |
| `dependency.unpack_error` | 1 |
| `template.kube_version_incompatible` | 1 |
| `template.values_merge_error` | 1 |

