package govec

// V2F

func (v V2F[T]) Dot(v2 V2F[T]) T {
	return v.X*v2.X + v.Y*v2.Y
}

func (v V2F[T]) DotComp(x T, y T) T {
	return v.X*x + v.Y*y
}

// V3F

func (v V3F[T]) Dot(v2 V3F[T]) T {
	return v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
}

func (v V3F[T]) DotComp(x T, y T, z T) T {
	return v.X*x + v.Y*y + v.Z*z
}

// V2I

func (v V2I[T]) Dot(v2 V2I[T]) T {
	return v.X*v2.X + v.Y*v2.Y
}

func (v V2I[T]) DotComp(x T, y T) T {
	return v.X*x + v.Y*y
}

// V3I

func (v V3I[T]) Dot(v2 V3I[T]) T {
	return v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
}

func (v V3I[T]) DotComp(x T, y T, z T) T {
	return v.X*x + v.Y*y + v.Z*z
}
