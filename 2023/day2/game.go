package day2

import (
	"fmt"
	"regexp"
	"strconv"
)

type Game struct {
	Id    int
	Red   int
	Green int
	Blue  int
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) SetParameters(input string) error {
	// Find Game ID
	re, err := regexp.Compile(`^Game (\d+):`)
	if err != nil {
		return err
	}

	arr := re.FindStringSubmatch(input)
	if len(arr) != 2 {
		return fmt.Errorf("game ID match cannot be found for: %s", input)
	}

	id, err := strconv.Atoi(arr[1])
	if err != nil {
		return err
	}

	g.Id = id

	// Find no. of reds, greens, blues
	colours := []string{
		"red",
		"green",
		"blue",
	}

	for _, c := range colours {

		re, err := regexp.Compile(fmt.Sprintf(`(\d+) %s`, c))
		if err != nil {
			return err
		}

		max := 0
		arr2 := re.FindAllStringSubmatch(input, -1)
		for _, arr := range arr2 {
			if num, err := strconv.Atoi(arr[1]); err == nil {
				if num > max {
					max = num
				}
			}
		}

		switch c {
		case "red":
			g.Red = max
		case "green":
			g.Green = max
		case "blue":
			g.Blue = max
		}
	}

	return nil
}

func (g *Game) WithinLimits(red, green, blue int) bool {
	return g.Red <= red && g.Green <= green && g.Blue <= blue
}

func (g *Game) Power() int {
	return g.Red * g.Green * g.Blue
}
