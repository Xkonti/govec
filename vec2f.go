package govec

import "golang.org/x/exp/constraints"

type V2F[T constraints.Float] struct {
	X T
	Y T
}

func (v V2F[T]) ToV2F64() V2F[float64] {
	return V2F[float64]{X: float64(v.X), Y: float64(v.Y)}
}

func (v V2F[T]) ToV2F32() V2F[float32] {
	return V2F[float32]{X: float32(v.X), Y: float32(v.Y)}
}

func (v V2F[T]) ToV2I64() V2I[int64] {
	return V2I[int64]{X: int64(v.X), Y: int64(v.Y)}
}

func (v V2F[T]) ToV2I32() V2I[int32] {
	return V2I[int32]{X: int32(v.X), Y: int32(v.Y)}
}

func (v V2F[T]) ToV2I16() V2I[int16] {
	return V2I[int16]{X: int16(v.X), Y: int16(v.Y)}
}

func (v V2F[T]) ToV2I8() V2I[int8] {
	return V2I[int8]{X: int8(v.X), Y: int8(v.Y)}
}
