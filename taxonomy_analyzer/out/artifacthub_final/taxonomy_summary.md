# Helm Error Taxonomy Report

Generated at: `2026-06-05 07:28:30 UTC`

Source catalog: `C:\Users\miabs\GolandProjects\helm-tests\helm_fetcher\runs\20260529_135859_artifacthub\catalog_by_project.json`

Fixed catalog: `C:\Users\miabs\GolandProjects\helm-tests\helm_fetcher\catalog_fixed_ah_final.json`

## Totals

| Metric | Value |
|---|---:|
| Repositories | 486 |
| Helm runs | 1557 |
| Template failures | 187 |
| Dependency failures | 86 |
| Classified errors | 255 |
| Unclassified errors | 18 |
| Fix attempts | 187 |
| Fix resolved | 27 |
| Fix unresolved | 160 |

## Taxonomy by Kind

| Kind | Count |
|---|---:|
| `template` | 169 |
| `dependency` | 86 |
| `unknown` | 18 |

## Taxonomy by SubKind

| SubKind | Count | Fix Resolved | Fix Unresolved |
|---|---:|---:|---:|
| `template.values_schema_validation` | 76 | 0 | 76 |
| `dependency.missing_repository` | 63 | 0 | 0 |
| `template.required_value` | 50 | 22 | 28 |
| `template.custom_validation` | 21 | 0 | 21 |
| `unknown.unclassified` | 18 | 0 | 18 |
| `template.library_chart_not_installable` | 7 | 0 | 7 |
| `dependency.missing_subchart` | 5 | 0 | 0 |
| `dependency.lock_file_out_of_sync` | 4 | 0 | 0 |
| `dependency.rate_limit` | 4 | 0 | 0 |
| `template.malformed_yaml` | 4 | 0 | 4 |
| `dependency.cache_index_missing` | 3 | 0 | 0 |
| `dependency.repo_update` | 3 | 0 | 0 |
| `template.nil_pointer` | 3 | 3 | 0 |
| `dependency.chart_validation` | 2 | 0 | 0 |
| `template.chart_metadata` | 2 | 0 | 2 |
| `template.kube_version_incompatible` | 2 | 2 | 0 |
| `template.missing_template` | 2 | 0 | 2 |
| `template.type_mismatch` | 2 | 0 | 2 |
| `dependency.malformed_yaml` | 1 | 0 | 0 |
| `dependency.unpack_error` | 1 | 0 | 0 |

## Unclassified Samples

- `influxdata/helm-charts` `D:\helm_clones_artifacthub\influxdata__helm-charts\charts\influxdb3-clustered` `template`: Error: execution error at (influxdb3-clustered/templates/app-instance.yml:90:3): missing catalog.dsn.SecretName  Use --debug flag to render out invalid YAML
- `camunda/camunda-platform-helm` `D:\helm_clones_artifacthub\camunda__camunda-platform-helm\charts\camunda-platform-8.10` `template`: Error: execution error at (camunda-platform/templates/common/constraints.tpl:40:6): Please configure an expected secondary storage type under `orchestration.data.secondaryStorage.t...
- `camunda/camunda-platform-helm` `D:\helm_clones_artifacthub\camunda__camunda-platform-helm\charts\camunda-platform-8.10` `template`: Error: execution error at (camunda-platform/templates/common/constraints.tpl:40:6): Please configure an expected secondary storage type under `orchestration.data.secondaryStorage.t...
- `camunda/camunda-platform-helm` `D:\helm_clones_artifacthub\camunda__camunda-platform-helm\charts\camunda-platform-8.10` `template`: Error: execution error at (camunda-platform/templates/common/constraints.tpl:40:6): Please configure an expected secondary storage type under `orchestration.data.secondaryStorage.t...
- `camunda/camunda-platform-helm` `D:\helm_clones_artifacthub\camunda__camunda-platform-helm\charts\camunda-platform-8.10` `template`: Error: execution error at (camunda-platform/templates/common/constraints.tpl:40:6): Please configure an expected secondary storage type under `orchestration.data.secondaryStorage.t...
- `camunda/camunda-platform-helm` `D:\helm_clones_artifacthub\camunda__camunda-platform-helm\charts\camunda-platform-8.10` `template`: Error: execution error at (camunda-platform/templates/common/constraints.tpl:40:6): Please configure an expected secondary storage type under `orchestration.data.secondaryStorage.t...
- `camunda/camunda-platform-helm` `D:\helm_clones_artifacthub\camunda__camunda-platform-helm\charts\camunda-platform-8.10` `template`: Error: execution error at (camunda-platform/templates/common/constraints.tpl:40:6): Please configure an expected secondary storage type under `orchestration.data.secondaryStorage.t...
- `camunda/camunda-platform-helm` `D:\helm_clones_artifacthub\camunda__camunda-platform-helm\charts\camunda-platform-8.10` `template`: Error: execution error at (camunda-platform/templates/common/constraints.tpl:40:6): Please configure an expected secondary storage type under `orchestration.data.secondaryStorage.t...
- `camunda/camunda-platform-helm` `D:\helm_clones_artifacthub\camunda__camunda-platform-helm\charts\camunda-platform-8.10` `template`: Error: execution error at (camunda-platform/templates/common/constraints.tpl:40:6): Please configure an expected secondary storage type under `orchestration.data.secondaryStorage.t...
- `camunda/camunda-platform-helm` `D:\helm_clones_artifacthub\camunda__camunda-platform-helm\charts\camunda-platform-8.10` `template`: Error: execution error at (camunda-platform/templates/common/constraints.tpl:40:6): Please configure an expected secondary storage type under `orchestration.data.secondaryStorage.t...
