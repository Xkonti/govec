package govec

// V2F

// PowN returns a new vector with each component raised to the power of n.
// Doesn't support fractional exponents.
func (v V2F[T]) PowN(n int16) V2F[T] {
	return V2F[T]{X: pow(v.X, n), Y: pow(v.Y, n)}
}

// PowNInPlace modifies vector by setting value of each component to the power of n.
// Doesn't support fractional exponents.
func (v *V2F[T]) PowNInPlace(n int16) {
	v.X = pow(v.X, n)
	v.Y = pow(v.Y, n)
}

// V3F

// PowN returns a new vector with each component raised to the power of n.
// Doesn't support fractional exponents.
func (v V3F[T]) PowN(n int16) V3F[T] {
	return V3F[T]{X: pow(v.X, n), Y: pow(v.Y, n), Z: pow(v.Z, n)}
}

// PowNInPlace modifies vector by setting value of each component to the power of n.
// Doesn't support fractional exponents.
func (v *V3F[T]) PowNInPlace(n int16) {
	v.X = pow(v.X, n)
	v.Y = pow(v.Y, n)
	v.Z = pow(v.Z, n)
}

// V2I

// PowN returns a new vector with each component raised to the power of n.
// Doesn't support fractional exponents.
func (v V2I[T]) PowN(n int16) V2I[T] {
	return V2I[T]{X: pow(v.X, n), Y: pow(v.Y, n)}
}

// PowNInPlace modifies vector by setting value of each component to the power of n.
// Doesn't support fractional exponents.
func (v *V2I[T]) PowNInPlace(n int16) {
	v.X = pow(v.X, n)
	v.Y = pow(v.Y, n)
}

// V3I

// PowN returns a new vector with each component raised to the power of n.
// Doesn't support fractional exponents.
func (v V3I[T]) PowN(n int16) V3I[T] {
	return V3I[T]{X: pow(v.X, n), Y: pow(v.Y, n), Z: pow(v.Z, n)}
}

// PowNInPlace modifies vector by setting value of each component to the power of n.
// Doesn't support fractional exponents.
func (v *V3I[T]) PowNInPlace(n int16) {
	v.X = pow(v.X, n)
	v.Y = pow(v.Y, n)
	v.Z = pow(v.Z, n)
}
