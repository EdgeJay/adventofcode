package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/EdgeJay/adventofcode/common/utils/data"
	"github.com/EdgeJay/adventofcode/common/utils/files"
)

type movement struct {
	numContainers int
	from          string
	to            string
}

func createStacks(lines []string) map[string]*data.Stack[string] {
	// last line contains stack number
	stacksNumLine := lines[len(lines)-1]
	stacksNumArr := strings.Split(strings.Trim(stacksNumLine, " "), "   ")
	totalStacks := len(stacksNumArr)

	// prepare map
	stacks := make(map[string]*data.Stack[string])
	for _, s := range stacksNumArr {
		stacks[s] = data.NewStack[string]()
	}

	for idx, ln := range lines {
		if idx < len(lines)-1 {
			r := []rune(ln)
			for i := 0; i < totalStacks; i++ {
				pos := i*4 + 1
				if r[pos] != ' ' {
					stack := stacks[strconv.Itoa(i+1)]
					stack.Add(string(r[pos]))
				}
			}
		}
	}

	for _, stack := range stacks {
		stack.ReverseItems()
	}

	return stacks
}

func createMovements(lines []string) []movement {
	movements := []movement{}
	re, _ := regexp.Compile(`^move\s(\d+)\sfrom\s(\d+)\sto\s(\d+)$`)
	for _, ln := range lines {
		arr := re.FindStringSubmatch(ln)[1:]
		num, _ := strconv.Atoi(arr[0])
		movements = append(movements, movement{num, arr[1], arr[2]})
	}
	return movements
}

func parseLines(lines []string) (stacks map[string]*data.Stack[string], movements []movement) {
	stacksRead := false
	var part1 []string
	var part2 []string

	for _, ln := range lines {
		if !stacksRead {
			if len(ln) > 0 {
				part1 = append(part1, ln)
			} else {
				stacksRead = true
			}
		} else {
			part2 = append(part2, ln)
		}
	}

	stacks = createStacks(part1)
	movements = createMovements(part2)
	return
}

func calculate(stacks map[string]*data.Stack[string], movements []movement) {
	for _, movement := range movements {
		removed := stacks[movement.from].RemoveItems(movement.numContainers, true)
		stacks[movement.to].AddItems(removed)
	}

	var b bytes.Buffer

	for idx := 1; idx <= len(stacks); idx++ {
		b.WriteString(stacks[strconv.Itoa(idx)].LastItem())
	}

	fmt.Println(b.String())
}

func main() {
	lines := files.ReadInputsFile("./input.txt")
	// lines := files.ReadInputsFile("./test.txt")
	parseLines(lines)
	stacks, movements := parseLines(lines)
	calculate(stacks, movements)
}
