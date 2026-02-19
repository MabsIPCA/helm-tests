# Test 13: String Operation Errors

**Generated:** 2026-02-19 19:53:08

**Total:** 300 | **Passed:** 110 (36.7%) | **Failed:** 190 (63.3%)

## Test Matrix

| Function | nilValue | stringValue | numberValue | boolValue | listValue | dictValue |
|----------|------|------|------|------|------|------|
| `print` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `println` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `printf` | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ |
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
| `repeat` | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ |
| `substr` | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ |
| `nospace` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `trunc` | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ |
| `abbrev` | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ |
| `abbrevboth` | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ |
| `initials` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `randAlphaNum` | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ |
| `randAlpha` | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ |
| `randNumeric` | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ |
| `randAscii` | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ |
| `wrap` | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ |
| `wrapWith` | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ |
| `contains` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `hasPrefix` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `hasSuffix` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `quote` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `squote` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `cat` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `indent` | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ |
| `nindent` | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ |
| `replace` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `plural` | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ |
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
| `printf` | wrong type for value; expected string; got interface | wrong type for value; expected string; got interface | wrong type for value; expected string; got interface | wrong type for value; expected string; got interface | wrong type for value; expected string; got interface | wrong type for value; expected string; got interface |
| `trim` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `trimAll` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `trimPrefix` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `trimSuffix` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `lower` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `upper` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `title` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `untitle` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `repeat` | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface |
| `substr` | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface |
| `nospace` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `trunc` | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface |
| `abbrev` | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface |
| `abbrevboth` | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface |
| `initials` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `randAlphaNum` | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface |
| `randAlpha` | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface |
| `randNumeric` | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface |
| `randAscii` | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface |
| `wrap` | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface |
| `wrapWith` | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface |
| `contains` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `hasPrefix` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `hasSuffix` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `indent` | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface |
| `nindent` | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface |
| `replace` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `plural` | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface |
| `snakecase` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `camelcase` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `kebabcase` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `swapcase` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `shuffle` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `atoi` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
