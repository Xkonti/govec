package govec

import (
	"golang.org/x/exp/constraints"
	"math"
)

func almostEqual[T constraints.Float](a, b, epsilon T) bool {
	return math.Abs(float64(a-b)) <= float64(epsilon)
}
