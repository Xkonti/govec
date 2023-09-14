package govec

func (v V2F[T]) apply(f func(T) T) V2F[T] {
	return V2F[T]{X: f(v.X), Y: f(v.Y)}
}

func (v V3F[T]) apply(f func(T) T) V3F[T] {
	return V3F[T]{X: f(v.X), Y: f(v.Y), Z: f(v.Z)}
}

func (v V2I[T]) apply(f func(T) T) V2I[T] {
	return V2I[T]{X: f(v.X), Y: f(v.Y)}
}

func (v V3I[T]) apply(f func(T) T) V3I[T] {
	return V3I[T]{X: f(v.X), Y: f(v.Y), Z: f(v.Z)}
}
