# Test 13: String Operation Errors

**Generated:** 2026-02-20 00:44:10

**Total:** 462 | **Passed:** 224 (48.5%) | **Failed:** 238 (51.5%)

## Test Matrix

| Function | nilValue | stringValue | numberValue | boolValue | listValue | dictValue |
|----------|------|------|------|------|------|------|
| `print` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `println` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `printf-%s` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `printf-%d` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `printf-%f` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `printf-%t` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `printf-%x` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `printf-%b` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `printf-%c` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `printf-%q` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `trim` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `trimAll` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `trimPrefix` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `trimSuffix` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `lower` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `upper` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `title` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `untitle` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `repeat-positive` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `repeat-zero` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `repeat-negative` | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ |
| `substr-positive` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `substr-zero` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `substr-negative` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `nospace` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `trunc-positive` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `trunc-zero` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `trunc-negative` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `abbrev-positive` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `abbrev-zero` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `abbrev-negative` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `abbrevboth-positive` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `abbrevboth-zero` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `abbrevboth-negative` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `initials` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `randAlphaNum-positive` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `randAlphaNum-zero` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `randAlphaNum-negative` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `randAlpha-positive` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `randAlpha-zero` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `randAlpha-negative` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `randNumeric-positive` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `randNumeric-zero` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `randNumeric-negative` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `randAscii-positive` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `randAscii-zero` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `randAscii-negative` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `wrap-positive` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `wrap-zero` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `wrap-negative` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `wrapWith-positive` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `wrapWith-zero` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `wrapWith-negative` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `contains` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `hasPrefix` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `hasSuffix` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `quote` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `squote` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `cat` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `indent-positive` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `indent-zero` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `indent-negative` | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ |
| `nindent-positive` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `nindent-zero` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `nindent-negative` | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ |
| `replace` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `plural-positive` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `plural-zero` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `plural-negative` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `snakecase` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `camelcase` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `kebabcase` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `swapcase` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `shuffle` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `atoi` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `int64` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `float64` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |

## Failure Details

| Function | nilValue | stringValue | numberValue | boolValue | listValue | dictValue |
|----------|------|------|------|------|------|------|
| `trim` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `trimAll` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `trimPrefix` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `trimSuffix` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `lower` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `upper` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `title` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `untitle` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `repeat-positive` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `repeat-zero` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `repeat-negative` | invalid value; expected string | error calling repeat: strings: negative Repeat count | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `substr-positive` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `substr-zero` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `substr-negative` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `nospace` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `trunc-positive` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `trunc-zero` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `trunc-negative` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `abbrev-positive` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `abbrev-zero` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `abbrev-negative` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `abbrevboth-positive` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `abbrevboth-zero` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `abbrevboth-negative` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `initials` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `wrap-positive` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `wrap-zero` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `wrap-negative` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `wrapWith-positive` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `wrapWith-zero` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `wrapWith-negative` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `contains` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `hasPrefix` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `hasSuffix` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `indent-positive` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `indent-zero` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `indent-negative` | invalid value; expected string | error calling indent: strings: negative Repeat count | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `nindent-positive` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `nindent-zero` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `nindent-negative` | invalid value; expected string | error calling nindent: strings: negative Repeat count | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `replace` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `snakecase` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `camelcase` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `kebabcase` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `swapcase` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `shuffle` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `atoi` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
