# KICS `--experimental-helm-scan` — Before/After Report

_Baseline: the 589 fixer-resolved Helm charts, each scanned twice — **before** = flag off (KICS's normal fail-closed render), **after** = `--experimental-helm-scan` on (fixer injects placeholder values so unrenderable charts still render)._

## TL;DR

- Turning the flag on takes total findings from **4,272 to 33,368** — a **7.8x increase** (+681%).
- **535 of 589 charts** surfaced more findings; only **5** dropped, **49** were unchanged, and **0** errored.
- The uplift is broad, not concentrated: it spans **10 categories**, **69 distinct queries**, and **19 CWEs**.
- The few *decreases* are not regressions — they are noisy checks (e.g. `Metadata Label Is Invalid`) that fire against raw, unrendered template text when the flag is off and largely disappear once the chart renders into proper manifests.

## Severity

Most of the newly-surfaced issues are LOW/MEDIUM hygiene findings, but HIGH-severity findings still jumped **8.4x**.

| Severity | Before | After | Change |
|---|---:|---:|---:|
| CRITICAL | 0 | 0 | **+0** (n/a) |
| HIGH | 359 | 3,033 | **+2,674** (+745%) |
| MEDIUM | 1,715 | 12,728 | **+11,013** (+642%) |
| LOW | 1,994 | 15,792 | **+13,798** (+692%) |
| INFO | 204 | 1,815 | **+1,611** (+790%) |
| **Total** | **4,272** | **33,368** | **+29,096** (+681%) |

## Where the findings land (category)

| Category | Before | After | Added | Share of uplift |
|---|---:|---:|---:|---:|
| Insecure Configurations | 1,799 | 13,887 | +12,088 | 42% |
| Best Practices | 735 | 4,738 | +4,003 | 14% |
| Resource Management | 311 | 4,180 | +3,869 | 13% |
| Access Control | 475 | 3,944 | +3,469 | 12% |
| Availability | 293 | 1,924 | +1,631 | 6% |
| Build Process | 221 | 1,740 | +1,519 | 5% |
| Secret Management | 122 | 1,411 | +1,289 | 4% |
| Insecure Defaults | 171 | 1,308 | +1,137 | 4% |
| Networking and Firewall | 12 | 80 | +68 | 0% |
| Supply-Chain | 133 | 156 | +23 | 0% |

## Top 15 checks driving the increase

| Added | Before | After | Severity | Query | Category |
|---:|---:|---:|:--|:--|:--|
| +1,853 | 226 | 2079 | LOW | Missing AppArmor Profile | Access Control |
| +1,813 | 226 | 2039 | LOW | Image Without Digest | Insecure Configurations |
| +1,676 | 214 | 1890 | LOW | Image Pull Policy Of The Container Is Not Set To Always | Insecure Configurations |
| +1,645 | 221 | 1866 | MEDIUM | Seccomp Profile Is Not Configured | Insecure Configurations |
| +1,611 | 220 | 1831 | MEDIUM | Container Running With Low UID | Best Practices |
| +1,469 | 210 | 1679 | HIGH | Privilege Escalation Allowed | Insecure Configurations |
| +1,440 | 214 | 1654 | LOW | Root Container Not Mounted Read-only | Build Process |
| +1,299 | 212 | 1511 | LOW | No Drop Capabilities for Containers | Best Practices |
| +1,289 | 212 | 1501 | MEDIUM | NET_RAW Capabilities Not Being Dropped | Insecure Configurations |
| +1,138 | 170 | 1308 | MEDIUM | Service Account Token Automount Not Disabled | Insecure Defaults |
| +1,134 | 204 | 1338 | MEDIUM | Container Running As Root | Best Practices |
| +1,104 | 74 | 1178 | LOW | CPU Limits Not Set | Resource Management |
| +986 | 172 | 1158 | LOW | Pod or Container Without LimitRange | Insecure Configurations |
| +986 | 172 | 1158 | LOW | Pod or Container Without ResourceQuota | Insecure Configurations |
| +944 | 73 | 1017 | MEDIUM | Memory Limits Not Defined | Resource Management |

## Weakness classes (CWE), top 10

| Added | Before | After | CWE | Theme |
|---:|---:|---:|:--|:--|
| +7,556 | 1013 | 8569 | CWE-665 | Improper initialization / insecure config |
| +5,627 | 650 | 6277 | CWE-400 | Uncontrolled resource consumption |
| +2,926 | 463 | 3389 | CWE-269 | Improper privilege management |
| +2,844 | 486 | 3330 | CWE-754 | Improper check of unusual conditions |
| +2,745 | 424 | 3169 | CWE-1188 | Insecure default resource initialization |
| +2,556 | 281 | 2837 | CWE-284 | Improper access control |
| +2,214 | 253 | 2467 | CWE-668 | Exposure of resource to wrong sphere |
| +868 | 168 | 1036 | CWE-285 | Improper authorization |
| +613 | 32 | 645 | CWE-526 | Sensitive info in environment variables |
| +325 | 41 | 366 | CWE-200 | Exposure of sensitive information |

## Decreases (all of them) — noise removed, not coverage lost

| Change | Before | After | Query | Why |
|---:|---:|---:|:--|:--|
| -89 | 90 | 1 | Metadata Label Is Invalid | fires on raw template placeholders; valid once rendered |
| -18 | 65 | 47 | RBAC Wildcard In Rule | matched unrendered RBAC text; precise on real manifests |
| -1 | 1 | 0 | Service Account Name Undefined Or Empty | resolved to a real value after render |

## Biggest per-chart wins (top 15)

| Added | Before | After | Chart |
|---:|---:|---:|:--|
| +1,216 | 14 | 1230 | `gh:junghoon2/k8s-class/gitlab/gitlab-7.0.1` |
| +1,057 | 32 | 1089 | `gh:vexxhost/atmosphere/charts/loki` |
| +1,010 | 25 | 1035 | `ah:grafana/loki/production/helm/loki` |
| +992 | 4 | 996 | `gh:ministryofjustice/cloud-platform-how-out-of-date-are-we/cloud-platform-reports-cronjobs` |
| +858 | 140 | 998 | `gh:Above-Os/terminus-apps/otmoiclp` |
| +812 | 24 | 836 | `gh:Above-Os/terminus-apps/mastodon` |
| +797 | 0 | 797 | `gh:cmgoffena13/etl-watcher/watcher` |
| +663 | 37 | 700 | `gh:Above-Os/terminus-apps/bisheng` |
| +463 | 395 | 858 | `gh:Above-Os/terminus-apps/otmoicrelay` |
| +331 | 1 | 332 | `gh:Above-Os/terminus-apps/langfuse` |
| +330 | 27 | 357 | `gh:Bahmni/helm-umbrella-chart` |
| +326 | 0 | 326 | `gh:Above-Os/terminus-apps/ragflow` |
| +289 | 19 | 308 | `gh:Clark1992/ECK1/infra/k8s/charts/elasticsearch/cluster` |
| +274 | 2 | 276 | `gh:Above-Os/terminus-apps/fastgpt` |
| +272 | 30 | 302 | `gh:Above-Os/terminus-apps/twenty` |

---
_Full per-query / per-CWE tables and every chart row: `kics_beforeafter_summary.md` (and `.json`)._