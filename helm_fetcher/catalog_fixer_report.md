# Catalog Fixer Report

**Source:** runs/20260529_135859_artifacthub/catalog_by_project.json
**Date:** 2026-06-02

## Summary

| Metric | Count |
|---|---:|
| Repos processed | 322 |
| Failing runs (before) | 252 |
| Resolved | 38 |
| Resolution rate | 15% |

## By Error Kind

| Kind | Before | Resolved | Still failing |
|---|---:|---:|---:|
| nil_pointer | 35 | 20 | 15 |
| required_value | 32 | 18 | 14 |
| type_mismatch | 0 | 0 | 0 |
| other | 185 | 0 | 185 |

## Stop Reasons

| Reason | Count |
|---|---:|
| resolved | 38 |
| unfixable_error_kind | 211 |
| loop_detected | 2 |
| max_iterations | 1 |
