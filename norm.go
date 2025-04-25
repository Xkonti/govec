package govec

// V2F

// Norm returns a new vector with the same direction and magnitude 1.
func (v V2F[T]) Norm() V2F[T] {
        magnitude := v.Len()
        return V2F[T]{X: v.X / T(magnitude), Y: v.Y / T(magnitude)}
}

// NormInPlace modifies the vector by normalizing it.
func (v *V2F[T]) NormInPlace() {
        magnitude := v.Len()
        v.X /= T(magnitude)
        v.Y /= T(magnitude)
}

// V3F

// Norm returns a new vector with the same direction and magnitude 1.
func (v V3F[T]) Norm() V3F[T] {
        magnitude := v.Len()
        return V3F[T]{X: v.X / T(magnitude), Y: v.Y / T(magnitude), Z: v.Z / T(magnitude)}
}

// NormInPlace modifies the vector by normalizing it.
func (v *V3F[T]) NormInPlace() {
        magnitude := v.Len()
        v.X /= T(magnitude)
        v.Y /= T(magnitude)
        v.Z /= T(magnitude)
}

// V2I

// Norm returns a new float vector with the same direction and magnitude 1.
func (v V2I[T]) Norm() V2F[float64] {
        magnitude := v.Len()
        return V2F[float64]{X: float64(v.X) / magnitude, Y: float64(v.Y) / magnitude}
}

// V3I

// Norm returns a new float vector with the same direction and magnitude 1.
func (v V3I[T]) Norm() V3F[float64] {
        magnitude := v.Len()
        return V3F[float64]{X: float64(v.X) / magnitude, Y: float64(v.Y) / magnitude, Z: float64(v.Z) / magnitude}
}
