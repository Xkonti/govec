package govec

func (v V2F[T]) extend(z T) V3F[T] {
	return V3F[T]{X: v.X, Y: v.Y, Z: z}
}

func (v V2I[T]) extend(z T) V3I[T] {
	return V3I[T]{X: v.X, Y: v.Y, Z: z}
}
