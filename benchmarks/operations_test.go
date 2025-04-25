package benchmarks

import (
        "math"
        "testing"

        "github.com/xkonti/govec"
        "golang.org/x/exp/constraints"
)

// Dot product benchmarks

// Inlined dot product for 2D vectors
func dotInlined2f[T constraints.Float](v1, v2 govec.V2F[T]) T {
        return v1.X*v2.X + v1.Y*v2.Y
}

// Inlined dot product for 3D vectors
func dotInlined3f[T constraints.Float](v1, v2 govec.V3F[T]) T {
        return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

// Benchmark dot product for 2D vectors with float32
func BenchmarkDot2F32_Inlined(b *testing.B) {
        v1 := govec.V2F[float32]{X: 1.0, Y: 2.0}
        v2 := govec.V2F[float32]{X: 3.0, Y: 4.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = dotInlined2f(v1, v2)
        }
}

func BenchmarkDot2F32_Library(b *testing.B) {
        v1 := govec.V2F[float32]{X: 1.0, Y: 2.0}
        v2 := govec.V2F[float32]{X: 3.0, Y: 4.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = v1.Dot(v2)
        }
}

// Benchmark dot product for 2D vectors with float64
func BenchmarkDot2F64_Inlined(b *testing.B) {
        v1 := govec.V2F[float64]{X: 1.0, Y: 2.0}
        v2 := govec.V2F[float64]{X: 3.0, Y: 4.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = dotInlined2f(v1, v2)
        }
}

func BenchmarkDot2F64_Library(b *testing.B) {
        v1 := govec.V2F[float64]{X: 1.0, Y: 2.0}
        v2 := govec.V2F[float64]{X: 3.0, Y: 4.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = v1.Dot(v2)
        }
}

// Benchmark dot product for 3D vectors with float32
func BenchmarkDot3F32_Inlined(b *testing.B) {
        v1 := govec.V3F[float32]{X: 1.0, Y: 2.0, Z: 3.0}
        v2 := govec.V3F[float32]{X: 4.0, Y: 5.0, Z: 6.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = dotInlined3f(v1, v2)
        }
}

func BenchmarkDot3F32_Library(b *testing.B) {
        v1 := govec.V3F[float32]{X: 1.0, Y: 2.0, Z: 3.0}
        v2 := govec.V3F[float32]{X: 4.0, Y: 5.0, Z: 6.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = v1.Dot(v2)
        }
}

// Benchmark dot product for 3D vectors with float64
func BenchmarkDot3F64_Inlined(b *testing.B) {
        v1 := govec.V3F[float64]{X: 1.0, Y: 2.0, Z: 3.0}
        v2 := govec.V3F[float64]{X: 4.0, Y: 5.0, Z: 6.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = dotInlined3f(v1, v2)
        }
}

func BenchmarkDot3F64_Library(b *testing.B) {
        v1 := govec.V3F[float64]{X: 1.0, Y: 2.0, Z: 3.0}
        v2 := govec.V3F[float64]{X: 4.0, Y: 5.0, Z: 6.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = v1.Dot(v2)
        }
}

// Magnitude benchmarks

// Inlined magnitude for 2D vectors
func magnitudeInlined2f[T constraints.Float](v govec.V2F[T]) float64 {
        return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}

// Inlined magnitude for 3D vectors
func magnitudeInlined3f[T constraints.Float](v govec.V3F[T]) float64 {
        return math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z))
}

// Benchmark magnitude for 2D vectors with float32
func BenchmarkMagnitude2F32_Inlined(b *testing.B) {
        v := govec.V2F[float32]{X: 1.0, Y: 2.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = magnitudeInlined2f(v)
        }
}

func BenchmarkMagnitude2F32_Library(b *testing.B) {
        v := govec.V2F[float32]{X: 1.0, Y: 2.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = v.Len()
        }
}

// Benchmark magnitude for 2D vectors with float64
func BenchmarkMagnitude2F64_Inlined(b *testing.B) {
        v := govec.V2F[float64]{X: 1.0, Y: 2.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = magnitudeInlined2f(v)
        }
}

func BenchmarkMagnitude2F64_Library(b *testing.B) {
        v := govec.V2F[float64]{X: 1.0, Y: 2.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = v.Len()
        }
}

// Benchmark magnitude for 3D vectors with float32
func BenchmarkMagnitude3F32_Inlined(b *testing.B) {
        v := govec.V3F[float32]{X: 1.0, Y: 2.0, Z: 3.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = magnitudeInlined3f(v)
        }
}

func BenchmarkMagnitude3F32_Library(b *testing.B) {
        v := govec.V3F[float32]{X: 1.0, Y: 2.0, Z: 3.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = v.Len()
        }
}

// Benchmark magnitude for 3D vectors with float64
func BenchmarkMagnitude3F64_Inlined(b *testing.B) {
        v := govec.V3F[float64]{X: 1.0, Y: 2.0, Z: 3.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = magnitudeInlined3f(v)
        }
}

func BenchmarkMagnitude3F64_Library(b *testing.B) {
        v := govec.V3F[float64]{X: 1.0, Y: 2.0, Z: 3.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = v.Len()
        }
}