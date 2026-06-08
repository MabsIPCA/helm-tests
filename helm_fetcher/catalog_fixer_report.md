# Catalog Fixer Report

**Source:** runs/20260529_135859_artifacthub/catalog_by_project.json
**Date:** 2026-06-08

## Summary

| Metric | Count |
|---|---:|
| Repos processed | 322 |
| Failing runs (before) | 248 |
| Resolved | 29 |
| Resolution rate | 11% |

## By Error Kind

| Kind | Before | Resolved | Still failing |
|---|---:|---:|---:|
| nil_pointer | 3 | 3 | 0 |
| required_value | 37 | 24 | 13 |
| type_mismatch | 0 | 0 | 0 |
| kube_version | 2 | 2 | 0 |
| other | 206 | 0 | 206 |

## Stop Reasons

| Reason | Count |
|---|---:|
| resolved | 29 |
| unfixable_error_kind | 212 |
| non_fixable_yaml | 4 |
| non_fixable_builtin | 0 |
| loop_detected | 3 |
| max_iterations | 0 |
