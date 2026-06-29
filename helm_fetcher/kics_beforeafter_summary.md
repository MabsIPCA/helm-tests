# KICS Before/After: `--experimental-helm-scan` on fixer-resolved charts

Baseline run: current (589 fixer-resolved charts). before = flag off, after = flag on.

**Total findings: 4272 -> 33368 (+29096, +681%)**

## By severity

| Severity | Before | After | Delta |
|---|---:|---:|---:|
| CRITICAL | 0 | 0 | +0 |
| HIGH | 359 | 3033 | +2674 |
| MEDIUM | 1715 | 12728 | +11013 |
| LOW | 1994 | 15792 | +13798 |
| INFO | 204 | 1815 | +1611 |
| TRACE | 0 | 0 | +0 |
| **Total** | **4272** | **33368** | **+29096** |

## By category

| Category | Before | After | Delta |
|---|---:|---:|---:|
| Insecure Configurations | 1799 | 13887 | +12088 |
| Best Practices | 735 | 4738 | +4003 |
| Resource Management | 311 | 4180 | +3869 |
| Access Control | 475 | 3944 | +3469 |
| Availability | 293 | 1924 | +1631 |
| Build Process | 221 | 1740 | +1519 |
| Secret Management | 122 | 1411 | +1289 |
| Insecure Defaults | 171 | 1308 | +1137 |
| Networking and Firewall | 12 | 80 | +68 |
| Supply-Chain | 133 | 156 | +23 |

## By query (all, sorted by delta)

