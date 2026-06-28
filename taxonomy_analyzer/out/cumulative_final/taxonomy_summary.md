# Helm Error Taxonomy Report

Generated at: `2026-06-24 18:18:50 UTC`

Source catalog: `C:\Users\miabs\GolandProjects\helm-tests\helm_fetcher\catalog_sources_merged.json`

Fixed catalog: `C:\Users\miabs\GolandProjects\helm-tests\helm_fetcher\catalog_fixed_cumulative.json`

## Totals

| Metric | Value |
|---|---:|
| Repositories | 808 |
| Helm runs | 15691 |
| Template failures | 1545 |
| Dependency failures | 113 |
| Classified errors | 873 |
| Unclassified errors | 105 |
| Fix attempts | 870 |
| Fix resolved | 479 |
| Fix non-reproduced (no fix applied) | 2 |
| Fix unresolved | 389 |

## Taxonomy by Kind

| Kind | Count |
|---|---:|
| `template` | 869 |
| `unknown` | 105 |
| `dependency` | 4 |

## Taxonomy by SubKind

| SubKind | Count | Fix Resolved | Fix Unresolved |
|---|---:|---:|---:|
| `template.nil_pointer` | 340 | 297 | 42 |
| `template.required_value` | 220 | 154 | 66 |
| `template.values_schema_validation` | 122 | 0 | 122 |
| `unknown.unclassified` | 105 | 0 | 1 |
| `template.custom_validation` | 35 | 0 | 35 |
| `template.malformed_yaml` | 31 | 0 | 31 |
| `template.kube_version_incompatible` | 27 | 27 | 0 |
| `template.missing_template` | 23 | 0 | 22 |
| `template.author_assertion` | 22 | 0 | 22 |
| `template.library_chart_not_installable` | 12 | 0 | 12 |
| `template.type_mismatch` | 9 | 1 | 8 |
| `template.unsupported_builtin` | 8 | 0 | 8 |
| `template.runtime_eval` | 7 | 0 | 7 |
| `template.parse_error` | 6 | 0 | 6 |
| `template.chart_metadata` | 4 | 0 | 4 |
| `dependency.missing_repository` | 3 | 0 | 0 |
| `template.invalid_value` | 2 | 0 | 2 |
| `dependency.missing_subchart` | 1 | 0 | 0 |
| `template.dependency_check_failed` | 1 | 0 | 1 |

## Unclassified Samples

- `rancher/rancher` `D:\helm_clones_github\rancher__rancher\chart` `dependency`: exit status 1
- `linode/apl-core` `D:\helm_clones_github\linode__apl-core\chart\chart-index` `dependency`: exit status 1
- `cozystack/cozystack` `D:\helm_clones_github\cozystack__cozystack\packages\apps\bucket` `dependency`: exit status 1
- `grafana/helm-charts` `D:\helm_clones_github\grafana__helm-charts\charts\enterprise-metrics` `dependency`: exit status 1
- `norwoodj/helm-docs` `D:\helm_clones_github\norwoodj__helm-docs\example-charts\custom-template` `dependency`: exit status 1
- `securitybunker/databunker` `D:\helm_clones_github\securitybunker__databunker\charts\databunker` `dependency`: exit status 1
- `WeBankFinTech/Prophecis` `D:\helm_clones_github\WeBankFinTech__Prophecis\install\Prophecis` `dependency`: exit status 1
- `rancher/charts` `D:\helm_clones_github\rancher__charts\charts\epinio\102.0.1+up1.6.2` `dependency`: exit status 1
- `cloudnativeapp/charts` `D:\helm_clones_github\cloudnativeapp__charts\curated\airflow` `dependency`: exit status 1
- `radondb/radondb-mysql-kubernetes` `D:\helm_clones_github\radondb__radondb-mysql-kubernetes\charts\mysql-operator` `dependency`: exit status 1
