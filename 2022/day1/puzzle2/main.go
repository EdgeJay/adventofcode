package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/EdgeJay/adventofcode/common/utils/files"
)

func calculate(arr []string) {
	sum := 0
	var calories []int

	for _, item := range arr {
		if cal, err := strconv.Atoi(item); err == nil {
			sum += cal
		} else if len(item) < 1 {
			calories = append(calories, sum)
			sum = 0
		}
	}

	calories = append(calories, sum)
	sort.Slice(calories, func(i, j int) bool {
		return calories[i] > calories[j]
	})
	fmt.Println(calories[0] + calories[1] + calories[2])
}

func main() {
	arr := files.ReadInputsFile("./input.txt")
	calculate(arr)
}
