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

func prevShape(s int) int {
	if s-1 > 0 {
		return s - 1
	} else {
		return 3
	}
}

func nextShape(s int) int {
	if s+1 <= 3 {
		return s + 1
	} else {
		return 1
	}
}

func calculate(arr []string) {
	total := 0

	for _, ln := range arr {
		chars := []rune(ln)
		// shape chosen by opponent
		c1 := int(chars[0]-'A') + 1
		// 'X' or 0 = lose, 'Y' or 3 = draw, 'Z' or 6 = 'win'
		c2 := int(chars[2]-'X') * 3

		// Points received for selecting shape
		// Rock = 1
		// Paper = 2
		// Scissors = 3

		if c2 == LosePts {
			total += prevShape(c1) + LosePts
		} else if c2 == DrawPts {
			total += c1 + DrawPts
		} else if c2 == WinPts {
			total += nextShape(c1) + WinPts
		}
	}

	fmt.Println(total)
}

func main() {
	arr := files.ReadInputsFile("./input.txt")
	calculate(arr)
}
