package govec

// V2F

// AddScalar returns a new V2F[T] with the scalar value added to each component.
func (v V2F[T]) AddScalar(scalar T) V2F[T] {
	return V2F[T]{X: v.X + scalar, Y: v.Y + scalar}
}

// AddScalarInPlace modifies v by adding the scalar value to each component.
func (v *V2F[T]) AddScalarInPlace(scalar T) {
	v.X += scalar
	v.Y += scalar
}

// V3F

// AddScalar returns a new V3F[T] with the scalar value added to each component.
func (v V3F[T]) AddScalar(scalar T) V3F[T] {
	return V3F[T]{X: v.X + scalar, Y: v.Y + scalar, Z: v.Z + scalar}
}

// AddScalarInPlace modifies v by adding the scalar value to each component.
func (v *V3F[T]) AddScalarInPlace(scalar T) {
	v.X += scalar
	v.Y += scalar
	v.Z += scalar
}

// V2I

// AddScalar returns a new V2I[T] with the scalar value added to each component.
func (v V2I[T]) AddScalar(scalar T) V2I[T] {
	return V2I[T]{X: v.X + scalar, Y: v.Y + scalar}
}

// AddScalarInPlace modifies v by adding the scalar value to each component.
func (v *V2I[T]) AddScalarInPlace(scalar T) {
	v.X += scalar
	v.Y += scalar
}

// V3I

// AddScalar returns a new V3I[T] with the scalar value added to each component.
func (v V3I[T]) AddScalar(scalar T) V3I[T] {
	return V3I[T]{X: v.X + scalar, Y: v.Y + scalar, Z: v.Z + scalar}
}

// AddScalarInPlace modifies v by adding the scalar value to each component.
func (v *V3I[T]) AddScalarInPlace(scalar T) {
	v.X += scalar
	v.Y += scalar
	v.Z += scalar
}
