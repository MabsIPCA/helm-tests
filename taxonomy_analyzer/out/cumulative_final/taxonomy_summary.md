# Helm Error Taxonomy Report

Generated at: `2026-06-02 00:33:08 UTC`

Source catalog: `C:\Users\miabs\GolandProjects\helm-tests\helm_fetcher\catalog_sources_merged.json`

Fixed catalog: `C:\Users\miabs\GolandProjects\helm-tests\helm_fetcher\catalog_fixed_cumulative.json`

## Totals

| Metric | Value |
|---|---:|
| Repositories | 808 |
| Helm runs | 9044 |
| Template failures | 1387 |
| Dependency failures | 222 |
| Classified errors | 1367 |
| Unclassified errors | 242 |
| Fix attempts | 1387 |
| Fix resolved | 486 |
| Fix unresolved | 901 |

## Taxonomy by Kind

| Kind | Count |
|---|---:|
| `template` | 1156 |
| `unknown` | 242 |
| `dependency` | 211 |

## Taxonomy by SubKind

| SubKind | Count | Fix Resolved | Fix Unresolved |
|---|---:|---:|---:|
| `template.nil_pointer` | 640 | 436 | 204 |
| `unknown.unclassified` | 242 | 6 | 225 |
| `template.kube_version_incompatible` | 142 | 0 | 142 |
| `dependency.missing_repository` | 131 | 0 | 0 |
| `template.required_value` | 118 | 42 | 76 |
| `template.values_schema_validation` | 115 | 0 | 115 |
| `template.library_chart_not_installable` | 50 | 0 | 50 |
| `dependency.missing_subchart` | 31 | 0 | 0 |
| `template.missing_template` | 29 | 0 | 29 |
| `template.runtime_eval` | 23 | 0 | 23 |
| `template.yaml_render` | 22 | 0 | 22 |
| `dependency.lock_file_out_of_sync` | 11 | 0 | 0 |
| `dependency.repo_update` | 10 | 0 | 0 |
| `template.type_mismatch` | 8 | 2 | 6 |
| `dependency.chart_validation` | 7 | 0 | 0 |
| `dependency.network_dns` | 7 | 0 | 0 |
| `template.invalid_value` | 7 | 0 | 7 |
| `dependency.cache_index_missing` | 4 | 0 | 0 |
| `dependency.rate_limit` | 4 | 0 | 0 |
| `dependency.unpack_error` | 3 | 0 | 0 |
| `dependency.version_resolution` | 3 | 0 | 0 |
| `template.parse_error` | 1 | 0 | 1 |
| `template.values_merge_error` | 1 | 0 | 1 |

## Unclassified Samples

- `rancher/rancher` `D:\helm_clones_github\rancher__rancher\chart` `dependency`: Error: cannot load Chart.yaml: error converting YAML to JSON: yaml: line 4: found character that cannot start any token
- `WeBankFinTech/Prophecis` `D:\helm_clones_github\WeBankFinTech__Prophecis\install\Prophecis` `dependency`: Error: cannot load values.yaml: error reading yaml document: invalid Yaml document separator: --END RSA PRIVATE KEY-----"
- `IBM/charts` `D:\helm_clones_github\IBM__charts\community\aqua-enforcer` `template`: Error: execution error at (aqua-enforcer/templates/enforcer-token-secret.yaml:14:13): A valid .Values.enforcerToken entry required!  Use --debug flag to render out invalid YAML
- `IBM/charts` `D:\helm_clones_github\IBM__charts\community\aqua-scanner` `template`: Error: execution error at (aqua-scanner/templates/scanner-deployment.yaml:31:14): Please specify a username associated with the Scanner role!  Use --debug flag to render out invali...
- `open-edge-platform/edge-ai-libraries` `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart` `template`: Error: failed to parse D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\values_ovms.yaml: cannot unmarshal yaml docume...
- `open-edge-platform/edge-ai-libraries` `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart` `template`: Error: failed to parse D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\values_ovms.yaml: cannot unmarshal yaml docume...
- `open-edge-platform/edge-ai-libraries` `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart` `template`: Error: failed to parse D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\values_ovms.yaml: cannot unmarshal yaml docume...
- `open-edge-platform/edge-ai-libraries` `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart` `template`: Error: failed to parse D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\values_ovms.yaml: cannot unmarshal yaml docume...
- `aws-samples/amazon-eks-machine-learning-with-terraform-and-kubeflow` `D:\helm_clones_github\aws-samples__amazon-eks-machine-learning-with-terraform-and-kubeflow\charts\machine-learning\testing\maskrcnn-jupyter` `template`: Error: execution error at (jupyter/templates/jupyter.yaml:27:7): .Values.global.source_cidr required!  Use --debug flag to render out invalid YAML
- `aws-samples/amazon-eks-machine-learning-with-terraform-and-kubeflow` `D:\helm_clones_github\aws-samples__amazon-eks-machine-learning-with-terraform-and-kubeflow\charts\machine-learning\testing\maskrcnn-optimized-jupyter` `template`: Error: execution error at (jupyter/templates/jupyter.yaml:27:7): .Values.global.source_cidr required!  Use --debug flag to render out invalid YAML
