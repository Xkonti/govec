package benchmarks

import (
        "math"
        "testing"

        "github.com/xkonti/govec"
        "golang.org/x/exp/constraints"
)

// Complex operation: Normalize a vector and then calculate the angle between it and another vector

// Using inlined code
func normalizeAndAngleInlined2f[T constraints.Float](v1, v2 govec.V2F[T]) float64 {
        // Normalize v1
        lenV1 := math.Sqrt(float64(v1.X*v1.X + v1.Y*v1.Y))
        normV1 := govec.V2F[T]{
                X: T(float64(v1.X) / lenV1),
                Y: T(float64(v1.Y) / lenV1),
        }
        
        // Calculate angle
        dotProduct := normV1.X*v2.X + normV1.Y*v2.Y
        magnitudeV2 := math.Sqrt(float64(v2.X*v2.X + v2.Y*v2.Y))
        cosTheta := float64(dotProduct) / magnitudeV2
        return math.Acos(math.Min(math.Max(cosTheta, -1.0), 1.0))
}

// Using helper functions
func normalizeAndAngleWithHelpers2f[T constraints.Float](v1, v2 govec.V2F[T]) float64 {
        // Normalize v1
        lenV1 := magnitude2f(v1)
        normV1 := govec.V2F[T]{
                X: T(float64(v1.X) / lenV1),
                Y: T(float64(v1.Y) / lenV1),
        }
        
        // Calculate angle
        dotProduct := dot2f(normV1, v2)
        magnitudeV2 := magnitude2f(v2)
        cosTheta := float64(dotProduct) / magnitudeV2
        return math.Acos(math.Min(math.Max(cosTheta, -1.0), 1.0))
}

// Using library functions
func normalizeAndAngleWithLibrary2f[T constraints.Float](v1, v2 govec.V2F[T]) float64 {
        // Normalize v1
        normV1 := v1.Norm()
        
        // Calculate angle
        return normV1.AngleBetweenRad(v2)
}

// Using inlined code for 3D vectors
func normalizeAndAngleInlined3f[T constraints.Float](v1, v2 govec.V3F[T]) float64 {
        // Normalize v1
        lenV1 := math.Sqrt(float64(v1.X*v1.X + v1.Y*v1.Y + v1.Z*v1.Z))
        normV1 := govec.V3F[T]{
                X: T(float64(v1.X) / lenV1),
                Y: T(float64(v1.Y) / lenV1),
                Z: T(float64(v1.Z) / lenV1),
        }
        
        // Calculate angle
        dotProduct := normV1.X*v2.X + normV1.Y*v2.Y + normV1.Z*v2.Z
        magnitudeV2 := math.Sqrt(float64(v2.X*v2.X + v2.Y*v2.Y + v2.Z*v2.Z))
        cosTheta := float64(dotProduct) / magnitudeV2
        return math.Acos(math.Min(math.Max(cosTheta, -1.0), 1.0))
}

// Using helper functions for 3D vectors
func normalizeAndAngleWithHelpers3f[T constraints.Float](v1, v2 govec.V3F[T]) float64 {
        // Normalize v1
        lenV1 := magnitude3f(v1)
        normV1 := govec.V3F[T]{
                X: T(float64(v1.X) / lenV1),
                Y: T(float64(v1.Y) / lenV1),
                Z: T(float64(v1.Z) / lenV1),
        }
        
        // Calculate angle
        dotProduct := dot3f(normV1, v2)
        magnitudeV2 := magnitude3f(v2)
        cosTheta := float64(dotProduct) / magnitudeV2
        return math.Acos(math.Min(math.Max(cosTheta, -1.0), 1.0))
}

// Using library functions for 3D vectors
func normalizeAndAngleWithLibrary3f[T constraints.Float](v1, v2 govec.V3F[T]) float64 {
        // Normalize v1
        normV1 := v1.Norm()
        
        // Calculate angle
        return normV1.AngleBetweenRad(v2)
}

// Benchmarks for 2D vectors with float32
func BenchmarkNormalizeAndAngle2F32_Inlined(b *testing.B) {
        v1 := govec.V2F[float32]{X: 1.0, Y: 2.0}
        v2 := govec.V2F[float32]{X: 3.0, Y: 4.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = normalizeAndAngleInlined2f(v1, v2)
        }
}

