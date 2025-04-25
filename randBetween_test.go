package govec

import (
	"testing"
)

func Test_V2F_RandBetween(t *testing.T) {
	t.Run("RandBetween should generate random values within range for V2F[float32]", func(t *testing.T) {
		from := V2F[float32]{10.0, 20.0}
		to := V2F[float32]{30.0, 40.0}
		
		// Generate multiple random vectors to ensure randomness
		for i := 0; i < 10; i++ {
			r := from.RandBetween(to)
			
			// Check that components are within range [from.X, to.X] and [from.Y, to.Y]
			if r.X < from.X || r.X > to.X {
				t.Errorf("Expected X component to be in range [%f, %f], got %f", from.X, to.X, r.X)
			}
			if r.Y < from.Y || r.Y > to.Y {
				t.Errorf("Expected Y component to be in range [%f, %f], got %f", from.Y, to.Y, r.Y)
			}
		}
	})
	
	t.Run("RandBetween should generate random values within range for V2F[float64]", func(t *testing.T) {
		from := V2F[float64]{10.0, 20.0}
		to := V2F[float64]{30.0, 40.0}
		
		// Generate multiple random vectors to ensure randomness
		for i := 0; i < 10; i++ {
			r := from.RandBetween(to)
			
			// Check that components are within range [from.X, to.X] and [from.Y, to.Y]
			if r.X < from.X || r.X > to.X {
				t.Errorf("Expected X component to be in range [%f, %f], got %f", from.X, to.X, r.X)
			}
			if r.Y < from.Y || r.Y > to.Y {
				t.Errorf("Expected Y component to be in range [%f, %f], got %f", from.Y, to.Y, r.Y)
			}
		}
	})
	
	t.Run("RandBetweenComp should generate random values within range for V2F[float32]", func(t *testing.T) {
		from := V2F[float32]{10.0, 20.0}
		toX, toY := float32(30.0), float32(40.0)
		
		// Generate multiple random vectors to ensure randomness
		for i := 0; i < 10; i++ {
			r := from.RandBetweenComp(toX, toY)
			
			// Check that components are within range [from.X, toX] and [from.Y, toY]
			if r.X < from.X || r.X > toX {
				t.Errorf("Expected X component to be in range [%f, %f], got %f", from.X, toX, r.X)
			}
			if r.Y < from.Y || r.Y > toY {
				t.Errorf("Expected Y component to be in range [%f, %f], got %f", from.Y, toY, r.Y)
			}
		}
	})
	
	t.Run("RandBetweenInPlace should modify the vector in-place for V2F[float32]", func(t *testing.T) {
		from := V2F[float32]{10.0, 20.0}
		to := V2F[float32]{30.0, 40.0}
		original := V2F[float32]{from.X, from.Y}
		
		from.RandBetweenInPlace(to)
		
		// Check that components are within range [original.X, to.X] and [original.Y, to.Y]
		if from.X < original.X || from.X > to.X {
			t.Errorf("Expected X component to be in range [%f, %f], got %f", original.X, to.X, from.X)
		}
		if from.Y < original.Y || from.Y > to.Y {
			t.Errorf("Expected Y component to be in range [%f, %f], got %f", original.Y, to.Y, from.Y)
		}
	})
	
	t.Run("RandBetweenCompInPlace should modify the vector in-place for V2F[float32]", func(t *testing.T) {
		from := V2F[float32]{10.0, 20.0}
		toX, toY := float32(30.0), float32(40.0)
		original := V2F[float32]{from.X, from.Y}
		
		from.RandBetweenCompInPlace(toX, toY)
		
		// Check that components are within range [original.X, toX] and [original.Y, toY]
		if from.X < original.X || from.X > toX {
			t.Errorf("Expected X component to be in range [%f, %f], got %f", original.X, toX, from.X)
		}
		if from.Y < original.Y || from.Y > toY {
			t.Errorf("Expected Y component to be in range [%f, %f], got %f", original.Y, toY, from.Y)
		}
	})
}

