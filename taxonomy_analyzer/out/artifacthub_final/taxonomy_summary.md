# Helm Error Taxonomy Report

Generated at: `2026-06-02 00:24:19 UTC`

Source catalog: `C:\Users\miabs\GolandProjects\helm-tests\helm_fetcher\runs\20260529_135859_artifacthub\catalog_by_project.json`

Fixed catalog: `C:\Users\miabs\GolandProjects\helm-tests\helm_fetcher\catalog_fixed_ah_final.json`

## Totals

| Metric | Value |
|---|---:|
| Repositories | 486 |
| Helm runs | 1728 |
| Template failures | 252 |
| Dependency failures | 89 |
| Classified errors | 288 |
| Unclassified errors | 53 |
| Fix attempts | 252 |
| Fix resolved | 38 |
| Fix unresolved | 214 |

## Taxonomy by Kind

| Kind | Count |
|---|---:|
| `template` | 200 |
| `dependency` | 88 |
| `unknown` | 53 |

## Taxonomy by SubKind

| SubKind | Count | Fix Resolved | Fix Unresolved |
|---|---:|---:|---:|
| `template.values_schema_validation` | 78 | 0 | 78 |
| `dependency.missing_repository` | 65 | 0 | 0 |
| `template.required_value` | 54 | 15 | 39 |
| `unknown.unclassified` | 53 | 3 | 49 |
| `template.nil_pointer` | 35 | 20 | 15 |
| `template.library_chart_not_installable` | 11 | 0 | 11 |
| `template.missing_template` | 10 | 0 | 10 |
| `dependency.missing_subchart` | 6 | 0 | 0 |
| `dependency.lock_file_out_of_sync` | 4 | 0 | 0 |
| `dependency.rate_limit` | 4 | 0 | 0 |
| `template.invalid_value` | 4 | 0 | 4 |
| `template.yaml_render` | 4 | 0 | 4 |
| `dependency.cache_index_missing` | 3 | 0 | 0 |
| `dependency.repo_update` | 3 | 0 | 0 |
| `dependency.chart_validation` | 2 | 0 | 0 |
| `template.kube_version_incompatible` | 2 | 0 | 2 |
| `template.type_mismatch` | 2 | 0 | 2 |
| `dependency.unpack_error` | 1 | 0 | 0 |

## Unclassified Samples

- `aws/eks-charts` `D:\helm_clones_artifacthub\aws__eks-charts\stable\aws-load-balancer-controller` `template`: Error: execution error at (aws-load-balancer-controller/templates/deployment.yaml:67:28): Chart cannot be installed without a valid clusterName!  Use --debug flag to render out inv...
- `rancher/rancher` `D:\helm_clones_artifacthub\rancher__rancher\chart` `dependency`: Error: cannot load Chart.yaml: error converting YAML to JSON: yaml: line 4: found character that cannot start any token
- `itzg/minecraft-server-charts` `D:\helm_clones_artifacthub\itzg__minecraft-server-charts\charts\rcon-web-admin` `template`: Error: execution error at (rcon-web-admin/templates/secrets.yaml:15:10): UI password required. Configure it either by using rconWeb.password or rconWeb.passwordExistingSecret  Use ...
- `VictoriaMetrics/helm-charts` `D:\helm_clones_artifacthub\VictoriaMetrics__helm-charts\charts\victoria-logs-agent` `template`: Error: execution error at (victoria-logs-agent/templates/server.yaml:64:20): at least one remoteWrite configuration must be provided  Use --debug flag to render out invalid YAML
- `actions/actions-runner-controller` `D:\helm_clones_artifacthub\actions__actions-runner-controller\charts\gha-runner-scale-set` `template`: Error: execution error at (gha-runner-scale-set/templates/manager_role_binding.yaml:42:11): No gha-rs-controller deployment found using label (app.kubernetes.io/part-of=gha-rs-cont...
- `linkerd/linkerd2` `D:\helm_clones_artifacthub\linkerd__linkerd2\charts\linkerd-control-plane` `template`: Error: execution error at (linkerd-control-plane/templates/identity.yaml:19:21): Please provide the identity issuer certificate  Use --debug flag to render out invalid YAML
- `linkerd/linkerd2` `D:\helm_clones_artifacthub\linkerd__linkerd2\charts\linkerd-control-plane` `template`: Error: execution error at (linkerd-control-plane/templates/identity.yaml:19:21): Please provide the identity issuer certificate  Use --debug flag to render out invalid YAML
- `piraeusdatastore/helm-charts` `D:\helm_clones_artifacthub\piraeusdatastore__helm-charts\charts\linstor-affinity-controller` `template`: level=WARN msg="this chart is deprecated" Error: execution error at (linstor-affinity-controller/templates/deployment.yaml:61:24): Please specify linstor.endpoint, no default URL c...
- `piraeusdatastore/helm-charts` `D:\helm_clones_artifacthub\piraeusdatastore__helm-charts\charts\linstor-scheduler` `template`: level=WARN msg="this chart is deprecated" Error: execution error at (linstor-scheduler/templates/deployment.yaml:91:24): Please specify linstor.endpoint, no default URL could be de...
- `influxdata/helm-charts` `D:\helm_clones_artifacthub\influxdata__helm-charts\charts\influxdb3-clustered` `template`: Error: execution error at (influxdb3-clustered/templates/app-instance.yml:90:3): missing catalog.dsn.SecretName  Use --debug flag to render out invalid YAML