| Delta | Before | After | Severity | Query | CWE | Category |
|---:|---:|---:|---|---|---|---|
| +1853 | 226 | 2079 | LOW | Missing AppArmor Profile | 284 | Access Control |
| +1813 | 226 | 2039 | LOW | Image Without Digest | 665 | Insecure Configurations |
| +1676 | 214 | 1890 | LOW | Image Pull Policy Of The Container Is Not Set To Always | 665 | Insecure Configurations |
| +1645 | 221 | 1866 | MEDIUM | Seccomp Profile Is Not Configured | 665 | Insecure Configurations |
| +1611 | 220 | 1831 | MEDIUM | Container Running With Low UID | 1188 | Best Practices |
| +1469 | 210 | 1679 | HIGH | Privilege Escalation Allowed | 269 | Insecure Configurations |
| +1440 | 214 | 1654 | LOW | Root Container Not Mounted Read-only | 668 | Build Process |
| +1299 | 212 | 1511 | LOW | No Drop Capabilities for Containers | 754 | Best Practices |
| +1289 | 212 | 1501 | MEDIUM | NET_RAW Capabilities Not Being Dropped | 269 | Insecure Configurations |
| +1138 | 170 | 1308 | MEDIUM | Service Account Token Automount Not Disabled | 665 | Insecure Defaults |
| +1134 | 204 | 1338 | MEDIUM | Container Running As Root | 1188 | Best Practices |
| +1104 | 74 | 1178 | LOW | CPU Limits Not Set | 400 | Resource Management |
| +986 | 172 | 1158 | LOW | Pod or Container Without LimitRange | 400 | Insecure Configurations |
| +986 | 172 | 1158 | LOW | Pod or Container Without ResourceQuota | 400 | Insecure Configurations |
| +944 | 73 | 1017 | MEDIUM | Memory Limits Not Defined | 400 | Resource Management |
| +928 | 93 | 1021 | MEDIUM | Using Unrecommended Namespace | 665 | Insecure Configurations |
| +868 | 168 | 1036 | LOW | Pod or Container Without Security Context | 285 | Insecure Configurations |
| +799 | 133 | 932 | INFO | Liveness Probe Is Not Defined | 754 | Availability |
| +791 | 69 | 860 | LOW | CPU Requests Not Set | 400 | Resource Management |
| +774 | 39 | 813 | HIGH | Non Kube System Pod With Host Mount | 668 | Access Control |
| +742 | 139 | 881 | MEDIUM | Readiness Probe Is Not Configured | 754 | Availability |
| +728 | 69 | 797 | MEDIUM | Memory Requests Not Defined | 400 | Resource Management |
| +613 | 32 | 645 | LOW | Secrets As Environment Variables | 526 | Secret Management |
| +501 | 44 | 545 | INFO | Ensure Administrative Boundaries Between Resources | 284 | Access Control |
| +308 | 27 | 335 | INFO | Using Kubernetes Native Secret Management | 311 | Secret Management |
| +277 | 54 | 331 | MEDIUM | RBAC Roles with Read Secrets Permissions | 732 | Access Control |
| +213 | 55 | 268 | MEDIUM | ServiceAccount Allows Access Secrets | 522 | Secret Management |
| +202 | 11 | 213 | HIGH | Volume Mount With OS Directory Write Permissions | 284 | Resource Management |
| +151 | 65 | 216 | LOW | Service Does Not Target Pod | 665 | Insecure Configurations |
| +151 | 8 | 159 | MEDIUM | Shared Service Account | 200 | Secret Management |
| +139 | 24 | 163 | HIGH | Workload Mounting With Sensitive OS Directory | 200 | Insecure Configurations |
| +69 | 7 | 76 | HIGH | Container Is Privileged | 269 | Insecure Configurations |
| +67 | 28 | 95 | MEDIUM | Permissive Access to Create Pods | 269 | Access Control |
| +65 | 3 | 68 | LOW | StatefulSet Requests Storage | 665 | Build Process |
| +61 | 0 | 61 | LOW | StatefulSet Without Service Name | 665 | Availability |
| +61 | 2 | 63 | LOW | CronJob Deadline Not Configured | 400 | Resource Management |
| +42 | 6 | 48 | MEDIUM | Service With External Load Balancer | 552 | Networking and Firewall |
| +40 | 7 | 47 | LOW | Object Is Using A Deprecated API Version | 665 | Best Practices |
| +28 | 5 | 33 | MEDIUM | Containers With Added Capabilities | 269 | Insecure Configurations |
| +26 | 9 | 35 | MEDIUM | Shared Host Network Namespace | 200 | Resource Management |
| +25 | 19 | 44 | LOW | Deployment Without PodDisruptionBudget | 400 | Availability |
| +23 | 8 | 31 | LOW | Invalid Image Tag | 665 | Supply-Chain |
| +16 | 0 | 16 | HIGH | Containers With Sys Admin Capabilities | 250 | Insecure Configurations |
| +14 | 0 | 14 | MEDIUM | Incorrect Volume Claim Access Mode ReadWriteOnce | 732 | Build Process |
| +11 | 17 | 28 | MEDIUM | RBAC Roles with Exec Permission | 732 | Access Control |
| +11 | 0 | 11 | LOW | Service Type is NodePort | 665 | Networking and Firewall |
| +8 | 0 | 8 | MEDIUM | Root Containers Admitted | 732 | Best Practices |
| +7 | 0 | 7 | HIGH | Shared Host PID Namespace | 200 | Insecure Configurations |
| +7 | 4 | 11 | LOW | Deployment Has No PodAntiAffinity | 710 | Resource Management |
| +6 | 5 | 11 | LOW | Network Policy Is Not Targeting Any Pod | 665 | Networking and Firewall |
| +4 | 2 | 6 | LOW | StatefulSet Without PodDisruptionBudget | 754 | Availability |
| +4 | 0 | 4 | HIGH | Etcd Peer TLS Certificate Files Not Properly Set | 287 | Networking and Firewall |
| +4 | 0 | 4 | MEDIUM | Etcd TLS Certificate Files Not Properly Set | 287 | Networking and Firewall |
| +4 | 4 | 8 | MEDIUM | Ingress Controller Exposes Workload | 779 | Insecure Configurations |
| +2 | 0 | 2 | LOW | StatefulSet Has No PodAntiAffinity | 400 | Resource Management |
| +2 | 0 | 2 | HIGH | PSP Allows Containers To Share The Host Network Namespace | 250 | Insecure Configurations |
| +2 | 0 | 2 | HIGH | PSP Allows Privilege Escalation | 269 | Insecure Configurations |
| +2 | 0 | 2 | HIGH | PSP Allows Sharing Host IPC | 250 | Insecure Configurations |
| +2 | 0 | 2 | HIGH | PSP Set To Privileged | 732 | Insecure Configurations |
| +2 | 0 | 2 | HIGH | PSP With Added Capabilities | 250 | Insecure Configurations |
| +2 | 0 | 2 | HIGH | PSP With Unrestricted Access to Host Path | 250 | Resource Management |
| +2 | 0 | 2 | MEDIUM | PSP Allows Sharing Host PID | 250 | Insecure Configurations |
| +2 | 0 | 2 | MEDIUM | Shared Host IPC Namespace | 200 | Resource Management |
| +2 | 0 | 2 | INFO | Not Limited Capabilities For Pod Security Policy | 770 | Insecure Configurations |
| +2 | 0 | 2 | MEDIUM | Etcd Client Certificate Authentication Set To False | 287 | Secret Management |
| +2 | 0 | 2 | MEDIUM | Etcd Peer Client Certificate Authentication Set To False | 287 | Secret Management |
| +2 | 1 | 3 | LOW | Cluster Admin Rolebinding With Superuser Permissions | 269 | Access Control |
| +2 | 1 | 3 | MEDIUM | RBAC Roles Allow Privilege Escalation | 288 | Access Control |
| +1 | 0 | 1 | INFO | Bind Address Not Properly Set | 710 | Networking and Firewall |
| +0 | 3 | 3 | HIGH | Missing User Instruction | 250 | Build Process |
| +0 | 6 | 6 | LOW | Healthcheck Instruction Missing | 710 | Insecure Configurations |
| +0 | 123 | 123 | MEDIUM | Apt Get Install Pin Version Not Defined | 1357 | Supply-Chain |
| +0 | 1 | 1 | MEDIUM | Unpinned Package Version in Apk Add | 1357 | Supply-Chain |
| +0 | 1 | 1 | LOW | RUN Instruction Using 'cd' Instead of WORKDIR | 710 | Build Process |
| +0 | 1 | 1 | MEDIUM | Add Instead of Copy | 610 | Supply-Chain |
| +0 | 1 | 1 | LOW | Curl or Wget Instead of Add | 610 | Best Practices |
| +0 | 1 | 1 | LOW | Multiple RUN, ADD, COPY, Instructions Listed | 710 | Best Practices |
| +0 | 1 | 1 | MEDIUM | CNI Plugin Does Not Support Network Policies | 923 | Networking and Firewall |
| -1 | 1 | 0 | MEDIUM | Service Account Name Undefined Or Empty | 665 | Insecure Defaults |
| -18 | 65 | 47 | HIGH | RBAC Wildcard In Rule | 732 | Access Control |
| -89 | 90 | 1 | LOW | Metadata Label Is Invalid | 710 | Best Practices |

