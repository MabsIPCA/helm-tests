# Helm Error Taxonomy Report

Generated at: `2026-06-10 17:45:26 UTC`

Source catalog: `C:\Users\miabs\GolandProjects\helm-tests\helm_fetcher\runs\20260529_135859_artifacthub\catalog_by_project.json`

Fixed catalog: `C:\Users\miabs\GolandProjects\helm-tests\helm_fetcher\catalog_fixed_ah_final.json`

## Totals

| Metric | Value |
|---|---:|
| Repositories | 486 |
| Helm runs | 2098 |
| Template failures | 248 |
| Dependency failures | 24 |
| Classified errors | 186 |
| Unclassified errors | 24 |
| Fix attempts | 186 |
| Fix resolved | 23 |
| Fix unresolved | 163 |

## Taxonomy by Kind

| Kind | Count |
|---|---:|
| `template` | 186 |
| `unknown` | 24 |

## Taxonomy by SubKind

| SubKind | Count | Fix Resolved | Fix Unresolved |
|---|---:|---:|---:|
| `template.values_schema_validation` | 79 | 0 | 79 |
| `template.required_value` | 44 | 19 | 25 |
| `unknown.unclassified` | 24 | 0 | 0 |
| `template.custom_validation` | 17 | 0 | 17 |
| `template.author_assertion` | 13 | 0 | 13 |
| `template.missing_template` | 8 | 0 | 8 |
| `template.library_chart_not_installable` | 7 | 0 | 7 |
| `template.malformed_yaml` | 4 | 0 | 4 |
| `template.nil_pointer` | 3 | 3 | 0 |
| `template.chart_metadata` | 2 | 0 | 2 |
| `template.parse_error` | 2 | 0 | 2 |
| `template.type_mismatch` | 2 | 0 | 2 |
| `template.unsupported_builtin` | 2 | 0 | 2 |
| `template.dependency_check_failed` | 1 | 0 | 1 |
| `template.kube_version_incompatible` | 1 | 1 | 0 |
| `template.runtime_eval` | 1 | 0 | 1 |

## Unclassified Samples

- `grafana/helm-charts` `D:\helm_clones_artifacthub\grafana__helm-charts\charts\enterprise-metrics` `dependency`: exit status 1
- `goharbor/harbor` `D:\helm_clones_artifacthub\goharbor__harbor\src\pkg\chart\testdata\harbor-schema1` `dependency`: exit status 1
- `rancher/rancher` `D:\helm_clones_artifacthub\rancher__rancher\chart` `dependency`: exit status 1
- `renovatebot/renovate` `D:\helm_clones_artifacthub\renovatebot__renovate\lib\modules\manager\helmv3\__fixtures__` `dependency`: exit status 1
- `rook/rook` `D:\helm_clones_artifacthub\rook__rook\deploy\charts\rook-ceph` `dependency`: exit status 1
- `kubernetes/charts` `D:\helm_clones_artifacthub\kubernetes__charts\incubator\distribution` `dependency`: exit status 1
- `jfrog/charts` `D:\helm_clones_artifacthub\jfrog__charts\stable\artifactory-cpp-ce` `dependency`: exit status 1
- `gabe565/charts` `D:\helm_clones_artifacthub\gabe565__charts\charts\adguard-home` `dependency`: exit status 1
- `kestra-io/kestra` `D:\helm_clones_artifacthub\kestra-io__kestra\charts\kestra-starter` `dependency`: exit status 1
- `helm/charts` `D:\helm_clones_artifacthub\helm__charts\incubator\distribution` `dependency`: exit status 1
