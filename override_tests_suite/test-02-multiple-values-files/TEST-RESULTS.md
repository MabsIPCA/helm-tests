# Test 2: Multiple Values Files Precedence

**Description:** Tests precedence when using multiple -f flags with values files

**Generated:** 2026-02-26 23:03:37

## Summary

| Test Case | Description |
|-----------|-------------|
| [single-override](#single-override) | Single override file over default values |
| [two-overrides](#two-overrides) | Two override files, second takes precedence |
| [three-overrides](#three-overrides) | Three override files, last takes precedence |
| [reverse-order](#reverse-order) | Same files in reverse order |

## Test Case Details

### single-override

Single override file over default values

**Command:**

```bash
helm template test-multiple-values . -f values-override-1.yaml
```

**Values Comparison:**

| Field | Default | `values-override-1.yaml` | **Final** |
|-------|---------|----------|----------|
| `config.logLevel` | `"info"` | `"debug"` | `"debug"` |
| `database.host` | `"localhost"` | `"override1-db.local"` | `"override1-db.local"` |
| `database.port` | `5432` | `5432` | `5432` |
| `environment` | `"default"` | `"override-1"` | `"override-1"` |
| `replicas` | `1` | `2` | `2` |
| `version` | `"1.0.0"` | `"1.0.0"` | `"1.0.0"` |

---

### two-overrides

Two override files, second takes precedence

**Command:**

```bash
helm template test-multiple-values . -f values-override-1.yaml -f values-override-2.yaml
```

**Values Comparison:**

| Field | Default | `values-override-1.yaml` | `values-override-2.yaml` | **Final** |
|-------|---------|----------|----------|----------|
| `config.features.caching` | `false` | `false` | `true` | `true` |
| `config.logLevel` | `"info"` | `"debug"` | `"debug"` | `"debug"` |
| `config.timeout` | `30` | `30` | `60` | `60` |
| `database.host` | `"localhost"` | `"override1-db.local"` | `"override2-db.local"` | `"override2-db.local"` |
| `database.name` | `"default_db"` | `"override1_db"` | `"override1_db"` | `"override1_db"` |
| `database.port` | `5432` | `5432` | `5433` | `5433` |
| `environment` | `"default"` | `"override-1"` | `"override-2"` | `"override-2"` |
| `replicas` | `1` | `2` | `2` | `2` |
| `version` | `"1.0.0"` | `"1.0.0"` | `"2.0.0"` | `"2.0.0"` |

---

### three-overrides

Three override files, last takes precedence

**Command:**

```bash
helm template test-multiple-values . -f values-override-1.yaml -f values-override-2.yaml -f values-override-3.yaml
```

**Values Comparison:**

| Field | Default | `values-override-1.yaml` | `values-override-2.yaml` | `values-override-3.yaml` | **Final** |
|-------|---------|----------|----------|----------|----------|
| `config.features.caching` | `false` | `false` | `true` | `true` | `true` |
| `config.features.monitoring` | `false` | `false` | `false` | `true` | `true` |
| `config.logLevel` | `"info"` | `"debug"` | `"debug"` | `"debug"` | `"debug"` |
| `config.timeout` | `30` | `30` | `60` | `60` | `60` |
| `database.host` | `"localhost"` | `"override1-db.local"` | `"override2-db.local"` | `"override2-db.local"` | `"override2-db.local"` |
| `database.name` | `"default_db"` | `"override1_db"` | `"override1_db"` | `"override3_db"` | `"override3_db"` |
| `database.port` | `5432` | `5432` | `5433` | `5433` | `5433` |
| `environment` | `"default"` | `"override-1"` | `"override-2"` | `"override-3"` | `"override-3"` |
| `replicas` | `1` | `2` | `2` | `5` | `5` |
| `version` | `"1.0.0"` | `"1.0.0"` | `"2.0.0"` | `"2.0.0"` | `"2.0.0"` |

---

### reverse-order

Same files in reverse order

**Command:**

```bash
helm template test-multiple-values . -f values-override-3.yaml -f values-override-2.yaml -f values-override-1.yaml
```

**Values Comparison:**

| Field | Default | `values-override-3.yaml` | `values-override-2.yaml` | `values-override-1.yaml` | **Final** |
|-------|---------|----------|----------|----------|----------|
| `database.host` | `"localhost"` | `"localhost"` | `"override2-db.local"` | `"override1-db.local"` | `"override1-db.local"` |
| `database.name` | `"default_db"` | `"override3_db"` | `"override3_db"` | `"override1_db"` | `"override1_db"` |
| `environment` | `"default"` | `"override-3"` | `"override-2"` | `"override-1"` | `"override-1"` |
| `replicas` | `1` | `5` | `5` | `2` | `2` |

---

## Expected Override Behaviors

| Field | Behavior | Description |
|-------|----------|-------------|
| `precedence` | left-to-right | Later -f flags override earlier ones |
| `maps` | deep-merge | Maps are deeply merged across all files |
| `scalars` | last-wins | Scalar values from the last file win |
