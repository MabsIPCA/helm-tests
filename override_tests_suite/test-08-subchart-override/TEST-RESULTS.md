# Test 8: Subchart Override

**Description:** Tests subchart values override from parent chart and external files

**Generated:** 2026-02-26 23:44:25

## Summary

| Test Case | Description |
|-----------|-------------|
| [parent-defaults](#parent-defaults) | Parent chart overrides subchart defaults |
| [subchart-override](#subchart-override) | External file overrides subchart values |
| [global-override](#global-override) | Override global values |
| [combined-override](#combined-override) | Multiple overrides together |

## Test Case Details

### parent-defaults

Parent chart overrides subchart defaults

**Command:**

```bash
helm template test-subchart-override .
```

**Default Values:**

| Field | Value |
|-------|-------|
| `global.environment` | `"development"` |
| `parentName` | `"parent-chart"` |
| `subchart1.config.enabled` | `true` |
| `subchart1.name` | `"subchart1-from-parent"` |
| `subchart1.replicas` | `3` |
| `subchart2.name` | `"subchart2-from-parent"` |
| `subchart2.port` | `9090` |

---

### subchart-override

External file overrides subchart values

**Command:**

```bash
helm template test-subchart-override . -f values-subchart-override.yaml
```

**Values Comparison:**

| Field | Default | Override | Final |
|-------|---------|---------|-------|
| `subchart1.config.timeout` | `30` | `60` | `60` |
| `subchart1.name` | `"subchart1-from-parent"` | `"subchart1-external-override"` | `"subchart1-external-override"` |
| `subchart1.replicas` | `3` | `5` | `5` |
| `subchart2.database.host` | `"localhost"` | `"external-db.local"` | `"external-db.local"` |
| `subchart2.name` | `"subchart2-from-parent"` | `"subchart2-external-override"` | `"subchart2-external-override"` |

---

### global-override

Override global values

**Command:**

```bash
helm template test-subchart-override . -f values-global-override.yaml
```

**Values Comparison:**

| Field | Default | Override | Final |
|-------|---------|---------|-------|
| `global.domain` | `"example.com"` | `"prod.example.com"` | `"prod.example.com"` |
| `global.environment` | `"development"` | `"production"` | `"production"` |
| `global.newGlobal` | _(not set)_ | `"added-global"` | `"added-global"` |

---

### combined-override

Multiple overrides together

**Command:**

```bash
helm template test-subchart-override . -f values-subchart-override.yaml -f values-combined-override.yaml
```

**Values Comparison:**

| Field | Default | `values-subchart-override.yaml` | `values-combined-override.yaml` | **Final** |
|-------|---------|----------|----------|----------|
| `global.environment` | `"development"` | `"development"` | `"staging"` | `"staging"` |
| `subchart1.replicas` | `3` | `5` | `10` | `10` |
| `subchart2.port` | `9090` | `9090` | `443` | `443` |

---

## Expected Override Behaviors

| Field | Behavior | Description |
|-------|----------|-------------|
| `subchart-namespace` | prefixed | Subchart values are under subchart name key |
| `global` | propagated | Global values are accessible by all subcharts |
| `precedence` | external > parent > subchart | External files override parent which overrides subchart defaults |
