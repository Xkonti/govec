package govec

import "golang.org/x/exp/constraints"

// V2FFromArray returns a new V2F[T] from the specified array.
func V2FFromArray[T constraints.Float](a [2]T) V2F[T] {
	return V2F[T]{X: a[0], Y: a[1]}
}

// V3FFromArray returns a new V3F[T] from the specified array.
func V3FFromArray[T constraints.Float](a [3]T) V3F[T] {
	return V3F[T]{X: a[0], Y: a[1], Z: a[2]}
}

// V2IFromArray returns a new V2I[T] from the specified array.
func V2IFromArray[T constraints.Integer](a [2]T) V2I[T] {
	return V2I[T]{X: a[0], Y: a[1]}
}

// V3IFromArray returns a new V3I[T] from the specified array.
func V3IFromArray[T constraints.Integer](a [3]T) V3I[T] {
	return V3I[T]{X: a[0], Y: a[1], Z: a[2]}
}

// ToArray returns an array of the components of the vector.
func (v V2F[T]) ToArray() [2]T {
	return [2]T{v.X, v.Y}
}

// ToArray returns an array of the components of the vector.
func (v V3F[T]) ToArray() [3]T {
	return [3]T{v.X, v.Y, v.Z}
}

// ToArray returns an array of the components of the vector.
func (v V2I[T]) ToArray() [2]T {
	return [2]T{v.X, v.Y}
}

// ToArray returns an array of the components of the vector.
func (v V3I[T]) ToArray() [3]T {
	return [3]T{v.X, v.Y, v.Z}
}

// ApplyToArray applies the components of the vector to the specified array.
func (v V2F[T]) ApplyToArray(a *[2]T) {
	a[0] = v.X
	a[1] = v.Y
}

// ApplyToArray applies the components of the vector to the specified array.
func (v V3F[T]) ApplyToArray(a *[3]T) {
	a[0] = v.X
	a[1] = v.Y
	a[2] = v.Z
}

// ApplyToArray applies the components of the vector to the specified array.
func (v V2I[T]) ApplyToArray(a *[2]T) {
	a[0] = v.X
	a[1] = v.Y
}

// ApplyToArray applies the components of the vector to the specified array.
func (v V3I[T]) ApplyToArray(a *[3]T) {
	a[0] = v.X
	a[1] = v.Y
	a[2] = v.Z
}
