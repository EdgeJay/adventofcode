package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/EdgeJay/adventofcode/common/utils/files"
)

func makeSectionsArray(sections string) []int {
	limits := strings.Split(sections, "-")
	lower, _ := strconv.Atoi(limits[0])
	upper, _ := strconv.Atoi(limits[1])
	arr := make([]int, upper-lower+1)
	for idx := 0; idx < len(arr); idx++ {
		arr[idx] = lower + idx
	}
	return arr
}

func makeSectionSlices(ln string) (sections1, sections2 []int) {
	arr := strings.Split(ln, ",")
	sections1 = makeSectionsArray(arr[0])
	sections2 = makeSectionsArray(arr[1])
	return
}

func calculate(arr []string) {
	sum := 0

	for _, ln := range arr {
		sections1, sections2 := makeSectionSlices(ln)

		if len(sections1) == len(sections2) && sections2[0] == sections1[0] && sections2[len(sections2)-1] == sections1[len(sections1)-1] {
			fmt.Println(ln)
			sum++
		} else if sections1[0] >= sections2[0] && sections1[0] <= sections2[len(sections2)-1] {
			fmt.Println(ln)
			sum++
		} else if sections2[0] >= sections1[0] && sections2[0] <= sections1[len(sections1)-1] {
			/**
			 * ..3456789  3-7 (sections 1)
			 * 12345678.  1-8 (sections 2)
			 *
			 * ...456...  4-6
			 * .....6...  6-6
			 */
			fmt.Println(ln)
			sum++
		}
	}

	fmt.Println(sum)
}

func main() {
	arr := files.ReadInputsFile("./input.txt")
	// arr := files.ReadInputsFile("./test.txt")
	calculate(arr)
}
