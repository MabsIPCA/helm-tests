# Helm Dependency Build Failures

Total dependency failures: **106**

| # | Repository | Chart Path | Error |
|---|------------|------------|-------|
| 1 | [ConductionNL/procces-engine](https://github.com/ConductionNL/procces-engine) | `cloned\ConductionNL__procces-engine\api\helm` | manager.go:122: warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash... Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies |
| 2 | [pluralsh/plural-helm-charts](https://github.com/pluralsh/plural-helm-charts) | `cloned\pluralsh__plural-helm-charts\charts\airbyte` | Error: no repository definition for https://charts.bitnami.com/bitnami, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts. Please add the missing repos via 'helm repo add' |
| 3 | [michaelw/common.itsumi](https://github.com/michaelw/common.itsumi) | `cloned\michaelw__common.itsumi` | Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 1 charts Downloading common from repo oci://registry-1.docker.io/bitnamicharts Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/common: failed to resolve registry-1.docker.io/bitnamicharts/common:2.31.3: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.31.3": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/common: failed to resolve registry-1.docker.io/bitnamicharts/common:2.31.3: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.31.3": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 4 | [solo-io/skv2](https://github.com/solo-io/skv2) | `cloned\solo-io__skv2\codegen\test\chart\conditional-crds` | Error: validation: chart.metadata.name is required |
| 5 | [flytrap/docker-admin](https://github.com/flytrap/docker-admin) | `cloned\flytrap__docker-admin\charts\apisix-1.3` | Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.apiseven.com, https://charts.apiseven.com. Please add the missing repos via 'helm repo add' |
| 6 | [akv-akv/de_research](https://github.com/akv-akv/de_research) | `cloned\akv-akv__de_research\kubernetes\charts\airbyte` | Error: no repository definition for https://charts.bitnami.com/bitnami, https://airbytehq.github.io/helm-charts/, https://airbytehq.github.io/helm-charts/, https://airbytehq.github.io/helm-charts/, https://airbytehq.github.io/helm-charts/, https://airbytehq.github.io/helm-charts/, https://airbytehq.github.io/helm-charts/, https://airbytehq.github.io/helm-charts/, https://airbytehq.github.io/helm-charts/, https://airbytehq.github.io/helm-charts/, https://airbytehq.github.io/helm-charts/, https://airbytehq.github.io/helm-charts/, https://airbytehq.github.io/helm-charts/, https://airbytehq.github.io/helm-charts/. Please add the missing repos via 'helm repo add' |
| 7 | [ConductionNL/Authorization-component](https://github.com/ConductionNL/Authorization-component) | `cloned\ConductionNL__Authorization-component\api\helm` | load.go:138: Warning: Dependency locking is handled in Chart.lock since apiVersion "v2". We recommend migrating to Chart.lock. Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies |
| 8 | [MirahImage/postfacto](https://github.com/MirahImage/postfacto) | `cloned\MirahImage__postfacto\tkg\postfacto` | Error: no repository definition for @bitnami, @bitnami. Please add them via 'helm repo add' |
| 9 | [noygal/helm](https://github.com/noygal/helm) | `cloned\noygal__helm\charts\example-dev-tools` | Getting updates for unmanaged Helm repositories... ...Unable to get an update from the "https://noygal.github.com/helm" chart repository: 	Get "https://noygal.github.com/helm/index.yaml": dial tcp: lookup noygal.github.com: no such host Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Error: no cached repository for helm-manager-19cefd96797d2c0b3ed74271dcee4ae7ad7d02cb4d4e47a98431938d24b11530 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-19cefd96797d2c0b3ed74271dcee4ae7ad7d02cb4d4e47a98431938d24b11530-index.yaml: O sistema não conseguiu localizar o ficheiro especificado. |
| 10 | [manics/zero-to-jupyterhub-k8s-chart-expanded](https://github.com/manics/zero-to-jupyterhub-k8s-chart-expanded) | `cloned\manics__zero-to-jupyterhub-k8s-chart-expanded\0.8.1` | manager.go:122: warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash... Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies |
| 11 | [unicef/magasin](https://github.com/unicef/magasin) | `cloned\unicef__magasin\helm\dagster` | Error: no repository definition for https://raw.githubusercontent.com/bitnami/charts/eb5f9a9513d987b519f0ecd732e7031241c50328/bitnami, https://raw.githubusercontent.com/bitnami/charts/eb5f9a9513d987b519f0ecd732e7031241c50328/bitnami, https://raw.githubusercontent.com/bitnami/charts/eb5f9a9513d987b519f0ecd732e7031241c50328/bitnami. Please add the missing repos via 'helm repo add' |
| 12 | [ConductionNL/berichtservice](https://github.com/ConductionNL/berichtservice) | `cloned\ConductionNL__berichtservice\api\helm` | load.go:138: Warning: Dependency locking is handled in Chart.lock since apiVersion "v2". We recommend migrating to Chart.lock. Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 13 | [ConductionNL/commonground-ui](https://github.com/ConductionNL/commonground-ui) | `cloned\ConductionNL__commonground-ui\api\helm` | manager.go:122: warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash... Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies |
| 14 | [Veronneau-Techno-Conseil/comax-commons](https://github.com/Veronneau-Techno-Conseil/comax-commons) | `cloned\Veronneau-Techno-Conseil__comax-commons\helm_referee` | Error: no repository definition for https://mongodb.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 15 | [waTeim/mcp-base](https://github.com/waTeim/mcp-base) | `cloned\waTeim__mcp-base\chart` | Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 1 charts Downloading redis from repo oci://registry-1.docker.io/cloudpirates Save error occurred:  could not download oci://registry-1.docker.io/cloudpirates/redis: failed to resolve registry-1.docker.io/cloudpirates/redis:0.16.4: GET "https://registry-1.docker.io/v2/cloudpirates/redis/manifests/0.16.4": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/cloudpirates/redis: failed to resolve registry-1.docker.io/cloudpirates/redis:0.16.4: GET "https://registry-1.docker.io/v2/cloudpirates/redis/manifests/0.16.4": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 16 | [datawizz/Firestream](https://github.com/datawizz/Firestream) | `cloned\datawizz__Firestream\src\charts\bitnami\airflow` | Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 3 charts Downloading redis from repo oci://registry-1.docker.io/bitnamicharts Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:21.2.5: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/21.2.5": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:21.2.5: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/21.2.5": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 17 | [limagbz/data-mesh-yelp](https://github.com/limagbz/data-mesh-yelp) | `cloned\limagbz__data-mesh-yelp\infra\monitoring\loki\helm` | Error: no repository definition for https://charts.min.io/, https://grafana.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 18 | [efremov-it/security](https://github.com/efremov-it/security) | `cloned\efremov-it__security\practice\weeks\2-3\secureCodeBox\operator` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 19 | [ConductionNL/zaak-registratie-component](https://github.com/ConductionNL/zaak-registratie-component) | `cloned\ConductionNL__zaak-registratie-component\api\helm` | manager.go:122: warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash... Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies |
| 20 | [ConductionNL/checkinservice](https://github.com/ConductionNL/checkinservice) | `cloned\ConductionNL__checkinservice\api\helm` | Getting updates for unmanaged Helm repositories... ...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Error: can't get a valid version for 1 subchart(s): "postgresql" (repository "https://charts.bitnami.com/bitnami", version "10.1.1"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml |
| 21 | [ConductionNL/digispoof-old](https://github.com/ConductionNL/digispoof-old) | `cloned\ConductionNL__digispoof-old\api\helm` | manager.go:122: warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash... Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies |
| 22 | [ConductionNL/proto-application-commonground](https://github.com/ConductionNL/proto-application-commonground) | `cloned\ConductionNL__proto-application-commonground\api\helm` | manager.go:122: warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash... Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies |
| 23 | [bitnami/charts](https://github.com/bitnami/charts) | `cloned\bitnami__charts\bitnami\airflow` | Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 3 charts Downloading redis from repo oci://registry-1.docker.io/bitnamicharts Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:22.0.4: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/22.0.4": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:22.0.4: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/22.0.4": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 24 | [grafana/helm-charts](https://github.com/grafana/helm-charts) | `cloned\grafana__helm-charts\charts\enterprise-metrics` | Getting updates for unmanaged Helm repositories... ...Successfully got an update from the "https://helm.min.io/" chart repository ...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Error: minio chart not found in repo https://helm.min.io/ |
| 25 | [cloudnativeapp/charts](https://github.com/cloudnativeapp/charts) | `cloned\cloudnativeapp__charts\curated\airflow` | manager.go:122: warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash... Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies |
| 26 | [krateoplatformops/composition-portal-starter](https://github.com/krateoplatformops/composition-portal-starter) | `cloned\krateoplatformops__composition-portal-starter\chart` | Error: validation: chart.metadata.version "CHART_VERSION" is invalid |
| 27 | [samcab28/Diseno-Software](https://github.com/samcab28/Diseno-Software) | `cloned\samcab28__Diseno-Software\proyecto\caso4\charts\databases` | Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies |
| 28 | [rancher/charts](https://github.com/rancher/charts) | `cloned\rancher__charts\charts\epinio\102.0.1+up1.6.2` | Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies |
| 29 | [gersonrs/hand-on-airflow-with-minio](https://github.com/gersonrs/hand-on-airflow-with-minio) | `cloned\gersonrs__hand-on-airflow-with-minio\airflow` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 30 | [musqat/lectures](https://github.com/musqat/lectures) | `cloned\musqat__lectures\MicroService\section_16\helm\grafana` | Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 1 charts Downloading common from repo oci://registry-1.docker.io/bitnamicharts Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/common: failed to resolve registry-1.docker.io/bitnamicharts/common:2.31.0: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.31.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/common: failed to resolve registry-1.docker.io/bitnamicharts/common:2.31.0: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.31.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 31 | [general-rj45/mini-ml-stand-GRJ45](https://github.com/general-rj45/mini-ml-stand-GRJ45) | `cloned\general-rj45__mini-ml-stand-GRJ45\airflow-14.2.5\airflow` | Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 3 charts Downloading redis from repo oci://registry-1.docker.io/bitnamicharts Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:17.11.3: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/17.11.3": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:17.11.3: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/17.11.3": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 32 | [shini4i/charts](https://github.com/shini4i/charts) | `cloned\shini4i__charts\charts\app` | Error: no repository definition for https://shini4i.github.io/charts/. Please add the missing repos via 'helm repo add' |
| 33 | [otus-kuber-2020-07/israodin_platform](https://github.com/otus-kuber-2020-07/israodin_platform) | `cloned\otus-kuber-2020-07__israodin_platform\kubernetes-templating\hipster-shop` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 34 | [robertbenzinger/charts](https://github.com/robertbenzinger/charts) | `cloned\robertbenzinger__charts\janusgraph` | Getting updates for unmanaged Helm repositories... ...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository ...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Error: can't get a valid version for 2 subchart(s): "cassandra" (repository "https://charts.bitnami.com/bitnami", version "5.6.1"), "elasticsearch" (repository "https://charts.bitnami.com/bitnami", version "12.6.2"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml |
| 35 | [openlit/helm](https://github.com/openlit/helm) | `cloned\openlit__helm\charts\openlit` | Error: no repository definition for https://openlit.github.io/helm/. Please add the missing repos via 'helm repo add' |
| 36 | [marcus-sa/helm-charts](https://github.com/marcus-sa/helm-charts) | `cloned\marcus-sa__helm-charts\repositories\banzaicloud-stable\anchore-engine\0.13.0\chart` | manager.go:122: warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash... Error: no repository definition for https://kubernetes-charts.storage.googleapis.com, https://kubernetes-charts.storage.googleapis.com, https://kubernetes-charts.storage.googleapis.com. Please add the missing repos via 'helm repo add' |
| 37 | [hey101/scale-catalog](https://github.com/hey101/scale-catalog) | `cloned\hey101__scale-catalog\incubator\archivebox\0.7.2` | Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 1 charts Downloading common from repo oci://tccr.io/truecharts Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to resolve tccr.io/truecharts/common:17.5.1: Get "https://tccr.io/v2/truecharts/common/manifests/17.5.1": dial tcp: lookup tccr.io: no such host Error: could not download oci://tccr.io/truecharts/common: failed to resolve tccr.io/truecharts/common:17.5.1: Get "https://tccr.io/v2/truecharts/common/manifests/17.5.1": dial tcp: lookup tccr.io: no such host |
| 38 | [unity-systems/spring-demo](https://github.com/unity-systems/spring-demo) | `cloned\unity-systems__spring-demo\charts\preview` | Getting updates for unmanaged Helm repositories... ...Unable to get an update from the "https://chartmuseum.build.cd.jenkins-x.io" chart repository: 	Get "https://chartmuseum.build.cd.jenkins-x.io/index.yaml": dial tcp 35.205.186.189:443: connectex: Uma tentativa de ligação falhou porque o componente ligado não respondeu corretamente após um período de tempo, ou a ligação estabelecida falhou porque o anfitrião ligado não respondeu. Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Error: no cached repository for helm-manager-e0b822155b4ae5f667d1ffc8875a97176819d25e9b8afb113dc2a031f0a96e31 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-e0b822155b4ae5f667d1ffc8875a97176819d25e9b8afb113dc2a031f0a96e31-index.yaml: O sistema não conseguiu localizar o ficheiro especificado. |
| 39 | [deresegetachew/playing-helm](https://github.com/deresegetachew/playing-helm) | `cloned\deresegetachew__playing-helm\kanban-backend` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 40 | [nadav07/helmTesting](https://github.com/nadav07/helmTesting) | `cloned\nadav07__helmTesting\chartTest` | Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies |
| 41 | [ibuildthecloud/rancher-charts](https://github.com/ibuildthecloud/rancher-charts) | `cloned\ibuildthecloud__rancher-charts\charts\anchore-engine\0.1.0` | manager.go:122: warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash... Error: no repository definition for https://kubernetes-charts.storage.googleapis.com. Please add the missing repos via 'helm repo add' |
| 42 | [czerwonk/hakanai](https://github.com/czerwonk/hakanai) | `cloned\czerwonk__hakanai\helm\hakanai` | Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 1 charts Downloading redis from repo oci://registry-1.docker.io/bitnamicharts Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:20.5.0: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/20.5.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:20.5.0: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/20.5.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 43 | [Abinesh0206/gitlab-ce](https://github.com/Abinesh0206/gitlab-ce) | `cloned\Abinesh0206__gitlab-ce` | Error: no repository definition for https://charts.jetstack.io/, https://prometheus-community.github.io/helm-charts, https://charts.bitnami.com/bitnami, https://charts.gitlab.io/, https://charts.bitnami.com/bitnami, https://charts.gitlab.io/, https://helm.traefik.io/traefik, https://haproxytech.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 44 | [samfms/truecharts](https://github.com/samfms/truecharts) | `cloned\samfms__truecharts\deprecated\bitwarden\1.2.5` | Error: no repository definition for https://truecharts.org/, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 45 | [Blankeeir/demo](https://github.com/Blankeeir/demo) | `cloned\Blankeeir__demo\infrastructure\helm\rafiki` | Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://interledger.github.io/helm-charts, https://interledger.github.io/helm-charts, https://interledger.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 46 | [randoli/helm-charts](https://github.com/randoli/helm-charts) | `cloned\randoli__helm-charts\charts\loki` | Error: no repository definition for https://charts.min.io/, https://grafana.github.io/helm-charts, https://grafana.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 47 | [dylanjustice/the-grid](https://github.com/dylanjustice/the-grid) | `cloned\dylanjustice__the-grid\gitops\workloads\kube-prometheus` | Error: no repository definition for https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 48 | [Mark4562/Truenas-Scale](https://github.com/Mark4562/Truenas-Scale) | `cloned\Mark4562__Truenas-Scale\premium\authelia\23.13.13` | Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 2 charts Downloading common from repo oci://tccr.io/truecharts Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to resolve tccr.io/truecharts/common:20.3.11: Get "https://tccr.io/v2/truecharts/common/manifests/20.3.11": dial tcp: lookup tccr.io: no such host Error: could not download oci://tccr.io/truecharts/common: failed to resolve tccr.io/truecharts/common:20.3.11: Get "https://tccr.io/v2/truecharts/common/manifests/20.3.11": dial tcp: lookup tccr.io: no such host |
| 49 | [praveenperera/matrix-element](https://github.com/praveenperera/matrix-element) | `cloned\praveenperera__matrix-element\matrix-synapse` | Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 50 | [suomitek/suomitek-charts](https://github.com/suomitek/suomitek-charts) | `cloned\suomitek__suomitek-charts\suomitek\airflow` | Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 51 | [sapcc/helm-charts](https://github.com/sapcc/helm-charts) | `cloned\sapcc__helm-charts\common\inventory-updater` | Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 1 charts Downloading owner-info from repo oci://keppel.eu-de-1.cloud.sap/ccloud-helm Save error occurred:  could not download oci://keppel.eu-de-1.cloud.sap/ccloud-helm/owner-info: failed to resolve keppel.eu-de-1.cloud.sap/ccloud-helm/owner-info:0.2.0: Get "https://keppel.eu-de-1.cloud.sap/v2/ccloud-helm/owner-info/manifests/0.2.0": dial tcp: lookup keppel.eu-de-1.cloud.sap: getaddrinfow: Este é geralmente um erro temporário durante a resolução de nomes de anfitrião e significa que o servidor local não recebeu uma resposta de um servidor autoritário. Error: could not download oci://keppel.eu-de-1.cloud.sap/ccloud-helm/owner-info: failed to resolve keppel.eu-de-1.cloud.sap/ccloud-helm/owner-info:0.2.0: Get "https://keppel.eu-de-1.cloud.sap/v2/ccloud-helm/owner-info/manifests/0.2.0": dial tcp: lookup keppel.eu-de-1.cloud.sap: getaddrinfow: Este é geralmente um erro temporário durante a resolução de nomes de anfitrião e significa que o servidor local não recebeu uma resposta de um servidor autoritário. |
| 52 | [feliux/containers-lab](https://github.com/feliux/containers-lab) | `cloned\feliux__containers-lab\k8s\sysdig` | Error: directory cloned\feliux__containers-lab\k8s\admission-controller not found |
| 53 | [quantumreasoning/kubeapps](https://github.com/quantumreasoning/kubeapps) | `cloned\quantumreasoning__kubeapps\chart\kubeapps` | Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 3 charts Downloading redis from repo oci://registry-1.docker.io/bitnamicharts Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:20.2.0: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/20.2.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:20.2.0: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/20.2.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 54 | [ndawpa/guiadodevops](https://github.com/ndawpa/guiadodevops) | `cloned\ndawpa__guiadodevops\kubernetes\helm\airflow-1.18.0\airflow` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 55 | [jstrachan/demo16](https://github.com/jstrachan/demo16) | `cloned\jstrachan__demo16\charts\preview` | Getting updates for unmanaged Helm repositories... ...Unable to get an update from the "https://chartmuseum.build.cd.jenkins-x.io" chart repository: 	Get "https://chartmuseum.build.cd.jenkins-x.io/index.yaml": dial tcp 35.205.186.189:443: connectex: Uma tentativa de ligação falhou porque o componente ligado não respondeu corretamente após um período de tempo, ou a ligação estabelecida falhou porque o anfitrião ligado não respondeu. Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Error: no cached repository for helm-manager-e0b822155b4ae5f667d1ffc8875a97176819d25e9b8afb113dc2a031f0a96e31 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-e0b822155b4ae5f667d1ffc8875a97176819d25e9b8afb113dc2a031f0a96e31-index.yaml: O sistema não conseguiu localizar o ficheiro especificado. |
| 56 | [t9k/apps](https://github.com/t9k/apps) | `cloned\t9k__apps\user-console\chatnio\chart` | Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 2 charts Downloading mysql from repo oci://registry-1.docker.io/bitnamicharts Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/mysql: failed to resolve registry-1.docker.io/bitnamicharts/mysql:9.18.2: GET "https://registry-1.docker.io/v2/bitnamicharts/mysql/manifests/9.18.2": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/mysql: failed to resolve registry-1.docker.io/bitnamicharts/mysql:9.18.2: GET "https://registry-1.docker.io/v2/bitnamicharts/mysql/manifests/9.18.2": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 57 | [TrevorSquillario/iDRAC-Telemetry-Ansible-Demo](https://github.com/TrevorSquillario/iDRAC-Telemetry-Ansible-Demo) | `cloned\TrevorSquillario__iDRAC-Telemetry-Ansible-Demo\k8s\helm\app` | Error: no repository definition for https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 58 | [DOME-Marketplace/dome-gitops](https://github.com/DOME-Marketplace/dome-gitops) | `cloned\DOME-Marketplace__dome-gitops\ionos\marketplace\kafka` | Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 1 charts Downloading kafka from repo oci://registry-1.docker.io/bitnamicharts Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/kafka: failed to resolve registry-1.docker.io/bitnamicharts/kafka:26.0.0: GET "https://registry-1.docker.io/v2/bitnamicharts/kafka/manifests/26.0.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/kafka: failed to resolve registry-1.docker.io/bitnamicharts/kafka:26.0.0: GET "https://registry-1.docker.io/v2/bitnamicharts/kafka/manifests/26.0.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 59 | [deepdivenow/rke2-charts](https://github.com/deepdivenow/rke2-charts) | `cloned\deepdivenow__rke2-charts\charts\harvester-cloud-provider\harvester-cloud-provider\0.2.100` | Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies |
| 60 | [aemneina/testcatalog](https://github.com/aemneina/testcatalog) | `cloned\aemneina__testcatalog\incubator\gogs` | manager.go:122: warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash... Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add' |
| 61 | [dynatrace-wwse/enablement-kubernetes-opentelemetry](https://github.com/dynatrace-wwse/enablement-kubernetes-opentelemetry) | `cloned\dynatrace-wwse__enablement-kubernetes-opentelemetry\.devcontainer\apps\astroshop\helm\dt-otel-demo-helm` | Error: no repository definition for https://open-telemetry.github.io/opentelemetry-helm-charts, https://open-telemetry.github.io/opentelemetry-helm-charts. Please add the missing repos via 'helm repo add' |
| 62 | [keonhoban/mlops-infra-gitops](https://github.com/keonhoban/mlops-infra-gitops) | `cloned\keonhoban__mlops-infra-gitops\baseline\charts\minio-wrapper` | Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 1 charts Downloading minio from repo oci://registry-1.docker.io/bitnamicharts Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/minio: failed to resolve registry-1.docker.io/bitnamicharts/minio:17.0.21: GET "https://registry-1.docker.io/v2/bitnamicharts/minio/manifests/17.0.21": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/minio: failed to resolve registry-1.docker.io/bitnamicharts/minio:17.0.21: GET "https://registry-1.docker.io/v2/bitnamicharts/minio/manifests/17.0.21": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 63 | [SBasak8/eu-val-loki](https://github.com/SBasak8/eu-val-loki) | `cloned\SBasak8__eu-val-loki\charts\airflow` | Getting updates for unmanaged Helm repositories... ...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Error: can't get a valid version for 1 subchart(s): "postgresql" (repository "https://charts.bitnami.com/bitnami", version "10.5.3"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml |
| 64 | [VadimShtukan/otus_homework](https://github.com/VadimShtukan/otus_homework) | `cloned\VadimShtukan__otus_homework\architect\lesson05\kubernetis\helm-chart` | Error: no repository definition for https://kubernetes-charts.storage.googleapis.com. Please add the missing repos via 'helm repo add' |
| 65 | [dariuszbolkowski/ezdrp_process](https://github.com/dariuszbolkowski/ezdrp_process) | `cloned\dariuszbolkowski__ezdrp_process\charts\ezd-backend\1.0.0\charts\mongodb` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 66 | [ValeriiVasianovych/argocd-kubernetes-infrastructure](https://github.com/ValeriiVasianovych/argocd-kubernetes-infrastructure) | `cloned\ValeriiVasianovych__argocd-kubernetes-infrastructure\icdev-01\helmCharts\prometheus-stack` | manager.go:122: warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash... Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies |
| 67 | [Jozefwl/applications-waldoserver](https://github.com/Jozefwl/applications-waldoserver) | `cloned\Jozefwl__applications-waldoserver\Applications\ERPNext\erpnext-official\erpnext` | Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 68 | [svtechnmaa/charts](https://github.com/svtechnmaa/charts) | `cloned\svtechnmaa__charts\kubernetes\airflow` | Getting updates for unmanaged Helm repositories... ...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Error: can't get a valid version for 1 subchart(s): "postgresql" (repository "https://charts.bitnami.com/bitnami", version "10.5.3"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml |
| 69 | [kingdonb/gitops-hephy](https://github.com/kingdonb/gitops-hephy) | `cloned\kingdonb__gitops-hephy\workflow` | Error: error unpacking subchart monitor in workflow: error unpacking subchart grafana in monitor: validation: chart.metadata.version "<Will be populated by the ci before publishing the chart>" is invalid |
| 70 | [NourBkh/k8s-config-repo](https://github.com/NourBkh/k8s-config-repo) | `cloned\NourBkh__k8s-config-repo\monitoring` | Error: no repository definition for https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 71 | [cyberrangecz/devops-helm](https://github.com/cyberrangecz/devops-helm) | `cloned\cyberrangecz__devops-helm\helm\crczp-head` | Error: directory cloned\cyberrangecz__devops-helm\helm\charts\common not found |
| 72 | [joostvdg/cmg](https://github.com/joostvdg/cmg) | `cloned\joostvdg__cmg\charts\preview` | Getting updates for unmanaged Helm repositories... ...Unable to get an update from the "http://chartmuseum.jenkins-x.io" chart repository: 	failed to fetch http://chartmuseum.jenkins-x.io/index.yaml : 404 Not Found Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Error: no cached repository for helm-manager-d8c90e93d45e753db0fa4ec457d893e08cd1e0827b77896ee636b714706cbbb2 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-d8c90e93d45e753db0fa4ec457d893e08cd1e0827b77896ee636b714706cbbb2-index.yaml: O sistema não conseguiu localizar o ficheiro especificado. |
| 73 | [superdoo/multitier-webapp](https://github.com/superdoo/multitier-webapp) | `cloned\superdoo__multitier-webapp\helm\postgresql` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 74 | [IvanciniGT/formacionHelm](https://github.com/IvanciniGT/formacionHelm) | `cloned\IvanciniGT__formacionHelm\charts\mariadb` | Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 1 charts Downloading common from repo oci://registry-1.docker.io/bitnamicharts Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/common: failed to resolve registry-1.docker.io/bitnamicharts/common:2.6.0: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.6.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/common: failed to resolve registry-1.docker.io/bitnamicharts/common:2.6.0: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.6.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 75 | [wadhekarpankaj/k8s-infrastructure](https://github.com/wadhekarpankaj/k8s-infrastructure) | `cloned\wadhekarpankaj__k8s-infrastructure\charts\argo-cd` | Error: no repository definition for https://dandydeveloper.github.io/charts/. Please add the missing repos via 'helm repo add' |
| 76 | [cb-kubecd/bdd-gh-1581604741](https://github.com/cb-kubecd/bdd-gh-1581604741) | `cloned\cb-kubecd__bdd-gh-1581604741\charts\preview` | Getting updates for unmanaged Helm repositories... ...Unable to get an update from the "http://chartmuseum.jenkins-x.io" chart repository: 	failed to fetch http://chartmuseum.jenkins-x.io/index.yaml : 404 Not Found Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Error: no cached repository for helm-manager-d8c90e93d45e753db0fa4ec457d893e08cd1e0827b77896ee636b714706cbbb2 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-d8c90e93d45e753db0fa4ec457d893e08cd1e0827b77896ee636b714706cbbb2-index.yaml: O sistema não conseguiu localizar o ficheiro especificado. |
| 77 | [Ammar-Taimoori/speechmatics-helm-chart](https://github.com/Ammar-Taimoori/speechmatics-helm-chart) | `cloned\Ammar-Taimoori__speechmatics-helm-chart\1.2.0\sm-realtime` | Error: directory cloned\Ammar-Taimoori__speechmatics-helm-chart\1.2.0\sm-proxy not found |
| 78 | [sookeke/jps-pidp](https://github.com/sookeke/jps-pidp) | `cloned\sookeke__jps-pidp\charts\nginx` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 79 | [HelmCI/ci-mon](https://github.com/HelmCI/ci-mon) | `cloned\HelmCI__ci-mon\charts\jaeger` | Error: no repository definition for https://charts.helm.sh/incubator, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 80 | [KevMCarp/truecharts-catalog-fork](https://github.com/KevMCarp/truecharts-catalog-fork) | `cloned\KevMCarp__truecharts-catalog-fork\dependency\clickhouse\5.0.54` | Getting updates for unmanaged Helm repositories... ...Unable to get an update from the "https://library-charts.truecharts.org" chart repository: 	Get "https://library-charts.truecharts.org/index.yaml": dial tcp: lookup library-charts.truecharts.org: no such host Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Error: no cached repository for helm-manager-024b189b59f6c6ccf0de6e5148db1578caf551c511f4eb220ece14cef00f80e0 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-024b189b59f6c6ccf0de6e5148db1578caf551c511f4eb220ece14cef00f80e0-index.yaml: O sistema não conseguiu localizar o ficheiro especificado. |
| 81 | [tetratelabs/charts](https://github.com/tetratelabs/charts) | `cloned\tetratelabs__charts\charts\demos\istio-monitoring-demo` | Error: no repository definition for https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 82 | [teamhephy/redis](https://github.com/teamhephy/redis) | `cloned\teamhephy__redis\charts\redis` | Error: validation: chart.metadata.version "<Will be populated by the ci before publishing the chart>" is invalid |
| 83 | [wesparish/rancher-catalog](https://github.com/wesparish/rancher-catalog) | `cloned\wesparish__rancher-catalog\charts\ark-wes` | Error: directory cloned\ark-server-charts\charts\ark-cluster not found |
| 84 | [cb-kubecd/bdd-nh-import-1586949778](https://github.com/cb-kubecd/bdd-nh-import-1586949778) | `cloned\cb-kubecd__bdd-nh-import-1586949778\charts\preview` | Getting updates for unmanaged Helm repositories... ...Unable to get an update from the "http://chartmuseum.jenkins-x.io" chart repository: 	failed to fetch http://chartmuseum.jenkins-x.io/index.yaml : 404 Not Found Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Error: no cached repository for helm-manager-d8c90e93d45e753db0fa4ec457d893e08cd1e0827b77896ee636b714706cbbb2 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-d8c90e93d45e753db0fa4ec457d893e08cd1e0827b77896ee636b714706cbbb2-index.yaml: O sistema não conseguiu localizar o ficheiro especificado. |
| 85 | [tetratelabs/helm-charts](https://github.com/tetratelabs/helm-charts) | `cloned\tetratelabs__helm-charts\charts\demos\istio-monitoring-demo` | Error: no repository definition for https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 86 | [ArvigEnterprises/integration-tests-bluegreen](https://github.com/ArvigEnterprises/integration-tests-bluegreen) | `cloned\ArvigEnterprises__integration-tests-bluegreen\.helm\charts\arvigtestcluster334d\canary` | Error: directory cloned\ArvigEnterprises__integration-tests-bluegreen\.helm\charts\arvigtestcluster334d\laravel-library not found |
| 87 | [tis-system/charts](https://github.com/tis-system/charts) | `cloned\tis-system__charts\charts\demos\istio-monitoring-demo\0.1.0` | Error: no repository definition for https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add' |
| 88 | [LMCache/LMBench](https://github.com/LMCache/LMBench) | `cloned\LMCache__LMBench\2-serving-engines\dynamo\dynamo\deploy\cloud\helm\platform` | Getting updates for unmanaged Helm repositories... ...Successfully got an update from the "https://nats-io.github.io/k8s/helm/charts/" chart repository ...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository ...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository ...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 6 charts Downloading nats from repo https://nats-io.github.io/k8s/helm/charts/ Downloading etcd from repo https://charts.bitnami.com/bitnami Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/etcd: failed to resolve registry-1.docker.io/bitnamicharts/etcd:11.1.0: GET "https://registry-1.docker.io/v2/bitnamicharts/etcd/manifests/11.1.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/etcd: failed to resolve registry-1.docker.io/bitnamicharts/etcd:11.1.0: GET "https://registry-1.docker.io/v2/bitnamicharts/etcd/manifests/11.1.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 89 | [tkestack/charts](https://github.com/tkestack/charts) | `cloned\tkestack__charts\incubator\airflow` | manager.go:122: warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash... Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/, https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add' |
| 90 | [shakedaskayo/kubiya-demo](https://github.com/shakedaskayo/kubiya-demo) | `cloned\shakedaskayo__kubiya-demo\helm\bitnami-charts\bitnami\airflow` | Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 3 charts Downloading redis from repo oci://registry-1.docker.io/bitnamicharts Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:18.11.0: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/18.11.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:18.11.0: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/18.11.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 91 | [moophlo/helm-charts](https://github.com/moophlo/helm-charts) | `cloned\moophlo__helm-charts\bitnami\airflow` | Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 3 charts Downloading redis from repo oci://registry-1.docker.io/bitnamicharts Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:22.0.4: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/22.0.4": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:22.0.4: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/22.0.4": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 92 | [aicanhelp/ai-cloud](https://github.com/aicanhelp/ai-cloud) | `cloned\aicanhelp__ai-cloud\bitnami\charts\bitnami-202503\airflow` | Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 3 charts Downloading redis from repo oci://registry-1.docker.io/bitnamicharts Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:20.7.1: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/20.7.1": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:20.7.1: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/20.7.1": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 93 | [itboon/bitnami-charts](https://github.com/itboon/bitnami-charts) | `cloned\itboon__bitnami-charts\charts-bitnami\airflow` | Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 3 charts Downloading redis from repo oci://registry-1.docker.io/bitnamicharts Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:22.0.4: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/22.0.4": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:22.0.4: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/22.0.4": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 94 | [p-/PyroDocker](https://github.com/p-/PyroDocker) | `cloned\p-__PyroDocker\compose\signoz\deploy\kubernetes\platform` | Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://charts.helm.sh/incubator. Please add the missing repos via 'helm repo add' |
| 95 | [yatvoikoresh/cloneranchercharts](https://github.com/yatvoikoresh/cloneranchercharts) | `cloned\yatvoikoresh__cloneranchercharts\charts\epinio\102.0.1+up1.6.2` | Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies |
| 96 | [otus-kuber-2020-04/israodin_platform](https://github.com/otus-kuber-2020-04/israodin_platform) | `cloned\otus-kuber-2020-04__israodin_platform\kubernetes-templating\hipster-shop` | Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add' |
| 97 | [ccojocar/spring-demo](https://github.com/ccojocar/spring-demo) | `cloned\ccojocar__spring-demo\charts\preview` | Getting updates for unmanaged Helm repositories... ...Unable to get an update from the "https://chartmuseum.build.cd.jenkins-x.io" chart repository: 	Get "https://chartmuseum.build.cd.jenkins-x.io/index.yaml": dial tcp 35.205.186.189:443: connectex: Uma tentativa de ligação falhou porque o componente ligado não respondeu corretamente após um período de tempo, ou a ligação estabelecida falhou porque o anfitrião ligado não respondeu. Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Error: no cached repository for helm-manager-e0b822155b4ae5f667d1ffc8875a97176819d25e9b8afb113dc2a031f0a96e31 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-e0b822155b4ae5f667d1ffc8875a97176819d25e9b8afb113dc2a031f0a96e31-index.yaml: O sistema não conseguiu localizar o ficheiro especificado. |
| 98 | [introproventures/spring-demo](https://github.com/introproventures/spring-demo) | `cloned\introproventures__spring-demo\charts\preview` | Getting updates for unmanaged Helm repositories... ...Unable to get an update from the "https://chartmuseum.build.cd.jenkins-x.io" chart repository: 	Get "https://chartmuseum.build.cd.jenkins-x.io/index.yaml": dial tcp 35.205.186.189:443: connectex: Uma tentativa de ligação falhou porque o componente ligado não respondeu corretamente após um período de tempo, ou a ligação estabelecida falhou porque o anfitrião ligado não respondeu. Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Error: no cached repository for helm-manager-e0b822155b4ae5f667d1ffc8875a97176819d25e9b8afb113dc2a031f0a96e31 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-e0b822155b4ae5f667d1ffc8875a97176819d25e9b8afb113dc2a031f0a96e31-index.yaml: O sistema não conseguiu localizar o ficheiro especificado. |
| 99 | [leandrofrs/spring_demo](https://github.com/leandrofrs/spring_demo) | `cloned\leandrofrs__spring_demo\charts\preview` | Getting updates for unmanaged Helm repositories... ...Unable to get an update from the "http://chartmuseum.jenkins-x.io" chart repository: 	failed to fetch http://chartmuseum.jenkins-x.io/index.yaml : 404 Not Found Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Error: no cached repository for helm-manager-d8c90e93d45e753db0fa4ec457d893e08cd1e0827b77896ee636b714706cbbb2 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-d8c90e93d45e753db0fa4ec457d893e08cd1e0827b77896ee636b714706cbbb2-index.yaml: O sistema não conseguiu localizar o ficheiro especificado. |
| 100 | [susnata1981/spring-demo](https://github.com/susnata1981/spring-demo) | `cloned\susnata1981__spring-demo\charts\preview` | Getting updates for unmanaged Helm repositories... ...Unable to get an update from the "http://chartmuseum.jenkins-x.io" chart repository: 	failed to fetch http://chartmuseum.jenkins-x.io/index.yaml : 404 Not Found Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Error: no cached repository for helm-manager-d8c90e93d45e753db0fa4ec457d893e08cd1e0827b77896ee636b714706cbbb2 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-d8c90e93d45e753db0fa4ec457d893e08cd1e0827b77896ee636b714706cbbb2-index.yaml: O sistema não conseguiu localizar o ficheiro especificado. |
| 101 | [Jaebytes/TrueCharts](https://github.com/Jaebytes/TrueCharts) | `cloned\Jaebytes__TrueCharts\premium\traefik\26.10.19` | Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 1 charts Downloading common from repo oci://tccr.io/truecharts Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to resolve tccr.io/truecharts/common:20.3.12: Get "https://tccr.io/v2/truecharts/common/manifests/20.3.12": dial tcp: lookup tccr.io: no such host Error: could not download oci://tccr.io/truecharts/common: failed to resolve tccr.io/truecharts/common:20.3.12: Get "https://tccr.io/v2/truecharts/common/manifests/20.3.12": dial tcp: lookup tccr.io: no such host |
| 102 | [PnsDev/Truechart-catalog-archive](https://github.com/PnsDev/Truechart-catalog-archive) | `cloned\PnsDev__Truechart-catalog-archive\incubator\archivebox\0.7.2` | Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 1 charts Downloading common from repo oci://tccr.io/truecharts Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to resolve tccr.io/truecharts/common:17.2.30: Get "https://tccr.io/v2/truecharts/common/manifests/17.2.30": dial tcp: lookup tccr.io: no such host Error: could not download oci://tccr.io/truecharts/common: failed to resolve tccr.io/truecharts/common:17.2.30: Get "https://tccr.io/v2/truecharts/common/manifests/17.2.30": dial tcp: lookup tccr.io: no such host |
| 103 | [petersandor/tc-archive](https://github.com/petersandor/tc-archive) | `cloned\petersandor__tc-archive\scale-catalog\incubator\archivebox\0.7.2` | Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 1 charts Downloading common from repo oci://tccr.io/truecharts Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to resolve tccr.io/truecharts/common:17.2.30: Get "https://tccr.io/v2/truecharts/common/manifests/17.2.30": dial tcp: lookup tccr.io: no such host Error: could not download oci://tccr.io/truecharts/common: failed to resolve tccr.io/truecharts/common:17.2.30: Get "https://tccr.io/v2/truecharts/common/manifests/17.2.30": dial tcp: lookup tccr.io: no such host |
| 104 | [adstanley/archive](https://github.com/adstanley/archive) | `cloned\adstanley__archive\scale-catalog\incubator\archivebox\0.7.2` | Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 1 charts Downloading common from repo oci://tccr.io/truecharts Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to resolve tccr.io/truecharts/common:17.2.30: Get "https://tccr.io/v2/truecharts/common/manifests/17.2.30": dial tcp: lookup tccr.io: no such host Error: could not download oci://tccr.io/truecharts/common: failed to resolve tccr.io/truecharts/common:17.2.30: Get "https://tccr.io/v2/truecharts/common/manifests/17.2.30": dial tcp: lookup tccr.io: no such host |
| 105 | [tdunlap607/kubeapps-fork](https://github.com/tdunlap607/kubeapps-fork) | `cloned\tdunlap607__kubeapps-fork\chart\kubeapps` | Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 3 charts Downloading redis from repo oci://registry-1.docker.io/bitnamicharts Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:20.6.1: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/20.6.1": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit Error: could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:20.6.1: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/20.6.1": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit |
| 106 | [adynetro/truecharts](https://github.com/adynetro/truecharts) | `cloned\adynetro__truecharts\incubator\archivebox\0.7.2` | Hang tight while we grab the latest from your chart repositories... ...Successfully got an update from the "traefik" chart repository Update Complete. ⎈Happy Helming!⎈ Saving 1 charts Downloading common from repo oci://tccr.io/truecharts Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to resolve tccr.io/truecharts/common:17.2.30: Get "https://tccr.io/v2/truecharts/common/manifests/17.2.30": dial tcp: lookup tccr.io: no such host Error: could not download oci://tccr.io/truecharts/common: failed to resolve tccr.io/truecharts/common:17.2.30: Get "https://tccr.io/v2/truecharts/common/manifests/17.2.30": dial tcp: lookup tccr.io: no such host |

## Full Error Details

### `cloned\ConductionNL__procces-engine\api\helm` — ConductionNL/procces-engine

```
manager.go:122: warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash...
Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies
```

### `cloned\pluralsh__plural-helm-charts\charts\airbyte` — pluralsh/plural-helm-charts

```
Error: no repository definition for https://charts.bitnami.com/bitnami, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts. Please add the missing repos via 'helm repo add'
```

### `cloned\michaelw__common.itsumi` — michaelw/common.itsumi

```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 1 charts
Downloading common from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/common: failed to resolve registry-1.docker.io/bitnamicharts/common:2.31.3: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.31.3": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/common: failed to resolve registry-1.docker.io/bitnamicharts/common:2.31.3: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.31.3": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `cloned\solo-io__skv2\codegen\test\chart\conditional-crds` — solo-io/skv2

```
Error: validation: chart.metadata.name is required
```

### `cloned\flytrap__docker-admin\charts\apisix-1.3` — flytrap/docker-admin

```
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.apiseven.com, https://charts.apiseven.com. Please add the missing repos via 'helm repo add'
```

### `cloned\akv-akv__de_research\kubernetes\charts\airbyte` — akv-akv/de_research

```
Error: no repository definition for https://charts.bitnami.com/bitnami, https://airbytehq.github.io/helm-charts/, https://airbytehq.github.io/helm-charts/, https://airbytehq.github.io/helm-charts/, https://airbytehq.github.io/helm-charts/, https://airbytehq.github.io/helm-charts/, https://airbytehq.github.io/helm-charts/, https://airbytehq.github.io/helm-charts/, https://airbytehq.github.io/helm-charts/, https://airbytehq.github.io/helm-charts/, https://airbytehq.github.io/helm-charts/, https://airbytehq.github.io/helm-charts/, https://airbytehq.github.io/helm-charts/, https://airbytehq.github.io/helm-charts/. Please add the missing repos via 'helm repo add'
```

### `cloned\ConductionNL__Authorization-component\api\helm` — ConductionNL/Authorization-component

```
load.go:138: Warning: Dependency locking is handled in Chart.lock since apiVersion "v2". We recommend migrating to Chart.lock.
Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies
```

### `cloned\MirahImage__postfacto\tkg\postfacto` — MirahImage/postfacto

```
Error: no repository definition for @bitnami, @bitnami. Please add them via 'helm repo add'
```

### `cloned\noygal__helm\charts\example-dev-tools` — noygal/helm

```
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "https://noygal.github.com/helm" chart repository:
	Get "https://noygal.github.com/helm/index.yaml": dial tcp: lookup noygal.github.com: no such host
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Error: no cached repository for helm-manager-19cefd96797d2c0b3ed74271dcee4ae7ad7d02cb4d4e47a98431938d24b11530 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-19cefd96797d2c0b3ed74271dcee4ae7ad7d02cb4d4e47a98431938d24b11530-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### `cloned\manics__zero-to-jupyterhub-k8s-chart-expanded\0.8.1` — manics/zero-to-jupyterhub-k8s-chart-expanded

```
manager.go:122: warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash...
Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies
```

### `cloned\unicef__magasin\helm\dagster` — unicef/magasin

```
Error: no repository definition for https://raw.githubusercontent.com/bitnami/charts/eb5f9a9513d987b519f0ecd732e7031241c50328/bitnami, https://raw.githubusercontent.com/bitnami/charts/eb5f9a9513d987b519f0ecd732e7031241c50328/bitnami, https://raw.githubusercontent.com/bitnami/charts/eb5f9a9513d987b519f0ecd732e7031241c50328/bitnami. Please add the missing repos via 'helm repo add'
```

### `cloned\ConductionNL__berichtservice\api\helm` — ConductionNL/berichtservice

```
load.go:138: Warning: Dependency locking is handled in Chart.lock since apiVersion "v2". We recommend migrating to Chart.lock.
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `cloned\ConductionNL__commonground-ui\api\helm` — ConductionNL/commonground-ui

```
manager.go:122: warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash...
Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies
```

### `cloned\Veronneau-Techno-Conseil__comax-commons\helm_referee` — Veronneau-Techno-Conseil/comax-commons

```
Error: no repository definition for https://mongodb.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `cloned\waTeim__mcp-base\chart` — waTeim/mcp-base

```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 1 charts
Downloading redis from repo oci://registry-1.docker.io/cloudpirates
Save error occurred:  could not download oci://registry-1.docker.io/cloudpirates/redis: failed to resolve registry-1.docker.io/cloudpirates/redis:0.16.4: GET "https://registry-1.docker.io/v2/cloudpirates/redis/manifests/0.16.4": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/cloudpirates/redis: failed to resolve registry-1.docker.io/cloudpirates/redis:0.16.4: GET "https://registry-1.docker.io/v2/cloudpirates/redis/manifests/0.16.4": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `cloned\datawizz__Firestream\src\charts\bitnami\airflow` — datawizz/Firestream

```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 3 charts
Downloading redis from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:21.2.5: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/21.2.5": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:21.2.5: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/21.2.5": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `cloned\limagbz__data-mesh-yelp\infra\monitoring\loki\helm` — limagbz/data-mesh-yelp

```
Error: no repository definition for https://charts.min.io/, https://grafana.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `cloned\efremov-it__security\practice\weeks\2-3\secureCodeBox\operator` — efremov-it/security

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `cloned\ConductionNL__zaak-registratie-component\api\helm` — ConductionNL/zaak-registratie-component

```
manager.go:122: warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash...
Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies
```

### `cloned\ConductionNL__checkinservice\api\helm` — ConductionNL/checkinservice

```
Getting updates for unmanaged Helm repositories...
...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Error: can't get a valid version for 1 subchart(s): "postgresql" (repository "https://charts.bitnami.com/bitnami", version "10.1.1"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml
```

### `cloned\ConductionNL__digispoof-old\api\helm` — ConductionNL/digispoof-old

```
manager.go:122: warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash...
Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies
```

### `cloned\ConductionNL__proto-application-commonground\api\helm` — ConductionNL/proto-application-commonground

```
manager.go:122: warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash...
Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies
```

### `cloned\bitnami__charts\bitnami\airflow` — bitnami/charts

```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 3 charts
Downloading redis from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:22.0.4: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/22.0.4": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:22.0.4: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/22.0.4": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `cloned\grafana__helm-charts\charts\enterprise-metrics` — grafana/helm-charts

```
Getting updates for unmanaged Helm repositories...
...Successfully got an update from the "https://helm.min.io/" chart repository
...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Error: minio chart not found in repo https://helm.min.io/
```

### `cloned\cloudnativeapp__charts\curated\airflow` — cloudnativeapp/charts

```
manager.go:122: warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash...
Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies
```

### `cloned\krateoplatformops__composition-portal-starter\chart` — krateoplatformops/composition-portal-starter

```
Error: validation: chart.metadata.version "CHART_VERSION" is invalid
```

### `cloned\samcab28__Diseno-Software\proyecto\caso4\charts\databases` — samcab28/Diseno-Software

```
Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies
```

### `cloned\rancher__charts\charts\epinio\102.0.1+up1.6.2` — rancher/charts

```
Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies
```

### `cloned\gersonrs__hand-on-airflow-with-minio\airflow` — gersonrs/hand-on-airflow-with-minio

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `cloned\musqat__lectures\MicroService\section_16\helm\grafana` — musqat/lectures

```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 1 charts
Downloading common from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/common: failed to resolve registry-1.docker.io/bitnamicharts/common:2.31.0: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.31.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/common: failed to resolve registry-1.docker.io/bitnamicharts/common:2.31.0: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.31.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `cloned\general-rj45__mini-ml-stand-GRJ45\airflow-14.2.5\airflow` — general-rj45/mini-ml-stand-GRJ45

```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 3 charts
Downloading redis from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:17.11.3: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/17.11.3": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:17.11.3: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/17.11.3": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `cloned\shini4i__charts\charts\app` — shini4i/charts

```
Error: no repository definition for https://shini4i.github.io/charts/. Please add the missing repos via 'helm repo add'
```

### `cloned\otus-kuber-2020-07__israodin_platform\kubernetes-templating\hipster-shop` — otus-kuber-2020-07/israodin_platform

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `cloned\robertbenzinger__charts\janusgraph` — robertbenzinger/charts

```
Getting updates for unmanaged Helm repositories...
...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository
...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Error: can't get a valid version for 2 subchart(s): "cassandra" (repository "https://charts.bitnami.com/bitnami", version "5.6.1"), "elasticsearch" (repository "https://charts.bitnami.com/bitnami", version "12.6.2"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml
```

### `cloned\openlit__helm\charts\openlit` — openlit/helm

```
Error: no repository definition for https://openlit.github.io/helm/. Please add the missing repos via 'helm repo add'
```

### `cloned\marcus-sa__helm-charts\repositories\banzaicloud-stable\anchore-engine\0.13.0\chart` — marcus-sa/helm-charts

```
manager.go:122: warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash...
Error: no repository definition for https://kubernetes-charts.storage.googleapis.com, https://kubernetes-charts.storage.googleapis.com, https://kubernetes-charts.storage.googleapis.com. Please add the missing repos via 'helm repo add'
```

### `cloned\hey101__scale-catalog\incubator\archivebox\0.7.2` — hey101/scale-catalog

```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 1 charts
Downloading common from repo oci://tccr.io/truecharts
Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to resolve tccr.io/truecharts/common:17.5.1: Get "https://tccr.io/v2/truecharts/common/manifests/17.5.1": dial tcp: lookup tccr.io: no such host
Error: could not download oci://tccr.io/truecharts/common: failed to resolve tccr.io/truecharts/common:17.5.1: Get "https://tccr.io/v2/truecharts/common/manifests/17.5.1": dial tcp: lookup tccr.io: no such host
```

### `cloned\unity-systems__spring-demo\charts\preview` — unity-systems/spring-demo

```
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "https://chartmuseum.build.cd.jenkins-x.io" chart repository:
	Get "https://chartmuseum.build.cd.jenkins-x.io/index.yaml": dial tcp 35.205.186.189:443: connectex: Uma tentativa de ligação falhou porque o componente ligado não respondeu
corretamente após um período de tempo, ou a ligação estabelecida falhou
porque o anfitrião ligado não respondeu.
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Error: no cached repository for helm-manager-e0b822155b4ae5f667d1ffc8875a97176819d25e9b8afb113dc2a031f0a96e31 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-e0b822155b4ae5f667d1ffc8875a97176819d25e9b8afb113dc2a031f0a96e31-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### `cloned\deresegetachew__playing-helm\kanban-backend` — deresegetachew/playing-helm

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `cloned\nadav07__helmTesting\chartTest` — nadav07/helmTesting

```
Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies
```

### `cloned\ibuildthecloud__rancher-charts\charts\anchore-engine\0.1.0` — ibuildthecloud/rancher-charts

```
manager.go:122: warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash...
Error: no repository definition for https://kubernetes-charts.storage.googleapis.com. Please add the missing repos via 'helm repo add'
```

### `cloned\czerwonk__hakanai\helm\hakanai` — czerwonk/hakanai

```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 1 charts
Downloading redis from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:20.5.0: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/20.5.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:20.5.0: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/20.5.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `cloned\Abinesh0206__gitlab-ce` — Abinesh0206/gitlab-ce

```
Error: no repository definition for https://charts.jetstack.io/, https://prometheus-community.github.io/helm-charts, https://charts.bitnami.com/bitnami, https://charts.gitlab.io/, https://charts.bitnami.com/bitnami, https://charts.gitlab.io/, https://helm.traefik.io/traefik, https://haproxytech.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `cloned\samfms__truecharts\deprecated\bitwarden\1.2.5` — samfms/truecharts

```
Error: no repository definition for https://truecharts.org/, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `cloned\Blankeeir__demo\infrastructure\helm\rafiki` — Blankeeir/demo

```
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://interledger.github.io/helm-charts, https://interledger.github.io/helm-charts, https://interledger.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `cloned\randoli__helm-charts\charts\loki` — randoli/helm-charts

```
Error: no repository definition for https://charts.min.io/, https://grafana.github.io/helm-charts, https://grafana.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `cloned\dylanjustice__the-grid\gitops\workloads\kube-prometheus` — dylanjustice/the-grid

```
Error: no repository definition for https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `cloned\Mark4562__Truenas-Scale\premium\authelia\23.13.13` — Mark4562/Truenas-Scale

```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 2 charts
Downloading common from repo oci://tccr.io/truecharts
Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to resolve tccr.io/truecharts/common:20.3.11: Get "https://tccr.io/v2/truecharts/common/manifests/20.3.11": dial tcp: lookup tccr.io: no such host
Error: could not download oci://tccr.io/truecharts/common: failed to resolve tccr.io/truecharts/common:20.3.11: Get "https://tccr.io/v2/truecharts/common/manifests/20.3.11": dial tcp: lookup tccr.io: no such host
```

### `cloned\praveenperera__matrix-element\matrix-synapse` — praveenperera/matrix-element

```
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `cloned\suomitek__suomitek-charts\suomitek\airflow` — suomitek/suomitek-charts

```
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `cloned\sapcc__helm-charts\common\inventory-updater` — sapcc/helm-charts

```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 1 charts
Downloading owner-info from repo oci://keppel.eu-de-1.cloud.sap/ccloud-helm
Save error occurred:  could not download oci://keppel.eu-de-1.cloud.sap/ccloud-helm/owner-info: failed to resolve keppel.eu-de-1.cloud.sap/ccloud-helm/owner-info:0.2.0: Get "https://keppel.eu-de-1.cloud.sap/v2/ccloud-helm/owner-info/manifests/0.2.0": dial tcp: lookup keppel.eu-de-1.cloud.sap: getaddrinfow: Este é geralmente um erro temporário durante a resolução de nomes de anfitrião e significa que o servidor local não recebeu uma resposta de um servidor autoritário.
Error: could not download oci://keppel.eu-de-1.cloud.sap/ccloud-helm/owner-info: failed to resolve keppel.eu-de-1.cloud.sap/ccloud-helm/owner-info:0.2.0: Get "https://keppel.eu-de-1.cloud.sap/v2/ccloud-helm/owner-info/manifests/0.2.0": dial tcp: lookup keppel.eu-de-1.cloud.sap: getaddrinfow: Este é geralmente um erro temporário durante a resolução de nomes de anfitrião e significa que o servidor local não recebeu uma resposta de um servidor autoritário.
```

### `cloned\feliux__containers-lab\k8s\sysdig` — feliux/containers-lab

```
Error: directory cloned\feliux__containers-lab\k8s\admission-controller not found
```

### `cloned\quantumreasoning__kubeapps\chart\kubeapps` — quantumreasoning/kubeapps

```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 3 charts
Downloading redis from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:20.2.0: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/20.2.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:20.2.0: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/20.2.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `cloned\ndawpa__guiadodevops\kubernetes\helm\airflow-1.18.0\airflow` — ndawpa/guiadodevops

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `cloned\jstrachan__demo16\charts\preview` — jstrachan/demo16

```
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "https://chartmuseum.build.cd.jenkins-x.io" chart repository:
	Get "https://chartmuseum.build.cd.jenkins-x.io/index.yaml": dial tcp 35.205.186.189:443: connectex: Uma tentativa de ligação falhou porque o componente ligado não respondeu
corretamente após um período de tempo, ou a ligação estabelecida falhou
porque o anfitrião ligado não respondeu.
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Error: no cached repository for helm-manager-e0b822155b4ae5f667d1ffc8875a97176819d25e9b8afb113dc2a031f0a96e31 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-e0b822155b4ae5f667d1ffc8875a97176819d25e9b8afb113dc2a031f0a96e31-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### `cloned\t9k__apps\user-console\chatnio\chart` — t9k/apps

```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 2 charts
Downloading mysql from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/mysql: failed to resolve registry-1.docker.io/bitnamicharts/mysql:9.18.2: GET "https://registry-1.docker.io/v2/bitnamicharts/mysql/manifests/9.18.2": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/mysql: failed to resolve registry-1.docker.io/bitnamicharts/mysql:9.18.2: GET "https://registry-1.docker.io/v2/bitnamicharts/mysql/manifests/9.18.2": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `cloned\TrevorSquillario__iDRAC-Telemetry-Ansible-Demo\k8s\helm\app` — TrevorSquillario/iDRAC-Telemetry-Ansible-Demo

```
Error: no repository definition for https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `cloned\DOME-Marketplace__dome-gitops\ionos\marketplace\kafka` — DOME-Marketplace/dome-gitops

```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 1 charts
Downloading kafka from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/kafka: failed to resolve registry-1.docker.io/bitnamicharts/kafka:26.0.0: GET "https://registry-1.docker.io/v2/bitnamicharts/kafka/manifests/26.0.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/kafka: failed to resolve registry-1.docker.io/bitnamicharts/kafka:26.0.0: GET "https://registry-1.docker.io/v2/bitnamicharts/kafka/manifests/26.0.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `cloned\deepdivenow__rke2-charts\charts\harvester-cloud-provider\harvester-cloud-provider\0.2.100` — deepdivenow/rke2-charts

```
Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies
```

### `cloned\aemneina__testcatalog\incubator\gogs` — aemneina/testcatalog

```
manager.go:122: warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash...
Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add'
```

### `cloned\dynatrace-wwse__enablement-kubernetes-opentelemetry\.devcontainer\apps\astroshop\helm\dt-otel-demo-helm` — dynatrace-wwse/enablement-kubernetes-opentelemetry

```
Error: no repository definition for https://open-telemetry.github.io/opentelemetry-helm-charts, https://open-telemetry.github.io/opentelemetry-helm-charts. Please add the missing repos via 'helm repo add'
```

### `cloned\keonhoban__mlops-infra-gitops\baseline\charts\minio-wrapper` — keonhoban/mlops-infra-gitops

```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 1 charts
Downloading minio from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/minio: failed to resolve registry-1.docker.io/bitnamicharts/minio:17.0.21: GET "https://registry-1.docker.io/v2/bitnamicharts/minio/manifests/17.0.21": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/minio: failed to resolve registry-1.docker.io/bitnamicharts/minio:17.0.21: GET "https://registry-1.docker.io/v2/bitnamicharts/minio/manifests/17.0.21": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `cloned\SBasak8__eu-val-loki\charts\airflow` — SBasak8/eu-val-loki

```
Getting updates for unmanaged Helm repositories...
...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Error: can't get a valid version for 1 subchart(s): "postgresql" (repository "https://charts.bitnami.com/bitnami", version "10.5.3"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml
```

### `cloned\VadimShtukan__otus_homework\architect\lesson05\kubernetis\helm-chart` — VadimShtukan/otus_homework

```
Error: no repository definition for https://kubernetes-charts.storage.googleapis.com. Please add the missing repos via 'helm repo add'
```

### `cloned\dariuszbolkowski__ezdrp_process\charts\ezd-backend\1.0.0\charts\mongodb` — dariuszbolkowski/ezdrp_process

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `cloned\ValeriiVasianovych__argocd-kubernetes-infrastructure\icdev-01\helmCharts\prometheus-stack` — ValeriiVasianovych/argocd-kubernetes-infrastructure

```
manager.go:122: warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash...
Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies
```

### `cloned\Jozefwl__applications-waldoserver\Applications\ERPNext\erpnext-official\erpnext` — Jozefwl/applications-waldoserver

```
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `cloned\svtechnmaa__charts\kubernetes\airflow` — svtechnmaa/charts

```
Getting updates for unmanaged Helm repositories...
...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Error: can't get a valid version for 1 subchart(s): "postgresql" (repository "https://charts.bitnami.com/bitnami", version "10.5.3"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml
```

### `cloned\kingdonb__gitops-hephy\workflow` — kingdonb/gitops-hephy

```
Error: error unpacking subchart monitor in workflow: error unpacking subchart grafana in monitor: validation: chart.metadata.version "<Will be populated by the ci before publishing the chart>" is invalid
```

### `cloned\NourBkh__k8s-config-repo\monitoring` — NourBkh/k8s-config-repo

```
Error: no repository definition for https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `cloned\cyberrangecz__devops-helm\helm\crczp-head` — cyberrangecz/devops-helm

```
Error: directory cloned\cyberrangecz__devops-helm\helm\charts\common not found
```

### `cloned\joostvdg__cmg\charts\preview` — joostvdg/cmg

```
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "http://chartmuseum.jenkins-x.io" chart repository:
	failed to fetch http://chartmuseum.jenkins-x.io/index.yaml : 404 Not Found
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Error: no cached repository for helm-manager-d8c90e93d45e753db0fa4ec457d893e08cd1e0827b77896ee636b714706cbbb2 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-d8c90e93d45e753db0fa4ec457d893e08cd1e0827b77896ee636b714706cbbb2-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### `cloned\superdoo__multitier-webapp\helm\postgresql` — superdoo/multitier-webapp

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `cloned\IvanciniGT__formacionHelm\charts\mariadb` — IvanciniGT/formacionHelm

```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 1 charts
Downloading common from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/common: failed to resolve registry-1.docker.io/bitnamicharts/common:2.6.0: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.6.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/common: failed to resolve registry-1.docker.io/bitnamicharts/common:2.6.0: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.6.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `cloned\wadhekarpankaj__k8s-infrastructure\charts\argo-cd` — wadhekarpankaj/k8s-infrastructure

```
Error: no repository definition for https://dandydeveloper.github.io/charts/. Please add the missing repos via 'helm repo add'
```

### `cloned\cb-kubecd__bdd-gh-1581604741\charts\preview` — cb-kubecd/bdd-gh-1581604741

```
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "http://chartmuseum.jenkins-x.io" chart repository:
	failed to fetch http://chartmuseum.jenkins-x.io/index.yaml : 404 Not Found
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Error: no cached repository for helm-manager-d8c90e93d45e753db0fa4ec457d893e08cd1e0827b77896ee636b714706cbbb2 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-d8c90e93d45e753db0fa4ec457d893e08cd1e0827b77896ee636b714706cbbb2-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### `cloned\Ammar-Taimoori__speechmatics-helm-chart\1.2.0\sm-realtime` — Ammar-Taimoori/speechmatics-helm-chart

```
Error: directory cloned\Ammar-Taimoori__speechmatics-helm-chart\1.2.0\sm-proxy not found
```

### `cloned\sookeke__jps-pidp\charts\nginx` — sookeke/jps-pidp

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `cloned\HelmCI__ci-mon\charts\jaeger` — HelmCI/ci-mon

```
Error: no repository definition for https://charts.helm.sh/incubator, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `cloned\KevMCarp__truecharts-catalog-fork\dependency\clickhouse\5.0.54` — KevMCarp/truecharts-catalog-fork

```
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "https://library-charts.truecharts.org" chart repository:
	Get "https://library-charts.truecharts.org/index.yaml": dial tcp: lookup library-charts.truecharts.org: no such host
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Error: no cached repository for helm-manager-024b189b59f6c6ccf0de6e5148db1578caf551c511f4eb220ece14cef00f80e0 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-024b189b59f6c6ccf0de6e5148db1578caf551c511f4eb220ece14cef00f80e0-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### `cloned\tetratelabs__charts\charts\demos\istio-monitoring-demo` — tetratelabs/charts

```
Error: no repository definition for https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `cloned\teamhephy__redis\charts\redis` — teamhephy/redis

```
Error: validation: chart.metadata.version "<Will be populated by the ci before publishing the chart>" is invalid
```

### `cloned\wesparish__rancher-catalog\charts\ark-wes` — wesparish/rancher-catalog

```
Error: directory cloned\ark-server-charts\charts\ark-cluster not found
```

### `cloned\cb-kubecd__bdd-nh-import-1586949778\charts\preview` — cb-kubecd/bdd-nh-import-1586949778

```
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "http://chartmuseum.jenkins-x.io" chart repository:
	failed to fetch http://chartmuseum.jenkins-x.io/index.yaml : 404 Not Found
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Error: no cached repository for helm-manager-d8c90e93d45e753db0fa4ec457d893e08cd1e0827b77896ee636b714706cbbb2 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-d8c90e93d45e753db0fa4ec457d893e08cd1e0827b77896ee636b714706cbbb2-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### `cloned\tetratelabs__helm-charts\charts\demos\istio-monitoring-demo` — tetratelabs/helm-charts

```
Error: no repository definition for https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `cloned\ArvigEnterprises__integration-tests-bluegreen\.helm\charts\arvigtestcluster334d\canary` — ArvigEnterprises/integration-tests-bluegreen

```
Error: directory cloned\ArvigEnterprises__integration-tests-bluegreen\.helm\charts\arvigtestcluster334d\laravel-library not found
```

### `cloned\tis-system__charts\charts\demos\istio-monitoring-demo\0.1.0` — tis-system/charts

```
Error: no repository definition for https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### `cloned\LMCache__LMBench\2-serving-engines\dynamo\dynamo\deploy\cloud\helm\platform` — LMCache/LMBench

```
Getting updates for unmanaged Helm repositories...
...Successfully got an update from the "https://nats-io.github.io/k8s/helm/charts/" chart repository
...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository
...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository
...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 6 charts
Downloading nats from repo https://nats-io.github.io/k8s/helm/charts/
Downloading etcd from repo https://charts.bitnami.com/bitnami
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/etcd: failed to resolve registry-1.docker.io/bitnamicharts/etcd:11.1.0: GET "https://registry-1.docker.io/v2/bitnamicharts/etcd/manifests/11.1.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/etcd: failed to resolve registry-1.docker.io/bitnamicharts/etcd:11.1.0: GET "https://registry-1.docker.io/v2/bitnamicharts/etcd/manifests/11.1.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `cloned\tkestack__charts\incubator\airflow` — tkestack/charts

```
manager.go:122: warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash...
Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/, https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add'
```

### `cloned\shakedaskayo__kubiya-demo\helm\bitnami-charts\bitnami\airflow` — shakedaskayo/kubiya-demo

```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 3 charts
Downloading redis from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:18.11.0: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/18.11.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:18.11.0: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/18.11.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `cloned\moophlo__helm-charts\bitnami\airflow` — moophlo/helm-charts

```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 3 charts
Downloading redis from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:22.0.4: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/22.0.4": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:22.0.4: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/22.0.4": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `cloned\aicanhelp__ai-cloud\bitnami\charts\bitnami-202503\airflow` — aicanhelp/ai-cloud

```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 3 charts
Downloading redis from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:20.7.1: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/20.7.1": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:20.7.1: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/20.7.1": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `cloned\itboon__bitnami-charts\charts-bitnami\airflow` — itboon/bitnami-charts

```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 3 charts
Downloading redis from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:22.0.4: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/22.0.4": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:22.0.4: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/22.0.4": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `cloned\p-__PyroDocker\compose\signoz\deploy\kubernetes\platform` — p-/PyroDocker

```
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://charts.helm.sh/incubator. Please add the missing repos via 'helm repo add'
```

### `cloned\yatvoikoresh__cloneranchercharts\charts\epinio\102.0.1+up1.6.2` — yatvoikoresh/cloneranchercharts

```
Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies
```

### `cloned\otus-kuber-2020-04__israodin_platform\kubernetes-templating\hipster-shop` — otus-kuber-2020-04/israodin_platform

```
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### `cloned\ccojocar__spring-demo\charts\preview` — ccojocar/spring-demo

```
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "https://chartmuseum.build.cd.jenkins-x.io" chart repository:
	Get "https://chartmuseum.build.cd.jenkins-x.io/index.yaml": dial tcp 35.205.186.189:443: connectex: Uma tentativa de ligação falhou porque o componente ligado não respondeu
corretamente após um período de tempo, ou a ligação estabelecida falhou
porque o anfitrião ligado não respondeu.
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Error: no cached repository for helm-manager-e0b822155b4ae5f667d1ffc8875a97176819d25e9b8afb113dc2a031f0a96e31 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-e0b822155b4ae5f667d1ffc8875a97176819d25e9b8afb113dc2a031f0a96e31-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### `cloned\introproventures__spring-demo\charts\preview` — introproventures/spring-demo

```
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "https://chartmuseum.build.cd.jenkins-x.io" chart repository:
	Get "https://chartmuseum.build.cd.jenkins-x.io/index.yaml": dial tcp 35.205.186.189:443: connectex: Uma tentativa de ligação falhou porque o componente ligado não respondeu
corretamente após um período de tempo, ou a ligação estabelecida falhou
porque o anfitrião ligado não respondeu.
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Error: no cached repository for helm-manager-e0b822155b4ae5f667d1ffc8875a97176819d25e9b8afb113dc2a031f0a96e31 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-e0b822155b4ae5f667d1ffc8875a97176819d25e9b8afb113dc2a031f0a96e31-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### `cloned\leandrofrs__spring_demo\charts\preview` — leandrofrs/spring_demo

```
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "http://chartmuseum.jenkins-x.io" chart repository:
	failed to fetch http://chartmuseum.jenkins-x.io/index.yaml : 404 Not Found
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Error: no cached repository for helm-manager-d8c90e93d45e753db0fa4ec457d893e08cd1e0827b77896ee636b714706cbbb2 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-d8c90e93d45e753db0fa4ec457d893e08cd1e0827b77896ee636b714706cbbb2-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### `cloned\susnata1981__spring-demo\charts\preview` — susnata1981/spring-demo

```
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "http://chartmuseum.jenkins-x.io" chart repository:
	failed to fetch http://chartmuseum.jenkins-x.io/index.yaml : 404 Not Found
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Error: no cached repository for helm-manager-d8c90e93d45e753db0fa4ec457d893e08cd1e0827b77896ee636b714706cbbb2 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-d8c90e93d45e753db0fa4ec457d893e08cd1e0827b77896ee636b714706cbbb2-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### `cloned\Jaebytes__TrueCharts\premium\traefik\26.10.19` — Jaebytes/TrueCharts

```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 1 charts
Downloading common from repo oci://tccr.io/truecharts
Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to resolve tccr.io/truecharts/common:20.3.12: Get "https://tccr.io/v2/truecharts/common/manifests/20.3.12": dial tcp: lookup tccr.io: no such host
Error: could not download oci://tccr.io/truecharts/common: failed to resolve tccr.io/truecharts/common:20.3.12: Get "https://tccr.io/v2/truecharts/common/manifests/20.3.12": dial tcp: lookup tccr.io: no such host
```

### `cloned\PnsDev__Truechart-catalog-archive\incubator\archivebox\0.7.2` — PnsDev/Truechart-catalog-archive

```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 1 charts
Downloading common from repo oci://tccr.io/truecharts
Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to resolve tccr.io/truecharts/common:17.2.30: Get "https://tccr.io/v2/truecharts/common/manifests/17.2.30": dial tcp: lookup tccr.io: no such host
Error: could not download oci://tccr.io/truecharts/common: failed to resolve tccr.io/truecharts/common:17.2.30: Get "https://tccr.io/v2/truecharts/common/manifests/17.2.30": dial tcp: lookup tccr.io: no such host
```

### `cloned\petersandor__tc-archive\scale-catalog\incubator\archivebox\0.7.2` — petersandor/tc-archive

```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 1 charts
Downloading common from repo oci://tccr.io/truecharts
Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to resolve tccr.io/truecharts/common:17.2.30: Get "https://tccr.io/v2/truecharts/common/manifests/17.2.30": dial tcp: lookup tccr.io: no such host
Error: could not download oci://tccr.io/truecharts/common: failed to resolve tccr.io/truecharts/common:17.2.30: Get "https://tccr.io/v2/truecharts/common/manifests/17.2.30": dial tcp: lookup tccr.io: no such host
```

### `cloned\adstanley__archive\scale-catalog\incubator\archivebox\0.7.2` — adstanley/archive

```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 1 charts
Downloading common from repo oci://tccr.io/truecharts
Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to resolve tccr.io/truecharts/common:17.2.30: Get "https://tccr.io/v2/truecharts/common/manifests/17.2.30": dial tcp: lookup tccr.io: no such host
Error: could not download oci://tccr.io/truecharts/common: failed to resolve tccr.io/truecharts/common:17.2.30: Get "https://tccr.io/v2/truecharts/common/manifests/17.2.30": dial tcp: lookup tccr.io: no such host
```

### `cloned\tdunlap607__kubeapps-fork\chart\kubeapps` — tdunlap607/kubeapps-fork

```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 3 charts
Downloading redis from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:20.6.1: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/20.6.1": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:20.6.1: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/20.6.1": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### `cloned\adynetro__truecharts\incubator\archivebox\0.7.2` — adynetro/truecharts

```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 1 charts
Downloading common from repo oci://tccr.io/truecharts
Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to resolve tccr.io/truecharts/common:17.2.30: Get "https://tccr.io/v2/truecharts/common/manifests/17.2.30": dial tcp: lookup tccr.io: no such host
Error: could not download oci://tccr.io/truecharts/common: failed to resolve tccr.io/truecharts/common:17.2.30: Get "https://tccr.io/v2/truecharts/common/manifests/17.2.30": dial tcp: lookup tccr.io: no such host
```

