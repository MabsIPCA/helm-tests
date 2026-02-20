# Test 7: File Read, Path, and Advanced Files Errors

**Generated:** 2026-02-20 02:06:38

**Total:** 66 | **Passed:** 6 (9.1%) | **Failed:** 60 (90.9%)

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
| `filesMissingRequired` | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ |
| `filesEmptyRequired` | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ |
| `filesNilPath` | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ |
| `filesTplError` | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ |

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
| `filesMissingRequired` | File content is required | File content is required | File content is required | File content is required | File content is required | File content is required |
| `filesEmptyRequired` | wrong type for value; expected string; got []uint8 | wrong type for value; expected string; got []uint8 | wrong type for value; expected string; got []uint8 | wrong type for value; expected string; got []uint8 | wrong type for value; expected string; got []uint8 | wrong type for value; expected string; got []uint8 |
| `filesNilPath` | wrong type for value; expected string; got interface | wrong type for value; expected string; got interface | wrong type for value; expected string; got interface | wrong type for value; expected string; got interface | wrong type for value; expected string; got interface | wrong type for value; expected string; got interface |
| `filesTplError` | error calling tpl: error during tpl function execution for "... | error calling tpl: error during tpl function execution for "... | error calling tpl: error during tpl function execution for "... | error calling tpl: error during tpl function execution for "... | error calling tpl: error during tpl function execution for "... | error calling tpl: error during tpl function execution for "... |
