package govec

import "math"

// V2F

func (v V2F[T]) Floor() V2F[T] {
	return V2F[T]{X: T(math.Floor(float64(v.X))), Y: T(math.Floor(float64(v.Y)))}
}

func (v *V2F[T]) FloorInPlace() {
	v.X = T(math.Floor(float64(v.X)))
	v.Y = T(math.Floor(float64(v.Y)))
}

// V3F

func (v V3F[T]) Floor() V3F[T] {
	return V3F[T]{X: T(math.Floor(float64(v.X))), Y: T(math.Floor(float64(v.Y))), Z: T(math.Floor(float64(v.Z)))}
}

func (v *V3F[T]) FloorInPlace() {
	v.X = T(math.Floor(float64(v.X)))
	v.Y = T(math.Floor(float64(v.Y)))
	v.Z = T(math.Floor(float64(v.Z)))
}
