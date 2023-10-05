package govec

import "golang.org/x/exp/constraints"

// V2FFromSlice returns a new V2F[T] with the given slice as components.
func V2FFromSlice[T constraints.Float](a []T) V2F[T] {
	if len(a) < 2 {
		panic("vector.V2FFromSlice: slice length < 2")
	}
	return V2F[T]{X: a[0], Y: a[1]}
}

// V3FFromSlice returns a new V3F with the given slice as components.
func V3FFromSlice[T constraints.Float](a []T) V3F[T] {
	if len(a) < 3 {
		panic("vector.V3FFromSlice: slice length < 3")
	}
	return V3F[T]{X: a[0], Y: a[1], Z: a[2]}
}

// V2IFromSlice returns a new V2I with the given slice as components.
func V2IFromSlice[T constraints.Integer](a []T) V2I[T] {
	if len(a) < 2 {
		panic("vector.V2IFromSlice: slice length < 2")
	}
	return V2I[T]{X: a[0], Y: a[1]}
}

// V3IFromSlice returns a new V3I with the given slice as components.
func V3IFromSlice[T constraints.Integer](a []T) V3I[T] {
	if len(a) < 3 {
		panic("vector.V3IFromSlice: slice length < 3")
	}
	return V3I[T]{X: a[0], Y: a[1], Z: a[2]}
}

// ToSlice returns a new slice with the components of the vector.
func (v V2F[T]) ToSlice() []T {
	return []T{v.X, v.Y}
}

// ToSlice returns a new slice with the components of the vector.
func (v V3F[T]) ToSlice() []T {
	return []T{v.X, v.Y, v.Z}
}

// ToSlice returns a new slice with the components of the vector.
func (v V2I[T]) ToSlice() []T {
	return []T{v.X, v.Y}
}

// ToSlice returns a new slice with the components of the vector.
func (v V3I[T]) ToSlice() []T {
	return []T{v.X, v.Y, v.Z}
}

// ApplyToSlice sets slice values to the components of the vector.
func (v V2F[T]) ApplyToSlice(a []T) {
	if len(a) < 2 {
		panic("vector.V2F.ApplyToSlice: slice length < 2")
	}
	a[0] = v.X
	a[1] = v.Y
}

// ApplyToSlice sets slice values to the components of the vector.
func (v V3F[T]) ApplyToSlice(a []T) {
	if len(a) < 3 {
		panic("vector.V3F.ApplyToSlice: slice length < 3")
	}
	a[0] = v.X
	a[1] = v.Y
	a[2] = v.Z
}

// ApplyToSlice sets slice values to the components of the vector.
func (v V2I[T]) ApplyToSlice(a []T) {
	if len(a) < 2 {
		panic("vector.V2I.ApplyToSlice: slice length < 2")
	}
	a[0] = v.X
	a[1] = v.Y
}

// ApplyToSlice sets slice values to the components of the vector.
func (v V3I[T]) ApplyToSlice(a []T) {
	if len(a) < 3 {
		panic("vector.V3I.ApplyToSlice: slice length < 3")
	}
	a[0] = v.X
	a[1] = v.Y
	a[2] = v.Z
}
