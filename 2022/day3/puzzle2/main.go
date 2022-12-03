package main

import (
	"fmt"

	"github.com/EdgeJay/adventofcode/common/utils/files"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getItemPriority(r rune) int {
	if r > 96 {
		return int(r) - 96
	} else {
		return int(r) - 38
	}
}

func calculate(arr []string) {
	sum := 0

	for lineNum := 0; lineNum < len(arr); lineNum += 3 {
		ruckSack1 := []rune(arr[lineNum])
		ruckSack2 := []rune(arr[lineNum+1])
		ruckSack3 := []rune(arr[lineNum+2])
		commonMap := make(map[rune]map[string]int)

		// get shortest rucksack length
		limit := max(max(len(ruckSack1), len(ruckSack2)), len(ruckSack3))

		// examine content of rucksack
		for idx := 0; idx < limit; idx++ {
			if idx < len(ruckSack1) {
				if commonMap[ruckSack1[idx]] == nil {
					commonMap[ruckSack1[idx]] = make(map[string]int)
				}
				commonMap[ruckSack1[idx]]["r1"] = commonMap[ruckSack1[idx]]["r1"] + 1
			}

			if idx < len(ruckSack2) {
				if commonMap[ruckSack2[idx]] == nil {
					commonMap[ruckSack2[idx]] = make(map[string]int)
				}
				commonMap[ruckSack2[idx]]["r2"] = commonMap[ruckSack2[idx]]["r2"] + 1
			}

			if idx < len(ruckSack3) {
				if commonMap[ruckSack3[idx]] == nil {
					commonMap[ruckSack3[idx]] = make(map[string]int)
				}
				commonMap[ruckSack3[idx]]["r3"] = commonMap[ruckSack3[idx]]["r3"] + 1
			}
		}

		for key, value := range commonMap {
			if value["r1"] > 0 && value["r2"] > 0 && value["r3"] > 0 {
				sum += getItemPriority(key)
				break
			}
		}
	}

	fmt.Println(sum)
}

func main() {
	arr := files.ReadInputsFile("./input.txt")
	calculate(arr)
}
