package govec

// Extend returns a new vector with the given component appended as Z.
func (v V2F[T]) Extend(z T) V3F[T] {
	return V3F[T]{X: v.X, Y: v.Y, Z: z}
}

// Extend returns a new vector with the given component appended as Z.
func (v V2I[T]) Extend(z T) V3I[T] {
	return V3I[T]{X: v.X, Y: v.Y, Z: z}
}
