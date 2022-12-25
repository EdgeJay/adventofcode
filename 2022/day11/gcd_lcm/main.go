package main

import (
	"fmt"

	"github.com/EdgeJay/adventofcode/common/math"
)

func gcd(a, b int) int {
	fmt.Printf("gcd(%d, %d)\n", a, b)

	if b == 0 {
		return a
	}

	return gcd(b, math.ModInt(a, b))
}

func lcm(a, b, gcd int) int {
	return (a * b) / gcd
}

func main() {
	a := 15
	b := 20
	result := gcd(a, b)
	fmt.Printf("GCD(%d, %d) = %d\n", a, b, result)

	// https://www.geeksforgeeks.org/program-to-find-lcm-of-two-numbers/
	// a x b = LCM(a, b) * GCD(a, b)
	// LCM(a, b) = (a x b) / GCD(a, b)
	result = lcm(a, b, result)
	fmt.Printf("LCM(%d, %d) = %d\n", a, b, result)
}
