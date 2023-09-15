package govec

// Apply applies a function to each component of a vector.
// The function passed to Apply takes the component value and index as arguments.
func (v V2F[T]) Apply(f func(compVal T, compIdx int8) T) V2F[T] {
	return V2F[T]{X: f(v.X, 0), Y: f(v.Y, 1)}
}

// Apply applies a function to each component of a vector.
// The function passed to Apply takes the component value and index as arguments.
func (v V3F[T]) Apply(f func(compVal T, compIdx int8) T) V3F[T] {
	return V3F[T]{X: f(v.X, 0), Y: f(v.Y, 1), Z: f(v.Z, 2)}
}

// Apply applies a function to each component of a vector.
// The function passed to Apply takes the component value and index as arguments.
func (v V2I[T]) Apply(f func(compVal T, compIdx int8) T) V2I[T] {
	return V2I[T]{X: f(v.X, 0), Y: f(v.Y, 1)}
}

// Apply applies a function to each component of a vector.
// The function passed to Apply takes the component value and index as arguments.
func (v V3I[T]) Apply(f func(compVal T, compIdx int8) T) V3I[T] {
	return V3I[T]{X: f(v.X, 0), Y: f(v.Y, 1), Z: f(v.Z, 2)}
}
