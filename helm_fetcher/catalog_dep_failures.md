# Helm Dependency Build Failures

Total dependency failures: **112**

| # | Repository | Chart Path | Error |
|---|------------|------------|-------|
| 1 | [prometheus-community/helm-charts](https://github.com/prometheus-community/helm-charts) | `E:\helm_clones_artifacthub\prometheus-community__helm-charts\charts\kube-prometheus-stack` | Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://grafana-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 2 | [argoproj/argo-helm](https://github.com/argoproj/argo-helm) | `E:\helm_clones_artifacthub\argoproj__argo-helm\charts\argo-cd` | Error: no repository definition for https://dandydeveloper.github.io/charts/. Please add the missing repos via 'helm repo add' |
| 3 | [bitnami/charts](https://github.com/bitnami/charts) | `E:\helm_clones_artifacthub\bitnami__charts\bitnami\grafana-loki` | Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 6 charts Downloading grafana-alloy from repo oci://registry-1.docker.io/bitnamicharts Pulled: registry-1.docker.io/bitnamicharts/grafana-alloy:1.0.7 Digest: sha256:ae23a727fe845affa83689d30d3311f46c80fbcb8a217b435bb61d67c11e98eb Downloading memcached from repo oci://registry-1.docker.io/bitnamicharts Pulled: registry-1.docker.io/bitnamicharts/memcached:7.9.7 Digest: sha256:ffb71ea353076658513efe8b46e3f0e6a19653b8b1ca8532f1ba7bb3a490a3e3 Downloading memcached from repo oci://registry-1.docker.io/bitnamicharts Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/memcached: failed to resolve registry-1.docker.io/bitnamicharts/memcached:7.9.7: GET "https://registry-1.docker.io/v2/bitnamicharts/memcached/manifests/7.9.7": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/memcached: failed to resolve registry-1.docker.io/bitnamicharts/memcached:7.9.7: GET "https://registry-1.docker.io/v2/bitnamicharts/memcached/manifests/7.9.7": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 4 | [kubernetes/dashboard](https://github.com/kubernetes/dashboard) | `E:\helm_clones_artifacthub\kubernetes__dashboard\charts\kubernetes-dashboard` | Error: no repository definition for https://kubernetes.github.io/ingress-nginx, https://charts.jetstack.io, https://kubernetes-sigs.github.io/metrics-server/, https://charts.konghq.com. Please add the missing repos via 'helm repo add' |
| 5 | [grafana/helm-charts](https://github.com/grafana/helm-charts) | `E:\helm_clones_artifacthub\grafana__helm-charts\charts\enterprise-logs` | Error: no repository definition for https://grafana.github.io/helm-charts, https://charts.min.io/. Please add the missing repos via 'helm repo add' |
| 6 | [goharbor/harbor](https://github.com/goharbor/harbor) | `E:\helm_clones_artifacthub\goharbor__harbor\src\pkg\chart\testdata\harbor-schema1` | Getting updates for unmanaged Helm repositories... ...Unable to get an update from the "https://kubernetes-charts.storage.googleapis.com/" chart repository: 	failed to fetch https://kubernetes-charts.storage.googleapis.com/index.yaml : 403 Forbidden ...Unable to get an update from the "https://kubernetes-charts.storage.googleapis.com/" chart repository: 	failed to fetch https://kubernetes-charts.storage.googleapis.com/index.yaml : 403 Forbidden Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Error: no cached repository for helm-manager-1067d9c6027b8c3f27b49e40521d64be96ea412858d8e45064fa44afd3966ddc found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-1067d9c6027b8c3f27b49e40521d64be96ea412858d8e45064fa44afd3966ddc-index.yaml: O sistema não conseguiu localizar o ficheiro especificado. |
| 7 | [apache/airflow](https://github.com/apache/airflow) | `E:\helm_clones_artifacthub\apache__airflow\chart` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 8 | [elastic/helm-charts](https://github.com/elastic/helm-charts) | `E:\helm_clones_artifacthub\elastic__helm-charts\metricbeat` | Error: no repository definition for https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 9 | [nextcloud/helm](https://github.com/nextcloud/helm) | `E:\helm_clones_artifacthub\nextcloud__helm\charts\nextcloud` | Error: no repository definition for https://collaboraonline.github.io/online. Please add the missing repos via 'helm repo add' |
| 10 | [metallb/metallb](https://github.com/metallb/metallb) | `E:\helm_clones_artifacthub\metallb__metallb\charts\metallb` | Error: no repository definition for https://metallb.github.io/frr-k8s. Please add the missing repos via 'helm repo add' |
| 11 | [SonarSource/helm-chart-sonarqube](https://github.com/SonarSource/helm-chart-sonarqube) | `E:\helm_clones_artifacthub\SonarSource__helm-chart-sonarqube\charts\sonarqube` | Error: no repository definition for https://kubernetes.github.io/ingress-nginx. Please add the missing repos via 'helm repo add' |
| 12 | [kyverno/kyverno](https://github.com/kyverno/kyverno) | `E:\helm_clones_artifacthub\kyverno__kyverno\charts\kyverno` | Error: no repository definition for https://kyverno.github.io/api, https://openreports.github.io/reports-api, https://kyverno.github.io/reports-server/. Please add the missing repos via 'helm repo add' |
| 13 | [cloudnative-pg/charts](https://github.com/cloudnative-pg/charts) | `E:\helm_clones_artifacthub\cloudnative-pg__charts\charts\cloudnative-pg` | Error: no repository definition for https://cloudnative-pg.github.io/grafana-dashboards. Please add the missing repos via 'helm repo add' |
| 14 | [DataDog/helm-charts](https://github.com/DataDog/helm-charts) | `E:\helm_clones_artifacthub\DataDog__helm-charts\charts\datadog` | Error: no repository definition for https://helm.datadoghq.com, https://prometheus-community.github.io/helm-charts, https://helm.datadoghq.com, https://helm.datadoghq.com. Please add the missing repos via 'helm repo add' |
| 15 | [oauth2-proxy/manifests](https://github.com/oauth2-proxy/manifests) | `E:\helm_clones_artifacthub\oauth2-proxy__manifests\helm\oauth2-proxy` | Error: no repository definition for https://dandydeveloper.github.io/charts. Please add the missing repos via 'helm repo add' |
| 16 | [artifacthub/hub](https://github.com/artifacthub/hub) | `E:\helm_clones_artifacthub\artifacthub__hub\charts\artifact-hub` | Getting updates for unmanaged Helm repositories... ...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 1 charts Downloading postgresql from repo https://charts.bitnami.com/bitnami Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/postgresql: failed to resolve registry-1.docker.io/bitnamicharts/postgresql:18.0.15: GET "https://registry-1.docker.io/v2/bitnamicharts/postgresql/manifests/18.0.15": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/postgresql: failed to resolve registry-1.docker.io/bitnamicharts/postgresql:18.0.15: GET "https://registry-1.docker.io/v2/bitnamicharts/postgresql/manifests/18.0.15": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 17 | [norwoodj/helm-docs](https://github.com/norwoodj/helm-docs) | `E:\helm_clones_artifacthub\norwoodj__helm-docs\example-charts\custom-template` | Error: no repository definition for @stable. Please add them via 'helm repo add' |
| 18 | [community-charts/helm-charts](https://github.com/community-charts/helm-charts) | `E:\helm_clones_artifacthub\community-charts__helm-charts\charts\kserve` | Error: no repository definition for https://charts.jetstack.io. Please add the missing repos via 'helm repo add' |
| 19 | [jaegertracing/helm-charts](https://github.com/jaegertracing/helm-charts) | `E:\helm_clones_artifacthub\jaegertracing__helm-charts\charts\jaeger` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 20 | [8gears/n8n-helm-chart](https://github.com/8gears/n8n-helm-chart) | `E:\helm_clones_artifacthub\8gears__n8n-helm-chart\charts\n8n` | Error: no repository definition for https://valkey.io/valkey-helm/. Please add the missing repos via 'helm repo add' |
| 21 | [airflow-helm/charts](https://github.com/airflow-helm/charts) | `E:\helm_clones_artifacthub\airflow-helm__charts\charts\airflow` | Error: no repository definition for https://charts.helm.sh/stable, https://charts.helm.sh/stable. Please add the missing repos via 'helm repo add' |
| 22 | [renovatebot/renovate](https://github.com/renovatebot/renovate) | `E:\helm_clones_artifacthub\renovatebot__renovate\lib\modules\manager\helmv3\__fixtures__` | Error: directory E:\helm_clones_artifacthub\renovatebot__renovate\lib\modules\manager\helmv3\dependency_chart\dask not found |
| 23 | [DandyDeveloper/charts](https://github.com/DandyDeveloper/charts) | `E:\helm_clones_artifacthub\DandyDeveloper__charts\charts\grafana-agent` | Error: no repository definition for https://helm.releases.hashicorp.com. Please add the missing repos via 'helm repo add' |
| 24 | [rook/rook](https://github.com/rook/rook) | `E:\helm_clones_artifacthub\rook__rook\deploy\charts\rook-ceph` | Error: error unpacking subchart library in rook-ceph: Chart.yaml file is missing |
| 25 | [codecentric/helm-charts](https://github.com/codecentric/helm-charts) | `E:\helm_clones_artifacthub\codecentric__helm-charts\charts\keycloak` | Error: no repository definition for https://raw.githubusercontent.com/bitnami/charts/archive-full-index/bitnami. Please add the missing repos via 'helm repo add' |
| 26 | [VictoriaMetrics/helm-charts](https://github.com/VictoriaMetrics/helm-charts) | `E:\helm_clones_artifacthub\VictoriaMetrics__helm-charts\charts\victoria-logs-agent` | Error: no repository definition for https://victoriametrics.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 27 | [backstage/charts](https://github.com/backstage/charts) | `E:\helm_clones_artifacthub\backstage__charts\charts\backstage` | Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 2 charts Downloading common from repo oci://registry-1.docker.io/bitnamicharts Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/common: failed to resolve registry-1.docker.io/bitnamicharts/common:2.10.0: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.10.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/common: failed to resolve registry-1.docker.io/bitnamicharts/common:2.10.0: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.10.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 28 | [open-webui/helm-charts](https://github.com/open-webui/helm-charts) | `E:\helm_clones_artifacthub\open-webui__helm-charts\charts\open-webui` | Error: no repository definition for https://otwld.github.io/ollama-helm/, https://helm.openwebui.com, https://apache.jfrog.io/artifactory/tika, https://helm.openwebui.com. Please add the missing repos via 'helm repo add' |
| 29 | [grafana/k8s-monitoring-helm](https://github.com/grafana/k8s-monitoring-helm) | `E:\helm_clones_artifacthub\grafana__k8s-monitoring-helm\charts\k8s-monitoring` | Error: no repository definition for https://grafana.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 30 | [CloudPirates-io/helm-charts](https://github.com/CloudPirates-io/helm-charts) | `E:\helm_clones_artifacthub\CloudPirates-io__helm-charts\charts\clusterpirate` | Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 2 charts Downloading common from repo oci://registry-1.docker.io/cloudpirates Save error occurred:  could not download oci://registry-1.docker.io/cloudpirates/common: failed to resolve registry-1.docker.io/cloudpirates/common:2.2.0: GET "https://registry-1.docker.io/v2/cloudpirates/common/manifests/2.2.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/cloudpirates/common: failed to resolve registry-1.docker.io/cloudpirates/common:2.2.0: GET "https://registry-1.docker.io/v2/cloudpirates/common/manifests/2.2.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 31 | [atlassian/data-center-helm-charts](https://github.com/atlassian/data-center-helm-charts) | `E:\helm_clones_artifacthub\atlassian__data-center-helm-charts\src\main\charts\bamboo` | Error: no repository definition for https://atlassian.github.io/data-center-helm-charts. Please add the missing repos via 'helm repo add' |
| 32 | [zitadel/zitadel-charts](https://github.com/zitadel/zitadel-charts) | `E:\helm_clones_artifacthub\zitadel__zitadel-charts\charts\zitadel` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 33 | [jp-gouin/helm-openldap](https://github.com/jp-gouin/helm-openldap) | `E:\helm_clones_artifacthub\jp-gouin__helm-openldap` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 34 | [bootc/netbox-chart](https://github.com/bootc/netbox-chart) | `E:\helm_clones_artifacthub\bootc__netbox-chart\charts\netbox` | Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 3 charts Downloading common from repo oci://registry-1.docker.io/bitnamicharts Pulled: registry-1.docker.io/bitnamicharts/common:2.38.0 Digest: sha256:fd148d63a289816d6af9bb3bd6fbe6317f24918a663ed824ad53f8de004c4234 Downloading postgresql from repo oci://registry-1.docker.io/bitnamicharts Pulled: registry-1.docker.io/bitnamicharts/postgresql:18.5.16 Digest: sha256:2e0eb26439c7f857677d33747b25636059f4189110e462f4a2c0ff73ae5314a7 Downloading valkey from repo oci://registry-1.docker.io/bitnamicharts Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/valkey: failed to resolve registry-1.docker.io/bitnamicharts/valkey:5.4.10: GET "https://registry-1.docker.io/v2/bitnamicharts/valkey/manifests/5.4.10": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/valkey: failed to resolve registry-1.docker.io/bitnamicharts/valkey:5.4.10: GET "https://registry-1.docker.io/v2/bitnamicharts/valkey/manifests/5.4.10": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 35 | [OpsMx/spinnaker-helm](https://github.com/OpsMx/spinnaker-helm) | `E:\helm_clones_artifacthub\OpsMx__spinnaker-helm\charts\spinnaker` | load.go:138: Warning: Dependency locking is handled in Chart.lock since apiVersion "v2". We recommend migrating to Chart.lock. load.go:120: Warning: Dependencies are handled in Chart.yaml since apiVersion "v2". We recommend migrating dependencies to Chart.yaml. Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 36 | [milvus-io/milvus-helm](https://github.com/milvus-io/milvus-helm) | `E:\helm_clones_artifacthub\milvus-io__milvus-helm\charts\milvus` | Error: no repository definition for https://raw.githubusercontent.com/bitnami/charts/pre-2022/bitnami, https://milvus-io.github.io/milvus-helm, https://pulsar.apache.org/charts, https://raw.githubusercontent.com/bitnami/charts/pre-2022/bitnami, https://raw.githubusercontent.com/bitnami/charts/pre-2022/bitnami. Please add the missing repos via 'helm repo add' |
| 37 | [camunda/camunda-platform-helm](https://github.com/camunda/camunda-platform-helm) | `E:\helm_clones_artifacthub\camunda__camunda-platform-helm\charts\camunda-platform-8.3` | Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 4 charts Dependency identity did not declare a repository. Assuming it exists in the charts directory Downloading elasticsearch from repo oci://registry-1.docker.io/bitnamicharts Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/elasticsearch: failed to resolve registry-1.docker.io/bitnamicharts/elasticsearch:19.21.2: GET "https://registry-1.docker.io/v2/bitnamicharts/elasticsearch/manifests/19.21.2": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/elasticsearch: failed to resolve registry-1.docker.io/bitnamicharts/elasticsearch:19.21.2: GET "https://registry-1.docker.io/v2/bitnamicharts/elasticsearch/manifests/19.21.2": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 38 | [WeblateOrg/helm](https://github.com/WeblateOrg/helm) | `E:\helm_clones_artifacthub\WeblateOrg__helm\charts\weblate` | Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 39 | [cowboysysop/charts](https://github.com/cowboysysop/charts) | `E:\helm_clones_artifacthub\cowboysysop__charts\archives\katib` | Error: no repository definition for https://charts.bitnami.com/bitnami/, https://charts.bitnami.com/bitnami/. Please add the missing repos via 'helm repo add' |
| 40 | [deliveryhero/helm-charts](https://github.com/deliveryhero/helm-charts) | `E:\helm_clones_artifacthub\deliveryhero__helm-charts\stable\cachet` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 41 | [groundhog2k/helm-charts](https://github.com/groundhog2k/helm-charts) | `E:\helm_clones_artifacthub\groundhog2k__helm-charts\charts\ghost` | Error: no repository definition for https://groundhog2k.github.io/helm-charts, https://groundhog2k.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 42 | [jouve/charts](https://github.com/jouve/charts) | `E:\helm_clones_artifacthub\jouve__charts\charts\extra` | Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 1 charts Downloading common from repo oci://registry-1.docker.io/bitnamicharts Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/common: failed to resolve registry-1.docker.io/bitnamicharts/common:2.38.0: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.38.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/common: failed to resolve registry-1.docker.io/bitnamicharts/common:2.38.0: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.38.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 43 | [jfrog/charts](https://github.com/jfrog/charts) | `E:\helm_clones_artifacthub\jfrog__charts\stable\artifactory` | Error: no repository definition for https://charts.jfrog.io/. Please add the missing repos via 'helm repo add' |
| 44 | [AdWerx/charts](https://github.com/AdWerx/charts) | `E:\helm_clones_artifacthub\AdWerx__charts\awx` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 45 | [SigNoz/charts](https://github.com/SigNoz/charts) | `E:\helm_clones_artifacthub\SigNoz__charts\charts\signoz` | Error: no repository definition for https://charts.signoz.io, https://charts.signoz.io, https://charts.redpanda.com, https://charts.signoz.io. Please add the missing repos via 'helm repo add' |
| 46 | [gabe565/charts](https://github.com/gabe565/charts) | `E:\helm_clones_artifacthub\gabe565__charts\charts\adguard-home` | Error: no repository definition for https://bjw-s.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 47 | [k8sgpt-ai/k8sgpt-operator](https://github.com/k8sgpt-ai/k8sgpt-operator) | `E:\helm_clones_artifacthub\k8sgpt-ai__k8sgpt-operator\chart\operator` | Error: no repository definition for https://charts.k8sgpt.ai/. Please add the missing repos via 'helm repo add' |
| 48 | [grafana-community/helm-charts](https://github.com/grafana-community/helm-charts) | `E:\helm_clones_artifacthub\grafana-community__helm-charts\charts\loki` | Error: no repository definition for https://charts.min.io/, https://grafana.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 49 | [kubernetes/charts](https://github.com/kubernetes/charts) | `E:\helm_clones_artifacthub\kubernetes__charts\incubator\distribution` | manager.go:122: warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash... Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/, https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add' |
| 50 | [langfuse/langfuse-k8s](https://github.com/langfuse/langfuse-k8s) | `E:\helm_clones_artifacthub\langfuse__langfuse-k8s\charts\langfuse` | Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 5 charts Downloading postgresql from repo oci://registry-1.docker.io/bitnamicharts Pulled: registry-1.docker.io/bitnamicharts/postgresql:16.4.9 Digest: sha256:ab4a0db005ca6a34c16c8ba72e76cb0241c4b856ec0699e956ef6c5f4f661444 Downloading clickhouse from repo oci://registry-1.docker.io/bitnamicharts Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/clickhouse: failed to resolve registry-1.docker.io/bitnamicharts/clickhouse:8.0.5: GET "https://registry-1.docker.io/v2/bitnamicharts/clickhouse/manifests/8.0.5": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/clickhouse: failed to resolve registry-1.docker.io/bitnamicharts/clickhouse:8.0.5: GET "https://registry-1.docker.io/v2/bitnamicharts/clickhouse/manifests/8.0.5": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 51 | [wiremind/wiremind-helm-charts](https://github.com/wiremind/wiremind-helm-charts) | `E:\helm_clones_artifacthub\wiremind__wiremind-helm-charts\charts\druid` | Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://machine424.github.io/kube-hpa-scale-to-zero, https://wiremind.github.io/wiremind-helm-charts, https://wiremind.github.io/wiremind-helm-charts. Please add the missing repos via 'helm repo add' |
| 52 | [getredash/contrib-helm-chart](https://github.com/getredash/contrib-helm-chart) | `E:\helm_clones_artifacthub\getredash__contrib-helm-chart\charts\redash` | Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 53 | [ngrok/ngrok-operator](https://github.com/ngrok/ngrok-operator) | `E:\helm_clones_artifacthub\ngrok__ngrok-operator\helm\ngrok-operator` | Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 2 charts Downloading common from repo oci://registry-1.docker.io/bitnamicharts Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/common: failed to resolve registry-1.docker.io/bitnamicharts/common:2.36.0: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.36.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/common: failed to resolve registry-1.docker.io/bitnamicharts/common:2.36.0: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.36.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 54 | [evryfs/helm-charts](https://github.com/evryfs/helm-charts) | `E:\helm_clones_artifacthub\evryfs__helm-charts\charts\dependency-track` | Error: no repository definition for https://raw.githubusercontent.com/bitnami/charts/eb5f9a9513d987b519f0ecd732e7031241c50328/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 55 | [k8gb-io/k8gb](https://github.com/k8gb-io/k8gb) | `E:\helm_clones_artifacthub\k8gb-io__k8gb\chart\k8gb` | Error: no repository definition for https://coredns.github.io/helm, https://kubernetes-sigs.github.io/external-dns. Please add the missing repos via 'helm repo add' |
| 56 | [localstack/helm-charts](https://github.com/localstack/helm-charts) | `E:\helm_clones_artifacthub\localstack__helm-charts\charts\localstack` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 57 | [percona/percona-helm-charts](https://github.com/percona/percona-helm-charts) | `E:\helm_clones_artifacthub\percona__percona-helm-charts\charts\everest` | Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://victoriametrics.github.io/helm-charts, https://percona.github.io/percona-helm-charts, https://percona.github.io/operator-lifecycle-manager. Please add the missing repos via 'helm repo add' |
| 58 | [Flagsmith/flagsmith-charts](https://github.com/Flagsmith/flagsmith-charts) | `E:\helm_clones_artifacthub\Flagsmith__flagsmith-charts\charts\flagsmith` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 59 | [glasskube/operator](https://github.com/glasskube/operator) | `E:\helm_clones_artifacthub\glasskube__operator\charts\glasskube-operator` | Error: no repository definition for https://charts.jetstack.io, https://prometheus-community.github.io/helm-charts, https://mariadb-operator.github.io/mariadb-operator, https://cloudnative-pg.io/charts/, https://charts.min.io/. Please add the missing repos via 'helm repo add' |
| 60 | [oauth2-proxy/oauth2-proxy](https://github.com/oauth2-proxy/oauth2-proxy) | `E:\helm_clones_artifacthub\oauth2-proxy__oauth2-proxy\contrib\local-environment\kubernetes` | Error: no repository definition for https://charts.dexidp.io, https://oauth2-proxy.github.io/manifests, https://conservis.github.io/helm-charts, https://conservis.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 61 | [aws/karpenter](https://github.com/aws/karpenter) | `E:\helm_clones_artifacthub\aws__karpenter\charts\karpenter` | Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies |
| 62 | [mvisonneau/helm-charts](https://github.com/mvisonneau/helm-charts) | `E:\helm_clones_artifacthub\mvisonneau__helm-charts\charts\gitlab-ci-pipelines-exporter` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 63 | [zilliztech/milvus-helm](https://github.com/zilliztech/milvus-helm) | `E:\helm_clones_artifacthub\zilliztech__milvus-helm\charts\milvus` | Error: no repository definition for https://raw.githubusercontent.com/bitnami/charts/archive-full-index/bitnami, https://zilliztech.github.io/milvus-helm, https://zilliztech.github.io/milvus-helm, https://pulsar.apache.org/charts, https://raw.githubusercontent.com/bitnami/charts/archive-full-index/bitnami, https://zilliztech.github.io/milvus-helm. Please add the missing repos via 'helm repo add' |
| 64 | [sentry-kubernetes/charts](https://github.com/sentry-kubernetes/charts) | `E:\helm_clones_artifacthub\sentry-kubernetes__charts\charts\sentry` | Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 5 charts Downloading memcached from repo oci://registry-1.docker.io/cloudpirates Pulled: registry-1.docker.io/cloudpirates/memcached:0.9.3 Digest: sha256:b5472e0c3f4ce4069b86dd16ba754ea78ef977ae5f9780ca1682557bdb02e94d Downloading redis from repo oci://registry-1.docker.io/bitnamicharts Pulled: registry-1.docker.io/bitnamicharts/redis:17.11.3 Digest: sha256:6472c71c6befe73673f584706a7924f6b57974a510149ed750b6f165ed7d30b2 Downloading kafka from repo oci://registry-1.docker.io/bitnamicharts Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/kafka: failed to resolve registry-1.docker.io/bitnamicharts/kafka:29.3.14: GET "https://registry-1.docker.io/v2/bitnamicharts/kafka/manifests/29.3.14": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/kafka: failed to resolve registry-1.docker.io/bitnamicharts/kafka:29.3.14: GET "https://registry-1.docker.io/v2/bitnamicharts/kafka/manifests/29.3.14": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 65 | [apache/incubator-superset](https://github.com/apache/incubator-superset) | `E:\helm_clones_artifacthub\apache__incubator-superset\helm\superset` | Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 2 charts Downloading postgresql from repo oci://registry-1.docker.io/bitnamicharts Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/postgresql: failed to resolve registry-1.docker.io/bitnamicharts/postgresql:16.7.27: GET "https://registry-1.docker.io/v2/bitnamicharts/postgresql/manifests/16.7.27": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/postgresql: failed to resolve registry-1.docker.io/bitnamicharts/postgresql:16.7.27: GET "https://registry-1.docker.io/v2/bitnamicharts/postgresql/manifests/16.7.27": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 66 | [hivemq/helm-charts](https://github.com/hivemq/helm-charts) | `E:\helm_clones_artifacthub\hivemq__helm-charts\charts\hivemq-operator` | Error: no repository definition for https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 67 | [Azure/secrets-store-csi-driver-provider-azure](https://github.com/Azure/secrets-store-csi-driver-provider-azure) | `E:\helm_clones_artifacthub\Azure__secrets-store-csi-driver-provider-azure\charts\csi-secrets-store-provider-azure` | Error: no repository definition for https://kubernetes-sigs.github.io/secrets-store-csi-driver/charts. Please add the missing repos via 'helm repo add' |
| 68 | [asdf2014/druid-helm](https://github.com/asdf2014/druid-helm) | `E:\helm_clones_artifacthub\asdf2014__druid-helm\charts\druid` | Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.helm.sh/stable, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 69 | [one-acre-fund/oaf-public-charts](https://github.com/one-acre-fund/oaf-public-charts) | `E:\helm_clones_artifacthub\one-acre-fund__oaf-public-charts\archive\geonode` | Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 70 | [kubernetes/helm](https://github.com/kubernetes/helm) | `E:\helm_clones_artifacthub\kubernetes__helm\internal\chart\v3\lint\rules\testdata\anotherbadchartfile` | Error: validation: chart.metadata.version "7.2445e+06" is invalid |
| 71 | [anza-labs/charts](https://github.com/anza-labs/charts) | `E:\helm_clones_artifacthub\anza-labs__charts\deprecated\hosted-control-plane` | Error: no repository definition for https://charts.jetstack.io. Please add the missing repos via 'helm repo add' |
| 72 | [pree/helm-charts](https://github.com/pree/helm-charts) | `E:\helm_clones_artifacthub\pree__helm-charts\charts\home-assistant` | Getting updates for unmanaged Helm repositories... ...Successfully got an update from the "https://bjw-s-labs.github.io/helm-charts/" chart repository ...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository ...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository ...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 4 charts Downloading common from repo https://bjw-s-labs.github.io/helm-charts/ Downloading postgresql from repo https://charts.bitnami.com/bitnami Pulled: registry-1.docker.io/bitnamicharts/postgresql:16.7.27 Digest: sha256:4ed63727a46ff6e6e1cc48b87c75c4b2abe9e5901d38079ba22768895c00de7a Downloading mariadb from repo https://charts.bitnami.com/bitnami Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/mariadb: failed to resolve registry-1.docker.io/bitnamicharts/mariadb:22.0.0: GET "https://registry-1.docker.io/v2/bitnamicharts/mariadb/manifests/22.0.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/mariadb: failed to resolve registry-1.docker.io/bitnamicharts/mariadb:22.0.0: GET "https://registry-1.docker.io/v2/bitnamicharts/mariadb/manifests/22.0.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 73 | [artur9010/charts](https://github.com/artur9010/charts) | `E:\helm_clones_artifacthub\artur9010__charts\charts\cloudflared` | Error: no repository definition for https://kubernetes.github.io/ingress-nginx. Please add the missing repos via 'helm repo add' |
| 74 | [helm/charts](https://github.com/helm/charts) | `E:\helm_clones_artifacthub\helm__charts\incubator\distribution` | manager.go:122: warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash... Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/, https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add' |
| 75 | [helm/helm](https://github.com/helm/helm) | `E:\helm_clones_artifacthub\helm__helm\internal\chart\v3\lint\rules\testdata\anotherbadchartfile` | Error: validation: chart.metadata.version "7.2445e+06" is invalid |
| 76 | [Gradiant/charts](https://github.com/Gradiant/charts) | `E:\helm_clones_artifacthub\Gradiant__charts\charts\hbase` | Getting updates for unmanaged Helm repositories... ...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Error: can't get a valid version for 1 subchart(s): "zookeeper" (repository "https://charts.bitnami.com/bitnami", version "~6.3.0"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml |
| 77 | [banzaicloud/banzai-charts](https://github.com/banzaicloud/banzai-charts) | `E:\helm_clones_artifacthub\banzaicloud__banzai-charts\argo` | Getting updates for unmanaged Helm repositories... ...Unable to get an update from the "https://kubernetes-charts.storage.googleapis.com/" chart repository: 	failed to fetch https://kubernetes-charts.storage.googleapis.com/index.yaml : 403 Forbidden Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Error: no cached repository for helm-manager-1067d9c6027b8c3f27b49e40521d64be96ea412858d8e45064fa44afd3966ddc found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-1067d9c6027b8c3f27b49e40521d64be96ea412858d8e45064fa44afd3966ddc-index.yaml: O sistema não conseguiu localizar o ficheiro especificado. |
| 78 | [beeinventor/charts](https://github.com/beeinventor/charts) | `E:\helm_clones_artifacthub\beeinventor__charts\beeinventor\keycloak` | Error: no repository definition for https://charts.bitnami.com/bitnami, https://raw.githubusercontent.com/bitnami/charts/archive-full-index/bitnami. Please add the missing repos via 'helm repo add' |
| 79 | [fmjstudios/helm](https://github.com/fmjstudios/helm) | `E:\helm_clones_artifacthub\fmjstudios__helm\charts\activepieces` | Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 2 charts Downloading postgresql from repo oci://registry-1.docker.io/bitnamicharts Pulled: registry-1.docker.io/bitnamicharts/postgresql:15.5.38 Digest: sha256:fd220eb22fa79a7c3568928d5eebba7b57ea00a9d10a6134e1d42bc51f0e3346 Downloading redis from repo oci://registry-1.docker.io/bitnamicharts Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:19.5.5: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/19.5.5": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:19.5.5: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/19.5.5": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 80 | [monicahq/helm](https://github.com/monicahq/helm) | `E:\helm_clones_artifacthub\monicahq__helm\charts\monica` | Error: no repository definition for https://meilisearch.github.io/meilisearch-kubernetes. Please add the missing repos via 'helm repo add' |
| 81 | [ortelius/ortelius-charts](https://github.com/ortelius/ortelius-charts) | `E:\helm_clones_artifacthub\ortelius__ortelius-charts\chart\ortelius` | Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies |
| 82 | [snowplow-devops/helm-charts](https://github.com/snowplow-devops/helm-charts) | `E:\helm_clones_artifacthub\snowplow-devops__helm-charts\charts\avalanche` | Error: no repository definition for https://snowplow-devops.github.io/helm-charts, https://snowplow-devops.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 83 | [enix/helm-charts](https://github.com/enix/helm-charts) | `E:\helm_clones_artifacthub\enix__helm-charts\charts\cnpg-monitoring` | Error: no repository definition for https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 84 | [cloudhippie/charts](https://github.com/cloudhippie/charts) | `E:\helm_clones_artifacthub\cloudhippie__charts\stable\yopass` | Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 2 charts Downloading redis from repo oci://registry-1.docker.io/cloudpirates Save error occurred:  could not download oci://registry-1.docker.io/cloudpirates/redis: failed to resolve registry-1.docker.io/cloudpirates/redis:0.26.8: GET "https://registry-1.docker.io/v2/cloudpirates/redis/manifests/0.26.8": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/cloudpirates/redis: failed to resolve registry-1.docker.io/cloudpirates/redis:0.26.8: GET "https://registry-1.docker.io/v2/cloudpirates/redis/manifests/0.26.8": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 85 | [andrenarchy/helm-charts](https://github.com/andrenarchy/helm-charts) | `E:\helm_clones_artifacthub\andrenarchy__helm-charts\charts\home-assistant` | Error: no repository definition for https://library-charts.k8s-at-home.com, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://andrenarchy.github.io/helm-charts/. Please add the missing repos via 'helm repo add' |
| 86 | [netrisai/charts](https://github.com/netrisai/charts) | `E:\helm_clones_artifacthub\netrisai__charts\charts\netris-controller` | Getting updates for unmanaged Helm repositories... ...Successfully got an update from the "https://raw.githubusercontent.com/rawfilescloud/charts/refs/heads/main" chart repository ...Successfully got an update from the "https://haproxytech.github.io/helm-charts" chart repository ...Successfully got an update from the "https://charts.ntppool.org" chart repository ...Successfully got an update from the "https://raw.githubusercontent.com/bitnami/charts/eb5f9a9513d987b519f0ecd732e7031241c50328/bitnami" chart repository ...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository ...Successfully got an update from the "https://raw.githubusercontent.com/bitnami/charts/eb5f9a9513d987b519f0ecd732e7031241c50328/bitnami" chart repository Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Error: can't get a valid version for 1 subchart(s): "smtp" (repository "https://charts.ntppool.org", version "1.0.1"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml |
| 87 | [epinio/helm-charts](https://github.com/epinio/helm-charts) | `E:\helm_clones_artifacthub\epinio__helm-charts\chart\epinio` | Error: no repository definition for https://charts.dexidp.io, https://charts.min.io/, https://s3gw-tech.github.io/s3gw-charts. Please add the missing repos via 'helm repo add' |
| 88 | [keptn/lifecycle-toolkit](https://github.com/keptn/lifecycle-toolkit) | `E:\helm_clones_artifacthub\keptn__lifecycle-toolkit\keptn-cert-manager\chart` | Error: no repository definition for https://charts.lifecycle.keptn.sh. Please add the missing repos via 'helm repo add' |
| 89 | [StrangeBeeCorp/helm-charts](https://github.com/StrangeBeeCorp/helm-charts) | `E:\helm_clones_artifacthub\StrangeBeeCorp__helm-charts\thehive-charts\thehive` | Getting updates for unmanaged Helm repositories... ...Successfully got an update from the "https://charts.min.io" chart repository Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 3 charts Downloading cassandra from repo oci://registry-1.docker.io/bitnamicharts Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/cassandra: failed to resolve registry-1.docker.io/bitnamicharts/cassandra:12.3.11: GET "https://registry-1.docker.io/v2/bitnamicharts/cassandra/manifests/12.3.11": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/cassandra: failed to resolve registry-1.docker.io/bitnamicharts/cassandra:12.3.11: GET "https://registry-1.docker.io/v2/bitnamicharts/cassandra/manifests/12.3.11": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 90 | [redhat-developer/rhdh-chart](https://github.com/redhat-developer/rhdh-chart) | `E:\helm_clones_artifacthub\redhat-developer__rhdh-chart\charts\backstage` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 91 | [adfinis/helm-charts](https://github.com/adfinis/helm-charts) | `E:\helm_clones_artifacthub\adfinis__helm-charts\charts\argo-apps` | Error: no repository definition for https://charts.adfinis.com. Please add the missing repos via 'helm repo add' |
| 92 | [apache/doris-operator](https://github.com/apache/doris-operator) | `E:\helm_clones_artifacthub\apache__doris-operator\helm-charts\doris` | Error: cannot load Chart.yaml: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal object into Go struct field Metadata.version of type string |
| 93 | [intel/helm-charts](https://github.com/intel/helm-charts) | `E:\helm_clones_artifacthub\intel__helm-charts\charts\evi-hbase-0.8.3\charts\hbase` | Error: directory E:\helm_clones_artifacthub\intel__helm-charts\charts\evi-hbase-0.8.3\charts\hadoop not found |
| 94 | [k8s-at-home/charts](https://github.com/k8s-at-home/charts) | `E:\helm_clones_artifacthub\k8s-at-home__charts\charts\stable\firefly-iii` | Getting updates for unmanaged Helm repositories... ...Successfully got an update from the "https://library-charts.k8s-at-home.com" chart repository ...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository ...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository ...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Error: can't get a valid version for 1 subchart(s): "mariadb" (repository "https://charts.bitnami.com/bitnami", version "11.0.2"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml |
| 95 | [netbox-community/netbox-chart](https://github.com/netbox-community/netbox-chart) | `E:\helm_clones_artifacthub\netbox-community__netbox-chart\charts\netbox` | Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 3 charts Downloading common from repo oci://registry-1.docker.io/bitnamicharts Pulled: registry-1.docker.io/bitnamicharts/common:2.38.0 Digest: sha256:fd148d63a289816d6af9bb3bd6fbe6317f24918a663ed824ad53f8de004c4234 Downloading postgresql from repo oci://registry-1.docker.io/bitnamicharts Pulled: registry-1.docker.io/bitnamicharts/postgresql:18.5.16 Digest: sha256:2e0eb26439c7f857677d33747b25636059f4189110e462f4a2c0ff73ae5314a7 Downloading valkey from repo oci://registry-1.docker.io/bitnamicharts Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/valkey: failed to resolve registry-1.docker.io/bitnamicharts/valkey:5.4.10: GET "https://registry-1.docker.io/v2/bitnamicharts/valkey/manifests/5.4.10": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/valkey: failed to resolve registry-1.docker.io/bitnamicharts/valkey:5.4.10: GET "https://registry-1.docker.io/v2/bitnamicharts/valkey/manifests/5.4.10": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 96 | [obeone/charts](https://github.com/obeone/charts) | `E:\helm_clones_artifacthub\obeone__charts\charts\cyberchef` | Error: no repository definition for https://library-charts.k8s-at-home.com. Please add the missing repos via 'helm repo add' |
| 97 | [icoretech/helm](https://github.com/icoretech/helm) | `E:\helm_clones_artifacthub\icoretech__helm\charts\tolgee` | Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 1 charts Downloading postgres from repo oci://registry-1.docker.io/cloudpirates Save error occurred:  could not download oci://registry-1.docker.io/cloudpirates/postgres: failed to resolve registry-1.docker.io/cloudpirates/postgres:0.17.2: GET "https://registry-1.docker.io/v2/cloudpirates/postgres/manifests/0.17.2": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/cloudpirates/postgres: failed to resolve registry-1.docker.io/cloudpirates/postgres:0.17.2: GET "https://registry-1.docker.io/v2/cloudpirates/postgres/manifests/0.17.2": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 98 | [liranme/redisinsight-secure](https://github.com/liranme/redisinsight-secure) | `E:\helm_clones_artifacthub\liranme__redisinsight-secure\helm\redisinsight-secure` | Error: no repository definition for https://oauth2-proxy.github.io/manifests. Please add the missing repos via 'helm repo add' |
| 99 | [kir4h/charts](https://github.com/kir4h/charts) | `E:\helm_clones_artifacthub\kir4h__charts\charts\process-exporter` | Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 1 charts Downloading common from repo oci://registry-1.docker.io/bitnamicharts Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/common: failed to resolve registry-1.docker.io/bitnamicharts/common:2.30.0: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.30.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/common: failed to resolve registry-1.docker.io/bitnamicharts/common:2.30.0: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.30.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 100 | [ectobit/charts](https://github.com/ectobit/charts) | `E:\helm_clones_artifacthub\ectobit__charts\rspamd` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 101 | [rstudio/helm](https://github.com/rstudio/helm) | `E:\helm_clones_artifacthub\rstudio__helm\charts\rstudio-connect` | Error: no repository definition for https://helm.rstudio.com. Please add the missing repos via 'helm repo add' |
| 102 | [hpe-storage/truenas-csp](https://github.com/hpe-storage/truenas-csp) | `E:\helm_clones_artifacthub\hpe-storage__truenas-csp\helm\charts\truenas-csp` | Error: no repository definition for https://hpe-storage.github.io/co-deployments. Please add the missing repos via 'helm repo add' |
| 103 | [christianhuth/helm-charts](https://github.com/christianhuth/helm-charts) | `E:\helm_clones_artifacthub\christianhuth__helm-charts\charts\baserow` | Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 2 charts Downloading postgresql from repo oci://registry-1.docker.io/bitnamicharts Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/postgresql: failed to resolve registry-1.docker.io/bitnamicharts/postgresql:18.5.16: GET "https://registry-1.docker.io/v2/bitnamicharts/postgresql/manifests/18.5.16": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/postgresql: failed to resolve registry-1.docker.io/bitnamicharts/postgresql:18.5.16: GET "https://registry-1.docker.io/v2/bitnamicharts/postgresql/manifests/18.5.16": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 104 | [sitewhere/sitewhere-k8s](https://github.com/sitewhere/sitewhere-k8s) | `E:\helm_clones_artifacthub\sitewhere__sitewhere-k8s\charts\sitewhere` | Error: cannot load Chart.yaml: error converting YAML to JSON: yaml: line 2: mapping values are not allowed in this context |
| 105 | [AvistoTelecom/charts](https://github.com/AvistoTelecom/charts) | `E:\helm_clones_artifacthub\AvistoTelecom__charts\charts\cloudbeaver` | Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 1 charts Downloading common from repo oci://registry-1.docker.io/bitnamicharts Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/common: failed to resolve registry-1.docker.io/bitnamicharts/common:2.36.0: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.36.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/common: failed to resolve registry-1.docker.io/bitnamicharts/common:2.36.0: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.36.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 106 | [epam/ai-dial-helm](https://github.com/epam/ai-dial-helm) | `E:\helm_clones_artifacthub\epam__ai-dial-helm\charts\dial` | Getting updates for unmanaged Helm repositories... ...Successfully got an update from the "https://charts.dialx.ai" chart repository ...Successfully got an update from the "https://charts.dialx.ai" chart repository ...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 9 charts Downloading common from repo oci://registry-1.docker.io/bitnamicharts Pulled: registry-1.docker.io/bitnamicharts/common:2.31.4 Digest: sha256:98cc992cb269f0b3f8f87c4bf933c1c0991b2d639f79e0bf62d1f35926efeb3f Downloading keycloak from repo https://charts.bitnami.com/bitnami Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/keycloak: failed to resolve registry-1.docker.io/bitnamicharts/keycloak:24.9.0: GET "https://registry-1.docker.io/v2/bitnamicharts/keycloak/manifests/24.9.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/keycloak: failed to resolve registry-1.docker.io/bitnamicharts/keycloak:24.9.0: GET "https://registry-1.docker.io/v2/bitnamicharts/keycloak/manifests/24.9.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 107 | [Kong/charts](https://github.com/Kong/charts) | `E:\helm_clones_artifacthub\Kong__charts\charts\ingress` | Error: no repository definition for https://charts.konghq.com, https://charts.konghq.com. Please add the missing repos via 'helm repo add' |
| 108 | [kubewarden/helm-charts](https://github.com/kubewarden/helm-charts) | `E:\helm_clones_artifacthub\kubewarden__helm-charts\charts\kubewarden-controller` | Error: no repository definition for https://kyverno.github.io/policy-reporter. Please add the missing repos via 'helm repo add' |
| 109 | [port-labs/helm-charts](https://github.com/port-labs/helm-charts) | `E:\helm_clones_artifacthub\port-labs__helm-charts\charts\port-ocean` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 110 | [relution-io/relution-kubernetes](https://github.com/relution-io/relution-kubernetes) | `E:\helm_clones_artifacthub\relution-io__relution-kubernetes\charts\relution` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 111 | [kubewarden/sbomscanner](https://github.com/kubewarden/sbomscanner) | `E:\helm_clones_artifacthub\kubewarden__sbomscanner\charts\sbomscanner` | Error: no repository definition for https://nats-io.github.io/k8s/helm/charts/. Please add the missing repos via 'helm repo add' |
| 112 | [shini4i/charts](https://github.com/shini4i/charts) | `E:\helm_clones_artifacthub\shini4i__charts\charts\app` | Error: no repository definition for https://shini4i.github.io/charts/. Please add the missing repos via 'helm repo add' |

## Full Error Details

### `E:\helm_clones_artifacthub\prometheus-community__helm-charts\charts\kube-prometheus-stack` — prometheus-community/helm-charts

```
Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://grafana-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\argoproj__argo-helm\charts\argo-cd` — argoproj/argo-helm

```
Error: no repository definition for https://dandydeveloper.github.io/charts/. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\bitnami__charts\bitnami\grafana-loki` — bitnami/charts

```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 6 charts
Downloading grafana-alloy from repo oci://registry-1.docker.io/bitnamicharts
Pulled: registry-1.docker.io/bitnamicharts/grafana-alloy:1.0.7
Digest: sha256:ae23a727fe845affa83689d30d3311f46c80fbcb8a217b435bb61d67c11e98eb
Downloading memcached from repo oci://registry-1.docker.io/bitnamicharts
Pulled: registry-1.docker.io/bitnamicharts/memcached:7.9.7
Digest: sha256:ffb71ea353076658513efe8b46e3f0e6a19653b8b1ca8532f1ba7bb3a490a3e3
Downloading memcached from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/memcached: failed to resolve registry-1.docker.io/bitnamicharts/memcached:7.9.7: GET "https://registry-1.docker.io/v2/bitnamicharts/memcached/manifests/7.9.7": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/memcached: failed to resolve registry-1.docker.io/bitnamicharts/memcached:7.9.7: GET "https://registry-1.docker.io/v2/bitnamicharts/memcached/manifests/7.9.7": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `E:\helm_clones_artifacthub\kubernetes__dashboard\charts\kubernetes-dashboard` — kubernetes/dashboard

```
Error: no repository definition for https://kubernetes.github.io/ingress-nginx, https://charts.jetstack.io, https://kubernetes-sigs.github.io/metrics-server/, https://charts.konghq.com. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\grafana__helm-charts\charts\enterprise-logs` — grafana/helm-charts

```
Error: no repository definition for https://grafana.github.io/helm-charts, https://charts.min.io/. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\goharbor__harbor\src\pkg\chart\testdata\harbor-schema1` — goharbor/harbor

```
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "https://kubernetes-charts.storage.googleapis.com/" chart repository:
	failed to fetch https://kubernetes-charts.storage.googleapis.com/index.yaml : 403 Forbidden
...Unable to get an update from the "https://kubernetes-charts.storage.googleapis.com/" chart repository:
	failed to fetch https://kubernetes-charts.storage.googleapis.com/index.yaml : 403 Forbidden
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Error: no cached repository for helm-manager-1067d9c6027b8c3f27b49e40521d64be96ea412858d8e45064fa44afd3966ddc found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-1067d9c6027b8c3f27b49e40521d64be96ea412858d8e45064fa44afd3966ddc-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### `E:\helm_clones_artifacthub\apache__airflow\chart` — apache/airflow

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\elastic__helm-charts\metricbeat` — elastic/helm-charts

```
Error: no repository definition for https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\nextcloud__helm\charts\nextcloud` — nextcloud/helm

```
Error: no repository definition for https://collaboraonline.github.io/online. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\metallb__metallb\charts\metallb` — metallb/metallb

```
Error: no repository definition for https://metallb.github.io/frr-k8s. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\SonarSource__helm-chart-sonarqube\charts\sonarqube` — SonarSource/helm-chart-sonarqube

```
Error: no repository definition for https://kubernetes.github.io/ingress-nginx. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\kyverno__kyverno\charts\kyverno` — kyverno/kyverno

```
Error: no repository definition for https://kyverno.github.io/api, https://openreports.github.io/reports-api, https://kyverno.github.io/reports-server/. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\cloudnative-pg__charts\charts\cloudnative-pg` — cloudnative-pg/charts

```
Error: no repository definition for https://cloudnative-pg.github.io/grafana-dashboards. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\DataDog__helm-charts\charts\datadog` — DataDog/helm-charts

```
Error: no repository definition for https://helm.datadoghq.com, https://prometheus-community.github.io/helm-charts, https://helm.datadoghq.com, https://helm.datadoghq.com. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\oauth2-proxy__manifests\helm\oauth2-proxy` — oauth2-proxy/manifests

```
Error: no repository definition for https://dandydeveloper.github.io/charts. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\artifacthub__hub\charts\artifact-hub` — artifacthub/hub

```
Getting updates for unmanaged Helm repositories...
...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 1 charts
Downloading postgresql from repo https://charts.bitnami.com/bitnami
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/postgresql: failed to resolve registry-1.docker.io/bitnamicharts/postgresql:18.0.15: GET "https://registry-1.docker.io/v2/bitnamicharts/postgresql/manifests/18.0.15": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/postgresql: failed to resolve registry-1.docker.io/bitnamicharts/postgresql:18.0.15: GET "https://registry-1.docker.io/v2/bitnamicharts/postgresql/manifests/18.0.15": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `E:\helm_clones_artifacthub\norwoodj__helm-docs\example-charts\custom-template` — norwoodj/helm-docs

```
Error: no repository definition for @stable. Please add them via 'helm repo add'
```

### `E:\helm_clones_artifacthub\community-charts__helm-charts\charts\kserve` — community-charts/helm-charts

```
Error: no repository definition for https://charts.jetstack.io. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\jaegertracing__helm-charts\charts\jaeger` — jaegertracing/helm-charts

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\8gears__n8n-helm-chart\charts\n8n` — 8gears/n8n-helm-chart

```
Error: no repository definition for https://valkey.io/valkey-helm/. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\airflow-helm__charts\charts\airflow` — airflow-helm/charts

```
Error: no repository definition for https://charts.helm.sh/stable, https://charts.helm.sh/stable. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\renovatebot__renovate\lib\modules\manager\helmv3\__fixtures__` — renovatebot/renovate

```
Error: directory E:\helm_clones_artifacthub\renovatebot__renovate\lib\modules\manager\helmv3\dependency_chart\dask not found
```

### `E:\helm_clones_artifacthub\DandyDeveloper__charts\charts\grafana-agent` — DandyDeveloper/charts

```
Error: no repository definition for https://helm.releases.hashicorp.com. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\rook__rook\deploy\charts\rook-ceph` — rook/rook

```
Error: error unpacking subchart library in rook-ceph: Chart.yaml file is missing
```

### `E:\helm_clones_artifacthub\codecentric__helm-charts\charts\keycloak` — codecentric/helm-charts

```
Error: no repository definition for https://raw.githubusercontent.com/bitnami/charts/archive-full-index/bitnami. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\VictoriaMetrics__helm-charts\charts\victoria-logs-agent` — VictoriaMetrics/helm-charts

```
Error: no repository definition for https://victoriametrics.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\backstage__charts\charts\backstage` — backstage/charts

```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 2 charts
Downloading common from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/common: failed to resolve registry-1.docker.io/bitnamicharts/common:2.10.0: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.10.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/common: failed to resolve registry-1.docker.io/bitnamicharts/common:2.10.0: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.10.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `E:\helm_clones_artifacthub\open-webui__helm-charts\charts\open-webui` — open-webui/helm-charts

```
Error: no repository definition for https://otwld.github.io/ollama-helm/, https://helm.openwebui.com, https://apache.jfrog.io/artifactory/tika, https://helm.openwebui.com. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\grafana__k8s-monitoring-helm\charts\k8s-monitoring` — grafana/k8s-monitoring-helm

```
Error: no repository definition for https://grafana.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\CloudPirates-io__helm-charts\charts\clusterpirate` — CloudPirates-io/helm-charts

```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 2 charts
Downloading common from repo oci://registry-1.docker.io/cloudpirates
Save error occurred:  could not download oci://registry-1.docker.io/cloudpirates/common: failed to resolve registry-1.docker.io/cloudpirates/common:2.2.0: GET "https://registry-1.docker.io/v2/cloudpirates/common/manifests/2.2.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/cloudpirates/common: failed to resolve registry-1.docker.io/cloudpirates/common:2.2.0: GET "https://registry-1.docker.io/v2/cloudpirates/common/manifests/2.2.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `E:\helm_clones_artifacthub\atlassian__data-center-helm-charts\src\main\charts\bamboo` — atlassian/data-center-helm-charts

```
Error: no repository definition for https://atlassian.github.io/data-center-helm-charts. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\zitadel__zitadel-charts\charts\zitadel` — zitadel/zitadel-charts

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\jp-gouin__helm-openldap` — jp-gouin/helm-openldap

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\bootc__netbox-chart\charts\netbox` — bootc/netbox-chart

```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 3 charts
Downloading common from repo oci://registry-1.docker.io/bitnamicharts
Pulled: registry-1.docker.io/bitnamicharts/common:2.38.0
Digest: sha256:fd148d63a289816d6af9bb3bd6fbe6317f24918a663ed824ad53f8de004c4234
Downloading postgresql from repo oci://registry-1.docker.io/bitnamicharts
Pulled: registry-1.docker.io/bitnamicharts/postgresql:18.5.16
Digest: sha256:2e0eb26439c7f857677d33747b25636059f4189110e462f4a2c0ff73ae5314a7
Downloading valkey from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/valkey: failed to resolve registry-1.docker.io/bitnamicharts/valkey:5.4.10: GET "https://registry-1.docker.io/v2/bitnamicharts/valkey/manifests/5.4.10": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/valkey: failed to resolve registry-1.docker.io/bitnamicharts/valkey:5.4.10: GET "https://registry-1.docker.io/v2/bitnamicharts/valkey/manifests/5.4.10": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `E:\helm_clones_artifacthub\OpsMx__spinnaker-helm\charts\spinnaker` — OpsMx/spinnaker-helm

```
load.go:138: Warning: Dependency locking is handled in Chart.lock since apiVersion "v2". We recommend migrating to Chart.lock.
load.go:120: Warning: Dependencies are handled in Chart.yaml since apiVersion "v2". We recommend migrating dependencies to Chart.yaml.
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\milvus-io__milvus-helm\charts\milvus` — milvus-io/milvus-helm

```
Error: no repository definition for https://raw.githubusercontent.com/bitnami/charts/pre-2022/bitnami, https://milvus-io.github.io/milvus-helm, https://pulsar.apache.org/charts, https://raw.githubusercontent.com/bitnami/charts/pre-2022/bitnami, https://raw.githubusercontent.com/bitnami/charts/pre-2022/bitnami. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\camunda__camunda-platform-helm\charts\camunda-platform-8.3` — camunda/camunda-platform-helm

```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 4 charts
Dependency identity did not declare a repository. Assuming it exists in the charts directory
Downloading elasticsearch from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/elasticsearch: failed to resolve registry-1.docker.io/bitnamicharts/elasticsearch:19.21.2: GET "https://registry-1.docker.io/v2/bitnamicharts/elasticsearch/manifests/19.21.2": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/elasticsearch: failed to resolve registry-1.docker.io/bitnamicharts/elasticsearch:19.21.2: GET "https://registry-1.docker.io/v2/bitnamicharts/elasticsearch/manifests/19.21.2": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `E:\helm_clones_artifacthub\WeblateOrg__helm\charts\weblate` — WeblateOrg/helm

```
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\cowboysysop__charts\archives\katib` — cowboysysop/charts

```
Error: no repository definition for https://charts.bitnami.com/bitnami/, https://charts.bitnami.com/bitnami/. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\deliveryhero__helm-charts\stable\cachet` — deliveryhero/helm-charts

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\groundhog2k__helm-charts\charts\ghost` — groundhog2k/helm-charts

```
Error: no repository definition for https://groundhog2k.github.io/helm-charts, https://groundhog2k.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\jouve__charts\charts\extra` — jouve/charts

```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 1 charts
Downloading common from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/common: failed to resolve registry-1.docker.io/bitnamicharts/common:2.38.0: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.38.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/common: failed to resolve registry-1.docker.io/bitnamicharts/common:2.38.0: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.38.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `E:\helm_clones_artifacthub\jfrog__charts\stable\artifactory` — jfrog/charts

```
Error: no repository definition for https://charts.jfrog.io/. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\AdWerx__charts\awx` — AdWerx/charts

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\SigNoz__charts\charts\signoz` — SigNoz/charts

```
Error: no repository definition for https://charts.signoz.io, https://charts.signoz.io, https://charts.redpanda.com, https://charts.signoz.io. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\gabe565__charts\charts\adguard-home` — gabe565/charts

```
Error: no repository definition for https://bjw-s.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\k8sgpt-ai__k8sgpt-operator\chart\operator` — k8sgpt-ai/k8sgpt-operator

```
Error: no repository definition for https://charts.k8sgpt.ai/. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\grafana-community__helm-charts\charts\loki` — grafana-community/helm-charts

```
Error: no repository definition for https://charts.min.io/, https://grafana.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\kubernetes__charts\incubator\distribution` — kubernetes/charts

```
manager.go:122: warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash...
Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/, https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\langfuse__langfuse-k8s\charts\langfuse` — langfuse/langfuse-k8s

```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 5 charts
Downloading postgresql from repo oci://registry-1.docker.io/bitnamicharts
Pulled: registry-1.docker.io/bitnamicharts/postgresql:16.4.9
Digest: sha256:ab4a0db005ca6a34c16c8ba72e76cb0241c4b856ec0699e956ef6c5f4f661444
Downloading clickhouse from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/clickhouse: failed to resolve registry-1.docker.io/bitnamicharts/clickhouse:8.0.5: GET "https://registry-1.docker.io/v2/bitnamicharts/clickhouse/manifests/8.0.5": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/clickhouse: failed to resolve registry-1.docker.io/bitnamicharts/clickhouse:8.0.5: GET "https://registry-1.docker.io/v2/bitnamicharts/clickhouse/manifests/8.0.5": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `E:\helm_clones_artifacthub\wiremind__wiremind-helm-charts\charts\druid` — wiremind/wiremind-helm-charts

```
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://machine424.github.io/kube-hpa-scale-to-zero, https://wiremind.github.io/wiremind-helm-charts, https://wiremind.github.io/wiremind-helm-charts. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\getredash__contrib-helm-chart\charts\redash` — getredash/contrib-helm-chart

```
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\ngrok__ngrok-operator\helm\ngrok-operator` — ngrok/ngrok-operator

```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 2 charts
Downloading common from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/common: failed to resolve registry-1.docker.io/bitnamicharts/common:2.36.0: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.36.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/common: failed to resolve registry-1.docker.io/bitnamicharts/common:2.36.0: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.36.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `E:\helm_clones_artifacthub\evryfs__helm-charts\charts\dependency-track` — evryfs/helm-charts

```
Error: no repository definition for https://raw.githubusercontent.com/bitnami/charts/eb5f9a9513d987b519f0ecd732e7031241c50328/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\k8gb-io__k8gb\chart\k8gb` — k8gb-io/k8gb

```
Error: no repository definition for https://coredns.github.io/helm, https://kubernetes-sigs.github.io/external-dns. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\localstack__helm-charts\charts\localstack` — localstack/helm-charts

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\percona__percona-helm-charts\charts\everest` — percona/percona-helm-charts

```
Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://victoriametrics.github.io/helm-charts, https://percona.github.io/percona-helm-charts, https://percona.github.io/operator-lifecycle-manager. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\Flagsmith__flagsmith-charts\charts\flagsmith` — Flagsmith/flagsmith-charts

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\glasskube__operator\charts\glasskube-operator` — glasskube/operator

```
Error: no repository definition for https://charts.jetstack.io, https://prometheus-community.github.io/helm-charts, https://mariadb-operator.github.io/mariadb-operator, https://cloudnative-pg.io/charts/, https://charts.min.io/. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\oauth2-proxy__oauth2-proxy\contrib\local-environment\kubernetes` — oauth2-proxy/oauth2-proxy

```
Error: no repository definition for https://charts.dexidp.io, https://oauth2-proxy.github.io/manifests, https://conservis.github.io/helm-charts, https://conservis.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\aws__karpenter\charts\karpenter` — aws/karpenter

```
Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies
```

### `E:\helm_clones_artifacthub\mvisonneau__helm-charts\charts\gitlab-ci-pipelines-exporter` — mvisonneau/helm-charts

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\zilliztech__milvus-helm\charts\milvus` — zilliztech/milvus-helm

```
Error: no repository definition for https://raw.githubusercontent.com/bitnami/charts/archive-full-index/bitnami, https://zilliztech.github.io/milvus-helm, https://zilliztech.github.io/milvus-helm, https://pulsar.apache.org/charts, https://raw.githubusercontent.com/bitnami/charts/archive-full-index/bitnami, https://zilliztech.github.io/milvus-helm. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\sentry-kubernetes__charts\charts\sentry` — sentry-kubernetes/charts

```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 5 charts
Downloading memcached from repo oci://registry-1.docker.io/cloudpirates
Pulled: registry-1.docker.io/cloudpirates/memcached:0.9.3
Digest: sha256:b5472e0c3f4ce4069b86dd16ba754ea78ef977ae5f9780ca1682557bdb02e94d
Downloading redis from repo oci://registry-1.docker.io/bitnamicharts
Pulled: registry-1.docker.io/bitnamicharts/redis:17.11.3
Digest: sha256:6472c71c6befe73673f584706a7924f6b57974a510149ed750b6f165ed7d30b2
Downloading kafka from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/kafka: failed to resolve registry-1.docker.io/bitnamicharts/kafka:29.3.14: GET "https://registry-1.docker.io/v2/bitnamicharts/kafka/manifests/29.3.14": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/kafka: failed to resolve registry-1.docker.io/bitnamicharts/kafka:29.3.14: GET "https://registry-1.docker.io/v2/bitnamicharts/kafka/manifests/29.3.14": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `E:\helm_clones_artifacthub\apache__incubator-superset\helm\superset` — apache/incubator-superset

```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 2 charts
Downloading postgresql from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/postgresql: failed to resolve registry-1.docker.io/bitnamicharts/postgresql:16.7.27: GET "https://registry-1.docker.io/v2/bitnamicharts/postgresql/manifests/16.7.27": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/postgresql: failed to resolve registry-1.docker.io/bitnamicharts/postgresql:16.7.27: GET "https://registry-1.docker.io/v2/bitnamicharts/postgresql/manifests/16.7.27": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `E:\helm_clones_artifacthub\hivemq__helm-charts\charts\hivemq-operator` — hivemq/helm-charts

```
Error: no repository definition for https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\Azure__secrets-store-csi-driver-provider-azure\charts\csi-secrets-store-provider-azure` — Azure/secrets-store-csi-driver-provider-azure

```
Error: no repository definition for https://kubernetes-sigs.github.io/secrets-store-csi-driver/charts. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\asdf2014__druid-helm\charts\druid` — asdf2014/druid-helm

```
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.helm.sh/stable, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\one-acre-fund__oaf-public-charts\archive\geonode` — one-acre-fund/oaf-public-charts

```
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\kubernetes__helm\internal\chart\v3\lint\rules\testdata\anotherbadchartfile` — kubernetes/helm

```
Error: validation: chart.metadata.version "7.2445e+06" is invalid
```

### `E:\helm_clones_artifacthub\anza-labs__charts\deprecated\hosted-control-plane` — anza-labs/charts

```
Error: no repository definition for https://charts.jetstack.io. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\pree__helm-charts\charts\home-assistant` — pree/helm-charts

```
Getting updates for unmanaged Helm repositories...
...Successfully got an update from the "https://bjw-s-labs.github.io/helm-charts/" chart repository
...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository
...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository
...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 4 charts
Downloading common from repo https://bjw-s-labs.github.io/helm-charts/
Downloading postgresql from repo https://charts.bitnami.com/bitnami
Pulled: registry-1.docker.io/bitnamicharts/postgresql:16.7.27
Digest: sha256:4ed63727a46ff6e6e1cc48b87c75c4b2abe9e5901d38079ba22768895c00de7a
Downloading mariadb from repo https://charts.bitnami.com/bitnami
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/mariadb: failed to resolve registry-1.docker.io/bitnamicharts/mariadb:22.0.0: GET "https://registry-1.docker.io/v2/bitnamicharts/mariadb/manifests/22.0.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/mariadb: failed to resolve registry-1.docker.io/bitnamicharts/mariadb:22.0.0: GET "https://registry-1.docker.io/v2/bitnamicharts/mariadb/manifests/22.0.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `E:\helm_clones_artifacthub\artur9010__charts\charts\cloudflared` — artur9010/charts

```
Error: no repository definition for https://kubernetes.github.io/ingress-nginx. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\helm__charts\incubator\distribution` — helm/charts

```
manager.go:122: warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash...
Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/, https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\helm__helm\internal\chart\v3\lint\rules\testdata\anotherbadchartfile` — helm/helm

```
Error: validation: chart.metadata.version "7.2445e+06" is invalid
```

### `E:\helm_clones_artifacthub\Gradiant__charts\charts\hbase` — Gradiant/charts

```
Getting updates for unmanaged Helm repositories...
...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Error: can't get a valid version for 1 subchart(s): "zookeeper" (repository "https://charts.bitnami.com/bitnami", version "~6.3.0"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml
```

### `E:\helm_clones_artifacthub\banzaicloud__banzai-charts\argo` — banzaicloud/banzai-charts

```
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "https://kubernetes-charts.storage.googleapis.com/" chart repository:
	failed to fetch https://kubernetes-charts.storage.googleapis.com/index.yaml : 403 Forbidden
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Error: no cached repository for helm-manager-1067d9c6027b8c3f27b49e40521d64be96ea412858d8e45064fa44afd3966ddc found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-1067d9c6027b8c3f27b49e40521d64be96ea412858d8e45064fa44afd3966ddc-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### `E:\helm_clones_artifacthub\beeinventor__charts\beeinventor\keycloak` — beeinventor/charts

```
Error: no repository definition for https://charts.bitnami.com/bitnami, https://raw.githubusercontent.com/bitnami/charts/archive-full-index/bitnami. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\fmjstudios__helm\charts\activepieces` — fmjstudios/helm

```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 2 charts
Downloading postgresql from repo oci://registry-1.docker.io/bitnamicharts
Pulled: registry-1.docker.io/bitnamicharts/postgresql:15.5.38
Digest: sha256:fd220eb22fa79a7c3568928d5eebba7b57ea00a9d10a6134e1d42bc51f0e3346
Downloading redis from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:19.5.5: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/19.5.5": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:19.5.5: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/19.5.5": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `E:\helm_clones_artifacthub\monicahq__helm\charts\monica` — monicahq/helm

```
Error: no repository definition for https://meilisearch.github.io/meilisearch-kubernetes. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\ortelius__ortelius-charts\chart\ortelius` — ortelius/ortelius-charts

```
Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies
```

### `E:\helm_clones_artifacthub\snowplow-devops__helm-charts\charts\avalanche` — snowplow-devops/helm-charts

```
Error: no repository definition for https://snowplow-devops.github.io/helm-charts, https://snowplow-devops.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\enix__helm-charts\charts\cnpg-monitoring` — enix/helm-charts

```
Error: no repository definition for https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\cloudhippie__charts\stable\yopass` — cloudhippie/charts

```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 2 charts
Downloading redis from repo oci://registry-1.docker.io/cloudpirates
Save error occurred:  could not download oci://registry-1.docker.io/cloudpirates/redis: failed to resolve registry-1.docker.io/cloudpirates/redis:0.26.8: GET "https://registry-1.docker.io/v2/cloudpirates/redis/manifests/0.26.8": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/cloudpirates/redis: failed to resolve registry-1.docker.io/cloudpirates/redis:0.26.8: GET "https://registry-1.docker.io/v2/cloudpirates/redis/manifests/0.26.8": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `E:\helm_clones_artifacthub\andrenarchy__helm-charts\charts\home-assistant` — andrenarchy/helm-charts

```
Error: no repository definition for https://library-charts.k8s-at-home.com, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://andrenarchy.github.io/helm-charts/. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\netrisai__charts\charts\netris-controller` — netrisai/charts

```
Getting updates for unmanaged Helm repositories...
...Successfully got an update from the "https://raw.githubusercontent.com/rawfilescloud/charts/refs/heads/main" chart repository
...Successfully got an update from the "https://haproxytech.github.io/helm-charts" chart repository
...Successfully got an update from the "https://charts.ntppool.org" chart repository
...Successfully got an update from the "https://raw.githubusercontent.com/bitnami/charts/eb5f9a9513d987b519f0ecd732e7031241c50328/bitnami" chart repository
...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository
...Successfully got an update from the "https://raw.githubusercontent.com/bitnami/charts/eb5f9a9513d987b519f0ecd732e7031241c50328/bitnami" chart repository
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Error: can't get a valid version for 1 subchart(s): "smtp" (repository "https://charts.ntppool.org", version "1.0.1"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml
```

### `E:\helm_clones_artifacthub\epinio__helm-charts\chart\epinio` — epinio/helm-charts

```
Error: no repository definition for https://charts.dexidp.io, https://charts.min.io/, https://s3gw-tech.github.io/s3gw-charts. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\keptn__lifecycle-toolkit\keptn-cert-manager\chart` — keptn/lifecycle-toolkit

```
Error: no repository definition for https://charts.lifecycle.keptn.sh. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\StrangeBeeCorp__helm-charts\thehive-charts\thehive` — StrangeBeeCorp/helm-charts

```
Getting updates for unmanaged Helm repositories...
...Successfully got an update from the "https://charts.min.io" chart repository
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 3 charts
Downloading cassandra from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/cassandra: failed to resolve registry-1.docker.io/bitnamicharts/cassandra:12.3.11: GET "https://registry-1.docker.io/v2/bitnamicharts/cassandra/manifests/12.3.11": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/cassandra: failed to resolve registry-1.docker.io/bitnamicharts/cassandra:12.3.11: GET "https://registry-1.docker.io/v2/bitnamicharts/cassandra/manifests/12.3.11": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `E:\helm_clones_artifacthub\redhat-developer__rhdh-chart\charts\backstage` — redhat-developer/rhdh-chart

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\adfinis__helm-charts\charts\argo-apps` — adfinis/helm-charts

```
Error: no repository definition for https://charts.adfinis.com. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\apache__doris-operator\helm-charts\doris` — apache/doris-operator

```
Error: cannot load Chart.yaml: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal object into Go struct field Metadata.version of type string
```

### `E:\helm_clones_artifacthub\intel__helm-charts\charts\evi-hbase-0.8.3\charts\hbase` — intel/helm-charts

```
Error: directory E:\helm_clones_artifacthub\intel__helm-charts\charts\evi-hbase-0.8.3\charts\hadoop not found
```

### `E:\helm_clones_artifacthub\k8s-at-home__charts\charts\stable\firefly-iii` — k8s-at-home/charts

```
Getting updates for unmanaged Helm repositories...
...Successfully got an update from the "https://library-charts.k8s-at-home.com" chart repository
...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository
...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository
...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Error: can't get a valid version for 1 subchart(s): "mariadb" (repository "https://charts.bitnami.com/bitnami", version "11.0.2"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml
```

### `E:\helm_clones_artifacthub\netbox-community__netbox-chart\charts\netbox` — netbox-community/netbox-chart

```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 3 charts
Downloading common from repo oci://registry-1.docker.io/bitnamicharts
Pulled: registry-1.docker.io/bitnamicharts/common:2.38.0
Digest: sha256:fd148d63a289816d6af9bb3bd6fbe6317f24918a663ed824ad53f8de004c4234
Downloading postgresql from repo oci://registry-1.docker.io/bitnamicharts
Pulled: registry-1.docker.io/bitnamicharts/postgresql:18.5.16
Digest: sha256:2e0eb26439c7f857677d33747b25636059f4189110e462f4a2c0ff73ae5314a7
Downloading valkey from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/valkey: failed to resolve registry-1.docker.io/bitnamicharts/valkey:5.4.10: GET "https://registry-1.docker.io/v2/bitnamicharts/valkey/manifests/5.4.10": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/valkey: failed to resolve registry-1.docker.io/bitnamicharts/valkey:5.4.10: GET "https://registry-1.docker.io/v2/bitnamicharts/valkey/manifests/5.4.10": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `E:\helm_clones_artifacthub\obeone__charts\charts\cyberchef` — obeone/charts

```
Error: no repository definition for https://library-charts.k8s-at-home.com. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\icoretech__helm\charts\tolgee` — icoretech/helm

```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 1 charts
Downloading postgres from repo oci://registry-1.docker.io/cloudpirates
Save error occurred:  could not download oci://registry-1.docker.io/cloudpirates/postgres: failed to resolve registry-1.docker.io/cloudpirates/postgres:0.17.2: GET "https://registry-1.docker.io/v2/cloudpirates/postgres/manifests/0.17.2": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/cloudpirates/postgres: failed to resolve registry-1.docker.io/cloudpirates/postgres:0.17.2: GET "https://registry-1.docker.io/v2/cloudpirates/postgres/manifests/0.17.2": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `E:\helm_clones_artifacthub\liranme__redisinsight-secure\helm\redisinsight-secure` — liranme/redisinsight-secure

```
Error: no repository definition for https://oauth2-proxy.github.io/manifests. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\kir4h__charts\charts\process-exporter` — kir4h/charts

```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 1 charts
Downloading common from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/common: failed to resolve registry-1.docker.io/bitnamicharts/common:2.30.0: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.30.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/common: failed to resolve registry-1.docker.io/bitnamicharts/common:2.30.0: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.30.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `E:\helm_clones_artifacthub\ectobit__charts\rspamd` — ectobit/charts

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\rstudio__helm\charts\rstudio-connect` — rstudio/helm

```
Error: no repository definition for https://helm.rstudio.com. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\hpe-storage__truenas-csp\helm\charts\truenas-csp` — hpe-storage/truenas-csp

```
Error: no repository definition for https://hpe-storage.github.io/co-deployments. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\christianhuth__helm-charts\charts\baserow` — christianhuth/helm-charts

```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 2 charts
Downloading postgresql from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/postgresql: failed to resolve registry-1.docker.io/bitnamicharts/postgresql:18.5.16: GET "https://registry-1.docker.io/v2/bitnamicharts/postgresql/manifests/18.5.16": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/postgresql: failed to resolve registry-1.docker.io/bitnamicharts/postgresql:18.5.16: GET "https://registry-1.docker.io/v2/bitnamicharts/postgresql/manifests/18.5.16": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `E:\helm_clones_artifacthub\sitewhere__sitewhere-k8s\charts\sitewhere` — sitewhere/sitewhere-k8s

```
Error: cannot load Chart.yaml: error converting YAML to JSON: yaml: line 2: mapping values are not allowed in this context
```

### `E:\helm_clones_artifacthub\AvistoTelecom__charts\charts\cloudbeaver` — AvistoTelecom/charts

```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 1 charts
Downloading common from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/common: failed to resolve registry-1.docker.io/bitnamicharts/common:2.36.0: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.36.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/common: failed to resolve registry-1.docker.io/bitnamicharts/common:2.36.0: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.36.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `E:\helm_clones_artifacthub\epam__ai-dial-helm\charts\dial` — epam/ai-dial-helm

```
Getting updates for unmanaged Helm repositories...
...Successfully got an update from the "https://charts.dialx.ai" chart repository
...Successfully got an update from the "https://charts.dialx.ai" chart repository
...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 9 charts
Downloading common from repo oci://registry-1.docker.io/bitnamicharts
Pulled: registry-1.docker.io/bitnamicharts/common:2.31.4
Digest: sha256:98cc992cb269f0b3f8f87c4bf933c1c0991b2d639f79e0bf62d1f35926efeb3f
Downloading keycloak from repo https://charts.bitnami.com/bitnami
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/keycloak: failed to resolve registry-1.docker.io/bitnamicharts/keycloak:24.9.0: GET "https://registry-1.docker.io/v2/bitnamicharts/keycloak/manifests/24.9.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/keycloak: failed to resolve registry-1.docker.io/bitnamicharts/keycloak:24.9.0: GET "https://registry-1.docker.io/v2/bitnamicharts/keycloak/manifests/24.9.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `E:\helm_clones_artifacthub\Kong__charts\charts\ingress` — Kong/charts

```
Error: no repository definition for https://charts.konghq.com, https://charts.konghq.com. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\kubewarden__helm-charts\charts\kubewarden-controller` — kubewarden/helm-charts

```
Error: no repository definition for https://kyverno.github.io/policy-reporter. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\port-labs__helm-charts\charts\port-ocean` — port-labs/helm-charts

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\relution-io__relution-kubernetes\charts\relution` — relution-io/relution-kubernetes

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\kubewarden__sbomscanner\charts\sbomscanner` — kubewarden/sbomscanner

```
Error: no repository definition for https://nats-io.github.io/k8s/helm/charts/. Please add the missing repos via 'helm repo add'
```

### `E:\helm_clones_artifacthub\shini4i__charts\charts\app` — shini4i/charts

```
Error: no repository definition for https://shini4i.github.io/charts/. Please add the missing repos via 'helm repo add'
```

