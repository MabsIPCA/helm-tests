# Helm Error Taxonomy Report

Generated at: `2026-07-17 00:00:56 UTC`

Source catalog: `C:\Users\miabs\GolandProjects\helm-tests\helm_fetcher\catalog_sources_v2.json`

Fixed catalog: `C:\Users\miabs\GolandProjects\helm-tests\helm_fetcher\catalog_fixed_cumulative_v2.json`

## Totals

| Metric | Value |
|---|---:|
| Repositories | 973 |
| Helm runs | 9600 |
| Template failures | 2141 |
| Dependency failures | 206 |
| Classified errors | 1388 |
| Unclassified errors | 29 |
| Fix attempts | 1211 |
| Fix resolved | 723 |
| Fix unresolved | 488 |

## Taxonomy by Kind

| Kind | Count |
|---|---:|
| `template` | 1211 |
| `dependency` | 177 |
| `unknown` | 29 |

## Taxonomy by SubKind

| SubKind | Count | Fix Resolved | Fix Unresolved |
|---|---:|---:|---:|
| `template.nil_pointer` | 587 | 540 | 47 |
| `template.required_value` | 263 | 159 | 104 |
| `template.missing_template` | 110 | 0 | 110 |
| `template.values_schema_validation` | 92 | 0 | 92 |
| `dependency.missing_repository` | 84 | 0 | 0 |
| `dependency.missing_subchart` | 41 | 0 | 0 |
| `unknown.unclassified` | 29 | 0 | 0 |
| `template.kube_version_incompatible` | 24 | 24 | 0 |
| `template.malformed_yaml` | 23 | 0 | 23 |
| `template.author_assertion` | 21 | 0 | 21 |
| `template.runtime_eval` | 21 | 0 | 21 |
| `template.custom_validation` | 20 | 0 | 20 |
| `template.type_mismatch` | 15 | 0 | 15 |
| `dependency.chart_validation` | 14 | 0 | 0 |
| `template.library_chart_not_installable` | 11 | 0 | 11 |
| `dependency.lock_file_out_of_sync` | 9 | 0 | 0 |
| `template.parse_error` | 9 | 0 | 9 |
| `template.unsupported_builtin` | 8 | 0 | 8 |
| `dependency.malformed_yaml` | 7 | 0 | 0 |
| `dependency.unpack_error` | 6 | 0 | 0 |
| `dependency.version_resolution` | 5 | 0 | 0 |
| `dependency.rate_limit` | 4 | 0 | 0 |
| `template.chart_metadata` | 4 | 0 | 4 |
| `dependency.network_dns` | 3 | 0 | 0 |
| `dependency.repo_update` | 3 | 0 | 0 |
| `template.invalid_value` | 2 | 0 | 2 |
| `dependency.cache_index_missing` | 1 | 0 | 0 |
| `template.dependency_check_failed` | 1 | 0 | 1 |

## Unclassified Samples

- `neuroxhq/helm-chart-neurox-control` `D:\helm_clones_github\neuroxhq__helm-chart-neurox-control` `dependency`: Saving 6 charts Downloading neurox-control-api from repo oci://ghcr.io/neuroxhq/helm-charts Save error occurred:  could not download oci://ghcr.io/neuroxhq/helm-charts/neurox-contr...
- `wenerme/kube-stub-cluster` `D:\helm_clones_github\wenerme__kube-stub-cluster\keycloak` `dependency`: Saving 1 charts Downloading keycloak from repo oci://dockercr.wener.me/bitnamicharts Save error occurred:  could not download oci://dockercr.wener.me/bitnamicharts/keycloak: failed...
- `hmcts/probate-frontend` `D:\helm_clones_github\hmcts__probate-frontend\charts\probate-frontend` `dependency`: Saving 2 charts Downloading nodejs from repo oci://hmctsprod.azurecr.io/helm Save error occurred:  could not download oci://hmctsprod.azurecr.io/helm/nodejs: failed to perform "Fet...
- `neuroxhq/helm-chart-neurox-workload` `D:\helm_clones_github\neuroxhq__helm-chart-neurox-workload` `dependency`: Saving 2 charts Downloading neurox-workload-agent from repo oci://ghcr.io/neuroxhq/helm-charts Save error occurred:  could not download oci://ghcr.io/neuroxhq/helm-charts/neurox-wo...
- `StopDenBus/helm-charts` `D:\helm_clones_github\StopDenBus__helm-charts\charts\external-dns` `dependency`: Saving 1 charts Downloading external-dns-chart from repo oci://dhi.io Save error occurred:  could not download oci://dhi.io/external-dns-chart: failed to perform "FetchReference" o...
- `grafana/helm-charts` `D:\helm_clones_artifacthub\grafana__helm-charts\charts\enterprise-metrics` `dependency`: exit status 1
- `goharbor/harbor` `D:\helm_clones_artifacthub\goharbor__harbor\src\pkg\chart\testdata\harbor-schema1` `dependency`: exit status 1
- `rancher/rancher` `D:\helm_clones_artifacthub\rancher__rancher\chart` `dependency`: exit status 1
- `renovatebot/renovate` `D:\helm_clones_artifacthub\renovatebot__renovate\lib\modules\manager\helmv3\__fixtures__` `dependency`: exit status 1
- `rook/rook` `D:\helm_clones_artifacthub\rook__rook\deploy\charts\rook-ceph` `dependency`: exit status 1
