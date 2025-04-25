package benchmarks

import (
        "math"
        "testing"

        "github.com/xkonti/govec"
        "golang.org/x/exp/constraints"
)

// Helper functions that could be used instead of inlining
func dot2f[T constraints.Float](v1, v2 govec.V2F[T]) T {
        return v1.X*v2.X + v1.Y*v2.Y
}

func dot3f[T constraints.Float](v1, v2 govec.V3F[T]) T {
        return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

func magnitude2f[T constraints.Float](v govec.V2F[T]) float64 {
        return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}

func magnitude3f[T constraints.Float](v govec.V3F[T]) float64 {
        return math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z))
}

// AngleBetweenRad with helper functions
func angleBetweenRadWithHelpers2f[T constraints.Float](v1, v2 govec.V2F[T]) float64 {
        dotProduct := dot2f(v1, v2)
        magnitudeV1 := magnitude2f(v1)
        magnitudeV2 := magnitude2f(v2)
        cosTheta := float64(dotProduct) / (magnitudeV1 * magnitudeV2)
        return math.Acos(math.Min(math.Max(cosTheta, -1.0), 1.0))
}

// AngleBetweenRad with inlined code (similar to current implementation)
func angleBetweenRadInlined2f[T constraints.Float](v1, v2 govec.V2F[T]) float64 {
        dotProduct := v1.X*v2.X + v1.Y*v2.Y
        magnitudeV1 := math.Sqrt(float64(v1.X*v1.X + v1.Y*v1.Y))
        magnitudeV2 := math.Sqrt(float64(v2.X*v2.X + v2.Y*v2.Y))
        cosTheta := float64(dotProduct) / (magnitudeV1 * magnitudeV2)
        return math.Acos(math.Min(math.Max(cosTheta, -1.0), 1.0))
}

// AngleBetweenRad with library functions
func angleBetweenRadWithLibrary2f[T constraints.Float](v1, v2 govec.V2F[T]) float64 {
        dotProduct := v1.Dot(v2)
        magnitudeV1 := v1.Len()
        magnitudeV2 := v2.Len()
        cosTheta := float64(dotProduct) / (magnitudeV1 * magnitudeV2)
        return math.Acos(math.Min(math.Max(cosTheta, -1.0), 1.0))
}

// AngleBetweenRad with helper functions for 3D vectors
func angleBetweenRadWithHelpers3f[T constraints.Float](v1, v2 govec.V3F[T]) float64 {
        dotProduct := dot3f(v1, v2)
        magnitudeV1 := magnitude3f(v1)
        magnitudeV2 := magnitude3f(v2)
        cosTheta := float64(dotProduct) / (magnitudeV1 * magnitudeV2)
        return math.Acos(math.Min(math.Max(cosTheta, -1.0), 1.0))
}

// AngleBetweenRad with inlined code for 3D vectors
func angleBetweenRadInlined3f[T constraints.Float](v1, v2 govec.V3F[T]) float64 {
        dotProduct := v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
        magnitudeV1 := math.Sqrt(float64(v1.X*v1.X + v1.Y*v1.Y + v1.Z*v1.Z))
        magnitudeV2 := math.Sqrt(float64(v2.X*v2.X + v2.Y*v2.Y + v2.Z*v2.Z))
        cosTheta := float64(dotProduct) / (magnitudeV1 * magnitudeV2)
        return math.Acos(math.Min(math.Max(cosTheta, -1.0), 1.0))
}

// AngleBetweenRad with library functions for 3D vectors
func angleBetweenRadWithLibrary3f[T constraints.Float](v1, v2 govec.V3F[T]) float64 {
        dotProduct := v1.Dot(v2)
        magnitudeV1 := v1.Len()
        magnitudeV2 := v2.Len()
        cosTheta := float64(dotProduct) / (magnitudeV1 * magnitudeV2)
        return math.Acos(math.Min(math.Max(cosTheta, -1.0), 1.0))
}

// Benchmarks for 2D vectors with float32
func BenchmarkAngleBetweenRad2F32_Inlined(b *testing.B) {
        v1 := govec.V2F[float32]{X: 1.0, Y: 2.0}
        v2 := govec.V2F[float32]{X: 3.0, Y: 4.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = angleBetweenRadInlined2f(v1, v2)
        }
}

func BenchmarkAngleBetweenRad2F32_WithHelpers(b *testing.B) {
        v1 := govec.V2F[float32]{X: 1.0, Y: 2.0}
        v2 := govec.V2F[float32]{X: 3.0, Y: 4.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = angleBetweenRadWithHelpers2f(v1, v2)
        }
}

