# Test 17: Math Operation Errors

**Generated:** 2026-02-20 01:48:13

**Total:** 114 | **Passed:** 111 (97.4%) | **Failed:** 3 (2.6%)

## Test Matrix

| Function | nilValue | stringValue | numberValue | boolValue | listValue | dictValue |
|----------|------|------|------|------|------|------|
| `add` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `add1` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `sub` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `div` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `mod` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `mul` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `max` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `min` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `len` | ❌ | ✅ | ❌ | ❌ | ✅ | ✅ |
| `addf` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `add1f` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `subf` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `divf` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `mulf` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `maxf` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `minf` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `floor` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `ceil` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `round` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |

## Failure Details

| Function | nilValue | stringValue | numberValue | boolValue | listValue | dictValue |
|----------|------|------|------|------|------|------|
| `len` | error calling len: reflect: call of reflect.Value.Type on ze... | - | error calling len: len of type float64 | error calling len: len of type bool | - | - |
