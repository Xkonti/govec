package govec

import (
	"math"
	"testing"
)

// Test V2F Lerp
func TestV2F_Lerp(t *testing.T) {
	v1 := V2F[float64]{X: 1.0, Y: 2.0}
	v2 := V2F[float64]{X: 5.0, Y: 10.0}

	// Test t = 0 (should return v1)
	v3 := v1.Lerp(v2, 0.0)
	if v3.X != 1.0 || v3.Y != 2.0 {
		t.Errorf("V2F Lerp with t=0 failed, got %v, expected %v", v3, v1)
	}

	// Test t = 1 (should return v2)
	v3 = v1.Lerp(v2, 1.0)
	if v3.X != 5.0 || v3.Y != 10.0 {
		t.Errorf("V2F Lerp with t=1 failed, got %v, expected %v", v3, v2)
	}

	// Test t = 0.5 (should return midpoint)
	v3 = v1.Lerp(v2, 0.5)
	if v3.X != 3.0 || v3.Y != 6.0 {
		t.Errorf("V2F Lerp with t=0.5 failed, got %v, expected {3.0, 6.0}", v3)
	}

	// Test t = 0.25
	v3 = v1.Lerp(v2, 0.25)
	if v3.X != 2.0 || v3.Y != 4.0 {
		t.Errorf("V2F Lerp with t=0.25 failed, got %v, expected {2.0, 4.0}", v3)
	}
}

func TestV2F_LerpInPlace(t *testing.T) {
	v1 := V2F[float64]{X: 1.0, Y: 2.0}
	v2 := V2F[float64]{X: 5.0, Y: 10.0}
	original := V2F[float64]{X: 1.0, Y: 2.0}

	// Test t = 0 (should leave v1 unchanged)
	v1.LerpInPlace(v2, 0.0)
	if v1.X != original.X || v1.Y != original.Y {
		t.Errorf("V2F LerpInPlace with t=0 failed, got %v, expected %v", v1, original)
	}

	// Test t = 1 (should set v1 to v2)
	v1 = original
	v1.LerpInPlace(v2, 1.0)
	if v1.X != v2.X || v1.Y != v2.Y {
		t.Errorf("V2F LerpInPlace with t=1 failed, got %v, expected %v", v1, v2)
	}

	// Test t = 0.5 (should set v1 to midpoint)
	v1 = original
	v1.LerpInPlace(v2, 0.5)
	if v1.X != 3.0 || v1.Y != 6.0 {
		t.Errorf("V2F LerpInPlace with t=0.5 failed, got %v, expected {3.0, 6.0}", v1)
	}
}

// Test V3F Lerp
func TestV3F_Lerp(t *testing.T) {
	v1 := V3F[float64]{X: 1.0, Y: 2.0, Z: 3.0}
	v2 := V3F[float64]{X: 5.0, Y: 10.0, Z: 15.0}

	// Test t = 0 (should return v1)
	v3 := v1.Lerp(v2, 0.0)
	if v3.X != 1.0 || v3.Y != 2.0 || v3.Z != 3.0 {
		t.Errorf("V3F Lerp with t=0 failed, got %v, expected %v", v3, v1)
	}

	// Test t = 1 (should return v2)
	v3 = v1.Lerp(v2, 1.0)
	if v3.X != 5.0 || v3.Y != 10.0 || v3.Z != 15.0 {
		t.Errorf("V3F Lerp with t=1 failed, got %v, expected %v", v3, v2)
	}

	// Test t = 0.5 (should return midpoint)
	v3 = v1.Lerp(v2, 0.5)
	if v3.X != 3.0 || v3.Y != 6.0 || v3.Z != 9.0 {
		t.Errorf("V3F Lerp with t=0.5 failed, got %v, expected {3.0, 6.0, 9.0}", v3)
	}
}

