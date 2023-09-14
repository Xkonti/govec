package govec

// V2F

func (v V2F[T]) ClampComp(min V2F[T], max V2F[T]) V2F[T] {
	return V2F[T]{X: clamp(v.X, min.X, max.X), Y: clamp(v.Y, min.Y, max.Y)}
}

func (v *V2F[T]) ClampCompInPlace(min V2F[T], max V2F[T]) {
	v.X = clamp(v.X, min.X, max.X)
	v.Y = clamp(v.Y, min.Y, max.Y)
}

// V3F

func (v V3F[T]) ClampComp(min V3F[T], max V3F[T]) V3F[T] {
	return V3F[T]{X: clamp(v.X, min.X, max.X), Y: clamp(v.Y, min.Y, max.Y), Z: clamp(v.Z, min.Z, max.Z)}
}

func (v *V3F[T]) ClampCompInPlace(min V3F[T], max V3F[T]) {
	v.X = clamp(v.X, min.X, max.X)
	v.Y = clamp(v.Y, min.Y, max.Y)
	v.Z = clamp(v.Z, min.Z, max.Z)
}

// V2I

func (v V2I[T]) ClampComp(min V2I[T], max V2I[T]) V2I[T] {
	return V2I[T]{X: clamp(v.X, min.X, max.X), Y: clamp(v.Y, min.Y, max.Y)}
}

func (v *V2I[T]) ClampCompInPlace(min V2I[T], max V2I[T]) {
	v.X = clamp(v.X, min.X, max.X)
	v.Y = clamp(v.Y, min.Y, max.Y)
}

// V3I

func (v V3I[T]) ClampComp(min V3I[T], max V3I[T]) V3I[T] {
	return V3I[T]{X: clamp(v.X, min.X, max.X), Y: clamp(v.Y, min.Y, max.Y), Z: clamp(v.Z, min.Z, max.Z)}
}

func (v *V3I[T]) ClampCompInPlace(min V3I[T], max V3I[T]) {
	v.X = clamp(v.X, min.X, max.X)
	v.Y = clamp(v.Y, min.Y, max.Y)
	v.Z = clamp(v.Z, min.Z, max.Z)
}
