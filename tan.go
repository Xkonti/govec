package govec

import "math"

// V2F

// Tan calculates the tangent of each component of the vector and returns a new vector.
func (v V2F[T]) Tan() V2F[T] {
	return V2F[T]{X: T(math.Tan(float64(v.X))), Y: T(math.Tan(float64(v.Y)))}
}

// TanInPlace modifies vector by calculating the tangent of each component of the vector and returns a new vector.
func (v *V2F[T]) TanInPlace() {
	v.X = T(math.Tan(float64(v.X)))
	v.Y = T(math.Tan(float64(v.Y)))
}

// V3F

// Tan calculates the tangent of each component of the vector and returns a new vector.
func (v V3F[T]) Tan() V3F[T] {
	return V3F[T]{X: T(math.Tan(float64(v.X))), Y: T(math.Tan(float64(v.Y))), Z: T(math.Tan(float64(v.Z)))}
}

// TanInPlace modifies vector by calculating the tangent of each component of the vector and returns a new vector.
func (v *V3F[T]) TanInPlace() {
	v.X = T(math.Tan(float64(v.X)))
	v.Y = T(math.Tan(float64(v.Y)))
	v.Z = T(math.Tan(float64(v.Z)))
}

// V2I

// Tan returns a new V2F by calculating the tangent of each component of the vector.
func (v V2I[T]) Tan() V2F[float64] {
	return V2F[float64]{X: math.Tan(float64(v.X)), Y: math.Tan(float64(v.Y))}
}

// V3I

// Tan returns a new V3F by calculating the tangent of each component of the vector.
func (v V3I[T]) Tan() V3F[float64] {
	return V3F[float64]{X: math.Tan(float64(v.X)), Y: math.Tan(float64(v.Y)), Z: math.Tan(float64(v.Z))}
}
