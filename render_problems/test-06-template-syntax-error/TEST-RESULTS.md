# Test 6: Template Syntax Errors

**Generated:** 2026-02-20 00:44:01

**Total:** 10 | **Passed:** 0 (0.0%) | **Failed:** 10 (100.0%)

## Test Matrix

| Function | default |
|----------|------|
| `unclosedAction` | ❌ |
| `invalidFunction` | ❌ |
| `wrongArgs` | ❌ |
| `undefinedVar` | ❌ |
| `invalidRange` | ❌ |
| `wrongPipelineType` | ❌ |
| `fieldOnNonMap` | ❌ |
| `rangeOverScalar` | ❌ |
| `indexOutOfRange` | ❌ |
| `lenOnScalar` | ❌ |

## Failure Details

| Function | default |
|----------|------|
| `unclosedAction` | parse error at (test-06a-unclosed-action/templates/test.yaml... |
| `invalidFunction` | parse error at (test-06b-invalid-function/templates/test.yam... |
| `wrongArgs` | wrong number of args |
| `undefinedVar` | parse error at (test-06d-undefined-var/templates/test.yaml:8... |
| `invalidRange` | parse error at (test-06e-invalid-range/templates/test.yaml:8... |
| `wrongPipelineType` | expected |
| `fieldOnNonMap` | error calling index: cannot index slice/array with type stri... |
| `rangeOverScalar` | range can't iterate over |
| `indexOutOfRange` | error calling index: index out of range: 999 |
| `lenOnScalar` | error calling len: len of type float64 |
