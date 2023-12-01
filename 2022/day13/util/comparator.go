package util

import (
	"fmt"
)

type Comparator struct {
	Left  string
	Right string
}

type ArraySearchContext struct {
	Runes  *[]rune
	Output *[]int
	Offset int
}

func isStartOfArray(r rune, offset int) bool {
	if offset == 0 {
		return false
	}
	return r == '['
}

func isEndOfArray(r rune) bool {
	return r == ']'
}

func (c *Comparator) packetsInRightOrder(leftOutput, rightOutput []int) bool {
	if len(leftOutput) > len(rightOutput) {
		return false
	}

	for n := 0; n < len(leftOutput); n++ {
		if leftOutput[n] > rightOutput[n] {
			return false
		}
	}

	return true
}

func (c *Comparator) doNextCheck(leftContext, rightContext *ArraySearchContext) bool {
	fmt.Println("left", string(*leftContext.Runes), "right", string(*rightContext.Runes))

	for leftContext.Offset < len(*leftContext.Runes) && rightContext.Offset < len(*rightContext.Runes) {
		charLt := (*leftContext.Runes)[leftContext.Offset]
		charRt := (*rightContext.Runes)[rightContext.Offset]

		if isStartOfArray(charLt, leftContext.Offset) || isStartOfArray(charRt, rightContext.Offset) {
			var newComparator *Comparator

			if isStartOfArray(charLt, leftContext.Offset) && !isStartOfArray(charRt, rightContext.Offset) {
				newComparator = NewComparator(
					string((*leftContext.Runes)[leftContext.Offset+1:]),
					fmt.Sprintf("[%s]", string((*rightContext.Runes)[rightContext.Offset])),
				)
			} else if !isStartOfArray(charLt, leftContext.Offset) && isStartOfArray(charRt, rightContext.Offset) {
				newComparator = NewComparator(
					fmt.Sprintf("[%s]", string((*leftContext.Runes)[leftContext.Offset])),
					string((*rightContext.Runes)[rightContext.Offset+1:]),
				)
			} else {
				newComparator = NewComparator(
					string((*leftContext.Runes)[leftContext.Offset+1:]),
					string((*rightContext.Runes)[rightContext.Offset+1:]),
				)
			}

			if newComparator.Check() {
				continue
			} else {
				return false
			}
		} else if isEndOfArray(charLt) || isEndOfArray(charRt) {
			break
		} else {
			*leftContext.Output = append(*leftContext.Output, int(charLt))
			*rightContext.Output = append(*rightContext.Output, int(charRt))
		}

		leftContext.Offset++
		rightContext.Offset++
	}

	return c.packetsInRightOrder(*leftContext.Output, *rightContext.Output)
}

func (c *Comparator) Check() bool {
	runesLeft := []rune(c.Left)
	runesRight := []rune(c.Right)

	leftContext := &ArraySearchContext{
		Runes:  &runesLeft,
		Output: &[]int{},
		Offset: 0,
	}

	rightContext := &ArraySearchContext{
		Runes:  &runesRight,
		Output: &[]int{},
		Offset: 0,
	}

	return c.doNextCheck(leftContext, rightContext)
}

func NewComparator(left, right string) *Comparator {
	fmt.Println("New comparator initialised", left, right)

	return &Comparator{
		Left:  left,
		Right: right,
	}
}
