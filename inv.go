package govec

// V2F

func (v V2F[T]) Inv() V2F[T] {
	return V2F[T]{X: 1 / v.X, Y: 1 / v.Y}
}

func (v *V2F[T]) InvInPlace() {
	v.X = 1 / v.X
	v.Y = 1 / v.Y
}

// V3F

func (v V3F[T]) Inv() V3F[T] {
	return V3F[T]{X: 1 / v.X, Y: 1 / v.Y, Z: 1 / v.Z}
}

func (v *V3F[T]) InvInPlace() {
	v.X = 1 / v.X
	v.Y = 1 / v.Y
	v.Z = 1 / v.Z
}

// V2I

func (v V2I[T]) Inv() V2I[T] {
	return V2I[T]{X: 1 / v.X, Y: 1 / v.Y}
}

func (v *V2I[T]) InvInPlace() {
	v.X = 1 / v.X
	v.Y = 1 / v.Y
}

// V3I

func (v V3I[T]) Inv() V3I[T] {
	return V3I[T]{X: 1 / v.X, Y: 1 / v.Y, Z: 1 / v.Z}
}

func (v *V3I[T]) InvInPlace() {
	v.X = 1 / v.X
	v.Y = 1 / v.Y
	v.Z = 1 / v.Z
}
