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

	// get nodes with certain elevation level only
	steps := 999
	nodes := grid.GetNodesWithElevationAt(1)
	for _, node := range nodes {
		newSteps := grid.FindShortestPath(node, grid.EndLocation)
		if newSteps != 0 && newSteps < steps {
			steps = newSteps
		}
	}
	fmt.Println(steps)
}
