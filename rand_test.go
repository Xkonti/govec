package govec

import (
	"math"
	"testing"
)

func Test_RandV2F(t *testing.T) {
	t.Run("generate 2 different random f32 vectors, shouldn't be equal", func(t *testing.T) {
		v1 := RandV2F[float32]()
		v2 := RandV2F[float32]()

		if v1.X == v2.X || v1.Y == v2.Y || v1.X == v1.Y || v2.X == v2.Y {
			t.Errorf("Expected [%f, %f] not equal to each other or with [%f, %f]", v1.X, v1.Y, v2.X, v2.Y)
		}
	})

	t.Run("generate 2 different random f64 vectors, shouldn't be equal", func(t *testing.T) {
		v1 := RandV2F[float64]()
		v2 := RandV2F[float64]()

		if v1.X == v2.X || v1.Y == v2.Y || v1.X == v1.Y || v2.X == v2.Y {
			t.Errorf("Expected [%f, %f] not equal to each other or with [%f, %f]", v1.X, v1.Y, v2.X, v2.Y)
		}
	})

	t.Run("randomized vectors should be normalized", func(t *testing.T) {
		v := RandV2F[float32]()

		if !almostEqual[float64](math.Sqrt(float64((v.X*v.X)+(v.Y*v.Y))), 1.0, 1e-5) {
			t.Errorf("Expected [%f, %f] to be normalized!", v.X, v.Y)
		}
	})
}

func Test_RandV3F(t *testing.T) {
	t.Run("generate 2 different random f32 vectors, shouldn't be equal", func(t *testing.T) {
		v1 := RandV3F[float32]()
		v2 := RandV3F[float32]()

		if v1.X == v2.X || v1.Y == v2.Y || v1.Z == v2.Z || v1.X == v1.Y || v1.X == v1.Z || v1.Y == v1.Z || v2.X == v2.Y || v2.X == v2.Z || v2.Y == v2.Z {
			t.Errorf("Expected [%f, %f, %f] not equal to each other or with [%f, %f, %f]", v1.X, v1.Y, v1.Z, v2.X, v2.Y, v2.Z)
		}
	})

	t.Run("generate 2 different random f64 vectors, shouldn't be equal", func(t *testing.T) {
		v1 := RandV3F[float64]()
		v2 := RandV3F[float64]()

		if v1.X == v2.X || v1.Y == v2.Y || v1.Z == v2.Z || v1.X == v1.Y || v1.X == v1.Z || v1.Y == v1.Z || v2.X == v2.Y || v2.X == v2.Z || v2.Y == v2.Z {
			t.Errorf("Expected [%f, %f, %f] not equal to each other or with [%f, %f, %f]", v1.X, v1.Y, v1.Z, v2.X, v2.Y, v2.Z)
		}
	})

	t.Run("randomized vectors should be normalized", func(t *testing.T) {
		v := RandV3F[float32]()

		if !almostEqual[float64](math.Sqrt(float64((v.X*v.X)+(v.Y*v.Y)+(v.Z*v.Z))), 1.0, 1e-5) {
			t.Errorf("Expected [%f, %f, %f] to be normalized!", v.X, v.Y, v.Z)
		}
	})
}
