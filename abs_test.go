package govec

import "testing"

func TestV2F_Abs(t *testing.T) {
	vectors := []V2F[float64]{
		{X: 0.0, Y: 0.0},
		{X: -1.0, Y: 0.0},
		{X: 0.0, Y: -1.0},
		{X: 1, Y: 1},
	}
	expectedAbsResult := []V2F[float64]{
		{X: 0.0, Y: 0.0},
		{X: 1.0, Y: 0.0},
		{X: 0.0, Y: 1.0},
		{X: 1.0, Y: 1.0},
	}
	for idx, vector := range vectors {
		absV := vector.Abs()
		expectedAbsV := expectedAbsResult[idx]
		if absV.X != expectedAbsV.X || absV.Y != expectedAbsV.Y {
			t.Errorf("%#v.Abs() incorrect. expected: %#v, got: %#v", vector, expectedAbsV, absV)
		}
	}
}

func TestV2F_AbsInPlace(t *testing.T) {
	vectors := []V2F[float64]{
		{X: 0.0, Y: 0.0},
		{X: -1.0, Y: 0.0},
		{X: 0.0, Y: -1.0},
		{X: 1, Y: 1},
	}
	expectedAbsInPlaceResult := []V2F[float64]{
		{X: 0.0, Y: 0.0},
		{X: 1.0, Y: 0.0},
		{X: 0.0, Y: 1.0},
		{X: 1.0, Y: 1.0},
	}
	for idx, vector := range vectors {
		vector.AbsInPlace()
		expectedAbsV := expectedAbsInPlaceResult[idx]
		if vector.X != expectedAbsV.X || vector.Y != expectedAbsV.Y {
			t.Errorf("%#v.AbsInPlace() incorrect. expected: %#v, got: %#v", vector, expectedAbsV, vector)
		}
	}
}

func TestV3F_Abs(t *testing.T) {
	vectors := []V3F[float64]{
		{X: 0.0, Y: 0.0, Z: 0.0},
		{X: -1.0, Y: 0.0, Z: 0.0},
		{X: 0.0, Y: -1.0, Z: 0.0},
		{X: 0.0, Y: 1.0, Z: -1.0},
		{X: 1, Y: 1, Z: 0},
	}
	expectedAbsResult := []V3F[float64]{
		{X: 0.0, Y: 0.0, Z: 0.0},
		{X: 1.0, Y: 0.0, Z: 0.0},
		{X: 0.0, Y: 1.0, Z: 0.0},
		{X: 0.0, Y: 1.0, Z: 1.0},
		{X: 1.0, Y: 1.0, Z: 0.0},
	}
	for idx, vector := range vectors {
		absV := vector.Abs()
		expectedAbsV := expectedAbsResult[idx]
		if absV.X != expectedAbsV.X || absV.Y != expectedAbsV.Y || absV.Z != expectedAbsV.Z {
			t.Errorf("%#v.Abs() incorrect. expected: %#v, got: %#v", vector, expectedAbsV, absV)
		}
	}
}

func TestV3F_AbsInPlace(t *testing.T) {
	vectors := []V3F[float64]{
		{X: 0.0, Y: 0.0, Z: 0.0},
		{X: -1.0, Y: 0.0, Z: 0.0},
		{X: 0.0, Y: -1.0, Z: 0.0},
		{X: 0.0, Y: 1.0, Z: -1.0},
		{X: 1, Y: 1, Z: 0},
	}
	expectedAbsInPlaceResult := []V3F[float64]{
		{X: 0.0, Y: 0.0, Z: 0.0},
		{X: 1.0, Y: 0.0, Z: 0.0},
		{X: 0.0, Y: 1.0, Z: 0.0},
		{X: 0.0, Y: 1.0, Z: 1.0},
		{X: 1.0, Y: 1.0, Z: 0.0},
	}
	for idx, vector := range vectors {
		vector.AbsInPlace()
		expectedAbsV := expectedAbsInPlaceResult[idx]
		if vector.X != expectedAbsV.X || vector.Y != expectedAbsV.Y || vector.Z != expectedAbsV.Z {
			t.Errorf("%#v.AbsInPlace() incorrect. expected: %#v, got: %#v", vector, expectedAbsV, vector)
		}
	}
}
