package govec

import (
        "math/rand"
        "time"
)

// RandBetween generates a vector with random components within a range specified by two vectors.
// Each component of the returned vector will be a random value between the corresponding components of the two vectors.
func (v V2F[T]) RandBetween(to V2F[T]) V2F[T] {
        r := rand.New(rand.NewSource(time.Now().UnixNano()))
        return V2F[T]{
                X: v.X + T(r.Float64())*(to.X-v.X),
                Y: v.Y + T(r.Float64())*(to.Y-v.Y),
        }
}

// RandBetweenComp generates a vector with random components within a range specified by the vector and the provided components.
// Each component of the returned vector will be a random value between the corresponding component of the vector and the provided component.
func (v V2F[T]) RandBetweenComp(toX, toY T) V2F[T] {
        r := rand.New(rand.NewSource(time.Now().UnixNano()))
        return V2F[T]{
                X: v.X + T(r.Float64())*(toX-v.X),
                Y: v.Y + T(r.Float64())*(toY-v.Y),
        }
}

// RandBetweenInPlace updates the vector with random components within a range specified by the vector and another vector.
// Each component of the vector will be set to a random value between its current value and the corresponding component of the provided vector.
func (v *V2F[T]) RandBetweenInPlace(to V2F[T]) {
        r := rand.New(rand.NewSource(time.Now().UnixNano()))
        v.X = v.X + T(r.Float64())*(to.X-v.X)
        v.Y = v.Y + T(r.Float64())*(to.Y-v.Y)
}

// RandBetweenCompInPlace updates the vector with random components within a range specified by the vector and the provided components.
// Each component of the vector will be set to a random value between its current value and the provided component.
func (v *V2F[T]) RandBetweenCompInPlace(toX, toY T) {
        r := rand.New(rand.NewSource(time.Now().UnixNano()))
        v.X = v.X + T(r.Float64())*(toX-v.X)
        v.Y = v.Y + T(r.Float64())*(toY-v.Y)
}

// RandBetween generates a vector with random components within a range specified by two vectors.
// Each component of the returned vector will be a random value between the corresponding components of the two vectors.
func (v V3F[T]) RandBetween(to V3F[T]) V3F[T] {
        r := rand.New(rand.NewSource(time.Now().UnixNano()))
        return V3F[T]{
                X: v.X + T(r.Float64())*(to.X-v.X),
                Y: v.Y + T(r.Float64())*(to.Y-v.Y),
                Z: v.Z + T(r.Float64())*(to.Z-v.Z),
        }
}

// RandBetweenComp generates a vector with random components within a range specified by the vector and the provided components.
// Each component of the returned vector will be a random value between the corresponding component of the vector and the provided component.
func (v V3F[T]) RandBetweenComp(toX, toY, toZ T) V3F[T] {
        r := rand.New(rand.NewSource(time.Now().UnixNano()))
        return V3F[T]{
                X: v.X + T(r.Float64())*(toX-v.X),
                Y: v.Y + T(r.Float64())*(toY-v.Y),
                Z: v.Z + T(r.Float64())*(toZ-v.Z),
        }
}

// RandBetweenInPlace updates the vector with random components within a range specified by the vector and another vector.
// Each component of the vector will be set to a random value between its current value and the corresponding component of the provided vector.
func (v *V3F[T]) RandBetweenInPlace(to V3F[T]) {
        r := rand.New(rand.NewSource(time.Now().UnixNano()))
        v.X = v.X + T(r.Float64())*(to.X-v.X)
        v.Y = v.Y + T(r.Float64())*(to.Y-v.Y)
        v.Z = v.Z + T(r.Float64())*(to.Z-v.Z)
}

// RandBetweenCompInPlace updates the vector with random components within a range specified by the vector and the provided components.
// Each component of the vector will be set to a random value between its current value and the provided component.
func (v *V3F[T]) RandBetweenCompInPlace(toX, toY, toZ T) {
        r := rand.New(rand.NewSource(time.Now().UnixNano()))
        v.X = v.X + T(r.Float64())*(toX-v.X)
        v.Y = v.Y + T(r.Float64())*(toY-v.Y)
        v.Z = v.Z + T(r.Float64())*(toZ-v.Z)
}

// RandBetween generates a vector with random components within a range specified by two vectors.
// Each component of the returned vector will be a random value between the corresponding components of the two vectors.
func (v V2I[T]) RandBetween(to V2I[T]) V2I[T] {
        r := rand.New(rand.NewSource(time.Now().UnixNano()))
        
        // Calculate the range for each component
        rangeX := int64(to.X) - int64(v.X)
        rangeY := int64(to.Y) - int64(v.Y)
        
        // Generate random values within the range
        randX := int64(v.X)
        if rangeX > 0 {
                randX += r.Int63n(rangeX + 1)
        }
        
        randY := int64(v.Y)
        if rangeY > 0 {
                randY += r.Int63n(rangeY + 1)
        }
        
        return V2I[T]{
                X: T(randX),
                Y: T(randY),
        }
}

// RandBetweenComp generates a vector with random components within a range specified by the vector and the provided components.
// Each component of the returned vector will be a random value between the corresponding component of the vector and the provided component.
func (v V2I[T]) RandBetweenComp(toX, toY T) V2I[T] {
        r := rand.New(rand.NewSource(time.Now().UnixNano()))
        
        // Calculate the range for each component
        rangeX := int64(toX) - int64(v.X)
        rangeY := int64(toY) - int64(v.Y)
        
        // Generate random values within the range
        randX := int64(v.X)
        if rangeX > 0 {
                randX += r.Int63n(rangeX + 1)
        }
        
        randY := int64(v.Y)
        if rangeY > 0 {
                randY += r.Int63n(rangeY + 1)
        }
        
        return V2I[T]{
                X: T(randX),
                Y: T(randY),
        }
}

