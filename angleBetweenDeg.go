package govec

// For V2F

// AngleBetweenDeg returns the angle between the two vectors in degrees.
func (v V2F[T]) AngleBetweenDeg(v2 V2F[T]) float64 {
	return v.AngleBetweenRad(v2) * radToDegFactor
}

// AngleBetweenDegComp returns the angle between the specified vector and the vector V2F[T]{X: x, Y: y} in degrees.
func (v V2F[T]) AngleBetweenDegComp(x T, y T) float64 {
	return v.AngleBetweenRadComp(x, y) * radToDegFactor
}

// For V3F

// AngleBetweenDeg returns the angle between the two vectors in degrees.
func (v V3F[T]) AngleBetweenDeg(v2 V3F[T]) float64 {
	return v.AngleBetweenRad(v2) * radToDegFactor
}

// AngleBetweenDegComp returns the angle between the specified vector and the vector V3F[T]{X: x, Y: y, Z: z} in degrees.
func (v V3F[T]) AngleBetweenDegComp(x T, y T, z T) float64 {
	return v.AngleBetweenRadComp(x, y, z) * radToDegFactor
}

// For V2I

// AngleBetweenDeg returns the angle between the two vectors in degrees.
func (v V2I[T]) AngleBetweenDeg(v2 V2I[T]) float64 {
	return v.AngleBetweenRad(v2) * radToDegFactor
}

// AngleBetweenDegComp returns the angle between the specified vector and the vector V2I[T]{X: x, Y: y} in degrees.
func (v V2I[T]) AngleBetweenDegComp(x T, y T) float64 {
	return v.AngleBetweenRadComp(x, y) * radToDegFactor
}

// For V3I

// AngleBetweenDeg returns the angle between the two vectors in degrees.
func (v V3I[T]) AngleBetweenDeg(v2 V3I[T]) float64 {
	return v.AngleBetweenRad(v2) * radToDegFactor
}

// AngleBetweenDegComp returns the angle between the specified vector and the vector V3I[T]{X: x, Y: y, Z: z} in degrees.
func (v V3I[T]) AngleBetweenDegComp(x T, y T, z T) float64 {
	return v.AngleBetweenRadComp(x, y, z) * radToDegFactor
}
