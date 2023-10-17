package govec

import (
	"testing"
)

func TestV2F_RotateDeg(t *testing.T) {
	type expected struct {
		X float64
		Y float64
	}

	tt := []struct {
		name     string
		degrees  float64
		expected expected
	}{
		{name: "0", expected: expected{X: 0, Y: 1}},
		{name: "90", degrees: 90, expected: expected{X: -1, Y: 0}},
		{name: "180", degrees: 180, expected: expected{X: 0, Y: -1}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v1 := V2F[float64]{X: 0, Y: 1}
			v2 := v1.RotateDeg(tc.degrees)

			if !almostEqual(v2.X, tc.expected.X, 1e-10) {
				t.Errorf("V2F RotateDeg failed expected %v: got %v", tc.expected.X, v2.X)
			}

			if !almostEqual(v2.Y, tc.expected.Y, 1e-10) {
				t.Errorf("V2F RotateDeg failed expected %v: got %v", tc.expected.Y, v2.Y)
			}
		})
	}
}

func TestV2I_RotateDeg(t *testing.T) {
	type expected struct {
		X float64
		Y float64
	}

	tt := []struct {
		name     string
		degrees  float64
		expected expected
	}{
		{name: "0", expected: expected{X: 0, Y: 1}},
		{name: "90", degrees: 90, expected: expected{X: -1, Y: 0}},
		{name: "180", degrees: 180, expected: expected{X: 0, Y: -1}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v1 := V2I[int]{X: 0, Y: 1}
			v2 := v1.RotateDeg(tc.degrees)

			if !almostEqual(v2.X, tc.expected.X, 1e-10) {
				t.Errorf("V2I RotateDeg failed expected %v: got %v", tc.expected.X, v2.X)
			}

			if !almostEqual(v2.Y, tc.expected.Y, 1e-10) {
				t.Errorf("V2I RotateDeg failed expected %v: got %v", tc.expected.Y, v2.Y)
			}
		})
	}
}

func TestV2F_RotateDegInPlace(t *testing.T) {
	type expected struct {
		X float64
		Y float64
	}

	tt := []struct {
		name     string
		degrees  float64
		expected expected
	}{
		{name: "0", expected: expected{X: 0, Y: 1}},
		{name: "90", degrees: 90, expected: expected{X: -1, Y: 0}},
		{name: "180", degrees: 180, expected: expected{X: 0, Y: -1}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v1 := V2F[float64]{X: 0, Y: 1}
			v1.RotateDegInPlace(tc.degrees)

			if !almostEqual(v1.X, tc.expected.X, 1e-10) {
				t.Errorf("V2F RotateDegInPlace failed expected %v: got %v", tc.expected.X, v1.X)
			}

			if !almostEqual(v1.Y, tc.expected.Y, 1e-10) {
				t.Errorf("V2F RotateDegInPlace failed expected %v: got %v", tc.expected.Y, v1.Y)
			}
		})
	}
}

func TestV3F_RotateDeg(t *testing.T) {
	type expected struct {
		X float64
		Y float64
		Z float64
	}

	tt := []struct {
		name     string
		degrees  float64
		axis     string
		expected expected
	}{
		{name: "z0", axis: "z", expected: expected{X: 0, Y: 1, Z: 0}},
		{name: "z90", degrees: 90, axis: "z", expected: expected{X: -1, Y: 0, Z: 0}},
		{name: "z180", degrees: 180, axis: "z", expected: expected{X: 0, Y: -1, Z: 0}},

		{name: "y0", axis: "y", expected: expected{X: 0, Y: 1, Z: 0}},
		{name: "y90", degrees: 90, axis: "y", expected: expected{X: -1, Y: 1, Z: 0}},
		{name: "y180", degrees: 180, axis: "y", expected: expected{X: 0, Y: 1, Z: 0}},

		{name: "x0", axis: "x", expected: expected{X: 0, Y: 1, Z: 0}},
		{name: "x90", degrees: 90, axis: "x", expected: expected{X: 0, Y: 0, Z: 0}},
		{name: "x180", degrees: 180, axis: "x", expected: expected{X: 0, Y: -1, Z: 0}},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v1 := V3F[float64]{X: 0, Y: 1, Z: 0}
			v2 := v1.RotateDeg(tc.degrees, tc.axis)

			if !almostEqual(v2.X, tc.expected.X, 1e-10) {
				t.Errorf("V3F RotateDeg failed expected %v: got %v", tc.expected.X, v2.X)
			}

			if !almostEqual(v2.Y, tc.expected.Y, 1e-10) {
				t.Errorf("V3F RotateDeg failed expected %v: got %v", tc.expected.Y, v2.Y)
			}

			if !almostEqual(v2.Z, tc.expected.Z, 1e-10) {
				t.Errorf("V3F RotateDeg failed expected %v: got %v", tc.expected.Z, v2.Z)
			}
		})
	}
}

