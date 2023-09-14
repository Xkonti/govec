package govec

import "math"

// V2F

func (v V2F[T]) Round() V2F[T] {
	return V2F[T]{X: T(math.Round(float64(v.X))), Y: T(math.Round(float64(v.Y)))}
}

func (v *V2F[T]) RoundInPlace() {
	v.X = T(math.Round(float64(v.X)))
	v.Y = T(math.Round(float64(v.Y)))
}

// V3F

func (v V3F[T]) Round() V3F[T] {
	return V3F[T]{X: T(math.Round(float64(v.X))), Y: T(math.Round(float64(v.Y))), Z: T(math.Round(float64(v.Z)))}
}

func (v *V3F[T]) RoundInPlace() {
	v.X = T(math.Round(float64(v.X)))
	v.Y = T(math.Round(float64(v.Y)))
	v.Z = T(math.Round(float64(v.Z)))
}
