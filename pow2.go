package govec

// Pow2 returns a new vector with each component raised to the power of 2.
// Doesn't support fractional exponents.
func (v V2F[T]) Pow2() V2F[T] {
	return V2F[T]{X: v.X * v.X, Y: v.Y * v.Y}
}

// Pow2InPlace modifies vector by setting value of each component to the power of 2.
// Doesn't support fractional exponents.
func (v *V2F[T]) Pow2InPlace() {
	v.X = v.X * v.X
	v.Y = v.Y * v.Y
}

// Pow2 returns a new vector with each component raised to the power of 2.
// Doesn't support fractional exponents.
func (v V3F[T]) Pow2() V3F[T] {
	return V3F[T]{X: v.X * v.X, Y: v.Y * v.Y, Z: v.Z * v.Z}
}

// Pow2InPlace modifies vector by setting value of each component to the power of 2.
// Doesn't support fractional exponents.
func (v *V3F[T]) Pow2InPlace() {
	v.X = v.X * v.X
	v.Y = v.Y * v.Y
	v.Z = v.Z * v.Z
}

// Pow2 returns a new vector with each component raised to the power of 2.
// Doesn't support fractional exponents.
func (v V2I[T]) Pow2() V2I[T] {
	return V2I[T]{X: v.X * v.X, Y: v.Y * v.Y}
}

// Pow2InPlace modifies vector by setting value of each component to the power of 2.
// Doesn't support fractional exponents.
func (v *V2I[T]) Pow2InPlace() {
	v.X = v.X * v.X
	v.Y = v.Y * v.Y
}

// Pow2 returns a new vector with each component raised to the power of 2.
// Doesn't support fractional exponents.
func (v V3I[T]) Pow2() V3I[T] {
	return V3I[T]{X: v.X * v.X, Y: v.Y * v.Y, Z: v.Z * v.Z}
}

// Pow2InPlace modifies vector by setting value of each component to the power of 2.
// Doesn't support fractional exponents.
func (v *V3I[T]) Pow2InPlace() {
	v.X = v.X * v.X
	v.Y = v.Y * v.Y
	v.Z = v.Z * v.Z
}
