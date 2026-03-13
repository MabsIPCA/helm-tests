# Complete Errors By Taxonomy Bucket

Generated at: `2026-03-13 00:39:39 UTC`

Source catalog: `C:\Users\miabs\GolandProjects\helm-tests\helm_fetcher\backup\run_002_274\catalog_by_project.json`

## `template.nil_pointer` (328)

### 1. `ayarinesrine/pfe-deployment`

- Chart: `cloned\ayarinesrine__pfe-deployment\helm-charts\backend\authentification`
- Source: `template`
- Command: `helm template test cloned\ayarinesrine__pfe-deployment\helm-charts\backend\authentification`

```text
Error: template: authentification/templates/service.yaml:6:18: executing "authentification/templates/service.yaml" at <.Values.service.type>: nil pointer evaluating interface {}.type

Use --debug flag to render out invalid YAML
```

### 2. `ayarinesrine/pfe-deployment`

- Chart: `cloned\ayarinesrine__pfe-deployment\helm-charts\backend\devis-en-ligne`
- Source: `template`
- Command: `helm template test cloned\ayarinesrine__pfe-deployment\helm-charts\backend\devis-en-ligne`

```text
Error: template: devis-en-ligne/templates/service.yaml:6:18: executing "devis-en-ligne/templates/service.yaml" at <.Values.service.type>: nil pointer evaluating interface {}.type

Use --debug flag to render out invalid YAML
```

### 3. `ayarinesrine/pfe-deployment`

- Chart: `cloned\ayarinesrine__pfe-deployment\helm-charts\backend\offre-d-emplois`
- Source: `template`
- Command: `helm template test cloned\ayarinesrine__pfe-deployment\helm-charts\backend\offre-d-emplois`

```text
Error: template: offre-d-emplois/templates/service.yaml:6:18: executing "offre-d-emplois/templates/service.yaml" at <.Values.service.type>: nil pointer evaluating interface {}.type

Use --debug flag to render out invalid YAML
```

### 4. `ayarinesrine/pfe-deployment`

- Chart: `cloned\ayarinesrine__pfe-deployment\helm-charts\backend\upload`
- Source: `template`
- Command: `helm template test cloned\ayarinesrine__pfe-deployment\helm-charts\backend\upload`

```text
Error: template: upload/templates/service.yaml:6:18: executing "upload/templates/service.yaml" at <.Values.service.type>: nil pointer evaluating interface {}.type

Use --debug flag to render out invalid YAML
```

### 5. `ayarinesrine/pfe-deployment`

- Chart: `cloned\ayarinesrine__pfe-deployment\helm-charts\frontend\dashbord`
- Source: `template`
- Command: `helm template test cloned\ayarinesrine__pfe-deployment\helm-charts\frontend\dashbord`

```text
Error: template: dashbord/templates/service.yaml:6:18: executing "dashbord/templates/service.yaml" at <.Values.service.type>: nil pointer evaluating interface {}.type

Use --debug flag to render out invalid YAML
```

### 6. `ayarinesrine/pfe-deployment`

- Chart: `cloned\ayarinesrine__pfe-deployment\helm-charts\frontend\dynamique`
- Source: `template`
- Command: `helm template test cloned\ayarinesrine__pfe-deployment\helm-charts\frontend\dynamique`

```text
Error: template: dynamique/templates/service.yaml:6:18: executing "dynamique/templates/service.yaml" at <.Values.service.type>: nil pointer evaluating interface {}.type

Use --debug flag to render out invalid YAML
```

### 7. `ayarinesrine/pfe-deployment`

- Chart: `cloned\ayarinesrine__pfe-deployment\helm-charts\frontend\statique`
- Source: `template`
- Command: `helm template test cloned\ayarinesrine__pfe-deployment\helm-charts\frontend\statique`

```text
Error: template: statique/templates/service.yaml:6:18: executing "statique/templates/service.yaml" at <.Values.service.type>: nil pointer evaluating interface {}.type

Use --debug flag to render out invalid YAML
```

### 8. `SAP/component-operator-runtime`

- Chart: `cloned\SAP__component-operator-runtime\internal\helm\testdata\main`
- Source: `template`
- Command: `helm template test cloned\SAP__component-operator-runtime\internal\helm\testdata\main`

```text
Error: template: main/templates/configmap.yaml:20:24: executing "main/templates/configmap.yaml" at <.Values.global.data>: nil pointer evaluating interface {}.data

Use --debug flag to render out invalid YAML
```

### 9. `SAP/component-operator-runtime`

- Chart: `cloned\SAP__component-operator-runtime\internal\helm\testdata\main\charts\sub11`
- Source: `template`
- Command: `helm template test cloned\SAP__component-operator-runtime\internal\helm\testdata\main\charts\sub11`

```text
Error: template: sub11/templates/configmap.yaml:16:24: executing "sub11/templates/configmap.yaml" at <.Values.global.data>: nil pointer evaluating interface {}.data

Use --debug flag to render out invalid YAML
```

### 10. `SAP/component-operator-runtime`

- Chart: `cloned\SAP__component-operator-runtime\internal\helm\testdata\main\charts\sub11\charts\sub21`
- Source: `template`
- Command: `helm template test cloned\SAP__component-operator-runtime\internal\helm\testdata\main\charts\sub11\charts\sub21`

```text
Error: template: sub21/templates/configmap.yaml:11:24: executing "sub21/templates/configmap.yaml" at <.Values.global.data>: nil pointer evaluating interface {}.data

Use --debug flag to render out invalid YAML
```

### 11. `SAP/component-operator-runtime`

- Chart: `cloned\SAP__component-operator-runtime\internal\helm\testdata\main\charts\sub12`
- Source: `template`
- Command: `helm template test cloned\SAP__component-operator-runtime\internal\helm\testdata\main\charts\sub12`

```text
Error: template: sub12/templates/configmap.yaml:11:24: executing "sub12/templates/configmap.yaml" at <.Values.global.data>: nil pointer evaluating interface {}.data

Use --debug flag to render out invalid YAML
```

### 12. `ericmelz/joplin-emelz`

- Chart: `cloned\ericmelz__joplin-emelz\joplin-cron\helm`
- Source: `template`
- Command: `helm template test cloned\ericmelz__joplin-emelz\joplin-cron\helm`

```text
Error: template: joplin-cron/templates/configmap-accounts.yaml:8:65: executing "joplin-cron/templates/configmap-accounts.yaml" at <.Values.secrets.joplinPasswords.ericAtEmelzDotOrg>: nil pointer evaluating interface {}.joplinPasswords

Use --debug flag to render out invalid YAML
```

### 13. `noygal/helm`

- Chart: `cloned\noygal__helm\charts\base-chart`
- Source: `template`
- Command: `helm template test cloned\noygal__helm\charts\base-chart`

```text
Error: template: base-chart/templates/tests/test-connection.yaml:14:62: executing "base-chart/templates/tests/test-connection.yaml" at <.Values.service.port>: nil pointer evaluating interface {}.port

Use --debug flag to render out invalid YAML
```

### 14. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`

```text
Error: template: bookshop/templates/cap-operator-cros-svc-ias.yaml:95:23: executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 15. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-svc-ias.yaml:95:23: executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 16. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-svc-ias.yaml:95:23: executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 17. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-svc-ias.yaml:95:23: executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 18. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-svc-ias.yaml:95:23: executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 19. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-svc-ias.yaml:95:23: executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 20. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-svc-ias.yaml:95:23: executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 21. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-svc-ias.yaml:95:23: executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 22. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 23. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 24. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 25. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 26. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 27. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 28. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 29. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 30. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-svc-ias.yaml:95:23: executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 31. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-svc-ias.yaml:95:23: executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 32. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-svc-ias.yaml:95:23: executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 33. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-svc-ias.yaml:95:23: executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 34. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-svc-ias.yaml:95:23: executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 35. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-svc-ias.yaml:95:23: executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 36. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-svc-ias.yaml:95:23: executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 37. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-svc-ias.yaml:95:23: executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 38. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 39. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 40. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 41. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 42. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 43. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 44. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 45. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 46. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-svc-ias.yaml:95:23: executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 47. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-svc-ias.yaml:95:23: executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 48. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-svc-ias.yaml:95:23: executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 49. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-svc-ias.yaml:95:23: executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 50. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-svc-ias.yaml:95:23: executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 51. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-svc-ias.yaml:95:23: executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 52. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-svc-ias.yaml:95:23: executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 53. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-svc-ias.yaml:95:23: executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 54. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 55. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 56. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 57. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 58. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 59. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 60. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 61. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 62. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-svc-ias.yaml:95:23: executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 63. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-svc-ias.yaml:95:23: executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 64. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-svc-ias.yaml:95:23: executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 65. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-svc-ias.yaml:95:23: executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 66. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-svc-ias.yaml:95:23: executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 67. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-svc-ias.yaml:95:23: executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 68. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-svc-ias.yaml:95:23: executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 69. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-svc-ias.yaml:95:23: executing "bookshop/templates/cap-operator-cros-svc-ias.yaml" at <.Values.workloads.amsDeployer.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 70. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 71. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 72. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 73. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 74. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 75. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 76. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 77. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 78. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 79. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 80. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 81. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 82. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 83. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 84. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 85. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 86. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 87. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 88. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 89. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 90. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 91. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 92. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 93. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 94. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 95. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 96. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 97. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 98. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 99. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 100. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 101. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 102. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 103. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 104. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 105. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 106. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 107. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 108. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 109. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 110. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 111. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 112. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 113. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-ias.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values-svc.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 114. `cap-js/cap-operator-plugin`

- Chart: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart`
- Source: `template`
- Values files: `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml`, `cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`
- Command: `helm template test cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\runtime-values.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-modified.yaml -f cloned\cap-js__cap-operator-plugin\test\files\expectedConfigurableTemplatesChart\values-svc-ias.yaml`

```text
Error: template: bookshop/templates/cap-operator-cros-mta.yaml:67:23: executing "bookshop/templates/cap-operator-cros-mta.yaml" at <.Values.workloads.authorReadingsApprouter.image>: nil pointer evaluating interface {}.image

Use --debug flag to render out invalid YAML
```

### 115. `steled/argocd-apps`

- Chart: `cloned\steled__argocd-apps\00_archive\charts\immich`
- Source: `template`
- Command: `helm template test cloned\steled__argocd-apps\00_archive\charts\immich`

```text
Error: template: immich-addition/templates/secret.yaml:7:31: executing "immich-addition/templates/secret.yaml" at <.Values.secret.postgresqlUserPassword>: nil pointer evaluating interface {}.postgresqlUserPassword

Use --debug flag to render out invalid YAML
```

### 116. `steled/argocd-apps`

- Chart: `cloned\steled__argocd-apps\00_archive\charts\immich`
- Source: `template`
- Values files: `cloned\steled__argocd-apps\00_archive\charts\immich\postgresql_values.yaml`
- Command: `helm template test cloned\steled__argocd-apps\00_archive\charts\immich -f cloned\steled__argocd-apps\00_archive\charts\immich\postgresql_values.yaml`

```text
Error: template: immich-addition/templates/secret.yaml:7:31: executing "immich-addition/templates/secret.yaml" at <.Values.secret.postgresqlUserPassword>: nil pointer evaluating interface {}.postgresqlUserPassword

Use --debug flag to render out invalid YAML
```

### 117. `steled/argocd-apps`

- Chart: `cloned\steled__argocd-apps\00_archive\charts\jellyfin`
- Source: `template`
- Command: `helm template test cloned\steled__argocd-apps\00_archive\charts\jellyfin`

```text
Error: template: jellyfin-addition/templates/httproute.yaml:7:16: executing "jellyfin-addition/templates/httproute.yaml" at <.Values.httproute.hostnames0>: nil pointer evaluating interface {}.hostnames0

Use --debug flag to render out invalid YAML
```

### 118. `steled/argocd-apps`

- Chart: `cloned\steled__argocd-apps\apps`
- Source: `template`
- Command: `helm template test cloned\steled__argocd-apps\apps`

```text
Error: template: applications/templates/helm-victoriametrics.yaml:23:29: executing "applications/templates/helm-victoriametrics.yaml" at <.Values.victoriametrics.alertmanager.route.hostnames0>: nil pointer evaluating interface {}.route

Use --debug flag to render out invalid YAML
```

### 119. `steled/argocd-apps`

- Chart: `cloned\steled__argocd-apps\charts\cert-manager`
- Source: `template`
- Command: `helm template test cloned\steled__argocd-apps\charts\cert-manager`

```text
Error: template: cert-manager-addition/templates/clusterissuer-stg.yaml:7:21: executing "cert-manager-addition/templates/clusterissuer-stg.yaml" at <.Values.clusterIssuer.certmanager.email>: nil pointer evaluating interface {}.certmanager

Use --debug flag to render out invalid YAML
```

### 120. `steled/argocd-apps`

- Chart: `cloned\steled__argocd-apps\charts\fritzbox-cloudflare-dyndns`
- Source: `template`
- Command: `helm template test cloned\steled__argocd-apps\charts\fritzbox-cloudflare-dyndns`

```text
Error: template: fritzbox-cloudflare-dyndns-addition/templates/httproute.yaml:7:16: executing "fritzbox-cloudflare-dyndns-addition/templates/httproute.yaml" at <.Values.httproute.hostnames0>: nil pointer evaluating interface {}.hostnames0

Use --debug flag to render out invalid YAML
```

### 121. `steled/argocd-apps`

- Chart: `cloned\steled__argocd-apps\charts\hassio`
- Source: `template`
- Command: `helm template test cloned\steled__argocd-apps\charts\hassio`

```text
Error: template: hassio/templates/httproute.yaml:7:16: executing "hassio/templates/httproute.yaml" at <.Values.httproute.hostnames0>: nil pointer evaluating interface {}.hostnames0

Use --debug flag to render out invalid YAML
```

### 122. `steled/argocd-apps`

- Chart: `cloned\steled__argocd-apps\charts\homepage`
- Source: `template`
- Command: `helm template test cloned\steled__argocd-apps\charts\homepage`

```text
Error: template: homepage-addition/templates/httproute.yaml:7:16: executing "homepage-addition/templates/httproute.yaml" at <.Values.httproute.hostnames0>: nil pointer evaluating interface {}.hostnames0

Use --debug flag to render out invalid YAML
```

### 123. `steled/argocd-apps`

- Chart: `cloned\steled__argocd-apps\charts\jellyfin-sftp\jellyfin`
- Source: `template`
- Command: `helm template test cloned\steled__argocd-apps\charts\jellyfin-sftp\jellyfin`

```text
Error: template: jellyfin-addition/templates/httproute.yaml:7:16: executing "jellyfin-addition/templates/httproute.yaml" at <.Values.httproute.hostnames0>: nil pointer evaluating interface {}.hostnames0

Use --debug flag to render out invalid YAML
```

### 124. `steled/argocd-apps`

- Chart: `cloned\steled__argocd-apps\charts\kiali`
- Source: `template`
- Command: `helm template test cloned\steled__argocd-apps\charts\kiali`

```text
Error: template: kiali-addition/templates/httproute.yaml:8:16: executing "kiali-addition/templates/httproute.yaml" at <.Values.httproute.hostnames0>: nil pointer evaluating interface {}.hostnames0

Use --debug flag to render out invalid YAML
```

### 125. `steled/argocd-apps`

- Chart: `cloned\steled__argocd-apps\charts\longhorn`
- Source: `template`
- Command: `helm template test cloned\steled__argocd-apps\charts\longhorn`

```text
Error: template: longhorn-addition/templates/httproute.yaml:7:16: executing "longhorn-addition/templates/httproute.yaml" at <.Values.httproute.hostnames0>: nil pointer evaluating interface {}.hostnames0

Use --debug flag to render out invalid YAML
```

### 126. `steled/argocd-apps`

- Chart: `cloned\steled__argocd-apps\charts\minio`
- Source: `template`
- Command: `helm template test cloned\steled__argocd-apps\charts\minio`

```text
Error: template: minio/templates/httproute.yaml:7:16: executing "minio/templates/httproute.yaml" at <.Values.httproute.minio.hostnames0>: nil pointer evaluating interface {}.minio

Use --debug flag to render out invalid YAML
```

### 127. `steled/argocd-apps`

- Chart: `cloned\steled__argocd-apps\charts\nextcloud`
- Source: `template`
- Command: `helm template test cloned\steled__argocd-apps\charts\nextcloud`

```text
Error: template: nextcloud-addition/templates/httproute.yaml:7:16: executing "nextcloud-addition/templates/httproute.yaml" at <.Values.httproute.hostnames0>: nil pointer evaluating interface {}.hostnames0

Use --debug flag to render out invalid YAML
```

### 128. `steled/argocd-apps`

- Chart: `cloned\steled__argocd-apps\charts\sealed-secrets-web`
- Source: `template`
- Command: `helm template test cloned\steled__argocd-apps\charts\sealed-secrets-web`

```text
Error: template: sealed-secrets-web-addition/templates/httproute.yaml:8:16: executing "sealed-secrets-web-addition/templates/httproute.yaml" at <.Values.httproute.hostnames0>: nil pointer evaluating interface {}.hostnames0

Use --debug flag to render out invalid YAML
```

### 129. `steled/argocd-apps`

- Chart: `cloned\steled__argocd-apps\charts\victoriametrics`
- Source: `template`
- Command: `helm template test cloned\steled__argocd-apps\charts\victoriametrics`

```text
Error: template: victoriametrics-addition/templates/httproute.yaml:7:16: executing "victoriametrics-addition/templates/httproute.yaml" at <.Values.grafana.httproute.hostnames0>: nil pointer evaluating interface {}.httproute

Use --debug flag to render out invalid YAML
```

### 130. `steled/argocd-apps`

- Chart: `cloned\steled__argocd-apps\charts\victoriametrics`
- Source: `template`
- Values files: `cloned\steled__argocd-apps\charts\victoriametrics\victoriametrics-addition-values.yaml`
- Command: `helm template test cloned\steled__argocd-apps\charts\victoriametrics -f cloned\steled__argocd-apps\charts\victoriametrics\victoriametrics-addition-values.yaml`

```text
Error: template: victoriametrics-addition/templates/httproute.yaml:7:16: executing "victoriametrics-addition/templates/httproute.yaml" at <.Values.grafana.httproute.hostnames0>: nil pointer evaluating interface {}.httproute

Use --debug flag to render out invalid YAML
```

### 131. `steled/argocd-apps`

- Chart: `cloned\steled__argocd-apps\charts\victoriametrics`
- Source: `template`
- Values files: `cloned\steled__argocd-apps\charts\victoriametrics\victoriametrics-values.yaml`
- Command: `helm template test cloned\steled__argocd-apps\charts\victoriametrics -f cloned\steled__argocd-apps\charts\victoriametrics\victoriametrics-values.yaml`

```text
Error: template: victoriametrics-addition/templates/httproute.yaml:7:16: executing "victoriametrics-addition/templates/httproute.yaml" at <.Values.grafana.httproute.hostnames0>: nil pointer evaluating interface {}.hostnames0

Use --debug flag to render out invalid YAML
```

### 132. `steled/argocd-apps`

- Chart: `cloned\steled__argocd-apps\charts\victoriametrics`
- Source: `template`
- Values files: `cloned\steled__argocd-apps\charts\victoriametrics\victoriametrics-addition-values.yaml`, `cloned\steled__argocd-apps\charts\victoriametrics\victoriametrics-values.yaml`
- Command: `helm template test cloned\steled__argocd-apps\charts\victoriametrics -f cloned\steled__argocd-apps\charts\victoriametrics\victoriametrics-addition-values.yaml -f cloned\steled__argocd-apps\charts\victoriametrics\victoriametrics-values.yaml`

```text
Error: template: victoriametrics-addition/templates/httproute.yaml:7:16: executing "victoriametrics-addition/templates/httproute.yaml" at <.Values.grafana.httproute.hostnames0>: nil pointer evaluating interface {}.hostnames0

Use --debug flag to render out invalid YAML
```

### 133. `microsoft/mssql-docker`

- Chart: `cloned\microsoft__mssql-docker\linux\rancher`
- Source: `template`
- Command: `helm template test cloned\microsoft__mssql-docker\linux\rancher`

```text
Error: template: sql-server-rancher/templates/secret.yaml:9:32: executing "sql-server-rancher/templates/secret.yaml" at <.Values.mssql.sa.password>: nil pointer evaluating interface {}.password

Use --debug flag to render out invalid YAML
```

### 134. `hyperledger-bevel/bevel-operator-fabric`

- Chart: `cloned\hyperledger-bevel__bevel-operator-fabric\charts\hlf-ca`
- Source: `template`
- Command: `helm template test cloned\hyperledger-bevel__bevel-operator-fabric\charts\hlf-ca`

```text
Error: template: hlf-ca/templates/traefikroute.yaml:1:13: executing "hlf-ca/templates/traefikroute.yaml" at <.Values.traefik.hosts>: nil pointer evaluating interface {}.hosts

Use --debug flag to render out invalid YAML
```

### 135. `hyperledger-bevel/bevel-operator-fabric`

- Chart: `cloned\hyperledger-bevel__bevel-operator-fabric\charts\hlf-ordnode`
- Source: `template`
- Command: `helm template test cloned\hyperledger-bevel__bevel-operator-fabric\charts\hlf-ordnode`

```text
Error: template: hlf-ordnode/templates/traefikroute.yaml:1:13: executing "hlf-ordnode/templates/traefikroute.yaml" at <.Values.traefik.hosts>: nil pointer evaluating interface {}.hosts

Use --debug flag to render out invalid YAML
```

### 136. `fp-mt-test-org/devx-helloworld-springboot-hn5mq`

- Chart: `cloned\fp-mt-test-org__devx-helloworld-springboot-hn5mq\src\main\helm`
- Source: `template`
- Command: `helm template test cloned\fp-mt-test-org__devx-helloworld-springboot-hn5mq\src\main\helm`

```text
Error: template: devx-helloworld-springboot-hn5mq/templates/virtual-service.yaml:5:24: executing "devx-helloworld-springboot-hn5mq/templates/virtual-service.yaml" at <.Values.config.namespace>: nil pointer evaluating interface {}.namespace

Use --debug flag to render out invalid YAML
```

### 137. `randoli/helm-charts`

- Chart: `cloned\randoli__helm-charts\charts\cost-management`
- Source: `template`
- Command: `helm template test cloned\randoli__helm-charts\charts\cost-management`

```text
Error: template: cost-management/templates/serviceaccount.yaml:4:16: executing "cost-management/templates/serviceaccount.yaml" at <.Values.global.imagePullSecrets>: nil pointer evaluating interface {}.imagePullSecrets

Use --debug flag to render out invalid YAML
```

### 138. `randoli/helm-charts`

- Chart: `cloned\randoli__helm-charts\charts\logs`
- Source: `template`
- Command: `helm template test cloned\randoli__helm-charts\charts\logs`

```text
Error: template: logs/templates/vector.yaml:13:20: executing "logs/templates/vector.yaml" at <.Values.global.imagePullSecrets>: nil pointer evaluating interface {}.imagePullSecrets

Use --debug flag to render out invalid YAML
```

### 139. `Gemeente-DenHaag/helm-charts`

- Chart: `cloned\Gemeente-DenHaag__helm-charts\charts\dh-nl-portal-frontend`
- Source: `template`
- Command: `helm template test cloned\Gemeente-DenHaag__helm-charts\charts\dh-nl-portal-frontend`

```text
Error: template: dh-nl-portal-frontend/charts/dh-lib/templates/_deployment.yaml:101:14: executing "dh-lib.deployment" at <include "common.tplvalues.render" (dict "value" .Values.env "context" $)>: error calling include: template: dh-nl-portal-frontend/charts/dh-lib/charts/common/templates/_tplvalues.tpl:19:8: executing "common.tplvalues.render" at <tpl $value .context>: error calling tpl: error during tpl function execution for "- name: TZ\n  value: Europe/Amsterdam\n- name: OIDC_URL\n  value: '{{ .Values.config.OIDC_URL }}'\n- name: NGINX_ENVSUBST_OUTPUT_DIR\n  value: /tmp/nginx/conf.d": template: gotpl:4:20: executing "gotpl" at <.Values.config.OIDC_URL>: nil pointer evaluating interface {}.OIDC_URL

Use --debug flag to render out invalid YAML
```

