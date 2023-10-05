package govec

// V2F

// ModScalar returns a new V2F[T] with same modulus applied to each component.
func (v V2F[T]) ModScalar(scalar T) V2F[T] {
	return V2F[T]{X: v.X - scalar*T(int64(v.X/scalar)), Y: v.Y - scalar*T(int64(v.Y/scalar))}
}

// ModScalarInPlace updates the vector in-place with same modulus applied to each component.
func (v *V2F[T]) ModScalarInPlace(scalar T) {
	v.X = v.X - scalar*T(int64(v.X/scalar))
	v.Y = v.Y - scalar*T(int64(v.Y/scalar))
}

// V3F

// ModScalar returns a new V3F[T] with same modulus applied to each component.
func (v V3F[T]) ModScalar(scalar T) V3F[T] {
	return V3F[T]{X: v.X - scalar*T(int64(v.X/scalar)), Y: v.Y - scalar*T(int64(v.Y/scalar)), Z: v.Z - scalar*T(int64(v.Z/scalar))}
}

// ModScalarInPlace updates the vector in-place with same modulus applied to each component.
func (v *V3F[T]) ModScalarInPlace(scalar T) {
	v.X = v.X - scalar*T(int64(v.X/scalar))
	v.Y = v.Y - scalar*T(int64(v.Y/scalar))
	v.Z = v.Z - scalar*T(int64(v.Z/scalar))
}

// V2I

// ModScalar returns a new V2I[T] with same modulus applied to each component.
func (v V2I[T]) ModScalar(scalar T) V2I[T] {
	return V2I[T]{X: v.X % scalar, Y: v.Y % scalar}
}

// ModScalarInPlace updates the vector in-place with same modulus applied to each component.
func (v *V2I[T]) ModScalarInPlace(scalar T) {
	v.X %= scalar
	v.Y %= scalar
}

// V3I

// ModScalar returns a new V3I[T] with same modulus applied to each component.
func (v V3I[T]) ModScalar(scalar T) V3I[T] {
	return V3I[T]{X: v.X % scalar, Y: v.Y % scalar, Z: v.Z % scalar}
}

// ModScalarInPlace updates the vector in-place with same modulus applied to each component.
func (v *V3I[T]) ModScalarInPlace(scalar T) {
	v.X %= scalar
	v.Y %= scalar
	v.Z %= scalar
}
