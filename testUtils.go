package govec

import (
	"golang.org/x/exp/constraints"
	"math"
)

// almostEqual returns true if the difference between a and b is less than epsilon.
func almostEqual[T constraints.Float](a, b, epsilon T) bool {
	return math.Abs(float64(a-b)) <= float64(epsilon)
}
