# Helm Error Taxonomy Report

Generated at: `2026-06-02 00:24:19 UTC`

Source catalog: `C:\Users\miabs\GolandProjects\helm-tests\helm_fetcher\results\github\catalog_by_project.json`

Fixed catalog: `C:\Users\miabs\GolandProjects\helm-tests\helm_fetcher\catalog_fixed_github_final.json`

## Totals

| Metric | Value |
|---|---:|
| Repositories | 322 |
| Helm runs | 7316 |
| Template failures | 1135 |
| Dependency failures | 133 |
| Classified errors | 1079 |
| Unclassified errors | 189 |
| Fix attempts | 1135 |
| Fix resolved | 448 |
| Fix unresolved | 687 |

## Taxonomy by Kind

| Kind | Count |
|---|---:|
| `template` | 956 |
| `unknown` | 189 |
| `dependency` | 123 |

## Taxonomy by SubKind

| SubKind | Count | Fix Resolved | Fix Unresolved |
|---|---:|---:|---:|
| `template.nil_pointer` | 605 | 416 | 189 |
| `unknown.unclassified` | 189 | 3 | 176 |
| `template.kube_version_incompatible` | 140 | 0 | 140 |
| `dependency.missing_repository` | 66 | 0 | 0 |
| `template.required_value` | 64 | 27 | 37 |
| `template.library_chart_not_installable` | 39 | 0 | 39 |
| `template.values_schema_validation` | 37 | 0 | 37 |
| `dependency.missing_subchart` | 25 | 0 | 0 |
| `template.runtime_eval` | 23 | 0 | 23 |
| `template.missing_template` | 19 | 0 | 19 |
| `template.yaml_render` | 18 | 0 | 18 |
| `dependency.lock_file_out_of_sync` | 7 | 0 | 0 |
| `dependency.network_dns` | 7 | 0 | 0 |
| `dependency.repo_update` | 7 | 0 | 0 |
| `template.type_mismatch` | 6 | 2 | 4 |
| `dependency.chart_validation` | 5 | 0 | 0 |
| `dependency.version_resolution` | 3 | 0 | 0 |
| `template.invalid_value` | 3 | 0 | 3 |
| `dependency.unpack_error` | 2 | 0 | 0 |
| `dependency.cache_index_missing` | 1 | 0 | 0 |
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
