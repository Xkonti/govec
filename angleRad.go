package govec

import (
	"math"
)

// AngleRad returns a normalized vector from an angle in radians.
func AngleRad(a float64) V2F[float64] {
	return V2F[float64]{X: math.Cos(a), Y: math.Sin(a)}
}
