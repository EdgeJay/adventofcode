package util

import (
	"strconv"
)

const (
	CMD_NOOP = "noop"
	CMD_ADDX = "addx"
)

type Command struct {
	Name  string
	Value int
}

func NewCommand(line string) *Command {
	if line[:4] == "noop" {
		return &Command{"noop", 0}
	}
	if line[:4] == "addx" {
		if val, err := strconv.Atoi(line[5:]); err == nil {
			return &Command{"addx", val}
		}
	}
	return nil
}
