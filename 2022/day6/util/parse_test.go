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
