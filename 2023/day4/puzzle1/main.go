package main

import (
	"fmt"

	"github.com/EdgeJay/adventofcode/2023/day4"
	"github.com/EdgeJay/adventofcode/common/utils/files"
)

func main() {

	sum := 0
	lines := files.ReadInputsFile("../in/input.txt")

	for _, line := range lines {
		c := day4.NewCard(line)
		if result, err := c.Calculate(); err == nil {
			sum += result
		}
	}

	fmt.Println(sum)
}
