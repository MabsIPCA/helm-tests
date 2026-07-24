# Five engines x three scenarios (findings)

| Scenario | Target | kics-default | kics-enhanced | kubescape | trivy | kube-linter |
|---|---|--:|--:|--:|--:|--:|
| k8s (raw manifests) | k8s | 408 | 408 | 279 | 291 | 54 |
| helm (production charts) | madgoat | 6 | 327 | 239 | 136 | 48 |
|  | madgoat-infra | 11 | 191 | 120 | 77 | 26 |
| render-problems (broken) | madgoat-render | 6 | 316 | 0 | 0 | 0 |
|  | madgoat-infra-render | 177 | 177 | 120 | 0 | 24 |

## OWASP K-Top-10 detected

| Scenario | Target | kics-default | kics-enhanced | kubescape | trivy | kube-linter |
|---|---|---|---|---|---|---|
| k8s (raw manifests) | k8s | K01/K05/K07 | K01/K05/K07 | K01/K03/K05 | K01/K03 | K01 |
| helm (production charts) | madgoat | — | K01/K02/K05/K09 | K01/K02/K03/K05 | K01/K02/K03 | K01 |
|  | madgoat-infra | K05 | K01/K02/K05/K09 | K01/K02/K03/K05 | K01/K02/K03 | K01 |
| render-problems (broken) | madgoat-render | — | K01/K02/K05/K09 | — | — | — |
|  | madgoat-infra-render | K01/K02/K05/K09 | K01/K02/K05/K09 | K01/K02/K03/K05 | — | K01 |
