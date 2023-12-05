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

func (c *Card) Calculate() (int, error) {

	// get card id
	re, err := regexp.Compile(`^Card\s*(\d+)\s*:(.*)\|(.*)$`)
	if err != nil {
		return 0, err
	}

	match := re.FindStringSubmatch(c.Input)
	num, err := strconv.Atoi(match[1])
	if err != nil {
		return 0, err
	}

	c.Id = num

	// get winning numbers
	re2, err := regexp.Compile(`\s+`)
	if err != nil {
		return 0, err
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

	// Calculate winning points
	count := -1

	for _, num := range c.CardNumbers {
		if slices.Contains(c.WinningNumbers, num) {
			count++
		}
	}

	result := int(math.Pow(2, float64(count)))

	return result, nil
}
