package govec

import "math"

func (v V2F[T]) RotateDeg(degrees float64) V2F[T] {
	d := degrees * (math.Pi / 180)

	return V2F[T]{
		X: T(math.Cos(d)*float64(v.X) - math.Sin(d)*float64(v.Y)),
		Y: T(math.Sin(d)*float64(v.X) + math.Cos(d)*float64(v.Y)),
	}
}

func (v V2I[T]) RotateDeg(degrees float64) V2F[float64] {
	d := degrees * (math.Pi / 180)

	t := V2F[float64]{
		X: math.Cos(d)*float64(v.X) - math.Sin(d)*float64(v.Y),
		Y: math.Sin(d)*float64(v.X) + math.Cos(d)*float64(v.Y),
	}

	return t
}

func (v *V2F[T]) RotateDegInPlace(degrees float64) {
	d := degrees * (math.Pi / 180)

	x1 := T(math.Cos(d)*float64(v.X) - math.Sin(d)*float64(v.Y))
	y1 := T(math.Sin(d)*float64(v.X) + math.Cos(d)*float64(v.Y))

	v.X = x1
	v.Y = y1
}

func (v V3F[T]) RotateDeg(degrees float64, axis string) V3F[T] {
	d := degrees * (math.Pi / 180)

	switch axis {
	case "z":
		return V3F[T]{
			X: T(math.Cos(d)*float64(v.X) - math.Sin(d)*float64(v.Y)),
			Y: T(math.Sin(d)*float64(v.X) + math.Cos(d)*float64(v.Y)),
			Z: v.Z,
		}
	case "y":
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

func (v V3I[T]) RotateDeg(degrees float64, axis string) V3F[float64] {
	d := degrees * (math.Pi / 180)

	switch axis {
	case "z":
		return V3F[float64]{
			X: math.Cos(d)*float64(v.X) - math.Sin(d)*float64(v.Y),
			Y: math.Sin(d)*float64(v.X) + math.Cos(d)*float64(v.Y),
			Z: float64(v.Z),
		}
	case "y":
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

func (v *V3F[T]) RotateDegInPlace(degrees float64, axis string) {
	d := degrees * (math.Pi / 180)

	switch axis {
	case "z":
		x := T(math.Cos(d)*float64(v.X) - math.Sin(d)*float64(v.Y))
		y := T(math.Sin(d)*float64(v.X) + math.Cos(d)*float64(v.Y))
		z := T(float64(v.Z))

		v.X = x
		v.Y = y
		v.Z = z
	case "y":
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
