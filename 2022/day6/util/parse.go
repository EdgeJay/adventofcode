package util

import (
	"fmt"
	"strings"

	"github.com/thoas/go-funk"
)

func ParseLine(line string, markerSize int) {
	t := []rune{}
	afterPos := 0

	for idx, c := range line {
		if len(t) == markerSize {
			afterPos = idx
			break
		}

		if pos := funk.IndexOfInt32(t, c); pos != -1 {
			if pos+1 < len(t) {
				t = t[pos+1:]
			} else {
				t = []rune{}
			}
		}

		t = append(t, c)
	}

	fmt.Println(string(t), afterPos)
}

func ParseLineImproved(line string, markerSize int) {
	for n := 0; n < len(line); n++ {
		str := line[n : n+markerSize]
		if len(funk.UniqString(strings.Split(str, ""))) == markerSize {
			fmt.Println(str, n+markerSize)
			break
		}
	}
}
