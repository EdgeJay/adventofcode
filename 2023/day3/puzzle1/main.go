package main

import (
	"fmt"
	"log"

	"github.com/EdgeJay/adventofcode/2023/day3"
	"github.com/EdgeJay/adventofcode/common/utils/files"
)

func main() {
	lines := files.ReadInputsFile("../in/input.txt")

	analyzer := day3.NewAnalyzer(lines)
	err := analyzer.Parse()
	if err != nil {
		log.Fatalln(err)
	}

	sum, err := analyzer.Calculate()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(sum)
}
