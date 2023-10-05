package govec

import "math"

// V2F

// Round returns a new vector with each component rounded to the nearest integer.
func (v V2F[T]) Round() V2F[T] {
	return V2F[T]{X: T(math.Round(float64(v.X))), Y: T(math.Round(float64(v.Y)))}
}

// RoundInPlace modifies vector by setting value of each component to the nearest integer.
func (v *V2F[T]) RoundInPlace() {
	v.X = T(math.Round(float64(v.X)))
	v.Y = T(math.Round(float64(v.Y)))
}

// V3F

// Round returns a new vector with each component rounded to the nearest integer.
func (v V3F[T]) Round() V3F[T] {
	return V3F[T]{X: T(math.Round(float64(v.X))), Y: T(math.Round(float64(v.Y))), Z: T(math.Round(float64(v.Z)))}
}

// RoundInPlace modifies vector by setting value of each component to the nearest integer.
func (v *V3F[T]) RoundInPlace() {
	v.X = T(math.Round(float64(v.X)))
	v.Y = T(math.Round(float64(v.Y)))
	v.Z = T(math.Round(float64(v.Z)))
}
