package govec

import "math"

// PowNFloat for V2F
func (v V2F[T]) PowNFloat(n float64) V2F[T] {
	return V2F[T]{X: T(math.Pow(float64(v.X), n)), Y: T(math.Pow(float64(v.Y), n))}
}

func (v *V2F[T]) PowNFloatInPlace(n float64) {
	v.X = T(math.Pow(float64(v.X), n))
	v.Y = T(math.Pow(float64(v.Y), n))
}

// PowNFloat for V3F
func (v V3F[T]) PowNFloat(n float64) V3F[T] {
	return V3F[T]{X: T(math.Pow(float64(v.X), n)), Y: T(math.Pow(float64(v.Y), n)), Z: T(math.Pow(float64(v.Z), n))}
}

func (v *V3F[T]) PowNFloatInPlace(n float64) {
	v.X = T(math.Pow(float64(v.X), n))
	v.Y = T(math.Pow(float64(v.Y), n))
	v.Z = T(math.Pow(float64(v.Z), n))
}

// PowNFloat for V2I
func (v V2I[T]) PowNFloat(n float64) V2I[T] {
	return V2I[T]{X: T(math.Pow(float64(v.X), n)), Y: T(math.Pow(float64(v.Y), n))}
}

func (v *V2I[T]) PowNFloatInPlace(n float64) {
	v.X = T(math.Pow(float64(v.X), n))
	v.Y = T(math.Pow(float64(v.Y), n))
}

// PowNFloat for V3I
func (v V3I[T]) PowNFloat(n float64) V3I[T] {
	return V3I[T]{X: T(math.Pow(float64(v.X), n)), Y: T(math.Pow(float64(v.Y), n)), Z: T(math.Pow(float64(v.Z), n))}
}

func (v *V3I[T]) PowNFloatInPlace(n float64) {
	v.X = T(math.Pow(float64(v.X), n))
	v.Y = T(math.Pow(float64(v.Y), n))
	v.Z = T(math.Pow(float64(v.Z), n))
}
