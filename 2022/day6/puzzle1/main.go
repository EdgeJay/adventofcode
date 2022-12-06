package main

import (
	"github.com/EdgeJay/adventofcode/2022/day6/util"
	"github.com/EdgeJay/adventofcode/common/utils/files"
)

const markerSize = 4

func calculate(lines []string) {
	for _, ln := range lines {
		util.ParseLineImproved(ln, markerSize)
	}
}

func main() {
	lines := files.ReadInputsFile("./input.txt")
	// lines := files.ReadInputsFile("./test.txt")
	calculate(lines)
}
