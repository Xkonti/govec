package govec

import "golang.org/x/exp/constraints"

// NEW

func NewV2F[T constraints.Float](x, y T) V2F[T] {
	return V2F[T]{X: x, Y: y}
}

func NewV3F[T constraints.Float](x, y, z T) V3F[T] {
	return V3F[T]{X: x, Y: y, Z: z}
}

func NewV2I[T constraints.Integer](x, y T) V2I[T] {
	return V2I[T]{X: x, Y: y}
}

func NewV3I[T constraints.Integer](x, y, z T) V3I[T] {
	return V3I[T]{X: x, Y: y, Z: z}
}

// Fill

func FillV2F[T constraints.Float](value T) V2F[T] {
	return V2F[T]{X: value, Y: value}
}

func FillV3F[T constraints.Float](value T) V3F[T] {
	return V3F[T]{X: value, Y: value, Z: value}
}

func FillV2I[T constraints.Integer](value T) V2I[T] {
	return V2I[T]{X: value, Y: value}
}

func FillV3I[T constraints.Integer](value T) V3I[T] {
	return V3I[T]{X: value, Y: value, Z: value}
}

// ZERO

func ZeroV2F[T constraints.Float]() V2F[T] {
	return V2F[T]{X: 0.0, Y: 0.0}
}

func ZeroV3F[T constraints.Float]() V3F[T] {
	return V3F[T]{X: 0.0, Y: 0.0, Z: 0.0}
}

func ZeroV2I[T constraints.Integer]() V2I[T] {
	return V2I[T]{X: 0, Y: 0}
}

func ZeroV3I[T constraints.Integer]() V3I[T] {
	return V3I[T]{X: 0, Y: 0, Z: 0}
}

// ONE

func OneV2F[T constraints.Float]() V2F[T] {
	return V2F[T]{X: 1.0, Y: 1.0}
}

func OneV3F[T constraints.Float]() V3F[T] {
	return V3F[T]{X: 1.0, Y: 1.0, Z: 1.0}
}

func OneV2I[T constraints.Integer]() V2I[T] {
	return V2I[T]{X: 1, Y: 1}
}

func OneV3I[T constraints.Integer]() V3I[T] {
	return V3I[T]{X: 1, Y: 1, Z: 1}
}
