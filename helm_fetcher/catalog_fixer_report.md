# Catalog Fixer Report

**Source:** runs/20260529_135859_artifacthub/catalog_by_project.json
**Date:** 2026-06-05

## Summary

| Metric | Count |
|---|---:|
| Repos processed | 322 |
| Failing runs (before) | 187 |
| Resolved | 27 |
| Resolution rate | 14% |

## By Error Kind

| Kind | Before | Resolved | Still failing |
|---|---:|---:|---:|
| nil_pointer | 3 | 3 | 0 |
| required_value | 35 | 22 | 13 |
| type_mismatch | 0 | 0 | 0 |
| kube_version | 2 | 2 | 0 |
| other | 147 | 0 | 147 |

## Stop Reasons

| Reason | Count |
|---|---:|
| resolved | 27 |
| unfixable_error_kind | 155 |
| non_fixable_yaml | 3 |
| loop_detected | 2 |
| max_iterations | 0 |
