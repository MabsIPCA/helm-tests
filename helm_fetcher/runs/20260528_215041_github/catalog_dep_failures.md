# Helm Dependency Build Failures

Total dependency failures: **194**

| # | Repository | Chart Path | Error |
|---|------------|------------|-------|
| 1 | [bitnami/charts](https://github.com/bitnami/charts) | `D:\helm_clones_github\bitnami__charts\bitnami\apache` | Saving 1 charts Downloading common from repo oci://registry-1.docker.io/bitnamicharts Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/common: failed to perform "FetchReference" on source: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.31.4": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/common: failed to perform "FetchReference" on source: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.31.4": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 2 | [higress-group/higress](https://github.com/higress-group/higress) | `D:\helm_clones_github\higress-group__higress\helm\core` | Error: directory D:\helm_clones_github\higress-group__higress\helm\redis not found |
| 3 | [feast-dev/feast](https://github.com/feast-dev/feast) | `D:\helm_clones_github\feast-dev__feast\infra\charts\feast` | Getting updates for unmanaged Helm repositories... ...Successfully got an update from the "https://feast-helm-charts.storage.googleapis.com" chart repository ...Successfully got an update from the "https://charts.helm.sh/stable" chart repository Error: feature-server chart not found in repo https://feast-helm-charts.storage.googleapis.com |
| 4 | [prometheus-community/helm-charts](https://github.com/prometheus-community/helm-charts) | `D:\helm_clones_github\prometheus-community__helm-charts\charts\prometheus` | Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 5 | [collabnix/kubelabs](https://github.com/collabnix/kubelabs) | `D:\helm_clones_github\collabnix__kubelabs\Helm101\Wordpress` | Error: cannot load Chart.yaml: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go struct field Metadata.deprecated of type bool |
| 6 | [cozystack/cozystack](https://github.com/cozystack/cozystack) | `D:\helm_clones_github\cozystack__cozystack\packages\apps\bucket` | Error: error unpacking subchart cozy-lib in bucket: Chart.yaml file is missing |
| 7 | [grafana/helm-charts](https://github.com/grafana/helm-charts) | `D:\helm_clones_github\grafana__helm-charts\charts\enterprise-logs` | Error: no repository definition for https://grafana.github.io/helm-charts, https://charts.min.io/. Please add the missing repos via 'helm repo add' |
| 8 | [pluralsh/plural](https://github.com/pluralsh/plural) | `D:\helm_clones_github\pluralsh__plural\plural\helm\plural` | Error: no repository definition for https://k8s.ory.sh/helm/charts, https://pluralsh.github.io/module-library. Please add the missing repos via 'helm repo add' |
| 9 | [trueforge-org/truecharts](https://github.com/trueforge-org/truecharts) | `D:\helm_clones_github\trueforge-org__truecharts\charts\stable\kubernetes-dashboard` | Error: no repository definition for https://charts.konghq.com. Please add the missing repos via 'helm repo add' |
| 10 | [slurm-personal/school-dev-k8s](https://github.com/slurm-personal/school-dev-k8s) | `D:\helm_clones_github\slurm-personal__school-dev-k8s\practice\18.templating\wordpress` | Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 11 | [kubeflow/arena](https://github.com/kubeflow/arena) | `D:\helm_clones_github\kubeflow__arena\arena-artifacts` | Error: no repository definition for @tf-operator, @tf-dashbard, @cron-operator, @et-operator, @mpi-operator, @pytorch-operator, @gpu-exporter, @elastic-job-supervisor. Please add them via 'helm repo add' |
| 12 | [fluxninja/aperture](https://github.com/fluxninja/aperture) | `D:\helm_clones_github\fluxninja__aperture\manifests\charts\aperture-agent` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 13 | [WeBankFinTech/Prophecis](https://github.com/WeBankFinTech/Prophecis) | `D:\helm_clones_github\WeBankFinTech__Prophecis\install\Prophecis` | Error: cannot load values.yaml: error reading yaml document: invalid Yaml document separator: --END RSA PRIVATE KEY-----" |
| 14 | [rancher/charts](https://github.com/rancher/charts) | `D:\helm_clones_github\rancher__charts\charts\epinio\102.0.1+up1.6.2` | Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update' |
| 15 | [DataDog/helm-charts](https://github.com/DataDog/helm-charts) | `D:\helm_clones_github\DataDog__helm-charts\charts\datadog` | Error: no repository definition for https://helm.datadoghq.com, https://prometheus-community.github.io/helm-charts, https://helm.datadoghq.com, https://helm.datadoghq.com. Please add the missing repos via 'helm repo add' |
| 16 | [radondb/radondb-mysql-kubernetes](https://github.com/radondb/radondb-mysql-kubernetes) | `D:\helm_clones_github\radondb__radondb-mysql-kubernetes\charts\mysql-operator` | Error: directory D:\helm_clones_github\radondb__radondb-mysql-kubernetes\charts\mysql-operator\charts\mysqlcluster not found |
| 17 | [truenas/charts](https://github.com/truenas/charts) | `D:\helm_clones_github\truenas__charts\charts\collabora\1.2.30` | Error: directory D:\helm_clones_github\truenas__charts\common\2304.0.1 not found |
| 18 | [instantlinux/docker-tools](https://github.com/instantlinux/docker-tools) | `D:\helm_clones_github\instantlinux__docker-tools\images\git-pull\helm` | Error: directory D:\helm_clones_github\instantlinux__docker-tools\images\git-pull\chartlib not found |
| 19 | [IBM/charts](https://github.com/IBM/charts) | `D:\helm_clones_github\IBM__charts\community\artifactory-ha` | level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..." Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add' |
| 20 | [mercedes-benz/DnA](https://github.com/mercedes-benz/DnA) | `D:\helm_clones_github\mercedes-benz__DnA\deployment\kubernetes\mysql-helm-chart` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 21 | [Hydrospheredata/hydro-serving](https://github.com/Hydrospheredata/hydro-serving) | `D:\helm_clones_github\Hydrospheredata__hydro-serving\helm` | Error: cannot load Chart.yaml: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type v2.Metadata |
| 22 | [JahstreetOrg/spark-on-kubernetes-helm](https://github.com/JahstreetOrg/spark-on-kubernetes-helm) | `D:\helm_clones_github\JahstreetOrg__spark-on-kubernetes-helm\charts\cluster-base` | level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..." Error: no repository definition for https://kubernetes.github.io/ingress-nginx, https://charts.jetstack.io, https://charts.helm.sh/stable, https://charts.helm.sh/stable. Please add the missing repos via 'helm repo add' |
| 23 | [osm-seed/osm-seed](https://github.com/osm-seed/osm-seed) | `D:\helm_clones_github\osm-seed__osm-seed\osm-seed` | Error: validation: chart.metadata.version is required |
| 24 | [quanxiang-cloud/quanxiang](https://github.com/quanxiang-cloud/quanxiang) | `D:\helm_clones_github\quanxiang-cloud__quanxiang\deployment\charts\elasticsearch` | Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 25 | [CenterForOpenScience/helm-charts](https://github.com/CenterForOpenScience/helm-charts) | `D:\helm_clones_github\CenterForOpenScience__helm-charts\elastic-stack` | level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..." Error: no repository definition for https://centerforopenscience.github.io/helm-charts/, https://centerforopenscience.github.io/helm-charts/. Please add the missing repos via 'helm repo add' |
| 26 | [companyinfo/helm-charts](https://github.com/companyinfo/helm-charts) | `D:\helm_clones_github\companyinfo__helm-charts\charts\helmet\examples\simple` | Error: no repository definition for https://charts.companyinfo.dev. Please add the missing repos via 'helm repo add' |
| 27 | [rancher/rke2-charts](https://github.com/rancher/rke2-charts) | `D:\helm_clones_github\rancher__rke2-charts\packages\rke2-cilium-legacy\charts` | Error: directory D:\helm_clones_github\rancher__rke2-charts\packages\rke2-cilium-legacy\charts\charts\cilium not found |
| 28 | [Obmondo/KubeAid](https://github.com/Obmondo/KubeAid) | `D:\helm_clones_github\Obmondo__KubeAid\argocd-helm-charts\argo-cd` | Error: error unpacking subchart kubeaid-addons in argo-cd: Chart.yaml file is missing |
| 29 | [rancher/partner-charts](https://github.com/rancher/partner-charts) | `D:\helm_clones_github\rancher__partner-charts\charts\amd\amd-gpu\0.10.0` | Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update' |
| 30 | [lianqingsec/NucleiPocGather](https://github.com/lianqingsec/NucleiPocGather) | `D:\helm_clones_github\lianqingsec__NucleiPocGather\poc\other` | Error: chart file "signatures_1.yaml" is larger than the maximum file size 5242880 |
| 31 | [elastisys/compliantkubernetes-apps](https://github.com/elastisys/compliantkubernetes-apps) | `D:\helm_clones_github\elastisys__compliantkubernetes-apps\helmfile.d\upstream\falcosecurity\falco` | Error: no repository definition for https://falcosecurity.github.io/charts, https://falcosecurity.github.io/charts, https://falcosecurity.github.io/charts. Please add the missing repos via 'helm repo add' |
| 32 | [wiremind/wiremind-helm-charts](https://github.com/wiremind/wiremind-helm-charts) | `D:\helm_clones_github\wiremind__wiremind-helm-charts\charts\druid` | Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://machine424.github.io/kube-hpa-scale-to-zero, https://wiremind.github.io/wiremind-helm-charts, https://wiremind.github.io/wiremind-helm-charts. Please add the missing repos via 'helm repo add' |
| 33 | [bytle/kubee](https://github.com/bytle/kubee) | `D:\helm_clones_github\bytle__kubee` | Error: error unpacking subchart README.md in Kubee: Chart.yaml file is missing |
| 34 | [tmforum-oda/oda-canvas](https://github.com/tmforum-oda/oda-canvas) | `D:\helm_clones_github\tmforum-oda__oda-canvas\charts\apisix-gateway` | Error: no repository definition for https://charts.apiseven.com. Please add the missing repos via 'helm repo add' |
| 35 | [IndustryFusion/DigitalTwin](https://github.com/IndustryFusion/DigitalTwin) | `D:\helm_clones_github\IndustryFusion__DigitalTwin\helm\charts\velero` | Error: directory D:\helm_clones_github\IndustryFusion__DigitalTwin\helm\airgap-deployment\helm-charts\charts\velero not found |
| 36 | [boozallen/aissemble](https://github.com/boozallen/aissemble) | `D:\helm_clones_github\boozallen__aissemble\extensions\extensions-helm\aissemble-airflow-chart` | Error: no repository definition for https://airflow.apache.org/. Please add the missing repos via 'helm repo add' |
| 37 | [tkestack/charts](https://github.com/tkestack/charts) | `D:\helm_clones_github\tkestack__charts\incubator\airflow` | level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..." Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/, https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add' |
| 38 | [sapcc/helm-charts](https://github.com/sapcc/helm-charts) | `D:\helm_clones_github\sapcc__helm-charts\common\inventory-updater` | Saving 1 charts Downloading owner-info from repo oci://keppel.eu-de-1.cloud.sap/ccloud-helm Save error occurred:  could not download oci://keppel.eu-de-1.cloud.sap/ccloud-helm/owner-info: failed to perform "FetchReference" on source: Get "https://keppel.eu-de-1.cloud.sap/v2/ccloud-helm/owner-info/manifests/0.2.0": dial tcp: lookup keppel.eu-de-1.cloud.sap: getaddrinfow: Este é geralmente um erro temporário durante a resolução de nomes de anfitrião e significa que o servidor local não recebeu uma resposta de um servidor autoritário. Error: could not download oci://keppel.eu-de-1.cloud.sap/ccloud-helm/owner-info: failed to perform "FetchReference" on source: Get "https://keppel.eu-de-1.cloud.sap/v2/ccloud-helm/owner-info/manifests/0.2.0": dial tcp: lookup keppel.eu-de-1.cloud.sap: getaddrinfow: Este é geralmente um erro temporário durante a resolução de nomes de anfitrião e significa que o servidor local não recebeu uma resposta de um servidor autoritário. |
| 39 | [k8s-home-lab/helm-charts](https://github.com/k8s-home-lab/helm-charts) | `D:\helm_clones_github\k8s-home-lab__helm-charts\charts\stable\bazarr` | Error: no repository definition for https://k8s-home-lab.github.io/helm-charts/. Please add the missing repos via 'helm repo add' |
| 40 | [neuroxhq/helm-chart-neurox-control](https://github.com/neuroxhq/helm-chart-neurox-control) | `D:\helm_clones_github\neuroxhq__helm-chart-neurox-control` | Saving 6 charts Downloading neurox-control-api from repo oci://ghcr.io/neuroxhq/helm-charts Save error occurred:  could not download oci://ghcr.io/neuroxhq/helm-charts/neurox-control-api: failed to perform "FetchReference" on source: GET "https://ghcr.io/v2/neuroxhq/helm-charts/neurox-control-api/manifests/2.233.1": GET "https://ghcr.io/token?scope=repository%3Aneuroxhq%2Fhelm-charts%2Fneurox-control-api%3Apull&service=ghcr.io": response status code 401: unauthorized: authentication required Error: could not download oci://ghcr.io/neuroxhq/helm-charts/neurox-control-api: failed to perform "FetchReference" on source: GET "https://ghcr.io/v2/neuroxhq/helm-charts/neurox-control-api/manifests/2.233.1": GET "https://ghcr.io/token?scope=repository%3Aneuroxhq%2Fhelm-charts%2Fneurox-control-api%3Apull&service=ghcr.io": response status code 401: unauthorized: authentication required |
| 41 | [pluralsh/console](https://github.com/pluralsh/console) | `D:\helm_clones_github\pluralsh__console\charts\console` | Error: no repository definition for https://fluxcd-community.github.io/helm-charts, https://charts.dexidp.io, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 42 | [YAKEcloud/yake](https://github.com/YAKEcloud/yake) | `D:\helm_clones_github\YAKEcloud__yake\helmcharts\acl` | Error: dependency "controller" has an invalid version/constraint format: improper constraint: "" |
| 43 | [GreptimeTeam/helm-charts](https://github.com/GreptimeTeam/helm-charts) | `D:\helm_clones_github\GreptimeTeam__helm-charts\charts\greptimedb-cluster` | Error: no repository definition for https://grafana.github.io/helm-charts, https://raw.githubusercontent.com/hansehe/jaeger-all-in-one/master/helm/charts, https://greptimeteam.github.io/helm-charts, https://greptimeteam.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 44 | [k0rdent/catalog](https://github.com/k0rdent/catalog) | `D:\helm_clones_github\k0rdent__catalog\apps\alloy\charts\alloy-1.6.1` | Error: no repository definition for https://grafana.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 45 | [junghoon2/k8s-class](https://github.com/junghoon2/k8s-class) | `D:\helm_clones_github\junghoon2__k8s-class\argo-cd\argo-cd-5.14.1` | Error: no repository definition for https://dandydeveloper.github.io/charts/. Please add the missing repos via 'helm repo add' |
| 46 | [silogen/cluster-forge](https://github.com/silogen/cluster-forge) | `D:\helm_clones_github\silogen__cluster-forge\sources\amd-gpu-operator\v1.3.1` | Error: no repository definition for https://kubernetes-sigs.github.io/node-feature-discovery/charts. Please add the missing repos via 'helm repo add' |
| 47 | [ai-solution-eng/frameworks](https://github.com/ai-solution-eng/frameworks) | `D:\helm_clones_github\ai-solution-eng__frameworks\appsmith\3.6.4` | Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 48 | [wenerme/kube-stub-cluster](https://github.com/wenerme/kube-stub-cluster) | `D:\helm_clones_github\wenerme__kube-stub-cluster\keycloak` | Saving 1 charts Downloading keycloak from repo oci://dockercr.wener.me/bitnamicharts Save error occurred:  could not download oci://dockercr.wener.me/bitnamicharts/keycloak: failed to perform "FetchReference" on source: GET "https://dockercr.wener.me/v2/bitnamicharts/keycloak/manifests/15.1.4": invalid response `Content-Type` header; mime: no media type Error: could not download oci://dockercr.wener.me/bitnamicharts/keycloak: failed to perform "FetchReference" on source: GET "https://dockercr.wener.me/v2/bitnamicharts/keycloak/manifests/15.1.4": invalid response `Content-Type` header; mime: no media type |
| 49 | [ThienAnTrinh/product-search](https://github.com/ThienAnTrinh/product-search) | `D:\helm_clones_github\ThienAnTrinh__product-search\monitoring\logs-metrics\helm-charts\charts\kube-prometheus-stack` | Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 50 | [nebius/nebius-k8s-applications](https://github.com/nebius/nebius-k8s-applications) | `D:\helm_clones_github\nebius__nebius-k8s-applications` | Getting updates for unmanaged Helm repositories... ...Unable to get an update from the "https://volcano-sh.github.io/charts/" chart repository: 	failed to fetch https://volcano-sh.github.io/charts/index.yaml : 404 Not Found ...Successfully got an update from the "https://mellanox.github.io/network-operator" chart repository ...Successfully got an update from the "https://ray-project.github.io/kuberay-helm/" chart repository ...Successfully got an update from the "https://qdrant.github.io/qdrant-helm" chart repository ...Successfully got an update from the "https://airflow.apache.org" chart repository ...Successfully got an update from the "https://otwld.github.io/ollama-helm/" chart repository ...Successfully got an update from the "https://zilliztech.github.io/milvus-helm/" chart repository ...Successfully got an update from the "https://cowboysysop.github.io/charts/" chart repository ...Successfully got an update from the "https://argoproj.github.io/argo-helm" chart repository ...Successfully got an update from the "https://hub.jupyter.org/helm-chart" chart repository ...Unable to get an update from the "https://github.com/weaviate/weaviate" chart repository: 	failed to fetch https://github.com/weaviate/weaviate/index.yaml : 404 Not Found ...Successfully got an update from the "https://grafana.github.io/helm-charts" chart repository ...Successfully got an update from the "https://helm.ngc.nvidia.com/nvidia" chart repository Error: directory D:\helm_clones_github\nebius__nebius-k8s-applications\charts\clearml-agent not found |
| 51 | [suse-edge/charts](https://github.com/suse-edge/charts) | `D:\helm_clones_github\suse-edge__charts\charts\kubevirt\0.1.0` | Error: dependency "cdi" has an invalid version/constraint format: improper constraint: "" |
| 52 | [wx-chevalier/k8s-examples](https://github.com/wx-chevalier/k8s-examples) | `D:\helm_clones_github\wx-chevalier__k8s-examples\helm-charts\backend-spring-app` | Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update' |
| 53 | [cnrancher/pandaria-catalog](https://github.com/cnrancher/pandaria-catalog) | `D:\helm_clones_github\cnrancher__pandaria-catalog\charts\rancher-hami\107.0.0+up2.5.2\charts\hami-webui` | Error: no repository definition for https://nvidia.github.io/dcgm-exporter/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 54 | [DaoCloud/dce-charts-repackage](https://github.com/DaoCloud/dce-charts-repackage) | `D:\helm_clones_github\DaoCloud__dce-charts-repackage\charts\argo-cd\argo-cd\charts\argo-cd` | Error: no repository definition for https://dandydeveloper.github.io/charts/. Please add the missing repos via 'helm repo add' |
| 55 | [sebolabs/eks-tf-gitops](https://github.com/sebolabs/eks-tf-gitops) | `D:\helm_clones_github\sebolabs__eks-tf-gitops\k8s\add-ons\csi-secrets-store-provider-aws` | Error: cannot load Chart.yaml: error converting YAML to JSON: yaml: line 3: mapping values are not allowed in this context |
| 56 | [StatCan/charts](https://github.com/StatCan/charts) | `D:\helm_clones_github\StatCan__charts\deprecated\cost-analyzer` | Error: no repository definition for https://kubecost.github.io/cost-analyzer/. Please add the missing repos via 'helm repo add' |
| 57 | [jordanopensource/charts](https://github.com/jordanopensource/charts) | `D:\helm_clones_github\jordanopensource__charts\charts\etherpad` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 58 | [jugatsu/microservices](https://github.com/jugatsu/microservices) | `D:\helm_clones_github\jugatsu__microservices\deploy\kubernetes\Charts\gitlab-omnibus` | level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..." Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies |
| 59 | [llajas/homelab](https://github.com/llajas/homelab) | `D:\helm_clones_github\llajas__homelab\apps\clusterplex` | Error: no repository definition for https://bjw-s-labs.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 60 | [replicatedhq/chartsmith](https://github.com/replicatedhq/chartsmith) | `D:\helm_clones_github\replicatedhq__chartsmith\chart\chartsmith` | Saving 1 charts Downloading replicated from repo oci://registry.chartsmith.ai/library Save error occurred:  could not download oci://registry.chartsmith.ai/library/replicated: failed to perform "FetchReference" on source: Get "https://registry.chartsmith.ai/v2/library/replicated/manifests/1.12.1": dial tcp: lookup registry.chartsmith.ai: no such host Error: could not download oci://registry.chartsmith.ai/library/replicated: failed to perform "FetchReference" on source: Get "https://registry.chartsmith.ai/v2/library/replicated/manifests/1.12.1": dial tcp: lookup registry.chartsmith.ai: no such host |
| 61 | [commercialhaskell/all-cabal-metadata](https://github.com/commercialhaskell/all-cabal-metadata) | `D:\helm_clones_github\commercialhaskell__all-cabal-metadata\packages\ch` | Error: validation: chart.metadata.name is required |
| 62 | [NeonGeckoCom/neon-diana-utils](https://github.com/NeonGeckoCom/neon-diana-utils) | `D:\helm_clones_github\NeonGeckoCom__neon-diana-utils\neon_diana_utils\helm_charts\http\libretranslate` | Error: can't get a valid version for 1 subchart(s): "base-http" (repository "file://../../base/base-http", version "0.0.6"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml |
| 63 | [easy2stake/thegraph](https://github.com/easy2stake/thegraph) | `D:\helm_clones_github\easy2stake__thegraph\graphprotocol` | Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://grafana.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 64 | [thoughtworks/byor-voting-infrastructure](https://github.com/thoughtworks/byor-voting-infrastructure) | `D:\helm_clones_github\thoughtworks__byor-voting-infrastructure\src\byor-voting-chart` | level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..." Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies |
| 65 | [TWilkin/powerpi](https://github.com/TWilkin/powerpi) | `D:\helm_clones_github\TWilkin__powerpi\kubernetes` | Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update' |
| 66 | [safesoftware/helm-charts](https://github.com/safesoftware/helm-charts) | `D:\helm_clones_github\safesoftware__helm-charts\chart-source\fmeserver-2018.1.1` | level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..." Error: no repository definition for https://kubernetes-charts.storage.googleapis.com. Please add the missing repos via 'helm repo add' |
| 67 | [anza-labs/charts](https://github.com/anza-labs/charts) | `D:\helm_clones_github\anza-labs__charts\deprecated\hosted-control-plane` | Error: no repository definition for https://charts.jetstack.io. Please add the missing repos via 'helm repo add' |
| 68 | [appscode-cloud/ui-wizards](https://github.com/appscode-cloud/ui-wizards) | `D:\helm_clones_github\appscode-cloud__ui-wizards\charts\kubedbcom-elasticsearch-editor` | Error: chart file "values.openapiv3_schema.yaml" is larger than the maximum file size 5242880 |
| 69 | [joostvdg/cmg](https://github.com/joostvdg/cmg) | `D:\helm_clones_github\joostvdg__cmg\charts\preview` | Getting updates for unmanaged Helm repositories... ...Unable to get an update from the "http://chartmuseum.jenkins-x.io" chart repository: 	failed to fetch http://chartmuseum.jenkins-x.io/index.yaml : 404 Not Found Error: no cached repository for helm-manager-d8c90e93d45e753db0fa4ec457d893e08cd1e0827b77896ee636b714706cbbb2 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-d8c90e93d45e753db0fa4ec457d893e08cd1e0827b77896ee636b714706cbbb2-index.yaml: O sistema não conseguiu localizar o ficheiro especificado. |
| 70 | [salmanmkc/agentverse](https://github.com/salmanmkc/agentverse) | `D:\helm_clones_github\salmanmkc__agentverse\mcp-backend\charts\rag-stack` | Error: no repository definition for https://helm.neo4j.com/neo4j, https://helm.neo4j.com/neo4j, https://zilliztech.github.io/milvus-helm/. Please add the missing repos via 'helm repo add' |
| 71 | [samgabrail/env0-argocd](https://github.com/samgabrail/env0-argocd) | `D:\helm_clones_github\samgabrail__env0-argocd\schoolapp-subchart` | Error: no repository definition for https://gitlab.com/api/v4/projects/34240616/packages/helm/stable, https://gitlab.com/api/v4/projects/34240616/packages/helm/stable, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 72 | [ai-on-gke/common-infra](https://github.com/ai-on-gke/common-infra) | `D:\helm_clones_github\ai-on-gke__common-infra\common\charts\gmp-engine` | Error: dependency "gmp-frontend" has an invalid version/constraint format: improper constraint: "" |
| 73 | [ibuildthecloud/rancher-charts](https://github.com/ibuildthecloud/rancher-charts) | `D:\helm_clones_github\ibuildthecloud__rancher-charts\charts\anchore-engine\0.1.0` | level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..." Error: no repository definition for https://kubernetes-charts.storage.googleapis.com. Please add the missing repos via 'helm repo add' |
| 74 | [wunderio/charts](https://github.com/wunderio/charts) | `D:\helm_clones_github\wunderio__charts\drupal` | Error: no repository definition for https://storage.googleapis.com/charts.wdr.io, https://percona.github.io/percona-helm-charts/, https://storage.googleapis.com/charts.wdr.io, https://storage.googleapis.com/charts.wdr.io, https://storage.googleapis.com/charts.wdr.io, https://jouve.github.io/charts/. Please add the missing repos via 'helm repo add' |
| 75 | [SocialGouv/no-package-malware](https://github.com/SocialGouv/no-package-malware) | `D:\helm_clones_github\SocialGouv__no-package-malware\charts\no-package-malware` | Error: no repository definition for https://groundhog2k.github.io/helm-charts/, https://cloudpirates-io.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 76 | [StackVista/helm-charts](https://github.com/StackVista/helm-charts) | `D:\helm_clones_github\StackVista__helm-charts\stable\otel-demo` | Error: no repository definition for https://open-telemetry.github.io/opentelemetry-helm-charts. Please add the missing repos via 'helm repo add' |
| 77 | [anup1384/helm-charts](https://github.com/anup1384/helm-charts) | `D:\helm_clones_github\anup1384__helm-charts\stable\kafka` | level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..." Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies |
| 78 | [hifer/devops](https://github.com/hifer/devops) | `D:\helm_clones_github\hifer__devops\monitor\prometheus-operator` | level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..." Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/, https://kubernetes-charts.storage.googleapis.com/, https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add' |
| 79 | [kubebb/components](https://github.com/kubebb/components) | `D:\helm_clones_github\kubebb__components\charts\cluster-component\charts\openebs` | Error: no repository definition for https://openebs.github.io/node-disk-manager, https://openebs.github.io/dynamic-localpv-provisioner, https://openebs.github.io/cstor-operators, https://openebs.github.io/jiva-operator, https://openebs.github.io/zfs-localpv, https://openebs.github.io/lvm-localpv, https://openebs.github.io/dynamic-nfs-provisioner. Please add the missing repos via 'helm repo add' |
| 80 | [marcosviniciusi/k3s-homelab](https://github.com/marcosviniciusi/k3s-homelab) | `D:\helm_clones_github\marcosviniciusi__k3s-homelab\kustomize\infisical-stack\infisical-server` | Error: no repository definition for https://kubernetes.github.io/ingress-nginx, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 81 | [AISphere/ffdl-trainer](https://github.com/AISphere/ffdl-trainer) | `D:\helm_clones_github\AISphere__ffdl-trainer` | Error: directory D:\helm_clones_github\ffdl-lcm not found |
| 82 | [Makhuta/truecharts-archive-scale-catalog](https://github.com/Makhuta/truecharts-archive-scale-catalog) | `D:\helm_clones_github\Makhuta__truecharts-archive-scale-catalog\incubator\archivebox\0.7.2` | Saving 1 charts Downloading common from repo oci://tccr.io/truecharts Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:17.2.30: not found Error: could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:17.2.30: not found |
| 83 | [MegaWiz-Dev-Team/Asgard](https://github.com/MegaWiz-Dev-Team/Asgard) | `D:\helm_clones_github\MegaWiz-Dev-Team__Asgard\charts\asgard` | Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update' |
| 84 | [Obsidian-Owl/floe](https://github.com/Obsidian-Owl/floe) | `D:\helm_clones_github\Obsidian-Owl__floe\charts\cognee-platform` | Error: no repository definition for https://helm.neo4j.com/neo4j, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 85 | [Yunjuzhen/charts](https://github.com/Yunjuzhen/charts) | `D:\helm_clones_github\Yunjuzhen__charts\stable\anchore-engine\0.1.3\anchore-engine` | level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..." Error: no repository definition for https://kubernetes-charts.storage.googleapis.com. Please add the missing repos via 'helm repo add' |
| 86 | [joshleecreates/clickhouse-opentelemetry-iac](https://github.com/joshleecreates/clickhouse-opentelemetry-iac) | `D:\helm_clones_github\joshleecreates__clickhouse-opentelemetry-iac\argo-apps\qryn` | Getting updates for unmanaged Helm repositories... ...Unable to get an update from the "https://metrico.github.io/qryn-helm/" chart repository: 	failed to fetch https://metrico.github.io/qryn-helm/index.yaml : 404 Not Found Error: no cached repository for helm-manager-fed7dc84065c3ccd251fae7cd72350b43bda8f8cb634374dd8b3a9318ac9d4e3 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-fed7dc84065c3ccd251fae7cd72350b43bda8f8cb634374dd8b3a9318ac9d4e3-index.yaml: O sistema não conseguiu localizar o ficheiro especificado. |
| 87 | [pluralsh/plural-helm-charts](https://github.com/pluralsh/plural-helm-charts) | `D:\helm_clones_github\pluralsh__plural-helm-charts\charts\airbyte` | Error: no repository definition for https://charts.bitnami.com/bitnami, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts. Please add the missing repos via 'helm repo add' |
| 88 | [rancher/ob-team-charts](https://github.com/rancher/ob-team-charts) | `D:\helm_clones_github\rancher__ob-team-charts\charts\prometheus-federator\0.0.1` | Error: dependency "helmProjectOperator" has an invalid version/constraint format: improper constraint: "" |
| 89 | [Andrew-Su-0718/zelos-image](https://github.com/Andrew-Su-0718/zelos-image) | `D:\helm_clones_github\Andrew-Su-0718__zelos-image\mmdet\image\charts\arena-artifacts` | Error: no repository definition for @tf-operator, @tf-dashbard, @cron-operator, @et-operator, @mpi-operator, @pytorch-operator, @gpu-exporter, @elastic-job-supervisor. Please add them via 'helm repo add' |
| 90 | [KevMCarp/truecharts-catalog-fork](https://github.com/KevMCarp/truecharts-catalog-fork) | `D:\helm_clones_github\KevMCarp__truecharts-catalog-fork\dependency\clickhouse\5.0.54` | Getting updates for unmanaged Helm repositories... ...Unable to get an update from the "https://library-charts.truecharts.org" chart repository: 	Get "https://library-charts.truecharts.org/index.yaml": dial tcp: lookup library-charts.truecharts.org: no such host Error: no cached repository for helm-manager-024b189b59f6c6ccf0de6e5148db1578caf551c511f4eb220ece14cef00f80e0 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-024b189b59f6c6ccf0de6e5148db1578caf551c511f4eb220ece14cef00f80e0-index.yaml: O sistema não conseguiu localizar o ficheiro especificado. |
| 91 | [NeuraLegion/helmcharts](https://github.com/NeuraLegion/helmcharts) | `D:\helm_clones_github\NeuraLegion__helmcharts\charts\altoroj` | Saving 1 charts Save error occurred:  can't get a valid version for dependency simple-service Error: can't get a valid version for dependency simple-service |
| 92 | [Sureya/airflow_k8s_executor](https://github.com/Sureya/airflow_k8s_executor) | `D:\helm_clones_github\Sureya__airflow_k8s_executor\helm_charts\official\charts\incubator\distribution` | level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..." Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/, https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add' |
| 93 | [ethpandaops/whisk-devnets](https://github.com/ethpandaops/whisk-devnets) | `D:\helm_clones_github\ethpandaops__whisk-devnets\kubernetes\devnet-0\blobscan` | Error: no repository definition for https://charts.bitnami.com/bitnami, https://ethpandaops.github.io/ethereum-helm-charts, https://ethpandaops.github.io/ethereum-helm-charts. Please add the missing repos via 'helm repo add' |
| 94 | [galserg/kubetest](https://github.com/galserg/kubetest) | `D:\helm_clones_github\galserg__kubetest\helmfile\charts\etcd` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 95 | [ishtiaqhimel/oms](https://github.com/ishtiaqhimel/oms) | `D:\helm_clones_github\ishtiaqhimel__oms\charts\oms-server` | Error: no repository definition for https://charts.konghq.com. Please add the missing repos via 'helm repo add' |
| 96 | [lucidworks/ocp-fusion-helm-charts](https://github.com/lucidworks/ocp-fusion-helm-charts) | `D:\helm_clones_github\lucidworks__ocp-fusion-helm-charts\5.3.4\fusion` | Error: no repository definition for https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com. Please add the missing repos via 'helm repo add' |
| 97 | [openshift-helm-charts/development](https://github.com/openshift-helm-charts/development) | `D:\helm_clones_github\openshift-helm-charts__development\charts\partners\embedded\tokenvisor\0.1.0\src` | Error: no repository definition for https://victoriametrics.github.io/helm-charts/. Please add the missing repos via 'helm repo add' |
| 98 | [openshift-helm-charts/sandbox](https://github.com/openshift-helm-charts/sandbox) | `D:\helm_clones_github\openshift-helm-charts__sandbox\charts\redhat\redhat\redhat-mysql-persistent\0.0.2\src` | Getting updates for unmanaged Helm repositories... ...Unable to get an update from the "https://github.com/openshift-helm-charts/charts" chart repository: 	failed to fetch https://github.com/openshift-helm-charts/charts/index.yaml : 404 Not Found Error: no cached repository for helm-manager-e8a50bb8edab0b7411123ba1d6ccffca210a3cb9ffb119170c254ad301b6b826 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-e8a50bb8edab0b7411123ba1d6ccffca210a3cb9ffb119170c254ad301b6b826-index.yaml: O sistema não conseguiu localizar o ficheiro especificado. |
| 99 | [oracle-cne/catalog](https://github.com/oracle-cne/catalog) | `D:\helm_clones_github\oracle-cne__catalog\charts\metallb-0.15.2` | Error: no repository definition for https://metallb.github.io/frr-k8s. Please add the missing repos via 'helm repo add' |
| 100 | [otterscale/charts](https://github.com/otterscale/charts) | `D:\helm_clones_github\otterscale__charts\charts\otterscale` | Error: no repository definition for https://codecentric.github.io/helm-charts, https://valkey.io/valkey-helm/, https://helm.goharbor.io. Please add the missing repos via 'helm repo add' |
| 101 | [platform-mesh/helm-charts](https://github.com/platform-mesh/helm-charts) | `D:\helm_clones_github\platform-mesh__helm-charts\charts\common\test-chart` | Saving 1 charts Save error occurred:  can't get a valid version for dependency common Error: can't get a valid version for dependency common |
| 102 | [sdelrio/homelab-k3s](https://github.com/sdelrio/homelab-k3s) | `D:\helm_clones_github\sdelrio__homelab-k3s\system\internal\postgres-operator` | Getting updates for unmanaged Helm repositories... ...Successfully got an update from the "https://opensource.zalando.com/postgres-operator/charts/postgres-operator" chart repository Error: can't get a valid version for 1 subchart(s): "postgres-operator" (repository "https://opensource.zalando.com/postgres-operator/charts/postgres-operator", version "1.10.0"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml |
| 103 | [tetratelabs/charts](https://github.com/tetratelabs/charts) | `D:\helm_clones_github\tetratelabs__charts\charts\demos\istio-monitoring-demo` | Error: no repository definition for https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 104 | [thomaswyrick/xdm-helm-chart](https://github.com/thomaswyrick/xdm-helm-chart) | `D:\helm_clones_github\thomaswyrick__xdm-helm-chart` | Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update' |
| 105 | [vineethvijay/prox-k8s-lab](https://github.com/vineethvijay/prox-k8s-lab) | `D:\helm_clones_github\vineethvijay__prox-k8s-lab\helm\charts\nfs-hdd-provisioner` | Error: no repository definition for https://kubernetes-sigs.github.io/nfs-subdir-external-provisioner/. Please add the missing repos via 'helm repo add' |
| 106 | [2694484453/helm-repo](https://github.com/2694484453/helm-repo) | `D:\helm_clones_github\2694484453__helm-repo\kube-prometheus-stack` | Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 107 | [Chinaxiaoming666/kube-prometheus-stack](https://github.com/Chinaxiaoming666/kube-prometheus-stack) | `D:\helm_clones_github\Chinaxiaoming666__kube-prometheus-stack` | Error: no repository definition for https://charts.helm.sh/stable, https://prometheus-community.github.io/helm-charts, https://grafana.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 108 | [DenisMarta/prodject1](https://github.com/DenisMarta/prodject1) | `D:\helm_clones_github\DenisMarta__prodject1\deployment\kubernetes\mysql-helm-chart` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 109 | [HonourHealth/JavaSpringbootRESTMicroservices](https://github.com/HonourHealth/JavaSpringbootRESTMicroservices) | `D:\helm_clones_github\HonourHealth__JavaSpringbootRESTMicroservices\springboot-microservices-springcloud-docker-kubernetes-helm\helm\bank-services\accounts` | Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update' |
| 110 | [Kapil-Bhalodiya/E-Commerce](https://github.com/Kapil-Bhalodiya/E-Commerce) | `D:\helm_clones_github\Kapil-Bhalodiya__E-Commerce\infra\addons\alloy` | Error: no repository definition for https://grafana.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 111 | [Kapil-Bhalodiya/E-commerce-Platform](https://github.com/Kapil-Bhalodiya/E-commerce-Platform) | `D:\helm_clones_github\Kapil-Bhalodiya__E-commerce-Platform\infra\addons\cert-manager` | Error: no repository definition for https://charts.jetstack.io. Please add the missing repos via 'helm repo add' |
| 112 | [PRO-Robotech/helmfile-manifests](https://github.com/PRO-Robotech/helmfile-manifests) | `D:\helm_clones_github\PRO-Robotech__helmfile-manifests\charts\argoproj\argo-cd-8.0.9\argo-cd` | Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update' |
| 113 | [PilotDataPlatform/helm-charts](https://github.com/PilotDataPlatform/helm-charts) | `D:\helm_clones_github\PilotDataPlatform__helm-charts\argo-cd-917` | Error: no repository definition for https://dandydeveloper.github.io/charts/. Please add the missing repos via 'helm repo add' |
| 114 | [SmartX-Team/smartx-k8s](https://github.com/SmartX-Team/smartx-k8s) | `D:\helm_clones_github\SmartX-Team__smartx-k8s\apps\exdns-2` | Error: cannot load Chart.yaml: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type v2.Metadata |
| 115 | [TSMC-NYCU-LAB-13/infrastructures](https://github.com/TSMC-NYCU-LAB-13/infrastructures) | `D:\helm_clones_github\TSMC-NYCU-LAB-13__infrastructures\argo\argo-cd` | Error: no repository definition for https://dandydeveloper.github.io/charts/. Please add the missing repos via 'helm repo add' |
| 116 | [adstanley/archive](https://github.com/adstanley/archive) | `D:\helm_clones_github\adstanley__archive\scale-catalog\incubator\archivebox\0.7.2` | Saving 1 charts Downloading common from repo oci://tccr.io/truecharts Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:17.2.30: not found Error: could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:17.2.30: not found |
| 117 | [azaurus1/homelab](https://github.com/azaurus1/homelab) | `D:\helm_clones_github\azaurus1__homelab\apps\api-coverage-server` | Getting updates for unmanaged Helm repositories... ...Unable to get an update from the "https://chartmuseum.azaurus.dev/" chart repository: 	failed to fetch https://chartmuseum.azaurus.dev/index.yaml : 530 <none> Error: no cached repository for helm-manager-23f58962e8ae66ac9eadc10ae209dc5cab99e97e33882fa9df1e156284414aea found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-23f58962e8ae66ac9eadc10ae209dc5cab99e97e33882fa9df1e156284414aea-index.yaml: O sistema não conseguiu localizar o ficheiro especificado. |
| 118 | [devops-360-online/go-with-me-ml](https://github.com/devops-360-online/go-with-me-ml) | `D:\helm_clones_github\devops-360-online__go-with-me-ml\infrastructure\kubernetes\bento` | Error: no repository definition for https://warpstreamlabs.github.io/bento-helm-chart, https://warpstreamlabs.github.io/bento-helm-chart, https://warpstreamlabs.github.io/bento-helm-chart. Please add the missing repos via 'helm repo add' |
| 119 | [dm3drummer/arducharts](https://github.com/dm3drummer/arducharts) | `D:\helm_clones_github\dm3drummer__arducharts\configs\schema\ahrs` | Error: validation: chart.metadata.version "schema" is invalid |
| 120 | [epoch8/helm-charts](https://github.com/epoch8/helm-charts) | `D:\helm_clones_github\epoch8__helm-charts\examples\multi-simple-app` | Saving 4 charts Save error occurred:  can't get a valid version for dependency simple-app Error: can't get a valid version for dependency simple-app |
| 121 | [flashiam12/strategic-next-best-offer](https://github.com/flashiam12/strategic-next-best-offer) | `D:\helm_clones_github\flashiam12__strategic-next-best-offer\dashboard\dependencies\loki` | Error: no repository definition for https://charts.min.io/, https://grafana.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 122 | [hey101/scale-catalog](https://github.com/hey101/scale-catalog) | `D:\helm_clones_github\hey101__scale-catalog\incubator\archivebox\0.7.2` | Saving 1 charts Downloading common from repo oci://tccr.io/truecharts Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:17.5.1: not found Error: could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:17.5.1: not found |
| 123 | [highmobility/helm-charts](https://github.com/highmobility/helm-charts) | `D:\helm_clones_github\highmobility__helm-charts\examples\sample-app\helm-web` | Saving 1 charts Save error occurred:  can't get a valid version for dependency hm-basic-webapp Error: can't get a valid version for dependency hm-basic-webapp |
| 124 | [hmcts/hmcts-charts](https://github.com/hmcts/hmcts-charts) | `D:\helm_clones_github\hmcts__hmcts-charts\stable\aac-manage-case-assignment` | Getting updates for unmanaged Helm repositories... ...Successfully got an update from the "https://helm.elastic.co" chart repository Saving 4 charts Downloading java from repo oci://hmctsprod.azurecr.io/helm Save error occurred:  could not download oci://hmctsprod.azurecr.io/helm/java: failed to perform "FetchReference" on source: GET "https://hmctsprod.azurecr.io/v2/helm/java/manifests/5.3.0": GET "https://hmctsprod.azurecr.io/oauth2/token?scope=repository%3Ahelm%2Fjava%3Apull&service=hmctsprod.azurecr.io": response status code 401: unauthorized: authentication required, visit https://aka.ms/acr/authorization for more information. CorrelationId: 46eab12f-4b7c-42e6-af30-f7c8865fbb3b Error: could not download oci://hmctsprod.azurecr.io/helm/java: failed to perform "FetchReference" on source: GET "https://hmctsprod.azurecr.io/v2/helm/java/manifests/5.3.0": GET "https://hmctsprod.azurecr.io/oauth2/token?scope=repository%3Ahelm%2Fjava%3Apull&service=hmctsprod.azurecr.io": response status code 401: unauthorized: authentication required, visit https://aka.ms/acr/authorization for more information. CorrelationId: 46eab12f-4b7c-42e6-af30-f7c8865fbb3b |
| 125 | [hmcts/probate-frontend](https://github.com/hmcts/probate-frontend) | `D:\helm_clones_github\hmcts__probate-frontend\charts\probate-frontend` | Saving 2 charts Downloading nodejs from repo oci://hmctsprod.azurecr.io/helm Save error occurred:  could not download oci://hmctsprod.azurecr.io/helm/nodejs: failed to perform "FetchReference" on source: GET "https://hmctsprod.azurecr.io/v2/helm/nodejs/manifests/3.2.1": GET "https://hmctsprod.azurecr.io/oauth2/token?scope=repository%3Ahelm%2Fnodejs%3Apull&service=hmctsprod.azurecr.io": response status code 401: unauthorized: authentication required, visit https://aka.ms/acr/authorization for more information. CorrelationId: 6316736e-9d81-4ccb-8adb-3c3e0792e1eb Error: could not download oci://hmctsprod.azurecr.io/helm/nodejs: failed to perform "FetchReference" on source: GET "https://hmctsprod.azurecr.io/v2/helm/nodejs/manifests/3.2.1": GET "https://hmctsprod.azurecr.io/oauth2/token?scope=repository%3Ahelm%2Fnodejs%3Apull&service=hmctsprod.azurecr.io": response status code 401: unauthorized: authentication required, visit https://aka.ms/acr/authorization for more information. CorrelationId: 6316736e-9d81-4ccb-8adb-3c3e0792e1eb |
| 126 | [iamkhattar/homelab](https://github.com/iamkhattar/homelab) | `D:\helm_clones_github\iamkhattar__homelab\cluster\core\cert-manager` | Error: no repository definition for https://charts.jetstack.io. Please add the missing repos via 'helm repo add' |
| 127 | [joctan-tec/2023-02-2021069671-IC4302](https://github.com/joctan-tec/2023-02-2021069671-IC4302) | `D:\helm_clones_github\joctan-tec__2023-02-2021069671-IC4302\TareaCorta1\HelmCharts\databases` | Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 128 | [lorenzofratini1998/moneyverse](https://github.com/lorenzofratini1998/moneyverse) | `D:\helm_clones_github\lorenzofratini1998__moneyverse\moneyverse-infrastructure\moneyverse-chart\apps\clickhouse` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 129 | [matic-insurance/helm-charts](https://github.com/matic-insurance/helm-charts) | `D:\helm_clones_github\matic-insurance__helm-charts\ops\workbench\app-deployment` | Saving 5 charts Save error occurred:  can't get a valid version for dependency application-component Error: can't get a valid version for dependency application-component |
| 130 | [neuroxhq/helm-chart-neurox-workload](https://github.com/neuroxhq/helm-chart-neurox-workload) | `D:\helm_clones_github\neuroxhq__helm-chart-neurox-workload` | Saving 2 charts Downloading neurox-workload-agent from repo oci://ghcr.io/neuroxhq/helm-charts Save error occurred:  could not download oci://ghcr.io/neuroxhq/helm-charts/neurox-workload-agent: failed to perform "FetchReference" on source: GET "https://ghcr.io/v2/neuroxhq/helm-charts/neurox-workload-agent/manifests/2.10.2": GET "https://ghcr.io/token?scope=repository%3Aneuroxhq%2Fhelm-charts%2Fneurox-workload-agent%3Apull&service=ghcr.io": response status code 401: unauthorized: authentication required Error: could not download oci://ghcr.io/neuroxhq/helm-charts/neurox-workload-agent: failed to perform "FetchReference" on source: GET "https://ghcr.io/v2/neuroxhq/helm-charts/neurox-workload-agent/manifests/2.10.2": GET "https://ghcr.io/token?scope=repository%3Aneuroxhq%2Fhelm-charts%2Fneurox-workload-agent%3Apull&service=ghcr.io": response status code 401: unauthorized: authentication required |
| 131 | [osherlevi7/dev](https://github.com/osherlevi7/dev) | `D:\helm_clones_github\osherlevi7__dev\dev-tools\helm-tools\litmus` | Error: no repository definition for https://raw.githubusercontent.com/bitnami/charts/archive-full-index/bitnami. Please add the missing repos via 'helm repo add' |
| 132 | [piliphulko/marketplace-example](https://github.com/piliphulko/marketplace-example) | `D:\helm_clones_github\piliphulko__marketplace-example\k8s\cores-api` | Error: error unpacking subchart Makefile in core-api: Chart.yaml file is missing |
| 133 | [roman009/go-rust](https://github.com/roman009/go-rust) | `D:\helm_clones_github\roman009__go-rust\infrastructure\infra\kafka` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 134 | [sujan060/testing111](https://github.com/sujan060/testing111) | `D:\helm_clones_github\sujan060__testing111\packages\helm-charts\attestation-service` | Error: no repository definition for @stable. Please add them via 'helm repo add' |
| 135 | [tetratelabs/helm-charts](https://github.com/tetratelabs/helm-charts) | `D:\helm_clones_github\tetratelabs__helm-charts\charts\demos\istio-monitoring-demo` | Error: no repository definition for https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 136 | [wiredquill/ai-demos](https://github.com/wiredquill/ai-demos) | `D:\helm_clones_github\wiredquill__ai-demos\charts\hr-assistant` | Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update' |
| 137 | [Ammar-Taimoori/speechmatics-helm-chart](https://github.com/Ammar-Taimoori/speechmatics-helm-chart) | `D:\helm_clones_github\Ammar-Taimoori__speechmatics-helm-chart\1.2.0\sm-realtime` | Error: directory D:\helm_clones_github\Ammar-Taimoori__speechmatics-helm-chart\1.2.0\sm-proxy not found |
| 138 | [ArieLevs/Kubernetes-Helm-Charts](https://github.com/ArieLevs/Kubernetes-Helm-Charts) | `D:\helm_clones_github\ArieLevs__Kubernetes-Helm-Charts\charts\nalkinscloud` | Error: no repository definition for https://charts.bitnami.com/bitnami, https://helm.releases.hashicorp.com, https://arielevs.github.io/Kubernetes-Helm-Charts/, https://arielevs.github.io/Kubernetes-Helm-Charts/. Please add the missing repos via 'helm repo add' |
| 139 | [Arthur-B-DevOps/old_helm_charts](https://github.com/Arthur-B-DevOps/old_helm_charts) | `D:\helm_clones_github\Arthur-B-DevOps__old_helm_charts\charts\Old_charts\charts\incubator\distribution` | level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..." Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/, https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add' |
| 140 | [ArvigEnterprises/integration-tests-bluegreen](https://github.com/ArvigEnterprises/integration-tests-bluegreen) | `D:\helm_clones_github\ArvigEnterprises__integration-tests-bluegreen\.helm\charts\arvigtestcluster334d\canary` | Error: directory D:\helm_clones_github\ArvigEnterprises__integration-tests-bluegreen\.helm\charts\arvigtestcluster334d\laravel-library not found |
| 141 | [AssemblerBossss/tools_and_manuals](https://github.com/AssemblerBossss/tools_and_manuals) | `D:\helm_clones_github\AssemblerBossss__tools_and_manuals\kubernetes\18.templating\wordpress` | Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 142 | [BYGX-wcr/Magma_NB-IoT](https://github.com/BYGX-wcr/Magma_NB-IoT) | `D:\helm_clones_github\BYGX-wcr__Magma_NB-IoT\cn\deploy\helm` | Error: chart file "virtctl" is larger than the maximum file size 5242880 |
| 143 | [Blastz13/quality_reducer](https://github.com/Blastz13/quality_reducer) | `D:\helm_clones_github\Blastz13__quality_reducer\k8s\kube-prometheus-stack` | Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 144 | [CERIT-SC/rancher-apps-secure](https://github.com/CERIT-SC/rancher-apps-secure) | `D:\helm_clones_github\CERIT-SC__rancher-apps-secure\charts\monitoring\v0.1` | level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..." Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies |
| 145 | [Chalupa-Tech/chalupa-tech-local](https://github.com/Chalupa-Tech/chalupa-tech-local) | `D:\helm_clones_github\Chalupa-Tech__chalupa-tech-local\gitops\apps\media\clonarr` | Error: no repository definition for https://bjw-s-labs.github.io/helm-charts/. Please add the missing repos via 'helm repo add' |
| 146 | [CloudLargeScale-UCLouvain/c2b2](https://github.com/CloudLargeScale-UCLouvain/c2b2) | `D:\helm_clones_github\CloudLargeScale-UCLouvain__c2b2\caliper` | Error: no repository definition for https://halkeye.github.io/helm-charts/. Please add the missing repos via 'helm repo add' |
| 147 | [DEL-ORG/s7mensah](https://github.com/DEL-ORG/s7mensah) | `D:\helm_clones_github\DEL-ORG__s7mensah\argo-helm\charts\argo-cd` | Error: no repository definition for https://dandydeveloper.github.io/charts/. Please add the missing repos via 'helm repo add' |
| 148 | [DakEnviy/my-terraform](https://github.com/DakEnviy/my-terraform) | `D:\helm_clones_github\DakEnviy__my-terraform\internal\cluster\helm\gitlab` | Error: no repository definition for https://charts.jetstack.io/, https://prometheus-community.github.io/helm-charts, https://charts.bitnami.com/bitnami, https://charts.gitlab.io/, https://charts.bitnami.com/bitnami, https://charts.gitlab.io/. Please add the missing repos via 'helm repo add' |
| 149 | [Dystopian00/truecharts](https://github.com/Dystopian00/truecharts) | `D:\helm_clones_github\Dystopian00__truecharts\charts\incubator\docuseal` | Saving 1 charts Downloading common from repo oci://tccr.io/truecharts Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:24.1.5: not found Error: could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:24.1.5: not found |
| 150 | [Feederhigh5/master-thesis-poc](https://github.com/Feederhigh5/master-thesis-poc) | `D:\helm_clones_github\Feederhigh5__master-thesis-poc\litmus` | Error: no repository definition for https://raw.githubusercontent.com/bitnami/charts/archive-full-index/bitnami. Please add the missing repos via 'helm repo add' |
| 151 | [Genocs/genocs-argocd](https://github.com/Genocs/genocs-argocd) | `D:\helm_clones_github\Genocs__genocs-argocd\helm-dependency` | Error: validation: chart.metadata.version is required |
| 152 | [GovStackWorkingGroup/sandbox-bb-identity](https://github.com/GovStackWorkingGroup/sandbox-bb-identity) | `D:\helm_clones_github\GovStackWorkingGroup__sandbox-bb-identity\Esignet Deployment\Esignet Service\helm\oidc-ui` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 153 | [GroomC4/c4ang-infra](https://github.com/GroomC4/c4ang-infra) | `D:\helm_clones_github\GroomC4__c4ang-infra\charts\schema-registry` | Getting updates for unmanaged Helm repositories... ...Successfully got an update from the "https://confluentinc.github.io/cp-helm-charts/" chart repository Error: cp-schema-registry chart not found in repo https://confluentinc.github.io/cp-helm-charts/ |
| 154 | [HelixDevelopment/HelixGitpx](https://github.com/HelixDevelopment/HelixGitpx) | `D:\helm_clones_github\HelixDevelopment__HelixGitpx\docs\specifications\main\main_implementation_material\HelixGitpx\18-manifests` | Getting updates for unmanaged Helm repositories... ...Successfully got an update from the "https://cloudnative-pg.io/charts" chart repository ...Successfully got an update from the "https://go.temporal.io/helm-charts" chart repository ...Successfully got an update from the "https://qdrant.github.io/qdrant-helm" chart repository ...Successfully got an update from the "https://meilisearch.github.io/meilisearch-kubernetes" chart repository ...Successfully got an update from the "https://opensearch-project.github.io/helm-charts" chart repository ...Successfully got an update from the "https://spiffe.github.io/helm-charts-hardened" chart repository ...Successfully got an update from the "https://kyverno.github.io/kyverno" chart repository ...Successfully got an update from the "https://strimzi.io/charts" chart repository ...Successfully got an update from the "https://grafana.github.io/helm-charts" chart repository ...Successfully got an update from the "https://prometheus-community.github.io/helm-charts" chart repository ...Successfully got an update from the "https://istio-release.storage.googleapis.com/charts" chart repository Error: istio-base chart not found in repo https://istio-release.storage.googleapis.com/charts |
| 155 | [ICICLE-ai/OpenPass](https://github.com/ICICLE-ai/OpenPass) | `D:\helm_clones_github\ICICLE-ai__OpenPass\helm-config` | Error: cannot load values.yaml: cannot unmarshal yaml document: error converting YAML to JSON: yaml: line 30: could not find expected ':' |
| 156 | [Jaebytes/TrueCharts](https://github.com/Jaebytes/TrueCharts) | `D:\helm_clones_github\Jaebytes__TrueCharts\premium\traefik\26.10.19` | Saving 1 charts Downloading common from repo oci://tccr.io/truecharts Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:20.3.12: not found Error: could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:20.3.12: not found |
| 157 | [Keshari07/Kubernetes_28_29_September_2024](https://github.com/Keshari07/Kubernetes_28_29_September_2024) | `D:\helm_clones_github\Keshari07__Kubernetes_28_29_September_2024\kubelabs\Helm101\Wordpress` | Error: cannot load Chart.yaml: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go struct field Metadata.deprecated of type bool |
| 158 | [LTKH/minikube](https://github.com/LTKH/minikube) | `D:\helm_clones_github\LTKH__minikube\argocd\charts\argo-cd` | Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update' |
| 159 | [LeeLisker/helm](https://github.com/LeeLisker/helm) | `D:\helm_clones_github\LeeLisker__helm\community\artifactory-ha` | level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..." Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add' |
| 160 | [MUST-CLOUD/osh](https://github.com/MUST-CLOUD/osh) | `D:\helm_clones_github\MUST-CLOUD__osh\openstack-helm\openstack` | Error: error unpacking subchart rabbitmq in openstack: Chart.yaml file is missing |
| 161 | [Mainuddin-Rizvi/devops](https://github.com/Mainuddin-Rizvi/devops) | `D:\helm_clones_github\Mainuddin-Rizvi__devops\kubernetes-k8s\12-deploying-to-multiple-environments\kluctl\.helm-charts\https_cloudnative-pg.github.io\charts\cloudnative-pg\0.21.4` | Error: no repository definition for https://cloudnative-pg.github.io/grafana-dashboards. Please add the missing repos via 'helm repo add' |
| 162 | [Mark4562/Truenas-Scale](https://github.com/Mark4562/Truenas-Scale) | `D:\helm_clones_github\Mark4562__Truenas-Scale\premium\authelia\23.13.13` | Saving 2 charts Downloading common from repo oci://tccr.io/truecharts Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:20.3.11: not found Error: could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:20.3.11: not found |
| 163 | [Musaiba/Research-investigation-tools](https://github.com/Musaiba/Research-investigation-tools) | `D:\helm_clones_github\Musaiba__Research-investigation-tools\theHarvester\deploy\charts\harvester` | Error: no repository definition for https://charts.harvesterhci.io, https://charts.harvesterhci.io, https://charts.longhorn.io, https://kube-vip.github.io/helm-charts, https://kube-vip.github.io/helm-charts, https://charts.harvesterhci.io, https://charts.harvesterhci.io. Please add the missing repos via 'helm repo add' |
| 164 | [MxBlu/tcharts](https://github.com/MxBlu/tcharts) | `D:\helm_clones_github\MxBlu__tcharts\charts\home-assistant\1.0.130` | Error: directory D:\helm_clones_github\MxBlu__tcharts\common\2304.0.1 not found |
| 165 | [My-DIGI-ID/ssi-helm-charts](https://github.com/My-DIGI-ID/ssi-helm-charts) | `D:\helm_clones_github\My-DIGI-ID__ssi-helm-charts\charts\ssi-aca-py` | Error: no repository definition for https://charts.bitnami.com/bitnami, https://my-digi-id.github.io/ssi-helm-charts/. Please add the missing repos via 'helm repo add' |
| 166 | [OpenSourceConsulting/playcekube-charts](https://github.com/OpenSourceConsulting/playcekube-charts) | `D:\helm_clones_github\OpenSourceConsulting__playcekube-charts\src\playcekube\keycloak\9.2.2\keycloak` | Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 167 | [OriHoch/sk8s-ops](https://github.com/OriHoch/sk8s-ops) | `D:\helm_clones_github\OriHoch__sk8s-ops` | Error: validation: chart.metadata.version is required |
| 168 | [Otus-DevOps-2022-02/saintjb_microservices](https://github.com/Otus-DevOps-2022-02/saintjb_microservices) | `D:\helm_clones_github\Otus-DevOps-2022-02__saintjb_microservices\kubernetes\Charts\gitlab` | Getting updates for unmanaged Helm repositories... ...Successfully got an update from the "https://charts.jetstack.io/" chart repository ...Successfully got an update from the "https://grafana.github.io/helm-charts" chart repository ...Successfully got an update from the "https://prometheus-community.github.io/helm-charts" chart repository ...Successfully got an update from the "https://charts.gitlab.io/" chart repository ...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository Error: can't get a valid version for 2 subchart(s): "postgresql" (repository "https://charts.bitnami.com/bitnami", version "8.9.4"), "redis" (repository "https://charts.bitnami.com/bitnami", version "11.3.4"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml |
| 169 | [PaulJequann/the-lab](https://github.com/PaulJequann/the-lab) | `D:\helm_clones_github\PaulJequann__the-lab\kubernetes\apps\glitchtip` | Error: no repository definition for https://gitlab.com/api/v4/projects/16325141/packages/helm/stable. Please add the missing repos via 'helm repo add' |
| 170 | [Riskin1999/arena-appwrapper-volcano](https://github.com/Riskin1999/arena-appwrapper-volcano) | `D:\helm_clones_github\Riskin1999__arena-appwrapper-volcano\arena-artifacts` | Error: no repository definition for @tf-operator, @tf-dashbard, @cron-operator, @et-operator, @mpi-operator, @pytorch-operator, @gpu-exporter, @elastic-job-supervisor. Please add them via 'helm repo add' |
| 171 | [SalesConnection/sc-helm-charts](https://github.com/SalesConnection/sc-helm-charts) | `D:\helm_clones_github\SalesConnection__sc-helm-charts\charts\caction` | Error: no repository definition for https://salesconnection.github.io/sc-helm-charts, https://salesconnection.github.io/sc-helm-charts, https://salesconnection.github.io/sc-helm-charts. Please add the missing repos via 'helm repo add' |
| 172 | [ShaharyarShakir/devops-practice](https://github.com/ShaharyarShakir/devops-practice) | `D:\helm_clones_github\ShaharyarShakir__devops-practice\16__gitops\projects\vprofile-action\helm\vprofilecharts` | Error: error unpacking subchart templates in vprofilecharts: Chart.yaml file is missing |
| 173 | [SirObi/airflow_k8s_executor](https://github.com/SirObi/airflow_k8s_executor) | `D:\helm_clones_github\SirObi__airflow_k8s_executor\helm_charts\official\charts\incubator\distribution` | level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..." Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/, https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add' |
| 174 | [StopDenBus/helm-charts](https://github.com/StopDenBus/helm-charts) | `D:\helm_clones_github\StopDenBus__helm-charts\charts\external-dns` | Saving 1 charts Downloading external-dns-chart from repo oci://dhi.io Save error occurred:  could not download oci://dhi.io/external-dns-chart: failed to perform "FetchReference" on source: GET "https://dhi.io/v2/external-dns-chart/manifests/1.21.1": GET "https://dhi.io/token?scope=repository%3Aexternal-dns-chart%3Apull&service=registry.docker.io": response status code 401: Unauthorized Error: could not download oci://dhi.io/external-dns-chart: failed to perform "FetchReference" on source: GET "https://dhi.io/v2/external-dns-chart/manifests/1.21.1": GET "https://dhi.io/token?scope=repository%3Aexternal-dns-chart%3Apull&service=registry.docker.io": response status code 401: Unauthorized |
| 175 | [SuperCoolAlan/homelab-manifests](https://github.com/SuperCoolAlan/homelab-manifests) | `D:\helm_clones_github\SuperCoolAlan__homelab-manifests\talos\authentik\charts\authentik-2026.2.1` | Error: no repository definition for https://charts.goauthentik.io. Please add the missing repos via 'helm repo add' |
| 176 | [Timkhch/school-dev-k8s](https://github.com/Timkhch/school-dev-k8s) | `D:\helm_clones_github\Timkhch__school-dev-k8s\practice\18.templating\wordpress` | Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 177 | [TrevorSquillario/iDRAC-Telemetry-Ansible-Demo](https://github.com/TrevorSquillario/iDRAC-Telemetry-Ansible-Demo) | `D:\helm_clones_github\TrevorSquillario__iDRAC-Telemetry-Ansible-Demo\k8s\helm\app` | Error: no repository definition for https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 178 | [TrivialJavaBackender/infrastructure-learn](https://github.com/TrivialJavaBackender/infrastructure-learn) | `D:\helm_clones_github\TrivialJavaBackender__infrastructure-learn\exercises\helm\Ex15_BasicChart\myapp` | Error: validation: chart.metadata.version is required |
| 179 | [Yuzu815/TruenasScaleArchive](https://github.com/Yuzu815/TruenasScaleArchive) | `D:\helm_clones_github\Yuzu815__TruenasScaleArchive\incubator\archivebox\0.7.2` | Saving 1 charts Downloading common from repo oci://tccr.io/truecharts Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:17.2.30: not found Error: could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:17.2.30: not found |
| 180 | [aashkulkarni/helm](https://github.com/aashkulkarni/helm) | `D:\helm_clones_github\aashkulkarni__helm\schoolapp-subchart` | Error: no repository definition for https://gitlab.com/api/v4/projects/34240616/packages/helm/stable, https://gitlab.com/api/v4/projects/34240616/packages/helm/stable, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 181 | [abhilashindulkar/misc](https://github.com/abhilashindulkar/misc) | `D:\helm_clones_github\abhilashindulkar__misc\helm\custom-charts\couponservice` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 182 | [abimaelrsergio/Abimoney](https://github.com/abimaelrsergio/Abimoney) | `D:\helm_clones_github\abimaelrsergio__Abimoney\charts-main\bitnami\sealed-secrets` | Saving 1 charts Downloading common from repo oci://registry-1.docker.io/bitnamicharts Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/common: failed to perform "FetchReference" on source: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.31.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/common: failed to perform "FetchReference" on source: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.31.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 183 | [acaisoft/bolt-platform-helm](https://github.com/acaisoft/bolt-platform-helm) | `D:\helm_clones_github\acaisoft__bolt-platform-helm` | Error: no repository definition for http://bolt-helm-charts.acaisoft.net, http://bolt-helm-charts.acaisoft.net, https://kubernetes.github.io/ingress-nginx, https://argoproj.github.io/argo-helm, https://hasura-extra.github.io/hasura-extra. Please add the missing repos via 'helm repo add' |
| 184 | [actions-marketplace-validations/LinuxSuRen_api-testing](https://github.com/actions-marketplace-validations/LinuxSuRen_api-testing) | `D:\helm_clones_github\actions-marketplace-validations__LinuxSuRen_api-testing\helm\api-testing` | Saving 1 charts Downloading mongodb from repo oci://registry-1.docker.io/bitnamicharts Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/mongodb: failed to perform "FetchReference" on source: GET "https://registry-1.docker.io/v2/bitnamicharts/mongodb/manifests/15.1.7": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/mongodb: failed to perform "FetchReference" on source: GET "https://registry-1.docker.io/v2/bitnamicharts/mongodb/manifests/15.1.7": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 185 | [aicanhelp/ai-cloud](https://github.com/aicanhelp/ai-cloud) | `D:\helm_clones_github\aicanhelp__ai-cloud\bitnami\charts\bitnami-202503\airflow` | Saving 3 charts Downloading redis from repo oci://registry-1.docker.io/bitnamicharts Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to perform "FetchReference" on source: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/20.7.1": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to perform "FetchReference" on source: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/20.7.1": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 186 | [albinoJimy/Neural-Hive-Mind](https://github.com/albinoJimy/Neural-Hive-Mind) | `D:\helm_clones_github\albinoJimy__Neural-Hive-Mind\helm\istio-base` | Error: no repository definition for https://istio-release.storage.googleapis.com/charts, https://istio-release.storage.googleapis.com/charts, https://istio-release.storage.googleapis.com/charts. Please add the missing repos via 'helm repo add' |
| 187 | [alnav3/argocd-charts](https://github.com/alnav3/argocd-charts) | `D:\helm_clones_github\alnav3__argocd-charts\charts\argo-cd` | Error: no repository definition for https://argoproj.github.io/argo-helm. Please add the missing repos via 'helm repo add' |
| 188 | [andgit7/repo2](https://github.com/andgit7/repo2) | `D:\helm_clones_github\andgit7__repo2\Kubernetes\slurm\practice\18.templating\wordpress` | Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 189 | [anmol372/matrix-chart](https://github.com/anmol372/matrix-chart) | `D:\helm_clones_github\anmol372__matrix-chart` | Error: no repository definition for https://kubernetes-charts.storage.googleapis.com, https://vmware-tanzu.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 190 | [anzerr/catalog](https://github.com/anzerr/catalog) | `D:\helm_clones_github\anzerr__catalog\dependency\mariadb\12.2.0` | Saving 1 charts Downloading common from repo oci://tccr.io/truecharts Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:17.4.2: not found Error: could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:17.4.2: not found |
| 191 | [artichoked1/FerriteCMS-chart](https://github.com/artichoked1/FerriteCMS-chart) | `D:\helm_clones_github\artichoked1__FerriteCMS-chart` | Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 192 | [arychj/argocd](https://github.com/arychj/argocd) | `D:\helm_clones_github\arychj__argocd\wordpress` | Error: validation: chart.metadata.version is required |
| 193 | [beantownpub/helm](https://github.com/beantownpub/helm) | `D:\helm_clones_github\beantownpub__helm\argo-cd` | Error: no repository definition for https://argoproj.github.io/argo-helm. Please add the missing repos via 'helm repo add' |
| 194 | [bgarcia-elastic/tfk8](https://github.com/bgarcia-elastic/tfk8) | `D:\helm_clones_github\bgarcia-elastic__tfk8` | Error: directory D:\helm_clones_github\subCharts\ui not found |

## Full Error Details

### `D:\helm_clones_github\bitnami__charts\bitnami\apache` — bitnami/charts

```
Saving 1 charts
Downloading common from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/common: failed to perform "FetchReference" on source: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.31.4": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/common: failed to perform "FetchReference" on source: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.31.4": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `D:\helm_clones_github\higress-group__higress\helm\core` — higress-group/higress

```
Error: directory D:\helm_clones_github\higress-group__higress\helm\redis not found
```

### `D:\helm_clones_github\feast-dev__feast\infra\charts\feast` — feast-dev/feast

```
Getting updates for unmanaged Helm repositories...
...Successfully got an update from the "https://feast-helm-charts.storage.googleapis.com" chart repository
...Successfully got an update from the "https://charts.helm.sh/stable" chart repository
Error: feature-server chart not found in repo https://feast-helm-charts.storage.googleapis.com
```

### `D:\helm_clones_github\prometheus-community__helm-charts\charts\prometheus` — prometheus-community/helm-charts

```
Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\collabnix__kubelabs\Helm101\Wordpress` — collabnix/kubelabs

```
Error: cannot load Chart.yaml: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go struct field Metadata.deprecated of type bool
```

### `D:\helm_clones_github\cozystack__cozystack\packages\apps\bucket` — cozystack/cozystack

```
Error: error unpacking subchart cozy-lib in bucket: Chart.yaml file is missing
```

### `D:\helm_clones_github\grafana__helm-charts\charts\enterprise-logs` — grafana/helm-charts

```
Error: no repository definition for https://grafana.github.io/helm-charts, https://charts.min.io/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\pluralsh__plural\plural\helm\plural` — pluralsh/plural

```
Error: no repository definition for https://k8s.ory.sh/helm/charts, https://pluralsh.github.io/module-library. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\trueforge-org__truecharts\charts\stable\kubernetes-dashboard` — trueforge-org/truecharts

```
Error: no repository definition for https://charts.konghq.com. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\slurm-personal__school-dev-k8s\practice\18.templating\wordpress` — slurm-personal/school-dev-k8s

```
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\kubeflow__arena\arena-artifacts` — kubeflow/arena

```
Error: no repository definition for @tf-operator, @tf-dashbard, @cron-operator, @et-operator, @mpi-operator, @pytorch-operator, @gpu-exporter, @elastic-job-supervisor. Please add them via 'helm repo add'
```

### `D:\helm_clones_github\fluxninja__aperture\manifests\charts\aperture-agent` — fluxninja/aperture

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\WeBankFinTech__Prophecis\install\Prophecis` — WeBankFinTech/Prophecis

```
Error: cannot load values.yaml: error reading yaml document: invalid Yaml document separator: --END RSA PRIVATE KEY-----"
```

### `D:\helm_clones_github\rancher__charts\charts\epinio\102.0.1+up1.6.2` — rancher/charts

```
Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update'
```

### `D:\helm_clones_github\DataDog__helm-charts\charts\datadog` — DataDog/helm-charts

```
Error: no repository definition for https://helm.datadoghq.com, https://prometheus-community.github.io/helm-charts, https://helm.datadoghq.com, https://helm.datadoghq.com. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\radondb__radondb-mysql-kubernetes\charts\mysql-operator` — radondb/radondb-mysql-kubernetes

```
Error: directory D:\helm_clones_github\radondb__radondb-mysql-kubernetes\charts\mysql-operator\charts\mysqlcluster not found
```

### `D:\helm_clones_github\truenas__charts\charts\collabora\1.2.30` — truenas/charts

```
Error: directory D:\helm_clones_github\truenas__charts\common\2304.0.1 not found
```

### `D:\helm_clones_github\instantlinux__docker-tools\images\git-pull\helm` — instantlinux/docker-tools

```
Error: directory D:\helm_clones_github\instantlinux__docker-tools\images\git-pull\chartlib not found
```

### `D:\helm_clones_github\IBM__charts\community\artifactory-ha` — IBM/charts

```
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\mercedes-benz__DnA\deployment\kubernetes\mysql-helm-chart` — mercedes-benz/DnA

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\Hydrospheredata__hydro-serving\helm` — Hydrospheredata/hydro-serving

```
Error: cannot load Chart.yaml: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type v2.Metadata
```

### `D:\helm_clones_github\JahstreetOrg__spark-on-kubernetes-helm\charts\cluster-base` — JahstreetOrg/spark-on-kubernetes-helm

```
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: no repository definition for https://kubernetes.github.io/ingress-nginx, https://charts.jetstack.io, https://charts.helm.sh/stable, https://charts.helm.sh/stable. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\osm-seed__osm-seed\osm-seed` — osm-seed/osm-seed

```
Error: validation: chart.metadata.version is required
```

### `D:\helm_clones_github\quanxiang-cloud__quanxiang\deployment\charts\elasticsearch` — quanxiang-cloud/quanxiang

```
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\CenterForOpenScience__helm-charts\elastic-stack` — CenterForOpenScience/helm-charts

```
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: no repository definition for https://centerforopenscience.github.io/helm-charts/, https://centerforopenscience.github.io/helm-charts/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\companyinfo__helm-charts\charts\helmet\examples\simple` — companyinfo/helm-charts

```
Error: no repository definition for https://charts.companyinfo.dev. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\rancher__rke2-charts\packages\rke2-cilium-legacy\charts` — rancher/rke2-charts

```
Error: directory D:\helm_clones_github\rancher__rke2-charts\packages\rke2-cilium-legacy\charts\charts\cilium not found
```

### `D:\helm_clones_github\Obmondo__KubeAid\argocd-helm-charts\argo-cd` — Obmondo/KubeAid

```
Error: error unpacking subchart kubeaid-addons in argo-cd: Chart.yaml file is missing
```

### `D:\helm_clones_github\rancher__partner-charts\charts\amd\amd-gpu\0.10.0` — rancher/partner-charts

```
Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update'
```

### `D:\helm_clones_github\lianqingsec__NucleiPocGather\poc\other` — lianqingsec/NucleiPocGather

```
Error: chart file "signatures_1.yaml" is larger than the maximum file size 5242880
```

### `D:\helm_clones_github\elastisys__compliantkubernetes-apps\helmfile.d\upstream\falcosecurity\falco` — elastisys/compliantkubernetes-apps

```
Error: no repository definition for https://falcosecurity.github.io/charts, https://falcosecurity.github.io/charts, https://falcosecurity.github.io/charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\wiremind__wiremind-helm-charts\charts\druid` — wiremind/wiremind-helm-charts

```
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://machine424.github.io/kube-hpa-scale-to-zero, https://wiremind.github.io/wiremind-helm-charts, https://wiremind.github.io/wiremind-helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\bytle__kubee` — bytle/kubee

```
Error: error unpacking subchart README.md in Kubee: Chart.yaml file is missing
```

### `D:\helm_clones_github\tmforum-oda__oda-canvas\charts\apisix-gateway` — tmforum-oda/oda-canvas

```
Error: no repository definition for https://charts.apiseven.com. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\IndustryFusion__DigitalTwin\helm\charts\velero` — IndustryFusion/DigitalTwin

```
Error: directory D:\helm_clones_github\IndustryFusion__DigitalTwin\helm\airgap-deployment\helm-charts\charts\velero not found
```

### `D:\helm_clones_github\boozallen__aissemble\extensions\extensions-helm\aissemble-airflow-chart` — boozallen/aissemble

```
Error: no repository definition for https://airflow.apache.org/. Please add the missing repos via 'helm repo add'
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

### `D:\helm_clones_github\k8s-home-lab__helm-charts\charts\stable\bazarr` — k8s-home-lab/helm-charts

```
Error: no repository definition for https://k8s-home-lab.github.io/helm-charts/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\neuroxhq__helm-chart-neurox-control` — neuroxhq/helm-chart-neurox-control

```
Saving 6 charts
Downloading neurox-control-api from repo oci://ghcr.io/neuroxhq/helm-charts
Save error occurred:  could not download oci://ghcr.io/neuroxhq/helm-charts/neurox-control-api: failed to perform "FetchReference" on source: GET "https://ghcr.io/v2/neuroxhq/helm-charts/neurox-control-api/manifests/2.233.1": GET "https://ghcr.io/token?scope=repository%3Aneuroxhq%2Fhelm-charts%2Fneurox-control-api%3Apull&service=ghcr.io": response status code 401: unauthorized: authentication required
Error: could not download oci://ghcr.io/neuroxhq/helm-charts/neurox-control-api: failed to perform "FetchReference" on source: GET "https://ghcr.io/v2/neuroxhq/helm-charts/neurox-control-api/manifests/2.233.1": GET "https://ghcr.io/token?scope=repository%3Aneuroxhq%2Fhelm-charts%2Fneurox-control-api%3Apull&service=ghcr.io": response status code 401: unauthorized: authentication required
```

### `D:\helm_clones_github\pluralsh__console\charts\console` — pluralsh/console

```
Error: no repository definition for https://fluxcd-community.github.io/helm-charts, https://charts.dexidp.io, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\YAKEcloud__yake\helmcharts\acl` — YAKEcloud/yake

```
Error: dependency "controller" has an invalid version/constraint format: improper constraint: ""
```

### `D:\helm_clones_github\GreptimeTeam__helm-charts\charts\greptimedb-cluster` — GreptimeTeam/helm-charts

```
Error: no repository definition for https://grafana.github.io/helm-charts, https://raw.githubusercontent.com/hansehe/jaeger-all-in-one/master/helm/charts, https://greptimeteam.github.io/helm-charts, https://greptimeteam.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\k0rdent__catalog\apps\alloy\charts\alloy-1.6.1` — k0rdent/catalog

```
Error: no repository definition for https://grafana.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\junghoon2__k8s-class\argo-cd\argo-cd-5.14.1` — junghoon2/k8s-class

```
Error: no repository definition for https://dandydeveloper.github.io/charts/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\silogen__cluster-forge\sources\amd-gpu-operator\v1.3.1` — silogen/cluster-forge

```
Error: no repository definition for https://kubernetes-sigs.github.io/node-feature-discovery/charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\ai-solution-eng__frameworks\appsmith\3.6.4` — ai-solution-eng/frameworks

```
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\wenerme__kube-stub-cluster\keycloak` — wenerme/kube-stub-cluster

```
Saving 1 charts
Downloading keycloak from repo oci://dockercr.wener.me/bitnamicharts
Save error occurred:  could not download oci://dockercr.wener.me/bitnamicharts/keycloak: failed to perform "FetchReference" on source: GET "https://dockercr.wener.me/v2/bitnamicharts/keycloak/manifests/15.1.4": invalid response `Content-Type` header; mime: no media type
Error: could not download oci://dockercr.wener.me/bitnamicharts/keycloak: failed to perform "FetchReference" on source: GET "https://dockercr.wener.me/v2/bitnamicharts/keycloak/manifests/15.1.4": invalid response `Content-Type` header; mime: no media type
```

### `D:\helm_clones_github\ThienAnTrinh__product-search\monitoring\logs-metrics\helm-charts\charts\kube-prometheus-stack` — ThienAnTrinh/product-search

```
Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\nebius__nebius-k8s-applications` — nebius/nebius-k8s-applications

```
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "https://volcano-sh.github.io/charts/" chart repository:
	failed to fetch https://volcano-sh.github.io/charts/index.yaml : 404 Not Found
...Successfully got an update from the "https://mellanox.github.io/network-operator" chart repository
...Successfully got an update from the "https://ray-project.github.io/kuberay-helm/" chart repository
...Successfully got an update from the "https://qdrant.github.io/qdrant-helm" chart repository
...Successfully got an update from the "https://airflow.apache.org" chart repository
...Successfully got an update from the "https://otwld.github.io/ollama-helm/" chart repository
...Successfully got an update from the "https://zilliztech.github.io/milvus-helm/" chart repository
...Successfully got an update from the "https://cowboysysop.github.io/charts/" chart repository
...Successfully got an update from the "https://argoproj.github.io/argo-helm" chart repository
...Successfully got an update from the "https://hub.jupyter.org/helm-chart" chart repository
...Unable to get an update from the "https://github.com/weaviate/weaviate" chart repository:
	failed to fetch https://github.com/weaviate/weaviate/index.yaml : 404 Not Found
...Successfully got an update from the "https://grafana.github.io/helm-charts" chart repository
...Successfully got an update from the "https://helm.ngc.nvidia.com/nvidia" chart repository
Error: directory D:\helm_clones_github\nebius__nebius-k8s-applications\charts\clearml-agent not found
```

### `D:\helm_clones_github\suse-edge__charts\charts\kubevirt\0.1.0` — suse-edge/charts

```
Error: dependency "cdi" has an invalid version/constraint format: improper constraint: ""
```

### `D:\helm_clones_github\wx-chevalier__k8s-examples\helm-charts\backend-spring-app` — wx-chevalier/k8s-examples

```
Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update'
```

### `D:\helm_clones_github\cnrancher__pandaria-catalog\charts\rancher-hami\107.0.0+up2.5.2\charts\hami-webui` — cnrancher/pandaria-catalog

```
Error: no repository definition for https://nvidia.github.io/dcgm-exporter/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\DaoCloud__dce-charts-repackage\charts\argo-cd\argo-cd\charts\argo-cd` — DaoCloud/dce-charts-repackage

```
Error: no repository definition for https://dandydeveloper.github.io/charts/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\sebolabs__eks-tf-gitops\k8s\add-ons\csi-secrets-store-provider-aws` — sebolabs/eks-tf-gitops

```
Error: cannot load Chart.yaml: error converting YAML to JSON: yaml: line 3: mapping values are not allowed in this context
```

### `D:\helm_clones_github\StatCan__charts\deprecated\cost-analyzer` — StatCan/charts

```
Error: no repository definition for https://kubecost.github.io/cost-analyzer/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\jordanopensource__charts\charts\etherpad` — jordanopensource/charts

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\jugatsu__microservices\deploy\kubernetes\Charts\gitlab-omnibus` — jugatsu/microservices

```
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies
```

### `D:\helm_clones_github\llajas__homelab\apps\clusterplex` — llajas/homelab

```
Error: no repository definition for https://bjw-s-labs.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\replicatedhq__chartsmith\chart\chartsmith` — replicatedhq/chartsmith

```
Saving 1 charts
Downloading replicated from repo oci://registry.chartsmith.ai/library
Save error occurred:  could not download oci://registry.chartsmith.ai/library/replicated: failed to perform "FetchReference" on source: Get "https://registry.chartsmith.ai/v2/library/replicated/manifests/1.12.1": dial tcp: lookup registry.chartsmith.ai: no such host
Error: could not download oci://registry.chartsmith.ai/library/replicated: failed to perform "FetchReference" on source: Get "https://registry.chartsmith.ai/v2/library/replicated/manifests/1.12.1": dial tcp: lookup registry.chartsmith.ai: no such host
```

### `D:\helm_clones_github\commercialhaskell__all-cabal-metadata\packages\ch` — commercialhaskell/all-cabal-metadata

```
Error: validation: chart.metadata.name is required
```

### `D:\helm_clones_github\NeonGeckoCom__neon-diana-utils\neon_diana_utils\helm_charts\http\libretranslate` — NeonGeckoCom/neon-diana-utils

```
Error: can't get a valid version for 1 subchart(s): "base-http" (repository "file://../../base/base-http", version "0.0.6"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml
```

### `D:\helm_clones_github\easy2stake__thegraph\graphprotocol` — easy2stake/thegraph

```
Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://grafana.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\thoughtworks__byor-voting-infrastructure\src\byor-voting-chart` — thoughtworks/byor-voting-infrastructure

```
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies
```

### `D:\helm_clones_github\TWilkin__powerpi\kubernetes` — TWilkin/powerpi

```
Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update'
```

### `D:\helm_clones_github\safesoftware__helm-charts\chart-source\fmeserver-2018.1.1` — safesoftware/helm-charts

```
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: no repository definition for https://kubernetes-charts.storage.googleapis.com. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\anza-labs__charts\deprecated\hosted-control-plane` — anza-labs/charts

```
Error: no repository definition for https://charts.jetstack.io. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\appscode-cloud__ui-wizards\charts\kubedbcom-elasticsearch-editor` — appscode-cloud/ui-wizards

```
Error: chart file "values.openapiv3_schema.yaml" is larger than the maximum file size 5242880
```

### `D:\helm_clones_github\joostvdg__cmg\charts\preview` — joostvdg/cmg

```
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "http://chartmuseum.jenkins-x.io" chart repository:
	failed to fetch http://chartmuseum.jenkins-x.io/index.yaml : 404 Not Found
Error: no cached repository for helm-manager-d8c90e93d45e753db0fa4ec457d893e08cd1e0827b77896ee636b714706cbbb2 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-d8c90e93d45e753db0fa4ec457d893e08cd1e0827b77896ee636b714706cbbb2-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### `D:\helm_clones_github\salmanmkc__agentverse\mcp-backend\charts\rag-stack` — salmanmkc/agentverse

```
Error: no repository definition for https://helm.neo4j.com/neo4j, https://helm.neo4j.com/neo4j, https://zilliztech.github.io/milvus-helm/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\samgabrail__env0-argocd\schoolapp-subchart` — samgabrail/env0-argocd

```
Error: no repository definition for https://gitlab.com/api/v4/projects/34240616/packages/helm/stable, https://gitlab.com/api/v4/projects/34240616/packages/helm/stable, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\ai-on-gke__common-infra\common\charts\gmp-engine` — ai-on-gke/common-infra

```
Error: dependency "gmp-frontend" has an invalid version/constraint format: improper constraint: ""
```

### `D:\helm_clones_github\ibuildthecloud__rancher-charts\charts\anchore-engine\0.1.0` — ibuildthecloud/rancher-charts

```
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: no repository definition for https://kubernetes-charts.storage.googleapis.com. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\wunderio__charts\drupal` — wunderio/charts

```
Error: no repository definition for https://storage.googleapis.com/charts.wdr.io, https://percona.github.io/percona-helm-charts/, https://storage.googleapis.com/charts.wdr.io, https://storage.googleapis.com/charts.wdr.io, https://storage.googleapis.com/charts.wdr.io, https://jouve.github.io/charts/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\SocialGouv__no-package-malware\charts\no-package-malware` — SocialGouv/no-package-malware

```
Error: no repository definition for https://groundhog2k.github.io/helm-charts/, https://cloudpirates-io.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\StackVista__helm-charts\stable\otel-demo` — StackVista/helm-charts

```
Error: no repository definition for https://open-telemetry.github.io/opentelemetry-helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\anup1384__helm-charts\stable\kafka` — anup1384/helm-charts

```
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies
```

### `D:\helm_clones_github\hifer__devops\monitor\prometheus-operator` — hifer/devops

```
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/, https://kubernetes-charts.storage.googleapis.com/, https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\kubebb__components\charts\cluster-component\charts\openebs` — kubebb/components

```
Error: no repository definition for https://openebs.github.io/node-disk-manager, https://openebs.github.io/dynamic-localpv-provisioner, https://openebs.github.io/cstor-operators, https://openebs.github.io/jiva-operator, https://openebs.github.io/zfs-localpv, https://openebs.github.io/lvm-localpv, https://openebs.github.io/dynamic-nfs-provisioner. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\marcosviniciusi__k3s-homelab\kustomize\infisical-stack\infisical-server` — marcosviniciusi/k3s-homelab

```
Error: no repository definition for https://kubernetes.github.io/ingress-nginx, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\AISphere__ffdl-trainer` — AISphere/ffdl-trainer

```
Error: directory D:\helm_clones_github\ffdl-lcm not found
```

### `D:\helm_clones_github\Makhuta__truecharts-archive-scale-catalog\incubator\archivebox\0.7.2` — Makhuta/truecharts-archive-scale-catalog

```
Saving 1 charts
Downloading common from repo oci://tccr.io/truecharts
Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:17.2.30: not found
Error: could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:17.2.30: not found
```

### `D:\helm_clones_github\MegaWiz-Dev-Team__Asgard\charts\asgard` — MegaWiz-Dev-Team/Asgard

```
Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update'
```

### `D:\helm_clones_github\Obsidian-Owl__floe\charts\cognee-platform` — Obsidian-Owl/floe

```
Error: no repository definition for https://helm.neo4j.com/neo4j, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\Yunjuzhen__charts\stable\anchore-engine\0.1.3\anchore-engine` — Yunjuzhen/charts

```
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: no repository definition for https://kubernetes-charts.storage.googleapis.com. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\joshleecreates__clickhouse-opentelemetry-iac\argo-apps\qryn` — joshleecreates/clickhouse-opentelemetry-iac

```
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "https://metrico.github.io/qryn-helm/" chart repository:
	failed to fetch https://metrico.github.io/qryn-helm/index.yaml : 404 Not Found
Error: no cached repository for helm-manager-fed7dc84065c3ccd251fae7cd72350b43bda8f8cb634374dd8b3a9318ac9d4e3 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-fed7dc84065c3ccd251fae7cd72350b43bda8f8cb634374dd8b3a9318ac9d4e3-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### `D:\helm_clones_github\pluralsh__plural-helm-charts\charts\airbyte` — pluralsh/plural-helm-charts

```
Error: no repository definition for https://charts.bitnami.com/bitnami, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\rancher__ob-team-charts\charts\prometheus-federator\0.0.1` — rancher/ob-team-charts

```
Error: dependency "helmProjectOperator" has an invalid version/constraint format: improper constraint: ""
```

### `D:\helm_clones_github\Andrew-Su-0718__zelos-image\mmdet\image\charts\arena-artifacts` — Andrew-Su-0718/zelos-image

```
Error: no repository definition for @tf-operator, @tf-dashbard, @cron-operator, @et-operator, @mpi-operator, @pytorch-operator, @gpu-exporter, @elastic-job-supervisor. Please add them via 'helm repo add'
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

### `D:\helm_clones_github\ethpandaops__whisk-devnets\kubernetes\devnet-0\blobscan` — ethpandaops/whisk-devnets

```
Error: no repository definition for https://charts.bitnami.com/bitnami, https://ethpandaops.github.io/ethereum-helm-charts, https://ethpandaops.github.io/ethereum-helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\galserg__kubetest\helmfile\charts\etcd` — galserg/kubetest

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\ishtiaqhimel__oms\charts\oms-server` — ishtiaqhimel/oms

```
Error: no repository definition for https://charts.konghq.com. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\lucidworks__ocp-fusion-helm-charts\5.3.4\fusion` — lucidworks/ocp-fusion-helm-charts

```
Error: no repository definition for https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com, https://charts.lucidworks.com. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\openshift-helm-charts__development\charts\partners\embedded\tokenvisor\0.1.0\src` — openshift-helm-charts/development

```
Error: no repository definition for https://victoriametrics.github.io/helm-charts/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\openshift-helm-charts__sandbox\charts\redhat\redhat\redhat-mysql-persistent\0.0.2\src` — openshift-helm-charts/sandbox

```
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "https://github.com/openshift-helm-charts/charts" chart repository:
	failed to fetch https://github.com/openshift-helm-charts/charts/index.yaml : 404 Not Found
Error: no cached repository for helm-manager-e8a50bb8edab0b7411123ba1d6ccffca210a3cb9ffb119170c254ad301b6b826 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-e8a50bb8edab0b7411123ba1d6ccffca210a3cb9ffb119170c254ad301b6b826-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### `D:\helm_clones_github\oracle-cne__catalog\charts\metallb-0.15.2` — oracle-cne/catalog

```
Error: no repository definition for https://metallb.github.io/frr-k8s. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\otterscale__charts\charts\otterscale` — otterscale/charts

```
Error: no repository definition for https://codecentric.github.io/helm-charts, https://valkey.io/valkey-helm/, https://helm.goharbor.io. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\platform-mesh__helm-charts\charts\common\test-chart` — platform-mesh/helm-charts

```
Saving 1 charts
Save error occurred:  can't get a valid version for dependency common
Error: can't get a valid version for dependency common
```

### `D:\helm_clones_github\sdelrio__homelab-k3s\system\internal\postgres-operator` — sdelrio/homelab-k3s

```
Getting updates for unmanaged Helm repositories...
...Successfully got an update from the "https://opensource.zalando.com/postgres-operator/charts/postgres-operator" chart repository
Error: can't get a valid version for 1 subchart(s): "postgres-operator" (repository "https://opensource.zalando.com/postgres-operator/charts/postgres-operator", version "1.10.0"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml
```

### `D:\helm_clones_github\tetratelabs__charts\charts\demos\istio-monitoring-demo` — tetratelabs/charts

```
Error: no repository definition for https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\thomaswyrick__xdm-helm-chart` — thomaswyrick/xdm-helm-chart

```
Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update'
```

### `D:\helm_clones_github\vineethvijay__prox-k8s-lab\helm\charts\nfs-hdd-provisioner` — vineethvijay/prox-k8s-lab

```
Error: no repository definition for https://kubernetes-sigs.github.io/nfs-subdir-external-provisioner/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\2694484453__helm-repo\kube-prometheus-stack` — 2694484453/helm-repo

```
Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\Chinaxiaoming666__kube-prometheus-stack` — Chinaxiaoming666/kube-prometheus-stack

```
Error: no repository definition for https://charts.helm.sh/stable, https://prometheus-community.github.io/helm-charts, https://grafana.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\DenisMarta__prodject1\deployment\kubernetes\mysql-helm-chart` — DenisMarta/prodject1

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\HonourHealth__JavaSpringbootRESTMicroservices\springboot-microservices-springcloud-docker-kubernetes-helm\helm\bank-services\accounts` — HonourHealth/JavaSpringbootRESTMicroservices

```
Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update'
```

### `D:\helm_clones_github\Kapil-Bhalodiya__E-Commerce\infra\addons\alloy` — Kapil-Bhalodiya/E-Commerce

```
Error: no repository definition for https://grafana.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\Kapil-Bhalodiya__E-commerce-Platform\infra\addons\cert-manager` — Kapil-Bhalodiya/E-commerce-Platform

```
Error: no repository definition for https://charts.jetstack.io. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\PRO-Robotech__helmfile-manifests\charts\argoproj\argo-cd-8.0.9\argo-cd` — PRO-Robotech/helmfile-manifests

```
Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update'
```

### `D:\helm_clones_github\PilotDataPlatform__helm-charts\argo-cd-917` — PilotDataPlatform/helm-charts

```
Error: no repository definition for https://dandydeveloper.github.io/charts/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\SmartX-Team__smartx-k8s\apps\exdns-2` — SmartX-Team/smartx-k8s

```
Error: cannot load Chart.yaml: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type v2.Metadata
```

### `D:\helm_clones_github\TSMC-NYCU-LAB-13__infrastructures\argo\argo-cd` — TSMC-NYCU-LAB-13/infrastructures

```
Error: no repository definition for https://dandydeveloper.github.io/charts/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\adstanley__archive\scale-catalog\incubator\archivebox\0.7.2` — adstanley/archive

```
Saving 1 charts
Downloading common from repo oci://tccr.io/truecharts
Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:17.2.30: not found
Error: could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:17.2.30: not found
```

### `D:\helm_clones_github\azaurus1__homelab\apps\api-coverage-server` — azaurus1/homelab

```
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "https://chartmuseum.azaurus.dev/" chart repository:
	failed to fetch https://chartmuseum.azaurus.dev/index.yaml : 530 <none>
Error: no cached repository for helm-manager-23f58962e8ae66ac9eadc10ae209dc5cab99e97e33882fa9df1e156284414aea found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-23f58962e8ae66ac9eadc10ae209dc5cab99e97e33882fa9df1e156284414aea-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### `D:\helm_clones_github\devops-360-online__go-with-me-ml\infrastructure\kubernetes\bento` — devops-360-online/go-with-me-ml

```
Error: no repository definition for https://warpstreamlabs.github.io/bento-helm-chart, https://warpstreamlabs.github.io/bento-helm-chart, https://warpstreamlabs.github.io/bento-helm-chart. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\dm3drummer__arducharts\configs\schema\ahrs` — dm3drummer/arducharts

```
Error: validation: chart.metadata.version "schema" is invalid
```

### `D:\helm_clones_github\epoch8__helm-charts\examples\multi-simple-app` — epoch8/helm-charts

```
Saving 4 charts
Save error occurred:  can't get a valid version for dependency simple-app
Error: can't get a valid version for dependency simple-app
```

### `D:\helm_clones_github\flashiam12__strategic-next-best-offer\dashboard\dependencies\loki` — flashiam12/strategic-next-best-offer

```
Error: no repository definition for https://charts.min.io/, https://grafana.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\hey101__scale-catalog\incubator\archivebox\0.7.2` — hey101/scale-catalog

```
Saving 1 charts
Downloading common from repo oci://tccr.io/truecharts
Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:17.5.1: not found
Error: could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:17.5.1: not found
```

### `D:\helm_clones_github\highmobility__helm-charts\examples\sample-app\helm-web` — highmobility/helm-charts

```
Saving 1 charts
Save error occurred:  can't get a valid version for dependency hm-basic-webapp
Error: can't get a valid version for dependency hm-basic-webapp
```

### `D:\helm_clones_github\hmcts__hmcts-charts\stable\aac-manage-case-assignment` — hmcts/hmcts-charts

```
Getting updates for unmanaged Helm repositories...
...Successfully got an update from the "https://helm.elastic.co" chart repository
Saving 4 charts
Downloading java from repo oci://hmctsprod.azurecr.io/helm
Save error occurred:  could not download oci://hmctsprod.azurecr.io/helm/java: failed to perform "FetchReference" on source: GET "https://hmctsprod.azurecr.io/v2/helm/java/manifests/5.3.0": GET "https://hmctsprod.azurecr.io/oauth2/token?scope=repository%3Ahelm%2Fjava%3Apull&service=hmctsprod.azurecr.io": response status code 401: unauthorized: authentication required, visit https://aka.ms/acr/authorization for more information. CorrelationId: 46eab12f-4b7c-42e6-af30-f7c8865fbb3b
Error: could not download oci://hmctsprod.azurecr.io/helm/java: failed to perform "FetchReference" on source: GET "https://hmctsprod.azurecr.io/v2/helm/java/manifests/5.3.0": GET "https://hmctsprod.azurecr.io/oauth2/token?scope=repository%3Ahelm%2Fjava%3Apull&service=hmctsprod.azurecr.io": response status code 401: unauthorized: authentication required, visit https://aka.ms/acr/authorization for more information. CorrelationId: 46eab12f-4b7c-42e6-af30-f7c8865fbb3b
```

### `D:\helm_clones_github\hmcts__probate-frontend\charts\probate-frontend` — hmcts/probate-frontend

```
Saving 2 charts
Downloading nodejs from repo oci://hmctsprod.azurecr.io/helm
Save error occurred:  could not download oci://hmctsprod.azurecr.io/helm/nodejs: failed to perform "FetchReference" on source: GET "https://hmctsprod.azurecr.io/v2/helm/nodejs/manifests/3.2.1": GET "https://hmctsprod.azurecr.io/oauth2/token?scope=repository%3Ahelm%2Fnodejs%3Apull&service=hmctsprod.azurecr.io": response status code 401: unauthorized: authentication required, visit https://aka.ms/acr/authorization for more information. CorrelationId: 6316736e-9d81-4ccb-8adb-3c3e0792e1eb
Error: could not download oci://hmctsprod.azurecr.io/helm/nodejs: failed to perform "FetchReference" on source: GET "https://hmctsprod.azurecr.io/v2/helm/nodejs/manifests/3.2.1": GET "https://hmctsprod.azurecr.io/oauth2/token?scope=repository%3Ahelm%2Fnodejs%3Apull&service=hmctsprod.azurecr.io": response status code 401: unauthorized: authentication required, visit https://aka.ms/acr/authorization for more information. CorrelationId: 6316736e-9d81-4ccb-8adb-3c3e0792e1eb
```

### `D:\helm_clones_github\iamkhattar__homelab\cluster\core\cert-manager` — iamkhattar/homelab

```
Error: no repository definition for https://charts.jetstack.io. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\joctan-tec__2023-02-2021069671-IC4302\TareaCorta1\HelmCharts\databases` — joctan-tec/2023-02-2021069671-IC4302

```
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\lorenzofratini1998__moneyverse\moneyverse-infrastructure\moneyverse-chart\apps\clickhouse` — lorenzofratini1998/moneyverse

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\matic-insurance__helm-charts\ops\workbench\app-deployment` — matic-insurance/helm-charts

```
Saving 5 charts
Save error occurred:  can't get a valid version for dependency application-component
Error: can't get a valid version for dependency application-component
```

### `D:\helm_clones_github\neuroxhq__helm-chart-neurox-workload` — neuroxhq/helm-chart-neurox-workload

```
Saving 2 charts
Downloading neurox-workload-agent from repo oci://ghcr.io/neuroxhq/helm-charts
Save error occurred:  could not download oci://ghcr.io/neuroxhq/helm-charts/neurox-workload-agent: failed to perform "FetchReference" on source: GET "https://ghcr.io/v2/neuroxhq/helm-charts/neurox-workload-agent/manifests/2.10.2": GET "https://ghcr.io/token?scope=repository%3Aneuroxhq%2Fhelm-charts%2Fneurox-workload-agent%3Apull&service=ghcr.io": response status code 401: unauthorized: authentication required
Error: could not download oci://ghcr.io/neuroxhq/helm-charts/neurox-workload-agent: failed to perform "FetchReference" on source: GET "https://ghcr.io/v2/neuroxhq/helm-charts/neurox-workload-agent/manifests/2.10.2": GET "https://ghcr.io/token?scope=repository%3Aneuroxhq%2Fhelm-charts%2Fneurox-workload-agent%3Apull&service=ghcr.io": response status code 401: unauthorized: authentication required
```

### `D:\helm_clones_github\osherlevi7__dev\dev-tools\helm-tools\litmus` — osherlevi7/dev

```
Error: no repository definition for https://raw.githubusercontent.com/bitnami/charts/archive-full-index/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\piliphulko__marketplace-example\k8s\cores-api` — piliphulko/marketplace-example

```
Error: error unpacking subchart Makefile in core-api: Chart.yaml file is missing
```

### `D:\helm_clones_github\roman009__go-rust\infrastructure\infra\kafka` — roman009/go-rust

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\sujan060__testing111\packages\helm-charts\attestation-service` — sujan060/testing111

```
Error: no repository definition for @stable. Please add them via 'helm repo add'
```

### `D:\helm_clones_github\tetratelabs__helm-charts\charts\demos\istio-monitoring-demo` — tetratelabs/helm-charts

```
Error: no repository definition for https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\wiredquill__ai-demos\charts\hr-assistant` — wiredquill/ai-demos

```
Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update'
```

### `D:\helm_clones_github\Ammar-Taimoori__speechmatics-helm-chart\1.2.0\sm-realtime` — Ammar-Taimoori/speechmatics-helm-chart

```
Error: directory D:\helm_clones_github\Ammar-Taimoori__speechmatics-helm-chart\1.2.0\sm-proxy not found
```

### `D:\helm_clones_github\ArieLevs__Kubernetes-Helm-Charts\charts\nalkinscloud` — ArieLevs/Kubernetes-Helm-Charts

```
Error: no repository definition for https://charts.bitnami.com/bitnami, https://helm.releases.hashicorp.com, https://arielevs.github.io/Kubernetes-Helm-Charts/, https://arielevs.github.io/Kubernetes-Helm-Charts/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\Arthur-B-DevOps__old_helm_charts\charts\Old_charts\charts\incubator\distribution` — Arthur-B-DevOps/old_helm_charts

```
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/, https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\ArvigEnterprises__integration-tests-bluegreen\.helm\charts\arvigtestcluster334d\canary` — ArvigEnterprises/integration-tests-bluegreen

```
Error: directory D:\helm_clones_github\ArvigEnterprises__integration-tests-bluegreen\.helm\charts\arvigtestcluster334d\laravel-library not found
```

### `D:\helm_clones_github\AssemblerBossss__tools_and_manuals\kubernetes\18.templating\wordpress` — AssemblerBossss/tools_and_manuals

```
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\BYGX-wcr__Magma_NB-IoT\cn\deploy\helm` — BYGX-wcr/Magma_NB-IoT

```
Error: chart file "virtctl" is larger than the maximum file size 5242880
```

### `D:\helm_clones_github\Blastz13__quality_reducer\k8s\kube-prometheus-stack` — Blastz13/quality_reducer

```
Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\CERIT-SC__rancher-apps-secure\charts\monitoring\v0.1` — CERIT-SC/rancher-apps-secure

```
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies
```

### `D:\helm_clones_github\Chalupa-Tech__chalupa-tech-local\gitops\apps\media\clonarr` — Chalupa-Tech/chalupa-tech-local

```
Error: no repository definition for https://bjw-s-labs.github.io/helm-charts/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\CloudLargeScale-UCLouvain__c2b2\caliper` — CloudLargeScale-UCLouvain/c2b2

```
Error: no repository definition for https://halkeye.github.io/helm-charts/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\DEL-ORG__s7mensah\argo-helm\charts\argo-cd` — DEL-ORG/s7mensah

```
Error: no repository definition for https://dandydeveloper.github.io/charts/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\DakEnviy__my-terraform\internal\cluster\helm\gitlab` — DakEnviy/my-terraform

```
Error: no repository definition for https://charts.jetstack.io/, https://prometheus-community.github.io/helm-charts, https://charts.bitnami.com/bitnami, https://charts.gitlab.io/, https://charts.bitnami.com/bitnami, https://charts.gitlab.io/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\Dystopian00__truecharts\charts\incubator\docuseal` — Dystopian00/truecharts

```
Saving 1 charts
Downloading common from repo oci://tccr.io/truecharts
Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:24.1.5: not found
Error: could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:24.1.5: not found
```

### `D:\helm_clones_github\Feederhigh5__master-thesis-poc\litmus` — Feederhigh5/master-thesis-poc

```
Error: no repository definition for https://raw.githubusercontent.com/bitnami/charts/archive-full-index/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\Genocs__genocs-argocd\helm-dependency` — Genocs/genocs-argocd

```
Error: validation: chart.metadata.version is required
```

### `D:\helm_clones_github\GovStackWorkingGroup__sandbox-bb-identity\Esignet Deployment\Esignet Service\helm\oidc-ui` — GovStackWorkingGroup/sandbox-bb-identity

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\GroomC4__c4ang-infra\charts\schema-registry` — GroomC4/c4ang-infra

```
Getting updates for unmanaged Helm repositories...
...Successfully got an update from the "https://confluentinc.github.io/cp-helm-charts/" chart repository
Error: cp-schema-registry chart not found in repo https://confluentinc.github.io/cp-helm-charts/
```

### `D:\helm_clones_github\HelixDevelopment__HelixGitpx\docs\specifications\main\main_implementation_material\HelixGitpx\18-manifests` — HelixDevelopment/HelixGitpx

```
Getting updates for unmanaged Helm repositories...
...Successfully got an update from the "https://cloudnative-pg.io/charts" chart repository
...Successfully got an update from the "https://go.temporal.io/helm-charts" chart repository
...Successfully got an update from the "https://qdrant.github.io/qdrant-helm" chart repository
...Successfully got an update from the "https://meilisearch.github.io/meilisearch-kubernetes" chart repository
...Successfully got an update from the "https://opensearch-project.github.io/helm-charts" chart repository
...Successfully got an update from the "https://spiffe.github.io/helm-charts-hardened" chart repository
...Successfully got an update from the "https://kyverno.github.io/kyverno" chart repository
...Successfully got an update from the "https://strimzi.io/charts" chart repository
...Successfully got an update from the "https://grafana.github.io/helm-charts" chart repository
...Successfully got an update from the "https://prometheus-community.github.io/helm-charts" chart repository
...Successfully got an update from the "https://istio-release.storage.googleapis.com/charts" chart repository
Error: istio-base chart not found in repo https://istio-release.storage.googleapis.com/charts
```

### `D:\helm_clones_github\ICICLE-ai__OpenPass\helm-config` — ICICLE-ai/OpenPass

```
Error: cannot load values.yaml: cannot unmarshal yaml document: error converting YAML to JSON: yaml: line 30: could not find expected ':'
```

### `D:\helm_clones_github\Jaebytes__TrueCharts\premium\traefik\26.10.19` — Jaebytes/TrueCharts

```
Saving 1 charts
Downloading common from repo oci://tccr.io/truecharts
Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:20.3.12: not found
Error: could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:20.3.12: not found
```

### `D:\helm_clones_github\Keshari07__Kubernetes_28_29_September_2024\kubelabs\Helm101\Wordpress` — Keshari07/Kubernetes_28_29_September_2024

```
Error: cannot load Chart.yaml: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go struct field Metadata.deprecated of type bool
```

### `D:\helm_clones_github\LTKH__minikube\argocd\charts\argo-cd` — LTKH/minikube

```
Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update'
```

### `D:\helm_clones_github\LeeLisker__helm\community\artifactory-ha` — LeeLisker/helm

```
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\MUST-CLOUD__osh\openstack-helm\openstack` — MUST-CLOUD/osh

```
Error: error unpacking subchart rabbitmq in openstack: Chart.yaml file is missing
```

### `D:\helm_clones_github\Mainuddin-Rizvi__devops\kubernetes-k8s\12-deploying-to-multiple-environments\kluctl\.helm-charts\https_cloudnative-pg.github.io\charts\cloudnative-pg\0.21.4` — Mainuddin-Rizvi/devops

```
Error: no repository definition for https://cloudnative-pg.github.io/grafana-dashboards. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\Mark4562__Truenas-Scale\premium\authelia\23.13.13` — Mark4562/Truenas-Scale

```
Saving 2 charts
Downloading common from repo oci://tccr.io/truecharts
Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:20.3.11: not found
Error: could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:20.3.11: not found
```

### `D:\helm_clones_github\Musaiba__Research-investigation-tools\theHarvester\deploy\charts\harvester` — Musaiba/Research-investigation-tools

```
Error: no repository definition for https://charts.harvesterhci.io, https://charts.harvesterhci.io, https://charts.longhorn.io, https://kube-vip.github.io/helm-charts, https://kube-vip.github.io/helm-charts, https://charts.harvesterhci.io, https://charts.harvesterhci.io. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\MxBlu__tcharts\charts\home-assistant\1.0.130` — MxBlu/tcharts

```
Error: directory D:\helm_clones_github\MxBlu__tcharts\common\2304.0.1 not found
```

### `D:\helm_clones_github\My-DIGI-ID__ssi-helm-charts\charts\ssi-aca-py` — My-DIGI-ID/ssi-helm-charts

```
Error: no repository definition for https://charts.bitnami.com/bitnami, https://my-digi-id.github.io/ssi-helm-charts/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\OpenSourceConsulting__playcekube-charts\src\playcekube\keycloak\9.2.2\keycloak` — OpenSourceConsulting/playcekube-charts

```
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\OriHoch__sk8s-ops` — OriHoch/sk8s-ops

```
Error: validation: chart.metadata.version is required
```

### `D:\helm_clones_github\Otus-DevOps-2022-02__saintjb_microservices\kubernetes\Charts\gitlab` — Otus-DevOps-2022-02/saintjb_microservices

```
Getting updates for unmanaged Helm repositories...
...Successfully got an update from the "https://charts.jetstack.io/" chart repository
...Successfully got an update from the "https://grafana.github.io/helm-charts" chart repository
...Successfully got an update from the "https://prometheus-community.github.io/helm-charts" chart repository
...Successfully got an update from the "https://charts.gitlab.io/" chart repository
...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository
Error: can't get a valid version for 2 subchart(s): "postgresql" (repository "https://charts.bitnami.com/bitnami", version "8.9.4"), "redis" (repository "https://charts.bitnami.com/bitnami", version "11.3.4"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml
```

### `D:\helm_clones_github\PaulJequann__the-lab\kubernetes\apps\glitchtip` — PaulJequann/the-lab

```
Error: no repository definition for https://gitlab.com/api/v4/projects/16325141/packages/helm/stable. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\Riskin1999__arena-appwrapper-volcano\arena-artifacts` — Riskin1999/arena-appwrapper-volcano

```
Error: no repository definition for @tf-operator, @tf-dashbard, @cron-operator, @et-operator, @mpi-operator, @pytorch-operator, @gpu-exporter, @elastic-job-supervisor. Please add them via 'helm repo add'
```

### `D:\helm_clones_github\SalesConnection__sc-helm-charts\charts\caction` — SalesConnection/sc-helm-charts

```
Error: no repository definition for https://salesconnection.github.io/sc-helm-charts, https://salesconnection.github.io/sc-helm-charts, https://salesconnection.github.io/sc-helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\ShaharyarShakir__devops-practice\16__gitops\projects\vprofile-action\helm\vprofilecharts` — ShaharyarShakir/devops-practice

```
Error: error unpacking subchart templates in vprofilecharts: Chart.yaml file is missing
```

### `D:\helm_clones_github\SirObi__airflow_k8s_executor\helm_charts\official\charts\incubator\distribution` — SirObi/airflow_k8s_executor

```
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/, https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\StopDenBus__helm-charts\charts\external-dns` — StopDenBus/helm-charts

```
Saving 1 charts
Downloading external-dns-chart from repo oci://dhi.io
Save error occurred:  could not download oci://dhi.io/external-dns-chart: failed to perform "FetchReference" on source: GET "https://dhi.io/v2/external-dns-chart/manifests/1.21.1": GET "https://dhi.io/token?scope=repository%3Aexternal-dns-chart%3Apull&service=registry.docker.io": response status code 401: Unauthorized
Error: could not download oci://dhi.io/external-dns-chart: failed to perform "FetchReference" on source: GET "https://dhi.io/v2/external-dns-chart/manifests/1.21.1": GET "https://dhi.io/token?scope=repository%3Aexternal-dns-chart%3Apull&service=registry.docker.io": response status code 401: Unauthorized
```

### `D:\helm_clones_github\SuperCoolAlan__homelab-manifests\talos\authentik\charts\authentik-2026.2.1` — SuperCoolAlan/homelab-manifests

```
Error: no repository definition for https://charts.goauthentik.io. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\Timkhch__school-dev-k8s\practice\18.templating\wordpress` — Timkhch/school-dev-k8s

```
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\TrevorSquillario__iDRAC-Telemetry-Ansible-Demo\k8s\helm\app` — TrevorSquillario/iDRAC-Telemetry-Ansible-Demo

```
Error: no repository definition for https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\TrivialJavaBackender__infrastructure-learn\exercises\helm\Ex15_BasicChart\myapp` — TrivialJavaBackender/infrastructure-learn

```
Error: validation: chart.metadata.version is required
```

### `D:\helm_clones_github\Yuzu815__TruenasScaleArchive\incubator\archivebox\0.7.2` — Yuzu815/TruenasScaleArchive

```
Saving 1 charts
Downloading common from repo oci://tccr.io/truecharts
Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:17.2.30: not found
Error: could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:17.2.30: not found
```

### `D:\helm_clones_github\aashkulkarni__helm\schoolapp-subchart` — aashkulkarni/helm

```
Error: no repository definition for https://gitlab.com/api/v4/projects/34240616/packages/helm/stable, https://gitlab.com/api/v4/projects/34240616/packages/helm/stable, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\abhilashindulkar__misc\helm\custom-charts\couponservice` — abhilashindulkar/misc

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\abimaelrsergio__Abimoney\charts-main\bitnami\sealed-secrets` — abimaelrsergio/Abimoney

```
Saving 1 charts
Downloading common from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/common: failed to perform "FetchReference" on source: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.31.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/common: failed to perform "FetchReference" on source: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.31.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `D:\helm_clones_github\acaisoft__bolt-platform-helm` — acaisoft/bolt-platform-helm

```
Error: no repository definition for http://bolt-helm-charts.acaisoft.net, http://bolt-helm-charts.acaisoft.net, https://kubernetes.github.io/ingress-nginx, https://argoproj.github.io/argo-helm, https://hasura-extra.github.io/hasura-extra. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\actions-marketplace-validations__LinuxSuRen_api-testing\helm\api-testing` — actions-marketplace-validations/LinuxSuRen_api-testing

```
Saving 1 charts
Downloading mongodb from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/mongodb: failed to perform "FetchReference" on source: GET "https://registry-1.docker.io/v2/bitnamicharts/mongodb/manifests/15.1.7": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/mongodb: failed to perform "FetchReference" on source: GET "https://registry-1.docker.io/v2/bitnamicharts/mongodb/manifests/15.1.7": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `D:\helm_clones_github\aicanhelp__ai-cloud\bitnami\charts\bitnami-202503\airflow` — aicanhelp/ai-cloud

```
Saving 3 charts
Downloading redis from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to perform "FetchReference" on source: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/20.7.1": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to perform "FetchReference" on source: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/20.7.1": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `D:\helm_clones_github\albinoJimy__Neural-Hive-Mind\helm\istio-base` — albinoJimy/Neural-Hive-Mind

```
Error: no repository definition for https://istio-release.storage.googleapis.com/charts, https://istio-release.storage.googleapis.com/charts, https://istio-release.storage.googleapis.com/charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\alnav3__argocd-charts\charts\argo-cd` — alnav3/argocd-charts

```
Error: no repository definition for https://argoproj.github.io/argo-helm. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\andgit7__repo2\Kubernetes\slurm\practice\18.templating\wordpress` — andgit7/repo2

```
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\anmol372__matrix-chart` — anmol372/matrix-chart

```
Error: no repository definition for https://kubernetes-charts.storage.googleapis.com, https://vmware-tanzu.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\anzerr__catalog\dependency\mariadb\12.2.0` — anzerr/catalog

```
Saving 1 charts
Downloading common from repo oci://tccr.io/truecharts
Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:17.4.2: not found
Error: could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:17.4.2: not found
```

### `D:\helm_clones_github\artichoked1__FerriteCMS-chart` — artichoked1/FerriteCMS-chart

```
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\arychj__argocd\wordpress` — arychj/argocd

```
Error: validation: chart.metadata.version is required
```

### `D:\helm_clones_github\beantownpub__helm\argo-cd` — beantownpub/helm

```
Error: no repository definition for https://argoproj.github.io/argo-helm. Please add the missing repos via 'helm repo add'
```

### `D:\helm_clones_github\bgarcia-elastic__tfk8` — bgarcia-elastic/tfk8

```
Error: directory D:\helm_clones_github\subCharts\ui not found
```

