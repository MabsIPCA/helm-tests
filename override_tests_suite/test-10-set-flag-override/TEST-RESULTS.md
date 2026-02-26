# Test 10: Set Flag Override

**Description:** Tests --set and --set-string flag behavior and precedence

**Generated:** 2026-02-26 23:44:29

## Summary

| Test Case | Description |
|-----------|-------------|
| [simple-set](#simple-set) | Simple --set flag override |
| [set-number](#set-number) | --set with numeric value |
| [set-boolean](#set-boolean) | --set with boolean value |
| [set-nested](#set-nested) | --set with nested path |
| [set-string-number](#set-string-number) | --set-string forces string type |
| [set-string-bool](#set-string-bool) | --set-string with boolean-like value |
| [set-over-file](#set-over-file) | --set overrides values file |
| [multiple-set](#multiple-set) | Multiple --set flags |
| [set-list-index](#set-list-index) | --set with list index creates a new list |
| [set-null](#set-null) | --set with null value sets actual null (key removed from output) |

## Test Case Details

### simple-set

Simple --set flag override

**Command:**

```bash
helm template test-set-flag . --set stringValue=set-value
```

**Values Comparison:**

| Field | Default | Override | Final |
|-------|---------|---------|-------|
| `stringValue` | `"default-string"` | `"set-value"` | `"set-value"` |

---

### set-number

--set with numeric value

**Command:**

```bash
helm template test-set-flag . --set numberValue=999
```

**Values Comparison:**

| Field | Default | Override | Final |
|-------|---------|---------|-------|
| `numberValue` | `100` | `999` | `999` |

---

### set-boolean

--set with boolean value

**Command:**

```bash
helm template test-set-flag . --set boolValue=true
```

**Values Comparison:**

| Field | Default | Override | Final |
|-------|---------|---------|-------|
| `boolValue` | `false` | `true` | `true` |

---

### set-nested

--set with nested path

**Command:**

```bash
helm template test-set-flag . --set nested.level1.value=set-nested-value --set nested.level1.number=999
```

**Values Comparison:**

| Field | Default | `--set nested.level1.value=set-nested-value` | `--set nested.level1.number=999` | **Final** |
|-------|---------|----------|----------|----------|
| `nested.level1.number` | `50` | `50` | `999` | `999` |
| `nested.level1.value` | `"nested-default"` | `"set-nested-value"` | `"set-nested-value"` | `"set-nested-value"` |

---

### set-string-number

--set-string forces string type

**Command:**

```bash
helm template test-set-flag . --set-string numberValue=123
```

**Values Comparison:**

| Field | Default | Override | Final |
|-------|---------|---------|-------|
| `numberValue` | `100` | `"123"` | `"123"` |

---

### set-string-bool

--set-string with boolean-like value

**Command:**

```bash
helm template test-set-flag . --set-string boolValue=true
```

**Values Comparison:**

| Field | Default | Override | Final |
|-------|---------|---------|-------|
| `boolValue` | `false` | `"true"` | `"true"` |

---

### set-over-file

--set overrides values file

**Command:**

```bash
helm template test-set-flag . -f values-base.yaml --set stringValue=set-wins
```

**Values Comparison:**

| Field | Default | `values-base.yaml` | `--set stringValue=set-wins` | **Final** |
|-------|---------|----------|----------|----------|
| `config.name` | `"default-config"` | `"from-values-file"` | `"from-values-file"` | `"from-values-file"` |
| `stringValue` | `"default-string"` | `"from-values-file"` | `"set-wins"` | `"set-wins"` |

---

### multiple-set

Multiple --set flags

**Command:**

```bash
helm template test-set-flag . --set stringValue=first --set numberValue=111 --set config.enabled=true
```

**Values Comparison:**

| Field | Default | `--set stringValue=first` | `--set numberValue=111` | `--set config.enabled=true` | **Final** |
|-------|---------|----------|----------|----------|----------|
| `config.enabled` | `false` | `false` | `false` | `true` | `true` |
| `numberValue` | `100` | `100` | `111` | `111` | `111` |
| `stringValue` | `"default-string"` | `"first"` | `"first"` | `"first"` | `"first"` |

---

### set-list-index

--set with list index creates a new list

**Command:**

```bash
helm template test-set-flag . --set list[0]=new-item1
```

**Values Comparison:**

| Field | Default | Override | Final |
|-------|---------|---------|-------|
| `list` | `["item1", "item2"]` | `["new-item1"]` | `["new-item1"]` |

---

### set-null

--set with null value sets actual null (key removed from output)

**Command:**

```bash
helm template test-set-flag . --set stringValue=null
```

**Values Comparison:**

| Field | Default | Override | Final |
|-------|---------|---------|-------|
| `stringValue` | `"default-string"` | _(not set)_ | _(not set)_ |

---

## Expected Override Behaviors

| Field | Behavior | Description |
|-------|----------|-------------|
| `--set` | type-inferred | --set infers type from value (number, boolean, string) |
| `--set-string` | always-string | --set-string always treats value as string |
| `precedence` | --set > -f | --set flags have higher precedence than values files |
| `nested-path` | dot-notation | Use dot notation for nested paths |
| `list-index` | bracket-notation | Use [index] notation for list items |
