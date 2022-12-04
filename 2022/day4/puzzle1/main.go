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

		var outer []int = nil
		var inner []int = nil

		if len(sections2) > len(sections1) {
			outer = sections2
			inner = sections1
		} else if len(sections1) > len(sections2) {
			outer = sections1
			inner = sections2
		}

		if outer != nil && inner != nil {
			if outer[0] <= inner[0] && outer[len(outer)-1] >= inner[len(inner)-1] {
				sum++
			}
		} else {
			// no. of sections for pair at the same
			if sections1[0] == sections2[0] && sections1[len(sections1)-1] == sections2[len(sections2)-1] {
				sum++
			}
		}
	}

	fmt.Println(sum)
}

func main() {
	arr := files.ReadInputsFile("./input.txt")
	// arr := files.ReadInputsFile("./test.txt")
	calculate(arr)
}
