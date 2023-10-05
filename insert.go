package govec

// InsertX returns a new vector with the given component inserted as X.
// Existing X and Y components are shifted to Y and Z.
func (v V2F[T]) InsertX(a T) V3F[T] {
	return V3F[T]{X: a, Y: v.X, Z: v.Y}
}

// InsertY returns a new vector with the given component inserted as Y.
// Existing Y component is shifted to Z.
func (v V2F[T]) InsertY(a T) V3F[T] {
	return V3F[T]{X: v.X, Y: a, Z: v.Y}
}

// InsertZ returns a new vector with the given component inserted as Z.
// Equivalent to Extend.
func (v V2F[T]) InsertZ(a T) V3F[T] {
	return V3F[T]{X: v.X, Y: v.Y, Z: a}
}

// InsertX returns a new vector with the given component inserted as X.
// Existing X and Y components are shifted to Y and Z.
func (v V2I[T]) InsertX(a T) V3I[T] {
	return V3I[T]{X: a, Y: v.X, Z: v.Y}
}

// InsertY returns a new vector with the given component inserted as Y.
// Existing Y component is shifted to Z.
func (v V2I[T]) InsertY(a T) V3I[T] {
	return V3I[T]{X: v.X, Y: a, Z: v.Y}
}

// InsertZ returns a new vector with the given component inserted as Z.
// Equivalent to Extend.
func (v V2I[T]) InsertZ(a T) V3I[T] {
	return V3I[T]{X: v.X, Y: v.Y, Z: a}
}
