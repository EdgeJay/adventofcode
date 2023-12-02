package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"unicode"

	"github.com/EdgeJay/adventofcode/common/utils/files"
)

func detectAndReplace(str *string) bool {

	if strings.Contains(*str, "one") {
		*str = strings.Replace(*str, "one", "1", 1)
		return true
	}

	if strings.Contains(*str, "two") {
		*str = strings.Replace(*str, "two", "2", 1)
		return true
	}

	if strings.Contains(*str, "three") {
		*str = strings.Replace(*str, "three", "3", 1)
		return true
	}

	if strings.Contains(*str, "four") {
		*str = strings.Replace(*str, "four", "4", 1)
		return true
	}

	if strings.Contains(*str, "five") {
		*str = strings.Replace(*str, "five", "5", 1)
		return true
	}

	if strings.Contains(*str, "six") {
		*str = strings.Replace(*str, "six", "6", 1)
		return true
	}

	if strings.Contains(*str, "seven") {
		*str = strings.Replace(*str, "seven", "7", 1)
		return true
	}

	if strings.Contains(*str, "eight") {
		*str = strings.Replace(*str, "eight", "8", 1)
		return true
	}

	if strings.Contains(*str, "nine") {
		*str = strings.Replace(*str, "nine", "9", 1)
		return true
	}

	return false
}

func main() {

	lines := files.ReadInputsFile("../input.txt")
	sum := 0

	for _, ln := range lines {

		left := ""
		arr := make([]rune, 0)
		for _, c := range ln {
			arr = append(arr, c)
			str := string(arr)
			if detectAndReplace(&str) {
				left = str
				break
			}
		}

		right := ""
		arr = make([]rune, 0)
		r := []rune(ln)
		for idx := len(r) - 1; idx >= 0; idx-- {
			arr = append([]rune{r[idx]}, arr...)
			str := string(arr)
			if detectAndReplace(&str) {
				right = str
				break
			}
		}

		fmt.Println(left, right)
		tmpStr := fmt.Sprintf("%s%s%s", left, ln[len(left):len(ln)-len(right)], right)

		ln1 := []rune(tmpStr)
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
		fmt.Println(ln, tmpStr, str)

		if num, err := strconv.Atoi(str); err == nil {
			sum += num
		}
	}

	fmt.Println(sum)
}
