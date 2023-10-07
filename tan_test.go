package govec

import (
	"math"
	"testing"
)

func TestV2F_Tan(t *testing.T) {
	v1 := V2F[float64]{X: 0, Y: math.Pi / 4}
	v2 := v1.Tan()

	if v2.X != 0 || !almostEqual(v2.Y, 1, 1e-10) {
		t.Errorf("V2F Tan failed")
	}
}

func TestV2F_TanInPlace(t *testing.T) {
	v1 := V2F[float64]{X: 0, Y: math.Pi / 4}
	v1.TanInPlace()

	if v1.X != 0 || !almostEqual(v1.Y, 1, 1e-10) {
		t.Errorf("V2F TanInPlace failed")
	}
}

func TestV3F_Tan(t *testing.T) {
	v1 := V3F[float64]{X: 0, Y: math.Pi / 4, Z: math.Pi}
	v2 := v1.Tan()

	if v2.X != 0 || !almostEqual(v2.Y, 1, 1e-10) || !almostEqual(v2.Z, 0, 1e-10) {
		t.Errorf("V3F Tan failed")
	}
}

func TestV3F_TanInPlace(t *testing.T) {
	v1 := V3F[float64]{X: 0, Y: math.Pi / 4, Z: math.Pi}
	v1.TanInPlace()

	if v1.X != 0 || !almostEqual(v1.Y, 1, 1e-10) || !almostEqual(v1.Z, 0, 1e-10) {
		t.Errorf("V3F TanInPlace failed")
	}
}

func TestV2I_Tan(t *testing.T) {
	v1 := V2I[int16]{X: 0, Y: 1}
	v2 := v1.Tan()

	if v2.X != 0 || v2.Y != 1.557407724654902 {
		t.Errorf("V2I Tan failed")
	}
}

func TestV3I_Tan(t *testing.T) {
	v1 := V3I[int16]{X: 0, Y: 1, Z: 3}
	v2 := v1.Tan()

	if v2.X != 0 || v2.Y != 1.557407724654902 || v2.Z != -0.1425465430742778 {
		t.Errorf("V3I Tan failed")
	}
}
