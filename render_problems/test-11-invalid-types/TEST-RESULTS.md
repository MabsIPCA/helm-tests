# Test 11: Type Conversion Errors

**Generated:** 2026-02-13 00:56:54

**Total:** 90 | **Passed:** 69 (76.7%) | **Failed:** 21 (23.3%)

## Test Matrix

| Function | nilValue | stringValue | numberValue | boolValue | listValue | dictValue |
|----------|------|------|------|------|------|------|
| `toStrings` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `toDecimal` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `toJson` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `mustToJson` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `toPrettyJson` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `mustToPrettyJson` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `toRawJson` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `mustToRawJson` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `fromYaml` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `fromJson` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `fromJsonArray` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `toYaml` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `toYamlPretty` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `fromYamlArray` | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| `toToml` | ❌ | ✅ | ✅ | ✅ | ✅ | ✅ |

## Failure Details

| Function | nilValue | stringValue | numberValue | boolValue | listValue | dictValue |
|----------|------|------|------|------|------|------|
| `fromYaml` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `fromJson` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `fromJsonArray` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `fromYamlArray` | invalid value; expected string | - | wrong type for value; expected string; got float64 | wrong type for value; expected string; got bool | wrong type for value; expected string; got []interface | wrong type for value; expected string; got map[string]interf... |
| `toToml` | error calling toToml: reflect: call of reflect.Value.Type on... | - | - | - | - | - |
