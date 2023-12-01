package main

import (
	"fmt"
	"strings"

	"github.com/EdgeJay/adventofcode/2022/day13/util"
	"github.com/EdgeJay/adventofcode/common/utils/files"
)

func calculate(input []string) {
	for _, ln := range input {
		fmt.Println("checking next pair of packets...")
		arr := strings.Split(ln, "\n")
		comparator := util.NewComparator(arr[0], arr[1])
		inRightOrder := comparator.Check()
		fmt.Println(inRightOrder)
	}
}

func main() {
	// lines := files.ReadInputsFileRaw("./input.txt")
	// lines := files.ReadInputsFileRaw("./test.txt")
	// lines := files.ReadInputsFileRaw("./test2.txt")
	lines := files.ReadInputsFileRaw("./test3.txt")
	input := strings.Split(lines, "\n\n")
	calculate(input)
}
