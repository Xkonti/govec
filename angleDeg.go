package govec

import (
	"math"
)

// AngleDeg returns a normalized vector from an angle in degrees.
func AngleDeg(a float64) V2F[float64] {
	return V2F[float64]{X: math.Cos(a / radToDegFactor), Y: math.Sin(a / radToDegFactor)}
}
