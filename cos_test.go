package govec

import (
	"math"
	"testing"
)

func TestV2F_Cos(t *testing.T) {
	v1 := V2F[float64]{X: 0, Y: math.Pi}
	v2 := v1.Cos()

	if !almostEqual(v2.X, 1, 1e-10) || !almostEqual(v2.Y, -1, 1e-10) {
		t.Errorf("V2F Cos failed")
	}
}

func TestV2F_CosInPlace(t *testing.T) {
	v1 := V2F[float64]{X: 0, Y: math.Pi}
	v1.CosInPlace()

	if !almostEqual(v1.X, 1, 1e-10) || !almostEqual(v1.Y, -1, 1e-10) {
		t.Errorf("V2F CosInPlace failed")
	}
}

func TestV3F_Cos(t *testing.T) {
	v1 := V3F[float32]{X: 0, Y: math.Pi, Z: 2 * math.Pi}
	v2 := v1.Cos()

	if !almostEqual(v2.X, 1, 1e-10) || !almostEqual(v2.Y, -1, 1e-10) || !almostEqual(v2.Z, 1, 1e-10) {
		t.Errorf("V3F Cos failed")
	}
}

func TestV3F_CosInPlace(t *testing.T) {
	v1 := V3F[float64]{X: 0, Y: math.Pi, Z: 2 * math.Pi}
	v1.CosInPlace()

	if !almostEqual(v1.X, 1, 1e-10) || !almostEqual(v1.Y, -1, 1e-10) || !almostEqual(v1.Z, 1, 1e-10) {
		t.Errorf("V3F CosInPlace failed")
	}
}

func TestV2I_Cos(t *testing.T) {
	v1 := V2I[int16]{X: 0, Y: 1}
	v2 := v1.Cos()

	if !almostEqual(v2.X, 1, 1e-10) || !almostEqual(v2.Y, math.Cos(1), 1e-10) {
		t.Errorf("V2I Cos failed")
	}
}

func TestV3I_Cos(t *testing.T) {
	v1 := V3I[int16]{X: 0, Y: 1, Z: 3}
	v2 := v1.Cos()

	if !almostEqual(v2.X, 1, 1e-10) || !almostEqual(v2.Y, math.Cos(1), 1e-10) || !almostEqual(v2.Z, math.Cos(3), 1e-10) {
		t.Errorf("V3I Cos failed")
	}
}