func Test_V3F_RandBetween(t *testing.T) {
	t.Run("RandBetween should generate random values within range for V3F[float32]", func(t *testing.T) {
		from := V3F[float32]{10.0, 20.0, 30.0}
		to := V3F[float32]{40.0, 50.0, 60.0}
		
		// Generate multiple random vectors to ensure randomness
		for i := 0; i < 10; i++ {
			r := from.RandBetween(to)
			
			// Check that components are within range [from.X, to.X], [from.Y, to.Y], and [from.Z, to.Z]
			if r.X < from.X || r.X > to.X {
				t.Errorf("Expected X component to be in range [%f, %f], got %f", from.X, to.X, r.X)
			}
			if r.Y < from.Y || r.Y > to.Y {
				t.Errorf("Expected Y component to be in range [%f, %f], got %f", from.Y, to.Y, r.Y)
			}
			if r.Z < from.Z || r.Z > to.Z {
				t.Errorf("Expected Z component to be in range [%f, %f], got %f", from.Z, to.Z, r.Z)
			}
		}
	})
	
	t.Run("RandBetweenComp should generate random values within range for V3F[float32]", func(t *testing.T) {
		from := V3F[float32]{10.0, 20.0, 30.0}
		toX, toY, toZ := float32(40.0), float32(50.0), float32(60.0)
		
		// Generate multiple random vectors to ensure randomness
		for i := 0; i < 10; i++ {
			r := from.RandBetweenComp(toX, toY, toZ)
			
			// Check that components are within range [from.X, toX], [from.Y, toY], and [from.Z, toZ]
			if r.X < from.X || r.X > toX {
				t.Errorf("Expected X component to be in range [%f, %f], got %f", from.X, toX, r.X)
			}
			if r.Y < from.Y || r.Y > toY {
				t.Errorf("Expected Y component to be in range [%f, %f], got %f", from.Y, toY, r.Y)
			}
			if r.Z < from.Z || r.Z > toZ {
				t.Errorf("Expected Z component to be in range [%f, %f], got %f", from.Z, toZ, r.Z)
			}
		}
	})
	
	t.Run("RandBetweenInPlace should modify the vector in-place for V3F[float32]", func(t *testing.T) {
		from := V3F[float32]{10.0, 20.0, 30.0}
		to := V3F[float32]{40.0, 50.0, 60.0}
		original := V3F[float32]{from.X, from.Y, from.Z}
		
		from.RandBetweenInPlace(to)
		
		// Check that components are within range [original.X, to.X], [original.Y, to.Y], and [original.Z, to.Z]
		if from.X < original.X || from.X > to.X {
			t.Errorf("Expected X component to be in range [%f, %f], got %f", original.X, to.X, from.X)
		}
		if from.Y < original.Y || from.Y > to.Y {
			t.Errorf("Expected Y component to be in range [%f, %f], got %f", original.Y, to.Y, from.Y)
		}
		if from.Z < original.Z || from.Z > to.Z {
			t.Errorf("Expected Z component to be in range [%f, %f], got %f", original.Z, to.Z, from.Z)
		}
	})
	
	t.Run("RandBetweenCompInPlace should modify the vector in-place for V3F[float32]", func(t *testing.T) {
		from := V3F[float32]{10.0, 20.0, 30.0}
		toX, toY, toZ := float32(40.0), float32(50.0), float32(60.0)
		original := V3F[float32]{from.X, from.Y, from.Z}
		
		from.RandBetweenCompInPlace(toX, toY, toZ)
		
		// Check that components are within range [original.X, toX], [original.Y, toY], and [original.Z, toZ]
		if from.X < original.X || from.X > toX {
			t.Errorf("Expected X component to be in range [%f, %f], got %f", original.X, toX, from.X)
		}
		if from.Y < original.Y || from.Y > toY {
			t.Errorf("Expected Y component to be in range [%f, %f], got %f", original.Y, toY, from.Y)
		}
		if from.Z < original.Z || from.Z > toZ {
			t.Errorf("Expected Z component to be in range [%f, %f], got %f", original.Z, toZ, from.Z)
		}
	})
}

