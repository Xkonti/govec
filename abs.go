package govec

// V2F

func (v V2F[T]) Abs() V2F[T] {
	absX := v.X
	absY := v.Y
	if absX < 0 {
		absX = -absX
	}
	if absY < 0 {
		absY = -absY
	}
	return V2F[T]{X: absX, Y: absY}
}

func (v *V2F[T]) AbsInPlace() {
	if v.X < 0 {
		v.X = -v.X
	}
	if v.Y < 0 {
		v.Y = -v.Y
	}
}

// V3F

func (v V3F[T]) Abs() V3F[T] {
	absX := v.X
	absY := v.Y
	absZ := v.Z
	if absX < 0 {
		absX = -absX
	}
	if absY < 0 {
		absY = -absY
	}
	if absZ < 0 {
		absZ = -absZ
	}
	return V3F[T]{X: absX, Y: absY, Z: absZ}
}

func (v *V3F[T]) AbsInPlace() {
	if v.X < 0 {
		v.X = -v.X
	}
	if v.Y < 0 {
		v.Y = -v.Y
	}
	if v.Z < 0 {
		v.Z = -v.Z
	}
}

// V2I

func (v V2I[T]) Abs() V2I[T] {
	absX := v.X
	absY := v.Y
	if absX < 0 {
		absX = -absX
	}
	if absY < 0 {
		absY = -absY
	}
	return V2I[T]{X: absX, Y: absY}
}

func (v *V2I[T]) AbsInPlace() {
	if v.X < 0 {
		v.X = -v.X
	}
	if v.Y < 0 {
		v.Y = -v.Y
	}
}

// V3I

func (v V3I[T]) Abs() V3I[T] {
	absX := v.X
	absY := v.Y
	absZ := v.Z
	if absX < 0 {
		absX = -absX
	}
	if absY < 0 {
		absY = -absY
	}
	if absZ < 0 {
		absZ = -absZ
	}
	return V3I[T]{X: absX, Y: absY, Z: absZ}
}

func (v *V3I[T]) AbsInPlace() {
	if v.X < 0 {
		v.X = -v.X
	}
	if v.Y < 0 {
		v.Y = -v.Y
	}
	if v.Z < 0 {
		v.Z = -v.Z
	}
}
