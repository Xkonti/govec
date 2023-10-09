package govec

import "math"

// V2F

// Cos calculates the cosine of each component of the vector and returns a new vector.
func (v V2F[T]) Cos() V2F[T] {
	return V2F[T]{X: T(math.Cos(float64(v.X))), Y: T(math.Cos(float64(v.Y)))}
}

// CosInPlace modifies vector by calculating the cosine of each component of the vector and returns a new vector.
func (v *V2F[T]) CosInPlace() {
	v.X = T(math.Cos(float64(v.X)))
	v.Y = T(math.Cos(float64(v.Y)))
}

// V3F

// Cos calculates the cosine of each component of the vector and returns a new vector.
func (v V3F[T]) Cos() V3F[T] {
	return V3F[T]{X: T(math.Cos(float64(v.X))), Y: T(math.Cos(float64(v.Y))), Z: T(math.Cos(float64(v.Z)))}
}

// CosInPlace modifies vector by calculating the cosine of each component of the vector and returns a new vector.
func (v *V3F[T]) CosInPlace() {
	v.X = T(math.Cos(float64(v.X)))
	v.Y = T(math.Cos(float64(v.Y)))
	v.Z = T(math.Cos(float64(v.Z)))
}

// V2I

// Cos returns a new V2F by calculating the cosine of each component of the vector.
func (v V2I[T]) Cos() V2F[float64] {
	return V2F[float64]{X: math.Cos(float64(v.X)), Y: math.Cos(float64(v.Y))}
}

// V3I

// Cos returns a new V3F by calculating the cosine of each component of the vector.
func (v V3I[T]) Cos() V3F[float64] {
	return V3F[float64]{X: math.Cos(float64(v.X)), Y: math.Cos(float64(v.Y)), Z: math.Cos(float64(v.Z))}
}