### 140. `Gemeente-DenHaag/helm-charts`

- Chart: `cloned\Gemeente-DenHaag__helm-charts\charts\zgw-dre`
- Source: `template`
- Command: `helm template test cloned\Gemeente-DenHaag__helm-charts\charts\zgw-dre`

```text
Error: template: zgw-dre/templates/zgwdreconfigmap.yaml:7:15: executing "zgw-dre/templates/zgwdreconfigmap.yaml" at <.Values.configurationFiles.production>: nil pointer evaluating interface {}.production

Use --debug flag to render out invalid YAML
```

### 141. `myysophia/k8s-yaml-repo`

- Chart: `cloned\myysophia__k8s-yaml-repo\iotdb\prod\iotdb-confignode`
- Source: `template`
- Command: `helm template test cloned\myysophia__k8s-yaml-repo\iotdb\prod\iotdb-confignode`

```text
Error: template: iotdb-confignode/templates/statefulset.yaml:20:24: executing "iotdb-confignode/templates/statefulset.yaml" at <.Values.podAntiAffinity.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 142. `myysophia/k8s-yaml-repo`

- Chart: `cloned\myysophia__k8s-yaml-repo\iotdb\prod\iotdb-datanode`
- Source: `template`
- Command: `helm template test cloned\myysophia__k8s-yaml-repo\iotdb\prod\iotdb-datanode`

```text
Error: template: iotdb-datanode/templates/statefulset.yaml:20:24: executing "iotdb-datanode/templates/statefulset.yaml" at <.Values.podAntiAffinity.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 143. `ministryofjustice/laa-amend-a-claim`

- Chart: `cloned\ministryofjustice__laa-amend-a-claim\.helm\laa-amend-a-claim`
- Source: `template`
- Command: `helm template test cloned\ministryofjustice__laa-amend-a-claim\.helm\laa-amend-a-claim`

```text
Error: template: laa-amend-a-claim/templates/grafana/grafana-dashboard.yaml:1:14: executing "laa-amend-a-claim/templates/grafana/grafana-dashboard.yaml" at <.Values.grafana.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 144. `axlgrosse/42c-Protect`

- Chart: `cloned\axlgrosse__42c-Protect\deployment\Pixi-helm\42c-pixiapi`
- Source: `template`
- Command: `helm template test cloned\axlgrosse__42c-Protect\deployment\Pixi-helm\42c-pixiapi`

```text
Error: template: 42c-pixi/templates/pixidb-service.yaml:12:20: executing "42c-pixi/templates/pixidb-service.yaml" at <.Values.pixidb.listen_port>: nil pointer evaluating interface {}.listen_port

Use --debug flag to render out invalid YAML
```

### 145. `axlgrosse/42c-Protect`

- Chart: `cloned\axlgrosse__42c-Protect\deployment\helm-artifacts\42c-proxy`
- Source: `template`
- Command: `helm template test cloned\axlgrosse__42c-Protect\deployment\helm-artifacts\42c-proxy`

```text
Error: template: 42c-proxy/templates/tests/test-connection.yaml:13:61: executing "42c-proxy/templates/tests/test-connection.yaml" at <.Values.apifirewall.svc_listen_port>: nil pointer evaluating interface {}.svc_listen_port

Use --debug flag to render out invalid YAML
```

### 146. `axlgrosse/42c-Protect`

- Chart: `cloned\axlgrosse__42c-Protect\deployment\helm-artifacts\capture-plugin`
- Source: `template`
- Command: `helm template test cloned\axlgrosse__42c-Protect\deployment\helm-artifacts\capture-plugin`

```text
Error: template: horus-agent/templates/serviceaccount.yaml:4:11: executing "horus-agent/templates/serviceaccount.yaml" at <include "horus.serviceAccountName" .>: error calling include: template: horus-agent/templates/_helpers.tpl:57:14: executing "horus.serviceAccountName" at <.Values.serviceAccount.create>: nil pointer evaluating interface {}.create

Use --debug flag to render out invalid YAML
```

### 147. `gridpp-Edi/xrootd-helm-charts`

- Chart: `cloned\gridpp-Edi__xrootd-helm-charts\charts\xrootd`
- Source: `template`
- Command: `helm template test cloned\gridpp-Edi__xrootd-helm-charts\charts\xrootd`

```text
Error: template: xrootd/templates/tests/test-connection.yaml:14:58: executing "xrootd/templates/tests/test-connection.yaml" at <.Values.service.port>: nil pointer evaluating interface {}.port

Use --debug flag to render out invalid YAML
```

### 148. `t9k/apps`

- Chart: `cloned\t9k__apps\user-console\argo-workflows\app\chart`
- Source: `template`
- Command: `helm template test cloned\t9k__apps\user-console\argo-workflows\app\chart`

```text
Error: template: argo-workflows/templates/argo-server.deploy.yaml:24:25: executing "argo-workflows/templates/argo-server.deploy.yaml" at <.Values.image.registry>: nil pointer evaluating interface {}.registry

Use --debug flag to render out invalid YAML
```

### 149. `0xPolygonID/issuer-node`

- Chart: `cloned\0xPolygonID__issuer-node\k8s\helm\charts\ingress`
- Source: `template`
- Command: `helm template test cloned\0xPolygonID__issuer-node\k8s\helm\charts\ingress`

```text
Error: template: ingress/templates/traefik.yaml:11:20: executing "ingress/templates/traefik.yaml" at <.Values.global.uidomain>: nil pointer evaluating interface {}.uidomain

Use --debug flag to render out invalid YAML
```

### 150. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.118\charts\app`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.118\charts\app`

```text
Error: template: app/templates/app-pvc-reports.yaml:13:14: executing "app/templates/app-pvc-reports.yaml" at <.Values.persistence.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 151. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.118\charts\db`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.118\charts\db`

```text
Error: template: db/templates/db-pvc-log.yaml:13:14: executing "db/templates/db-pvc-log.yaml" at <.Values.persistence.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 152. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.118\charts\search`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.118\charts\search`

```text
Error: template: search/templates/search-pvc-data.yaml:13:14: executing "search/templates/search-pvc-data.yaml" at <.Values.persistence.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 153. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.118\charts\webapi`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.118\charts\webapi`

```text
Error: template: webapi/templates/webapi-pvc-log.yaml:13:14: executing "webapi/templates/webapi-pvc-log.yaml" at <.Values.persistence.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 154. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.121\charts\app`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.121\charts\app`

```text
Error: template: app/templates/app-statefulset.yaml:41:24: executing "app/templates/app-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 155. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.121\charts\db`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.121\charts\db`

```text
Error: template: db/templates/db-statefulset.yaml:41:24: executing "db/templates/db-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 156. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.121\charts\search`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.121\charts\search`

```text
Error: template: search/templates/search-statefulset.yaml:41:24: executing "search/templates/search-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 157. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.121\charts\webapi`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.121\charts\webapi`

```text
Error: template: webapi/templates/webapi-pvc-log.yaml:13:14: executing "webapi/templates/webapi-pvc-log.yaml" at <.Values.persistence.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 158. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.123\charts\app`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.123\charts\app`

```text
Error: template: app/templates/app-statefulset.yaml:41:24: executing "app/templates/app-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 159. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.123\charts\db`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.123\charts\db`

```text
Error: template: db/templates/db-statefulset.yaml:41:24: executing "db/templates/db-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 160. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.123\charts\search`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.123\charts\search`

```text
Error: template: search/templates/search-statefulset.yaml:41:24: executing "search/templates/search-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 161. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.123\charts\webapi`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.123\charts\webapi`

```text
Error: template: webapi/templates/webapi-pvc-log.yaml:13:14: executing "webapi/templates/webapi-pvc-log.yaml" at <.Values.persistence.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 162. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.125\charts\app`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.125\charts\app`

```text
Error: template: app/templates/app-statefulset.yaml:41:24: executing "app/templates/app-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 163. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.125\charts\db`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.125\charts\db`

```text
Error: template: db/templates/db-statefulset.yaml:41:24: executing "db/templates/db-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 164. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.125\charts\search`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.125\charts\search`

```text
Error: template: search/templates/search-statefulset.yaml:41:24: executing "search/templates/search-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 165. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.125\charts\webapi`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.125\charts\webapi`

```text
Error: template: webapi/templates/webapi-pvc-log.yaml:13:14: executing "webapi/templates/webapi-pvc-log.yaml" at <.Values.persistence.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 166. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.128\charts\app`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.128\charts\app`

```text
Error: template: app/templates/app-statefulset.yaml:41:24: executing "app/templates/app-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 167. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.128\charts\db`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.128\charts\db`

```text
Error: template: db/templates/db-statefulset.yaml:41:24: executing "db/templates/db-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 168. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.128\charts\search`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.128\charts\search`

```text
Error: template: search/templates/search-statefulset.yaml:41:24: executing "search/templates/search-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 169. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.128\charts\webapi`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.128\charts\webapi`

```text
Error: template: webapi/templates/webapi-pvc-log.yaml:13:14: executing "webapi/templates/webapi-pvc-log.yaml" at <.Values.persistence.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 170. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.138\charts\app`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.138\charts\app`

```text
Error: template: app/templates/app-statefulset.yaml:41:20: executing "app/templates/app-statefulset.yaml" at <.Values.reports.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 171. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.138\charts\db`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.138\charts\db`

```text
Error: template: db/templates/db-statefulset.yaml:41:24: executing "db/templates/db-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 172. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.138\charts\search`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.138\charts\search`

```text
Error: template: search/templates/search-statefulset.yaml:41:24: executing "search/templates/search-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 173. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.138\charts\webapi`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.138\charts\webapi`

```text
Error: template: webapi/templates/webapi-pvc-log.yaml:13:14: executing "webapi/templates/webapi-pvc-log.yaml" at <.Values.persistence.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 174. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144\charts\app`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144\charts\app`

```text
Error: template: app/templates/app-statefulset.yaml:41:20: executing "app/templates/app-statefulset.yaml" at <.Values.reports.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 175. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144\charts\db`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144\charts\db`

```text
Error: template: db/templates/db-statefulset.yaml:41:24: executing "db/templates/db-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 176. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144\charts\search`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144\charts\search`

```text
Error: template: search/templates/search-statefulset.yaml:41:24: executing "search/templates/search-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 177. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144\charts\webapi`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144\charts\webapi`

```text
Error: template: webapi/templates/webapi-pvc-log.yaml:13:14: executing "webapi/templates/webapi-pvc-log.yaml" at <.Values.persistence.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 178. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.0-kommch\charts\app`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.0-kommch\charts\app`

```text
Error: template: app/templates/app-statefulset.yaml:41:20: executing "app/templates/app-statefulset.yaml" at <.Values.reports.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 179. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.0-kommch\charts\db`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.0-kommch\charts\db`

```text
Error: template: db/templates/db-statefulset.yaml:41:24: executing "db/templates/db-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 180. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.0-kommch\charts\search`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.0-kommch\charts\search`

```text
Error: template: search/templates/search-statefulset.yaml:41:24: executing "search/templates/search-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 181. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.0-kommch\charts\webapi`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.0-kommch\charts\webapi`

```text
Error: template: webapi/templates/webapi-pvc-log.yaml:13:14: executing "webapi/templates/webapi-pvc-log.yaml" at <.Values.persistence.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 182. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.2-kommch\charts\app`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.2-kommch\charts\app`

```text
Error: template: app/templates/app-statefulset.yaml:41:20: executing "app/templates/app-statefulset.yaml" at <.Values.reports.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 183. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.2-kommch\charts\db`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.2-kommch\charts\db`

```text
Error: template: db/templates/db-statefulset.yaml:41:24: executing "db/templates/db-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 184. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.2-kommch\charts\search`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.2-kommch\charts\search`

```text
Error: template: search/templates/search-statefulset.yaml:41:24: executing "search/templates/search-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 185. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.2-kommch\charts\webapi`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.2-kommch\charts\webapi`

```text
Error: template: webapi/templates/webapi-pvc-log.yaml:13:14: executing "webapi/templates/webapi-pvc-log.yaml" at <.Values.persistence.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 186. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.4-kommch\charts\app`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.4-kommch\charts\app`

```text
Error: template: app/templates/app-statefulset.yaml:41:20: executing "app/templates/app-statefulset.yaml" at <.Values.reports.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 187. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.4-kommch\charts\db`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.4-kommch\charts\db`

```text
Error: template: db/templates/db-statefulset.yaml:41:24: executing "db/templates/db-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 188. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.4-kommch\charts\search`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.4-kommch\charts\search`

```text
Error: template: search/templates/search-statefulset.yaml:41:24: executing "search/templates/search-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 189. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.4-kommch\charts\webapi`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.4-kommch\charts\webapi`

```text
Error: template: webapi/templates/webapi-pvc-log.yaml:13:14: executing "webapi/templates/webapi-pvc-log.yaml" at <.Values.persistence.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 190. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.6-1-kommch\charts\app`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.6-1-kommch\charts\app`

```text
Error: template: app/templates/app-statefulset.yaml:41:20: executing "app/templates/app-statefulset.yaml" at <.Values.reports.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 191. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.6-1-kommch\charts\db`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.6-1-kommch\charts\db`

```text
Error: template: db/templates/db-statefulset.yaml:41:24: executing "db/templates/db-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 192. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.6-1-kommch\charts\search`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.6-1-kommch\charts\search`

```text
Error: template: search/templates/search-statefulset.yaml:41:24: executing "search/templates/search-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 193. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.6-1-kommch\charts\webapi`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.6-1-kommch\charts\webapi`

```text
Error: template: webapi/templates/webapi-pvc-log.yaml:13:14: executing "webapi/templates/webapi-pvc-log.yaml" at <.Values.persistence.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 194. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.6-1.kommch\charts\app`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.6-1.kommch\charts\app`

```text
Error: template: app/templates/app-statefulset.yaml:41:20: executing "app/templates/app-statefulset.yaml" at <.Values.reports.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 195. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.6-1.kommch\charts\db`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.6-1.kommch\charts\db`

```text
Error: template: db/templates/db-statefulset.yaml:41:24: executing "db/templates/db-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 196. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.6-1.kommch\charts\search`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.6-1.kommch\charts\search`

```text
Error: template: search/templates/search-statefulset.yaml:41:24: executing "search/templates/search-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 197. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.6-1.kommch\charts\webapi`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.6-1.kommch\charts\webapi`

```text
Error: template: webapi/templates/webapi-pvc-log.yaml:13:14: executing "webapi/templates/webapi-pvc-log.yaml" at <.Values.persistence.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 198. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.7-1-kommch\charts\app`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.7-1-kommch\charts\app`

```text
Error: template: app/templates/app-statefulset.yaml:41:20: executing "app/templates/app-statefulset.yaml" at <.Values.reports.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 199. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.7-1-kommch\charts\db`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.7-1-kommch\charts\db`

```text
Error: template: db/templates/db-statefulset.yaml:41:24: executing "db/templates/db-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 200. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.7-1-kommch\charts\search`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.7-1-kommch\charts\search`

```text
Error: template: search/templates/search-statefulset.yaml:41:24: executing "search/templates/search-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 201. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.7-1-kommch\charts\webapi`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.7-1-kommch\charts\webapi`

```text
Error: template: webapi/templates/webapi-pvc-log.yaml:13:14: executing "webapi/templates/webapi-pvc-log.yaml" at <.Values.persistence.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 202. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.7-2-kommch\charts\app`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.7-2-kommch\charts\app`

```text
Error: template: app/templates/app-statefulset.yaml:41:20: executing "app/templates/app-statefulset.yaml" at <.Values.reports.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 203. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.7-2-kommch\charts\db`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.7-2-kommch\charts\db`

```text
Error: template: db/templates/db-statefulset.yaml:41:24: executing "db/templates/db-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 204. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.7-2-kommch\charts\search`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.7-2-kommch\charts\search`

```text
Error: template: search/templates/search-statefulset.yaml:41:24: executing "search/templates/search-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 205. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.7-2-kommch\charts\webapi`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.7-2-kommch\charts\webapi`

```text
Error: template: webapi/templates/webapi-pvc-log.yaml:13:14: executing "webapi/templates/webapi-pvc-log.yaml" at <.Values.persistence.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 206. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-1-kommch\charts\app`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-1-kommch\charts\app`

```text
Error: template: app/templates/app-statefulset.yaml:41:20: executing "app/templates/app-statefulset.yaml" at <.Values.reports.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 207. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-1-kommch\charts\db`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-1-kommch\charts\db`

```text
Error: template: db/templates/db-statefulset.yaml:41:24: executing "db/templates/db-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 208. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-1-kommch\charts\search`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-1-kommch\charts\search`

```text
Error: template: search/templates/search-statefulset.yaml:41:24: executing "search/templates/search-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 209. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-1-kommch\charts\webapi`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-1-kommch\charts\webapi`

```text
Error: template: webapi/templates/webapi-pvc-log.yaml:13:14: executing "webapi/templates/webapi-pvc-log.yaml" at <.Values.persistence.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 210. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-10-kommch`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-10-kommch`

```text
Error: template: kommunikativch-metasfresh/charts/app/templates/app-statefulset.yaml:57:28: executing "kommunikativch-metasfresh/charts/app/templates/app-statefulset.yaml" at <.Values.reports.custom.gitCredentials>: nil pointer evaluating interface {}.gitCredentials

Use --debug flag to render out invalid YAML
```

### 211. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-10-kommch\charts\app`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-10-kommch\charts\app`

```text
Error: template: app/templates/app-statefulset.yaml:41:20: executing "app/templates/app-statefulset.yaml" at <.Values.reports.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 212. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-10-kommch\charts\db`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-10-kommch\charts\db`

```text
Error: template: db/templates/db-statefulset.yaml:41:24: executing "db/templates/db-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 213. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-10-kommch\charts\search`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-10-kommch\charts\search`

```text
Error: template: search/templates/search-statefulset.yaml:41:24: executing "search/templates/search-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 214. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-10-kommch\charts\webapi`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-10-kommch\charts\webapi`

```text
Error: template: webapi/templates/webapi-pvc-log.yaml:13:14: executing "webapi/templates/webapi-pvc-log.yaml" at <.Values.persistence.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 215. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-11-kommch`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-11-kommch`

```text
Error: template: kommunikativch-metasfresh/charts/app/templates/app-statefulset.yaml:57:28: executing "kommunikativch-metasfresh/charts/app/templates/app-statefulset.yaml" at <.Values.reports.custom.gitCredentials>: nil pointer evaluating interface {}.gitCredentials

Use --debug flag to render out invalid YAML
```

### 216. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-11-kommch\charts\app`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-11-kommch\charts\app`

```text
Error: template: app/templates/app-statefulset.yaml:41:20: executing "app/templates/app-statefulset.yaml" at <.Values.reports.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 217. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-11-kommch\charts\db`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-11-kommch\charts\db`

```text
Error: template: db/templates/db-statefulset.yaml:41:24: executing "db/templates/db-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 218. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-11-kommch\charts\search`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-11-kommch\charts\search`

```text
Error: template: search/templates/search-statefulset.yaml:41:24: executing "search/templates/search-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 219. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-11-kommch\charts\webapi`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-11-kommch\charts\webapi`

```text
Error: template: webapi/templates/webapi-pvc-log.yaml:13:14: executing "webapi/templates/webapi-pvc-log.yaml" at <.Values.persistence.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 220. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-2-kommch\charts\app`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-2-kommch\charts\app`

```text
Error: template: app/templates/app-statefulset.yaml:41:20: executing "app/templates/app-statefulset.yaml" at <.Values.reports.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 221. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-2-kommch\charts\db`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-2-kommch\charts\db`

```text
Error: template: db/templates/db-statefulset.yaml:41:24: executing "db/templates/db-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 222. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-2-kommch\charts\search`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-2-kommch\charts\search`

```text
Error: template: search/templates/search-statefulset.yaml:41:24: executing "search/templates/search-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 223. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-2-kommch\charts\webapi`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-2-kommch\charts\webapi`

```text
Error: template: webapi/templates/webapi-pvc-log.yaml:13:14: executing "webapi/templates/webapi-pvc-log.yaml" at <.Values.persistence.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 224. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-3-kommch\charts\app`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-3-kommch\charts\app`

```text
Error: template: app/templates/app-statefulset.yaml:41:20: executing "app/templates/app-statefulset.yaml" at <.Values.reports.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 225. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-3-kommch\charts\db`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-3-kommch\charts\db`

```text
Error: template: db/templates/db-statefulset.yaml:41:24: executing "db/templates/db-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 226. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-3-kommch\charts\search`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-3-kommch\charts\search`

```text
Error: template: search/templates/search-statefulset.yaml:41:24: executing "search/templates/search-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 227. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-3-kommch\charts\webapi`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-3-kommch\charts\webapi`

```text
Error: template: webapi/templates/webapi-pvc-log.yaml:13:14: executing "webapi/templates/webapi-pvc-log.yaml" at <.Values.persistence.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 228. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-6-kommch\charts\app`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-6-kommch\charts\app`

```text
Error: template: app/templates/app-statefulset.yaml:41:20: executing "app/templates/app-statefulset.yaml" at <.Values.reports.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 229. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-6-kommch\charts\db`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-6-kommch\charts\db`

```text
Error: template: db/templates/db-statefulset.yaml:41:24: executing "db/templates/db-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 230. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-6-kommch\charts\search`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-6-kommch\charts\search`

```text
Error: template: search/templates/search-statefulset.yaml:41:24: executing "search/templates/search-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 231. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-6-kommch\charts\webapi`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-6-kommch\charts\webapi`

```text
Error: template: webapi/templates/webapi-pvc-log.yaml:13:14: executing "webapi/templates/webapi-pvc-log.yaml" at <.Values.persistence.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 232. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-8-kommch\charts\app`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-8-kommch\charts\app`

```text
Error: template: app/templates/app-statefulset.yaml:41:20: executing "app/templates/app-statefulset.yaml" at <.Values.reports.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 233. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-8-kommch\charts\db`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-8-kommch\charts\db`

```text
Error: template: db/templates/db-statefulset.yaml:41:24: executing "db/templates/db-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 234. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-8-kommch\charts\search`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-8-kommch\charts\search`

```text
Error: template: search/templates/search-statefulset.yaml:41:24: executing "search/templates/search-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 235. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-8-kommch\charts\webapi`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-8-kommch\charts\webapi`

```text
Error: template: webapi/templates/webapi-pvc-log.yaml:13:14: executing "webapi/templates/webapi-pvc-log.yaml" at <.Values.persistence.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 236. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-9-kommch\charts\app`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-9-kommch\charts\app`

```text
Error: template: app/templates/app-statefulset.yaml:41:20: executing "app/templates/app-statefulset.yaml" at <.Values.reports.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 237. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-9-kommch\charts\db`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-9-kommch\charts\db`

```text
Error: template: db/templates/db-statefulset.yaml:41:24: executing "db/templates/db-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 238. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-9-kommch\charts\search`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-9-kommch\charts\search`

```text
Error: template: search/templates/search-statefulset.yaml:41:24: executing "search/templates/search-statefulset.yaml" at <.Values.transfer.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 239. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-9-kommch\charts\webapi`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-9-kommch\charts\webapi`

```text
Error: template: webapi/templates/webapi-pvc-log.yaml:13:14: executing "webapi/templates/webapi-pvc-log.yaml" at <.Values.persistence.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 240. `openidl-org/openidl-aais-gitops`

- Chart: `cloned\openidl-org__openidl-aais-gitops\openidl-k8s`
- Source: `template`
- Command: `helm template test cloned\openidl-org__openidl-aais-gitops\openidl-k8s`

```text
Error: template: openidl-k8s/charts/openidl-utilities/templates/service.yml:1:14: executing "openidl-k8s/charts/openidl-utilities/templates/service.yml" at <.Values.global.utilities.install>: nil pointer evaluating interface {}.install

