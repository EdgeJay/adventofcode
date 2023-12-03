package main

import (
	"fmt"

	"github.com/EdgeJay/adventofcode/2023/day2"
	"github.com/EdgeJay/adventofcode/common/utils/files"
)

func main() {
	lines := files.ReadInputsFile("../in/input.txt")

	sum := 0

	for _, ln := range lines {
		g := day2.NewGame()
		if err := g.SetParameters(ln); err == nil {
			sum += g.Power()
		}
	}

	fmt.Println(sum)
}
