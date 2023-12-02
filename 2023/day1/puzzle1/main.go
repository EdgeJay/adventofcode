package main

import (
	"fmt"
	"slices"
	"strconv"
	"unicode"

	"github.com/EdgeJay/adventofcode/common/utils/files"
)

func main() {

	lines := files.ReadInputsFile("../input.txt")
	sum := 0

	for _, ln := range lines {

		ln1 := []rune(ln)
		ln2 := slices.Clone(ln1)
		slices.Reverse(ln2)

		digits := []rune{' ', ' '}

		for idx := 0; idx < len(ln1); idx++ {

			if digits[0] == ' ' && unicode.IsDigit(ln1[idx]) {
				digits[0] = ln1[idx]
			}

			if digits[1] == ' ' && unicode.IsDigit(ln2[idx]) {
				digits[1] = ln2[idx]
			}

			if digits[0] != ' ' && digits[1] != ' ' {
				break
			}
		}

		str := string(digits)
		if num, err := strconv.Atoi(str); err == nil {
			sum += num
		}
	}

	fmt.Println(sum)
}
