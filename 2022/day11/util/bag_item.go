package util

type BagItem struct {
	WorryLevel int
}

func NewBagItem(worryLevel int) *BagItem {
	return &BagItem{WorryLevel: worryLevel}
}
