package govec

// V2F

// Min returns a vector with the minimum of each component of the two vectors.
func (v V2F[T]) Min(v2 V2F[T]) V2F[T] {
	return V2F[T]{X: min(v.X, v2.X), Y: min(v.Y, v2.Y)}
}

// MinInPlace updates the vector in-place with the minimum of each component of the two vectors.
func (v *V2F[T]) MinInPlace(v2 V2F[T]) {
	v.X = min(v.X, v2.X)
	v.Y = min(v.Y, v2.Y)
}

// MinComp returns a vector with the minimum of each component of the vector and the given components.
func (v V2F[T]) MinComp(x T, y T) V2F[T] {
	return V2F[T]{X: min(v.X, x), Y: min(v.Y, y)}
}

// MinCompInPlace updates the vector in-place with the minimum of each component of the vector and the given components.
func (v *V2F[T]) MinCompInPlace(x T, y T) {
	v.X = min(v.X, x)
	v.Y = min(v.Y, y)
}

// V3F

// Min returns a vector with the minimum of each component of the two vectors.
func (v V3F[T]) Min(v2 V3F[T]) V3F[T] {
	return V3F[T]{X: min(v.X, v2.X), Y: min(v.Y, v2.Y), Z: min(v.Z, v2.Z)}
}

// MinInPlace updates the vector in-place with the minimum of each component of the two vectors.
func (v *V3F[T]) MinInPlace(v2 V3F[T]) {
	v.X = min(v.X, v2.X)
	v.Y = min(v.Y, v2.Y)
	v.Z = min(v.Z, v2.Z)
}

// MinComp returns a vector with the minimum of each component of the vector and the given components.
func (v V3F[T]) MinComp(x T, y T, z T) V3F[T] {
	return V3F[T]{X: min(v.X, x), Y: min(v.Y, y), Z: min(v.Z, z)}
}

// MinCompInPlace updates the vector in-place with the minimum of each component of the vector and the given components.
func (v *V3F[T]) MinCompInPlace(x T, y T, z T) {
	v.X = min(v.X, x)
	v.Y = min(v.Y, y)
	v.Z = min(v.Z, z)
}

// V2I

// Min returns a vector with the minimum of each component of the two vectors.
func (v V2I[T]) Min(v2 V2I[T]) V2I[T] {
	return V2I[T]{X: min(v.X, v2.X), Y: min(v.Y, v2.Y)}
}

// MinInPlace updates the vector in-place with the minimum of each component of the two vectors.
func (v *V2I[T]) MinInPlace(v2 V2I[T]) {
	v.X = min(v.X, v2.X)
	v.Y = min(v.Y, v2.Y)
}

// MinComp returns a vector with the minimum of each component of the vector and the given components.
func (v V2I[T]) MinComp(x T, y T) V2I[T] {
	return V2I[T]{X: min(v.X, x), Y: min(v.Y, y)}
}

// MinCompInPlace updates the vector in-place with the minimum of each component of the vector and the given components.
func (v *V2I[T]) MinCompInPlace(x T, y T) {
	v.X = min(v.X, x)
	v.Y = min(v.Y, y)
}

// V3I

// Min returns a vector with the minimum of each component of the two vectors.
func (v V3I[T]) Min(v2 V3I[T]) V3I[T] {
	return V3I[T]{X: min(v.X, v2.X), Y: min(v.Y, v2.Y), Z: min(v.Z, v2.Z)}
}

// MinInPlace updates the vector in-place with the minimum of each component of the two vectors.
func (v *V3I[T]) MinInPlace(v2 V3I[T]) {
	v.X = min(v.X, v2.X)
	v.Y = min(v.Y, v2.Y)
	v.Z = min(v.Z, v2.Z)
}

// MinComp returns a vector with the minimum of each component of the vector and the given components.
func (v V3I[T]) MinComp(x T, y T, z T) V3I[T] {
	return V3I[T]{X: min(v.X, x), Y: min(v.Y, y), Z: min(v.Z, z)}
}

// MinCompInPlace updates the vector in-place with the minimum of each component of the vector and the given components.
func (v *V3I[T]) MinCompInPlace(x T, y T, z T) {
	v.X = min(v.X, x)
	v.Y = min(v.Y, y)
	v.Z = min(v.Z, z)
}
