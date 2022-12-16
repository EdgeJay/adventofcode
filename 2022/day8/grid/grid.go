package grid

import (
	"strconv"
	"strings"
)

type Grid [][]int

type SearchOptions struct {
	heightLimit int
	row         int
	col         int
	direction   string
}

const (
	DIR_UP    = "up"
	DIR_DOWN  = "down"
	DIR_LEFT  = "left"
	DIR_RIGHT = "right"
)

func NewGrid(rows int) *Grid {
	grid := make(Grid, rows)
	return &grid
}

func (g *Grid) GetTreeHeight(row, col int) int {
	return (*g)[row][col]
}

func (g *Grid) SetTreesHeightAtRow(rowNum int, input string) {
	row := make([]int, len(input))
	arr := strings.Split(input, "")
	for pos, n := range arr {
		num, _ := strconv.Atoi(n)
		row[pos] = num
	}
	(*g)[rowNum] = row
}

func (g *Grid) IsAboveOtherTrees(opts *SearchOptions) bool {
	treeRow := opts.row
	treeCol := opts.col
	switch opts.direction {
	case DIR_UP:
		treeRow -= 1
	case DIR_DOWN:
		treeRow += 1
	case DIR_LEFT:
		treeCol -= 1
	case DIR_RIGHT:
		treeCol += 1
	}

	if g.IsOutOfBounds(treeRow, treeCol) {
		return true
	}

	treeHeight := g.GetTreeHeight(treeRow, treeCol)
	if opts.heightLimit > treeHeight {
		opts.row = treeRow
		opts.col = treeCol
		return g.IsAboveOtherTrees(opts)
	}

	return false
}

func (g *Grid) IsTreeVisible(row, col int) bool {
	if row == 0 || col == 0 {
		return true
	}

	if row+1 >= len(*g) || col+1 >= len((*g)[0]) {
		return true
	}

	heightLimit := g.GetTreeHeight(row, col)
	results := g.IsAboveOtherTrees(&SearchOptions{
		heightLimit: heightLimit,
		row:         row,
		col:         col,
		direction:   DIR_UP,
	}) ||
		g.IsAboveOtherTrees(&SearchOptions{
			heightLimit: heightLimit,
			row:         row,
			col:         col,
			direction:   DIR_DOWN,
		}) ||
		g.IsAboveOtherTrees(&SearchOptions{
			heightLimit: heightLimit,
			row:         row,
			col:         col,
			direction:   DIR_LEFT,
		}) ||
		g.IsAboveOtherTrees(&SearchOptions{
			heightLimit: heightLimit,
			row:         row,
			col:         col,
			direction:   DIR_RIGHT,
		})

	return results
}

func (g *Grid) IsOutOfBounds(row, col int) bool {
	if row < 0 || col < 0 {
		return true
	}

	if row >= len(*g) || col >= len((*g)[0]) {
		return true
	}

	return false
}

func (g *Grid) GetRowsCount() int {
	return len(*g)
}

func (g *Grid) GetColumnsCount() int {
	return len((*g)[0])
}

func (g *Grid) GetTotalCount() int {
	return g.GetRowsCount() * g.GetColumnsCount()
}
