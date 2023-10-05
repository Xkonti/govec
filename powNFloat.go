package govec

import "math"

// PowNFloat returns a new vector with each component raised to the power of n.
// Supports fractional exponents.
func (v V2F[T]) PowNFloat(n float64) V2F[T] {
	return V2F[T]{X: T(math.Pow(float64(v.X), n)), Y: T(math.Pow(float64(v.Y), n))}
}

// PowNFloatInPlace modifies vector by setting value of each component to the power of n.
// Supports fractional exponents.
func (v *V2F[T]) PowNFloatInPlace(n float64) {
	v.X = T(math.Pow(float64(v.X), n))
	v.Y = T(math.Pow(float64(v.Y), n))
}

// PowNFloat returns a new vector with each component raised to the power of n.
// Supports fractional exponents.
func (v V3F[T]) PowNFloat(n float64) V3F[T] {
	return V3F[T]{X: T(math.Pow(float64(v.X), n)), Y: T(math.Pow(float64(v.Y), n)), Z: T(math.Pow(float64(v.Z), n))}
}

// PowNFloatInPlace modifies vector by setting value of each component to the power of n.
// Supports fractional exponents.
func (v *V3F[T]) PowNFloatInPlace(n float64) {
	v.X = T(math.Pow(float64(v.X), n))
	v.Y = T(math.Pow(float64(v.Y), n))
	v.Z = T(math.Pow(float64(v.Z), n))
}

// PowNFloat returns a new vector with each component raised to the power of n.
// Supports fractional exponents.
func (v V2I[T]) PowNFloat(n float64) V2I[T] {
	return V2I[T]{X: T(math.Pow(float64(v.X), n)), Y: T(math.Pow(float64(v.Y), n))}
}

// PowNFloatInPlace modifies vector by setting value of each component to the power of n.
// Supports fractional exponents.
func (v *V2I[T]) PowNFloatInPlace(n float64) {
	v.X = T(math.Pow(float64(v.X), n))
	v.Y = T(math.Pow(float64(v.Y), n))
}

// PowNFloat returns a new vector with each component raised to the power of n.
func (v V3I[T]) PowNFloat(n float64) V3I[T] {
	return V3I[T]{X: T(math.Pow(float64(v.X), n)), Y: T(math.Pow(float64(v.Y), n)), Z: T(math.Pow(float64(v.Z), n))}
}

// PowNFloatInPlace modifies vector by setting value of each component to the power of n.
func (v *V3I[T]) PowNFloatInPlace(n float64) {
	v.X = T(math.Pow(float64(v.X), n))
	v.Y = T(math.Pow(float64(v.Y), n))
	v.Z = T(math.Pow(float64(v.Z), n))
}
