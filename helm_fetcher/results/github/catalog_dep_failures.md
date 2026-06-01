# Helm Dependency Build Failures

Total dependency failures: **133**

| # | Repository | Chart Path | Error |
|---|------------|------------|-------|
| 1 | [rancher/rancher](https://github.com/rancher/rancher) | `D:\helm_clones_github\rancher__rancher\chart` | Error: cannot load Chart.yaml: error converting YAML to JSON: yaml: line 4: found character that cannot start any token |
| 2 | [refly-ai/refly](https://github.com/refly-ai/refly) | `D:\helm_clones_github\refly-ai__refly\deploy\helm\refly-stack` | Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 3 | [linode/apl-core](https://github.com/linode/apl-core) | `D:\helm_clones_github\linode__apl-core\chart\chart-index` | Getting updates for unmanaged Helm repositories... ...Successfully got an update from the "https://linode.github.io/cloud-firewall-controller" chart repository ...Successfully got an update from the "https://prometheus-msteams.github.io/prometheus-msteams/" chart repository ...Successfully got an update from the "https://kubernetes-sigs.github.io/metrics-server" chart repository ...Successfully got an update from the "https://cdfoundation.github.io/tekton-helm-chart/" chart repository ...Successfully got an update from the "https://knative.github.io/operator" chart repository ...Successfully got an update from the "https://kubernetes-sigs.github.io/external-dns" chart repository ...Successfully got an update from the "https://bitnami-labs.github.io/sealed-secrets/" chart repository ...Successfully got an update from the "https://kyverno.github.io/policy-reporter" chart repository ...Successfully got an update from the "https://aquasecurity.github.io/helm-charts/" chart repository ...Successfully got an update from the "https://oauth2-proxy.github.io/manifests" chart repository ...Successfully got an update from the "https://codecentric.github.io/helm-charts" chart repository ...Successfully got an update from the "https://kyverno.github.io/kyverno/" chart repository ...Successfully got an update from the "https://cloudnative-pg.github.io/charts" chart repository ...Successfully got an update from the "https://charts.jetstack.io" chart repository ...Successfully got an update from the "https://open-telemetry.github.io/opentelemetry-helm-charts" chart repository ...Successfully got an update from the "https://argoproj.github.io/argo-helm" chart repository ...Successfully got an update from the "https://dl.gitea.io/charts" chart repository ...Successfully got an update from the "https://helm.goharbor.io" chart repository ...Successfully got an update from the "https://grafana.github.io/helm-charts" chart repository ...Successfully got an update from the "https://prometheus-community.github.io/helm-charts" chart repository ...Successfully got an update from the "https://charts.external-secrets.io" chart repository ...Successfully got an update from the "https://istio-release.storage.googleapis.com/charts" chart repository ...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository Saving 31 charts Downloading argo-cd from repo https://argoproj.github.io/argo-helm Downloading argocd-image-updater from repo oci://ghcr.io/argoproj/argo-helm/argocd-image-updater Save error occurred:  could not download oci://ghcr.io/argoproj/argo-helm/argocd-image-updater/argocd-image-updater: failed to perform "FetchReference" on source: GET "https://ghcr.io/v2/argoproj/argo-helm/argocd-image-updater/argocd-image-updater/manifests/1.1.5": GET "https://ghcr.io/token?scope=repository%3Aargoproj%2Fargo-helm%2Fargocd-image-updater%2Fargocd-image-updater%3Apull&service=ghcr.io": response status code 403: denied: requested access to the resource is denied Error: could not download oci://ghcr.io/argoproj/argo-helm/argocd-image-updater/argocd-image-updater: failed to perform "FetchReference" on source: GET "https://ghcr.io/v2/argoproj/argo-helm/argocd-image-updater/argocd-image-updater/manifests/1.1.5": GET "https://ghcr.io/token?scope=repository%3Aargoproj%2Fargo-helm%2Fargocd-image-updater%2Fargocd-image-updater%3Apull&service=ghcr.io": response status code 403: denied: requested access to the resource is denied |
| 4 | [cozystack/cozystack](https://github.com/cozystack/cozystack) | `D:\helm_clones_github\cozystack__cozystack\packages\apps\bucket` | Error: error unpacking subchart cozy-lib in bucket: Chart.yaml file is missing |
| 5 | [grafana/helm-charts](https://github.com/grafana/helm-charts) | `D:\helm_clones_github\grafana__helm-charts\charts\enterprise-metrics` | Getting updates for unmanaged Helm repositories... ...Successfully got an update from the "https://helm.min.io/" chart repository ...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository Error: minio chart not found in repo https://helm.min.io/ |
| 6 | [norwoodj/helm-docs](https://github.com/norwoodj/helm-docs) | `D:\helm_clones_github\norwoodj__helm-docs\example-charts\custom-template` | Error: no repository definition for @stable. Please add them via 'helm repo add' |
| 7 | [securitybunker/databunker](https://github.com/securitybunker/databunker) | `D:\helm_clones_github\securitybunker__databunker\charts\databunker` | Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 8 | [WeBankFinTech/Prophecis](https://github.com/WeBankFinTech/Prophecis) | `D:\helm_clones_github\WeBankFinTech__Prophecis\install\Prophecis` | Error: cannot load values.yaml: error reading yaml document: invalid Yaml document separator: --END RSA PRIVATE KEY-----" |
| 9 | [rancher/charts](https://github.com/rancher/charts) | `D:\helm_clones_github\rancher__charts\charts\epinio\102.0.1+up1.6.2` | Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update' |
| 10 | [cloudnativeapp/charts](https://github.com/cloudnativeapp/charts) | `D:\helm_clones_github\cloudnativeapp__charts\curated\airflow` | level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..." Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies |
| 11 | [BigKAA/youtube](https://github.com/BigKAA/youtube) | `D:\helm_clones_github\BigKAA__youtube\tracing\for_admins\charts\jaeger\jaeger` | Error: no repository definition for https://charts.helm.sh/incubator, https://helm.elastic.co, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 12 | [radondb/radondb-mysql-kubernetes](https://github.com/radondb/radondb-mysql-kubernetes) | `D:\helm_clones_github\radondb__radondb-mysql-kubernetes\charts\mysql-operator` | Error: directory D:\helm_clones_github\radondb__radondb-mysql-kubernetes\charts\mysql-operator\charts\mysqlcluster not found |
| 13 | [IBM/charts](https://github.com/IBM/charts) | `D:\helm_clones_github\IBM__charts\community\artifactory-ha` | level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..." Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add' |
| 14 | [JahstreetOrg/spark-on-kubernetes-helm](https://github.com/JahstreetOrg/spark-on-kubernetes-helm) | `D:\helm_clones_github\JahstreetOrg__spark-on-kubernetes-helm\charts\cluster-base` | level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..." Error: no repository definition for https://kubernetes.github.io/ingress-nginx, https://charts.jetstack.io, https://charts.helm.sh/stable, https://charts.helm.sh/stable. Please add the missing repos via 'helm repo add' |
| 15 | [vexxhost/atmosphere](https://github.com/vexxhost/atmosphere) | `D:\helm_clones_github\vexxhost__atmosphere\charts\barbican` | level=INFO msg="Warning: Dependency locking is handled in Chart.lock since apiVersion \"v2\". We recommend migrating to Chart.lock." Error: no repository definition for https://tarballs.openstack.org/openstack-helm. Please add the missing repos via 'helm repo add' |
| 16 | [unixhot/devops-x](https://github.com/unixhot/devops-x) | `D:\helm_clones_github\unixhot__devops-x\helm\gitlab` | level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..." Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies |
| 17 | [rancher/rke2-charts](https://github.com/rancher/rke2-charts) | `D:\helm_clones_github\rancher__rke2-charts\packages\rke2-cilium-legacy\charts` | Error: directory D:\helm_clones_github\rancher__rke2-charts\packages\rke2-cilium-legacy\charts\charts\cilium not found |
| 18 | [stackrox/helm-charts](https://github.com/stackrox/helm-charts) | `D:\helm_clones_github\stackrox__helm-charts\3.0.41.0` | Error: validation: chart.metadata.version "3.0.41.0" is invalid |
| 19 | [project-sunbird/sunbird-devops](https://github.com/project-sunbird/sunbird-devops) | `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\monitoring\oauth2-proxy` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 20 | [rancher/partner-charts](https://github.com/rancher/partner-charts) | `D:\helm_clones_github\rancher__partner-charts\charts\amd\amd-gpu\0.10.0` | Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update' |
| 21 | [CARV-ICS-FORTH/frisbee](https://github.com/CARV-ICS-FORTH/frisbee) | `D:\helm_clones_github\CARV-ICS-FORTH__frisbee\charts\platform` | Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update' |
| 22 | [wikibook/kubepractice](https://github.com/wikibook/kubepractice) | `D:\helm_clones_github\wikibook__kubepractice\ch06\nginx-12.0.0` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 23 | [pluralsh/plural-artifacts](https://github.com/pluralsh/plural-artifacts) | `D:\helm_clones_github\pluralsh__plural-artifacts\airbyte\helm\airbyte` | Error: directory D:\helm_clones_github\airbyte\charts\airbyte not found |
| 24 | [tmforum-oda/oda-canvas](https://github.com/tmforum-oda/oda-canvas) | `D:\helm_clones_github\tmforum-oda__oda-canvas\charts\observability-stack` | Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://open-telemetry.github.io/opentelemetry-helm-charts, https://jaegertracing.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 25 | [sa-mw-dach/bobbycar](https://github.com/sa-mw-dach/bobbycar) | `D:\helm_clones_github\sa-mw-dach__bobbycar\helm\bobbycar-core-infra` | Error: no repository definition for https://drogue-iot.github.io/drogue-cloud-helm-charts/. Please add the missing repos via 'helm repo add' |
| 26 | [boozallen/aissemble](https://github.com/boozallen/aissemble) | `D:\helm_clones_github\boozallen__aissemble\foundation\foundation-archetype\src\main\resources\archetype-resources\__rootArtifactId__-deploy\src\main\resources\apps\common-infrastructure` | Error: validation: chart.metadata.version "${version}" is invalid |
| 27 | [tkestack/charts](https://github.com/tkestack/charts) | `D:\helm_clones_github\tkestack__charts\incubator\airflow` | level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..." Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/, https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add' |
| 28 | [sapcc/helm-charts](https://github.com/sapcc/helm-charts) | `D:\helm_clones_github\sapcc__helm-charts\common\inventory-updater` | Saving 1 charts Downloading owner-info from repo oci://keppel.eu-de-1.cloud.sap/ccloud-helm Save error occurred:  could not download oci://keppel.eu-de-1.cloud.sap/ccloud-helm/owner-info: failed to perform "FetchReference" on source: Get "https://keppel.eu-de-1.cloud.sap/v2/ccloud-helm/owner-info/manifests/0.2.0": dial tcp: lookup keppel.eu-de-1.cloud.sap: getaddrinfow: Este é geralmente um erro temporário durante a resolução de nomes de anfitrião e significa que o servidor local não recebeu uma resposta de um servidor autoritário. Error: could not download oci://keppel.eu-de-1.cloud.sap/ccloud-helm/owner-info: failed to perform "FetchReference" on source: Get "https://keppel.eu-de-1.cloud.sap/v2/ccloud-helm/owner-info/manifests/0.2.0": dial tcp: lookup keppel.eu-de-1.cloud.sap: getaddrinfow: Este é geralmente um erro temporário durante a resolução de nomes de anfitrião e significa que o servidor local não recebeu uma resposta de um servidor autoritário. |
| 29 | [platyplus/platyplus](https://github.com/platyplus/platyplus) | `D:\helm_clones_github\platyplus__platyplus\charts\hasura` | Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 30 | [opendevstack/ods-quickstarters](https://github.com/opendevstack/ods-quickstarters) | `D:\helm_clones_github\opendevstack__ods-quickstarters\be-rust-axum\rust-template\chart` | Error: cannot load Chart.yaml: error converting YAML to JSON: yaml: invalid map key: map[interface {}]interface {}{"project-name":interface {}(nil)} |
| 31 | [k8s-home-lab/helm-charts](https://github.com/k8s-home-lab/helm-charts) | `D:\helm_clones_github\k8s-home-lab__helm-charts\unmaintained\audiobookshelf` | Getting updates for unmanaged Helm repositories... ...Successfully got an update from the "https://library-charts.k8s-at-home.com" chart repository Error: can't get a valid version for 1 subchart(s): "common" (repository "https://library-charts.k8s-at-home.com", version "4.5.3"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml |
| 32 | [mojaloop/helm](https://github.com/mojaloop/helm) | `D:\helm_clones_github\mojaloop__helm\perf-test-harness` | Error: can't get a valid version for 1 subchart(s): "ml-testing-toolkit-cli" (repository "file://../ml-testing-toolkit-cli", version "15.9.0"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml |
| 33 | [camptocamp/charts](https://github.com/camptocamp/charts) | `D:\helm_clones_github\camptocamp__charts\common-build-code` | Getting updates for unmanaged Helm repositories... ...Unable to get an update from the "https://kubernetes-charts-incubator.storage.googleapis.com" chart repository: 	failed to fetch https://kubernetes-charts-incubator.storage.googleapis.com/index.yaml : 403 Forbidden Error: no cached repository for helm-manager-53271637451a5b2439ffd0af71673734b808e371a8a6aed9bf100a8f219a3006 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-53271637451a5b2439ffd0af71673734b808e371a8a6aed9bf100a8f219a3006-index.yaml: O sistema não conseguiu localizar o ficheiro especificado. |
| 34 | [devtron-labs/charts](https://github.com/devtron-labs/charts) | `D:\helm_clones_github\devtron-labs__charts\charts\cluster-essentials` | Getting updates for unmanaged Helm repositories... ...Successfully got an update from the "https://kedacore.github.io/charts" chart repository ...Successfully got an update from the "https://kubernetes.github.io/autoscaler" chart repository ...Successfully got an update from the "https://helm.devtron.ai" chart repository ...Successfully got an update from the "https://aws.github.io/eks-charts" chart repository ...Successfully got an update from the "https://kubernetes-sigs.github.io/metrics-server/" chart repository ...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository Error: can't get a valid version for 1 subchart(s): "kubernetes-event-exporter" (repository "https://charts.bitnami.com/bitnami", version "1.2.*"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml |
| 35 | [YAKEcloud/yake](https://github.com/YAKEcloud/yake) | `D:\helm_clones_github\YAKEcloud__yake\helmcharts\acl` | Error: dependency "controller" has an invalid version/constraint format: improper constraint: "" |
| 36 | [bflance/proxmox-talos](https://github.com/bflance/proxmox-talos) | `D:\helm_clones_github\bflance__proxmox-talos\charts\kube-prometheus-stack` | Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 37 | [k0rdent/catalog](https://github.com/k0rdent/catalog) | `D:\helm_clones_github\k0rdent__catalog\apps\alloy\charts\alloy-1.6.1` | Error: no repository definition for https://grafana.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 38 | [Loongson-Cloud-Community/dockerfiles](https://github.com/Loongson-Cloud-Community/dockerfiles) | `D:\helm_clones_github\Loongson-Cloud-Community__dockerfiles\kubesphere\ks-installer\v3.2.1\roles\ks-multicluster\files\kubefed\kubefed` | Getting updates for unmanaged Helm repositories... ...Unable to get an update from the "https://localhost/" chart repository: 	Get "https://localhost/index.yaml": dial tcp [::1]:443: connectex: Nenhuma ligação pôde ser feita porque o computador de destino as recusou ativamente. Error: no cached repository for helm-manager-f2b99ce05b94599549c70dbbe7a891b278e7c3cacad02334fa44682fca36c740 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-f2b99ce05b94599549c70dbbe7a891b278e7c3cacad02334fa44682fca36c740-index.yaml: O sistema não conseguiu localizar o ficheiro especificado. |
| 39 | [junghoon2/k8s-class](https://github.com/junghoon2/k8s-class) | `D:\helm_clones_github\junghoon2__k8s-class\argo-cd\argo-cd-5.14.1` | Error: no repository definition for https://dandydeveloper.github.io/charts/. Please add the missing repos via 'helm repo add' |
| 40 | [cloud-native-toolkit/toolkit-charts](https://github.com/cloud-native-toolkit/toolkit-charts) | `D:\helm_clones_github\cloud-native-toolkit__toolkit-charts\stable\cloud-setup` | Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update' |
| 41 | [suse-edge/charts](https://github.com/suse-edge/charts) | `D:\helm_clones_github\suse-edge__charts\charts\kubevirt\0.1.0` | Error: dependency "cdi" has an invalid version/constraint format: improper constraint: "" |
| 42 | [ai-solution-eng/frameworks](https://github.com/ai-solution-eng/frameworks) | `D:\helm_clones_github\ai-solution-eng__frameworks\appsmith\3.6.4` | Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 43 | [cnrancher/pandaria-catalog](https://github.com/cnrancher/pandaria-catalog) | `D:\helm_clones_github\cnrancher__pandaria-catalog\charts\rancher-hami\107.0.0+up2.5.2\charts\hami-webui` | Error: no repository definition for https://nvidia.github.io/dcgm-exporter/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 44 | [Sagar2366/tech-talks](https://github.com/Sagar2366/tech-talks) | `D:\helm_clones_github\Sagar2366__tech-talks\k8s_pune_oct22\prometheus-comunity-helm-chart\charts\kube-prometheus-stack` | Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://grafana.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 45 | [boozallen/sdp-helm-chart](https://github.com/boozallen/sdp-helm-chart) | `D:\helm_clones_github\boozallen__sdp-helm-chart` | Error: dependency "jenkins" has an invalid version/constraint format: improper constraint: "" |
| 46 | [randoli/helm-charts](https://github.com/randoli/helm-charts) | `D:\helm_clones_github\randoli__helm-charts\charts\cost-management` | Error: no repository definition for https://opencost.github.io/opencost-helm-chart. Please add the missing repos via 'helm repo add' |
| 47 | [unixfox/k8s](https://github.com/unixfox/k8s) | `D:\helm_clones_github\unixfox__k8s\charts\bibliogram` | Error: no repository definition for https://library-charts.k8s-at-home.com. Please add the missing repos via 'helm repo add' |
| 48 | [bcgov/OCWA](https://github.com/bcgov/OCWA) | `D:\helm_clones_github\bcgov__OCWA\helm\ocwa` | level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..." Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies |
| 49 | [henrywhitaker3/homelab](https://github.com/henrywhitaker3/homelab) | `D:\helm_clones_github\henrywhitaker3__homelab\kubernetes\k3s\apps\databases\nats\cluster\chart` | Error: no repository definition for https://nats-io.github.io/k8s/helm/charts/. Please add the missing repos via 'helm repo add' |
| 50 | [kubero-dev/kubero-operator](https://github.com/kubero-dev/kubero-operator) | `D:\helm_clones_github\kubero-dev__kubero-operator\helm-charts\kuberoaddonmongodb` | Error: no repository definition for https://groundhog2k.github.io/helm-charts/. Please add the missing repos via 'helm repo add' |
| 51 | [logicalisuki/ubiquity-open](https://github.com/logicalisuki/ubiquity-open) | `D:\helm_clones_github\logicalisuki__ubiquity-open\disabled\platform\opensm` | Getting updates for unmanaged Helm repositories... ...Unable to get an update from the "https://github.com/Mellanox/network-operator" chart repository: 	failed to fetch https://github.com/Mellanox/network-operator/index.yaml : 404 Not Found Error: no cached repository for helm-manager-bedd4d6f25f9f14d254639d8224675502e27945a13484e5a9e1499a78d72770e found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-bedd4d6f25f9f14d254639d8224675502e27945a13484e5a9e1499a78d72770e-index.yaml: O sistema não conseguiu localizar o ficheiro especificado. |
| 52 | [junghoon2/kube-books](https://github.com/junghoon2/kube-books) | `D:\helm_clones_github\junghoon2__kube-books\ch06\nginx-12.0.0` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 53 | [llajas/homelab](https://github.com/llajas/homelab) | `D:\helm_clones_github\llajas__homelab\apps\plex-apps` | Error: can't get a valid version for 1 subchart(s): "overseerr" (repository "file://./charts/overseerr", version "5.4.2"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml |
| 54 | [preloop/preloop](https://github.com/preloop/preloop) | `D:\helm_clones_github\preloop__preloop\helm\preloop` | Error: no repository definition for https://nats-io.github.io/k8s/helm/charts. Please add the missing repos via 'helm repo add' |
| 55 | [SimCubeLtd/simcube-helm-charts](https://github.com/SimCubeLtd/simcube-helm-charts) | `D:\helm_clones_github\SimCubeLtd__simcube-helm-charts\charts\bytesafe` | Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 56 | [stakater/nordmart-apps-gitops-config](https://github.com/stakater/nordmart-apps-gitops-config) | `D:\helm_clones_github\stakater__nordmart-apps-gitops-config\01-arsenal\01-stakater-nordmart-review-api\01-dev` | Getting updates for unmanaged Helm repositories... ...Unable to get an update from the "https://nexus-helm-stakater-nexus.apps.devtest.vxdqgl7u.kubeapp.cloud/repository/helm-charts/" chart repository: 	Get "https://nexus-helm-stakater-nexus.apps.devtest.vxdqgl7u.kubeapp.cloud/repository/helm-charts/index.yaml": dial tcp: lookup nexus-helm-stakater-nexus.apps.devtest.vxdqgl7u.kubeapp.cloud: no such host Error: no cached repository for helm-manager-7ebb8ed6883774d2c679cf4b093eaa1b7bd49e3f4401e7427ef1456c3315f23d found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-7ebb8ed6883774d2c679cf4b093eaa1b7bd49e3f4401e7427ef1456c3315f23d-index.yaml: O sistema não conseguiu localizar o ficheiro especificado. |
| 57 | [thoughtworks/byor-voting-infrastructure](https://github.com/thoughtworks/byor-voting-infrastructure) | `D:\helm_clones_github\thoughtworks__byor-voting-infrastructure\src\byor-voting-chart` | level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..." Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies |
| 58 | [starlingx/openstack-armada-app](https://github.com/starlingx/openstack-armada-app) | `D:\helm_clones_github\starlingx__openstack-armada-app\stx-openstack-helm-fluxcd\stx-openstack-helm-fluxcd\helm-charts\clients` | Getting updates for unmanaged Helm repositories... ...Unable to get an update from the "http://localhost:8879/charts" chart repository: 	Get "http://localhost:8879/charts/index.yaml": dial tcp [::1]:8879: connectex: Nenhuma ligação pôde ser feita porque o computador de destino as recusou ativamente. Error: no cached repository for helm-manager-878d619eb15837b169144dfaab3a7d6c5e800dd40daf0369bbe2b101f2275284 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-878d619eb15837b169144dfaab3a7d6c5e800dd40daf0369bbe2b101f2275284-index.yaml: O sistema não conseguiu localizar o ficheiro especificado. |
| 59 | [appscode-cloud/ui-wizards](https://github.com/appscode-cloud/ui-wizards) | `D:\helm_clones_github\appscode-cloud__ui-wizards\charts\kubedbcom-elasticsearch-editor` | Error: chart file "values.openapiv3_schema.yaml" is larger than the maximum file size 5242880 |
| 60 | [cloudstark/helm-charts](https://github.com/cloudstark/helm-charts) | `D:\helm_clones_github\cloudstark__helm-charts\postgrest` | level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..." Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add' |
| 61 | [dungdm93/shipyard](https://github.com/dungdm93/shipyard) | `D:\helm_clones_github\dungdm93__shipyard\helm\druid` | Getting updates for unmanaged Helm repositories... ...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository Error: can't get a valid version for 2 subchart(s): "zookeeper" (repository "https://charts.bitnami.com/bitnami", version "7.x.x"), "postgresql" (repository "https://charts.bitnami.com/bitnami", version "10.x.x"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml |
| 62 | [atsip76/k8s_asterisk_project](https://github.com/atsip76/k8s_asterisk_project) | `D:\helm_clones_github\atsip76__k8s_asterisk_project\k8s\gitlab` | level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..." Error: no repository definition for https://charts.gitlab.io/. Please add the missing repos via 'helm repo add' |
| 63 | [ibuildthecloud/rancher-charts](https://github.com/ibuildthecloud/rancher-charts) | `D:\helm_clones_github\ibuildthecloud__rancher-charts\charts\anchore-engine\0.1.0` | level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..." Error: no repository definition for https://kubernetes-charts.storage.googleapis.com. Please add the missing repos via 'helm repo add' |
| 64 | [kaikodata/canton-tooling](https://github.com/kaikodata/canton-tooling) | `D:\helm_clones_github\kaikodata__canton-tooling\kubernetes\templates\canton-validator-template` | Error: validation: chart.metadata.version "TEMPLATE_VERSION" is invalid |
| 65 | [teddy-ambona/kind-e2e](https://github.com/teddy-ambona/kind-e2e) | `D:\helm_clones_github\teddy-ambona__kind-e2e\helm\loki` | Error: no repository definition for https://grafana.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 66 | [CDCgov/NEDSS-Helm](https://github.com/CDCgov/NEDSS-Helm) | `D:\helm_clones_github\CDCgov__NEDSS-Helm\charts\strimzi` | Error: no repository definition for https://strimzi.io/charts/. Please add the missing repos via 'helm repo add' |
| 67 | [Otus-DevOps-2019-08/sgremyachikh_microservices](https://github.com/Otus-DevOps-2019-08/sgremyachikh_microservices) | `D:\helm_clones_github\Otus-DevOps-2019-08__sgremyachikh_microservices\kubernetes\Charts\gitlab-omnibus` | level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..." Error: no repository definition for https://charts.gitlab.io/. Please add the missing repos via 'helm repo add' |
| 68 | [SpechtLabs/k8s-deployment](https://github.com/SpechtLabs/k8s-deployment) | `D:\helm_clones_github\SpechtLabs__k8s-deployment\charts\cert-checker` | Error: no repository definition for https://mogensen.github.io/cert-checker. Please add the missing repos via 'helm repo add' |
| 69 | [claytono/infra](https://github.com/claytono/infra) | `D:\helm_clones_github\claytono__infra\kubernetes\crowdsec` | Error: no repository definition for https://crowdsecurity.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 70 | [griggheo/blogomatic](https://github.com/griggheo/blogomatic) | `D:\helm_clones_github\griggheo__blogomatic\devops\bootstrap_kind_cluster\helm_charts\signoz\signoz` | Error: no repository definition for https://signoz.github.io/charts, https://signoz.github.io/charts, https://charts.jetstack.io, https://kubernetes.github.io/ingress-nginx, https://charts.min.io, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 71 | [punchplatform/punch-helm](https://github.com/punchplatform/punch-helm) | `D:\helm_clones_github\punchplatform__punch-helm\operator` | Error: directory D:\helm_clones_github\punchplatform__punch-helm\operator\charts\operator.certificate not found |
| 72 | [EamonKeane/k8s-cluster-services](https://github.com/EamonKeane/k8s-cluster-services) | `D:\helm_clones_github\EamonKeane__k8s-cluster-services\cluster-svc` | level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..." Error: no repository definition for http://kubernetes-charts.storage.googleapis.com/, http://kubernetes-charts.storage.googleapis.com/, http://storage.googleapis.com/kubernetes-charts-incubator, http://kubernetes-charts.storage.googleapis.com/, http://storage.googleapis.com/kubernetes-charts-incubator, http://storage.googleapis.com/kubernetes-charts-incubator, http://kubernetes-charts.storage.googleapis.com/, http://kubernetes-charts.storage.googleapis.com/, http://kubernetes-charts.storage.googleapis.com/, http://kubernetes-charts.storage.googleapis.com/, http://kubernetes-charts.storage.googleapis.com/, https://opensource-helm.squareroute.io/, http://kubernetes-charts.storage.googleapis.com/, https://helm.github.io/monocular, http://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add' |
| 73 | [Makhuta/truecharts-archive-scale-catalog](https://github.com/Makhuta/truecharts-archive-scale-catalog) | `D:\helm_clones_github\Makhuta__truecharts-archive-scale-catalog\incubator\archivebox\0.7.2` | Saving 1 charts Downloading common from repo oci://tccr.io/truecharts Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:17.2.30: not found Error: could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:17.2.30: not found |
| 74 | [cnieg/helm-charts](https://github.com/cnieg/helm-charts) | `D:\helm_clones_github\cnieg__helm-charts\charts\clamapi` | Error: no repository definition for https://wiremind.github.io/wiremind-helm-charts. Please add the missing repos via 'helm repo add' |
| 75 | [dynatrace-wwse/enablement-kubernetes-opentelemetry-openpipeline](https://github.com/dynatrace-wwse/enablement-kubernetes-opentelemetry-openpipeline) | `D:\helm_clones_github\dynatrace-wwse__enablement-kubernetes-opentelemetry-openpipeline\.devcontainer\astroshop\helm\dt-otel-demo-helm` | Error: no repository definition for https://open-telemetry.github.io/opentelemetry-helm-charts, https://open-telemetry.github.io/opentelemetry-helm-charts. Please add the missing repos via 'helm repo add' |
| 76 | [helxplatform/translator-devops](https://github.com/helxplatform/translator-devops) | `D:\helm_clones_github\helxplatform__translator-devops\helm\answer-appraiser` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 77 | [kast-spells/kast-system](https://github.com/kast-spells/kast-system) | `D:\helm_clones_github\kast-spells__kast-system\librarian` | Error: error unpacking subchart common in librarian: Chart.yaml file is missing |
| 78 | [opspresso/argocd-env-addons](https://github.com/opspresso/argocd-env-addons) | `D:\helm_clones_github\opspresso__argocd-env-addons\charts\dashboard` | Getting updates for unmanaged Helm repositories... ...Unable to get an update from the "https://kubernetes.github.io/dashboard" chart repository: 	failed to fetch https://kubernetes.github.io/dashboard/index.yaml : 404 Not Found ...Successfully got an update from the "https://charts.helm.sh/incubator" chart repository Error: no cached repository for helm-manager-fc08c6c0f466a809ed2b24637e970ca3cd7bc1d7524efc4832f2405812f07ab0 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-fc08c6c0f466a809ed2b24637e970ca3cd7bc1d7524efc4832f2405812f07ab0-index.yaml: O sistema não conseguiu localizar o ficheiro especificado. |
| 79 | [pluralsh/plural-helm-charts](https://github.com/pluralsh/plural-helm-charts) | `D:\helm_clones_github\pluralsh__plural-helm-charts\charts\airbyte` | Error: no repository definition for https://charts.bitnami.com/bitnami, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts. Please add the missing repos via 'helm repo add' |
| 80 | [AchillesChan/memo](https://github.com/AchillesChan/memo) | `D:\helm_clones_github\AchillesChan__memo\helm-demo\prometheus-charts\charts\kube-prometheus-stack` | Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://grafana.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 81 | [AntSan813/react-hasura-keycloak-app](https://github.com/AntSan813/react-hasura-keycloak-app) | `D:\helm_clones_github\AntSan813__react-hasura-keycloak-app\api\hasura` | Error: no repository definition for https://hasura.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 82 | [KevMCarp/truecharts-catalog-fork](https://github.com/KevMCarp/truecharts-catalog-fork) | `D:\helm_clones_github\KevMCarp__truecharts-catalog-fork\dependency\clickhouse\5.0.54` | Getting updates for unmanaged Helm repositories... ...Unable to get an update from the "https://library-charts.truecharts.org" chart repository: 	Get "https://library-charts.truecharts.org/index.yaml": dial tcp: lookup library-charts.truecharts.org: no such host Error: no cached repository for helm-manager-024b189b59f6c6ccf0de6e5148db1578caf551c511f4eb220ece14cef00f80e0 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-024b189b59f6c6ccf0de6e5148db1578caf551c511f4eb220ece14cef00f80e0-index.yaml: O sistema não conseguiu localizar o ficheiro especificado. |
| 83 | [NeuraLegion/helmcharts](https://github.com/NeuraLegion/helmcharts) | `D:\helm_clones_github\NeuraLegion__helmcharts\charts\altoroj` | Saving 1 charts Save error occurred:  can't get a valid version for dependency simple-service Error: can't get a valid version for dependency simple-service |
| 84 | [Sureya/airflow_k8s_executor](https://github.com/Sureya/airflow_k8s_executor) | `D:\helm_clones_github\Sureya__airflow_k8s_executor\helm_charts\official\charts\incubator\distribution` | level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..." Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/, https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add' |
| 85 | [TheOpsDev/homelab](https://github.com/TheOpsDev/homelab) | `D:\helm_clones_github\TheOpsDev__homelab\charts\k8s-dashboard` | Getting updates for unmanaged Helm repositories... ...Unable to get an update from the "https://kubernetes.github.io/dashboard/" chart repository: 	failed to fetch https://kubernetes.github.io/dashboard/index.yaml : 404 Not Found Error: no cached repository for helm-manager-602693e8f5d1a68dc0300eb544f8e9829d89b7af15ee517b5231c07768425e69 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-602693e8f5d1a68dc0300eb544f8e9829d89b7af15ee517b5231c07768425e69-index.yaml: O sistema não conseguiu localizar o ficheiro especificado. |
| 86 | [devops4solutions/guestbook](https://github.com/devops4solutions/guestbook) | `D:\helm_clones_github\devops4solutions__guestbook` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 87 | [edixos/ekp-helm](https://github.com/edixos/ekp-helm) | `D:\helm_clones_github\edixos__ekp-helm\charts\alertmanager` | Error: no repository definition for https://oauth2-proxy.github.io/manifests. Please add the missing repos via 'helm repo add' |
| 88 | [elastic/k8s-integration-infra](https://github.com/elastic/k8s-integration-infra) | `D:\helm_clones_github\elastic__k8s-integration-infra\infra\charts\elastic-agent` | Error: no repository definition for @stable. Please add them via 'helm repo add' |
| 89 | [ishtiaqhimel/oms](https://github.com/ishtiaqhimel/oms) | `D:\helm_clones_github\ishtiaqhimel__oms\charts\oms-server` | Error: no repository definition for https://charts.konghq.com. Please add the missing repos via 'helm repo add' |
| 90 | [kalavai-net/helm-charts](https://github.com/kalavai-net/helm-charts) | `D:\helm_clones_github\kalavai-net__helm-charts\deployments\monitoring` | Error: no repository definition for https://grafana.github.io/helm-charts, https://ckotzbauer.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 91 | [lucidworks/ocp-fusion-helm-charts](https://github.com/lucidworks/ocp-fusion-helm-charts) | `D:\helm_clones_github\lucidworks__ocp-fusion-helm-charts\5.3.4\fusion\charts\admin-ui` | Getting updates for unmanaged Helm repositories... ...Successfully got an update from the "https://charts.lucidworks.com/" chart repository Error: can't get a valid version for 1 subchart(s): "fusion-common-utils" (repository "https://charts.lucidworks.com/", version "1.5.1"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml |
| 92 | [rancher/ob-team-charts](https://github.com/rancher/ob-team-charts) | `D:\helm_clones_github\rancher__ob-team-charts\charts\prometheus-federator\0.0.1` | Error: dependency "helmProjectOperator" has an invalid version/constraint format: improper constraint: "" |
| 93 | [shelleg/ac-k8s](https://github.com/shelleg/ac-k8s) | `D:\helm_clones_github\shelleg__ac-k8s\helm\ant-umbrella` | Error: no repository definition for @incubator, @ac-charts, @ac-charts, @ac-charts, @ac-charts. Please add them via 'helm repo add' |
| 94 | [shini4i/charts](https://github.com/shini4i/charts) | `D:\helm_clones_github\shini4i__charts\charts\app` | Error: no repository definition for https://shini4i.github.io/charts/. Please add the missing repos via 'helm repo add' |
| 95 | [tetratelabs/charts](https://github.com/tetratelabs/charts) | `D:\helm_clones_github\tetratelabs__charts\charts\demos\istio-monitoring-demo` | Error: no repository definition for https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 96 | [vlab-research/fly](https://github.com/vlab-research/fly) | `D:\helm_clones_github\vlab-research__fly\devops\vlab` | Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update' |
| 97 | [yunzck8s/cloudNative](https://github.com/yunzck8s/cloudNative) | `D:\helm_clones_github\yunzck8s__cloudNative\charts\deepflow` | Error: no repository definition for https://grafana.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 98 | [FIWARE-Ops/fiware-gitops](https://github.com/FIWARE-Ops/fiware-gitops) | `D:\helm_clones_github\FIWARE-Ops__fiware-gitops\aws\token\mongodb` | Getting updates for unmanaged Helm repositories... ...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository Error: can't get a valid version for 1 subchart(s): "mongodb" (repository "https://charts.bitnami.com/bitnami", version "11.0.4"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml |
| 99 | [Frndo1203/stack_iceberg_starrocks_k8s](https://github.com/Frndo1203/stack_iceberg_starrocks_k8s) | `D:\helm_clones_github\Frndo1203__stack_iceberg_starrocks_k8s\infra\src\helm-charts\airflow` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 100 | [JonasHess/homelab-iac](https://github.com/JonasHess/homelab-iac) | `D:\helm_clones_github\JonasHess__homelab-iac\apps\nextcloud` | Error: no repository definition for https://nextcloud.github.io/helm/. Please add the missing repos via 'helm repo add' |
| 101 | [Kapil-Bhalodiya/E-Commerce](https://github.com/Kapil-Bhalodiya/E-Commerce) | `D:\helm_clones_github\Kapil-Bhalodiya__E-Commerce\infra\addons\nginx-ingress` | Error: no repository definition for https://kubernetes.github.io/ingress-nginx. Please add the missing repos via 'helm repo add' |
| 102 | [Kapil-Bhalodiya/E-commerce-Platform](https://github.com/Kapil-Bhalodiya/E-commerce-Platform) | `D:\helm_clones_github\Kapil-Bhalodiya__E-commerce-Platform\infra\addons\frontend` | Error: cannot load values.yaml: cannot unmarshal yaml document: error converting YAML to JSON: yaml: line 8: could not find expected ':' |
| 103 | [LuukHors/homelab](https://github.com/LuukHors/homelab) | `D:\helm_clones_github\LuukHors__homelab\products\_base` | Error: validation: chart.metadata.name is required |
| 104 | [MrE-Fog/ks-installer2](https://github.com/MrE-Fog/ks-installer2) | `D:\helm_clones_github\MrE-Fog__ks-installer2\roles\ks-multicluster\files\kubefed\kubefed` | Getting updates for unmanaged Helm repositories... ...Unable to get an update from the "https://localhost/" chart repository: 	Get "https://localhost/index.yaml": dial tcp [::1]:443: connectex: Nenhuma ligação pôde ser feita porque o computador de destino as recusou ativamente. Error: no cached repository for helm-manager-f2b99ce05b94599549c70dbbe7a891b278e7c3cacad02334fa44682fca36c740 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-f2b99ce05b94599549c70dbbe7a891b278e7c3cacad02334fa44682fca36c740-index.yaml: O sistema não conseguiu localizar o ficheiro especificado. |
| 105 | [PilotDataPlatform/helm-charts](https://github.com/PilotDataPlatform/helm-charts) | `D:\helm_clones_github\PilotDataPlatform__helm-charts\argo-cd-917` | Error: no repository definition for https://dandydeveloper.github.io/charts/. Please add the missing repos via 'helm repo add' |
| 106 | [SpelmanDevops/retail-store](https://github.com/SpelmanDevops/retail-store) | `D:\helm_clones_github\SpelmanDevops__retail-store\helm\monitoring` | Error: no repository definition for https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 107 | [TSMC-NYCU-LAB-13/infrastructures](https://github.com/TSMC-NYCU-LAB-13/infrastructures) | `D:\helm_clones_github\TSMC-NYCU-LAB-13__infrastructures\argo\argo-cd` | Error: no repository definition for https://dandydeveloper.github.io/charts/. Please add the missing repos via 'helm repo add' |
| 108 | [VadimShtukan/otus_homework](https://github.com/VadimShtukan/otus_homework) | `D:\helm_clones_github\VadimShtukan__otus_homework\architect\lesson05\kubernetis\helm-chart` | Error: no repository definition for https://kubernetes-charts.storage.googleapis.com. Please add the missing repos via 'helm repo add' |
| 109 | [Vaibhav2goyal/alertmanager](https://github.com/Vaibhav2goyal/alertmanager) | `D:\helm_clones_github\Vaibhav2goyal__alertmanager\scripts\kube-prometheus-stack` | Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 110 | [WesleyJw/modern-data-stack](https://github.com/WesleyJw/modern-data-stack) | `D:\helm_clones_github\WesleyJw__modern-data-stack\infra\src\helm-charts\airflow` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 111 | [Yahya-rabii/soge-plus](https://github.com/Yahya-rabii/soge-plus) | `D:\helm_clones_github\Yahya-rabii__soge-plus\helm-charts\kube-prometheus-stack` | Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 112 | [adstanley/archive](https://github.com/adstanley/archive) | `D:\helm_clones_github\adstanley__archive\scale-catalog\incubator\archivebox\0.7.2` | Saving 1 charts Downloading common from repo oci://tccr.io/truecharts Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:17.2.30: not found Error: could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:17.2.30: not found |
| 113 | [codefuturist/helm-charts](https://github.com/codefuturist/helm-charts) | `D:\helm_clones_github\codefuturist__helm-charts\templates\chart-template` | Error: validation: chart.metadata.name is required |
| 114 | [dan1dan12345678/Helm_charts](https://github.com/dan1dan12345678/Helm_charts) | `D:\helm_clones_github\dan1dan12345678__Helm_charts\kube-prometheus-stack` | Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 115 | [davidlesicnik/homelab-argo](https://github.com/davidlesicnik/homelab-argo) | `D:\helm_clones_github\davidlesicnik__homelab-argo\apps\grafana` | Error: no repository definition for https://grafana.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 116 | [dboeckli/spring-6-icecold-micro-service](https://github.com/dboeckli/spring-6-icecold-micro-service) | `D:\helm_clones_github\dboeckli__spring-6-icecold-micro-service\helm-charts` | Error: cannot load Chart.yaml: error converting YAML to JSON: yaml: line 4: found character that cannot start any token |
| 117 | [hey101/scale-catalog](https://github.com/hey101/scale-catalog) | `D:\helm_clones_github\hey101__scale-catalog\incubator\archivebox\0.7.2` | Saving 1 charts Downloading common from repo oci://tccr.io/truecharts Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:17.5.1: not found Error: could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:17.5.1: not found |
| 118 | [hmcts/hmcts-charts](https://github.com/hmcts/hmcts-charts) | `D:\helm_clones_github\hmcts__hmcts-charts\stable\aac-manage-case-assignment` | Getting updates for unmanaged Helm repositories... ...Successfully got an update from the "https://helm.elastic.co" chart repository Saving 4 charts Downloading java from repo oci://hmctsprod.azurecr.io/helm Save error occurred:  could not download oci://hmctsprod.azurecr.io/helm/java: failed to perform "FetchReference" on source: GET "https://hmctsprod.azurecr.io/v2/helm/java/manifests/5.3.0": GET "https://hmctsprod.azurecr.io/oauth2/token?scope=repository%3Ahelm%2Fjava%3Apull&service=hmctsprod.azurecr.io": response status code 401: unauthorized: authentication required, visit https://aka.ms/acr/authorization for more information. CorrelationId: f65eaf6b-2681-42ff-b433-3789c00e163f Error: could not download oci://hmctsprod.azurecr.io/helm/java: failed to perform "FetchReference" on source: GET "https://hmctsprod.azurecr.io/v2/helm/java/manifests/5.3.0": GET "https://hmctsprod.azurecr.io/oauth2/token?scope=repository%3Ahelm%2Fjava%3Apull&service=hmctsprod.azurecr.io": response status code 401: unauthorized: authentication required, visit https://aka.ms/acr/authorization for more information. CorrelationId: f65eaf6b-2681-42ff-b433-3789c00e163f |
| 119 | [legion-platform/legion-infrastructure](https://github.com/legion-platform/legion-infrastructure) | `D:\helm_clones_github\legion-platform__legion-infrastructure\helms\monitoring` | Error: no repository definition for @stable, @stable. Please add them via 'helm repo add' |
| 120 | [merlindorin/charts](https://github.com/merlindorin/charts) | `D:\helm_clones_github\merlindorin__charts\charts\pinniped` | Error: no repository definition for https://merlindorin.github.io/charts, https://merlindorin.github.io/charts. Please add the missing repos via 'helm repo add' |
| 121 | [nwthomas/gitops](https://github.com/nwthomas/gitops) | `D:\helm_clones_github\nwthomas__gitops\helm\longhorn` | Error: no repository definition for https://charts.longhorn.io. Please add the missing repos via 'helm repo add' |
| 122 | [otus-kuber-2019-12/gidmaster_platform](https://github.com/otus-kuber-2019-12/gidmaster_platform) | `D:\helm_clones_github\otus-kuber-2019-12__gidmaster_platform\kubernetes-gitops\deploy\charts\cartservice` | Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add' |
| 123 | [rtang03/helm-charts](https://github.com/rtang03/helm-charts) | `D:\helm_clones_github\rtang03__helm-charts\charts\argocd` | Error: no repository definition for https://argoproj.github.io/argo-helm. Please add the missing repos via 'helm repo add' |
| 124 | [tetratelabs/helm-charts](https://github.com/tetratelabs/helm-charts) | `D:\helm_clones_github\tetratelabs__helm-charts\charts\demos\istio-monitoring-demo` | Error: no repository definition for https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 125 | [100CallsToEurop/otus](https://github.com/100CallsToEurop/otus) | `D:\helm_clones_github\100CallsToEurop__otus\k8s\auth` | Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update' |
| 126 | [100rd/platform-design](https://github.com/100rd/platform-design) | `D:\helm_clones_github\100rd__platform-design\apps\infra\cilium` | Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update' |
| 127 | [AdmanTIC/helm-charts](https://github.com/AdmanTIC/helm-charts) | `D:\helm_clones_github\AdmanTIC__helm-charts\charts\cremecrm` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 128 | [Arthur-B-DevOps/old_helm_charts](https://github.com/Arthur-B-DevOps/old_helm_charts) | `D:\helm_clones_github\Arthur-B-DevOps__old_helm_charts\charts\Old_charts\charts\incubator\distribution` | level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..." Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/, https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add' |
| 129 | [Avichayef/datateam_calculator](https://github.com/Avichayef/datateam_calculator) | `D:\helm_clones_github\Avichayef__datateam_calculator\helm\jenkins` | Error: no repository definition for https://charts.jenkins.io. Please add the missing repos via 'helm repo add' |
| 130 | [Backstage-Epitech/cltest](https://github.com/Backstage-Epitech/cltest) | `D:\helm_clones_github\Backstage-Epitech__cltest\kustomize\kube-prometheus-stack\charts\kube-prometheus-stack-73.2.2\kube-prometheus-stack` | Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 131 | [Bernardpro/ClusterGCP](https://github.com/Bernardpro/ClusterGCP) | `D:\helm_clones_github\Bernardpro__ClusterGCP\kustomize\kube-prometheus-stack\charts\kube-prometheus-stack-73.2.2\kube-prometheus-stack` | Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 132 | [Bernardpro/ClusterGKE](https://github.com/Bernardpro/ClusterGKE) | `D:\helm_clones_github\Bernardpro__ClusterGKE\kustomize\kube-prometheus-stack\charts\kube-prometheus-stack-73.2.2\kube-prometheus-stack` | Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 133 | [Clark1992/ECK1](https://github.com/Clark1992/ECK1) | `D:\helm_clones_github\Clark1992__ECK1\src\Integration\ECK1.FailedViewRebuilder\Deploy\service` | Saving 1 charts Downloading config-watcher from repo oci://localhost:5000/helm Save error occurred:  could not download oci://localhost:5000/helm/config-watcher: failed to perform "FetchReference" on source: Get "https://localhost:5000/v2/helm/config-watcher/manifests/0.1.0": dial tcp [::1]:5000: connectex: Nenhuma ligação pôde ser feita porque o computador de destino as recusou ativamente. Error: could not download oci://localhost:5000/helm/config-watcher: failed to perform "FetchReference" on source: Get "https://localhost:5000/v2/helm/config-watcher/manifests/0.1.0": dial tcp [::1]:5000: connectex: Nenhuma ligação pôde ser feita porque o computador de destino as recusou ativamente. |

## Full Error Details

### `D:\helm_clones_github\rancher__rancher\chart` — rancher/rancher

```
Error: cannot load Chart.yaml: error converting YAML to JSON: yaml: line 4: found character that cannot start any token
```

### `D:\helm_clones_github\refly-ai__refly\deploy\helm\refly-stack` — refly-ai/refly

```
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\linode__apl-core\chart\chart-index` — linode/apl-core

```
Getting updates for unmanaged Helm repositories...
...Successfully got an update from the "https://linode.github.io/cloud-firewall-controller" chart repository
...Successfully got an update from the "https://prometheus-msteams.github.io/prometheus-msteams/" chart repository
...Successfully got an update from the "https://kubernetes-sigs.github.io/metrics-server" chart repository
...Successfully got an update from the "https://cdfoundation.github.io/tekton-helm-chart/" chart repository
...Successfully got an update from the "https://knative.github.io/operator" chart repository
...Successfully got an update from the "https://kubernetes-sigs.github.io/external-dns" chart repository
...Successfully got an update from the "https://bitnami-labs.github.io/sealed-secrets/" chart repository
...Successfully got an update from the "https://kyverno.github.io/policy-reporter" chart repository
...Successfully got an update from the "https://aquasecurity.github.io/helm-charts/" chart repository
...Successfully got an update from the "https://oauth2-proxy.github.io/manifests" chart repository
...Successfully got an update from the "https://codecentric.github.io/helm-charts" chart repository
...Successfully got an update from the "https://kyverno.github.io/kyverno/" chart repository
...Successfully got an update from the "https://cloudnative-pg.github.io/charts" chart repository
...Successfully got an update from the "https://charts.jetstack.io" chart repository
...Successfully got an update from the "https://open-telemetry.github.io/opentelemetry-helm-charts" chart repository
...Successfully got an update from the "https://argoproj.github.io/argo-helm" chart repository
...Successfully got an update from the "https://dl.gitea.io/charts" chart repository
...Successfully got an update from the "https://helm.goharbor.io" chart repository
...Successfully got an update from the "https://grafana.github.io/helm-charts" chart repository
...Successfully got an update from the "https://prometheus-community.github.io/helm-charts" chart repository
...Successfully got an update from the "https://charts.external-secrets.io" chart repository
...Successfully got an update from the "https://istio-release.storage.googleapis.com/charts" chart repository
...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository
Saving 31 charts
Downloading argo-cd from repo https://argoproj.github.io/argo-helm
Downloading argocd-image-updater from repo oci://ghcr.io/argoproj/argo-helm/argocd-image-updater
Save error occurred:  could not download oci://ghcr.io/argoproj/argo-helm/argocd-image-updater/argocd-image-updater: failed to perform "FetchReference" on source: GET "https://ghcr.io/v2/argoproj/argo-helm/argocd-image-updater/argocd-image-updater/manifests/1.1.5": GET "https://ghcr.io/token?scope=repository%3Aargoproj%2Fargo-helm%2Fargocd-image-updater%2Fargocd-image-updater%3Apull&service=ghcr.io": response status code 403: denied: requested access to the resource is denied
Error: could not download oci://ghcr.io/argoproj/argo-helm/argocd-image-updater/argocd-image-updater: failed to perform "FetchReference" on source: GET "https://ghcr.io/v2/argoproj/argo-helm/argocd-image-updater/argocd-image-updater/manifests/1.1.5": GET "https://ghcr.io/token?scope=repository%3Aargoproj%2Fargo-helm%2Fargocd-image-updater%2Fargocd-image-updater%3Apull&service=ghcr.io": response status code 403: denied: requested access to the resource is denied
```

### `D:\helm_clones_github\cozystack__cozystack\packages\apps\bucket` — cozystack/cozystack

```
Error: error unpacking subchart cozy-lib in bucket: Chart.yaml file is missing
```

### `D:\helm_clones_github\grafana__helm-charts\charts\enterprise-metrics` — grafana/helm-charts

```
Getting updates for unmanaged Helm repositories...
...Successfully got an update from the "https://helm.min.io/" chart repository
...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository
Error: minio chart not found in repo https://helm.min.io/
```

### `D:\helm_clones_github\norwoodj__helm-docs\example-charts\custom-template` — norwoodj/helm-docs

```
Error: no repository definition for @stable. Please add them via 'helm repo add'
```

### `D:\helm_clones_github\securitybunker__databunker\charts\databunker` — securitybunker/databunker

```
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\WeBankFinTech__Prophecis\install\Prophecis` — WeBankFinTech/Prophecis

```
Error: cannot load values.yaml: error reading yaml document: invalid Yaml document separator: --END RSA PRIVATE KEY-----"
```

### `D:\helm_clones_github\rancher__charts\charts\epinio\102.0.1+up1.6.2` — rancher/charts

```
Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update'
```

### `D:\helm_clones_github\cloudnativeapp__charts\curated\airflow` — cloudnativeapp/charts

```
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies
```

### `D:\helm_clones_github\BigKAA__youtube\tracing\for_admins\charts\jaeger\jaeger` — BigKAA/youtube

```
Error: no repository definition for https://charts.helm.sh/incubator, https://helm.elastic.co, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\radondb__radondb-mysql-kubernetes\charts\mysql-operator` — radondb/radondb-mysql-kubernetes

```
Error: directory D:\helm_clones_github\radondb__radondb-mysql-kubernetes\charts\mysql-operator\charts\mysqlcluster not found
```

### `D:\helm_clones_github\IBM__charts\community\artifactory-ha` — IBM/charts

```
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\JahstreetOrg__spark-on-kubernetes-helm\charts\cluster-base` — JahstreetOrg/spark-on-kubernetes-helm

```
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: no repository definition for https://kubernetes.github.io/ingress-nginx, https://charts.jetstack.io, https://charts.helm.sh/stable, https://charts.helm.sh/stable. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\vexxhost__atmosphere\charts\barbican` — vexxhost/atmosphere

```
level=INFO msg="Warning: Dependency locking is handled in Chart.lock since apiVersion \"v2\". We recommend migrating to Chart.lock."
Error: no repository definition for https://tarballs.openstack.org/openstack-helm. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\unixhot__devops-x\helm\gitlab` — unixhot/devops-x

```
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies
```

### `D:\helm_clones_github\rancher__rke2-charts\packages\rke2-cilium-legacy\charts` — rancher/rke2-charts

```
Error: directory D:\helm_clones_github\rancher__rke2-charts\packages\rke2-cilium-legacy\charts\charts\cilium not found
```

### `D:\helm_clones_github\stackrox__helm-charts\3.0.41.0` — stackrox/helm-charts

```
Error: validation: chart.metadata.version "3.0.41.0" is invalid
```

### `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\monitoring\oauth2-proxy` — project-sunbird/sunbird-devops

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\rancher__partner-charts\charts\amd\amd-gpu\0.10.0` — rancher/partner-charts

```
Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update'
```

### `D:\helm_clones_github\CARV-ICS-FORTH__frisbee\charts\platform` — CARV-ICS-FORTH/frisbee

```
Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update'
```

### `D:\helm_clones_github\wikibook__kubepractice\ch06\nginx-12.0.0` — wikibook/kubepractice

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\pluralsh__plural-artifacts\airbyte\helm\airbyte` — pluralsh/plural-artifacts

```
Error: directory D:\helm_clones_github\airbyte\charts\airbyte not found
```

### `D:\helm_clones_github\tmforum-oda__oda-canvas\charts\observability-stack` — tmforum-oda/oda-canvas

```
Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://open-telemetry.github.io/opentelemetry-helm-charts, https://jaegertracing.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\sa-mw-dach__bobbycar\helm\bobbycar-core-infra` — sa-mw-dach/bobbycar

```
Error: no repository definition for https://drogue-iot.github.io/drogue-cloud-helm-charts/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\boozallen__aissemble\foundation\foundation-archetype\src\main\resources\archetype-resources\__rootArtifactId__-deploy\src\main\resources\apps\common-infrastructure` — boozallen/aissemble

```
Error: validation: chart.metadata.version "${version}" is invalid
```

### `D:\helm_clones_github\tkestack__charts\incubator\airflow` — tkestack/charts

```
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/, https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\sapcc__helm-charts\common\inventory-updater` — sapcc/helm-charts

```
Saving 1 charts
Downloading owner-info from repo oci://keppel.eu-de-1.cloud.sap/ccloud-helm
Save error occurred:  could not download oci://keppel.eu-de-1.cloud.sap/ccloud-helm/owner-info: failed to perform "FetchReference" on source: Get "https://keppel.eu-de-1.cloud.sap/v2/ccloud-helm/owner-info/manifests/0.2.0": dial tcp: lookup keppel.eu-de-1.cloud.sap: getaddrinfow: Este é geralmente um erro temporário durante a resolução de nomes de anfitrião e significa que o servidor local não recebeu uma resposta de um servidor autoritário.
Error: could not download oci://keppel.eu-de-1.cloud.sap/ccloud-helm/owner-info: failed to perform "FetchReference" on source: Get "https://keppel.eu-de-1.cloud.sap/v2/ccloud-helm/owner-info/manifests/0.2.0": dial tcp: lookup keppel.eu-de-1.cloud.sap: getaddrinfow: Este é geralmente um erro temporário durante a resolução de nomes de anfitrião e significa que o servidor local não recebeu uma resposta de um servidor autoritário.
```

### `D:\helm_clones_github\platyplus__platyplus\charts\hasura` — platyplus/platyplus

```
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\opendevstack__ods-quickstarters\be-rust-axum\rust-template\chart` — opendevstack/ods-quickstarters

```
Error: cannot load Chart.yaml: error converting YAML to JSON: yaml: invalid map key: map[interface {}]interface {}{"project-name":interface {}(nil)}
```

### `D:\helm_clones_github\k8s-home-lab__helm-charts\unmaintained\audiobookshelf` — k8s-home-lab/helm-charts

```
Getting updates for unmanaged Helm repositories...
...Successfully got an update from the "https://library-charts.k8s-at-home.com" chart repository
Error: can't get a valid version for 1 subchart(s): "common" (repository "https://library-charts.k8s-at-home.com", version "4.5.3"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml
```

### `D:\helm_clones_github\mojaloop__helm\perf-test-harness` — mojaloop/helm

```
Error: can't get a valid version for 1 subchart(s): "ml-testing-toolkit-cli" (repository "file://../ml-testing-toolkit-cli", version "15.9.0"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml
```

### `D:\helm_clones_github\camptocamp__charts\common-build-code` — camptocamp/charts

```
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "https://kubernetes-charts-incubator.storage.googleapis.com" chart repository:
	failed to fetch https://kubernetes-charts-incubator.storage.googleapis.com/index.yaml : 403 Forbidden
Error: no cached repository for helm-manager-53271637451a5b2439ffd0af71673734b808e371a8a6aed9bf100a8f219a3006 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-53271637451a5b2439ffd0af71673734b808e371a8a6aed9bf100a8f219a3006-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### `D:\helm_clones_github\devtron-labs__charts\charts\cluster-essentials` — devtron-labs/charts

```
Getting updates for unmanaged Helm repositories...
...Successfully got an update from the "https://kedacore.github.io/charts" chart repository
...Successfully got an update from the "https://kubernetes.github.io/autoscaler" chart repository
...Successfully got an update from the "https://helm.devtron.ai" chart repository
...Successfully got an update from the "https://aws.github.io/eks-charts" chart repository
...Successfully got an update from the "https://kubernetes-sigs.github.io/metrics-server/" chart repository
...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository
Error: can't get a valid version for 1 subchart(s): "kubernetes-event-exporter" (repository "https://charts.bitnami.com/bitnami", version "1.2.*"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml
```

### `D:\helm_clones_github\YAKEcloud__yake\helmcharts\acl` — YAKEcloud/yake

```
Error: dependency "controller" has an invalid version/constraint format: improper constraint: ""
```

### `D:\helm_clones_github\bflance__proxmox-talos\charts\kube-prometheus-stack` — bflance/proxmox-talos

```
Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\k0rdent__catalog\apps\alloy\charts\alloy-1.6.1` — k0rdent/catalog

```
Error: no repository definition for https://grafana.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\Loongson-Cloud-Community__dockerfiles\kubesphere\ks-installer\v3.2.1\roles\ks-multicluster\files\kubefed\kubefed` — Loongson-Cloud-Community/dockerfiles

```
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "https://localhost/" chart repository:
	Get "https://localhost/index.yaml": dial tcp [::1]:443: connectex: Nenhuma ligação pôde ser feita porque o computador de destino
as recusou ativamente.
Error: no cached repository for helm-manager-f2b99ce05b94599549c70dbbe7a891b278e7c3cacad02334fa44682fca36c740 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-f2b99ce05b94599549c70dbbe7a891b278e7c3cacad02334fa44682fca36c740-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### `D:\helm_clones_github\junghoon2__k8s-class\argo-cd\argo-cd-5.14.1` — junghoon2/k8s-class

```
Error: no repository definition for https://dandydeveloper.github.io/charts/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\cloud-native-toolkit__toolkit-charts\stable\cloud-setup` — cloud-native-toolkit/toolkit-charts

```
Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update'
```

### `D:\helm_clones_github\suse-edge__charts\charts\kubevirt\0.1.0` — suse-edge/charts

```
Error: dependency "cdi" has an invalid version/constraint format: improper constraint: ""
```

### `D:\helm_clones_github\ai-solution-eng__frameworks\appsmith\3.6.4` — ai-solution-eng/frameworks

```
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\cnrancher__pandaria-catalog\charts\rancher-hami\107.0.0+up2.5.2\charts\hami-webui` — cnrancher/pandaria-catalog

```
Error: no repository definition for https://nvidia.github.io/dcgm-exporter/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\Sagar2366__tech-talks\k8s_pune_oct22\prometheus-comunity-helm-chart\charts\kube-prometheus-stack` — Sagar2366/tech-talks

```
Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://grafana.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\boozallen__sdp-helm-chart` — boozallen/sdp-helm-chart

```
Error: dependency "jenkins" has an invalid version/constraint format: improper constraint: ""
```

### `D:\helm_clones_github\randoli__helm-charts\charts\cost-management` — randoli/helm-charts

```
Error: no repository definition for https://opencost.github.io/opencost-helm-chart. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\unixfox__k8s\charts\bibliogram` — unixfox/k8s

```
Error: no repository definition for https://library-charts.k8s-at-home.com. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\bcgov__OCWA\helm\ocwa` — bcgov/OCWA

```
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies
```

### `D:\helm_clones_github\henrywhitaker3__homelab\kubernetes\k3s\apps\databases\nats\cluster\chart` — henrywhitaker3/homelab

```
Error: no repository definition for https://nats-io.github.io/k8s/helm/charts/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\kubero-dev__kubero-operator\helm-charts\kuberoaddonmongodb` — kubero-dev/kubero-operator

```
Error: no repository definition for https://groundhog2k.github.io/helm-charts/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\logicalisuki__ubiquity-open\disabled\platform\opensm` — logicalisuki/ubiquity-open

```
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "https://github.com/Mellanox/network-operator" chart repository:
	failed to fetch https://github.com/Mellanox/network-operator/index.yaml : 404 Not Found
Error: no cached repository for helm-manager-bedd4d6f25f9f14d254639d8224675502e27945a13484e5a9e1499a78d72770e found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-bedd4d6f25f9f14d254639d8224675502e27945a13484e5a9e1499a78d72770e-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### `D:\helm_clones_github\junghoon2__kube-books\ch06\nginx-12.0.0` — junghoon2/kube-books

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\llajas__homelab\apps\plex-apps` — llajas/homelab

```
Error: can't get a valid version for 1 subchart(s): "overseerr" (repository "file://./charts/overseerr", version "5.4.2"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml
```

### `D:\helm_clones_github\preloop__preloop\helm\preloop` — preloop/preloop

```
Error: no repository definition for https://nats-io.github.io/k8s/helm/charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\SimCubeLtd__simcube-helm-charts\charts\bytesafe` — SimCubeLtd/simcube-helm-charts

```
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\stakater__nordmart-apps-gitops-config\01-arsenal\01-stakater-nordmart-review-api\01-dev` — stakater/nordmart-apps-gitops-config

```
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "https://nexus-helm-stakater-nexus.apps.devtest.vxdqgl7u.kubeapp.cloud/repository/helm-charts/" chart repository:
	Get "https://nexus-helm-stakater-nexus.apps.devtest.vxdqgl7u.kubeapp.cloud/repository/helm-charts/index.yaml": dial tcp: lookup nexus-helm-stakater-nexus.apps.devtest.vxdqgl7u.kubeapp.cloud: no such host
Error: no cached repository for helm-manager-7ebb8ed6883774d2c679cf4b093eaa1b7bd49e3f4401e7427ef1456c3315f23d found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-7ebb8ed6883774d2c679cf4b093eaa1b7bd49e3f4401e7427ef1456c3315f23d-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### `D:\helm_clones_github\thoughtworks__byor-voting-infrastructure\src\byor-voting-chart` — thoughtworks/byor-voting-infrastructure

```
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies
```

### `D:\helm_clones_github\starlingx__openstack-armada-app\stx-openstack-helm-fluxcd\stx-openstack-helm-fluxcd\helm-charts\clients` — starlingx/openstack-armada-app

```
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "http://localhost:8879/charts" chart repository:
	Get "http://localhost:8879/charts/index.yaml": dial tcp [::1]:8879: connectex: Nenhuma ligação pôde ser feita porque o computador de destino
as recusou ativamente.
Error: no cached repository for helm-manager-878d619eb15837b169144dfaab3a7d6c5e800dd40daf0369bbe2b101f2275284 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-878d619eb15837b169144dfaab3a7d6c5e800dd40daf0369bbe2b101f2275284-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### `D:\helm_clones_github\appscode-cloud__ui-wizards\charts\kubedbcom-elasticsearch-editor` — appscode-cloud/ui-wizards

```
Error: chart file "values.openapiv3_schema.yaml" is larger than the maximum file size 5242880
```

### `D:\helm_clones_github\cloudstark__helm-charts\postgrest` — cloudstark/helm-charts

```
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\dungdm93__shipyard\helm\druid` — dungdm93/shipyard

```
Getting updates for unmanaged Helm repositories...
...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository
Error: can't get a valid version for 2 subchart(s): "zookeeper" (repository "https://charts.bitnami.com/bitnami", version "7.x.x"), "postgresql" (repository "https://charts.bitnami.com/bitnami", version "10.x.x"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml
```

### `D:\helm_clones_github\atsip76__k8s_asterisk_project\k8s\gitlab` — atsip76/k8s_asterisk_project

```
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: no repository definition for https://charts.gitlab.io/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\ibuildthecloud__rancher-charts\charts\anchore-engine\0.1.0` — ibuildthecloud/rancher-charts

```
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: no repository definition for https://kubernetes-charts.storage.googleapis.com. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\kaikodata__canton-tooling\kubernetes\templates\canton-validator-template` — kaikodata/canton-tooling

```
Error: validation: chart.metadata.version "TEMPLATE_VERSION" is invalid
```

### `D:\helm_clones_github\teddy-ambona__kind-e2e\helm\loki` — teddy-ambona/kind-e2e

```
Error: no repository definition for https://grafana.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\CDCgov__NEDSS-Helm\charts\strimzi` — CDCgov/NEDSS-Helm

```
Error: no repository definition for https://strimzi.io/charts/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\Otus-DevOps-2019-08__sgremyachikh_microservices\kubernetes\Charts\gitlab-omnibus` — Otus-DevOps-2019-08/sgremyachikh_microservices

```
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: no repository definition for https://charts.gitlab.io/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\SpechtLabs__k8s-deployment\charts\cert-checker` — SpechtLabs/k8s-deployment

```
Error: no repository definition for https://mogensen.github.io/cert-checker. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\claytono__infra\kubernetes\crowdsec` — claytono/infra

```
Error: no repository definition for https://crowdsecurity.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\griggheo__blogomatic\devops\bootstrap_kind_cluster\helm_charts\signoz\signoz` — griggheo/blogomatic

```
Error: no repository definition for https://signoz.github.io/charts, https://signoz.github.io/charts, https://charts.jetstack.io, https://kubernetes.github.io/ingress-nginx, https://charts.min.io, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\punchplatform__punch-helm\operator` — punchplatform/punch-helm

```
Error: directory D:\helm_clones_github\punchplatform__punch-helm\operator\charts\operator.certificate not found
```

### `D:\helm_clones_github\EamonKeane__k8s-cluster-services\cluster-svc` — EamonKeane/k8s-cluster-services

```
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: no repository definition for http://kubernetes-charts.storage.googleapis.com/, http://kubernetes-charts.storage.googleapis.com/, http://storage.googleapis.com/kubernetes-charts-incubator, http://kubernetes-charts.storage.googleapis.com/, http://storage.googleapis.com/kubernetes-charts-incubator, http://storage.googleapis.com/kubernetes-charts-incubator, http://kubernetes-charts.storage.googleapis.com/, http://kubernetes-charts.storage.googleapis.com/, http://kubernetes-charts.storage.googleapis.com/, http://kubernetes-charts.storage.googleapis.com/, http://kubernetes-charts.storage.googleapis.com/, https://opensource-helm.squareroute.io/, http://kubernetes-charts.storage.googleapis.com/, https://helm.github.io/monocular, http://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\Makhuta__truecharts-archive-scale-catalog\incubator\archivebox\0.7.2` — Makhuta/truecharts-archive-scale-catalog

```
Saving 1 charts
Downloading common from repo oci://tccr.io/truecharts
Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:17.2.30: not found
Error: could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:17.2.30: not found
```

### `D:\helm_clones_github\cnieg__helm-charts\charts\clamapi` — cnieg/helm-charts

```
Error: no repository definition for https://wiremind.github.io/wiremind-helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\dynatrace-wwse__enablement-kubernetes-opentelemetry-openpipeline\.devcontainer\astroshop\helm\dt-otel-demo-helm` — dynatrace-wwse/enablement-kubernetes-opentelemetry-openpipeline

```
Error: no repository definition for https://open-telemetry.github.io/opentelemetry-helm-charts, https://open-telemetry.github.io/opentelemetry-helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\helxplatform__translator-devops\helm\answer-appraiser` — helxplatform/translator-devops

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\kast-spells__kast-system\librarian` — kast-spells/kast-system

```
Error: error unpacking subchart common in librarian: Chart.yaml file is missing
```

### `D:\helm_clones_github\opspresso__argocd-env-addons\charts\dashboard` — opspresso/argocd-env-addons

```
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "https://kubernetes.github.io/dashboard" chart repository:
	failed to fetch https://kubernetes.github.io/dashboard/index.yaml : 404 Not Found
...Successfully got an update from the "https://charts.helm.sh/incubator" chart repository
Error: no cached repository for helm-manager-fc08c6c0f466a809ed2b24637e970ca3cd7bc1d7524efc4832f2405812f07ab0 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-fc08c6c0f466a809ed2b24637e970ca3cd7bc1d7524efc4832f2405812f07ab0-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### `D:\helm_clones_github\pluralsh__plural-helm-charts\charts\airbyte` — pluralsh/plural-helm-charts

```
Error: no repository definition for https://charts.bitnami.com/bitnami, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\AchillesChan__memo\helm-demo\prometheus-charts\charts\kube-prometheus-stack` — AchillesChan/memo

```
Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://grafana.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\AntSan813__react-hasura-keycloak-app\api\hasura` — AntSan813/react-hasura-keycloak-app

```
Error: no repository definition for https://hasura.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\KevMCarp__truecharts-catalog-fork\dependency\clickhouse\5.0.54` — KevMCarp/truecharts-catalog-fork

```
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "https://library-charts.truecharts.org" chart repository:
	Get "https://library-charts.truecharts.org/index.yaml": dial tcp: lookup library-charts.truecharts.org: no such host
Error: no cached repository for helm-manager-024b189b59f6c6ccf0de6e5148db1578caf551c511f4eb220ece14cef00f80e0 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-024b189b59f6c6ccf0de6e5148db1578caf551c511f4eb220ece14cef00f80e0-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### `D:\helm_clones_github\NeuraLegion__helmcharts\charts\altoroj` — NeuraLegion/helmcharts

```
Saving 1 charts
Save error occurred:  can't get a valid version for dependency simple-service
Error: can't get a valid version for dependency simple-service
```

### `D:\helm_clones_github\Sureya__airflow_k8s_executor\helm_charts\official\charts\incubator\distribution` — Sureya/airflow_k8s_executor

```
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/, https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\TheOpsDev__homelab\charts\k8s-dashboard` — TheOpsDev/homelab

```
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "https://kubernetes.github.io/dashboard/" chart repository:
	failed to fetch https://kubernetes.github.io/dashboard/index.yaml : 404 Not Found
Error: no cached repository for helm-manager-602693e8f5d1a68dc0300eb544f8e9829d89b7af15ee517b5231c07768425e69 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-602693e8f5d1a68dc0300eb544f8e9829d89b7af15ee517b5231c07768425e69-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### `D:\helm_clones_github\devops4solutions__guestbook` — devops4solutions/guestbook

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\edixos__ekp-helm\charts\alertmanager` — edixos/ekp-helm

```
Error: no repository definition for https://oauth2-proxy.github.io/manifests. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\elastic__k8s-integration-infra\infra\charts\elastic-agent` — elastic/k8s-integration-infra

```
Error: no repository definition for @stable. Please add them via 'helm repo add'
```

### `D:\helm_clones_github\ishtiaqhimel__oms\charts\oms-server` — ishtiaqhimel/oms

```
Error: no repository definition for https://charts.konghq.com. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\kalavai-net__helm-charts\deployments\monitoring` — kalavai-net/helm-charts

```
Error: no repository definition for https://grafana.github.io/helm-charts, https://ckotzbauer.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\lucidworks__ocp-fusion-helm-charts\5.3.4\fusion\charts\admin-ui` — lucidworks/ocp-fusion-helm-charts

```
Getting updates for unmanaged Helm repositories...
...Successfully got an update from the "https://charts.lucidworks.com/" chart repository
Error: can't get a valid version for 1 subchart(s): "fusion-common-utils" (repository "https://charts.lucidworks.com/", version "1.5.1"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml
```

### `D:\helm_clones_github\rancher__ob-team-charts\charts\prometheus-federator\0.0.1` — rancher/ob-team-charts

```
Error: dependency "helmProjectOperator" has an invalid version/constraint format: improper constraint: ""
```

### `D:\helm_clones_github\shelleg__ac-k8s\helm\ant-umbrella` — shelleg/ac-k8s

```
Error: no repository definition for @incubator, @ac-charts, @ac-charts, @ac-charts, @ac-charts. Please add them via 'helm repo add'
```

### `D:\helm_clones_github\shini4i__charts\charts\app` — shini4i/charts

```
Error: no repository definition for https://shini4i.github.io/charts/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\tetratelabs__charts\charts\demos\istio-monitoring-demo` — tetratelabs/charts

```
Error: no repository definition for https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\vlab-research__fly\devops\vlab` — vlab-research/fly

```
Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update'
```

### `D:\helm_clones_github\yunzck8s__cloudNative\charts\deepflow` — yunzck8s/cloudNative

```
Error: no repository definition for https://grafana.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\FIWARE-Ops__fiware-gitops\aws\token\mongodb` — FIWARE-Ops/fiware-gitops

```
Getting updates for unmanaged Helm repositories...
...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository
Error: can't get a valid version for 1 subchart(s): "mongodb" (repository "https://charts.bitnami.com/bitnami", version "11.0.4"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml
```

### `D:\helm_clones_github\Frndo1203__stack_iceberg_starrocks_k8s\infra\src\helm-charts\airflow` — Frndo1203/stack_iceberg_starrocks_k8s

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\JonasHess__homelab-iac\apps\nextcloud` — JonasHess/homelab-iac

```
Error: no repository definition for https://nextcloud.github.io/helm/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\Kapil-Bhalodiya__E-Commerce\infra\addons\nginx-ingress` — Kapil-Bhalodiya/E-Commerce

```
Error: no repository definition for https://kubernetes.github.io/ingress-nginx. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\Kapil-Bhalodiya__E-commerce-Platform\infra\addons\frontend` — Kapil-Bhalodiya/E-commerce-Platform

```
Error: cannot load values.yaml: cannot unmarshal yaml document: error converting YAML to JSON: yaml: line 8: could not find expected ':'
```

### `D:\helm_clones_github\LuukHors__homelab\products\_base` — LuukHors/homelab

```
Error: validation: chart.metadata.name is required
```

### `D:\helm_clones_github\MrE-Fog__ks-installer2\roles\ks-multicluster\files\kubefed\kubefed` — MrE-Fog/ks-installer2

```
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "https://localhost/" chart repository:
	Get "https://localhost/index.yaml": dial tcp [::1]:443: connectex: Nenhuma ligação pôde ser feita porque o computador de destino
as recusou ativamente.
Error: no cached repository for helm-manager-f2b99ce05b94599549c70dbbe7a891b278e7c3cacad02334fa44682fca36c740 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-f2b99ce05b94599549c70dbbe7a891b278e7c3cacad02334fa44682fca36c740-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### `D:\helm_clones_github\PilotDataPlatform__helm-charts\argo-cd-917` — PilotDataPlatform/helm-charts

```
Error: no repository definition for https://dandydeveloper.github.io/charts/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\SpelmanDevops__retail-store\helm\monitoring` — SpelmanDevops/retail-store

```
Error: no repository definition for https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\TSMC-NYCU-LAB-13__infrastructures\argo\argo-cd` — TSMC-NYCU-LAB-13/infrastructures

```
Error: no repository definition for https://dandydeveloper.github.io/charts/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\VadimShtukan__otus_homework\architect\lesson05\kubernetis\helm-chart` — VadimShtukan/otus_homework

```
Error: no repository definition for https://kubernetes-charts.storage.googleapis.com. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\Vaibhav2goyal__alertmanager\scripts\kube-prometheus-stack` — Vaibhav2goyal/alertmanager

```
Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\WesleyJw__modern-data-stack\infra\src\helm-charts\airflow` — WesleyJw/modern-data-stack

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\Yahya-rabii__soge-plus\helm-charts\kube-prometheus-stack` — Yahya-rabii/soge-plus

```
Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\adstanley__archive\scale-catalog\incubator\archivebox\0.7.2` — adstanley/archive

```
Saving 1 charts
Downloading common from repo oci://tccr.io/truecharts
Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:17.2.30: not found
Error: could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:17.2.30: not found
```

### `D:\helm_clones_github\codefuturist__helm-charts\templates\chart-template` — codefuturist/helm-charts

```
Error: validation: chart.metadata.name is required
```

### `D:\helm_clones_github\dan1dan12345678__Helm_charts\kube-prometheus-stack` — dan1dan12345678/Helm_charts

```
Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\davidlesicnik__homelab-argo\apps\grafana` — davidlesicnik/homelab-argo

```
Error: no repository definition for https://grafana.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\dboeckli__spring-6-icecold-micro-service\helm-charts` — dboeckli/spring-6-icecold-micro-service

```
Error: cannot load Chart.yaml: error converting YAML to JSON: yaml: line 4: found character that cannot start any token
```

### `D:\helm_clones_github\hey101__scale-catalog\incubator\archivebox\0.7.2` — hey101/scale-catalog

```
Saving 1 charts
Downloading common from repo oci://tccr.io/truecharts
Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:17.5.1: not found
Error: could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:17.5.1: not found
```

### `D:\helm_clones_github\hmcts__hmcts-charts\stable\aac-manage-case-assignment` — hmcts/hmcts-charts

```
Getting updates for unmanaged Helm repositories...
...Successfully got an update from the "https://helm.elastic.co" chart repository
Saving 4 charts
Downloading java from repo oci://hmctsprod.azurecr.io/helm
Save error occurred:  could not download oci://hmctsprod.azurecr.io/helm/java: failed to perform "FetchReference" on source: GET "https://hmctsprod.azurecr.io/v2/helm/java/manifests/5.3.0": GET "https://hmctsprod.azurecr.io/oauth2/token?scope=repository%3Ahelm%2Fjava%3Apull&service=hmctsprod.azurecr.io": response status code 401: unauthorized: authentication required, visit https://aka.ms/acr/authorization for more information. CorrelationId: f65eaf6b-2681-42ff-b433-3789c00e163f
Error: could not download oci://hmctsprod.azurecr.io/helm/java: failed to perform "FetchReference" on source: GET "https://hmctsprod.azurecr.io/v2/helm/java/manifests/5.3.0": GET "https://hmctsprod.azurecr.io/oauth2/token?scope=repository%3Ahelm%2Fjava%3Apull&service=hmctsprod.azurecr.io": response status code 401: unauthorized: authentication required, visit https://aka.ms/acr/authorization for more information. CorrelationId: f65eaf6b-2681-42ff-b433-3789c00e163f
```

### `D:\helm_clones_github\legion-platform__legion-infrastructure\helms\monitoring` — legion-platform/legion-infrastructure

```
Error: no repository definition for @stable, @stable. Please add them via 'helm repo add'
```

### `D:\helm_clones_github\merlindorin__charts\charts\pinniped` — merlindorin/charts

```
Error: no repository definition for https://merlindorin.github.io/charts, https://merlindorin.github.io/charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\nwthomas__gitops\helm\longhorn` — nwthomas/gitops

```
Error: no repository definition for https://charts.longhorn.io. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\otus-kuber-2019-12__gidmaster_platform\kubernetes-gitops\deploy\charts\cartservice` — otus-kuber-2019-12/gidmaster_platform

```
Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\rtang03__helm-charts\charts\argocd` — rtang03/helm-charts

```
Error: no repository definition for https://argoproj.github.io/argo-helm. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\tetratelabs__helm-charts\charts\demos\istio-monitoring-demo` — tetratelabs/helm-charts

```
Error: no repository definition for https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\100CallsToEurop__otus\k8s\auth` — 100CallsToEurop/otus

```
Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update'
```

### `D:\helm_clones_github\100rd__platform-design\apps\infra\cilium` — 100rd/platform-design

```
Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update'
```

### `D:\helm_clones_github\AdmanTIC__helm-charts\charts\cremecrm` — AdmanTIC/helm-charts

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\Arthur-B-DevOps__old_helm_charts\charts\Old_charts\charts\incubator\distribution` — Arthur-B-DevOps/old_helm_charts

```
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/, https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\Avichayef__datateam_calculator\helm\jenkins` — Avichayef/datateam_calculator

```
Error: no repository definition for https://charts.jenkins.io. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\Backstage-Epitech__cltest\kustomize\kube-prometheus-stack\charts\kube-prometheus-stack-73.2.2\kube-prometheus-stack` — Backstage-Epitech/cltest

```
Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\Bernardpro__ClusterGCP\kustomize\kube-prometheus-stack\charts\kube-prometheus-stack-73.2.2\kube-prometheus-stack` — Bernardpro/ClusterGCP

```
Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\Bernardpro__ClusterGKE\kustomize\kube-prometheus-stack\charts\kube-prometheus-stack-73.2.2\kube-prometheus-stack` — Bernardpro/ClusterGKE

```
Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\Clark1992__ECK1\src\Integration\ECK1.FailedViewRebuilder\Deploy\service` — Clark1992/ECK1

```
Saving 1 charts
Downloading config-watcher from repo oci://localhost:5000/helm
Save error occurred:  could not download oci://localhost:5000/helm/config-watcher: failed to perform "FetchReference" on source: Get "https://localhost:5000/v2/helm/config-watcher/manifests/0.1.0": dial tcp [::1]:5000: connectex: Nenhuma ligação pôde ser feita porque o computador de destino
as recusou ativamente.
Error: could not download oci://localhost:5000/helm/config-watcher: failed to perform "FetchReference" on source: Get "https://localhost:5000/v2/helm/config-watcher/manifests/0.1.0": dial tcp [::1]:5000: connectex: Nenhuma ligação pôde ser feita porque o computador de destino
as recusou ativamente.
```

