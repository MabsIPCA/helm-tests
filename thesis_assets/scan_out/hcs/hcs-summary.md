## 🔎 HCS Helm Chart scan

### Misconfigurations

| Critical | High | Medium | Low | Info |
|:-:|:-:|:-:|:-:|:-:|
| 0 | 22 | 134 | 152 | 14 |

<details><summary>Top misconfigurations</summary>

- **HIGH** Container Is Privileged `../../../../mad-deployment-service/helm/render-problems/madgoat-render/templates/deployment.yaml:45`
- **HIGH** Privilege Escalation Allowed `../../../../mad-deployment-service/helm/render-problems/madgoat-render/charts/madgoat-infra-render/templates/core.yaml:66`
- **HIGH** Volume Mount With OS Directory Write Permissions `../../../../mad-deployment-service/helm/render-problems/madgoat-render/charts/madgoat-infra-render/templates/db.yaml:116`
- **MEDIUM** Container Running As Root `../../../../mad-deployment-service/helm/render-problems/madgoat-render/charts/madgoat-infra-render/templates/core.yaml:40`
- **MEDIUM** Container Running With Low UID `../../../../mad-deployment-service/helm/render-problems/madgoat-render/templates/deployment.yaml:45`
- **MEDIUM** Memory Limits Not Defined `../../../../mad-deployment-service/helm/render-problems/madgoat-render/templates/deployment.yaml:45`
- **MEDIUM** Memory Requests Not Defined `../../../../mad-deployment-service/helm/render-problems/madgoat-render/charts/madgoat-infra-render/templates/db.yaml:38`
- **MEDIUM** NET_RAW Capabilities Not Being Dropped `../../../../mad-deployment-service/helm/render-problems/madgoat-render/charts/madgoat-infra-render/templates/db.yaml:38`
- **MEDIUM** RBAC Roles with Read Secrets Permissions `../../../../mad-deployment-service/helm/render-problems/madgoat-render/charts/madgoat-infra-render/templates/rbac.yaml:27`
- **MEDIUM** Readiness Probe Is Not Configured `../../../../mad-deployment-service/helm/render-problems/madgoat-render/templates/deployment.yaml:45`
- **MEDIUM** Role Binding To Default Service Account `../../../../mad-deployment-service/helm/render-problems/madgoat-render/charts/madgoat-infra-render/templates/rbac.yaml:45`
- **MEDIUM** Seccomp Profile Is Not Configured `../../../../mad-deployment-service/helm/render-problems/madgoat-render/charts/madgoat-infra-render/templates/db.yaml:48`
- **MEDIUM** Service Account Token Automount Not Disabled `../../../../mad-deployment-service/helm/render-problems/madgoat-render/charts/madgoat-infra-render/templates/db.yaml:36`
- **MEDIUM** ServiceAccount Allows Access Secrets `../../../../mad-deployment-service/helm/render-problems/madgoat-render/charts/madgoat-infra-render/templates/rbac.yaml:27`
- **MEDIUM** Using Unrecommended Namespace `../../../../mad-deployment-service/helm/render-problems/madgoat-render/charts/madgoat-infra-render/templates/configmap.yaml:6`
- **LOW** CPU Limits Not Set `../../../../mad-deployment-service/helm/render-problems/madgoat-render/charts/madgoat-infra-render/templates/db.yaml:38`
- **LOW** CPU Requests Not Set `../../../../mad-deployment-service/helm/render-problems/madgoat-render/charts/madgoat-infra-render/templates/db.yaml:38`
- **LOW** Cluster Admin Rolebinding With Superuser Permissions `../../../../mad-deployment-service/helm/render-problems/madgoat-render/charts/madgoat-infra-render/templates/rbac.yaml:9`
- **LOW** Image Pull Policy Of The Container Is Not Set To Always `../../../../mad-deployment-service/helm/render-problems/madgoat-render/templates/deployment.yaml:45`
- **LOW** Image Without Digest `../../../../mad-deployment-service/helm/render-problems/madgoat-render/charts/madgoat-infra-render/templates/db.yaml:48`

