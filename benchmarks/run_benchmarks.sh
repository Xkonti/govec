#!/bin/bash

# Run all benchmarks and save the results to a file
echo "Running benchmarks..."
go test -bench=. -benchmem ./benchmarks > benchmark_results.txt

# Display the results
echo "Benchmark results:"
cat benchmark_results.txt

# Generate a summary report
echo "Generating summary report..."
echo "# Benchmark Results Summary" > benchmark_summary.md
echo "" >> benchmark_summary.md
echo "## Simple Operations" >> benchmark_summary.md
echo "" >> benchmark_summary.md
echo "### Dot Product" >> benchmark_summary.md
echo "" >> benchmark_summary.md
echo "| Vector Type | Implementation | Operations/sec | Bytes/op | Allocs/op |" >> benchmark_summary.md
echo "|-------------|---------------|---------------|----------|-----------|" >> benchmark_summary.md
grep "BenchmarkDot" benchmark_results.txt | awk '{printf "| %s | %s | %s | %s | %s |\n", $1, $2, $3, $4, $5}' | sed 's/BenchmarkDot//g' | sed 's/_/ /g' >> benchmark_summary.md

echo "" >> benchmark_summary.md
echo "### Magnitude (Length)" >> benchmark_summary.md
echo "" >> benchmark_summary.md
echo "| Vector Type | Implementation | Operations/sec | Bytes/op | Allocs/op |" >> benchmark_summary.md
echo "|-------------|---------------|---------------|----------|-----------|" >> benchmark_summary.md
grep "BenchmarkMagnitude" benchmark_results.txt | awk '{printf "| %s | %s | %s | %s | %s |\n", $1, $2, $3, $4, $5}' | sed 's/BenchmarkMagnitude//g' | sed 's/_/ /g' >> benchmark_summary.md

echo "" >> benchmark_summary.md
echo "## Angle Between Vectors" >> benchmark_summary.md
echo "" >> benchmark_summary.md
echo "| Vector Type | Implementation | Operations/sec | Bytes/op | Allocs/op |" >> benchmark_summary.md
echo "|-------------|---------------|---------------|----------|-----------|" >> benchmark_summary.md
grep "BenchmarkAngleBetweenRad" benchmark_results.txt | awk '{printf "| %s | %s | %s | %s | %s |\n", $1, $2, $3, $4, $5}' | sed 's/BenchmarkAngleBetweenRad//g' | sed 's/_/ /g' >> benchmark_summary.md

echo "" >> benchmark_summary.md
echo "## Complex Operations (Normalize and Calculate Angle)" >> benchmark_summary.md
echo "" >> benchmark_summary.md
echo "| Vector Type | Implementation | Operations/sec | Bytes/op | Allocs/op |" >> benchmark_summary.md
echo "|-------------|---------------|---------------|----------|-----------|" >> benchmark_summary.md
grep "BenchmarkNormalizeAndAngle" benchmark_results.txt | awk '{printf "| %s | %s | %s | %s | %s |\n", $1, $2, $3, $4, $5}' | sed 's/BenchmarkNormalizeAndAngle//g' | sed 's/_/ /g' >> benchmark_summary.md

echo "Summary report generated: benchmark_summary.md"