Use --debug flag to render out invalid YAML
```

### 241. `openidl-org/openidl-aais-gitops`

- Chart: `cloned\openidl-org__openidl-aais-gitops\openidl-k8s\charts\openidl-carrier-ui`
- Source: `template`
- Command: `helm template test cloned\openidl-org__openidl-aais-gitops\openidl-k8s\charts\openidl-carrier-ui`

```text
Error: template: openidl-carrier-ui/templates/service.yml:1:14: executing "openidl-carrier-ui/templates/service.yml" at <.Values.global.carrierui.install>: nil pointer evaluating interface {}.carrierui

Use --debug flag to render out invalid YAML
```

### 242. `openidl-org/openidl-aais-gitops`

- Chart: `cloned\openidl-org__openidl-aais-gitops\openidl-k8s\charts\openidl-data-call-app`
- Source: `template`
- Command: `helm template test cloned\openidl-org__openidl-aais-gitops\openidl-k8s\charts\openidl-data-call-app`

```text
Error: template: openidl-data-call-app/templates/service.yml:1:14: executing "openidl-data-call-app/templates/service.yml" at <.Values.global.datacallapp.install>: nil pointer evaluating interface {}.datacallapp

Use --debug flag to render out invalid YAML
```

### 243. `openidl-org/openidl-aais-gitops`

- Chart: `cloned\openidl-org__openidl-aais-gitops\openidl-k8s\charts\openidl-data-call-mood-listener`
- Source: `template`
- Command: `helm template test cloned\openidl-org__openidl-aais-gitops\openidl-k8s\charts\openidl-data-call-mood-listener`

```text
Error: template: openidl-data-call-mood-listener/templates/service.yaml:1:14: executing "openidl-data-call-mood-listener/templates/service.yaml" at <.Values.global.datacallmoodlistener.install>: nil pointer evaluating interface {}.datacallmoodlistener

Use --debug flag to render out invalid YAML
```

### 244. `openidl-org/openidl-aais-gitops`

- Chart: `cloned\openidl-org__openidl-aais-gitops\openidl-k8s\charts\openidl-data-call-processor`
- Source: `template`
- Command: `helm template test cloned\openidl-org__openidl-aais-gitops\openidl-k8s\charts\openidl-data-call-processor`

```text
Error: template: openidl-data-call-processor/templates/service.yaml:1:14: executing "openidl-data-call-processor/templates/service.yaml" at <.Values.global.datacallprocessor.install>: nil pointer evaluating interface {}.datacallprocessor

Use --debug flag to render out invalid YAML
```

### 245. `openidl-org/openidl-aais-gitops`

- Chart: `cloned\openidl-org__openidl-aais-gitops\openidl-k8s\charts\openidl-insurance-data-manager`
- Source: `template`
- Command: `helm template test cloned\openidl-org__openidl-aais-gitops\openidl-k8s\charts\openidl-insurance-data-manager`

```text
Error: template: openidl-insurance-data-manager/templates/service.yml:1:14: executing "openidl-insurance-data-manager/templates/service.yml" at <.Values.global.insurancedatamanager.install>: nil pointer evaluating interface {}.insurancedatamanager

Use --debug flag to render out invalid YAML
```

### 246. `openidl-org/openidl-aais-gitops`

- Chart: `cloned\openidl-org__openidl-aais-gitops\openidl-k8s\charts\openidl-secrets`
- Source: `template`
- Command: `helm template test cloned\openidl-org__openidl-aais-gitops\openidl-k8s\charts\openidl-secrets`

```text
Error: template: openidl-secrets/templates/secrets.yaml:1:14: executing "openidl-secrets/templates/secrets.yaml" at <.Values.global.secrets.install>: nil pointer evaluating interface {}.secrets

Use --debug flag to render out invalid YAML
```

### 247. `openidl-org/openidl-aais-gitops`

- Chart: `cloned\openidl-org__openidl-aais-gitops\openidl-k8s\charts\openidl-transactional-data-event-listener`
- Source: `template`
- Command: `helm template test cloned\openidl-org__openidl-aais-gitops\openidl-k8s\charts\openidl-transactional-data-event-listener`

```text
Error: template: openidl-transactional-data-event-listener/templates/service.yaml:1:14: executing "openidl-transactional-data-event-listener/templates/service.yaml" at <.Values.global.transactionaldataeventlistener.install>: nil pointer evaluating interface {}.transactionaldataeventlistener

Use --debug flag to render out invalid YAML
```

### 248. `openidl-org/openidl-aais-gitops`

- Chart: `cloned\openidl-org__openidl-aais-gitops\openidl-k8s\charts\openidl-ui`
- Source: `template`
- Command: `helm template test cloned\openidl-org__openidl-aais-gitops\openidl-k8s\charts\openidl-ui`

```text
Error: template: openidl-ui/templates/service.yml:1:14: executing "openidl-ui/templates/service.yml" at <.Values.global.ui.install>: nil pointer evaluating interface {}.ui

Use --debug flag to render out invalid YAML
```

### 249. `openidl-org/openidl-aais-gitops`

- Chart: `cloned\openidl-org__openidl-aais-gitops\openidl-k8s\charts\openidl-upload`
- Source: `template`
- Command: `helm template test cloned\openidl-org__openidl-aais-gitops\openidl-k8s\charts\openidl-upload`

```text
Error: template: openidl-upload/templates/service.yml:1:14: executing "openidl-upload/templates/service.yml" at <.Values.global.upload.install>: nil pointer evaluating interface {}.upload

Use --debug flag to render out invalid YAML
```

### 250. `kubesphere/ksbuilder`

- Chart: `cloned\kubesphere__ksbuilder\pkg\extension\templatessimple\charts\base`
- Source: `template`
- Command: `helm template test cloned\kubesphere__ksbuilder\pkg\extension\templatessimple\charts\base`

```text
Error: template: base/templates/ingress.yaml:10:30: executing "base/templates/ingress.yaml" at <.Values.global.extension.ingress.ingressClassName>: nil pointer evaluating interface {}.extension

Use --debug flag to render out invalid YAML
```

### 251. `cdv18/helm-charts`

- Chart: `cloned\cdv18__helm-charts\charts\calculate-money`
- Source: `template`
- Command: `helm template test cloned\cdv18__helm-charts\charts\calculate-money`

```text
Error: template: calculate-money/templates/hpa.yaml:35:43: executing "calculate-money/templates/hpa.yaml" at <.Values.frontend.autoscaling.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 252. `cdv18/helm-charts`

- Chart: `cloned\cdv18__helm-charts\umbrella`
- Source: `template`
- Command: `helm template test cloned\cdv18__helm-charts\umbrella`

```text
Error: template: money-system/charts/calculate-money/templates/hpa.yaml:35:43: executing "money-system/charts/calculate-money/templates/hpa.yaml" at <.Values.frontend.autoscaling.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 253. `cdv18/helm-charts`

- Chart: `cloned\cdv18__helm-charts\umbrella`
- Source: `template`
- Values files: `cloned\cdv18__helm-charts\umbrella\values-dev.yaml`
- Command: `helm template test cloned\cdv18__helm-charts\umbrella -f cloned\cdv18__helm-charts\umbrella\values-dev.yaml`

```text
Error: template: money-system/charts/calculate-money/templates/hpa.yaml:35:43: executing "money-system/charts/calculate-money/templates/hpa.yaml" at <.Values.frontend.autoscaling.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 254. `cdv18/helm-charts`

- Chart: `cloned\cdv18__helm-charts\umbrella`
- Source: `template`
- Values files: `cloned\cdv18__helm-charts\umbrella\values-staging.yaml`
- Command: `helm template test cloned\cdv18__helm-charts\umbrella -f cloned\cdv18__helm-charts\umbrella\values-staging.yaml`

```text
Error: template: money-system/charts/calculate-money/templates/hpa.yaml:35:43: executing "money-system/charts/calculate-money/templates/hpa.yaml" at <.Values.frontend.autoscaling.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 255. `cdv18/helm-charts`

- Chart: `cloned\cdv18__helm-charts\umbrella`
- Source: `template`
- Values files: `cloned\cdv18__helm-charts\umbrella\values-dev.yaml`, `cloned\cdv18__helm-charts\umbrella\values-staging.yaml`
- Command: `helm template test cloned\cdv18__helm-charts\umbrella -f cloned\cdv18__helm-charts\umbrella\values-dev.yaml -f cloned\cdv18__helm-charts\umbrella\values-staging.yaml`

```text
Error: template: money-system/charts/calculate-money/templates/hpa.yaml:35:43: executing "money-system/charts/calculate-money/templates/hpa.yaml" at <.Values.frontend.autoscaling.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 256. `NourBkh/k8s-config-repo`

- Chart: `cloned\NourBkh__k8s-config-repo\helmTestingP\testingprojectHelm`
- Source: `template`
- Command: `helm template test cloned\NourBkh__k8s-config-repo\helmTestingP\testingprojectHelm`

```text
Error: template: pfeprojecthelm/templates/tests/test-connection.yaml:14:70: executing "pfeprojecthelm/templates/tests/test-connection.yaml" at <.Values.service.port>: nil pointer evaluating interface {}.port

Use --debug flag to render out invalid YAML
```

### 257. `superdoo/multitier-webapp`

- Chart: `cloned\superdoo__multitier-webapp\helm\backend`
- Source: `template`
- Command: `helm template test cloned\superdoo__multitier-webapp\helm\backend`

```text
Error: template: mt-backend/templates/service.yaml:10:22: executing "mt-backend/templates/service.yaml" at <.Values.service.port>: nil pointer evaluating interface {}.port

Use --debug flag to render out invalid YAML
```

### 258. `small-hack/argocd-apps`

- Chart: `cloned\small-hack__argocd-apps\ghost\storage`
- Source: `template`
- Command: `helm template test cloned\small-hack__argocd-apps\ghost\storage`

```text
Error: template: ghost-persistence-chart/templates/pod_config.yaml:1:46: executing "ghost-persistence-chart/templates/pod_config.yaml" at <.Values.affinity.key>: nil pointer evaluating interface {}.key

Use --debug flag to render out invalid YAML
```

### 259. `small-hack/argocd-apps`

- Chart: `cloned\small-hack__argocd-apps\home-assistant\storage`
- Source: `template`
- Command: `helm template test cloned\small-hack__argocd-apps\home-assistant\storage`

```text
Error: template: home-assistance-persistence-chart/templates/pod_config.yaml:1:46: executing "home-assistance-persistence-chart/templates/pod_config.yaml" at <.Values.affinity.key>: nil pointer evaluating interface {}.key

Use --debug flag to render out invalid YAML
```

### 260. `small-hack/argocd-apps`

- Chart: `cloned\small-hack__argocd-apps\writefreely\storage`
- Source: `template`
- Command: `helm template test cloned\small-hack__argocd-apps\writefreely\storage`

```text
Error: template: writefreely-persistence-chart/templates/pod_config.yaml:1:46: executing "writefreely-persistence-chart/templates/pod_config.yaml" at <.Values.affinity.key>: nil pointer evaluating interface {}.key

Use --debug flag to render out invalid YAML
```

### 261. `trikorasolns/helm-charts`

- Chart: `cloned\trikorasolns__helm-charts\charts\redis`
- Source: `template`
- Command: `helm template test cloned\trikorasolns__helm-charts\charts\redis`

```text
Error: template: redis/templates/pvc.yaml:1:18: executing "redis/templates/pvc.yaml" at <.Values.persistence.iscsi>: nil pointer evaluating interface {}.iscsi

Use --debug flag to render out invalid YAML
```

### 262. `danskernesdigitalebibliotek/dpl-platform`

- Chart: `cloned\danskernesdigitalebibliotek__dpl-platform\gitops\charts\moduletest-reset-cronjob`
- Source: `template`
- Command: `helm template test cloned\danskernesdigitalebibliotek__dpl-platform\gitops\charts\moduletest-reset-cronjob`

```text
Error: template: moduletest-reset/templates/moduletest-reset-cronjob.yaml:23:60: executing "moduletest-reset/templates/moduletest-reset-cronjob.yaml" at <.Values.syncConfig.source>: nil pointer evaluating interface {}.source

Use --debug flag to render out invalid YAML
```

### 263. `kieranajp/homelab`

- Chart: `cloned\kieranajp__homelab\charts\calibre-web-automated`
- Source: `template`
- Command: `helm template test cloned\kieranajp__homelab\charts\calibre-web-automated`

```text
Error: template: calibre-web-automated/templates/pvc.yaml:12:25: executing "calibre-web-automated/templates/pvc.yaml" at <.Values.persistence.configSize>: nil pointer evaluating interface {}.configSize

Use --debug flag to render out invalid YAML
```

### 264. `kieranajp/homelab`

- Chart: `cloned\kieranajp__homelab\charts\lazylibrarian`
- Source: `template`
- Command: `helm template test cloned\kieranajp__homelab\charts\lazylibrarian`

```text
Error: template: lazylibrarian/templates/pvc.yaml:12:25: executing "lazylibrarian/templates/pvc.yaml" at <.Values.persistence.size>: nil pointer evaluating interface {}.size

Use --debug flag to render out invalid YAML
```

### 265. `kieranajp/homelab`

- Chart: `cloned\kieranajp__homelab\charts\mcp-server`
- Source: `template`
- Command: `helm template test cloned\kieranajp__homelab\charts\mcp-server`

```text
Error: template: mcp-server/templates/service.yaml:11:20: executing "mcp-server/templates/service.yaml" at <.Values.service.port>: nil pointer evaluating interface {}.port

Use --debug flag to render out invalid YAML
```

### 266. `kieranajp/homelab`

- Chart: `cloned\kieranajp__homelab\charts\openclaw`
- Source: `template`
- Command: `helm template test cloned\kieranajp__homelab\charts\openclaw`

```text
Error: template: openclaw/templates/pvc.yaml:12:25: executing "openclaw/templates/pvc.yaml" at <.Values.persistence.size>: nil pointer evaluating interface {}.size

Use --debug flag to render out invalid YAML
```

### 267. `kieranajp/homelab`

- Chart: `cloned\kieranajp__homelab\charts\paperless-ngx`
- Source: `template`
- Command: `helm template test cloned\kieranajp__homelab\charts\paperless-ngx`

```text
Error: template: paperless-ngx/templates/ingressroute.yaml:1:14: executing "paperless-ngx/templates/ingressroute.yaml" at <.Values.ingress.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 268. `kieranajp/homelab`

- Chart: `cloned\kieranajp__homelab\charts\prowlarr`
- Source: `template`
- Command: `helm template test cloned\kieranajp__homelab\charts\prowlarr`

```text
Error: template: prowlarr/templates/pvc.yaml:12:25: executing "prowlarr/templates/pvc.yaml" at <.Values.persistence.size>: nil pointer evaluating interface {}.size

Use --debug flag to render out invalid YAML
```

### 269. `kieranajp/homelab`

- Chart: `cloned\kieranajp__homelab\charts\sonarr`
- Source: `template`
- Command: `helm template test cloned\kieranajp__homelab\charts\sonarr`

```text
Error: template: sonarr/templates/pvc.yaml:12:25: executing "sonarr/templates/pvc.yaml" at <.Values.persistence.size>: nil pointer evaluating interface {}.size

Use --debug flag to render out invalid YAML
```

### 270. `kieranajp/homelab`

- Chart: `cloned\kieranajp__homelab\charts\syncthing`
- Source: `template`
- Command: `helm template test cloned\kieranajp__homelab\charts\syncthing`

```text
Error: template: syncthing/templates/pvc.yaml:12:25: executing "syncthing/templates/pvc.yaml" at <.Values.persistence.size>: nil pointer evaluating interface {}.size

Use --debug flag to render out invalid YAML
```

### 271. `kieranajp/homelab`

- Chart: `cloned\kieranajp__homelab\charts\transmission`
- Source: `template`
- Command: `helm template test cloned\kieranajp__homelab\charts\transmission`

```text
Error: template: transmission/templates/pvc.yaml:12:25: executing "transmission/templates/pvc.yaml" at <.Values.persistence.size>: nil pointer evaluating interface {}.size

Use --debug flag to render out invalid YAML
```

### 272. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\adp-app-config`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\adp-app-config`

```text
Error: template: adp-app-config/templates/deployment.yaml:13:19: executing "adp-app-config/templates/deployment.yaml" at <.Values.global.clusterSize>: nil pointer evaluating interface {}.clusterSize

Use --debug flag to render out invalid YAML
```

### 273. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\agent-config-server`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\agent-config-server`

```text
Error: template: agent-config-server/templates/deployment.yaml:12:19: executing "agent-config-server/templates/deployment.yaml" at <.Values.global.clusterSize>: nil pointer evaluating interface {}.clusterSize

Use --debug flag to render out invalid YAML
```

### 274. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\ai-gateway`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\ai-gateway`

```text
Error: template: ai-gateway/templates/deployment.yaml:12:19: executing "ai-gateway/templates/deployment.yaml" at <.Values.global.clusterSize>: nil pointer evaluating interface {}.clusterSize

Use --debug flag to render out invalid YAML
```

### 275. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\apex`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\apex`

```text
Error: template: apex/templates/apex.yaml:12:19: executing "apex/templates/apex.yaml" at <.Values.global.clusterSize>: nil pointer evaluating interface {}.clusterSize

Use --debug flag to render out invalid YAML
```

### 276. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\bot-agent-server`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\bot-agent-server`

```text
Error: template: bot-agent-server/templates/deployment.yaml:12:19: executing "bot-agent-server/templates/deployment.yaml" at <.Values.global.clusterSize>: nil pointer evaluating interface {}.clusterSize

Use --debug flag to render out invalid YAML
```

### 277. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\bot-data-statistics-server`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\bot-data-statistics-server`

```text
Error: template: bot-data-statistics-server/templates/deployment.yaml:13:19: executing "bot-data-statistics-server/templates/deployment.yaml" at <.Values.global.clusterSize>: nil pointer evaluating interface {}.clusterSize

Use --debug flag to render out invalid YAML
```

### 278. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\bot-evaluate-config-server`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\bot-evaluate-config-server`

```text
Error: template: bot-evaluate-config-server/templates/deployment.yaml:12:19: executing "bot-evaluate-config-server/templates/deployment.yaml" at <.Values.global.clusterSize>: nil pointer evaluating interface {}.clusterSize

Use --debug flag to render out invalid YAML
```

### 279. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\bot-knowledge-config-server`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\bot-knowledge-config-server`

```text
Error: template: bot-knowledge-config-server/templates/deployment.yaml:12:19: executing "bot-knowledge-config-server/templates/deployment.yaml" at <.Values.global.clusterSize>: nil pointer evaluating interface {}.clusterSize

Use --debug flag to render out invalid YAML
```

### 280. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\bot-llm-server`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\bot-llm-server`

```text
Error: template: bot-llm-server/templates/deployment.yaml:12:19: executing "bot-llm-server/templates/deployment.yaml" at <.Values.global.clusterSize>: nil pointer evaluating interface {}.clusterSize

Use --debug flag to render out invalid YAML
```

### 281. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\bot-memory-server`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\bot-memory-server`

```text
Error: template: bot-memory-server/templates/deployment.yaml:12:19: executing "bot-memory-server/templates/deployment.yaml" at <.Values.global.clusterSize>: nil pointer evaluating interface {}.clusterSize

Use --debug flag to render out invalid YAML
```

### 282. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\bot-rag-server`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\bot-rag-server`

```text
Error: template: bot-rag-server/templates/deployment.yaml:12:19: executing "bot-rag-server/templates/deployment.yaml" at <.Values.global.clusterSize>: nil pointer evaluating interface {}.clusterSize

Use --debug flag to render out invalid YAML
```

### 283. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\bot-retrieval-server`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\bot-retrieval-server`

```text
Error: template: bot-retrieval-server/templates/deployment.yaml:12:19: executing "bot-retrieval-server/templates/deployment.yaml" at <.Values.global.clusterSize>: nil pointer evaluating interface {}.clusterSize

Use --debug flag to render out invalid YAML
```

### 284. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\bot-task-config-server`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\bot-task-config-server`

```text
Error: template: bot-task-config-server/templates/deployment.yaml:12:19: executing "bot-task-config-server/templates/deployment.yaml" at <.Values.global.clusterSize>: nil pointer evaluating interface {}.clusterSize

Use --debug flag to render out invalid YAML
```

### 285. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\bot-workflow-dm-server`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\bot-workflow-dm-server`

```text
Error: template: bot-workflow-dm-server/templates/deployment.yaml:13:19: executing "bot-workflow-dm-server/templates/deployment.yaml" at <.Values.global.clusterSize>: nil pointer evaluating interface {}.clusterSize

Use --debug flag to render out invalid YAML
```

### 286. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\channel-msg-svr`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\channel-msg-svr`

```text
Error: template: channel-msg-svr/templates/deployment.yaml:13:19: executing "channel-msg-svr/templates/deployment.yaml" at <.Values.global.clusterSize>: nil pointer evaluating interface {}.clusterSize

Use --debug flag to render out invalid YAML
```

### 287. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\channel-proxy-svr`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\channel-proxy-svr`

```text
Error: template: channel-proxy-svr/templates/deployment.yaml:13:19: executing "channel-proxy-svr/templates/deployment.yaml" at <.Values.global.clusterSize>: nil pointer evaluating interface {}.clusterSize

Use --debug flag to render out invalid YAML
```

### 288. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\channel-token-svr`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\channel-token-svr`

```text
Error: template: channel-token-svr/templates/deployment.yaml:13:19: executing "channel-token-svr/templates/deployment.yaml" at <.Values.global.clusterSize>: nil pointer evaluating interface {}.clusterSize

Use --debug flag to render out invalid YAML
```

### 289. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\chat`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\chat`

```text
Error: template: chat/templates/deployment.yaml:11:19: executing "chat/templates/deployment.yaml" at <.Values.global.clusterSize>: nil pointer evaluating interface {}.clusterSize

Use --debug flag to render out invalid YAML
```

### 290. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\code-interpreter-dispatcher`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\code-interpreter-dispatcher`

```text
Error: template: code-interpreter-dispatcher/templates/deployment.yaml:12:19: executing "code-interpreter-dispatcher/templates/deployment.yaml" at <.Values.global.clusterSize>: nil pointer evaluating interface {}.clusterSize

Use --debug flag to render out invalid YAML
```

### 291. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\common-config-map`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\common-config-map`

```text
Error: template: common-config-map/templates/config-scurl.yaml:12:32: executing "common-config-map/templates/config-scurl.yaml" at <.Values.global.components.db.user>: nil pointer evaluating interface {}.components

Use --debug flag to render out invalid YAML
```

### 292. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\crawler`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\crawler`

```text
Error: template: crawler/templates/deployment-crawler.yaml:10:19: executing "crawler/templates/deployment-crawler.yaml" at <.Values.global.clusterSize>: nil pointer evaluating interface {}.clusterSize

Use --debug flag to render out invalid YAML
```

### 293. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\deploy-data-init`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\deploy-data-init`

```text
Error: template: deploy-data-init/templates/import-vdb-job.yaml:37:34: executing "deploy-data-init/templates/import-vdb-job.yaml" at <.Values.global.components.vdb>: nil pointer evaluating interface {}.components

Use --debug flag to render out invalid YAML
```

### 294. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\deploy-pre-check`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\deploy-pre-check`

```text
Error: template: deploy-pre-check/templates/job.yaml:59:27: executing "deploy-pre-check/templates/job.yaml" at <.Values.global.components.db.host>: nil pointer evaluating interface {}.components

Use --debug flag to render out invalid YAML
```

### 295. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\entity-extractor-server`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\entity-extractor-server`

```text
Error: template: entity-extractor-server/templates/deployment.yaml:11:19: executing "entity-extractor-server/templates/deployment.yaml" at <.Values.global.clusterSize>: nil pointer evaluating interface {}.clusterSize

