# Test 1: Basic Data Types Override

**Description:** Tests override behavior for all basic data types: strings, integers, floats, booleans, null, lists, and maps

**Generated:** 2026-02-26 23:03:36

## Summary

| Test Case | Description |
|-----------|-------------|
| [default-values](#default-values) | Verify default values without any override |
| [string-override](#string-override) | Override string values |
| [number-override](#number-override) | Override numeric values (integers and floats) |
| [boolean-override](#boolean-override) | Override boolean values |
| [list-override](#list-override) | Override list values (lists are replaced, not merged) |
| [map-override](#map-override) | Override map values (maps are merged) |

## Test Case Details

### default-values

Verify default values without any override

**Command:**

```bash
helm template test-basic-data-types .
```

**Default Values:**

| Field | Value |
|-------|-------|
| `boolFalse` | `false` |
| `boolTrue` | `true` |
| `floatValue` | `3.14` |
| `integerValue` | `100` |
| `nullValue` | `null` |
| `stringValue` | `"default-string"` |

---

### string-override

Override string values

**Command:**

```bash
helm template test-basic-data-types . -f values-string-override.yaml
```

**Values Comparison:**

| Field | Default | `values-string-override.yaml` | **Final** |
|-------|---------|----------|----------|
| `emptyString` | `""` | `"no-longer-empty"` | `"no-longer-empty"` |
| `nullValue` | `null` | `"was-null"` | `"was-null"` |
| `simpleMap.key4` | _(not set)_ | `"new-key"` | `"new-key"` |
| `stringValue` | `"default-string"` | `"overridden-string"` | `"overridden-string"` |

---

### number-override

Override numeric values (integers and floats)

**Command:**

```bash
helm template test-basic-data-types . -f values-number-override.yaml
```

**Values Comparison:**

| Field | Default | `values-number-override.yaml` | **Final** |
|-------|---------|----------|----------|
| `floatValue` | `3.14` | `99.99` | `99.99` |
| `integerValue` | `100` | `999` | `999` |
| `negativeInteger` | `-50` | `100` | `100` |
| `zeroInteger` | `0` | `42` | `42` |

---

### boolean-override

Override boolean values

**Command:**

```bash
helm template test-basic-data-types . -f values-boolean-override.yaml
```

**Values Comparison:**

| Field | Default | `values-boolean-override.yaml` | **Final** |
|-------|---------|----------|----------|
| `boolFalse` | `false` | `true` | `true` |
| `boolTrue` | `true` | `false` | `false` |
| `config.enabled` | `false` | `true` | `true` |

---

### list-override

Override list values (lists are replaced, not merged)

**Command:**

```bash
helm template test-basic-data-types . -f values-list-override.yaml
```

**Values Comparison:**

| Field | Default | `values-list-override.yaml` | **Final** |
|-------|---------|----------|----------|
| `mixedList` | `["string", 123, true, null]` | `["all", "strings", "now"]` | `["all", "strings", "now"]` |
| `simpleList` | `["item1", "item2", "item3"]` | `["override1", "override2"]` | `["override1", "override2"]` |

---

### map-override

Override map values (maps are merged)

**Command:**

```bash
helm template test-basic-data-types . -f values-map-override.yaml
```

**Values Comparison:**

| Field | Default | `values-map-override.yaml` | **Final** |
|-------|---------|----------|----------|
| `nestedMap.level1.level2.level3` | `"deep-value"` | `"overridden-deep"` | `"overridden-deep"` |
| `nestedMap.level1.level2.newLevel3` | _(not set)_ | `"added-nested"` | `"added-nested"` |
| `nestedMap.level1.newLevel2` | _(not set)_ | `"added-at-level2"` | `"added-at-level2"` |
| `simpleMap.key1` | `"value1"` | `"new-value1"` | `"new-value1"` |
| `simpleMap.key2` | `"value2"` | `"value2"` | `"value2"` |
| `simpleMap.newKey` | _(not set)_ | `"added-key"` | `"added-key"` |

---

## Expected Override Behaviors

| Field | Behavior | Description |
|-------|----------|-------------|
| `strings` | replaced | String values are completely replaced |
| `numbers` | replaced | Numeric values are completely replaced |
| `booleans` | replaced | Boolean values are completely replaced |
| `lists` | replaced | Lists are completely replaced, not merged |
| `maps` | merged | Maps are recursively merged |
| `null` | replaced | Null can be overridden with any type |
