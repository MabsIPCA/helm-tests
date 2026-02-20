# Test 9: Regex Operation Errors

**Generated:** 2026-02-20 02:06:37

**Total:** 72 | **Passed:** 12 (16.7%) | **Failed:** 60 (83.3%)

## Test Matrix

| Function | nilValue | stringValue | numberValue | boolValue | listValue | dictValue |
|----------|------|------|------|------|------|------|
| `regexMatch` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `mustRegexMatch` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `regexFindAll` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `mustRegexFindAll` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `regexFind` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `mustRegexFind` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `regexReplaceAll` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `mustRegexReplaceAll` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `regexReplaceAllLiteral` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `mustRegexReplaceAllLiteral` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `regexSplit` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `mustRegexSplit` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |

## Failure Details

| Function | nilValue | stringValue | numberValue | boolValue | listValue | dictValue |
|----------|------|------|------|------|------|------|
| `regexMatch` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `mustRegexMatch` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `regexFindAll` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `mustRegexFindAll` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `regexFind` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `mustRegexFind` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `regexReplaceAll` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `mustRegexReplaceAll` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `regexReplaceAllLiteral` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `mustRegexReplaceAllLiteral` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `regexSplit` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `mustRegexSplit` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