func BenchmarkNormalizeAndAngle2F32_WithHelpers(b *testing.B) {
        v1 := govec.V2F[float32]{X: 1.0, Y: 2.0}
        v2 := govec.V2F[float32]{X: 3.0, Y: 4.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = normalizeAndAngleWithHelpers2f(v1, v2)
        }
}

func BenchmarkNormalizeAndAngle2F32_WithLibrary(b *testing.B) {
        v1 := govec.V2F[float32]{X: 1.0, Y: 2.0}
        v2 := govec.V2F[float32]{X: 3.0, Y: 4.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = normalizeAndAngleWithLibrary2f(v1, v2)
        }
}

// Benchmarks for 2D vectors with float64
func BenchmarkNormalizeAndAngle2F64_Inlined(b *testing.B) {
        v1 := govec.V2F[float64]{X: 1.0, Y: 2.0}
        v2 := govec.V2F[float64]{X: 3.0, Y: 4.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = normalizeAndAngleInlined2f(v1, v2)
        }
}

func BenchmarkNormalizeAndAngle2F64_WithHelpers(b *testing.B) {
        v1 := govec.V2F[float64]{X: 1.0, Y: 2.0}
        v2 := govec.V2F[float64]{X: 3.0, Y: 4.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = normalizeAndAngleWithHelpers2f(v1, v2)
        }
}

func BenchmarkNormalizeAndAngle2F64_WithLibrary(b *testing.B) {
        v1 := govec.V2F[float64]{X: 1.0, Y: 2.0}
        v2 := govec.V2F[float64]{X: 3.0, Y: 4.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = normalizeAndAngleWithLibrary2f(v1, v2)
        }
}

// Benchmarks for 3D vectors with float32
func BenchmarkNormalizeAndAngle3F32_Inlined(b *testing.B) {
        v1 := govec.V3F[float32]{X: 1.0, Y: 2.0, Z: 3.0}
        v2 := govec.V3F[float32]{X: 4.0, Y: 5.0, Z: 6.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = normalizeAndAngleInlined3f(v1, v2)
        }
}

func BenchmarkNormalizeAndAngle3F32_WithHelpers(b *testing.B) {
        v1 := govec.V3F[float32]{X: 1.0, Y: 2.0, Z: 3.0}
        v2 := govec.V3F[float32]{X: 4.0, Y: 5.0, Z: 6.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = normalizeAndAngleWithHelpers3f(v1, v2)
        }
}

func BenchmarkNormalizeAndAngle3F32_WithLibrary(b *testing.B) {
        v1 := govec.V3F[float32]{X: 1.0, Y: 2.0, Z: 3.0}
        v2 := govec.V3F[float32]{X: 4.0, Y: 5.0, Z: 6.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = normalizeAndAngleWithLibrary3f(v1, v2)
        }
}

// Benchmarks for 3D vectors with float64
func BenchmarkNormalizeAndAngle3F64_Inlined(b *testing.B) {
        v1 := govec.V3F[float64]{X: 1.0, Y: 2.0, Z: 3.0}
        v2 := govec.V3F[float64]{X: 4.0, Y: 5.0, Z: 6.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = normalizeAndAngleInlined3f(v1, v2)
        }
}

func BenchmarkNormalizeAndAngle3F64_WithHelpers(b *testing.B) {
        v1 := govec.V3F[float64]{X: 1.0, Y: 2.0, Z: 3.0}
        v2 := govec.V3F[float64]{X: 4.0, Y: 5.0, Z: 6.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = normalizeAndAngleWithHelpers3f(v1, v2)
        }
}

func BenchmarkNormalizeAndAngle3F64_WithLibrary(b *testing.B) {
        v1 := govec.V3F[float64]{X: 1.0, Y: 2.0, Z: 3.0}
        v2 := govec.V3F[float64]{X: 4.0, Y: 5.0, Z: 6.0}
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                _ = normalizeAndAngleWithLibrary3f(v1, v2)
        }
}