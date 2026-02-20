# Test 25: Semver Function Errors

**Generated:** 2026-02-20 00:44:30

**Total:** 48 | **Passed:** 8 (16.7%) | **Failed:** 40 (83.3%)

## Test Matrix

| Function | nilValue | stringValue | numberValue | boolValue | listValue | dictValue |
|----------|------|------|------|------|------|------|
| `semver` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `semverCompare` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `semverCompareBasic` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `semverComparePrerelease` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `semverCompareHyphen` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `semverCompareWildcard` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `semverCompareTilde` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `semverCompareCaret` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |

## Failure Details

| Function | nilValue | stringValue | numberValue | boolValue | listValue | dictValue |
|----------|------|------|------|------|------|------|
| `semver` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `semverCompare` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `semverCompareBasic` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `semverComparePrerelease` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `semverCompareHyphen` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `semverCompareWildcard` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `semverCompareTilde` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `semverCompareCaret` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
