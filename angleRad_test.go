package govec

import (
	"fmt"
	"math"
	"testing"
)

func TestAngleRad(t *testing.T) {
	testCases := []struct {
		a, X, Y float64
	}{
		{-180.0, -1, 0},
		{-135.0, -1 / math.Sqrt2, -1 / math.Sqrt2},
		{-90.0, 0, -1},
		{-45.0, 1 / math.Sqrt2, -1 / math.Sqrt2},
		{0, 1, 0},
		{45.0, 1 / math.Sqrt2, 1 / math.Sqrt2},
		{90.0, 0, 1},
		{135.0, -1 / math.Sqrt2, 1 / math.Sqrt2},
		{180, -1, 0},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%f", tc.a), func(t *testing.T) {
			v := AngleRad(tc.a / radToDegFactor)

			if !almostEqual(v.X, tc.X, 1e-10) {
				t.Errorf("AngleRad %f radians | v.X failed", tc.a)
			}
			if !almostEqual(v.Y, tc.Y, 1e-10) {
				t.Errorf("AngleRad %f radians | v.Y failed", tc.a)
			}
		})
	}

}
