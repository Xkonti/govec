package govec

// DiscardX returns a new V2F[T] with the X component omitted.
func (v V3F[T]) DiscardX() V2F[T] {
	return V2F[T]{X: v.Y, Y: v.Z}
}

// DiscardY returns a new V2F[T] with the Y component omitted.
func (v V3F[T]) DiscardY() V2F[T] {
	return V2F[T]{X: v.X, Y: v.Z}
}

// DiscardZ returns a new V2F[T] with the Z component omitted.
func (v V3F[T]) DiscardZ() V2F[T] {
	return V2F[T]{X: v.X, Y: v.Y}
}

// DiscardX returns a new V2I[T] with the X component omitted.
func (v V3I[T]) DiscardX() V2I[T] {
	return V2I[T]{X: v.Y, Y: v.Z}
}

// DiscardY returns a new V2I[T] with the Y component omitted.
func (v V3I[T]) DiscardY() V2I[T] {
	return V2I[T]{X: v.X, Y: v.Z}
}

// DiscardZ returns a new V2I[T] with the Z component omitted.
func (v V3I[T]) DiscardZ() V2I[T] {
	return V2I[T]{X: v.X, Y: v.Y}
}
