package govec

// Split returns the components of the vector as separate values.
func (v V2F[T]) split() (T, T) {
	return v.X, v.Y
}

// Split returns the components of the vector as separate values.
func (v V3F[T]) split() (T, T, T) {
	return v.X, v.Y, v.Z
}

// Split returns the components of the vector as separate values.
func (v V2I[T]) split() (T, T) {
	return v.X, v.Y
}

// Split returns the components of the vector as separate values.
func (v V3I[T]) split() (T, T, T) {
	return v.X, v.Y, v.Z
}
