package day4

import (
	"math"
	"regexp"
	"slices"
	"strconv"
)

type Card struct {
	Id             int
	Input          string
	WinningNumbers []int
	CardNumbers    []int
}

func NewCard(input string) *Card {
	return &Card{
		Input: input,
	}
}

func (c *Card) parseInput() error {
	// get card id
	re, err := regexp.Compile(`^Card\s*(\d+)\s*:(.*)\|(.*)$`)
	if err != nil {
		return err
	}

	match := re.FindStringSubmatch(c.Input)
	num, err := strconv.Atoi(match[1])
	if err != nil {
		return err
	}

	c.Id = num

	// get winning numbers
	re2, err := regexp.Compile(`\s+`)
	if err != nil {
		return err
	}

	arr := re2.Split(match[2], -1)
	nums := make([]int, 0)
	for _, v := range arr {
		if num, err := strconv.Atoi(v); err == nil {
			nums = append(nums, num)
		}
	}

	c.WinningNumbers = nums

	// get card numbers
	arr = re2.Split(match[3], -1)
	nums = make([]int, 0)
	for _, v := range arr {
		if num, err := strconv.Atoi(v); err == nil {
			nums = append(nums, num)
		}
	}

	c.CardNumbers = nums

	return nil
}

func (c *Card) Calculate() (int, error) {

	if err := c.parseInput(); err != nil {
		return 0, err
	}

	// Calculate winning points
	count := -1

	for _, num := range c.CardNumbers {
		if slices.Contains(c.WinningNumbers, num) {
			count++
		}
	}

	if count == -1 {
		return 0, nil
	}

	result := int(math.Pow(2, float64(count)))

	return result, nil
}

func (c *Card) NumberOfCardsWon() (int, error) {

	if err := c.parseInput(); err != nil {
		return 0, err
	}

	// Calculate winning points
	count := 0

	for _, num := range c.CardNumbers {
		if slices.Contains(c.WinningNumbers, num) {
			count++
		}
	}

	return count, nil
}
