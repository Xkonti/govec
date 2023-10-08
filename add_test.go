package govec

import "testing"

// Test V2F Add
func TestV2F_Add(t *testing.T) {
	v1 := V2F[float64]{X: 1.4, Y: 2.4}
	v2 := V2F[float64]{X: 3.5, Y: 4.6}
	v3 := v1.Add(v2)
	if v3.X != 4.8 || v3.Y != 7.0 {
		t.Errorf("V2F Add failed")
	}
}

func TestV2F_AddComp(t *testing.T) {
	v1 := V2F[float64]{X: 1.3, Y: 2.4}
	v3 := v1.AddComp(3.5, 4.6)
	if v3.X != 4.8 || v3.Y != 7.0 {
		t.Errorf("V2F AddComp failed")
	}
}

func TestV2F_AddInPlace(t *testing.T) {
	v1 := V2F[float64]{X: 1.3, Y: 2.4}
	v2 := V2F[float64]{X: 3.5, Y: 4.6}
	v1.AddInPlace(v2)
	if v1.X != 4.8 || v1.Y != 7.0 {
		t.Errorf("V2F AddInPlace failed")
	}
}

func TestV2F_AddCompInPlace(t *testing.T) {
	v1 := V2F[float64]{X: 1.3, Y: 2.4}
	v1.AddCompInPlace(3.5, 4.6)
	if v1.X != 4.8 || v1.Y != 7.0 {
		t.Errorf("V2F AddCompInPlace failed")
	}
}

// Test V2I Add
func TestV2I_Add(t *testing.T) {
	v1 := V2I[int16]{X: 3, Y: 24}
	v2 := V2I[int16]{X: 67, Y: 100}
	v3 := v1.Add(v2)
	if v3.X != 70 || v3.Y != 124 {
		t.Errorf("V2I Add failed")
	}
}

func TestV2I_AddComp(t *testing.T) {
	v1 := V2I[int16]{X: 3, Y: 24}
	v3 := v1.AddComp(67, 100)
	if v3.X != 70 || v3.Y != 124 {
		t.Errorf("V2I AddComp failed")
	}
}

func TestV2I_AddInPlace(t *testing.T) {
	v1 := V2I[int16]{X: 3, Y: 24}
	v2 := V2I[int16]{X: 67, Y: 100}
	v1.AddInPlace(v2)
	if v1.X != 70 || v1.Y != 124 {
		t.Errorf("V2I AddInPlace failed")
	}
}

func TestV2I_AddCompInPlace(t *testing.T) {
	v1 := V2I[int16]{X: 3, Y: 24}
	v1.AddCompInPlace(67, 100)
	if v1.X != 70 || v1.Y != 124 {
		t.Errorf("V2I AddCompInPlace failed")
	}
}
