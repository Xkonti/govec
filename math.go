package govec

import "golang.org/x/exp/constraints"

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
