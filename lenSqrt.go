package govec

func (v V2F[T]) LenSqrt() T {
	return v.X*v.X + v.Y*v.Y
}

func (v V3F[T]) LenSqrt() T {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v V2I[T]) LenSqrt() T {
	return v.X*v.X + v.Y*v.Y
}

func (v V3I[T]) LenSqrt() T {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}
