package govec

// V2F

// MulScalar returns the component-wise product of the vector and a scalar.
func (v V2F[T]) MulScalar(scalar T) V2F[T] {
	return V2F[T]{X: v.X * scalar, Y: v.Y * scalar}
}

// MulScalarInPlace modifies the vector by multiplying it with a scalar.
func (v *V2F[T]) MulScalarInPlace(scalar T) {
	v.X *= scalar
	v.Y *= scalar
}

// V3F

// MulScalar returns the component-wise product of the vector and a scalar.
func (v V3F[T]) MulScalar(scalar T) V3F[T] {
	return V3F[T]{X: v.X * scalar, Y: v.Y * scalar, Z: v.Z * scalar}
}

// MulScalarInPlace modifies the vector by multiplying it with a scalar.
func (v *V3F[T]) MulScalarInPlace(scalar T) {
	v.X *= scalar
	v.Y *= scalar
	v.Z *= scalar
}

// V2I

// MulScalar returns the component-wise product of the vector and a scalar.
func (v V2I[T]) MulScalar(scalar T) V2I[T] {
	return V2I[T]{X: v.X * scalar, Y: v.Y * scalar}
}

// MulScalarInPlace modifies the vector by multiplying it with a scalar.
func (v *V2I[T]) MulScalarInPlace(scalar T) {
	v.X *= scalar
	v.Y *= scalar
}

// V3I

// MulScalar returns the component-wise product of the vector and a scalar.
func (v V3I[T]) MulScalar(scalar T) V3I[T] {
	return V3I[T]{X: v.X * scalar, Y: v.Y * scalar, Z: v.Z * scalar}
}

// MulScalarInPlace modifies the vector by multiplying it with a scalar.
func (v *V3I[T]) MulScalarInPlace(scalar T) {
	v.X *= scalar
	v.Y *= scalar
	v.Z *= scalar
}
