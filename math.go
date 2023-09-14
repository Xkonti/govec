package govec

import "golang.org/x/exp/constraints"

// clamp clamps a value between min and max
func clamp[T constraints.Float | constraints.Integer](val, min, max T) T {
	if val < min {
		return min
	} else if val > max {
		return max
	}
	return val
}

// pow computes x^n for integer n
func pow[T constraints.Float | constraints.Integer](x T, n int16) T {
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}

	result := T(1.0)
	for n > 0 {
		if n%2 == 1 {
			result *= x
		}
		x *= x
		n /= 2
	}

	return result
}
