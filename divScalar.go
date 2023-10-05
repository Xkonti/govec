package govec

// V2F

// DivScalar returns a new V2F[T] with components divided by scalar.
func (v V2F[T]) DivScalar(scalar T) V2F[T] {
	return V2F[T]{X: v.X / scalar, Y: v.Y / scalar}
}

// DivScalarInPlace modifies vector by setting components to their quotient with scalar.
func (v *V2F[T]) DivScalarInPlace(scalar T) {
	v.X /= scalar
	v.Y /= scalar
}

// V3F

// DivScalar returns a new V3F[T] with components divided by scalar.
func (v V3F[T]) DivScalar(scalar T) V3F[T] {
	return V3F[T]{X: v.X / scalar, Y: v.Y / scalar, Z: v.Z / scalar}
}

// DivScalarInPlace modifies vector by setting components to their quotient with scalar.
func (v *V3F[T]) DivScalarInPlace(scalar T) {
	v.X /= scalar
	v.Y /= scalar
	v.Z /= scalar
}

// V2I

// DivScalar returns a new V2I[T] with components divided by scalar.
func (v V2I[T]) DivScalar(scalar T) V2I[T] {
	return V2I[T]{X: v.X / scalar, Y: v.Y / scalar}
}

// DivScalarInPlace modifies vector by setting components to their quotient with scalar.
func (v *V2I[T]) DivScalarInPlace(scalar T) {
	v.X /= scalar
	v.Y /= scalar
}

// V3I

// DivScalar returns a new V3I[T] with components divided by scalar.
func (v V3I[T]) DivScalar(scalar T) V3I[T] {
	return V3I[T]{X: v.X / scalar, Y: v.Y / scalar, Z: v.Z / scalar}
}

// DivScalarInPlace modifies vector by setting components to their quotient with scalar.
func (v *V3I[T]) DivScalarInPlace(scalar T) {
	v.X /= scalar
	v.Y /= scalar
	v.Z /= scalar
}
