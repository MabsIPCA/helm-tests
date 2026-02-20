# Test 14: Helper Template Errors

**Generated:** 2026-02-20 02:07:05

**Total:** 6 | **Passed:** 0 (0.0%) | **Failed:** 6 (100.0%)

## Test Matrix

| Function | default |
|----------|------|
| `helperNilPointer` | ❌ |
| `helperRequired` | ❌ |
| `helperRangeOverScalar` | ❌ |
| `helperDivZero` | ❌ |
| `helperMissingInclude` | ❌ |
| `helperExplicitFail` | ❌ |

## Failure Details

| Function | default |
|----------|------|
| `helperNilPointer` | nil pointer evaluating |
| `helperRequired` | missingRequiredField is required |
| `helperRangeOverScalar` | range can't iterate over |
| `helperDivZero` | integer divide by zero |
| `helperMissingInclude` | no template "non.existent.helper.template" associated |
| `helperExplicitFail` | This helper explicitly fails |
