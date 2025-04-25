package govec

// V2F

// Distance returns the norm of the difference between the two vectors.
func (v V2F[T]) Distance(v2 V2F[T]) float64 {
        return v.Sub(v2).Len()
}

// DistanceComp returns the norm of the difference between
// the specified vector and the vector V2F[T]{X: x, Y: y}.
func (v V2F[T]) DistanceComp(x T, y T) float64 {
        return v.SubComp(x, y).Len()
}

// V3F

// Distance returns the norm of the difference between the two vectors.
func (v V3F[T]) Distance(v2 V3F[T]) float64 {
        return v.Sub(v2).Len()
}

// DistanceComp returns the norm of the difference between
// the specified vector and the vector V3F[T]{X: x, Y: y, z: Z}.
func (v V3F[T]) DistanceComp(x T, y T, z T) float64 {
        return v.SubComp(x, y, z).Len()
}

// V2I

// Distance returns the norm of the difference between the two vectors.
func (v V2I[T]) Distance(v2 V2I[T]) float64 {
        return v.Sub(v2).Len()
}

// DistanceComp returns the norm of the difference between
// the specified vector and the vector V2I[T]{X: x, Y: y}.
func (v V2I[T]) DistanceComp(x T, y T) float64 {
        return v.SubComp(x, y).Len()
}

// V3I

// Distance returns the norm of the difference between the two vectors.
func (v V3I[T]) Distance(v2 V3I[T]) float64 {
        return v.Sub(v2).Len()
}

// DistanceComp returns the norm of the difference between
// the specified vector and the vector V3I[T]{X: x, Y: y, Z: z}.
func (v V3I[T]) DistanceComp(x T, y T, z T) float64 {
        return v.SubComp(x, y, z).Len()
}