func Test_V2I_RandBetween(t *testing.T) {
	t.Run("RandBetween should generate random values within range for V2I[int]", func(t *testing.T) {
		from := V2I[int]{10, 20}
		to := V2I[int]{30, 40}
		
		// Generate multiple random vectors to ensure randomness
		for i := 0; i < 10; i++ {
			r := from.RandBetween(to)
			
			// Check that components are within range [from.X, to.X] and [from.Y, to.Y]
			if r.X < from.X || r.X > to.X {
				t.Errorf("Expected X component to be in range [%d, %d], got %d", from.X, to.X, r.X)
			}
			if r.Y < from.Y || r.Y > to.Y {
				t.Errorf("Expected Y component to be in range [%d, %d], got %d", from.Y, to.Y, r.Y)
			}
		}
	})
	
	t.Run("RandBetweenComp should generate random values within range for V2I[int]", func(t *testing.T) {
		from := V2I[int]{10, 20}
		toX, toY := 30, 40
		
		// Generate multiple random vectors to ensure randomness
		for i := 0; i < 10; i++ {
			r := from.RandBetweenComp(toX, toY)
			
			// Check that components are within range [from.X, toX] and [from.Y, toY]
			if r.X < from.X || r.X > toX {
				t.Errorf("Expected X component to be in range [%d, %d], got %d", from.X, toX, r.X)
			}
			if r.Y < from.Y || r.Y > toY {
				t.Errorf("Expected Y component to be in range [%d, %d], got %d", from.Y, toY, r.Y)
			}
		}
	})
	
	t.Run("RandBetweenInPlace should modify the vector in-place for V2I[int]", func(t *testing.T) {
		from := V2I[int]{10, 20}
		to := V2I[int]{30, 40}
		original := V2I[int]{from.X, from.Y}
		
		from.RandBetweenInPlace(to)
		
		// Check that components are within range [original.X, to.X] and [original.Y, to.Y]
		if from.X < original.X || from.X > to.X {
			t.Errorf("Expected X component to be in range [%d, %d], got %d", original.X, to.X, from.X)
		}
		if from.Y < original.Y || from.Y > to.Y {
			t.Errorf("Expected Y component to be in range [%d, %d], got %d", original.Y, to.Y, from.Y)
		}
	})
	
	t.Run("RandBetweenCompInPlace should modify the vector in-place for V2I[int]", func(t *testing.T) {
		from := V2I[int]{10, 20}
		toX, toY := 30, 40
		original := V2I[int]{from.X, from.Y}
		
		from.RandBetweenCompInPlace(toX, toY)
		
		// Check that components are within range [original.X, toX] and [original.Y, toY]
		if from.X < original.X || from.X > toX {
			t.Errorf("Expected X component to be in range [%d, %d], got %d", original.X, toX, from.X)
		}
		if from.Y < original.Y || from.Y > toY {
			t.Errorf("Expected Y component to be in range [%d, %d], got %d", original.Y, toY, from.Y)
		}
	})
	
	t.Run("RandBetween should handle reversed ranges correctly for V2I[int]", func(t *testing.T) {
		from := V2I[int]{30, 40}
		to := V2I[int]{10, 20}
		
		// Generate multiple random vectors to ensure randomness
		for i := 0; i < 10; i++ {
			r := from.RandBetween(to)
			
			// When from > to, the result should be equal to from
			if r.X != from.X {
				t.Errorf("Expected X component to be %d (from.X), got %d", from.X, r.X)
			}
			if r.Y != from.Y {
				t.Errorf("Expected Y component to be %d (from.Y), got %d", from.Y, r.Y)
			}
		}
	})
}

