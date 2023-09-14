package govec

import "golang.org/x/exp/constraints"

func V2FFromArray[T constraints.Float](a [2]T) V2F[T] {
	return V2F[T]{X: a[0], Y: a[1]}
}

func V3FFromArray[T constraints.Float](a [3]T) V3F[T] {
	return V3F[T]{X: a[0], Y: a[1], Z: a[2]}
}

func V2IFromArray[T constraints.Integer](a [2]T) V2I[T] {
	return V2I[T]{X: a[0], Y: a[1]}
}

func V3IFromArray[T constraints.Integer](a [3]T) V3I[T] {
	return V3I[T]{X: a[0], Y: a[1], Z: a[2]}
}

func (v V2F[T]) ToArray() [2]T {
	return [2]T{v.X, v.Y}
}

func (v V3F[T]) ToArray() [3]T {
	return [3]T{v.X, v.Y, v.Z}
}

func (v V2I[T]) ToArray() [2]T {
	return [2]T{v.X, v.Y}
}

func (v V3I[T]) ToArray() [3]T {
	return [3]T{v.X, v.Y, v.Z}
}

func (v V2F[T]) ApplyToArray(a *[2]T) {
	a[0] = v.X
	a[1] = v.Y
}

func (v V3F[T]) ApplyToArray(a *[3]T) {
	a[0] = v.X
	a[1] = v.Y
	a[2] = v.Z
}

func (v V2I[T]) ApplyToArray(a *[2]T) {
	a[0] = v.X
	a[1] = v.Y
}

func (v V3I[T]) ApplyToArray(a *[3]T) {
	a[0] = v.X
	a[1] = v.Y
	a[2] = v.Z
}
