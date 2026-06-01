# Catalog Fixer Report

**Source:** backup/run_002_274/catalog_kept.json
**Date:** 2026-05-17

## Summary

| Metric | Count |
|---|---:|
| Repos processed | 55 |
| Failing runs (before) | 373 |
| Resolved | 225 |
| Resolution rate | 60% |

## By Error Kind

| Kind | Before | Resolved | Still failing |
|---|---:|---:|---:|
| nil_pointer | 252 | 224 | 28 |
| required_value | 0 | 0 | 0 |
| other | 121 | 1 | 120 |

## Stop Reasons

| Reason | Count |
|---|---:|
| resolved | 225 |
| unfixable_error_kind | 144 |
| loop_detected | 2 |
| max_iterations | 2 |
