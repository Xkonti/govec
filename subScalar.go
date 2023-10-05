package govec

// V2F

// SubScalar returns a new vector with each component subtracted by scalar.
func (v V2F[T]) SubScalar(scalar T) V2F[T] {
	return V2F[T]{X: v.X - scalar, Y: v.Y - scalar}
}

// SubScalarInPlace modifies vector by subtracting each component by scalar.
func (v *V2F[T]) SubScalarInPlace(scalar T) {
	v.X -= scalar
	v.Y -= scalar
}

// V3F

// SubScalar returns a new vector with each component subtracted by scalar.
func (v V3F[T]) SubScalar(scalar T) V3F[T] {
	return V3F[T]{X: v.X - scalar, Y: v.Y - scalar, Z: v.Z - scalar}
}

// SubScalarInPlace modifies vector by subtracting each component by scalar.
func (v *V3F[T]) SubScalarInPlace(scalar T) {
	v.X -= scalar
	v.Y -= scalar
	v.Z -= scalar
}

// V2I

// SubScalar returns a new vector with each component subtracted by scalar.
func (v V2I[T]) SubScalar(scalar T) V2I[T] {
	return V2I[T]{X: v.X - scalar, Y: v.Y - scalar}
}

// SubScalarInPlace modifies vector by subtracting each component by scalar.
func (v *V2I[T]) SubScalarInPlace(scalar T) {
	v.X -= scalar
	v.Y -= scalar
}

// V3I

// SubScalar returns a new vector with each component subtracted by scalar.
func (v V3I[T]) SubScalar(scalar T) V3I[T] {
	return V3I[T]{X: v.X - scalar, Y: v.Y - scalar, Z: v.Z - scalar}
}

// SubScalarInPlace modifies vector by subtracting each component by scalar.
func (v *V3I[T]) SubScalarInPlace(scalar T) {
	v.X -= scalar
	v.Y -= scalar
	v.Z -= scalar
}
