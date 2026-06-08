# Helm Error Taxonomy Report

Generated at: `2026-06-08 21:34:29 UTC`

Source catalog: `C:\Users\miabs\GolandProjects\helm-tests\helm_fetcher\results\github\catalog_by_project.json`

Fixed catalog: `C:\Users\miabs\GolandProjects\helm-tests\helm_fetcher\catalog_fixed_github_final.json`

## Totals

| Metric | Value |
|---|---:|
| Repositories | 322 |
| Helm runs | 15014 |
| Template failures | 2637 |
| Dependency failures | 87 |
| Classified errors | 852 |
| Unclassified errors | 98 |
| Fix attempts | 863 |
| Fix resolved | 474 |
| Fix unresolved | 389 |

## Taxonomy by Kind

| Kind | Count |
|---|---:|
| `template` | 852 |
| `unknown` | 98 |

## Taxonomy by SubKind

| SubKind | Count | Fix Resolved | Fix Unresolved |
|---|---:|---:|---:|
| `template.nil_pointer` | 356 | 312 | 44 |
| `template.malformed_yaml` | 182 | 0 | 182 |
| `template.required_value` | 177 | 135 | 42 |
| `unknown.unclassified` | 98 | 0 | 11 |
| `template.values_schema_validation` | 43 | 0 | 43 |
| `template.kube_version_incompatible` | 26 | 26 | 0 |
| `template.custom_validation` | 19 | 0 | 19 |
| `template.missing_template` | 15 | 0 | 15 |
| `template.type_mismatch` | 7 | 1 | 6 |
| `template.library_chart_not_installable` | 6 | 0 | 6 |
| `template.runtime_eval` | 6 | 0 | 6 |
| `template.unsupported_builtin` | 6 | 0 | 6 |
| `template.parse_error` | 4 | 0 | 4 |
| `template.invalid_value` | 3 | 0 | 3 |
| `template.chart_metadata` | 2 | 0 | 2 |

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