</details>

### Image vulnerabilities

| Image | Source | Critical | High | Medium | Low |
|-------|--------|:-:|:-:|:-:|:-:|
| library/mongo:6.0 | /mnt/c/Users/miabs/GolandProjects/mad-deployment-service/helm/render-problems/madgoat-render/charts/madgoat-infra-render/templates/db.yaml:37 | 1 | 184 | 171 | 79 |
| library/postgres:14.1-alpine | /mnt/c/Users/miabs/GolandProjects/mad-deployment-service/helm/render-problems/madgoat-render/charts/madgoat-infra-render/templates/db.yaml:47 | 3 | 27 | 19 | 0 |
| library/rabbitmq:3-management-alpine | /mnt/c/Users/miabs/GolandProjects/mad-deployment-service/helm/render-problems/madgoat-render/charts/madgoat-infra-render/templates/core.yaml:23 | 2 | 22 | 36 | 41 |
| mad-goat-project/db-lesson:data | /mnt/c/Users/miabs/GolandProjects/mad-deployment-service/helm/render-problems/madgoat-render/charts/madgoat-infra-render/templates/db.yaml:26 | 2 | 17 | 26 | 41 |
| mad-goat-project/keycloak:main | /mnt/c/Users/miabs/GolandProjects/mad-deployment-service/helm/render-problems/madgoat-render/charts/madgoat-infra-render/templates/core.yaml:52 | 1 | 105 | 252 | 71 |
| mad-goat-project/mad-goat-docs:main | /mnt/c/Users/miabs/GolandProjects/mad-deployment-service/helm/render-problems/madgoat-render/templates/deployment.yaml:23 | 0 | 8 | 17 | 3 |
| mad-goat-project/mad-goat4shell-service:safe | /mnt/c/Users/miabs/GolandProjects/mad-deployment-service/helm/render-problems/madgoat-render/templates/deployment.yaml:23 | 8 | 47 | 477 | 199 |
| mad-goat-project/mad-goat4shell-service:unsafe | /mnt/c/Users/miabs/GolandProjects/mad-deployment-service/helm/render-problems/madgoat-render/templates/deployment.yaml:23 | 10 | 48 | 477 | 200 |
| mad-goat-project/mad-lessons-service:main | /mnt/c/Users/miabs/GolandProjects/mad-deployment-service/helm/render-problems/madgoat-render/templates/deployment.yaml:23 | 7 | 71 | 71 | 38 |
| mad-goat-project/mad-profile-service:main | /mnt/c/Users/miabs/GolandProjects/mad-deployment-service/helm/render-problems/madgoat-render/templates/deployment.yaml:23 | 3 | 45 | 75 | 22 |
| mad-goat-project/mad-scoreboard-service:main | /mnt/c/Users/miabs/GolandProjects/mad-deployment-service/helm/render-problems/madgoat-render/templates/deployment.yaml:23 | 4 | 65 | 56 | 34 |
| mad-goat-project/mad-web-app:madgoat-tech | /mnt/c/Users/miabs/GolandProjects/mad-deployment-service/helm/render-problems/madgoat-render/templates/deployment.yaml:23 | 9 | 52 | 28 | 6 |
| mad-goat-project/mc-minio:mc-minio | /mnt/c/Users/miabs/GolandProjects/mad-deployment-service/helm/render-problems/madgoat-render/charts/madgoat-infra-render/templates/db.yaml:47 | 5 | 159 | 483 | 242 |
| mad-goat-project/minio:data | /mnt/c/Users/miabs/GolandProjects/mad-deployment-service/helm/render-problems/madgoat-render/charts/madgoat-infra-render/templates/db.yaml:26 | 2 | 17 | 26 | 41 |

<!-- hcs -->