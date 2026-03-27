# Helm Dependency Build Failures

Total dependency failures: **2**

| # | Repository | Chart Path | Error |
|---|------------|------------|-------|
| 1 | [liranme/redisinsight-secure](https://github.com/liranme/redisinsight-secure) | `cloned\liranme__redisinsight-secure\helm\redisinsight-secure` | Error: no repository definition for https://oauth2-proxy.github.io/manifests. Please add the missing repos via 'helm repo add' |
| 2 | [helm/charts](https://github.com/helm/charts) | `cloned\helm__charts\incubator\distribution` | manager.go:122: warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash... Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/, https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add' |

## Full Error Details

### `cloned\liranme__redisinsight-secure\helm\redisinsight-secure` — liranme/redisinsight-secure

```
Error: no repository definition for https://oauth2-proxy.github.io/manifests. Please add the missing repos via 'helm repo add'
```

### `cloned\helm__charts\incubator\distribution` — helm/charts

```
manager.go:122: warning: a valid Helm v3 hash was not found. Checking against Helm v2 hash...
Error: no repository definition for https://kubernetes-charts.storage.googleapis.com/, https://kubernetes-charts.storage.googleapis.com/. Please add the missing repos via 'helm repo add'
```

