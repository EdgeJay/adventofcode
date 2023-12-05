package main

import (
	"fmt"

	"github.com/EdgeJay/adventofcode/2023/day4"
	"github.com/EdgeJay/adventofcode/common/utils/files"
)

func main() {

	lines := files.ReadInputsFile("../in/input.txt")
	cardPile := day4.NewCardPile(lines)
	result := cardPile.TotalWinningCards()

	fmt.Println(result)
}