Use --debug flag to render out invalid YAML
```

### 296. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\entity-extractor-server-sand-box`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\entity-extractor-server-sand-box`

```text
Error: template: entity-extractor-server-sand-box/templates/deployment.yaml:11:19: executing "entity-extractor-server-sand-box/templates/deployment.yaml" at <.Values.global.clusterSize>: nil pointer evaluating interface {}.clusterSize

Use --debug flag to render out invalid YAML
```

### 297. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\etcd`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\etcd`

```text
Error: template: etcd/templates/statefulset.yaml:11:19: executing "etcd/templates/statefulset.yaml" at <.Values.global.clusterSize>: nil pointer evaluating interface {}.clusterSize

Use --debug flag to render out invalid YAML
```

### 298. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\file-converter`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\file-converter`

```text
Error: template: file-converter/templates/deployment-fileconverter.yaml:10:19: executing "file-converter/templates/deployment-fileconverter.yaml" at <.Values.global.clusterSize>: nil pointer evaluating interface {}.clusterSize

Use --debug flag to render out invalid YAML
```

### 299. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\finance`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\finance`

```text
Error: template: finance/templates/deployment.yaml:11:19: executing "finance/templates/deployment.yaml" at <.Values.global.clusterSize>: nil pointer evaluating interface {}.clusterSize

Use --debug flag to render out invalid YAML
```

### 300. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\infosec`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\infosec`

```text
Error: template: infosec/templates/deployment.yaml:11:19: executing "infosec/templates/deployment.yaml" at <.Values.global.clusterSize>: nil pointer evaluating interface {}.clusterSize

Use --debug flag to render out invalid YAML
```

### 301. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\kernel`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\kernel`

```text
Error: template: kernel/templates/deployment.yaml:12:19: executing "kernel/templates/deployment.yaml" at <.Values.global.clusterSize>: nil pointer evaluating interface {}.clusterSize

Use --debug flag to render out invalid YAML
```

### 302. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\lke-accessmanagerserver`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\lke-accessmanagerserver`

```text
Error: template: lke-accessmanagerserver/templates/deployment.yaml:1:31: executing "lke-accessmanagerserver/templates/deployment.yaml" at <.Values.global.helmfileEnvName>: nil pointer evaluating interface {}.helmfileEnvName

Use --debug flag to render out invalid YAML
```

### 303. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\lke-admin`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\lke-admin`

```text
Error: template: lke-admin/templates/deployment.yaml:12:19: executing "lke-admin/templates/deployment.yaml" at <.Values.global.clusterSize>: nil pointer evaluating interface {}.clusterSize

Use --debug flag to render out invalid YAML
```

### 304. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\lke-aiconfmanagerserver`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\lke-aiconfmanagerserver`

```text
Error: template: lke-aiconfmanagerserver/templates/deployment.yaml:1:31: executing "lke-aiconfmanagerserver/templates/deployment.yaml" at <.Values.global.helmfileEnvName>: nil pointer evaluating interface {}.helmfileEnvName

Use --debug flag to render out invalid YAML
```

### 305. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\lke-ti-auth-web`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\lke-ti-auth-web`

```text
Error: template: lke-ti-auth-web/templates/deployment.yaml:1:31: executing "lke-ti-auth-web/templates/deployment.yaml" at <.Values.global.helmfileEnvName>: nil pointer evaluating interface {}.helmfileEnvName

Use --debug flag to render out invalid YAML
```

### 306. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\lke-ti-gateway`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\lke-ti-gateway`

```text
Error: template: lke-ti-gateway/templates/service.yaml:3:17: executing "lke-ti-gateway/templates/service.yaml" at <.Values.global.helmfileEnvName>: nil pointer evaluating interface {}.helmfileEnvName

Use --debug flag to render out invalid YAML
```

### 307. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\lke-ti-ingress-controller`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\lke-ti-ingress-controller`

```text
Error: template: lke-ti-ingress-controller/templates/deployment.yaml:1:31: executing "lke-ti-ingress-controller/templates/deployment.yaml" at <.Values.global.helmfileEnvName>: nil pointer evaluating interface {}.helmfileEnvName

Use --debug flag to render out invalid YAML
```

### 308. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\lke-ti-perm-auth`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\lke-ti-perm-auth`

```text
Error: template: lke-ti-perm-auth/templates/secret.yaml:8:22: executing "lke-ti-perm-auth/templates/secret.yaml" at <$.Values.global.rsa.publicKey>: nil pointer evaluating interface {}.rsa

Use --debug flag to render out invalid YAML
```

### 309. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\llm-manager-server`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\llm-manager-server`

```text
Error: template: llm-manager-server/templates/deployment.yaml:11:19: executing "llm-manager-server/templates/deployment.yaml" at <.Values.global.clusterSize>: nil pointer evaluating interface {}.clusterSize

Use --debug flag to render out invalid YAML
```

### 310. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\op`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\op`

```text
Error: template: op/templates/deployment.yaml:11:19: executing "op/templates/deployment.yaml" at <.Values.global.clusterSize>: nil pointer evaluating interface {}.clusterSize

Use --debug flag to render out invalid YAML
```

### 311. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\platform-manager`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\platform-manager`

```text
Error: template: platform-manager/templates/deployment.yaml:13:19: executing "platform-manager/templates/deployment.yaml" at <.Values.global.clusterSize>: nil pointer evaluating interface {}.clusterSize

Use --debug flag to render out invalid YAML
```

### 312. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\plugin-config-server`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\plugin-config-server`

```text
Error: template: plugin-config-server/templates/deployment.yaml:12:19: executing "plugin-config-server/templates/deployment.yaml" at <.Values.global.clusterSize>: nil pointer evaluating interface {}.clusterSize

Use --debug flag to render out invalid YAML
```

### 313. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\plugin-cos-util`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\plugin-cos-util`

```text
Error: template: plugin-cos-util/templates/deployment.yaml:11:19: executing "plugin-cos-util/templates/deployment.yaml" at <.Values.global.clusterSize>: nil pointer evaluating interface {}.clusterSize

Use --debug flag to render out invalid YAML
```

### 314. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\plugin-exec-server`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\plugin-exec-server`

```text
Error: template: plugin-exec-server/templates/deployment.yaml:12:19: executing "plugin-exec-server/templates/deployment.yaml" at <.Values.global.clusterSize>: nil pointer evaluating interface {}.clusterSize

Use --debug flag to render out invalid YAML
```

### 315. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\pyright`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\pyright`

```text
Error: template: pyright/templates/deployment-pyright.yaml:10:19: executing "pyright/templates/deployment-pyright.yaml" at <.Values.global.clusterSize>: nil pointer evaluating interface {}.clusterSize

Use --debug flag to render out invalid YAML
```

### 316. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\resource-gallery`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\resource-gallery`

```text
Error: template: resource-gallery/templates/deployment.yaml:12:19: executing "resource-gallery/templates/deployment.yaml" at <.Values.global.clusterSize>: nil pointer evaluating interface {}.clusterSize

Use --debug flag to render out invalid YAML
```

### 317. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\resource-schedule-server`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\resource-schedule-server`

```text
Error: template: resource-schedule-server/templates/deployment.yaml:12:19: executing "resource-schedule-server/templates/deployment.yaml" at <.Values.global.clusterSize>: nil pointer evaluating interface {}.clusterSize

Use --debug flag to render out invalid YAML
```

### 318. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\short-url`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\short-url`

```text
Error: template: short-url/templates/deployment.yaml:11:19: executing "short-url/templates/deployment.yaml" at <.Values.global.clusterSize>: nil pointer evaluating interface {}.clusterSize

Use --debug flag to render out invalid YAML
```

### 319. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\vector-db-manager`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\vector-db-manager`

```text
Error: template: vector-db-manager/templates/deployment.yaml:11:19: executing "vector-db-manager/templates/deployment.yaml" at <.Values.global.clusterSize>: nil pointer evaluating interface {}.clusterSize

Use --debug flag to render out invalid YAML
```

### 320. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\web-app`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\web-app`

```text
Error: template: web-app/templates/deployment-webappchatbot.yaml:11:19: executing "web-app/templates/deployment-webappchatbot.yaml" at <.Values.global.clusterSize>: nil pointer evaluating interface {}.clusterSize

Use --debug flag to render out invalid YAML
```

### 321. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\web-parser-server`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\web-parser-server`

```text
Error: template: web-parser-server/templates/deployment.yaml:11:19: executing "web-parser-server/templates/deployment.yaml" at <.Values.global.clusterSize>: nil pointer evaluating interface {}.clusterSize

Use --debug flag to render out invalid YAML
```

### 322. `pickledish/witw-chart`

- Chart: `cloned\pickledish__witw-chart\witw`
- Source: `template`
- Command: `helm template test cloned\pickledish__witw-chart\witw`

```text
Error: template: witw/templates/service.yaml:12:14: executing "witw/templates/service.yaml" at <$.Values.service.port>: nil pointer evaluating interface {}.port

Use --debug flag to render out invalid YAML
```

### 323. `chayeonhee/msa2`

- Chart: `cloned\chayeonhee__msa2\environments\prod-env`
- Source: `template`
- Command: `helm template test cloned\chayeonhee__msa2\environments\prod-env`

```text
Error: template: prod-env/charts/eazybankcommon/templates/service.yaml:12:22: executing "common.service" at <.Values.eazybankcommon.service.port>: nil pointer evaluating interface {}.service

Use --debug flag to render out invalid YAML
```

### 324. `huggingface/responses.js`

- Chart: `cloned\huggingface__responses.js\chart`
- Source: `template`
- Command: `helm template test cloned\huggingface__responses.js\chart`

```text
Error: template: responses-js/templates/hpa.yaml:1:8: executing "responses-js/templates/hpa.yaml" at <$.Values.autoscaling.enabled>: nil pointer evaluating interface {}.enabled

Use --debug flag to render out invalid YAML
```

### 325. `LasseRapo/fabric-cti-sharing`

- Chart: `cloned\LasseRapo__fabric-cti-sharing\bevel-operator-fabric\charts\hlf-ca`
- Source: `template`
- Command: `helm template test cloned\LasseRapo__fabric-cti-sharing\bevel-operator-fabric\charts\hlf-ca`

```text
Error: template: hlf-ca/templates/traefikroute.yaml:1:13: executing "hlf-ca/templates/traefikroute.yaml" at <.Values.traefik.hosts>: nil pointer evaluating interface {}.hosts

Use --debug flag to render out invalid YAML
```

### 326. `LasseRapo/fabric-cti-sharing`

- Chart: `cloned\LasseRapo__fabric-cti-sharing\bevel-operator-fabric\charts\hlf-ordnode`
- Source: `template`
- Command: `helm template test cloned\LasseRapo__fabric-cti-sharing\bevel-operator-fabric\charts\hlf-ordnode`

```text
Error: template: hlf-ordnode/templates/traefikroute.yaml:1:13: executing "hlf-ordnode/templates/traefikroute.yaml" at <.Values.traefik.hosts>: nil pointer evaluating interface {}.hosts

Use --debug flag to render out invalid YAML
```

### 327. `42c-arhayem/Demo`

- Chart: `cloned\42c-arhayem__Demo\artifacts\42c-pixiapi`
- Source: `template`
- Command: `helm template test cloned\42c-arhayem__Demo\artifacts\42c-pixiapi`

```text
Error: template: 42c-pixi/templates/pixidb-service.yaml:12:20: executing "42c-pixi/templates/pixidb-service.yaml" at <.Values.pixidb.listen_port>: nil pointer evaluating interface {}.listen_port

Use --debug flag to render out invalid YAML
```

### 328. `42c-arhayem/Pixi`

- Chart: `cloned\42c-arhayem__Pixi\artifacts\42c-pixiapi`
- Source: `template`
- Command: `helm template test cloned\42c-arhayem__Pixi\artifacts\42c-pixiapi`

```text
Error: template: 42c-pixi/templates/pixidb-service.yaml:12:20: executing "42c-pixi/templates/pixidb-service.yaml" at <.Values.pixidb.listen_port>: nil pointer evaluating interface {}.listen_port

Use --debug flag to render out invalid YAML
```

## `template.required_value` (39)

### 1. `bakseter/whpah`

- Chart: `cloned\bakseter__whpah\manifests\cluster-addons\monitoring`
- Source: `template`
- Command: `helm template test cloned\bakseter__whpah\manifests\cluster-addons\monitoring`

```text
Error: execution error at (monitoring-umbrella/charts/loki/templates/validate.yaml:46:4): Please define loki.storage.bucketName.chunks

Use --debug flag to render out invalid YAML
```

### 2. `AlexisMtr/amphitrite`

- Chart: `cloned\AlexisMtr__amphitrite\helm\amphitrite`
- Source: `template`
- Command: `helm template test cloned\AlexisMtr__amphitrite\helm\amphitrite`

```text
Error: execution error at (amphitrite/templates/secret.yaml:11:27): A connection string to SQLServer is required

Use --debug flag to render out invalid YAML
```

### 3. `Gemeente-DenHaag/helm-charts`

- Chart: `cloned\Gemeente-DenHaag__helm-charts\charts\generic`
- Source: `template`
- Command: `helm template test cloned\Gemeente-DenHaag__helm-charts\charts\generic`

```text
Error: execution error at (generic/templates/NOTES.txt:1:4): 
VALUES VALIDATION:
app: no-name
    You did not specify the application name for the generic app chart. Please
    set name.
app: no-image-repository
    You did not specify the application image repository. Please
    set image.repository.

Use --debug flag to render out invalid YAML
```

### 4. `BIRU-Scop/helm-charts`

- Chart: `cloned\BIRU-Scop__helm-charts\charts\tenzu-back`
- Source: `template`
- Command: `helm template test cloned\BIRU-Scop__helm-charts\charts\tenzu-back`

```text
Error: execution error at (tenzu-back/templates/job.yaml:32:14): A host is mandatory for redis 

Use --debug flag to render out invalid YAML
```

### 5. `DEFRA/coreai-mcu-knowledge-pgv`

- Chart: `cloned\DEFRA__coreai-mcu-knowledge-pgv\helm\coreai-mcu-knowledge-pgv-infra`
- Source: `template`
- Command: `helm template test cloned\DEFRA__coreai-mcu-knowledge-pgv\helm\coreai-mcu-knowledge-pgv-infra`

```text
Error: execution error at (coreai-mcu-knowledge-pgv-infra/templates/userassignedidentity.yaml:1:4): No value found for 'teamMIPrefix' in adp-aso-helm-library template

Use --debug flag to render out invalid YAML
```

### 6. `praveenperera/matrix-element`

- Chart: `cloned\praveenperera__matrix-element\element-web`
- Source: `template`
- Command: `helm template test cloned\praveenperera__matrix-element\element-web`

```text
Error: execution error at (element-web/templates/deployment.yaml:18:28): Must specify a default homeserver

Use --debug flag to render out invalid YAML
```

### 7. `elasticio/helm-charts`

- Chart: `cloned\elasticio__helm-charts\cluster`
- Source: `template`
- Command: `helm template test cloned\elasticio__helm-charts\cluster`

```text
Error: execution error at (cluster/charts/platform-storage-slugs/templates/service-loadbalancer.yaml:15:21): You must provide load balancer IP for slugs storage

Use --debug flag to render out invalid YAML
```

### 8. `elasticio/helm-charts`

- Chart: `cloned\elasticio__helm-charts\gitreceiver`
- Source: `template`
- Command: `helm template test cloned\elasticio__helm-charts\gitreceiver`

```text
Error: execution error at (gitreceiver/templates/service-loadbalancer.yaml:13:21): You need provide load balancer IP for gitreceiver

Use --debug flag to render out invalid YAML
```

### 9. `elasticio/helm-charts`

- Chart: `cloned\elasticio__helm-charts\platform-storage-slugs`
- Source: `template`
- Command: `helm template test cloned\elasticio__helm-charts\platform-storage-slugs`

```text
Error: execution error at (platform-storage-slugs/templates/service-loadbalancer.yaml:15:21): You must provide load balancer IP for slugs storage

Use --debug flag to render out invalid YAML
```

### 10. `LegitSecurity/helm-charts`

- Chart: `cloned\LegitSecurity__helm-charts\charts\legit-kubernetes-agent`
- Source: `template`
- Command: `helm template test cloned\LegitSecurity__helm-charts\charts\legit-kubernetes-agent`

```text
Error: execution error at (legit-kubernetes-agent/templates/secret.yaml:7:17): agent.identifier is required!

Use --debug flag to render out invalid YAML
```

### 11. `JJimmyFlynn/laravel-eks`

- Chart: `cloned\JJimmyFlynn__laravel-eks\k8s\helm\laravel-app`
- Source: `template`
- Command: `helm template test cloned\JJimmyFlynn__laravel-eks\k8s\helm\laravel-app`

```text
Error: execution error at (laravel-app/templates/sa-external-secrets.yaml:7:35): A role ARN for the secrets provider service account is required

Use --debug flag to render out invalid YAML
```

### 12. `LeoShivas/GitOps`

- Chart: `cloned\LeoShivas__GitOps\kubernetes\transmission\helm\transmission`
- Source: `template`
- Command: `helm template test cloned\LeoShivas__GitOps\kubernetes\transmission\helm\transmission`

```text
Error: execution error at (transmission/templates/deployment.yaml:94:26): An existing claim name for storing data is required !

Use --debug flag to render out invalid YAML
```

### 13. `Outburn-IL/fume-deployments`

- Chart: `cloned\Outburn-IL__fume-deployments\helm\fume`
- Source: `template`
- Command: `helm template test cloned\Outburn-IL__fume-deployments\helm\fume`

```text
Error: execution error at (fume/templates/frontend-deployment.yaml:23:28): CANONICAL_BASE_URL is required. Set via --set configMap.CANONICAL_BASE_URL="https://fume.your-company.com"

Use --debug flag to render out invalid YAML
```

### 14. `Outburn-IL/fume-deployments`

- Chart: `cloned\Outburn-IL__fume-deployments\helm\fume`
- Source: `template`
- Values files: `cloned\Outburn-IL__fume-deployments\helm\fume\values.prod.yaml`
- Command: `helm template test cloned\Outburn-IL__fume-deployments\helm\fume -f cloned\Outburn-IL__fume-deployments\helm\fume\values.prod.yaml`

```text
Error: execution error at (fume/templates/configmap.yaml:1:4): CANONICAL_BASE_URL is required. Set via --set configMap.CANONICAL_BASE_URL="https://fume.your-company.com"

Use --debug flag to render out invalid YAML
```

### 15. `dariuszbolkowski/ezdrp_process`

- Chart: `cloned\dariuszbolkowski__ezdrp_process\charts\ezd-backend\1.0.0`
- Source: `template`
- Command: `helm template test cloned\dariuszbolkowski__ezdrp_process\charts\ezd-backend\1.0.0`

```text
Error: execution error at (ezd-backend/templates/ezdrp-app/relationaldb-secret.yaml:11:38): Password is required.

Use --debug flag to render out invalid YAML
```

### 16. `ValeriiVasianovych/argocd-kubernetes-infrastructure`

- Chart: `cloned\ValeriiVasianovych__argocd-kubernetes-infrastructure\icdev-01\helmCharts\eck-stack\charts\eck-agent`
- Source: `template`
- Command: `helm template test cloned\ValeriiVasianovych__argocd-kubernetes-infrastructure\icdev-01\helmCharts\eck-stack\charts\eck-agent`

```text
Error: execution error at (eck-agent/templates/elastic-agent.yaml:16:5): At least one of daemonSet or deployment or statefulSet is required

Use --debug flag to render out invalid YAML
```

### 17. `ValeriiVasianovych/argocd-kubernetes-infrastructure`

- Chart: `cloned\ValeriiVasianovych__argocd-kubernetes-infrastructure\icdev-01\helmCharts\eck-stack\charts\eck-beats`
- Source: `template`
- Command: `helm template test cloned\ValeriiVasianovych__argocd-kubernetes-infrastructure\icdev-01\helmCharts\eck-stack\charts\eck-beats`

```text
Error: execution error at (eck-beats/templates/beats.yaml:16:5): At least one of daemonSet or deployment is required for a functional Beat

Use --debug flag to render out invalid YAML
```

### 18. `ValeriiVasianovych/argocd-kubernetes-infrastructure`

- Chart: `cloned\ValeriiVasianovych__argocd-kubernetes-infrastructure\icdev-01\helmCharts\eck-stack\charts\eck-enterprise-search`
- Source: `template`
- Command: `helm template test cloned\ValeriiVasianovych__argocd-kubernetes-infrastructure\icdev-01\helmCharts\eck-stack\charts\eck-enterprise-search`

```text
Error: execution error at (eck-enterprise-search/templates/enterprisesearch.yaml:24:5): At least one of configRef or elasticsearchRef is required

Use --debug flag to render out invalid YAML
```

### 19. `colearendt/helm`

- Chart: `cloned\colearendt__helm\charts\postgrest`
- Source: `template`
- Command: `helm template test cloned\colearendt__helm\charts\postgrest`

```text
Error: execution error at (postgrest/templates/secret.yaml:8:8): Missing required value postgrest.dbUri

Use --debug flag to render out invalid YAML
```

### 20. `pschichtel/helm-charts`

- Chart: `cloned\pschichtel__helm-charts\charts\coturn`
- Source: `template`
- Command: `helm template test cloned\pschichtel__helm-charts\charts\coturn`

```text
Error: execution error at (coturn/templates/deployment.yaml:23:31): Either static users or a static secret is required!

Use --debug flag to render out invalid YAML
```

### 21. `pschichtel/helm-charts`

- Chart: `cloned\pschichtel__helm-charts\charts\docspell`
- Source: `template`
- Command: `helm template test cloned\pschichtel__helm-charts\charts\docspell`

```text
Error: execution error at (docspell/templates/secret-restserver.yml:10:14): The server secret must be a long random string!

Use --debug flag to render out invalid YAML
```

### 22. `pschichtel/helm-charts`

- Chart: `cloned\pschichtel__helm-charts\charts\graylog`
- Source: `template`
- Command: `helm template test cloned\pschichtel__helm-charts\charts\graylog`

```text
Error: execution error at (graylog/templates/secret.yaml:16:6): No root password and no existing secret was provided!

Use --debug flag to render out invalid YAML
```

### 23. `pschichtel/helm-charts`

- Chart: `cloned\pschichtel__helm-charts\charts\keycloak`
- Source: `template`
- Command: `helm template test cloned\pschichtel__helm-charts\charts\keycloak`

```text
Error: execution error at (keycloak/templates/keycloak.yaml:13:23): .db is required!

Use --debug flag to render out invalid YAML
```

### 24. `pschichtel/helm-charts`

- Chart: `cloned\pschichtel__helm-charts\charts\keycloak-backup`
- Source: `template`
- Command: `helm template test cloned\pschichtel__helm-charts\charts\keycloak-backup`

```text
Error: execution error at (keycloak-backup/templates/cronjob.yaml:42:57): .backup.targetKeycloak is required!

Use --debug flag to render out invalid YAML
```

### 25. `pschichtel/helm-charts`

- Chart: `cloned\pschichtel__helm-charts\charts\mediawiki`
- Source: `template`
- Command: `helm template test cloned\pschichtel__helm-charts\charts\mediawiki`

```text
Error: execution error at (mediawiki/templates/secret.yaml:10:20): localSettings is required!

Use --debug flag to render out invalid YAML
```

### 26. `pschichtel/helm-charts`

- Chart: `cloned\pschichtel__helm-charts\charts\nextcloud-notify_push`
- Source: `template`
- Command: `helm template test cloned\pschichtel__helm-charts\charts\nextcloud-notify_push`

```text
Error: execution error at (nextcloud-notify_push/templates/secret.yaml:7:53): databaseUrl is required!

Use --debug flag to render out invalid YAML
```

### 27. `pschichtel/helm-charts`

- Chart: `cloned\pschichtel__helm-charts\charts\powerdns`
- Source: `template`
- Command: `helm template test cloned\pschichtel__helm-charts\charts\powerdns`

```text
Error: execution error at (powerdns/templates/deployment.yaml:70:14): A value or secretRef must be provided for PDNS_DB_HOST!

