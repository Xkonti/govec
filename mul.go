package govec

// V2F

// Mul returns the component-wise product of two vectors.
func (v V2F[T]) Mul(v2 V2F[T]) V2F[T] {
	return V2F[T]{X: v.X * v2.X, Y: v.Y * v2.Y}
}

// MulInPlace modifies the vector by multiplying it with another vector.
func (v *V2F[T]) MulInPlace(v2 V2F[T]) {
	v.X *= v2.X
	v.Y *= v2.Y
}

// MulComp returns the component-wise product of a vector and three scalars.
func (v V2F[T]) MulComp(x T, y T) V2F[T] {
	return V2F[T]{X: v.X * x, Y: v.Y * y}
}

// MulCompInPlace modifies the vector by multiplying it with three scalars.
func (v *V2F[T]) MulCompInPlace(x T, y T) {
	v.X *= x
	v.Y *= y
}

// V3F

// Mul returns the component-wise product of two vectors.
func (v V3F[T]) Mul(v2 V3F[T]) V3F[T] {
	return V3F[T]{X: v.X * v2.X, Y: v.Y * v2.Y, Z: v.Z * v2.Z}
}

// MulInPlace modifies the vector by multiplying it with another vector.
func (v *V3F[T]) MulInPlace(v2 V3F[T]) {
	v.X *= v2.X
	v.Y *= v2.Y
	v.Z *= v2.Z
}

// MulComp returns the component-wise product of a vector and three scalars.
func (v V3F[T]) MulComp(x T, y T, z T) V3F[T] {
	return V3F[T]{X: v.X * x, Y: v.Y * y, Z: v.Z * z}
}

// MulCompInPlace modifies the vector by multiplying it with three scalars.
func (v *V3F[T]) MulCompInPlace(x T, y T, z T) {
	v.X *= x
	v.Y *= y
	v.Z *= z
}

// V2I

// Mul returns the component-wise product of two vectors.
func (v V2I[T]) Mul(v2 V2I[T]) V2I[T] {
	return V2I[T]{X: v.X * v2.X, Y: v.Y * v2.Y}
}

// MulInPlace modifies the vector by multiplying it with another vector.
func (v *V2I[T]) MulInPlace(v2 V2I[T]) {
	v.X *= v2.X
	v.Y *= v2.Y
}

// MulComp returns the component-wise product of a vector and three scalars.
func (v V2I[T]) MulComp(x T, y T) V2I[T] {
	return V2I[T]{X: v.X * x, Y: v.Y * y}
}

// MulCompInPlace modifies the vector by multiplying it with three scalars.
func (v *V2I[T]) MulCompInPlace(x T, y T) {
	v.X *= x
	v.Y *= y
}

// V3I

// Mul returns the component-wise product of two vectors.
func (v V3I[T]) Mul(v2 V3I[T]) V3I[T] {
	return V3I[T]{X: v.X * v2.X, Y: v.Y * v2.Y, Z: v.Z * v2.Z}
}

// MulInPlace modifies the vector by multiplying it with another vector.
func (v *V3I[T]) MulInPlace(v2 V3I[T]) {
	v.X *= v2.X
	v.Y *= v2.Y
	v.Z *= v2.Z
}

// MulComp returns the component-wise product of a vector and three scalars.
func (v V3I[T]) MulComp(x T, y T, z T) V3I[T] {
	return V3I[T]{X: v.X * x, Y: v.Y * y, Z: v.Z * z}
}

// MulCompInPlace modifies the vector by multiplying it with three scalars.
func (v *V3I[T]) MulCompInPlace(x T, y T, z T) {
	v.X *= x
	v.Y *= y
	v.Z *= z
}
