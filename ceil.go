package govec

import "math"

// V2F

// Ceil returns a new V2F[T] with the ceil (round up) of each component.
func (v V2F[T]) Ceil() V2F[T] {
	return V2F[T]{X: T(math.Ceil(float64(v.X))), Y: T(math.Ceil(float64(v.Y)))}
}

// CeilInPlace modifies vector by setting each component to its ceil (round up).
func (v *V2F[T]) CeilInPlace() {
	v.X = T(math.Ceil(float64(v.X)))
	v.Y = T(math.Ceil(float64(v.Y)))
}

// V3F

// Ceil returns a new V3F[T] with the ceil (round up) of each component.
func (v V3F[T]) Ceil() V3F[T] {
	return V3F[T]{X: T(math.Ceil(float64(v.X))), Y: T(math.Ceil(float64(v.Y))), Z: T(math.Ceil(float64(v.Z)))}
}

// CeilInPlace modifies vector by setting each component to its ceil (round up).
func (v *V3F[T]) CeilInPlace() {
	v.X = T(math.Ceil(float64(v.X)))
	v.Y = T(math.Ceil(float64(v.Y)))
	v.Z = T(math.Ceil(float64(v.Z)))
}
