# Test 12: List Operation Errors

**Generated:** 2026-02-20 00:44:06

**Total:** 180 | **Passed:** 29 (16.1%) | **Failed:** 151 (83.9%)

## Test Matrix

| Function | nilValue | stringValue | numberValue | boolValue | listValue | dictValue |
|----------|------|------|------|------|------|------|
| `first` | ❌ | ❌ | ❌ | ❌ | ✅ | ❌ |
| `mustFirst` | ❌ | ❌ | ❌ | ❌ | ✅ | ❌ |
| `rest` | ❌ | ❌ | ❌ | ❌ | ✅ | ❌ |
| `mustRest` | ❌ | ❌ | ❌ | ❌ | ✅ | ❌ |
| `last` | ❌ | ❌ | ❌ | ❌ | ✅ | ❌ |
| `mustLast` | ❌ | ❌ | ❌ | ❌ | ✅ | ❌ |
| `initial` | ❌ | ❌ | ❌ | ❌ | ✅ | ❌ |
| `mustInitial` | ❌ | ❌ | ❌ | ❌ | ✅ | ❌ |
| `append` | ❌ | ❌ | ❌ | ❌ | ✅ | ❌ |
| `mustAppend` | ❌ | ❌ | ❌ | ❌ | ✅ | ❌ |
| `prepend` | ❌ | ❌ | ❌ | ❌ | ✅ | ❌ |
| `mustPrepend` | ❌ | ❌ | ❌ | ❌ | ✅ | ❌ |
| `concat` | ❌ | ❌ | ❌ | ❌ | ✅ | ❌ |
| `reverse` | ❌ | ❌ | ❌ | ❌ | ✅ | ❌ |
| `mustReverse` | ❌ | ❌ | ❌ | ❌ | ✅ | ❌ |
| `uniq` | ❌ | ❌ | ❌ | ❌ | ✅ | ❌ |
| `mustUniq` | ❌ | ❌ | ❌ | ❌ | ✅ | ❌ |
| `without` | ❌ | ❌ | ❌ | ❌ | ✅ | ❌ |
| `mustWithout` | ❌ | ❌ | ❌ | ❌ | ✅ | ❌ |
| `has` | ✅ | ❌ | ❌ | ❌ | ✅ | ❌ |
| `mustHas` | ✅ | ❌ | ❌ | ❌ | ✅ | ❌ |
| `compact` | ❌ | ❌ | ❌ | ❌ | ✅ | ❌ |
| `mustCompact` | ❌ | ❌ | ❌ | ❌ | ✅ | ❌ |
| `index` | ❌ | ✅ | ❌ | ❌ | ✅ | ❌ |
| `slice` | ❌ | ❌ | ❌ | ❌ | ✅ | ❌ |
| `mustSlice` | ❌ | ❌ | ❌ | ❌ | ✅ | ❌ |
| `until` | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ |
| `untilStep` | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ |
| `seq` | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ |
| `chunk` | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ |

## Failure Details

