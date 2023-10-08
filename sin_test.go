package govec

import (
	"math"
	"testing"
)

func TestV2F_Sin(t *testing.T) {
	v1 := V2F[float64]{X: 0, Y: math.Pi / 2}
	v2 := v1.Sin()

	if !almostEqual(v2.X, 0, 1e-10) || !almostEqual(v2.Y, 1, 1e-10) {
		t.Errorf("V2F Sin failed")
	}
}

func TestV2F_SinInPlace(t *testing.T) {
	v1 := V2F[float64]{X: 0, Y: math.Pi / 2}
	v1.SinInPlace()

	if !almostEqual(v1.X, 0, 1e-10) || !almostEqual(v1.Y, 1, 1e-10) {
		t.Errorf("V2F SinInPlace failed")
	}
}

func TestV3F_Sin(t *testing.T) {
	v1 := V3F[float64]{X: 0, Y: math.Pi / 2, Z: 3 * math.Pi / 2}
	v2 := v1.Sin()

	if !almostEqual(v2.X, 0, 1e-10) || !almostEqual(v2.Y, 1, 1e-10) || !almostEqual(v2.Z, -1, 1e-10) {
		t.Errorf("V3F Sin failed")
	}
}

func TestV3F_SinInPlace(t *testing.T) {
	v1 := V3F[float64]{X: 0, Y: math.Pi / 2, Z: 3 * math.Pi / 2}
	v1.SinInPlace()

	if !almostEqual(v1.X, 0, 1e-10) || !almostEqual(v1.Y, 1, 1e-10) || !almostEqual(v1.Z, -1, 1e-10) {
		t.Errorf("V3F SinInPlace failed")
	}
}

func TestV2I_Sin(t *testing.T) {
	v1 := V2I[int16]{X: 0, Y: 1}
	v2 := v1.Sin()

	if !almostEqual(v2.X, 0, 1e-10) || !almostEqual(v2.Y, math.Sin(1), 1e-10) {
		t.Errorf("V2I Sin failed")
	}
}

func TestV3I_Sin(t *testing.T) {
	v1 := V3I[int16]{X: 0, Y: 1, Z: 3}
	v2 := v1.Sin()

	if !almostEqual(v2.X, 0, 1e-10) || !almostEqual(v2.Y, math.Sin(1), 1e-10) || !almostEqual(v2.Z, math.Sin(3), 1e-10) {
		t.Errorf("V3I Sin failed")
	}
}
