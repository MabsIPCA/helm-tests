# Test 5: Nil Pointer Dereference Scenarios

**Generated:** 2026-02-20 00:44:01

**Total:** 90 | **Passed:** 60 (66.7%) | **Failed:** 30 (33.3%)

## Test Matrix

| Function | nilMap | nilNested | emptyMap | partialMap | nilInList | nilInDict |
|----------|------|------|------|------|------|------|
| `nil-map-access` | ✅ | ✅ | ✅ | ✅ | ❌ | ✅ |
| `nil-nested-map` | ✅ | ❌ | ❌ | ❌ | ❌ | ❌ |
| `nil-deep-nest` | ✅ | ❌ | ❌ | ❌ | ❌ | ❌ |
| `nil-list-index` | ❌ | ❌ | ❌ | ❌ | ✅ | ❌ |
| `nil-range` | ✅ | ❌ | ✅ | ❌ | ❌ | ❌ |
| `nil-with-default` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `nil-index` | ❌ | ✅ | ✅ | ✅ | ❌ | ✅ |
| `nil-pluck` | ✅ | ✅ | ✅ | ✅ | ❌ | ✅ |
| `nil-keys` | ✅ | ✅ | ✅ | ✅ | ❌ | ✅ |
| `nil-values` | ✅ | ✅ | ✅ | ✅ | ❌ | ✅ |
| `nil-pick` | ✅ | ✅ | ✅ | ✅ | ❌ | ✅ |
| `nil-omit` | ✅ | ✅ | ✅ | ✅ | ❌ | ✅ |
| `nil-merge` | ✅ | ✅ | ✅ | ✅ | ❌ | ✅ |
| `nil-get` | ✅ | ✅ | ✅ | ✅ | ❌ | ✅ |
| `nil-hasKey` | ✅ | ✅ | ✅ | ✅ | ❌ | ✅ |

## Failure Details

| Function | nilMap | nilNested | emptyMap | partialMap | nilInList | nilInDict |
|----------|------|------|------|------|------|------|
| `nil-map-access` | - | - | - | - | Error: template: test-nil-pointer/templates/nil-pointer-test... | - |
| `nil-nested-map` | - | nil pointer | nil pointer | nil pointer | Error: template: test-nil-pointer/templates/nil-pointer-test... | nil pointer |
| `nil-deep-nest` | - | nil pointer | nil pointer | nil pointer | Error: template: test-nil-pointer/templates/nil-pointer-test... | nil pointer |
| `nil-list-index` | error calling | error calling | error calling | error calling | - | error calling |
| `nil-range` | - | nil pointer | - | Error: template: test-nil-pointer/templates/nil-pointer-test... | Error: template: test-nil-pointer/templates/nil-pointer-test... | Error: template: test-nil-pointer/templates/nil-pointer-test... |
| `nil-index` | error calling | - | - | - | cannot index | - |
| `nil-pluck` | - | - | - | - | wrong type | - |
| `nil-keys` | - | - | - | - | wrong type | - |
| `nil-values` | - | - | - | - | wrong type | - |
| `nil-pick` | - | - | - | - | wrong type | - |
| `nil-omit` | - | - | - | - | wrong type | - |
| `nil-merge` | - | - | - | - | wrong type | - |
| `nil-get` | - | - | - | - | wrong type | - |
| `nil-hasKey` | - | - | - | - | wrong type | - |
