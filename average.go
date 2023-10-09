package govec

// V2F

// Average calculates the average of all components of the vector and returns a T value.
func (v V2F[T]) Average() T {
	return (v.X + v.Y) / 2
}

// V3F

// Average calculates the average of all components of the vector and returns a T value.
func (v V3F[T]) Average() T {
	return (v.X + v.Y + v.Z) / 3
}

// V2I

// Average calculates the average of all components of the vector and returns a float64 value.
func (v V2I[T]) Average() float64 {
	return (float64(v.X) + float64(v.Y)) / 2
}

// V3I

// Average calculates the average of all components of the vector and returns a float64 value.
func (v V3I[T]) Average() float64 {
	return (float64(v.X) + float64(v.Y) + float64(v.Z)) / 3
}
