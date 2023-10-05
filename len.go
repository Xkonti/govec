package govec

import "math"

// V2F

// Len calculates the length of the vector.
func (v V2F[T]) Len() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}

// V3F

// Len calculates the length of the vector.
func (v V3F[T]) Len() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z))
}

// V2I

// Len calculates the length of the vector.
func (v V2I[T]) Len() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}

// V3I

// Len calculates the length of the vector.
func (v V3I[T]) Len() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z))
}
