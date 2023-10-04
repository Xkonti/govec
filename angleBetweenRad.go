package govec

import (
	"math"
)

// For V2F

// AngleBetweenRad returns the angle between the two vectors in radians.
func (v V2F[T]) AngleBetweenRad(v2 V2F[T]) float64 {
	dotProduct := v.X*v2.X + v.Y*v2.Y
	magnitudeV := math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
	magnitudeV2 := math.Sqrt(float64(v2.X*v2.X + v2.Y*v2.Y))
	cosTheta := float64(dotProduct) / (magnitudeV * magnitudeV2)
	return math.Acos(math.Min(math.Max(cosTheta, -1.0), 1.0))
}

// AngleBetweenRadComp returns the angle between the specified vector and the vector V2F[T]{X: x, Y: y} in radians.
func (v V2F[T]) AngleBetweenRadComp(x T, y T) float64 {
	dotProduct := v.X*x + v.Y*y
	magnitudeV := math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
	magnitudeV2 := math.Sqrt(float64(x*x + y*y))
	cosTheta := float64(dotProduct) / (magnitudeV * magnitudeV2)
	return math.Acos(math.Min(math.Max(cosTheta, -1.0), 1.0))
}

// For V3F

// AngleBetweenRad returns the angle between the two vectors in radians.
func (v V3F[T]) AngleBetweenRad(v2 V3F[T]) float64 {
	dotProduct := v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
	magnitudeV := math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z))
	magnitudeV2 := math.Sqrt(float64(v2.X*v2.X + v2.Y*v2.Y + v2.Z*v2.Z))
	cosTheta := float64(dotProduct) / (magnitudeV * magnitudeV2)
	return math.Acos(math.Min(math.Max(cosTheta, -1.0), 1.0))
}

// AngleBetweenRadComp returns the angle between the specified vector and the vector V3F[T]{X: x, Y: y, Z: z} in radians.
func (v V3F[T]) AngleBetweenRadComp(x T, y T, z T) float64 {
	dotProduct := v.X*x + v.Y*y + v.Z*z
	magnitudeV := math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z))
	magnitudeV2 := math.Sqrt(float64(x*x + y*y + z*z))
	cosTheta := float64(dotProduct) / (magnitudeV * magnitudeV2)
	return math.Acos(math.Min(math.Max(cosTheta, -1.0), 1.0))
}

// For V2I

// AngleBetweenRad returns the angle between the two vectors in radians.
func (v V2I[T]) AngleBetweenRad(v2 V2I[T]) float64 {
	dotProduct := v.X*v2.X + v.Y*v2.Y
	magnitudeV := math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
	magnitudeV2 := math.Sqrt(float64(v2.X*v2.X + v2.Y*v2.Y))
	cosTheta := float64(dotProduct) / (magnitudeV * magnitudeV2)
	return math.Acos(math.Min(math.Max(cosTheta, -1.0), 1.0))
}

// AngleBetweenRadComp returns the angle between the specified vector and the vector V2I[T]{X: x, Y: y} in radians.
func (v V2I[T]) AngleBetweenRadComp(x T, y T) float64 {
	dotProduct := v.X*x + v.Y*y
	magnitudeV := math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
	magnitudeV2 := math.Sqrt(float64(x*x + y*y))
	cosTheta := float64(dotProduct) / (magnitudeV * magnitudeV2)
	return math.Acos(math.Min(math.Max(cosTheta, -1.0), 1.0))
}

// For V3I

// AngleBetweenRad returns the angle between the two vectors in radians.
func (v V3I[T]) AngleBetweenRad(v2 V3I[T]) float64 {
	dotProduct := v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
	magnitudeV := math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z))
	magnitudeV2 := math.Sqrt(float64(v2.X*v2.X + v2.Y*v2.Y + v2.Z*v2.Z))
	cosTheta := float64(dotProduct) / (magnitudeV * magnitudeV2)
	return math.Acos(math.Min(math.Max(cosTheta, -1.0), 1.0))
}

// AngleBetweenRadComp returns the angle between the specified vector and the vector V3I[T]{X: x, Y: y, Z: z} in radians.
func (v V3I[T]) AngleBetweenRadComp(x T, y T, z T) float64 {
	dotProduct := v.X*x + v.Y*y + v.Z*z
	magnitudeV := math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z))
	magnitudeV2 := math.Sqrt(float64(x*x + y*y + z*z))
	cosTheta := float64(dotProduct) / (magnitudeV * magnitudeV2)
	return math.Acos(math.Min(math.Max(cosTheta, -1.0), 1.0))
}
