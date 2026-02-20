# Test 3: Template Syntax Errors

**Generated:** 2026-02-20 01:17:41

**Total:** 10 | **Passed:** 0 (0.0%) | **Failed:** 10 (100.0%)

## Test Matrix

| Function | default |
|----------|------|
| `fieldOnNonMap` | ❌ |
| `indexOutOfRange` | ❌ |
| `invalidFunction` | ❌ |
| `invalidRange` | ❌ |
| `lenOnScalar` | ❌ |
| `rangeOverScalar` | ❌ |
| `unclosedAction` | ❌ |
| `undefinedVar` | ❌ |
| `wrongArgs` | ❌ |
| `wrongPipelineType` | ❌ |

## Failure Details

| Function | default |
|----------|------|
| `fieldOnNonMap` | error calling index: cannot index slice/array with type stri... |
| `indexOutOfRange` | error calling index: index out of range: 999 |
| `invalidFunction` | parse error at (test-03c-invalid-function/templates/test.yam... |
| `invalidRange` | parse error at (test-03d-invalid-range/templates/test.yaml:8... |
| `lenOnScalar` | error calling len: len of type float64 |
| `rangeOverScalar` | range can't iterate over |
| `unclosedAction` | parse error at (test-03g-unclosed-action/templates/test.yaml... |
| `undefinedVar` | parse error at (test-03h-undefined-var/templates/test.yaml:8... |
| `wrongArgs` | wrong number of args |
| `wrongPipelineType` | expected |
