package govec

func (v V2F[T]) split() (T, T) {
	return v.X, v.Y
}

func (v V3F[T]) split() (T, T, T) {
	return v.X, v.Y, v.Z
}

func (v V2I[T]) split() (T, T) {
	return v.X, v.Y
}

func (v V3I[T]) split() (T, T, T) {
	return v.X, v.Y, v.Z
}
