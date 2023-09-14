package govec

// V2F

func (v V2F[T]) MulScalar(scalar T) V2F[T] {
	return V2F[T]{X: v.X * scalar, Y: v.Y * scalar}
}

func (v *V2F[T]) MulScalarInPlace(scalar T) {
	v.X *= scalar
	v.Y *= scalar
}

// V3F

func (v V3F[T]) MulScalar(scalar T) V3F[T] {
	return V3F[T]{X: v.X * scalar, Y: v.Y * scalar, Z: v.Z * scalar}
}

func (v *V3F[T]) MulScalarInPlace(scalar T) {
	v.X *= scalar
	v.Y *= scalar
	v.Z *= scalar
}

// V2I

func (v V2I[T]) MulScalar(scalar T) V2I[T] {
	return V2I[T]{X: v.X * scalar, Y: v.Y * scalar}
}

func (v *V2I[T]) MulScalarInPlace(scalar T) {
	v.X *= scalar
	v.Y *= scalar
}

// V3I

func (v V3I[T]) MulScalar(scalar T) V3I[T] {
	return V3I[T]{X: v.X * scalar, Y: v.Y * scalar, Z: v.Z * scalar}
}

func (v *V3I[T]) MulScalarInPlace(scalar T) {
	v.X *= scalar
	v.Y *= scalar
	v.Z *= scalar
}
