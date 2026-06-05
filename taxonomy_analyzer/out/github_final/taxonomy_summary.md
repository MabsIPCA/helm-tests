# Helm Error Taxonomy Report

Generated at: `2026-06-05 07:28:30 UTC`

Source catalog: `C:\Users\miabs\GolandProjects\helm-tests\helm_fetcher\results\github\catalog_by_project.json`

Fixed catalog: `C:\Users\miabs\GolandProjects\helm-tests\helm_fetcher\catalog_fixed_github_final.json`

## Totals

| Metric | Value |
|---|---:|
| Repositories | 322 |
| Helm runs | 6895 |
| Template failures | 916 |
| Dependency failures | 128 |
| Classified errors | 1040 |
| Unclassified errors | 4 |
| Fix attempts | 916 |
| Fix resolved | 664 |
| Fix unresolved | 252 |

## Taxonomy by Kind

| Kind | Count |
|---|---:|
| `template` | 912 |
| `dependency` | 128 |
| `unknown` | 4 |

## Taxonomy by SubKind

| SubKind | Count | Fix Resolved | Fix Unresolved |
|---|---:|---:|---:|
| `template.nil_pointer` | 446 | 394 | 52 |
| `template.required_value` | 172 | 140 | 32 |
| `template.kube_version_incompatible` | 132 | 129 | 3 |
| `template.malformed_yaml` | 65 | 0 | 65 |
| `dependency.missing_repository` | 63 | 0 | 0 |
| `template.values_schema_validation` | 33 | 0 | 33 |
| `dependency.missing_subchart` | 25 | 0 | 0 |
| `template.runtime_eval` | 23 | 0 | 23 |
| `template.custom_validation` | 15 | 0 | 15 |
| `template.missing_template` | 10 | 0 | 10 |
| `dependency.chart_validation` | 9 | 0 | 0 |
| `dependency.lock_file_out_of_sync` | 7 | 0 | 0 |
| `dependency.network_dns` | 7 | 0 | 0 |
| `dependency.repo_update` | 6 | 0 | 0 |
| `dependency.malformed_yaml` | 5 | 0 | 0 |
| `template.library_chart_not_installable` | 5 | 0 | 5 |
| `template.type_mismatch` | 5 | 1 | 4 |
| `unknown.unclassified` | 4 | 0 | 4 |
| `dependency.version_resolution` | 3 | 0 | 0 |
| `template.invalid_value` | 3 | 0 | 3 |
| `dependency.unpack_error` | 2 | 0 | 0 |
| `template.chart_metadata` | 2 | 0 | 2 |
| `dependency.cache_index_missing` | 1 | 0 | 0 |
| `template.parse_error` | 1 | 0 | 1 |

## Unclassified Samples

- `opea-project/Enterprise-Inference` `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\genai-gateway` `template`: Error: execution error at (genaigateway/charts/redis/templates/NOTES.txt:216:4):   ⚠ ERROR: Original containers have been substituted for unrecognized ones. Deploying this chart ...
- `dungdm93/shipyard` `D:\helm_clones_github\dungdm93__shipyard\helm\cloudflared` `template`: Error: execution error at (cloudflared/templates/deployment.yaml:50:41): Missing .Values.token  Use --debug flag to render out invalid YAML
- `dungdm93/shipyard` `D:\helm_clones_github\dungdm93__shipyard\helm\datahub` `template`: Error: execution error at (datahub/templates/setup/secret.yaml:7:55): missing 'ebean.host'  Use --debug flag to render out invalid YAML
- `kast-spells/kast-system` `D:\helm_clones_github\kast-spells__kast-system\covenant` `template`: Error: execution error at (covenant/templates/covenant.yaml:39:6): covenant/index.yaml not found in bookrack/test  Use --debug flag to render out invalid YAML
