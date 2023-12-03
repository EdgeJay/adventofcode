package day3

import (
	"fmt"
	"regexp"
	"strconv"
)

type EnginePart struct {
	Number   int
	Row      int
	ColStart int
	ColEnd   int
}

func (p *EnginePart) IsNearPosition(row, col int) bool {

	if row == p.Row {
		return p.ColEnd == col || p.ColStart-1 == col
	}

	if row == p.Row-1 || row == p.Row+1 {
		return col >= p.ColStart-1 && col <= p.ColEnd
	}

	return false
}

type Analyzer struct {
	Lines        []string
	EngineParts  []EnginePart
	TotalSymbols int
}

func NewAnalyzer(lines []string) *Analyzer {
	return &Analyzer{
		Lines: lines,
	}
}

func (a *Analyzer) Parse() error {

	re, err := regexp.Compile(`\W*(\d+)\W*`)
	if err != nil {
		return err
	}

	parts := make([]EnginePart, 0)

	for row, ln := range a.Lines {

		matches := re.FindAllStringSubmatch(ln, -1)
		matchIndices := re.FindAllStringSubmatchIndex(ln, -1)

		for idx, match := range matches {

			num, err := strconv.Atoi(match[1])

			if err == nil {
				part := EnginePart{
					Number:   num,
					Row:      row,
					ColStart: matchIndices[idx][2],
					ColEnd:   matchIndices[idx][3],
				}
				parts = append(parts, part)
			}
		}
	}

	a.EngineParts = parts

	return nil
}

func (a *Analyzer) Calculate() (int, error) {

	a.TotalSymbols = 0

	totalSymbols := 0

	sum := 0

	re, err := regexp.Compile(`[^0-9.]`)
	if err != nil {
		return 0, err
	}

	for row, ln := range a.Lines {

		matchIndices := re.FindAllStringSubmatchIndex(ln, -1)
		for _, match := range matchIndices {

			col := match[0]
			parts := a.EnginePartsNearPosition(row, col)

			for _, p := range parts {
				sum += p.Number
			}

			totalSymbols++
		}
	}

	a.TotalSymbols = totalSymbols

	return sum, nil
}

func (a *Analyzer) FindGearRatios(gearSymbol string, numAdj int) (int, error) {

	sum := 0

	re, err := regexp.Compile(fmt.Sprintf(`[%s]`, gearSymbol))
	if err != nil {
		return 0, err
	}

	for row, ln := range a.Lines {

		matchIndices := re.FindAllStringSubmatchIndex(ln, -1)
		for _, match := range matchIndices {

			col := match[0]
			parts := a.EnginePartsNearPosition(row, col)

			if len(parts) == numAdj {
				sum += parts[0].Number * parts[1].Number
			}
		}
	}

	return sum, nil
}

func (a *Analyzer) EnginePartsNearPosition(row, col int) []EnginePart {

	parts := make([]EnginePart, 0)

	for _, p := range a.EngineParts {
		if p.IsNearPosition(row, col) {
			parts = append(parts, p)
		}
	}

	return parts
}
