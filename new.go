package govec

import "golang.org/x/exp/constraints"

// NEW

// NewV2F returns a new V2F[T] with the given components.
func NewV2F[T constraints.Float](x, y T) V2F[T] {
	return V2F[T]{X: x, Y: y}
}

// NewV3F returns a new V3F[T] with the given components.
func NewV3F[T constraints.Float](x, y, z T) V3F[T] {
	return V3F[T]{X: x, Y: y, Z: z}
}

// NewV2I returns a new V2I[T] with the given components.
func NewV2I[T constraints.Integer](x, y T) V2I[T] {
	return V2I[T]{X: x, Y: y}
}

// NewV3I returns a new V3I[T] with the given components.
func NewV3I[T constraints.Integer](x, y, z T) V3I[T] {
	return V3I[T]{X: x, Y: y, Z: z}
}

// Fill

// FillV2F returns a new V2F[T] with the components set to the provided value.
func FillV2F[T constraints.Float](value T) V2F[T] {
	return V2F[T]{X: value, Y: value}
}

// FillV3F returns a new V3F[T] with the components set to the provided value.
func FillV3F[T constraints.Float](value T) V3F[T] {
	return V3F[T]{X: value, Y: value, Z: value}
}

// FillV2I returns a new V2I[T] with the components set to the provided value.
func FillV2I[T constraints.Integer](value T) V2I[T] {
	return V2I[T]{X: value, Y: value}
}

// FillV3I returns a new V3I[T] with the components set to the provided value.
func FillV3I[T constraints.Integer](value T) V3I[T] {
	return V3I[T]{X: value, Y: value, Z: value}
}

// ZERO

// ZeroV2F returns a new V2F[T] with the components set to zero.
func ZeroV2F[T constraints.Float]() V2F[T] {
	return V2F[T]{X: 0.0, Y: 0.0}
}

// ZeroV3F returns a new V3F[T] with the components set to zero.
func ZeroV3F[T constraints.Float]() V3F[T] {
	return V3F[T]{X: 0.0, Y: 0.0, Z: 0.0}
}

// ZeroV2I returns a new V2I[T] with the components set to zero.
func ZeroV2I[T constraints.Integer]() V2I[T] {
	return V2I[T]{X: 0, Y: 0}
}

// ZeroV3I returns a new V3I[T] with the components set to zero.
func ZeroV3I[T constraints.Integer]() V3I[T] {
	return V3I[T]{X: 0, Y: 0, Z: 0}
}

// ONE

// OneV2F returns a new V2F[T] with the components set to one.
func OneV2F[T constraints.Float]() V2F[T] {
	return V2F[T]{X: 1.0, Y: 1.0}
}

// OneV3F returns a new V3F[T] with the components set to one.
func OneV3F[T constraints.Float]() V3F[T] {
	return V3F[T]{X: 1.0, Y: 1.0, Z: 1.0}
}

// OneV2I returns a new V2I[T] with the components set to one.
func OneV2I[T constraints.Integer]() V2I[T] {
	return V2I[T]{X: 1, Y: 1}
}

// OneV3I returns a new V3I[T] with the components set to one.
func OneV3I[T constraints.Integer]() V3I[T] {
	return V3I[T]{X: 1, Y: 1, Z: 1}
}
