package util

import (
	"testing"

	"github.com/EdgeJay/adventofcode/common/utils"
	"github.com/stretchr/testify/assert"
)

func TestNewMonkey(t *testing.T) {
	input := `Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0
	`

	m := NewMonkey(input, nil)

	if m.ID != 1 {
		t.Errorf("Expected ID to be %d, got %d", 1, m.ID)
	}

	if m.GetRemainingBagItems() != 4 {
		t.Errorf("Expected no. of bag items to be %d, got %d", 4, m.GetRemainingBagItems())
	}

	assert.Contains(t, m.OperationFuncParams, "old")
	assert.Contains(t, m.OperationFuncParams, "+")
	assert.Contains(t, m.OperationFuncParams, "6")
	assert.Contains(t, utils.GetFunctionName(m.OperationFunc), "AdditionFunc")

	if m.PassMonkeyID != 2 {
		t.Errorf("Expected PassMonkeyID to be %d, got %d", 2, m.PassMonkeyID)
	}

	if m.FailMonkeyID != 0 {
		t.Errorf("Expected FailMonkeyID to be %d, got %d", 2, m.FailMonkeyID)
	}
}

func TestMonkeyObserve(t *testing.T) {
	ob := &Observation{
		RoundNumber: 0,
	}

	input := `Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0
	`
	m1 := NewMonkey(input, ob)
	assert.EqualValues(t, 4, m1.GetRemainingBagItems())

	input = `Monkey 2:
	Starting items: 79, 60, 97
	Operation: new = old * old
	Test: divisible by 13
	  If true: throw to monkey 1
	  If false: throw to monkey 3
	`
	m2 := NewMonkey(input, ob)
	assert.EqualValues(t, 3, m2.GetRemainingBagItems())

	input = `Monkey 0:
	Starting items: 79, 98
	Operation: new = old * 19
	Test: divisible by 23
	  If true: throw to monkey 2
	  If false: throw to monkey 3
	`
	m3 := NewMonkey(input, ob)
	assert.EqualValues(t, 2, m3.GetRemainingBagItems())

	ob.Monkeys = []*Monkey{m1, m2, m3}

	// Round #1
	bagItem := m1.ObserveNextItem()

	assert.EqualValues(t, 20, bagItem.WorryLevel)
	// after 1 observation Monkey 1 should have passed item to Monkey 0
	assert.EqualValues(t, 3, len(m3.BagItems))

	// Round #2
	bagItem = m1.ObserveNextItem()
	assert.EqualValues(t, 24, bagItem.WorryLevel)
	// after 1 observation Monkey 1 should have passed item to Monkey 0
	assert.EqualValues(t, 4, len(m3.BagItems))
}

func TestObservationNextRound(t *testing.T) {
	ob := &Observation{
		RoundNumber: 0,
	}

	input := `Monkey 0:
	Starting items: 79, 98
	Operation: new = old * 19
	Test: divisible by 23
	  If true: throw to monkey 2
	  If false: throw to monkey 3
	`
	m0 := NewMonkey(input, ob)
	assert.EqualValues(t, 2, m0.GetRemainingBagItems())

	input = `Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0
	`
	m1 := NewMonkey(input, ob)
	assert.EqualValues(t, 4, m1.GetRemainingBagItems())

	input = `Monkey 2:
	Starting items: 79, 60, 97
	Operation: new = old * old
	Test: divisible by 13
	  If true: throw to monkey 1
	  If false: throw to monkey 3
	`
	m2 := NewMonkey(input, ob)
	assert.EqualValues(t, 3, m2.GetRemainingBagItems())

	input = `Monkey 3:
	Starting items: 74
	Operation: new = old + 3
	Test: divisible by 17
	  If true: throw to monkey 0
	  If false: throw to monkey 1
	`
	m3 := NewMonkey(input, ob)
	assert.EqualValues(t, 1, m3.GetRemainingBagItems())

	ob.Monkeys = []*Monkey{m0, m1, m2, m3}

	// Round #1
	ob.NextRound()

	assert.EqualValues(t, 1, ob.RoundNumber)
	assert.EqualValues(t, 5, m1.GetRemainingBagItems())
	assert.EqualValues(t, 2080, m1.BagItems[3].WorryLevel)
	assert.EqualValues(t, 26, m1.BagItems[4].WorryLevel)

	assert.EqualValues(t, 1, m3.GetRemainingBagItems())
	assert.EqualValues(t, 500, m3.BagItems[0].WorryLevel)
}
