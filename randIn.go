package govec

import (
        "math/rand"
        "time"
)

// RandIn generates a vector with random components within a range specified by the vector and a zero vector.
// Each component of the returned vector will be a random value between 0 and the corresponding component of the vector.
func (v V2F[T]) RandIn() V2F[T] {
        r := rand.New(rand.NewSource(time.Now().UnixNano()))
        return V2F[T]{
                X: T(r.Float64()) * v.X,
                Y: T(r.Float64()) * v.Y,
        }
}

// RandInInPlace updates the vector with random components within a range specified by the vector and a zero vector.
// Each component of the vector will be set to a random value between 0 and its current value.
func (v *V2F[T]) RandInInPlace() {
        r := rand.New(rand.NewSource(time.Now().UnixNano()))
        v.X = T(r.Float64()) * v.X
        v.Y = T(r.Float64()) * v.Y
}

// RandIn generates a vector with random components within a range specified by the vector and a zero vector.
// Each component of the returned vector will be a random value between 0 and the corresponding component of the vector.
func (v V3F[T]) RandIn() V3F[T] {
        r := rand.New(rand.NewSource(time.Now().UnixNano()))
        return V3F[T]{
                X: T(r.Float64()) * v.X,
                Y: T(r.Float64()) * v.Y,
                Z: T(r.Float64()) * v.Z,
        }
}

// RandInInPlace updates the vector with random components within a range specified by the vector and a zero vector.
// Each component of the vector will be set to a random value between 0 and its current value.
func (v *V3F[T]) RandInInPlace() {
        r := rand.New(rand.NewSource(time.Now().UnixNano()))
        v.X = T(r.Float64()) * v.X
        v.Y = T(r.Float64()) * v.Y
        v.Z = T(r.Float64()) * v.Z
}

// RandIn generates a vector with random components within a range specified by the vector and a zero vector.
// Each component of the returned vector will be a random value between 0 and the corresponding component of the vector.
func (v V2I[T]) RandIn() V2I[T] {
        r := rand.New(rand.NewSource(time.Now().UnixNano()))
        return V2I[T]{
                X: T(r.Int63n(int64(v.X) + 1)),
                Y: T(r.Int63n(int64(v.Y) + 1)),
        }
}

// RandInInPlace updates the vector with random components within a range specified by the vector and a zero vector.
// Each component of the vector will be set to a random value between 0 and its current value.
func (v *V2I[T]) RandInInPlace() {
        r := rand.New(rand.NewSource(time.Now().UnixNano()))
        v.X = T(r.Int63n(int64(v.X) + 1))
        v.Y = T(r.Int63n(int64(v.Y) + 1))
}

// RandIn generates a vector with random components within a range specified by the vector and a zero vector.
// Each component of the returned vector will be a random value between 0 and the corresponding component of the vector.
func (v V3I[T]) RandIn() V3I[T] {
        r := rand.New(rand.NewSource(time.Now().UnixNano()))
        return V3I[T]{
                X: T(r.Int63n(int64(v.X) + 1)),
                Y: T(r.Int63n(int64(v.Y) + 1)),
                Z: T(r.Int63n(int64(v.Z) + 1)),
        }
}

// RandInInPlace updates the vector with random components within a range specified by the vector and a zero vector.
// Each component of the vector will be set to a random value between 0 and its current value.
func (v *V3I[T]) RandInInPlace() {
        r := rand.New(rand.NewSource(time.Now().UnixNano()))
        v.X = T(r.Int63n(int64(v.X) + 1))
        v.Y = T(r.Int63n(int64(v.Y) + 1))
        v.Z = T(r.Int63n(int64(v.Z) + 1))
}