package govec

// Neg for V2F
func (v V2F[T]) Neg(v2 V2F[T]) V2F[T] {
	return V2F[T]{X: v.X - v2.X, Y: v.Y - v2.Y}
}

func (v *V2F[T]) NegInPlace(v2 V2F[T]) {
	v.X -= v2.X
	v.Y -= v2.Y
}

// Neg for V3F
func (v V3F[T]) Neg(v2 V3F[T]) V3F[T] {
	return V3F[T]{X: v.X - v2.X, Y: v.Y - v2.Y, Z: v.Z - v2.Z}
}

func (v *V3F[T]) NegInPlace(v2 V3F[T]) {
	v.X -= v2.X
	v.Y -= v2.Y
	v.Z -= v2.Z
}

// Neg for V2I
func (v V2I[T]) Neg(v2 V2I[T]) V2I[T] {
	return V2I[T]{X: v.X - v2.X, Y: v.Y - v2.Y}
}

func (v *V2I[T]) NegInPlace(v2 V2I[T]) {
	v.X -= v2.X
	v.Y -= v2.Y
}

// Neg for V3I
func (v V3I[T]) Neg(v2 V3I[T]) V3I[T] {
	return V3I[T]{X: v.X - v2.X, Y: v.Y - v2.Y, Z: v.Z - v2.Z}
}

func (v *V3I[T]) NegInPlace(v2 V3I[T]) {
	v.X -= v2.X
	v.Y -= v2.Y
	v.Z -= v2.Z
}
