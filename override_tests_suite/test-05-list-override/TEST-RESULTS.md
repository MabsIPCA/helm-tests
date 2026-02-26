# Test 5: List Override Behavior

**Description:** Tests that lists are completely replaced, not merged

**Generated:** 2026-02-26 23:03:41

## Summary

| Test Case | Description |
|-----------|-------------|
| [fewer-items](#fewer-items) | Override list with fewer items |
| [more-items](#more-items) | Override list with more items |
| [empty-lists](#empty-lists) | Override with empty lists |
| [nested-list](#nested-list) | Override nested lists |
| [sequential-list-override](#sequential-list-override) | Multiple files overriding same list |

## Test Case Details

### fewer-items

Override list with fewer items

**Command:**

```bash
helm template test-list-override . -f values-fewer-items.yaml
```

**Values Comparison:**

| Field | Default | `values-fewer-items.yaml` | **Final** |
|-------|---------|----------|----------|
| `numberList` | `[1, 2, 3]` | `[100, 200]` | `[100, 200]` |
| `simpleList` | `["item1", "item2", "item3"]` | `["override1", "override2"]` | `["override1", "override2"]` |

---

### more-items

Override list with more items

**Command:**

```bash
helm template test-list-override . -f values-more-items.yaml
```

**Values Comparison:**

| Field | Default | `values-more-items.yaml` | **Final** |
|-------|---------|----------|----------|
| `simpleList` | `["item1", "item2", "item3"]` | `["new1", "new2", "new3", "new4", "new5"]` | `["new1", "new2", "new3", "new4", "new5"]` |

---

### empty-lists

Override with empty lists

**Command:**

```bash
helm template test-list-override . -f values-empty-lists.yaml
```

**Values Comparison:**

| Field | Default | `values-empty-lists.yaml` | **Final** |
|-------|---------|----------|----------|
| `numberList` | `[1, 2, 3]` | `[]` | `[]` |
| `objectList` | `[{enabled: true, name: "obj1", value: "value1"}, {enabled: false, name: "obj2", value: "value2"}]` | `[]` | `[]` |
| `simpleList` | `["item1", "item2", "item3"]` | `[]` | `[]` |

---

### nested-list

Override nested lists

**Command:**

```bash
helm template test-list-override . -f values-nested-list.yaml
```

**Values Comparison:**

| Field | Default | `values-nested-list.yaml` | **Final** |
|-------|---------|----------|----------|
| `nestedLists.level1` | `["nested1", "nested2"]` | `["replaced1", "replaced2", "replaced3"]` | `["replaced1", "replaced2", "replaced3"]` |
| `nestedLists.level2` | _(not set)_ | `["new-nested"]` | `["new-nested"]` |

---

### sequential-list-override

Multiple files overriding same list

**Command:**

```bash
helm template test-list-override . -f values-fewer-items.yaml -f values-more-items.yaml
```

**Values Comparison:**

| Field | Default | `values-fewer-items.yaml` | `values-more-items.yaml` | **Final** |
|-------|---------|----------|----------|----------|
| `simpleList` | `["item1", "item2", "item3"]` | `["override1", "override2"]` | `["new1", "new2", "new3", "new4", "new5"]` | `["new1", "new2", "new3", "new4", "new5"]` |

---

## Expected Override Behaviors

| Field | Behavior | Description |
|-------|----------|-------------|
| `lists` | full-replacement | Lists are completely replaced, never merged |
| `list-length` | no-constraint | Override list can have any length |
| `empty-list` | valid | Empty list [] is a valid override |