func Test_V3I_RandBetween(t *testing.T) {
	t.Run("RandBetween should generate random values within range for V3I[int]", func(t *testing.T) {
		from := V3I[int]{10, 20, 30}
		to := V3I[int]{40, 50, 60}
		
		// Generate multiple random vectors to ensure randomness
		for i := 0; i < 10; i++ {
			r := from.RandBetween(to)
			
			// Check that components are within range [from.X, to.X], [from.Y, to.Y], and [from.Z, to.Z]
			if r.X < from.X || r.X > to.X {
				t.Errorf("Expected X component to be in range [%d, %d], got %d", from.X, to.X, r.X)
			}
			if r.Y < from.Y || r.Y > to.Y {
				t.Errorf("Expected Y component to be in range [%d, %d], got %d", from.Y, to.Y, r.Y)
			}
			if r.Z < from.Z || r.Z > to.Z {
				t.Errorf("Expected Z component to be in range [%d, %d], got %d", from.Z, to.Z, r.Z)
			}
		}
	})
	
	t.Run("RandBetweenComp should generate random values within range for V3I[int]", func(t *testing.T) {
		from := V3I[int]{10, 20, 30}
		toX, toY, toZ := 40, 50, 60
		
		// Generate multiple random vectors to ensure randomness
		for i := 0; i < 10; i++ {
			r := from.RandBetweenComp(toX, toY, toZ)
			
			// Check that components are within range [from.X, toX], [from.Y, toY], and [from.Z, toZ]
			if r.X < from.X || r.X > toX {
				t.Errorf("Expected X component to be in range [%d, %d], got %d", from.X, toX, r.X)
			}
			if r.Y < from.Y || r.Y > toY {
				t.Errorf("Expected Y component to be in range [%d, %d], got %d", from.Y, toY, r.Y)
			}
			if r.Z < from.Z || r.Z > toZ {
				t.Errorf("Expected Z component to be in range [%d, %d], got %d", from.Z, toZ, r.Z)
			}
		}
	})
	
	t.Run("RandBetweenInPlace should modify the vector in-place for V3I[int]", func(t *testing.T) {
		from := V3I[int]{10, 20, 30}
		to := V3I[int]{40, 50, 60}
		original := V3I[int]{from.X, from.Y, from.Z}
		
		from.RandBetweenInPlace(to)
		
		// Check that components are within range [original.X, to.X], [original.Y, to.Y], and [original.Z, to.Z]
		if from.X < original.X || from.X > to.X {
			t.Errorf("Expected X component to be in range [%d, %d], got %d", original.X, to.X, from.X)
		}
		if from.Y < original.Y || from.Y > to.Y {
			t.Errorf("Expected Y component to be in range [%d, %d], got %d", original.Y, to.Y, from.Y)
		}
		if from.Z < original.Z || from.Z > to.Z {
			t.Errorf("Expected Z component to be in range [%d, %d], got %d", original.Z, to.Z, from.Z)
		}
	})
	
	t.Run("RandBetweenCompInPlace should modify the vector in-place for V3I[int]", func(t *testing.T) {
		from := V3I[int]{10, 20, 30}
		toX, toY, toZ := 40, 50, 60
		original := V3I[int]{from.X, from.Y, from.Z}
		
		from.RandBetweenCompInPlace(toX, toY, toZ)
		
		// Check that components are within range [original.X, toX], [original.Y, toY], and [original.Z, toZ]
		if from.X < original.X || from.X > toX {
			t.Errorf("Expected X component to be in range [%d, %d], got %d", original.X, toX, from.X)
		}
		if from.Y < original.Y || from.Y > toY {
			t.Errorf("Expected Y component to be in range [%d, %d], got %d", original.Y, toY, from.Y)
		}
		if from.Z < original.Z || from.Z > toZ {
			t.Errorf("Expected Z component to be in range [%d, %d], got %d", original.Z, toZ, from.Z)
		}
	})
	
	t.Run("RandBetween should handle reversed ranges correctly for V3I[int]", func(t *testing.T) {
		from := V3I[int]{40, 50, 60}
		to := V3I[int]{10, 20, 30}
		
		// Generate multiple random vectors to ensure randomness
		for i := 0; i < 10; i++ {
			r := from.RandBetween(to)
			
			// When from > to, the result should be equal to from
			if r.X != from.X {
				t.Errorf("Expected X component to be %d (from.X), got %d", from.X, r.X)
			}
			if r.Y != from.Y {
				t.Errorf("Expected Y component to be %d (from.Y), got %d", from.Y, r.Y)
			}
			if r.Z != from.Z {
				t.Errorf("Expected Z component to be %d (from.Z), got %d", from.Z, r.Z)
			}
		}
	})
}