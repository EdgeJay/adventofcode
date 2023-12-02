package main

import (
	"fmt"

	"github.com/EdgeJay/adventofcode/2023/day2"
	"github.com/EdgeJay/adventofcode/common/utils/files"
)

const (
	redLimit   = 12
	greenLimit = 13
	blueLimit  = 14
)

func main() {
	lines := files.ReadInputsFile("../in/input.txt")

	sum := 0

	for _, ln := range lines {
		g := day2.NewGame()
		if err := g.SetParameters(ln); err == nil {
			if g.WithinLimits(redLimit, greenLimit, blueLimit) {
				sum += g.Id
			}
		}
	}

	fmt.Println(sum)
}
