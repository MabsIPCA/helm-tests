# Helm Chart Catalog Results

## Summary by Repository

| # | Repository | Charts | Runs | ✅ Success | ❌ Failures | 🔧 Dep Failures | Success Rate | Status |
|---|------------|--------|------|------------|-------------|-----------------|--------------|--------|
| 1 | [featurehub-io/featurehub](https://github.com/featurehub-io/featurehub) | 5 | 6 | 5 | 1 | 0 | 83.3% | Kept |
|---|------------|--------|------|------------|-------------|-----------------|--------------|--------|
| **Total** | **1 repos** | **5** | **6** | **5** | **1** | **0** | **83.3%** | - |

## Failure Details

### featurehub-io/featurehub

#### `cloned\featurehub-io__featurehub\pipeline\deploy\src\main\resources\featurehub\charts\common`

| Values Files | Command | Error |
|--------------|---------|-------|
| (default) | `helm template test cloned\featurehub-io__featurehub\pipeline\deploy\src\main\resources\featurehub\charts\common` | Error: template: common-config/templates/configmap.yaml:10:17: executing "common-config/templates/configmap.yaml" at <.Values.global.testEnvironment>: nil pointer evaluating interface {}.testEnvironment  Use --debug flag to render out invalid YAML  |

