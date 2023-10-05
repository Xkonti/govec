package govec

import "math"

// V2F

// Pow returns a new vector with each component raised to the power of the corresponding component in v2.
// Supports fractional exponents.
func (v V2F[T]) Pow(v2 V2F[T]) V2F[T] {
	return V2F[T]{X: T(math.Pow(float64(v.X), float64(v2.X))), Y: T(math.Pow(float64(v.Y), float64(v2.Y)))}
}

// PowInPlace modifies vector by setting value of each component to the power of the corresponding component in v2.
// Supports fractional exponents.
func (v *V2F[T]) PowInPlace(v2 V2F[T]) {
	v.X = T(math.Pow(float64(v.X), float64(v2.X)))
	v.Y = T(math.Pow(float64(v.Y), float64(v2.Y)))
}

// PowComp returns a new vector with each component raised to the power of the corresponding values.
// Supports fractional exponents.
func (v V2F[T]) PowComp(x T, y T) V2F[T] {
	return V2F[T]{X: T(math.Pow(float64(v.X), float64(x))), Y: T(math.Pow(float64(v.Y), float64(y)))}
}

// PowCompInPlace modifies vector by setting value of each component to the power of the corresponding values.
// Supports fractional exponents.
func (v *V2F[T]) PowCompInPlace(x T, y T) {
	v.X = T(math.Pow(float64(v.X), float64(x)))
	v.Y = T(math.Pow(float64(v.Y), float64(y)))
}

// V3F

// Pow returns a new vector with each component raised to the power of the corresponding component in v2.
// Supports fractional exponents.
func (v V3F[T]) Pow(v2 V3F[T]) V3F[T] {
	return V3F[T]{X: T(math.Pow(float64(v.X), float64(v2.X))), Y: T(math.Pow(float64(v.Y), float64(v2.Y))), Z: T(math.Pow(float64(v.Z), float64(v2.Z)))}
}

// PowInPlace modifies vector by setting value of each component to the power of the corresponding component in v2.
// Supports fractional exponents.
func (v *V3F[T]) PowInPlace(v2 V3F[T]) {
	v.X = T(math.Pow(float64(v.X), float64(v2.X)))
	v.Y = T(math.Pow(float64(v.Y), float64(v2.Y)))
	v.Z = T(math.Pow(float64(v.Z), float64(v2.Z)))
}

// PowComp returns a new vector with each component raised to the power of the corresponding values.
// Supports fractional exponents.
func (v V3F[T]) PowComp(x T, y T, z T) V3F[T] {
	return V3F[T]{X: T(math.Pow(float64(v.X), float64(x))), Y: T(math.Pow(float64(v.Y), float64(y))), Z: T(math.Pow(float64(v.Z), float64(z)))}
}

// PowCompInPlace modifies vector by setting value of each component to the power of the corresponding values.
// Supports fractional exponents.
func (v *V3F[T]) PowCompInPlace(x T, y T, z T) {
	v.X = T(math.Pow(float64(v.X), float64(x)))
	v.Y = T(math.Pow(float64(v.Y), float64(y)))
	v.Z = T(math.Pow(float64(v.Z), float64(z)))
}

// V2I

// Pow returns a new vector with each component raised to the power of the corresponding component in v2.
// Doesn't support fractional exponents.
func (v V2I[T]) Pow(v2 V2I[T]) V2I[T] {
	return V2I[T]{X: pow(v.X, int16(v2.X)), Y: pow(v.Y, int16(v2.Y))}
}

// PowInPlace modifies vector by setting value of each component to the power of the corresponding component in v2.
// Doesn't support fractional exponents.
func (v *V2I[T]) PowInPlace(v2 V2I[T]) {
	v.X = pow(v.X, int16(v2.X))
	v.Y = pow(v.Y, int16(v2.Y))
}

// PowComp returns a new vector with each component raised to the power of the corresponding values.
// Doesn't support fractional exponents.
func (v V2I[T]) PowComp(x T, y T) V2I[T] {
	return V2I[T]{X: pow(v.X, int16(x)), Y: pow(v.Y, int16(y))}
}

// PowCompInPlace modifies vector by setting value of each component to the power of the corresponding values.
// Doesn't support fractional exponents.
func (v *V2I[T]) PowCompInPlace(x T, y T) {
	v.X = pow(v.X, int16(x))
	v.Y = pow(v.Y, int16(y))
}

// V3I

// Pow returns a new vector with each component raised to the power of the corresponding component in v2.
// Doesn't support fractional exponents.
func (v V3I[T]) Pow(v2 V3I[T]) V3I[T] {
	return V3I[T]{X: pow(v.X, int16(v2.X)), Y: pow(v.Y, int16(v2.Y)), Z: pow(v.Z, int16(v2.Z))}
}

// PowInPlace modifies vector by setting value of each component to the power of the corresponding component in v2.
// Doesn't support fractional exponents.
func (v *V3I[T]) PowInPlace(v2 V3I[T]) {
	v.X = pow(v.X, int16(v2.X))
	v.Y = pow(v.Y, int16(v2.Y))
	v.Z = pow(v.Z, int16(v2.Z))
}

// PowComp returns a new vector with each component raised to the power of the corresponding values.
// Doesn't support fractional exponents.
func (v V3I[T]) PowComp(x T, y T, z T) V3I[T] {
	return V3I[T]{X: pow(v.X, int16(x)), Y: pow(v.Y, int16(y)), Z: pow(v.Z, int16(z))}
}

// PowCompInPlace modifies vector by setting value of each component to the power of the corresponding values.
// Doesn't support fractional exponents.
func (v *V3I[T]) PowCompInPlace(x T, y T, z T) {
	v.X = pow(v.X, int16(x))
	v.Y = pow(v.Y, int16(y))
	v.Z = pow(v.Z, int16(z))
}
