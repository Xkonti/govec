package govec

import (
	"testing"
)

func Test_V2F_RandIn(t *testing.T) {
	t.Run("RandIn should generate random values within range for V2F[float32]", func(t *testing.T) {
		v := V2F[float32]{10.0, 20.0}
		
		// Generate multiple random vectors to ensure randomness
		for i := 0; i < 10; i++ {
			r := v.RandIn()
			
			// Check that components are within range [0, v.X] and [0, v.Y]
			if r.X < 0 || r.X > v.X {
				t.Errorf("Expected X component to be in range [0, %f], got %f", v.X, r.X)
			}
			if r.Y < 0 || r.Y > v.Y {
				t.Errorf("Expected Y component to be in range [0, %f], got %f", v.Y, r.Y)
			}
		}
	})
	
	t.Run("RandIn should generate random values within range for V2F[float64]", func(t *testing.T) {
		v := V2F[float64]{10.0, 20.0}
		
		// Generate multiple random vectors to ensure randomness
		for i := 0; i < 10; i++ {
			r := v.RandIn()
			
			// Check that components are within range [0, v.X] and [0, v.Y]
			if r.X < 0 || r.X > v.X {
				t.Errorf("Expected X component to be in range [0, %f], got %f", v.X, r.X)
			}
			if r.Y < 0 || r.Y > v.Y {
				t.Errorf("Expected Y component to be in range [0, %f], got %f", v.Y, r.Y)
			}
		}
	})
	
	t.Run("RandInInPlace should modify the vector in-place for V2F[float32]", func(t *testing.T) {
		v := V2F[float32]{10.0, 20.0}
		original := V2F[float32]{v.X, v.Y}
		
		v.RandInInPlace()
		
		// Check that components are within range [0, original.X] and [0, original.Y]
		if v.X < 0 || v.X > original.X {
			t.Errorf("Expected X component to be in range [0, %f], got %f", original.X, v.X)
		}
		if v.Y < 0 || v.Y > original.Y {
			t.Errorf("Expected Y component to be in range [0, %f], got %f", original.Y, v.Y)
		}
	})
}

func Test_V3F_RandIn(t *testing.T) {
	t.Run("RandIn should generate random values within range for V3F[float32]", func(t *testing.T) {
		v := V3F[float32]{10.0, 20.0, 30.0}
		
		// Generate multiple random vectors to ensure randomness
		for i := 0; i < 10; i++ {
			r := v.RandIn()
			
			// Check that components are within range [0, v.X], [0, v.Y], and [0, v.Z]
			if r.X < 0 || r.X > v.X {
				t.Errorf("Expected X component to be in range [0, %f], got %f", v.X, r.X)
			}
			if r.Y < 0 || r.Y > v.Y {
				t.Errorf("Expected Y component to be in range [0, %f], got %f", v.Y, r.Y)
			}
			if r.Z < 0 || r.Z > v.Z {
				t.Errorf("Expected Z component to be in range [0, %f], got %f", v.Z, r.Z)
			}
		}
	})
	
	t.Run("RandIn should generate random values within range for V3F[float64]", func(t *testing.T) {
		v := V3F[float64]{10.0, 20.0, 30.0}
		
		// Generate multiple random vectors to ensure randomness
		for i := 0; i < 10; i++ {
			r := v.RandIn()
			
			// Check that components are within range [0, v.X], [0, v.Y], and [0, v.Z]
			if r.X < 0 || r.X > v.X {
				t.Errorf("Expected X component to be in range [0, %f], got %f", v.X, r.X)
			}
			if r.Y < 0 || r.Y > v.Y {
				t.Errorf("Expected Y component to be in range [0, %f], got %f", v.Y, r.Y)
			}
			if r.Z < 0 || r.Z > v.Z {
				t.Errorf("Expected Z component to be in range [0, %f], got %f", v.Z, r.Z)
			}
		}
	})
	
	t.Run("RandInInPlace should modify the vector in-place for V3F[float32]", func(t *testing.T) {
		v := V3F[float32]{10.0, 20.0, 30.0}
		original := V3F[float32]{v.X, v.Y, v.Z}
		
		v.RandInInPlace()
		
		// Check that components are within range [0, original.X], [0, original.Y], and [0, original.Z]
		if v.X < 0 || v.X > original.X {
			t.Errorf("Expected X component to be in range [0, %f], got %f", original.X, v.X)
		}
		if v.Y < 0 || v.Y > original.Y {
			t.Errorf("Expected Y component to be in range [0, %f], got %f", original.Y, v.Y)
		}
		if v.Z < 0 || v.Z > original.Z {
			t.Errorf("Expected Z component to be in range [0, %f], got %f", original.Z, v.Z)
		}
	})
}

