package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/EdgeJay/adventofcode/2022/day9/util"
	"github.com/EdgeJay/adventofcode/common/utils/files"
)

func parseSteps(lines []string) []util.Step {
	steps := make([]util.Step, 0)

	for _, ln := range lines {
		arr := strings.Split(ln, " ")
		if units, err := strconv.Atoi(arr[1]); err == nil {
			steps = append(steps, util.Step{Direction: arr[0], Units: units})
		}
	}

	return steps
}

func calculate(steps []util.Step) {
	head := util.NewPosition("head", 0, 0)
	tail := util.NewPosition("tail", 0, 0)
	tail.RecordVisit()

	for _, step := range steps {
		for n := 0; n < step.Units; n++ {
			head.Move(util.Step{Direction: step.Direction, Units: 1}, true)
			tail.MoveNextTo(head)
		}

		fmt.Printf("%s moved to (%d, %d)\n\n", head.Name, head.X, head.Y)
	}

	fmt.Println("Positions visited at least once:", len(tail.VisitedPositions))
}

func main() {
	lines := files.ReadInputsFile("./input.txt")
	// lines := files.ReadInputsFile("./test.txt")
	steps := parseSteps(lines)
	calculate(steps)
}
