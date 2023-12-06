package day5

import (
	"regexp"
	"strconv"
	"strings"
)

type SeedMap struct {
	Seed        int64
	Soil        int64
	Fertilizer  int64
	Water       int64
	Light       int64
	Temperature int64
	Humidity    int64
	Location    int64
}

type Almanac struct {
	input    string
	Seeds    []int64
	SeedMaps []*SeedMap
}

func NewAlmanac(input string) (*Almanac, error) {

	re, err := regexp.Compile(`^(?:seeds:\s*)(.*)(?:\n\n.*)`)
	if err != nil {
		return nil, err
	}

	// find seeds
	seeds := make([]int64, 0)
	arr := re.FindStringSubmatch(input)
	arr2 := strings.Split(arr[1], " ")
	for _, s := range arr2 {
		if v, err := strconv.Atoi(s); err == nil {
			seeds = append(seeds, int64(v))
		}
	}

	return &Almanac{
		input: input,
		Seeds: seeds,
	}, nil
}
