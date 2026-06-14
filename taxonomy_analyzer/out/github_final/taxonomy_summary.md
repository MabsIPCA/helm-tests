# Helm Error Taxonomy Report

Generated at: `2026-06-10 17:45:26 UTC`

Source catalog: `C:\Users\miabs\GolandProjects\helm-tests\helm_fetcher\results\github\catalog_by_project.json`

Fixed catalog: `C:\Users\miabs\GolandProjects\helm-tests\helm_fetcher\catalog_fixed_github_final.json`

## Totals

| Metric | Value |
|---|---:|
| Repositories | 322 |
| Helm runs | 13593 |
| Template failures | 1297 |
| Dependency failures | 89 |
| Classified errors | 688 |
| Unclassified errors | 86 |
| Fix attempts | 685 |
| Fix resolved | 458 |
| Fix unresolved | 227 |

## Taxonomy by Kind

| Kind | Count |
|---|---:|
| `template` | 684 |
| `unknown` | 86 |
| `dependency` | 4 |

## Taxonomy by SubKind

| SubKind | Count | Fix Resolved | Fix Unresolved |
|---|---:|---:|---:|
| `template.nil_pointer` | 337 | 295 | 42 |
| `template.required_value` | 176 | 135 | 41 |
| `unknown.unclassified` | 86 | 0 | 1 |
| `template.values_schema_validation` | 43 | 0 | 43 |
| `template.malformed_yaml` | 27 | 0 | 27 |
| `template.kube_version_incompatible` | 26 | 26 | 0 |
| `template.custom_validation` | 18 | 0 | 18 |
| `template.missing_template` | 15 | 1 | 14 |
| `template.author_assertion` | 9 | 0 | 9 |
| `template.type_mismatch` | 7 | 1 | 6 |
| `template.library_chart_not_installable` | 6 | 0 | 6 |
| `template.runtime_eval` | 6 | 0 | 6 |
| `template.unsupported_builtin` | 6 | 0 | 6 |
| `template.parse_error` | 4 | 0 | 4 |
| `dependency.missing_repository` | 3 | 0 | 0 |
| `template.chart_metadata` | 2 | 0 | 2 |
| `template.invalid_value` | 2 | 0 | 2 |
| `dependency.missing_subchart` | 1 | 0 | 0 |

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
