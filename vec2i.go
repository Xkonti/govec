package govec

import "golang.org/x/exp/constraints"

// V2I is a 2-dimensional vector with components of type T, where T is an integer type.
type V2I[T constraints.Integer] struct {
	X T
	Y T
}

// ToV2F64 returns a new V2F with the components converted to float64.
func (v V2I[T]) ToV2F64() V2F[float64] {
	return V2F[float64]{X: float64(v.X), Y: float64(v.Y)}
}

// ToV2F32 returns a new V2F with the components converted to float32.
func (v V2I[T]) ToV2F32() V2F[float32] {
	return V2F[float32]{X: float32(v.X), Y: float32(v.Y)}
}

// ToV2I64 returns a new V2I with the components converted to int64.
func (v V2I[T]) ToV2I64() V2I[int64] {
	return V2I[int64]{X: int64(v.X), Y: int64(v.Y)}
}

// ToV2I32 returns a new V2I with the components converted to int32.
func (v V2I[T]) ToV2I32() V2I[int32] {
	return V2I[int32]{X: int32(v.X), Y: int32(v.Y)}
}

// ToV2I16 returns a new V2I with the components converted to int16.
func (v V2I[T]) ToV2I16() V2I[int16] {
	return V2I[int16]{X: int16(v.X), Y: int16(v.Y)}
}

// ToV2I8 returns a new V2I with the components converted to int8.
func (v V2I[T]) ToV2I8() V2I[int8] {
	return V2I[int8]{X: int8(v.X), Y: int8(v.Y)}
}

// ToV2I returns a new V2I with the components converted to int.
func (v V2I[T]) ToV2I() V2I[int] {
	return V2I[int]{X: int(v.X), Y: int(v.Y)}
}
