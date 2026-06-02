# Catalog Fixer Report

**Source:** results/github/catalog_by_project.json
**Date:** 2026-06-02

## Summary

| Metric | Count |
|---|---:|
| Repos processed | 289 |
| Failing runs (before) | 1135 |
| Resolved | 412 |
| Resolution rate | 36% |

## By Error Kind

| Kind | Before | Resolved | Still failing |
|---|---:|---:|---:|
| nil_pointer | 586 | 411 | 175 |
| required_value | 0 | 0 | 0 |
| type_mismatch | 0 | 0 | 0 |
| other | 549 | 1 | 548 |

## Stop Reasons

| Reason | Count |
|---|---:|
| resolved | 412 |
| unfixable_error_kind | 716 |
| loop_detected | 1 |
| max_iterations | 6 |