| Function | nilValue | stringValue | numberValue | boolValue | listValue | dictValue |
|----------|------|------|------|------|------|------|
| `first` | error calling first: runtime error: invalid memory address o... | error calling first: Cannot find first on type string | error calling first: Cannot find first on type float64 | error calling first: Cannot find first on type bool | - | error calling first: Cannot find first on type map |
| `mustFirst` | error calling mustFirst: runtime error: invalid memory addre... | error calling mustFirst: Cannot find first on type string | error calling mustFirst: Cannot find first on type float64 | error calling mustFirst: Cannot find first on type bool | - | error calling mustFirst: Cannot find first on type map |
| `rest` | error calling rest: runtime error: invalid memory address or... | error calling rest: Cannot find rest on type string | error calling rest: Cannot find rest on type float64 | error calling rest: Cannot find rest on type bool | - | error calling rest: Cannot find rest on type map |
| `mustRest` | error calling mustRest: runtime error: invalid memory addres... | error calling mustRest: Cannot find rest on type string | error calling mustRest: Cannot find rest on type float64 | error calling mustRest: Cannot find rest on type bool | - | error calling mustRest: Cannot find rest on type map |
| `last` | error calling last: runtime error: invalid memory address or... | error calling last: Cannot find last on type string | error calling last: Cannot find last on type float64 | error calling last: Cannot find last on type bool | - | error calling last: Cannot find last on type map |
| `mustLast` | error calling mustLast: runtime error: invalid memory addres... | error calling mustLast: Cannot find last on type string | error calling mustLast: Cannot find last on type float64 | error calling mustLast: Cannot find last on type bool | - | error calling mustLast: Cannot find last on type map |
| `initial` | error calling initial: runtime error: invalid memory address... | error calling initial: Cannot find initial on type string | error calling initial: Cannot find initial on type float64 | error calling initial: Cannot find initial on type bool | - | error calling initial: Cannot find initial on type map |
| `mustInitial` | error calling mustInitial: runtime error: invalid memory add... | error calling mustInitial: Cannot find initial on type strin... | error calling mustInitial: Cannot find initial on type float... | error calling mustInitial: Cannot find initial on type bool | - | error calling mustInitial: Cannot find initial on type map |
| `append` | error calling append: runtime error: invalid memory address ... | error calling append: Cannot push on type string | error calling append: Cannot push on type float64 | error calling append: Cannot push on type bool | - | error calling append: Cannot push on type map |
| `mustAppend` | error calling mustAppend: runtime error: invalid memory addr... | error calling mustAppend: Cannot push on type string | error calling mustAppend: Cannot push on type float64 | error calling mustAppend: Cannot push on type bool | - | error calling mustAppend: Cannot push on type map |
| `prepend` | error calling prepend: runtime error: invalid memory address... | error calling prepend: Cannot prepend on type string | error calling prepend: Cannot prepend on type float64 | error calling prepend: Cannot prepend on type bool | - | error calling prepend: Cannot prepend on type map |
| `mustPrepend` | error calling mustPrepend: runtime error: invalid memory add... | error calling mustPrepend: Cannot prepend on type string | error calling mustPrepend: Cannot prepend on type float64 | error calling mustPrepend: Cannot prepend on type bool | - | error calling mustPrepend: Cannot prepend on type map |
| `concat` | error calling concat: runtime error: invalid memory address ... | error calling concat: Cannot concat type string as list | error calling concat: Cannot concat type float64 as list | error calling concat: Cannot concat type bool as list | - | error calling concat: Cannot concat type map as list |
| `reverse` | error calling reverse: runtime error: invalid memory address... | error calling reverse: Cannot find reverse on type string | error calling reverse: Cannot find reverse on type float64 | error calling reverse: Cannot find reverse on type bool | - | error calling reverse: Cannot find reverse on type map |
| `mustReverse` | error calling mustReverse: runtime error: invalid memory add... | error calling mustReverse: Cannot find reverse on type strin... | error calling mustReverse: Cannot find reverse on type float... | error calling mustReverse: Cannot find reverse on type bool | - | error calling mustReverse: Cannot find reverse on type map |
| `uniq` | error calling uniq: runtime error: invalid memory address or... | error calling uniq: Cannot find uniq on type string | error calling uniq: Cannot find uniq on type float64 | error calling uniq: Cannot find uniq on type bool | - | error calling uniq: Cannot find uniq on type map |
| `mustUniq` | error calling mustUniq: runtime error: invalid memory addres... | error calling mustUniq: Cannot find uniq on type string | error calling mustUniq: Cannot find uniq on type float64 | error calling mustUniq: Cannot find uniq on type bool | - | error calling mustUniq: Cannot find uniq on type map |
| `without` | error calling without: runtime error: invalid memory address... | error calling without: Cannot find without on type string | error calling without: Cannot find without on type float64 | error calling without: Cannot find without on type bool | - | error calling without: Cannot find without on type map |
| `mustWithout` | error calling mustWithout: runtime error: invalid memory add... | error calling mustWithout: Cannot find without on type strin... | error calling mustWithout: Cannot find without on type float... | error calling mustWithout: Cannot find without on type bool | - | error calling mustWithout: Cannot find without on type map |
| `has` | - | error calling has: Cannot find has on type string | error calling has: Cannot find has on type float64 | error calling has: Cannot find has on type bool | - | error calling has: Cannot find has on type map |
| `mustHas` | - | error calling mustHas: Cannot find has on type string | error calling mustHas: Cannot find has on type float64 | error calling mustHas: Cannot find has on type bool | - | error calling mustHas: Cannot find has on type map |
| `compact` | error calling compact: runtime error: invalid memory address... | error calling compact: Cannot compact on type string | error calling compact: Cannot compact on type float64 | error calling compact: Cannot compact on type bool | - | error calling compact: Cannot compact on type map |
| `mustCompact` | error calling mustCompact: runtime error: invalid memory add... | error calling mustCompact: Cannot compact on type string | error calling mustCompact: Cannot compact on type float64 | error calling mustCompact: Cannot compact on type bool | - | error calling mustCompact: Cannot compact on type map |
| `index` | error calling index: index of untyped nil | - | error calling index: can't index item of type float64 | error calling index: can't index item of type bool | - | error calling index: value has type int; should be string |
| `slice` | error calling slice: runtime error: invalid memory address o... | error calling slice: list should be type of slice or array b... | error calling slice: list should be type of slice or array b... | error calling slice: list should be type of slice or array b... | - | error calling slice: list should be type of slice or array b... |
| `mustSlice` | error calling mustSlice: runtime error: invalid memory addre... | error calling mustSlice: list should be type of slice or arr... | error calling mustSlice: list should be type of slice or arr... | error calling mustSlice: list should be type of slice or arr... | - | error calling mustSlice: list should be type of slice or arr... |
| `until` | invalid value; expected int | wrong type for value; expected int; got string | wrong type for value; expected int; got float64 | wrong type for value; expected int; got bool | wrong type for value; expected int; got []interface | wrong type for value; expected int; got map[string]interface |
| `untilStep` | invalid value; expected int | wrong type for value; expected int; got string | wrong type for value; expected int; got float64 | wrong type for value; expected int; got bool | wrong type for value; expected int; got []interface | wrong type for value; expected int; got map[string]interface |
| `seq` | invalid value; expected int | wrong type for value; expected int; got string | wrong type for value; expected int; got float64 | wrong type for value; expected int; got bool | wrong type for value; expected int; got []interface | wrong type for value; expected int; got map[string]interface |
| `chunk` | invalid value; expected int | wrong type for value; expected int; got string | wrong type for value; expected int; got float64 | wrong type for value; expected int; got bool | wrong type for value; expected int; got []interface | wrong type for value; expected int; got map[string]interface |
