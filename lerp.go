package govec

// V2F

// Lerp returns a new V2F[T] that is a linear interpolation between v and v2 using t as the interpolation parameter.
// t = 0 returns v, t = 1 returns v2, t = 0.5 returns the midpoint between v and v2.
func (v V2F[T]) Lerp(v2 V2F[T], t T) V2F[T] {
	return V2F[T]{
		X: v.X + (v2.X-v.X)*t,
		Y: v.Y + (v2.Y-v.Y)*t,
	}
}

// LerpInPlace modifies v by linearly interpolating between v and v2 using t as the interpolation parameter.
// t = 0 leaves v unchanged, t = 1 sets v to v2, t = 0.5 sets v to the midpoint between v and v2.
func (v *V2F[T]) LerpInPlace(v2 V2F[T], t T) {
	v.X = v.X + (v2.X-v.X)*t
	v.Y = v.Y + (v2.Y-v.Y)*t
}

// V3F

// Lerp returns a new V3F[T] that is a linear interpolation between v and v2 using t as the interpolation parameter.
// t = 0 returns v, t = 1 returns v2, t = 0.5 returns the midpoint between v and v2.
func (v V3F[T]) Lerp(v2 V3F[T], t T) V3F[T] {
	return V3F[T]{
		X: v.X + (v2.X-v.X)*t,
		Y: v.Y + (v2.Y-v.Y)*t,
		Z: v.Z + (v2.Z-v.Z)*t,
	}
}

// LerpInPlace modifies v by linearly interpolating between v and v2 using t as the interpolation parameter.
// t = 0 leaves v unchanged, t = 1 sets v to v2, t = 0.5 sets v to the midpoint between v and v2.
func (v *V3F[T]) LerpInPlace(v2 V3F[T], t T) {
	v.X = v.X + (v2.X-v.X)*t
	v.Y = v.Y + (v2.Y-v.Y)*t
	v.Z = v.Z + (v2.Z-v.Z)*t
}

// V2I

// Lerp returns a new V2F[float64] that is a linear interpolation between v and v2 using t as the interpolation parameter.
// t = 0 returns v, t = 1 returns v2, t = 0.5 returns the midpoint between v and v2.
// Note: For integer vectors, Lerp returns a float vector.
func (v V2I[T]) Lerp(v2 V2I[T], t float64) V2F[float64] {
	return V2F[float64]{
		X: float64(v.X) + (float64(v2.X)-float64(v.X))*t,
		Y: float64(v.Y) + (float64(v2.Y)-float64(v.Y))*t,
	}
}

// V3I

// Lerp returns a new V3F[float64] that is a linear interpolation between v and v2 using t as the interpolation parameter.
// t = 0 returns v, t = 1 returns v2, t = 0.5 returns the midpoint between v and v2.
// Note: For integer vectors, Lerp returns a float vector.
func (v V3I[T]) Lerp(v2 V3I[T], t float64) V3F[float64] {
	return V3F[float64]{
		X: float64(v.X) + (float64(v2.X)-float64(v.X))*t,
		Y: float64(v.Y) + (float64(v2.Y)-float64(v.Y))*t,
		Z: float64(v.Z) + (float64(v2.Z)-float64(v.Z))*t,
	}
}
