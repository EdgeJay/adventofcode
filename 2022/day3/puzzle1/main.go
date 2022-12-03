package main

import (
	"fmt"

	"github.com/EdgeJay/adventofcode/common/utils/files"
)

func getItemPriority(r rune) int {
	if r > 96 {
		return int(r) - 96
	} else {
		return int(r) - 38
	}
}

func calculate(arr []string) {
	sum := 0

	for _, ln := range arr {
		items := []rune(ln)
		m1 := make(map[rune]int) // track compartment 1
		m2 := make(map[rune]int) // track compartment 2

		// examine content of rucksack
		for idx, item := range items {
			if idx == len(items)/2 {
				break
			}

			m1[item] = m1[item] + 1
			m2[items[idx+(len(items)/2)]] = m2[items[idx+(len(items)/2)]] + 1
		}

		for key, val := range m1 {
			if val > 0 && m2[key] > 0 {
				sum += getItemPriority(key)
			}
		}
	}

	fmt.Println(sum)
}

func main() {
	arr := files.ReadInputsFile("./input.txt")
	calculate(arr)
}
