package govec

// V2F

func (v V2F[T]) Div(v2 V2F[T]) V2F[T] {
	return V2F[T]{X: v.X / v2.X, Y: v.Y / v2.Y}
}

func (v *V2F[T]) DivInPlace(v2 V2F[T]) {
	v.X /= v2.X
	v.Y /= v2.Y
}

func (v V2F[T]) DivComp(x T, y T) V2F[T] {
	return V2F[T]{X: v.X / x, Y: v.Y / y}
}

func (v *V2F[T]) DivCompInPlace(x T, y T) {
	v.X /= x
	v.Y /= y
}

// V3F

func (v V3F[T]) Div(v2 V3F[T]) V3F[T] {
	return V3F[T]{X: v.X / v2.X, Y: v.Y / v2.Y, Z: v.Z / v2.Z}
}

func (v *V3F[T]) DivInPlace(v2 V3F[T]) {
	v.X /= v2.X
	v.Y /= v2.Y
	v.Z /= v2.Z
}

func (v V3F[T]) DivComp(x T, y T, z T) V3F[T] {
	return V3F[T]{X: v.X / x, Y: v.Y / y, Z: v.Z / z}
}

func (v *V3F[T]) DivCompInPlace(x T, y T, z T) {
	v.X /= x
	v.Y /= y
	v.Z /= z
}

// V2I

func (v V2I[T]) Div(v2 V2I[T]) V2I[T] {
	return V2I[T]{X: v.X / v2.X, Y: v.Y / v2.Y}
}

func (v *V2I[T]) DivInPlace(v2 V2I[T]) {
	v.X /= v2.X
	v.Y /= v2.Y
}

func (v V2I[T]) DivComp(x T, y T) V2I[T] {
	return V2I[T]{X: v.X / x, Y: v.Y / y}
}

func (v *V2I[T]) DivCompInPlace(x T, y T) {
	v.X /= x
	v.Y /= y
}

// V3I

func (v V3I[T]) Div(v2 V3I[T]) V3I[T] {
	return V3I[T]{X: v.X / v2.X, Y: v.Y / v2.Y, Z: v.Z / v2.Z}
}

func (v *V3I[T]) DivInPlace(v2 V3I[T]) {
	v.X /= v2.X
	v.Y /= v2.Y
	v.Z /= v2.Z
}

func (v V3I[T]) DivComp(x T, y T, z T) V3I[T] {
	return V3I[T]{X: v.X / x, Y: v.Y / y, Z: v.Z / z}
}

func (v *V3I[T]) DivCompInPlace(x T, y T, z T) {
	v.X /= x
	v.Y /= y
	v.Z /= z
}