// RandBetweenInPlace updates the vector with random components within a range specified by the vector and another vector.
// Each component of the vector will be set to a random value between its current value and the corresponding component of the provided vector.
func (v *V2I[T]) RandBetweenInPlace(to V2I[T]) {
        r := rand.New(rand.NewSource(time.Now().UnixNano()))
        
        // Calculate the range for each component
        rangeX := int64(to.X) - int64(v.X)
        rangeY := int64(to.Y) - int64(v.Y)
        
        // Generate random values within the range and update the vector
        if rangeX > 0 {
                v.X = T(int64(v.X) + r.Int63n(rangeX+1))
        }
        
        if rangeY > 0 {
                v.Y = T(int64(v.Y) + r.Int63n(rangeY+1))
        }
}

// RandBetweenCompInPlace updates the vector with random components within a range specified by the vector and the provided components.
// Each component of the vector will be set to a random value between its current value and the provided component.
func (v *V2I[T]) RandBetweenCompInPlace(toX, toY T) {
        r := rand.New(rand.NewSource(time.Now().UnixNano()))
        
        // Calculate the range for each component
        rangeX := int64(toX) - int64(v.X)
        rangeY := int64(toY) - int64(v.Y)
        
        // Generate random values within the range and update the vector
        if rangeX > 0 {
                v.X = T(int64(v.X) + r.Int63n(rangeX+1))
        }
        
        if rangeY > 0 {
                v.Y = T(int64(v.Y) + r.Int63n(rangeY+1))
        }
}

// RandBetween generates a vector with random components within a range specified by two vectors.
// Each component of the returned vector will be a random value between the corresponding components of the two vectors.
func (v V3I[T]) RandBetween(to V3I[T]) V3I[T] {
        r := rand.New(rand.NewSource(time.Now().UnixNano()))
        
        // Calculate the range for each component
        rangeX := int64(to.X) - int64(v.X)
        rangeY := int64(to.Y) - int64(v.Y)
        rangeZ := int64(to.Z) - int64(v.Z)
        
        // Generate random values within the range
        randX := int64(v.X)
        if rangeX > 0 {
                randX += r.Int63n(rangeX + 1)
        }
        
        randY := int64(v.Y)
        if rangeY > 0 {
                randY += r.Int63n(rangeY + 1)
        }
        
        randZ := int64(v.Z)
        if rangeZ > 0 {
                randZ += r.Int63n(rangeZ + 1)
        }
        
        return V3I[T]{
                X: T(randX),
                Y: T(randY),
                Z: T(randZ),
        }
}

// RandBetweenComp generates a vector with random components within a range specified by the vector and the provided components.
// Each component of the returned vector will be a random value between the corresponding component of the vector and the provided component.
func (v V3I[T]) RandBetweenComp(toX, toY, toZ T) V3I[T] {
        r := rand.New(rand.NewSource(time.Now().UnixNano()))
        
        // Calculate the range for each component
        rangeX := int64(toX) - int64(v.X)
        rangeY := int64(toY) - int64(v.Y)
        rangeZ := int64(toZ) - int64(v.Z)
        
        // Generate random values within the range
        randX := int64(v.X)
        if rangeX > 0 {
                randX += r.Int63n(rangeX + 1)
        }
        
        randY := int64(v.Y)
        if rangeY > 0 {
                randY += r.Int63n(rangeY + 1)
        }
        
        randZ := int64(v.Z)
        if rangeZ > 0 {
                randZ += r.Int63n(rangeZ + 1)
        }
        
        return V3I[T]{
                X: T(randX),
                Y: T(randY),
                Z: T(randZ),
        }
}

// RandBetweenInPlace updates the vector with random components within a range specified by the vector and another vector.
// Each component of the vector will be set to a random value between its current value and the corresponding component of the provided vector.
func (v *V3I[T]) RandBetweenInPlace(to V3I[T]) {
        r := rand.New(rand.NewSource(time.Now().UnixNano()))
        
        // Calculate the range for each component
        rangeX := int64(to.X) - int64(v.X)
        rangeY := int64(to.Y) - int64(v.Y)
        rangeZ := int64(to.Z) - int64(v.Z)
        
        // Generate random values within the range and update the vector
        if rangeX > 0 {
                v.X = T(int64(v.X) + r.Int63n(rangeX+1))
        }
        
        if rangeY > 0 {
                v.Y = T(int64(v.Y) + r.Int63n(rangeY+1))
        }
        
        if rangeZ > 0 {
                v.Z = T(int64(v.Z) + r.Int63n(rangeZ+1))
        }
}

// RandBetweenCompInPlace updates the vector with random components within a range specified by the vector and the provided components.
// Each component of the vector will be set to a random value between its current value and the provided component.
func (v *V3I[T]) RandBetweenCompInPlace(toX, toY, toZ T) {
        r := rand.New(rand.NewSource(time.Now().UnixNano()))
        
        // Calculate the range for each component
        rangeX := int64(toX) - int64(v.X)
        rangeY := int64(toY) - int64(v.Y)
        rangeZ := int64(toZ) - int64(v.Z)
        
        // Generate random values within the range and update the vector
        if rangeX > 0 {
                v.X = T(int64(v.X) + r.Int63n(rangeX+1))
        }
        
        if rangeY > 0 {
                v.Y = T(int64(v.Y) + r.Int63n(rangeY+1))
        }
        
        if rangeZ > 0 {
                v.Z = T(int64(v.Z) + r.Int63n(rangeZ+1))
        }
}