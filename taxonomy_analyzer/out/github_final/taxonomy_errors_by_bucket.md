# Complete Errors By Taxonomy Bucket

Generated at: `2026-06-02 00:24:19 UTC`

Source catalog: `C:\Users\miabs\GolandProjects\helm-tests\helm_fetcher\results\github\catalog_by_project.json`

## `template.nil_pointer` (605)

### 1. `labring/sealos`

- Chart: `D:\helm_clones_github\labring__sealos\frontend\providers\kubepanel\deploy\charts\kubepanel-frontend`
- Source: `template`
- Values files: `D:\helm_clones_github\labring__sealos\frontend\providers\kubepanel\deploy\charts\kubepanel-frontend\kubepanel-frontend-values.yaml`
- Command: `helm template test D:\helm_clones_github\labring__sealos\frontend\providers\kubepanel\deploy\charts\kubepanel-frontend -f D:\helm_clones_github\labring__sealos\frontend\providers\kubepanel\deploy\charts\kubepanel-frontend\kubepanel-frontend-values.yaml`

```text
Error: kubepanel-frontend/templates/ingress.yaml:10:55
  executing "kubepanel-frontend/templates/ingress.yaml" at <include "kubepanel-frontend.cloudOrigin" .>:
    error calling include:
kubepanel-frontend/templates/_helpers.tpl:71:4
  executing "kubepanel-frontend.cloudOrigin" at <include "kubepanel-frontend.scheme" .>:
    error calling include:
kubepanel-frontend/templates/_helpers.tpl:44:27
  executing "kubepanel-frontend.scheme" at <.Values.kubepanelConfig.disableHttps>:
    nil pointer evaluating interface {}.disableHttps

Use --debug flag to render out invalid YAML
```

### 2. `microsoft/FluidFramework`

- Chart: `D:\helm_clones_github\microsoft__FluidFramework\server\routerlicious\kubernetes\routerlicious`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\microsoft__FluidFramework\server\routerlicious\kubernetes\routerlicious`

```text
Error: routerlicious/templates/scriptorium-deployment.yaml:21:28
  executing "routerlicious/templates/scriptorium-deployment.yaml" at <include (print $.Template.BasePath "/fluid-configmap.yaml") .>:
    error calling include:
routerlicious/templates/fluid-configmap.yaml:117:39
  executing "routerlicious/templates/fluid-configmap.yaml" at <.Values.alfred.api.patchRoot>:
    nil pointer evaluating interface {}.patchRoot

Use --debug flag to render out invalid YAML
```

### 3. `microsoft/mssql-docker`

- Chart: `D:\helm_clones_github\microsoft__mssql-docker\linux\rancher`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\microsoft__mssql-docker\linux\rancher`

```text
Error: sql-server-rancher/templates/secret.yaml:9:32
  executing "sql-server-rancher/templates/secret.yaml" at <.Values.mssql.sa.password>:
    nil pointer evaluating interface {}.password

Use --debug flag to render out invalid YAML
```

### 4. `medic/cht-core`

- Chart: `D:\helm_clones_github\medic__cht-core\scripts\build\helm`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\medic__cht-core\scripts\build\helm`

```text
Error: cht-chart/templates/sentinel/deployment.yaml:34:28
  executing "cht-chart/templates/sentinel/deployment.yaml" at <.Values.upstream_servers.docker_registry>:
    nil pointer evaluating interface {}.docker_registry

Use --debug flag to render out invalid YAML
```

### 5. `WeBankFinTech/Prophecis`

- Chart: `D:\helm_clones_github\WeBankFinTech__Prophecis\di\jobmonitor\charts\custom-serving`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\WeBankFinTech__Prophecis\di\jobmonitor\charts\custom-serving`

```text
Error: custom-serving/templates/deployment.yaml:17:34
  executing "custom-serving/templates/deployment.yaml" at <.Release.Time.Seconds>:
    nil pointer evaluating interface {}.Seconds

Use --debug flag to render out invalid YAML
```

### 6. `WeBankFinTech/Prophecis`

- Chart: `D:\helm_clones_github\WeBankFinTech__Prophecis\di\jobmonitor\charts\tfserving`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\WeBankFinTech__Prophecis\di\jobmonitor\charts\tfserving`

```text
Error: tensorflow-serving/templates/deployment.yaml:17:34
  executing "tensorflow-serving/templates/deployment.yaml" at <.Release.Time.Seconds>:
    nil pointer evaluating interface {}.Seconds

Use --debug flag to render out invalid YAML
```

### 7. `WeBankFinTech/Prophecis`

- Chart: `D:\helm_clones_github\WeBankFinTech__Prophecis\di\jobmonitor\charts\trtserving`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\WeBankFinTech__Prophecis\di\jobmonitor\charts\trtserving`

```text
Error: tensorrt-serving/templates/deployment.yaml:16:34
  executing "tensorrt-serving/templates/deployment.yaml" at <.Release.Time.Seconds>:
    nil pointer evaluating interface {}.Seconds

Use --debug flag to render out invalid YAML
```

### 8. `WeBankFinTech/Prophecis`

- Chart: `D:\helm_clones_github\WeBankFinTech__Prophecis\di\lcm\charts\custom-serving`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\WeBankFinTech__Prophecis\di\lcm\charts\custom-serving`

```text
Error: custom-serving/templates/deployment.yaml:17:34
  executing "custom-serving/templates/deployment.yaml" at <.Release.Time.Seconds>:
    nil pointer evaluating interface {}.Seconds

Use --debug flag to render out invalid YAML
```

### 9. `WeBankFinTech/Prophecis`

- Chart: `D:\helm_clones_github\WeBankFinTech__Prophecis\di\lcm\charts\tfserving`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\WeBankFinTech__Prophecis\di\lcm\charts\tfserving`

```text
Error: tensorflow-serving/templates/deployment.yaml:17:34
  executing "tensorflow-serving/templates/deployment.yaml" at <.Release.Time.Seconds>:
    nil pointer evaluating interface {}.Seconds

Use --debug flag to render out invalid YAML
```

### 10. `WeBankFinTech/Prophecis`

- Chart: `D:\helm_clones_github\WeBankFinTech__Prophecis\di\lcm\charts\trtserving`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\WeBankFinTech__Prophecis\di\lcm\charts\trtserving`

```text
Error: tensorrt-serving/templates/deployment.yaml:16:34
  executing "tensorrt-serving/templates/deployment.yaml" at <.Release.Time.Seconds>:
    nil pointer evaluating interface {}.Seconds

Use --debug flag to render out invalid YAML
```

### 11. `open-edge-platform/geti`

- Chart: `D:\helm_clones_github\open-edge-platform__geti\deploy\charts\impt\chart\charts\xpu-manager`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\open-edge-platform__geti\deploy\charts\impt\chart\charts\xpu-manager`

```text
Error: xpu-manager/templates/serviceaccount.yaml:1:14
  executing "xpu-manager/templates/serviceaccount.yaml" at <.Values.global.install_telemetry_stack>:
    nil pointer evaluating interface {}.install_telemetry_stack

Use --debug flag to render out invalid YAML
```

### 12. `Thakurvaibhav/k8s`

- Chart: `D:\helm_clones_github\Thakurvaibhav__k8s\charts\envoy-gateway`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Thakurvaibhav__k8s\charts\envoy-gateway`

```text
Error: envoy-gateway/templates/gatewayAPI/gateway-infra-shared-01.yaml:23:27
  executing "envoy-gateway/templates/gatewayAPI/gateway-infra-shared-01.yaml" at <.Values.gateway.defaultHttps.hostname>:
    nil pointer evaluating interface {}.hostname

Use --debug flag to render out invalid YAML
```

### 13. `Thakurvaibhav/k8s`

- Chart: `D:\helm_clones_github\Thakurvaibhav__k8s\charts\sealed-secrets`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Thakurvaibhav__k8s\charts\sealed-secrets`

```text
Error: sealed-secrets/templates/sealed-my-global-secret.yaml:11:24
  executing "sealed-secrets/templates/sealed-my-global-secret.yaml" at <.Values.globalExampleSecret.data.username>:
    nil pointer evaluating interface {}.username

Use --debug flag to render out invalid YAML
```

### 14. `hyperledger-bevel/bevel-operator-fabric`

- Chart: `D:\helm_clones_github\hyperledger-bevel__bevel-operator-fabric\charts\hlf-ca`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\hyperledger-bevel__bevel-operator-fabric\charts\hlf-ca`

```text
Error: hlf-ca/templates/traefikroute.yaml:1:13
  executing "hlf-ca/templates/traefikroute.yaml" at <.Values.traefik.hosts>:
    nil pointer evaluating interface {}.hosts

Use --debug flag to render out invalid YAML
```

### 15. `hyperledger-bevel/bevel-operator-fabric`

- Chart: `D:\helm_clones_github\hyperledger-bevel__bevel-operator-fabric\charts\hlf-ordnode`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\hyperledger-bevel__bevel-operator-fabric\charts\hlf-ordnode`

```text
Error: hlf-ordnode/templates/traefikroute.yaml:1:13
  executing "hlf-ordnode/templates/traefikroute.yaml" at <.Values.traefik.hosts>:
    nil pointer evaluating interface {}.hosts

Use --debug flag to render out invalid YAML
```

### 16. `vmware/secrets-manager`

- Chart: `D:\helm_clones_github\vmware__secrets-manager\helm-charts\0.28.1\charts\keystone`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\vmware__secrets-manager\helm-charts\0.28.1\charts\keystone`

```text
Error: keystone/templates/ServiceAccount.yaml:16:23
  executing "keystone/templates/ServiceAccount.yaml" at <.Values.global.vsecm.namespace>:
    nil pointer evaluating interface {}.vsecm

Use --debug flag to render out invalid YAML
```

### 17. `vmware/secrets-manager`

- Chart: `D:\helm_clones_github\vmware__secrets-manager\helm-charts\0.28.1\charts\safe`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\vmware__secrets-manager\helm-charts\0.28.1\charts\safe`

```text
Error: safe/templates/hook-preinstall-role.yaml:24:27
  executing "safe/templates/hook-preinstall-role.yaml" at <.Values.global.vsecm.namespace>:
    nil pointer evaluating interface {}.vsecm

Use --debug flag to render out invalid YAML
```

### 18. `vmware/secrets-manager`

- Chart: `D:\helm_clones_github\vmware__secrets-manager\helm-charts\0.28.1\charts\scout`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\vmware__secrets-manager\helm-charts\0.28.1\charts\scout`

```text
Error: scout/templates/ServiceAccount.yaml:17:23
  executing "scout/templates/ServiceAccount.yaml" at <.Values.global.vsecm.namespace>:
    nil pointer evaluating interface {}.vsecm

Use --debug flag to render out invalid YAML
```

### 19. `vmware/secrets-manager`

- Chart: `D:\helm_clones_github\vmware__secrets-manager\helm-charts\0.28.1\charts\sentinel`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\vmware__secrets-manager\helm-charts\0.28.1\charts\sentinel`

```text
Error: sentinel/templates/ServiceAccount.yaml:16:23
  executing "sentinel/templates/ServiceAccount.yaml" at <.Values.global.vsecm.namespace>:
    nil pointer evaluating interface {}.vsecm

Use --debug flag to render out invalid YAML
```

### 20. `vmware/secrets-manager`

- Chart: `D:\helm_clones_github\vmware__secrets-manager\helm-charts\0.28.1\charts\spire`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\vmware__secrets-manager\helm-charts\0.28.1\charts\spire`

```text
Error: spire/templates/validatingwebhookconfiguration-spire-server-spire-controller-manager-webhook.yaml:20:29
  executing "spire/templates/validatingwebhookconfiguration-spire-server-spire-controller-manager-webhook.yaml" at <.Values.global.spire.serverNamespace>:
    nil pointer evaluating interface {}.spire

Use --debug flag to render out invalid YAML
```

### 21. `open-edge-platform/edge-ai-libraries`

- Chart: `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\subchart\embedding\ovms`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\subchart\embedding\ovms`

```text
Error: ovmsEmbeddingService/templates/ovms-embed-configmap.yaml:18:35
  executing "ovmsEmbeddingService/templates/ovms-embed-configmap.yaml" at <.Values.global.modelDownload.serviceName>:
    nil pointer evaluating interface {}.serviceName

Use --debug flag to render out invalid YAML
```

### 22. `open-edge-platform/edge-ai-libraries`

- Chart: `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\subchart\llm\ovms`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\subchart\llm\ovms`

```text
Error: ovmsService/templates/ovms-deployment.yaml:15:19
  executing "ovmsService/templates/ovms-deployment.yaml" at <.Values.global.affinity.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 23. `open-edge-platform/edge-ai-libraries`

- Chart: `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\subchart\llm\tgi`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\subchart\llm\tgi`

```text
Error: tgiService/templates/pvc.yaml:14:25
  executing "tgiService/templates/pvc.yaml" at <.Values.global.tgi_pvc.size>:
    nil pointer evaluating interface {}.size

Use --debug flag to render out invalid YAML
```

### 24. `open-edge-platform/edge-ai-libraries`

- Chart: `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\subchart\llm\vllm`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\subchart\llm\vllm`

```text
Error: vllmService/templates/pvc.yaml:14:25
  executing "vllmService/templates/pvc.yaml" at <.Values.global.vllm_pvc.size>:
    nil pointer evaluating interface {}.size

Use --debug flag to render out invalid YAML
```

### 25. `open-edge-platform/edge-ai-libraries`

- Chart: `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\subchart\minioserver`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\subchart\minioserver`

```text
Error: minioServer/templates/pvc.yaml:14:25
  executing "minioServer/templates/pvc.yaml" at <.Values.global.minio_pvc.size>:
    nil pointer evaluating interface {}.size

Use --debug flag to render out invalid YAML
```

### 26. `open-edge-platform/edge-ai-libraries`

- Chart: `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\subchart\reranker`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\subchart\reranker`

```text
Error: reranker/templates/pvc.yaml:14:25
  executing "reranker/templates/pvc.yaml" at <.Values.global.reranker_pvc.size>:
    nil pointer evaluating interface {}.size

Use --debug flag to render out invalid YAML
```

### 27. `open-edge-platform/edge-ai-libraries`

- Chart: `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\document-summarization\chart\subchart\backend\docsum-api`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\document-summarization\chart\subchart\backend\docsum-api`

```text
Error: docsum-api/templates/docsum-api-deployment.yaml:19:28
  executing "docsum-api/templates/docsum-api-deployment.yaml" at <.Values.global.docSum.image.repository>:
    nil pointer evaluating interface {}.docSum

Use --debug flag to render out invalid YAML
```

### 28. `open-edge-platform/edge-ai-libraries`

- Chart: `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\document-summarization\chart\subchart\ovms\ovms-service`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\document-summarization\chart\subchart\ovms\ovms-service`

```text
Error: ovms-service/templates/ovm-deployment.yaml:5:16
  executing "ovms-service/templates/ovm-deployment.yaml" at <.Values.global.keeppvc>:
    nil pointer evaluating interface {}.keeppvc

Use --debug flag to render out invalid YAML
```

### 29. `open-edge-platform/edge-ai-libraries`

- Chart: `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\video-search-and-summarization\chart\subchart\multimodal-embedding-ms`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\video-search-and-summarization\chart\subchart\multimodal-embedding-ms`

```text
Error: multimodalembeddingms/templates/pvc.yaml:1:38
  executing "multimodalembeddingms/templates/pvc.yaml" at <.Values.global.devices.multimodalEmbedding.device>:
    nil pointer evaluating interface {}.multimodalEmbedding

Use --debug flag to render out invalid YAML
```

### 30. `open-edge-platform/edge-ai-libraries`

- Chart: `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\video-search-and-summarization\chart\subchart\postgresql`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\video-search-and-summarization\chart\subchart\postgresql`

```text
Error: postgresql/templates/postgresql-deployment.yaml:39:31
  executing "postgresql/templates/postgresql-deployment.yaml" at <.Values.global.proxy.http_proxy>:
    nil pointer evaluating interface {}.http_proxy

Use --debug flag to render out invalid YAML
```

### 31. `open-edge-platform/edge-ai-libraries`

- Chart: `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\video-search-and-summarization\chart\subchart\vdms-dataprep`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\video-search-and-summarization\chart\subchart\vdms-dataprep`

```text
Error: vdmsdataprep/templates/deployment.yaml:1:38
  executing "vdmsdataprep/templates/deployment.yaml" at <.Values.global.devices.vdmsDataprep.device>:
    nil pointer evaluating interface {}.vdmsDataprep

Use --debug flag to render out invalid YAML
```

### 32. `open-edge-platform/edge-ai-libraries`

- Chart: `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\video-search-and-summarization\chart\subchart\vllm`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\video-search-and-summarization\chart\subchart\vllm`

```text
Error: vllm-server/templates/deployment.yaml:59:31
  executing "vllm-server/templates/deployment.yaml" at <$global.proxy.http_proxy>:
    nil pointer evaluating interface {}.http_proxy

Use --debug flag to render out invalid YAML
```

### 33. `aws-samples/amazon-eks-machine-learning-with-terraform-and-kubeflow`

- Chart: `D:\helm_clones_github\aws-samples__amazon-eks-machine-learning-with-terraform-and-kubeflow\charts\pv-efs`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\aws-samples__amazon-eks-machine-learning-with-terraform-and-kubeflow\charts\pv-efs`

```text
Error: pv-efs/templates/storage-class.yaml:4:18
  executing "pv-efs/templates/storage-class.yaml" at <.Values.efs.class_name>:
    nil pointer evaluating interface {}.class_name

Use --debug flag to render out invalid YAML
```

### 34. `aws-samples/amazon-eks-machine-learning-with-terraform-and-kubeflow`

- Chart: `D:\helm_clones_github\aws-samples__amazon-eks-machine-learning-with-terraform-and-kubeflow\charts\pv-fsx`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\aws-samples__amazon-eks-machine-learning-with-terraform-and-kubeflow\charts\pv-fsx`

```text
Error: pv-fsx/templates/storage-class.yaml:4:18
  executing "pv-fsx/templates/storage-class.yaml" at <.Values.fsx.class_name>:
    nil pointer evaluating interface {}.class_name

Use --debug flag to render out invalid YAML
```

### 35. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\adminutils`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\adminutils`

```text
Error: adminutils/templates/hpa.yaml:1:14
  executing "adminutils/templates/hpa.yaml" at <.Values.autoscaling.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 36. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\analytics`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\analytics`

```text
Error: analytics/templates/serviceMonitor.yaml:1:14
  executing "analytics/templates/serviceMonitor.yaml" at <.Values.serviceMonitor.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 37. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\apimanager`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\apimanager`

```text
Error: apimanager/templates/recordingRules.yaml:1:14
  executing "apimanager/templates/recordingRules.yaml" at <.Values.serviceMonitor.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 38. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\apimanagerecho`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\apimanagerecho`

```text
Error: apimanagerecho/templates/hpa.yaml:1:14
  executing "apimanagerecho/templates/hpa.yaml" at <.Values.autoscaling.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 39. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\assessment`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\assessment`

```text
Error: assessment/templates/hpa.yaml:1:14
  executing "assessment/templates/hpa.yaml" at <.Values.autoscaling.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 40. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\cert`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\cert`

```text
Error: cert/templates/hpa.yaml:1:14
  executing "cert/templates/hpa.yaml" at <.Values.autoscaling.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 41. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\certregistry`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\certregistry`

```text
Error: certregistry/templates/hpa.yaml:1:14
  executing "certregistry/templates/hpa.yaml" at <.Values.autoscaling.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 42. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\content`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\content`

```text
Error: content/templates/hpa.yaml:1:14
  executing "content/templates/hpa.yaml" at <.Values.autoscaling.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 43. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\dhiti`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\dhiti`

```text
Error: dhiti/templates/deployment.yaml:13:28
  executing "dhiti/templates/deployment.yaml" at <.Values.strategy.maxsurge>:
    nil pointer evaluating interface {}.maxsurge

Use --debug flag to render out invalid YAML
```

### 44. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\dial`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\dial`

```text
Error: dial/templates/hpa.yaml:1:14
  executing "dial/templates/hpa.yaml" at <.Values.autoscaling.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 45. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\discussionsmw`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\discussionsmw`

```text
Error: discussionsmw/templates/hpa.yaml:1:14
  executing "discussionsmw/templates/hpa.yaml" at <.Values.autoscaling.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 46. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\enc`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\enc`

```text
Error: enc/templates/hpa.yaml:1:14
  executing "enc/templates/hpa.yaml" at <.Values.autoscaling.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 47. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\gotenberg`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\gotenberg`

```text
Error: gotenberg/templates/hpa.yaml:1:14
  executing "gotenberg/templates/hpa.yaml" at <.Values.autoscaling.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 48. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\gql`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\gql`

```text
Error: gql/templates/hpa.yaml:1:14
  executing "gql/templates/hpa.yaml" at <.Values.autoscaling.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 49. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\groups`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\groups`

```text
Error: groups/templates/hpa.yaml:1:14
  executing "groups/templates/hpa.yaml" at <.Values.autoscaling.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 50. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\hawkeyesuperset`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\hawkeyesuperset`

```text
Error: superset/templates/service.yaml:15:18
  executing "superset/templates/service.yaml" at <.Values.service.type>:
    nil pointer evaluating interface {}.type

Use --debug flag to render out invalid YAML
```

### 51. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\inbound`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\inbound`

```text
Error: inbound/templates/hpa.yaml:1:14
  executing "inbound/templates/hpa.yaml" at <.Values.autoscaling.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 52. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\knowledgemw`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\knowledgemw`

```text
Error: knowledgemw/templates/hpa.yaml:1:14
  executing "knowledgemw/templates/hpa.yaml" at <.Values.autoscaling.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 53. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\learner`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\learner`

```text
Error: learner/templates/hpa.yaml:1:14
  executing "learner/templates/hpa.yaml" at <.Values.autoscaling.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 54. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\lms`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\lms`

```text
Error: lms/templates/hpa.yaml:1:14
  executing "lms/templates/hpa.yaml" at <.Values.autoscaling.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 55. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\ml-core-service`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\ml-core-service`

```text
Error: ml-core-service/templates/hpa.yaml:1:14
  executing "ml-core-service/templates/hpa.yaml" at <.Values.autoscaling.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 56. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\ml-projects-service`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\ml-projects-service`

```text
Error: ml-projects-service/templates/hpa.yaml:1:14
  executing "ml-projects-service/templates/hpa.yaml" at <.Values.autoscaling.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 57. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\ml-reports-service`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\ml-reports-service`

```text
Error: ml-reports-service/templates/hpa.yaml:1:14
  executing "ml-reports-service/templates/hpa.yaml" at <.Values.autoscaling.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 58. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\ml-survey-service`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\ml-survey-service`

```text
Error: ml-survey-service/templates/hpa.yaml:1:14
  executing "ml-survey-service/templates/hpa.yaml" at <.Values.autoscaling.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 59. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\nginx-public-ingress`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\nginx-public-ingress`

```text
Error: nginx-public-ingress/templates/serviceMonitor.yml:1:14
  executing "nginx-public-ingress/templates/serviceMonitor.yml" at <.Values.serviceMonitor.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 60. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\nodebb`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\nodebb`

```text
Error: nodebb/templates/hpa.yaml:1:14
  executing "nodebb/templates/hpa.yaml" at <.Values.autoscaling.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 61. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\notification`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\notification`

```text
Error: notification/templates/hpa.yaml:1:14
  executing "notification/templates/hpa.yaml" at <.Values.autoscaling.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 62. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\odk`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\odk`

```text
Error: odk/templates/hpa.yaml:1:14
  executing "odk/templates/hpa.yaml" at <.Values.autoscaling.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 63. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\orchestrator`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\orchestrator`

```text
Error: orchestrator/templates/hpa.yaml:1:14
  executing "orchestrator/templates/hpa.yaml" at <.Values.autoscaling.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 64. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\outbound`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\outbound`

```text
Error: outbound/templates/hpa.yaml:1:14
  executing "outbound/templates/hpa.yaml" at <.Values.autoscaling.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 65. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\player`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\player`

```text
Error: player/templates/hpa.yaml:1:14
  executing "player/templates/hpa.yaml" at <.Values.autoscaling.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 66. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\print`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\print`

```text
Error: print/templates/hpa.yaml:1:14
  executing "print/templates/hpa.yaml" at <.Values.autoscaling.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 67. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\report`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\report`

```text
Error: report/templates/hpa.yaml:1:14
  executing "report/templates/hpa.yaml" at <.Values.autoscaling.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 68. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\search`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\search`

```text
Error: search/templates/hpa.yaml:1:14
  executing "search/templates/hpa.yaml" at <.Values.autoscaling.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 69. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\taxonomy`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\taxonomy`

```text
Error: taxonomy/templates/hpa.yaml:1:14
  executing "taxonomy/templates/hpa.yaml" at <.Values.autoscaling.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 70. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\telemetry`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\telemetry`

```text
Error: telemetry/templates/hpa.yaml:1:14
  executing "telemetry/templates/hpa.yaml" at <.Values.autoscaling.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 71. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\telemetry-dp-logstash`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\telemetry-dp-logstash`

```text
Error: telemetry-logstash-dock/templates/deployment.yaml:13:28
  executing "telemetry-logstash-dock/templates/deployment.yaml" at <.Values.strategy.maxsurge>:
    nil pointer evaluating interface {}.maxsurge

Use --debug flag to render out invalid YAML
```

### 72. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\transformer`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\transformer`

```text
Error: transformer/templates/hpa.yaml:1:14
  executing "transformer/templates/hpa.yaml" at <.Values.autoscaling.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 73. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\uci`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\uci`

```text
Error: uci/templates/hpa.yaml:1:14
  executing "uci/templates/hpa.yaml" at <.Values.autoscaling.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 74. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\userorg`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\userorg`

```text
Error: userorg/templates/hpa.yaml:1:14
  executing "userorg/templates/hpa.yaml" at <.Values.autoscaling.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 75. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\charts\certmanager`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\charts\certmanager`

```text
Error: certmanager/templates/serviceaccount.yaml:3:14
  executing "certmanager/templates/serviceaccount.yaml" at <.Values.global.imagePullSecrets>:
    nil pointer evaluating interface {}.imagePullSecrets

Use --debug flag to render out invalid YAML
```

### 76. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\charts\galley`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\charts\galley`

```text
Error: galley/templates/serviceaccount.yaml:3:14
  executing "galley/templates/serviceaccount.yaml" at <.Values.global.imagePullSecrets>:
    nil pointer evaluating interface {}.imagePullSecrets

Use --debug flag to render out invalid YAML
```

### 77. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\charts\gateways`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\charts\gateways`

```text
Error: gateways/templates/serviceaccount.yaml:6:8
  executing "gateways/templates/serviceaccount.yaml" at <$.Values.global.imagePullSecrets>:
    nil pointer evaluating interface {}.imagePullSecrets

Use --debug flag to render out invalid YAML
```

### 78. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\charts\grafana`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\charts\grafana`

```text
Error: grafana/templates/tests/test-grafana-connection.yaml:1:14
  executing "grafana/templates/tests/test-grafana-connection.yaml" at <.Values.global.enableHelmTest>:
    nil pointer evaluating interface {}.enableHelmTest

Use --debug flag to render out invalid YAML
```

### 79. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\charts\istiocoredns`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\charts\istiocoredns`

```text
Error: istiocoredns/templates/serviceaccount.yaml:3:14
  executing "istiocoredns/templates/serviceaccount.yaml" at <.Values.global.imagePullSecrets>:
    nil pointer evaluating interface {}.imagePullSecrets

Use --debug flag to render out invalid YAML
```

### 80. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\charts\kiali`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\charts\kiali`

```text
Error: kiali/templates/tests/test-kiali-connection.yaml:1:14
  executing "kiali/templates/tests/test-kiali-connection.yaml" at <.Values.global.enableHelmTest>:
    nil pointer evaluating interface {}.enableHelmTest

Use --debug flag to render out invalid YAML
```

### 81. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\charts\mixer`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\charts\mixer`

```text
Error: mixer/templates/serviceaccount.yaml:4:14
  executing "mixer/templates/serviceaccount.yaml" at <.Values.global.imagePullSecrets>:
    nil pointer evaluating interface {}.imagePullSecrets

Use --debug flag to render out invalid YAML
```

### 82. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\charts\nodeagent`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\charts\nodeagent`

```text
Error: nodeagent/templates/serviceaccount.yaml:3:14
  executing "nodeagent/templates/serviceaccount.yaml" at <.Values.global.imagePullSecrets>:
    nil pointer evaluating interface {}.imagePullSecrets

Use --debug flag to render out invalid YAML
```

### 83. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\charts\pilot`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\charts\pilot`

```text
Error: pilot/templates/serviceaccount.yaml:3:14
  executing "pilot/templates/serviceaccount.yaml" at <.Values.global.imagePullSecrets>:
    nil pointer evaluating interface {}.imagePullSecrets

Use --debug flag to render out invalid YAML
```

### 84. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\charts\prometheus`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\charts\prometheus`

```text
Error: prometheus/templates/tests/test-prometheus-connection.yaml:1:14
  executing "prometheus/templates/tests/test-prometheus-connection.yaml" at <.Values.global.enableHelmTest>:
    nil pointer evaluating interface {}.enableHelmTest

Use --debug flag to render out invalid YAML
```

### 85. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\charts\security`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\charts\security`

```text
Error: security/templates/tests/test-citadel-connection.yaml:1:14
  executing "security/templates/tests/test-citadel-connection.yaml" at <.Values.global.enableHelmTest>:
    nil pointer evaluating interface {}.enableHelmTest

Use --debug flag to render out invalid YAML
```

### 86. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\charts\sidecarInjectorWebhook`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\charts\sidecarInjectorWebhook`

```text
Error: sidecarInjectorWebhook/templates/serviceaccount.yaml:3:14
  executing "sidecarInjectorWebhook/templates/serviceaccount.yaml" at <.Values.global.imagePullSecrets>:
    nil pointer evaluating interface {}.imagePullSecrets

Use --debug flag to render out invalid YAML
```

### 87. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\charts\tracing`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\charts\tracing`

```text
Error: tracing/templates/tests/test-tracing-connection.yaml:1:14
  executing "tracing/templates/tests/test-tracing-connection.yaml" at <.Values.global.enableHelmTest>:
    nil pointer evaluating interface {}.enableHelmTest

Use --debug flag to render out invalid YAML
```

### 88. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\logging\filebeat`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\logging\filebeat`

```text
Error: filebeat/templates/graylog-service.yaml:23:21
  executing "filebeat/templates/graylog-service.yaml" at <.Values.graylog.hosts>:
    nil pointer evaluating interface {}.hosts

Use --debug flag to render out invalid YAML
```

### 89. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\monitoring\azure-ambari-prometheus-exporter`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\monitoring\azure-ambari-prometheus-exporter`

```text
Error: azure-ambari-prometheus-exporter/templates/deployment.yaml:28:30
  executing "azure-ambari-prometheus-exporter/templates/deployment.yaml" at <.Values.env.ambari.rm_url>:
    nil pointer evaluating interface {}.ambari

Use --debug flag to render out invalid YAML
```

### 90. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\monitoring\kafka-lag-exporter`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\monitoring\kafka-lag-exporter`

```text
Error: kafka-lag-exporter/templates/serviceMonitor.yml:1:14
  executing "kafka-lag-exporter/templates/serviceMonitor.yml" at <.Values.serviceMonitor.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 91. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\monitoring\kafka-topic-exporter`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\monitoring\kafka-topic-exporter`

```text
Error: kafka-topic-exporter/templates/deployment.yaml:28:30
  executing "kafka-topic-exporter/templates/deployment.yaml" at <.Values.env.kafka.host>:
    nil pointer evaluating interface {}.kafka

Use --debug flag to render out invalid YAML
```

### 92. `Borjis131/docker-open5gs`

- Chart: `D:\helm_clones_github\Borjis131__docker-open5gs\helm\open5gs`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Borjis131__docker-open5gs\helm\open5gs`

```text
Error: template: open5gs/charts/upf/templates/_helpers.tpl:15:3: executing "upf.config" at <tpl (.Files.Get "configs/upf.yaml") .>: error calling tpl: error during tpl function execution for "logger:\r\n  file:\r\n    path: /var/log/open5gs/upf.log\r\n\r\nglobal:\r\n\r\nupf:\r\n  pfcp:\r\n    server:\r\n      - dev: eth0\r\n        {{- if eq .Values.services.pfcp.type \"ClusterIP\" }}\r\n        advertise: {{ include \"common.names.fullname\" . }}-service-pfcp\r\n        {{- end }}\r\n    client:\r\n      smf:\r\n        {{- if eq .Release.Name \"smf\" }}\r\n        - address: smf-service-pfcp\r\n        {{- else }}\r\n        - address: {{ .Release.Name }}-smf-service-pfcp\r\n        {{- end }}\r\n  gtpu:\r\n    server:\r\n      - dev: eth0\r\n        {{- if eq .Values.services.gtpu.type \"LoadBalancer\" }}\r\n        advertise: {{ .Values.services.gtpu.loadBalancerIP }}\r\n        {{- end }}\r\n  session:\r\n    - subnet: {{ .Values.global.mobileNetwork.dataNetwork.subnet }}\r\n      gateway: {{ .Values.global.mobileNetwork.dataNetwork.gateway }}\r\n      {{- if .Values.global.mobileNetwork.dataNetwork.dnn }}\r\n      dnn: {{ .Values.global.mobileNetwork.dataNetwork.dnn }}\r\n      {{- end }}\r\n": template: gotpl:11:25: executing "gotpl" at <.Values.services.pfcp.type>: nil pointer evaluating interface {}.type

Use --debug flag to render out invalid YAML
```

### 93. `Borjis131/docker-open5gs`

- Chart: `D:\helm_clones_github\Borjis131__docker-open5gs\helm\upf`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Borjis131__docker-open5gs\helm\upf`

```text
Error: template: upf/templates/_helpers.tpl:15:3: executing "upf.config" at <tpl (.Files.Get "configs/upf.yaml") .>: error calling tpl: error during tpl function execution for "logger:\r\n  file:\r\n    path: /var/log/open5gs/upf.log\r\n\r\nglobal:\r\n\r\nupf:\r\n  pfcp:\r\n    server:\r\n      - dev: eth0\r\n        {{- if eq .Values.services.pfcp.type \"ClusterIP\" }}\r\n        advertise: {{ include \"common.names.fullname\" . }}-service-pfcp\r\n        {{- end }}\r\n    client:\r\n      smf:\r\n        {{- if eq .Release.Name \"smf\" }}\r\n        - address: smf-service-pfcp\r\n        {{- else }}\r\n        - address: {{ .Release.Name }}-smf-service-pfcp\r\n        {{- end }}\r\n  gtpu:\r\n    server:\r\n      - dev: eth0\r\n        {{- if eq .Values.services.gtpu.type \"LoadBalancer\" }}\r\n        advertise: {{ .Values.services.gtpu.loadBalancerIP }}\r\n        {{- end }}\r\n  session:\r\n    - subnet: {{ .Values.global.mobileNetwork.dataNetwork.subnet }}\r\n      gateway: {{ .Values.global.mobileNetwork.dataNetwork.gateway }}\r\n      {{- if .Values.global.mobileNetwork.dataNetwork.dnn }}\r\n      dnn: {{ .Values.global.mobileNetwork.dataNetwork.dnn }}\r\n      {{- end }}\r\n": template: gotpl:11:25: executing "gotpl" at <.Values.services.pfcp.type>: nil pointer evaluating interface {}.type

Use --debug flag to render out invalid YAML
```

### 94. `tmforum-oda/oda-canvas`

- Chart: `D:\helm_clones_github\tmforum-oda__oda-canvas\charts\api-operator-istio`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tmforum-oda__oda-canvas\charts\api-operator-istio`

```text
Error: api-operator-istio/templates/deployment.yaml:21:20
  executing "api-operator-istio/templates/deployment.yaml" at <.Values.global.imagePullSecrets>:
    nil pointer evaluating interface {}.imagePullSecrets

Use --debug flag to render out invalid YAML
```

### 95. `tmforum-oda/oda-canvas`

- Chart: `D:\helm_clones_github\tmforum-oda__oda-canvas\charts\apigee-gateway`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tmforum-oda__oda-canvas\charts\apigee-gateway`

```text
Error: api-operator-apigee/templates/Deployment.yaml:17:20
  executing "api-operator-apigee/templates/Deployment.yaml" at <.Values.global.imagePullSecrets>:
    nil pointer evaluating interface {}.imagePullSecrets

Use --debug flag to render out invalid YAML
```

### 96. `tmforum-oda/oda-canvas`

- Chart: `D:\helm_clones_github\tmforum-oda__oda-canvas\charts\apisix-gateway`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tmforum-oda__oda-canvas\charts\apisix-gateway`

```text
Error: api-operator-apisix/templates/DisableIstioLB.yaml:13:20
  executing "api-operator-apisix/templates/DisableIstioLB.yaml" at <.Values.global.imagePullSecrets>:
    nil pointer evaluating interface {}.imagePullSecrets

Use --debug flag to render out invalid YAML
```

### 97. `tmforum-oda/oda-canvas`

- Chart: `D:\helm_clones_github\tmforum-oda__oda-canvas\charts\canvas-info-service`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tmforum-oda__oda-canvas\charts\canvas-info-service`

```text
Error: canvas-info-service/templates/mongodb-deployment.yaml:17:20
  executing "canvas-info-service/templates/mongodb-deployment.yaml" at <.Values.global.imagePullSecrets>:
    nil pointer evaluating interface {}.imagePullSecrets

Use --debug flag to render out invalid YAML
```

### 98. `tmforum-oda/oda-canvas`

- Chart: `D:\helm_clones_github\tmforum-oda__oda-canvas\charts\canvas-vault`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tmforum-oda__oda-canvas\charts\canvas-vault`

```text
Error: canvas-vault/templates/post-install-hook.yaml:26:20
  executing "canvas-vault/templates/post-install-hook.yaml" at <.Values.global.imagePullSecrets>:
    nil pointer evaluating interface {}.imagePullSecrets

Use --debug flag to render out invalid YAML
```

### 99. `tmforum-oda/oda-canvas`

- Chart: `D:\helm_clones_github\tmforum-oda__oda-canvas\charts\component-operator`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tmforum-oda__oda-canvas\charts\component-operator`

```text
Error: component-operator/templates/deployment.yaml:21:20
  executing "component-operator/templates/deployment.yaml" at <.Values.global.imagePullSecrets>:
    nil pointer evaluating interface {}.imagePullSecrets

Use --debug flag to render out invalid YAML
```

### 100. `tmforum-oda/oda-canvas`

- Chart: `D:\helm_clones_github\tmforum-oda__oda-canvas\charts\dependentapi-simple-operator`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tmforum-oda__oda-canvas\charts\dependentapi-simple-operator`

```text
Error: dependentapi-simple-operator/templates/depapi-operator-deployment.yaml:18:20
  executing "dependentapi-simple-operator/templates/depapi-operator-deployment.yaml" at <.Values.global.imagePullSecrets>:
    nil pointer evaluating interface {}.imagePullSecrets

Use --debug flag to render out invalid YAML
```

### 101. `tmforum-oda/oda-canvas`

- Chart: `D:\helm_clones_github\tmforum-oda__oda-canvas\charts\identityconfig-operator-keycloak`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tmforum-oda__oda-canvas\charts\identityconfig-operator-keycloak`

```text
Error: identityconfig-operator-keycloak/templates/deployment.yaml:21:20
  executing "identityconfig-operator-keycloak/templates/deployment.yaml" at <.Values.global.imagePullSecrets>:
    nil pointer evaluating interface {}.imagePullSecrets

Use --debug flag to render out invalid YAML
```

### 102. `tmforum-oda/oda-canvas`

- Chart: `D:\helm_clones_github\tmforum-oda__oda-canvas\charts\kong-gateway`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tmforum-oda__oda-canvas\charts\kong-gateway`

```text
Error: api-operator-kong/templates/DisableIstioLB.yaml:13:20
  executing "api-operator-kong/templates/DisableIstioLB.yaml" at <.Values.global.imagePullSecrets>:
    nil pointer evaluating interface {}.imagePullSecrets

Use --debug flag to render out invalid YAML
```

### 103. `boozallen/aissemble`

- Chart: `D:\helm_clones_github\boozallen__aissemble\extensions\extensions-helm\aissemble-spark-operator-chart`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\boozallen__aissemble\extensions\extensions-helm\aissemble-spark-operator-chart`

```text
Error: aissemble-spark-operator-chart/templates/ivy-pvc.yaml:1:14
  executing "aissemble-spark-operator-chart/templates/ivy-pvc.yaml" at <.Values.ivyCache.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 104. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-agent-config`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-agent-config`

```text
Error: adp-agent-config/templates/deployment.yaml:9:18
  executing "adp-agent-config/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 105. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-agent-exec`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-agent-exec`

```text
Error: adp-agent-exec/templates/deployment.yaml:9:18
  executing "adp-agent-exec/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 106. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-ai-gateway`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-ai-gateway`

```text
Error: adp-ai-gateway/templates/deployment.yaml:7:18
  executing "adp-ai-gateway/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 107. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-apex`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-apex`

```text
Error: adp-apex/templates/apex.yaml:8:18
  executing "adp-apex/templates/apex.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 108. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-app-channel-msg`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-app-channel-msg`

```text
Error: adp-app-channel-msg/templates/deployment.yaml:7:18
  executing "adp-app-channel-msg/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 109. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-app-channel-proxy-svr`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-app-channel-proxy-svr`

```text
Error: adp-app-channel-proxy-svr/templates/deployment.yaml:7:18
  executing "adp-app-channel-proxy-svr/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 110. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-app-channel-token`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-app-channel-token`

```text
Error: adp-app-channel-token/templates/deployment.yaml:7:18
  executing "adp-app-channel-token/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 111. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-app-config`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-app-config`

```text
Error: adp-app-config/templates/deployment.yaml:7:18
  executing "adp-app-config/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 112. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-app-evaluate`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-app-evaluate`

```text
Error: adp-app-evaluate/templates/deployment.yaml:7:18
  executing "adp-app-evaluate/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 113. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-app-memory`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-app-memory`

```text
Error: adp-app-memory/templates/deployment.yaml:9:18
  executing "adp-app-memory/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 114. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-app-shorturl`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-app-shorturl`

```text
Error: adp-app-shorturl/templates/deployment.yaml:7:18
  executing "adp-app-shorturl/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 115. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-chat-manage`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-chat-manage`

```text
Error: adp-chat-manage/templates/ingress.yaml:15:28
  executing "adp-chat-manage/templates/ingress.yaml" at <.Values.global.scheme>:
    nil pointer evaluating interface {}.scheme

Use --debug flag to render out invalid YAML
```

### 116. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-crawler`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-crawler`

```text
Error: adp-crawler/templates/deployment.yaml:9:18
  executing "adp-crawler/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 117. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-file-converter`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-file-converter`

```text
Error: adp-file-converter/templates/deployment-fileconverter.yaml:9:18
  executing "adp-file-converter/templates/deployment-fileconverter.yaml" at <.Values.global.customLabels>:
    nil pointer evaluating interface {}.customLabels

Use --debug flag to render out invalid YAML
```

### 118. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-kb-config`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-kb-config`

```text
Error: adp-kb-config/templates/deployment.yaml:9:18
  executing "adp-kb-config/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 119. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-kb-retrieval`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-kb-retrieval`

```text
Error: adp-kb-retrieval/templates/deployment.yaml:7:18
  executing "adp-kb-retrieval/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 120. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-kb-vdb-proxy`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-kb-vdb-proxy`

```text
Error: adp-kb-vdb-proxy/templates/deployment.yaml:7:18
  executing "adp-kb-vdb-proxy/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 121. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-markmap-service`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-markmap-service`

```text
Error: adp-markmap-service/templates/deployment.yaml:9:18
  executing "adp-markmap-service/templates/deployment.yaml" at <.Values.global.customLabels>:
    nil pointer evaluating interface {}.customLabels

Use --debug flag to render out invalid YAML
```

### 122. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-model-token-count`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-model-token-count`

```text
Error: adp-model-token-count/templates/deployment.yaml:9:18
  executing "adp-model-token-count/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 123. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-platform-authenticator`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-platform-authenticator`

```text
Error: adp-platform-authenticator/templates/deployment.yaml:7:18
  executing "adp-platform-authenticator/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 124. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-platform-charger`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-platform-charger`

```text
Error: adp-platform-charger/templates/deployment.yaml:7:18
  executing "adp-platform-charger/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 125. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-platform-content-moderation`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-platform-content-moderation`

```text
Error: adp-platform-content-moderation/templates/deployment.yaml:7:18
  executing "adp-platform-content-moderation/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 126. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-platform-manager`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-platform-manager`

```text
Error: adp-platform-manager/templates/deployment.yaml:7:18
  executing "adp-platform-manager/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 127. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-platform-metrology`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-platform-metrology`

```text
Error: adp-platform-metrology/templates/deployment.yaml:7:18
  executing "adp-platform-metrology/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 128. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-platform-op`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-platform-op`

```text
Error: adp-platform-op/templates/deployment.yaml:7:18
  executing "adp-platform-op/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 129. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-platform-permission-config`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-platform-permission-config`

```text
Error: adp-platform-permission-config/templates/deployment.yaml:7:18
  executing "adp-platform-permission-config/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 130. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-plugin-code-interpret`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-plugin-code-interpret`

```text
Error: adp-plugin-code-interpret/templates/deployment.yaml:9:18
  executing "adp-plugin-code-interpret/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 131. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-plugin-code-kernel`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-plugin-code-kernel`

```text
Error: adp-plugin-code-kernel/templates/deployment.yaml:9:18
  executing "adp-plugin-code-kernel/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 132. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-plugin-config`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-plugin-config`

```text
Error: adp-plugin-config/templates/deployment.yaml:9:18
  executing "adp-plugin-config/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 133. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-plugin-cos-util`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-plugin-cos-util`

```text
Error: adp-plugin-cos-util/templates/deployment.yaml:8:18
  executing "adp-plugin-cos-util/templates/deployment.yaml" at <.Values.global.customLabels>:
    nil pointer evaluating interface {}.customLabels

Use --debug flag to render out invalid YAML
```

### 134. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-plugin-exec`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-plugin-exec`

```text
Error: adp-plugin-exec/templates/deployment.yaml:9:18
  executing "adp-plugin-exec/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 135. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-plugin-pod-schedule`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-plugin-pod-schedule`

```text
Error: adp-plugin-pod-schedule/templates/deployment.yaml:7:18
  executing "adp-plugin-pod-schedule/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 136. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-pyright`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-pyright`

```text
Error: adp-pyright/templates/deployment-pyright.yaml:9:18
  executing "adp-pyright/templates/deployment-pyright.yaml" at <.Values.global.customLabels>:
    nil pointer evaluating interface {}.customLabels

Use --debug flag to render out invalid YAML
```

### 137. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-rag-exec`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-rag-exec`

```text
Error: adp-rag-exec/templates/deployment.yaml:7:18
  executing "adp-rag-exec/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 138. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-resource-gallery`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-resource-gallery`

```text
Error: adp-resource-gallery/templates/deployment.yaml:7:18
  executing "adp-resource-gallery/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 139. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-web-app`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-web-app`

```text
Error: adp-web-app/templates/deployment-webappchatbot.yaml:9:18
  executing "adp-web-app/templates/deployment-webappchatbot.yaml" at <.Values.global.customLabels>:
    nil pointer evaluating interface {}.customLabels

Use --debug flag to render out invalid YAML
```

### 140. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-widget-config`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-widget-config`

```text
Error: adp-widget-config/templates/deployment.yaml:7:18
  executing "adp-widget-config/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 141. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-workflow-config`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-workflow-config`

```text
Error: adp-workflow-config/templates/deployment.yaml:9:18
  executing "adp-workflow-config/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 142. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-workflow-exec`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\adp-workflow-exec`

```text
Error: adp-workflow-exec/templates/deployment.yaml:9:18
  executing "adp-workflow-exec/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 143. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\common-config-map`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\common-config-map`

```text
Error: common-config-map/templates/infrastructure-configmap.yaml:11:36
  executing "common-config-map/templates/infrastructure-configmap.yaml" at <.Values.global.contentSecurity.type>:
    nil pointer evaluating interface {}.contentSecurity

Use --debug flag to render out invalid YAML
```

### 144. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\deploy-data-init`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\deploy-data-init`

```text
Error: deploy-data-init/templates/data-init-job.yaml:10:18
  executing "deploy-data-init/templates/data-init-job.yaml" at <.Values.global.customLabels>:
    nil pointer evaluating interface {}.customLabels

Use --debug flag to render out invalid YAML
```

### 145. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\deploy-pre-check`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\deploy-pre-check`

```text
Error: deploy-pre-check/templates/job.yaml:9:18
  executing "deploy-pre-check/templates/job.yaml" at <.Values.global.customLabels>:
    nil pointer evaluating interface {}.customLabels

Use --debug flag to render out invalid YAML
```

### 146. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\entity-extractor-server`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\entity-extractor-server`

```text
Error: entity-extractor-server/templates/deployment.yaml:7:18
  executing "entity-extractor-server/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 147. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\entity-extractor-server-sand-box`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\entity-extractor-server-sand-box`

```text
Error: entity-extractor-server-sand-box/templates/deployment.yaml:7:18
  executing "entity-extractor-server-sand-box/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 148. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\etcd`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\etcd`

```text
Error: etcd/templates/statefulset.yaml:9:18
  executing "etcd/templates/statefulset.yaml" at <.Values.global.customLabels>:
    nil pointer evaluating interface {}.customLabels

Use --debug flag to render out invalid YAML
```

### 149. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\lke-access-manager`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\lke-access-manager`

```text
Error: lke-access-manager/templates/deployment.yaml:10:18
  executing "lke-access-manager/templates/deployment.yaml" at <.Values.global.customLabels>:
    nil pointer evaluating interface {}.customLabels

Use --debug flag to render out invalid YAML
```

### 150. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\lke-aiconf-manager`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\lke-aiconf-manager`

```text
Error: lke-aiconf-manager/templates/deployment.yaml:6:18
  executing "lke-aiconf-manager/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 151. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\lke-bot-admin-config-server`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\lke-bot-admin-config-server`

```text
Error: lke-bot-admin-config-server/templates/deployment.yaml:7:18
  executing "lke-bot-admin-config-server/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 152. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\lke-ti-auth-web`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\lke-ti-auth-web`

```text
Error: lke-ti-auth-web/templates/deployment.yaml:8:18
  executing "lke-ti-auth-web/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 153. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\lke-ti-ingress-controller`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\lke-ti-ingress-controller`

```text
Error: lke-ti-ingress-controller/templates/deployment.yaml:10:18
  executing "lke-ti-ingress-controller/templates/deployment.yaml" at <.Values.global.customLabels>:
    nil pointer evaluating interface {}.customLabels

Use --debug flag to render out invalid YAML
```

### 154. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\lke-ti-perm-auth`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\lke-ti-perm-auth`

```text
Error: lke-ti-perm-auth/templates/secret.yaml:8:22
  executing "lke-ti-perm-auth/templates/secret.yaml" at <$.Values.global.rsa.publicKey>:
    nil pointer evaluating interface {}.rsa

Use --debug flag to render out invalid YAML
```

### 155. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\lke-ti-perm-auth-proxy`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\lke-ti-perm-auth-proxy`

```text
Error: lke-ti-perm-auth-proxy/templates/deployment.yaml:9:18
  executing "lke-ti-perm-auth-proxy/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 156. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\web-parser-server`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\web-parser-server`

```text
Error: web-parser-server/templates/deployment.yaml:7:18
  executing "web-parser-server/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 157. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-agent-config`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-agent-config`

```text
Error: adp-agent-config/templates/deployment.yaml:9:18
  executing "adp-agent-config/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 158. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-agent-exec`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-agent-exec`

```text
Error: adp-agent-exec/templates/deployment.yaml:9:18
  executing "adp-agent-exec/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 159. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-ai-gateway`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-ai-gateway`

```text
Error: adp-ai-gateway/templates/deployment.yaml:7:18
  executing "adp-ai-gateway/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 160. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-apex`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-apex`

```text
Error: adp-apex/templates/apex.yaml:8:18
  executing "adp-apex/templates/apex.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 161. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-app-channel-msg`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-app-channel-msg`

```text
Error: adp-app-channel-msg/templates/deployment.yaml:7:18
  executing "adp-app-channel-msg/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 162. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-app-channel-proxy-svr`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-app-channel-proxy-svr`

```text
Error: adp-app-channel-proxy-svr/templates/deployment.yaml:7:18
  executing "adp-app-channel-proxy-svr/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 163. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-app-channel-token`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-app-channel-token`

```text
Error: adp-app-channel-token/templates/deployment.yaml:7:18
  executing "adp-app-channel-token/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 164. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-app-config`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-app-config`

```text
Error: adp-app-config/templates/deployment.yaml:7:18
  executing "adp-app-config/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 165. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-app-evaluate`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-app-evaluate`

```text
Error: adp-app-evaluate/templates/deployment.yaml:7:18
  executing "adp-app-evaluate/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 166. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-app-memory`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-app-memory`

```text
Error: adp-app-memory/templates/deployment.yaml:9:18
  executing "adp-app-memory/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 167. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-app-shorturl`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-app-shorturl`

```text
Error: adp-app-shorturl/templates/deployment.yaml:7:18
  executing "adp-app-shorturl/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 168. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-chat-manage`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-chat-manage`

```text
Error: adp-chat-manage/templates/ingress.yaml:15:28
  executing "adp-chat-manage/templates/ingress.yaml" at <.Values.global.scheme>:
    nil pointer evaluating interface {}.scheme

Use --debug flag to render out invalid YAML
```

### 169. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-crawler`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-crawler`

```text
Error: adp-crawler/templates/deployment.yaml:9:18
  executing "adp-crawler/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 170. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-file-converter`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-file-converter`

```text
Error: adp-file-converter/templates/deployment-fileconverter.yaml:9:18
  executing "adp-file-converter/templates/deployment-fileconverter.yaml" at <.Values.global.customLabels>:
    nil pointer evaluating interface {}.customLabels

Use --debug flag to render out invalid YAML
```

### 171. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-kb-config`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-kb-config`

```text
Error: adp-kb-config/templates/deployment.yaml:9:18
  executing "adp-kb-config/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 172. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-kb-retrieval`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-kb-retrieval`

```text
Error: adp-kb-retrieval/templates/deployment.yaml:7:18
  executing "adp-kb-retrieval/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 173. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-kb-vdb-proxy`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-kb-vdb-proxy`

```text
Error: adp-kb-vdb-proxy/templates/deployment.yaml:7:18
  executing "adp-kb-vdb-proxy/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 174. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-markmap-service`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-markmap-service`

```text
Error: adp-markmap-service/templates/deployment.yaml:9:18
  executing "adp-markmap-service/templates/deployment.yaml" at <.Values.global.customLabels>:
    nil pointer evaluating interface {}.customLabels

Use --debug flag to render out invalid YAML
```

### 175. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-model-token-count`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-model-token-count`

```text
Error: adp-model-token-count/templates/deployment.yaml:9:18
  executing "adp-model-token-count/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 176. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-platform-authenticator`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-platform-authenticator`

```text
Error: adp-platform-authenticator/templates/deployment.yaml:7:18
  executing "adp-platform-authenticator/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 177. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-platform-charger`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-platform-charger`

```text
Error: adp-platform-charger/templates/deployment.yaml:7:18
  executing "adp-platform-charger/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 178. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-platform-content-moderation`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-platform-content-moderation`

```text
Error: adp-platform-content-moderation/templates/deployment.yaml:7:18
  executing "adp-platform-content-moderation/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 179. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-platform-manager`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-platform-manager`

```text
Error: adp-platform-manager/templates/deployment.yaml:7:18
  executing "adp-platform-manager/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 180. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-platform-metrology`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-platform-metrology`

```text
Error: adp-platform-metrology/templates/deployment.yaml:7:18
  executing "adp-platform-metrology/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 181. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-platform-op`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-platform-op`

```text
Error: adp-platform-op/templates/deployment.yaml:7:18
  executing "adp-platform-op/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 182. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-platform-permission-config`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-platform-permission-config`

```text
Error: adp-platform-permission-config/templates/deployment.yaml:7:18
  executing "adp-platform-permission-config/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 183. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-plugin-code-interpret`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-plugin-code-interpret`

```text
Error: adp-plugin-code-interpret/templates/deployment.yaml:9:18
  executing "adp-plugin-code-interpret/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 184. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-plugin-code-kernel`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-plugin-code-kernel`

```text
Error: adp-plugin-code-kernel/templates/deployment.yaml:9:18
  executing "adp-plugin-code-kernel/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 185. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-plugin-config`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-plugin-config`

```text
Error: adp-plugin-config/templates/deployment.yaml:9:18
  executing "adp-plugin-config/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 186. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-plugin-cos-util`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-plugin-cos-util`

```text
Error: adp-plugin-cos-util/templates/deployment.yaml:8:18
  executing "adp-plugin-cos-util/templates/deployment.yaml" at <.Values.global.customLabels>:
    nil pointer evaluating interface {}.customLabels

Use --debug flag to render out invalid YAML
```

### 187. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-plugin-exec`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-plugin-exec`

```text
Error: adp-plugin-exec/templates/deployment.yaml:9:18
  executing "adp-plugin-exec/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 188. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-plugin-pod-schedule`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-plugin-pod-schedule`

```text
Error: adp-plugin-pod-schedule/templates/deployment.yaml:7:18
  executing "adp-plugin-pod-schedule/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 189. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-pyright`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-pyright`

```text
Error: adp-pyright/templates/deployment-pyright.yaml:9:18
  executing "adp-pyright/templates/deployment-pyright.yaml" at <.Values.global.customLabels>:
    nil pointer evaluating interface {}.customLabels

Use --debug flag to render out invalid YAML
```

### 190. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-rag-exec`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-rag-exec`

```text
Error: adp-rag-exec/templates/deployment.yaml:7:18
  executing "adp-rag-exec/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 191. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-resource-gallery`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-resource-gallery`

```text
Error: adp-resource-gallery/templates/deployment.yaml:7:18
  executing "adp-resource-gallery/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 192. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-web-app`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-web-app`

```text
Error: adp-web-app/templates/deployment-webappchatbot.yaml:9:18
  executing "adp-web-app/templates/deployment-webappchatbot.yaml" at <.Values.global.customLabels>:
    nil pointer evaluating interface {}.customLabels

Use --debug flag to render out invalid YAML
```

### 193. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-widget-config`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-widget-config`

```text
Error: adp-widget-config/templates/deployment.yaml:7:18
  executing "adp-widget-config/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 194. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-workflow-config`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-workflow-config`

```text
Error: adp-workflow-config/templates/deployment.yaml:9:18
  executing "adp-workflow-config/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 195. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-workflow-exec`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\adp-workflow-exec`

```text
Error: adp-workflow-exec/templates/deployment.yaml:9:18
  executing "adp-workflow-exec/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 196. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\common-config-map`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\common-config-map`

```text
Error: common-config-map/templates/infrastructure-configmap.yaml:11:36
  executing "common-config-map/templates/infrastructure-configmap.yaml" at <.Values.global.contentSecurity.type>:
    nil pointer evaluating interface {}.contentSecurity

Use --debug flag to render out invalid YAML
```

### 197. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\deploy-data-init`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\deploy-data-init`

```text
Error: deploy-data-init/templates/data-init-job.yaml:10:18
  executing "deploy-data-init/templates/data-init-job.yaml" at <.Values.global.customLabels>:
    nil pointer evaluating interface {}.customLabels

Use --debug flag to render out invalid YAML
```

### 198. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\deploy-pre-check`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\deploy-pre-check`

```text
Error: deploy-pre-check/templates/job.yaml:9:18
  executing "deploy-pre-check/templates/job.yaml" at <.Values.global.customLabels>:
    nil pointer evaluating interface {}.customLabels

Use --debug flag to render out invalid YAML
```

### 199. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\entity-extractor-server`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\entity-extractor-server`

```text
Error: entity-extractor-server/templates/deployment.yaml:7:18
  executing "entity-extractor-server/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 200. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\entity-extractor-server-sand-box`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\entity-extractor-server-sand-box`

```text
Error: entity-extractor-server-sand-box/templates/deployment.yaml:7:18
  executing "entity-extractor-server-sand-box/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 201. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\etcd`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\etcd`

```text
Error: etcd/templates/statefulset.yaml:9:18
  executing "etcd/templates/statefulset.yaml" at <.Values.global.customLabels>:
    nil pointer evaluating interface {}.customLabels

Use --debug flag to render out invalid YAML
```

### 202. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\lke-access-manager`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\lke-access-manager`

```text
Error: lke-access-manager/templates/deployment.yaml:10:18
  executing "lke-access-manager/templates/deployment.yaml" at <.Values.global.customLabels>:
    nil pointer evaluating interface {}.customLabels

Use --debug flag to render out invalid YAML
```

### 203. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\lke-aiconf-manager`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\lke-aiconf-manager`

```text
Error: lke-aiconf-manager/templates/deployment.yaml:6:18
  executing "lke-aiconf-manager/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 204. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\lke-bot-admin-config-server`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\lke-bot-admin-config-server`

```text
Error: lke-bot-admin-config-server/templates/deployment.yaml:7:18
  executing "lke-bot-admin-config-server/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 205. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\lke-ti-auth-web`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\lke-ti-auth-web`

```text
Error: lke-ti-auth-web/templates/deployment.yaml:8:18
  executing "lke-ti-auth-web/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 206. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\lke-ti-ingress-controller`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\lke-ti-ingress-controller`

```text
Error: lke-ti-ingress-controller/templates/deployment.yaml:10:18
  executing "lke-ti-ingress-controller/templates/deployment.yaml" at <.Values.global.customLabels>:
    nil pointer evaluating interface {}.customLabels

Use --debug flag to render out invalid YAML
```

### 207. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\lke-ti-perm-auth`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\lke-ti-perm-auth`

```text
Error: lke-ti-perm-auth/templates/secret.yaml:8:22
  executing "lke-ti-perm-auth/templates/secret.yaml" at <$.Values.global.rsa.publicKey>:
    nil pointer evaluating interface {}.rsa

Use --debug flag to render out invalid YAML
```

### 208. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\lke-ti-perm-auth-proxy`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\lke-ti-perm-auth-proxy`

```text
Error: lke-ti-perm-auth-proxy/templates/deployment.yaml:9:18
  executing "lke-ti-perm-auth-proxy/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 209. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\web-parser-server`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\web-parser-server`

```text
Error: web-parser-server/templates/deployment.yaml:7:18
  executing "web-parser-server/templates/deployment.yaml" at <.Values.global.customAnnotations>:
    nil pointer evaluating interface {}.customAnnotations

Use --debug flag to render out invalid YAML
```

### 210. `ODIM-Project/ODIM`

- Chart: `D:\helm_clones_github\ODIM-Project__ODIM\odim-controller\helmcharts\reloader`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\ODIM-Project__ODIM\odim-controller\helmcharts\reloader`

```text
Error: reloader/templates/deployment.yaml:56:26
  executing "reloader/templates/deployment.yaml" at <.Values.odimra.imageRegistryAddress>:
    nil pointer evaluating interface {}.imageRegistryAddress

Use --debug flag to render out invalid YAML
```

### 211. `YAKEcloud/yake`

- Chart: `D:\helm_clones_github\YAKEcloud__yake\configuration\configuration`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\YAKEcloud__yake\configuration\configuration`

```text
Error: configuration/templates/kube-apiserver.yaml:13:19
  executing "configuration/templates/kube-apiserver.yaml" at <include "api.domain" .>:
    error calling include:
configuration/templates/_domains.tpl:28:15
  executing "api.domain" at <.Values.domains.api>:
    nil pointer evaluating interface {}.api

Use --debug flag to render out invalid YAML
```

### 212. `YAKEcloud/yake`

- Chart: `D:\helm_clones_github\YAKEcloud__yake\gardener\garden-content`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\YAKEcloud__yake\gardener\garden-content`

```text
Error: garden-content/templates/secret-openvpn-diffie-hellman.yaml:20:4
  executing "garden-content/templates/secret-openvpn-diffie-hellman.yaml" at <include "gardener.secret-openvpn-diffie-hellman" .>:
    error calling include:
garden-content/templates/secret-openvpn-diffie-hellman.yaml:3:14
  executing "gardener.secret-openvpn-diffie-hellman" at <.Values.global.openVPNDiffieHellmanKey>:
    nil pointer evaluating interface {}.openVPNDiffieHellmanKey

Use --debug flag to render out invalid YAML
```

### 213. `Loongson-Cloud-Community/dockerfiles`

- Chart: `D:\helm_clones_github\Loongson-Cloud-Community__dockerfiles\kubesphere\ks-installer\v3.2.1\roles\common\files\openldap-ha`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Loongson-Cloud-Community__dockerfiles\kubesphere\ks-installer\v3.2.1\roles\common\files\openldap-ha`

```text
Error: openldap-ha/templates/statefulset.yaml:144:18
  executing "openldap-ha/templates/statefulset.yaml" at <.Values.persistence.storageClass>:
    nil pointer evaluating interface {}.storageClass

Use --debug flag to render out invalid YAML
```

### 214. `Loongson-Cloud-Community/dockerfiles`

- Chart: `D:\helm_clones_github\Loongson-Cloud-Community__dockerfiles\kubesphere\ks-installer\v3.2.1\roles\ks-auditing\files\kube-auditing`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Loongson-Cloud-Community__dockerfiles\kubesphere\ks-installer\v3.2.1\roles\ks-auditing\files\kube-auditing`

```text
Error: kube-auditing/templates/webhook.yaml:6:16
  executing "kube-auditing/templates/webhook.yaml" at <.Values.webhook.replicas>:
    nil pointer evaluating interface {}.replicas

Use --debug flag to render out invalid YAML
```

### 215. `Loongson-Cloud-Community/dockerfiles`

- Chart: `D:\helm_clones_github\Loongson-Cloud-Community__dockerfiles\kubesphere\ks-installer\v3.2.1\roles\ks-monitor\files\notification-manager`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Loongson-Cloud-Community__dockerfiles\kubesphere\ks-installer\v3.2.1\roles\ks-monitor\files\notification-manager`

```text
Error: notification-manager/templates/operator.yaml:23:25
  executing "notification-manager/templates/operator.yaml" at <.Values.operator.containers.proxy.image.repo>:
    nil pointer evaluating interface {}.containers

Use --debug flag to render out invalid YAML
```

### 216. `FIWARE-Ops/marinera`

- Chart: `D:\helm_clones_github\FIWARE-Ops__marinera\applications\grafana-metrics\chart`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\FIWARE-Ops__marinera\applications\grafana-metrics\chart`

```text
Error: grafana-metrics/templates/tests/test-config.yaml:23:29
  executing "grafana-metrics/templates/tests/test-config.yaml" at <.Values.admin.username>:
    nil pointer evaluating interface {}.username

Use --debug flag to render out invalid YAML
```

### 217. `oneconcern/datamon`

- Chart: `D:\helm_clones_github\oneconcern__datamon\k8s\purge\gsutil`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oneconcern__datamon\k8s\purge\gsutil`

```text
Error: datamon-gsutil/templates/job.yaml:28:26
  executing "datamon-gsutil/templates/job.yaml" at <.Values.image.repository>:
    nil pointer evaluating interface {}.repository

Use --debug flag to render out invalid YAML
```

### 218. `cnrancher/pandaria-catalog`

- Chart: `D:\helm_clones_github\cnrancher__pandaria-catalog\charts\mcs-ext-chart\0.0.1\charts\mcs-addon-crd`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\cnrancher__pandaria-catalog\charts\mcs-ext-chart\0.0.1\charts\mcs-addon-crd`

```text
Error: mcs-addon-crd/templates/serviceimport.yaml:1:17
  executing "mcs-addon-crd/templates/serviceimport.yaml" at <.Values.global.installationType>:
    nil pointer evaluating interface {}.installationType

Use --debug flag to render out invalid YAML
```

### 219. `cnrancher/pandaria-catalog`

- Chart: `D:\helm_clones_github\cnrancher__pandaria-catalog\charts\mcs-ext-chart\0.0.1\charts\submariner-operator`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\cnrancher__pandaria-catalog\charts\mcs-ext-chart\0.0.1\charts\submariner-operator`

```text
Error: submariner-operator/templates/submariner.yaml:1:17
  executing "submariner-operator/templates/submariner.yaml" at <.Values.global.installationType>:
    nil pointer evaluating interface {}.installationType

Use --debug flag to render out invalid YAML
```

### 220. `mrybas/k8s-bootstrap`

- Chart: `D:\helm_clones_github\mrybas__k8s-bootstrap\backend\definitions\charts\kubevirt-operator`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\mrybas__k8s-bootstrap\backend\definitions\charts\kubevirt-operator`

```text
Error: kubevirt-operator/templates/servicemonitor.yaml:1:14
  executing "kubevirt-operator/templates/servicemonitor.yaml" at <.Values.serviceMonitor.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 221. `mrybas/k8s-bootstrap`

- Chart: `D:\helm_clones_github\mrybas__k8s-bootstrap\backend\definitions\charts\piraeus-operator`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\mrybas__k8s-bootstrap\backend\definitions\charts\piraeus-operator`

```text
Error: piraeus-operator/templates/servicemonitor.yaml:1:14
  executing "piraeus-operator/templates/servicemonitor.yaml" at <.Values.serviceMonitor.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 222. `DOME-Marketplace/dome-gitops`

- Chart: `D:\helm_clones_github\DOME-Marketplace__dome-gitops\ionos_common\cert-manager`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\DOME-Marketplace__dome-gitops\ionos_common\cert-manager`

```text
Error: cert-manager/templates/issuer.yaml:4:30
  executing "cert-manager/templates/issuer.yaml" at <.Values.dome.env>:
    nil pointer evaluating interface {}.env

Use --debug flag to render out invalid YAML
```

### 223. `DOME-Marketplace/dome-gitops`

- Chart: `D:\helm_clones_github\DOME-Marketplace__dome-gitops\ionos_common\external-dns-ionos`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\DOME-Marketplace__dome-gitops\ionos_common\external-dns-ionos`

```text
Error: external-dns-ionos/templates/ionos-token-sealed-secret.yaml:9:21
  executing "external-dns-ionos/templates/ionos-token-sealed-secret.yaml" at <.Values.ionos.token>:
    nil pointer evaluating interface {}.token

Use --debug flag to render out invalid YAML
```

### 224. `btr1975/automation-framework`

- Chart: `D:\helm_clones_github\btr1975__automation-framework\helm\teamcity\charts\teamcity-agent`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\btr1975__automation-framework\helm\teamcity\charts\teamcity-agent`

```text
Error: teamcity-agent/templates/service.yaml:8:21
  executing "teamcity-agent/templates/service.yaml" at <.Values.global.namespace>:
    nil pointer evaluating interface {}.namespace

Use --debug flag to render out invalid YAML
```

### 225. `btr1975/automation-framework`

- Chart: `D:\helm_clones_github\btr1975__automation-framework\helm\teamcity\charts\teamcity-server`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\btr1975__automation-framework\helm\teamcity\charts\teamcity-server`

```text
Error: teamcity-server/templates/service.yaml:8:21
  executing "teamcity-server/templates/service.yaml" at <.Values.global.namespace>:
    nil pointer evaluating interface {}.namespace

Use --debug flag to render out invalid YAML
```

### 226. `cmgoffena13/etl-watcher`

- Chart: `D:\helm_clones_github\cmgoffena13__etl-watcher\watcher`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\cmgoffena13__etl-watcher\watcher`

```text
Error: watcher/templates/serviceaccount.yaml:1:14
  executing "watcher/templates/serviceaccount.yaml" at <.Values.gcp.secretManager.enabled>:
    nil pointer evaluating interface {}.secretManager

Use --debug flag to render out invalid YAML
```

### 227. `cmgoffena13/etl-watcher`

- Chart: `D:\helm_clones_github\cmgoffena13__etl-watcher\watcher`
- Source: `template`
- Values files: `D:\helm_clones_github\cmgoffena13__etl-watcher\watcher\values-prod.yaml`
- Command: `helm template test D:\helm_clones_github\cmgoffena13__etl-watcher\watcher -f D:\helm_clones_github\cmgoffena13__etl-watcher\watcher\values-prod.yaml`

```text
Error: watcher/templates/serviceaccount.yaml:1:14
  executing "watcher/templates/serviceaccount.yaml" at <.Values.gcp.secretManager.enabled>:
    nil pointer evaluating interface {}.secretManager

Use --debug flag to render out invalid YAML
```

### 228. `thoughtworks/byor-voting-infrastructure`

- Chart: `D:\helm_clones_github\thoughtworks__byor-voting-infrastructure\src\byor-voting-cert`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\thoughtworks__byor-voting-infrastructure\src\byor-voting-cert`

```text
Error: byor-voting-cert/templates/cluster-issuer.yaml:8:22
  executing "byor-voting-cert/templates/cluster-issuer.yaml" at <.Values.letsencrypt.Url>:
    nil pointer evaluating interface {}.Url

Use --debug flag to render out invalid YAML
```

### 229. `SAP/component-operator-runtime`

- Chart: `D:\helm_clones_github\SAP__component-operator-runtime\internal\helm\testdata\main`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\SAP__component-operator-runtime\internal\helm\testdata\main`

```text
Error: main/templates/configmap.yaml:20:24
  executing "main/templates/configmap.yaml" at <.Values.global.data>:
    nil pointer evaluating interface {}.data

Use --debug flag to render out invalid YAML
```

### 230. `SAP/component-operator-runtime`

- Chart: `D:\helm_clones_github\SAP__component-operator-runtime\internal\helm\testdata\main\charts\sub11`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\SAP__component-operator-runtime\internal\helm\testdata\main\charts\sub11`

```text
Error: sub11/templates/configmap.yaml:16:24
  executing "sub11/templates/configmap.yaml" at <.Values.global.data>:
    nil pointer evaluating interface {}.data

Use --debug flag to render out invalid YAML
```

### 231. `SAP/component-operator-runtime`

- Chart: `D:\helm_clones_github\SAP__component-operator-runtime\internal\helm\testdata\main\charts\sub11\charts\sub21`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\SAP__component-operator-runtime\internal\helm\testdata\main\charts\sub11\charts\sub21`

```text
Error: sub21/templates/configmap.yaml:11:24
  executing "sub21/templates/configmap.yaml" at <.Values.global.data>:
    nil pointer evaluating interface {}.data

Use --debug flag to render out invalid YAML
```

### 232. `SAP/component-operator-runtime`

- Chart: `D:\helm_clones_github\SAP__component-operator-runtime\internal\helm\testdata\main\charts\sub12`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\SAP__component-operator-runtime\internal\helm\testdata\main\charts\sub12`

```text
Error: sub12/templates/configmap.yaml:11:24
  executing "sub12/templates/configmap.yaml" at <.Values.global.data>:
    nil pointer evaluating interface {}.data

Use --debug flag to render out invalid YAML
```

### 233. `batleforc/weebo-si`

- Chart: `D:\helm_clones_github\batleforc__weebo-si\all-in-one.argo\helm\cluster-tpl-object`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\batleforc__weebo-si\all-in-one.argo\helm\cluster-tpl-object`

```text
Error: cluster-tpl-object/templates/monitoring/coroot/secret-value.yaml:1:17
  executing "cluster-tpl-object/templates/monitoring/coroot/secret-value.yaml" at <.Values.monitoring.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 234. `celo-org/charts`

- Chart: `D:\helm_clones_github\celo-org__charts\charts\clean-pvcs`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\celo-org__charts\charts\clean-pvcs`

```text
Error: clean-pvcs/templates/cronjob.yaml:53:59
  executing "clean-pvcs/templates/cronjob.yaml" at <.Release.Namespace>:
    nil pointer evaluating interface {}.Namespace

Use --debug flag to render out invalid YAML
```

### 235. `dungdm93/shipyard`

- Chart: `D:\helm_clones_github\dungdm93__shipyard\helm\airflow`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\dungdm93__shipyard\helm\airflow`

```text
Error: airflow/charts/redis/templates/master/service.yaml:12:14
  executing "airflow/charts/redis/templates/master/service.yaml" at <include "common.labels.standard" (dict "customLabels" .Values.commonLabels "context" $)>:
    error calling include:
airflow/charts/postgresql/charts/common/templates/_labels.tpl:6:27
  executing "common.labels.standard" at <include "common.names.name" .>:
    error calling include:
airflow/charts/postgresql/charts/common/templates/_names.tpl:6:18
  executing "common.names.name" at <.Chart.Name>:
    nil pointer evaluating interface {}.Name

Use --debug flag to render out invalid YAML
```

### 236. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\acestepv3\acestepv3server`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\acestepv3\acestepv3server`

```text
Error: acestepv3server/templates/deployment.yaml:78:29
  executing "acestepv3server/templates/deployment.yaml" at <.Values.userspace.userData>:
    nil pointer evaluating interface {}.userData

Use --debug flag to render out invalid YAML
```

### 237. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\affine`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\affine`

```text
Error: affine/templates/redis/deployment.yaml:67:29
  executing "affine/templates/redis/deployment.yaml" at <.Values.userspace.appData>:
    nil pointer evaluating interface {}.appData

Use --debug flag to render out invalid YAML
```

### 238. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\agentzero`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\agentzero`

```text
Error: agentzero/templates/agentzero.yaml:71:29
  executing "agentzero/templates/agentzero.yaml" at <.Values.userspace.appData>:
    nil pointer evaluating interface {}.appData

Use --debug flag to render out invalid YAML
```

### 239. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\alist`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\alist`

```text
Error: alist/templates/alist.yaml:70:29
  executing "alist/templates/alist.yaml" at <.Values.userspace.appData>:
    nil pointer evaluating interface {}.appData

Use --debug flag to render out invalid YAML
```

### 240. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\answer`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\answer`

```text
Error: answer/templates/answer.yaml:1:43
  executing "answer/templates/answer.yaml" at <.Values.domain.answer>:
    nil pointer evaluating interface {}.answer

Use --debug flag to render out invalid YAML
```

### 241. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\anythingllm`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\anythingllm`

```text
Error: anythingllm/templates/configmap.yaml:1:46
  executing "anythingllm/templates/configmap.yaml" at <.Values.domain.anythingllm>:
    nil pointer evaluating interface {}.anythingllm

Use --debug flag to render out invalid YAML
```

### 242. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\appsmith`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\appsmith`

```text
Error: appsmith/templates/appsmith.yaml:73:29
  executing "appsmith/templates/appsmith.yaml" at <.Values.userspace.appData>:
    nil pointer evaluating interface {}.appData

Use --debug flag to render out invalid YAML
```

### 243. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\astrbot`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\astrbot`

```text
Error: astrbot/templates/deployment.yaml:55:27
  executing "astrbot/templates/deployment.yaml" at <.Values.userspace.appData>:
    nil pointer evaluating interface {}.appData

Use --debug flag to render out invalid YAML
```

### 244. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\audiobookshelf`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\audiobookshelf`

```text
Error: audiobookshelf/templates/deployment.yaml:163:29
  executing "audiobookshelf/templates/deployment.yaml" at <.Values.userspace.userData>:
    nil pointer evaluating interface {}.userData

Use --debug flag to render out invalid YAML
```

### 245. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\awesomedigitalhuman`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\awesomedigitalhuman`

```text
Error: awesomedigitalhuman/templates/deployment.yaml:1:41
  executing "awesomedigitalhuman/templates/deployment.yaml" at <.Values.domain.awesomedigitalhuman>:
    nil pointer evaluating interface {}.awesomedigitalhuman

Use --debug flag to render out invalid YAML
```

### 246. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\bazarr`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\bazarr`

```text
Error: bazarr/templates/deployment.yaml:74:26
  executing "bazarr/templates/deployment.yaml" at <.Values.userspace.appData>:
    nil pointer evaluating interface {}.appData

Use --debug flag to render out invalid YAML
```

### 247. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\bifrost`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\bifrost`

```text
Error: bifrost/templates/configmap.yaml:18:29
  executing "bifrost/templates/configmap.yaml" at <.Values.postgres.host>:
    nil pointer evaluating interface {}.host

Use --debug flag to render out invalid YAML
```

### 248. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\bisheng`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\bisheng`

```text
Error: bisheng/templates/redis.yaml:66:29
  executing "bisheng/templates/redis.yaml" at <.Values.userspace.appData>:
    nil pointer evaluating interface {}.appData

Use --debug flag to render out invalid YAML
```

### 249. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\blinko`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\blinko`

```text
Error: blinko/templates/deployment.yaml:33:39
  executing "blinko/templates/deployment.yaml" at <.Values.postgres.username>:
    nil pointer evaluating interface {}.username

Use --debug flag to render out invalid YAML
```

### 250. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\bytebase`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\bytebase`

```text
Error: bytebase/templates/bytebase.yaml:1:45
  executing "bytebase/templates/bytebase.yaml" at <.Values.domain.bytebase>:
    nil pointer evaluating interface {}.bytebase

Use --debug flag to render out invalid YAML
```

### 251. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\calendar`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\calendar`

```text
Error: calendar/templates/calendar_deploy.yaml:91:27
  executing "calendar/templates/calendar_deploy.yaml" at <.Values.postgres.databases.bloben_api>:
    nil pointer evaluating interface {}.bloben_api

Use --debug flag to render out invalid YAML
```

### 252. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\calibre`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\calibre`

```text
Error: calibre/templates/deployment.yaml:61:26
  executing "calibre/templates/deployment.yaml" at <.Values.userspace.appCache>:
    nil pointer evaluating interface {}.appCache

Use --debug flag to render out invalid YAML
```

### 253. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\calibreweb`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\calibreweb`

```text
Error: calibreweb/templates/deployment.yaml:49:26
  executing "calibreweb/templates/deployment.yaml" at <.Values.userspace.userData>:
    nil pointer evaluating interface {}.userData

Use --debug flag to render out invalid YAML
```

### 254. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\chinesesubfinder`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\chinesesubfinder`

```text
Error: chinesesubfinder/templates/chinesesubfinder.yaml:84:28
  executing "chinesesubfinder/templates/chinesesubfinder.yaml" at <.Values.userspace.appData>:
    nil pointer evaluating interface {}.appData

Use --debug flag to render out invalid YAML
```

### 255. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\chromium`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\chromium`

```text
Error: chromium/templates/chromium.yaml:139:28
  executing "chromium/templates/chromium.yaml" at <.Values.userspace.appCache>:
    nil pointer evaluating interface {}.appCache

Use --debug flag to render out invalid YAML
```

### 256. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\claudecode`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\claudecode`

```text
Error: claudecode/templates/deployment.yaml:157:28
  executing "claudecode/templates/deployment.yaml" at <.Values.olaresEnv.ANTHROPIC_AUTH_TOKEN>:
    nil pointer evaluating interface {}.ANTHROPIC_AUTH_TOKEN

Use --debug flag to render out invalid YAML
```

### 257. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\clawdbot`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\clawdbot`

```text
Error: clawdbot/templates/deployment.yaml:368:29
  executing "clawdbot/templates/deployment.yaml" at <.Values.userspace.appData>:
    nil pointer evaluating interface {}.appData

Use --debug flag to render out invalid YAML
```

### 258. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\cloudreve`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\cloudreve`

```text
Error: cloudreve/templates/redis/deployment.yaml:68:29
  executing "cloudreve/templates/redis/deployment.yaml" at <.Values.userspace.appData>:
    nil pointer evaluating interface {}.appData

Use --debug flag to render out invalid YAML
```

### 259. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\coder`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\coder`

```text
Error: coder/templates/coder.yaml:1:40
  executing "coder/templates/coder.yaml" at <.Values.domain.coder>:
    nil pointer evaluating interface {}.coder

Use --debug flag to render out invalid YAML
```

### 260. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\codeserver`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\codeserver`

```text
Error: codeserver/templates/codeserver.yaml:70:28
  executing "codeserver/templates/codeserver.yaml" at <.Values.userspace.userData>:
    nil pointer evaluating interface {}.userData

Use --debug flag to render out invalid YAML
```

### 261. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\codex`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\codex`

```text
Error: codex/templates/deployment.yaml:174:32
  executing "codex/templates/deployment.yaml" at <.Values.olaresEnv.OPENAI_API_KEY>:
    nil pointer evaluating interface {}.OPENAI_API_KEY

Use --debug flag to render out invalid YAML
```

### 262. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\comfyuishare`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\comfyuishare`

```text
Error: comfyuishare/templates/clientproxy.yaml:22:46
  executing "comfyuishare/templates/clientproxy.yaml" at <.Values.bfl.username>:
    nil pointer evaluating interface {}.username

Use --debug flag to render out invalid YAML
```

### 263. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\comfyuisharev2\comfyuisharev2`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\comfyuisharev2\comfyuisharev2`

```text
Error: comfyuisharev2/templates/clientproxy.yaml:22:46
  executing "comfyuisharev2/templates/clientproxy.yaml" at <.Values.bfl.username>:
    nil pointer evaluating interface {}.username

Use --debug flag to render out invalid YAML
```

### 264. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\context7`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\context7`

```text
Error: context7/templates/context7.yaml:56:26
  executing "context7/templates/context7.yaml" at <.Values.olaresEnv.CONTEXT7_API_KEY>:
    nil pointer evaluating interface {}.CONTEXT7_API_KEY

Use --debug flag to render out invalid YAML
```

### 265. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\deepseekocrwebuiv2\deepseekocrwebuiv2server`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\deepseekocrwebuiv2\deepseekocrwebuiv2server`

```text
Error: deepseekocrwebuiv2server/templates/download.yaml:83:29
  executing "deepseekocrwebuiv2server/templates/download.yaml" at <.Values.userspace.userData>:
    nil pointer evaluating interface {}.userData

Use --debug flag to render out invalid YAML
```

### 266. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\deerflow`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\deerflow`

```text
Error: deerflow/templates/frontend.yaml:1:43
  executing "deerflow/templates/frontend.yaml" at <.Values.domain.deerflow>:
    nil pointer evaluating interface {}.deerflow

Use --debug flag to render out invalid YAML
```

### 267. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\deerflowv2`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\deerflowv2`

```text
Error: deerflowv2/templates/deployment.yaml:326:45
  executing "deerflowv2/templates/deployment.yaml" at <.Values.postgres.username>:
    nil pointer evaluating interface {}.username

Use --debug flag to render out invalid YAML
```

### 268. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\deluge`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\deluge`

```text
Error: deluge/templates/deployment.yaml:63:26
  executing "deluge/templates/deployment.yaml" at <.Values.userspace.appData>:
    nil pointer evaluating interface {}.appData

Use --debug flag to render out invalid YAML
```

### 269. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\didgate`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\didgate`

```text
Error: didgate/templates/pg-client.yaml:31:28
  executing "didgate/templates/pg-client.yaml" at <.Values.postgres.host>:
    nil pointer evaluating interface {}.host

Use --debug flag to render out invalid YAML
```

### 270. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\dify`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\dify`

```text
Error: dify/templates/client/client.yaml:22:46
  executing "dify/templates/client/client.yaml" at <.Values.bfl.username>:
    nil pointer evaluating interface {}.username

Use --debug flag to render out invalid YAML
```

### 271. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\difyv2\difyv2server`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\difyv2\difyv2server`

```text
Error: difyv2server/templates/api/configmap.yaml:1:39
  executing "difyv2server/templates/api/configmap.yaml" at <.Values.domain.client>:
    nil pointer evaluating interface {}.client

Use --debug flag to render out invalid YAML
```

### 272. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\directus`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\directus`

```text
Error: directus/templates/directus.yaml:49:32
  executing "directus/templates/directus.yaml" at <.Values.olaresEnv.SECRET>:
    nil pointer evaluating interface {}.SECRET

Use --debug flag to render out invalid YAML
```

### 273. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\dman`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\dman`

```text
Error: dman/templates/signalling.yaml:117:28
  executing "dman/templates/signalling.yaml" at <.Values.userspace.appData>:
    nil pointer evaluating interface {}.appData

Use --debug flag to render out invalid YAML
```

### 274. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\docmost`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\docmost`

```text
Error: docmost/templates/docmost.yaml:33:60
  executing "docmost/templates/docmost.yaml" at <.Values.domain.docmost>:
    nil pointer evaluating interface {}.docmost

Use --debug flag to render out invalid YAML
```

### 275. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\documenso`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\documenso`

```text
Error: documenso/templates/documenso.yaml:49:32
  executing "documenso/templates/documenso.yaml" at <.Values.olaresEnv.NEXT_PRIVATE_SMTP_FROM_ADDRESS>:
    nil pointer evaluating interface {}.NEXT_PRIVATE_SMTP_FROM_ADDRESS

Use --debug flag to render out invalid YAML
```

### 276. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\duplicati`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\duplicati`

```text
Error: duplicati/templates/deployment.yaml:23:27
  executing "duplicati/templates/deployment.yaml" at <.Values.userspace.userData>:
    nil pointer evaluating interface {}.userData

Use --debug flag to render out invalid YAML
```

### 277. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\facefusion`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\facefusion`

```text
Error: facefusion/templates/facefusion.yaml:1:45
  executing "facefusion/templates/facefusion.yaml" at <.Values.domain.facefusion>:
    nil pointer evaluating interface {}.facefusion

Use --debug flag to render out invalid YAML
```

### 278. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\farcasterhubble`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\farcasterhubble`

```text
Error: farcasterhubble/templates/deployment.yaml:1:52
  executing "farcasterhubble/templates/deployment.yaml" at <.Values.domain.farcasterhubble>:
    nil pointer evaluating interface {}.farcasterhubble

Use --debug flag to render out invalid YAML
```

### 279. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\fastgpt`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\fastgpt`

```text
Error: fastgpt/templates/baseconfig.yaml:1:42
  executing "fastgpt/templates/baseconfig.yaml" at <.Values.domain.fastgpt>:
    nil pointer evaluating interface {}.fastgpt

Use --debug flag to render out invalid YAML
```

### 280. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\filebrowser`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\filebrowser`

```text
Error: filebrowser/templates/deployment.yaml:59:28
  executing "filebrowser/templates/deployment.yaml" at <.Values.os.appSecret>:
    nil pointer evaluating interface {}.appSecret

Use --debug flag to render out invalid YAML
```

### 281. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\firecrawl`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\firecrawl`

```text
Error: firecrawl/templates/redis/deployment.yaml:69:28
  executing "firecrawl/templates/redis/deployment.yaml" at <.Values.userspace.appCache>:
    nil pointer evaluating interface {}.appCache

Use --debug flag to render out invalid YAML
```

### 282. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\fireflyiii`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\fireflyiii`

```text
Error: fireflyiii/templates/fireflyiii.yaml:1:47
  executing "fireflyiii/templates/fireflyiii.yaml" at <.Values.domain.fireflyiii>:
    nil pointer evaluating interface {}.fireflyiii

Use --debug flag to render out invalid YAML
```

### 283. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\firefox`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\firefox`

```text
Error: firefox/templates/firefox.yaml:102:28
  executing "firefox/templates/firefox.yaml" at <.Values.userspace.appCache>:
    nil pointer evaluating interface {}.appCache

Use --debug flag to render out invalid YAML
```

### 284. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\flowise`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\flowise`

```text
Error: flowise/templates/configmap.yaml:24:28
  executing "flowise/templates/configmap.yaml" at <.Values.postgres.databases.flowise>:
    nil pointer evaluating interface {}.flowise

Use --debug flag to render out invalid YAML
```

### 285. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\focalboard`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\focalboard`

```text
Error: focalboard/templates/deployment.yaml:41:50
  executing "focalboard/templates/deployment.yaml" at <.Values.postgres.username>:
    nil pointer evaluating interface {}.username

Use --debug flag to render out invalid YAML
```

### 286. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\formbricks`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\formbricks`

```text
Error: formbricks/templates/formbricks.yaml:1:47
  executing "formbricks/templates/formbricks.yaml" at <.Values.domain.formbricks>:
    nil pointer evaluating interface {}.formbricks

Use --debug flag to render out invalid YAML
```

### 287. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\freshrss`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\freshrss`

```text
Error: freshrss/templates/freshrss.yaml:1:45
  executing "freshrss/templates/freshrss.yaml" at <.Values.domain.freshrss>:
    nil pointer evaluating interface {}.freshrss

Use --debug flag to render out invalid YAML
```

### 288. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\geth`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\geth`

```text
Error: geth/templates/website.yaml:26:57
  executing "geth/templates/website.yaml" at <.Values.bfl.username>:
    nil pointer evaluating interface {}.username

Use --debug flag to render out invalid YAML
```

### 289. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\ghost`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\ghost`

```text
Error: ghost/templates/deployment.yaml:1:42
  executing "ghost/templates/deployment.yaml" at <.Values.domain.ghost>:
    nil pointer evaluating interface {}.ghost

Use --debug flag to render out invalid YAML
```

### 290. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\gitea`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\gitea`

```text
Error: gitea/templates/deployment.yaml:39:32
  executing "gitea/templates/deployment.yaml" at <.Values.postgres.databases.gitea>:
    nil pointer evaluating interface {}.gitea

Use --debug flag to render out invalid YAML
```

### 291. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\gitlab`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\gitlab`

```text
Error: gitlab/templates/gitlab.yaml:7756:18
  executing "gitlab/templates/gitlab.yaml" at <.Values.olaresEnv.GITLAB_INITIAL_ROOT_PASSWORD>:
    nil pointer evaluating interface {}.GITLAB_INITIAL_ROOT_PASSWORD

Use --debug flag to render out invalid YAML
```

### 292. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\gitlab`
- Source: `template`
- Values files: `D:\helm_clones_github\Above-Os__terminus-apps\gitlab\values.helm.gitlab.yaml`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\gitlab -f D:\helm_clones_github\Above-Os__terminus-apps\gitlab\values.helm.gitlab.yaml`

```text
Error: gitlab/templates/gitlab.yaml:7756:18
  executing "gitlab/templates/gitlab.yaml" at <.Values.olaresEnv.GITLAB_INITIAL_ROOT_PASSWORD>:
    nil pointer evaluating interface {}.GITLAB_INITIAL_ROOT_PASSWORD

Use --debug flag to render out invalid YAML
```

### 293. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\gitlabpure`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\gitlabpure`

```text
Error: gitlabpure/templates/gitlab.yaml:3:43
  executing "gitlabpure/templates/gitlab.yaml" at <.Values.domain.gitlab>:
    nil pointer evaluating interface {}.gitlab

Use --debug flag to render out invalid YAML
```

### 294. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\grafana`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\grafana`

```text
Error: grafana/templates/granafa.yaml:51:31
  executing "grafana/templates/granafa.yaml" at <.Values.olaresEnv.GF_USERNAME>:
    nil pointer evaluating interface {}.GF_USERNAME

Use --debug flag to render out invalid YAML
```

### 295. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\halo`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\halo`

```text
Error: halo/templates/deployment.yaml:1:39
  executing "halo/templates/deployment.yaml" at <.Values.domain.halo>:
    nil pointer evaluating interface {}.halo

Use --debug flag to render out invalid YAML
```

### 296. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\hasura`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\hasura`

```text
Error: hasura/templates/hasura.yaml:140:32
  executing "hasura/templates/hasura.yaml" at <.Values.postgres.host>:
    nil pointer evaluating interface {}.host

Use --debug flag to render out invalid YAML
```

### 297. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\hermesagent`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\hermesagent`

```text
Error: hermesagent/templates/deployment.yaml:424:39
  executing "hermesagent/templates/deployment.yaml" at <.Values.olaresEnv.ALLOW_HOME_DIR_ACCESS>:
    nil pointer evaluating interface {}.ALLOW_HOME_DIR_ACCESS

Use --debug flag to render out invalid YAML
```

### 298. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\homeassistant`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\homeassistant`

```text
Error: homeassistant/templates/deployment.yaml:198:92
  executing "homeassistant/templates/deployment.yaml" at <.Values.userspace.appData>:
    nil pointer evaluating interface {}.appData

Use --debug flag to render out invalid YAML
```

### 299. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\homebox`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\homebox`

```text
Error: homebox/templates/homebox.yaml:62:32
  executing "homebox/templates/homebox.yaml" at <.Values.postgres.databases.homebox>:
    nil pointer evaluating interface {}.homebox

Use --debug flag to render out invalid YAML
```

### 300. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\hoppscotch`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\hoppscotch`

```text
Error: hoppscotch/templates/configmap.yaml:1:45
  executing "hoppscotch/templates/configmap.yaml" at <.Values.domain.hoppscotch>:
    nil pointer evaluating interface {}.hoppscotch

Use --debug flag to render out invalid YAML
```

### 301. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\immich`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\immich`

```text
Error: immich/templates/redis/deployment.yaml:67:29
  executing "immich/templates/redis/deployment.yaml" at <.Values.userspace.appData>:
    nil pointer evaluating interface {}.appData

Use --debug flag to render out invalid YAML
```

### 302. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\indexttsv2\indexttsv2server`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\indexttsv2\indexttsv2server`

```text
Error: indexttsv2server/templates/deployment.yaml:36:32
  executing "indexttsv2server/templates/deployment.yaml" at <.Values.olaresEnv.OLARES_USER_HUGGINGFACE_SERVICE>:
    nil pointer evaluating interface {}.OLARES_USER_HUGGINGFACE_SERVICE

Use --debug flag to render out invalid YAML
```

### 303. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\ipfs`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\ipfs`

```text
Error: ipfs/templates/deployment.yaml:1:41
  executing "ipfs/templates/deployment.yaml" at <.Values.domain.rpc>:
    nil pointer evaluating interface {}.rpc

Use --debug flag to render out invalid YAML
```

### 304. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\isaaclab`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\isaaclab`

```text
Error: isaaclab/templates/isaaclab.yaml:98:29
  executing "isaaclab/templates/isaaclab.yaml" at <.Values.userspace.appCache>:
    nil pointer evaluating interface {}.appCache

Use --debug flag to render out invalid YAML
```

### 305. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\jackett`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\jackett`

```text
Error: jackett/templates/jackett.yaml:62:29
  executing "jackett/templates/jackett.yaml" at <.Values.userspace.appCache>:
    nil pointer evaluating interface {}.appCache

Use --debug flag to render out invalid YAML
```

### 306. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\jdownloader2`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\jdownloader2`

```text
Error: jdownloader2/templates/deployment.yaml:104:28
  executing "jdownloader2/templates/deployment.yaml" at <.Values.userspace.userData>:
    nil pointer evaluating interface {}.userData

Use --debug flag to render out invalid YAML
```

### 307. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\jellyfin`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\jellyfin`

```text
Error: jellyfin/templates/deployment.yaml:46:28
  executing "jellyfin/templates/deployment.yaml" at <.Values.userspace.userData>:
    nil pointer evaluating interface {}.userData

Use --debug flag to render out invalid YAML
```

### 308. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\jupyterhub`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\jupyterhub`

```text
Error: jupyterhub/templates/ingress.yaml:29:44
  executing "jupyterhub/templates/ingress.yaml" at <.Values.bfl.username>:
    nil pointer evaluating interface {}.username

Use --debug flag to render out invalid YAML
```

### 309. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\karakeep`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\karakeep`

```text
Error: karakeep/templates/deployment.yaml:1:43
  executing "karakeep/templates/deployment.yaml" at <.Values.domain.karakeep>:
    nil pointer evaluating interface {}.karakeep

Use --debug flag to render out invalid YAML
```

### 310. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\koboldcpp\koboldcppserver`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\koboldcpp\koboldcppserver`

```text
Error: koboldcppserver/templates/download.yaml:23:32
  executing "koboldcppserver/templates/download.yaml" at <.Values.olaresEnv.OLARES_USER_HUGGINGFACE_SERVICE>:
    nil pointer evaluating interface {}.OLARES_USER_HUGGINGFACE_SERVICE

Use --debug flag to render out invalid YAML
```

### 311. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\komga`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\komga`

```text
Error: komga/templates/deployment.yaml:57:29
  executing "komga/templates/deployment.yaml" at <.Values.userspace.appData>:
    nil pointer evaluating interface {}.appData

Use --debug flag to render out invalid YAML
```

### 312. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\langbot`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\langbot`

```text
Error: langbot/templates/langbot.yaml:58:29
  executing "langbot/templates/langbot.yaml" at <.Values.userspace.appData>:
    nil pointer evaluating interface {}.appData

Use --debug flag to render out invalid YAML
```

### 313. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\langfuse`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\langfuse`

```text
Error: langfuse/templates/redis/deployment.yaml:87:29
  executing "langfuse/templates/redis/deployment.yaml" at <.Values.userspace.appData>:
    nil pointer evaluating interface {}.appData

Use --debug flag to render out invalid YAML
```

### 314. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\larescompanion`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\larescompanion`

```text
Error: larescompanion/templates/deployment.yaml:30:28
  executing "larescompanion/templates/deployment.yaml" at <.Values.olaresEnv.VERIFIABLE_CREDENTIAL>:
    nil pointer evaluating interface {}.VERIFIABLE_CREDENTIAL

Use --debug flag to render out invalid YAML
```

### 315. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\leantime`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\leantime`

```text
Error: leantime/templates/leantime.yaml:46:32
  executing "leantime/templates/leantime.yaml" at <.Values.mysql.host>:
    nil pointer evaluating interface {}.host

Use --debug flag to render out invalid YAML
```

### 316. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\librechat`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\librechat`

```text
Error: librechat/templates/configmap.yaml:1:38
  executing "librechat/templates/configmap.yaml" at <.Values.domain.librechatclient>:
    nil pointer evaluating interface {}.librechatclient

Use --debug flag to render out invalid YAML
```

### 317. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\lidarr`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\lidarr`

```text
Error: lidarr/templates/deployment.yaml:106:26
  executing "lidarr/templates/deployment.yaml" at <.Values.userspace.appData>:
    nil pointer evaluating interface {}.appData

Use --debug flag to render out invalid YAML
```

### 318. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\listmonk`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\listmonk`

```text
Error: listmonk/templates/listmonk.yaml:44:32
  executing "listmonk/templates/listmonk.yaml" at <.Values.postgres.databases.listmonk>:
    nil pointer evaluating interface {}.listmonk

Use --debug flag to render out invalid YAML
```

### 319. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\litellm`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\litellm`

```text
Error: litellm/templates/deployment.yaml:33:28
  executing "litellm/templates/deployment.yaml" at <.Values.olaresEnv.LITELLM_MASTER_KEY>:
    nil pointer evaluating interface {}.LITELLM_MASTER_KEY

Use --debug flag to render out invalid YAML
```

### 320. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\llamafactory`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\llamafactory`

```text
Error: llamafactory/templates/deployment.yaml:88:27
  executing "llamafactory/templates/deployment.yaml" at <.Values.userspace.appCache>:
    nil pointer evaluating interface {}.appCache

Use --debug flag to render out invalid YAML
```

### 321. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\llmgateway`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\llmgateway`

```text
Error: llmgateway/templates/configmap.yaml:25:23
  executing "llmgateway/templates/configmap.yaml" at <.Values.postgres.host>:
    nil pointer evaluating interface {}.host

Use --debug flag to render out invalid YAML
```

### 322. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\lobechat`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\lobechat`

```text
Error: lobechat/templates/postgres.yaml:74:29
  executing "lobechat/templates/postgres.yaml" at <.Values.userspace.appData>:
    nil pointer evaluating interface {}.appData

Use --debug flag to render out invalid YAML
```

### 323. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\macos`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\macos`

```text
Error: macos/templates/deployment.yaml:101:32
  executing "macos/templates/deployment.yaml" at <.Values.olaresEnv.DISK_SIZE>:
    nil pointer evaluating interface {}.DISK_SIZE

Use --debug flag to render out invalid YAML
```

### 324. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\magicpig`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\magicpig`

```text
Error: magicpig/templates/deployment.yaml:28:28
  executing "magicpig/templates/deployment.yaml" at <.Values.olaresEnv.VERIFIABLE_CREDENTIAL>:
    nil pointer evaluating interface {}.VERIFIABLE_CREDENTIAL

Use --debug flag to render out invalid YAML
```

### 325. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\mastodon`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\mastodon`

```text
Error: mastodon/templates/redis/statefulset.yaml:166:29
  executing "mastodon/templates/redis/statefulset.yaml" at <.Values.userspace.appCache>:
    nil pointer evaluating interface {}.appCache

Use --debug flag to render out invalid YAML
```

### 326. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\mattermost`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\mattermost`

```text
Error: mattermost/templates/mattermost.yaml:1:47
  executing "mattermost/templates/mattermost.yaml" at <.Values.domain.mattermost>:
    nil pointer evaluating interface {}.mattermost

Use --debug flag to render out invalid YAML
```

### 327. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\mealie`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\mealie`

```text
Error: mealie/templates/mealie.yaml:1:43
  executing "mealie/templates/mealie.yaml" at <.Values.domain.mealie>:
    nil pointer evaluating interface {}.mealie

Use --debug flag to render out invalid YAML
```

### 328. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\medusa`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\medusa`

```text
Error: medusa/templates/deployment.yaml:113:40
  executing "medusa/templates/deployment.yaml" at <.Values.redis.host>:
    nil pointer evaluating interface {}.host

Use --debug flag to render out invalid YAML
```

### 329. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\memos`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\memos`

```text
Error: memos/templates/memos.yaml:39:117
  executing "memos/templates/memos.yaml" at <.Values.postgres.databases.memosdb>:
    nil pointer evaluating interface {}.memosdb

Use --debug flag to render out invalid YAML
```

### 330. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\merchant`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\merchant`

```text
Error: merchant/templates/vc-system.yaml:48:29
  executing "merchant/templates/vc-system.yaml" at <.Values.postgres.host>:
    nil pointer evaluating interface {}.host

Use --debug flag to render out invalid YAML
```

### 331. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\metabase`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\metabase`

```text
Error: metabase/templates/configmap.yaml:1:43
  executing "metabase/templates/configmap.yaml" at <.Values.domain.metabase>:
    nil pointer evaluating interface {}.metabase

Use --debug flag to render out invalid YAML
```

### 332. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\migpt`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\migpt`

```text
Error: migpt/templates/configmap.yaml:8:23
  executing "migpt/templates/configmap.yaml" at <.Values.olaresEnv.MIGPT_USERNAME>:
    nil pointer evaluating interface {}.MIGPT_USERNAME

Use --debug flag to render out invalid YAML
```

### 333. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\miniflux`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\miniflux`

```text
Error: miniflux/templates/miniflux-deployment.yaml:28:32
  executing "miniflux/templates/miniflux-deployment.yaml" at <.Values.olaresEnv.ADMIN_PASSWORD>:
    nil pointer evaluating interface {}.ADMIN_PASSWORD

Use --debug flag to render out invalid YAML
```

### 334. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\n8n`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\n8n`

```text
Error: n8n/templates/n8n.yaml:45:28
  executing "n8n/templates/n8n.yaml" at <.Values.userspace.appData>:
    nil pointer evaluating interface {}.appData

Use --debug flag to render out invalid YAML
```

### 335. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\navidrome`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\navidrome`

```text
Error: navidrome/templates/deployment.yaml:70:27
  executing "navidrome/templates/deployment.yaml" at <.Values.userspace.appData>:
    nil pointer evaluating interface {}.appData

Use --debug flag to render out invalid YAML
```

### 336. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\nemoclaw`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\nemoclaw`

```text
Error: nemoclaw/templates/deployment.yaml:2094:29
  executing "nemoclaw/templates/deployment.yaml" at <.Values.userspace.appData>:
    nil pointer evaluating interface {}.appData

Use --debug flag to render out invalid YAML
```

### 337. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\netdata`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\netdata`

```text
Error: netdata/templates/configmap.yaml:15:29
  executing "netdata/templates/configmap.yaml" at <.Values.olaresEnv.SMTP_HOST>:
    nil pointer evaluating interface {}.SMTP_HOST

Use --debug flag to render out invalid YAML
```

### 338. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\nextcloud`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\nextcloud`

```text
Error: nextcloud/templates/deployment.yaml:1:44
  executing "nextcloud/templates/deployment.yaml" at <.Values.domain.nextcloud>:
    nil pointer evaluating interface {}.nextcloud

Use --debug flag to render out invalid YAML
```

### 339. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\nocobase`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\nocobase`

```text
Error: nocobase/templates/deployment.yaml:40:27
  executing "nocobase/templates/deployment.yaml" at <.Values.postgres.host>:
    nil pointer evaluating interface {}.host

Use --debug flag to render out invalid YAML
```

### 340. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\nocodb`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\nocodb`

```text
Error: nocodb/templates/deployment.yaml:53:37
  executing "nocodb/templates/deployment.yaml" at <.Values.postgres.host>:
    nil pointer evaluating interface {}.host

Use --debug flag to render out invalid YAML
```

### 341. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\nomad`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\nomad`

```text
Error: nomad/templates/deployment.yaml:154:27
  executing "nomad/templates/deployment.yaml" at <.Values.userspace.appData>:
    nil pointer evaluating interface {}.appData

Use --debug flag to render out invalid YAML
```

### 342. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\nostream`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\nostream`

```text
Error: nostream/templates/deployment.yaml:33:28
  executing "nostream/templates/deployment.yaml" at <.Values.postgres.username>:
    nil pointer evaluating interface {}.username

Use --debug flag to render out invalid YAML
```

### 343. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\novella`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\novella`

```text
Error: novella/templates/deployment.yaml:9:43
  executing "novella/templates/deployment.yaml" at <.Values.bfl.username>:
    nil pointer evaluating interface {}.username

Use --debug flag to render out invalid YAML
```

### 344. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\ntfy`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\ntfy`

```text
Error: ntfy/templates/ntfy.yaml:73:29
  executing "ntfy/templates/ntfy.yaml" at <.Values.userspace.appCache>:
    nil pointer evaluating interface {}.appCache

Use --debug flag to render out invalid YAML
```

### 345. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\nzbget`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\nzbget`

```text
Error: nzbget/templates/deployment.yaml:75:26
  executing "nzbget/templates/deployment.yaml" at <.Values.userspace.appData>:
    nil pointer evaluating interface {}.appData

Use --debug flag to render out invalid YAML
```

### 346. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\obsidian`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\obsidian`

```text
Error: obsidian/templates/deployment.yaml:1:35
  executing "obsidian/templates/deployment.yaml" at <.Values.userspace.appData>:
    nil pointer evaluating interface {}.appData

Use --debug flag to render out invalid YAML
```

### 347. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\odoo`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\odoo`

```text
Error: odoo/templates/odoo.yaml:42:32
  executing "odoo/templates/odoo.yaml" at <.Values.postgres.databases.odoo>:
    nil pointer evaluating interface {}.odoo

Use --debug flag to render out invalid YAML
```

### 348. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\onlyoffice`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\onlyoffice`

```text
Error: onlyoffice/templates/proxy.yaml:18:44
  executing "onlyoffice/templates/proxy.yaml" at <.Values.bfl.username>:
    nil pointer evaluating interface {}.username

Use --debug flag to render out invalid YAML
```

### 349. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\onlyofficev2\onlyofficev2`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\onlyofficev2\onlyofficev2`

```text
Error: onlyofficev2/templates/clientproxy.yaml:23:44
  executing "onlyofficev2/templates/clientproxy.yaml" at <.Values.bfl.username>:
    nil pointer evaluating interface {}.username

Use --debug flag to render out invalid YAML
```

### 350. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\opencode`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\opencode`

```text
Error: opencode/templates/opencode.yaml:1:45
  executing "opencode/templates/opencode.yaml" at <.Values.domain.opencode>:
    nil pointer evaluating interface {}.opencode

Use --debug flag to render out invalid YAML
```

### 351. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\openedaispeech`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\openedaispeech`

```text
Error: openedaispeech/templates/deployment.yaml:95:27
  executing "openedaispeech/templates/deployment.yaml" at <.Values.userspace.appCache>:
    nil pointer evaluating interface {}.appCache

Use --debug flag to render out invalid YAML
```

### 352. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\openllm`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\openllm`

```text
Error: openllm/templates/traefik.yaml:296:59
  executing "openllm/templates/traefik.yaml" at <.Values.user.zone>:
    nil pointer evaluating interface {}.zone

Use --debug flag to render out invalid YAML
```

### 353. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\opennotebook`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\opennotebook`

```text
Error: opennotebook/templates/deployment.yaml:113:27
  executing "opennotebook/templates/deployment.yaml" at <.Values.userspace.appData>:
    nil pointer evaluating interface {}.appData

Use --debug flag to render out invalid YAML
```

### 354. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\openwebui`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\openwebui`

```text
Error: openwebui/templates/deployment.yaml:1:38
  executing "openwebui/templates/deployment.yaml" at <.Values.domain.openwebui>:
    nil pointer evaluating interface {}.openwebui

Use --debug flag to render out invalid YAML
```

### 355. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\otmoiclp`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\otmoiclp`

```text
Error: otmoiclp/templates/traefik.yaml:1:44
  executing "otmoiclp/templates/traefik.yaml" at <.Values.domain.traefik>:
    nil pointer evaluating interface {}.traefik

Use --debug flag to render out invalid YAML
```

### 356. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\otmoicrelay`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\otmoicrelay`

```text
Error: otmoicrelay/templates/traefik.yaml:1:44
  executing "otmoicrelay/templates/traefik.yaml" at <.Values.domain.traefik>:
    nil pointer evaluating interface {}.traefik

Use --debug flag to render out invalid YAML
```

### 357. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\paperclip`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\paperclip`

```text
Error: paperclip/templates/deployment.yaml:1:46
  executing "paperclip/templates/deployment.yaml" at <.Values.domain.paperclip>:
    nil pointer evaluating interface {}.paperclip

Use --debug flag to render out invalid YAML
```

### 358. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\pdfmathtranslate`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\pdfmathtranslate`

```text
Error: pdfmathtranslate/templates/deployment.yaml:46:29
  executing "pdfmathtranslate/templates/deployment.yaml" at <.Values.userspace.appData>:
    nil pointer evaluating interface {}.appData

Use --debug flag to render out invalid YAML
```

### 359. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\pds`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\pds`

```text
Error: pds/templates/deployment.yaml:1:38
  executing "pds/templates/deployment.yaml" at <.Values.domain.pds>:
    nil pointer evaluating interface {}.pds

Use --debug flag to render out invalid YAML
```

### 360. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\penpot`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\penpot`

```text
Error: penpot/templates/redis/deployment.yaml:62:29
  executing "penpot/templates/redis/deployment.yaml" at <.Values.userspace.appData>:
    nil pointer evaluating interface {}.appData

Use --debug flag to render out invalid YAML
```

### 361. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\perplexica`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\perplexica`

```text
Error: perplexica/templates/deployment.yaml:1:38
  executing "perplexica/templates/deployment.yaml" at <.Values.domain.pbe>:
    nil pointer evaluating interface {}.pbe

Use --debug flag to render out invalid YAML
```

### 362. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\photoprism`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\photoprism`

```text
Error: photoprism/templates/deployment.yaml:64:28
  executing "photoprism/templates/deployment.yaml" at <.Values.olaresEnv.PHOTOPRISM_ADMIN_PASSWORD>:
    nil pointer evaluating interface {}.PHOTOPRISM_ADMIN_PASSWORD

Use --debug flag to render out invalid YAML
```

### 363. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\photoview`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\photoview`

```text
Error: photoview/templates/deployment.yaml:59:28
  executing "photoview/templates/deployment.yaml" at <.Values.mariadb.username>:
    nil pointer evaluating interface {}.username

Use --debug flag to render out invalid YAML
```

### 364. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\plane`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\plane`

```text
Error: plane/templates/deployment.yaml:1:42
  executing "plane/templates/deployment.yaml" at <.Values.domain.plane>:
    nil pointer evaluating interface {}.plane

Use --debug flag to render out invalid YAML
```

### 365. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\prometheusclient`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\prometheusclient`

```text
Error: prometheusclient/templates/deployment.yaml:66:18
  executing "prometheusclient/templates/deployment.yaml" at <.Values.bfl.username>:
    nil pointer evaluating interface {}.username

Use --debug flag to render out invalid YAML
```

### 366. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\prowlarr`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\prowlarr`

```text
Error: prowlarr/templates/deployment.yaml:57:26
  executing "prowlarr/templates/deployment.yaml" at <.Values.userspace.appData>:
    nil pointer evaluating interface {}.appData

Use --debug flag to render out invalid YAML
```

### 367. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\qbittorrent`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\qbittorrent`

```text
Error: qbittorrent/templates/deployment.yaml:82:28
  executing "qbittorrent/templates/deployment.yaml" at <.Values.userspace.userData>:
    nil pointer evaluating interface {}.userData

Use --debug flag to render out invalid YAML
```

### 368. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\qinglong`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\qinglong`

```text
Error: qinglong/templates/qinglong.yaml:62:26
  executing "qinglong/templates/qinglong.yaml" at <.Values.userspace.appCache>:
    nil pointer evaluating interface {}.appCache

Use --debug flag to render out invalid YAML
```

### 369. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\radarr`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\radarr`

```text
Error: radarr/templates/deployment.yaml:111:26
  executing "radarr/templates/deployment.yaml" at <.Values.userspace.appData>:
    nil pointer evaluating interface {}.appData

Use --debug flag to render out invalid YAML
```

### 370. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\radicale`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\radicale`

```text
Error: radicale/templates/all.yaml:131:29
  executing "radicale/templates/all.yaml" at <.Values.userspace.appData>:
    nil pointer evaluating interface {}.appData

Use --debug flag to render out invalid YAML
```

### 371. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\ragflow`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\ragflow`

```text
Error: ragflow/templates/redis.yaml:60:27
  executing "ragflow/templates/redis.yaml" at <.Values.userspace.appCache>:
    nil pointer evaluating interface {}.appCache

Use --debug flag to render out invalid YAML
```

### 372. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\rallly`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\rallly`

```text
Error: rallly/templates/rallly.yaml:34:165
  executing "rallly/templates/rallly.yaml" at <.Values.postgres.databases.rallly>:
    nil pointer evaluating interface {}.rallly

Use --debug flag to render out invalid YAML
```

### 373. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\rembgv2\rembgv2server`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\rembgv2\rembgv2server`

```text
Error: rembgv2server/templates/deployment.yaml:63:29
  executing "rembgv2server/templates/deployment.yaml" at <.Values.userspace.userData>:
    nil pointer evaluating interface {}.userData

Use --debug flag to render out invalid YAML
```

### 374. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\rnasequencing`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\rnasequencing`

```text
Error: rnasequencing/templates/rnasequencing.yaml:88:29
  executing "rnasequencing/templates/rnasequencing.yaml" at <.Values.userspace.appCache>:
    nil pointer evaluating interface {}.appCache

Use --debug flag to render out invalid YAML
```

### 375. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\rssubscribe\rsserver`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\rssubscribe\rsserver`

```text
Error: rsserver/templates/rssubscribe.yaml:124:30
  executing "rsserver/templates/rssubscribe.yaml" at <.Values.postgres.username>:
    nil pointer evaluating interface {}.username

Use --debug flag to render out invalid YAML
```

### 376. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\rssubscribe\rssubscribe`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\rssubscribe\rssubscribe`

```text
Error: rssubscribe/templates/deployment.yaml:22:46
  executing "rssubscribe/templates/deployment.yaml" at <.Values.bfl.username>:
    nil pointer evaluating interface {}.username

Use --debug flag to render out invalid YAML
```

### 377. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\sdwebuisharev2\sdwebuisharev2server`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\sdwebuisharev2\sdwebuisharev2server`

```text
Error: sdwebuisharev2server/templates/deployment.yaml:206:26
  executing "sdwebuisharev2server/templates/deployment.yaml" at <.Values.userspace.appCache>:
    nil pointer evaluating interface {}.appCache

Use --debug flag to render out invalid YAML
```

### 378. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\searxngv2\searxngv2server`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\searxngv2\searxngv2server`

```text
Error: searxngv2server/templates/searxng-config.yaml:91:30
  executing "searxngv2server/templates/searxng-config.yaml" at <.Values.redis.password>:
    nil pointer evaluating interface {}.password

Use --debug flag to render out invalid YAML
```

### 379. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\seatable`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\seatable`

```text
Error: seatable/templates/mariadb/deployment.yaml:87:29
  executing "seatable/templates/mariadb/deployment.yaml" at <.Values.userspace.appData>:
    nil pointer evaluating interface {}.appData

Use --debug flag to render out invalid YAML
```

### 380. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\secondme`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\secondme`

```text
Error: secondme/templates/secondme.yaml:130:29
  executing "secondme/templates/secondme.yaml" at <.Values.userspace.appData>:
    nil pointer evaluating interface {}.appData

Use --debug flag to render out invalid YAML
```

### 381. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\showdoc`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\showdoc`

```text
Error: showdoc/templates/showdoc.yaml:3:16
  executing "showdoc/templates/showdoc.yaml" at <.Values.cluster.arch>:
    nil pointer evaluating interface {}.arch

Use --debug flag to render out invalid YAML
```

### 382. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\sickchill`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\sickchill`

```text
Error: sickchill/templates/deployment.yaml:82:26
  executing "sickchill/templates/deployment.yaml" at <.Values.userspace.appData>:
    nil pointer evaluating interface {}.appData

Use --debug flag to render out invalid YAML
```

### 383. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\sillytavern`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\sillytavern`

```text
Error: sillytavern/templates/deployment.yaml:1:42
  executing "sillytavern/templates/deployment.yaml" at <.Values.domain.sillytavern>:
    nil pointer evaluating interface {}.sillytavern

Use --debug flag to render out invalid YAML
```

### 384. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\solidtime`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\solidtime`

```text
Error: solidtime/templates/solidtime.yaml:1:46
  executing "solidtime/templates/solidtime.yaml" at <.Values.domain.solidtime>:
    nil pointer evaluating interface {}.solidtime

Use --debug flag to render out invalid YAML
```

### 385. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\sonarr`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\sonarr`

```text
Error: sonarr/templates/deployment.yaml:106:26
  executing "sonarr/templates/deployment.yaml" at <.Values.userspace.appData>:
    nil pointer evaluating interface {}.appData

Use --debug flag to render out invalid YAML
```

### 386. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\steamheadless`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\steamheadless`

```text
Error: steamheadless/templates/steam.yaml:70:31
  executing "steamheadless/templates/steam.yaml" at <.Values.domain.steamheadlessaudio>:
    nil pointer evaluating interface {}.steamheadlessaudio

Use --debug flag to render out invalid YAML
```

### 387. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\stirlingpdf`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\stirlingpdf`

```text
Error: stirlingpdf/templates/stirlingpdf.yaml:83:105
  executing "stirlingpdf/templates/stirlingpdf.yaml" at <.Values.userspace.appData>:
    nil pointer evaluating interface {}.appData

Use --debug flag to render out invalid YAML
```

### 388. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\studio\studio`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\studio\studio`

```text
Error: studio/templates/front.yaml:294:71
  executing "studio/templates/front.yaml" at <.Values.bfl.username>:
    nil pointer evaluating interface {}.username

Use --debug flag to render out invalid YAML
```

### 389. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\studio\studioserver`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\studio\studioserver`

```text
Error: studioserver/templates/studio_server_deploy.yaml:268:31
  executing "studioserver/templates/studio_server_deploy.yaml" at <.Values.postgres.host>:
    nil pointer evaluating interface {}.host

Use --debug flag to render out invalid YAML
```

### 390. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\teable`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\teable`

```text
Error: teable/templates/redis/deployment.yaml:83:29
  executing "teable/templates/redis/deployment.yaml" at <.Values.userspace.appData>:
    nil pointer evaluating interface {}.appData

Use --debug flag to render out invalid YAML
```

### 391. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\tensorzero`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\tensorzero`

```text
Error: tensorzero/templates/tensorzero.yaml:203:43
  executing "tensorzero/templates/tensorzero.yaml" at <.Values.postgres.username>:
    nil pointer evaluating interface {}.username

Use --debug flag to render out invalid YAML
```

### 392. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\testapp1`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\testapp1`

```text
Error: testapp1/templates/testapp1.yaml:62:32
  executing "testapp1/templates/testapp1.yaml" at <.Values.postgres.databases.testapp1>:
    nil pointer evaluating interface {}.testapp1

Use --debug flag to render out invalid YAML
```

### 393. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\testapp5`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\testapp5`

```text
Error: testapp5/templates/testapp5.yaml:62:32
  executing "testapp5/templates/testapp5.yaml" at <.Values.postgres.databases.testapp5>:
    nil pointer evaluating interface {}.testapp5

Use --debug flag to render out invalid YAML
```

### 394. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\testenv`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\testenv`

```text
Error: testenv/templates/firefox.yaml:63:32
  executing "testenv/templates/firefox.yaml" at <.Values.olaresEnv.TESTENV_CUDA>:
    nil pointer evaluating interface {}.TESTENV_CUDA

Use --debug flag to render out invalid YAML
```

### 395. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\testnsfw`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\testnsfw`

```text
Error: testnsfw/templates/homebox.yaml:62:32
  executing "testnsfw/templates/homebox.yaml" at <.Values.postgres.databases.testnsfw>:
    nil pointer evaluating interface {}.testnsfw

Use --debug flag to render out invalid YAML
```

### 396. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\testpayment`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\testpayment`

```text
Error: testpayment/templates/firefox.yaml:59:32
  executing "testpayment/templates/firefox.yaml" at <.Values.olaresEnv.VERIFIABLE_CREDENTIAL>:
    nil pointer evaluating interface {}.VERIFIABLE_CREDENTIAL

Use --debug flag to render out invalid YAML
```

### 397. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\testpid`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\testpid`

```text
Error: testpid/templates/firefox.yaml:95:28
  executing "testpid/templates/firefox.yaml" at <.Values.userspace.appCache>:
    nil pointer evaluating interface {}.appCache

Use --debug flag to render out invalid YAML
```

### 398. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\texttoimagesearch`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\texttoimagesearch`

```text
Error: texttoimagesearch/templates/service.yaml:25:18
  executing "texttoimagesearch/templates/service.yaml" at <.Values.bfl.username>:
    nil pointer evaluating interface {}.username

Use --debug flag to render out invalid YAML
```

### 399. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\tradingagents`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\tradingagents`

```text
Error: tradingagents/templates/tradingagents.yaml:150:29
  executing "tradingagents/templates/tradingagents.yaml" at <.Values.userspace.appData>:
    nil pointer evaluating interface {}.appData

Use --debug flag to render out invalid YAML
```

### 400. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\transmission`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\transmission`

```text
Error: transmission/templates/deployment.yaml:91:26
  executing "transmission/templates/deployment.yaml" at <.Values.userspace.appData>:
    nil pointer evaluating interface {}.appData

Use --debug flag to render out invalid YAML
```

### 401. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\trek`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\trek`

```text
Error: trek/templates/deployment.yaml:1:41
  executing "trek/templates/deployment.yaml" at <.Values.domain.trek>:
    nil pointer evaluating interface {}.trek

Use --debug flag to render out invalid YAML
```

### 402. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\twenty`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\twenty`

```text
Error: twenty/templates/worker-deployment.yaml:1:38
  executing "twenty/templates/worker-deployment.yaml" at <.Values.domain.crm>:
    nil pointer evaluating interface {}.crm

Use --debug flag to render out invalid YAML
```

### 403. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\twitter\twitter`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\twitter\twitter`

```text
Error: twitter/templates/deployment.yaml:22:46
  executing "twitter/templates/deployment.yaml" at <.Values.bfl.username>:
    nil pointer evaluating interface {}.username

Use --debug flag to render out invalid YAML
```

### 404. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\uptimekuma`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\uptimekuma`

```text
Error: uptimekuma/templates/uptimekuma.yaml:50:29
  executing "uptimekuma/templates/uptimekuma.yaml" at <.Values.userspace.appData>:
    nil pointer evaluating interface {}.appData

Use --debug flag to render out invalid YAML
```

### 405. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\vane`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\vane`

```text
Error: vane/templates/deployment.yaml:1:38
  executing "vane/templates/deployment.yaml" at <.Values.domain.vane>:
    nil pointer evaluating interface {}.vane

Use --debug flag to render out invalid YAML
```

### 406. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\vllmdeepseekocr3bv2\vllmserver`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\vllmdeepseekocr3bv2\vllmserver`

```text
Error: vllmserver/templates/deployment.yaml:1:39
  executing "vllmserver/templates/deployment.yaml" at <.Values.domain.vllmclient>:
    nil pointer evaluating interface {}.vllmclient

Use --debug flag to render out invalid YAML
```

### 407. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\vllmgemma312bitv2\vllmgemma312bitv2server`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\vllmgemma312bitv2\vllmgemma312bitv2server`

```text
Error: vllmgemma312bitv2server/templates/download.yaml:1:40
  executing "vllmgemma312bitv2server/templates/download.yaml" at <.Values.domain.vllmclient>:
    nil pointer evaluating interface {}.vllmclient

Use --debug flag to render out invalid YAML
```

### 408. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\vllmgemma327bqatv2\vllmgemma327bqatv2server`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\vllmgemma327bqatv2\vllmgemma327bqatv2server`

```text
Error: vllmgemma327bqatv2server/templates/download.yaml:1:39
  executing "vllmgemma327bqatv2server/templates/download.yaml" at <.Values.domain.vllmclient>:
    nil pointer evaluating interface {}.vllmclient

Use --debug flag to render out invalid YAML
```

### 409. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\vllmgptoss20bv2\vllmgptoss20bv2server`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\vllmgptoss20bv2\vllmgptoss20bv2server`

```text
Error: vllmgptoss20bv2server/templates/download.yaml:1:39
  executing "vllmgptoss20bv2server/templates/download.yaml" at <.Values.domain.vllmclient>:
    nil pointer evaluating interface {}.vllmclient

Use --debug flag to render out invalid YAML
```

### 410. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\vllmhymt1518bv2\vllmhymt1518bv2server`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\vllmhymt1518bv2\vllmhymt1518bv2server`

```text
Error: vllmhymt1518bv2server/templates/download.yaml:1:39
  executing "vllmhymt1518bv2server/templates/download.yaml" at <.Values.domain.vllmclient>:
    nil pointer evaluating interface {}.vllmclient

Use --debug flag to render out invalid YAML
```

### 411. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\vllmhymt157bv2\vllmhymt157bv2server`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\vllmhymt157bv2\vllmhymt157bv2server`

```text
Error: vllmhymt157bv2server/templates/download.yaml:1:39
  executing "vllmhymt157bv2server/templates/download.yaml" at <.Values.domain.vllmclient>:
    nil pointer evaluating interface {}.vllmclient

Use --debug flag to render out invalid YAML
```

### 412. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\vllmqwen330ba3binstruct4bitv2\vllmqwen330ba3bv2server`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\vllmqwen330ba3binstruct4bitv2\vllmqwen330ba3bv2server`

```text
Error: vllmqwen330ba3bv2server/templates/download.yaml:1:39
  executing "vllmqwen330ba3bv2server/templates/download.yaml" at <.Values.domain.vllmclient>:
    nil pointer evaluating interface {}.vllmclient

Use --debug flag to render out invalid YAML
```

### 413. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\wewerss`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\wewerss`

```text
Error: wewerss/templates/wewerss.yaml:36:31
  executing "wewerss/templates/wewerss.yaml" at <.Values.mysql.password>:
    nil pointer evaluating interface {}.password

Use --debug flag to render out invalid YAML
```

### 414. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\windows`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\windows`

```text
Error: windows/templates/deployment.yaml:103:32
  executing "windows/templates/deployment.yaml" at <.Values.olaresEnv.RAM_SIZE>:
    nil pointer evaluating interface {}.RAM_SIZE

Use --debug flag to render out invalid YAML
```

### 415. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\windowsarm`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\windowsarm`

```text
Error: windowsarm/templates/deployment.yaml:102:32
  executing "windowsarm/templates/deployment.yaml" at <.Values.olaresEnv.RAM_SIZE>:
    nil pointer evaluating interface {}.RAM_SIZE

Use --debug flag to render out invalid YAML
```

### 416. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\wise\knowledge`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\wise\knowledge`

```text
Error: knowledge/templates/knowledge_deployment.yaml:127:29
  executing "knowledge/templates/knowledge_deployment.yaml" at <.Values.postgres.host>:
    nil pointer evaluating interface {}.host

Use --debug flag to render out invalid YAML
```

### 417. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\wise\wise`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\wise\wise`

```text
Error: wise/templates/front.yaml:61:35
  executing "wise/templates/front.yaml" at <.Values.bfl.username>:
    nil pointer evaluating interface {}.username

Use --debug flag to render out invalid YAML
```

### 418. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\wordpresspure`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\wordpresspure`

```text
Error: wordpresspure/templates/wordpress.yaml:246:31
  executing "wordpresspure/templates/wordpress.yaml" at <.Values.mariadb.host>:
    nil pointer evaluating interface {}.host

Use --debug flag to render out invalid YAML
```

### 419. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\xinference`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\xinference`

```text
Error: xinference/templates/deployment.yaml:63:30
  executing "xinference/templates/deployment.yaml" at <.Values.olaresEnv.OLARES_USER_HUGGINGFACE_SERVICE>:
    nil pointer evaluating interface {}.OLARES_USER_HUGGINGFACE_SERVICE

Use --debug flag to render out invalid YAML
```

### 420. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\xybotv2`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\xybotv2`

```text
Error: xybotv2/templates/xybotv2.yaml:73:29
  executing "xybotv2/templates/xybotv2.yaml" at <.Values.userspace.appData>:
    nil pointer evaluating interface {}.appData

Use --debug flag to render out invalid YAML
```

### 421. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\ytdlp\ytdlp`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\ytdlp\ytdlp`

```text
Error: ytdlp/templates/deployment.yaml:22:46
  executing "ytdlp/templates/deployment.yaml" at <.Values.bfl.username>:
    nil pointer evaluating interface {}.username

Use --debug flag to render out invalid YAML
```

### 422. `Above-Os/terminus-apps`

- Chart: `D:\helm_clones_github\Above-Os__terminus-apps\ytnavigator`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Above-Os__terminus-apps\ytnavigator`

```text
Error: ytnavigator/templates/deployment.yaml:28:28
  executing "ytnavigator/templates/deployment.yaml" at <.Values.postgres.databases.ytnavigator>:
    nil pointer evaluating interface {}.databases

Use --debug flag to render out invalid YAML
```

### 423. `CDCgov/NEDSS-Helm`

- Chart: `D:\helm_clones_github\CDCgov__NEDSS-Helm\argocd`
- Source: `template`
- Values files: `D:\helm_clones_github\CDCgov__NEDSS-Helm\argocd\template-values.yaml`
- Command: `helm template test D:\helm_clones_github\CDCgov__NEDSS-Helm\argocd -f D:\helm_clones_github\CDCgov__NEDSS-Helm\argocd\template-values.yaml`

```text
Error: NBS7/templates/rtr.yaml:30:27
  executing "NBS7/templates/rtr.yaml" at <.Values.rtr.threadPoolSize>:
    nil pointer evaluating interface {}.threadPoolSize

Use --debug flag to render out invalid YAML
```

### 424. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`

```text
Error: bookshop/templates/cap-operator-cros-svc-ias.yaml:93:23
  executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 425. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-svc-ias.yaml:93:23
  executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 426. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`

```text
Error: bookshop/templates/cap-operator-cros-svc-ias.yaml:93:23
  executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 427. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`

```text
Error: bookshop/templates/cap-operator-cros-svc-ias.yaml:93:23
  executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 428. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`

```text
Error: bookshop/templates/cap-operator-cros-svc-ias.yaml:93:23
  executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 429. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`

```text
Error: bookshop/templates/cap-operator-cros-svc-ias.yaml:93:23
  executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 430. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`

```text
Error: bookshop/templates/cap-operator-cros-svc-ias.yaml:93:23
  executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 431. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`

```text
Error: bookshop/templates/cap-operator-cros-svc-ias.yaml:93:23
  executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 432. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 433. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 434. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 435. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 436. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 437. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 438. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 439. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 440. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`

```text
Error: bookshop/templates/cap-operator-cros-svc-ias.yaml:93:23
  executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 441. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`

```text
Error: bookshop/templates/cap-operator-cros-svc-ias.yaml:93:23
  executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 442. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`

```text
Error: bookshop/templates/cap-operator-cros-svc-ias.yaml:93:23
  executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 443. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`

```text
Error: bookshop/templates/cap-operator-cros-svc-ias.yaml:93:23
  executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 444. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`

```text
Error: bookshop/templates/cap-operator-cros-svc-ias.yaml:93:23
  executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 445. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`

```text
Error: bookshop/templates/cap-operator-cros-svc-ias.yaml:93:23
  executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 446. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`

```text
Error: bookshop/templates/cap-operator-cros-svc-ias.yaml:93:23
  executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 447. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`

```text
Error: bookshop/templates/cap-operator-cros-svc-ias.yaml:93:23
  executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 448. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 449. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 450. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 451. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 452. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 453. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 454. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 455. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 456. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: bookshop/templates/cap-operator-cros-svc-ias.yaml:93:23
  executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 457. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: bookshop/templates/cap-operator-cros-svc-ias.yaml:93:23
  executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 458. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: bookshop/templates/cap-operator-cros-svc-ias.yaml:93:23
  executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 459. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: bookshop/templates/cap-operator-cros-svc-ias.yaml:93:23
  executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 460. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: bookshop/templates/cap-operator-cros-svc-ias.yaml:93:23
  executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 461. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: bookshop/templates/cap-operator-cros-svc-ias.yaml:93:23
  executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 462. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: bookshop/templates/cap-operator-cros-svc-ias.yaml:93:23
  executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 463. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: bookshop/templates/cap-operator-cros-svc-ias.yaml:93:23
  executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 464. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 465. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 466. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 467. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 468. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 469. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 470. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 471. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 472. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: bookshop/templates/cap-operator-cros-svc-ias.yaml:93:23
  executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 473. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: bookshop/templates/cap-operator-cros-svc-ias.yaml:93:23
  executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 474. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: bookshop/templates/cap-operator-cros-svc-ias.yaml:93:23
  executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 475. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: bookshop/templates/cap-operator-cros-svc-ias.yaml:93:23
  executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 476. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: bookshop/templates/cap-operator-cros-svc-ias.yaml:93:23
  executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 477. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: bookshop/templates/cap-operator-cros-svc-ias.yaml:93:23
  executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 478. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: bookshop/templates/cap-operator-cros-svc-ias.yaml:93:23
  executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 479. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: bookshop/templates/cap-operator-cros-svc-ias.yaml:93:23
  executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 480. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 481. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 482. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 483. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 484. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 485. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 486. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 487. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 488. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 489. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 490. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 491. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 492. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 493. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 494. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 495. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 496. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 497. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 498. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 499. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 500. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 501. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 502. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 503. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 504. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 505. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 506. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 507. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 508. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 509. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 510. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 511. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 512. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 513. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 514. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 515. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 516. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 517. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 518. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 519. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 520. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 521. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 522. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 523. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 524. `cap-js/cap-operator-plugin`

- Chart: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`, `D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml -f D:\helm_clones_github\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: bookshop/templates/cap-operator-cros-mta.yaml:69:23
  executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 525. `punchplatform/punch-helm`

- Chart: `D:\helm_clones_github\punchplatform__punch-helm\injector`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\punchplatform__punch-helm\injector`

```text
Error: log-injector/templates/deployment.yaml:21:22
  executing "log-injector/templates/deployment.yaml" at <.Values.image.imagePullSecrets>:
    nil pointer evaluating interface {}.imagePullSecrets

Use --debug flag to render out invalid YAML
```

### 526. `CloudFitSoftware/PubSec-Info-Assistant`

- Chart: `D:\helm_clones_github\CloudFitSoftware__PubSec-Info-Assistant\charts\cert-manager-chart`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\CloudFitSoftware__PubSec-Info-Assistant\charts\cert-manager-chart`

```text
Error: cert-manager/templates/webhook-validating-webhook.yaml:10:8
  executing "cert-manager/templates/webhook-validating-webhook.yaml" at <include "labels" .>:
    error calling include:
cert-manager/templates/_helpers.tpl:159:14
  executing "labels" at <.Values.global.commonLabels>:
    nil pointer evaluating interface {}.commonLabels

Use --debug flag to render out invalid YAML
```

### 527. `CloudFitSoftware/PubSec-Info-Assistant`

- Chart: `D:\helm_clones_github\CloudFitSoftware__PubSec-Info-Assistant\charts\cert-manager-chart`
- Source: `template`
- Values files: `D:\helm_clones_github\CloudFitSoftware__PubSec-Info-Assistant\charts\cert-manager-chart\values-cfs.yaml`
- Command: `helm template test D:\helm_clones_github\CloudFitSoftware__PubSec-Info-Assistant\charts\cert-manager-chart -f D:\helm_clones_github\CloudFitSoftware__PubSec-Info-Assistant\charts\cert-manager-chart\values-cfs.yaml`

```text
Error: cert-manager/templates/webhook-validating-webhook.yaml:13:20
  executing "cert-manager/templates/webhook-validating-webhook.yaml" at <.Values.webhook.validatingWebhookConfigurationAnnotations>:
    nil pointer evaluating interface {}.validatingWebhookConfigurationAnnotations

Use --debug flag to render out invalid YAML
```

### 528. `CloudFitSoftware/PubSec-Info-Assistant`

- Chart: `D:\helm_clones_github\CloudFitSoftware__PubSec-Info-Assistant\charts\infoasst-enrichment`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\CloudFitSoftware__PubSec-Info-Assistant\charts\infoasst-enrichment`

```text
Error: infoasst-enrichment/templates/enrichment-deployment.yaml:18:25
  executing "infoasst-enrichment/templates/enrichment-deployment.yaml" at <.Values.images.enrichment.repository>:
    nil pointer evaluating interface {}.enrichment

Use --debug flag to render out invalid YAML
```

### 529. `CloudFitSoftware/PubSec-Info-Assistant`

- Chart: `D:\helm_clones_github\CloudFitSoftware__PubSec-Info-Assistant\charts\infoasst-llm`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\CloudFitSoftware__PubSec-Info-Assistant\charts\infoasst-llm`

```text
Error: infoasst-llm/templates/llm-deployment.yaml:18:25
  executing "infoasst-llm/templates/llm-deployment.yaml" at <.Values.images.llm.repository>:
    nil pointer evaluating interface {}.llm

Use --debug flag to render out invalid YAML
```

### 530. `CloudFitSoftware/PubSec-Info-Assistant`

- Chart: `D:\helm_clones_github\CloudFitSoftware__PubSec-Info-Assistant\charts\infoasst-reranker`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\CloudFitSoftware__PubSec-Info-Assistant\charts\infoasst-reranker`

```text
Error: infoasst-reranker/templates/reranker-deployment.yaml:18:25
  executing "infoasst-reranker/templates/reranker-deployment.yaml" at <.Values.images.reranker.repository>:
    nil pointer evaluating interface {}.reranker

Use --debug flag to render out invalid YAML
```

### 531. `CloudFitSoftware/PubSec-Info-Assistant`

- Chart: `D:\helm_clones_github\CloudFitSoftware__PubSec-Info-Assistant\charts\infoasst-t2v`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\CloudFitSoftware__PubSec-Info-Assistant\charts\infoasst-t2v`

```text
Error: infoasst-enrichment/templates/t2v-deployment.yaml:18:25
  executing "infoasst-enrichment/templates/t2v-deployment.yaml" at <.Values.images.t2v.repository>:
    nil pointer evaluating interface {}.t2v

Use --debug flag to render out invalid YAML
```

### 532. `CloudFitSoftware/PubSec-Info-Assistant`

- Chart: `D:\helm_clones_github\CloudFitSoftware__PubSec-Info-Assistant\charts\infoasst-weaviate-container`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\CloudFitSoftware__PubSec-Info-Assistant\charts\infoasst-weaviate-container`

```text
Error: infoasst-weaviate/templates/weaviate-container-deployment.yaml:18:25
  executing "infoasst-weaviate/templates/weaviate-container-deployment.yaml" at <.Values.images.weaviate.repository>:
    nil pointer evaluating interface {}.weaviate

Use --debug flag to render out invalid YAML
```

### 533. `CloudFitSoftware/PubSec-Info-Assistant`

- Chart: `D:\helm_clones_github\CloudFitSoftware__PubSec-Info-Assistant\charts\infoasst-webapp`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\CloudFitSoftware__PubSec-Info-Assistant\charts\infoasst-webapp`

```text
Error: infoasst-webapp/templates/webapp-deployment.yaml:18:25
  executing "infoasst-webapp/templates/webapp-deployment.yaml" at <.Values.images.webapp.repository>:
    nil pointer evaluating interface {}.webapp

Use --debug flag to render out invalid YAML
```

### 534. `CloudFitSoftware/PubSec-Info-Assistant`

- Chart: `D:\helm_clones_github\CloudFitSoftware__PubSec-Info-Assistant\charts\istio-controlplane-chart`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\CloudFitSoftware__PubSec-Info-Assistant\charts\istio-controlplane-chart`

```text
Error: istio-controlplane/templates/readyjob.yaml:16:27
  executing "istio-controlplane/templates/readyjob.yaml" at <.Values.images.readyJob.repository>:
    nil pointer evaluating interface {}.readyJob

Use --debug flag to render out invalid YAML
```

### 535. `CloudFitSoftware/PubSec-Info-Assistant`

- Chart: `D:\helm_clones_github\CloudFitSoftware__PubSec-Info-Assistant\charts\istio-operator-chart`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\CloudFitSoftware__PubSec-Info-Assistant\charts\istio-operator-chart`

```text
Error: istio-operator/templates/deployment.yaml:50:14
  executing "istio-operator/templates/deployment.yaml" at <.Values.operator.seccompProfile>:
    nil pointer evaluating interface {}.seccompProfile

Use --debug flag to render out invalid YAML
```

### 536. `Taipei-HUG/kubernetes-microservices`

- Chart: `D:\helm_clones_github\Taipei-HUG__kubernetes-microservices\examples\helm\details`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Taipei-HUG__kubernetes-microservices\examples\helm\details`

```text
Error: details/templates/tests/test-connection.yaml:6:8
  executing "details/templates/tests/test-connection.yaml" at <include "details.labels" .>:
    error calling include:
details/templates/_helpers.tpl:40:3
  executing "details.labels" at <include "details.istio.labels" .>:
    error calling include:
details/templates/_helpers.tpl:67:14
  executing "details.istio.labels" at <.Values.global.istioEnabled>:
    nil pointer evaluating interface {}.istioEnabled

Use --debug flag to render out invalid YAML
```

### 537. `Taipei-HUG/kubernetes-microservices`

- Chart: `D:\helm_clones_github\Taipei-HUG__kubernetes-microservices\examples\helm\productpage`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Taipei-HUG__kubernetes-microservices\examples\helm\productpage`

```text
Error: productpage/templates/tests/test-connection.yaml:6:8
  executing "productpage/templates/tests/test-connection.yaml" at <include "productpage.labels" .>:
    error calling include:
productpage/templates/_helpers.tpl:40:3
  executing "productpage.labels" at <include "productpage.istio.labels" .>:
    error calling include:
productpage/templates/_helpers.tpl:67:14
  executing "productpage.istio.labels" at <.Values.global.istioEnabled>:
    nil pointer evaluating interface {}.istioEnabled

Use --debug flag to render out invalid YAML
```

### 538. `Taipei-HUG/kubernetes-microservices`

- Chart: `D:\helm_clones_github\Taipei-HUG__kubernetes-microservices\examples\helm\ratings`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Taipei-HUG__kubernetes-microservices\examples\helm\ratings`

```text
Error: ratings/templates/tests/test-connection.yaml:6:8
  executing "ratings/templates/tests/test-connection.yaml" at <include "ratings.labels" .>:
    error calling include:
ratings/templates/_helpers.tpl:39:3
  executing "ratings.labels" at <include "ratings.selectorLabels" .>:
    error calling include:
ratings/templates/_helpers.tpl:52:3
  executing "ratings.selectorLabels" at <include "ratings.istio.labels" .>:
    error calling include:
ratings/templates/_helpers.tpl:67:14
  executing "ratings.istio.labels" at <.Values.global.istioEnabled>:
    nil pointer evaluating interface {}.istioEnabled

Use --debug flag to render out invalid YAML
```

### 539. `Taipei-HUG/kubernetes-microservices`

- Chart: `D:\helm_clones_github\Taipei-HUG__kubernetes-microservices\examples\helm\reviews`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Taipei-HUG__kubernetes-microservices\examples\helm\reviews`

```text
Error: reviews/templates/deployment.yaml:16:12
  executing "reviews/templates/deployment.yaml" at <include "reviews.istio.labels" .>:
    error calling include:
reviews/templates/_helpers.tpl:66:14
  executing "reviews.istio.labels" at <.Values.global.istioEnabled>:
    nil pointer evaluating interface {}.istioEnabled

Use --debug flag to render out invalid YAML
```

### 540. `cheyang/arena-spark`

- Chart: `D:\helm_clones_github\cheyang__arena-spark\charts\tfserving`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\cheyang__arena-spark\charts\tfserving`

```text
Error: tensorflow-serving/templates/deployment.yaml:16:34
  executing "tensorflow-serving/templates/deployment.yaml" at <.Release.Time.Seconds>:
    nil pointer evaluating interface {}.Seconds

Use --debug flag to render out invalid YAML
```

### 541. `cheyang/arena-spark`

- Chart: `D:\helm_clones_github\cheyang__arena-spark\charts\trtserving`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\cheyang__arena-spark\charts\trtserving`

```text
Error: tensorrt-serving/templates/deployment.yaml:15:34
  executing "tensorrt-serving/templates/deployment.yaml" at <.Release.Time.Seconds>:
    nil pointer evaluating interface {}.Seconds

Use --debug flag to render out invalid YAML
```

### 542. `kalavai-net/helm-charts`

- Chart: `D:\helm_clones_github\kalavai-net__helm-charts\charts\kalavai-api`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\kalavai-net__helm-charts\charts\kalavai-api`

```text
Error: kalavai-api/templates/secrets.yaml:9:32
  executing "kalavai-api/templates/secrets.yaml" at <.Values.system.jobId>:
    nil pointer evaluating interface {}.jobId

Use --debug flag to render out invalid YAML
```

### 543. `lucidworks/ocp-fusion-helm-charts`

- Chart: `D:\helm_clones_github\lucidworks__ocp-fusion-helm-charts\5.3.4\fusion`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\lucidworks__ocp-fusion-helm-charts\5.3.4\fusion`

```text
level=INFO msg="warning: destination for pulsar.zookeeper.ports.client is a table. Ignoring non-table value (2181)"
Error: fusion/charts/templating/templates/deployment.yaml:61:4
  executing "fusion/charts/templating/templates/deployment.yaml" at <include "fusion.initContainers.checkZk-v2" .>:
    error calling include:
fusion/charts/admin-ui/charts/fusion-common-utils/templates/_helpers.tpl:342:25
  executing "fusion.initContainers.checkZk-v2" at <.Values.securityContext.runAsUser>:
    nil pointer evaluating interface {}.runAsUser

Use --debug flag to render out invalid YAML
```

### 544. `ministryofjustice/cloud-platform-how-out-of-date-are-we`

- Chart: `D:\helm_clones_github\ministryofjustice__cloud-platform-how-out-of-date-are-we\cloud-platform-reports-cronjobs`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\ministryofjustice__cloud-platform-how-out-of-date-are-we\cloud-platform-reports-cronjobs`

```text
Error: cloud-platform-reports-cronjobs/templates/terraform-modules.yaml:14:14
  executing "cloud-platform-reports-cronjobs/templates/terraform-modules.yaml" at <include "cloud-platform-reports-cronjobs.imagePullSecrets" .>:
    error calling include:
cloud-platform-reports-cronjobs/templates/_helpers.tpl:25:20
  executing "cloud-platform-reports-cronjobs.imagePullSecrets" at <.Values.dockerhubCredentials.secretName>:
    nil pointer evaluating interface {}.secretName

Use --debug flag to render out invalid YAML
```

### 545. `ministryofjustice/cloud-platform-how-out-of-date-are-we`

- Chart: `D:\helm_clones_github\ministryofjustice__cloud-platform-how-out-of-date-are-we\cloud-platform-reports-cronjobs`
- Source: `template`
- Values files: `D:\helm_clones_github\ministryofjustice__cloud-platform-how-out-of-date-are-we\cloud-platform-reports-cronjobs\values-dev.yaml`
- Command: `helm template test D:\helm_clones_github\ministryofjustice__cloud-platform-how-out-of-date-are-we\cloud-platform-reports-cronjobs -f D:\helm_clones_github\ministryofjustice__cloud-platform-how-out-of-date-are-we\cloud-platform-reports-cronjobs\values-dev.yaml`

```text
Error: cloud-platform-reports-cronjobs/templates/terraform-modules.yaml:14:14
  executing "cloud-platform-reports-cronjobs/templates/terraform-modules.yaml" at <include "cloud-platform-reports-cronjobs.imagePullSecrets" .>:
    error calling include:
cloud-platform-reports-cronjobs/templates/_helpers.tpl:25:20
  executing "cloud-platform-reports-cronjobs.imagePullSecrets" at <.Values.dockerhubCredentials.secretName>:
    nil pointer evaluating interface {}.secretName

Use --debug flag to render out invalid YAML
```

### 546. `onasunnymorning/domain-os`

- Chart: `D:\helm_clones_github\onasunnymorning__domain-os\deploy\helm\dos`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\onasunnymorning__domain-os\deploy\helm\dos`

```text
Error: dos/templates/sync-worker-deployment.yaml:6:22
  executing "dos/templates/sync-worker-deployment.yaml" at <.Values.syncworker.replicaCount>:
    nil pointer evaluating interface {}.replicaCount

Use --debug flag to render out invalid YAML
```

### 547. `onasunnymorning/domain-os`

- Chart: `D:\helm_clones_github\onasunnymorning__domain-os\deploy\helm\echo`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\onasunnymorning__domain-os\deploy\helm\echo`

```text
Error: streamclient/templates/secrets.yaml:6:22
  executing "streamclient/templates/secrets.yaml" at <.Values.rabbitmq.host>:
    nil pointer evaluating interface {}.host

Use --debug flag to render out invalid YAML
```

### 548. `onasunnymorning/domain-os`

- Chart: `D:\helm_clones_github\onasunnymorning__domain-os\deploy\helm\metabase`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\onasunnymorning__domain-os\deploy\helm\metabase`

```text
Error: metabase-helm-chart/templates/metabase-service.yaml:8:18
  executing "metabase-helm-chart/templates/metabase-service.yaml" at <.Values.metabase.service.type>:
    nil pointer evaluating interface {}.service

Use --debug flag to render out invalid YAML
```

### 549. `onasunnymorning/domain-os`

- Chart: `D:\helm_clones_github\onasunnymorning__domain-os\deploy\helm\traefik`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\onasunnymorning__domain-os\deploy\helm\traefik`

```text
Error: traefik-helm/templates/traefik-service.yaml:9:18
  executing "traefik-helm/templates/traefik-service.yaml" at <.Values.service.type>:
    nil pointer evaluating interface {}.type

Use --debug flag to render out invalid YAML
```

### 550. `onasunnymorning/domain-os`

- Chart: `D:\helm_clones_github\onasunnymorning__domain-os\deploy\helm\v2\admin-api`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\onasunnymorning__domain-os\deploy\helm\v2\admin-api`

```text
Error: admin-api/templates/admin-api-service.yaml:4:18
  executing "admin-api/templates/admin-api-service.yaml" at <.Values.service.name>:
    nil pointer evaluating interface {}.name

Use --debug flag to render out invalid YAML
```

### 551. `onasunnymorning/domain-os`

- Chart: `D:\helm_clones_github\onasunnymorning__domain-os\deploy\helm\v2\jump`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\onasunnymorning__domain-os\deploy\helm\v2\jump`

```text
Error: bastion-host/templates/junp-service.yaml:4:18
  executing "bastion-host/templates/junp-service.yaml" at <.Values.service.name>:
    nil pointer evaluating interface {}.name

Use --debug flag to render out invalid YAML
```

### 552. `onasunnymorning/domain-os`

- Chart: `D:\helm_clones_github\onasunnymorning__domain-os\deploy\helm\v2\metabase`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\onasunnymorning__domain-os\deploy\helm\v2\metabase`

```text
Error: metabase-helm-chart/templates/metabase-service.yaml:8:18
  executing "metabase-helm-chart/templates/metabase-service.yaml" at <.Values.metabase.service.type>:
    nil pointer evaluating interface {}.service

Use --debug flag to render out invalid YAML
```

### 553. `onasunnymorning/domain-os`

- Chart: `D:\helm_clones_github\onasunnymorning__domain-os\deploy\helm\v2\sftp`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\onasunnymorning__domain-os\deploy\helm\v2\sftp`

```text
Error: sftp/templates/serviceaccount.yaml:1:14
  executing "sftp/templates/serviceaccount.yaml" at <.Values.serviceAccount.create>:
    nil pointer evaluating interface {}.create

Use --debug flag to render out invalid YAML
```

### 554. `onasunnymorning/domain-os`

- Chart: `D:\helm_clones_github\onasunnymorning__domain-os\deploy\helm\v2\workers`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\onasunnymorning__domain-os\deploy\helm\v2\workers`

```text
Error: dos-workers/templates/domain-os-worker-deployment.yaml:4:18
  executing "dos-workers/templates/domain-os-worker-deployment.yaml" at <.Values.worker.name>:
    nil pointer evaluating interface {}.name

Use --debug flag to render out invalid YAML
```

### 555. `vlab-research/fly`

- Chart: `D:\helm_clones_github\vlab-research__fly\dashboard-server\chart`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\vlab-research__fly\dashboard-server\chart`

```text
Error: dashboard/templates/service.yaml:7:18
  executing "dashboard/templates/service.yaml" at <.Values.service.type>:
    nil pointer evaluating interface {}.type

Use --debug flag to render out invalid YAML
```

### 556. `GuanceDemo/guance-java-ruoyi-demo`

- Chart: `D:\helm_clones_github\GuanceDemo__guance-java-ruoyi-demo\deployment\helm\charts\middleware`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\GuanceDemo__guance-java-ruoyi-demo\deployment\helm\charts\middleware`

```text
Error: middleware/templates/redis/deploy.yaml:24:27
  executing "middleware/templates/redis/deploy.yaml" at <.Values.global.docker_registry>:
    nil pointer evaluating interface {}.docker_registry

Use --debug flag to render out invalid YAML
```

### 557. `JonasHess/homelab-iac`

- Chart: `D:\helm_clones_github\JonasHess__homelab-iac\apps\argocd`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\JonasHess__homelab-iac\apps\argocd`

```text
Error: argocd/templates/argocd-httproute.yaml:1:18
  executing "argocd/templates/argocd-httproute.yaml" at <$.Values.global.security>:
    nil pointer evaluating interface {}.security

Use --debug flag to render out invalid YAML
```

### 558. `JonasHess/homelab-iac`

- Chart: `D:\helm_clones_github\JonasHess__homelab-iac\apps\crossplane`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\JonasHess__homelab-iac\apps\crossplane`

```text
Error: crossplane/templates/external-secret-aws-creds.yaml:9:20
  executing "crossplane/templates/external-secret-aws-creds.yaml" at <.Values.global.externalSecrets.clusterSecretStoreRef.name>:
    nil pointer evaluating interface {}.externalSecrets

Use --debug flag to render out invalid YAML
```

### 559. `JonasHess/homelab-iac`

- Chart: `D:\helm_clones_github\JonasHess__homelab-iac\apps\mosquitto`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\JonasHess__homelab-iac\apps\mosquitto`

```text
Error: mosquitto/charts/generic/templates/external-secrets.yaml:14:23
  executing "mosquitto/charts/generic/templates/external-secrets.yaml" at <$.Values.global.externalSecrets.refreshInterval>:
    nil pointer evaluating interface {}.refreshInterval

Use --debug flag to render out invalid YAML
```

### 560. `Kapil-Bhalodiya/E-Commerce`

- Chart: `D:\helm_clones_github\Kapil-Bhalodiya__E-Commerce\infra\addons\backend`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Kapil-Bhalodiya__E-Commerce\infra\addons\backend`

```text
Error: backend/templates/ingress.yaml:1:14
  executing "backend/templates/ingress.yaml" at <.Values.ingress.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 561. `Kapil-Bhalodiya/E-Commerce`

- Chart: `D:\helm_clones_github\Kapil-Bhalodiya__E-Commerce\infra\addons\frontend`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Kapil-Bhalodiya__E-Commerce\infra\addons\frontend`

```text
Error: frontend/templates/ingress.yaml:1:14
  executing "frontend/templates/ingress.yaml" at <.Values.ingress.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 562. `Kapil-Bhalodiya/E-Commerce`

- Chart: `D:\helm_clones_github\Kapil-Bhalodiya__E-Commerce\infra\addons\gateway-api`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Kapil-Bhalodiya__E-Commerce\infra\addons\gateway-api`

```text
Error: gateway-api/templates/httproute-backend.yaml:29:23
  executing "gateway-api/templates/httproute-backend.yaml" at <.service.name>:
    nil pointer evaluating interface {}.name

Use --debug flag to render out invalid YAML
```

### 563. `LasseRapo/fabric-cti-sharing`

- Chart: `D:\helm_clones_github\LasseRapo__fabric-cti-sharing\bevel-operator-fabric\charts\hlf-ca`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\LasseRapo__fabric-cti-sharing\bevel-operator-fabric\charts\hlf-ca`

```text
Error: hlf-ca/templates/traefikroute.yaml:1:13
  executing "hlf-ca/templates/traefikroute.yaml" at <.Values.traefik.hosts>:
    nil pointer evaluating interface {}.hosts

Use --debug flag to render out invalid YAML
```

### 564. `LasseRapo/fabric-cti-sharing`

- Chart: `D:\helm_clones_github\LasseRapo__fabric-cti-sharing\bevel-operator-fabric\charts\hlf-ordnode`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\LasseRapo__fabric-cti-sharing\bevel-operator-fabric\charts\hlf-ordnode`

```text
Error: hlf-ordnode/templates/traefikroute.yaml:1:13
  executing "hlf-ordnode/templates/traefikroute.yaml" at <.Values.traefik.hosts>:
    nil pointer evaluating interface {}.hosts

Use --debug flag to render out invalid YAML
```

### 565. `MrE-Fog/ks-installer2`

- Chart: `D:\helm_clones_github\MrE-Fog__ks-installer2\roles\common\files\openldap-ha`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\MrE-Fog__ks-installer2\roles\common\files\openldap-ha`

```text
Error: openldap-ha/templates/statefulset.yaml:144:18
  executing "openldap-ha/templates/statefulset.yaml" at <.Values.persistence.storageClass>:
    nil pointer evaluating interface {}.storageClass

Use --debug flag to render out invalid YAML
```

### 566. `MrE-Fog/ks-installer2`

- Chart: `D:\helm_clones_github\MrE-Fog__ks-installer2\roles\edgeruntime\files\kubeedge\cloudcore`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\MrE-Fog__ks-installer2\roles\edgeruntime\files\kubeedge\cloudcore`

```text
Error: cloudcore/templates/service_edgeservice.yaml:4:18
  executing "cloudcore/templates/service_edgeservice.yaml" at <.Values.edgeService.labels>:
    nil pointer evaluating interface {}.labels

Use --debug flag to render out invalid YAML
```

### 567. `MrE-Fog/ks-installer2`

- Chart: `D:\helm_clones_github\MrE-Fog__ks-installer2\roles\ks-auditing\files\kube-auditing`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\MrE-Fog__ks-installer2\roles\ks-auditing\files\kube-auditing`

```text
Error: kube-auditing/templates/webhook.yaml:6:16
  executing "kube-auditing/templates/webhook.yaml" at <.Values.webhook.replicas>:
    nil pointer evaluating interface {}.replicas

Use --debug flag to render out invalid YAML
```

### 568. `MrE-Fog/ks-installer2`

- Chart: `D:\helm_clones_github\MrE-Fog__ks-installer2\roles\ks-monitor\files\notification-manager`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\MrE-Fog__ks-installer2\roles\ks-monitor\files\notification-manager`

```text
Error: notification-manager/templates/operator.yaml:23:25
  executing "notification-manager/templates/operator.yaml" at <.Values.operator.containers.proxy.image.repo>:
    nil pointer evaluating interface {}.containers

Use --debug flag to render out invalid YAML
```

### 569. `PilotDataPlatform/helm-charts`

- Chart: `D:\helm_clones_github\PilotDataPlatform__helm-charts\ai-chat-service`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\PilotDataPlatform__helm-charts\ai-chat-service`

```text
Error: ai-chat-service/templates/deployment.yaml:51:39
  executing "ai-chat-service/templates/deployment.yaml" at <.Values.container.port>:
    nil pointer evaluating interface {}.port

Use --debug flag to render out invalid YAML
```

### 570. `cisco-open/appdynamics-k8s-webhook-instrumentor`

- Chart: `D:\helm_clones_github\cisco-open__appdynamics-k8s-webhook-instrumentor\buildEnv\helm\webhook-instrumentor`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\cisco-open__appdynamics-k8s-webhook-instrumentor\buildEnv\helm\webhook-instrumentor`

```text
Error: webhook-instrumentor/templates/d-webhook-instrumentor.yaml:31:25
  executing "webhook-instrumentor/templates/d-webhook-instrumentor.yaml" at <.Values.image.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 571. `cisco-open/appdynamics-k8s-webhook-instrumentor`

- Chart: `D:\helm_clones_github\cisco-open__appdynamics-k8s-webhook-instrumentor\webhook\helm\instrumentor`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\cisco-open__appdynamics-k8s-webhook-instrumentor\webhook\helm\instrumentor`

```text
Error: webhook-instrumentor/templates/d-webhook-instrumentor.yaml:31:25
  executing "webhook-instrumentor/templates/d-webhook-instrumentor.yaml" at <.Values.image.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 572. `codefuturist/helm-charts`

- Chart: `D:\helm_clones_github\codefuturist__helm-charts\charts\apps\actualbudget`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\codefuturist__helm-charts\charts\apps\actualbudget`

```text
Error: actualbudget/templates/statefulset.yaml:1:17
  executing "actualbudget/templates/statefulset.yaml" at <.Values.controller.type>:
    nil pointer evaluating interface {}.type

Use --debug flag to render out invalid YAML
```

### 573. `codefuturist/helm-charts`

- Chart: `D:\helm_clones_github\codefuturist__helm-charts\charts\apps\bitwarden-eso-provider`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\codefuturist__helm-charts\charts\apps\bitwarden-eso-provider`

```text
Error: bitwarden-eso-provider/templates/serviceaccount.yaml:1:14
  executing "bitwarden-eso-provider/templates/serviceaccount.yaml" at <.Values.serviceAccount.create>:
    nil pointer evaluating interface {}.create

Use --debug flag to render out invalid YAML
```

### 574. `codefuturist/helm-charts`

- Chart: `D:\helm_clones_github\codefuturist__helm-charts\charts\apps\homarr`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\codefuturist__helm-charts\charts\apps\homarr`

```text
Error: homarr/templates/deployment.yaml:222:22
  executing "homarr/templates/deployment.yaml" at <.Values.deployment.livenessProbe.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 575. `codefuturist/helm-charts`

- Chart: `D:\helm_clones_github\codefuturist__helm-charts\charts\apps\home-assistant`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\codefuturist__helm-charts\charts\apps\home-assistant`

```text
Error: home-assistant/templates/statefulset.yaml:1:4
  executing "home-assistant/templates/statefulset.yaml" at <include "home-assistant.validateController" .>:
    error calling include:
home-assistant/templates/_helpers.tpl:130:23
  executing "home-assistant.validateController" at <.Values.controller.type>:
    nil pointer evaluating interface {}.type

Use --debug flag to render out invalid YAML
```

### 576. `codefuturist/helm-charts`

- Chart: `D:\helm_clones_github\codefuturist__helm-charts\charts\apps\metube`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\codefuturist__helm-charts\charts\apps\metube`

```text
Error: metube/templates/statefulset.yaml:1:17
  executing "metube/templates/statefulset.yaml" at <.Values.controller.type>:
    nil pointer evaluating interface {}.type

Use --debug flag to render out invalid YAML
```

### 577. `codefuturist/helm-charts`

- Chart: `D:\helm_clones_github\codefuturist__helm-charts\charts\apps\paperless-ngx`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\codefuturist__helm-charts\charts\apps\paperless-ngx`

```text
Error: paperless-ngx/templates/worker-deployment.yaml:1:14
  executing "paperless-ngx/templates/worker-deployment.yaml" at <.Values.worker.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 578. `codefuturist/helm-charts`

- Chart: `D:\helm_clones_github\codefuturist__helm-charts\charts\apps\shlink`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\codefuturist__helm-charts\charts\apps\shlink`

```text
Error: shlink/templates/statefulset.yaml:1:4
  executing "shlink/templates/statefulset.yaml" at <include "shlink.validateValues" .>:
    error calling include:
shlink/templates/_helpers.tpl:153:23
  executing "shlink.validateValues" at <.Values.database.existingSecret>:
    nil pointer evaluating interface {}.existingSecret

Use --debug flag to render out invalid YAML
```

### 579. `stfc/cloud-helm-charts`

- Chart: `D:\helm_clones_github\stfc__cloud-helm-charts\old\stfc-cloud-chatops`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\stfc__cloud-helm-charts\old\stfc-cloud-chatops`

```text
Error: stfc-cloud-chatops/templates/secrets.yaml:9:31
  executing "stfc-cloud-chatops/templates/secrets.yaml" at <.Values.secrets.slackBotToken>:
    nil pointer evaluating interface {}.slackBotToken

Use --debug flag to render out invalid YAML
```

### 580. `AlexanderBabel/helm-charts`

- Chart: `D:\helm_clones_github\AlexanderBabel__helm-charts\charts\matrix`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\AlexanderBabel__helm-charts\charts\matrix`

```text
Error: matrix/templates/synapse/deployment.yaml:20:36
  executing "matrix/templates/synapse/deployment.yaml" at <include (print $.Template.BasePath "/synapse/configmap.yaml") .>:
    error calling include:
matrix/templates/synapse/configmap.yaml:10:7
  executing "matrix/templates/synapse/configmap.yaml" at <include "homeserver.yaml" .>:
    error calling include:
matrix/templates/synapse/_homeserver.yaml:421:16
  executing "homeserver.yaml" at <.Values.ldap.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 581. `Bahmni/helm-umbrella-chart`

- Chart: `D:\helm_clones_github\Bahmni__helm-umbrella-chart`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Bahmni__helm-umbrella-chart`

```text
Error: bahmni/templates/ingress.yaml:6:27
  executing "bahmni/templates/ingress.yaml" at <.Values.metadata.labels.environment>:
    nil pointer evaluating interface {}.labels

Use --debug flag to render out invalid YAML
```

### 582. `Clark1992/ECK1`

- Chart: `D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\analytics\clickhouse`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\analytics\clickhouse`

```text
Error: clickhouse/templates/stateful-set.yaml:60:25
  executing "clickhouse/templates/stateful-set.yaml" at <$.Values.clickhouse.persistence.size>:
    nil pointer evaluating interface {}.size

Use --debug flag to render out invalid YAML
```

### 583. `Clark1992/ECK1`

- Chart: `D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\elasticsearch\cert`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\elasticsearch\cert`

```text
Error: elasticsearch-certs/templates/renew-cert-job-once.yaml:17:27
  executing "elasticsearch-certs/templates/renew-cert-job-once.yaml" at <.Values.elasticsearch.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 584. `Clark1992/ECK1`

- Chart: `D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\elasticsearch\cluster`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\elasticsearch\cluster`

```text
Error: elasticsearch/templates/statefulset-master.yaml:4:18
  executing "elasticsearch/templates/statefulset-master.yaml" at <.Values.elasticsearch.resourcePrefix>:
    nil pointer evaluating interface {}.resourcePrefix

Use --debug flag to render out invalid YAML
```

### 585. `Clark1992/ECK1`

- Chart: `D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\elasticsearch\cluster`
- Source: `template`
- Values files: `D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\elasticsearch\cluster\values.common.yaml`
- Command: `helm template test D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\elasticsearch\cluster -f D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\elasticsearch\cluster\values.common.yaml`

```text
Error: elasticsearch/templates/statefulset-master.yaml:4:18
  executing "elasticsearch/templates/statefulset-master.yaml" at <.Values.elasticsearch.resourcePrefix>:
    nil pointer evaluating interface {}.resourcePrefix

Use --debug flag to render out invalid YAML
```

### 586. `Clark1992/ECK1`

- Chart: `D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\elasticsearch\cluster`
- Source: `template`
- Values files: `D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\elasticsearch\cluster\values.local.yaml`
- Command: `helm template test D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\elasticsearch\cluster -f D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\elasticsearch\cluster\values.local.yaml`

```text
Error: elasticsearch/templates/statefulset-master.yaml:4:18
  executing "elasticsearch/templates/statefulset-master.yaml" at <.Values.elasticsearch.resourcePrefix>:
    nil pointer evaluating interface {}.resourcePrefix

Use --debug flag to render out invalid YAML
```

### 587. `Clark1992/ECK1`

- Chart: `D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\elasticsearch\cluster`
- Source: `template`
- Values files: `D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\elasticsearch\cluster\values.common.yaml`, `D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\elasticsearch\cluster\values.local.yaml`
- Command: `helm template test D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\elasticsearch\cluster -f D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\elasticsearch\cluster\values.common.yaml -f D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\elasticsearch\cluster\values.local.yaml`

```text
Error: elasticsearch/templates/statefulset-master.yaml:4:18
  executing "elasticsearch/templates/statefulset-master.yaml" at <.Values.elasticsearch.resourcePrefix>:
    nil pointer evaluating interface {}.resourcePrefix

Use --debug flag to render out invalid YAML
```

### 588. `Clark1992/ECK1`

- Chart: `D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\elasticsearch\external-access`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\elasticsearch\external-access`

```text
Error: elasticsearch-external-access/templates/ingress.yaml:1:14
  executing "elasticsearch-external-access/templates/ingress.yaml" at <.Values.ingress.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 589. `Clark1992/ECK1`

- Chart: `D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\elasticsearch\external-access`
- Source: `template`
- Values files: `D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\elasticsearch\external-access\values.local.yaml`
- Command: `helm template test D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\elasticsearch\external-access -f D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\elasticsearch\external-access\values.local.yaml`

```text
Error: elasticsearch-external-access/templates/ingress.yaml:28:32
  executing "elasticsearch-external-access/templates/ingress.yaml" at <.Values.elasticsearch.resourcePrefix>:
    nil pointer evaluating interface {}.resourcePrefix

Use --debug flag to render out invalid YAML
```

### 590. `Clark1992/ECK1`

- Chart: `D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\elasticsearch\internal-access`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\elasticsearch\internal-access`

```text
Error: elasticsearch-internal-access/templates/service.yaml:4:18
  executing "elasticsearch-internal-access/templates/service.yaml" at <.Values.elasticsearch.resourcePrefix>:
    nil pointer evaluating interface {}.resourcePrefix

Use --debug flag to render out invalid YAML
```

### 591. `Clark1992/ECK1`

- Chart: `D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\elasticsearch\internal-access`
- Source: `template`
- Values files: `D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\elasticsearch\internal-access\values.local.yaml`
- Command: `helm template test D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\elasticsearch\internal-access -f D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\elasticsearch\internal-access\values.local.yaml`

```text
Error: elasticsearch-internal-access/templates/service.yaml:4:18
  executing "elasticsearch-internal-access/templates/service.yaml" at <.Values.elasticsearch.resourcePrefix>:
    nil pointer evaluating interface {}.resourcePrefix

Use --debug flag to render out invalid YAML
```

### 592. `Clark1992/ECK1`

- Chart: `D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\global-vars`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\global-vars`

```text
Error: global-config/templates/configmap.yaml:13:48
  executing "global-config/templates/configmap.yaml" at <.Values.observability.otlpService>:
    nil pointer evaluating interface {}.otlpService

Use --debug flag to render out invalid YAML
```

### 593. `Clark1992/ECK1`

- Chart: `D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\global-vars`
- Source: `template`
- Values files: `D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\global-vars\values.local.yaml`
- Command: `helm template test D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\global-vars -f D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\global-vars\values.local.yaml`

```text
Error: global-config/templates/configmap.yaml:13:48
  executing "global-config/templates/configmap.yaml" at <.Values.observability.otlpService>:
    nil pointer evaluating interface {}.otlpService

Use --debug flag to render out invalid YAML
```

### 594. `Clark1992/ECK1`

- Chart: `D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\kafka\cluster`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\kafka\cluster`

```text
Error: strimzi-kafka/templates/user.yaml:4:18
  executing "strimzi-kafka/templates/user.yaml" at <.Values.kafka.userSecretName>:
    nil pointer evaluating interface {}.userSecretName

Use --debug flag to render out invalid YAML
```

### 595. `Clark1992/ECK1`

- Chart: `D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\kafka\schema-registry\apicurio`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\kafka\schema-registry\apicurio`

```text
Error: apicurio-registry/templates/service-topics.yaml:8:34
  executing "apicurio-registry/templates/service-topics.yaml" at <.Values.storage.clusterName>:
    nil pointer evaluating interface {}.clusterName

Use --debug flag to render out invalid YAML
```

### 596. `Clark1992/ECK1`

- Chart: `D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\kafka\schema-registry\apicurio`
- Source: `template`
- Values files: `D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\kafka\schema-registry\apicurio\values.local.yaml`
- Command: `helm template test D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\kafka\schema-registry\apicurio -f D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\kafka\schema-registry\apicurio\values.local.yaml`

```text
Error: apicurio-registry/templates/registry-cr.yaml:13:34
  executing "apicurio-registry/templates/registry-cr.yaml" at <.Values.kafka.bootstrapServers>:
    nil pointer evaluating interface {}.bootstrapServers

Use --debug flag to render out invalid YAML
```

### 597. `Clark1992/ECK1`

- Chart: `D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\metrics-stack\metrics-access`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\metrics-stack\metrics-access`

```text
Error: metrics-access/templates/clusterrolebinding.yaml:10:20
  executing "metrics-access/templates/clusterrolebinding.yaml" at <.Values.subject.kind>:
    nil pointer evaluating interface {}.kind

Use --debug flag to render out invalid YAML
```

### 598. `Clark1992/ECK1`

- Chart: `D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\namespace`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\namespace`

```text
Error: namespaces/templates/namespace.yaml:8:19
  executing "namespaces/templates/namespace.yaml" at <.Values.elasticsearch.caSecretName>:
    nil pointer evaluating interface {}.caSecretName

Use --debug flag to render out invalid YAML
```

### 599. `Clark1992/ECK1`

- Chart: `D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\observability\grafana\ingress`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\observability\grafana\ingress`

```text
Error: grafana-ingress/templates/ingress.yaml:1:14
  executing "grafana-ingress/templates/ingress.yaml" at <.Values.ingress.enabled>:
    nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 600. `Clark1992/ECK1`

- Chart: `D:\helm_clones_github\Clark1992__ECK1\src\ECK1.CommandsAPI\Deploy\integration-manifests`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Clark1992__ECK1\src\ECK1.CommandsAPI\Deploy\integration-manifests`

```text
Error: eck1-commands-integration-manifests/templates/elasticsearch/es-mappings-job.yaml:16:27
  executing "eck1-commands-integration-manifests/templates/elasticsearch/es-mappings-job.yaml" at <.Values.toolbox.image>:
    nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 601. `Clark1992/ECK1`

- Chart: `D:\helm_clones_github\Clark1992__ECK1\src\ECK1.CommandsAPI\Deploy\integration-manifests`
- Source: `template`
- Values files: `D:\helm_clones_github\Clark1992__ECK1\src\ECK1.CommandsAPI\Deploy\integration-manifests\values.local.yaml`
- Command: `helm template test D:\helm_clones_github\Clark1992__ECK1\src\ECK1.CommandsAPI\Deploy\integration-manifests -f D:\helm_clones_github\Clark1992__ECK1\src\ECK1.CommandsAPI\Deploy\integration-manifests\values.local.yaml`

```text
Error: eck1-commands-integration-manifests/templates/elasticsearch/es-mappings-job.yaml:23:31
  executing "eck1-commands-integration-manifests/templates/elasticsearch/es-mappings-job.yaml" at <.Values.elasticsearch.clusterUrl>:
    nil pointer evaluating interface {}.clusterUrl

Use --debug flag to render out invalid YAML
```

### 602. `Clark1992/ECK1`

- Chart: `D:\helm_clones_github\Clark1992__ECK1\src\ECK1.CommandsAPI\Deploy\service`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Clark1992__ECK1\src\ECK1.CommandsAPI\Deploy\service`

```text
Error: eck1-commandsapi/templates/deployment.yaml:80:32
  executing "eck1-commandsapi/templates/deployment.yaml" at <.Values.redis.host>:
    nil pointer evaluating interface {}.host

Use --debug flag to render out invalid YAML
```

### 603. `Clark1992/ECK1`

- Chart: `D:\helm_clones_github\Clark1992__ECK1\src\ECK1.CommandsAPI\Deploy\service`
- Source: `template`
- Values files: `D:\helm_clones_github\Clark1992__ECK1\src\ECK1.CommandsAPI\Deploy\service\values.local.yaml`
- Command: `helm template test D:\helm_clones_github\Clark1992__ECK1\src\ECK1.CommandsAPI\Deploy\service -f D:\helm_clones_github\Clark1992__ECK1\src\ECK1.CommandsAPI\Deploy\service\values.local.yaml`

```text
Error: eck1-commandsapi/templates/deployment.yaml:80:32
  executing "eck1-commandsapi/templates/deployment.yaml" at <.Values.redis.host>:
    nil pointer evaluating interface {}.host

Use --debug flag to render out invalid YAML
```

### 604. `Clark1992/ECK1`

- Chart: `D:\helm_clones_github\Clark1992__ECK1\src\ECK1.FE\Deploy`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Clark1992__ECK1\src\ECK1.FE\Deploy`

```text
Error: eck1-fe/templates/service.yaml:8:18
  executing "eck1-fe/templates/service.yaml" at <.Values.service.type>:
    nil pointer evaluating interface {}.type

Use --debug flag to render out invalid YAML
```

### 605. `Clark1992/ECK1`

- Chart: `D:\helm_clones_github\Clark1992__ECK1\src\ECK1.Gateway\Deploy`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Clark1992__ECK1\src\ECK1.Gateway\Deploy`

```text
Error: eck1-gateway/templates/service.yaml:8:18
  executing "eck1-gateway/templates/service.yaml" at <.Values.service.type>:
    nil pointer evaluating interface {}.type

Use --debug flag to render out invalid YAML
```

## `unknown.unclassified` (189)

### 1. `rancher/rancher`

- Chart: `D:\helm_clones_github\rancher__rancher\chart`
- Source: `dependency`

```text
Error: cannot load Chart.yaml: error converting YAML to JSON: yaml: line 4: found character that cannot start any token
```

### 2. `WeBankFinTech/Prophecis`

- Chart: `D:\helm_clones_github\WeBankFinTech__Prophecis\install\Prophecis`
- Source: `dependency`

```text
Error: cannot load values.yaml: error reading yaml document: invalid Yaml document separator: --END RSA PRIVATE KEY-----"
```

### 3. `IBM/charts`

- Chart: `D:\helm_clones_github\IBM__charts\community\aqua-enforcer`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\IBM__charts\community\aqua-enforcer`

```text
Error: execution error at (aqua-enforcer/templates/enforcer-token-secret.yaml:14:13): A valid .Values.enforcerToken entry required!

Use --debug flag to render out invalid YAML
```

### 4. `IBM/charts`

- Chart: `D:\helm_clones_github\IBM__charts\community\aqua-scanner`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\IBM__charts\community\aqua-scanner`

```text
Error: execution error at (aqua-scanner/templates/scanner-deployment.yaml:31:14): Please specify a username associated with the Scanner role!

Use --debug flag to render out invalid YAML
```

### 5. `open-edge-platform/edge-ai-libraries`

- Chart: `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart`
- Source: `template`
- Values files: `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\values_ovms.yaml`
- Command: `helm template test D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart -f D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\values_ovms.yaml`

```text
Error: failed to parse D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\values_ovms.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 6. `open-edge-platform/edge-ai-libraries`

- Chart: `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart`
- Source: `template`
- Values files: `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\values_ovms.yaml`, `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\values_tgi.yaml`
- Command: `helm template test D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart -f D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\values_ovms.yaml -f D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\values_tgi.yaml`

```text
Error: failed to parse D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\values_ovms.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 7. `open-edge-platform/edge-ai-libraries`

- Chart: `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart`
- Source: `template`
- Values files: `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\values_ovms.yaml`, `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\values_vllm.yaml`
- Command: `helm template test D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart -f D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\values_ovms.yaml -f D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\values_vllm.yaml`

```text
Error: failed to parse D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\values_ovms.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 8. `open-edge-platform/edge-ai-libraries`

- Chart: `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart`
- Source: `template`
- Values files: `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\values_ovms.yaml`, `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\values_tgi.yaml`, `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\values_vllm.yaml`
- Command: `helm template test D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart -f D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\values_ovms.yaml -f D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\values_tgi.yaml -f D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\values_vllm.yaml`

```text
Error: failed to parse D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\values_ovms.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 9. `aws-samples/amazon-eks-machine-learning-with-terraform-and-kubeflow`

- Chart: `D:\helm_clones_github\aws-samples__amazon-eks-machine-learning-with-terraform-and-kubeflow\charts\machine-learning\testing\maskrcnn-jupyter`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\aws-samples__amazon-eks-machine-learning-with-terraform-and-kubeflow\charts\machine-learning\testing\maskrcnn-jupyter`

```text
Error: execution error at (jupyter/templates/jupyter.yaml:27:7): .Values.global.source_cidr required!

Use --debug flag to render out invalid YAML
```

### 10. `aws-samples/amazon-eks-machine-learning-with-terraform-and-kubeflow`

- Chart: `D:\helm_clones_github\aws-samples__amazon-eks-machine-learning-with-terraform-and-kubeflow\charts\machine-learning\testing\maskrcnn-optimized-jupyter`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\aws-samples__amazon-eks-machine-learning-with-terraform-and-kubeflow\charts\machine-learning\testing\maskrcnn-optimized-jupyter`

```text
Error: execution error at (jupyter/templates/jupyter.yaml:27:7): .Values.global.source_cidr required!

Use --debug flag to render out invalid YAML
```

### 11. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\genai-gateway`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\genai-gateway`

```text
Error: execution error at (genaigateway/charts/redis/templates/NOTES.txt:216:4): 

⚠ ERROR: Original containers have been substituted for unrecognized ones. Deploying this chart with non-standard containers is likely to cause degraded security and performance, broken chart features, and missing environment variables.

Unrecognized images:
  - docker.io/bitnamilegacy/redis:8.0.1-debian-12-r1

If you are sure you want to proceed with non-standard containers, you can skip container image verification by setting the global parameter 'global.security.allowInsecureImages' to true.
Further information can be obtained at https://github.com/bitnami/charts/issues/30850

Use --debug flag to render out invalid YAML
```

### 12. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tei`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tei\ci-gaudi-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tei -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tei\ci-gaudi-values.yaml`

```text
Error: failed to parse D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tei\ci-gaudi-values.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 13. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tei`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tei\ci-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tei -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tei\ci-values.yaml`

```text
Error: failed to parse D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tei\ci-values.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 14. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tei`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tei\ci-gaudi-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tei\ci-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tei -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tei\ci-gaudi-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tei\ci-values.yaml`

```text
Error: failed to parse D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tei\ci-gaudi-values.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 15. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tei`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tei\ci-gaudi-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tei\gaudi-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tei -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tei\ci-gaudi-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tei\gaudi-values.yaml`

```text
Error: failed to parse D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tei\ci-gaudi-values.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 16. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tei`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tei\ci-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tei\gaudi-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tei -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tei\ci-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tei\gaudi-values.yaml`

```text
Error: failed to parse D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tei\ci-values.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 17. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tei`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tei\ci-gaudi-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tei\ci-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tei\gaudi-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tei -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tei\ci-gaudi-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tei\ci-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tei\gaudi-values.yaml`

```text
Error: failed to parse D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tei\ci-gaudi-values.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 18. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\teirerank`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\teirerank\ci-gaudi-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\teirerank -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\teirerank\ci-gaudi-values.yaml`

```text
Error: failed to parse D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\teirerank\ci-gaudi-values.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 19. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\teirerank`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\teirerank\ci-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\teirerank -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\teirerank\ci-values.yaml`

```text
Error: failed to parse D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\teirerank\ci-values.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 20. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\teirerank`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\teirerank\ci-gaudi-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\teirerank\ci-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\teirerank -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\teirerank\ci-gaudi-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\teirerank\ci-values.yaml`

```text
Error: failed to parse D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\teirerank\ci-gaudi-values.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 21. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\teirerank`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\teirerank\ci-gaudi-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\teirerank\gaudi-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\teirerank -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\teirerank\ci-gaudi-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\teirerank\gaudi-values.yaml`

```text
Error: failed to parse D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\teirerank\ci-gaudi-values.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 22. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\teirerank`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\teirerank\ci-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\teirerank\gaudi-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\teirerank -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\teirerank\ci-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\teirerank\gaudi-values.yaml`

```text
Error: failed to parse D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\teirerank\ci-values.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 23. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\teirerank`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\teirerank\ci-gaudi-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\teirerank\ci-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\teirerank\gaudi-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\teirerank -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\teirerank\ci-gaudi-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\teirerank\ci-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\teirerank\gaudi-values.yaml`

```text
Error: failed to parse D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\teirerank\ci-gaudi-values.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 24. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tgi`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tgi\ci-gaudi-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tgi -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tgi\ci-gaudi-values.yaml`

```text
Error: failed to parse D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tgi\ci-gaudi-values.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 25. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tgi`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tgi\ci-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tgi -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tgi\ci-values.yaml`

```text
Error: failed to parse D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tgi\ci-values.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 26. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tgi`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tgi\ci-gaudi-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tgi\ci-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tgi -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tgi\ci-gaudi-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tgi\ci-values.yaml`

```text
Error: failed to parse D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tgi\ci-gaudi-values.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 27. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tgi`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tgi\ci-gaudi-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tgi\gaudi-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tgi -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tgi\ci-gaudi-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tgi\gaudi-values.yaml`

```text
Error: failed to parse D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tgi\ci-gaudi-values.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 28. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tgi`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tgi\ci-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tgi\gaudi-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tgi -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tgi\ci-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tgi\gaudi-values.yaml`

```text
Error: failed to parse D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tgi\ci-values.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 29. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tgi`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tgi\ci-gaudi-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tgi\ci-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tgi\gaudi-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tgi -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tgi\ci-gaudi-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tgi\ci-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tgi\gaudi-values.yaml`

```text
Error: failed to parse D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\tgi\ci-gaudi-values.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 30. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml`

```text
Error: failed to parse D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 31. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-values.yaml`

```text
Error: failed to parse D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-values.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 32. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-values.yaml`

```text
Error: failed to parse D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 33. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi-values.yaml`

```text
Error: failed to parse D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 34. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi-values.yaml`

```text
Error: failed to parse D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-values.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 35. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi-values.yaml`

```text
Error: failed to parse D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 36. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi3-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi3-values.yaml`

```text
Error: failed to parse D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 37. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi3-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi3-values.yaml`

```text
Error: failed to parse D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-values.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 38. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi3-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi3-values.yaml`

```text
Error: failed to parse D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 39. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi3-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi3-values.yaml`

```text
Error: failed to parse D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 40. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi3-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi3-values.yaml`

```text
Error: failed to parse D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-values.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 41. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi3-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi3-values.yaml`

```text
Error: failed to parse D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 42. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\xeon-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\xeon-values.yaml`

```text
Error: failed to parse D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 43. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\xeon-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\xeon-values.yaml`

```text
Error: failed to parse D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-values.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 44. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\xeon-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\xeon-values.yaml`

```text
Error: failed to parse D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 45. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\xeon-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\xeon-values.yaml`

```text
Error: failed to parse D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 46. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\xeon-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\xeon-values.yaml`

```text
Error: failed to parse D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-values.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 47. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\xeon-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\xeon-values.yaml`

```text
Error: failed to parse D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 48. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi3-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\xeon-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi3-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\xeon-values.yaml`

```text
Error: failed to parse D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 49. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi3-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\xeon-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi3-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\xeon-values.yaml`

```text
Error: failed to parse D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-values.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 50. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi3-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\xeon-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi3-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\xeon-values.yaml`

```text
Error: failed to parse D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 51. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi3-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\xeon-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi3-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\xeon-values.yaml`

```text
Error: failed to parse D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 52. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi3-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\xeon-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi3-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\xeon-values.yaml`

```text
Error: failed to parse D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-values.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 53. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi3-values.yaml`, `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\xeon-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\gaudi3-values.yaml -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\xeon-values.yaml`

```text
Error: failed to parse D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm\ci-gaudi-values.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}
```

### 54. `opendevstack/ods-quickstarters`

- Chart: `D:\helm_clones_github\opendevstack__ods-quickstarters\be-rust-axum\rust-template\chart`
- Source: `dependency`

```text
Error: cannot load Chart.yaml: error converting YAML to JSON: yaml: invalid map key: map[interface {}]interface {}{"project-name":interface {}(nil)}
```

### 55. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-21st-dev-magic\charts\mcp-server-21st-dev-magic`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-21st-dev-magic\charts\mcp-server-21st-dev-magic`

```text
Error: execution error at (mcp-server-21st-dev-magic/templates/secrets.yaml:10:9): required value for secrets.TWENTY_FIRST_API_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 56. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-adfin\charts\mcp-server-adfin`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-adfin\charts\mcp-server-adfin`

```text
Error: execution error at (mcp-server-adfin/templates/secrets.yaml:10:9): required value for secrets.ADFIN_PASSWORD either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 57. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-agentql\charts\mcp-server-agentql`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-agentql\charts\mcp-server-agentql`

```text
Error: execution error at (mcp-server-agentql/templates/secrets.yaml:10:9): required value for secrets.AGENTQL_API_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 58. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-agentrpc\charts\mcp-server-agentrpc`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-agentrpc\charts\mcp-server-agentrpc`

```text
Error: execution error at (mcp-server-agentrpc/templates/secrets.yaml:10:9): required value for secrets.AGENTRPC_API_SECRET either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 59. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-aiven\charts\mcp-server-aiven`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-aiven\charts\mcp-server-aiven`

```text
Error: execution error at (mcp-server-aiven/templates/secrets.yaml:10:9): required value for secrets.AIVEN_TOKEN either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 60. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-alchemy\charts\mcp-server-alchemy`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-alchemy\charts\mcp-server-alchemy`

```text
Error: execution error at (mcp-server-alchemy/templates/secrets.yaml:10:9): required value for secrets.ALCHEMY_API_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 61. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-algolia\charts\mcp-server-algolia`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-algolia\charts\mcp-server-algolia`

```text
Error: execution error at (mcp-server-algolia/templates/secrets.yaml:10:9): required value for secrets.ALGOLIA_CREDENTIALS either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 62. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-alibaba-cloud-ops\charts\mcp-server-alibaba-cloud-ops`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-alibaba-cloud-ops\charts\mcp-server-alibaba-cloud-ops`

```text
Error: execution error at (mcp-server-alibaba-cloud-ops/templates/secrets.yaml:10:9): required value for secrets.ALIBABA_CLOUD_ACCESS_KEY_ID either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 63. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-alibabacloud-adb-mysql\charts\mcp-server-alibabacloud-adb-mysql`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-alibabacloud-adb-mysql\charts\mcp-server-alibabacloud-adb-mysql`

```text
Error: execution error at (mcp-server-alibabacloud-adb-mysql/templates/secrets.yaml:10:9): required value for secrets.ADB_MYSQL_PASSWORD either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 64. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-alibabacloud-hologres\charts\mcp-server-alibaba-hologres`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-alibabacloud-hologres\charts\mcp-server-alibaba-hologres`

```text
Error: execution error at (mcp-server-alibaba-hologres/templates/secrets.yaml:10:9): required value for secrets.HOLOGRES_PASSWORD either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 65. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-alibabacloud-hologres\charts\mcp-server-alibabacloud-hologres`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-alibabacloud-hologres\charts\mcp-server-alibabacloud-hologres`

```text
Error: execution error at (mcp-server-alibabacloud-hologres/templates/secrets.yaml:10:9): required value for secrets.HOLOGRES_PASSWORD either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 66. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-alibabacloud-opensearch-ai-search\charts\mcp-server-alibabacloud-opensearch-ai-search`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-alibabacloud-opensearch-ai-search\charts\mcp-server-alibabacloud-opensearch-ai-search`

```text
Error: execution error at (mcp-server-alibabacloud-opensearch-ai-search/templates/secrets.yaml:10:9): required value for secrets.AISEARCH_API_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 67. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-alibabacloud-opensearch-vector-search\charts\mcp-server-alibabacloud-opensearch-vector-search`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-alibabacloud-opensearch-vector-search\charts\mcp-server-alibabacloud-opensearch-vector-search`

```text
Error: execution error at (mcp-server-alibabacloud-opensearch-vector-search/templates/secrets.yaml:10:9): required value for secrets.OPENSEARCH_VECTOR_PASSWORD either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 68. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-alibabacloud-ops\charts\mcp-server-alibaba-cloud-ops`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-alibabacloud-ops\charts\mcp-server-alibaba-cloud-ops`

```text
Error: execution error at (mcp-server-alibaba-cloud-ops/templates/secrets.yaml:10:9): required value for secrets.ALIBABA_CLOUD_ACCESS_KEY_ID either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 69. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-alibabacloud-ops\charts\mcp-server-alibabacloud-ops`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-alibabacloud-ops\charts\mcp-server-alibabacloud-ops`

```text
Error: execution error at (mcp-server-alibabacloud-ops/templates/secrets.yaml:10:9): required value for secrets.ALIBABA_CLOUD_ACCESS_KEY_ID either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 70. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-alibabacloud-rds\charts\mcp-server-alibaba-cloud-rds`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-alibabacloud-rds\charts\mcp-server-alibaba-cloud-rds`

```text
Error: execution error at (mcp-server-alibaba-cloud-rds/templates/secrets.yaml:10:9): required value for secrets.ALIBABA_CLOUD_ACCESS_KEY_ID either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 71. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-alibabacloud-rds\charts\mcp-server-alibabacloud-rds`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-alibabacloud-rds\charts\mcp-server-alibabacloud-rds`

```text
Error: execution error at (mcp-server-alibabacloud-rds/templates/secrets.yaml:10:9): required value for secrets.ALIBABA_CLOUD_ACCESS_KEY_ID either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 72. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-anilist\charts\mcp-server-anilist`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-anilist\charts\mcp-server-anilist`

```text
Error: execution error at (mcp-server-anilist/templates/secrets.yaml:10:9): required value for secrets.ANILIST_TOKEN either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 73. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-apache-airflow\charts\mcp-server-apache-airflow`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-apache-airflow\charts\mcp-server-apache-airflow`

```text
Error: execution error at (mcp-server-apache-airflow/templates/secrets.yaml:10:9): required value for secrets.AIRFLOW_PASSWORD either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 74. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-apify-actors\charts\mcp-server-apify-actors`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-apify-actors\charts\mcp-server-apify-actors`

```text
Error: execution error at (mcp-server-apify-actors/templates/secrets.yaml:10:9): required value for secrets.APIFY_TOKEN either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 75. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-apify-rag-web-browser\charts\mcp-server-apify-rag-web-browser`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-apify-rag-web-browser\charts\mcp-server-apify-rag-web-browser`

```text
Error: execution error at (mcp-server-apify-rag-web-browser/templates/secrets.yaml:10:9): required value for secrets.APIFY_TOKEN either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 76. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-apimatic-validator\charts\mcp-server-apimatic-validator`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-apimatic-validator\charts\mcp-server-apimatic-validator`

```text
Error: execution error at (mcp-server-apimatic-validator/templates/secrets.yaml:10:9): required value for secrets.APIMATIC_API_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 77. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-asana\charts\mcp-server-asana`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-asana\charts\mcp-server-asana`

```text
Error: execution error at (mcp-server-asana/templates/secrets.yaml:10:9): required value for secrets.ASANA_ACCESS_TOKEN either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 78. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-asgardeo\charts\mcp-server-asgardeo`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-asgardeo\charts\mcp-server-asgardeo`

```text
Error: execution error at (mcp-server-asgardeo/templates/secrets.yaml:10:9): required value for secrets.ASGARDEO_CLIENT_ID either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 79. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-astra-db-mcp\charts\mcp-server-astra-db-mcp`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-astra-db-mcp\charts\mcp-server-astra-db-mcp`

```text
Error: execution error at (mcp-server-astra-db-mcp/templates/secrets.yaml:10:9): required value for secrets.ASTRA_DB_API_ENDPOINT either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 80. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-atla\charts\mcp-server-atla`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-atla\charts\mcp-server-atla`

```text
Error: execution error at (mcp-server-atla/templates/secrets.yaml:10:9): required value for secrets.ATLA_API_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 81. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-atlan\charts\mcp-server-atlan`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-atlan\charts\mcp-server-atlan`

```text
Error: execution error at (mcp-server-atlan/templates/secrets.yaml:10:9): required value for secrets.ATLAN_API_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 82. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-atlassian\charts\mcp-server-atlassian`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-atlassian\charts\mcp-server-atlassian`

```text
Error: execution error at (mcp-server-atlassian/templates/secrets.yaml:10:9): required value for secrets.CONFLUENCE_API_TOKEN either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 83. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-attio\charts\mcp-server-attio`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-attio\charts\mcp-server-attio`

```text
Error: execution error at (mcp-server-attio/templates/secrets.yaml:10:9): required value for secrets.ATTIO_API_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 84. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-audiense-insights\charts\mcp-server-audiense-insights`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-audiense-insights\charts\mcp-server-audiense-insights`

```text
Error: execution error at (mcp-server-audiense-insights/templates/secrets.yaml:10:9): required value for secrets.AUDIENSE_CLIENT_ID either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 85. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-aws-keyspaces\charts\mcp-server-aws-keyspaces`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-aws-keyspaces\charts\mcp-server-aws-keyspaces`

```text
Error: execution error at (mcp-server-aws-keyspaces/templates/secrets.yaml:10:9): required value for secrets.DB_CASSANDRA_PASSWORD either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 86. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-aws-valkey\charts\mcp-server-aws-valkey`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-aws-valkey\charts\mcp-server-aws-valkey`

```text
Error: execution error at (mcp-server-aws-valkey/templates/secrets.yaml:10:9): required value for secrets.VALKEY_PWD either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 87. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-axiom\charts\mcp-server-axiom`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-axiom\charts\mcp-server-axiom`

```text
Error: execution error at (mcp-server-axiom/templates/secrets.yaml:10:9): required value for secrets.AXIOM_ORG either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 88. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-azure\charts\mcp-server-azure`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-azure\charts\mcp-server-azure`

```text
Error: execution error at (mcp-server-azure/templates/secrets.yaml:10:9): required value for secrets.AZURE_CLIENT_ID either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 89. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-bankless-onchain\charts\mcp-server-bankless-onchain`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-bankless-onchain\charts\mcp-server-bankless-onchain`

```text
Error: execution error at (mcp-server-bankless-onchain/templates/secrets.yaml:10:9): required value for secrets.BANKLESS_API_TOKEN either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 90. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-benborla-mysql\charts\mcp-server-benborla-mysql`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-benborla-mysql\charts\mcp-server-benborla-mysql`

```text
Error: execution error at (mcp-server-benborla-mysql/templates/secrets.yaml:10:9): required value for secrets.MYSQL_PASS either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 91. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-bicscan\charts\mcp-server-bicscan`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-bicscan\charts\mcp-server-bicscan`

```text
Error: execution error at (mcp-server-bicscan/templates/secrets.yaml:10:9): required value for secrets.BICSCAN_API_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 92. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-bing-search\charts\mcp-server-bing-search`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-bing-search\charts\mcp-server-bing-search`

```text
Error: execution error at (mcp-server-bing-search/templates/secrets.yaml:10:9): required value for secrets.BING_API_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 93. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-bitrefill\charts\mcp-server-bitrefill`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-bitrefill\charts\mcp-server-bitrefill`

```text
Error: execution error at (mcp-server-bitrefill/templates/secrets.yaml:10:9): required value for secrets.BITREFILL_API_ID either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 94. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-bitrise\charts\mcp-server-bitrise`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-bitrise\charts\mcp-server-bitrise`

```text
Error: execution error at (mcp-server-bitrise/templates/secrets.yaml:10:9): required value for secrets.BITRISE_TOKEN either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 95. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-brave-search\charts\mcp-server-brave-search`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-brave-search\charts\mcp-server-brave-search`

```text
Error: execution error at (mcp-server-brave-search/templates/secrets.yaml:10:9): required value for secrets.BRAVE_API_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 96. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-brightdata\charts\mcp-server-brightdata`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-brightdata\charts\mcp-server-brightdata`

```text
Error: execution error at (mcp-server-brightdata/templates/secrets.yaml:10:9): required value for secrets.API_TOKEN either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 97. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-browserbase\charts\mcp-server-browserbase`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-browserbase\charts\mcp-server-browserbase`

```text
Error: execution error at (mcp-server-browserbase/templates/secrets.yaml:10:9): required value for secrets.BROWSERBASE_API_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 98. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-browserstack\charts\mcp-server-browserstack`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-browserstack\charts\mcp-server-browserstack`

```text
Error: execution error at (mcp-server-browserstack/templates/secrets.yaml:10:9): required value for secrets.BROWSERSTACK_ACCESS_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 99. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-bugsnag\charts\mcp-server-bugsnag`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-bugsnag\charts\mcp-server-bugsnag`

```text
Error: execution error at (mcp-server-bugsnag/templates/secrets.yaml:10:9): required value for secrets.BUGSNAG_API_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 100. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-cartesia\charts\mcp-server-cartesia`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-cartesia\charts\mcp-server-cartesia`

```text
Error: execution error at (mcp-server-cartesia/templates/secrets.yaml:10:9): required value for secrets.CARTESIA_API_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 101. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-chronulus-ai\charts\mcp-server-chronulus-ai`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-chronulus-ai\charts\mcp-server-chronulus-ai`

```text
Error: execution error at (mcp-server-chronulus-ai/templates/secrets.yaml:10:9): required value for secrets.CHRONULUS_API_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 102. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-clickhouse\charts\mcp-server-clickhouse`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-clickhouse\charts\mcp-server-clickhouse`

```text
Error: execution error at (mcp-server-clickhouse/templates/secrets.yaml:10:9): required value for secrets.CLICKHOUSE_PASSWORD either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 103. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-coda\charts\mcp-server-coda`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-coda\charts\mcp-server-coda`

```text
Error: execution error at (mcp-server-coda/templates/secrets.yaml:10:9): required value for secrets.API_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 104. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-codelogic\charts\mcp-server-codelogic`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-codelogic\charts\mcp-server-codelogic`

```text
Error: execution error at (mcp-server-codelogic/templates/secrets.yaml:10:9): required value for secrets.CODELOGIC_PASSWORD either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 105. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-codex\charts\mcp-server-codex`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-codex\charts\mcp-server-codex`

```text
Error: execution error at (mcp-server-codex/templates/secrets.yaml:10:9): required value for secrets.CODEX_API_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 106. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-contentful\charts\mcp-server-contentful`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-contentful\charts\mcp-server-contentful`

```text
Error: execution error at (mcp-server-contentful/templates/secrets.yaml:10:9): required value for secrets.CONTENTFUL_MANAGEMENT_ACCESS_TOKEN either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 107. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-couchbase\charts\mcp-server-couchbase`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-couchbase\charts\mcp-server-couchbase`

```text
Error: execution error at (mcp-server-couchbase/templates/secrets.yaml:10:9): required value for secrets.CB_PASSWORD either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 108. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-dappier\charts\mcp-server-dappier`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-dappier\charts\mcp-server-dappier`

```text
Error: execution error at (mcp-server-dappier/templates/secrets.yaml:10:9): required value for secrets.DAPPIER_API_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 109. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-dart\charts\mcp-server-dart`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-dart\charts\mcp-server-dart`

```text
Error: execution error at (mcp-server-dart/templates/secrets.yaml:10:9): required value for secrets.DART_TOKEN either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 110. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-datahub\charts\mcp-server-datahub`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-datahub\charts\mcp-server-datahub`

```text
Error: execution error at (mcp-server-datahub/templates/secrets.yaml:10:9): required value for secrets.DATAHUB_GMS_TOKEN either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 111. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-devhub\charts\mcp-server-devhub`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-devhub\charts\mcp-server-devhub`

```text
Error: execution error at (mcp-server-devhub/templates/secrets.yaml:10:9): required value for secrets.DEVHUB_API_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 112. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-doit\charts\mcp-server-doit`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-doit\charts\mcp-server-doit`

```text
Error: execution error at (mcp-server-doit/templates/secrets.yaml:10:9): required value for secrets.DOIT_API_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 113. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-e2b\charts\mcp-server-e2b`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-e2b\charts\mcp-server-e2b`

```text
Error: execution error at (mcp-server-e2b/templates/secrets.yaml:10:9): required value for secrets.E2B_API_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 114. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-elasticsearch\charts\mcp-server-elasticsearch`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-elasticsearch\charts\mcp-server-elasticsearch`

```text
Error: execution error at (mcp-server-elasticsearch/templates/secrets.yaml:10:9): required value for secrets.ES_API_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 115. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-elevenlabs\charts\mcp-server-elevenlabs`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-elevenlabs\charts\mcp-server-elevenlabs`

```text
Error: execution error at (mcp-server-elevenlabs/templates/secrets.yaml:10:9): required value for secrets.ELEVENLABS_API_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 116. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-exa\charts\mcp-server-exa`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-exa\charts\mcp-server-exa`

```text
Error: execution error at (mcp-server-exa/templates/secrets.yaml:10:9): required value for secrets.EXA_API_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 117. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-fibery\charts\mcp-server-fibery`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-fibery\charts\mcp-server-fibery`

```text
Error: execution error at (mcp-server-fibery/templates/secrets.yaml:10:9): required value for secrets.FIBERY_API_TOKEN either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 118. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-figma\charts\mcp-server-figma`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-figma\charts\mcp-server-figma`

```text
Error: execution error at (mcp-server-figma/templates/secrets.yaml:10:9): required value for secrets.FIGMA_API_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 119. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-firecrawl\charts\mcp-server-firecrawl`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-firecrawl\charts\mcp-server-firecrawl`

```text
Error: execution error at (mcp-server-firecrawl/templates/secrets.yaml:10:9): required value for secrets.FIRECRAWL_API_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 120. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-github\charts\mcp-server-github`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-github\charts\mcp-server-github`

```text
Error: execution error at (mcp-server-github/templates/secrets.yaml:10:9): required value for secrets.GITHUB_PERSONAL_ACCESS_TOKEN either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 121. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-github-chat\charts\mcp-server-github-chat`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-github-chat\charts\mcp-server-github-chat`

```text
Error: execution error at (mcp-server-github-chat/templates/secrets.yaml:10:9): required value for secrets.GITHUB_API_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 122. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-gitlab\charts\mcp-server-gitlab`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-gitlab\charts\mcp-server-gitlab`

```text
Error: execution error at (mcp-server-gitlab/templates/secrets.yaml:10:9): required value for secrets.GITLAB_PERSONAL_ACCESS_TOKEN either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 123. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-glif\charts\mcp-server-glif`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-glif\charts\mcp-server-glif`

```text
Error: execution error at (mcp-server-glif/templates/secrets.yaml:10:9): required value for secrets.GLIF_API_TOKEN either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 124. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-google-bigquery\charts\mcp-server-google-bigquery`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-google-bigquery\charts\mcp-server-google-bigquery`

```text
Error: execution error at (mcp-server-google-bigquery/templates/secrets.yaml:10:9): required value for secrets.GCP_PROJECT_ID either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 125. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-google-maps\charts\mcp-server-google-maps`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-google-maps\charts\mcp-server-google-maps`

```text
Error: execution error at (mcp-server-google-maps/templates/secrets.yaml:10:9): required value for secrets.GOOGLE_MAPS_API_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 126. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-grafana\charts\mcp-server-grafana`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-grafana\charts\mcp-server-grafana`

```text
Error: execution error at (mcp-server-grafana/templates/secrets.yaml:10:9): required value for secrets.GRAFANA_API_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 127. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-graphlit\charts\mcp-server-graphlit`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-graphlit\charts\mcp-server-graphlit`

```text
Error: execution error at (mcp-server-graphlit/templates/secrets.yaml:10:9): required value for secrets.GRAPHLIT_JWT_SECRET either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 128. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-gyazo\charts\mcp-server-gyazo`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-gyazo\charts\mcp-server-gyazo`

```text
Error: execution error at (mcp-server-gyazo/templates/secrets.yaml:10:9): required value for secrets.GYAZO_ACCESS_TOKEN either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 129. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-hackle-mcp\charts\mcp-server-hackle-mcp`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-hackle-mcp\charts\mcp-server-hackle-mcp`

```text
Error: execution error at (mcp-server-hackle-mcp/templates/secrets.yaml:10:9): required value for secrets.API_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 130. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-harness\charts\mcp-server-harness`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-harness\charts\mcp-server-harness`

```text
Error: execution error at (mcp-server-harness/templates/secrets.yaml:10:9): required value for secrets.HARNESS_API_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 131. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-heroku\charts\mcp-server-heroku`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-heroku\charts\mcp-server-heroku`

```text
Error: execution error at (mcp-server-heroku/templates/secrets.yaml:10:9): required value for secrets.HEROKU_API_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 132. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-hyperbrowser\charts\mcp-server-hyperbrowser`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-hyperbrowser\charts\mcp-server-hyperbrowser`

```text
Error: execution error at (mcp-server-hyperbrowser/templates/secrets.yaml:10:9): required value for secrets.HYPERBROWSER_API_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 133. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-hyperspell\charts\mcp-server-hyperspell`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-hyperspell\charts\mcp-server-hyperspell`

```text
Error: execution error at (mcp-server-hyperspell/templates/secrets.yaml:10:9): required value for secrets.HYPERSPELL_TOKEN either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 134. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-iaptic\charts\mcp-server-iaptic`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-iaptic\charts\mcp-server-iaptic`

```text
Error: execution error at (mcp-server-iaptic/templates/secrets.yaml:10:9): required value for secrets.IAPTIC_API_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 135. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-iotdb\charts\mcp-server-iotdb`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-iotdb\charts\mcp-server-iotdb`

```text
Error: execution error at (mcp-server-iotdb/templates/secrets.yaml:10:9): required value for secrets.IOTDB_PASSWORD either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 136. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-kagisearch\charts\mcp-server-kagisearch`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-kagisearch\charts\mcp-server-kagisearch`

```text
Error: execution error at (mcp-server-kagisearch/templates/secrets.yaml:10:9): required value for secrets.KAGI_API_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 137. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-lara-translate\charts\mcp-server-lara-translate`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-lara-translate\charts\mcp-server-lara-translate`

```text
Error: execution error at (mcp-server-lara-translate/templates/secrets.yaml:10:9): required value for secrets.LARA_ACCESS_KEY_ID either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 138. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-line-bot\charts\mcp-server-line-bot`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-line-bot\charts\mcp-server-line-bot`

```text
Error: execution error at (mcp-server-line-bot/templates/secrets.yaml:10:9): required value for secrets.CHANNEL_ACCESS_TOKEN either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 139. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-metatool-app\charts\mcp-server-metatool-app`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-metatool-app\charts\mcp-server-metatool-app`

```text
Error: execution error at (mcp-server-metatool-app/templates/secrets.yaml:10:9): required value for secrets.METAMCP_API_BASE_URL either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 140. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-microsoft-graph\charts\mcp-server-microsoft-graph`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-microsoft-graph\charts\mcp-server-microsoft-graph`

```text
Error: execution error at (mcp-server-microsoft-graph/templates/secrets.yaml:10:9): required value for secrets.MCP_SERVER_MICROSOFT_GRAPH_CLIENT_ID either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 141. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-mobsf\charts\mcp-server-mobsf`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-mobsf\charts\mcp-server-mobsf`

```text
Error: execution error at (mcp-server-mobsf/templates/secrets.yaml:10:9): required value for secrets.MOBSF_API_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 142. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-neo4j-aura\charts\mcp-server-neo4j-aura`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-neo4j-aura\charts\mcp-server-neo4j-aura`

```text
Error: execution error at (mcp-server-neo4j-aura/templates/secrets.yaml:10:9): required value for secrets.NEO4J_AURA_CLIENT_ID either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 143. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-neo4j-cypher\charts\mcp-server-neo4j-cypher`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-neo4j-cypher\charts\mcp-server-neo4j-cypher`

```text
Error: execution error at (mcp-server-neo4j-cypher/templates/secrets.yaml:10:9): required value for secrets.NEO4J_PASSWORD either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 144. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-neo4j-memory\charts\mcp-server-neo4j-memory`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-neo4j-memory\charts\mcp-server-neo4j-memory`

```text
Error: execution error at (mcp-server-neo4j-memory/templates/secrets.yaml:10:9): required value for secrets.NEO4J_PASSWORD either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 145. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-neon\charts\mcp-server-neon`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-neon\charts\mcp-server-neon`

```text
Error: execution error at (mcp-server-neon/templates/secrets.yaml:10:9): required value for secrets.NEON_API_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 146. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-notion\charts\mcp-server-notion`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-notion\charts\mcp-server-notion`

```text
Error: execution error at (mcp-server-notion/templates/secrets.yaml:10:9): required value for secrets.OPENAPI_MCP_HEADERS either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 147. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-obsidian\charts\mcp-server-obsidian`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-obsidian\charts\mcp-server-obsidian`

```text
Error: execution error at (mcp-server-obsidian/templates/secrets.yaml:10:9): required value for secrets.OBSIDIAN_API_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 148. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-oxylabs\charts\mcp-server-oxylabs`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-oxylabs\charts\mcp-server-oxylabs`

```text
Error: execution error at (mcp-server-oxylabs/templates/secrets.yaml:10:9): required value for secrets.OXYLABS_PASSWORD either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 149. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-paypal\charts\mcp-server-paypal`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-paypal\charts\mcp-server-paypal`

```text
Error: execution error at (mcp-server-paypal/templates/secrets.yaml:10:9): required value for secrets.PAYPAL_ACCESS_TOKEN either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 150. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-postgres\charts\mcp-server-postgres`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-postgres\charts\mcp-server-postgres`

```text
Error: execution error at (mcp-server-postgres/templates/secrets.yaml:10:9): required value for secrets.DATABASE_URI either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 151. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-razorpay\charts\mcp-server-razorpay`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-razorpay\charts\mcp-server-razorpay`

```text
Error: execution error at (mcp-server-razorpay/templates/secrets.yaml:10:9): required value for secrets.RAZORPAY_KEY_ID either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 152. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-risken\charts\mcp-server-risken`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-risken\charts\mcp-server-risken`

```text
Error: execution error at (mcp-server-risken/templates/secrets.yaml:10:9): required value for secrets.RISKEN_ACCESS_TOKEN either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 153. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-scrapezy\charts\mcp-server-scrapezy`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-scrapezy\charts\mcp-server-scrapezy`

```text
Error: execution error at (mcp-server-scrapezy/templates/secrets.yaml:10:9): required value for secrets.SCRAPEZY_API_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 154. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-send-email\charts\mcp-server-send-email`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-send-email\charts\mcp-server-send-email`

```text
Error: execution error at (mcp-server-send-email/templates/secrets.yaml:10:9): required value for secrets.RESEND_API_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 155. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-sentry\charts\mcp-server-sentry`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-sentry\charts\mcp-server-sentry`

```text
Error: execution error at (mcp-server-sentry/templates/secrets.yaml:10:9): required value for secrets.SENTRY_TOKEN either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 156. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-slack\charts\mcp-server-slack`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-slack\charts\mcp-server-slack`

```text
Error: execution error at (mcp-server-slack/templates/secrets.yaml:10:9): required value for secrets.SLACK_BOT_TOKEN either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 157. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-stripe\charts\mcp-server-stripe`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-stripe\charts\mcp-server-stripe`

```text
Error: execution error at (mcp-server-stripe/templates/secrets.yaml:10:9): required value for secrets.STRIPE_SECRET_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 158. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-tavily\charts\mcp-server-tavily`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-tavily\charts\mcp-server-tavily`

```text
Error: execution error at (mcp-server-tavily/templates/secrets.yaml:10:9): required value for secrets.TAVILY_API_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 159. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-triplewhale\charts\mcp-server-triplewhale`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-triplewhale\charts\mcp-server-triplewhale`

```text
Error: execution error at (mcp-server-triplewhale/templates/secrets.yaml:10:9): required value for secrets.TRIPLEWHALE_API_KEY either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 160. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-webflow\charts\mcp-server-webflow`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-webflow\charts\mcp-server-webflow`

```text
Error: execution error at (mcp-server-webflow/templates/secrets.yaml:10:9): required value for secrets.WEBFLOW_TOKEN either as .value or .valueFrom.name and .valueFrom.key

Use --debug flag to render out invalid YAML
```

### 161. `YAKEcloud/yake`

- Chart: `D:\helm_clones_github\YAKEcloud__yake\helmcharts\acl`
- Source: `dependency`

```text
Error: dependency "controller" has an invalid version/constraint format: improper constraint: ""
```

### 162. `cloud-native-toolkit/toolkit-charts`

- Chart: `D:\helm_clones_github\cloud-native-toolkit__toolkit-charts\stable\cloud-portworx`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\cloud-native-toolkit__toolkit-charts\stable\cloud-portworx`

```text
Error: execution error at (cloud-portworx/templates/subscription.yaml:6:16): Portworx must be installed in kube-system namespace

Use --debug flag to render out invalid YAML
```

### 163. `suse-edge/charts`

- Chart: `D:\helm_clones_github\suse-edge__charts\charts\kubevirt\0.1.0`
- Source: `dependency`

```text
Error: dependency "cdi" has an invalid version/constraint format: improper constraint: ""
```

### 164. `cnrancher/pandaria-catalog`

- Chart: `D:\helm_clones_github\cnrancher__pandaria-catalog\charts\rancher-f5-cis\107.0.0+up0.0.29`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\cnrancher__pandaria-catalog\charts\rancher-f5-cis\107.0.0+up0.0.29`

```text
Error: execution error at (rancher-f5-cis/templates/f5-bigip-ctlr-deploy.yaml:5:4): BIG-IP url not specified - add to Values or pass with `--set` 

Use --debug flag to render out invalid YAML
```

### 165. `cnrancher/pandaria-catalog`

- Chart: `D:\helm_clones_github\cnrancher__pandaria-catalog\charts\rancher-f5-cis\108.0.0+up0.0.29`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\cnrancher__pandaria-catalog\charts\rancher-f5-cis\108.0.0+up0.0.29`

```text
Error: execution error at (rancher-f5-cis/templates/f5-bigip-ctlr-deploy.yaml:5:4): BIG-IP url not specified - add to Values or pass with `--set` 

Use --debug flag to render out invalid YAML
```

### 166. `boozallen/sdp-helm-chart`

- Chart: `D:\helm_clones_github\boozallen__sdp-helm-chart`
- Source: `dependency`

```text
Error: dependency "jenkins" has an invalid version/constraint format: improper constraint: ""
```

### 167. `llajas/homelab`

- Chart: `D:\helm_clones_github\llajas__homelab\apps\clusterplex`
- Source: `template`
- Values files: `D:\helm_clones_github\llajas__homelab\apps\clusterplex\custom-values.yaml`
- Command: `helm template test D:\helm_clones_github\llajas__homelab\apps\clusterplex -f D:\helm_clones_github\llajas__homelab\apps\clusterplex\custom-values.yaml`

```text
Error: execution error at (clusterplex/templates/worker.yaml:161:9): Invalid value for worker.config.replicas. Must be greater than or equal to 1 when worker.enabled is true.

Use --debug flag to render out invalid YAML
```

### 168. `Merck/Data-Profiler`

- Chart: `D:\helm_clones_github\Merck__Data-Profiler\infrastructure\kube\helm-charts\namespace-setup\dp-cluster-proxies`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Merck__Data-Profiler\infrastructure\kube\helm-charts\namespace-setup\dp-cluster-proxies`

```text
Error: execution error at (dp-cluster-proxies/templates/spark-master.yml:62:13): A valid .values.loadBalancer entry required!

Use --debug flag to render out invalid YAML
```

### 169. `hsmade/velero-ui`

- Chart: `D:\helm_clones_github\hsmade__velero-ui\chart`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\hsmade__velero-ui\chart`

```text
Error: unsupported chart version
```

### 170. `appscode-cloud/ui-wizards`

- Chart: `D:\helm_clones_github\appscode-cloud__ui-wizards\charts\kubedbcom-elasticsearch-editor`
- Source: `dependency`

```text
Error: chart file "values.openapiv3_schema.yaml" is larger than the maximum file size 5242880
```

### 171. `dungdm93/shipyard`

- Chart: `D:\helm_clones_github\dungdm93__shipyard\helm\cloudflared`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\dungdm93__shipyard\helm\cloudflared`

```text
Error: execution error at (cloudflared/templates/deployment.yaml:50:41): Missing .Values.token

Use --debug flag to render out invalid YAML
```

### 172. `dungdm93/shipyard`

- Chart: `D:\helm_clones_github\dungdm93__shipyard\helm\datahub`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\dungdm93__shipyard\helm\datahub`

```text
Error: execution error at (datahub/templates/setup/secret.yaml:7:55): missing 'ebean.host'

Use --debug flag to render out invalid YAML
```

### 173. `Unique-AG/helm-charts`

- Chart: `D:\helm_clones_github\Unique-AG__helm-charts\charts\backend-service`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Unique-AG__helm-charts\charts\backend-service`

```text
Error: execution error at (backend-service/templates/routes.yaml:228:4): You can't use routes without gateway.networking.k8s.io CRDs installed. Install CRDs first.

Use --debug flag to render out invalid YAML
```

### 174. `Unique-AG/helm-charts`

- Chart: `D:\helm_clones_github\Unique-AG__helm-charts\charts\web-app`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Unique-AG__helm-charts\charts\web-app`

```text
Error: execution error at (web-app/templates/routes.yaml:71:4): You can't use routes without gateway.networking.k8s.io CRDs installed. Install CRDs first.

Use --debug flag to render out invalid YAML
```

### 175. `claytono/infra`

- Chart: `D:\helm_clones_github\claytono__infra\kubernetes\codex-runners`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\claytono__infra\kubernetes\codex-runners`

```text
Error: execution error at (codex-runners-meta/charts/gha-runner-scale-set/templates/manager_role_binding.yaml:42:11): No gha-rs-controller deployment found using label (app.kubernetes.io/part-of=gha-rs-controller). Consider setting controllerServiceAccount.name in values.yaml to be explicit if you think the discovery is wrong.

Use --debug flag to render out invalid YAML
```

### 176. `JuniorJPDJ/charts`

- Chart: `D:\helm_clones_github\JuniorJPDJ__charts\charts\deluge`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\JuniorJPDJ__charts\charts\deluge`

```text
Error: execution error at (deluge/templates/common.yaml:13:4): Duplicate port 6881/TCP found in Service. (service: 'torrent', ports: 'tcp' and 'udp')

Use --debug flag to render out invalid YAML
```

### 177. `kast-spells/kast-system`

- Chart: `D:\helm_clones_github\kast-spells__kast-system\covenant`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\kast-spells__kast-system\covenant`

```text
Error: execution error at (covenant/templates/covenant.yaml:39:6): covenant/index.yaml not found in bookrack/test

Use --debug flag to render out invalid YAML
```

### 178. `opspresso/argocd-env-addons`

- Chart: `D:\helm_clones_github\opspresso__argocd-env-addons\charts\aws-load-balancer-controller`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\opspresso__argocd-env-addons\charts\aws-load-balancer-controller`

```text
Error: execution error at (aws-load-balancer-controller/charts/aws-load-balancer-controller/templates/deployment.yaml:65:28): Chart cannot be installed without a valid clusterName!

Use --debug flag to render out invalid YAML
```

### 179. `rancher/ob-team-charts`

- Chart: `D:\helm_clones_github\rancher__ob-team-charts\charts\prometheus-federator\0.0.1`
- Source: `dependency`

```text
Error: dependency "helmProjectOperator" has an invalid version/constraint format: improper constraint: ""
```

### 180. `thelande/charts`

- Chart: `D:\helm_clones_github\thelande__charts\charts\sdtd-exporter`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\thelande__charts\charts\sdtd-exporter`

```text
Error: execution error at (sdtd-exporter/templates/deployment.yaml:53:24): apiUrl must be provided

Use --debug flag to render out invalid YAML
```

### 181. `K-FOSS/CoRE-CNTRL`

- Chart: `D:\helm_clones_github\K-FOSS__CoRE-CNTRL\Archive\CNTRLPlane`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\K-FOSS__CoRE-CNTRL\Archive\CNTRLPlane`

```text
level=ERROR msg="chart dependencies processing failed" error="type mismatch on kamaji: %!t(<nil>)"
Error: chart dependencies processing failed: type mismatch on kamaji: %!t(<nil>)
```

### 182. `Kapil-Bhalodiya/E-commerce-Platform`

- Chart: `D:\helm_clones_github\Kapil-Bhalodiya__E-commerce-Platform\infra\addons\frontend`
- Source: `dependency`

```text
Error: cannot load values.yaml: cannot unmarshal yaml document: error converting YAML to JSON: yaml: line 8: could not find expected ':'
```

### 183. `codefuturist/helm-charts`

- Chart: `D:\helm_clones_github\codefuturist__helm-charts\charts\application`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\codefuturist__helm-charts\charts\application`

```text
Error: execution error at (application/templates/deployment.yaml:144:20): Undefined image for application container

Use --debug flag to render out invalid YAML
```

### 184. `codefuturist/helm-charts`

- Chart: `D:\helm_clones_github\codefuturist__helm-charts\charts\application`
- Source: `template`
- Values files: `D:\helm_clones_github\codefuturist__helm-charts\charts\application\values-test.yaml`
- Command: `helm template test D:\helm_clones_github\codefuturist__helm-charts\charts\application -f D:\helm_clones_github\codefuturist__helm-charts\charts\application\values-test.yaml`

```text
Error: execution error at (application/templates/vpa.yaml:3:6): There is no VerticalPodAutoscaler resource definition in the target cluster!

Use --debug flag to render out invalid YAML
```

### 185. `codefuturist/helm-charts`

- Chart: `D:\helm_clones_github\codefuturist__helm-charts\charts\apps\restic-backup`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\codefuturist__helm-charts\charts\apps\restic-backup`

```text
Error: execution error at (restic-backup/templates/job-init.yaml:2:10): At least one volume must be specified in .Values.volumes when backup is enabled

Use --debug flag to render out invalid YAML
```

### 186. `codefuturist/helm-charts`

- Chart: `D:\helm_clones_github\codefuturist__helm-charts\charts\homarr`
- Source: `template`
- Values files: `D:\helm_clones_github\codefuturist__helm-charts\charts\homarr\values-test.yaml`
- Command: `helm template test D:\helm_clones_github\codefuturist__helm-charts\charts\homarr -f D:\helm_clones_github\codefuturist__helm-charts\charts\homarr\values-test.yaml`

```text
Error: execution error at (homarr/templates/vpa.yaml:3:6): There is no VerticalPodAutoscaler resource definition in the target cluster!

Use --debug flag to render out invalid YAML
```

### 187. `dboeckli/spring-6-icecold-micro-service`

- Chart: `D:\helm_clones_github\dboeckli__spring-6-icecold-micro-service\helm-charts`
- Source: `dependency`

```text
Error: cannot load Chart.yaml: error converting YAML to JSON: yaml: line 4: found character that cannot start any token
```

### 188. `100rd/platform-design`

- Chart: `D:\helm_clones_github\100rd__platform-design\apps\infra\aws-lb-controller`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\100rd__platform-design\apps\infra\aws-lb-controller`

```text
Error: execution error at (aws-load-balancer-controller/charts/aws-load-balancer-controller/templates/deployment.yaml:67:28): Chart cannot be installed without a valid clusterName!

Use --debug flag to render out invalid YAML
```

### 189. `Clark1992/ECK1`

- Chart: `D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\zitadel-upstream`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Clark1992__ECK1\infra\k8s\charts\zitadel-upstream`

```text
Error: execution error at (zitadel/templates/secret_zitadel-masterkey.yaml:2:4): Either set .Values.zitadel.masterkey xor .Values.zitadel.masterkeySecretName

Use --debug flag to render out invalid YAML
```

## `template.kube_version_incompatible` (140)

### 1. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\aerospike`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\aerospike`

```text
Error: chart requires kubeVersion: <= 1.22.5-x which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 2. `openshift-helm-charts/stage`

- Chart: `D:\helm_clones_github\openshift-helm-charts__stage\charts\partners\redhat-sp\helmstage1\0.1.0\src`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\openshift-helm-charts__stage\charts\partners\redhat-sp\helmstage1\0.1.0\src`

```text
Error: chart requires kubeVersion: 1.20.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 3. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\bootstrap-capi-1.10.8`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\bootstrap-capi-1.10.8`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 4. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\bootstrap-capi-1.7.1`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\bootstrap-capi-1.7.1`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.32.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 5. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\bootstrap-capi-1.8.12`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\bootstrap-capi-1.8.12`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.32.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 6. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\bootstrap-capi-1.9.4`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\bootstrap-capi-1.9.4`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.33.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 7. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\bootstrap-capi-1.9.9`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\bootstrap-capi-1.9.9`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 8. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\cert-manager-1.14.5`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\cert-manager-1.14.5`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.32.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 9. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\cert-manager-1.16.3`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\cert-manager-1.16.3`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.32.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 10. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\cert-manager-1.17.2`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\cert-manager-1.17.2`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 11. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\cert-manager-1.19.1`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\cert-manager-1.19.1`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 12. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\cert-manager-1.19.4`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\cert-manager-1.19.4`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 13. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\cert-manager-1.9.1`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\cert-manager-1.9.1`

```text
Error: chart requires kubeVersion: < 1.31.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 14. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\control-plane-capi-1.10.8`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\control-plane-capi-1.10.8`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 15. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\control-plane-capi-1.7.1`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\control-plane-capi-1.7.1`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.32.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 16. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\control-plane-capi-1.8.12`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\control-plane-capi-1.8.12`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.32.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 17. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\control-plane-capi-1.9.4`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\control-plane-capi-1.9.4`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.33.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 18. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\control-plane-capi-1.9.9`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\control-plane-capi-1.9.9`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 19. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\core-capi-1.10.8`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\core-capi-1.10.8`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 20. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\core-capi-1.7.1`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\core-capi-1.7.1`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.32.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 21. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\core-capi-1.8.12`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\core-capi-1.8.12`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.32.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 22. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\core-capi-1.9.4`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\core-capi-1.9.4`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.33.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 23. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\core-capi-1.9.9`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\core-capi-1.9.9`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 24. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\csi-driver-nfs-4.11.0`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\csi-driver-nfs-4.11.0`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.33.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 25. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\csi-driver-nfs-4.12.1`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\csi-driver-nfs-4.12.1`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 26. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\csi-driver-nfs-4.13.1`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\csi-driver-nfs-4.13.1`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 27. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\dex-2.39.1`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\dex-2.39.1`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.33.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 28. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\dex-2.43.1`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\dex-2.43.1`

```text
Error: chart requires kubeVersion: >= 1.30 < 1.34 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 29. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\dex-2.44.0`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\dex-2.44.0`

```text
Error: chart requires kubeVersion: >= 1.30 < 1.34 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 30. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\flannel-0.22.3`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\flannel-0.22.3`

```text
Error: chart requires kubeVersion: < 1.31 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 31. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\fluent-operator-3.2.0`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\fluent-operator-3.2.0`

```text
Error: chart requires kubeVersion: >= 1.26.0 < 1.33.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 32. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\fluent-operator-3.4.0`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\fluent-operator-3.4.0`

```text
Error: chart requires kubeVersion: >= 1.26.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 33. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\fluent-operator-3.6.0`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\fluent-operator-3.6.0`

```text
Error: chart requires kubeVersion: >= 1.26.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 34. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\fluentd-1.14.5`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\fluentd-1.14.5`

```text
level=WARN msg="this chart is deprecated"
Error: chart requires kubeVersion: >= 1.26.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 35. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\grafana-10.2.6`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\grafana-10.2.6`

```text
Error: chart requires kubeVersion: >= 1.26.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 36. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\grafana-7.5.17`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\grafana-7.5.17`

```text
Error: chart requires kubeVersion: >= 1.26.0 < 1.33.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 37. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\grafana-9.2.10`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\grafana-9.2.10`

```text
Error: chart requires kubeVersion: >= 1.26.0 < 1.33.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 38. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\ingress-nginx-1.12.1`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\ingress-nginx-1.12.1`

```text
Error: chart requires kubeVersion: >= 1.26.0 < 1.33.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 39. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\ingress-nginx-1.12.5`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\ingress-nginx-1.12.5`

```text
Error: chart requires kubeVersion: >= 1.26.0 < 1.33.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 40. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\ingress-nginx-1.13.1`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\ingress-nginx-1.13.1`

```text
Error: chart requires kubeVersion: >= 1.26.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 41. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\ingress-nginx-1.14.0`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\ingress-nginx-1.14.0`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 42. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\ingress-nginx-1.14.4`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\ingress-nginx-1.14.4`

```text
level=WARN msg="this chart is deprecated"
Error: chart requires kubeVersion: >= 1.30.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 43. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\ingress-nginx-1.15.1`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\ingress-nginx-1.15.1`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 44. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\ingress-nginx-1.9.6`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\ingress-nginx-1.9.6`

```text
Error: chart requires kubeVersion: >= 1.26.0 < 1.32.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 45. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\istio-base-1.19.9`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\istio-base-1.19.9`

```text
Error: chart requires kubeVersion: < 1.31.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 46. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\istio-base-1.20.5`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\istio-base-1.20.5`

```text
Error: chart requires kubeVersion: >= 1.29.0 < 1.32.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 47. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\istio-base-1.20.8`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\istio-base-1.20.8`

```text
Error: chart requires kubeVersion: >= 1.29.0 < 1.32.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 48. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\istio-base-1.22.6`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\istio-base-1.22.6`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.32.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 49. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\istio-base-1.22.8`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\istio-base-1.22.8`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.32.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 50. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\istio-base-1.24.1`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\istio-base-1.24.1`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.33.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 51. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\istio-base-1.24.6`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\istio-base-1.24.6`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 52. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\istio-cni-1.24.1`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\istio-cni-1.24.1`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.33.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 53. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\istio-cni-1.24.6`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\istio-cni-1.24.6`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 54. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\istio-egress-1.19.9`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\istio-egress-1.19.9`

```text
Error: chart requires kubeVersion: < 1.31.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 55. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\istio-egress-1.20.5`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\istio-egress-1.20.5`

```text
Error: chart requires kubeVersion: >= 1.29.0 < 1.32.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 56. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\istio-egress-1.20.8`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\istio-egress-1.20.8`

```text
Error: chart requires kubeVersion: >= 1.29.0 < 1.32.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 57. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\istio-egress-1.22.6`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\istio-egress-1.22.6`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.32.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 58. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\istio-egress-1.22.8`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\istio-egress-1.22.8`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.32.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 59. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\istio-egress-1.24.1`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\istio-egress-1.24.1`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.33.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 60. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\istio-egress-1.24.6`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\istio-egress-1.24.6`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 61. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\istio-ingress-1.19.9`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\istio-ingress-1.19.9`

```text
Error: chart requires kubeVersion: < 1.31.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 62. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\istio-ingress-1.20.5`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\istio-ingress-1.20.5`

```text
Error: chart requires kubeVersion: >= 1.29.0 < 1.32.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 63. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\istio-ingress-1.20.8`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\istio-ingress-1.20.8`

```text
Error: chart requires kubeVersion: >= 1.29.0 < 1.32.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 64. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\istio-ingress-1.22.6`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\istio-ingress-1.22.6`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.32.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 65. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\istio-ingress-1.22.8`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\istio-ingress-1.22.8`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.32.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 66. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\istio-ingress-1.24.1`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\istio-ingress-1.24.1`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.33.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 67. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\istio-ingress-1.24.6`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\istio-ingress-1.24.6`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 68. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\istio-ztunnel-1.24.1`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\istio-ztunnel-1.24.1`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.33.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 69. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\istio-ztunnel-1.24.6`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\istio-ztunnel-1.24.6`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 70. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\istiod-1.19.9`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\istiod-1.19.9`

```text
Error: chart requires kubeVersion: < 1.31.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 71. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\istiod-1.20.5`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\istiod-1.20.5`

```text
Error: chart requires kubeVersion: >= 1.29.0 < 1.32.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 72. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\istiod-1.20.5`
- Source: `template`
- Values files: `D:\helm_clones_github\oracle-cne__catalog\charts\istiod-1.20.5\ambient-values.yaml`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\istiod-1.20.5 -f D:\helm_clones_github\oracle-cne__catalog\charts\istiod-1.20.5\ambient-values.yaml`

```text
Error: chart requires kubeVersion: >= 1.29.0 < 1.32.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 73. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\istiod-1.20.8`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\istiod-1.20.8`

```text
Error: chart requires kubeVersion: >= 1.29.0 < 1.32.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 74. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\istiod-1.20.8`
- Source: `template`
- Values files: `D:\helm_clones_github\oracle-cne__catalog\charts\istiod-1.20.8\ambient-values.yaml`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\istiod-1.20.8 -f D:\helm_clones_github\oracle-cne__catalog\charts\istiod-1.20.8\ambient-values.yaml`

```text
Error: chart requires kubeVersion: >= 1.29.0 < 1.32.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 75. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\istiod-1.22.6`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\istiod-1.22.6`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.32.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 76. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\istiod-1.22.8`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\istiod-1.22.8`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.32.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 77. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\istiod-1.24.1`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\istiod-1.24.1`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.33.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 78. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\istiod-1.24.6`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\istiod-1.24.6`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 79. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\keycloak-21.1.2`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\keycloak-21.1.2`

```text
level=WARN msg="this chart is deprecated"
Error: chart requires kubeVersion: >= 1.26.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 80. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\kube-prometheus-stack-0.63.0`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\kube-prometheus-stack-0.63.0`

```text
Error: chart requires kubeVersion: >= 1.26.0 < 1.33.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 81. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\kube-prometheus-stack-0.85.0`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\kube-prometheus-stack-0.85.0`

```text
Error: chart requires kubeVersion: >= 1.26.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 82. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\kube-prometheus-stack-0.89.0`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\kube-prometheus-stack-0.89.0`

```text
Error: chart requires kubeVersion: >= 1.26.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 83. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\kube-state-metrics-2.17.0`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\kube-state-metrics-2.17.0`

```text
Error: chart requires kubeVersion: >= 1.26.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 84. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\kube-state-metrics-2.18.0`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\kube-state-metrics-2.18.0`

```text
Error: chart requires kubeVersion: >= 1.26.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 85. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\kube-state-metrics-2.8.2`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\kube-state-metrics-2.8.2`

```text
Error: chart requires kubeVersion: >= 1.26.0 < 1.33.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 86. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\kubernetes-gateway-api-crds-1.2.1`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\kubernetes-gateway-api-crds-1.2.1`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.33.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 87. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\kubernetes-gateway-api-crds-1.3.0`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\kubernetes-gateway-api-crds-1.3.0`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.33.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 88. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\kubernetes-gateway-api-crds-1.4.0`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\kubernetes-gateway-api-crds-1.4.0`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 89. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\kubevirt-0.58.0`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\kubevirt-0.58.0`

```text
Error: chart requires kubeVersion: <= 1.28.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 90. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\kubevirt-0.59.0`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\kubevirt-0.59.0`

```text
Error: chart requires kubeVersion: <= 1.29.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 91. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\kubevirt-1.0.1`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\kubevirt-1.0.1`

```text
Error: chart requires kubeVersion: <= 1.30.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 92. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\kubevirt-1.1.1`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\kubevirt-1.1.1`

```text
Error: chart requires kubeVersion: >= 1.29.0 < 1.33.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 93. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\kubevirt-1.2.2`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\kubevirt-1.2.2`

```text
Error: chart requires kubeVersion: >= 1.28.0 < 1.30.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 94. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\kubevirt-1.3.1`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\kubevirt-1.3.1`

```text
Error: chart requires kubeVersion: >= 1.28.0 < 1.31.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 95. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\kubevirt-1.4.1`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\kubevirt-1.4.1`

```text
Error: chart requires kubeVersion: >= 1.29.0 < 1.32.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 96. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\kubevirt-1.5.2`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\kubevirt-1.5.2`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 97. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\metallb-0.12.1`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\metallb-0.12.1`

```text
Error: chart requires kubeVersion: <= 1.31.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 98. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\metallb-0.13.10`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\metallb-0.13.10`

```text
Error: chart requires kubeVersion: >= 1.26.0 < 1.32.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 99. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\metallb-0.15.2`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\metallb-0.15.2`

```text
Error: chart requires kubeVersion: > 1.29.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 100. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\metallb-0.15.3`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\metallb-0.15.3`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 101. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\multus-4.0.2`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\multus-4.0.2`

```text
Error: chart requires kubeVersion: >= 1.26.0 < 1.32.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 102. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\multus-4.2.1`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\multus-4.2.1`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 103. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\oauth2-proxy-7.12.0`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\oauth2-proxy-7.12.0`

```text
Error: chart requires kubeVersion: >= 1.26.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 104. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\oauth2-proxy-7.14.3`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\oauth2-proxy-7.14.3`

```text
Error: chart requires kubeVersion: >= 1.26.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 105. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\oauth2-proxy-7.8.0`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\oauth2-proxy-7.8.0`

```text
Error: chart requires kubeVersion: >= 1.26.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 106. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\oci-capi-0.15.0`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\oci-capi-0.15.0`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.32.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 107. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\oci-capi-0.16.0`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\oci-capi-0.16.0`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.33.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 108. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\oci-capi-0.17.0`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\oci-capi-0.17.0`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.33.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 109. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\oci-capi-0.19.0`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\oci-capi-0.19.0`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 110. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\oci-capi-0.21.0`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\oci-capi-0.21.0`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 111. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\oci-ccm-1.27.2`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\oci-ccm-1.27.2`

```text
Error: chart requires kubeVersion: <= 1.31.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 112. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\oci-ccm-1.28.0`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\oci-ccm-1.28.0`

```text
Error: chart requires kubeVersion: <= 1.31.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 113. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\oci-ccm-1.30.0`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\oci-ccm-1.30.0`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 114. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\olvm-capi-1.0.0`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\olvm-capi-1.0.0`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 115. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\opensearch-2.15.0`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\opensearch-2.15.0`

```text
Error: chart requires kubeVersion: >= 1.26.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 116. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\opensearch-dashboards-2.15.0`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\opensearch-dashboards-2.15.0`

```text
Error: chart requires kubeVersion: >= 1.26.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 117. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\ovirt-csi-driver-4.20.0`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\ovirt-csi-driver-4.20.0`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.33.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 118. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\ovirt-csi-driver-4.21.0`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\ovirt-csi-driver-4.21.0`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 119. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\ovirt-csi-driver-4.21.0-alpha1`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\ovirt-csi-driver-4.21.0-alpha1`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 120. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\prometheus-2.31.1`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\prometheus-2.31.1`

```text
Error: chart requires kubeVersion: >= 1.26.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 121. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\prometheus-adapter-0.10.0`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\prometheus-adapter-0.10.0`

```text
Error: chart requires kubeVersion: >= 1.26.0 < 1.33.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 122. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\prometheus-adapter-0.12.0`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\prometheus-adapter-0.12.0`

```text
Error: chart requires kubeVersion: >= 1.26.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 123. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\prometheus-node-exporter-1.10.2`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\prometheus-node-exporter-1.10.2`

```text
Error: chart requires kubeVersion: >= 1.26.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 124. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\prometheus-node-exporter-1.6.1`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\prometheus-node-exporter-1.6.1`

```text
Error: chart requires kubeVersion: >= 1.26.0 < 1.33.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 125. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\prometheus-node-exporter-1.9.1`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\prometheus-node-exporter-1.9.1`

```text
Error: chart requires kubeVersion: >= 1.26.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 126. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\rook-1.10.9`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\rook-1.10.9`

```text
Error: chart requires kubeVersion: <= 1.31.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 127. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\rook-1.11.6`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\rook-1.11.6`

```text
Error: chart requires kubeVersion: <= 1.29.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 128. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\rook-1.12.3`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\rook-1.12.3`

```text
Error: chart requires kubeVersion: >= 1.26.0 < 1.33.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 129. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\rook-1.13.10`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\rook-1.13.10`

```text
Error: chart requires kubeVersion: >= 1.26.0 < 1.33.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 130. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\rook-1.14.12`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\rook-1.14.12`

```text
Error: chart requires kubeVersion: >= 1.26.0 < 1.33.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 131. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\rook-1.15.9`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\rook-1.15.9`

```text
Error: chart requires kubeVersion: >= 1.26.0 < 1.33.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 132. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\rook-1.16.6`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\rook-1.16.6`

```text
Error: chart requires kubeVersion: >= 1.27.0 < 1.33.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 133. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\rook-1.17.7`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\rook-1.17.7`

```text
Error: chart requires kubeVersion: >= 1.28.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 134. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\rook-1.18.0`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\rook-1.18.0`

```text
Error: chart requires kubeVersion: >= 1.29.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 135. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\tigera-operator-1.29.3`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\tigera-operator-1.29.3`

```text
Error: chart requires kubeVersion: <= 1.31.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 136. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\tigera-operator-1.32.12`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\tigera-operator-1.32.12`

```text
Error: chart requires kubeVersion: >= 1.28.0 < 1.33.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 137. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\tigera-operator-1.32.4`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\tigera-operator-1.32.4`

```text
Error: chart requires kubeVersion: >= 1.28.0 < 1.33.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 138. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\tigera-operator-1.38.1`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\tigera-operator-1.38.1`

```text
Error: chart requires kubeVersion: >= 1.30.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 139. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\ui-2.2.0`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\ui-2.2.0`

```text
Error: chart requires kubeVersion: >= 1.29.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

### 140. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\ui-2.3.0`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\ui-2.3.0`

```text
Error: chart requires kubeVersion: >= 1.29.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0

Use --debug flag to render out invalid YAML
```

## `dependency.missing_repository` (66)

### 1. `refly-ai/refly`

- Chart: `D:\helm_clones_github\refly-ai__refly\deploy\helm\refly-stack`
- Source: `dependency`

```text
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### 2. `norwoodj/helm-docs`

- Chart: `D:\helm_clones_github\norwoodj__helm-docs\example-charts\custom-template`
- Source: `dependency`

```text
Error: no repository definition for @stable. Please add them via 'helm repo add'
```

### 3. `securitybunker/databunker`

- Chart: `D:\helm_clones_github\securitybunker__databunker\charts\databunker`
- Source: `dependency`

```text
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### 4. `BigKAA/youtube`

- Chart: `D:\helm_clones_github\BigKAA__youtube\tracing\for_admins\charts\jaeger\jaeger`
- Source: `dependency`

```text
Error: no repository definition for https://charts.helm.sh/incubator, https://helm.elastic.co, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### 5. `vexxhost/atmosphere`

- Chart: `D:\helm_clones_github\vexxhost__atmosphere\charts\barbican`
- Source: `dependency`

```text
level=INFO msg="Warning: Dependency locking is handled in Chart.lock since apiVersion \"v2\". We recommend migrating to Chart.lock."
Error: no repository definition for https://tarballs.openstack.org/openstack-helm. Please add the missing repos via 'helm repo add'
```

### 6. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\monitoring\oauth2-proxy`
- Source: `dependency`

```text
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### 7. `wikibook/kubepractice`

- Chart: `D:\helm_clones_github\wikibook__kubepractice\ch06\nginx-12.0.0`
- Source: `dependency`

```text
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### 8. `tmforum-oda/oda-canvas`

- Chart: `D:\helm_clones_github\tmforum-oda__oda-canvas\charts\observability-stack`
- Source: `dependency`

```text
Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://open-telemetry.github.io/opentelemetry-helm-charts, https://jaegertracing.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### 9. `sa-mw-dach/bobbycar`

- Chart: `D:\helm_clones_github\sa-mw-dach__bobbycar\helm\bobbycar-core-infra`
- Source: `dependency`

```text
Error: no repository definition for https://drogue-iot.github.io/drogue-cloud-helm-charts/. Please add the missing repos via 'helm repo add'
```

### 10. `platyplus/platyplus`

- Chart: `D:\helm_clones_github\platyplus__platyplus\charts\hasura`
- Source: `dependency`

```text
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### 11. `bflance/proxmox-talos`

- Chart: `D:\helm_clones_github\bflance__proxmox-talos\charts\kube-prometheus-stack`
- Source: `dependency`

```text
Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### 12. `k0rdent/catalog`

- Chart: `D:\helm_clones_github\k0rdent__catalog\apps\alloy\charts\alloy-1.6.1`
- Source: `dependency`

```text
Error: no repository definition for https://grafana.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### 13. `junghoon2/k8s-class`

- Chart: `D:\helm_clones_github\junghoon2__k8s-class\argo-cd\argo-cd-5.14.1`
- Source: `dependency`

```text
Error: no repository definition for https://dandydeveloper.github.io/charts/. Please add the missing repos via 'helm repo add'
```

### 14. `ai-solution-eng/frameworks`

- Chart: `D:\helm_clones_github\ai-solution-eng__frameworks\appsmith\3.6.4`
- Source: `dependency`

```text
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### 15. `cnrancher/pandaria-catalog`

- Chart: `D:\helm_clones_github\cnrancher__pandaria-catalog\charts\rancher-hami\107.0.0+up2.5.2\charts\hami-webui`
- Source: `dependency`

```text
Error: no repository definition for https://nvidia.github.io/dcgm-exporter/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### 16. `Sagar2366/tech-talks`

- Chart: `D:\helm_clones_github\Sagar2366__tech-talks\k8s_pune_oct22\prometheus-comunity-helm-chart\charts\kube-prometheus-stack`
- Source: `dependency`

```text
Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://grafana.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### 17. `randoli/helm-charts`

- Chart: `D:\helm_clones_github\randoli__helm-charts\charts\cost-management`
- Source: `dependency`

```text
Error: no repository definition for https://opencost.github.io/opencost-helm-chart. Please add the missing repos via 'helm repo add'
```

### 18. `unixfox/k8s`

- Chart: `D:\helm_clones_github\unixfox__k8s\charts\bibliogram`
- Source: `dependency`

```text
Error: no repository definition for https://library-charts.k8s-at-home.com. Please add the missing repos via 'helm repo add'
```

### 19. `henrywhitaker3/homelab`

- Chart: `D:\helm_clones_github\henrywhitaker3__homelab\kubernetes\k3s\apps\databases\nats\cluster\chart`
- Source: `dependency`

```text
Error: no repository definition for https://nats-io.github.io/k8s/helm/charts/. Please add the missing repos via 'helm repo add'
```

### 20. `kubero-dev/kubero-operator`

- Chart: `D:\helm_clones_github\kubero-dev__kubero-operator\helm-charts\kuberoaddonmongodb`
- Source: `dependency`

```text
Error: no repository definition for https://groundhog2k.github.io/helm-charts/. Please add the missing repos via 'helm repo add'
```

### 21. `junghoon2/kube-books`

- Chart: `D:\helm_clones_github\junghoon2__kube-books\ch06\nginx-12.0.0`
- Source: `dependency`

```text
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### 22. `preloop/preloop`

- Chart: `D:\helm_clones_github\preloop__preloop\helm\preloop`
- Source: `dependency`

```text
Error: no repository definition for https://nats-io.github.io/k8s/helm/charts. Please add the missing repos via 'helm repo add'
```

### 23. `SimCubeLtd/simcube-helm-charts`

- Chart: `D:\helm_clones_github\SimCubeLtd__simcube-helm-charts\charts\bytesafe`
- Source: `dependency`

```text
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### 24. `teddy-ambona/kind-e2e`

- Chart: `D:\helm_clones_github\teddy-ambona__kind-e2e\helm\loki`
- Source: `dependency`

```text
Error: no repository definition for https://grafana.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### 25. `CDCgov/NEDSS-Helm`

- Chart: `D:\helm_clones_github\CDCgov__NEDSS-Helm\charts\strimzi`
- Source: `dependency`

```text
Error: no repository definition for https://strimzi.io/charts/. Please add the missing repos via 'helm repo add'
```

### 26. `SpechtLabs/k8s-deployment`

- Chart: `D:\helm_clones_github\SpechtLabs__k8s-deployment\charts\cert-checker`
- Source: `dependency`

```text
Error: no repository definition for https://mogensen.github.io/cert-checker. Please add the missing repos via 'helm repo add'
```

### 27. `claytono/infra`

- Chart: `D:\helm_clones_github\claytono__infra\kubernetes\crowdsec`
- Source: `dependency`

```text
Error: no repository definition for https://crowdsecurity.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### 28. `griggheo/blogomatic`

- Chart: `D:\helm_clones_github\griggheo__blogomatic\devops\bootstrap_kind_cluster\helm_charts\signoz\signoz`
- Source: `dependency`

```text
Error: no repository definition for https://signoz.github.io/charts, https://signoz.github.io/charts, https://charts.jetstack.io, https://kubernetes.github.io/ingress-nginx, https://charts.min.io, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### 29. `cnieg/helm-charts`

- Chart: `D:\helm_clones_github\cnieg__helm-charts\charts\clamapi`
- Source: `dependency`

```text
Error: no repository definition for https://wiremind.github.io/wiremind-helm-charts. Please add the missing repos via 'helm repo add'
```

### 30. `dynatrace-wwse/enablement-kubernetes-opentelemetry-openpipeline`

- Chart: `D:\helm_clones_github\dynatrace-wwse__enablement-kubernetes-opentelemetry-openpipeline\.devcontainer\astroshop\helm\dt-otel-demo-helm`
- Source: `dependency`

```text
Error: no repository definition for https://open-telemetry.github.io/opentelemetry-helm-charts, https://open-telemetry.github.io/opentelemetry-helm-charts. Please add the missing repos via 'helm repo add'
```

### 31. `helxplatform/translator-devops`

- Chart: `D:\helm_clones_github\helxplatform__translator-devops\helm\answer-appraiser`
- Source: `dependency`

```text
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### 32. `pluralsh/plural-helm-charts`

- Chart: `D:\helm_clones_github\pluralsh__plural-helm-charts\charts\airbyte`
- Source: `dependency`

```text
Error: no repository definition for https://charts.bitnami.com/bitnami, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts. Please add the missing repos via 'helm repo add'
```

### 33. `AchillesChan/memo`

- Chart: `D:\helm_clones_github\AchillesChan__memo\helm-demo\prometheus-charts\charts\kube-prometheus-stack`
- Source: `dependency`

```text
Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://grafana.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### 34. `AntSan813/react-hasura-keycloak-app`

- Chart: `D:\helm_clones_github\AntSan813__react-hasura-keycloak-app\api\hasura`
- Source: `dependency`

```text
Error: no repository definition for https://hasura.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### 35. `devops4solutions/guestbook`

- Chart: `D:\helm_clones_github\devops4solutions__guestbook`
- Source: `dependency`

```text
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### 36. `edixos/ekp-helm`

- Chart: `D:\helm_clones_github\edixos__ekp-helm\charts\alertmanager`
- Source: `dependency`

```text
Error: no repository definition for https://oauth2-proxy.github.io/manifests. Please add the missing repos via 'helm repo add'
```

### 37. `elastic/k8s-integration-infra`

- Chart: `D:\helm_clones_github\elastic__k8s-integration-infra\infra\charts\elastic-agent`
- Source: `dependency`

```text
Error: no repository definition for @stable. Please add them via 'helm repo add'
```

### 38. `ishtiaqhimel/oms`

- Chart: `D:\helm_clones_github\ishtiaqhimel__oms\charts\oms-server`
- Source: `dependency`

```text
Error: no repository definition for https://charts.konghq.com. Please add the missing repos via 'helm repo add'
```

### 39. `kalavai-net/helm-charts`

- Chart: `D:\helm_clones_github\kalavai-net__helm-charts\deployments\monitoring`
- Source: `dependency`

```text
Error: no repository definition for https://grafana.github.io/helm-charts, https://ckotzbauer.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### 40. `shelleg/ac-k8s`

- Chart: `D:\helm_clones_github\shelleg__ac-k8s\helm\ant-umbrella`
- Source: `dependency`

```text
Error: no repository definition for @incubator, @ac-charts, @ac-charts, @ac-charts, @ac-charts. Please add them via 'helm repo add'
```

### 41. `shini4i/charts`

- Chart: `D:\helm_clones_github\shini4i__charts\charts\app`
- Source: `dependency`

```text
Error: no repository definition for https://shini4i.github.io/charts/. Please add the missing repos via 'helm repo add'
```

### 42. `tetratelabs/charts`

- Chart: `D:\helm_clones_github\tetratelabs__charts\charts\demos\istio-monitoring-demo`
- Source: `dependency`

```text
Error: no repository definition for https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### 43. `yunzck8s/cloudNative`

- Chart: `D:\helm_clones_github\yunzck8s__cloudNative\charts\deepflow`
- Source: `dependency`

```text
Error: no repository definition for https://grafana.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### 44. `Frndo1203/stack_iceberg_starrocks_k8s`

- Chart: `D:\helm_clones_github\Frndo1203__stack_iceberg_starrocks_k8s\infra\src\helm-charts\airflow`
- Source: `dependency`

```text
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### 45. `JonasHess/homelab-iac`

- Chart: `D:\helm_clones_github\JonasHess__homelab-iac\apps\nextcloud`
- Source: `dependency`

```text
Error: no repository definition for https://nextcloud.github.io/helm/. Please add the missing repos via 'helm repo add'
```

### 46. `Kapil-Bhalodiya/E-Commerce`

- Chart: `D:\helm_clones_github\Kapil-Bhalodiya__E-Commerce\infra\addons\nginx-ingress`
- Source: `dependency`

```text
Error: no repository definition for https://kubernetes.github.io/ingress-nginx. Please add the missing repos via 'helm repo add'
```

### 47. `PilotDataPlatform/helm-charts`

- Chart: `D:\helm_clones_github\PilotDataPlatform__helm-charts\argo-cd-917`
- Source: `dependency`

```text
Error: no repository definition for https://dandydeveloper.github.io/charts/. Please add the missing repos via 'helm repo add'
```

### 48. `SpelmanDevops/retail-store`

- Chart: `D:\helm_clones_github\SpelmanDevops__retail-store\helm\monitoring`
- Source: `dependency`

```text
Error: no repository definition for https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### 49. `TSMC-NYCU-LAB-13/infrastructures`

- Chart: `D:\helm_clones_github\TSMC-NYCU-LAB-13__infrastructures\argo\argo-cd`
- Source: `dependency`

```text
Error: no repository definition for https://dandydeveloper.github.io/charts/. Please add the missing repos via 'helm repo add'
```

### 50. `VadimShtukan/otus_homework`

- Chart: `D:\helm_clones_github\VadimShtukan__otus_homework\architect\lesson05\kubernetis\helm-chart`
- Source: `dependency`

```text
Error: no repository definition for https://kubernetes-charts.storage.googleapis.com. Please add the missing repos via 'helm repo add'
```

### 51. `Vaibhav2goyal/alertmanager`

- Chart: `D:\helm_clones_github\Vaibhav2goyal__alertmanager\scripts\kube-prometheus-stack`
- Source: `dependency`

```text
Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### 52. `WesleyJw/modern-data-stack`

- Chart: `D:\helm_clones_github\WesleyJw__modern-data-stack\infra\src\helm-charts\airflow`
- Source: `dependency`

```text
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### 53. `Yahya-rabii/soge-plus`

- Chart: `D:\helm_clones_github\Yahya-rabii__soge-plus\helm-charts\kube-prometheus-stack`
- Source: `dependency`

```text
Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### 54. `dan1dan12345678/Helm_charts`

- Chart: `D:\helm_clones_github\dan1dan12345678__Helm_charts\kube-prometheus-stack`
- Source: `dependency`

```text
Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### 55. `davidlesicnik/homelab-argo`

- Chart: `D:\helm_clones_github\davidlesicnik__homelab-argo\apps\grafana`
- Source: `dependency`

```text
Error: no repository definition for https://grafana.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### 56. `legion-platform/legion-infrastructure`

- Chart: `D:\helm_clones_github\legion-platform__legion-infrastructure\helms\monitoring`
- Source: `dependency`

```text
Error: no repository definition for @stable, @stable. Please add them via 'helm repo add'
```

### 57. `merlindorin/charts`

- Chart: `D:\helm_clones_github\merlindorin__charts\charts\pinniped`
- Source: `dependency`

```text
Error: no repository definition for https://merlindorin.github.io/charts, https://merlindorin.github.io/charts. Please add the missing repos via 'helm repo add'
```

### 58. `nwthomas/gitops`

- Chart: `D:\helm_clones_github\nwthomas__gitops\helm\longhorn`
- Source: `dependency`

```text
Error: no repository definition for https://charts.longhorn.io. Please add the missing repos via 'helm repo add'
```

### 59. `otus-kuber-2019-12/gidmaster_platform`

- Chart: `D:\helm_clones_github\otus-kuber-2019-12__gidmaster_platform\kubernetes-gitops\deploy\charts\cartservice`
- Source: `dependency`

```text
Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add'
```

### 60. `rtang03/helm-charts`

- Chart: `D:\helm_clones_github\rtang03__helm-charts\charts\argocd`
- Source: `dependency`

```text
Error: no repository definition for https://argoproj.github.io/argo-helm. Please add the missing repos via 'helm repo add'
```

### 61. `tetratelabs/helm-charts`

- Chart: `D:\helm_clones_github\tetratelabs__helm-charts\charts\demos\istio-monitoring-demo`
- Source: `dependency`

```text
Error: no repository definition for https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### 62. `AdmanTIC/helm-charts`

- Chart: `D:\helm_clones_github\AdmanTIC__helm-charts\charts\cremecrm`
- Source: `dependency`

```text
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### 63. `Avichayef/datateam_calculator`

- Chart: `D:\helm_clones_github\Avichayef__datateam_calculator\helm\jenkins`
- Source: `dependency`

```text
Error: no repository definition for https://charts.jenkins.io. Please add the missing repos via 'helm repo add'
```

### 64. `Backstage-Epitech/cltest`

- Chart: `D:\helm_clones_github\Backstage-Epitech__cltest\kustomize\kube-prometheus-stack\charts\kube-prometheus-stack-73.2.2\kube-prometheus-stack`
- Source: `dependency`

```text
Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### 65. `Bernardpro/ClusterGCP`

- Chart: `D:\helm_clones_github\Bernardpro__ClusterGCP\kustomize\kube-prometheus-stack\charts\kube-prometheus-stack-73.2.2\kube-prometheus-stack`
- Source: `dependency`

```text
Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### 66. `Bernardpro/ClusterGKE`

- Chart: `D:\helm_clones_github\Bernardpro__ClusterGKE\kustomize\kube-prometheus-stack\charts\kube-prometheus-stack-73.2.2\kube-prometheus-stack`
- Source: `dependency`

```text
Error: no repository definition for https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

## `template.required_value` (64)

### 1. `trueforge-org/truecharts`

- Chart: `D:\helm_clones_github\trueforge-org__truecharts\charts\stable\app-template`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\trueforge-org__truecharts\charts\stable\app-template`

```text
Error: execution error at (app-template/templates/common.yaml:1:3): Service - Expected non-empty [port.port]

Use --debug flag to render out invalid YAML
```

### 2. `trueforge-org/truecharts`

- Chart: `D:\helm_clones_github\trueforge-org__truecharts\charts\stable\authentik`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\trueforge-org__truecharts\charts\stable\authentik`

```text
Error: execution error at (authentik/templates/common.yaml:97:3): Ingress - Expected ingress [main] to be enabled. This chart is designed to work only with ingress enabled.

Use --debug flag to render out invalid YAML
```

### 3. `trueforge-org/truecharts`

- Chart: `D:\helm_clones_github\trueforge-org__truecharts\charts\stable\nextcloud`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\trueforge-org__truecharts\charts\stable\nextcloud`

```text
Error: execution error at (nextcloud/templates/common.yaml:92:4): Expected non-empty [ip] value on [hostAliases].

Use --debug flag to render out invalid YAML
```

### 4. `trueforge-org/truecharts`

- Chart: `D:\helm_clones_github\trueforge-org__truecharts\charts\stable\slink`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\trueforge-org__truecharts\charts\stable\slink`

```text
Error: execution error at (slink/templates/common.yaml:1:3): Ingress - Expected ingress [main] to be enabled. This chart is designed to work only with ingress enabled.

Use --debug flag to render out invalid YAML
```

### 5. `trueforge-org/truecharts`

- Chart: `D:\helm_clones_github\trueforge-org__truecharts\charts\stable\vaultwarden`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\trueforge-org__truecharts\charts\stable\vaultwarden`

```text
Error: execution error at (vaultwarden/templates/common.yaml:17:3): Ingress - Expected ingress [main] to be enabled. This chart is designed to work only with ingress enabled.

Use --debug flag to render out invalid YAML
```

### 6. `trueforge-org/truecharts`

- Chart: `D:\helm_clones_github\trueforge-org__truecharts\charts\stable\wireguard`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\trueforge-org__truecharts\charts\stable\wireguard`

```text
Error: execution error at (wireguard/templates/common.yaml:19:3): Volumes - Expected the key [enabled] in [persistence.configfile] to exist

Use --debug flag to render out invalid YAML
```

### 7. `BigKAA/youtube`

- Chart: `D:\helm_clones_github\BigKAA__youtube\monitoring\charts\vmalert`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\BigKAA__youtube\monitoring\charts\vmalert`

```text
Error: execution error at (vmalert/charts/victoria-metrics-alert/templates/server-deployment.yaml:4:4): at least one item in `.server.config.alerts.groups` or `.server.extraArgs.rule` must be set 

Use --debug flag to render out invalid YAML
```

### 8. `JahstreetOrg/spark-on-kubernetes-helm`

- Chart: `D:\helm_clones_github\JahstreetOrg__spark-on-kubernetes-helm\charts\azure-keyvault-secret-reloader`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\JahstreetOrg__spark-on-kubernetes-helm\charts\azure-keyvault-secret-reloader`

```text
Error: execution error at (azure-keyvault-secret-reloader/templates/secret.yaml:13:19): .Values.servicePrincipal.clientId is required.

Use --debug flag to render out invalid YAML
```

### 9. `open-edge-platform/edge-ai-libraries`

- Chart: `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart`

```text
Error: execution error at (chat-question-and-answer/charts/dataprepPgvector/templates/datapreppgvector-deployment.yaml:57:24): ALLOWED_HOSTS must be set

Use --debug flag to render out invalid YAML
```

### 10. `open-edge-platform/edge-ai-libraries`

- Chart: `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart`
- Source: `template`
- Values files: `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\values_tgi.yaml`
- Command: `helm template test D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart -f D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\values_tgi.yaml`

```text
Error: execution error at (chat-question-and-answer/charts/dataprepPgvector/templates/datapreppgvector-deployment.yaml:57:24): ALLOWED_HOSTS must be set

Use --debug flag to render out invalid YAML
```

### 11. `open-edge-platform/edge-ai-libraries`

- Chart: `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart`
- Source: `template`
- Values files: `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\values_vllm.yaml`
- Command: `helm template test D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart -f D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\values_vllm.yaml`

```text
Error: execution error at (chat-question-and-answer/charts/dataprepPgvector/templates/datapreppgvector-deployment.yaml:57:24): ALLOWED_HOSTS must be set

Use --debug flag to render out invalid YAML
```

### 12. `open-edge-platform/edge-ai-libraries`

- Chart: `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart`
- Source: `template`
- Values files: `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\values_tgi.yaml`, `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\values_vllm.yaml`
- Command: `helm template test D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart -f D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\values_tgi.yaml -f D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\values_vllm.yaml`

```text
Error: execution error at (chat-question-and-answer/charts/dataprepPgvector/templates/datapreppgvector-deployment.yaml:57:24): ALLOWED_HOSTS must be set

Use --debug flag to render out invalid YAML
```

### 13. `open-edge-platform/edge-ai-libraries`

- Chart: `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\subchart\dataprep`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer\chart\subchart\dataprep`

```text
Error: execution error at (dataprepPgvector/templates/datapreppgvector-deployment.yaml:57:24): ALLOWED_HOSTS must be set

Use --debug flag to render out invalid YAML
```

### 14. `open-edge-platform/edge-ai-libraries`

- Chart: `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\video-search-and-summarization\chart`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\video-search-and-summarization\chart`

```text
Error: execution error at (video-search-and-summarization/templates/pipeline-manager-deployment.yaml:124:24): Value for MINIO_ROOT_USER is required and cannot be empty!

Use --debug flag to render out invalid YAML
```

### 15. `open-edge-platform/edge-ai-libraries`

- Chart: `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\video-search-and-summarization\chart`
- Source: `template`
- Values files: `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\video-search-and-summarization\chart\user_values_override.yaml`
- Command: `helm template test D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\video-search-and-summarization\chart -f D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\video-search-and-summarization\chart\user_values_override.yaml`

```text
Error: execution error at (video-search-and-summarization/templates/pipeline-manager-deployment.yaml:124:24): Value for MINIO_ROOT_USER is required and cannot be empty!

Use --debug flag to render out invalid YAML
```

### 16. `open-edge-platform/edge-ai-libraries`

- Chart: `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\video-search-and-summarization\chart`
- Source: `template`
- Values files: `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\video-search-and-summarization\chart\xeon_vllm_values.yaml`
- Command: `helm template test D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\video-search-and-summarization\chart -f D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\video-search-and-summarization\chart\xeon_vllm_values.yaml`

```text
Error: execution error at (video-search-and-summarization/templates/pipeline-manager-deployment.yaml:124:24): Value for MINIO_ROOT_USER is required and cannot be empty!

Use --debug flag to render out invalid YAML
```

### 17. `open-edge-platform/edge-ai-libraries`

- Chart: `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\video-search-and-summarization\chart`
- Source: `template`
- Values files: `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\video-search-and-summarization\chart\user_values_override.yaml`, `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\video-search-and-summarization\chart\xeon_vllm_values.yaml`
- Command: `helm template test D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\video-search-and-summarization\chart -f D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\video-search-and-summarization\chart\user_values_override.yaml -f D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\video-search-and-summarization\chart\xeon_vllm_values.yaml`

```text
Error: execution error at (video-search-and-summarization/templates/pipeline-manager-deployment.yaml:124:24): Value for MINIO_ROOT_USER is required and cannot be empty!

Use --debug flag to render out invalid YAML
```

### 18. `open-edge-platform/edge-ai-libraries`

- Chart: `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\video-search-and-summarization\chart\subchart\ovms`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\video-search-and-summarization\chart\subchart\ovms`

```text
Error: execution error at (ovms/templates/ovms-deployment.yaml:135:20): Value for `global.vlmName` is required!

Use --debug flag to render out invalid YAML
```

### 19. `open-edge-platform/edge-ai-libraries`

- Chart: `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\video-search-and-summarization\chart\subchart\video-ingestion`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\video-search-and-summarization\chart\subchart\video-ingestion`

```text
Error: execution error at (videoingestion/templates/video-ingestion-deployment.yaml:45:24): Value for odModelName is required!

Use --debug flag to render out invalid YAML
```

### 20. `open-edge-platform/edge-ai-libraries`

- Chart: `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\video-search-and-summarization\chart\subchart\video-search`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\video-search-and-summarization\chart\subchart\video-search`

```text
Error: execution error at (videosearch/templates/video-search-deployment.yaml:65:24): global.embeddingModelName is required

Use --debug flag to render out invalid YAML
```

### 21. `aws-samples/amazon-eks-machine-learning-with-terraform-and-kubeflow`

- Chart: `D:\helm_clones_github\aws-samples__amazon-eks-machine-learning-with-terraform-and-kubeflow\charts\machine-learning\agentic\kagent-agent`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\aws-samples__amazon-eks-machine-learning-with-terraform-and-kubeflow\charts\machine-learning\agentic\kagent-agent`

```text
Error: execution error at (kagent-agent/templates/agent.yaml:1:4): name is required

Use --debug flag to render out invalid YAML
```

### 22. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\blueprints\finetuning_service\src\api\helm-charts\finetuning-api`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\blueprints\finetuning_service\src\api\helm-charts\finetuning-api`

```text
Error: execution error at (finetuning-service/templates/secret.yaml:9:19): Database URL is required

Use --debug flag to render out invalid YAML
```

### 23. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\blueprints\finetuning_service\src\api\helm-charts\postgresql`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\blueprints\finetuning_service\src\api\helm-charts\postgresql`

```text
Error: execution error at (postgres/templates/secret.yaml:12:24): PostgreSQL password is required

Use --debug flag to render out invalid YAML
```

### 24. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\observability\logs-stack`
- Source: `template`
- Values files: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\observability\logs-stack\aws-s3-values.yaml`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\observability\logs-stack -f D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\observability\logs-stack\aws-s3-values.yaml`

```text
Error: execution error at (logs/charts/loki/templates/write/statefulset-write.yaml:50:28): Please define loki.storage.bucketNames.chunks

Use --debug flag to render out invalid YAML
```

### 25. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-alation\charts\mcp-server-alation`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-alation\charts\mcp-server-alation`

```text
Error: execution error at (mcp-server-alation/templates/deployment.yaml:43:24): env.ALATION_BASE_URL is required

Use --debug flag to render out invalid YAML
```

### 26. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-aws-aurora-dsql\charts\mcp-server-aws-aurora-dsql`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-aws-aurora-dsql\charts\mcp-server-aws-aurora-dsql`

```text
Error: execution error at (mcp-server-aws-aurora-dsql/templates/deployment.yaml:50:24): env.CLUSTER_ENDPOINT is required

Use --debug flag to render out invalid YAML
```

### 27. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-aws-bedrock-data-automation\charts\mcp-server-aws-bedrock-data-automation`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-aws-bedrock-data-automation\charts\mcp-server-aws-bedrock-data-automation`

```text
Error: execution error at (mcp-server-aws-bedrock-data-automation/templates/deployment.yaml:47:24): env.AWS_BUCKET_NAME is required

Use --debug flag to render out invalid YAML
```

### 28. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-aws-kendra-index\charts\mcp-server-aws-kendra-index`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-aws-kendra-index\charts\mcp-server-aws-kendra-index`

```text
Error: execution error at (mcp-server-aws-kendra-index/templates/deployment.yaml:43:24): env.KENDRA_INDEX_ID is required

Use --debug flag to render out invalid YAML
```

### 29. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-aws-memcached\charts\mcp-server-aws-memcached`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-aws-memcached\charts\mcp-server-aws-memcached`

```text
Error: execution error at (mcp-server-aws-memcached/templates/deployment.yaml:43:24): env.MEMCACHED_HOST is required

Use --debug flag to render out invalid YAML
```

### 30. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-aws-mysql\charts\mcp-server-aws-mysql`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-aws-mysql\charts\mcp-server-aws-mysql`

```text
Error: execution error at (mcp-server-aws-mysql/templates/deployment.yaml:51:24): env.RESOURCE_ARN is required

Use --debug flag to render out invalid YAML
```

### 31. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-aws-neptune\charts\mcp-server-aws-neptune`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-aws-neptune\charts\mcp-server-aws-neptune`

```text
Error: execution error at (mcp-server-aws-neptune/templates/deployment.yaml:43:24): env.NEPTUNE_ENDPOINT is required

Use --debug flag to render out invalid YAML
```

### 32. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-aws-postgres\charts\mcp-server-aws-postgres`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-aws-postgres\charts\mcp-server-aws-postgres`

```text
Error: execution error at (mcp-server-aws-postgres/templates/deployment.yaml:51:24): env.RESOURCE_ARN is required

Use --debug flag to render out invalid YAML
```

### 33. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-aws-prometheus\charts\mcp-server-aws-prometheus`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-aws-prometheus\charts\mcp-server-aws-prometheus`

```text
Error: execution error at (mcp-server-aws-prometheus/templates/deployment.yaml:43:24): env.PROMETHEUS_URL is required

Use --debug flag to render out invalid YAML
```

### 34. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-azure-devops\charts\mcp-server-azure-devops`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-azure-devops\charts\mcp-server-azure-devops`

```text
Error: execution error at (mcp-server-azure-devops/templates/deployment.yaml:43:24): env.AZURE_DEVOPS_AUTH_METHOD is required

Use --debug flag to render out invalid YAML
```

### 35. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-chroma\charts\mcp-server-chroma`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-chroma\charts\mcp-server-chroma`

```text
Error: execution error at (mcp-server-chroma/templates/deployment.yaml:44:24): env.CHROMA_CLIENT_TYPE is required

Use --debug flag to render out invalid YAML
```

### 36. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-google-drive\charts\mcp-server-google-drive`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-google-drive\charts\mcp-server-google-drive`

```text
Error: execution error at (mcp-server-google-drive/templates/deployment.yaml:43:24): env.GDRIVE_CREDENTIALS_PATH is required

Use --debug flag to render out invalid YAML
```

### 37. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-jetbrains\charts\mcp-server-jetbrains`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-jetbrains\charts\mcp-server-jetbrains`

```text
Error: execution error at (mcp-server-jetbrains/templates/deployment.yaml:43:24): env.IDE_PORT is required

Use --debug flag to render out invalid YAML
```

### 38. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-multiversx\charts\mcp-server-multiversx`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-multiversx\charts\mcp-server-multiversx`

```text
Error: execution error at (mcp-server-multiversx/templates/deployment.yaml:43:24): env.MVX_NETWORK is required

Use --debug flag to render out invalid YAML
```

### 39. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-openapi-schema\charts\mcp-server-openapi-schema`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-openapi-schema\charts\mcp-server-openapi-schema`

```text
Error: execution error at (mcp-server-openapi-schema/templates/deployment.yaml:43:24): env.SCHEMA_PATH is required

Use --debug flag to render out invalid YAML
```

### 40. `acuvity/mcp-servers-registry`

- Chart: `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-phoenix\charts\mcp-server-phoenix`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-phoenix\charts\mcp-server-phoenix`

```text
Error: execution error at (mcp-server-phoenix/templates/deployment.yaml:43:24): env.PHOENIX_BASE_URL is required

Use --debug flag to render out invalid YAML
```

### 41. `openkruise/charts`

- Chart: `D:\helm_clones_github\openkruise__charts\versions\kruise-agents-sandbox-manager\0.2.0`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\openkruise__charts\versions\kruise-agents-sandbox-manager\0.2.0`

```text
Error: execution error at (agents-sandbox-manager/templates/ingress.yaml:13:23): ingress.className is required

Use --debug flag to render out invalid YAML
```

### 42. `openkruise/charts`

- Chart: `D:\helm_clones_github\openkruise__charts\versions\kruise-agents-sandbox-manager\0.2.0-rc1`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\openkruise__charts\versions\kruise-agents-sandbox-manager\0.2.0-rc1`

```text
Error: execution error at (agents-sandbox-manager/templates/ingress.yaml:13:23): ingress.className is required

Use --debug flag to render out invalid YAML
```

### 43. `openkruise/charts`

- Chart: `D:\helm_clones_github\openkruise__charts\versions\kruise-agents-sandbox-manager\0.3.0`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\openkruise__charts\versions\kruise-agents-sandbox-manager\0.3.0`

```text
Error: execution error at (agents-sandbox-manager/templates/ingress.yaml:13:23): ingress.className is required

Use --debug flag to render out invalid YAML
```

### 44. `openkruise/charts`

- Chart: `D:\helm_clones_github\openkruise__charts\versions\kruise-agents-sandbox-manager\next`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\openkruise__charts\versions\kruise-agents-sandbox-manager\next`

```text
Error: execution error at (agents-sandbox-manager/templates/ingress.yaml:13:23): ingress.className is required

Use --debug flag to render out invalid YAML
```

### 45. `LeoShivas/GitOps`

- Chart: `D:\helm_clones_github\LeoShivas__GitOps\kubernetes\transmission\helm\transmission`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\LeoShivas__GitOps\kubernetes\transmission\helm\transmission`

```text
Error: execution error at (transmission/templates/deployment.yaml:94:26): An existing claim name for storing data is required !

Use --debug flag to render out invalid YAML
```

### 46. `celo-org/charts`

- Chart: `D:\helm_clones_github\celo-org__charts\charts\sync-test`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\celo-org__charts\charts\sync-test`

```text
Error: execution error at (sync-test/templates/NOTES.txt:1:4): op-geth.secrets.nodeKey.value is required: --set op-geth.secrets.nodeKey.value=0x$(openssl rand -hex 32)

Use --debug flag to render out invalid YAML
```

### 47. `pascalnaber/ignite-tour-k8s-lessons-learned`

- Chart: `D:\helm_clones_github\pascalnaber__ignite-tour-k8s-lessons-learned\app-keyvault\provisioning\aad-pod-identity`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\pascalnaber__ignite-tour-k8s-lessons-learned\app-keyvault\provisioning\aad-pod-identity`

```text
Error: execution error at (aad-pod-identity/templates/identities.yaml:8:15): .Values.azureIdentity.resourceID is required!

Use --debug flag to render out invalid YAML
```

### 48. `pascalnaber/ignite-tour-k8s-lessons-learned`

- Chart: `D:\helm_clones_github\pascalnaber__ignite-tour-k8s-lessons-learned\provisioning\scripts\configuration\aad-pod-identity`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\pascalnaber__ignite-tour-k8s-lessons-learned\provisioning\scripts\configuration\aad-pod-identity`

```text
Error: execution error at (aad-pod-identity/templates/identities.yaml:8:15): .Values.azureIdentity.resourceID is required!

Use --debug flag to render out invalid YAML
```

### 49. `bakseter/whpah`

- Chart: `D:\helm_clones_github\bakseter__whpah\manifests\cluster-addons\monitoring`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\bakseter__whpah\manifests\cluster-addons\monitoring`

```text
Error: execution error at (monitoring-umbrella/charts/loki/templates/validate.yaml:46:4): Please define loki.storage.bucketName.chunks

Use --debug flag to render out invalid YAML
```

### 50. `elasticio/helm-charts`

- Chart: `D:\helm_clones_github\elasticio__helm-charts\cluster`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\elasticio__helm-charts\cluster`

```text
Error: execution error at (cluster/charts/platform-storage-slugs/templates/service-loadbalancer.yaml:15:21): You must provide load balancer IP for slugs storage

Use --debug flag to render out invalid YAML
```

### 51. `elasticio/helm-charts`

- Chart: `D:\helm_clones_github\elasticio__helm-charts\gitreceiver`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\elasticio__helm-charts\gitreceiver`

```text
Error: execution error at (gitreceiver/templates/service-loadbalancer.yaml:13:21): You need provide load balancer IP for gitreceiver

Use --debug flag to render out invalid YAML
```

### 52. `elasticio/helm-charts`

- Chart: `D:\helm_clones_github\elasticio__helm-charts\platform-storage-slugs`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\elasticio__helm-charts\platform-storage-slugs`

```text
Error: execution error at (platform-storage-slugs/templates/service-loadbalancer.yaml:15:21): You must provide load balancer IP for slugs storage

Use --debug flag to render out invalid YAML
```

### 53. `thelande/charts`

- Chart: `D:\helm_clones_github\thelande__charts\charts\apcupsd-exporter`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\thelande__charts\charts\apcupsd-exporter`

```text
Error: execution error at (apcupsd-exporter/templates/common.yml:81:34): A valid apcupsd target is required.

Use --debug flag to render out invalid YAML
```

### 54. `thelande/charts`

- Chart: `D:\helm_clones_github\thelande__charts\charts\bar-assistant`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\thelande__charts\charts\bar-assistant`

```text
Error: execution error at (bar-assistant/templates/statefulset.yaml:94:25): baseUrl must be set

Use --debug flag to render out invalid YAML
```

### 55. `thelande/charts`

- Chart: `D:\helm_clones_github\thelande__charts\charts\docmost`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\thelande__charts\charts\docmost`

```text
Error: execution error at (docmost/templates/deployment.yaml:75:25): An app URL is required

Use --debug flag to render out invalid YAML
```

### 56. `CSCfi/helm-charts`

- Chart: `D:\helm_clones_github\CSCfi__helm-charts\charts\hedgedoc`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\CSCfi__helm-charts\charts\hedgedoc`

```text
Error: execution error at (hedgedoc/templates/route.yaml:9:42): A valid domain must be set

Use --debug flag to render out invalid YAML
```

### 57. `CSCfi/helm-charts`

- Chart: `D:\helm_clones_github\CSCfi__helm-charts\charts\mattermost`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\CSCfi__helm-charts\charts\mattermost`

```text
Error: execution error at (mattermost/templates/route.yaml:9:11): A value for 'route.host' is required.

Use --debug flag to render out invalid YAML
```

### 58. `CSCfi/helm-charts`

- Chart: `D:\helm_clones_github\CSCfi__helm-charts\charts\minio`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\CSCfi__helm-charts\charts\minio`

```text
Error: execution error at (minio/templates/route.yaml:10:43): A valid domainSuffix is required

Use --debug flag to render out invalid YAML
```

### 59. `CSCfi/helm-charts`

- Chart: `D:\helm_clones_github\CSCfi__helm-charts\charts\rocketchat`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\CSCfi__helm-charts\charts\rocketchat`

```text
Error: execution error at (rocketchat/templates/route.yaml:12:36): A value is required for host

Use --debug flag to render out invalid YAML
```

### 60. `CSCfi/helm-charts`

- Chart: `D:\helm_clones_github\CSCfi__helm-charts\charts\rstudio`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\CSCfi__helm-charts\charts\rstudio`

```text
Error: execution error at (rstudio/templates/routes.yaml:9:11): A valid .Values.rstudio.route.host entry is required!

Use --debug flag to render out invalid YAML
```

### 61. `JonasHess/homelab-iac`

- Chart: `D:\helm_clones_github\JonasHess__homelab-iac\apps\envoy-gateway`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\JonasHess__homelab-iac\apps\envoy-gateway`

```text
Error: execution error at (envoy-gateway/templates/oidc-security-policy.yaml:21:15): global.domain must be set

Use --debug flag to render out invalid YAML
```

### 62. `codefuturist/helm-charts`

- Chart: `D:\helm_clones_github\codefuturist__helm-charts\charts\apps\compass-web`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\codefuturist__helm-charts\charts\apps\compass-web`

```text
Error: execution error at (compass-web/templates/statefulset.yaml:1:4): Either compassWeb.mongoUri or compassWeb.existingSecret must be set

Use --debug flag to render out invalid YAML
```

### 63. `merlindorin/charts`

- Chart: `D:\helm_clones_github\merlindorin__charts\charts\exporter-unifi-protect`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\merlindorin__charts\charts\exporter-unifi-protect`

```text
Error: execution error at (exporter-unifi-protect/templates/secret.yaml:10:59): unifiProtect.username is required when not using existingSecret

Use --debug flag to render out invalid YAML
```

### 64. `AlexanderBabel/helm-charts`

- Chart: `D:\helm_clones_github\AlexanderBabel__helm-charts\charts\limesurvey`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\AlexanderBabel__helm-charts\charts\limesurvey`

```text
Error: execution error at (limesurvey/templates/secrets.yaml:45:18): externalDatabase.password is required

Use --debug flag to render out invalid YAML
```

## `template.library_chart_not_installable` (39)

### 1. `bitnami/charts`

- Chart: `D:\helm_clones_github\bitnami__charts\bitnami\common`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\bitnami__charts\bitnami\common`

```text
Error: library charts are not installable
```

### 2. `trueforge-org/truecharts`

- Chart: `D:\helm_clones_github\trueforge-org__truecharts\charts\library\common`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\trueforge-org__truecharts\charts\library\common`

```text
Error: library charts are not installable
```

### 3. `open-edge-platform/geti`

- Chart: `D:\helm_clones_github\open-edge-platform__geti\deploy\charts\control-plane\chart\charts\control-plane-common`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\open-edge-platform__geti\deploy\charts\control-plane\chart\charts\control-plane-common`

```text
Error: library charts are not installable
```

### 4. `open-edge-platform/geti`

- Chart: `D:\helm_clones_github\open-edge-platform__geti\deploy\charts\impt\chart\charts\geti-common`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\open-edge-platform__geti\deploy\charts\impt\chart\charts\geti-common`

```text
Error: library charts are not installable
```

### 5. `open-edge-platform/geti`

- Chart: `D:\helm_clones_github\open-edge-platform__geti\deploy\charts\impt\chart\charts\geti-common-labels`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\open-edge-platform__geti\deploy\charts\impt\chart\charts\geti-common-labels`

```text
Error: library charts are not installable
```

### 6. `apecloud/kubeblocks-addons`

- Chart: `D:\helm_clones_github\apecloud__kubeblocks-addons\addons\kblib`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\apecloud__kubeblocks-addons\addons\kblib`

```text
Error: library charts are not installable
```

### 7. `apecloud/kubeblocks-addons`

- Chart: `D:\helm_clones_github\apecloud__kubeblocks-addons\addons-cluster\kblib`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\apecloud__kubeblocks-addons\addons-cluster\kblib`

```text
Error: library charts are not installable
```

### 8. `boozallen/aissemble`

- Chart: `D:\helm_clones_github\boozallen__aissemble\extensions\extensions-helm\extensions-helm-pipeline-invocation\extensions-helm-pipeline-invocation-lib\aissemble-pipeline-invocation-lib-src-chart`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\boozallen__aissemble\extensions\extensions-helm\extensions-helm-pipeline-invocation\extensions-helm-pipeline-invocation-lib\aissemble-pipeline-invocation-lib-src-chart`

```text
Error: library charts are not installable
```

### 9. `sapcc/helm-charts`

- Chart: `D:\helm_clones_github\sapcc__helm-charts\common\helm3-helper`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\sapcc__helm-charts\common\helm3-helper`

```text
Error: library charts are not installable
```

### 10. `k8s-home-lab/helm-charts`

- Chart: `D:\helm_clones_github\k8s-home-lab__helm-charts\charts\stable\common`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\k8s-home-lab__helm-charts\charts\stable\common`

```text
Error: library charts are not installable
```

### 11. `bflance/proxmox-talos`

- Chart: `D:\helm_clones_github\bflance__proxmox-talos\charts\grafana-loki\charts\common`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\bflance__proxmox-talos\charts\grafana-loki\charts\common`

```text
Error: library charts are not installable
```

### 12. `bflance/proxmox-talos`

- Chart: `D:\helm_clones_github\bflance__proxmox-talos\charts\grafana-loki\charts\memcached\charts\common`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\bflance__proxmox-talos\charts\grafana-loki\charts\memcached\charts\common`

```text
Error: library charts are not installable
```

### 13. `llajas/homelab`

- Chart: `D:\helm_clones_github\llajas__homelab\apps\clusterplex\charts\common`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\llajas__homelab\apps\clusterplex\charts\common`

```text
Error: library charts are not installable
```

### 14. `SAP/component-operator-runtime`

- Chart: `D:\helm_clones_github\SAP__component-operator-runtime\internal\helm\testdata\main\charts\lib11`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\SAP__component-operator-runtime\internal\helm\testdata\main\charts\lib11`

```text
Error: library charts are not installable
```

### 15. `celo-org/charts`

- Chart: `D:\helm_clones_github\celo-org__charts\charts\common`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\celo-org__charts\charts\common`

```text
Error: library charts are not installable
```

### 16. `general-rj45/mini-ml-stand-GRJ45`

- Chart: `D:\helm_clones_github\general-rj45__mini-ml-stand-GRJ45\airflow-14.2.5\airflow\charts\common`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\general-rj45__mini-ml-stand-GRJ45\airflow-14.2.5\airflow\charts\common`

```text
Error: library charts are not installable
```

### 17. `general-rj45/mini-ml-stand-GRJ45`

- Chart: `D:\helm_clones_github\general-rj45__mini-ml-stand-GRJ45\airflow-14.2.5\airflow\charts\postgresql\charts\common`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\general-rj45__mini-ml-stand-GRJ45\airflow-14.2.5\airflow\charts\postgresql\charts\common`

```text
Error: library charts are not installable
```

### 18. `general-rj45/mini-ml-stand-GRJ45`

- Chart: `D:\helm_clones_github\general-rj45__mini-ml-stand-GRJ45\airflow-14.2.5\airflow\charts\redis\charts\common`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\general-rj45__mini-ml-stand-GRJ45\airflow-14.2.5\airflow\charts\redis\charts\common`

```text
Error: library charts are not installable
```

### 19. `general-rj45/mini-ml-stand-GRJ45`

- Chart: `D:\helm_clones_github\general-rj45__mini-ml-stand-GRJ45\gitea-8.3.0\gitea\charts\memcached\charts\common`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\general-rj45__mini-ml-stand-GRJ45\gitea-8.3.0\gitea\charts\memcached\charts\common`

```text
Error: library charts are not installable
```

### 20. `general-rj45/mini-ml-stand-GRJ45`

- Chart: `D:\helm_clones_github\general-rj45__mini-ml-stand-GRJ45\gitea-8.3.0\gitea\charts\postgresql\charts\common`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\general-rj45__mini-ml-stand-GRJ45\gitea-8.3.0\gitea\charts\postgresql\charts\common`

```text
Error: library charts are not installable
```

### 21. `general-rj45/mini-ml-stand-GRJ45`

- Chart: `D:\helm_clones_github\general-rj45__mini-ml-stand-GRJ45\minio-12.6.4\minio\charts\common`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\general-rj45__mini-ml-stand-GRJ45\minio-12.6.4\minio\charts\common`

```text
Error: library charts are not installable
```

### 22. `general-rj45/mini-ml-stand-GRJ45`

- Chart: `D:\helm_clones_github\general-rj45__mini-ml-stand-GRJ45\mlflow\charts\postgresql\charts\common`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\general-rj45__mini-ml-stand-GRJ45\mlflow\charts\postgresql\charts\common`

```text
Error: library charts are not installable
```

### 23. `kast-spells/kast-system`

- Chart: `D:\helm_clones_github\kast-spells__kast-system\charts\glyphs\s3`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\kast-spells__kast-system\charts\glyphs\s3`

```text
Error: library charts are not installable
```

### 24. `yapily/helm-charts`

- Chart: `D:\helm_clones_github\yapily__helm-charts\charts\base`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\yapily__helm-charts\charts\base`

```text
Error: library charts are not installable
```

### 25. `erost/vdz26-demo-fleet-commander`

- Chart: `D:\helm_clones_github\erost__vdz26-demo-fleet-commander\common\generic-composition-chart`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\erost__vdz26-demo-fleet-commander\common\generic-composition-chart`

```text
Error: library charts are not installable
```

### 26. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\rook-1.10.9\charts\library`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\rook-1.10.9\charts\library`

```text
Error: library charts are not installable
```

### 27. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\rook-1.11.6\charts\library`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\rook-1.11.6\charts\library`

```text
Error: library charts are not installable
```

### 28. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\rook-1.12.3\charts\library`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\rook-1.12.3\charts\library`

```text
Error: library charts are not installable
```

### 29. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\rook-1.13.10\charts\library`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\rook-1.13.10\charts\library`

```text
Error: library charts are not installable
```

### 30. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\rook-1.14.12\charts\library`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\rook-1.14.12\charts\library`

```text
Error: library charts are not installable
```

### 31. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\rook-1.15.9\charts\library`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\rook-1.15.9\charts\library`

```text
Error: library charts are not installable
```

### 32. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\rook-1.16.6\charts\library`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\rook-1.16.6\charts\library`

```text
Error: library charts are not installable
```

### 33. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\rook-1.17.7\charts\library`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\rook-1.17.7\charts\library`

```text
Error: library charts are not installable
```

### 34. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\rook-1.18.0\charts\library`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\rook-1.18.0\charts\library`

```text
Error: library charts are not installable
```

### 35. `Coollision/homelab-config`

- Chart: `D:\helm_clones_github\Coollision__homelab-config\lib\longhorn-storage-lib`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Coollision__homelab-config\lib\longhorn-storage-lib`

```text
Error: library charts are not installable
```

### 36. `Coollision/homelab-config`

- Chart: `D:\helm_clones_github\Coollision__homelab-config\lib\shared-lib`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Coollision__homelab-config\lib\shared-lib`

```text
Error: library charts are not installable
```

### 37. `codefuturist/helm-charts`

- Chart: `D:\helm_clones_github\codefuturist__helm-charts\charts\libs\common`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\codefuturist__helm-charts\charts\libs\common`

```text
Error: library charts are not installable
```

### 38. `kast-spells/kaster`

- Chart: `D:\helm_clones_github\kast-spells__kaster\charts\s3`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\kast-spells__kaster\charts\s3`

```text
Error: library charts are not installable
```

### 39. `kast-spells/summon`

- Chart: `D:\helm_clones_github\kast-spells__summon\charts\s3`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\kast-spells__summon\charts\s3`

```text
Error: library charts are not installable
```

## `template.values_schema_validation` (37)

### 1. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\fusionauth`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\fusionauth`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
fusionauth:
- at '': missing properties 'replicaCount', 'image', 'imagePullSecrets', 'initImage', 'nameOverride', 'fullnameOverride', 'service', 'database', 'search', 'app', 'environment', 'kickstart', 'podDisruptionBudget', 'resources', 'nodeSelector', 'tolerations', 'affinity', 'dnsConfig', 'dnsPolicy', 'annotations', 'podAnnotations', 'livenessProbe', 'readinessProbe', 'startupProbe'
```

### 2. `dynatrace-wwse/enablement-kubernetes-opentelemetry`

- Chart: `D:\helm_clones_github\dynatrace-wwse__enablement-kubernetes-opentelemetry\cluster-manifests\istio-1.22.1\manifests\charts\gateway`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\dynatrace-wwse__enablement-kubernetes-opentelemetry\cluster-manifests\istio-1.22.1\manifests\charts\gateway`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
gateway:
- at '': additional properties 'defaults' not allowed
```

### 3. `opspresso/argocd-env-addons`

- Chart: `D:\helm_clones_github\opspresso__argocd-env-addons\charts\atlantis`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\opspresso__argocd-env-addons\charts\atlantis`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
atlantis:
"file:///values.schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2019-09/schema#'
- at '': 'allOf' failed
  - at '/properties/extraArgs': 'allOf' failed
    - at '/properties/extraArgs/examples': got object, want array
```

### 4. `kamu-data/helm-charts`

- Chart: `D:\helm_clones_github\kamu-data__helm-charts\charts\kamu-oracle-provider`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\kamu-data__helm-charts\charts\kamu-oracle-provider`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
kamu-oracle-provider:
- at '/app/config': missing properties 'oracleContractAddress', 'providerAddress', 'providerPrivateKey', 'transactionConfirmations'
```

### 5. `kamu-data/helm-charts`

- Chart: `D:\helm_clones_github\kamu-data__helm-charts\charts\kamu-web-ui`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\kamu-data__helm-charts\charts\kamu-web-ui`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
kamu-web-ui:
- at '/app/config/apiServerGqlUrl': '' is not valid uri: relative url
- at '/app/config/apiServerHttpUrl': '' is not valid uri: relative url
- at '/app/config/githubClientId': minLength: got 0, want 1
```

### 6. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\istio-gateway-1.22.6`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\istio-gateway-1.22.6`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
istio-gateway:
- at '': additional properties 'defaults' not allowed
```

### 7. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\istio-gateway-1.22.8`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\istio-gateway-1.22.8`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
istio-gateway:
- at '': additional properties 'defaults' not allowed
```

### 8. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\istio-gateway-1.24.1`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\istio-gateway-1.24.1`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
istio-gateway:
- at '': additional properties '_internal_defaults_do_not_set' not allowed
```

### 9. `oracle-cne/catalog`

- Chart: `D:\helm_clones_github\oracle-cne__catalog\charts\istio-gateway-1.24.6`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oracle-cne__catalog\charts\istio-gateway-1.24.6`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
istio-gateway:
- at '': additional properties '_internal_defaults_do_not_set' not allowed
```

### 10. `CSCfi/helm-charts`

- Chart: `D:\helm_clones_github\CSCfi__helm-charts\charts\airflow`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\CSCfi__helm-charts\charts\airflow`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
airflow:
- at '/airflow/pgbouncer/uid': got number, want null
- at '/airflow/gid': got number, want null
- at '/airflow/uid': got number, want null
- at '/airflow/redis/uid': got number, want null
```

### 11. `CSCfi/helm-charts`

- Chart: `D:\helm_clones_github\CSCfi__helm-charts\charts\prometheus-grafana`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\CSCfi__helm-charts\charts\prometheus-grafana`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
prometheus-grafana-helm:
- at '/prometheus/serverFiles/prometheus.yml/scrape_configs/1/kubernetes_sd_configs/0/namespaces/names/0': got null, want string
```

### 12. `JonasHess/homelab-iac`

- Chart: `D:\helm_clones_github\JonasHess__homelab-iac\apps\adguard`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\JonasHess__homelab-iac\apps\adguard`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
generic:
- at '': missing property 'appName'
- at '/deployment/pvcMounts': validation failed
  - at '/deployment/pvcMounts/config': missing property 'hostPath'
  - at '/deployment/pvcMounts/data': missing property 'hostPath'
```

### 13. `JonasHess/homelab-iac`

- Chart: `D:\helm_clones_github\JonasHess__homelab-iac\apps\alexa-custom-skill`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\JonasHess__homelab-iac\apps\alexa-custom-skill`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
generic:
- at '': missing property 'appName'
```

### 14. `JonasHess/homelab-iac`

- Chart: `D:\helm_clones_github\JonasHess__homelab-iac\apps\alexa-smarthome-skill`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\JonasHess__homelab-iac\apps\alexa-smarthome-skill`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
generic:
- at '': missing property 'appName'
```

### 15. `JonasHess/homelab-iac`

- Chart: `D:\helm_clones_github\JonasHess__homelab-iac\apps\asn`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\JonasHess__homelab-iac\apps\asn`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
generic:
- at '': missing property 'appName'
```

### 16. `JonasHess/homelab-iac`

- Chart: `D:\helm_clones_github\JonasHess__homelab-iac\apps\audiobookshelf`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\JonasHess__homelab-iac\apps\audiobookshelf`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
generic:
- at '': missing property 'appName'
- at '/deployment/pvcMounts': validation failed
  - at '/deployment/pvcMounts/metadata': missing property 'hostPath'
  - at '/deployment/pvcMounts/config': missing property 'hostPath'
  - at '/deployment/pvcMounts/media': missing property 'hostPath'
```

### 17. `JonasHess/homelab-iac`

- Chart: `D:\helm_clones_github\JonasHess__homelab-iac\apps\backrest`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\JonasHess__homelab-iac\apps\backrest`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
generic:
- at '': missing property 'appName'
- at '/persistentVolumeClaims': validation failed
  - at '/persistentVolumeClaims/cache': missing property 'hostPath'
  - at '/persistentVolumeClaims/data': missing property 'hostPath'
  - at '/persistentVolumeClaims/restore': missing property 'hostPath'
  - at '/persistentVolumeClaims/tmp': missing property 'hostPath'
```

### 18. `JonasHess/homelab-iac`

- Chart: `D:\helm_clones_github\JonasHess__homelab-iac\apps\cert-manager`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\JonasHess__homelab-iac\apps\cert-manager`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
generic:
- at '': missing property 'appName'
```

### 19. `JonasHess/homelab-iac`

- Chart: `D:\helm_clones_github\JonasHess__homelab-iac\apps\cloudflareddns`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\JonasHess__homelab-iac\apps\cloudflareddns`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
generic:
- at '': missing property 'appName'
```

### 20. `JonasHess/homelab-iac`

- Chart: `D:\helm_clones_github\JonasHess__homelab-iac\apps\duplicatiprometheusexporter`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\JonasHess__homelab-iac\apps\duplicatiprometheusexporter`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
generic:
- at '': missing property 'appName'
```

### 21. `JonasHess/homelab-iac`

- Chart: `D:\helm_clones_github\JonasHess__homelab-iac\apps\filecleanup`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\JonasHess__homelab-iac\apps\filecleanup`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
generic:
- at '': missing property 'appName'
```

### 22. `JonasHess/homelab-iac`

- Chart: `D:\helm_clones_github\JonasHess__homelab-iac\apps\freshrss`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\JonasHess__homelab-iac\apps\freshrss`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
generic:
- at '': missing property 'appName'
- at '/deployment/pvcMounts/data': missing property 'hostPath'
```

### 23. `JonasHess/homelab-iac`

- Chart: `D:\helm_clones_github\JonasHess__homelab-iac\apps\generic`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\JonasHess__homelab-iac\apps\generic`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
generic:
- at '': missing property 'appName'
```

### 24. `JonasHess/homelab-iac`

- Chart: `D:\helm_clones_github\JonasHess__homelab-iac\apps\gotenberg`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\JonasHess__homelab-iac\apps\gotenberg`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
generic:
- at '': missing property 'appName'
```

### 25. `JonasHess/homelab-iac`

- Chart: `D:\helm_clones_github\JonasHess__homelab-iac\apps\homeassistant`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\JonasHess__homelab-iac\apps\homeassistant`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
generic:
- at '': missing property 'appName'
- at '/deployment/pvcMounts/config': missing property 'hostPath'
```

### 26. `JonasHess/homelab-iac`

- Chart: `D:\helm_clones_github\JonasHess__homelab-iac\apps\homematic`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\JonasHess__homelab-iac\apps\homematic`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
generic:
- at '': missing property 'appName'
```

### 27. `JonasHess/homelab-iac`

- Chart: `D:\helm_clones_github\JonasHess__homelab-iac\apps\homer`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\JonasHess__homelab-iac\apps\homer`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
generic:
- at '': missing property 'appName'
```

### 28. `JonasHess/homelab-iac`

- Chart: `D:\helm_clones_github\JonasHess__homelab-iac\apps\immich`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\JonasHess__homelab-iac\apps\immich`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
generic:
- at '': missing property 'appName'
- at '/persistentVolumeClaims': validation failed
  - at '/persistentVolumeClaims/cli': missing property 'hostPath'
  - at '/persistentVolumeClaims/library': missing property 'hostPath'
  - at '/persistentVolumeClaims/postgresql': missing property 'hostPath'
  - at '/persistentVolumeClaims/redis': missing property 'hostPath'
```

### 29. `JonasHess/homelab-iac`

- Chart: `D:\helm_clones_github\JonasHess__homelab-iac\apps\jellyfin`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\JonasHess__homelab-iac\apps\jellyfin`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
generic:
- at '': missing property 'appName'
- at '/deployment/pvcMounts': validation failed
  - at '/deployment/pvcMounts/config': missing property 'hostPath'
  - at '/deployment/pvcMounts/media': missing property 'hostPath'
```

### 30. `JonasHess/homelab-iac`

- Chart: `D:\helm_clones_github\JonasHess__homelab-iac\apps\mealie`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\JonasHess__homelab-iac\apps\mealie`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
generic:
- at '': missing property 'appName'
- at '/deployment/pvcMounts/data': missing property 'hostPath'
```

### 31. `JonasHess/homelab-iac`

- Chart: `D:\helm_clones_github\JonasHess__homelab-iac\apps\n8n`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\JonasHess__homelab-iac\apps\n8n`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
generic:
- at '': missing property 'appName'
- at '/deployment/pvcMounts/data': missing property 'hostPath'
```

### 32. `K-FOSS/CoRE-CNTRL`

- Chart: `D:\helm_clones_github\K-FOSS__CoRE-CNTRL\Archive\CNTRL-Plane`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\K-FOSS__CoRE-CNTRL\Archive\CNTRL-Plane`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
cert-manager:
- at '/startupapicheck/rbac': additional properties 'create' not allowed
```

### 33. `gregorwolf/cap-python`

- Chart: `D:\helm_clones_github\gregorwolf__cap-python\chart\charts\content-deployment`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\gregorwolf__cap-python\chart\charts\content-deployment`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
content-deployment:
- at '': missing property 'bindings'
- at '/image': missing property 'repository'
```

### 34. `gregorwolf/cap-python`

- Chart: `D:\helm_clones_github\gregorwolf__cap-python\chart\charts\service-instance`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\gregorwolf__cap-python\chart\charts\service-instance`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
service-instance:
- at '': missing properties 'serviceOfferingName', 'servicePlanName'
```

### 35. `gregorwolf/cap-python`

- Chart: `D:\helm_clones_github\gregorwolf__cap-python\chart\charts\web-application`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\gregorwolf__cap-python\chart\charts\web-application`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
web-application:
- at '': missing property 'resources'
- at '/image': missing property 'repository'
```

### 36. `100rd/platform-design`

- Chart: `D:\helm_clones_github\100rd__platform-design\apps\infra\cert-manager`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\100rd__platform-design\apps\infra\cert-manager`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
cert-manager:
- at '': additional properties 'priorityClassName', 'logLevel' not allowed
```

### 37. `Amaterassu17/RL_Microservice_Grouped_Scaler`

- Chart: `D:\helm_clones_github\Amaterassu17__RL_Microservice_Grouped_Scaler\Istio_Mesh\istio-1.21.0\manifests\charts\gateway`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Amaterassu17__RL_Microservice_Grouped_Scaler\Istio_Mesh\istio-1.21.0\manifests\charts\gateway`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
gateway:
- at '': additional properties 'defaults' not allowed
```

## `dependency.missing_subchart` (25)

### 1. `grafana/helm-charts`

- Chart: `D:\helm_clones_github\grafana__helm-charts\charts\enterprise-metrics`
- Source: `dependency`

```text
Getting updates for unmanaged Helm repositories...
...Successfully got an update from the "https://helm.min.io/" chart repository
...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository
Error: minio chart not found in repo https://helm.min.io/
```

### 2. `cloudnativeapp/charts`

- Chart: `D:\helm_clones_github\cloudnativeapp__charts\curated\airflow`
- Source: `dependency`

```text
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies
```

### 3. `radondb/radondb-mysql-kubernetes`

- Chart: `D:\helm_clones_github\radondb__radondb-mysql-kubernetes\charts\mysql-operator`
- Source: `dependency`

```text
Error: directory D:\helm_clones_github\radondb__radondb-mysql-kubernetes\charts\mysql-operator\charts\mysqlcluster not found
```

### 4. `IBM/charts`

- Chart: `D:\helm_clones_github\IBM__charts\community\artifactory-ha`
- Source: `dependency`

```text
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add'
```

### 5. `JahstreetOrg/spark-on-kubernetes-helm`

- Chart: `D:\helm_clones_github\JahstreetOrg__spark-on-kubernetes-helm\charts\cluster-base`
- Source: `dependency`

```text
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: no repository definition for https://kubernetes.github.io/ingress-nginx, https://charts.jetstack.io, https://charts.helm.sh/stable, https://charts.helm.sh/stable. Please add the missing repos via 'helm repo add'
```

### 6. `unixhot/devops-x`

- Chart: `D:\helm_clones_github\unixhot__devops-x\helm\gitlab`
- Source: `dependency`

```text
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies
```

### 7. `rancher/rke2-charts`

- Chart: `D:\helm_clones_github\rancher__rke2-charts\packages\rke2-cilium-legacy\charts`
- Source: `dependency`

```text
Error: directory D:\helm_clones_github\rancher__rke2-charts\packages\rke2-cilium-legacy\charts\charts\cilium not found
```

### 8. `pluralsh/plural-artifacts`

- Chart: `D:\helm_clones_github\pluralsh__plural-artifacts\airbyte\helm\airbyte`
- Source: `dependency`

```text
Error: directory D:\helm_clones_github\airbyte\charts\airbyte not found
```

### 9. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\airflow`
- Source: `dependency`

```text
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/, https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add'
```

### 10. `bcgov/OCWA`

- Chart: `D:\helm_clones_github\bcgov__OCWA\helm\ocwa`
- Source: `dependency`

```text
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies
```

### 11. `logicalisuki/ubiquity-open`

- Chart: `D:\helm_clones_github\logicalisuki__ubiquity-open\disabled\platform\opensm`
- Source: `dependency`

```text
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "https://github.com/Mellanox/network-operator" chart repository:
	failed to fetch https://github.com/Mellanox/network-operator/index.yaml : 404 Not Found
Error: no cached repository for helm-manager-bedd4d6f25f9f14d254639d8224675502e27945a13484e5a9e1499a78d72770e found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-bedd4d6f25f9f14d254639d8224675502e27945a13484e5a9e1499a78d72770e-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### 12. `thoughtworks/byor-voting-infrastructure`

- Chart: `D:\helm_clones_github\thoughtworks__byor-voting-infrastructure\src\byor-voting-chart`
- Source: `dependency`

```text
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies
```

### 13. `cloudstark/helm-charts`

- Chart: `D:\helm_clones_github\cloudstark__helm-charts\postgrest`
- Source: `dependency`

```text
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add'
```

### 14. `atsip76/k8s_asterisk_project`

- Chart: `D:\helm_clones_github\atsip76__k8s_asterisk_project\k8s\gitlab`
- Source: `dependency`

```text
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: no repository definition for https://charts.gitlab.io/. Please add the missing repos via 'helm repo add'
```

### 15. `ibuildthecloud/rancher-charts`

- Chart: `D:\helm_clones_github\ibuildthecloud__rancher-charts\charts\anchore-engine\0.1.0`
- Source: `dependency`

```text
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: no repository definition for https://kubernetes-charts.storage.googleapis.com. Please add the missing repos via 'helm repo add'
```

### 16. `Otus-DevOps-2019-08/sgremyachikh_microservices`

- Chart: `D:\helm_clones_github\Otus-DevOps-2019-08__sgremyachikh_microservices\kubernetes\Charts\gitlab-omnibus`
- Source: `dependency`

```text
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: no repository definition for https://charts.gitlab.io/. Please add the missing repos via 'helm repo add'
```

### 17. `punchplatform/punch-helm`

- Chart: `D:\helm_clones_github\punchplatform__punch-helm\operator`
- Source: `dependency`

```text
Error: directory D:\helm_clones_github\punchplatform__punch-helm\operator\charts\operator.certificate not found
```

### 18. `EamonKeane/k8s-cluster-services`

- Chart: `D:\helm_clones_github\EamonKeane__k8s-cluster-services\cluster-svc`
- Source: `dependency`

```text
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: no repository definition for http://kubernetes-charts.storage.googleapis.com/, http://kubernetes-charts.storage.googleapis.com/, http://storage.googleapis.com/kubernetes-charts-incubator, http://kubernetes-charts.storage.googleapis.com/, http://storage.googleapis.com/kubernetes-charts-incubator, http://storage.googleapis.com/kubernetes-charts-incubator, http://kubernetes-charts.storage.googleapis.com/, http://kubernetes-charts.storage.googleapis.com/, http://kubernetes-charts.storage.googleapis.com/, http://kubernetes-charts.storage.googleapis.com/, http://kubernetes-charts.storage.googleapis.com/, https://opensource-helm.squareroute.io/, http://kubernetes-charts.storage.googleapis.com/, https://helm.github.io/monocular, http://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add'
```

### 19. `Makhuta/truecharts-archive-scale-catalog`

- Chart: `D:\helm_clones_github\Makhuta__truecharts-archive-scale-catalog\incubator\archivebox\0.7.2`
- Source: `dependency`

```text
Saving 1 charts
Downloading common from repo oci://tccr.io/truecharts
Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:17.2.30: not found
Error: could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:17.2.30: not found
```

### 20. `opspresso/argocd-env-addons`

- Chart: `D:\helm_clones_github\opspresso__argocd-env-addons\charts\dashboard`
- Source: `dependency`

```text
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "https://kubernetes.github.io/dashboard" chart repository:
	failed to fetch https://kubernetes.github.io/dashboard/index.yaml : 404 Not Found
...Successfully got an update from the "https://charts.helm.sh/incubator" chart repository
Error: no cached repository for helm-manager-fc08c6c0f466a809ed2b24637e970ca3cd7bc1d7524efc4832f2405812f07ab0 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-fc08c6c0f466a809ed2b24637e970ca3cd7bc1d7524efc4832f2405812f07ab0-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### 21. `Sureya/airflow_k8s_executor`

- Chart: `D:\helm_clones_github\Sureya__airflow_k8s_executor\helm_charts\official\charts\incubator\distribution`
- Source: `dependency`

```text
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/, https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add'
```

### 22. `TheOpsDev/homelab`

- Chart: `D:\helm_clones_github\TheOpsDev__homelab\charts\k8s-dashboard`
- Source: `dependency`

```text
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "https://kubernetes.github.io/dashboard/" chart repository:
	failed to fetch https://kubernetes.github.io/dashboard/index.yaml : 404 Not Found
Error: no cached repository for helm-manager-602693e8f5d1a68dc0300eb544f8e9829d89b7af15ee517b5231c07768425e69 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-602693e8f5d1a68dc0300eb544f8e9829d89b7af15ee517b5231c07768425e69-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### 23. `adstanley/archive`

- Chart: `D:\helm_clones_github\adstanley__archive\scale-catalog\incubator\archivebox\0.7.2`
- Source: `dependency`

```text
Saving 1 charts
Downloading common from repo oci://tccr.io/truecharts
Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:17.2.30: not found
Error: could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:17.2.30: not found
```

### 24. `hey101/scale-catalog`

- Chart: `D:\helm_clones_github\hey101__scale-catalog\incubator\archivebox\0.7.2`
- Source: `dependency`

```text
Saving 1 charts
Downloading common from repo oci://tccr.io/truecharts
Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:17.5.1: not found
Error: could not download oci://tccr.io/truecharts/common: failed to perform "FetchReference" on source: tccr.io/truecharts/common:17.5.1: not found
```

### 25. `Arthur-B-DevOps/old_helm_charts`

- Chart: `D:\helm_clones_github\Arthur-B-DevOps__old_helm_charts\charts\Old_charts\charts\incubator\distribution`
- Source: `dependency`

```text
level=INFO msg="warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash..."
Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/, https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add'
```

## `template.runtime_eval` (23)

### 1. `WeBankFinTech/Prophecis`

- Chart: `D:\helm_clones_github\WeBankFinTech__Prophecis\di\jobmonitor\charts\mpijob`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\WeBankFinTech__Prophecis\di\jobmonitor\charts\mpijob`

```text
Error: template: mpijob/templates/mpijob.yaml:37:17: executing "mpijob/templates/mpijob.yaml" at <len .Values.nodeSelectors>: error calling len: len of nil pointer

Use --debug flag to render out invalid YAML
```

### 2. `WeBankFinTech/Prophecis`

- Chart: `D:\helm_clones_github\WeBankFinTech__Prophecis\di\jobmonitor\charts\tfjob`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\WeBankFinTech__Prophecis\di\jobmonitor\charts\tfjob`

```text
Error: template: tfjob/templates/tfjob.yaml:300:21: executing "tfjob/templates/tfjob.yaml" at <len $tfNodeSelectors.Worker>: error calling len: reflect: call of reflect.Value.Type on zero Value

Use --debug flag to render out invalid YAML
```

### 3. `WeBankFinTech/Prophecis`

- Chart: `D:\helm_clones_github\WeBankFinTech__Prophecis\di\lcm\charts\mpijob`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\WeBankFinTech__Prophecis\di\lcm\charts\mpijob`

```text
Error: template: mpijob/templates/mpijob.yaml:37:17: executing "mpijob/templates/mpijob.yaml" at <len .Values.nodeSelectors>: error calling len: len of nil pointer

Use --debug flag to render out invalid YAML
```

### 4. `WeBankFinTech/Prophecis`

- Chart: `D:\helm_clones_github\WeBankFinTech__Prophecis\di\lcm\charts\tfjob`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\WeBankFinTech__Prophecis\di\lcm\charts\tfjob`

```text
Error: template: tfjob/templates/tfjob.yaml:300:21: executing "tfjob/templates/tfjob.yaml" at <len $tfNodeSelectors.Worker>: error calling len: reflect: call of reflect.Value.Type on zero Value

Use --debug flag to render out invalid YAML
```

### 5. `IBM/charts`

- Chart: `D:\helm_clones_github\IBM__charts\community\aqua-server`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\IBM__charts\community\aqua-server`

```text
Error: aqua-server/templates/web-secrets.yaml:2:11
  executing "aqua-server/templates/web-secrets.yaml" at <(.Values.admin.password) .Values.admin.token>:
    can't give argument to non-function .Values.admin.password

Use --debug flag to render out invalid YAML
```

### 6. `Thakurvaibhav/k8s`

- Chart: `D:\helm_clones_github\Thakurvaibhav__k8s\.archive\consul`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Thakurvaibhav__k8s\.archive\consul`

```text
Error: consul/templates/client-daemonset.yaml:125:23
  executing "consul/templates/client-daemonset.yaml" at <(.Values.client.join) and (gt (len .Values.client.join) 0)>:
    can't give argument to non-function .Values.client.join

Use --debug flag to render out invalid YAML
```

### 7. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio`

```text
Error: istio/charts/gateways/templates/rolebindings.yaml:4:7
  executing "istio/charts/gateways/templates/rolebindings.yaml" at <($spec.sds) and (eq $spec.sds.enabled true)>:
    can't give argument to non-function $spec.sds

Use --debug flag to render out invalid YAML
```

### 8. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio`
- Source: `template`
- Values files: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-demo-auth.yaml`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio -f D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-demo-auth.yaml`

```text
Error: istio/charts/gateways/templates/rolebindings.yaml:4:7
  executing "istio/charts/gateways/templates/rolebindings.yaml" at <($spec.sds) and (eq $spec.sds.enabled true)>:
    can't give argument to non-function $spec.sds

Use --debug flag to render out invalid YAML
```

### 9. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio`
- Source: `template`
- Values files: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-demo.yaml`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio -f D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-demo.yaml`

```text
Error: istio/charts/gateways/templates/rolebindings.yaml:4:7
  executing "istio/charts/gateways/templates/rolebindings.yaml" at <($spec.sds) and (eq $spec.sds.enabled true)>:
    can't give argument to non-function $spec.sds

Use --debug flag to render out invalid YAML
```

### 10. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio`
- Source: `template`
- Values files: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-demo-auth.yaml`, `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-demo.yaml`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio -f D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-demo-auth.yaml -f D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-demo.yaml`

```text
Error: istio/charts/gateways/templates/rolebindings.yaml:4:7
  executing "istio/charts/gateways/templates/rolebindings.yaml" at <($spec.sds) and (eq $spec.sds.enabled true)>:
    can't give argument to non-function $spec.sds

Use --debug flag to render out invalid YAML
```

### 11. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio`
- Source: `template`
- Values files: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-sds-auth-control-plane-auth-disabled.yaml`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio -f D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-sds-auth-control-plane-auth-disabled.yaml`

```text
Error: istio/charts/gateways/templates/rolebindings.yaml:4:7
  executing "istio/charts/gateways/templates/rolebindings.yaml" at <($spec.sds) and (eq $spec.sds.enabled true)>:
    can't give argument to non-function $spec.sds

Use --debug flag to render out invalid YAML
```

### 12. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio`
- Source: `template`
- Values files: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-demo-auth.yaml`, `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-sds-auth-control-plane-auth-disabled.yaml`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio -f D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-demo-auth.yaml -f D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-sds-auth-control-plane-auth-disabled.yaml`

```text
Error: istio/charts/gateways/templates/rolebindings.yaml:4:7
  executing "istio/charts/gateways/templates/rolebindings.yaml" at <($spec.sds) and (eq $spec.sds.enabled true)>:
    can't give argument to non-function $spec.sds

Use --debug flag to render out invalid YAML
```

### 13. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio`
- Source: `template`
- Values files: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-demo.yaml`, `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-sds-auth-control-plane-auth-disabled.yaml`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio -f D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-demo.yaml -f D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-sds-auth-control-plane-auth-disabled.yaml`

```text
Error: istio/charts/gateways/templates/rolebindings.yaml:4:7
  executing "istio/charts/gateways/templates/rolebindings.yaml" at <($spec.sds) and (eq $spec.sds.enabled true)>:
    can't give argument to non-function $spec.sds

Use --debug flag to render out invalid YAML
```

### 14. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio`
- Source: `template`
- Values files: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-demo-auth.yaml`, `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-demo.yaml`, `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-sds-auth-control-plane-auth-disabled.yaml`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio -f D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-demo-auth.yaml -f D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-demo.yaml -f D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-sds-auth-control-plane-auth-disabled.yaml`

```text
Error: istio/charts/gateways/templates/rolebindings.yaml:4:7
  executing "istio/charts/gateways/templates/rolebindings.yaml" at <($spec.sds) and (eq $spec.sds.enabled true)>:
    can't give argument to non-function $spec.sds

Use --debug flag to render out invalid YAML
```

### 15. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio`
- Source: `template`
- Values files: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-sds-auth.yaml`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio -f D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-sds-auth.yaml`

```text
Error: istio/charts/gateways/templates/rolebindings.yaml:4:7
  executing "istio/charts/gateways/templates/rolebindings.yaml" at <($spec.sds) and (eq $spec.sds.enabled true)>:
    can't give argument to non-function $spec.sds

Use --debug flag to render out invalid YAML
```

### 16. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio`
- Source: `template`
- Values files: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-demo-auth.yaml`, `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-sds-auth.yaml`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio -f D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-demo-auth.yaml -f D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-sds-auth.yaml`

```text
Error: istio/charts/gateways/templates/rolebindings.yaml:4:7
  executing "istio/charts/gateways/templates/rolebindings.yaml" at <($spec.sds) and (eq $spec.sds.enabled true)>:
    can't give argument to non-function $spec.sds

Use --debug flag to render out invalid YAML
```

### 17. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio`
- Source: `template`
- Values files: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-demo.yaml`, `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-sds-auth.yaml`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio -f D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-demo.yaml -f D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-sds-auth.yaml`

```text
Error: istio/charts/gateways/templates/rolebindings.yaml:4:7
  executing "istio/charts/gateways/templates/rolebindings.yaml" at <($spec.sds) and (eq $spec.sds.enabled true)>:
    can't give argument to non-function $spec.sds

Use --debug flag to render out invalid YAML
```

### 18. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio`
- Source: `template`
- Values files: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-demo-auth.yaml`, `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-demo.yaml`, `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-sds-auth.yaml`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio -f D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-demo-auth.yaml -f D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-demo.yaml -f D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-sds-auth.yaml`

```text
Error: istio/charts/gateways/templates/rolebindings.yaml:4:7
  executing "istio/charts/gateways/templates/rolebindings.yaml" at <($spec.sds) and (eq $spec.sds.enabled true)>:
    can't give argument to non-function $spec.sds

Use --debug flag to render out invalid YAML
```

### 19. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio`
- Source: `template`
- Values files: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-sds-auth-control-plane-auth-disabled.yaml`, `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-sds-auth.yaml`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio -f D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-sds-auth-control-plane-auth-disabled.yaml -f D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-sds-auth.yaml`

```text
Error: istio/charts/gateways/templates/rolebindings.yaml:4:7
  executing "istio/charts/gateways/templates/rolebindings.yaml" at <($spec.sds) and (eq $spec.sds.enabled true)>:
    can't give argument to non-function $spec.sds

Use --debug flag to render out invalid YAML
```

### 20. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio`
- Source: `template`
- Values files: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-demo-auth.yaml`, `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-sds-auth-control-plane-auth-disabled.yaml`, `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-sds-auth.yaml`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio -f D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-demo-auth.yaml -f D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-sds-auth-control-plane-auth-disabled.yaml -f D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-sds-auth.yaml`

```text
Error: istio/charts/gateways/templates/rolebindings.yaml:4:7
  executing "istio/charts/gateways/templates/rolebindings.yaml" at <($spec.sds) and (eq $spec.sds.enabled true)>:
    can't give argument to non-function $spec.sds

Use --debug flag to render out invalid YAML
```

### 21. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio`
- Source: `template`
- Values files: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-demo.yaml`, `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-sds-auth-control-plane-auth-disabled.yaml`, `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-sds-auth.yaml`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio -f D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-demo.yaml -f D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-sds-auth-control-plane-auth-disabled.yaml -f D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-sds-auth.yaml`

```text
Error: istio/charts/gateways/templates/rolebindings.yaml:4:7
  executing "istio/charts/gateways/templates/rolebindings.yaml" at <($spec.sds) and (eq $spec.sds.enabled true)>:
    can't give argument to non-function $spec.sds

Use --debug flag to render out invalid YAML
```

### 22. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio`
- Source: `template`
- Values files: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-demo-auth.yaml`, `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-demo.yaml`, `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-sds-auth-control-plane-auth-disabled.yaml`, `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-sds-auth.yaml`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio -f D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-demo-auth.yaml -f D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-demo.yaml -f D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-sds-auth-control-plane-auth-disabled.yaml -f D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\istio-system\istio\values-istio-sds-auth.yaml`

```text
Error: istio/charts/gateways/templates/rolebindings.yaml:4:7
  executing "istio/charts/gateways/templates/rolebindings.yaml" at <($spec.sds) and (eq $spec.sds.enabled true)>:
    can't give argument to non-function $spec.sds

Use --debug flag to render out invalid YAML
```

### 23. `yunzck8s/cloudNative`

- Chart: `D:\helm_clones_github\yunzck8s__cloudNative\charts\cilium`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\yunzck8s__cloudNative\charts\cilium`

```text
Error: template: cilium/templates/cilium-envoy/configmap.yaml:15:5: executing "cilium/templates/cilium-envoy/configmap.yaml" at <tpl (.Files.Glob "files/cilium-envoy/configmap/bootstrap-config.json").AsConfig .>: error calling tpl: cannot parse template "bootstrap-config.json: \"{\\r\\n  \\\"node\\\": {\\r\\n    \\\"id\\\": \\\"host~127.0.0.1~no-id~localdomain\\\",\\r\\n\n  \\   \\\"cluster\\\": \\\"ingress-cluster\\\"\\r\\n  },\\r\\n  \\\"staticResources\\\": {\\r\\n    \\\"listeners\\\":\n  [\\r\\n      {{- if .Values.envoy.prometheus.enabled }}\\r\\n      {\\r\\n        \\\"name\\\":\n  \\\"envoy-prometheus-metrics-listener\\\",\\r\\n        \\\"address\\\": {\\r\\n          \\\"socket_address\\\":\n  {\\r\\n            \\\"address\\\": \\\"0.0.0.0\\\",\\r\\n            \\\"port_value\\\": {{ .Values.envoy.prometheus.port\n  }}\\r\\n          }\\r\\n        },\\r\\n        \\\"filter_chains\\\": [\\r\\n          {\\r\\n\n  \\           \\\"filters\\\": [\\r\\n              {\\r\\n                \\\"name\\\": \\\"envoy.filters.network.http_connection_manager\\\",\\r\\n\n  \\               \\\"typed_config\\\": {\\r\\n                  \\\"@type\\\": \\\"type.googleapis.com/envoy.extensions.filters.network.http_connection_manager.v3.HttpConnectionManager\\\",\\r\\n\n  \\                 \\\"stat_prefix\\\": \\\"envoy-prometheus-metrics-listener\\\",\\r\\n                  \\\"route_config\\\":\n  {\\r\\n                    \\\"virtual_hosts\\\": [\\r\\n                      {\\r\\n                        \\\"name\\\":\n  \\\"prometheus_metrics_route\\\",\\r\\n                        \\\"domains\\\": [\\r\\n                          \\\"*\\\"\\r\\n\n  \\                       ],\\r\\n                        \\\"routes\\\": [\\r\\n                          {\\r\\n\n  \\                           \\\"name\\\": \\\"prometheus_metrics_route\\\",\\r\\n                            \\\"match\\\":\n  {\\r\\n                              \\\"prefix\\\": \\\"/metrics\\\"\\r\\n                            },\\r\\n\n  \\                           \\\"route\\\": {\\r\\n                              \\\"cluster\\\":\n  \\\"/envoy-admin\\\",\\r\\n                              \\\"prefix_rewrite\\\": \\\"/stats/prometheus\\\"\\r\\n\n  \\                           }\\r\\n                          }\\r\\n                        ]\\r\\n\n  \\                     }\\r\\n                    ]\\r\\n                  },\\r\\n                  \\\"http_filters\\\":\n  [\\r\\n                    {\\r\\n                      \\\"name\\\": \\\"envoy.filters.http.router\\\",\\r\\n\n  \\                     \\\"typed_config\\\": {\\r\\n                        \\\"@type\\\":\n  \\\"type.googleapis.com/envoy.extensions.filters.http.router.v3.Router\\\"\\r\\n                      }\\r\\n\n  \\                   }\\r\\n                  ],\\r\\n                  \\\"stream_idle_timeout\\\":\n  \\\"0s\\\"\\r\\n                }\\r\\n              }\\r\\n            ]\\r\\n          }\\r\\n\n  \\       ]\\r\\n      },\\r\\n      {{- end }}\\r\\n      {{- if and .Values.envoy.debug.admin.enabled\n  }}\\r\\n      {\\r\\n        \\\"name\\\": \\\"envoy-admin-listener\\\",\\r\\n        \\\"address\\\":\n  {\\r\\n          \\\"socket_address\\\": {\\r\\n            \\\"address\\\": {{ .Values.ipv4.enabled\n  | ternary \\\"127.0.0.1\\\" \\\"::1\\\" | quote }},\\r\\n            \\\"port_value\\\": {{ .Values.envoy.debug.admin.port\n  }}\\r\\n          }\\r\\n        },\\r\\n        {{- if and .Values.ipv4.enabled .Values.ipv6.enabled\n  }}\\r\\n        \\\"additional_addresses\\\": [\\r\\n          {\\r\\n            \\\"address\\\":\n  {\\r\\n              \\\"socket_address\\\": {\\r\\n                \\\"address\\\": \\\"::1\\\",\\r\\n\n  \\               \\\"port_value\\\": {{ .Values.envoy.debug.admin.port }}\\r\\n              }\\r\\n\n  \\           }\\r\\n          }\\r\\n        ],\\r\\n        {{- end }}\\r\\n        \\\"filter_chains\\\":\n  [\\r\\n          {\\r\\n            \\\"filters\\\": [\\r\\n              {\\r\\n                \\\"name\\\":\n  \\\"envoy.filters.network.http_connection_manager\\\",\\r\\n                \\\"typed_config\\\":\n  {\\r\\n                  \\\"@type\\\": \\\"type.googleapis.com/envoy.extensions.filters.network.http_connection_manager.v3.HttpConnectionManager\\\",\\r\\n\n  \\                 \\\"stat_prefix\\\": \\\"envoy-admin-listener\\\",\\r\\n                  \\\"route_config\\\":\n  {\\r\\n                    \\\"virtual_hosts\\\": [\\r\\n                      {\\r\\n                        \\\"name\\\":\n  \\\"admin_route\\\",\\r\\n                        \\\"domains\\\": [\\r\\n                          \\\"*\\\"\\r\\n\n  \\                       ],\\r\\n                        \\\"routes\\\": [\\r\\n                          {\\r\\n\n  \\                           \\\"name\\\": \\\"admin_route\\\",\\r\\n                            \\\"match\\\":\n  {\\r\\n                              \\\"prefix\\\": \\\"/\\\"\\r\\n                            },\\r\\n\n  \\                           \\\"route\\\": {\\r\\n                              \\\"cluster\\\":\n  \\\"/envoy-admin\\\",\\r\\n                              \\\"prefix_rewrite\\\": \\\"/\\\"\\r\\n\n  \\                           }\\r\\n                          }\\r\\n                        ]\\r\\n\n  \\                     }\\r\\n                    ]\\r\\n                  },\\r\\n                  \\\"http_filters\\\":\n  [\\r\\n                    {\\r\\n                      \\\"name\\\": \\\"envoy.filters.http.router\\\",\\r\\n\n  \\                     \\\"typed_config\\\": {\\r\\n                        \\\"@type\\\":\n  \\\"type.googleapis.com/envoy.extensions.filters.http.router.v3.Router\\\"\\r\\n                      }\\r\\n\n  \\                   }\\r\\n                  ],\\r\\n                  \\\"stream_idle_timeout\\\":\n  \\\"0s\\\"\\r\\n                }\\r\\n              }\\r\\n            ]\\r\\n          }\\r\\n\n  \\       ]\\r\\n      },\\r\\n      {{- end }}\\r\\n      {\\r\\n        \\\"name\\\": \\\"envoy-health-listener\\\",\\r\\n\n  \\       \\\"address\\\": {\\r\\n          \\\"socket_address\\\": {\\r\\n            \\\"address\\\":\n  {{ .Values.ipv4.enabled | ternary \\\"127.0.0.1\\\" \\\"::1\\\" | quote }},\\r\\n            \\\"port_value\\\":\n  {{ .Values.envoy.healthPort }}\\r\\n          }\\r\\n        },\\r\\n        {{- if and\n  .Values.ipv4.enabled .Values.ipv6.enabled }}\\r\\n        \\\"additional_addresses\\\":\n  [\\r\\n          {\\r\\n            \\\"address\\\": {\\r\\n              \\\"socket_address\\\":\n  {\\r\\n                \\\"address\\\": \\\"::1\\\",\\r\\n                \\\"port_value\\\": {{\n  .Values.envoy.healthPort }}\\r\\n              }\\r\\n            }\\r\\n          }\\r\\n\n  \\       ],\\r\\n        {{- end }}\\r\\n        \\\"filter_chains\\\": [\\r\\n          {\\r\\n\n  \\           \\\"filters\\\": [\\r\\n              {\\r\\n                \\\"name\\\": \\\"envoy.filters.network.http_connection_manager\\\",\\r\\n\n  \\               \\\"typed_config\\\": {\\r\\n                  \\\"@type\\\": \\\"type.googleapis.com/envoy.extensions.filters.network.http_connection_manager.v3.HttpConnectionManager\\\",\\r\\n\n  \\                 \\\"stat_prefix\\\": \\\"envoy-health-listener\\\",\\r\\n                  \\\"route_config\\\":\n  {\\r\\n                    \\\"virtual_hosts\\\": [\\r\\n                      {\\r\\n                        \\\"name\\\":\n  \\\"health\\\",\\r\\n                        \\\"domains\\\": [\\r\\n                          \\\"*\\\"\\r\\n\n  \\                       ],\\r\\n                        \\\"routes\\\": [\\r\\n                          {\\r\\n\n  \\                           \\\"name\\\": \\\"health\\\",\\r\\n                            \\\"match\\\":\n  {\\r\\n                              \\\"prefix\\\": \\\"/healthz\\\"\\r\\n                            },\\r\\n\n  \\                           \\\"route\\\": {\\r\\n                              \\\"cluster\\\":\n  \\\"/envoy-admin\\\",\\r\\n                              \\\"prefix_rewrite\\\": \\\"/ready\\\"\\r\\n\n  \\                           }\\r\\n                          }\\r\\n                        ]\\r\\n\n  \\                     }\\r\\n                    ]\\r\\n                  },\\r\\n                  \\\"http_filters\\\":\n  [\\r\\n                    {\\r\\n                      \\\"name\\\": \\\"envoy.filters.http.router\\\",\\r\\n\n  \\                     \\\"typed_config\\\": {\\r\\n                        \\\"@type\\\":\n  \\\"type.googleapis.com/envoy.extensions.filters.http.router.v3.Router\\\"\\r\\n                      }\\r\\n\n  \\                   }\\r\\n                  ],\\r\\n                  \\\"stream_idle_timeout\\\":\n  \\\"0s\\\"\\r\\n                }\\r\\n              }\\r\\n            ]\\r\\n          }\\r\\n\n  \\       ]\\r\\n      }\\r\\n    ],\\r\\n    \\\"clusters\\\": [\\r\\n      {\\r\\n        \\\"name\\\":\n  \\\"ingress-cluster\\\",\\r\\n        \\\"type\\\": \\\"ORIGINAL_DST\\\",\\r\\n        \\\"connectTimeout\\\":\n  \\\"{{ .Values.envoy.connectTimeoutSeconds }}s\\\",\\r\\n        \\\"lbPolicy\\\": \\\"CLUSTER_PROVIDED\\\",\\r\\n\n  \\       \\\"typedExtensionProtocolOptions\\\": {\\r\\n          \\\"envoy.extensions.upstreams.http.v3.HttpProtocolOptions\\\":\n  {\\r\\n            \\\"@type\\\": \\\"type.googleapis.com/envoy.extensions.upstreams.http.v3.HttpProtocolOptions\\\",\\r\\n\n  \\           \\\"commonHttpProtocolOptions\\\": {\\r\\n              \\\"idleTimeout\\\": \\\"{{\n  .Values.envoy.idleTimeoutDurationSeconds }}s\\\",\\r\\n              \\\"maxConnectionDuration\\\":\n  \\\"{{ .Values.envoy.maxConnectionDurationSeconds }}s\\\",\\r\\n              \\\"maxRequestsPerConnection\\\":\n  {{ .Values.envoy.maxRequestsPerConnection }}\\r\\n            },\\r\\n            \\\"useDownstreamProtocolConfig\\\":\n  {}\\r\\n          }\\r\\n        },\\r\\n        \\\"cleanupInterval\\\": \\\"{{ .Values.envoy.connectTimeoutSeconds\n  }}.500s\\\"\\r\\n      },\\r\\n      {\\r\\n        \\\"name\\\": \\\"egress-cluster-tls\\\",\\r\\n\n  \\       \\\"type\\\": \\\"ORIGINAL_DST\\\",\\r\\n        \\\"connectTimeout\\\": \\\"{{ .Values.envoy.connectTimeoutSeconds\n  }}s\\\",\\r\\n        \\\"lbPolicy\\\": \\\"CLUSTER_PROVIDED\\\",\\r\\n        \\\"typedExtensionProtocolOptions\\\":\n  {\\r\\n          \\\"envoy.extensions.upstreams.http.v3.HttpProtocolOptions\\\": {\\r\\n\n  \\           \\\"@type\\\": \\\"type.googleapis.com/envoy.extensions.upstreams.http.v3.HttpProtocolOptions\\\",\\r\\n\n  \\           \\\"commonHttpProtocolOptions\\\": {\\r\\n              \\\"idleTimeout\\\": \\\"{{\n  .Values.envoy.idleTimeoutDurationSeconds }}s\\\",\\r\\n              \\\"maxConnectionDuration\\\":\n  \\\"{{ .Values.envoy.maxConnectionDurationSeconds }}s\\\",\\r\\n              \\\"maxRequestsPerConnection\\\":\n  {{ .Values.envoy.maxRequestsPerConnection }}\\r\\n            },\\r\\n            \\\"upstreamHttpProtocolOptions\\\":\n  {},\\r\\n            \\\"useDownstreamProtocolConfig\\\": {}\\r\\n          }\\r\\n        },\\r\\n\n  \\       \\\"cleanupInterval\\\": \\\"{{ .Values.envoy.connectTimeoutSeconds }}.500s\\\",\\r\\n\n  \\       \\\"transportSocket\\\": {\\r\\n          \\\"name\\\": \\\"cilium.tls_wrapper\\\",\\r\\n\n  \\         \\\"typedConfig\\\": {\\r\\n            \\\"@type\\\": \\\"type.googleapis.com/cilium.UpstreamTlsWrapperContext\\\"\\r\\n\n  \\         }\\r\\n        }\\r\\n      },\\r\\n      {\\r\\n        \\\"name\\\": \\\"egress-cluster\\\",\\r\\n\n  \\       \\\"type\\\": \\\"ORIGINAL_DST\\\",\\r\\n        \\\"connectTimeout\\\": \\\"{{ .Values.envoy.connectTimeoutSeconds\n  }}s\\\",\\r\\n        \\\"lbPolicy\\\": \\\"CLUSTER_PROVIDED\\\",\\r\\n        \\\"typedExtensionProtocolOptions\\\":\n  {\\r\\n          \\\"envoy.extensions.upstreams.http.v3.HttpProtocolOptions\\\": {\\r\\n\n  \\           \\\"@type\\\": \\\"type.googleapis.com/envoy.extensions.upstreams.http.v3.HttpProtocolOptions\\\",\\r\\n\n  \\           \\\"commonHttpProtocolOptions\\\": {\\r\\n              \\\"idleTimeout\\\": \\\"{{\n  .Values.envoy.idleTimeoutDurationSeconds }}s\\\",\\r\\n              \\\"maxConnectionDuration\\\":\n  \\\"{{ .Values.envoy.maxConnectionDurationSeconds }}s\\\",\\r\\n              \\\"maxRequestsPerConnection\\\":\n  {{ .Values.envoy.maxRequestsPerConnection }}\\r\\n            },\\r\\n            \\\"useDownstreamProtocolConfig\\\":\n  {}\\r\\n          }\\r\\n        },\\r\\n        \\\"cleanupInterval\\\": \\\"{{ .Values.envoy.connectTimeoutSeconds\n  }}.500s\\\"\\r\\n      },\\r\\n      {\\r\\n        \\\"name\\\": \\\"ingress-cluster-tls\\\",\\r\\n\n  \\       \\\"type\\\": \\\"ORIGINAL_DST\\\",\\r\\n        \\\"connectTimeout\\\": \\\"{{ .Values.envoy.connectTimeoutSeconds\n  }}s\\\",\\r\\n        \\\"lbPolicy\\\": \\\"CLUSTER_PROVIDED\\\",\\r\\n        \\\"typedExtensionProtocolOptions\\\":\n  {\\r\\n          \\\"envoy.extensions.upstreams.http.v3.HttpProtocolOptions\\\": {\\r\\n\n  \\           \\\"@type\\\": \\\"type.googleapis.com/envoy.extensions.upstreams.http.v3.HttpProtocolOptions\\\",\\r\\n\n  \\           \\\"commonHttpProtocolOptions\\\": {\\r\\n              \\\"idleTimeout\\\": \\\"{{\n  .Values.envoy.idleTimeoutDurationSeconds }}s\\\",\\r\\n              \\\"maxConnectionDuration\\\":\n  \\\"{{ .Values.envoy.maxConnectionDurationSeconds }}s\\\",\\r\\n              \\\"maxRequestsPerConnection\\\":\n  {{ .Values.envoy.maxRequestsPerConnection }}\\r\\n            },\\r\\n            \\\"upstreamHttpProtocolOptions\\\":\n  {},\\r\\n            \\\"useDownstreamProtocolConfig\\\": {}\\r\\n          }\\r\\n        },\\r\\n\n  \\       \\\"cleanupInterval\\\": \\\"{{ .Values.envoy.connectTimeoutSeconds }}.500s\\\",\\r\\n\n  \\       \\\"transportSocket\\\": {\\r\\n          \\\"name\\\": \\\"cilium.tls_wrapper\\\",\\r\\n\n  \\         \\\"typedConfig\\\": {\\r\\n            \\\"@type\\\": \\\"type.googleapis.com/cilium.UpstreamTlsWrapperContext\\\"\\r\\n\n  \\         }\\r\\n        }\\r\\n      },\\r\\n      {\\r\\n        \\\"name\\\": \\\"xds-grpc-cilium\\\",\\r\\n\n  \\       \\\"type\\\": \\\"STATIC\\\",\\r\\n        \\\"connectTimeout\\\": \\\"{{ .Values.envoy.connectTimeoutSeconds\n  }}s\\\",\\r\\n        \\\"loadAssignment\\\": {\\r\\n          \\\"clusterName\\\": \\\"xds-grpc-cilium\\\",\\r\\n\n  \\         \\\"endpoints\\\": [\\r\\n            {\\r\\n              \\\"lbEndpoints\\\": [\\r\\n\n  \\               {\\r\\n                  \\\"endpoint\\\": {\\r\\n                    \\\"address\\\":\n  {\\r\\n                      \\\"pipe\\\": {\\r\\n                        \\\"path\\\": \\\"/var/run/cilium/envoy/sockets/xds.sock\\\"\\r\\n\n  \\                     }\\r\\n                    }\\r\\n                  }\\r\\n                }\\r\\n\n  \\             ]\\r\\n            }\\r\\n          ]\\r\\n        },\\r\\n        \\\"typedExtensionProtocolOptions\\\":\n  {\\r\\n          \\\"envoy.extensions.upstreams.http.v3.HttpProtocolOptions\\\": {\\r\\n\n  \\           \\\"@type\\\": \\\"type.googleapis.com/envoy.extensions.upstreams.http.v3.HttpProtocolOptions\\\",\\r\\n\n  \\           \\\"explicitHttpConfig\\\": {\\r\\n              \\\"http2ProtocolOptions\\\":\n  {}\\r\\n            }\\r\\n          }\\r\\n        }\\r\\n      },\\r\\n      {\\r\\n        \\\"name\\\":\n  \\\"/envoy-admin\\\",\\r\\n        \\\"type\\\": \\\"STATIC\\\",\\r\\n        \\\"connectTimeout\\\":\n  \\\"{{ .Values.envoy.connectTimeoutSeconds }}s\\\",\\r\\n        \\\"loadAssignment\\\": {\\r\\n\n  \\         \\\"clusterName\\\": \\\"/envoy-admin\\\",\\r\\n          \\\"endpoints\\\": [\\r\\n            {\\r\\n\n  \\             \\\"lbEndpoints\\\": [\\r\\n                {\\r\\n                  \\\"endpoint\\\":\n  {\\r\\n                    \\\"address\\\": {\\r\\n                      \\\"pipe\\\": {\\r\\n\n  \\                       \\\"path\\\": \\\"/var/run/cilium/envoy/sockets/admin.sock\\\"\\r\\n\n  \\                     }\\r\\n                    }\\r\\n                  }\\r\\n                }\\r\\n\n  \\             ]\\r\\n            }\\r\\n          ]\\r\\n        }\\r\\n      }\\r\\n    ]\\r\\n\n  \\ },\\r\\n  \\\"dynamicResources\\\": {\\r\\n    \\\"ldsConfig\\\": {\\r\\n      \\\"apiConfigSource\\\":\n  {\\r\\n        \\\"apiType\\\": \\\"GRPC\\\",\\r\\n        \\\"transportApiVersion\\\": \\\"V3\\\",\\r\\n\n  \\       \\\"grpcServices\\\": [\\r\\n          {\\r\\n            \\\"envoyGrpc\\\": {\\r\\n              \\\"clusterName\\\":\n  \\\"xds-grpc-cilium\\\"\\r\\n            }\\r\\n          }\\r\\n        ],\\r\\n        \\\"setNodeOnFirstMessageOnly\\\":\n  true\\r\\n      },\\r\\n      \\\"resourceApiVersion\\\": \\\"V3\\\"\\r\\n    },\\r\\n    \\\"cdsConfig\\\":\n  {\\r\\n      \\\"apiConfigSource\\\": {\\r\\n        \\\"apiType\\\": \\\"GRPC\\\",\\r\\n        \\\"transportApiVersion\\\":\n  \\\"V3\\\",\\r\\n        \\\"grpcServices\\\": [\\r\\n          {\\r\\n            \\\"envoyGrpc\\\":\n  {\\r\\n              \\\"clusterName\\\": \\\"xds-grpc-cilium\\\"\\r\\n            }\\r\\n          }\\r\\n\n  \\       ],\\r\\n        \\\"setNodeOnFirstMessageOnly\\\": true\\r\\n      },\\r\\n      \\\"resourceApiVersion\\\":\n  \\\"V3\\\"\\r\\n    }\\r\\n  },\\r\\n  \\\"bootstrapExtensions\\\": [\\r\\n    {\\r\\n      \\\"name\\\":\n  \\\"envoy.bootstrap.internal_listener\\\",\\r\\n      \\\"typed_config\\\": {\\r\\n        \\\"@type\\\":\n  \\\"type.googleapis.com/envoy.extensions.bootstrap.internal_listener.v3.InternalListener\\\"\\r\\n\n  \\     }\\r\\n    }\\r\\n  ],\\r\\n  \\\"layeredRuntime\\\": {\\r\\n    \\\"layers\\\": [\\r\\n      {\\r\\n\n  \\       \\\"name\\\": \\\"static_layer_0\\\",\\r\\n        \\\"staticLayer\\\": {\\r\\n          \\\"overload\\\":\n  {\\r\\n            \\\"global_downstream_max_connections\\\": 50000\\r\\n          }\\r\\n\n  \\       }\\r\\n      }\\r\\n    ]\\r\\n  },\\r\\n  \\\"admin\\\": {\\r\\n    \\\"address\\\": {\\r\\n\n  \\     \\\"pipe\\\": {\\r\\n        \\\"path\\\": \\\"/var/run/cilium/envoy/sockets/admin.sock\\\"\\r\\n\n  \\     }\\r\\n    }\\r\\n  }\\r\\n}\\r\\n\"": template: gotpl:27: unexpected "\\" in operand

Use --debug flag to render out invalid YAML
```

## `template.missing_template` (19)

### 1. `open-edge-platform/edge-ai-libraries`

- Chart: `D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer-core\chart`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\open-edge-platform__edge-ai-libraries\sample-applications\chat-question-and-answer-core\chart`

```text
Error: chat-question-and-answer-core/templates/deployment.yaml:1:4
  executing "chat-question-and-answer-core/templates/deployment.yaml" at <include "chatqna-core.validateGpuSettings" .>:
    error calling include:
chat-question-and-answer-core/templates/_helpers.tpl:30:47
  executing "chatqna-core.validateGpuSettings" at <lower>:
    invalid value; expected string

Use --debug flag to render out invalid YAML
```

### 2. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\logging\elasticsearch`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\logging\elasticsearch`

```text
Error: elasticsearch/templates/statefulset.yaml:16:24
  executing "elasticsearch/templates/statefulset.yaml" at <include "esMajorVersion" .>:
    error calling include:
elasticsearch/templates/_helpers.tpl:42:57
  executing "esMajorVersion" at <".">:
    invalid value; expected string

Use --debug flag to render out invalid YAML
```

### 3. `tmforum-oda/oda-canvas`

- Chart: `D:\helm_clones_github\tmforum-oda__oda-canvas\charts\cert-manager-init`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tmforum-oda__oda-canvas\charts\cert-manager-init`

```text
Error: cert-manager-init/templates/webhook.yaml:24:20
  executing "cert-manager-init/templates/webhook.yaml" at <include "docker.registry" .>:
    error calling include:
template: no template "docker.registry" associated with template "gotpl"

Use --debug flag to render out invalid YAML
```

### 4. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\lke-ti-gateway`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\lke-ti-gateway`

```text
Error: lke-ti-gateway/templates/deployment.yaml:83:17
  executing "lke-ti-gateway/templates/deployment.yaml" at <include "needInitContainer" .>:
    error calling include:
template: no template "needInitContainer" associated with template "gotpl"

Use --debug flag to render out invalid YAML
```

### 5. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp\charts\ti-common`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp\charts\ti-common`

```text
Error: ti-common/templates/serviceaccount_reader.yaml:1:11
  executing "ti-common/templates/serviceaccount_reader.yaml" at <include "needInitContainer" .>:
    error calling include:
template: no template "needInitContainer" associated with template "gotpl"

Use --debug flag to render out invalid YAML
```

### 6. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\lke-ti-gateway`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\lke-ti-gateway`

```text
Error: lke-ti-gateway/templates/deployment.yaml:83:17
  executing "lke-ti-gateway/templates/deployment.yaml" at <include "needInitContainer" .>:
    error calling include:
template: no template "needInitContainer" associated with template "gotpl"

Use --debug flag to render out invalid YAML
```

### 7. `tkestack/charts`

- Chart: `D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\ti-common`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tkestack__charts\incubator\adp-intl\charts\ti-common`

```text
Error: ti-common/templates/serviceaccount_reader.yaml:1:11
  executing "ti-common/templates/serviceaccount_reader.yaml" at <include "needInitContainer" .>:
    error calling include:
template: no template "needInitContainer" associated with template "gotpl"

Use --debug flag to render out invalid YAML
```

### 8. `mojaloop/helm`

- Chart: `D:\helm_clones_github\mojaloop__helm\bulk-api-adapter\chart-handler-notification`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\mojaloop__helm\bulk-api-adapter\chart-handler-notification`

```text
Error: template: bulk-api-adapter-handler-notification/templates/deployment.yaml:59:14: executing "bulk-api-adapter-handler-notification/templates/deployment.yaml" at <include "common.tplvalues.render" (dict "value" .Values.initContainers "context" $)>: error calling include: template: bulk-api-adapter-handler-notification/charts/common/templates/_tplvalues.tpl:9:12: executing "common.tplvalues.render" at <tpl .value .context>: error calling tpl: error during tpl function execution for "{{- include \"mojaloop-common.waitForKafkaInitContainer\" . | nindent 2 }}\n{{- include \"mojaloop-common.waitForMongodbInitContainer\" . | nindent 2 }}\n": template: gotpl:1:4: executing "gotpl" at <include "mojaloop-common.waitForKafkaInitContainer" .>: error calling include: template: no template "mojaloop-common.waitForKafkaInitContainer" associated with template "gotpl"

Use --debug flag to render out invalid YAML
```

### 9. `mojaloop/helm`

- Chart: `D:\helm_clones_github\mojaloop__helm\bulk-api-adapter\chart-service`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\mojaloop__helm\bulk-api-adapter\chart-service`

```text
Error: template: bulk-api-adapter-service/templates/deployment.yaml:50:14: executing "bulk-api-adapter-service/templates/deployment.yaml" at <include "common.tplvalues.render" (dict "value" .Values.initContainers "context" $)>: error calling include: template: bulk-api-adapter-service/charts/common/templates/_tplvalues.tpl:9:12: executing "common.tplvalues.render" at <tpl .value .context>: error calling tpl: error during tpl function execution for "{{- include \"mojaloop-common.waitForKafkaInitContainer\" . | nindent 2 }}\n{{- include \"mojaloop-common.waitForMongodbInitContainer\" . | nindent 2 }}\n": template: gotpl:1:4: executing "gotpl" at <include "mojaloop-common.waitForKafkaInitContainer" .>: error calling include: template: no template "mojaloop-common.waitForKafkaInitContainer" associated with template "gotpl"

Use --debug flag to render out invalid YAML
```

### 10. `mojaloop/helm`

- Chart: `D:\helm_clones_github\mojaloop__helm\ml-api-adapter\chart-handler-notification`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\mojaloop__helm\ml-api-adapter\chart-handler-notification`

```text
Error: ml-api-adapter-handler-notification/templates/deployment.yaml:99:14
  executing "ml-api-adapter-handler-notification/templates/deployment.yaml" at <include "mojaloop-common.probes" .>:
    error calling include:
template: no template "mojaloop-common.probes" associated with template "gotpl"

Use --debug flag to render out invalid YAML
```

### 11. `mojaloop/helm`

- Chart: `D:\helm_clones_github\mojaloop__helm\ml-api-adapter\chart-service`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\mojaloop__helm\ml-api-adapter\chart-service`

```text
Error: ml-api-adapter-service/templates/deployment.yaml:92:14
  executing "ml-api-adapter-service/templates/deployment.yaml" at <include "mojaloop-common.probes" .>:
    error calling include:
template: no template "mojaloop-common.probes" associated with template "gotpl"

Use --debug flag to render out invalid YAML
```

### 12. `oneconcern/datamon`

- Chart: `D:\helm_clones_github\oneconcern__datamon\k8s\purge\build-index`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oneconcern__datamon\k8s\purge\build-index`

```text
Error: build-datamon-index/templates/job.yaml:20:28
  executing "build-datamon-index/templates/job.yaml" at <include (print $.Template.BasePath "/config.yaml") .>:
    error calling include:
build-datamon-index/templates/config.yaml:12:27
  executing "build-datamon-index/templates/config.yaml" at <4>:
    invalid value; expected string

Use --debug flag to render out invalid YAML
```

### 13. `oneconcern/datamon`

- Chart: `D:\helm_clones_github\oneconcern__datamon\k8s\purge\delete-repo`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oneconcern__datamon\k8s\purge\delete-repo`

```text
Error: datamon-delete/templates/job.yaml:20:28
  executing "datamon-delete/templates/job.yaml" at <include (print $.Template.BasePath "/config.yaml") .>:
    error calling include:
datamon-delete/templates/config.yaml:12:27
  executing "datamon-delete/templates/config.yaml" at <4>:
    invalid value; expected string

Use --debug flag to render out invalid YAML
```

### 14. `oneconcern/datamon`

- Chart: `D:\helm_clones_github\oneconcern__datamon\k8s\purge\delete-unused`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oneconcern__datamon\k8s\purge\delete-unused`

```text
Error: delete-datamon-unused/templates/job.yaml:20:28
  executing "delete-datamon-unused/templates/job.yaml" at <include (print $.Template.BasePath "/config.yaml") .>:
    error calling include:
delete-datamon-unused/templates/config.yaml:12:27
  executing "delete-datamon-unused/templates/config.yaml" at <4>:
    invalid value; expected string

Use --debug flag to render out invalid YAML
```

### 15. `oneconcern/datamon`

- Chart: `D:\helm_clones_github\oneconcern__datamon\k8s\purge\squash`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oneconcern__datamon\k8s\purge\squash`

```text
Error: build-datamon-index/templates/job.yaml:20:28
  executing "build-datamon-index/templates/job.yaml" at <include (print $.Template.BasePath "/config.yaml") .>:
    error calling include:
build-datamon-index/templates/config.yaml:12:27
  executing "build-datamon-index/templates/config.yaml" at <4>:
    invalid value; expected string

Use --debug flag to render out invalid YAML
```

### 16. `kast-spells/kast-system`

- Chart: `D:\helm_clones_github\kast-spells__kast-system\charts\summon`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\kast-spells__kast-system\charts\summon`

```text
Error: summon/templates/summon.yaml:39:4
  executing "summon/templates/summon.yaml" at <include (printf "summon.workload.%s" .Values.workload.type) .>:
    error calling include:
template: no template "summon.workload.deployment" associated with template "gotpl"

Use --debug flag to render out invalid YAML
```

### 17. `kast-spells/kast-system`

- Chart: `D:\helm_clones_github\kast-spells__kast-system\charts\trinkets\microspell`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\kast-spells__kast-system\charts\trinkets\microspell`

```text
Error: microspell/templates/microservice.yaml:10:30
  executing "microspell/templates/microservice.yaml" at <include "common.name" $root>:
    error calling include:
template: no template "common.name" associated with template "gotpl"

Use --debug flag to render out invalid YAML
```

### 18. `kast-spells/kast-system`

- Chart: `D:\helm_clones_github\kast-spells__kast-system\charts\trinkets\tarot`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\kast-spells__kast-system\charts\trinkets\tarot`

```text
Error: tarot/templates/workflow.yaml:83:19
  executing "tarot/templates/workflow.yaml" at <include "tarot.workflowName" .>:
    error calling include:
tarot/templates/_helpers.tpl:43:18
  executing "tarot.workflowName" at <include "common.name" .>:
    error calling include:
template: no template "common.name" associated with template "gotpl"

Use --debug flag to render out invalid YAML
```

### 19. `helverinio/misw4406-14-desacopla2`

- Chart: `D:\helm_clones_github\helverinio__misw4406-14-desacopla2\helm\shared-gateway`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\helverinio__misw4406-14-desacopla2\helm\shared-gateway`

```text
Error: shared-gateway/templates/gateway.yaml:2:17
  executing "shared-gateway/templates/gateway.yaml" at <include "shared-ingress.fullname" .>:
    error calling include:
template: no template "shared-ingress.fullname" associated with template "gotpl"

Use --debug flag to render out invalid YAML
```

## `template.yaml_render` (18)

### 1. `camptocamp/charts`

- Chart: `D:\helm_clones_github\camptocamp__charts\bivac`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\camptocamp__charts\bivac`

```text
Error: YAML parse error on bivac/templates/cronjob.yaml: error converting YAML to JSON: yaml: line 32: could not find expected ':'

Use --debug flag to render out invalid YAML
```

### 2. `Loongson-Cloud-Community/dockerfiles`

- Chart: `D:\helm_clones_github\Loongson-Cloud-Community__dockerfiles\kubesphere\ks-installer\v3.2.1\roles\gatekeeper\files\gatekeeper`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\Loongson-Cloud-Community__dockerfiles\kubesphere\ks-installer\v3.2.1\roles\gatekeeper\files\gatekeeper`

```text
Error: YAML parse error on gatekeeper/templates/gatekeeper-validating-webhook-configuration-validatingwebhookconfiguration.yaml: error converting YAML to JSON: yaml: invalid map key: map[interface {}]interface {}{".Values.validatingWebhookFailurePolicy":interface {}(nil)}

Use --debug flag to render out invalid YAML
```

### 3. `cloud-native-toolkit/toolkit-charts`

- Chart: `D:\helm_clones_github\cloud-native-toolkit__toolkit-charts\stable\assign-group-cronjob`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\cloud-native-toolkit__toolkit-charts\stable\assign-group-cronjob`

```text
Error: YAML parse error on assign-group-cronjob/templates/cronjob.yaml: error converting YAML to JSON: yaml: mapping values are not allowed in this context

Use --debug flag to render out invalid YAML
```

### 4. `jharmison-redhat/openshift-setup`

- Chart: `D:\helm_clones_github\jharmison-redhat__openshift-setup\charts\aws-efs-csi-setup`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\jharmison-redhat__openshift-setup\charts\aws-efs-csi-setup`

```text
Error: YAML parse error on aws-efs-csi-set/templates/csi-driver/storageclass-job.yaml: error converting YAML to JSON: yaml: line 9: did not find expected key

Use --debug flag to render out invalid YAML
```

### 5. `erost/vdz26-demo-fleet-commander`

- Chart: `D:\helm_clones_github\erost__vdz26-demo-fleet-commander\commander\public-function-bucket\chart`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\erost__vdz26-demo-fleet-commander\commander\public-function-bucket\chart`

```text
Error: YAML parse error on public-function-bucket/templates/function.yaml: error converting YAML to JSON: yaml: line 6: mapping values are not allowed in this context

Use --debug flag to render out invalid YAML
```

### 6. `erost/vdz26-demo-fleet-commander`

- Chart: `D:\helm_clones_github\erost__vdz26-demo-fleet-commander\commander\public-function-numbers\chart`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\erost__vdz26-demo-fleet-commander\commander\public-function-numbers\chart`

```text
Error: YAML parse error on public-function-numbers/templates/function.yaml: error converting YAML to JSON: yaml: line 6: mapping values are not allowed in this context

Use --debug flag to render out invalid YAML
```

### 7. `erost/vdz26-demo-fleet-commander`

- Chart: `D:\helm_clones_github\erost__vdz26-demo-fleet-commander\commander\public-function-strings\chart`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\erost__vdz26-demo-fleet-commander\commander\public-function-strings\chart`

```text
Error: YAML parse error on public-function-strings/templates/function.yaml: error converting YAML to JSON: yaml: line 6: mapping values are not allowed in this context

Use --debug flag to render out invalid YAML
```

### 8. `erost/vdz26-demo-fleet-commander`

- Chart: `D:\helm_clones_github\erost__vdz26-demo-fleet-commander\units\unit-aws\function-bucket\chart`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\erost__vdz26-demo-fleet-commander\units\unit-aws\function-bucket\chart`

```text
Error: YAML parse error on function-bucket/templates/function.yaml: error converting YAML to JSON: yaml: line 6: mapping values are not allowed in this context

Use --debug flag to render out invalid YAML
```

### 9. `erost/vdz26-demo-fleet-commander`

- Chart: `D:\helm_clones_github\erost__vdz26-demo-fleet-commander\units\unit-numbers\function-numbers\chart`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\erost__vdz26-demo-fleet-commander\units\unit-numbers\function-numbers\chart`

```text
Error: YAML parse error on function-numbers/templates/function.yaml: error converting YAML to JSON: yaml: line 6: mapping values are not allowed in this context

Use --debug flag to render out invalid YAML
```

### 10. `erost/vdz26-demo-fleet-commander`

- Chart: `D:\helm_clones_github\erost__vdz26-demo-fleet-commander\units\unit-strings\function-strings\chart`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\erost__vdz26-demo-fleet-commander\units\unit-strings\function-strings\chart`

```text
Error: YAML parse error on function-strings/templates/function.yaml: error converting YAML to JSON: yaml: line 6: mapping values are not allowed in this context

Use --debug flag to render out invalid YAML
```

### 11. `ndebuhr/isidro`

- Chart: `D:\helm_clones_github\ndebuhr__isidro\chart`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\ndebuhr__isidro\chart`

```text
Error: YAML parse error on isidro/templates/prometheus.yaml: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type util.SimpleHead

Use --debug flag to render out invalid YAML
```

### 12. `thelande/charts`

- Chart: `D:\helm_clones_github\thelande__charts\charts\opencloud`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\thelande__charts\charts\opencloud`

```text
Error: YAML parse error on opencloud/templates/secret.yaml: error converting YAML to JSON: yaml: line 5: mapping values are not allowed in this context

Use --debug flag to render out invalid YAML
```

### 13. `alercebroker/web-services`

- Chart: `D:\helm_clones_github\alercebroker__web-services\charts\alerts-api`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\alercebroker__web-services\charts\alerts-api`

```text
Error: YAML parse error on alerts-api/templates/ingress.yaml: error converting YAML to JSON: yaml: line 36: found character that cannot start any token

Use --debug flag to render out invalid YAML
```

### 14. `cisco-open/appdynamics-k8s-webhook-instrumentor`

- Chart: `D:\helm_clones_github\cisco-open__appdynamics-k8s-webhook-instrumentor\buildEnv\helm\webhook-instrumentor`
- Source: `template`
- Values files: `D:\helm_clones_github\cisco-open__appdynamics-k8s-webhook-instrumentor\buildEnv\helm\webhook-instrumentor\values-sample.yaml`
- Command: `helm template test D:\helm_clones_github\cisco-open__appdynamics-k8s-webhook-instrumentor\buildEnv\helm\webhook-instrumentor -f D:\helm_clones_github\cisco-open__appdynamics-k8s-webhook-instrumentor\buildEnv\helm\webhook-instrumentor\values-sample.yaml`

```text
Error: YAML parse error on webhook-instrumentor/templates/otel/cm-otel-collector-config.yaml: error converting YAML to JSON: yaml: line 10: mapping values are not allowed in this context

Use --debug flag to render out invalid YAML
```

### 15. `cisco-open/appdynamics-k8s-webhook-instrumentor`

- Chart: `D:\helm_clones_github\cisco-open__appdynamics-k8s-webhook-instrumentor\buildEnv\helm\webhook-instrumentor`
- Source: `template`
- Values files: `D:\helm_clones_github\cisco-open__appdynamics-k8s-webhook-instrumentor\buildEnv\helm\webhook-instrumentor\values-sample-otel.yaml`, `D:\helm_clones_github\cisco-open__appdynamics-k8s-webhook-instrumentor\buildEnv\helm\webhook-instrumentor\values-sample.yaml`
- Command: `helm template test D:\helm_clones_github\cisco-open__appdynamics-k8s-webhook-instrumentor\buildEnv\helm\webhook-instrumentor -f D:\helm_clones_github\cisco-open__appdynamics-k8s-webhook-instrumentor\buildEnv\helm\webhook-instrumentor\values-sample-otel.yaml -f D:\helm_clones_github\cisco-open__appdynamics-k8s-webhook-instrumentor\buildEnv\helm\webhook-instrumentor\values-sample.yaml`

```text
Error: YAML parse error on webhook-instrumentor/templates/otel/cm-otel-collector-config.yaml: error converting YAML to JSON: yaml: line 10: mapping values are not allowed in this context

Use --debug flag to render out invalid YAML
```

### 16. `cisco-open/appdynamics-k8s-webhook-instrumentor`

- Chart: `D:\helm_clones_github\cisco-open__appdynamics-k8s-webhook-instrumentor\webhook\helm\instrumentor`
- Source: `template`
- Values files: `D:\helm_clones_github\cisco-open__appdynamics-k8s-webhook-instrumentor\webhook\helm\instrumentor\values-sample.yaml`
- Command: `helm template test D:\helm_clones_github\cisco-open__appdynamics-k8s-webhook-instrumentor\webhook\helm\instrumentor -f D:\helm_clones_github\cisco-open__appdynamics-k8s-webhook-instrumentor\webhook\helm\instrumentor\values-sample.yaml`

```text
Error: YAML parse error on webhook-instrumentor/templates/otel/cm-otel-collector-config.yaml: error converting YAML to JSON: yaml: line 10: mapping values are not allowed in this context

Use --debug flag to render out invalid YAML
```

### 17. `cisco-open/appdynamics-k8s-webhook-instrumentor`

- Chart: `D:\helm_clones_github\cisco-open__appdynamics-k8s-webhook-instrumentor\webhook\helm\instrumentor`
- Source: `template`
- Values files: `D:\helm_clones_github\cisco-open__appdynamics-k8s-webhook-instrumentor\webhook\helm\instrumentor\values-sample-otel.yaml`, `D:\helm_clones_github\cisco-open__appdynamics-k8s-webhook-instrumentor\webhook\helm\instrumentor\values-sample.yaml`
- Command: `helm template test D:\helm_clones_github\cisco-open__appdynamics-k8s-webhook-instrumentor\webhook\helm\instrumentor -f D:\helm_clones_github\cisco-open__appdynamics-k8s-webhook-instrumentor\webhook\helm\instrumentor\values-sample-otel.yaml -f D:\helm_clones_github\cisco-open__appdynamics-k8s-webhook-instrumentor\webhook\helm\instrumentor\values-sample.yaml`

```text
Error: YAML parse error on webhook-instrumentor/templates/otel/cm-otel-collector-config.yaml: error converting YAML to JSON: yaml: line 10: mapping values are not allowed in this context

Use --debug flag to render out invalid YAML
```

### 18. `sydney900/CQRS-ES-Microservices`

- Chart: `D:\helm_clones_github\sydney900__CQRS-ES-Microservices\Charts\V1`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\sydney900__CQRS-ES-Microservices\Charts\V1`

```text
Error: YAML parse error on V1/templates/ingress.yaml: error converting YAML to JSON: yaml: line 8: mapping values are not allowed in this context

Use --debug flag to render out invalid YAML
```

## `dependency.lock_file_out_of_sync` (7)

### 1. `rancher/charts`

- Chart: `D:\helm_clones_github\rancher__charts\charts\epinio\102.0.1+up1.6.2`
- Source: `dependency`

```text
Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update'
```

### 2. `rancher/partner-charts`

- Chart: `D:\helm_clones_github\rancher__partner-charts\charts\amd\amd-gpu\0.10.0`
- Source: `dependency`

```text
Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update'
```

### 3. `CARV-ICS-FORTH/frisbee`

- Chart: `D:\helm_clones_github\CARV-ICS-FORTH__frisbee\charts\platform`
- Source: `dependency`

```text
Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update'
```

### 4. `cloud-native-toolkit/toolkit-charts`

- Chart: `D:\helm_clones_github\cloud-native-toolkit__toolkit-charts\stable\cloud-setup`
- Source: `dependency`

```text
Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update'
```

### 5. `vlab-research/fly`

- Chart: `D:\helm_clones_github\vlab-research__fly\devops\vlab`
- Source: `dependency`

```text
Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update'
```

### 6. `100CallsToEurop/otus`

- Chart: `D:\helm_clones_github\100CallsToEurop__otus\k8s\auth`
- Source: `dependency`

```text
Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update'
```

### 7. `100rd/platform-design`

- Chart: `D:\helm_clones_github\100rd__platform-design\apps\infra\cilium`
- Source: `dependency`

```text
Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies with 'helm dependency update'
```

## `dependency.network_dns` (7)

### 1. `sapcc/helm-charts`

- Chart: `D:\helm_clones_github\sapcc__helm-charts\common\inventory-updater`
- Source: `dependency`

```text
Saving 1 charts
Downloading owner-info from repo oci://keppel.eu-de-1.cloud.sap/ccloud-helm
Save error occurred:  could not download oci://keppel.eu-de-1.cloud.sap/ccloud-helm/owner-info: failed to perform "FetchReference" on source: Get "https://keppel.eu-de-1.cloud.sap/v2/ccloud-helm/owner-info/manifests/0.2.0": dial tcp: lookup keppel.eu-de-1.cloud.sap: getaddrinfow: Este é geralmente um erro temporário durante a resolução de nomes de anfitrião e significa que o servidor local não recebeu uma resposta de um servidor autoritário.
Error: could not download oci://keppel.eu-de-1.cloud.sap/ccloud-helm/owner-info: failed to perform "FetchReference" on source: Get "https://keppel.eu-de-1.cloud.sap/v2/ccloud-helm/owner-info/manifests/0.2.0": dial tcp: lookup keppel.eu-de-1.cloud.sap: getaddrinfow: Este é geralmente um erro temporário durante a resolução de nomes de anfitrião e significa que o servidor local não recebeu uma resposta de um servidor autoritário.
```

### 2. `Loongson-Cloud-Community/dockerfiles`

- Chart: `D:\helm_clones_github\Loongson-Cloud-Community__dockerfiles\kubesphere\ks-installer\v3.2.1\roles\ks-multicluster\files\kubefed\kubefed`
- Source: `dependency`

```text
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "https://localhost/" chart repository:
	Get "https://localhost/index.yaml": dial tcp [::1]:443: connectex: Nenhuma ligação pôde ser feita porque o computador de destino
as recusou ativamente.
Error: no cached repository for helm-manager-f2b99ce05b94599549c70dbbe7a891b278e7c3cacad02334fa44682fca36c740 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-f2b99ce05b94599549c70dbbe7a891b278e7c3cacad02334fa44682fca36c740-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### 3. `stakater/nordmart-apps-gitops-config`

- Chart: `D:\helm_clones_github\stakater__nordmart-apps-gitops-config\01-arsenal\01-stakater-nordmart-review-api\01-dev`
- Source: `dependency`

```text
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "https://nexus-helm-stakater-nexus.apps.devtest.vxdqgl7u.kubeapp.cloud/repository/helm-charts/" chart repository:
	Get "https://nexus-helm-stakater-nexus.apps.devtest.vxdqgl7u.kubeapp.cloud/repository/helm-charts/index.yaml": dial tcp: lookup nexus-helm-stakater-nexus.apps.devtest.vxdqgl7u.kubeapp.cloud: no such host
Error: no cached repository for helm-manager-7ebb8ed6883774d2c679cf4b093eaa1b7bd49e3f4401e7427ef1456c3315f23d found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-7ebb8ed6883774d2c679cf4b093eaa1b7bd49e3f4401e7427ef1456c3315f23d-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### 4. `starlingx/openstack-armada-app`

- Chart: `D:\helm_clones_github\starlingx__openstack-armada-app\stx-openstack-helm-fluxcd\stx-openstack-helm-fluxcd\helm-charts\clients`
- Source: `dependency`

```text
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "http://localhost:8879/charts" chart repository:
	Get "http://localhost:8879/charts/index.yaml": dial tcp [::1]:8879: connectex: Nenhuma ligação pôde ser feita porque o computador de destino
as recusou ativamente.
Error: no cached repository for helm-manager-878d619eb15837b169144dfaab3a7d6c5e800dd40daf0369bbe2b101f2275284 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-878d619eb15837b169144dfaab3a7d6c5e800dd40daf0369bbe2b101f2275284-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### 5. `KevMCarp/truecharts-catalog-fork`

- Chart: `D:\helm_clones_github\KevMCarp__truecharts-catalog-fork\dependency\clickhouse\5.0.54`
- Source: `dependency`

```text
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "https://library-charts.truecharts.org" chart repository:
	Get "https://library-charts.truecharts.org/index.yaml": dial tcp: lookup library-charts.truecharts.org: no such host
Error: no cached repository for helm-manager-024b189b59f6c6ccf0de6e5148db1578caf551c511f4eb220ece14cef00f80e0 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-024b189b59f6c6ccf0de6e5148db1578caf551c511f4eb220ece14cef00f80e0-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### 6. `MrE-Fog/ks-installer2`

- Chart: `D:\helm_clones_github\MrE-Fog__ks-installer2\roles\ks-multicluster\files\kubefed\kubefed`
- Source: `dependency`

```text
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "https://localhost/" chart repository:
	Get "https://localhost/index.yaml": dial tcp [::1]:443: connectex: Nenhuma ligação pôde ser feita porque o computador de destino
as recusou ativamente.
Error: no cached repository for helm-manager-f2b99ce05b94599549c70dbbe7a891b278e7c3cacad02334fa44682fca36c740 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-f2b99ce05b94599549c70dbbe7a891b278e7c3cacad02334fa44682fca36c740-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### 7. `Clark1992/ECK1`

- Chart: `D:\helm_clones_github\Clark1992__ECK1\src\Integration\ECK1.FailedViewRebuilder\Deploy\service`
- Source: `dependency`

```text
Saving 1 charts
Downloading config-watcher from repo oci://localhost:5000/helm
Save error occurred:  could not download oci://localhost:5000/helm/config-watcher: failed to perform "FetchReference" on source: Get "https://localhost:5000/v2/helm/config-watcher/manifests/0.1.0": dial tcp [::1]:5000: connectex: Nenhuma ligação pôde ser feita porque o computador de destino
as recusou ativamente.
Error: could not download oci://localhost:5000/helm/config-watcher: failed to perform "FetchReference" on source: Get "https://localhost:5000/v2/helm/config-watcher/manifests/0.1.0": dial tcp [::1]:5000: connectex: Nenhuma ligação pôde ser feita porque o computador de destino
as recusou ativamente.
```

## `dependency.repo_update` (7)

### 1. `linode/apl-core`

- Chart: `D:\helm_clones_github\linode__apl-core\chart\chart-index`
- Source: `dependency`

```text
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

### 2. `k8s-home-lab/helm-charts`

- Chart: `D:\helm_clones_github\k8s-home-lab__helm-charts\unmaintained\audiobookshelf`
- Source: `dependency`

```text
Getting updates for unmanaged Helm repositories...
...Successfully got an update from the "https://library-charts.k8s-at-home.com" chart repository
Error: can't get a valid version for 1 subchart(s): "common" (repository "https://library-charts.k8s-at-home.com", version "4.5.3"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml
```

### 3. `devtron-labs/charts`

- Chart: `D:\helm_clones_github\devtron-labs__charts\charts\cluster-essentials`
- Source: `dependency`

```text
Getting updates for unmanaged Helm repositories...
...Successfully got an update from the "https://kedacore.github.io/charts" chart repository
...Successfully got an update from the "https://kubernetes.github.io/autoscaler" chart repository
...Successfully got an update from the "https://helm.devtron.ai" chart repository
...Successfully got an update from the "https://aws.github.io/eks-charts" chart repository
...Successfully got an update from the "https://kubernetes-sigs.github.io/metrics-server/" chart repository
...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository
Error: can't get a valid version for 1 subchart(s): "kubernetes-event-exporter" (repository "https://charts.bitnami.com/bitnami", version "1.2.*"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml
```

### 4. `dungdm93/shipyard`

- Chart: `D:\helm_clones_github\dungdm93__shipyard\helm\druid`
- Source: `dependency`

```text
Getting updates for unmanaged Helm repositories...
...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository
Error: can't get a valid version for 2 subchart(s): "zookeeper" (repository "https://charts.bitnami.com/bitnami", version "7.x.x"), "postgresql" (repository "https://charts.bitnami.com/bitnami", version "10.x.x"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml
```

### 5. `lucidworks/ocp-fusion-helm-charts`

- Chart: `D:\helm_clones_github\lucidworks__ocp-fusion-helm-charts\5.3.4\fusion\charts\admin-ui`
- Source: `dependency`

```text
Getting updates for unmanaged Helm repositories...
...Successfully got an update from the "https://charts.lucidworks.com/" chart repository
Error: can't get a valid version for 1 subchart(s): "fusion-common-utils" (repository "https://charts.lucidworks.com/", version "1.5.1"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml
```

### 6. `FIWARE-Ops/fiware-gitops`

- Chart: `D:\helm_clones_github\FIWARE-Ops__fiware-gitops\aws\token\mongodb`
- Source: `dependency`

```text
Getting updates for unmanaged Helm repositories...
...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository
Error: can't get a valid version for 1 subchart(s): "mongodb" (repository "https://charts.bitnami.com/bitnami", version "11.0.4"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml
```

### 7. `hmcts/hmcts-charts`

- Chart: `D:\helm_clones_github\hmcts__hmcts-charts\stable\aac-manage-case-assignment`
- Source: `dependency`

```text
Getting updates for unmanaged Helm repositories...
...Successfully got an update from the "https://helm.elastic.co" chart repository
Saving 4 charts
Downloading java from repo oci://hmctsprod.azurecr.io/helm
Save error occurred:  could not download oci://hmctsprod.azurecr.io/helm/java: failed to perform "FetchReference" on source: GET "https://hmctsprod.azurecr.io/v2/helm/java/manifests/5.3.0": GET "https://hmctsprod.azurecr.io/oauth2/token?scope=repository%3Ahelm%2Fjava%3Apull&service=hmctsprod.azurecr.io": response status code 401: unauthorized: authentication required, visit https://aka.ms/acr/authorization for more information. CorrelationId: f65eaf6b-2681-42ff-b433-3789c00e163f
Error: could not download oci://hmctsprod.azurecr.io/helm/java: failed to perform "FetchReference" on source: GET "https://hmctsprod.azurecr.io/v2/helm/java/manifests/5.3.0": GET "https://hmctsprod.azurecr.io/oauth2/token?scope=repository%3Ahelm%2Fjava%3Apull&service=hmctsprod.azurecr.io": response status code 401: unauthorized: authentication required, visit https://aka.ms/acr/authorization for more information. CorrelationId: f65eaf6b-2681-42ff-b433-3789c00e163f
```

## `template.type_mismatch` (6)

### 1. `project-sunbird/sunbird-devops`

- Chart: `D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\nginx-private-ingress`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\project-sunbird__sunbird-devops\kubernetes\helm_charts\core\nginx-private-ingress`

```text
Error: nginx-private-ingress/templates/secrets.yaml:9:34
  executing "nginx-private-ingress/templates/secrets.yaml" at <b64enc>:
    invalid value; expected string

Use --debug flag to render out invalid YAML
```

### 2. `tmforum-oda/oda-canvas`

- Chart: `D:\helm_clones_github\tmforum-oda__oda-canvas\charts\canvas-namespaces`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tmforum-oda__oda-canvas\charts\canvas-namespaces`

```text
Error: canvas-namespaces/templates/namespace-default.yaml:1:49
  executing "canvas-namespaces/templates/namespace-default.yaml" at <.Values.componentNamespace>:
    wrong type for value; expected string; got interface {}

Use --debug flag to render out invalid YAML
```

### 3. `tmforum-oda/oda-canvas`

- Chart: `D:\helm_clones_github\tmforum-oda__oda-canvas\charts\credentialsmanagement-operator`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\tmforum-oda__oda-canvas\charts\credentialsmanagement-operator`

```text
Error: credentialsmanagement-operator/templates/secret.yaml:10:56
  executing "credentialsmanagement-operator/templates/secret.yaml" at <b64enc>:
    invalid value; expected string

Use --debug flag to render out invalid YAML
```

### 4. `ODIM-Project/ODIM`

- Chart: `D:\helm_clones_github\ODIM-Project__ODIM\odim-controller\helmcharts\odimra-secret`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\ODIM-Project__ODIM\odim-controller\helmcharts\odimra-secret`

```text
Error: odimra-secret/templates/secret.yaml:11:44
  executing "odimra-secret/templates/secret.yaml" at <b64enc>:
    invalid value; expected string

Use --debug flag to render out invalid YAML
```

### 5. `oneconcern/datamon`

- Chart: `D:\helm_clones_github\oneconcern__datamon\k8s\migratev2`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\oneconcern__datamon\k8s\migratev2`

```text
Error: migratev2/templates/secret.yaml:7:18
  executing "migratev2/templates/secret.yaml" at <.Values.credentials>:
    wrong type for value; expected string; got interface {}

Use --debug flag to render out invalid YAML
```

### 6. `AlexanderBabel/helm-charts`

- Chart: `D:\helm_clones_github\AlexanderBabel__helm-charts\charts\dendrite`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\AlexanderBabel__helm-charts\charts\dendrite`

```text
Error: dendrite/templates/dendrite-config.yaml:5:153
  executing "dendrite/templates/dendrite-config.yaml" at <.Values.database.connection_string>:
    wrong type for value; expected string; got interface {}

Use --debug flag to render out invalid YAML
```

## `dependency.chart_validation` (5)

### 1. `stackrox/helm-charts`

- Chart: `D:\helm_clones_github\stackrox__helm-charts\3.0.41.0`
- Source: `dependency`

```text
Error: validation: chart.metadata.version "3.0.41.0" is invalid
```

### 2. `boozallen/aissemble`

- Chart: `D:\helm_clones_github\boozallen__aissemble\foundation\foundation-archetype\src\main\resources\archetype-resources\__rootArtifactId__-deploy\src\main\resources\apps\common-infrastructure`
- Source: `dependency`

```text
Error: validation: chart.metadata.version "${version}" is invalid
```

### 3. `kaikodata/canton-tooling`

- Chart: `D:\helm_clones_github\kaikodata__canton-tooling\kubernetes\templates\canton-validator-template`
- Source: `dependency`

```text
Error: validation: chart.metadata.version "TEMPLATE_VERSION" is invalid
```

### 4. `LuukHors/homelab`

- Chart: `D:\helm_clones_github\LuukHors__homelab\products\_base`
- Source: `dependency`

```text
Error: validation: chart.metadata.name is required
```

### 5. `codefuturist/helm-charts`

- Chart: `D:\helm_clones_github\codefuturist__helm-charts\templates\chart-template`
- Source: `dependency`

```text
Error: validation: chart.metadata.name is required
```

## `dependency.version_resolution` (3)

### 1. `mojaloop/helm`

- Chart: `D:\helm_clones_github\mojaloop__helm\perf-test-harness`
- Source: `dependency`

```text
Error: can't get a valid version for 1 subchart(s): "ml-testing-toolkit-cli" (repository "file://../ml-testing-toolkit-cli", version "15.9.0"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml
```

### 2. `llajas/homelab`

- Chart: `D:\helm_clones_github\llajas__homelab\apps\plex-apps`
- Source: `dependency`

```text
Error: can't get a valid version for 1 subchart(s): "overseerr" (repository "file://./charts/overseerr", version "5.4.2"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml
```

### 3. `NeuraLegion/helmcharts`

- Chart: `D:\helm_clones_github\NeuraLegion__helmcharts\charts\altoroj`
- Source: `dependency`

```text
Saving 1 charts
Save error occurred:  can't get a valid version for dependency simple-service
Error: can't get a valid version for dependency simple-service
```

## `template.invalid_value` (3)

### 1. `ODIM-Project/ODIM`

- Chart: `D:\helm_clones_github\ODIM-Project__ODIM\odim-controller\helmcharts\etcd-ha`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\ODIM-Project__ODIM\odim-controller\helmcharts\etcd-ha`

```text
Error: template: etcd-ha/templates/etcd-ha-deployment.yaml:40:29: executing "etcd-ha/templates/etcd-ha-deployment.yaml" at <index .Values.hostname 0>: error calling index: index of untyped nil

Use --debug flag to render out invalid YAML
```

### 2. `opea-project/Enterprise-Inference`

- Chart: `D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\opea-project__Enterprise-Inference\core\helm-charts\vllm`

```text
Error: template: vllm/templates/deployment.yaml:135:33: executing "vllm/templates/deployment.yaml" at <index .Values.modelConfigs $modelName>: error calling index: index of untyped nil

Use --debug flag to render out invalid YAML
```

### 3. `stfc/cloud-helm-charts`

- Chart: `D:\helm_clones_github\stfc__cloud-helm-charts\charts\stfc-cloud-openstack-cluster`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\stfc__cloud-helm-charts\charts\stfc-cloud-openstack-cluster`

```text
Error: template: stfc-cloud-openstack-cluster/templates/credentials-secret.yaml:18:12: executing "stfc-cloud-openstack-cluster/templates/credentials-secret.yaml" at <index .Values.clouds $cloudName>: error calling index: index of untyped nil

Use --debug flag to render out invalid YAML
```

## `dependency.unpack_error` (2)

### 1. `cozystack/cozystack`

- Chart: `D:\helm_clones_github\cozystack__cozystack\packages\apps\bucket`
- Source: `dependency`

```text
Error: error unpacking subchart cozy-lib in bucket: Chart.yaml file is missing
```

### 2. `kast-spells/kast-system`

- Chart: `D:\helm_clones_github\kast-spells__kast-system\librarian`
- Source: `dependency`

```text
Error: error unpacking subchart common in librarian: Chart.yaml file is missing
```

## `dependency.cache_index_missing` (1)

### 1. `camptocamp/charts`

- Chart: `D:\helm_clones_github\camptocamp__charts\common-build-code`
- Source: `dependency`

```text
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "https://kubernetes-charts-incubator.storage.googleapis.com" chart repository:
	failed to fetch https://kubernetes-charts-incubator.storage.googleapis.com/index.yaml : 403 Forbidden
Error: no cached repository for helm-manager-53271637451a5b2439ffd0af71673734b808e371a8a6aed9bf100a8f219a3006 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-53271637451a5b2439ffd0af71673734b808e371a8a6aed9bf100a8f219a3006-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

## `template.parse_error` (1)

### 1. `ApasoftTraining/cursoHelm`

- Chart: `D:\helm_clones_github\ApasoftTraining__cursoHelm\Comentarios`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\ApasoftTraining__cursoHelm\Comentarios`

```text
Error: parse error at (If/templates/Comentarios.txt:3): function "CONDICION" not defined

Use --debug flag to render out invalid YAML
```

## `template.values_merge_error` (1)

### 1. `trueforge-org/truecharts`

- Chart: `D:\helm_clones_github\trueforge-org__truecharts\charts\stable\clamav`
- Source: `template`
- Command: `helm template test D:\helm_clones_github\trueforge-org__truecharts\charts\stable\clamav`

```text
Error: execution error at (clamav/templates/common.yaml:13:3): Chart - Values contain an error that may be a result of merging. Make sure you don't have any invalid YAML characters starting a value. 
 Renderd Values containing the error: 

 ============================================================================================= 
 TZ: UTC
addons:
  codeserver:
    container:
      enabled: true
      env:
        DEFAULT_WORKSPACE: /
        PORT: 12321
      imageSelector: codeserverImage
      probes:
        liveness:
          enabled: true
          path: /
          port: 12321
        readiness:
          enabled: true
          path: /
          port: 12321
        startup:
          enabled: true
          path: /
          port: 12321
      resources:
        excludeExtra: true
      securityContext:
        readOnlyRootFilesystem: false
        runAsGroup: 0
        runAsNonRoot: false
        runAsUser: 0
      targetSelector:
      - main
    enabled: false
    ingress:
      annotations: {}
      enabled: false
      hosts:
      - host: code.chart-example.local
        paths:
        - path: /
          pathType: Prefix
      labels: {}
      tls: []
    service:
      enabled: true
      ports:
        codeserver:
          enabled: true
          port: 12321
          primary: true
          protocol: http
          targetPort: 12321
      type: ClusterIP
  gluetun:
    container:
      enabled: true
      env:
        DNS_KEEP_NAMESERVER: "on"
        DOT: "off"
        FIREWALL: "off"
        FIREWALL_INPUT_PORTS: ""
        FIREWALL_OUTBOUND_SUBNETS: ""
      imageSelector: gluetunImage
      probes:
        liveness:
          enabled: false
        readiness:
          enabled: false
        startup:
          enabled: false
      resources:
        excludeExtra: true
      securityContext:
        capabilities:
          add:
          - NET_ADMIN
          - NET_RAW
          - MKNOD
        readOnlyRootFilesystem: false
        runAsGroup: 568
        runAsNonRoot: false
        runAsUser: 0
    enabled: false
    targetSelector:
    - main
  netshoot:
    container:
      command:
      - /bin/sh
      - -c
      - sleep infinity
      enabled: true
      imageSelector: netshootImage
      probes:
        liveness:
          enabled: false
        readiness:
          enabled: false
        startup:
          enabled: false
      resources:
        excludeExtra: true
      securityContext:
        capabilities:
          add:
          - NET_ADMIN
          - NET_RAW
        readOnlyRootFilesystem: false
        runAsGroup: 0
        runAsNonRoot: false
        runAsUser: 0
    enabled: false
  tailscale:
    accept_dns: false
    annotations: {}
    auth_once: true
    authkey: ""
    config: ""
    container:
      command:
      - /usr/local/bin/containerboot
      enabled: true
      env:
        TS_ACCEPT_DNS: false
        TS_AUTH_KEY: ""
        TS_AUTH_ONCE: true
        TS_DEST_IP: ""
        TS_EXTRA_ARGS: ""
        TS_KUBE_SECRET: ""
        TS_OUTBOUND_HTTP_PROXY_LISTEN: ""
        TS_ROUTES: ""
        TS_SOCKET: /var/run/tailscale/tailscaled.sock
        TS_SOCKS5_SERVER: ""
        TS_STATE_DIR: /var/lib/tailscale/state
        TS_TAILSCALED_EXTRA_ARGS: ""
        TS_USERSPACE: true
      imageSelector: tailscaleImage
      probes:
        liveness:
          enabled: false
        readiness:
          enabled: false
        startup:
          enabled: false
      resources:
        excludeExtra: true
      securityContext:
        capabilities:
          add:
          - NET_ADMIN
          - NET_RAW
    daemon_extra_args: ""
    dest_ip: ""
    enabled: false
    extra_args: ""
    outbound_http_proxy_listen: ""
    routes: ""
    sock5_server: ""
    targetSelector:
    - main
    userspace: true
alpineImage:
  pullPolicy: IfNotPresent
  repository: oci.trueforge.org/tccr/alpine
  tag: v3.22.1@sha256:6dc807ae4f2867cb2d00d061f8f579f1966420ad792c179ac68072ab235109f8
certificate: {}
chartContext:
  appHost: 127.0.0.1
  appPath: /
  appPort: "3310"
  appProtocol: http
  appUrl: http://127.0.0.1:3310
  appUrlWithPortAndPath: http://127.0.0.1:3310/
  podCIDR: 172.16.0.0/16
  svcCIDR: 172.17.0.0/16
clamav:
  cron_enabled: true
  cron_schedule: '* * * * *'
  date_format: +%m-%d-%Y_%H.%M.%S
  extra_args: ""
  log_file_name: clamscan_report
  report_path: /logs
clickhouse:
  creds: {}
  enabled: false
  includeCommon: false
  password: PLACEHOLDERPASSWORD
cnpg:
  main:
    annotations: {}
    backups:
      credentials: ""
      destinationPath: ""
      enabled: false
      encryption:
        enabled: false
      manualBackups: []
      retentionPolicy: 30d
      scheduledBackups:
      - backupOwnerReference: self
        immediate: true
        name: daily-backup
        schedule: 0 0 0 * * *
        suspend: false
      servername: ""
      target: ""
    cluster:
      annotations: {}
      env: {}
      envFrom: {}
      initdb: {}
      instances: 2
      labels: {}
      logLevel: info
      primaryUpdateMethod: switchover
      primaryUpdateStrategy: unsupervised
      singleNode: false
    creds: {}
    database: app
    enabled: false
    hibernate: false
    labels: {}
    mode: standalone
    monitoring:
      customQueries: []
      disableDefaultQueries: false
      enablePodMonitor: false
    password: PLACEHOLDERPASSWORD
    pgVersion: 16
    pooler:
      annotations: {}
      createRO: false
      enabled: false
      instances: 2
      labels: {}
      poolMode: session
    primary: true
    recovery:
      backupName: ""
      clusterName: ""
      destinationPath: ""
      method: object_store
      pitrTarget:
        time: ""
      servername: ""
    type: postgres
    user: app
codeserverImage:
  pullPolicy: IfNotPresent
  repository: oci.trueforge.org/containerforge/code-server
  tag: 4.118.0@sha256:b02714ac2d0ee60df7dbcb2aa8b323779d97b71d2d40deeb412831a9530f7a9c
configmap: {}
containerOptions:
  NVIDIA_CAPS:
  - all
credentials: {}
cronjob:
  annotations: {}
  failedJobsHistoryLimit: 5
  successfulJobsHistoryLimit: 2
diagnosticMode:
  enabled: false
extraTpl: []
global:
  annotations: {}
  diagnosticMode:
    enabled: false
  fallbackDefaults:
    accessModes:
    - ReadWriteOnce
    cnpg:
      pgVersion: 16
      skipEmptyWalArchiveCheck: true
    persistenceType: pvc
    probeTimeouts:
      liveness:
        failureThreshold: 5
        initialDelaySeconds: 12
        periodSeconds: 15
        successThreshold: 1
        timeoutSeconds: 5
      readiness:
        failureThreshold: 4
        initialDelaySeconds: 10
        periodSeconds: 12
        successThreshold: 2
        timeoutSeconds: 5
      startup:
        failureThreshold: 60
        initialDelaySeconds: 10
        periodSeconds: 5
        successThreshold: 1
        timeoutSeconds: 3
    probeType: http
    pvcRetain: false
    pvcSize: 100Gi
    serviceProtocol: tcp
    serviceType: ClusterIP
    topologyKey: kubernetes.io/hostname
    vctAccessModes:
    - ReadWriteOnce
    vctSize: 100Gi
  labels: {}
  minNodePort: 9000
  namespace: ""
  stopAll: false
  traefik:
    commonMiddlewares:
    - name: tc-basic-secure-headers
gluetunImage:
  pullPolicy: IfNotPresent
  repository: oci.trueforge.org/tccr/gluetun
  tag: v3.40.0@sha256:a8189e29155e0f8142be1500ae068a92b189b1b25abbba036321e74d6389bf2b
hpa:
  main:
    enabled: false
    targetSelector: []
image:
  pullPolicy: IfNotPresent
  repository: docker.io/clamav/clamav
  tag: 1.5.2@sha256:f7954ca7ca13f0ebc301bb8a8b492a88db793192c48763a0b9a2ffb64dcf4bee
imagePullSecret: {}
ingress:
  main:
    annotations: {}
    enabled: false
    hosts: []
    ingressClassName: ""
    integrations:
      certManager:
        certificateIssuer: ""
        enabled: false
      homepage:
        description: ""
        enabled: false
        group: ""
        icon: ""
        name: ""
        widget:
          type: ""
          url: ""
      nginx:
        auth:
          externalHost: ""
          internalHost: ""
          responseHeaders: []
          type: ""
        enabled: false
        ipWhitelist: []
        themepark:
          css: ""
          enabled: false
      traefik:
        enabled: false
        entrypoints:
        - websecure
        forceTLS: true
        middlewares: []
    labels: {}
    primary: true
    required: false
    tls: []
ingressMiddlewares:
  traefik:
    tc-basic-secure-headers:
      data:
        accessControlAllowMethods:
        - GET
        - OPTIONS
        - HEAD
        - PUT
        accessControlMaxAge: 100
        browserXssFilter: true
        contentTypeNosniff: true
        customRequestHeaders:
          X-Forwarded-Proto: https
        forceSTSHeader: true
        referrerPolicy: same-origin
        stsSeconds: 63072000
      enabled: false
      type: headers
kubectlImage:
  pullPolicy: IfNotPresent
  repository: oci.trueforge.org/containerforge/kubectl
  tag: 1.31.1@sha256:47fa8a164386b5320c18140236109baec862d7ed778215263c473965cecd18ac
mariadb:
  creds: {}
  enabled: false
  includeCommon: false
  password: PLACEHOLDERPASSWORD
  rootPassword: PLACEHOLDERROOTPASSWORD
mariadbClientImage:
  pullPolicy: IfNotPresent
  repository: oci.trueforge.org/containerforge/mariadb-client
  tag: 12.2.2@sha256:a731429e1a17c24609eea50b0705ff5ad2e74fde98545d465f41bf2175a9c2e4
metrics:
  main:
    enabled: false
    endpoints:
    - honorLabels: false
      interval: 5s
      path: /
      port: main
      scrapeTimeout: 5s
    primary: true
    prometheusRule:
      enabled: false
      groups: {}
    selector: {}
    type: servicemonitor
mongodb:
  creds: {}
  enabled: false
  includeCommon: false
  password: PLACEHOLDERPASSWORD
  rootPassword: PLACEHOLDERROOTPASSWORD
mongodbClientImage:
  pullPolicy: IfNotPresent
  repository: oci.trueforge.org/containerforge/mongosh
  tag: 2.8.3@sha256:f414e9cdfe773400f2a38577c2595f4005723d5e4156b71d5dd5e3c15772658a
namespace: ""
netshootImage:
  pullPolicy: IfNotPresent
  repository: oci.trueforge.org/tccr/netshoot
  tag: v0.14.0@sha256:28ede4317d22391e7d89a15eb78dc2afc3587ece02c76c983dde7239a0e43679
notes:
  custom: |
    {{- if .Values.chartContext.appUrl }}
    ## Connecting externally
    You can use this Chart by opening the following links in your browser:
    - {{ toYaml .Values.chartContext.appUrl }}
    {{- end }}

    {{ if .Chart.Dependencies }}
    ## Dependencies for {{ .Chart.Name }}

    {{- range .Chart.Dependencies }}
    - Chart: {{ .Repository }}/{{ .Name }}
      Version: {{ .Version }}
    {{- end }}
    {{- end }}
    {{- if .Values.chartContext.internalUrls }}
    ## Connecting Internally

    You can reach this chart inside your cluster, using the following service URLS:
    {{- range $url := .Values.chartContext.internalUrls -}}
    - {{ $url }}
    {{- end }}
    {{- end }}

    ## Sources for {{ .Chart.Name }}

    {{- range .Chart.Sources }}
    - {{ . }}
    {{- end -}}

    {{- $link := .Chart.Annotations.docs -}}
    {{- if not $link -}}
      {{- $link = .Chart.Home -}}
    {{- end }}

    See more for **{{ $.Chart.Name }}** at ({{ $link }})
  footer: |
    ## Documentation
    Please check out the TrueCharts documentation on:
    https://truecharts.org

    OpenSource can only exist with your help, please consider supporting TrueCharts:
    https://trueforge.org/sponsor
  header: |
    # Thank you for installing {{ .Chart.Name }} by TrueCharts.
  warnings: []
openvpnImage:
  pullPolicy: IfNotPresent
  repository: oci.trueforge.org/tccr/openvpn-client
  tag: latest@sha256:9bfdf50791d6e51056e31c03f73c9db329b2b72e7746155cfdc63e0c8b49b55a
persistence:
  devshm:
    enabled: true
    medium: Memory
    mountPath: /dev/shm
    targetSelectAll: true
    type: emptyDir
  logs:
    enabled: true
    mountPath: /logs
    targetSelectAll: true
  scandir:
    enabled: true
    mountPath: /scandir
    readOnly: true
    targetSelectAll: true
  shared:
    enabled: true
    mountPath: /shared
    targetSelectAll: true
    type: emptyDir
  sigdatabase:
    enabled: true
    mountPath: /var/lib/clamav
    targetSelectAll: true
  tmp:
    enabled: true
    medium: Memory
    mountPath: /tmp
    targetSelectAll: true
    type: emptyDir
  varlogs:
    enabled: true
    medium: Memory
    mountPath: /var/logs
    targetSelectAll: true
    type: emptyDir
  varrun:
    enabled: true
    medium: Memory
    mountPath: /var/run
    targetSelectAll: true
    type: emptyDir
podDisruptionBudget:
  main:
    enabled: false
    targetSelector: main
podOptions:
  affinity: {}
  automountServiceAccountToken: false
  defaultAffinity: true
  defaultSpread: true
  dnsConfig:
    options:
    - name: ndots
      value: "1"
  dnsPolicy: ClusterFirst
  enableServiceLinks: false
  hostAliases: []
  hostIPC: false
  hostNetwork: false
  hostPID: false
  nodeSelector:
    kubernetes.io/arch: amd64
  priorityClassName: ""
  runtimeClassName: ""
  schedulerName: ""
  shareProcessNamespace: false
  terminationGracePeriodSeconds: 60
  tolerations: []
  topologySpreadConstraints: []
postgres15Image:
  pullPolicy: IfNotPresent
  repository: ghcr.io/cloudnative-pg/postgresql
  tag: 15.17@sha256:23d56b4e6232954ba7ae9eda453a9630946f382325e3bb63af3c2c32fce9a119
postgres16Image:
  pullPolicy: IfNotPresent
  repository: ghcr.io/cloudnative-pg/postgresql
  tag: 16.13@sha256:da09a5927dd7af1572591df6c2c1a192d441054ff73e6df823402fcbf08c5ff0
postgresClientImage:
  pullPolicy: IfNotPresent
  repository: oci.trueforge.org/containerforge/postgresql-client
  tag: 9.6.24@sha256:c7ea20fbec42f93e8c45d05676673654dfba12665f63b9dc7020a00e01cff3cc
postgresPostgis15Image:
  pullPolicy: IfNotPresent
  repository: ghcr.io/cloudnative-pg/postgis
  tag: 15-3.4@sha256:4596b193991cd2463a07d40e3d3d62c59f046a7dece8f163ad1aba15925a38e4
postgresPostgis16Image:
  pullPolicy: IfNotPresent
  repository: ghcr.io/cloudnative-pg/postgis
  tag: 16-3.4@sha256:bb5a8590a8c934767482e34e1d103253f412aec703b77bbc52ad9044bf6e56f9
postgresVectorchord15Image:
  pullPolicy: IfNotPresent
  repository: ghcr.io/tensorchord/cloudnative-vectorchord
  tag: 15.14-0.5.3@sha256:1978732dc1e7e9ef94b9e806a094fcb123afab1e50a7e878e2d29de8b849cf47
postgresVectorchord16Image:
  pullPolicy: IfNotPresent
  repository: ghcr.io/tensorchord/cloudnative-vectorchord
  tag: 16.10-0.5.3@sha256:a0776b514bb23858d8aa59b08587223f8b449bba1b278ffbcf85e097d6504eb0
postgresVectors15Image:
  pullPolicy: IfNotPresent
  repository: ghcr.io/tensorchord/cloudnative-pgvecto.rs
  tag: 15.7-v0.2.1@sha256:dbdeddf0d635f76df41f745407816c87c7468df35e3b7b0665ca4e0500ff3048
postgresVectors16Image:
  pullPolicy: IfNotPresent
  repository: ghcr.io/tensorchord/cloudnative-pgvecto.rs
  tag: 16.3-v0.2.1@sha256:f1a19d4fc4073b0671a72ad34ef012aa20d21b3ddf5b4b0c9077d54450db679a
priorityClass: {}
rbac: {}
redis:
  creds: {}
  enabled: false
  includeCommon: false
  password: PLACEHOLDERPASSWORD
  secret:
    credentials:
      enabled: false
redisClientImage:
  pullPolicy: IfNotPresent
  repository: oci.trueforge.org/containerforge/valkey-tools
  tag: 1.1.0@sha256:10dfaf1f3bcc49a58eb1c298d3433efbe8e05e17b7e856699381efb2e4338362
resources:
  limits:
    cpu: 1500m
    memory: 2400Mi
  requests:
    cpu: 75m
    memory: 200Mi
route:
  main:
    annotations: {}
    enabled: false
    hostnames: []
    kind: HTTPRoute
    labels: {}
    parentRefs:
    - group: gateway.networking.k8s.io
      kind: Gateway
      name: null
      namespace: null
      sectionName: null
    rules:
    - backendRefs:
      - group: ""
        kind: Service
        name: null
        namespace: null
        port: null
        weight: 1
      matches:
      - path:
          type: PathPrefix
          value: /
scratchImage:
  pullPolicy: IfNotPresent
  repository: oci.trueforge.org/containerforge/scratch
  tag: 1.0.0@sha256:84182a2371dd7584ad752326b7150b4e928c640b7f70a249a3994b40f1db11fd
secret: {}
securityContext:
  container:
    PUID: 568
    UMASK: "0022"
    allowPrivilegeEscalation: false
    capabilities:
      add: []
      disableS6Caps: false
      drop:
      - ALL
    privileged: false
    readOnlyRootFilesystem: false
    runAsGroup: 0
    runAsNonRoot: false
    runAsUser: 0
    seccompProfile:
      type: RuntimeDefault
  pod:
    fsGroup: 568
    fsGroupChangePolicy: OnRootMismatch
    supplementalGroups: []
    sysctls: []
service:
  main:
    enabled: true
    ports:
      main:
        enabled: true
        port: 3310
        primary: true
        protocol: http
        targetPort: 3310
    primary: true
  milter:
    enabled: false
    ports:
      milter:
        enabled: true
        port: 7357
        protocol: http
        targetPort: 7357
serviceAccount: {}
solr:
  creds: {}
  enabled: false
  includeCommon: false
  password: PLACEHOLDERPASSWORD
  solrCores: 1
  solrEnableAuthentication: "no"
storageClass: {}
tailscaleImage:
  pullPolicy: IfNotPresent
  repository: oci.trueforge.org/tccr/tailscale
  tag: v1.88.3@sha256:878612592f133bc0728e978558b10a1c457371ac5949985d0584664c8e92c2f9
ubuntuImage:
  pullPolicy: IfNotPresent
  repository: oci.trueforge.org/containerforge/ubuntu
  tag: 26.4@sha256:17312da6b52fe7ab6e3bdc8e4a1f99df63c2494b2bad4f4b634a9065d585898e
valkeyClientImage:
  pullPolicy: IfNotPresent
  repository: oci.trueforge.org/containerforge/valkey-tools
  tag: 1.1.0@sha256:10dfaf1f3bcc49a58eb1c298d3433efbe8e05e17b7e856699381efb2e4338362
volumeSnapshotClass: {}
volumeSnapshots: {}
vpa:
  main:
    enabled: false
    resourcePolicy:
      containerPolicies:
      - containerName: '*'
        controlledResources:
        - cpu
        - memory
        maxAllowed:
          cpu: 8000m
          memory: 20Gi
        minAllowed:
          cpu: 50m
          memory: 50Mi
    targetSelector: []
webhook:
  mutating:
    enabled: false
    type: mutating
    webhooks: []
  validating:
    enabled: false
    type: validating
    webhooks: []
wgetImage:
  pullPolicy: IfNotPresent
  repository: oci.trueforge.org/containerforge/ubuntu
  tag: 26.4@sha256:17312da6b52fe7ab6e3bdc8e4a1f99df63c2494b2bad4f4b634a9065d585898e
wireguardImage:
  pullPolicy: IfNotPresent
  repository: oci.trueforge.org/tccr/wireguard
  tag: v1.0.20210914@sha256:683b8b74d64ebd07f9955147539834c2a4b60fee51d2a36fa76b9aba689601bf
workload:
  cron:
    Error: 'error converting YAML to JSON: yaml: line 3: did not find expected alphabetic
      or numeric character'
  main:
    dbWait: true
    enabled: true
    podSpec:
      containers:
        main:
          enabled: true
          env:
            CLAMAV_NO_CLAMD: false
            CLAMAV_NO_FRESHCLAMD: false
            CLAMAV_NO_MILTERD: '{{ not .Values.service.milter.enabled }}'
            CLAMD_STARTUP_TIMEOUT: 1800
            FRESHCLAM_CHECKS: 1
          imageSelector: image
          primary: true
          probes:
            liveness:
              command:
              - clamdcheck.sh
              enabled: true
              port: '{{ $.Values.service.main.ports.main.targetPort | default .Values.service.main.ports.main.port
                }}'
              type: exec
            readiness:
              command:
              - clamdcheck.sh
              enabled: true
              port: '{{ $.Values.service.main.ports.main.targetPort | default .Values.service.main.ports.main.port
                }}'
              type: exec
            startup:
              command:
              - clamdcheck.sh
              enabled: true
              port: '{{ $.Values.service.main.ports.main.targetPort | default .Values.service.main.ports.main.port
                }}'
              type: exec
    primary: true
    type: Deployment
yqImage:
  pullPolicy: IfNotPresent
  repository: oci.trueforge.org/containerforge/go-yq
  tag: 4.53.2@sha256:297b088fc4e93a81e782eb3b641345c71d43dff3107753c0468c1ef9498c63f3 
 ============================================================================================= 

 See error above values.

Use --debug flag to render out invalid YAML
```

