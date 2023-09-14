package govec

// V3F

func (v V3F[T]) Cross(v2 V3F[T]) V3F[T] {
	return V3F[T]{
		X: v.Y*v2.Z - v.Z*v2.Y,
		Y: v.Z*v2.X - v.X*v2.Z,
		Z: v.X*v2.Y - v.Y*v2.X,
	}
}

func (v *V3F[T]) CrossInPlace(v2 V3F[T]) {
	x := v.Y*v2.Z - v.Z*v2.Y
	y := v.Z*v2.X - v.X*v2.Z
	z := v.X*v2.Y - v.Y*v2.X
	v.X, v.Y, v.Z = x, y, z
}

func (v V3F[T]) CrossComp(x T, y T, z T) V3F[T] {
	return V3F[T]{
		X: v.Y*z - v.Z*y,
		Y: v.Z*x - v.X*z,
		Z: v.X*y - v.Y*x,
	}
}

func (v *V3F[T]) CrossCompInPlace(x T, y T, z T) {
	nx := v.Y*z - v.Z*y
	ny := v.Z*x - v.X*z
	nz := v.X*y - v.Y*x
	v.X, v.Y, v.Z = nx, ny, nz
}

// V3I

func (v V3I[T]) Cross(v2 V3I[T]) V3I[T] {
	return V3I[T]{
		X: v.Y*v2.Z - v.Z*v2.Y,
		Y: v.Z*v2.X - v.X*v2.Z,
		Z: v.X*v2.Y - v.Y*v2.X,
	}
}

func (v *V3I[T]) CrossInPlace(v2 V3I[T]) {
	x := v.Y*v2.Z - v.Z*v2.Y
	y := v.Z*v2.X - v.X*v2.Z
	z := v.X*v2.Y - v.Y*v2.X
	v.X, v.Y, v.Z = x, y, z
}

func (v V3I[T]) CrossComp(x T, y T, z T) V3I[T] {
	return V3I[T]{
		X: v.Y*z - v.Z*y,
		Y: v.Z*x - v.X*z,
		Z: v.X*y - v.Y*x,
	}
}

func (v *V3I[T]) CrossCompInPlace(x T, y T, z T) {
	nx := v.Y*z - v.Z*y
	ny := v.Z*x - v.X*z
	nz := v.X*y - v.Y*x
	v.X, v.Y, v.Z = nx, ny, nz
}
