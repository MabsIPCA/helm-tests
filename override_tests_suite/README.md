# Helm Values Override Test Suite

This test suite validates Helm values file override behavior across different scenarios.

## Overview

The test suite covers:
- **Test 01**: Basic Data Types Override - Tests all primitive data types
- **Test 02**: Multiple Values Files - Tests precedence with multiple `-f` flags
- **Test 03**: Multiple Values Files with Default - Tests default values file combined with overrides
- **Test 04**: Type Changes - Tests changing data types during override
- **Test 05**: List Override Behavior - Tests list replacement vs merging
- **Test 06**: Deep Nested Override - Tests deeply nested structure overrides
- **Test 07**: Null Value Override - Tests null value handling
- **Test 08**: Sub-Chart Override - Tests sub-chart values override
- **Test 09**: Global Values Override - Tests global values propagation
- **Test 10**: Set Flag Override - Tests `--set` and `--set-string` flags

## Running Tests

```bash
# Run all tests
make test

# Build test runner
make build

# Clean generated files
make clean
```

## Test Structure

Each test directory contains:
- `config.json` - Test configuration and expected results
- `Chart.yaml` - Helm chart definition
- `values.yaml` - Default values file
- `values-*.yaml` - Override values files
- `templates/` - Template files for rendering
- `TEST-RESULTS.md` - Generated test results (after running tests)

## Override Precedence

Helm values override precedence (lowest to highest):
1. `values.yaml` in chart
2. Parent chart's `values.yaml`
3. Values file passed with `-f` or `--values` (in order)
4. `--set` parameters
5. `--set-string` parameters
6. `--set-file` parameters

## Data Types Tested

- String
- Integer
- Float
- Boolean
- Null
- List/Array
- Map/Object
- Nested structures