func TestV3F_LerpInPlace(t *testing.T) {
	v1 := V3F[float64]{X: 1.0, Y: 2.0, Z: 3.0}
	v2 := V3F[float64]{X: 5.0, Y: 10.0, Z: 15.0}
	original := V3F[float64]{X: 1.0, Y: 2.0, Z: 3.0}

	// Test t = 0 (should leave v1 unchanged)
	v1.LerpInPlace(v2, 0.0)
	if v1.X != original.X || v1.Y != original.Y || v1.Z != original.Z {
		t.Errorf("V3F LerpInPlace with t=0 failed, got %v, expected %v", v1, original)
	}

	// Test t = 1 (should set v1 to v2)
	v1 = original
	v1.LerpInPlace(v2, 1.0)
	if v1.X != v2.X || v1.Y != v2.Y || v1.Z != v2.Z {
		t.Errorf("V3F LerpInPlace with t=1 failed, got %v, expected %v", v1, v2)
	}

	// Test t = 0.5 (should set v1 to midpoint)
	v1 = original
	v1.LerpInPlace(v2, 0.5)
	if v1.X != 3.0 || v1.Y != 6.0 || v1.Z != 9.0 {
		t.Errorf("V3F LerpInPlace with t=0.5 failed, got %v, expected {3.0, 6.0, 9.0}", v1)
	}
}

// Test V2I Lerp
func TestV2I_Lerp(t *testing.T) {
	v1 := V2I[int]{X: 1, Y: 2}
	v2 := V2I[int]{X: 5, Y: 10}

	// Test t = 0 (should return v1 as float)
	v3 := v1.Lerp(v2, 0.0)
	if v3.X != 1.0 || v3.Y != 2.0 {
		t.Errorf("V2I Lerp with t=0 failed, got %v, expected {1.0, 2.0}", v3)
	}

	// Test t = 1 (should return v2 as float)
	v3 = v1.Lerp(v2, 1.0)
	if v3.X != 5.0 || v3.Y != 10.0 {
		t.Errorf("V2I Lerp with t=1 failed, got %v, expected {5.0, 10.0}", v3)
	}

	// Test t = 0.5 (should return midpoint as float)
	v3 = v1.Lerp(v2, 0.5)
	if v3.X != 3.0 || v3.Y != 6.0 {
		t.Errorf("V2I Lerp with t=0.5 failed, got %v, expected {3.0, 6.0}", v3)
	}

	// Test t = 0.25 (should return quarter point as float)
	v3 = v1.Lerp(v2, 0.25)
	if math.Abs(v3.X-2.0) > 1e-10 || math.Abs(v3.Y-4.0) > 1e-10 {
		t.Errorf("V2I Lerp with t=0.25 failed, got %v, expected {2.0, 4.0}", v3)
	}
}

// Test V3I Lerp
func TestV3I_Lerp(t *testing.T) {
	v1 := V3I[int]{X: 1, Y: 2, Z: 3}
	v2 := V3I[int]{X: 5, Y: 10, Z: 15}

	// Test t = 0 (should return v1 as float)
	v3 := v1.Lerp(v2, 0.0)
	if v3.X != 1.0 || v3.Y != 2.0 || v3.Z != 3.0 {
		t.Errorf("V3I Lerp with t=0 failed, got %v, expected {1.0, 2.0, 3.0}", v3)
	}

	// Test t = 1 (should return v2 as float)
	v3 = v1.Lerp(v2, 1.0)
	if v3.X != 5.0 || v3.Y != 10.0 || v3.Z != 15.0 {
		t.Errorf("V3I Lerp with t=1 failed, got %v, expected {5.0, 10.0, 15.0}", v3)
	}

	// Test t = 0.5 (should return midpoint as float)
	v3 = v1.Lerp(v2, 0.5)
	if v3.X != 3.0 || v3.Y != 6.0 || v3.Z != 9.0 {
		t.Errorf("V3I Lerp with t=0.5 failed, got %v, expected {3.0, 6.0, 9.0}", v3)
	}
}
