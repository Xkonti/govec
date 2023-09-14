package govec

import "golang.org/x/exp/constraints"

type V3F[T constraints.Float] struct {
	X T
	Y T
	Z T
}

func (v V3F[T]) ToV3F64() V3F[float64] {
	return V3F[float64]{X: float64(v.X), Y: float64(v.Y), Z: float64(v.Z)}
}

func (v V3F[T]) ToV3F32() V3F[float32] {
	return V3F[float32]{X: float32(v.X), Y: float32(v.Y), Z: float32(v.Z)}
}

func (v V3F[T]) ToV3I64() V3I[int64] {
	return V3I[int64]{X: int64(v.X), Y: int64(v.Y), Z: int64(v.Z)}
}

func (v V3F[T]) ToV3I32() V3I[int32] {
	return V3I[int32]{X: int32(v.X), Y: int32(v.Y), Z: int32(v.Z)}
}

func (v V3F[T]) ToV3I16() V3I[int16] {
	return V3I[int16]{X: int16(v.X), Y: int16(v.Y), Z: int16(v.Z)}
}

func (v V3F[T]) ToV3I8() V3I[int8] {
	return V3I[int8]{X: int8(v.X), Y: int8(v.Y), Z: int8(v.Z)}
}
