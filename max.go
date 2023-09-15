package govec

// V2F

// Max returns a vector with the maximum of each component of the two vectors.
func (v V2F[T]) Max(v2 V2F[T]) V2F[T] {
	return V2F[T]{X: max(v.X, v2.X), Y: max(v.Y, v2.Y)}
}

// MaxInPlace updates the vector in-place with the maximum of each component of the two vectors.
func (v *V2F[T]) MaxInPlace(v2 V2F[T]) {
	v.X = max(v.X, v2.X)
	v.Y = max(v.Y, v2.Y)
}

// MaxComp returns a vector with the maximum of each component of the vector and the given components.
func (v V2F[T]) MaxComp(x T, y T) V2F[T] {
	return V2F[T]{X: max(v.X, x), Y: max(v.Y, y)}
}

// MaxCompInPlace updates the vector in-place with the maximum of each component of the vector and the given components.
func (v *V2F[T]) MaxCompInPlace(x T, y T) {
	v.X = max(v.X, x)
	v.Y = max(v.Y, y)
}

// V3F

// Max returns a vector with the maximum of each component of the two vectors.
func (v V3F[T]) Max(v2 V3F[T]) V3F[T] {
	return V3F[T]{X: max(v.X, v2.X), Y: max(v.Y, v2.Y), Z: max(v.Z, v2.Z)}
}

// MaxInPlace updates the vector in-place with the maximum of each component of the two vectors.
func (v *V3F[T]) MaxInPlace(v2 V3F[T]) {
	v.X = max(v.X, v2.X)
	v.Y = max(v.Y, v2.Y)
	v.Z = max(v.Z, v2.Z)
}

// MaxComp returns a vector with the maximum of each component of the vector and the given components.
func (v V3F[T]) MaxComp(x T, y T, z T) V3F[T] {
	return V3F[T]{X: max(v.X, x), Y: max(v.Y, y), Z: max(v.Z, z)}
}

// MaxCompInPlace updates the vector in-place with the maximum of each component of the vector and the given components.
func (v *V3F[T]) MaxCompInPlace(x T, y T, z T) {
	v.X = max(v.X, x)
	v.Y = max(v.Y, y)
	v.Z = max(v.Z, z)
}

// V2I

// Max returns a vector with the maximum of each component of the two vectors.
func (v V2I[T]) Max(v2 V2I[T]) V2I[T] {
	return V2I[T]{X: max(v.X, v2.X), Y: max(v.Y, v2.Y)}
}

// MaxInPlace updates the vector in-place with the maximum of each component of the two vectors.
func (v *V2I[T]) MaxInPlace(v2 V2I[T]) {
	v.X = max(v.X, v2.X)
	v.Y = max(v.Y, v2.Y)
}

// MaxComp returns a vector with the maximum of each component of the vector and the given components.
func (v V2I[T]) MaxComp(x T, y T) V2I[T] {
	return V2I[T]{X: max(v.X, x), Y: max(v.Y, y)}
}

// MaxCompInPlace updates the vector in-place with the maximum of each component of the vector and the given components.
func (v *V2I[T]) MaxCompInPlace(x T, y T) {
	v.X = max(v.X, x)
	v.Y = max(v.Y, y)
}

// V3I

// Max returns a vector with the maximum of each component of the two vectors.
func (v V3I[T]) Max(v2 V3I[T]) V3I[T] {
	return V3I[T]{X: max(v.X, v2.X), Y: max(v.Y, v2.Y), Z: max(v.Z, v2.Z)}
}

// MaxInPlace updates the vector in-place with the maximum of each component of the two vectors.
func (v *V3I[T]) MaxInPlace(v2 V3I[T]) {
	v.X = max(v.X, v2.X)
	v.Y = max(v.Y, v2.Y)
	v.Z = max(v.Z, v2.Z)
}

// MaxComp returns a vector with the maximum of each component of the vector and the given components.
func (v V3I[T]) MaxComp(x T, y T, z T) V3I[T] {
	return V3I[T]{X: max(v.X, x), Y: max(v.Y, y), Z: max(v.Z, z)}
}

// MaxCompInPlace updates the vector in-place with the maximum of each component of the vector and the given components.
func (v *V3I[T]) MaxCompInPlace(x T, y T, z T) {
	v.X = max(v.X, x)
	v.Y = max(v.Y, y)
	v.Z = max(v.Z, z)
}
