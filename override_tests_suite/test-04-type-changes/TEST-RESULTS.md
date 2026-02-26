# Test 4: Type Changes During Override

**Description:** Tests behavior when value types change during override

**Generated:** 2026-02-26 23:03:39

## Summary

| Test Case | Description |
|-----------|-------------|
| [scalar-type-changes](#scalar-type-changes) | Change types of scalar values |
| [collection-type-changes](#collection-type-changes) | Change between list and map types |
| [null-override-all](#null-override-all) | Override all values to null |
| [chained-type-changes](#chained-type-changes) | Multiple type changes in sequence |

## Test Case Details

### scalar-type-changes

Change types of scalar values

**Command:**

```bash
helm template test-type-changes . -f values-type-change-1.yaml
```

**Values Comparison:**

| Field | Default | `values-type-change-1.yaml` | **Final** |
|-------|---------|----------|----------|
| `booleanValue` | `true` | `"yes"` | `"yes"` |
| `integerValue` | `42` | `"forty-two"` | `"forty-two"` |
| `nullValue.key` | _(not set)_ | `"now-has-value"` | `"now-has-value"` |
| `stringValue` | `"original-string"` | `12345` | `12345` |

---

### collection-type-changes

Change between list and map types

**Command:**

```bash
helm template test-type-changes . -f values-type-change-2.yaml
```

**Values Comparison:**

| Field | Default | `values-type-change-2.yaml` | **Final** |
|-------|---------|----------|----------|
| `listValue.first` | _(not set)_ | `"item1"` | `"item1"` |
| `listValue.second` | _(not set)_ | `"item2"` | `"item2"` |
| `mapValue` | `{key1: "value1", key2: "value2"}` | `["key1-value1", "key2-value2"]` | `["key1-value1", "key2-value2"]` |
| `nestedMap` | `{level1: {level2: "nested-value"}}` | `"flattened-value"` | `"flattened-value"` |

---

### null-override-all

Override all values to null

**Command:**

```bash
helm template test-type-changes . -f values-null-override.yaml
```

**Values Comparison:**

| Field | Default | `values-null-override.yaml` | **Final** |
|-------|---------|----------|----------|
| `booleanValue` | `true` | _(not set)_ | _(not set)_ |
| `floatValue` | `3.14` | _(not set)_ | _(not set)_ |
| `integerValue` | `42` | _(not set)_ | _(not set)_ |
| `listValue` | `["item1", "item2"]` | _(not set)_ | _(not set)_ |
| `mapValue` | `{key1: "value1", key2: "value2"}` | _(not set)_ | _(not set)_ |
| `stringValue` | `"original-string"` | _(not set)_ | _(not set)_ |

---

### chained-type-changes

Multiple type changes in sequence

**Command:**

```bash
helm template test-type-changes . -f values-type-change-1.yaml -f values-null-override.yaml
```

**Values Comparison:**

| Field | Default | `values-type-change-1.yaml` | `values-null-override.yaml` | **Final** |
|-------|---------|----------|----------|----------|
| `booleanValue` | `true` | `"yes"` | _(not set)_ | _(not set)_ |
| `integerValue` | `42` | `"forty-two"` | _(not set)_ | _(not set)_ |
| `stringValue` | `"original-string"` | `12345` | _(not set)_ | _(not set)_ |

---

## Expected Override Behaviors

| Field | Behavior | Description |
|-------|----------|-------------|
| `scalar-to-scalar` | replaced | Scalar type changes are allowed and replace completely |
| `list-to-map` | replaced | List can be replaced with map |
| `map-to-list` | replaced | Map can be replaced with list |
| `any-to-null` | replaced | Any type can be set to null |
| `null-to-any` | replaced | Null can be replaced with any type |
