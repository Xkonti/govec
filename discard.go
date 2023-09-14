package govec

func (v V3F[T]) discardX() V2F[T] {
	return V2F[T]{X: v.Y, Y: v.Z}
}

func (v V3F[T]) discardY() V2F[T] {
	return V2F[T]{X: v.X, Y: v.Z}
}

func (v V3F[T]) discardZ() V2F[T] {
	return V2F[T]{X: v.X, Y: v.Y}
}

func (v V3I[T]) discardX() V2I[T] {
	return V2I[T]{X: v.Y, Y: v.Z}
}

func (v V3I[T]) discardY() V2I[T] {
	return V2I[T]{X: v.X, Y: v.Z}
}

func (v V3I[T]) discardZ() V2I[T] {
	return V2I[T]{X: v.X, Y: v.Y}
}
