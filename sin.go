package govec

import "math"

// V2F

// Sin calculates the sine of each component of the vector and returns a new vector.
func (v V2F[T]) Sin() V2F[T] {
	return V2F[T]{X: T(math.Sin(float64(v.X))), Y: T(math.Sin(float64(v.Y)))}
}

// SinInPlace modifies vector by calculating the sine of each component of the vector and returns a new vector.
func (v *V2F[T]) SinInPlace() {
	v.X = T(math.Sin(float64(v.X)))
	v.Y = T(math.Sin(float64(v.Y)))
}

// V3F

// Sin calculates the sine of each component of the vector and returns a new vector.
func (v V3F[T]) Sin() V3F[T] {
	return V3F[T]{X: T(math.Sin(float64(v.X))), Y: T(math.Sin(float64(v.Y))), Z: T(math.Sin(float64(v.Z)))}
}

// SinInPlace modifies vector by calculating the sine of each component of the vector and returns a new vector.
func (v *V3F[T]) SinInPlace() {
	v.X = T(math.Sin(float64(v.X)))
	v.Y = T(math.Sin(float64(v.Y)))
	v.Z = T(math.Sin(float64(v.Z)))
}

// V2I

// Sin returns a new V2F by calculating the sine of each component of the vector.
func (v V2I[T]) Sin() V2F[float64] {
	return V2F[float64]{X: math.Sin(float64(v.X)), Y: math.Sin(float64(v.Y))}
}

// V3I

// Sin returns a new V3F by calculating the sine of each component of the vector.
func (v V3I[T]) Sin() V3F[float64] {
	return V3F[float64]{X: math.Sin(float64(v.X)), Y: math.Sin(float64(v.Y)), Z: math.Sin(float64(v.Z))}
}
