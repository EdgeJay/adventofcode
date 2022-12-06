package main

import (
	"fmt"

	"github.com/EdgeJay/adventofcode/common/utils/files"
	"github.com/thoas/go-funk"
)

const markerSize = 14

func parseLine(line string) {
	t := []rune{}
	afterPos := 0

	for idx, c := range line {
		if len(t) == markerSize {
			afterPos = idx
			break
		}

		if pos := funk.IndexOfInt32(t, c); pos != -1 {
			if pos+1 < len(t) {
				t = t[pos+1:]
			} else {
				t = []rune{}
			}
		}

		t = append(t, c)
	}

	fmt.Println(string(t), afterPos)
}

func calculate(lines []string) {
	for _, ln := range lines {
		parseLine(ln)
	}
}

func main() {
	lines := files.ReadInputsFile("./input.txt")
	// lines := files.ReadInputsFile("./test.txt")
	calculate(lines)
}
