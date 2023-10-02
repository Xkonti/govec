package govec

// V2F

// SwizzleYZ returns a new V2F with the X and Y components swapped.
func (v V2F[T]) SwizzleYZ() V2F[T] {
	return V2F[T]{X: v.Y, Y: v.X}
}

// SwizzleInPlaceYZ swaps the X and Y components of the V2F in place.
func (v V2F[T]) SwizzleInPlaceYZ() {
	v.X, v.Y = v.Y, v.X
}

// V3F

// SwizzleXZY returns a new V3F with components swapped: X->X Z->Y Y->Z
func (v V3F[T]) SwizzleXZY() V3F[T] {
	return V3F[T]{X: v.X, Y: v.Z, Z: v.Y}
}

// SwizzleInPlaceXZY swaps components in place: X->X Z->Y Y->Z
func (v V3F[T]) SwizzleInPlaceXZY() {
	v.Y, v.Z = v.Z, v.Y
}

// SwizzleYXZ returns a new V3F with components swapped: Y->X X->Y Z->Z
func (v V3F[T]) SwizzleYXZ() V3F[T] {
	return V3F[T]{X: v.Y, Y: v.X, Z: v.Z}
}

// SwizzleInPlaceYXZ swaps components in place: Y->X X->Y Z->Z
func (v V3F[T]) SwizzleInPlaceYXZ() {
	v.X, v.Y = v.Y, v.X
}

// SwizzleYZX returns a new V3F with components swapped: Y->X Z->Y X->Z
func (v V3F[T]) SwizzleYZX() V3F[T] {
	return V3F[T]{X: v.Y, Y: v.Z, Z: v.X}
}

// SwizzleInPlaceYZX swaps components in place: Y->X Z->Y X->Z
func (v V3F[T]) SwizzleInPlaceYZX() {
	v.X, v.Y, v.Z = v.Y, v.Z, v.X
}

// SwizzleZXY returns a new V3F with components swapped: Z->X X->Y Y->Z
func (v V3F[T]) SwizzleZXY() V3F[T] {
	return V3F[T]{X: v.Z, Y: v.X, Z: v.Y}
}

// SwizzleInPlaceZXY swaps components in place: Z->X X->Y Y->Z
func (v V3F[T]) SwizzleInPlaceZXY() {
	v.X, v.Y, v.Z = v.Z, v.X, v.Y
}

// SwizzleZYX returns a new V3F with components swapped: Z->X Y->Y X->Z
func (v V3F[T]) SwizzleZYX() V3F[T] {
	return V3F[T]{X: v.Z, Y: v.Y, Z: v.X}
}

// SwizzleInPlaceZYX swaps components in place: Z->X Y->Y X->Z
func (v V3F[T]) SwizzleInPlaceZYX() {
	v.X, v.Z = v.Z, v.X
}

// V2I

// SwizzleYZ returns a new V2I with the X and Y components swapped.
func (v V2I[T]) SwizzleYZ() V2I[T] {
	return V2I[T]{X: v.Y, Y: v.X}
}

// SwizzleInPlaceYZ swaps the X and Y components of the V2I in place.
func (v V2I[T]) SwizzleInPlaceYZ() {
	v.X, v.Y = v.Y, v.X
}

// V3I

// SwizzleXZY returns a new V3I with components swapped: X->X Z->Y Y->Z
func (v V3I[T]) SwizzleXZY() V3I[T] {
	return V3I[T]{X: v.X, Y: v.Z, Z: v.Y}
}

// SwizzleInPlaceXZY swaps components in place: X->X Z->Y Y->Z
func (v V3I[T]) SwizzleInPlaceXZY() {
	v.Y, v.Z = v.Z, v.Y
}

// SwizzleYXZ returns a new V3I with components swapped: Y->X X->Y Z->Z
func (v V3I[T]) SwizzleYXZ() V3I[T] {
	return V3I[T]{X: v.Y, Y: v.X, Z: v.Z}
}

// SwizzleInPlaceYXZ swaps components in place: Y->X X->Y Z->Z
func (v V3I[T]) SwizzleInPlaceYXZ() {
	v.X, v.Y = v.Y, v.X
}

// SwizzleYZX returns a new V3I with components swapped: Y->X Z->Y X->Z
func (v V3I[T]) SwizzleYZX() V3I[T] {
	return V3I[T]{X: v.Y, Y: v.Z, Z: v.X}
}

// SwizzleInPlaceYZX swaps components in place: Y->X Z->Y X->Z
func (v V3I[T]) SwizzleInPlaceYZX() {
	v.X, v.Y, v.Z = v.Y, v.Z, v.X
}

// SwizzleZXY returns a new V3I with components swapped: Z->X X->Y Y->Z
func (v V3I[T]) SwizzleZXY() V3I[T] {
	return V3I[T]{X: v.Z, Y: v.X, Z: v.Y}
}

// SwizzleInPlaceZXY swaps components in place: Z->X X->Y Y->Z
func (v V3I[T]) SwizzleInPlaceZXY() {
	v.X, v.Y, v.Z = v.Z, v.X, v.Y
}

// SwizzleZYX returns a new V3I with components swapped: Z->X Y->Y X->Z
func (v V3I[T]) SwizzleZYX() V3I[T] {
	return V3I[T]{X: v.Z, Y: v.Y, Z: v.X}
}

// SwizzleInPlaceZYX swaps components in place: Z->X Y->Y X->Z
func (v V3I[T]) SwizzleInPlaceZYX() {
	v.X, v.Z = v.Z, v.X
}
