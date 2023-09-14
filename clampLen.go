package govec

// V2F

func (v V2F[T]) ClampLen(min float64, max float64) V2F[T] {
	length := v.Len()
	if length == 0 {
		return V2F[T]{X: 0, Y: 0}
	}
	if length < min {
		scale := min / length
		return V2F[T]{X: v.X * T(scale), Y: v.Y * T(scale)}
	} else if length > max {
		scale := max / length
		return V2F[T]{X: v.X * T(scale), Y: v.Y * T(scale)}
	}
	return v
}

func (v *V2F[T]) ClampLenInPlace(min float64, max float64) {
	length := v.Len()
	if length == 0 {
		v.X = 0
		v.Y = 0
	}
	if length < min {
		scale := min / length
		v.X *= T(scale)
		v.Y *= T(scale)
	} else if length > max {
		scale := max / length
		v.X *= T(scale)
		v.Y *= T(scale)
	}
}

// V3F

func (v V3F[T]) ClampLen(min float64, max float64) V3F[T] {
	length := v.Len()
	if length == 0 {
		return V3F[T]{X: 0, Y: 0, Z: 0}
	}
	if length < min {
		scale := min / length
		return V3F[T]{X: v.X * T(scale), Y: v.Y * T(scale), Z: v.Z * T(scale)}
	} else if length > max {
		scale := max / length
		return V3F[T]{X: v.X * T(scale), Y: v.Y * T(scale), Z: v.Z * T(scale)}
	}
	return v
}

func (v *V3F[T]) ClampLenInPlace(min float64, max float64) {
	length := v.Len()
	if length == 0 {
		v.X = 0
		v.Y = 0
		v.Z = 0
	}
	if length < min {
		scale := min / length
		v.X *= T(scale)
		v.Y *= T(scale)
		v.Z *= T(scale)
	} else if length > max {
		scale := max / length
		v.X *= T(scale)
		v.Y *= T(scale)
		v.Z *= T(scale)
	}
}

// V2I

func (v V2I[T]) ClampLen(min float64, max float64) V2F[float64] {
	length := v.Len()
	if length == 0 {
		return V2F[float64]{X: 0, Y: 0}
	}
	if length < min {
		scale := min / length
		return V2F[float64]{X: float64(v.X) * scale, Y: float64(v.Y) * scale}
	} else if length > max {
		scale := max / length
		return V2F[float64]{X: float64(v.X) * scale, Y: float64(v.Y) * scale}
	}
	return V2F[float64]{X: float64(v.X), Y: float64(v.Y)}
}

// V3I

func (v V3I[T]) ClampLen(min float64, max float64) V3F[float64] {
	length := v.Len()
	if length == 0 {
		return V3F[float64]{X: 0, Y: 0, Z: 0}
	}
	if length < min {
		scale := min / length
		return V3F[float64]{X: float64(v.X) * scale, Y: float64(v.Y) * scale, Z: float64(v.Z) * scale}
	} else if length > max {
		scale := max / length
		return V3F[float64]{X: float64(v.X) * scale, Y: float64(v.Y) * scale, Z: float64(v.Z) * scale}
	}
	return V3F[float64]{X: float64(v.X), Y: float64(v.Y), Z: float64(v.Z)}
}
