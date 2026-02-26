# Test 7: Null Value Override

**Description:** Tests behavior when overriding null values and setting values to null

**Generated:** 2026-02-26 23:03:42

## Summary

| Test Case | Description |
|-----------|-------------|
| [null-to-value](#null-to-value) | Override null values with actual values |
| [value-to-null](#value-to-null) | Override values to null |
| [nested-null](#nested-null) | Override nested null values |
| [round-trip](#round-trip) | Set to value then back to null - values set in first file remain if not overridden |

## Test Case Details

### null-to-value

Override null values with actual values

**Command:**

```bash
helm template test-null-override . -f values-null-to-value.yaml
```

**Values Comparison:**

| Field | Default | `values-null-to-value.yaml` | **Final** |
|-------|---------|----------|----------|
| `nullBool` | `null` | `true` | `true` |
| `nullList` | `null` | `["new-item"]` | `["new-item"]` |
| `nullMap.newKey` | _(not set)_ | `"new-value"` | `"new-value"` |
| `nullNumber` | `null` | `100` | `100` |
| `nullString` | `null` | `"now-has-string"` | `"now-has-string"` |

---

### value-to-null

Override values to null

**Command:**

```bash
helm template test-null-override . -f values-value-to-null.yaml
```

**Values Comparison:**

| Field | Default | `values-value-to-null.yaml` | **Final** |
|-------|---------|----------|----------|
| `boolValue` | `true` | _(not set)_ | _(not set)_ |
| `listValue` | `["item1", "item2"]` | _(not set)_ | _(not set)_ |
| `mapValue` | `{key1: "value1", key2: "value2"}` | _(not set)_ | _(not set)_ |
| `numberValue` | `42` | _(not set)_ | _(not set)_ |
| `stringValue` | `"has-value"` | _(not set)_ | _(not set)_ |

---

### nested-null

Override nested null values

**Command:**

```bash
helm template test-null-override . -f values-nested-null.yaml
```

**Values Comparison:**

| Field | Default | `values-nested-null.yaml` | **Final** |
|-------|---------|----------|----------|
| `nested.deepNested.alsoHasValue` | `"deep-exists"` | `"deep-exists"` | `"deep-exists"` |
| `nested.deepNested.alsoNull` | `null` | `"also-has-value"` | `"also-has-value"` |
| `nested.hasValue` | `"exists"` | `"exists"` | `"exists"` |
| `nested.nullChild` | `null` | `"now-has-value"` | `"now-has-value"` |

---

### round-trip

Set to value then back to null - values set in first file remain if not overridden

**Command:**

```bash
helm template test-null-override . -f values-null-to-value.yaml -f values-value-to-null.yaml
```

**Values Comparison:**

| Field | Default | `values-null-to-value.yaml` | `values-value-to-null.yaml` | **Final** |
|-------|---------|----------|----------|----------|
| `nullNumber` | `null` | `100` | `100` | `100` |
| `nullString` | `null` | `"now-has-string"` | `"now-has-string"` | `"now-has-string"` |
| `stringValue` | `"has-value"` | `"has-value"` | _(not set)_ | _(not set)_ |

---

## Expected Override Behaviors

| Field | Behavior | Description |
|-------|----------|-------------|
| `null-to-value` | replaced | Null values can be overridden with any type |
| `value-to-null` | replaced | Any value can be set to null |
| `nested-null` | merged | Nested structures with null are properly merged |
