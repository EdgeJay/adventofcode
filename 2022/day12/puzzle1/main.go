package main

import (
	"fmt"

	"github.com/EdgeJay/adventofcode/2022/day12/util"
	"github.com/EdgeJay/adventofcode/common/utils/files"
)

func createGrid(lines []string) *util.Grid {
	return util.MakeGrid(lines)
}

func main() {
	lines := files.ReadInputsFile("./input.txt")
	// lines := files.ReadInputsFile("./test.txt")
	grid := createGrid(lines)
	steps := grid.GetShortestPathSteps()
	fmt.Println("Shortest path steps: ", steps)
}
