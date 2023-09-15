package govec

import (
	"golang.org/x/exp/constraints"
	"testing"
)

func runMinTestsFloat[T constraints.Float](t *testing.T, unitName string) {
	testCases := []struct {
		name                            string
		v1x, v1y, v1z                   T
		v2x, v2y, v2z                   T
		expectedX, expectedY, expectedZ T
	}{
		{"all same", 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{"v1 max", 2, 3, 4, 1, 2, 3, 1, 2, 3},
		{"v2 max", 1, 2, 3, 2, 3, 4, 1, 2, 3},
		{"v1 max neg", -1, -2, -3, -2, -3, -4, -2, -3, -4},
		{"v2 max neg", -2, -3, -4, -1, -2, -3, -2, -3, -4},
		{"v1 max mixed", -1, 2, -3, -2, 1, -4, -2, 1, -4},
		{"v2 max mixed", -2, 1, -4, -1, 2, -3, -2, 1, -4},
		{"rand 1", 3, 5, -3, 15, 2, 4, 3, 2, -3},
		{"rand 2", 120, 127, -2, -51, 36, 1, -51, 36, -2},
	}

	for _, tc := range testCases {
		t.Run(tc.name+" V2F["+unitName+"]", func(t *testing.T) {
			// Bring to proper type
			v1x, v1y := tc.v1x, tc.v1y
			v2x, v2y := tc.v2x, tc.v2y
			expectedX, expectedY := tc.expectedX, tc.expectedY

			expected := V2F[T]{X: expectedX, Y: expectedY}

			// Results of operation variants
			vMin := V2F[T]{X: v1x, Y: v1y}.Min(V2F[T]{X: v2x, Y: v2y})
			vMinComp := V2F[T]{X: v1x, Y: v1y}.MinComp(v2x, v2y)
			vMinInPlace := V2F[T]{X: v1x, Y: v1y}
			vMinInPlace.MinInPlace(V2F[T]{X: v2x, Y: v2y})
			vMinCompInPlace := V2F[T]{X: v1x, Y: v1y}
			vMinCompInPlace.MinCompInPlace(v2x, v2y)
			results := []V2F[T]{vMin, vMinComp, vMinInPlace, vMinCompInPlace}

			// Check results
			for i, result := range results {
				if !almostEqual(result.X, expected.X, 1e-9) || !almostEqual(result.Y, expected.Y, 1e-9) {
					t.Errorf("Expected %v, got %v when testing variant #%v", expected, result, i)
				}
			}
		})

		t.Run(tc.name+" V3F["+unitName+"]", func(t *testing.T) {
			// Bring to proper type
			v1x, v1y, v1z := tc.v1x, tc.v1y, tc.v1z
			v2x, v2y, v2z := tc.v2x, tc.v2y, tc.v2z
			expectedX, expectedY, expectedZ := tc.expectedX, tc.expectedY, tc.expectedZ

			expected := V3F[T]{X: expectedX, Y: expectedY, Z: expectedZ}

			// Results of operation variants
			vMin := V3F[T]{X: v1x, Y: v1y, Z: v1z}.Min(V3F[T]{X: v2x, Y: v2y, Z: v2z})
			vMinComp := V3F[T]{X: v1x, Y: v1y, Z: v1z}.MinComp(v2x, v2y, v2z)
			vMinInPlace := V3F[T]{X: v1x, Y: v1y, Z: v1z}
			vMinInPlace.MinInPlace(V3F[T]{X: v2x, Y: v2y, Z: v2z})
			vMinCompInPlace := V3F[T]{X: v1x, Y: v1y, Z: v1z}
			vMinCompInPlace.MinCompInPlace(v2x, v2y, v2z)
			results := []V3F[T]{vMin, vMinComp, vMinInPlace, vMinCompInPlace}

			// Check results
			for i, result := range results {
				if !almostEqual(result.X, expected.X, 1e-9) || !almostEqual(result.Y, expected.Y, 1e-9) || !almostEqual(result.Z, expected.Z, 1e-9) {
					t.Errorf("Expected %v, got %v when testing variant #%v", expected, result, i)
				}
			}
		})
	}
}

func runMinTestsInteger[T constraints.Signed](t *testing.T, unitName string) {
	testCases := []struct {
		name                            string
		v1x, v1y, v1z                   T
		v2x, v2y, v2z                   T
		expectedX, expectedY, expectedZ T
	}{
		{"all same", 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{"v1 max", 2, 3, 4, 1, 2, 3, 1, 2, 3},
		{"v2 max", 1, 2, 3, 2, 3, 4, 1, 2, 3},
		{"v1 max neg", -1, -2, -3, -2, -3, -4, -2, -3, -4},
		{"v2 max neg", -2, -3, -4, -1, -2, -3, -2, -3, -4},
		{"v1 max mixed", -1, 2, -3, -2, 1, -4, -2, 1, -4},
		{"v2 max mixed", -2, 1, -4, -1, 2, -3, -2, 1, -4},
		{"rand 1", 3, 5, -3, 15, 2, 4, 3, 2, -3},
		{"rand 2", 120, 127, -2, -51, 36, 1, -51, 36, -2},
	}

	for _, tc := range testCases {
		t.Run(tc.name+" V2I["+unitName+"]", func(t *testing.T) {
			// Bring to proper type
			v1x, v1y := tc.v1x, tc.v1y
			v2x, v2y := tc.v2x, tc.v2y
			expectedX, expectedY := tc.expectedX, tc.expectedY

			expected := V2I[T]{X: expectedX, Y: expectedY}

			// Results of operation variants
			vMin := V2I[T]{X: v1x, Y: v1y}.Min(V2I[T]{X: v2x, Y: v2y})
			vMinComp := V2I[T]{X: v1x, Y: v1y}.MinComp(v2x, v2y)
			vMinInPlace := V2I[T]{X: v1x, Y: v1y}
			vMinInPlace.MinInPlace(V2I[T]{X: v2x, Y: v2y})
			vMinCompInPlace := V2I[T]{X: v1x, Y: v1y}
			vMinCompInPlace.MinCompInPlace(v2x, v2y)
			results := []V2I[T]{vMin, vMinComp, vMinInPlace, vMinCompInPlace}

			// Check results
			for i, result := range results {
				if result.X != expected.X || result.Y != expected.Y {
					t.Errorf("Expected %v, got %v when testing variant #%v", expected, result, i)
				}
			}
		})

		t.Run(tc.name+" V3I["+unitName+"]", func(t *testing.T) {
			// Bring to proper type
			v1x, v1y, v1z := tc.v1x, tc.v1y, tc.v1z
			v2x, v2y, v2z := tc.v2x, tc.v2y, tc.v2z
			expectedX, expectedY, expectedZ := tc.expectedX, tc.expectedY, tc.expectedZ

			expected := V3I[T]{X: expectedX, Y: expectedY, Z: expectedZ}

			// Results of operation variants
			vMin := V3I[T]{X: v1x, Y: v1y, Z: v1z}.Min(V3I[T]{X: v2x, Y: v2y, Z: v2z})
			vMinComp := V3I[T]{X: v1x, Y: v1y, Z: v1z}.MinComp(v2x, v2y, v2z)
			vMinInPlace := V3I[T]{X: v1x, Y: v1y, Z: v1z}
			vMinInPlace.MinInPlace(V3I[T]{X: v2x, Y: v2y, Z: v2z})
			vMinCompInPlace := V3I[T]{X: v1x, Y: v1y, Z: v1z}
			vMinCompInPlace.MinCompInPlace(v2x, v2y, v2z)
			results := []V3I[T]{vMin, vMinComp, vMinInPlace, vMinCompInPlace}

			// Check results
			for i, result := range results {
				if result.X != expected.X || result.Y != expected.Y || result.Z != expected.Z {
					t.Errorf("Expected %v, got %v when testing variant #%v", expected, result, i)
				}
			}
		})
	}
}

func TestMin(t *testing.T) {
	runMinTestsFloat[float64](t, "float64")
	runMinTestsFloat[float32](t, "float32")
	runMinTestsInteger[int64](t, "int64")
	runMinTestsInteger[int32](t, "int32")
	runMinTestsInteger[int16](t, "int16")
	runMinTestsInteger[int8](t, "int8")
	runMinTestsInteger[int](t, "int")
}
