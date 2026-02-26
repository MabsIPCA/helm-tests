# Test 3: Multiple Values Files with Default Values

**Description:** Tests interaction between chart defaults and external values files

**Generated:** 2026-02-26 23:03:38

## Summary

| Test Case | Description |
|-----------|-------------|
| [chart-defaults-only](#chart-defaults-only) | Only chart default values.yaml |
| [external-defaults](#external-defaults) | External defaults over chart defaults |
| [external-then-dev](#external-then-dev) | External defaults then dev override |
| [dev-only](#dev-only) | Dev override directly on chart defaults |
| [external-then-prod](#external-then-prod) | External defaults then prod override |

## Test Case Details

### chart-defaults-only

Only chart default values.yaml

**Command:**

```bash
helm template test-values-with-defaults .
```

**Default Values:**

| Field | Value |
|-------|-------|
| `appName` | `"chart-default"` |
| `chartOnlyValue` | `"exists-only-in-chart"` |
| `environment` | `"default"` |
| `replicas` | `1` |

---

### external-defaults

External defaults over chart defaults

**Command:**

```bash
helm template test-values-with-defaults . -f values-external-defaults.yaml
```

**Values Comparison:**

| Field | Default | `values-external-defaults.yaml` | **Final** |
|-------|---------|----------|----------|
| `appName` | `"chart-default"` | `"external-default"` | `"external-default"` |
| `chartOnlyValue` | `"exists-only-in-chart"` | `"exists-only-in-chart"` | `"exists-only-in-chart"` |
| `database.host` | `"chart-default-db"` | `"external-default-db"` | `"external-default-db"` |
| `database.port` | `5432` | `5432` | `5432` |
| `enabled` | `false` | `true` | `true` |
| `environment` | `"default"` | `"base"` | `"base"` |
| `externalOnlyValue` | _(not set)_ | `"exists-only-in-external"` | `"exists-only-in-external"` |
| `replicas` | `1` | `2` | `2` |

---

### external-then-dev

External defaults then dev override

**Command:**

```bash
helm template test-values-with-defaults . -f values-external-defaults.yaml -f values-dev.yaml
```

**Values Comparison:**

| Field | Default | `values-external-defaults.yaml` | `values-dev.yaml` | **Final** |
|-------|---------|----------|----------|----------|
| `appName` | `"chart-default"` | `"external-default"` | `"external-default"` | `"external-default"` |
| `chartOnlyValue` | `"exists-only-in-chart"` | `"exists-only-in-chart"` | `"exists-only-in-chart"` | `"exists-only-in-chart"` |
| `database.host` | `"chart-default-db"` | `"external-default-db"` | `"dev-db.local"` | `"dev-db.local"` |
| `database.name` | _(not set)_ | `"external_db"` | `"dev_db"` | `"dev_db"` |
| `environment` | `"default"` | `"base"` | `"development"` | `"development"` |
| `externalOnlyValue` | _(not set)_ | `"exists-only-in-external"` | `"exists-only-in-external"` | `"exists-only-in-external"` |
| `replicas` | `1` | `2` | `3` | `3` |
| `resources.requests.cpu` | `"100m"` | `"100m"` | `"200m"` | `"200m"` |

---

### dev-only

Dev override directly on chart defaults

**Command:**

```bash
helm template test-values-with-defaults . -f values-dev.yaml
```

**Values Comparison:**

| Field | Default | `values-dev.yaml` | **Final** |
|-------|---------|----------|----------|
| `appName` | `"chart-default"` | `"chart-default"` | `"chart-default"` |
| `chartOnlyValue` | `"exists-only-in-chart"` | `"exists-only-in-chart"` | `"exists-only-in-chart"` |
| `database.host` | `"chart-default-db"` | `"dev-db.local"` | `"dev-db.local"` |
| `environment` | `"default"` | `"development"` | `"development"` |
| `replicas` | `1` | `3` | `3` |

---

### external-then-prod

External defaults then prod override

**Command:**

```bash
helm template test-values-with-defaults . -f values-external-defaults.yaml -f values-prod.yaml
```

**Values Comparison:**

| Field | Default | `values-external-defaults.yaml` | `values-prod.yaml` | **Final** |
|-------|---------|----------|----------|----------|
| `chartOnlyValue` | `"exists-only-in-chart"` | `"exists-only-in-chart"` | `"exists-only-in-chart"` | `"exists-only-in-chart"` |
| `database.host` | `"chart-default-db"` | `"external-default-db"` | `"prod-db.example.com"` | `"prod-db.example.com"` |
| `database.port` | `5432` | `5432` | `5433` | `5433` |
| `database.username` | `"chart_user"` | `"chart_user"` | `"prod_user"` | `"prod_user"` |
| `environment` | `"default"` | `"base"` | `"production"` | `"production"` |
| `externalOnlyValue` | _(not set)_ | `"exists-only-in-external"` | `"exists-only-in-external"` | `"exists-only-in-external"` |
| `replicas` | `1` | `2` | `10` | `10` |

---

## Expected Override Behaviors

| Field | Behavior | Description |
|-------|----------|-------------|
| `chart values.yaml` | base | Chart values.yaml is always the base layer |
| `external files` | overlay | External files are merged on top of chart defaults |
| `unique keys` | preserved | Keys unique to each layer are preserved through merging |