Use --debug flag to render out invalid YAML
```

### 28. `pschichtel/helm-charts`

- Chart: `cloned\pschichtel__helm-charts\charts\powerdns-admin`
- Source: `template`
- Command: `helm template test cloned\pschichtel__helm-charts\charts\powerdns-admin`

```text
Error: execution error at (powerdns-admin/templates/deployment.yaml:47:63): .Values.powerdnsAdmin.config.secretKey or .Values.powerdnsAdmin.config.secretKeySecret is required!

Use --debug flag to render out invalid YAML
```

### 29. `pschichtel/helm-charts`

- Chart: `cloned\pschichtel__helm-charts\charts\spicedb`
- Source: `template`
- Command: `helm template test cloned\pschichtel__helm-charts\charts\spicedb`

```text
Error: execution error at (spicedb/templates/deployment.yaml:81:14): .spicedb.grpcPresharedKey.value or .secretName is required!

Use --debug flag to render out invalid YAML
```

### 30. `pschichtel/helm-charts`

- Chart: `cloned\pschichtel__helm-charts\charts\web-app`
- Source: `template`
- Command: `helm template test cloned\pschichtel__helm-charts\charts\web-app`

```text
Error: execution error at (web-app/templates/deployment.yaml:40:48): the image repository is required!

Use --debug flag to render out invalid YAML
```

### 31. `pschichtel/helm-charts`

- Chart: `cloned\pschichtel__helm-charts\charts\webdav`
- Source: `template`
- Command: `helm template test cloned\pschichtel__helm-charts\charts\webdav`

```text
Error: execution error at (webdav/templates/deployment.yaml:64:44): storage.volume is required!

Use --debug flag to render out invalid YAML
```

### 32. `blinklabs-io/helm-charts`

- Chart: `cloned\blinklabs-io__helm-charts\charts\demeter-fabric`
- Source: `template`
- Command: `helm template test cloned\blinklabs-io__helm-charts\charts\demeter-fabric`

```text
Error: execution error at (demeter-fabric/templates/sts.yaml:23:28): config.kafka.consumerCacheName is required

Use --debug flag to render out invalid YAML
```

### 33. `blinklabs-io/helm-charts`

- Chart: `cloned\blinklabs-io__helm-charts\charts\wireguard`
- Source: `template`
- Command: `helm template test cloned\blinklabs-io__helm-charts\charts\wireguard`

```text
Error: execution error at (wireguard/templates/secret.yaml:2:11): wireguard.privateKey is required when wireguard.existingSecret is not set

Use --debug flag to render out invalid YAML
```

### 34. `trueforge-org/truecharts`

- Chart: `cloned\trueforge-org__truecharts\charts\stable\app-template`
- Source: `template`
- Command: `helm template test cloned\trueforge-org__truecharts\charts\stable\app-template`

```text
Error: execution error at (app-template/templates/common.yaml:1:3): Service - Expected non-empty [port.port]

Use --debug flag to render out invalid YAML
```

### 35. `trueforge-org/truecharts`

- Chart: `cloned\trueforge-org__truecharts\charts\stable\authentik`
- Source: `template`
- Command: `helm template test cloned\trueforge-org__truecharts\charts\stable\authentik`

```text
Error: execution error at (authentik/templates/common.yaml:97:3): Ingress - Expected ingress [main] to be enabled. This chart is designed to work only with ingress enabled.

Use --debug flag to render out invalid YAML
```

### 36. `trueforge-org/truecharts`

- Chart: `cloned\trueforge-org__truecharts\charts\stable\nextcloud`
- Source: `template`
- Command: `helm template test cloned\trueforge-org__truecharts\charts\stable\nextcloud`

```text
Error: execution error at (nextcloud/templates/common.yaml:92:4): Expected non-empty [ip] value on [hostAliases].

Use --debug flag to render out invalid YAML
```

### 37. `trueforge-org/truecharts`

- Chart: `cloned\trueforge-org__truecharts\charts\stable\slink`
- Source: `template`
- Command: `helm template test cloned\trueforge-org__truecharts\charts\stable\slink`

```text
Error: execution error at (slink/templates/common.yaml:1:3): Ingress - Expected ingress [main] to be enabled. This chart is designed to work only with ingress enabled.

Use --debug flag to render out invalid YAML
```

### 38. `trueforge-org/truecharts`

- Chart: `cloned\trueforge-org__truecharts\charts\stable\vaultwarden`
- Source: `template`
- Command: `helm template test cloned\trueforge-org__truecharts\charts\stable\vaultwarden`

```text
Error: execution error at (vaultwarden/templates/common.yaml:17:3): Ingress - Expected ingress [main] to be enabled. This chart is designed to work only with ingress enabled.

Use --debug flag to render out invalid YAML
```

### 39. `trueforge-org/truecharts`

- Chart: `cloned\trueforge-org__truecharts\charts\stable\wireguard`
- Source: `template`
- Command: `helm template test cloned\trueforge-org__truecharts\charts\stable\wireguard`

```text
Error: execution error at (wireguard/templates/common.yaml:19:3): Volumes - Expected the key [enabled] in [persistence.configfile] to exist

Use --debug flag to render out invalid YAML
```

## `dependency.missing_repository` (37)

### 1. `pluralsh/plural-helm-charts`

- Chart: `cloned\pluralsh__plural-helm-charts\charts\airbyte`
- Source: `dependency`

```text
Error: no repository definition for https://charts.bitnami.com/bitnami, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts, https://pluralsh.github.io/plural-helm-charts. Please add the missing repos via 'helm repo add'
```

### 2. `flytrap/docker-admin`

- Chart: `cloned\flytrap__docker-admin\charts\apisix-1.3`
- Source: `dependency`

```text
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.apiseven.com, https://charts.apiseven.com. Please add the missing repos via 'helm repo add'
```

### 3. `akv-akv/de_research`

- Chart: `cloned\akv-akv__de_research\kubernetes\charts\airbyte`
- Source: `dependency`

```text
Error: no repository definition for https://charts.bitnami.com/bitnami, https://airbytehq.github.io/helm-charts/, https://airbytehq.github.io/helm-charts/, https://airbytehq.github.io/helm-charts/, https://airbytehq.github.io/helm-charts/, https://airbytehq.github.io/helm-charts/, https://airbytehq.github.io/helm-charts/, https://airbytehq.github.io/helm-charts/, https://airbytehq.github.io/helm-charts/, https://airbytehq.github.io/helm-charts/, https://airbytehq.github.io/helm-charts/, https://airbytehq.github.io/helm-charts/, https://airbytehq.github.io/helm-charts/, https://airbytehq.github.io/helm-charts/. Please add the missing repos via 'helm repo add'
```

### 4. `MirahImage/postfacto`

- Chart: `cloned\MirahImage__postfacto\tkg\postfacto`
- Source: `dependency`

```text
Error: no repository definition for @bitnami, @bitnami. Please add them via 'helm repo add'
```

### 5. `unicef/magasin`

- Chart: `cloned\unicef__magasin\helm\dagster`
- Source: `dependency`

```text
Error: no repository definition for https://raw.githubusercontent.com/bitnami/charts/eb5f9a9513d987b519f0ecd732e7031241c50328/bitnami, https://raw.githubusercontent.com/bitnami/charts/eb5f9a9513d987b519f0ecd732e7031241c50328/bitnami, https://raw.githubusercontent.com/bitnami/charts/eb5f9a9513d987b519f0ecd732e7031241c50328/bitnami. Please add the missing repos via 'helm repo add'
```

### 6. `ConductionNL/berichtservice`

- Chart: `cloned\ConductionNL__berichtservice\api\helm`
- Source: `dependency`

```text
load.go:138: Warning: Dependency locking is handled in Chart.lock since apiVersion "v2". We recommend migrating to Chart.lock.
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### 7. `Veronneau-Techno-Conseil/comax-commons`

- Chart: `cloned\Veronneau-Techno-Conseil__comax-commons\helm_referee`
- Source: `dependency`

```text
Error: no repository definition for https://mongodb.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### 8. `limagbz/data-mesh-yelp`

- Chart: `cloned\limagbz__data-mesh-yelp\infra\monitoring\loki\helm`
- Source: `dependency`

```text
Error: no repository definition for https://charts.min.io/, https://grafana.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### 9. `efremov-it/security`

- Chart: `cloned\efremov-it__security\practice\weeks\2-3\secureCodeBox\operator`
- Source: `dependency`

```text
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### 10. `gersonrs/hand-on-airflow-with-minio`

- Chart: `cloned\gersonrs__hand-on-airflow-with-minio\airflow`
- Source: `dependency`

```text
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### 11. `shini4i/charts`

- Chart: `cloned\shini4i__charts\charts\app`
- Source: `dependency`

```text
Error: no repository definition for https://shini4i.github.io/charts/. Please add the missing repos via 'helm repo add'
```

### 12. `otus-kuber-2020-07/israodin_platform`

- Chart: `cloned\otus-kuber-2020-07__israodin_platform\kubernetes-templating\hipster-shop`
- Source: `dependency`

```text
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### 13. `openlit/helm`

- Chart: `cloned\openlit__helm\charts\openlit`
- Source: `dependency`

```text
Error: no repository definition for https://openlit.github.io/helm/. Please add the missing repos via 'helm repo add'
```

### 14. `deresegetachew/playing-helm`

- Chart: `cloned\deresegetachew__playing-helm\kanban-backend`
- Source: `dependency`

```text
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### 15. `Abinesh0206/gitlab-ce`

- Chart: `cloned\Abinesh0206__gitlab-ce`
- Source: `dependency`

```text
Error: no repository definition for https://charts.jetstack.io/, https://prometheus-community.github.io/helm-charts, https://charts.bitnami.com/bitnami, https://charts.gitlab.io/, https://charts.bitnami.com/bitnami, https://charts.gitlab.io/, https://helm.traefik.io/traefik, https://haproxytech.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### 16. `samfms/truecharts`

- Chart: `cloned\samfms__truecharts\deprecated\bitwarden\1.2.5`
- Source: `dependency`

```text
Error: no repository definition for https://truecharts.org/, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### 17. `Blankeeir/demo`

- Chart: `cloned\Blankeeir__demo\infrastructure\helm\rafiki`
- Source: `dependency`

```text
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://interledger.github.io/helm-charts, https://interledger.github.io/helm-charts, https://interledger.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### 18. `randoli/helm-charts`

- Chart: `cloned\randoli__helm-charts\charts\loki`
- Source: `dependency`

```text
Error: no repository definition for https://charts.min.io/, https://grafana.github.io/helm-charts, https://grafana.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### 19. `dylanjustice/the-grid`

- Chart: `cloned\dylanjustice__the-grid\gitops\workloads\kube-prometheus`
- Source: `dependency`

```text
Error: no repository definition for https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### 20. `praveenperera/matrix-element`

- Chart: `cloned\praveenperera__matrix-element\matrix-synapse`
- Source: `dependency`

```text
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### 21. `suomitek/suomitek-charts`

- Chart: `cloned\suomitek__suomitek-charts\suomitek\airflow`
- Source: `dependency`

```text
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### 22. `ndawpa/guiadodevops`

- Chart: `cloned\ndawpa__guiadodevops\kubernetes\helm\airflow-1.18.0\airflow`
- Source: `dependency`

```text
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### 23. `TrevorSquillario/iDRAC-Telemetry-Ansible-Demo`

- Chart: `cloned\TrevorSquillario__iDRAC-Telemetry-Ansible-Demo\k8s\helm\app`
- Source: `dependency`

```text
Error: no repository definition for https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### 24. `dynatrace-wwse/enablement-kubernetes-opentelemetry`

- Chart: `cloned\dynatrace-wwse__enablement-kubernetes-opentelemetry\.devcontainer\apps\astroshop\helm\dt-otel-demo-helm`
- Source: `dependency`

```text
Error: no repository definition for https://open-telemetry.github.io/opentelemetry-helm-charts, https://open-telemetry.github.io/opentelemetry-helm-charts. Please add the missing repos via 'helm repo add'
```

### 25. `VadimShtukan/otus_homework`

- Chart: `cloned\VadimShtukan__otus_homework\architect\lesson05\kubernetis\helm-chart`
- Source: `dependency`

```text
Error: no repository definition for https://kubernetes-charts.storage.googleapis.com. Please add the missing repos via 'helm repo add'
```

### 26. `dariuszbolkowski/ezdrp_process`

- Chart: `cloned\dariuszbolkowski__ezdrp_process\charts\ezd-backend\1.0.0\charts\mongodb`
- Source: `dependency`

```text
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### 27. `Jozefwl/applications-waldoserver`

- Chart: `cloned\Jozefwl__applications-waldoserver\Applications\ERPNext\erpnext-official\erpnext`
- Source: `dependency`

```text
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### 28. `NourBkh/k8s-config-repo`

- Chart: `cloned\NourBkh__k8s-config-repo\monitoring`
- Source: `dependency`

```text
Error: no repository definition for https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### 29. `superdoo/multitier-webapp`

- Chart: `cloned\superdoo__multitier-webapp\helm\postgresql`
- Source: `dependency`

```text
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### 30. `wadhekarpankaj/k8s-infrastructure`

- Chart: `cloned\wadhekarpankaj__k8s-infrastructure\charts\argo-cd`
- Source: `dependency`

```text
Error: no repository definition for https://dandydeveloper.github.io/charts/. Please add the missing repos via 'helm repo add'
```

### 31. `sookeke/jps-pidp`

- Chart: `cloned\sookeke__jps-pidp\charts\nginx`
- Source: `dependency`

```text
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### 32. `HelmCI/ci-mon`

- Chart: `cloned\HelmCI__ci-mon\charts\jaeger`
- Source: `dependency`

```text
Error: no repository definition for https://charts.helm.sh/incubator, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

### 33. `tetratelabs/charts`

- Chart: `cloned\tetratelabs__charts\charts\demos\istio-monitoring-demo`
- Source: `dependency`

```text
Error: no repository definition for https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### 34. `tetratelabs/helm-charts`

- Chart: `cloned\tetratelabs__helm-charts\charts\demos\istio-monitoring-demo`
- Source: `dependency`

```text
Error: no repository definition for https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### 35. `tis-system/charts`

- Chart: `cloned\tis-system__charts\charts\demos\istio-monitoring-demo\0.1.0`
- Source: `dependency`

```text
Error: no repository definition for https://grafana.github.io/helm-charts, https://prometheus-community.github.io/helm-charts, https://prometheus-community.github.io/helm-charts. Please add the missing repos via 'helm repo add'
```

### 36. `p-/PyroDocker`

- Chart: `cloned\p-__PyroDocker\compose\signoz\deploy\kubernetes\platform`
- Source: `dependency`

```text
Error: no repository definition for https://charts.bitnami.com/bitnami, https://charts.bitnami.com/bitnami, https://charts.helm.sh/incubator. Please add the missing repos via 'helm repo add'
```

### 37. `otus-kuber-2020-04/israodin_platform`

- Chart: `cloned\otus-kuber-2020-04__israodin_platform\kubernetes-templating\hipster-shop`
- Source: `dependency`

```text
Error: no repository definition for https://charts.bitnami.com/bitnami. Please add the missing repos via 'helm repo add'
```

## `template.runtime_eval` (27)

### 1. `t9k/apps`

- Chart: `cloned\t9k__apps\archived\vllm-smls-dcu\chart`
- Source: `template`
- Command: `helm template test cloned\t9k__apps\archived\vllm-smls-dcu\chart`

```text
Error: template: vllm-dcu/templates/simplemlservice.yaml:50:15: executing "vllm-dcu/templates/simplemlservice.yaml" at <eq (index .Values.resources.limits "hygon.com/dcu") 1>: error calling eq: incompatible types for comparison

Use --debug flag to render out invalid YAML
```

### 2. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.118\charts\rabbitmq`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.118\charts\rabbitmq`

```text
Error: template: rabbitmq/templates/rabbitmq-statefulset.yaml:4:20: executing "rabbitmq/templates/rabbitmq-statefulset.yaml" at <{{template "search.fullname" .}}>: template "search.fullname" not defined

Use --debug flag to render out invalid YAML
```

### 3. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.121\charts\rabbitmq`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.121\charts\rabbitmq`

```text
Error: template: rabbitmq/templates/rabbitmq-statefulset.yaml:4:20: executing "rabbitmq/templates/rabbitmq-statefulset.yaml" at <{{template "search.fullname" .}}>: template "search.fullname" not defined

Use --debug flag to render out invalid YAML
```

### 4. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.123\charts\rabbitmq`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.123\charts\rabbitmq`

```text
Error: template: rabbitmq/templates/rabbitmq-statefulset.yaml:4:20: executing "rabbitmq/templates/rabbitmq-statefulset.yaml" at <{{template "search.fullname" .}}>: template "search.fullname" not defined

Use --debug flag to render out invalid YAML
```

### 5. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.125\charts\rabbitmq`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.125\charts\rabbitmq`

```text
Error: template: rabbitmq/templates/rabbitmq-statefulset.yaml:4:20: executing "rabbitmq/templates/rabbitmq-statefulset.yaml" at <{{template "search.fullname" .}}>: template "search.fullname" not defined

Use --debug flag to render out invalid YAML
```

### 6. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.128\charts\rabbitmq`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.128\charts\rabbitmq`

```text
Error: template: rabbitmq/templates/rabbitmq-statefulset.yaml:4:20: executing "rabbitmq/templates/rabbitmq-statefulset.yaml" at <{{template "search.fullname" .}}>: template "search.fullname" not defined

Use --debug flag to render out invalid YAML
```

### 7. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.138\charts\rabbitmq`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.138\charts\rabbitmq`

```text
Error: template: rabbitmq/templates/rabbitmq-statefulset.yaml:4:20: executing "rabbitmq/templates/rabbitmq-statefulset.yaml" at <{{template "search.fullname" .}}>: template "search.fullname" not defined

Use --debug flag to render out invalid YAML
```

### 8. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144\charts\rabbitmq`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144\charts\rabbitmq`

```text
Error: template: rabbitmq/templates/rabbitmq-statefulset.yaml:4:20: executing "rabbitmq/templates/rabbitmq-statefulset.yaml" at <{{template "search.fullname" .}}>: template "search.fullname" not defined

Use --debug flag to render out invalid YAML
```

### 9. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.0-kommch\charts\rabbitmq`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.0-kommch\charts\rabbitmq`

```text
Error: template: rabbitmq/templates/rabbitmq-statefulset.yaml:4:20: executing "rabbitmq/templates/rabbitmq-statefulset.yaml" at <{{template "search.fullname" .}}>: template "search.fullname" not defined

Use --debug flag to render out invalid YAML
```

### 10. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.2-kommch\charts\rabbitmq`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.2-kommch\charts\rabbitmq`

```text
Error: template: rabbitmq/templates/rabbitmq-statefulset.yaml:4:20: executing "rabbitmq/templates/rabbitmq-statefulset.yaml" at <{{template "search.fullname" .}}>: template "search.fullname" not defined

Use --debug flag to render out invalid YAML
```

### 11. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.4-kommch\charts\rabbitmq`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.4-kommch\charts\rabbitmq`

```text
Error: template: rabbitmq/templates/rabbitmq-statefulset.yaml:4:20: executing "rabbitmq/templates/rabbitmq-statefulset.yaml" at <{{template "search.fullname" .}}>: template "search.fullname" not defined

Use --debug flag to render out invalid YAML
```

### 12. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.6-1-kommch\charts\rabbitmq`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.6-1-kommch\charts\rabbitmq`

```text
Error: template: rabbitmq/templates/rabbitmq-statefulset.yaml:4:20: executing "rabbitmq/templates/rabbitmq-statefulset.yaml" at <{{template "search.fullname" .}}>: template "search.fullname" not defined

Use --debug flag to render out invalid YAML
```

### 13. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.6-1.kommch\charts\rabbitmq`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.6-1.kommch\charts\rabbitmq`

```text
Error: template: rabbitmq/templates/rabbitmq-statefulset.yaml:4:20: executing "rabbitmq/templates/rabbitmq-statefulset.yaml" at <{{template "search.fullname" .}}>: template "search.fullname" not defined

Use --debug flag to render out invalid YAML
```

### 14. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.7-1-kommch\charts\rabbitmq`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.7-1-kommch\charts\rabbitmq`

```text
Error: template: rabbitmq/templates/rabbitmq-statefulset.yaml:4:20: executing "rabbitmq/templates/rabbitmq-statefulset.yaml" at <{{template "search.fullname" .}}>: template "search.fullname" not defined

Use --debug flag to render out invalid YAML
```

### 15. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.7-2-kommch\charts\rabbitmq`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.144.7-2-kommch\charts\rabbitmq`

```text
Error: template: rabbitmq/templates/rabbitmq-statefulset.yaml:4:20: executing "rabbitmq/templates/rabbitmq-statefulset.yaml" at <{{template "search.fullname" .}}>: template "search.fullname" not defined

Use --debug flag to render out invalid YAML
```

### 16. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-1-kommch\charts\rabbitmq`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-1-kommch\charts\rabbitmq`

```text
Error: template: rabbitmq/templates/rabbitmq-statefulset.yaml:4:20: executing "rabbitmq/templates/rabbitmq-statefulset.yaml" at <{{template "search.fullname" .}}>: template "search.fullname" not defined

Use --debug flag to render out invalid YAML
```

### 17. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-10-kommch\charts\rabbitmq`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-10-kommch\charts\rabbitmq`

```text
Error: template: rabbitmq/templates/rabbitmq-statefulset.yaml:4:20: executing "rabbitmq/templates/rabbitmq-statefulset.yaml" at <{{template "search.fullname" .}}>: template "search.fullname" not defined

Use --debug flag to render out invalid YAML
```

### 18. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-11-kommch\charts\rabbitmq`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-11-kommch\charts\rabbitmq`

```text
Error: template: rabbitmq/templates/rabbitmq-statefulset.yaml:4:20: executing "rabbitmq/templates/rabbitmq-statefulset.yaml" at <{{template "search.fullname" .}}>: template "search.fullname" not defined

Use --debug flag to render out invalid YAML
```

### 19. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-2-kommch\charts\rabbitmq`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-2-kommch\charts\rabbitmq`

```text
Error: template: rabbitmq/templates/rabbitmq-statefulset.yaml:4:20: executing "rabbitmq/templates/rabbitmq-statefulset.yaml" at <{{template "search.fullname" .}}>: template "search.fullname" not defined

Use --debug flag to render out invalid YAML
```

### 20. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-3-kommch\charts\rabbitmq`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-3-kommch\charts\rabbitmq`

```text
Error: template: rabbitmq/templates/rabbitmq-statefulset.yaml:4:20: executing "rabbitmq/templates/rabbitmq-statefulset.yaml" at <{{template "search.fullname" .}}>: template "search.fullname" not defined

Use --debug flag to render out invalid YAML
```

### 21. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-6-kommch\charts\rabbitmq`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-6-kommch\charts\rabbitmq`

```text
Error: template: rabbitmq/templates/rabbitmq-statefulset.yaml:4:20: executing "rabbitmq/templates/rabbitmq-statefulset.yaml" at <{{template "search.fullname" .}}>: template "search.fullname" not defined

Use --debug flag to render out invalid YAML
```

### 22. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-8-kommch\charts\rabbitmq`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-8-kommch\charts\rabbitmq`

```text
Error: template: rabbitmq/templates/rabbitmq-statefulset.yaml:4:20: executing "rabbitmq/templates/rabbitmq-statefulset.yaml" at <{{template "search.fullname" .}}>: template "search.fullname" not defined

Use --debug flag to render out invalid YAML
```

### 23. `KommunikativCh/rancher-charts`

- Chart: `cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-9-kommch\charts\rabbitmq`
- Source: `template`
- Command: `helm template test cloned\KommunikativCh__rancher-charts\charts\metasfresh\5.145.1-9-kommch\charts\rabbitmq`

```text
Error: template: rabbitmq/templates/rabbitmq-statefulset.yaml:4:20: executing "rabbitmq/templates/rabbitmq-statefulset.yaml" at <{{template "search.fullname" .}}>: template "search.fullname" not defined

