# Test 9: Global Values Override

**Description:** Tests global values propagation and override behavior

**Generated:** 2026-02-26 23:44:26

## Summary

| Test Case | Description |
|-----------|-------------|
| [default-globals](#default-globals) | Default global values from parent |
| [override-globals](#override-globals) | Override global values |
| [add-globals](#add-globals) | Add new global values |
| [override-nested-globals](#override-nested-globals) | Override nested global values like labels |
| [chained-global-override](#chained-global-override) | Multiple global overrides in sequence |

## Test Case Details

### default-globals

Default global values from parent

**Command:**

```bash
helm template test-global-values .
```

**Default Values:**

| Field | Value |
|-------|-------|
| `global.config.debug` | `false` |
| `global.environment` | `"default"` |
| `global.imageTag` | `"latest"` |
| `global.region` | `"us-east-1"` |

---

### override-globals

Override global values

**Command:**

```bash
helm template test-global-values . -f values-global-prod.yaml
```

**Values Comparison:**

| Field | Default | Override | Final |
|-------|---------|---------|-------|
| `global.config.debug` | `false` | `true` | `true` |
| `global.config.logLevel` | `"info"` | `"info"` | `"info"` |
| `global.environment` | `"default"` | `"production"` | `"production"` |
| `global.region` | `"us-east-1"` | `"eu-west-1"` | `"eu-west-1"` |

---

### add-globals

Add new global values

**Command:**

```bash
helm template test-global-values . -f values-global-add.yaml
```

**Values Comparison:**

| Field | Default | Override | Final |
|-------|---------|---------|-------|
| `global.config.newConfig` | _(not set)_ | `"added-config"` | `"added-config"` |
| `global.environment` | `"default"` | `"default"` | `"default"` |
| `global.newValue` | _(not set)_ | `"added-via-override"` | `"added-via-override"` |

---

### override-nested-globals

Override nested global values like labels

**Command:**

```bash
helm template test-global-values . -f values-global-labels.yaml
```

**Values Comparison:**

| Field | Default | Override | Final |
|-------|---------|---------|-------|
| `global.labels.app` | `"test"` | `"test"` | `"test"` |
| `global.labels.newLabel` | _(not set)_ | `"added"` | `"added"` |
| `global.labels.team` | `"platform"` | `"devops"` | `"devops"` |

---

### chained-global-override

Multiple global overrides in sequence

**Command:**

```bash
helm template test-global-values . -f values-global-prod.yaml -f values-global-add.yaml
```

**Values Comparison:**

| Field | Default | `values-global-prod.yaml` | `values-global-add.yaml` | **Final** |
|-------|---------|----------|----------|----------|
| `global.environment` | `"default"` | `"production"` | `"production"` | `"production"` |
| `global.newValue` | _(not set)_ | _(not set)_ | `"added-via-override"` | `"added-via-override"` |
| `global.region` | `"us-east-1"` | `"eu-west-1"` | `"eu-west-1"` | `"eu-west-1"` |

---

## Expected Override Behaviors

| Field | Behavior | Description |
|-------|----------|-------------|
| `global` | propagated | Global values propagate to all subcharts |
| `global-merge` | deep-merge | Global values are deeply merged |
| `child-global` | overridden | Parent global values override child global defaults |
