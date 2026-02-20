# Test 15: tpl Function Errors

**Generated:** 2026-02-20 01:48:12

**Total:** 7 | **Passed:** 0 (0.0%) | **Failed:** 7 (100.0%)

## Test Matrix

| Function | default |
|----------|------|
| `tplSyntaxError` | ❌ |
| `tplNilPointer` | ❌ |
| `tplRequired` | ❌ |
| `tplDivZero` | ❌ |
| `tplTypeMismatch` | ❌ |
| `tplInvalidFunc` | ❌ |
| `tplFail` | ❌ |

## Failure Details

| Function | default |
|----------|------|
| `tplSyntaxError` | unclosed action |
| `tplNilPointer` | nil pointer evaluating |
| `tplRequired` | tpl required value |
| `tplDivZero` | integer divide by zero |
| `tplTypeMismatch` | range can't iterate over |
| `tplInvalidFunc` | function "invalidFunc" not defined |
| `tplFail` | tpl triggered failure |
