# Helm Error Taxonomy Report

Generated at: `2026-06-08 21:34:29 UTC`

Source catalog: `C:\Users\miabs\GolandProjects\helm-tests\helm_fetcher\runs\20260529_135859_artifacthub\catalog_by_project.json`

Fixed catalog: `C:\Users\miabs\GolandProjects\helm-tests\helm_fetcher\catalog_fixed_ah_final.json`

## Totals

| Metric | Value |
|---|---:|
| Repositories | 486 |
| Helm runs | 2098 |
| Template failures | 248 |
| Dependency failures | 24 |
| Classified errors | 173 |
| Unclassified errors | 37 |
| Fix attempts | 186 |
| Fix resolved | 23 |
| Fix unresolved | 163 |

## Taxonomy by Kind

| Kind | Count |
|---|---:|
| `template` | 173 |
| `unknown` | 37 |

## Taxonomy by SubKind

| SubKind | Count | Fix Resolved | Fix Unresolved |
|---|---:|---:|---:|
| `template.values_schema_validation` | 79 | 0 | 79 |
| `template.required_value` | 44 | 19 | 25 |
| `unknown.unclassified` | 37 | 0 | 13 |
| `template.custom_validation` | 17 | 0 | 17 |
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
- `VictoriaMetrics/helm-charts` `D:\helm_clones_artifacthub\VictoriaMetrics__helm-charts\charts\victoria-logs-mcp` `template`: Error: execution error at (victoria-logs-mcp/templates/deployment.yaml:54:25): .Values.vl.entrypoint should be set  Use --debug flag to render out invalid YAML
- `VictoriaMetrics/helm-charts` `D:\helm_clones_artifacthub\VictoriaMetrics__helm-charts\charts\victoria-metrics-alert` `template`: Error: execution error at (victoria-metrics-alert/templates/alert-server.yaml:8:6): server.datasource.url datasource URL must be specified  Use --debug flag to render out invalid Y...
- `VictoriaMetrics/helm-charts` `D:\helm_clones_artifacthub\VictoriaMetrics__helm-charts\charts\victoria-metrics-anomaly` `template`: Error: execution error at (victoria-metrics-anomaly/templates/server.yaml:2:4): Pass valid license at .Values.license or .Values.global.license if you have an enterprise license fo...
- `VictoriaMetrics/helm-charts` `D:\helm_clones_artifacthub\VictoriaMetrics__helm-charts\charts\victoria-metrics-gateway` `template`: Error: execution error at (victoria-metrics-gateway/templates/server.yaml:1:4): Pass valid license at .Values.license or .Values.global.license if you have an enterprise license fo...
- `VictoriaMetrics/helm-charts` `D:\helm_clones_artifacthub\VictoriaMetrics__helm-charts\charts\victoria-metrics-mcp` `template`: Error: execution error at (victoria-metrics-mcp/templates/deployment.yaml:64:25): either .Values.vm.cloudAPIKey or .Values.vm.entrypoint should be set  Use --debug flag to render o...