func Test_V2I_RandIn(t *testing.T) {
	t.Run("RandIn should generate random values within range for V2I[int]", func(t *testing.T) {
		v := V2I[int]{10, 20}
		
		// Generate multiple random vectors to ensure randomness
		for i := 0; i < 10; i++ {
			r := v.RandIn()
			
			// Check that components are within range [0, v.X] and [0, v.Y]
			if r.X < 0 || r.X > v.X {
				t.Errorf("Expected X component to be in range [0, %d], got %d", v.X, r.X)
			}
			if r.Y < 0 || r.Y > v.Y {
				t.Errorf("Expected Y component to be in range [0, %d], got %d", v.Y, r.Y)
			}
		}
	})
	
	t.Run("RandInInPlace should modify the vector in-place for V2I[int]", func(t *testing.T) {
		v := V2I[int]{10, 20}
		original := V2I[int]{v.X, v.Y}
		
		v.RandInInPlace()
		
		// Check that components are within range [0, original.X] and [0, original.Y]
		if v.X < 0 || v.X > original.X {
			t.Errorf("Expected X component to be in range [0, %d], got %d", original.X, v.X)
		}
		if v.Y < 0 || v.Y > original.Y {
			t.Errorf("Expected Y component to be in range [0, %d], got %d", original.Y, v.Y)
		}
	})
	
	t.Run("RandIn should generate random values within range for V2I[int64]", func(t *testing.T) {
		v := V2I[int64]{10, 20}
		
		// Generate multiple random vectors to ensure randomness
		for i := 0; i < 10; i++ {
			r := v.RandIn()
			
			// Check that components are within range [0, v.X] and [0, v.Y]
			if r.X < 0 || r.X > v.X {
				t.Errorf("Expected X component to be in range [0, %d], got %d", v.X, r.X)
			}
			if r.Y < 0 || r.Y > v.Y {
				t.Errorf("Expected Y component to be in range [0, %d], got %d", v.Y, r.Y)
			}
		}
	})
}

func Test_V3I_RandIn(t *testing.T) {
	t.Run("RandIn should generate random values within range for V3I[int]", func(t *testing.T) {
		v := V3I[int]{10, 20, 30}
		
		// Generate multiple random vectors to ensure randomness
		for i := 0; i < 10; i++ {
			r := v.RandIn()
			
			// Check that components are within range [0, v.X], [0, v.Y], and [0, v.Z]
			if r.X < 0 || r.X > v.X {
				t.Errorf("Expected X component to be in range [0, %d], got %d", v.X, r.X)
			}
			if r.Y < 0 || r.Y > v.Y {
				t.Errorf("Expected Y component to be in range [0, %d], got %d", v.Y, r.Y)
			}
			if r.Z < 0 || r.Z > v.Z {
				t.Errorf("Expected Z component to be in range [0, %d], got %d", v.Z, r.Z)
			}
		}
	})
	
	t.Run("RandInInPlace should modify the vector in-place for V3I[int]", func(t *testing.T) {
		v := V3I[int]{10, 20, 30}
		original := V3I[int]{v.X, v.Y, v.Z}
		
		v.RandInInPlace()
		
		// Check that components are within range [0, original.X], [0, original.Y], and [0, original.Z]
		if v.X < 0 || v.X > original.X {
			t.Errorf("Expected X component to be in range [0, %d], got %d", original.X, v.X)
		}
		if v.Y < 0 || v.Y > original.Y {
			t.Errorf("Expected Y component to be in range [0, %d], got %d", original.Y, v.Y)
		}
		if v.Z < 0 || v.Z > original.Z {
			t.Errorf("Expected Z component to be in range [0, %d], got %d", original.Z, v.Z)
		}
	})
	
	t.Run("RandIn should generate random values within range for V3I[int64]", func(t *testing.T) {
		v := V3I[int64]{10, 20, 30}
		
		// Generate multiple random vectors to ensure randomness
		for i := 0; i < 10; i++ {
			r := v.RandIn()
			
			// Check that components are within range [0, v.X], [0, v.Y], and [0, v.Z]
			if r.X < 0 || r.X > v.X {
				t.Errorf("Expected X component to be in range [0, %d], got %d", v.X, r.X)
			}
			if r.Y < 0 || r.Y > v.Y {
				t.Errorf("Expected Y component to be in range [0, %d], got %d", v.Y, r.Y)
			}
			if r.Z < 0 || r.Z > v.Z {
				t.Errorf("Expected Z component to be in range [0, %d], got %d", v.Z, r.Z)
			}
		}
	})
}