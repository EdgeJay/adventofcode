package main

import (
	"fmt"

	"github.com/EdgeJay/adventofcode/2023/day5"
	"github.com/EdgeJay/adventofcode/common/utils/files"
)

func main() {

	lines := files.ReadInputsFileRaw("../in/input.txt")

	almanac, err := day5.NewAlmanac(lines)
	fmt.Println(almanac, err)
}
