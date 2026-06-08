# Helm Error Taxonomy Report

Generated at: `2026-06-08 21:34:22 UTC`

Source catalog: `C:\Users\miabs\GolandProjects\helm-tests\helm_fetcher\catalog_sources_merged.json`

Fixed catalog: `C:\Users\miabs\GolandProjects\helm-tests\helm_fetcher\catalog_fixed_cumulative.json`

## Totals

| Metric | Value |
|---|---:|
| Repositories | 808 |
| Helm runs | 17112 |
| Template failures | 2885 |
| Dependency failures | 111 |
| Classified errors | 1024 |
| Unclassified errors | 130 |
| Fix attempts | 1048 |
| Fix resolved | 497 |
| Fix unresolved | 551 |

## Taxonomy by Kind

| Kind | Count |
|---|---:|
| `template` | 1024 |
| `unknown` | 130 |

## Taxonomy by SubKind

| SubKind | Count | Fix Resolved | Fix Unresolved |
|---|---:|---:|---:|
| `template.nil_pointer` | 359 | 315 | 44 |
| `template.required_value` | 221 | 154 | 67 |
| `template.malformed_yaml` | 186 | 0 | 186 |
| `unknown.unclassified` | 130 | 0 | 24 |
| `template.values_schema_validation` | 122 | 0 | 122 |
| `template.custom_validation` | 36 | 0 | 36 |
| `template.kube_version_incompatible` | 27 | 27 | 0 |
| `template.missing_template` | 23 | 0 | 23 |
| `template.library_chart_not_installable` | 12 | 0 | 12 |
| `template.type_mismatch` | 9 | 1 | 8 |
| `template.unsupported_builtin` | 8 | 0 | 8 |
| `template.runtime_eval` | 7 | 0 | 7 |
| `template.parse_error` | 6 | 0 | 6 |
| `template.chart_metadata` | 4 | 0 | 4 |
| `template.invalid_value` | 3 | 0 | 3 |
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
