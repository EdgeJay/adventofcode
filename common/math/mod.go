package math

import "math"

func ModInt(x, y int) int {
	return int(math.Mod(float64(x), float64(y)))
}