Use --debug flag to render out invalid YAML
```

### 24. `small-hack/argocd-apps`

- Chart: `cloned\small-hack__argocd-apps\demo\pixelfed\storage\pvc`
- Source: `template`
- Command: `helm template test cloned\small-hack__argocd-apps\demo\pixelfed\storage\pvc`

```text
Error: template: pixelfed-pvc/templates/peertube_pvc.yaml:1:7: executing "pixelfed-pvc/templates/peertube_pvc.yaml" at <eq .Values.pvc.enabled "true">: error calling eq: incompatible types for comparison

Use --debug flag to render out invalid YAML
```

### 25. `small-hack/argocd-apps`

- Chart: `cloned\small-hack__argocd-apps\mastodon\small-hack\storage\pvc`
- Source: `template`
- Command: `helm template test cloned\small-hack__argocd-apps\mastodon\small-hack\storage\pvc`

```text
Error: template: mastodon-pvcs/templates/valkey_pvc.yaml:1:7: executing "mastodon-pvcs/templates/valkey_pvc.yaml" at <eq .Values.valkey_pvc.enabled "true">: error calling eq: incompatible types for comparison

Use --debug flag to render out invalid YAML
```

### 26. `small-hack/argocd-apps`

- Chart: `cloned\small-hack__argocd-apps\nextcloud\storage\pvc`
- Source: `template`
- Command: `helm template test cloned\small-hack__argocd-apps\nextcloud\storage\pvc`

```text
Error: template: nextcloud-pvcs/templates/nextcloud_files_pvc.yaml:1:7: executing "nextcloud-pvcs/templates/nextcloud_files_pvc.yaml" at <eq .Values.files_pvc.enabled "true">: error calling eq: incompatible types for comparison

Use --debug flag to render out invalid YAML
```

### 27. `small-hack/argocd-apps`

- Chart: `cloned\small-hack__argocd-apps\peertube\storage\pvc`
- Source: `template`
- Command: `helm template test cloned\small-hack__argocd-apps\peertube\storage\pvc`

```text
Error: template: peertube-pvc/templates/peertube_pvc.yaml:1:7: executing "peertube-pvc/templates/peertube_pvc.yaml" at <eq .Values.pvc.enabled "true">: error calling eq: incompatible types for comparison

Use --debug flag to render out invalid YAML
```

## `dependency.missing_subchart` (23)

### 1. `ConductionNL/procces-engine`

- Chart: `cloned\ConductionNL__procces-engine\api\helm`
- Source: `dependency`

```text
manager.go:122: warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash...
Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies
```

### 2. `manics/zero-to-jupyterhub-k8s-chart-expanded`

- Chart: `cloned\manics__zero-to-jupyterhub-k8s-chart-expanded\0.8.1`
- Source: `dependency`

```text
manager.go:122: warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash...
Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies
```

### 3. `ConductionNL/commonground-ui`

- Chart: `cloned\ConductionNL__commonground-ui\api\helm`
- Source: `dependency`

```text
manager.go:122: warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash...
Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies
```

### 4. `ConductionNL/zaak-registratie-component`

- Chart: `cloned\ConductionNL__zaak-registratie-component\api\helm`
- Source: `dependency`

```text
manager.go:122: warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash...
Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies
```

### 5. `ConductionNL/digispoof-old`

- Chart: `cloned\ConductionNL__digispoof-old\api\helm`
- Source: `dependency`

```text
manager.go:122: warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash...
Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies
```

### 6. `ConductionNL/proto-application-commonground`

- Chart: `cloned\ConductionNL__proto-application-commonground\api\helm`
- Source: `dependency`

```text
manager.go:122: warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash...
Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies
```

### 7. `grafana/helm-charts`

- Chart: `cloned\grafana__helm-charts\charts\enterprise-metrics`
- Source: `dependency`

```text
Getting updates for unmanaged Helm repositories...
...Successfully got an update from the "https://helm.min.io/" chart repository
...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Error: minio chart not found in repo https://helm.min.io/
```

### 8. `cloudnativeapp/charts`

- Chart: `cloned\cloudnativeapp__charts\curated\airflow`
- Source: `dependency`

```text
manager.go:122: warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash...
Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies
```

### 9. `marcus-sa/helm-charts`

- Chart: `cloned\marcus-sa__helm-charts\repositories\banzaicloud-stable\anchore-engine\0.13.0\chart`
- Source: `dependency`

```text
manager.go:122: warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash...
Error: no repository definition for https://kubernetes-charts.storage.googleapis.com, https://kubernetes-charts.storage.googleapis.com, https://kubernetes-charts.storage.googleapis.com. Please add the missing repos via 'helm repo add'
```

### 10. `ibuildthecloud/rancher-charts`

- Chart: `cloned\ibuildthecloud__rancher-charts\charts\anchore-engine\0.1.0`
- Source: `dependency`

```text
manager.go:122: warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash...
Error: no repository definition for https://kubernetes-charts.storage.googleapis.com. Please add the missing repos via 'helm repo add'
```

### 11. `feliux/containers-lab`

- Chart: `cloned\feliux__containers-lab\k8s\sysdig`
- Source: `dependency`

```text
Error: directory cloned\feliux__containers-lab\k8s\admission-controller not found
```

### 12. `aemneina/testcatalog`

- Chart: `cloned\aemneina__testcatalog\incubator\gogs`
- Source: `dependency`

```text
manager.go:122: warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash...
Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add'
```

### 13. `ValeriiVasianovych/argocd-kubernetes-infrastructure`

- Chart: `cloned\ValeriiVasianovych__argocd-kubernetes-infrastructure\icdev-01\helmCharts\prometheus-stack`
- Source: `dependency`

```text
manager.go:122: warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash...
Error: the lock file (requirements.lock) is out of sync with the dependencies file (requirements.yaml). Please update the dependencies
```

### 14. `cyberrangecz/devops-helm`

- Chart: `cloned\cyberrangecz__devops-helm\helm\crczp-head`
- Source: `dependency`

```text
Error: directory cloned\cyberrangecz__devops-helm\helm\charts\common not found
```

### 15. `joostvdg/cmg`

- Chart: `cloned\joostvdg__cmg\charts\preview`
- Source: `dependency`

```text
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "http://chartmuseum.jenkins-x.io" chart repository:
	failed to fetch http://chartmuseum.jenkins-x.io/index.yaml : 404 Not Found
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Error: no cached repository for helm-manager-d8c90e93d45e753db0fa4ec457d893e08cd1e0827b77896ee636b714706cbbb2 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-d8c90e93d45e753db0fa4ec457d893e08cd1e0827b77896ee636b714706cbbb2-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### 16. `cb-kubecd/bdd-gh-1581604741`

- Chart: `cloned\cb-kubecd__bdd-gh-1581604741\charts\preview`
- Source: `dependency`

```text
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "http://chartmuseum.jenkins-x.io" chart repository:
	failed to fetch http://chartmuseum.jenkins-x.io/index.yaml : 404 Not Found
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Error: no cached repository for helm-manager-d8c90e93d45e753db0fa4ec457d893e08cd1e0827b77896ee636b714706cbbb2 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-d8c90e93d45e753db0fa4ec457d893e08cd1e0827b77896ee636b714706cbbb2-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### 17. `Ammar-Taimoori/speechmatics-helm-chart`

- Chart: `cloned\Ammar-Taimoori__speechmatics-helm-chart\1.2.0\sm-realtime`
- Source: `dependency`

```text
Error: directory cloned\Ammar-Taimoori__speechmatics-helm-chart\1.2.0\sm-proxy not found
```

### 18. `wesparish/rancher-catalog`

- Chart: `cloned\wesparish__rancher-catalog\charts\ark-wes`
- Source: `dependency`

```text
Error: directory cloned\ark-server-charts\charts\ark-cluster not found
```

### 19. `cb-kubecd/bdd-nh-import-1586949778`

- Chart: `cloned\cb-kubecd__bdd-nh-import-1586949778\charts\preview`
- Source: `dependency`

```text
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "http://chartmuseum.jenkins-x.io" chart repository:
	failed to fetch http://chartmuseum.jenkins-x.io/index.yaml : 404 Not Found
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Error: no cached repository for helm-manager-d8c90e93d45e753db0fa4ec457d893e08cd1e0827b77896ee636b714706cbbb2 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-d8c90e93d45e753db0fa4ec457d893e08cd1e0827b77896ee636b714706cbbb2-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### 20. `ArvigEnterprises/integration-tests-bluegreen`

- Chart: `cloned\ArvigEnterprises__integration-tests-bluegreen\.helm\charts\arvigtestcluster334d\canary`
- Source: `dependency`

```text
Error: directory cloned\ArvigEnterprises__integration-tests-bluegreen\.helm\charts\arvigtestcluster334d\laravel-library not found
```

### 21. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\airflow`
- Source: `dependency`

```text
manager.go:122: warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash...
Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/, https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add'
```

### 22. `leandrofrs/spring_demo`

- Chart: `cloned\leandrofrs__spring_demo\charts\preview`
- Source: `dependency`

```text
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "http://chartmuseum.jenkins-x.io" chart repository:
	failed to fetch http://chartmuseum.jenkins-x.io/index.yaml : 404 Not Found
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Error: no cached repository for helm-manager-d8c90e93d45e753db0fa4ec457d893e08cd1e0827b77896ee636b714706cbbb2 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-d8c90e93d45e753db0fa4ec457d893e08cd1e0827b77896ee636b714706cbbb2-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### 23. `susnata1981/spring-demo`

- Chart: `cloned\susnata1981__spring-demo\charts\preview`
- Source: `dependency`

```text
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "http://chartmuseum.jenkins-x.io" chart repository:
	failed to fetch http://chartmuseum.jenkins-x.io/index.yaml : 404 Not Found
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Error: no cached repository for helm-manager-d8c90e93d45e753db0fa4ec457d893e08cd1e0827b77896ee636b714706cbbb2 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-d8c90e93d45e753db0fa4ec457d893e08cd1e0827b77896ee636b714706cbbb2-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

## `dependency.rate_limit` (18)

### 1. `michaelw/common.itsumi`

- Chart: `cloned\michaelw__common.itsumi`
- Source: `dependency`

```text
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 1 charts
Downloading common from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/common: failed to resolve registry-1.docker.io/bitnamicharts/common:2.31.3: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.31.3": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/common: failed to resolve registry-1.docker.io/bitnamicharts/common:2.31.3: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.31.3": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### 2. `waTeim/mcp-base`

- Chart: `cloned\waTeim__mcp-base\chart`
- Source: `dependency`

```text
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 1 charts
Downloading redis from repo oci://registry-1.docker.io/cloudpirates
Save error occurred:  could not download oci://registry-1.docker.io/cloudpirates/redis: failed to resolve registry-1.docker.io/cloudpirates/redis:0.16.4: GET "https://registry-1.docker.io/v2/cloudpirates/redis/manifests/0.16.4": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/cloudpirates/redis: failed to resolve registry-1.docker.io/cloudpirates/redis:0.16.4: GET "https://registry-1.docker.io/v2/cloudpirates/redis/manifests/0.16.4": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### 3. `datawizz/Firestream`

- Chart: `cloned\datawizz__Firestream\src\charts\bitnami\airflow`
- Source: `dependency`

```text
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 3 charts
Downloading redis from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:21.2.5: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/21.2.5": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:21.2.5: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/21.2.5": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### 4. `bitnami/charts`

- Chart: `cloned\bitnami__charts\bitnami\airflow`
- Source: `dependency`

```text
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 3 charts
Downloading redis from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:22.0.4: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/22.0.4": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:22.0.4: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/22.0.4": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### 5. `musqat/lectures`

- Chart: `cloned\musqat__lectures\MicroService\section_16\helm\grafana`
- Source: `dependency`

```text
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 1 charts
Downloading common from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/common: failed to resolve registry-1.docker.io/bitnamicharts/common:2.31.0: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.31.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/common: failed to resolve registry-1.docker.io/bitnamicharts/common:2.31.0: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.31.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### 6. `general-rj45/mini-ml-stand-GRJ45`

- Chart: `cloned\general-rj45__mini-ml-stand-GRJ45\airflow-14.2.5\airflow`
- Source: `dependency`

```text
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 3 charts
Downloading redis from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:17.11.3: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/17.11.3": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:17.11.3: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/17.11.3": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### 7. `czerwonk/hakanai`

- Chart: `cloned\czerwonk__hakanai\helm\hakanai`
- Source: `dependency`

```text
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 1 charts
Downloading redis from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:20.5.0: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/20.5.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:20.5.0: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/20.5.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### 8. `quantumreasoning/kubeapps`

- Chart: `cloned\quantumreasoning__kubeapps\chart\kubeapps`
- Source: `dependency`

```text
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 3 charts
Downloading redis from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:20.2.0: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/20.2.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:20.2.0: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/20.2.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### 9. `t9k/apps`

- Chart: `cloned\t9k__apps\user-console\chatnio\chart`
- Source: `dependency`

```text
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 2 charts
Downloading mysql from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/mysql: failed to resolve registry-1.docker.io/bitnamicharts/mysql:9.18.2: GET "https://registry-1.docker.io/v2/bitnamicharts/mysql/manifests/9.18.2": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/mysql: failed to resolve registry-1.docker.io/bitnamicharts/mysql:9.18.2: GET "https://registry-1.docker.io/v2/bitnamicharts/mysql/manifests/9.18.2": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### 10. `DOME-Marketplace/dome-gitops`

- Chart: `cloned\DOME-Marketplace__dome-gitops\ionos\marketplace\kafka`
- Source: `dependency`

```text
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 1 charts
Downloading kafka from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/kafka: failed to resolve registry-1.docker.io/bitnamicharts/kafka:26.0.0: GET "https://registry-1.docker.io/v2/bitnamicharts/kafka/manifests/26.0.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/kafka: failed to resolve registry-1.docker.io/bitnamicharts/kafka:26.0.0: GET "https://registry-1.docker.io/v2/bitnamicharts/kafka/manifests/26.0.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### 11. `keonhoban/mlops-infra-gitops`

- Chart: `cloned\keonhoban__mlops-infra-gitops\baseline\charts\minio-wrapper`
- Source: `dependency`

```text
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 1 charts
Downloading minio from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/minio: failed to resolve registry-1.docker.io/bitnamicharts/minio:17.0.21: GET "https://registry-1.docker.io/v2/bitnamicharts/minio/manifests/17.0.21": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/minio: failed to resolve registry-1.docker.io/bitnamicharts/minio:17.0.21: GET "https://registry-1.docker.io/v2/bitnamicharts/minio/manifests/17.0.21": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### 12. `IvanciniGT/formacionHelm`

- Chart: `cloned\IvanciniGT__formacionHelm\charts\mariadb`
- Source: `dependency`

```text
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 1 charts
Downloading common from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/common: failed to resolve registry-1.docker.io/bitnamicharts/common:2.6.0: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.6.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/common: failed to resolve registry-1.docker.io/bitnamicharts/common:2.6.0: GET "https://registry-1.docker.io/v2/bitnamicharts/common/manifests/2.6.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### 13. `LMCache/LMBench`

- Chart: `cloned\LMCache__LMBench\2-serving-engines\dynamo\dynamo\deploy\cloud\helm\platform`
- Source: `dependency`

```text
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

### 14. `shakedaskayo/kubiya-demo`

- Chart: `cloned\shakedaskayo__kubiya-demo\helm\bitnami-charts\bitnami\airflow`
- Source: `dependency`

```text
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 3 charts
Downloading redis from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:18.11.0: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/18.11.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:18.11.0: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/18.11.0": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### 15. `moophlo/helm-charts`

- Chart: `cloned\moophlo__helm-charts\bitnami\airflow`
- Source: `dependency`

```text
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 3 charts
Downloading redis from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:22.0.4: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/22.0.4": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:22.0.4: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/22.0.4": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### 16. `aicanhelp/ai-cloud`

- Chart: `cloned\aicanhelp__ai-cloud\bitnami\charts\bitnami-202503\airflow`
- Source: `dependency`

```text
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 3 charts
Downloading redis from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:20.7.1: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/20.7.1": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:20.7.1: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/20.7.1": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### 17. `itboon/bitnami-charts`

- Chart: `cloned\itboon__bitnami-charts\charts-bitnami\airflow`
- Source: `dependency`

```text
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 3 charts
Downloading redis from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:22.0.4: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/22.0.4": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:22.0.4: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/22.0.4": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

### 18. `tdunlap607/kubeapps-fork`

- Chart: `cloned\tdunlap607__kubeapps-fork\chart\kubeapps`
- Source: `dependency`

```text
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 3 charts
Downloading redis from repo oci://registry-1.docker.io/bitnamicharts
Save error occurred:  could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:20.6.1: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/20.6.1": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
Error: could not download oci://registry-1.docker.io/bitnamicharts/redis: failed to resolve registry-1.docker.io/bitnamicharts/redis:20.6.1: GET "https://registry-1.docker.io/v2/bitnamicharts/redis/manifests/20.6.1": response status code 429: toomanyrequests: You have reached your unauthenticated pull rate limit. https://www.docker.com/increase-rate-limit
```

## `template.invalid_value` (16)

### 1. `manics/zero-to-jupyterhub-k8s-chart-expanded`

- Chart: `cloned\manics__zero-to-jupyterhub-k8s-chart-expanded\0.10.0`
- Source: `template`
- Command: `helm template test cloned\manics__zero-to-jupyterhub-k8s-chart-expanded\0.10.0`

```text
Error: execution error at (jupyterhub/templates/proxy/deployment.yaml:28:32): Proxy token must be a 32 byte random string generated with `openssl rand -hex 32`!

Use --debug flag to render out invalid YAML
```

### 2. `manics/zero-to-jupyterhub-k8s-chart-expanded`

- Chart: `cloned\manics__zero-to-jupyterhub-k8s-chart-expanded\0.10.0-beta.1`
- Source: `template`
- Command: `helm template test cloned\manics__zero-to-jupyterhub-k8s-chart-expanded\0.10.0-beta.1`

```text
Error: execution error at (jupyterhub/templates/proxy/deployment.yaml:28:32): Proxy token must be a 32 byte random string generated with `openssl rand -hex 32`!

Use --debug flag to render out invalid YAML
```

### 3. `manics/zero-to-jupyterhub-k8s-chart-expanded`

- Chart: `cloned\manics__zero-to-jupyterhub-k8s-chart-expanded\0.10.1`
- Source: `template`
- Command: `helm template test cloned\manics__zero-to-jupyterhub-k8s-chart-expanded\0.10.1`

```text
Error: execution error at (jupyterhub/templates/proxy/deployment.yaml:28:32): Proxy token must be a 32 byte random string generated with `openssl rand -hex 32`!

Use --debug flag to render out invalid YAML
```

### 4. `manics/zero-to-jupyterhub-k8s-chart-expanded`

- Chart: `cloned\manics__zero-to-jupyterhub-k8s-chart-expanded\0.10.2`
- Source: `template`
- Command: `helm template test cloned\manics__zero-to-jupyterhub-k8s-chart-expanded\0.10.2`

```text
Error: execution error at (jupyterhub/templates/proxy/deployment.yaml:28:32): Proxy token must be a 32 byte random string generated with `openssl rand -hex 32`!

Use --debug flag to render out invalid YAML
```

### 5. `manics/zero-to-jupyterhub-k8s-chart-expanded`

- Chart: `cloned\manics__zero-to-jupyterhub-k8s-chart-expanded\0.10.3`
- Source: `template`
- Command: `helm template test cloned\manics__zero-to-jupyterhub-k8s-chart-expanded\0.10.3`

```text
Error: execution error at (jupyterhub/templates/proxy/deployment.yaml:28:32): Proxy token must be a 32 byte random string generated with `openssl rand -hex 32`!

Use --debug flag to render out invalid YAML
```

### 6. `manics/zero-to-jupyterhub-k8s-chart-expanded`

- Chart: `cloned\manics__zero-to-jupyterhub-k8s-chart-expanded\0.10.4`
- Source: `template`
- Command: `helm template test cloned\manics__zero-to-jupyterhub-k8s-chart-expanded\0.10.4`

```text
Error: execution error at (jupyterhub/templates/proxy/deployment.yaml:28:32): Proxy token must be a 32 byte random string generated with `openssl rand -hex 32`!

Use --debug flag to render out invalid YAML
```

### 7. `manics/zero-to-jupyterhub-k8s-chart-expanded`

- Chart: `cloned\manics__zero-to-jupyterhub-k8s-chart-expanded\0.10.5`
- Source: `template`
- Command: `helm template test cloned\manics__zero-to-jupyterhub-k8s-chart-expanded\0.10.5`

```text
Error: execution error at (jupyterhub/templates/proxy/deployment.yaml:28:32): Proxy token must be a 32 byte random string generated with `openssl rand -hex 32`!

Use --debug flag to render out invalid YAML
```

### 8. `manics/zero-to-jupyterhub-k8s-chart-expanded`

- Chart: `cloned\manics__zero-to-jupyterhub-k8s-chart-expanded\0.10.6`
- Source: `template`
- Command: `helm template test cloned\manics__zero-to-jupyterhub-k8s-chart-expanded\0.10.6`

```text
Error: execution error at (jupyterhub/templates/proxy/deployment.yaml:28:32): Proxy token must be a 32 byte random string generated with `openssl rand -hex 32`!

Use --debug flag to render out invalid YAML
```

### 9. `manics/zero-to-jupyterhub-k8s-chart-expanded`

- Chart: `cloned\manics__zero-to-jupyterhub-k8s-chart-expanded\0.11.0`
- Source: `template`
- Command: `helm template test cloned\manics__zero-to-jupyterhub-k8s-chart-expanded\0.11.0`

```text
Error: execution error at (jupyterhub/templates/proxy/deployment.yaml:28:32): Proxy token must be a 32 byte random string generated with `openssl rand -hex 32`!

Use --debug flag to render out invalid YAML
```

### 10. `manics/zero-to-jupyterhub-k8s-chart-expanded`

- Chart: `cloned\manics__zero-to-jupyterhub-k8s-chart-expanded\0.11.1`
- Source: `template`
- Command: `helm template test cloned\manics__zero-to-jupyterhub-k8s-chart-expanded\0.11.1`

```text
Error: execution error at (jupyterhub/templates/proxy/deployment.yaml:28:32): Proxy token must be a 32 byte random string generated with `openssl rand -hex 32`!

Use --debug flag to render out invalid YAML
```

### 11. `manics/zero-to-jupyterhub-k8s-chart-expanded`

- Chart: `cloned\manics__zero-to-jupyterhub-k8s-chart-expanded\0.5.0`
- Source: `template`
- Command: `helm template test cloned\manics__zero-to-jupyterhub-k8s-chart-expanded\0.5.0`

```text
Error: execution error at (jupyterhub/templates/proxy/deployment.yaml:16:32): proxy.secretToken must be a 32 byte random string generated with `openssl rand -hex 32`

Use --debug flag to render out invalid YAML
```

### 12. `manics/zero-to-jupyterhub-k8s-chart-expanded`

- Chart: `cloned\manics__zero-to-jupyterhub-k8s-chart-expanded\0.7.0`
- Source: `template`
- Command: `helm template test cloned\manics__zero-to-jupyterhub-k8s-chart-expanded\0.7.0`

```text
Error: execution error at (jupyterhub/templates/proxy/deployment.yaml:26:32): Proxy token must be a 32 byte random string generated with `openssl rand -hex 32`!

Use --debug flag to render out invalid YAML
```

### 13. `manics/zero-to-jupyterhub-k8s-chart-expanded`

- Chart: `cloned\manics__zero-to-jupyterhub-k8s-chart-expanded\0.7.0-beta.1`
- Source: `template`
- Command: `helm template test cloned\manics__zero-to-jupyterhub-k8s-chart-expanded\0.7.0-beta.1`

```text
Error: execution error at (jupyterhub/templates/proxy/deployment.yaml:26:32): Proxy token must be a 32 byte random string generated with `openssl rand -hex 32`!

