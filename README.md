# helm-tests

This repository contains Helm chart test suites and tools used during the development of a Master thesis. It is designed to validate Helm values overrides, rendering behaviors, and edge cases for Helm charts.

## Prerequisites

- [Go](https://golang.org/) (>= 1.18)
- [Helm](https://helm.sh/) (>= 3.x)

## Structure

- `override_tests_suite/` — Automated test suite for override precedence, type changes, multiple values files, subchart overrides, and more.
- `render_problems/` — Test cases for rendering errors, template edge cases, and Helm rendering failures.

## Running Override Tests

1. Ensure prerequisites are installed.
2. Run the Go test runner in the `override_tests_suite` directory:

```powershell
cd override_tests_suite
go run run-tests.go
```

Or run a specific test:

```powershell
cd override_tests_suite
go run run-tests.go --test 01
```

## Running Render Problem Tests

The `render_problems` suite is designed to test Helm rendering failures, template errors, and edge cases. To run these tests:

```powershell
cd render_problems
go run run-tests.go
```

## Test Reports

Each test directory generates a `TEST-RESULTS.md` file summarizing the test cases, values before/after overrides, rendering errors, and the Helm command used. Render problem tests will highlight template errors and expected failures.


For more information about Helm, see the [Helm documentation](https://helm.sh/docs/).
