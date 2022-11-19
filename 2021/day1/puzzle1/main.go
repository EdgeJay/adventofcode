package main

import (
	"fmt"
	"strconv"

	"github.com/EdgeJay/adventofcode/2021/common/utils/files"
)

func calculate(arr []string) (count int) {
	var prevNum int = 99999

	for _, s := range arr {
		n, err := strconv.Atoi(s)
		if err == nil {
			if n > prevNum {
				count += 1
			}
			prevNum = n
		}
	}

	return
}

func main() {
	arr := files.ReadInputsFile("./input.txt")
	count := calculate(arr)
	fmt.Printf("%d\n", count)
}