Use --debug flag to render out invalid YAML
```

### 14. `manics/zero-to-jupyterhub-k8s-chart-expanded`

- Chart: `cloned\manics__zero-to-jupyterhub-k8s-chart-expanded\0.7.0-beta.2`
- Source: `template`
- Command: `helm template test cloned\manics__zero-to-jupyterhub-k8s-chart-expanded\0.7.0-beta.2`

```text
Error: execution error at (jupyterhub/templates/proxy/deployment.yaml:26:32): Proxy token must be a 32 byte random string generated with `openssl rand -hex 32`!

Use --debug flag to render out invalid YAML
```

### 15. `manics/zero-to-jupyterhub-k8s-chart-expanded`

- Chart: `cloned\manics__zero-to-jupyterhub-k8s-chart-expanded\0.8.0`
- Source: `template`
- Command: `helm template test cloned\manics__zero-to-jupyterhub-k8s-chart-expanded\0.8.0`

```text
Error: execution error at (jupyterhub/templates/proxy/deployment.yaml:26:32): Proxy token must be a 32 byte random string generated with `openssl rand -hex 32`!

Use --debug flag to render out invalid YAML
```

### 16. `manics/zero-to-jupyterhub-k8s-chart-expanded`

- Chart: `cloned\manics__zero-to-jupyterhub-k8s-chart-expanded\0.8.0-beta.1`
- Source: `template`
- Command: `helm template test cloned\manics__zero-to-jupyterhub-k8s-chart-expanded\0.8.0-beta.1`

```text
Error: execution error at (jupyterhub/templates/proxy/deployment.yaml:26:32): Proxy token must be a 32 byte random string generated with `openssl rand -hex 32`!

Use --debug flag to render out invalid YAML
```

## `template.type_mismatch` (16)

### 1. `25th-Night/lion-k8s`

- Chart: `cloned\25th-Night__lion-k8s\lionp`
- Source: `template`
- Command: `helm template test cloned\25th-Night__lion-k8s\lionp`

```text
Error: template: lion/templates/django-secret.yaml:7:66: executing "lion/templates/django-secret.yaml" at <b64enc>: invalid value; expected string

Use --debug flag to render out invalid YAML
```

### 2. `redhat-cop/global-load-balancer-operator`

- Chart: `cloned\redhat-cop__global-load-balancer-operator\docs\scripts\acm-aws-cluster`
- Source: `template`
- Command: `helm template test cloned\redhat-cop__global-load-balancer-operator\docs\scripts\acm-aws-cluster`

```text
Error: template: acm-aws-cluster/templates/ssh-private-key-secret.yaml:7:19: executing "acm-aws-cluster/templates/ssh-private-key-secret.yaml" at <.Values.sshKey>: wrong type for value; expected string; got interface {}

Use --debug flag to render out invalid YAML
```

### 3. `otus-kuber-2020-07/israodin_platform`

- Chart: `cloned\otus-kuber-2020-07__israodin_platform\kubernetes-templating\frontend`
- Source: `template`
- Command: `helm template test cloned\otus-kuber-2020-07__israodin_platform\kubernetes-templating\frontend`

```text
Error: template: frontend/templates/secret.yaml:7:40: executing "frontend/templates/secret.yaml" at <b64enc>: invalid value; expected string

Use --debug flag to render out invalid YAML
```

### 4. `ArseniiT/spring-microservices`

- Chart: `cloned\ArseniiT__spring-microservices\helm\api-gateway-ingress`
- Source: `template`
- Command: `helm template test cloned\ArseniiT__spring-microservices\helm\api-gateway-ingress`

```text
Error: template: api-gateway-ingress/templates/ingress.yaml:18:54: executing "api-gateway-ingress/templates/ingress.yaml" at <"-">: invalid value; expected string

Use --debug flag to render out invalid YAML
```

### 5. `ArseniiT/spring-microservices`

- Chart: `cloned\ArseniiT__spring-microservices\helm\certificate`
- Source: `template`
- Command: `helm template test cloned\ArseniiT__spring-microservices\helm\certificate`

```text
Error: template: certificate/templates/certificate.yaml:4:44: executing "certificate/templates/certificate.yaml" at <"-">: invalid value; expected string

Use --debug flag to render out invalid YAML
```

### 6. `ArseniiT/spring-microservices`

- Chart: `cloned\ArseniiT__spring-microservices\helm\certificate`
- Source: `template`
- Values files: `cloned\ArseniiT__spring-microservices\helm\certificate\values.example.yaml`
- Command: `helm template test cloned\ArseniiT__spring-microservices\helm\certificate -f cloned\ArseniiT__spring-microservices\helm\certificate\values.example.yaml`

```text
Error: template: certificate/templates/certificate.yaml:4:44: executing "certificate/templates/certificate.yaml" at <"-">: invalid value; expected string

Use --debug flag to render out invalid YAML
```

### 7. `axlgrosse/42c-Protect`

- Chart: `cloned\axlgrosse__42c-Protect\deployment\Pixi-helm\42c-firewall`
- Source: `template`
- Command: `helm template test cloned\axlgrosse__42c-Protect\deployment\Pixi-helm\42c-firewall`

```text
Error: template: firewall/templates/protection-token.yaml:9:62: executing "firewall/templates/protection-token.yaml" at <b64enc>: invalid value; expected string

Use --debug flag to render out invalid YAML
```

### 8. `axlgrosse/42c-Protect`

- Chart: `cloned\axlgrosse__42c-Protect\deployment\helm-artifacts\capture-plugin`
- Source: `template`
- Values files: `cloned\axlgrosse__42c-Protect\deployment\helm-artifacts\capture-plugin\azure-values.yaml`
- Command: `helm template test cloned\axlgrosse__42c-Protect\deployment\helm-artifacts\capture-plugin -f cloned\axlgrosse__42c-Protect\deployment\helm-artifacts\capture-plugin\azure-values.yaml`

```text
Error: template: horus-agent/templates/access-token.yaml:9:48: executing "horus-agent/templates/access-token.yaml" at <b64enc>: invalid value; expected string

Use --debug flag to render out invalid YAML
```

### 9. `danskernesdigitalebibliotek/dpl-platform`

- Chart: `cloned\danskernesdigitalebibliotek__dpl-platform\gitops\charts\proxy-pass-service`
- Source: `template`
- Command: `helm template test cloned\danskernesdigitalebibliotek__dpl-platform\gitops\charts\proxy-pass-service`

```text
Error: template: proxy-pass-service/templates/ingress-go.yaml:4:38: executing "proxy-pass-service/templates/ingress-go.yaml" at <.Values.host>: wrong type for value; expected string; got interface {}

Use --debug flag to render out invalid YAML
```

### 10. `danskernesdigitalebibliotek/dpl-platform`

- Chart: `cloned\danskernesdigitalebibliotek__dpl-platform\gitops\dplplat01\sites\albertslund-moduletest`
- Source: `template`
- Command: `helm template test cloned\danskernesdigitalebibliotek__dpl-platform\gitops\dplplat01\sites\albertslund-moduletest`

```text
Error: template: albertslund-moduletest/charts/proxy-pass-service/templates/ingress-go.yaml:4:38: executing "albertslund-moduletest/charts/proxy-pass-service/templates/ingress-go.yaml" at <.Values.host>: wrong type for value; expected string; got interface {}

Use --debug flag to render out invalid YAML
```

### 11. `danskernesdigitalebibliotek/dpl-platform`

- Chart: `cloned\danskernesdigitalebibliotek__dpl-platform\gitops\dplplat01\sites\faxe-moduletest`
- Source: `template`
- Command: `helm template test cloned\danskernesdigitalebibliotek__dpl-platform\gitops\dplplat01\sites\faxe-moduletest`

```text
Error: template: faxe-moduletest/charts/proxy-pass-service/templates/ingress-go.yaml:4:38: executing "faxe-moduletest/charts/proxy-pass-service/templates/ingress-go.yaml" at <.Values.host>: wrong type for value; expected string; got interface {}

Use --debug flag to render out invalid YAML
```

### 12. `danskernesdigitalebibliotek/dpl-platform`

- Chart: `cloned\danskernesdigitalebibliotek__dpl-platform\gitops\dplplat01\sites\kobenhavn-main`
- Source: `template`
- Command: `helm template test cloned\danskernesdigitalebibliotek__dpl-platform\gitops\dplplat01\sites\kobenhavn-main`

```text
Error: template: kobenhavn-main/charts/proxy-pass-service/templates/ingress-go.yaml:4:38: executing "kobenhavn-main/charts/proxy-pass-service/templates/ingress-go.yaml" at <.Values.host>: wrong type for value; expected string; got interface {}

Use --debug flag to render out invalid YAML
```

### 13. `danskernesdigitalebibliotek/dpl-platform`

- Chart: `cloned\danskernesdigitalebibliotek__dpl-platform\gitops\dplplat01\sites\kobenhavn-moduletest`
- Source: `template`
- Command: `helm template test cloned\danskernesdigitalebibliotek__dpl-platform\gitops\dplplat01\sites\kobenhavn-moduletest`

```text
Error: template: kobenhavn-moduletest/charts/proxy-pass-service/templates/ingress-go.yaml:4:38: executing "kobenhavn-moduletest/charts/proxy-pass-service/templates/ingress-go.yaml" at <.Values.host>: wrong type for value; expected string; got interface {}

Use --debug flag to render out invalid YAML
```

### 14. `otus-kuber-2020-04/israodin_platform`

- Chart: `cloned\otus-kuber-2020-04__israodin_platform\kubernetes-templating\frontend`
- Source: `template`
- Command: `helm template test cloned\otus-kuber-2020-04__israodin_platform\kubernetes-templating\frontend`

```text
Error: template: frontend/templates/secret.yaml:7:40: executing "frontend/templates/secret.yaml" at <b64enc>: invalid value; expected string

Use --debug flag to render out invalid YAML
```

### 15. `42c-arhayem/Demo`

- Chart: `cloned\42c-arhayem__Demo\artifacts\42c-firewall`
- Source: `template`
- Command: `helm template test cloned\42c-arhayem__Demo\artifacts\42c-firewall`

```text
Error: template: firewall/templates/protection-token.yaml:9:62: executing "firewall/templates/protection-token.yaml" at <b64enc>: invalid value; expected string

Use --debug flag to render out invalid YAML
```

### 16. `42c-arhayem/Pixi`

- Chart: `cloned\42c-arhayem__Pixi\artifacts\42c-firewall`
- Source: `template`
- Command: `helm template test cloned\42c-arhayem__Pixi\artifacts\42c-firewall`

```text
Error: template: firewall/templates/protection-token.yaml:9:62: executing "firewall/templates/protection-token.yaml" at <b64enc>: invalid value; expected string

Use --debug flag to render out invalid YAML
```

## `dependency.network_dns` (14)

### 1. `noygal/helm`

- Chart: `cloned\noygal__helm\charts\example-dev-tools`
- Source: `dependency`

```text
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "https://noygal.github.com/helm" chart repository:
	Get "https://noygal.github.com/helm/index.yaml": dial tcp: lookup noygal.github.com: no such host
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Error: no cached repository for helm-manager-19cefd96797d2c0b3ed74271dcee4ae7ad7d02cb4d4e47a98431938d24b11530 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-19cefd96797d2c0b3ed74271dcee4ae7ad7d02cb4d4e47a98431938d24b11530-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### 2. `hey101/scale-catalog`

- Chart: `cloned\hey101__scale-catalog\incubator\archivebox\0.7.2`
- Source: `dependency`

```text
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 1 charts
Downloading common from repo oci://tccr.io/truecharts
Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to resolve tccr.io/truecharts/common:17.5.1: Get "https://tccr.io/v2/truecharts/common/manifests/17.5.1": dial tcp: lookup tccr.io: no such host
Error: could not download oci://tccr.io/truecharts/common: failed to resolve tccr.io/truecharts/common:17.5.1: Get "https://tccr.io/v2/truecharts/common/manifests/17.5.1": dial tcp: lookup tccr.io: no such host
```

### 3. `unity-systems/spring-demo`

- Chart: `cloned\unity-systems__spring-demo\charts\preview`
- Source: `dependency`

```text
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

### 4. `Mark4562/Truenas-Scale`

- Chart: `cloned\Mark4562__Truenas-Scale\premium\authelia\23.13.13`
- Source: `dependency`

```text
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 2 charts
Downloading common from repo oci://tccr.io/truecharts
Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to resolve tccr.io/truecharts/common:20.3.11: Get "https://tccr.io/v2/truecharts/common/manifests/20.3.11": dial tcp: lookup tccr.io: no such host
Error: could not download oci://tccr.io/truecharts/common: failed to resolve tccr.io/truecharts/common:20.3.11: Get "https://tccr.io/v2/truecharts/common/manifests/20.3.11": dial tcp: lookup tccr.io: no such host
```

### 5. `sapcc/helm-charts`

- Chart: `cloned\sapcc__helm-charts\common\inventory-updater`
- Source: `dependency`

```text
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 1 charts
Downloading owner-info from repo oci://keppel.eu-de-1.cloud.sap/ccloud-helm
Save error occurred:  could not download oci://keppel.eu-de-1.cloud.sap/ccloud-helm/owner-info: failed to resolve keppel.eu-de-1.cloud.sap/ccloud-helm/owner-info:0.2.0: Get "https://keppel.eu-de-1.cloud.sap/v2/ccloud-helm/owner-info/manifests/0.2.0": dial tcp: lookup keppel.eu-de-1.cloud.sap: getaddrinfow: Este é geralmente um erro temporário durante a resolução de nomes de anfitrião e significa que o servidor local não recebeu uma resposta de um servidor autoritário.
Error: could not download oci://keppel.eu-de-1.cloud.sap/ccloud-helm/owner-info: failed to resolve keppel.eu-de-1.cloud.sap/ccloud-helm/owner-info:0.2.0: Get "https://keppel.eu-de-1.cloud.sap/v2/ccloud-helm/owner-info/manifests/0.2.0": dial tcp: lookup keppel.eu-de-1.cloud.sap: getaddrinfow: Este é geralmente um erro temporário durante a resolução de nomes de anfitrião e significa que o servidor local não recebeu uma resposta de um servidor autoritário.
```

### 6. `jstrachan/demo16`

- Chart: `cloned\jstrachan__demo16\charts\preview`
- Source: `dependency`

```text
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

### 7. `KevMCarp/truecharts-catalog-fork`

- Chart: `cloned\KevMCarp__truecharts-catalog-fork\dependency\clickhouse\5.0.54`
- Source: `dependency`

```text
Getting updates for unmanaged Helm repositories...
...Unable to get an update from the "https://library-charts.truecharts.org" chart repository:
	Get "https://library-charts.truecharts.org/index.yaml": dial tcp: lookup library-charts.truecharts.org: no such host
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Error: no cached repository for helm-manager-024b189b59f6c6ccf0de6e5148db1578caf551c511f4eb220ece14cef00f80e0 found. (try 'helm repo update'): open C:\Users\miabs\AppData\Local\Temp\helm\repository\helm-manager-024b189b59f6c6ccf0de6e5148db1578caf551c511f4eb220ece14cef00f80e0-index.yaml: O sistema não conseguiu localizar o ficheiro especificado.
```

### 8. `ccojocar/spring-demo`

- Chart: `cloned\ccojocar__spring-demo\charts\preview`
- Source: `dependency`

```text
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

### 9. `introproventures/spring-demo`

- Chart: `cloned\introproventures__spring-demo\charts\preview`
- Source: `dependency`

```text
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

### 10. `Jaebytes/TrueCharts`

- Chart: `cloned\Jaebytes__TrueCharts\premium\traefik\26.10.19`
- Source: `dependency`

```text
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 1 charts
Downloading common from repo oci://tccr.io/truecharts
Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to resolve tccr.io/truecharts/common:20.3.12: Get "https://tccr.io/v2/truecharts/common/manifests/20.3.12": dial tcp: lookup tccr.io: no such host
Error: could not download oci://tccr.io/truecharts/common: failed to resolve tccr.io/truecharts/common:20.3.12: Get "https://tccr.io/v2/truecharts/common/manifests/20.3.12": dial tcp: lookup tccr.io: no such host
```

### 11. `PnsDev/Truechart-catalog-archive`

- Chart: `cloned\PnsDev__Truechart-catalog-archive\incubator\archivebox\0.7.2`
- Source: `dependency`

```text
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 1 charts
Downloading common from repo oci://tccr.io/truecharts
Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to resolve tccr.io/truecharts/common:17.2.30: Get "https://tccr.io/v2/truecharts/common/manifests/17.2.30": dial tcp: lookup tccr.io: no such host
Error: could not download oci://tccr.io/truecharts/common: failed to resolve tccr.io/truecharts/common:17.2.30: Get "https://tccr.io/v2/truecharts/common/manifests/17.2.30": dial tcp: lookup tccr.io: no such host
```

### 12. `petersandor/tc-archive`

- Chart: `cloned\petersandor__tc-archive\scale-catalog\incubator\archivebox\0.7.2`
- Source: `dependency`

```text
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 1 charts
Downloading common from repo oci://tccr.io/truecharts
Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to resolve tccr.io/truecharts/common:17.2.30: Get "https://tccr.io/v2/truecharts/common/manifests/17.2.30": dial tcp: lookup tccr.io: no such host
Error: could not download oci://tccr.io/truecharts/common: failed to resolve tccr.io/truecharts/common:17.2.30: Get "https://tccr.io/v2/truecharts/common/manifests/17.2.30": dial tcp: lookup tccr.io: no such host
```

### 13. `adstanley/archive`

- Chart: `cloned\adstanley__archive\scale-catalog\incubator\archivebox\0.7.2`
- Source: `dependency`

```text
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 1 charts
Downloading common from repo oci://tccr.io/truecharts
Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to resolve tccr.io/truecharts/common:17.2.30: Get "https://tccr.io/v2/truecharts/common/manifests/17.2.30": dial tcp: lookup tccr.io: no such host
Error: could not download oci://tccr.io/truecharts/common: failed to resolve tccr.io/truecharts/common:17.2.30: Get "https://tccr.io/v2/truecharts/common/manifests/17.2.30": dial tcp: lookup tccr.io: no such host
```

### 14. `adynetro/truecharts`

- Chart: `cloned\adynetro__truecharts\incubator\archivebox\0.7.2`
- Source: `dependency`

```text
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 1 charts
Downloading common from repo oci://tccr.io/truecharts
Save error occurred:  could not download oci://tccr.io/truecharts/common: failed to resolve tccr.io/truecharts/common:17.2.30: Get "https://tccr.io/v2/truecharts/common/manifests/17.2.30": dial tcp: lookup tccr.io: no such host
Error: could not download oci://tccr.io/truecharts/common: failed to resolve tccr.io/truecharts/common:17.2.30: Get "https://tccr.io/v2/truecharts/common/manifests/17.2.30": dial tcp: lookup tccr.io: no such host
```

## `template.values_schema_validation` (12)

### 1. `maheshlakkammanavar/InciCW26`

- Chart: `cloned\maheshlakkammanavar__InciCW26\chart\charts\content-deployment`
- Source: `template`
- Command: `helm template test cloned\maheshlakkammanavar__InciCW26\chart\charts\content-deployment`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
content-deployment:
- (root): bindings is required
- image: repository is required
```

### 2. `maheshlakkammanavar/InciCW26`

- Chart: `cloned\maheshlakkammanavar__InciCW26\chart\charts\service-instance`
- Source: `template`
- Command: `helm template test cloned\maheshlakkammanavar__InciCW26\chart\charts\service-instance`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
service-instance:
- (root): serviceOfferingName is required
- (root): servicePlanName is required
```

### 3. `maheshlakkammanavar/InciCW26`

- Chart: `cloned\maheshlakkammanavar__InciCW26\chart\charts\web-application`
- Source: `template`
- Command: `helm template test cloned\maheshlakkammanavar__InciCW26\chart\charts\web-application`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
web-application:
- (root): resources is required
- image: repository is required
```

### 4. `cdv18/helm-charts`

- Chart: `cloned\cdv18__helm-charts\umbrella`
- Source: `template`
- Values files: `cloned\cdv18__helm-charts\umbrella\values-prod.yaml`
- Command: `helm template test cloned\cdv18__helm-charts\umbrella -f cloned\cdv18__helm-charts\umbrella\values-prod.yaml`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
redis:
- architecture: architecture must be one of the following: "standalone", "replication"
```

### 5. `cdv18/helm-charts`

- Chart: `cloned\cdv18__helm-charts\umbrella`
- Source: `template`
- Values files: `cloned\cdv18__helm-charts\umbrella\values-dev.yaml`, `cloned\cdv18__helm-charts\umbrella\values-prod.yaml`
- Command: `helm template test cloned\cdv18__helm-charts\umbrella -f cloned\cdv18__helm-charts\umbrella\values-dev.yaml -f cloned\cdv18__helm-charts\umbrella\values-prod.yaml`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
redis:
- architecture: architecture must be one of the following: "standalone", "replication"
```

### 6. `cdv18/helm-charts`

- Chart: `cloned\cdv18__helm-charts\umbrella`
- Source: `template`
- Values files: `cloned\cdv18__helm-charts\umbrella\values-prod.yaml`, `cloned\cdv18__helm-charts\umbrella\values-staging.yaml`
- Command: `helm template test cloned\cdv18__helm-charts\umbrella -f cloned\cdv18__helm-charts\umbrella\values-prod.yaml -f cloned\cdv18__helm-charts\umbrella\values-staging.yaml`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
redis:
- architecture: architecture must be one of the following: "standalone", "replication"
```

### 7. `cdv18/helm-charts`

- Chart: `cloned\cdv18__helm-charts\umbrella`
- Source: `template`
- Values files: `cloned\cdv18__helm-charts\umbrella\values-dev.yaml`, `cloned\cdv18__helm-charts\umbrella\values-prod.yaml`, `cloned\cdv18__helm-charts\umbrella\values-staging.yaml`
- Command: `helm template test cloned\cdv18__helm-charts\umbrella -f cloned\cdv18__helm-charts\umbrella\values-dev.yaml -f cloned\cdv18__helm-charts\umbrella\values-prod.yaml -f cloned\cdv18__helm-charts\umbrella\values-staging.yaml`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
redis:
- architecture: architecture must be one of the following: "standalone", "replication"
```

### 8. `gregorwolf/cap-python`

- Chart: `cloned\gregorwolf__cap-python\chart\charts\content-deployment`
- Source: `template`
- Command: `helm template test cloned\gregorwolf__cap-python\chart\charts\content-deployment`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
content-deployment:
- (root): bindings is required
- image: repository is required
```

### 9. `gregorwolf/cap-python`

- Chart: `cloned\gregorwolf__cap-python\chart\charts\service-instance`
- Source: `template`
- Command: `helm template test cloned\gregorwolf__cap-python\chart\charts\service-instance`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
service-instance:
- (root): serviceOfferingName is required
- (root): servicePlanName is required
```

### 10. `gregorwolf/cap-python`

- Chart: `cloned\gregorwolf__cap-python\chart\charts\web-application`
- Source: `template`
- Command: `helm template test cloned\gregorwolf__cap-python\chart\charts\web-application`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
web-application:
- (root): resources is required
- image: repository is required
```

### 11. `brampurnot/sap-cap-mtx`

