package govec

import (
	"math"
	"testing"
)

func TestV2F_Average(t *testing.T) {
	v1 := V2F[float64]{X: 0, Y: math.Pi / 2}
	avg := v1.Average()

	if !almostEqual(avg, math.Pi/4, 1e-10) {
		t.Errorf("V2F Average failed")
	}
}

func TestV3F_Average(t *testing.T) {
	v1 := V3F[float64]{X: 0, Y: math.Pi / 2, Z: 3 * math.Pi / 2}
	avg := v1.Average()

	if !almostEqual(avg, 2.0/3.0*math.Pi, 1e-10) {
		t.Errorf("V3F Average failed")
	}
}

func TestV2I_Average(t *testing.T) {
	v1 := V2I[int16]{X: 0, Y: 1}
	avg := v1.Average()

	if !almostEqual(avg, 0.5, 1e-10) {
		t.Errorf("V2I Average failed")
	}
}

func TestV3I_Average(t *testing.T) {
	v1 := V3I[int16]{X: 0, Y: 1, Z: 3}
	avg := v1.Average()

	if !almostEqual(avg, 4.0/3.0, 1e-10) {
		t.Errorf("V3I Average failed")
	}
}
