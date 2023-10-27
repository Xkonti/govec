package govec

import (
	"math/rand"
	"time"

	"golang.org/x/exp/constraints"
)

// RandV2F returns a new random normalized V2F vector
func RandV2F[T constraints.Float]() V2F[T] {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return V2F[T]{T(r.Float32()), T(r.Float32())}.Norm()
}

// RandV3F returns a new random normalized V3F vector
func RandV3F[T constraints.Float]() V3F[T] {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return V3F[T]{T(r.Float32()), T(r.Float32()), T(r.Float32())}.Norm()
}