func TestV3I_RotateDeg(t *testing.T) {
	type expected struct {
		X float64
		Y float64
		Z float64
	}

	tt := []struct {
		name     string
		degrees  float64
		axis     string
		expected expected
	}{
		{name: "z0", axis: "z", expected: expected{X: 0, Y: 1, Z: 0}},
		{name: "z90", degrees: 90, axis: "z", expected: expected{X: -1, Y: 0, Z: 0}},
		{name: "z180", degrees: 180, axis: "z", expected: expected{X: 0, Y: -1, Z: 0}},

		{name: "y0", axis: "y", expected: expected{X: 0, Y: 1, Z: 0}},
		{name: "y90", degrees: 90, axis: "y", expected: expected{X: -1, Y: 1, Z: 0}},
		{name: "y180", degrees: 180, axis: "y", expected: expected{X: 0, Y: 1, Z: 0}},

		{name: "x0", axis: "x", expected: expected{X: 0, Y: 1, Z: 0}},
		{name: "x90", degrees: 90, axis: "x", expected: expected{X: 0, Y: 0, Z: 0}},
		{name: "x180", degrees: 180, axis: "x", expected: expected{X: 0, Y: -1, Z: 0}},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v1 := V3I[int]{X: 0, Y: 1, Z: 0}
			v2 := v1.RotateDeg(tc.degrees, tc.axis)

			if !almostEqual(v2.X, tc.expected.X, 1e-10) {
				t.Errorf("V3I RotateDeg failed expected %v: got %v", tc.expected.X, v2.X)
			}

			if !almostEqual(v2.Y, tc.expected.Y, 1e-10) {
				t.Errorf("V3I RotateDeg failed expected %v: got %v", tc.expected.Y, v2.Y)
			}

			if !almostEqual(v2.Z, tc.expected.Z, 1e-10) {
				t.Errorf("V3I RotateDeg failed expected %v: got %v", tc.expected.Z, v2.Z)
			}
		})
	}
}

func TestV3F_RotateDegInPlace(t *testing.T) {
	type expected struct {
		X float64
		Y float64
		Z float64
	}

	tt := []struct {
		name     string
		degrees  float64
		axis     string
		expected expected
	}{
		{name: "z0", axis: "z", expected: expected{X: 0, Y: 1, Z: 0}},
		{name: "z90", degrees: 90, axis: "z", expected: expected{X: -1, Y: 0, Z: 0}},
		{name: "z180", degrees: 180, axis: "z", expected: expected{X: 0, Y: -1, Z: 0}},

		{name: "y0", axis: "y", expected: expected{X: 0, Y: 1, Z: 0}},
		{name: "y90", degrees: 90, axis: "y", expected: expected{X: -1, Y: 1, Z: 0}},
		{name: "y180", degrees: 180, axis: "y", expected: expected{X: 0, Y: 1, Z: 0}},

		{name: "x0", axis: "x", expected: expected{X: 0, Y: 1, Z: 0}},
		{name: "x90", degrees: 90, axis: "x", expected: expected{X: 0, Y: 0, Z: 0}},
		{name: "x180", degrees: 180, axis: "x", expected: expected{X: 0, Y: -1, Z: 0}},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v1 := V3F[float64]{X: 0, Y: 1, Z: 0}
			v1.RotateDegInPlace(tc.degrees, tc.axis)

			if !almostEqual(v1.X, tc.expected.X, 1e-10) {
				t.Errorf("V3F RotateDegInPlace failed expected %v: got %v", tc.expected.X, v1.X)
			}

			if !almostEqual(v1.Y, tc.expected.Y, 1e-10) {
				t.Errorf("V3F RotateDegInPlace failed expected %v: got %v", tc.expected.Y, v1.Y)
			}

			if !almostEqual(v1.Z, tc.expected.Z, 1e-10) {
				t.Errorf("V3F RotateDegInPlace failed expected %v: got %v", tc.expected.Z, v1.Z)
			}
		})
	}
}
