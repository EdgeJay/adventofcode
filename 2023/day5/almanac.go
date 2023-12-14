package day5

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

type ItemName string

const SeedItem ItemName = "seed"

const SoilItem ItemName = "soil"

const FertilizerItem ItemName = "fertilizer"

const WaterItem ItemName = "water"

const LightItem ItemName = "light"

const TemperatureItem ItemName = "temperature"

const HumidityItem ItemName = "humidity"

const LocationItem ItemName = "location"

type SrcDstMap struct {
	Source           ItemName
	Destination      ItemName
	SourceStart      int64
	DestinationStart int64
	Range            int64
}

func (m *SrcDstMap) IsWithinSourceRange(v int64) bool {
	return v >= m.SourceStart && v <= m.SourceStart+m.Range-1
}

func (m *SrcDstMap) IsWithinDestinationRange(v int64) bool {
	return v >= m.DestinationStart && v <= m.DestinationStart+m.Range-1
}

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
	input         string
	SeedMaps      []*SeedMap
	SrcDstMapList []*SrcDstMap
}

func NewAlmanac(input string) (*Almanac, error) {

	re, err := regexp.Compile(`^(?:seeds:\s*)(.*)(?:\n\n.*)`)
	if err != nil {
		return nil, err
	}

	// find seeds
	seedMaps := make([]*SeedMap, 0)
	arr := re.FindStringSubmatch(input)
	arr2 := strings.Split(arr[1], " ")
	for _, s := range arr2 {
		if v, err := strconv.Atoi(s); err == nil {
			seedMap := &SeedMap{
				Seed: int64(v),
			}
			seedMaps = append(seedMaps, seedMap)
		}
	}

	// find src-to-dst maps
	list := make([]*SrcDstMap, 0)
	arr = strings.Split(input, "\n\n")
	for _, str := range arr[1:] {

		arr2 = strings.Split(str, "\n")
		re, err := regexp.Compile(`(.*)-to-(.*)\s+map:`)
		if err != nil {
			return nil, err
		}

		arr3 := re.FindStringSubmatch(arr2[0])

		for _, ln := range arr2[1:] {

			srcDstMap := &SrcDstMap{
				Source:      ItemName(arr3[1]),
				Destination: ItemName(arr3[2]),
			}

			arr4 := strings.Split(ln, " ")
			if n, err := strconv.Atoi(arr4[0]); err == nil {
				srcDstMap.DestinationStart = int64(n)
			}
			if n, err := strconv.Atoi(arr4[1]); err == nil {
				srcDstMap.SourceStart = int64(n)
			}
			if n, err := strconv.Atoi(arr4[2]); err == nil {
				srcDstMap.Range = int64(n)
			}

			list = append(list, srcDstMap)
		}
	}

	return &Almanac{
		input:         input,
		SeedMaps:      seedMaps,
		SrcDstMapList: list,
	}, nil
}

func (a *Almanac) SrcDstMapListBySource(source ItemName) []*SrcDstMap {

	list := make([]*SrcDstMap, 0)
	for _, seedMap := range a.SrcDstMapList {
		if seedMap.Source == source {
			list = append(list, seedMap)
		}
	}

	return list
}

func (a *Almanac) FindLowestLocationNumber() int64 {

	var location int64 = math.MaxInt64

	/*
	for _, seedMap := range a.SeedMaps {

		// seed to soil
		list := a.SrcDstMapListBySource(SoilItem)
		for _, srcDstMap := range list {

		}
	}
	*/
	
	return location
}
