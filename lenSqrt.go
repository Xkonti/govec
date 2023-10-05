package govec

// LenSqrt calculates the length of the vector without the square root.
// This is useful for comparing lengths without the overhead of a square root.
func (v V2F[T]) LenSqrt() T {
	return v.X*v.X + v.Y*v.Y
}

// LenSqrt calculates the length of the vector without the square root.
// This is useful for comparing lengths without the overhead of a square root.
func (v V3F[T]) LenSqrt() T {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// LenSqrt calculates the length of the vector without the square root.
// This is useful for comparing lengths without the overhead of a square root.
func (v V2I[T]) LenSqrt() T {
	return v.X*v.X + v.Y*v.Y
}

// LenSqrt calculates the length of the vector without the square root.
// This is useful for comparing lengths without the overhead of a square root.
func (v V3I[T]) LenSqrt() T {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}
