package govec

// V2F

func (v V2F[T]) PowN(n int16) V2F[T] {
	return V2F[T]{X: pow(v.X, n), Y: pow(v.Y, n)}
}

func (v *V2F[T]) PowNInPlace(n int16) {
	v.X = pow(v.X, n)
	v.Y = pow(v.Y, n)
}

// V3F

func (v V3F[T]) PowN(n int16) V3F[T] {
	return V3F[T]{X: pow(v.X, n), Y: pow(v.Y, n), Z: pow(v.Z, n)}
}

func (v *V3F[T]) PowNInPlace(n int16) {
	v.X = pow(v.X, n)
	v.Y = pow(v.Y, n)
	v.Z = pow(v.Z, n)
}

// V2I

func (v V2I[T]) PowN(n int16) V2I[T] {
	return V2I[T]{X: pow(v.X, n), Y: pow(v.Y, n)}
}

func (v *V2I[T]) PowNInPlace(n int16) {
	v.X = pow(v.X, n)
	v.Y = pow(v.Y, n)
}

// V3I

func (v V3I[T]) PowN(n int16) V3I[T] {
	return V3I[T]{X: pow(v.X, n), Y: pow(v.Y, n), Z: pow(v.Z, n)}
}

func (v *V3I[T]) PowNInPlace(n int16) {
	v.X = pow(v.X, n)
	v.Y = pow(v.Y, n)
	v.Z = pow(v.Z, n)
}
