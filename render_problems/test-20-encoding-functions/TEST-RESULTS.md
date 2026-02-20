# Test 20: Encoding Function Errors

**Generated:** 2026-02-20 01:48:14

**Total:** 24 | **Passed:** 4 (16.7%) | **Failed:** 20 (83.3%)

## Test Matrix

| Function | nilValue | stringValue | numberValue | boolValue | listValue | dictValue |
|----------|------|------|------|------|------|------|
| `b64enc` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `b64dec` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `b32enc` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `b32dec` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |

## Failure Details

| Function | nilValue | stringValue | numberValue | boolValue | listValue | dictValue |
|----------|------|------|------|------|------|------|
| `b64enc` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `b64dec` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `b32enc` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `b32dec` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
