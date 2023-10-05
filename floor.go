package govec

import "math"

// V2F

// Floor returns a new V2F[T] with components rounded down to the nearest integer.
func (v V2F[T]) Floor() V2F[T] {
	return V2F[T]{X: T(math.Floor(float64(v.X))), Y: T(math.Floor(float64(v.Y)))}
}

// FloorInPlace modifies vector by setting components to their rounded down value.
func (v *V2F[T]) FloorInPlace() {
	v.X = T(math.Floor(float64(v.X)))
	v.Y = T(math.Floor(float64(v.Y)))
}

// V3F

// Floor returns a new V3F[T] with components rounded down to the nearest integer.
func (v V3F[T]) Floor() V3F[T] {
	return V3F[T]{X: T(math.Floor(float64(v.X))), Y: T(math.Floor(float64(v.Y))), Z: T(math.Floor(float64(v.Z)))}
}

// FloorInPlace modifies vector by setting components to their rounded down value.
func (v *V3F[T]) FloorInPlace() {
	v.X = T(math.Floor(float64(v.X)))
	v.Y = T(math.Floor(float64(v.Y)))
	v.Z = T(math.Floor(float64(v.Z)))
}
