# Five engines x three scenarios (grouped)

Charts per scenario: raw = `k8s/`; production = `madgoat` + `madgoat-infra`; render-problems = `madgoat-render` + `madgoat-infra-render`.

## Table 1 - Findings (summed per scenario)

| Scenario | kics-default | kics-enhanced | kubescape | trivy | kube-linter |
|---|--:|--:|--:|--:|--:|
| raw (k8s manifests) | 408 | 408 | 279 | 291 | 54 |
| production (helm) | 17 | 518 | 359 | 213 | 74 |
| render-problems | 183 | 493 | 120 | 0 | 24 |

## Table 2 - OWASP Kubernetes Top 10 (union per scenario)

| Scenario | kics-default | kics-enhanced | kubescape | trivy | kube-linter |
|---|---|---|---|---|---|
| raw (k8s manifests) | K01/K05/K07 | K01/K05/K07 | K01/K03/K05 | K01/K03 | K01 |
| production (helm) | K05 | K01/K02/K05/K09 | K01/K02/K03/K05 | K01/K02/K03 | K01 |
| render-problems | K01/K02/K05/K09 | K01/K02/K05/K09 | K01/K02/K03/K05 | — | K01 |
