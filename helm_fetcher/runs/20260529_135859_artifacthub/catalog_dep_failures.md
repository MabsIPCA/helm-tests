# Helm Dependency Build Failures

Total dependency failures: **90**

| # | Repository | Chart Path | Error |
|---|------------|------------|-------|
| 1 | [argoproj/argo-helm](https://github.com/argoproj/argo-helm) | `D:\helm_clones_artifacthub\argoproj__argo-helm\charts\argo-cd` | Error: no repository definition for https://dandydeveloper.github.io/charts/. Please add the missing repos via 'helm repo add' |
| 2 | [prometheus-community/helm-charts](https://github.com/prometheus-community/helm-charts) | `D:\helm_clones_github\prometheus-community__helm-charts\charts\prometheus` | Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 3 | [bitnami/charts](https://github.com/bitnami/charts) | `D:\helm_clones_github\bitnami__charts\bitnami\apache` | Saving 1 charts Downloading common from repo oci://registry-1.docker.io/bitnamicharts Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/common: failed to perform "FetchReference" on source: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.31.4": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/common: failed to perform "FetchReference" on source: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.31.4": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 4 | [kubernetes/dashboard](https://github.com/kubernetes/dashboard) | `D:\helm_clones_artifacthub\kubernetes__dashboard\charts\kubernetes-dashboard` | Error: no repository definition for https://kubernetes.github.io/ingress-nginx, https://charts.jetstack.io, https://kubernetes-sigs.github.io/metrics-server/, https://charts.konghq.com. Please add the missing repos via 'helm repo add' |
| 5 | [grafana/helm-charts](https://github.com/grafana/helm-charts) | `D:\helm_clones_github\grafana__helm-charts\charts\enterprise-logs` | Error: no repository definition for https://grafana.github.io/helm-charts, https://charts.min.io/. Please add the missing repos via 'helm repo add' |
| 6 | [goharbor/harbor](https://github.com/goharbor/harbor) | `D:\helm_clones_artifacthub\goharbor__harbor\src\pkg\chart\testdata\harbor-schema1` | Getting updates for unmanaged Helm repositories... ...Unable to get an update from the "https://kubernetes-charts.storage.googleapis.com/" chart repository: 	failed to fetch https://kubernetes-charts.storage.googleapis.com/index.yaml : 403 Forbidden Error: no cached repository for helm-manager-1067d9c6027b8c3f27b49e40521d64be96ea412858d8e45064fa44afd3966ddc found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-1067d9c6027b8c3f27b49e40521d64be96ea412858d8e45064fa44afd3966ddc-index.yaml: O sistema não conseguiu localizar o ficheiro especificado. |
| 7 | [apache/airflow](https://github.com/apache/airflow) | `D:\helm_clones_artifacthub\apache__airflow\chart` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 8 | [elastic/helm-charts](https://github.com/elastic/helm-charts) | `D:\helm_clones_artifacthub\elastic__helm-charts\metricbeat` | Error: no repository definition for https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 9 | [nextcloud/helm](https://github.com/nextcloud/helm) | `D:\helm_clones_artifacthub\nextcloud__helm\charts\nextcloud` | Error: no repository definition for https://collaboraonline.github.io/online. Please add the missing repos via 'helm repo add' |
| 10 | [metallb/metallb](https://github.com/metallb/metallb) | `D:\helm_clones_artifacthub\metallb__metallb\charts\metallb` | Error: no repository definition for https://metallb.github.io/frr-k8s. Please add the missing repos via 'helm repo add' |
| 11 | [rancher/rancher](https://github.com/rancher/rancher) | `D:\helm_clones_artifacthub\rancher__rancher\chart` | Error: cannot load Chart.yaml: error converting YAML to JSON: yaml: line 4: found character that cannot start any token |
| 12 | [kyverno/kyverno](https://github.com/kyverno/kyverno) | `D:\helm_clones_artifacthub\kyverno__kyverno\charts\kyverno` | Error: no repository definition for https://kyverno.github.io/api, https://openreports.github.io/reports-api, https://kyverno.github.io/reports-server/. Please add the missing repos via 'helm repo add' |
| 13 | [SonarSource/helm-chart-sonarqube](https://github.com/SonarSource/helm-chart-sonarqube) | `D:\helm_clones_artifacthub\SonarSource__helm-chart-sonarqube\charts\sonarqube` | Error: no repository definition for https://kubernetes.github.io/ingress-nginx. Please add the missing repos via 'helm repo add' |
| 14 | [cloudnative-pg/charts](https://github.com/cloudnative-pg/charts) | `D:\helm_clones_artifacthub\cloudnative-pg__charts\charts\cloudnative-pg` | Error: no repository definition for https://cloudnative-pg.github.io/grafana-dashboards. Please add the missing repos via 'helm repo add' |
| 15 | [oauth2-proxy/manifests](https://github.com/oauth2-proxy/manifests) | `D:\helm_clones_artifacthub\oauth2-proxy__manifests\helm\oauth2-proxy` | Error: no repository definition for https://dandydeveloper.github.io/charts. Please add the missing repos via 'helm repo add' |
| 16 | [community-charts/helm-charts](https://github.com/community-charts/helm-charts) | `D:\helm_clones_artifacthub\community-charts__helm-charts\charts\kserve` | Error: no repository definition for https://charts.jetstack.io. Please add the missing repos via 'helm repo add' |
| 17 | [jaegertracing/helm-charts](https://github.com/jaegertracing/helm-charts) | `D:\helm_clones_artifacthub\jaegertracing__helm-charts\charts\jaeger` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 18 | [8gears/n8n-helm-chart](https://github.com/8gears/n8n-helm-chart) | `D:\helm_clones_artifacthub\8gears__n8n-helm-chart\charts\n8n` | Error: no repository definition for https://valkey.io/valkey-helm/. Please add the missing repos via 'helm repo add' |
| 19 | [airflow-helm/charts](https://github.com/airflow-helm/charts) | `D:\helm_clones_artifacthub\airflow-helm__charts\charts\airflow` | Error: no repository definition for https://charts.helm.sh/stable, https://charts.helm.sh/stable. Please add the missing repos via 'helm repo add' |
| 20 | [renovatebot/renovate](https://github.com/renovatebot/renovate) | `D:\helm_clones_artifacthub\renovatebot__renovate\lib\modules\manager\helmv3\__fixtures__` | Error: directory D:\helm_clones_artifacthub\renovatebot__renovate\lib\modules\manager\helmv3\dependency_chart\dask not found |
| 21 | [Kong/charts](https://github.com/Kong/charts) | `D:\helm_clones_artifacthub\Kong__charts\charts\ingress` | Error: no repository definition for https://charts.konghq.com, https://charts.konghq.com. Please add the missing repos via 'helm repo add' |
| 22 | [rook/rook](https://github.com/rook/rook) | `D:\helm_clones_artifacthub\rook__rook\deploy\charts\rook-ceph` | Error: error unpacking subchart library in rook-ceph: Chart.yaml file is missing |
| 23 | [codecentric/helm-charts](https://github.com/codecentric/helm-charts) | `D:\helm_clones_artifacthub\codecentric__helm-charts\charts\keycloak` | Error: no repository definition for https://raw.githubusercontent.com/bitnami/charts/archive-full-index/bitnami. Please add the missing repos via 'helm repo add' |
| 24 | [VictoriaMetrics/helm-charts](https://github.com/VictoriaMetrics/helm-charts) | `D:\helm_clones_artifacthub\VictoriaMetrics__helm-charts\charts\victoria-logs-cluster` | Error: no repository definition for https://helm.vector.dev. Please add the missing repos via 'helm repo add' |
| 25 | [kubernetes/charts](https://github.com/kubernetes/charts) | `D:\helm_clones_artifacthub\kubernetes__charts\incubator\distribution` | level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..." Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/, https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add' |
| 26 | [open-webui/helm-charts](https://github.com/open-webui/helm-charts) | `D:\helm_clones_artifacthub\open-webui__helm-charts\charts\open-webui` | Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update' |
| 27 | [CloudPirates-io/helm-charts](https://github.com/CloudPirates-io/helm-charts) | `D:\helm_clones_artifacthub\CloudPirates-io__helm-charts\charts\ghost` | Saving 2 charts Downloading common from repo oci://registry-1.docker.io/cloudpirates Save error occurred:  could not download oci://registry-1.docker.io/cloudpirates/common: failed to perform "FetchReference" on source: GET "https://registry-1.docker.io/v2/cloudpirates/common/manifests/2.2.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/cloudpirates/common: failed to perform "FetchReference" on source: GET "https://registry-1.docker.io/v2/cloudpirates/common/manifests/2.2.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 28 | [atlassian/data-center-helm-charts](https://github.com/atlassian/data-center-helm-charts) | `D:\helm_clones_artifacthub\atlassian__data-center-helm-charts\src\main\charts\bamboo` | Error: no repository definition for https://atlassian.github.io/data-center-helm-charts. Please add the missing repos via 'helm repo add' |
| 29 | [grafana/loki](https://github.com/grafana/loki) | `D:\helm_clones_artifacthub\grafana__loki\production\helm\loki` | Error: no repository definition for https://charts.min.io/, https://grafana.github.io/helm-charts, https://grafana.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 30 | [jp-gouin/helm-openldap](https://github.com/jp-gouin/helm-openldap) | `D:\helm_clones_artifacthub\jp-gouin__helm-openldap` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 31 | [camunda/camunda-platform-helm](https://github.com/camunda/camunda-platform-helm) | `D:\helm_clones_artifacthub\camunda__camunda-platform-helm\charts\camunda-platform-8.3` | Saving 4 charts Dependency identity did not declare a repository. Assuming it exists in the charts directory Downloading elasticsearch from repo oci://registry-1.docker.io/bitnamicharts Pulled: registry-1.docker.io/bitnamicharts/elasticsearch:19.21.2 Digest: sha256:266b3a787798808763d559b3b45af0a14f741f76510dfaae65862f845491f40d Downloading postgresql from repo oci://registry-1.docker.io/bitnamicharts Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/postgresql: failed to perform "FetchReference" on source: GET "https://registry-1.docker.io/v2/bitnamicharts/postgresql/manifests/11.9.13": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/postgresql: failed to perform "FetchReference" on source: GET "https://registry-1.docker.io/v2/bitnamicharts/postgresql/manifests/11.9.13": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 32 | [wiremind/wiremind-helm-charts](https://github.com/wiremind/wiremind-helm-charts) | `D:\helm_clones_github\wiremind__wiremind-helm-charts\charts\druid` | Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://machine424.github.io/kube-hpa-scale-to-zero, https://wiremind.github.io/wiremind-helm-charts, https://wiremind.github.io/wiremind-helm-charts. Please add the missing repos via 'helm repo add' |
| 33 | [bootc/netbox-chart](https://github.com/bootc/netbox-chart) | `D:\helm_clones_artifacthub\bootc__netbox-chart\charts\netbox` | Saving 3 charts Downloading common from repo oci://registry-1.docker.io/bitnamicharts Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/common: failed to perform "FetchReference" on source: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.40.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/common: failed to perform "FetchReference" on source: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.40.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 34 | [OpsMx/spinnaker-helm](https://github.com/OpsMx/spinnaker-helm) | `D:\helm_clones_artifacthub\OpsMx__spinnaker-helm\charts\spinnaker` | level=INFO msg="Warning: Dependency locking is handled in Chart.lock since apiVersion \"v2\". We recommend migrating to Chart.lock." level=INFO msg="Warning: Dependencies are handled in Chart.yaml since apiVersion \"v2\". We recommend migrating dependencies to Chart.yaml." Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 35 | [milvus-io/milvus-helm](https://github.com/milvus-io/milvus-helm) | `D:\helm_clones_artifacthub\milvus-io__milvus-helm\charts\milvus` | Error: no repository definition for https://raw.githubusercontent.com/bitnami/charts/pre-2022/bitnami, https://milvus-io.github.io/milvus-helm, https://pulsar.apache.org/charts, https://raw.githubusercontent.com/bitnami/charts/pre-2022/bitnami, https://raw.githubusercontent.com/bitnami/charts/pre-2022/bitnami. Please add the missing repos via 'helm repo add' |
| 36 | [clearml/clearml-helm-charts](https://github.com/clearml/clearml-helm-charts) | `D:\helm_clones_artifacthub\clearml__clearml-helm-charts\charts\clearml` | Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://helm.elastic.co. Please add the missing repos via 'helm repo add' |
| 37 | [DataDog/helm-charts](https://github.com/DataDog/helm-charts) | `D:\helm_clones_github\DataDog__helm-charts\charts\datadog` | Error: no repository definition for https://helm.datadoghq.com, https://prometheus-community.github.io/helm-charts, https://helm.datadoghq.com, https://helm.datadoghq.com. Please add the missing repos via 'helm repo add' |
| 38 | [WeblateOrg/helm](https://github.com/WeblateOrg/helm) | `D:\helm_clones_artifacthub\WeblateOrg__helm\charts\weblate` | Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 39 | [jouve/charts](https://github.com/jouve/charts) | `D:\helm_clones_artifacthub\jouve__charts\charts\mailpit` | Saving 1 charts Downloading common from repo oci://registry-1.docker.io/bitnamicharts Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/common: failed to perform "FetchReference" on source: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.39.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/common: failed to perform "FetchReference" on source: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.39.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 40 | [groundhog2k/helm-charts](https://github.com/groundhog2k/helm-charts) | `D:\helm_clones_artifacthub\groundhog2k__helm-charts\charts\ghost` | Error: no repository definition for https://groundhog2k.github.io/helm-charts, https://groundhog2k.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 41 | [jfrog/charts](https://github.com/jfrog/charts) | `D:\helm_clones_artifacthub\jfrog__charts\stable\artifactory` | Error: no repository definition for https://charts.jfrog.io/. Please add the missing repos via 'helm repo add' |
| 42 | [SigNoz/charts](https://github.com/SigNoz/charts) | `D:\helm_clones_artifacthub\SigNoz__charts\charts\signoz` | Error: no repository definition for https://charts.signoz.io, https://charts.signoz.io, https://charts.redpanda.com, https://charts.signoz.io. Please add the missing repos via 'helm repo add' |
| 43 | [gabe565/charts](https://github.com/gabe565/charts) | `D:\helm_clones_artifacthub\gabe565__charts\charts\adguard-home` | Error: no repository definition for https://bjw-s.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 44 | [k8sgpt-ai/k8sgpt-operator](https://github.com/k8sgpt-ai/k8sgpt-operator) | `D:\helm_clones_artifacthub\k8sgpt-ai__k8sgpt-operator\chart\operator` | Error: no repository definition for https://charts.k8sgpt.ai/. Please add the missing repos via 'helm repo add' |
| 45 | [kestra-io/kestra](https://github.com/kestra-io/kestra) | `D:\helm_clones_artifacthub\kestra-io__kestra\charts\kestra-starter` | Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update' |
| 46 | [helm/charts](https://github.com/helm/charts) | `D:\helm_clones_artifacthub\helm__charts\incubator\distribution` | level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..." Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/, https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add' |
| 47 | [norwoodj/helm-docs](https://github.com/norwoodj/helm-docs) | `D:\helm_clones_artifacthub\norwoodj__helm-docs\example-charts\custom-template` | Error: no repository definition for @stable. Please add them via 'helm repo add' |
| 48 | [chatwoot/charts](https://github.com/chatwoot/charts) | `D:\helm_clones_artifacthub\chatwoot__charts\charts\chatwoot` | Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 49 | [hirosystems/charts](https://github.com/hirosystems/charts) | `D:\helm_clones_artifacthub\hirosystems__charts\hirosystems\bitcoin-core` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 50 | [evryfs/helm-charts](https://github.com/evryfs/helm-charts) | `D:\helm_clones_artifacthub\evryfs__helm-charts\charts\dependency-track` | Error: no repository definition for https://raw.githubusercontent.com/bitnami/charts/eb5f9a9513d987b519f0ecd732e7031241c50328/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 51 | [k8gb-io/k8gb](https://github.com/k8gb-io/k8gb) | `D:\helm_clones_artifacthub\k8gb-io__k8gb\chart\k8gb` | Error: no repository definition for https://coredns.github.io/helm, https://kubernetes-sigs.github.io/external-dns. Please add the missing repos via 'helm repo add' |
| 52 | [kiwigrid/helm-charts](https://github.com/kiwigrid/helm-charts) | `D:\helm_clones_artifacthub\kiwigrid__helm-charts\charts\ditto-digital-twins` | Getting updates for unmanaged Helm repositories... ...Unable to get an update from the "https://kubernetes-charts.storage.googleapis.com" chart repository: 	failed to fetch https://kubernetes-charts.storage.googleapis.com/index.yaml : 403 Forbidden Error: no cached repository for helm-manager-f57ac438e6d97e3defb8e0378330d3cde64dcb65ed39e560646a0813a34e0581 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-f57ac438e6d97e3defb8e0378330d3cde64dcb65ed39e560646a0813a34e0581-index.yaml: O sistema não conseguiu localizar o ficheiro especificado. |
| 53 | [Flagsmith/flagsmith-charts](https://github.com/Flagsmith/flagsmith-charts) | `D:\helm_clones_artifacthub\Flagsmith__flagsmith-charts\charts\flagsmith` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 54 | [glasskube/operator](https://github.com/glasskube/operator) | `D:\helm_clones_artifacthub\glasskube__operator\charts\glasskube-operator` | Error: no repository definition for https://charts.jetstack.io, https://prometheus-community.github.io/helm-charts, https://mariadb-operator.github.io/mariadb-operator, https://cloudnative-pg.io/charts/, https://charts.min.io/. Please add the missing repos via 'helm repo add' |
| 55 | [oauth2-proxy/oauth2-proxy](https://github.com/oauth2-proxy/oauth2-proxy) | `D:\helm_clones_artifacthub\oauth2-proxy__oauth2-proxy\contrib\local-environment\kubernetes` | Error: no repository definition for https://charts.dexidp.io, https://oauth2-proxy.github.io/manifests, https://conservis.github.io/helm-charts, https://conservis.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 56 | [aws/karpenter](https://github.com/aws/karpenter) | `D:\helm_clones_artifacthub\aws__karpenter\charts\karpenter` | Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update' |
| 57 | [mogenius/renovate-operator](https://github.com/mogenius/renovate-operator) | `D:\helm_clones_artifacthub\mogenius__renovate-operator\charts\renovate-operator` | Error: no repository definition for https://valkey.io/valkey-helm/. Please add the missing repos via 'helm repo add' |
| 58 | [hivemq/helm-charts](https://github.com/hivemq/helm-charts) | `D:\helm_clones_artifacthub\hivemq__helm-charts\charts\hivemq-operator` | Error: no repository definition for https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 59 | [Azure/secrets-store-csi-driver-provider-azure](https://github.com/Azure/secrets-store-csi-driver-provider-azure) | `D:\helm_clones_artifacthub\Azure__secrets-store-csi-driver-provider-azure\charts\csi-secrets-store-provider-azure` | Error: no repository definition for https://kubernetes-sigs.github.io/secrets-store-csi-driver/charts. Please add the missing repos via 'helm repo add' |
| 60 | [asdf2014/druid-helm](https://github.com/asdf2014/druid-helm) | `D:\helm_clones_artifacthub\asdf2014__druid-helm\charts\druid` | Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.helm.sh/stable, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 61 | [one-acre-fund/oaf-public-charts](https://github.com/one-acre-fund/oaf-public-charts) | `D:\helm_clones_artifacthub\one-acre-fund__oaf-public-charts\archive\geonode` | Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 62 | [kubernetes/helm](https://github.com/kubernetes/helm) | `D:\helm_clones_artifacthub\kubernetes__helm\internal\chart\v3\lint\rules\testdata\anotherbadchartfile` | Error: validation: chart.metadata.version "7.2445e+06" is invalid |
| 63 | [helm/helm](https://github.com/helm/helm) | `D:\helm_clones_artifacthub\helm__helm\internal\chart\v3\lint\rules\testdata\anotherbadchartfile` | Error: validation: chart.metadata.version "7.2445e+06" is invalid |
| 64 | [banzaicloud/banzai-charts](https://github.com/banzaicloud/banzai-charts) | `D:\helm_clones_artifacthub\banzaicloud__banzai-charts\argo` | Getting updates for unmanaged Helm repositories... ...Unable to get an update from the "https://kubernetes-charts.storage.googleapis.com/" chart repository: 	failed to fetch https://kubernetes-charts.storage.googleapis.com/index.yaml : 403 Forbidden Error: no cached repository for helm-manager-1067d9c6027b8c3f27b49e40521d64be96ea412858d8e45064fa44afd3966ddc found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-1067d9c6027b8c3f27b49e40521d64be96ea412858d8e45064fa44afd3966ddc-index.yaml: O sistema não conseguiu localizar o ficheiro especificado. |
| 65 | [beeinventor/charts](https://github.com/beeinventor/charts) | `D:\helm_clones_artifacthub\beeinventor__charts\beeinventor\keycloak` | Error: no repository definition for https://charts.bitnami.com/bitnami, https://raw.githubusercontent.com/bitnami/charts/archive-full-index/bitnami. Please add the missing repos via 'helm repo add' |
| 66 | [ortelius/ortelius-charts](https://github.com/ortelius/ortelius-charts) | `D:\helm_clones_artifacthub\ortelius__ortelius-charts\chart\ortelius` | Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update' |
| 67 | [snowplow-devops/helm-charts](https://github.com/snowplow-devops/helm-charts) | `D:\helm_clones_artifacthub\snowplow-devops__helm-charts\charts\avalanche` | Error: no repository definition for https://snowplow-devops.github.io/helm-charts, https://snowplow-devops.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 68 | [enix/helm-charts](https://github.com/enix/helm-charts) | `D:\helm_clones_artifacthub\enix__helm-charts\charts\cnpg-monitoring` | Error: no repository definition for https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 69 | [andrenarchy/helm-charts](https://github.com/andrenarchy/helm-charts) | `D:\helm_clones_artifacthub\andrenarchy__helm-charts\charts\home-assistant` | Error: no repository definition for https://library-charts.k8s-at-home.com, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://andrenarchy.github.io/helm-charts/. Please add the missing repos via 'helm repo add' |
| 70 | [devtron-labs/charts](https://github.com/devtron-labs/charts) | `D:\helm_clones_artifacthub\devtron-labs__charts\charts\cluster-essentials` | Getting updates for unmanaged Helm repositories... ...Successfully got an update from the "https://kubernetes-sigs.github.io/metrics-server/" chart repository ...Successfully got an update from the "https://kubernetes.github.io/autoscaler" chart repository ...Successfully got an update from the "https://helm.devtron.ai" chart repository ...Successfully got an update from the "https://aws.github.io/eks-charts" chart repository ...Successfully got an update from the "https://kedacore.github.io/charts" chart repository ...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository Error: can't get a valid version for 1 subchart(s): "kubernetes-event-exporter" (repository "https://charts.bitnami.com/bitnami", version "1.2.*"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml |
| 71 | [litmuschaos/litmus-helm](https://github.com/litmuschaos/litmus-helm) | `D:\helm_clones_artifacthub\litmuschaos__litmus-helm\charts\litmus` | Error: no repository definition for https://raw.githubusercontent.com/bitnami/charts/archive-full-index/bitnami. Please add the missing repos via 'helm repo add' |
| 72 | [netrisai/charts](https://github.com/netrisai/charts) | `D:\helm_clones_artifacthub\netrisai__charts\charts\netris-controller` | Getting updates for unmanaged Helm repositories... ...Successfully got an update from the "https://raw.githubusercontent.com/rawfilescloud/charts/refs/heads/main" chart repository ...Successfully got an update from the "https://haproxytech.github.io/helm-charts" chart repository ...Successfully got an update from the "https://charts.ntppool.org" chart repository ...Successfully got an update from the "https://raw.githubusercontent.com/bitnami/charts/eb5f9a9513d987b519f0ecd732e7031241c50328/bitnami" chart repository ...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository Error: can't get a valid version for 1 subchart(s): "smtp" (repository "https://charts.ntppool.org", version "1.0.1"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml |
| 73 | [k8s-home-lab/helm-charts](https://github.com/k8s-home-lab/helm-charts) | `D:\helm_clones_github\k8s-home-lab__helm-charts\charts\stable\bazarr` | Error: no repository definition for https://k8s-home-lab.github.io/helm-charts/. Please add the missing repos via 'helm repo add' |
| 74 | [keptn/lifecycle-toolkit](https://github.com/keptn/lifecycle-toolkit) | `D:\helm_clones_artifacthub\keptn__lifecycle-toolkit\keptn-cert-manager\chart` | Error: no repository definition for https://charts.lifecycle.keptn.sh. Please add the missing repos via 'helm repo add' |
| 75 | [fosrl/helm-charts](https://github.com/fosrl/helm-charts) | `D:\helm_clones_artifacthub\fosrl__helm-charts\charts\pangolin` | Error: no repository definition for https://cloudnative-pg.github.io/charts, https://cloudnative-pg.github.io/charts. Please add the missing repos via 'helm repo add' |
| 76 | [relution-io/relution-kubernetes](https://github.com/relution-io/relution-kubernetes) | `D:\helm_clones_artifacthub\relution-io__relution-kubernetes\charts\relution` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 77 | [solarwinds/swi-k8s-opentelemetry-collector](https://github.com/solarwinds/swi-k8s-opentelemetry-collector) | `D:\helm_clones_artifacthub\solarwinds__swi-k8s-opentelemetry-collector\deploy\helm` | Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://open-telemetry.github.io/opentelemetry-helm-charts, https://charts.jetstack.io, https://aquasecurity.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 78 | [rhdh-bot/openshift-helm-charts](https://github.com/rhdh-bot/openshift-helm-charts) | `D:\helm_clones_artifacthub\rhdh-bot__openshift-helm-charts\charts\redhat\redhat\redhat-mysql-persistent\0.0.2\src` | Getting updates for unmanaged Helm repositories... ...Unable to get an update from the "https://github.com/openshift-helm-charts/charts" chart repository: 	failed to fetch https://github.com/openshift-helm-charts/charts/index.yaml : 404 Not Found Error: no cached repository for helm-manager-e8a50bb8edab0b7411123ba1d6ccffca210a3cb9ffb119170c254ad301b6b826 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-e8a50bb8edab0b7411123ba1d6ccffca210a3cb9ffb119170c254ad301b6b826-index.yaml: O sistema não conseguiu localizar o ficheiro especificado. |
| 79 | [intel/helm-charts](https://github.com/intel/helm-charts) | `D:\helm_clones_artifacthub\intel__helm-charts\charts\evi-hbase-0.8.3\charts\hbase` | Error: directory D:\helm_clones_artifacthub\intel__helm-charts\charts\evi-hbase-0.8.3\charts\hadoop not found |
| 80 | [k8s-at-home/charts](https://github.com/k8s-at-home/charts) | `D:\helm_clones_artifacthub\k8s-at-home__charts\charts\stable\firefly-iii` | Getting updates for unmanaged Helm repositories... ...Successfully got an update from the "https://library-charts.k8s-at-home.com" chart repository ...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository Error: can't get a valid version for 1 subchart(s): "mariadb" (repository "https://charts.bitnami.com/bitnami", version "11.0.2"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml |
| 81 | [loeken/helm-charts](https://github.com/loeken/helm-charts) | `D:\helm_clones_artifacthub\loeken__helm-charts\charts\home-assistant` | Error: no repository definition for https://bjw-s-labs.github.io/helm-charts, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 82 | [kube-logging/helm-charts](https://github.com/kube-logging/helm-charts) | `D:\helm_clones_artifacthub\kube-logging__helm-charts\charts\logging-demo` | Error: no repository definition for https://helm.min.io/, https://kube-logging.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 83 | [obeone/charts](https://github.com/obeone/charts) | `D:\helm_clones_artifacthub\obeone__charts\charts\cyberchef` | Error: no repository definition for https://library-charts.k8s-at-home.com. Please add the missing repos via 'helm repo add' |
| 84 | [liranme/redisinsight-secure](https://github.com/liranme/redisinsight-secure) | `D:\helm_clones_artifacthub\liranme__redisinsight-secure\helm\redisinsight-secure` | Error: no repository definition for https://oauth2-proxy.github.io/manifests. Please add the missing repos via 'helm repo add' |
| 85 | [cnieg/helm-charts](https://github.com/cnieg/helm-charts) | `D:\helm_clones_artifacthub\cnieg__helm-charts\charts\clamapi` | Error: no repository definition for https://wiremind.github.io/wiremind-helm-charts. Please add the missing repos via 'helm repo add' |
| 86 | [fonzdm/servarr](https://github.com/fonzdm/servarr) | `D:\helm_clones_artifacthub\fonzdm__servarr\servarr` | Saving 7 charts Downloading sonarr from repo oci://tccr.io/truecharts Save error occurred:  could not download oci://tccr.io/truecharts/sonarr: failed to perform "FetchReference" on source: tccr.io/truecharts/sonarr:21.2.1: not found Error: could not download oci://tccr.io/truecharts/sonarr: failed to perform "FetchReference" on source: tccr.io/truecharts/sonarr:21.2.1: not found |
| 87 | [kubewarden/helm-charts](https://github.com/kubewarden/helm-charts) | `D:\helm_clones_artifacthub\kubewarden__helm-charts\charts\kubewarden-controller` | Error: no repository definition for https://kyverno.github.io/policy-reporter. Please add the missing repos via 'helm repo add' |
| 88 | [openmeterio/openmeter](https://github.com/openmeterio/openmeter) | `D:\helm_clones_artifacthub\openmeterio__openmeter\deploy\charts\openmeter` | Error: no repository definition for https://docs.altinity.com/clickhouse-operator/. Please add the missing repos via 'helm repo add' |
| 89 | [kubewarden/sbomscanner](https://github.com/kubewarden/sbomscanner) | `D:\helm_clones_artifacthub\kubewarden__sbomscanner\charts\sbomscanner` | Error: no repository definition for https://nats-io.github.io/k8s/helm/charts/. Please add the missing repos via 'helm repo add' |
| 90 | [kuoss/venti](https://github.com/kuoss/venti) | `D:\helm_clones_artifacthub\kuoss__venti\hack\venti-stack-dev` | Error: no repository definition for https://kuoss.github.io/helm-charts, https://kuoss.github.io/helm-charts. Please add the missing repos via 'helm repo add' |

## Full Error Details

### `D:\helm_clones_artifacthub\argoproj__argo-helm\charts\argo-cd` — argoproj/argo-helm

```
Error: no repository definition for https://dandydeveloper.github.io/charts/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\prometheus-community__helm-charts\charts\prometheus` — prometheus-community/helm-charts

```
Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\bitnami__charts\bitnami\apache` — bitnami/charts

```
Saving 1 charts
Downloading common from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/common: failed to perform "FetchReference" on source: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.31.4": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/common: failed to perform "FetchReference" on source: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.31.4": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `D:\helm_clones_artifacthub\kubernetes__dashboard\charts\kubernetes-dashboard` — kubernetes/dashboard

```
Error: no repository definition for https://kubernetes.github.io/ingress-nginx, https://charts.jetstack.io, https://kubernetes-sigs.github.io/metrics-server/, https://charts.konghq.com. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\grafana__helm-charts\charts\enterprise-logs` — grafana/helm-charts

```
Error: no repository definition for https://grafana.github.io/helm-charts, https://charts.min.io/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\goharbor__harbor\src\pkg\chart\testdata\harbor-schema1` — goharbor/harbor

```
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "https://kubernetes-charts.storage.googleapis.com/" chart repository:
	failed to fetch https://kubernetes-charts.storage.googleapis.com/index.yaml : 403 Forbidden
Error: no cached repository for helm-manager-1067d9c6027b8c3f27b49e40521d64be96ea412858d8e45064fa44afd3966ddc found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-1067d9c6027b8c3f27b49e40521d64be96ea412858d8e45064fa44afd3966ddc-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### `D:\helm_clones_artifacthub\apache__airflow\chart` — apache/airflow

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\elastic__helm-charts\metricbeat` — elastic/helm-charts

```
Error: no repository definition for https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\nextcloud__helm\charts\nextcloud` — nextcloud/helm

```
Error: no repository definition for https://collaboraonline.github.io/online. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\metallb__metallb\charts\metallb` — metallb/metallb

```
Error: no repository definition for https://metallb.github.io/frr-k8s. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\rancher__rancher\chart` — rancher/rancher

```
Error: cannot load Chart.yaml: error converting YAML to JSON: yaml: line 4: found character that cannot start any token
```

### `D:\helm_clones_artifacthub\kyverno__kyverno\charts\kyverno` — kyverno/kyverno

```
Error: no repository definition for https://kyverno.github.io/api, https://openreports.github.io/reports-api, https://kyverno.github.io/reports-server/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\SonarSource__helm-chart-sonarqube\charts\sonarqube` — SonarSource/helm-chart-sonarqube

```
Error: no repository definition for https://kubernetes.github.io/ingress-nginx. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\cloudnative-pg__charts\charts\cloudnative-pg` — cloudnative-pg/charts

```
Error: no repository definition for https://cloudnative-pg.github.io/grafana-dashboards. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\oauth2-proxy__manifests\helm\oauth2-proxy` — oauth2-proxy/manifests

```
Error: no repository definition for https://dandydeveloper.github.io/charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\community-charts__helm-charts\charts\kserve` — community-charts/helm-charts

```
Error: no repository definition for https://charts.jetstack.io. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\jaegertracing__helm-charts\charts\jaeger` — jaegertracing/helm-charts

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\8gears__n8n-helm-chart\charts\n8n` — 8gears/n8n-helm-chart

```
Error: no repository definition for https://valkey.io/valkey-helm/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\airflow-helm__charts\charts\airflow` — airflow-helm/charts

```
Error: no repository definition for https://charts.helm.sh/stable, https://charts.helm.sh/stable. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\renovatebot__renovate\lib\modules\manager\helmv3\__fixtures__` — renovatebot/renovate

```
Error: directory D:\helm_clones_artifacthub\renovatebot__renovate\lib\modules\manager\helmv3\dependency_chart\dask not found
```

### `D:\helm_clones_artifacthub\Kong__charts\charts\ingress` — Kong/charts

```
Error: no repository definition for https://charts.konghq.com, https://charts.konghq.com. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\rook__rook\deploy\charts\rook-ceph` — rook/rook

```
Error: error unpacking subchart library in rook-ceph: Chart.yaml file is missing
```

### `D:\helm_clones_artifacthub\codecentric__helm-charts\charts\keycloak` — codecentric/helm-charts

```
Error: no repository definition for https://raw.githubusercontent.com/bitnami/charts/archive-full-index/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\VictoriaMetrics__helm-charts\charts\victoria-logs-cluster` — VictoriaMetrics/helm-charts

```
Error: no repository definition for https://helm.vector.dev. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\kubernetes__charts\incubator\distribution` — kubernetes/charts

```
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/, https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\open-webui__helm-charts\charts\open-webui` — open-webui/helm-charts

```
Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update'
```

### `D:\helm_clones_artifacthub\CloudPirates-io__helm-charts\charts\ghost` — CloudPirates-io/helm-charts

```
Saving 2 charts
Downloading common from repo oci://registry-1.docker.io/cloudpirates
Save error occurred:  could not download oci://registry-1.docker.io/cloudpirates/common: failed to perform "FetchReference" on source: GET "https://registry-1.docker.io/v2/cloudpirates/common/manifests/2.2.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/cloudpirates/common: failed to perform "FetchReference" on source: GET "https://registry-1.docker.io/v2/cloudpirates/common/manifests/2.2.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `D:\helm_clones_artifacthub\atlassian__data-center-helm-charts\src\main\charts\bamboo` — atlassian/data-center-helm-charts

```
Error: no repository definition for https://atlassian.github.io/data-center-helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\grafana__loki\production\helm\loki` — grafana/loki

```
Error: no repository definition for https://charts.min.io/, https://grafana.github.io/helm-charts, https://grafana.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\jp-gouin__helm-openldap` — jp-gouin/helm-openldap

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\camunda__camunda-platform-helm\charts\camunda-platform-8.3` — camunda/camunda-platform-helm

```
Saving 4 charts
Dependency identity did not declare a repository. Assuming it exists in the charts directory
Downloading elasticsearch from repo oci://registry-1.docker.io/bitnamicharts
Pulled: registry-1.docker.io/bitnamicharts/elasticsearch:19.21.2
Digest: sha256:266b3a787798808763d559b3b45af0a14f741f76510dfaae65862f845491f40d
Downloading postgresql from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/postgresql: failed to perform "FetchReference" on source: GET "https://registry-1.docker.io/v2/bitnamicharts/postgresql/manifests/11.9.13": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/postgresql: failed to perform "FetchReference" on source: GET "https://registry-1.docker.io/v2/bitnamicharts/postgresql/manifests/11.9.13": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `D:\helm_clones_github\wiremind__wiremind-helm-charts\charts\druid` — wiremind/wiremind-helm-charts

```
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://machine424.github.io/kube-hpa-scale-to-zero, https://wiremind.github.io/wiremind-helm-charts, https://wiremind.github.io/wiremind-helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\bootc__netbox-chart\charts\netbox` — bootc/netbox-chart

```
Saving 3 charts
Downloading common from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/common: failed to perform "FetchReference" on source: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.40.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/common: failed to perform "FetchReference" on source: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.40.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `D:\helm_clones_artifacthub\OpsMx__spinnaker-helm\charts\spinnaker` — OpsMx/spinnaker-helm

```
level=INFO msg="Warning: Dependency locking is handled in Chart.lock since apiVersion \"v2\". We recommend migrating to Chart.lock."
level=INFO msg="Warning: Dependencies are handled in Chart.yaml since apiVersion \"v2\". We recommend migrating dependencies to Chart.yaml."
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\milvus-io__milvus-helm\charts\milvus` — milvus-io/milvus-helm

```
Error: no repository definition for https://raw.githubusercontent.com/bitnami/charts/pre-2022/bitnami, https://milvus-io.github.io/milvus-helm, https://pulsar.apache.org/charts, https://raw.githubusercontent.com/bitnami/charts/pre-2022/bitnami, https://raw.githubusercontent.com/bitnami/charts/pre-2022/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\clearml__clearml-helm-charts\charts\clearml` — clearml/clearml-helm-charts

```
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://helm.elastic.co. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\DataDog__helm-charts\charts\datadog` — DataDog/helm-charts

```
Error: no repository definition for https://helm.datadoghq.com, https://prometheus-community.github.io/helm-charts, https://helm.datadoghq.com, https://helm.datadoghq.com. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\WeblateOrg__helm\charts\weblate` — WeblateOrg/helm

```
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\jouve__charts\charts\mailpit` — jouve/charts

```
Saving 1 charts
Downloading common from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/common: failed to perform "FetchReference" on source: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.39.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/common: failed to perform "FetchReference" on source: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.39.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `D:\helm_clones_artifacthub\groundhog2k__helm-charts\charts\ghost` — groundhog2k/helm-charts

```
Error: no repository definition for https://groundhog2k.github.io/helm-charts, https://groundhog2k.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\jfrog__charts\stable\artifactory` — jfrog/charts

```
Error: no repository definition for https://charts.jfrog.io/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\SigNoz__charts\charts\signoz` — SigNoz/charts

```
Error: no repository definition for https://charts.signoz.io, https://charts.signoz.io, https://charts.redpanda.com, https://charts.signoz.io. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\gabe565__charts\charts\adguard-home` — gabe565/charts

```
Error: no repository definition for https://bjw-s.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\k8sgpt-ai__k8sgpt-operator\chart\operator` — k8sgpt-ai/k8sgpt-operator

```
Error: no repository definition for https://charts.k8sgpt.ai/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\kestra-io__kestra\charts\kestra-starter` — kestra-io/kestra

```
Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update'
```

### `D:\helm_clones_artifacthub\helm__charts\incubator\distribution` — helm/charts

```
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/, https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\norwoodj__helm-docs\example-charts\custom-template` — norwoodj/helm-docs

```
Error: no repository definition for @stable. Please add them via 'helm repo add'
```

### `D:\helm_clones_artifacthub\chatwoot__charts\charts\chatwoot` — chatwoot/charts

```
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\hirosystems__charts\hirosystems\bitcoin-core` — hirosystems/charts

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\evryfs__helm-charts\charts\dependency-track` — evryfs/helm-charts

```
Error: no repository definition for https://raw.githubusercontent.com/bitnami/charts/eb5f9a9513d987b519f0ecd732e7031241c50328/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\k8gb-io__k8gb\chart\k8gb` — k8gb-io/k8gb

```
Error: no repository definition for https://coredns.github.io/helm, https://kubernetes-sigs.github.io/external-dns. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\kiwigrid__helm-charts\charts\ditto-digital-twins` — kiwigrid/helm-charts

```
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "https://kubernetes-charts.storage.googleapis.com" chart repository:
	failed to fetch https://kubernetes-charts.storage.googleapis.com/index.yaml : 403 Forbidden
Error: no cached repository for helm-manager-f57ac438e6d97e3defb8e0378330d3cde64dcb65ed39e560646a0813a34e0581 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-f57ac438e6d97e3defb8e0378330d3cde64dcb65ed39e560646a0813a34e0581-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### `D:\helm_clones_artifacthub\Flagsmith__flagsmith-charts\charts\flagsmith` — Flagsmith/flagsmith-charts

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\glasskube__operator\charts\glasskube-operator` — glasskube/operator

```
Error: no repository definition for https://charts.jetstack.io, https://prometheus-community.github.io/helm-charts, https://mariadb-operator.github.io/mariadb-operator, https://cloudnative-pg.io/charts/, https://charts.min.io/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\oauth2-proxy__oauth2-proxy\contrib\local-environment\kubernetes` — oauth2-proxy/oauth2-proxy

```
Error: no repository definition for https://charts.dexidp.io, https://oauth2-proxy.github.io/manifests, https://conservis.github.io/helm-charts, https://conservis.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\aws__karpenter\charts\karpenter` — aws/karpenter

```
Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update'
```

### `D:\helm_clones_artifacthub\mogenius__renovate-operator\charts\renovate-operator` — mogenius/renovate-operator

```
Error: no repository definition for https://valkey.io/valkey-helm/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\hivemq__helm-charts\charts\hivemq-operator` — hivemq/helm-charts

```
Error: no repository definition for https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\Azure__secrets-store-csi-driver-provider-azure\charts\csi-secrets-store-provider-azure` — Azure/secrets-store-csi-driver-provider-azure

```
Error: no repository definition for https://kubernetes-sigs.github.io/secrets-store-csi-driver/charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\asdf2014__druid-helm\charts\druid` — asdf2014/druid-helm

```
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.helm.sh/stable, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\one-acre-fund__oaf-public-charts\archive\geonode` — one-acre-fund/oaf-public-charts

```
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\kubernetes__helm\internal\chart\v3\lint\rules\testdata\anotherbadchartfile` — kubernetes/helm

```
Error: validation: chart.metadata.version "7.2445e+06" is invalid
```

### `D:\helm_clones_artifacthub\helm__helm\internal\chart\v3\lint\rules\testdata\anotherbadchartfile` — helm/helm

```
Error: validation: chart.metadata.version "7.2445e+06" is invalid
```

### `D:\helm_clones_artifacthub\banzaicloud__banzai-charts\argo` — banzaicloud/banzai-charts

```
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "https://kubernetes-charts.storage.googleapis.com/" chart repository:
	failed to fetch https://kubernetes-charts.storage.googleapis.com/index.yaml : 403 Forbidden
Error: no cached repository for helm-manager-1067d9c6027b8c3f27b49e40521d64be96ea412858d8e45064fa44afd3966ddc found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-1067d9c6027b8c3f27b49e40521d64be96ea412858d8e45064fa44afd3966ddc-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### `D:\helm_clones_artifacthub\beeinventor__charts\beeinventor\keycloak` — beeinventor/charts

```
Error: no repository definition for https://charts.bitnami.com/bitnami, https://raw.githubusercontent.com/bitnami/charts/archive-full-index/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\ortelius__ortelius-charts\chart\ortelius` — ortelius/ortelius-charts

```
Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update'
```

### `D:\helm_clones_artifacthub\snowplow-devops__helm-charts\charts\avalanche` — snowplow-devops/helm-charts

```
Error: no repository definition for https://snowplow-devops.github.io/helm-charts, https://snowplow-devops.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\enix__helm-charts\charts\cnpg-monitoring` — enix/helm-charts

```
Error: no repository definition for https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\andrenarchy__helm-charts\charts\home-assistant` — andrenarchy/helm-charts

```
Error: no repository definition for https://library-charts.k8s-at-home.com, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://andrenarchy.github.io/helm-charts/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\devtron-labs__charts\charts\cluster-essentials` — devtron-labs/charts

```
Getting updates for unmanaged Helm repositories...
...Successfully got an update from the "https://kubernetes-sigs.github.io/metrics-server/" chart repository
...Successfully got an update from the "https://kubernetes.github.io/autoscaler" chart repository
...Successfully got an update from the "https://helm.devtron.ai" chart repository
...Successfully got an update from the "https://aws.github.io/eks-charts" chart repository
...Successfully got an update from the "https://kedacore.github.io/charts" chart repository
...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository
Error: can't get a valid version for 1 subchart(s): "kubernetes-event-exporter" (repository "https://charts.bitnami.com/bitnami", version "1.2.*"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml
```

### `D:\helm_clones_artifacthub\litmuschaos__litmus-helm\charts\litmus` — litmuschaos/litmus-helm

```
Error: no repository definition for https://raw.githubusercontent.com/bitnami/charts/archive-full-index/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\netrisai__charts\charts\netris-controller` — netrisai/charts

```
Getting updates for unmanaged Helm repositories...
...Successfully got an update from the "https://raw.githubusercontent.com/rawfilescloud/charts/refs/heads/main" chart repository
...Successfully got an update from the "https://haproxytech.github.io/helm-charts" chart repository
...Successfully got an update from the "https://charts.ntppool.org" chart repository
...Successfully got an update from the "https://raw.githubusercontent.com/bitnami/charts/eb5f9a9513d987b519f0ecd732e7031241c50328/bitnami" chart repository
...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository
Error: can't get a valid version for 1 subchart(s): "smtp" (repository "https://charts.ntppool.org", version "1.0.1"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml
```

### `D:\helm_clones_github\k8s-home-lab__helm-charts\charts\stable\bazarr` — k8s-home-lab/helm-charts

```
Error: no repository definition for https://k8s-home-lab.github.io/helm-charts/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\keptn__lifecycle-toolkit\keptn-cert-manager\chart` — keptn/lifecycle-toolkit

```
Error: no repository definition for https://charts.lifecycle.keptn.sh. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\fosrl__helm-charts\charts\pangolin` — fosrl/helm-charts

```
Error: no repository definition for https://cloudnative-pg.github.io/charts, https://cloudnative-pg.github.io/charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\relution-io__relution-kubernetes\charts\relution` — relution-io/relution-kubernetes

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\solarwinds__swi-k8s-opentelemetry-collector\deploy\helm` — solarwinds/swi-k8s-opentelemetry-collector

```
Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://open-telemetry.github.io/opentelemetry-helm-charts, https://charts.jetstack.io, https://aquasecurity.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\rhdh-bot__openshift-helm-charts\charts\redhat\redhat\redhat-mysql-persistent\0.0.2\src` — rhdh-bot/openshift-helm-charts

```
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "https://github.com/openshift-helm-charts/charts" chart repository:
	failed to fetch https://github.com/openshift-helm-charts/charts/index.yaml : 404 Not Found
Error: no cached repository for helm-manager-e8a50bb8edab0b7411123ba1d6ccffca210a3cb9ffb119170c254ad301b6b826 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-e8a50bb8edab0b7411123ba1d6ccffca210a3cb9ffb119170c254ad301b6b826-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### `D:\helm_clones_artifacthub\intel__helm-charts\charts\evi-hbase-0.8.3\charts\hbase` — intel/helm-charts

```
Error: directory D:\helm_clones_artifacthub\intel__helm-charts\charts\evi-hbase-0.8.3\charts\hadoop not found
```

### `D:\helm_clones_artifacthub\k8s-at-home__charts\charts\stable\firefly-iii` — k8s-at-home/charts

```
Getting updates for unmanaged Helm repositories...
...Successfully got an update from the "https://library-charts.k8s-at-home.com" chart repository
...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository
Error: can't get a valid version for 1 subchart(s): "mariadb" (repository "https://charts.bitnami.com/bitnami", version "11.0.2"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml
```

### `D:\helm_clones_artifacthub\loeken__helm-charts\charts\home-assistant` — loeken/helm-charts

```
Error: no repository definition for https://bjw-s-labs.github.io/helm-charts, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\kube-logging__helm-charts\charts\logging-demo` — kube-logging/helm-charts

```
Error: no repository definition for https://helm.min.io/, https://kube-logging.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\obeone__charts\charts\cyberchef` — obeone/charts

```
Error: no repository definition for https://library-charts.k8s-at-home.com. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\liranme__redisinsight-secure\helm\redisinsight-secure` — liranme/redisinsight-secure

```
Error: no repository definition for https://oauth2-proxy.github.io/manifests. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\cnieg__helm-charts\charts\clamapi` — cnieg/helm-charts

```
Error: no repository definition for https://wiremind.github.io/wiremind-helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\fonzdm__servarr\servarr` — fonzdm/servarr

```
Saving 7 charts
Downloading sonarr from repo oci://tccr.io/truecharts
Save error occurred:  could not download oci://tccr.io/truecharts/sonarr: failed to perform "FetchReference" on source: tccr.io/truecharts/sonarr:21.2.1: not found
Error: could not download oci://tccr.io/truecharts/sonarr: failed to perform "FetchReference" on source: tccr.io/truecharts/sonarr:21.2.1: not found
```

### `D:\helm_clones_artifacthub\kubewarden__helm-charts\charts\kubewarden-controller` — kubewarden/helm-charts

```
Error: no repository definition for https://kyverno.github.io/policy-reporter. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\openmeterio__openmeter\deploy\charts\openmeter` — openmeterio/openmeter

```
Error: no repository definition for https://docs.altinity.com/clickhouse-operator/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\kubewarden__sbomscanner\charts\sbomscanner` — kubewarden/sbomscanner

```
Error: no repository definition for https://nats-io.github.io/k8s/helm/charts/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_artifacthub\kuoss__venti\hack\venti-stack-dev` — kuoss/venti

```
Error: no repository definition for https://kuoss.github.io/helm-charts, https://kuoss.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

