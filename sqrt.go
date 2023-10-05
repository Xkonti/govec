package govec

import "math"

// V2F

// Sqrt returns a new vector with each component square rooted.
func (v V2F[T]) Sqrt() V2F[T] {
	return V2F[T]{X: T(math.Sqrt(float64(v.X))), Y: T(math.Sqrt(float64(v.Y)))}
}

// SqrtInPlace modifies vector by setting value of each component to the square root.
func (v *V2F[T]) SqrtInPlace() {
	v.X = T(math.Sqrt(float64(v.X)))
	v.Y = T(math.Sqrt(float64(v.Y)))
}

// V3F

// Sqrt returns a new vector with each component square rooted.
func (v V3F[T]) Sqrt() V3F[T] {
	return V3F[T]{X: T(math.Sqrt(float64(v.X))), Y: T(math.Sqrt(float64(v.Y))), Z: T(math.Sqrt(float64(v.Z)))}
}

// SqrtInPlace modifies vector by setting value of each component to the square root.
func (v *V3F[T]) SqrtInPlace() {
	v.X = T(math.Sqrt(float64(v.X)))
	v.Y = T(math.Sqrt(float64(v.Y)))
	v.Z = T(math.Sqrt(float64(v.Z)))
}

// V2I

// Sqrt returns a new V2F with each component square rooted.
func (v V2I[T]) Sqrt() V2F[float64] {
	return V2F[float64]{X: math.Sqrt(float64(v.X)), Y: math.Sqrt(float64(v.Y))}
}

// V3I

// Sqrt returns a new V3F with each component square rooted.
func (v V3I[T]) Sqrt() V3F[float64] {
	return V3F[float64]{X: math.Sqrt(float64(v.X)), Y: math.Sqrt(float64(v.Y)), Z: math.Sqrt(float64(v.Z))}
}