- Chart: `cloned\brampurnot__sap-cap-mtx\chart\charts\service-instance`
- Source: `template`
- Command: `helm template test cloned\brampurnot__sap-cap-mtx\chart\charts\service-instance`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
service-instance:
- (root): serviceOfferingName is required
- (root): servicePlanName is required
```

### 12. `brampurnot/sap-cap-mtx`

- Chart: `cloned\brampurnot__sap-cap-mtx\chart\charts\web-application`
- Source: `template`
- Command: `helm template test cloned\brampurnot__sap-cap-mtx\chart\charts\web-application`

```text
Error: values don't meet the specifications of the schema(s) in the following chart(s):
web-application:
- (root): resources is required
- image: repository is required
```

## `template.missing_template` (11)

### 1. `ESGF/esgf-helm`

- Chart: `cloned\ESGF__esgf-helm`
- Source: `template`
- Command: `helm template test cloned\ESGF__esgf-helm`

```text
Error: template: esgf-helm/templates/tds/deployment.yaml:60:34: executing "esgf-helm/templates/tds/deployment.yaml" at <include (print $.Template.BasePath "/auth/secret.yaml") .>: error calling include: template: esgf-helm/templates/auth/secret.yaml:9:30: executing "esgf-helm/templates/auth/secret.yaml" at <index .Values.secrets "auth-database-password">: error calling index: index of untyped nil

Use --debug flag to render out invalid YAML
```

### 2. `stsurya/Kubernetes`

- Chart: `cloned\stsurya__Kubernetes\HelmCharts\replicasLoop`
- Source: `template`
- Command: `helm template test cloned\stsurya__Kubernetes\HelmCharts\replicasLoop`

```text
Error: template: replicasLoop/templates/serviceaccount.yaml:5:11: executing "replicasLoop/templates/serviceaccount.yaml" at <include "replicasLoop.serviceAccountName" .>: error calling include: template: no template "replicasLoop.serviceAccountName" associated with template "gotpl"

Use --debug flag to render out invalid YAML
```

### 3. `EOEPCA/um-identity-service`

- Chart: `cloned\EOEPCA__um-identity-service\helm\charts\identity-spring-boot-echo`
- Source: `template`
- Command: `helm template test cloned\EOEPCA__um-identity-service\helm\charts\identity-spring-boot-echo`

```text
Error: template: identity-spring-boot-echo/templates/deployment.yaml:4:11: executing "identity-spring-boot-echo/templates/deployment.yaml" at <include "identity-api.name" .>: error calling include: template: no template "identity-api.name" associated with template "gotpl"

Use --debug flag to render out invalid YAML
```

### 4. `0xPolygonID/issuer-node`

- Chart: `cloned\0xPolygonID__issuer-node\k8s\helm`
- Source: `template`
- Command: `helm template test cloned\0xPolygonID__issuer-node\k8s\helm`

```text
Error: template: privadoid-issuer/templates/issuer-node-service-account.yml:7:10: executing "privadoid-issuer/templates/issuer-node-service-account.yml" at <include "polygon-id-issuer.staticLabel" .>: error calling include: template: no template "polygon-id-issuer.staticLabel" associated with template "gotpl"

Use --debug flag to render out invalid YAML
```

### 5. `0xPolygonID/issuer-node`

- Chart: `cloned\0xPolygonID__issuer-node\k8s\helm\charts\postgres`
- Source: `template`
- Command: `helm template test cloned\0xPolygonID__issuer-node\k8s\helm\charts\postgres`

```text
Error: template: postgres/templates/service.yaml:7:8: executing "postgres/templates/service.yaml" at <include "privadoid-issuer.postgresIssuerNode.common.labels" .>: error calling include: template: postgres/templates/_helpers.tpl:5:18: executing "privadoid-issuer.postgresIssuerNode.common.labels" at <include "privadoid-issuer.chart" .>: error calling include: template: no template "privadoid-issuer.chart" associated with template "gotpl"

Use --debug flag to render out invalid YAML
```

### 6. `0xPolygonID/issuer-node`

- Chart: `cloned\0xPolygonID__issuer-node\k8s\helm\charts\redis`
- Source: `template`
- Command: `helm template test cloned\0xPolygonID__issuer-node\k8s\helm\charts\redis`

```text
Error: template: redis/templates/service.yaml:7:8: executing "redis/templates/service.yaml" at <include "privadoid-issuer.redisIssuerNode.common.labels" .>: error calling include: template: redis/templates/_helpers.tpl:5:18: executing "privadoid-issuer.redisIssuerNode.common.labels" at <include "privadoid-issuer.chart" .>: error calling include: template: no template "privadoid-issuer.chart" associated with template "gotpl"

Use --debug flag to render out invalid YAML
```

### 7. `0xPolygonID/issuer-node`

- Chart: `cloned\0xPolygonID__issuer-node\k8s\helm\charts\vault`
- Source: `template`
- Command: `helm template test cloned\0xPolygonID__issuer-node\k8s\helm\charts\vault`

```text
Error: template: vault/templates/service.yaml:7:8: executing "vault/templates/service.yaml" at <include "privadoid-issuer.vaultIssuerNode.common.labels" .>: error calling include: template: vault/templates/_helpers.tpl:5:18: executing "privadoid-issuer.vaultIssuerNode.common.labels" at <include "privadoid-issuer.chart" .>: error calling include: template: no template "privadoid-issuer.chart" associated with template "gotpl"

Use --debug flag to render out invalid YAML
```

### 8. `DOME-Marketplace/dome-gitops`

- Chart: `cloned\DOME-Marketplace__dome-gitops\ionos\dome-trust\trusted-issuers-list`
- Source: `template`
- Command: `helm template test cloned\DOME-Marketplace__dome-gitops\ionos\dome-trust\trusted-issuers-list`

```text
Error: template: trusted-issuers-list/charts/trusted-issuers-list/templates/deployment.yaml:85:26: executing "trusted-issuers-list/charts/trusted-issuers-list/templates/deployment.yaml" at <include "til.passwordKey" .>: error calling include: template: trusted-issuers-list/charts/trusted-issuers-list/templates/_helpers.tpl:71:36: executing "til.passwordKey" at <.Values.database.existingSecret.key>: wrong type for value; expected string; got interface {}

Use --debug flag to render out invalid YAML
```

### 9. `DOME-Marketplace/dome-gitops`

- Chart: `cloned\DOME-Marketplace__dome-gitops\ionos\marketplace\iam\trusted-issuers-list`
- Source: `template`
- Command: `helm template test cloned\DOME-Marketplace__dome-gitops\ionos\marketplace\iam\trusted-issuers-list`

```text
Error: template: trusted-issuers-list/charts/trusted-issuers-list/templates/deployment.yaml:85:26: executing "trusted-issuers-list/charts/trusted-issuers-list/templates/deployment.yaml" at <include "til.passwordKey" .>: error calling include: template: trusted-issuers-list/charts/trusted-issuers-list/templates/_helpers.tpl:71:36: executing "til.passwordKey" at <.Values.database.existingSecret.key>: wrong type for value; expected string; got interface {}

Use --debug flag to render out invalid YAML
```

### 10. `openidl-org/openidl-aais-gitops`

- Chart: `cloned\openidl-org__openidl-aais-gitops\openidl-k8s\charts\openidl-utilities`
- Source: `template`
- Command: `helm template test cloned\openidl-org__openidl-aais-gitops\openidl-k8s\charts\openidl-utilities`

```text
Error: template: openidl-utilities/templates/tests/test-connection.yaml:4:12: executing "openidl-utilities/templates/tests/test-connection.yaml" at <include "openidl-data-call-app.fullname" .>: error calling include: template: no template "openidl-data-call-app.fullname" associated with template "gotpl"

Use --debug flag to render out invalid YAML
```

### 11. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\adp\charts\ti-common`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\adp\charts\ti-common`

```text
Error: template: ti-common/templates/serviceaccount_reader.yaml:1:11: executing "ti-common/templates/serviceaccount_reader.yaml" at <include "needInitContainer" .>: error calling include: template: no template "needInitContainer" associated with template "gotpl"

Use --debug flag to render out invalid YAML
```

## `template.yaml_render` (7)

### 1. `aicacia/ts-auth`

- Chart: `cloned\aicacia__ts-auth\helm\auth-ui`
- Source: `template`
- Command: `helm template test cloned\aicacia__ts-auth\helm\auth-ui`

```text
Error: YAML parse error on auth-ui/templates/service.yaml: error converting YAML to JSON: yaml: line 4: did not find expected node content

Use --debug flag to render out invalid YAML
```

### 2. `clayrisser/charts`

- Chart: `cloned\clayrisser__charts\stable\crossplane-on-eks`
- Source: `template`
- Command: `helm template test cloned\clayrisser__charts\stable\crossplane-on-eks`

```text
Error: YAML parse error on crossplane-on-eks/templates/delegates/README.md: error converting YAML to JSON: yaml: line 3: did not find expected comment or line break

Use --debug flag to render out invalid YAML
```

### 3. `ndebuhr/isidro`

- Chart: `cloned\ndebuhr__isidro\chart`
- Source: `template`
- Command: `helm template test cloned\ndebuhr__isidro\chart`

```text
Error: YAML parse error on isidro/templates/prometheus.yaml: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type releaseutil.SimpleHead

Use --debug flag to render out invalid YAML
```

### 4. `kubesphere/ksbuilder`

- Chart: `cloned\kubesphere__ksbuilder\pkg\extension\templates\charts\backend`
- Source: `template`
- Command: `helm template test cloned\kubesphere__ksbuilder\pkg\extension\templates\charts\backend`

```text
Error: YAML parse error on backend/templates/extensions.yaml: error converting YAML to JSON: yaml: line 5: did not find expected key

Use --debug flag to render out invalid YAML
```

### 5. `kubesphere/ksbuilder`

- Chart: `cloned\kubesphere__ksbuilder\pkg\extension\templatesapp\charts\base`
- Source: `template`
- Command: `helm template test cloned\kubesphere__ksbuilder\pkg\extension\templatesapp\charts\base`

```text
Error: YAML parse error on base/templates/app.yaml: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal array into Go struct field .metadata.annotations of type string

Use --debug flag to render out invalid YAML
```

### 6. `chandan-kumar-11/PHP-deployment`

- Chart: `cloned\chandan-kumar-11__PHP-deployment\php-app-helm-chat\php-app-chart`
- Source: `template`
- Command: `helm template test cloned\chandan-kumar-11__PHP-deployment\php-app-helm-chat\php-app-chart`

```text
Error: YAML parse error on php-app-chart/templates/tms.sql: error converting YAML to JSON: yaml: line 5: mapping values are not allowed in this context

Use --debug flag to render out invalid YAML
```

### 7. `trikorasolns/helm-charts`

- Chart: `cloned\trikorasolns__helm-charts\charts\nodered`
- Source: `template`
- Command: `helm template test cloned\trikorasolns__helm-charts\charts\nodered`

```text
Error: YAML parse error on nodered/templates/deployment.yaml: error converting YAML to JSON: yaml: line 36: mapping values are not allowed in this context

Use --debug flag to render out invalid YAML
```

## `dependency.lock_file_out_of_sync` (6)

### 1. `ConductionNL/Authorization-component`

- Chart: `cloned\ConductionNL__Authorization-component\api\helm`
- Source: `dependency`

```text
load.go:138: Warning: Dependency locking is handled in Chart.lock since apiVersion "v2". We recommend migrating to Chart.lock.
Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies
```

### 2. `samcab28/Diseno-Software`

- Chart: `cloned\samcab28__Diseno-Software\proyecto\caso4\charts\databases`
- Source: `dependency`

```text
Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies
```

### 3. `rancher/charts`

- Chart: `cloned\rancher__charts\charts\epinio\102.0.1+up1.6.2`
- Source: `dependency`

```text
Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies
```

### 4. `nadav07/helmTesting`

- Chart: `cloned\nadav07__helmTesting\chartTest`
- Source: `dependency`

```text
Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies
```

### 5. `deepdivenow/rke2-charts`

- Chart: `cloned\deepdivenow__rke2-charts\charts\harvester-cloud-provider\harvester-cloud-provider\0.2.100`
- Source: `dependency`

```text
Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies
```

### 6. `yatvoikoresh/cloneranchercharts`

- Chart: `cloned\yatvoikoresh__cloneranchercharts\charts\epinio\102.0.1+up1.6.2`
- Source: `dependency`

```text
Error: the lock file (Chart.lock) is out of sync with the dependencies file (Chart.yaml). Please update the dependencies
```

## `template.dependency_check_failed` (6)

### 1. `chayeonhee/msa2`

- Chart: `cloned\chayeonhee__msa2\accounts\helm-chart`
- Source: `template`
- Command: `helm template test cloned\chayeonhee__msa2\accounts\helm-chart`

```text
Error: An error occurred while checking for chart dependencies. You may need to run `helm dependency build` to fetch missing dependencies: found in Chart.yaml, but missing in charts/ directory: eazybank-common
```

### 2. `chayeonhee/msa2`

- Chart: `cloned\chayeonhee__msa2\cards\helm-chart`
- Source: `template`
- Command: `helm template test cloned\chayeonhee__msa2\cards\helm-chart`

```text
Error: An error occurred while checking for chart dependencies. You may need to run `helm dependency build` to fetch missing dependencies: found in Chart.yaml, but missing in charts/ directory: eazybank-common
```

### 3. `chayeonhee/msa2`

- Chart: `cloned\chayeonhee__msa2\configserver\helm-chart`
- Source: `template`
- Command: `helm template test cloned\chayeonhee__msa2\configserver\helm-chart`

```text
Error: An error occurred while checking for chart dependencies. You may need to run `helm dependency build` to fetch missing dependencies: found in Chart.yaml, but missing in charts/ directory: eazybank-common
```

### 4. `chayeonhee/msa2`

- Chart: `cloned\chayeonhee__msa2\eurekaserver\helm-chart`
- Source: `template`
- Command: `helm template test cloned\chayeonhee__msa2\eurekaserver\helm-chart`

```text
Error: An error occurred while checking for chart dependencies. You may need to run `helm dependency build` to fetch missing dependencies: found in Chart.yaml, but missing in charts/ directory: eazybank-common
```

### 5. `chayeonhee/msa2`

- Chart: `cloned\chayeonhee__msa2\gatewayserver\helm-chart`
- Source: `template`
- Command: `helm template test cloned\chayeonhee__msa2\gatewayserver\helm-chart`

```text
Error: An error occurred while checking for chart dependencies. You may need to run `helm dependency build` to fetch missing dependencies: found in Chart.yaml, but missing in charts/ directory: eazybank-common
```

### 6. `chayeonhee/msa2`

- Chart: `cloned\chayeonhee__msa2\loans\helm-chart`
- Source: `template`
- Command: `helm template test cloned\chayeonhee__msa2\loans\helm-chart`

```text
Error: An error occurred while checking for chart dependencies. You may need to run `helm dependency build` to fetch missing dependencies: found in Chart.yaml, but missing in charts/ directory: eazybank-common
```

## `template.library_chart_not_installable` (6)

### 1. `SAP/component-operator-runtime`

- Chart: `cloned\SAP__component-operator-runtime\internal\helm\testdata\main\charts\lib11`
- Source: `template`
- Command: `helm template test cloned\SAP__component-operator-runtime\internal\helm\testdata\main\charts\lib11`

```text
Error: library charts are not installable
```

### 2. `Gemeente-DenHaag/helm-charts`

- Chart: `cloned\Gemeente-DenHaag__helm-charts\charts\dh-lib`
- Source: `template`
- Command: `helm template test cloned\Gemeente-DenHaag__helm-charts\charts\dh-lib`

```text
Error: library charts are not installable
```

### 3. `sapcc/helm-charts`

- Chart: `cloned\sapcc__helm-charts\common\helm3-helper`
- Source: `template`
- Command: `helm template test cloned\sapcc__helm-charts\common\helm3-helper`

```text
Error: library charts are not installable
```

### 4. `nagcloudlab/java-devsecops`

- Chart: `cloned\nagcloudlab__java-devsecops\docker-k8s\2-k8s-lab\10-helm\kubernetes\helm\common`
- Source: `template`
- Command: `helm template test cloned\nagcloudlab__java-devsecops\docker-k8s\2-k8s-lab\10-helm\kubernetes\helm\common`

```text
Error: library charts are not installable
```

### 5. `nagcloudlab/java-devsecops`

- Chart: `cloned\nagcloudlab__java-devsecops\week-4\k8s-lab\10-helm\kubernetes\helm\common`
- Source: `template`
- Command: `helm template test cloned\nagcloudlab__java-devsecops\week-4\k8s-lab\10-helm\kubernetes\helm\common`

```text
Error: library charts are not installable
```

### 6. `trueforge-org/truecharts`

- Chart: `cloned\trueforge-org__truecharts\charts\library\common`
- Source: `template`
- Command: `helm template test cloned\trueforge-org__truecharts\charts\library\common`

```text
Error: library charts are not installable
```

## `dependency.repo_update` (4)

### 1. `ConductionNL/checkinservice`

- Chart: `cloned\ConductionNL__checkinservice\api\helm`
- Source: `dependency`

```text
Getting updates for unmanaged Helm repositories...
...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Error: can't get a valid version for 1 subchart(s): "postgresql" (repository "https://charts.bitnami.com/bitnami", version "10.1.1"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml
```

### 2. `robertbenzinger/charts`

- Chart: `cloned\robertbenzinger__charts\janusgraph`
- Source: `dependency`

```text
Getting updates for unmanaged Helm repositories...
...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository
...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Error: can't get a valid version for 2 subchart(s): "cassandra" (repository "https://charts.bitnami.com/bitnami", version "5.6.1"), "elasticsearch" (repository "https://charts.bitnami.com/bitnami", version "12.6.2"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml
```

### 3. `SBasak8/eu-val-loki`

- Chart: `cloned\SBasak8__eu-val-loki\charts\airflow`
- Source: `dependency`

```text
Getting updates for unmanaged Helm repositories...
...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Error: can't get a valid version for 1 subchart(s): "postgresql" (repository "https://charts.bitnami.com/bitnami", version "10.5.3"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml
```

### 4. `svtechnmaa/charts`

- Chart: `cloned\svtechnmaa__charts\kubernetes\airflow`
- Source: `dependency`

```text
Getting updates for unmanaged Helm repositories...
...Successfully got an update from the "https://charts.bitnami.com/bitnami" chart repository
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "traefik" chart repository
Update Complete. ⎈Happy Helming!⎈
Error: can't get a valid version for 1 subchart(s): "postgresql" (repository "https://charts.bitnami.com/bitnami", version "10.5.3"). Make sure a matching chart version exists in the repo, or change the version constraint in Chart.yaml
```

## `dependency.chart_validation` (3)

### 1. `solo-io/skv2`

- Chart: `cloned\solo-io__skv2\codegen\test\chart\conditional-crds`
- Source: `dependency`

```text
Error: validation: chart.metadata.name is required
```

### 2. `krateoplatformops/composition-portal-starter`

- Chart: `cloned\krateoplatformops__composition-portal-starter\chart`
- Source: `dependency`

```text
Error: validation: chart.metadata.version "CHART_VERSION" is invalid
```

### 3. `teamhephy/redis`

- Chart: `cloned\teamhephy__redis\charts\redis`
- Source: `dependency`

```text
Error: validation: chart.metadata.version "<Will be populated by the ci before publishing the chart>" is invalid
```

## `template.parse_error` (3)

### 1. `sostheim/chart-krak8s-api`

- Chart: `cloned\sostheim__chart-krak8s-api\krak8s-api`
- Source: `template`
- Command: `helm template test cloned\sostheim__chart-krak8s-api\krak8s-api`

```text
Error: parse error at (krak8s/templates/aws_secrets.yaml:1): unexpected "=" in operand

Use --debug flag to render out invalid YAML
```

### 2. `marcus-sa/helm-charts`

- Chart: `cloned\marcus-sa__helm-charts\repositories\appscode\searchlight\7.0.0\chart`
- Source: `template`
- Command: `helm template test cloned\marcus-sa__helm-charts\repositories\appscode\searchlight\7.0.0\chart`

```text
Error: parse error at (searchlight/templates/secret.yaml:78): unexpected EOF

Use --debug flag to render out invalid YAML
```

### 3. `marcus-sa/helm-charts`

- Chart: `cloned\marcus-sa__helm-charts\repositories\appscode\searchlight\7.0.0-rc.0\chart`
- Source: `template`
- Command: `helm template test cloned\marcus-sa__helm-charts\repositories\appscode\searchlight\7.0.0-rc.0\chart`

```text
Error: parse error at (searchlight/templates/secret.yaml:78): unexpected EOF

Use --debug flag to render out invalid YAML
```

## `dependency.unpack_error` (1)

### 1. `kingdonb/gitops-hephy`

- Chart: `cloned\kingdonb__gitops-hephy\workflow`
- Source: `dependency`

```text
Error: error unpacking subchart monitor in workflow: error unpacking subchart grafana in monitor: validation: chart.metadata.version "<Will be populated by the ci before publishing the chart>" is invalid
```

## `template.kube_version_incompatible` (1)

### 1. `tkestack/charts`

- Chart: `cloned\tkestack__charts\incubator\aerospike`
- Source: `template`
- Command: `helm template test cloned\tkestack__charts\incubator\aerospike`

```text
Error: chart requires kubeVersion: <= 1.22.5-x which is incompatible with Kubernetes v1.33.0

Use --debug flag to render out invalid YAML
```

## `template.values_merge_error` (1)

### 1. `trueforge-org/truecharts`

- Chart: `cloned\trueforge-org__truecharts\charts\stable\clamav`
- Source: `template`
- Command: `helm template test cloned\trueforge-org__truecharts\charts\stable\clamav`

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
  tag: 4.109.2@sha256:841cf00f26ed44db1d0be33cf2e9afff96ac298fd655d484d28be5d86b945930
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
  tag: 1.5.2@sha256:f8347147079772cc2f0b2d3cf1059b9c7031c4afb543b93733df78463c0cc6aa
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
  tag: 1.31.1@sha256:49b02f4dad562e2fb4031ebcd9d47d432ff760a96b66c4d35a10ad21451a22f2
mariadb:
  creds: {}
  enabled: false
  includeCommon: false
  password: PLACEHOLDERPASSWORD
  rootPassword: PLACEHOLDERROOTPASSWORD
mariadbClientImage:
  pullPolicy: IfNotPresent
  repository: oci.trueforge.org/containerforge/mariadb-client
  tag: 1.1.0@sha256:fa063910d33f0196cdf3db7c4c375fa59e295feee33a1cc71ebf71f23eca5a16
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
  tag: 1.10.6@sha256:a51d3296f173394635c88ab483410f1d68b572e76554a02d256fbd39f43f8ba0
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
  tag: 15.17@sha256:e2c52cb374b218f4ea175f331530316d5c0000e56a7089114f54a222db879ff8
postgres16Image:
  pullPolicy: IfNotPresent
  repository: ghcr.io/cloudnative-pg/postgresql
  tag: 16.12@sha256:7fe820e900df4492243a1d4aeff88e1eb2b9706ac30bb1f10565e4a71099023f
postgresClientImage:
  pullPolicy: IfNotPresent
  repository: oci.trueforge.org/containerforge/postgresql-client
  tag: 1.1.0@sha256:4df4c77540485bbb4a681268b9a7b33b72a51cf9e6c4d4cd8f41db077afe4f8a
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
  tag: 1.1.0@sha256:8bbd28e12efc6ee35fc1be5220ea9553739008fb2a4d292e88306150e7e52e79
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
  tag: 1.0.0@sha256:fcb59c5d301102db2f167dcb0fd7db22acf24a3d77c5acc7f464c563006fd638
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
  tag: 24.04@sha256:1c874769d7dc656030d189c30532396d444ab56a283f2828480b4c8a06359ee8
valkeyClientImage:
  pullPolicy: IfNotPresent
  repository: oci.trueforge.org/containerforge/valkey-tools
  tag: 1.1.0@sha256:8bbd28e12efc6ee35fc1be5220ea9553739008fb2a4d292e88306150e7e52e79
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
  tag: 24.04@sha256:1c874769d7dc656030d189c30532396d444ab56a283f2828480b4c8a06359ee8
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
  tag: 4.52.4@sha256:95132ba7539ee4692c2a33d8ceaa49c49a9106893e2c35dcadba9c540fa3a9dc 
 ============================================================================================= 

 See error above values.

Use --debug flag to render out invalid YAML
```