## By CWE (sorted by delta)

| Delta | Before | After | CWE | Queries |
|---:|---:|---:|---|---|
| +7556 | 1013 | 8569 | CWE-665 | Image Pull Policy Of The Container Is Not Set To Always, Image With... |
| +5627 | 650 | 6277 | CWE-400 | CPU Limits Not Set, CPU Requests Not Set, CronJob Deadline Not Conf... |
| +2926 | 463 | 3389 | CWE-269 | Cluster Admin Rolebinding With Superuser Permissions, Container Is ... |
| +2844 | 486 | 3330 | CWE-754 | Liveness Probe Is Not Defined, No Drop Capabilities for Containers,... |
| +2745 | 424 | 3169 | CWE-1188 | Container Running As Root, Container Running With Low UID |
| +2556 | 281 | 2837 | CWE-284 | Ensure Administrative Boundaries Between Resources, Missing AppArmo... |
| +2214 | 253 | 2467 | CWE-668 | Non Kube System Pod With Host Mount, Root Container Not Mounted Rea... |
| +868 | 168 | 1036 | CWE-285 | Pod or Container Without Security Context |
| +613 | 32 | 645 | CWE-526 | Secrets As Environment Variables |
| +325 | 41 | 366 | CWE-200 | Shared Host IPC Namespace, Shared Host Network Namespace, Shared Ho... |
| +308 | 27 | 335 | CWE-311 | Using Kubernetes Native Secret Management |
| +294 | 136 | 430 | CWE-732 | Incorrect Volume Claim Access Mode ReadWriteOnce, PSP Set To Privil... |
| +213 | 55 | 268 | CWE-522 | ServiceAccount Allows Access Secrets |
| +42 | 6 | 48 | CWE-552 | Service With External Load Balancer |
| +26 | 3 | 29 | CWE-250 | Containers With Sys Admin Capabilities, Missing User Instruction, P... |
| +12 | 0 | 12 | CWE-287 | Etcd Client Certificate Authentication Set To False, Etcd Peer Clie... |
| +4 | 4 | 8 | CWE-779 | Ingress Controller Exposes Workload |
| +2 | 0 | 2 | CWE-770 | Not Limited Capabilities For Pod Security Policy |
| +2 | 1 | 3 | CWE-288 | RBAC Roles Allow Privilege Escalation |
| +0 | 124 | 124 | CWE-1357 | Apt Get Install Pin Version Not Defined, Unpinned Package Version i... |
| +0 | 2 | 2 | CWE-610 | Add Instead of Copy, Curl or Wget Instead of Add |
| +0 | 1 | 1 | CWE-923 | CNI Plugin Does Not Support Network Policies |
| -81 | 102 | 21 | CWE-710 | Bind Address Not Properly Set, Deployment Has No PodAntiAffinity, H... |

