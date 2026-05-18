# Helm Error Taxonomy Report

Generated at: `2026-05-17 23:13:27 UTC`

Source catalog: `C:\Users\miabs\GolandProjects\helm-tests\helm_fetcher\catalog_fixed.json`

Fixed catalog: `C:\Users\miabs\GolandProjects\helm-tests\helm_fetcher\catalog_fixed.json`

## Totals

| Metric | Value |
|---|---:|
| Repositories | 55 |
| Helm runs | 1805 |
| Template failures | 373 |
| Dependency failures | 0 |
| Classified errors | 373 |
| Unclassified errors | 0 |
| Fix attempts | 373 |
| Fix resolved | 225 |
| Fix unresolved | 148 |

## Taxonomy by Kind

| Kind | Count |
|---|---:|
| `template` | 373 |

## Taxonomy by SubKind

| SubKind | Count | Fix Resolved | Fix Unresolved |
|---|---:|---:|---:|
| `template.nil_pointer` | 263 | 225 | 38 |
| `template.required_value` | 34 | 0 | 34 |
| `template.runtime_eval` | 26 | 0 | 26 |
| `template.values_schema_validation` | 12 | 0 | 12 |
| `template.type_mismatch` | 10 | 0 | 10 |
| `template.missing_template` | 8 | 0 | 8 |
| `template.yaml_render` | 7 | 0 | 7 |
| `template.dependency_check_failed` | 6 | 0 | 6 |
| `template.library_chart_not_installable` | 5 | 0 | 5 |
| `template.parse_error` | 1 | 0 | 1 |
| `template.values_merge_error` | 1 | 0 | 1 |

