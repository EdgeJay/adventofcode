package main

import (
	"fmt"

	"github.com/EdgeJay/adventofcode/2022/day8/grid"
	"github.com/EdgeJay/adventofcode/common/utils/files"
)

type treePosition struct {
	row int
	col int
}

func parseGrid(lines []string) *grid.Grid {
	grid := grid.NewGrid(len(lines))
	for idx, ln := range lines {
		grid.SetTreesHeightAtRow(idx, ln)
	}
	return grid
}

func calculate(grid *grid.Grid) {
	visibleTrees := make([]treePosition, 0)

	offset := 0
	row := offset
	col := offset

	for {
		if col >= grid.GetColumnsCount()-offset {
			col = offset
			row++
		}

		if row >= grid.GetRowsCount()-offset {
			break
		}

		if grid.IsTreeVisible(row, col) {
			visibleTrees = append(visibleTrees, treePosition{row, col})
		}

		col++
	}

	fmt.Println(visibleTrees, len(visibleTrees))
}

func main() {
	lines := files.ReadInputsFile("./input.txt")
	// lines := files.ReadInputsFile("./test.txt")
	grid := parseGrid(lines)
	calculate(grid)
}
