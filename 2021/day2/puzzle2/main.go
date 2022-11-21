package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/EdgeJay/adventofcode/common"
	"github.com/EdgeJay/adventofcode/common/structs"
	"github.com/EdgeJay/adventofcode/common/utils/files"
	"github.com/EdgeJay/adventofcode/common/utils/input"
)

type instruction struct {
	direction string
	value     int
}

func calculate(instructions []instruction) {
	sub := structs.NewSubmarine()
	for _, in := range instructions {
		if in.direction == string(common.Forward) {
			sub.Horizontal += in.value
			sub.Depth += sub.Aim * in.value
		} else if in.direction == string(common.Down) {
			sub.Aim += in.value
		} else if in.direction == string(common.Up) {
			sub.Aim -= in.value
		}
	}
	fmt.Println(sub.Horizontal * sub.Depth)
}

func main() {
	arr := files.ReadInputsFile("./input.txt")
	processed := input.ProcessEachLine(arr, func(line string) instruction {
		parts := strings.Split(line, " ")
		in := instruction{}
		if len(parts) >= 2 {
			in.direction = parts[0]
			if val, err := strconv.Atoi(parts[1]); err == nil {
				in.value = val
			}
		}
		return in
	})
	calculate(processed)
}
