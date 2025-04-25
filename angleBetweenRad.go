package govec

import (
        "math"
)

// For V2F

// AngleBetweenRad returns the angle between the two vectors in radians.
func (v V2F[T]) AngleBetweenRad(v2 V2F[T]) float64 {
        dotProduct := v.Dot(v2)
        magnitudeV := v.Len()
        magnitudeV2 := v2.Len()
        cosTheta := float64(dotProduct) / (magnitudeV * magnitudeV2)
        return math.Acos(math.Min(math.Max(cosTheta, -1.0), 1.0))
}

// AngleBetweenRadComp returns the angle between the specified vector and the vector V2F[T]{X: x, Y: y} in radians.
func (v V2F[T]) AngleBetweenRadComp(x T, y T) float64 {
        v2 := V2F[T]{X: x, Y: y}
        return v.AngleBetweenRad(v2)
}

// For V3F

// AngleBetweenRad returns the angle between the two vectors in radians.
func (v V3F[T]) AngleBetweenRad(v2 V3F[T]) float64 {
        dotProduct := v.Dot(v2)
        magnitudeV := v.Len()
        magnitudeV2 := v2.Len()
        cosTheta := float64(dotProduct) / (magnitudeV * magnitudeV2)
        return math.Acos(math.Min(math.Max(cosTheta, -1.0), 1.0))
}

// AngleBetweenRadComp returns the angle between the specified vector and the vector V3F[T]{X: x, Y: y, Z: z} in radians.
func (v V3F[T]) AngleBetweenRadComp(x T, y T, z T) float64 {
        v2 := V3F[T]{X: x, Y: y, Z: z}
        return v.AngleBetweenRad(v2)
}

// For V2I

// AngleBetweenRad returns the angle between the two vectors in radians.
func (v V2I[T]) AngleBetweenRad(v2 V2I[T]) float64 {
        dotProduct := v.Dot(v2)
        magnitudeV := v.Len()
        magnitudeV2 := v2.Len()
        cosTheta := float64(dotProduct) / (magnitudeV * magnitudeV2)
        return math.Acos(math.Min(math.Max(cosTheta, -1.0), 1.0))
}

// AngleBetweenRadComp returns the angle between the specified vector and the vector V2I[T]{X: x, Y: y} in radians.
func (v V2I[T]) AngleBetweenRadComp(x T, y T) float64 {
        v2 := V2I[T]{X: x, Y: y}
        return v.AngleBetweenRad(v2)
}

// For V3I

// AngleBetweenRad returns the angle between the two vectors in radians.
func (v V3I[T]) AngleBetweenRad(v2 V3I[T]) float64 {
        dotProduct := v.Dot(v2)
        magnitudeV := v.Len()
        magnitudeV2 := v2.Len()
        cosTheta := float64(dotProduct) / (magnitudeV * magnitudeV2)
        return math.Acos(math.Min(math.Max(cosTheta, -1.0), 1.0))
}

// AngleBetweenRadComp returns the angle between the specified vector and the vector V3I[T]{X: x, Y: y, Z: z} in radians.
func (v V3I[T]) AngleBetweenRadComp(x T, y T, z T) float64 {
        v2 := V3I[T]{X: x, Y: y, Z: z}
        return v.AngleBetweenRad(v2)
}
