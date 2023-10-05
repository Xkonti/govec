package govec

// V2F

// Dot returns the dot product of given vectors.
func (v V2F[T]) Dot(v2 V2F[T]) T {
	return v.X*v2.X + v.Y*v2.Y
}

// DotComp returns the dot product of given vectors.
func (v V2F[T]) DotComp(x T, y T) T {
	return v.X*x + v.Y*y
}

// V3F

// Dot returns the dot product of given vectors.
func (v V3F[T]) Dot(v2 V3F[T]) T {
	return v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
}

// DotComp returns the dot product of given vectors.
func (v V3F[T]) DotComp(x T, y T, z T) T {
	return v.X*x + v.Y*y + v.Z*z
}

// V2I

// Dot returns the dot product of given vectors.
func (v V2I[T]) Dot(v2 V2I[T]) T {
	return v.X*v2.X + v.Y*v2.Y
}

// DotComp returns the dot product of given vectors.
func (v V2I[T]) DotComp(x T, y T) T {
	return v.X*x + v.Y*y
}

// V3I

// Dot returns the dot product of given vectors.
func (v V3I[T]) Dot(v2 V3I[T]) T {
	return v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
}

// DotComp returns the dot product of given vectors.
func (v V3I[T]) DotComp(x T, y T, z T) T {
	return v.X*x + v.Y*y + v.Z*z
}
