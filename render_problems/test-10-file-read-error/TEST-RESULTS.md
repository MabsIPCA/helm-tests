# Test 10: File and Path Function Errors

**Generated:** 2026-02-19 19:53:00

**Total:** 42 | **Passed:** 6 (14.3%) | **Failed:** 36 (85.7%)

## Test Matrix

| Function | nilValue | stringValue | numberValue | boolValue | listValue | dictValue |
|----------|------|------|------|------|------|------|
| `base` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `dir` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `clean` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `ext` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `isAbs` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `globRequired` | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ |
| `asConfigEmpty` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |

## Failure Details

| Function | nilValue | stringValue | numberValue | boolValue | listValue | dictValue |
|----------|------|------|------|------|------|------|
| `base` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `dir` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `clean` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `ext` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `isAbs` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `globRequired` | invalid value; expected string | No files matched the pattern | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `asConfigEmpty` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
