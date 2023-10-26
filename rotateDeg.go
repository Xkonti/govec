package govec

import "math"

// V2F

// RotateDeg rotates the vector counterclockwise by the specified number of degrees and returns a new vector.
func (v V2F[T]) RotateDeg(degrees float64) V2F[T] {
	d := degrees * degToRadFactor

	return V2F[T]{
		X: T(math.Cos(d)*float64(v.X) - math.Sin(d)*float64(v.Y)),
		Y: T(math.Sin(d)*float64(v.X) + math.Cos(d)*float64(v.Y)),
	}
}

// RotateDegInPlace modifies v by rotating the vector counterclockwise by the specified number of degrees.
func (v *V2F[T]) RotateDegInPlace(degrees float64) {
	d := degrees * degToRadFactor

	x1 := T(math.Cos(d)*float64(v.X) - math.Sin(d)*float64(v.Y))
	y1 := T(math.Sin(d)*float64(v.X) + math.Cos(d)*float64(v.Y))

	v.X = x1
	v.Y = y1
}

// V2I

// RotateDeg rotates the vector counterclockwise by the specified number of degrees and returns a new V2F vector.
func (v V2I[T]) RotateDeg(degrees float64) V2F[float64] {
	d := degrees * degToRadFactor

	t := V2F[float64]{
		X: math.Cos(d)*float64(v.X) - math.Sin(d)*float64(v.Y),
		Y: math.Sin(d)*float64(v.X) + math.Cos(d)*float64(v.Y),
	}

	return t
}

// V3F

// RotateDeg rotates the vector counterclockwise by the specified number of degrees, around axis and returns a new vector.
func (v V3F[T]) RotateDeg(degrees float64, axis axis) V3F[T] {
	d := degrees * degToRadFactor

	switch axis {
	case zAxis:
		return V3F[T]{
			X: T(math.Cos(d)*float64(v.X) - math.Sin(d)*float64(v.Y)),
			Y: T(math.Sin(d)*float64(v.X) + math.Cos(d)*float64(v.Y)),
			Z: v.Z,
		}
	case yAxis:
		return V3F[T]{
			X: T(math.Cos(d)*float64(v.X) - math.Sin(d)*float64(v.Y)),
			Y: v.Y,
			Z: T(-math.Sin(d)*float64(v.X) + math.Cos(d)*float64(v.Z)),
		}
	default:
		return V3F[T]{
			X: v.X,
			Y: T(math.Cos(d)*float64(v.Y) - math.Sin(d)*float64(v.Z)),
			Z: T(-math.Sin(d)*float64(v.Z) + math.Cos(d)*float64(v.Z)),
		}
	}
}

// RotateDegInPlace modifies v by rotating the vector counterclockwise by the specified number of degrees, around axis.
func (v *V3F[T]) RotateDegInPlace(degrees float64, axis axis) {
	d := degrees * degToRadFactor

	switch axis {
	case zAxis:
		x := T(math.Cos(d)*float64(v.X) - math.Sin(d)*float64(v.Y))
		y := T(math.Sin(d)*float64(v.X) + math.Cos(d)*float64(v.Y))
		z := T(float64(v.Z))

		v.X = x
		v.Y = y
		v.Z = z
	case yAxis:
		x := T(math.Cos(d)*float64(v.X) - math.Sin(d)*float64(v.Y))
		y := T(float64(v.Y))
		z := T(-math.Sin(d)*float64(v.X) + math.Cos(d)*float64(v.Z))

		v.X = x
		v.Y = y
		v.Z = z
	default:
		x := T(float64(v.X))
		y := T(math.Cos(d)*float64(v.Y) - math.Sin(d)*float64(v.Z))
		z := T(-math.Sin(d)*float64(v.Z) + math.Cos(d)*float64(v.Z))

		v.X = x
		v.Y = y
		v.Z = z
	}
}

// V3I

// RotateDeg rotates the vector counterclockwise by the specified number of degrees, around axis and returns a new V3F vector.
func (v V3I[T]) RotateDeg(degrees float64, axis axis) V3F[float64] {
	d := degrees * degToRadFactor

	switch axis {
	case zAxis:
		return V3F[float64]{
			X: math.Cos(d)*float64(v.X) - math.Sin(d)*float64(v.Y),
			Y: math.Sin(d)*float64(v.X) + math.Cos(d)*float64(v.Y),
			Z: float64(v.Z),
		}
	case yAxis:
		return V3F[float64]{
			X: math.Cos(d)*float64(v.X) - math.Sin(d)*float64(v.Y),
			Y: float64(v.Y),
			Z: -math.Sin(d)*float64(v.X) + math.Cos(d)*float64(v.Z),
		}
	default:
		return V3F[float64]{
			X: float64(v.X),
			Y: math.Cos(d)*float64(v.Y) - math.Sin(d)*float64(v.Z),
			Z: -math.Sin(d)*float64(v.Z) + math.Cos(d)*float64(v.Z),
		}
	}
}
