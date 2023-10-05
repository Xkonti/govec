package govec

// Neg returns a new V2F[T] with the negated components.
func (v V2F[T]) Neg(v2 V2F[T]) V2F[T] {
	return V2F[T]{X: v.X - v2.X, Y: v.Y - v2.Y}
}

// NegInPlace modifies the vector by negating its components.
func (v *V2F[T]) NegInPlace(v2 V2F[T]) {
	v.X -= v2.X
	v.Y -= v2.Y
}

// Neg returns a new V3F[T] with the negated components.
func (v V3F[T]) Neg(v2 V3F[T]) V3F[T] {
	return V3F[T]{X: v.X - v2.X, Y: v.Y - v2.Y, Z: v.Z - v2.Z}
}

// NegInPlace modifies the vector by negating its components.
func (v *V3F[T]) NegInPlace(v2 V3F[T]) {
	v.X -= v2.X
	v.Y -= v2.Y
	v.Z -= v2.Z
}

// Neg returns a new V2I[T] with the negated components.
func (v V2I[T]) Neg(v2 V2I[T]) V2I[T] {
	return V2I[T]{X: v.X - v2.X, Y: v.Y - v2.Y}
}

// NegInPlace modifies the vector by negating its components.
func (v *V2I[T]) NegInPlace(v2 V2I[T]) {
	v.X -= v2.X
	v.Y -= v2.Y
}

// Neg returns a new V3I[T] with the negated components.
func (v V3I[T]) Neg(v2 V3I[T]) V3I[T] {
	return V3I[T]{X: v.X - v2.X, Y: v.Y - v2.Y, Z: v.Z - v2.Z}
}

// NegInPlace modifies the vector by negating its components.
func (v *V3I[T]) NegInPlace(v2 V3I[T]) {
	v.X -= v2.X
	v.Y -= v2.Y
	v.Z -= v2.Z
}
