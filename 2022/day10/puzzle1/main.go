package main

import (
	"fmt"

	"github.com/EdgeJay/adventofcode/2022/day10/util"
	"github.com/EdgeJay/adventofcode/common/utils/files"
)

func parseCommands(lines []string) []*util.Command {
	commands := make([]*util.Command, len(lines))

	for idx, ln := range lines {
		commands[idx] = util.NewCommand(ln)
	}

	return commands
}

func calculate(commands []*util.Command) {
	register := util.NewRegister("X", 1)

	for _, c := range commands {
		switch c.Name {
		case util.CMD_NOOP:
			register.Noop()
		case util.CMD_ADDX:
			register.Addx(c.Value)
		}
	}

	fmt.Println(register.SumOfValuesAtCycles(20, 60, 100, 140, 180, 220))
}

func main() {
	lines := files.ReadInputsFile("./input.txt")
	// lines := files.ReadInputsFile("./test2.txt")
	// lines := files.ReadInputsFile("./test.txt")
	commands := parseCommands(lines)
	calculate(commands)
}
