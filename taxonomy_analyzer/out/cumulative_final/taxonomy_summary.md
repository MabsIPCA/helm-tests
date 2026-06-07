# Helm Error Taxonomy Report

Generated at: `2026-06-07 21:12:05 UTC`

Source catalog: `C:\Users\miabs\GolandProjects\helm-tests\helm_fetcher\catalog_sources_merged.json`

Fixed catalog: `C:\Users\miabs\GolandProjects\helm-tests\helm_fetcher\catalog_fixed_cumulative.json`

## Totals

| Metric | Value |
|---|---:|
| Repositories | 808 |
| Helm runs | 8452 |
| Template failures | 1103 |
| Dependency failures | 214 |
| Classified errors | 1295 |
| Unclassified errors | 22 |
| Fix attempts | 1103 |
| Fix resolved | 693 |
| Fix unresolved | 410 |

## Taxonomy by Kind

| Kind | Count |
|---|---:|
| `template` | 1081 |
| `dependency` | 214 |
| `unknown` | 22 |

## Taxonomy by SubKind

| SubKind | Count | Fix Resolved | Fix Unresolved |
|---|---:|---:|---:|
| `template.nil_pointer` | 440 | 399 | 41 |
| `template.required_value` | 222 | 162 | 60 |
| `template.kube_version_incompatible` | 134 | 131 | 3 |
| `dependency.missing_repository` | 126 | 0 | 0 |
| `template.values_schema_validation` | 109 | 0 | 109 |
| `template.malformed_yaml` | 69 | 0 | 69 |
| `template.custom_validation` | 36 | 0 | 36 |
| `dependency.missing_subchart` | 30 | 0 | 0 |
| `template.runtime_eval` | 23 | 0 | 23 |
| `unknown.unclassified` | 22 | 0 | 22 |
| `template.library_chart_not_installable` | 12 | 0 | 12 |
| `template.missing_template` | 12 | 0 | 12 |
| `dependency.chart_validation` | 11 | 0 | 0 |
| `dependency.lock_file_out_of_sync` | 11 | 0 | 0 |
| `dependency.repo_update` | 9 | 0 | 0 |
| `template.unsupported_builtin` | 9 | 0 | 9 |
| `dependency.network_dns` | 7 | 0 | 0 |
| `template.type_mismatch` | 7 | 1 | 6 |
| `dependency.malformed_yaml` | 6 | 0 | 0 |
| `dependency.cache_index_missing` | 4 | 0 | 0 |
| `dependency.rate_limit` | 4 | 0 | 0 |
| `template.chart_metadata` | 4 | 0 | 4 |
| `dependency.unpack_error` | 3 | 0 | 0 |
| `dependency.version_resolution` | 3 | 0 | 0 |
| `template.invalid_value` | 3 | 0 | 3 |
| `template.parse_error` | 1 | 0 | 1 |

## Unclassified Samples

- `opea-project/Enterprise-Inference` `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\genai-gateway` `template`: Error: execution error at (genaigateway/charts/redis/templates/NOTES.txt:216:4):   ⚠ ERROR: Original containers have been substituted for unrecognized ones. Deploying this chart ...
- `dungdm93/shipyard` `D:\helm_clones_github\dungdm93__shipyard\helm\cloudflared` `template`: Error: execution error at (cloudflared/templates/deployment.yaml:50:41): Missing .Values.token  Use --debug flag to render out invalid YAML
- `dungdm93/shipyard` `D:\helm_clones_github\dungdm93__shipyard\helm\datahub` `template`: Error: execution error at (datahub/templates/setup/secret.yaml:7:55): missing 'ebean.host'  Use --debug flag to render out invalid YAML
- `kast-spells/kast-system` `D:\helm_clones_github\kast-spells__kast-system\covenant` `template`: Error: execution error at (covenant/templates/covenant.yaml:39:6): covenant/index.yaml not found in bookrack/test  Use --debug flag to render out invalid YAML
- `influxdata/helm-charts` `D:\helm_clones_artifacthub\influxdata__helm-charts\charts\influxdb3-clustered` `template`: Error: execution error at (influxdb3-clustered/templates/app-instance.yml:90:3): missing catalog.dsn.SecretName  Use --debug flag to render out invalid YAML
- `camunda/camunda-platform-helm` `D:\helm_clones_artifacthub\camunda__camunda-platform-helm\charts\camunda-platform-8.10` `template`: Error: execution error at (camunda-platform/templates/common/constraints.tpl:40:6): Please configure an expected secondary storage type under `orchestration.data.secondaryStorage.t...
- `camunda/camunda-platform-helm` `D:\helm_clones_artifacthub\camunda__camunda-platform-helm\charts\camunda-platform-8.10` `template`: Error: execution error at (camunda-platform/templates/common/constraints.tpl:40:6): Please configure an expected secondary storage type under `orchestration.data.secondaryStorage.t...
- `camunda/camunda-platform-helm` `D:\helm_clones_artifacthub\camunda__camunda-platform-helm\charts\camunda-platform-8.10` `template`: Error: execution error at (camunda-platform/templates/common/constraints.tpl:40:6): Please configure an expected secondary storage type under `orchestration.data.secondaryStorage.t...
- `camunda/camunda-platform-helm` `D:\helm_clones_artifacthub\camunda__camunda-platform-helm\charts\camunda-platform-8.10` `template`: Error: execution error at (camunda-platform/templates/common/constraints.tpl:40:6): Please configure an expected secondary storage type under `orchestration.data.secondaryStorage.t...
- `camunda/camunda-platform-helm` `D:\helm_clones_artifacthub\camunda__camunda-platform-helm\charts\camunda-platform-8.10` `template`: Error: execution error at (camunda-platform/templates/common/constraints.tpl:40:6): Please configure an expected secondary storage type under `orchestration.data.secondaryStorage.t...
