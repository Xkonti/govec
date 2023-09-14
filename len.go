package govec

import "math"

// V2F

func (v V2F[T]) Len() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}

// V3F

func (v V3F[T]) Len() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z))
}

// V2I

func (v V2I[T]) Len() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}

// V3I

func (v V3I[T]) Len() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z))
}
