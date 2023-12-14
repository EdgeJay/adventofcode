package day5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAlmanac(t *testing.T) {

	input := `seeds: 1848591090 462385043 2611025720 154883670 1508373603 11536371 3692308424 16905163 1203540561 280364121 3755585679 337861951 93589727 738327409 3421539474 257441906 3119409201 243224070 50985980 7961058

seed-to-soil map:
3305253869 1699909104 39566623
3344820492 1130725752 384459310
3244681427 1739475727 60572442
951517531 1800048169 868898709
1820416240 951517531 179208221
1999624461 2668946878 219310925
3729279802 1515185062 184724042
2218935386 2898481077 1015522767
3234458153 2888257803 10223274`

	almanac, err := NewAlmanac(input)
	assert.NoError(t, err)
	assert.Equal(t, 20, len(almanac.SeedMaps))
}

func TestSrcDstMapIsWithinRange(t *testing.T) {

	input := `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

	almanac, err := NewAlmanac(input)
	assert.NoError(t, err)
	assert.Equal(t, false, almanac.SrcDstMapList[0].IsWithinSourceRange(97))
	assert.Equal(t, true, almanac.SrcDstMapList[0].IsWithinSourceRange(98))
	assert.Equal(t, true, almanac.SrcDstMapList[0].IsWithinSourceRange(99))
	assert.Equal(t, false, almanac.SrcDstMapList[0].IsWithinSourceRange(100))
	assert.Equal(t, false, almanac.SrcDstMapList[0].IsWithinDestinationRange(49))
	assert.Equal(t, true, almanac.SrcDstMapList[0].IsWithinDestinationRange(50))
	assert.Equal(t, true, almanac.SrcDstMapList[0].IsWithinDestinationRange(51))
	assert.Equal(t, false, almanac.SrcDstMapList[0].IsWithinDestinationRange(52))
}