## Impact (per chart)

| Metric | Count |
|---|---:|
| Charts | 589 |
| Findings unlocked (after > before) | 535 |
| No change | 49 |
| Fewer (after < before) | 5 |
| Unscannable before (0 files) | 1 |
| After-scan error/timeout | 0 |

## Top 25 charts by findings unlocked

| Delta | Before | After | Chart |
|---:|---:|---:|---|
| +1216 | 14 | 1230 | `helm_clones_github__junghoon2__k8s-class__gitlab__gitlab-7.0.1` |
| +1057 | 32 | 1089 | `helm_clones_github__vexxhost__atmosphere__charts__loki` |
| +1010 | 25 | 1035 | `helm_clones_artifacthub__grafana__loki__production__helm__loki` |
| +992 | 4 | 996 | `helm_clones_github__ministryofjustice__cloud-platform-how-out-of-date-are-we__cloud-platform-reports-cronjobs` |
| +858 | 140 | 998 | `helm_clones_github__Above-Os__terminus-apps__otmoiclp` |
| +812 | 24 | 836 | `helm_clones_github__Above-Os__terminus-apps__mastodon` |
| +797 | 0 | 797 | `helm_clones_github__cmgoffena13__etl-watcher__watcher` |
| +663 | 37 | 700 | `helm_clones_github__Above-Os__terminus-apps__bisheng` |
| +463 | 395 | 858 | `helm_clones_github__Above-Os__terminus-apps__otmoicrelay` |
| +331 | 1 | 332 | `helm_clones_github__Above-Os__terminus-apps__langfuse` |
| +330 | 27 | 357 | `helm_clones_github__Bahmni__helm-umbrella-chart` |
| +326 | 0 | 326 | `helm_clones_github__Above-Os__terminus-apps__ragflow` |
| +289 | 19 | 308 | `helm_clones_github__Clark1992__ECK1__infra__k8s__charts__elasticsearch__cluster` |
| +274 | 2 | 276 | `helm_clones_github__Above-Os__terminus-apps__fastgpt` |
| +272 | 30 | 302 | `helm_clones_github__Above-Os__terminus-apps__twenty` |
| +261 | 15 | 276 | `helm_clones_github__Above-Os__terminus-apps__opennotebook` |
| +249 | 1 | 250 | `helm_clones_github__Above-Os__terminus-apps__affine` |
| +235 | 119 | 354 | `helm_clones_github__Above-Os__terminus-apps__firecrawl` |
| +224 | 4 | 228 | `helm_clones_github__Above-Os__terminus-apps__penpot` |
| +201 | 3 | 204 | `helm_clones_github__oracle-cne__catalog__charts__oci-ccm-1.30.0` |
| +199 | 17 | 216 | `helm_clones_github__Above-Os__terminus-apps__seatable` |
| +198 | 0 | 198 | `helm_clones_artifacthub__codefresh-io__venona__charts__cf-runtime` |
| +194 | 0 | 194 | `helm_clones_github__Above-Os__terminus-apps__merchant` |
| +188 | 42 | 230 | `helm_clones_github__Above-Os__terminus-apps__nemoclaw` |
| +159 | 0 | 159 | `helm_clones_github__oracle-cne__catalog__charts__ovirt-csi-driver-4.20.0` |
