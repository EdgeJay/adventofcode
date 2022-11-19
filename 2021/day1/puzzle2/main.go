package main

import (
	"fmt"
	"strconv"

	"github.com/EdgeJay/adventofcode/common/utils/files"
)

/**
 * This puzzle requires the understanding of the "Window Sliding Technique":
 * https://www.geeksforgeeks.org/window-sliding-technique/
 */

func calculate(arr []string, k int) (count int) {
	tempSum := 0

	// Get sum of 1st k elements
	for _, s := range arr[0:k] {
		n, err := strconv.Atoi(s)
		if err == nil {
			tempSum += n
		}
	}

	for index, s := range arr {
		if index+k >= len(arr) {
			break
		}

		if prevNum, err := strconv.Atoi(s); err == nil {
			if nextNum, err := strconv.Atoi(arr[index+k]); err == nil {
				curSum := tempSum - prevNum + nextNum
				if curSum > tempSum {
					tempSum = curSum
					count += 1
				}
			}
		}
	}

	return
}

func main() {
	arr := files.ReadInputsFile("./input.txt")
	count := calculate(arr, 3)
	fmt.Printf("%d\n", count)
}
