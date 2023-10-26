package govec

import (
	"math"
)

const radToDegFactor float64 = 180 / math.Pi

const degToRadFactor float64 = math.Pi / 180

type axis int

const (
	xAxis axis = iota
	yAxis
	zAxis
)
