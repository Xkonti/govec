package govec

import (
	"math"
)

// V2F

func (v V2F[T]) Norm() V2F[T] {
	magnitude := math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
	return V2F[T]{X: v.X / T(magnitude), Y: v.Y / T(magnitude)}
}

func (v *V2F[T]) NormInPlace() {
	magnitude := math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
	v.X /= T(magnitude)
	v.Y /= T(magnitude)
}

// V3F

func (v V3F[T]) Norm() V3F[T] {
	magnitude := math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z))
	return V3F[T]{X: v.X / T(magnitude), Y: v.Y / T(magnitude), Z: v.Z / T(magnitude)}
}

func (v *V3F[T]) NormInPlace() {
	magnitude := math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z))
	v.X /= T(magnitude)
	v.Y /= T(magnitude)
	v.Z /= T(magnitude)
}

// V2I

func (v V2I[T]) Norm() V2F[float64] {
	magnitude := math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
	return V2F[float64]{X: float64(v.X) / magnitude, Y: float64(v.Y) / magnitude}
}

// V3I

func (v V3I[T]) Norm() V3F[float64] {
	magnitude := math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z))
	return V3F[float64]{X: float64(v.X) / magnitude, Y: float64(v.Y) / magnitude, Z: float64(v.Z) / magnitude}
}
