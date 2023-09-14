package govec

// Pow3 for V2F
func (v V2F[T]) Pow3() V2F[T] {
	return V2F[T]{X: v.X * v.X * v.X, Y: v.Y * v.Y * v.Y}
}

func (v *V2F[T]) Pow3InPlace() {
	v.X = v.X * v.X * v.X
	v.Y = v.Y * v.Y * v.Y
}

// Pow3 for V3F
func (v V3F[T]) Pow3() V3F[T] {
	return V3F[T]{X: v.X * v.X * v.X, Y: v.Y * v.Y * v.Y, Z: v.Z * v.Z * v.Z}
}

func (v *V3F[T]) Pow3InPlace() {
	v.X = v.X * v.X * v.X
	v.Y = v.Y * v.Y * v.Y
	v.Z = v.Z * v.Z * v.Z
}

// Pow3 for V2I
func (v V2I[T]) Pow3() V2I[T] {
	return V2I[T]{X: v.X * v.X * v.X, Y: v.Y * v.Y * v.Y}
}

func (v *V2I[T]) Pow3InPlace() {
	v.X = v.X * v.X * v.X
	v.Y = v.Y * v.Y * v.Y
}

// Pow3 for V3I
func (v V3I[T]) Pow3() V3I[T] {
	return V3I[T]{X: v.X * v.X * v.X, Y: v.Y * v.Y * v.Y, Z: v.Z * v.Z * v.Z}
}

func (v *V3I[T]) Pow3InPlace() {
	v.X = v.X * v.X * v.X
	v.Y = v.Y * v.Y * v.Y
	v.Z = v.Z * v.Z * v.Z
}
