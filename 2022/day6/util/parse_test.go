package util

import (
	"testing"

	"github.com/EdgeJay/adventofcode/common/utils/files"
)

func BenchmarkParseLine(b *testing.B) {
	lines := files.ReadInputsFile("./input.txt")
	for n := 0; n < b.N; n++ {
		ParseLine(lines[0], 14)
	}
}

func BenchmarkParseLineImproved(b *testing.B) {
	lines := files.ReadInputsFile("./input.txt")
	for n := 0; n < b.N; n++ {
		ParseLineImproved(lines[0], 14)
	}
}
