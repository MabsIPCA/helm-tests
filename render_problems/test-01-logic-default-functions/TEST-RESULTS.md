# Test 1: Logic and Default Functions

**Generated:** 2026-02-20 00:44:00

**Total:** 90 | **Passed:** 50 (55.6%) | **Failed:** 40 (44.4%)

## Test Matrix

| Function | nilValue | stringValue | numberValue | boolValue | listValue | dictValue |
|----------|------|------|------|------|------|------|
| `and` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `or` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `not` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `eq` | ✅ | ❌ | ✅ | ❌ | ❌ | ❌ |
| `ne` | ✅ | ❌ | ✅ | ❌ | ❌ | ❌ |
| `lt` | ❌ | ❌ | ✅ | ❌ | ❌ | ❌ |
| `le` | ❌ | ❌ | ✅ | ❌ | ❌ | ❌ |
| `gt` | ❌ | ❌ | ✅ | ❌ | ❌ | ❌ |
| `ge` | ❌ | ❌ | ✅ | ❌ | ❌ | ❌ |
| `default` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `required` | ❌ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `empty` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `fail` | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ |
| `coalesce` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `ternary` | ❌ | ❌ | ❌ | ✅ | ❌ | ❌ |

## Failure Details

| Function | nilValue | stringValue | numberValue | boolValue | listValue | dictValue |
|----------|------|------|------|------|------|------|
| `eq` | - | error calling eq: incompatible types for comparison | - | error calling eq: incompatible types for comparison | error calling eq: incompatible types for comparison | error calling eq: incompatible types for comparison |
| `ne` | - | error calling ne: incompatible types for comparison | - | error calling ne: incompatible types for comparison | error calling ne: incompatible types for comparison | error calling ne: incompatible types for comparison |
| `lt` | error calling lt: invalid type for comparison | error calling lt: incompatible types for comparison | - | error calling lt: incompatible types for comparison | error calling lt: invalid type for comparison | error calling lt: invalid type for comparison |
| `le` | error calling le: invalid type for comparison | error calling le: incompatible types for comparison | - | error calling le: incompatible types for comparison | error calling le: invalid type for comparison | error calling le: invalid type for comparison |
| `gt` | error calling gt: invalid type for comparison | error calling gt: incompatible types for comparison | - | error calling gt: incompatible types for comparison | error calling gt: invalid type for comparison | error calling gt: invalid type for comparison |
| `ge` | error calling ge: invalid type for comparison | error calling ge: incompatible types for comparison | - | error calling ge: incompatible types for comparison | error calling ge: invalid type for comparison | error calling ge: invalid type for comparison |
| `required` | execution error | - | - | - | - | - |
| `fail` | execution error | execution error | execution error | execution error | execution error | execution error |
| `ternary` | invalid value; expected bool | wrong type for value; expected bool; got string | wrong type for value; expected bool; got float64 | - | wrong type for value; expected bool; got []interface | wrong type for value; expected bool; got map[string]interfac... |
