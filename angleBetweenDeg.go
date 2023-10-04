package govec

import (
	"math"
)

// For V2F

// AngleBetweenDeg returns the angle between the two vectors in degrees.
func (v V2F[T]) AngleBetweenDeg(v2 V2F[T]) float64 {
	dotProduct := v.X*v2.X + v.Y*v2.Y
	magnitudeV := math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
	magnitudeV2 := math.Sqrt(float64(v2.X*v2.X + v2.Y*v2.Y))
	cosTheta := float64(dotProduct) / (magnitudeV * magnitudeV2)
	return math.Acos(math.Min(math.Max(cosTheta, -1.0), 1.0)) * 180 / math.Pi
}

// AngleBetweenDegComp returns the angle between the specified vector and the vector V2F[T]{X: x, Y: y} in degrees.
func (v V2F[T]) AngleBetweenDegComp(x T, y T) float64 {
	dotProduct := v.X*x + v.Y*y
	magnitudeV := math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
	magnitudeV2 := math.Sqrt(float64(x*x + y*y))
	cosTheta := float64(dotProduct) / (magnitudeV * magnitudeV2)
	return math.Acos(math.Min(math.Max(cosTheta, -1.0), 1.0)) * 180 / math.Pi
}

// For V3F

// AngleBetweenDeg returns the angle between the two vectors in degrees.
func (v V3F[T]) AngleBetweenDeg(v2 V3F[T]) float64 {
	dotProduct := v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
	magnitudeV := math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z))
	magnitudeV2 := math.Sqrt(float64(v2.X*v2.X + v2.Y*v2.Y + v2.Z*v2.Z))
	cosTheta := float64(dotProduct) / (magnitudeV * magnitudeV2)
	return math.Acos(math.Min(math.Max(cosTheta, -1.0), 1.0)) * 180 / math.Pi
}

// AngleBetweenDegComp returns the angle between the specified vector and the vector V3F[T]{X: x, Y: y, Z: z} in degrees.
func (v V3F[T]) AngleBetweenDegComp(x T, y T, z T) float64 {
	dotProduct := v.X*x + v.Y*y + v.Z*z
	magnitudeV := math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z))
	magnitudeV2 := math.Sqrt(float64(x*x + y*y + z*z))
	cosTheta := float64(dotProduct) / (magnitudeV * magnitudeV2)
	return math.Acos(math.Min(math.Max(cosTheta, -1.0), 1.0)) * 180 / math.Pi
}

// For V2I

// AngleBetweenDeg returns the angle between the two vectors in degrees.
func (v V2I[T]) AngleBetweenDeg(v2 V2I[T]) float64 {
	dotProduct := v.X*v2.X + v.Y*v2.Y
	magnitudeV := math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
	magnitudeV2 := math.Sqrt(float64(v2.X*v2.X + v2.Y*v2.Y))
	cosTheta := float64(dotProduct) / (magnitudeV * magnitudeV2)
	return math.Acos(math.Min(math.Max(cosTheta, -1.0), 1.0)) * 180 / math.Pi
}

// AngleBetweenDegComp returns the angle between the specified vector and the vector V2I[T]{X: x, Y: y} in degrees.
func (v V2I[T]) AngleBetweenDegComp(x T, y T) float64 {
	dotProduct := v.X*x + v.Y*y
	magnitudeV := math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
	magnitudeV2 := math.Sqrt(float64(x*x + y*y))
	cosTheta := float64(dotProduct) / (magnitudeV * magnitudeV2)
	return math.Acos(math.Min(math.Max(cosTheta, -1.0), 1.0)) * 180 / math.Pi
}

// For V3I

// AngleBetweenDeg returns the angle between the two vectors in degrees.
func (v V3I[T]) AngleBetweenDeg(v2 V3I[T]) float64 {
	dotProduct := v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
	magnitudeV := math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z))
	magnitudeV2 := math.Sqrt(float64(v2.X*v2.X + v2.Y*v2.Y + v2.Z*v2.Z))
	cosTheta := float64(dotProduct) / (magnitudeV * magnitudeV2)
	return math.Acos(math.Min(math.Max(cosTheta, -1.0), 1.0)) * 180 / math.Pi
}

// AngleBetweenDegComp returns the angle between the specified vector and the vector V3I[T]{X: x, Y: y, Z: z} in degrees.
func (v V3I[T]) AngleBetweenDegComp(x T, y T, z T) float64 {
	dotProduct := v.X*x + v.Y*y + v.Z*z
	magnitudeV := math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z))
	magnitudeV2 := math.Sqrt(float64(x*x + y*y + z*z))
	cosTheta := float64(dotProduct) / (magnitudeV * magnitudeV2)
	return math.Acos(math.Min(math.Max(cosTheta, -1.0), 1.0)) * 180 / math.Pi
}
