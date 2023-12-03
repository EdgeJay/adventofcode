package day3

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnalyzerParse(t *testing.T) {

	lines := []string{
		"..............423....688..934............970................................95.728..........896...113..................153..972.............",
	}

	analyzer := NewAnalyzer(lines)
	analyzer.Parse()

	assert.Equal(t, 10, len(analyzer.EngineParts))
	assert.Equal(t, 14, analyzer.EngineParts[0].ColStart)
	assert.Equal(t, 17, analyzer.EngineParts[0].ColEnd)

	assert.Equal(t, "423", lines[0][analyzer.EngineParts[0].ColStart:analyzer.EngineParts[0].ColEnd])
	assert.Equal(t, "688", lines[0][analyzer.EngineParts[1].ColStart:analyzer.EngineParts[1].ColEnd])
}

func TestAnalyzerCalculate(t *testing.T) {

	input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

	lines := strings.Split(input, "\n")

	analyzer := NewAnalyzer(lines)
	analyzer.Parse()
	sum, err := analyzer.Calculate()
	assert.NoError(t, err)
	assert.Equal(t, 6, analyzer.TotalSymbols)
	assert.Equal(t, 4361, sum)
}

func TestEnginePartIsNearPosition(t *testing.T) {

	part := EnginePart{
		Row:      2,
		ColStart: 5,
		ColEnd:   7,
	}

	assert.Equal(t, true, part.IsNearPosition(2, 4))
	assert.Equal(t, true, part.IsNearPosition(2, 7))
	assert.Equal(t, false, part.IsNearPosition(2, 8))

	assert.Equal(t, true, part.IsNearPosition(3, 4))
	assert.Equal(t, true, part.IsNearPosition(3, 7))
	assert.Equal(t, false, part.IsNearPosition(3, 8))
}

func TestAnalyzerFindGearRatios(t *testing.T) {

	input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

	lines := strings.Split(input, "\n")

	analyzer := NewAnalyzer(lines)
	analyzer.Parse()

	sum, err := analyzer.FindGearRatios("*", 2)
	assert.NoError(t, err)
	assert.Equal(t, 467835, sum)
}
