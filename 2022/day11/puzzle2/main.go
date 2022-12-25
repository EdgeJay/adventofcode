package main

import (
	"fmt"
	"strings"

	"github.com/EdgeJay/adventofcode/2022/day11/util"
	"github.com/EdgeJay/adventofcode/common/utils/files"
)

func parseNotes(lines string) *util.Observation {
	monkeys := strings.Split(lines, "\n\n")
	ob := util.NewObservation(monkeys, 0)
	return ob
}

func calculate(ob *util.Observation) {
	for n := 0; n < 10000; n++ {
		ob.NextRound()
	}

	fmt.Println(ob.GetMonkeybyID(0).ObservationCount)
	fmt.Println(ob.GetMonkeybyID(3).ObservationCount)

	monkeys := ob.GetMonkeysSortedByObservationCount()
	total := monkeys[0].ObservationCount * monkeys[1].ObservationCount
	fmt.Println(total)
}

func main() {
	// lines := files.ReadInputsFileRaw("./input.txt")
	lines := files.ReadInputsFileRaw("./test.txt")
	observation := parseNotes(lines)
	calculate(observation)
}
