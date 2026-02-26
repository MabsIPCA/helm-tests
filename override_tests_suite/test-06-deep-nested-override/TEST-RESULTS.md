# Test 6: Deep Nested Structure Override

**Description:** Tests override behavior for deeply nested structures

**Generated:** 2026-02-26 23:44:23

## Summary

| Test Case | Description |
|-----------|-------------|
| [deepest-level](#deepest-level) | Override only the deepest nested level |
| [multi-level](#multi-level) | Override at multiple nesting levels |
| [complex-nested](#complex-nested) | Override complex nested structure with mixed types |

## Test Case Details

### deepest-level

Override only the deepest nested level

**Command:**

```bash
helm template test-deep-nested . -f values-deep-override.yaml
```

**Values Comparison:**

| Field | Default | Override | Final |
|-------|---------|---------|-------|
| `level1.level2.level3.level4.level5.deepest` | `"original-deepest"` | `"overridden-deepest"` | `"overridden-deepest"` |
| `level1.level2.level3.level4.level5.newKey` | _(not set)_ | `"added-at-deepest"` | `"added-at-deepest"` |
| `level1.level2.level3.level4.level5.value` | `"l5-value"` | `"l5-value"` | `"l5-value"` |
| `level1.value` | `"l1-value"` | `"l1-value"` | `"l1-value"` |

---

### multi-level

Override at multiple nesting levels

**Command:**

```bash
helm template test-deep-nested . -f values-multi-level.yaml
```

**Values Comparison:**

| Field | Default | Override | Final |
|-------|---------|---------|-------|
| `level1.level2.level3.level4.level5.value` | `"l5-value"` | `"l5-value"` | `"l5-value"` |
| `level1.level2.level3.level4.value` | `"l4-value"` | `"l4-overridden"` | `"l4-overridden"` |
| `level1.level2.level3.value` | `"l3-value"` | `"l3-value"` | `"l3-value"` |
| `level1.level2.value` | `"l2-value"` | `"l2-overridden"` | `"l2-overridden"` |
| `level1.value` | `"l1-value"` | `"l1-overridden"` | `"l1-overridden"` |

---

### complex-nested

Override complex nested structure with mixed types

**Command:**

```bash
helm template test-deep-nested . -f values-complex-override.yaml
```

**Values Comparison:**

| Field | Default | Override | Final |
|-------|---------|---------|-------|
| `complex.name` | `"complex-structure"` | `"complex-structure"` | `"complex-structure"` |
| `complex.settings.cache.enabled` | `false` | `true` | `true` |
| `complex.settings.cache.ttl` | `3600` | `3600` | `3600` |
| `complex.settings.database.primary.credentials.password` | `"secret"` | `"new-secret"` | `"new-secret"` |
| `complex.settings.database.primary.credentials.username` | `"admin"` | `"admin"` | `"admin"` |
| `complex.settings.database.primary.host` | `"primary-db"` | `"new-primary-db"` | `"new-primary-db"` |
| `complex.settings.database.primary.port` | `5432` | `5432` | `5432` |

---

## Expected Override Behaviors

| Field | Behavior | Description |
|-------|----------|-------------|
| `nested-maps` | deep-merge | Maps are merged at all nesting levels |
| `partial-path` | preserved | Unspecified intermediate keys are preserved |
| `new-keys` | added | New keys can be added at any depth |
