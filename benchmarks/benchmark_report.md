# Inlining Helper Functions Benchmark Report

## Overview

This report presents the results of benchmarking different implementation approaches for vector operations in the govec library. The goal is to determine whether using helper functions has a significant performance impact compared to inlining code.

Three implementation approaches were tested:
1. **Inlined**: All calculations are performed inline without calling helper functions
2. **With Helpers**: Calculations are performed by calling small helper functions
3. **With Library**: Calculations are performed by calling existing library functions
4. **Original**: The current implementation in the library (for comparison)

## Benchmark Environment

- Go version: 1.21.8
- OS: Linux
- Architecture: amd64
- CPU: AMD EPYC 9B14

## Simple Operations

### Dot Product

| Vector Type | Implementation | Operations/sec | ns/op | Bytes/op | Allocs/op |
|-------------|---------------|---------------|-------|----------|-----------|
| 2F32 | Inlined | 1,000,000,000 | 0.3297 | 0 | 0 |
| 2F32 | Library | 1,000,000,000 | 0.3254 | 0 | 0 |
| 2F64 | Inlined | 1,000,000,000 | 0.2990 | 0 | 0 |
| 2F64 | Library | 1,000,000,000 | 0.2979 | 0 | 0 |
| 3F32 | Inlined | 1,000,000,000 | 0.2966 | 0 | 0 |
| 3F32 | Library | 1,000,000,000 | 0.3047 | 0 | 0 |
| 3F64 | Inlined | 1,000,000,000 | 0.2980 | 0 | 0 |
| 3F64 | Library | 1,000,000,000 | 0.3359 | 0 | 0 |

### Magnitude (Length)

| Vector Type | Implementation | Operations/sec | ns/op | Bytes/op | Allocs/op |
|-------------|---------------|---------------|-------|----------|-----------|
| 2F32 | Inlined | 1,000,000,000 | 0.3004 | 0 | 0 |
| 2F32 | Library | 1,000,000,000 | 0.2999 | 0 | 0 |
| 2F64 | Inlined | 1,000,000,000 | 0.2987 | 0 | 0 |
| 2F64 | Library | 1,000,000,000 | 0.2957 | 0 | 0 |
| 3F32 | Inlined | 1,000,000,000 | 0.3004 | 0 | 0 |
| 3F32 | Library | 1,000,000,000 | 0.2994 | 0 | 0 |
| 3F64 | Inlined | 1,000,000,000 | 0.2997 | 0 | 0 |
| 3F64 | Library | 1,000,000,000 | 0.3008 | 0 | 0 |

## Complex Operations

### Angle Between Vectors

| Vector Type | Implementation | Operations/sec | ns/op | Bytes/op | Allocs/op |
|-------------|---------------|---------------|-------|----------|-----------|
| 2F32 | Inlined | 41,315,996 | 28.52 | 0 | 0 |
| 2F32 | WithHelpers | 41,840,582 | 28.49 | 0 | 0 |
| 2F32 | WithLibrary | 42,446,563 | 28.81 | 0 | 0 |
| 2F32 | Original | 42,785,186 | 28.14 | 0 | 0 |
| 2F64 | Inlined | 44,044,164 | 27.53 | 0 | 0 |
| 2F64 | WithHelpers | 43,575,973 | 27.76 | 0 | 0 |
| 2F64 | WithLibrary | 43,318,500 | 28.70 | 0 | 0 |
| 2F64 | Original | 43,531,239 | 27.73 | 0 | 0 |
| 3F32 | Inlined | 40,783,796 | 29.43 | 0 | 0 |
| 3F32 | WithHelpers | 40,754,348 | 29.47 | 0 | 0 |
| 3F32 | WithLibrary | 40,270,104 | 30.03 | 0 | 0 |
| 3F32 | Original | 40,466,319 | 29.68 | 0 | 0 |
| 3F64 | Inlined | 42,194,403 | 29.27 | 0 | 0 |
| 3F64 | WithHelpers | 40,665,880 | 29.25 | 0 | 0 |
| 3F64 | WithLibrary | 41,212,428 | 28.70 | 0 | 0 |
| 3F64 | Original | 42,280,959 | 29.09 | 0 | 0 |

### Normalize and Calculate Angle

| Vector Type | Implementation | Operations/sec | ns/op | Bytes/op | Allocs/op |
|-------------|---------------|---------------|-------|----------|-----------|
| 2F32 | Inlined | 36,659,467 | 32.50 | 0 | 0 |
| 2F32 | WithHelpers | 37,068,900 | 33.15 | 0 | 0 |
| 2F32 | WithLibrary | 37,452,622 | 32.43 | 0 | 0 |
| 2F64 | Inlined | 39,997,599 | 30.45 | 0 | 0 |
| 2F64 | WithHelpers | 39,587,145 | 30.41 | 0 | 0 |
| 2F64 | WithLibrary | 35,035,196 | 34.46 | 0 | 0 |
| 3F32 | Inlined | 34,494,950 | 35.26 | 0 | 0 |
| 3F32 | WithHelpers | 34,821,163 | 34.66 | 0 | 0 |
| 3F32 | WithLibrary | 35,683,932 | 34.13 | 0 | 0 |
| 3F64 | Inlined | 36,159,946 | 33.34 | 0 | 0 |
| 3F64 | WithHelpers | 30,545,706 | 34.25 | 0 | 0 |
| 3F64 | WithLibrary | 33,386,064 | 36.61 | 0 | 0 |

## Analysis

### Simple Operations

For simple operations like dot product and magnitude calculation:

1. **Dot Product**: The performance difference between inlined code and library functions is negligible, with variations of less than 0.04 ns/op for most cases. The only exception is the 3F64 (3D vector with float64) case, where the library function is about 0.038 ns/op slower than the inlined version, which is still a very small difference.

2. **Magnitude**: There is virtually no performance difference between inlined code and library functions, with all implementations performing within 0.005 ns/op of each other.

### Complex Operations

For more complex operations:

1. **Angle Between Vectors**: All four implementations (inlined, with helpers, with library functions, and the original) perform very similarly, with differences of less than 1 ns/op between them. This suggests that the Go compiler is effectively optimizing the code regardless of whether helper functions are used.

2. **Normalize and Calculate Angle**: Again, the performance differences are minimal, with most variations within 1-2 ns/op. The largest difference is for 2F64 vectors using library functions, which is about 4 ns/op slower than the inlined version.

## Conclusion

Based on the benchmark results, we can draw the following conclusions:

1. **Helper Functions vs. Inlining**: Using helper functions instead of inlining code has a negligible impact on performance. The Go compiler appears to be effectively optimizing the code, possibly inlining small helper functions automatically.

2. **Code Maintainability**: Given the minimal performance differences, it would be beneficial to prioritize code maintainability by using helper functions where appropriate. This would reduce code duplication and make the codebase easier to maintain.

3. **Recommendations**:
   - Extract common calculations into helper functions to improve code maintainability
   - Use the existing library functions (like `Dot` and `Len`) in more complex operations
   - Focus on algorithm optimization rather than micro-optimizations like function inlining

These findings suggest that the library can safely extract common calculations into reusable functions without significant performance penalties, which would improve code maintainability and reduce the risk of bugs due to duplicated code.