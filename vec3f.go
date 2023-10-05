package govec

import "golang.org/x/exp/constraints"

// V3F is a 3-dimensional vector with components of type T, where T is a float type.
type V3F[T constraints.Float] struct {
	X T
	Y T
	Z T
}

// ToV3F64 returns a new V3F with the components converted to float64.
func (v V3F[T]) ToV3F64() V3F[float64] {
	return V3F[float64]{X: float64(v.X), Y: float64(v.Y), Z: float64(v.Z)}
}

// ToV3F32 returns a new V3F with the components converted to float32.
func (v V3F[T]) ToV3F32() V3F[float32] {
	return V3F[float32]{X: float32(v.X), Y: float32(v.Y), Z: float32(v.Z)}
}

// ToV3I64 returns a new V3I with the components converted to int64.
func (v V3F[T]) ToV3I64() V3I[int64] {
	return V3I[int64]{X: int64(v.X), Y: int64(v.Y), Z: int64(v.Z)}
}

// ToV3I32 returns a new V3I with the components converted to int32.
func (v V3F[T]) ToV3I32() V3I[int32] {
	return V3I[int32]{X: int32(v.X), Y: int32(v.Y), Z: int32(v.Z)}
}

// ToV3I16 returns a new V3I with the components converted to int16.
func (v V3F[T]) ToV3I16() V3I[int16] {
	return V3I[int16]{X: int16(v.X), Y: int16(v.Y), Z: int16(v.Z)}
}

// ToV3I8 returns a new V3I with the components converted to int8.
func (v V3F[T]) ToV3I8() V3I[int8] {
	return V3I[int8]{X: int8(v.X), Y: int8(v.Y), Z: int8(v.Z)}
}

// ToV3I returns a new V3I with the components converted to int.
func (v V3F[T]) ToV3I() V3I[int] {
	return V3I[int]{X: int(v.X), Y: int(v.Y), Z: int(v.Z)}
}
