package govec

// V2F

// ClampComp returns a new V2F[T] with each component clamped between min and max.
func (v V2F[T]) ClampComp(min V2F[T], max V2F[T]) V2F[T] {
	return V2F[T]{X: clamp(v.X, min.X, max.X), Y: clamp(v.Y, min.Y, max.Y)}
}

// ClampCompComp returns a new V2F[T] with each component clamped between the respective min and max.
func (v V2F[T]) ClampCompComp(minX T, maxX T, minY T, maxY T) V2F[T] {
	return V2F[T]{X: clamp(v.X, minX, maxX), Y: clamp(v.Y, minY, maxY)}
}

// ClampCompInPlace modifies vector by clamping each component between min and max.
func (v *V2F[T]) ClampCompInPlace(min V2F[T], max V2F[T]) {
	v.X = clamp(v.X, min.X, max.X)
	v.Y = clamp(v.Y, min.Y, max.Y)
}

// ClampCompCompInPlace modifies vector by clamping each component between the respective min and max.
func (v *V2F[T]) ClampCompCompInPlace(minX T, maxX T, minY T, maxY T) {
	v.X = clamp(v.X, minX, maxX)
	v.Y = clamp(v.Y, minY, maxY)
}

// V3F

// ClampComp returns a new V3F[T] with each component clamped between min and max.
func (v V3F[T]) ClampComp(min V3F[T], max V3F[T]) V3F[T] {
	return V3F[T]{X: clamp(v.X, min.X, max.X), Y: clamp(v.Y, min.Y, max.Y), Z: clamp(v.Z, min.Z, max.Z)}
}

// ClampCompComp returns a new V3F[T] with each component clamped between the respective min and max.
func (v V3F[T]) ClampCompComp(minX T, maxX T, minY T, maxY T, minZ T, maxZ T) V3F[T] {
	return V3F[T]{X: clamp(v.X, minX, maxX), Y: clamp(v.Y, minY, maxY), Z: clamp(v.Z, minZ, maxZ)}
}

// ClampCompInPlace modifies vector by clamping each component between min and max.
func (v *V3F[T]) ClampCompInPlace(min V3F[T], max V3F[T]) {
	v.X = clamp(v.X, min.X, max.X)
	v.Y = clamp(v.Y, min.Y, max.Y)
	v.Z = clamp(v.Z, min.Z, max.Z)
}

// ClampCompCompInPlace modifies vector by clamping each component between the respective min and max.
func (v *V3F[T]) ClampCompCompInPlace(minX T, maxX T, minY T, maxY T, minZ T, maxZ T) {
	v.X = clamp(v.X, minX, maxX)
	v.Y = clamp(v.Y, minY, maxY)
	v.Z = clamp(v.Z, minZ, maxZ)
}

// V2I

// ClampComp returns a new V2I[T] with each component clamped between min and max.
func (v V2I[T]) ClampComp(min V2I[T], max V2I[T]) V2I[T] {
	return V2I[T]{X: clamp(v.X, min.X, max.X), Y: clamp(v.Y, min.Y, max.Y)}
}

// ClampCompComp returns a new V2I[T] with each component clamped between the respective min and max.
func (v V2I[T]) ClampCompComp(minX T, maxX T, minY T, maxY T) V2I[T] {
	return V2I[T]{X: clamp(v.X, minX, maxX), Y: clamp(v.Y, minY, maxY)}
}

// ClampCompInPlace modifies vector by clamping each component between min and max.
func (v *V2I[T]) ClampCompInPlace(min V2I[T], max V2I[T]) {
	v.X = clamp(v.X, min.X, max.X)
	v.Y = clamp(v.Y, min.Y, max.Y)
}

// ClampCompCompInPlace modifies vector by clamping each component between the respective min and max.
func (v *V2I[T]) ClampCompCompInPlace(minX T, maxX T, minY T, maxY T) {
	v.X = clamp(v.X, minX, maxX)
	v.Y = clamp(v.Y, minY, maxY)
}

// V3I

// ClampComp returns a new V3I[T] with each component clamped between min and max.
func (v V3I[T]) ClampComp(min V3I[T], max V3I[T]) V3I[T] {
	return V3I[T]{X: clamp(v.X, min.X, max.X), Y: clamp(v.Y, min.Y, max.Y), Z: clamp(v.Z, min.Z, max.Z)}
}

// ClampCompComp returns a new V3I[T] with each component clamped between the respective min and max.
func (v V3I[T]) ClampCompComp(minX T, maxX T, minY T, maxY T, minZ T, maxZ T) V3I[T] {
	return V3I[T]{X: clamp(v.X, minX, maxX), Y: clamp(v.Y, minY, maxY), Z: clamp(v.Z, minZ, maxZ)}
}

// ClampCompInPlace modifies vector by clamping each component between min and max.
func (v *V3I[T]) ClampCompInPlace(min V3I[T], max V3I[T]) {
	v.X = clamp(v.X, min.X, max.X)
	v.Y = clamp(v.Y, min.Y, max.Y)
	v.Z = clamp(v.Z, min.Z, max.Z)
}

// ClampCompCompInPlace modifies vector by clamping each component between the respective min and max.
func (v *V3I[T]) ClampCompCompInPlace(minX T, maxX T, minY T, maxY T, minZ T, maxZ T) {
	v.X = clamp(v.X, minX, maxX)
	v.Y = clamp(v.Y, minY, maxY)
	v.Z = clamp(v.Z, minZ, maxZ)
}
