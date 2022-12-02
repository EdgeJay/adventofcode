package main

import (
	"fmt"

	"github.com/EdgeJay/adventofcode/common/utils/files"
)

const (
	LosePts = 0
	DrawPts = 3
	WinPts  = 6
)

func calculate(arr []string) {
	total := 0

	for _, ln := range arr {
		chars := []rune(ln)
		c1 := chars[0]
		c2 := chars[2] - 23
		// Points received for selecting shape
		// Rock = 1
		// Paper = 2
		// Scissors = 3
		shapePts := int(chars[2] - 87)

		if c1-1 == c2 || c2-2 == c1 {
			total += shapePts + LosePts
		} else if c2-1 == c1 || c1-2 == c2 {
			total += shapePts + WinPts
		} else if c1 == c2 {
			total += shapePts + DrawPts
		}
	}

	fmt.Println(total)
}

func main() {
	arr := files.ReadInputsFile("./input.txt")
	calculate(arr)
}
