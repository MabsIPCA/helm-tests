# Test 22: URL Function Errors

**Generated:** 2026-02-19 19:53:28

**Total:** 18 | **Passed:** 8 (44.4%) | **Failed:** 10 (55.6%)

## Test Matrix

| Function | nilValue | stringValue | numberValue | boolValue | listValue | dictValue |
|----------|------|------|------|------|------|------|
| `urlParse` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `urlJoin` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `urlquery` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |

## Failure Details

| Function | nilValue | stringValue | numberValue | boolValue | listValue | dictValue |
|----------|------|------|------|------|------|------|
| `urlParse` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `urlJoin` | error calling urlJoin: runtime error: invalid memory address... | - | error calling urlJoin: unable to parse path key, must be of ... | error calling urlJoin: unable to parse path key, must be of ... | error calling urlJoin: unable to parse path key, must be of ... | error calling urlJoin: unable to parse path key, must be of ... |
