package govec

import (
	"math"
)

// RotateRad returns the vector obtained by rotating
// the specified 2D vector by the angle specified in
// radians.
func (v V2F[T]) RotateRad(angleRad float64) V2F[float64] {
	return V2F[float64]{
		X: float64(v.X)*math.Cos(angleRad) - float64(v.Y)*math.Sin(angleRad),
		Y: float64(v.X)*math.Sin(angleRad) + float64(v.Y)*math.Cos(angleRad),
	}
}

// RotateRad returns the vector obtained by rotating
// the specified 2D vector by the angle specified in
// radians.
func (v V2I[T]) RotateRad(angleRad float64) V2F[float64] {
	return V2F[float64]{
		X: float64(v.X)*math.Cos(angleRad) - float64(v.Y)*math.Sin(angleRad),
		Y: float64(v.X)*math.Sin(angleRad) + float64(v.Y)*math.Cos(angleRad),
	}
}
