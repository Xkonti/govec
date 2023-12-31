package govec

// V2F

// Mod returns a new V2F[T] with modulus applied to each component.
func (v V2F[T]) Mod(v2 V2F[T]) V2F[T] {
	return V2F[T]{X: v.X - v2.X*T(int64(v.X/v2.X)), Y: v.Y - v2.Y*T(int64(v.Y/v2.Y))}
}

// ModInPlace updates the vector in-place with modulus applied to each component.
func (v *V2F[T]) ModInPlace(v2 V2F[T]) {
	v.X = v.X - v2.X*T(int64(v.X/v2.X))
	v.Y = v.Y - v2.Y*T(int64(v.Y/v2.Y))
}

// ModComp returns a new V2F[T] with modulus applied to each component.
func (v V2F[T]) ModComp(x T, y T) V2F[T] {
	return V2F[T]{X: v.X - x*T(int64(v.X/x)), Y: v.Y - y*T(int64(v.Y/y))}
}

// ModCompInPlace updates the vector in-place with modulus applied to each component.
func (v *V2F[T]) ModCompInPlace(x T, y T) {
	v.X = v.X - x*T(int64(v.X/x))
	v.Y = v.Y - y*T(int64(v.Y/y))
}

// V3F

// Mod returns a new V3F[T] with modulus applied to each component.
func (v V3F[T]) Mod(v2 V3F[T]) V3F[T] {
	return V3F[T]{X: v.X - v2.X*T(int64(v.X/v2.X)), Y: v.Y - v2.Y*T(int64(v.Y/v2.Y)), Z: v.Z - v2.Z*T(int64(v.Z/v2.Z))}
}

// ModInPlace updates the vector in-place with modulus applied to each component.
func (v *V3F[T]) ModInPlace(v2 V3F[T]) {
	v.X = v.X - v2.X*T(int64(v.X/v2.X))
	v.Y = v.Y - v2.Y*T(int64(v.Y/v2.Y))
	v.Z = v.Z - v2.Z*T(int64(v.Z/v2.Z))
}

// ModComp returns a new V3F[T] with modulus applied to each component.
func (v V3F[T]) ModComp(x T, y T, z T) V3F[T] {
	return V3F[T]{X: v.X - x*T(int64(v.X/x)), Y: v.Y - y*T(int64(v.Y/y)), Z: v.Z - z*T(int64(v.Z/z))}
}

// ModCompInPlace updates the vector in-place with modulus applied to each component.
func (v *V3F[T]) ModCompInPlace(x T, y T, z T) {
	v.X = v.X - x*T(int64(v.X/x))
	v.Y = v.Y - y*T(int64(v.Y/y))
	v.Z = v.Z - z*T(int64(v.Z/z))
}

// V2I

// Mod returns a new V2I[T] with modulus applied to each component.
func (v V2I[T]) Mod(v2 V2I[T]) V2I[T] {
	return V2I[T]{X: v.X % v2.X, Y: v.Y % v2.Y}
}

// ModInPlace updates the vector in-place with modulus applied to each component.
func (v *V2I[T]) ModInPlace(v2 V2I[T]) {
	v.X %= v2.X
	v.Y %= v2.Y
}

// ModComp returns a new V2I[T] with modulus applied to each component.
func (v V2I[T]) ModComp(x T, y T) V2I[T] {
	return V2I[T]{X: v.X % x, Y: v.Y % y}
}

// ModCompInPlace updates the vector in-place with modulus applied to each component.
func (v *V2I[T]) ModCompInPlace(x T, y T) {
	v.X %= x
	v.Y %= y
}

// V3I

// Mod returns a new V3I[T] with modulus applied to each component.
func (v V3I[T]) Mod(v2 V3I[T]) V3I[T] {
	return V3I[T]{X: v.X % v2.X, Y: v.Y % v2.Y, Z: v.Z % v2.Z}
}

// ModInPlace updates the vector in-place with modulus applied to each component.
func (v *V3I[T]) ModInPlace(v2 V3I[T]) {
	v.X %= v2.X
	v.Y %= v2.Y
	v.Z %= v2.Z
}

// ModComp returns a new V3I[T] with modulus applied to each component.
func (v V3I[T]) ModComp(x T, y T, z T) V3I[T] {
	return V3I[T]{X: v.X % x, Y: v.Y % y, Z: v.Z % z}
}

// ModCompInPlace updates the vector in-place with modulus applied to each component.
func (v *V3I[T]) ModCompInPlace(x T, y T, z T) {
	v.X %= x
	v.Y %= y
	v.Z %= z
}
