package util

import (
	"math"
)

type Register struct {
	Name      string
	Cycle     int
	V         int
	ValuesMap map[int]int
}

func NewRegister(name string, startValue int) *Register {
	return &Register{
		Name:      name,
		Cycle:     0,
		V:         startValue,
		ValuesMap: make(map[int]int),
	}
}

func (r *Register) RecordCycle() {
	r.Cycle++
	if math.Mod(float64(r.Cycle), 20) == 0 {
		r.ValuesMap[r.Cycle] = r.V
	}
}

func (r *Register) Noop() {
	r.RecordCycle()
}

func (r *Register) Addx(value int) {
	r.RecordCycle()

	r.RecordCycle()
	r.V += value
}

func (r *Register) GetValue() int {
	return r.V
}

func (r *Register) SumOfValuesAtCycles(cycles ...int) int {
	sum := 0
	for _, c := range cycles {
		sum += c * r.ValuesMap[c]
	}
	return sum
}