func BenchmarkAngleBetweenRad2F32_WithLibrary(b *testing.B) {
        v1 := govec.V2F[float32]{X: 1.0, Y: 2.0}
        v2 := govec.V2F[float32]{X: 3.0, Y: 4.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = angleBetweenRadWithLibrary2f(v1, v2)
        }
}

func BenchmarkAngleBetweenRad2F32_Original(b *testing.B) {
        v1 := govec.V2F[float32]{X: 1.0, Y: 2.0}
        v2 := govec.V2F[float32]{X: 3.0, Y: 4.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = v1.AngleBetweenRad(v2)
        }
}

// Benchmarks for 2D vectors with float64
func BenchmarkAngleBetweenRad2F64_Inlined(b *testing.B) {
        v1 := govec.V2F[float64]{X: 1.0, Y: 2.0}
        v2 := govec.V2F[float64]{X: 3.0, Y: 4.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = angleBetweenRadInlined2f(v1, v2)
        }
}

func BenchmarkAngleBetweenRad2F64_WithHelpers(b *testing.B) {
        v1 := govec.V2F[float64]{X: 1.0, Y: 2.0}
        v2 := govec.V2F[float64]{X: 3.0, Y: 4.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = angleBetweenRadWithHelpers2f(v1, v2)
        }
}

func BenchmarkAngleBetweenRad2F64_WithLibrary(b *testing.B) {
        v1 := govec.V2F[float64]{X: 1.0, Y: 2.0}
        v2 := govec.V2F[float64]{X: 3.0, Y: 4.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = angleBetweenRadWithLibrary2f(v1, v2)
        }
}

func BenchmarkAngleBetweenRad2F64_Original(b *testing.B) {
        v1 := govec.V2F[float64]{X: 1.0, Y: 2.0}
        v2 := govec.V2F[float64]{X: 3.0, Y: 4.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = v1.AngleBetweenRad(v2)
        }
}

// Benchmarks for 3D vectors with float32
func BenchmarkAngleBetweenRad3F32_Inlined(b *testing.B) {
        v1 := govec.V3F[float32]{X: 1.0, Y: 2.0, Z: 3.0}
        v2 := govec.V3F[float32]{X: 4.0, Y: 5.0, Z: 6.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = angleBetweenRadInlined3f(v1, v2)
        }
}

func BenchmarkAngleBetweenRad3F32_WithHelpers(b *testing.B) {
        v1 := govec.V3F[float32]{X: 1.0, Y: 2.0, Z: 3.0}
        v2 := govec.V3F[float32]{X: 4.0, Y: 5.0, Z: 6.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = angleBetweenRadWithHelpers3f(v1, v2)
        }
}

func BenchmarkAngleBetweenRad3F32_WithLibrary(b *testing.B) {
        v1 := govec.V3F[float32]{X: 1.0, Y: 2.0, Z: 3.0}
        v2 := govec.V3F[float32]{X: 4.0, Y: 5.0, Z: 6.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = angleBetweenRadWithLibrary3f(v1, v2)
        }
}

func BenchmarkAngleBetweenRad3F32_Original(b *testing.B) {
        v1 := govec.V3F[float32]{X: 1.0, Y: 2.0, Z: 3.0}
        v2 := govec.V3F[float32]{X: 4.0, Y: 5.0, Z: 6.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = v1.AngleBetweenRad(v2)
        }
}

// Benchmarks for 3D vectors with float64
func BenchmarkAngleBetweenRad3F64_Inlined(b *testing.B) {
        v1 := govec.V3F[float64]{X: 1.0, Y: 2.0, Z: 3.0}
        v2 := govec.V3F[float64]{X: 4.0, Y: 5.0, Z: 6.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = angleBetweenRadInlined3f(v1, v2)
        }
}

func BenchmarkAngleBetweenRad3F64_WithHelpers(b *testing.B) {
        v1 := govec.V3F[float64]{X: 1.0, Y: 2.0, Z: 3.0}
        v2 := govec.V3F[float64]{X: 4.0, Y: 5.0, Z: 6.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = angleBetweenRadWithHelpers3f(v1, v2)
        }
}

func BenchmarkAngleBetweenRad3F64_WithLibrary(b *testing.B) {
        v1 := govec.V3F[float64]{X: 1.0, Y: 2.0, Z: 3.0}
        v2 := govec.V3F[float64]{X: 4.0, Y: 5.0, Z: 6.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = angleBetweenRadWithLibrary3f(v1, v2)
        }
}

func BenchmarkAngleBetweenRad3F64_Original(b *testing.B) {
        v1 := govec.V3F[float64]{X: 1.0, Y: 2.0, Z: 3.0}
        v2 := govec.V3F[float64]{X: 4.0, Y: 5.0, Z: 6.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = v1.AngleBetweenRad(v2)
        }
}