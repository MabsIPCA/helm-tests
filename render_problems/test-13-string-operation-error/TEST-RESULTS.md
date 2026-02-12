# Test 13: String Operation Errors

**Generated:** 2026-02-11 00:50:50

**Total:** 252 | **Passed:** 62 (24.6%) | **Failed:** 190 (75.4%)

## Test Matrix

| Function | nilValue | stringValue | numberValue | boolValue | listValue | dictValue |
|----------|------|------|------|------|------|------|
| `print` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `println` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `printf` | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ |
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
| `printf` | wrong type for value; expected string; got interfa... | wrong type for value; expected string; got interfa... | wrong type for value; expected string; got interfa... | wrong type for value; expected string; got interfa... | wrong type for value; expected string; got interfa... | wrong type for value; expected string; got interfa... |
| `trim` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []inter... | wrong type for value; expected string; got map[str... |
| `trimAll` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []inter... | wrong type for value; expected string; got map[str... |
| `trimPrefix` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []inter... | wrong type for value; expected string; got map[str... |
| `trimSuffix` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []inter... | wrong type for value; expected string; got map[str... |
| `lower` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []inter... | wrong type for value; expected string; got map[str... |
| `upper` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []inter... | wrong type for value; expected string; got map[str... |
| `title` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []inter... | wrong type for value; expected string; got map[str... |
| `untitle` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []inter... | wrong type for value; expected string; got map[str... |
| `repeat` | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface |
| `substr` | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface |
| `nospace` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []inter... | wrong type for value; expected string; got map[str... |
| `trunc` | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface |
| `abbrev` | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface |
| `abbrevboth` | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface |
| `initials` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []inter... | wrong type for value; expected string; got map[str... |
| `randAlphaNum` | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface |
| `randAlpha` | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface |
| `randNumeric` | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface |
| `randAscii` | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface |
| `wrap` | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface |
| `wrapWith` | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface |
| `contains` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []inter... | wrong type for value; expected string; got map[str... |
| `hasPrefix` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []inter... | wrong type for value; expected string; got map[str... |
| `hasSuffix` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []inter... | wrong type for value; expected string; got map[str... |
| `indent` | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface |
| `nindent` | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface |
| `replace` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []inter... | wrong type for value; expected string; got map[str... |
| `plural` | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface | wrong type for value; expected int; got interface |
| `snakecase` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []inter... | wrong type for value; expected string; got map[str... |
| `camelcase` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []inter... | wrong type for value; expected string; got map[str... |
| `kebabcase` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []inter... | wrong type for value; expected string; got map[str... |
| `swapcase` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []inter... | wrong type for value; expected string; got map[str... |
| `shuffle` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []inter... | wrong type for value; expected string; got map[str... |
| `atoi` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []inter... | wrong type for value; expected string; got map[str... |
