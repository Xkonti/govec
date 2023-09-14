package govec

func (v V2F[T]) insertX(a T) V3F[T] {
	return V3F[T]{X: a, Y: v.X, Z: v.Y}
}

func (v V2F[T]) insertY(a T) V3F[T] {
	return V3F[T]{X: v.X, Y: a, Z: v.Y}
}

func (v V2F[T]) insertZ(a T) V3F[T] {
	return V3F[T]{X: v.X, Y: v.Y, Z: a}
}

func (v V2I[T]) insertX(a T) V3I[T] {
	return V3I[T]{X: a, Y: v.X, Z: v.Y}
}

func (v V2I[T]) insertY(a T) V3I[T] {
	return V3I[T]{X: v.X, Y: a, Z: v.Y}
}

func (v V2I[T]) insertZ(a T) V3I[T] {
	return V3I[T]{X: v.X, Y: v.Y, Z: a}
}
