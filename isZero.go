package govec

// V2F

// IsZero returns whether the vector is
// a zero vector or not.
func (v V2F[T]) IsZero() bool {
	return v.X == 0 && v.Y == 0
}

// V3F

// IsZero returns whether the vector is
// a zero vector or not.
func (v V3F[T]) IsZero() bool {
	return v.X == 0 && v.Y == 0 && v.Z == 0
}

// V2I

// IsZero returns whether the vector is
// a zero vector or not.
func (v V2I[T]) IsZero() bool {
	return v.X == 0 && v.Y == 0
}

// V3I

// IsZero returns whether the vector is
// a zero vector or not.
func (v V3I[T]) IsZero() bool {
	return v.X == 0 && v.Y == 0 && v.Z == 0
}
