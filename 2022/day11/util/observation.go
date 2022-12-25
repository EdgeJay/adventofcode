package util

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"

	mathUtils "github.com/EdgeJay/adventofcode/common/math"
	"github.com/EdgeJay/adventofcode/common/utils/str"
)

type Monkey struct {
	ID                  int
	BagItems            []*BagItem
	OperationFuncParams []string
	OperationFunc       OperationFunc
	TestFunc            func(*BagItem)
	PassMonkeyID        int
	FailMonkeyID        int
	Observation         *Observation
	ObservationCount    int
}

func (m *Monkey) GetOpFuncParamValue(param string, bagItem *BagItem) int {
	if param == "old" {
		return bagItem.WorryLevel
	} else if val, err := strconv.Atoi(param); err == nil {
		return val
	}
	return 0
}

func (m *Monkey) Observe() {
	fmt.Printf("Monkey %d observing items (%d remaining)\n", m.ID, m.GetRemainingBagItems())
	for {
		// skip if nothing to observe
		if m.GetRemainingBagItems() < 1 {
			break
		}

		m.ObserveNextItem()
	}
}

func (m *Monkey) ObserveNextItem() *BagItem {
	bagItem := m.RemoveFirstBagItem()
	a := m.GetOpFuncParamValue(m.OperationFuncParams[0], bagItem)
	b := m.GetOpFuncParamValue(m.OperationFuncParams[2], bagItem)

	reliefLevel := m.Observation.ReliefLevel
	if reliefLevel > 0 {
		bagItem.WorryLevel = int(math.Floor(float64(m.OperationFunc(a, b)) / float64(reliefLevel)))
	} else {
		bagItem.WorryLevel = m.OperationFunc(a, b)
	}
	m.TestFunc(bagItem)

	m.ObservationCount++

	return bagItem
}

func (m *Monkey) SetOperationFunc(parts ...string) {
	m.OperationFuncParams = parts
	switch parts[1] {
	case "+":
		m.OperationFunc = AdditionFunc
	case "-":
		m.OperationFunc = SubtractionFunc
	case "*":
		m.OperationFunc = MultiplicationFunc
	case "/":
		m.OperationFunc = DivisionFunc
	}
}

func (m *Monkey) SetWorryTestFunc(num string, passMonkeyID, failMonkeyID string) {
	val, err := strconv.Atoi(num)
	if err != nil {
		val = 1
	}

	if val, err := strconv.Atoi(passMonkeyID); err == nil {
		m.PassMonkeyID = val
	} else {
		log.Fatal(err)
	}

	if val, err := strconv.Atoi(failMonkeyID); err == nil {
		m.FailMonkeyID = val
	} else {
		log.Fatal(err)
	}

	m.TestFunc = func(bagItem *BagItem) {
		if m.Observation == nil {
			log.Fatal("m.Observation is nil")
		}

		if mathUtils.ModInt(bagItem.WorryLevel, val) == 0 {
			m.Observation.GetMonkeybyID(m.PassMonkeyID).AddBagItem(bagItem)
		} else {
			m.Observation.GetMonkeybyID(m.FailMonkeyID).AddBagItem(bagItem)
		}
	}
}

func (m *Monkey) GetRemainingBagItems() int {
	return len(m.BagItems)
}

func (m *Monkey) AddBagItem(item *BagItem) {
	m.BagItems = append(m.BagItems, item)
}

func (m *Monkey) RemoveFirstBagItem() *BagItem {
	var bagItem *BagItem
	numRemainingBagItems := m.GetRemainingBagItems()

	if numRemainingBagItems == 1 {
		bagItem = m.BagItems[0]
		m.BagItems = []*BagItem{}
	}

	if numRemainingBagItems > 1 {
		bagItem = m.BagItems[0]
		m.BagItems = m.BagItems[1:]
	}

	return bagItem
}

func (m *Monkey) ReportItemsWorryLevels() string {
	levels := make([]string, len(m.BagItems))
	for idx, item := range m.BagItems {
		levels[idx] = strconv.Itoa(item.WorryLevel)
	}
	return strings.Join(levels, ", ")
}

func NewMonkey(data string, ob *Observation) *Monkey {
	arr := strings.Split(data, "\n")

	// get Monkey ID
	if id, err := strconv.Atoi(arr[0][7:8]); err == nil {
		// get items
		items := strings.Split(str.TrimWhitespace(arr[1])[16:], ", ")
		bagItems := make([]*BagItem, 0)
		for _, item := range items {
			if num, err := strconv.Atoi(item); err == nil {
				bagItems = append(bagItems, NewBagItem(num))
			}
		}

		m := &Monkey{
			ID:               id,
			BagItems:         bagItems,
			Observation:      ob,
			ObservationCount: 0,
		}

		// get operation func
		opStr := strings.Split(strings.Trim(arr[2], " "), "Operation: new = ")[1]
		opArr := strings.Split(opStr, " ")
		m.SetOperationFunc(opArr...)

		// get worry test func
		testStr := strings.Split(strings.Trim(arr[3], " "), "Test: ")[1]
		testArr := strings.Split(testStr, "divisible by ")

		passMonkeyID := strings.Split(str.TrimWhitespace(arr[4]), " ")[5]
		failMonkeyID := strings.Split(str.TrimWhitespace(arr[5]), " ")[5]

		m.SetWorryTestFunc(testArr[1], passMonkeyID, failMonkeyID)

		return m
	}
	return nil
}

type Observation struct {
	Monkeys     []*Monkey
	RoundNumber int
	ReliefLevel int
}

func NewObservation(monkeys []string, reliefLevel int) *Observation {
	ob := &Observation{
		RoundNumber: 0,
		ReliefLevel: reliefLevel,
	}

	m := make([]*Monkey, len(monkeys))

	for n, data := range monkeys {
		m[n] = NewMonkey(data, ob)
	}

	ob.Monkeys = m

	return ob
}

func (ob *Observation) NextRound() {
	ob.RoundNumber++
	fmt.Printf("Starting round %d\n", ob.RoundNumber)

	for _, m := range ob.Monkeys {
		m.Observe()
	}

	fmt.Printf("End of round %d\n", ob.RoundNumber)
	ob.PrintRoundReport()
}

func (ob *Observation) PrintRoundReport() {
	report := ""
	for _, m := range ob.Monkeys {
		report += fmt.Sprintf("Monkey %d: %s\n", m.ID, m.ReportItemsWorryLevels())
	}
	fmt.Println(report)
}

func (ob *Observation) GetMonkeybyID(id int) *Monkey {
	for _, m := range ob.Monkeys {
		if m.ID == id {
			return m
		}
	}
	return nil
}

func (ob *Observation) GetMonkeysSortedByObservationCount() []*Monkey {
	arr := make([]*Monkey, len(ob.Monkeys))
	copy(arr, ob.Monkeys)

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].ObservationCount > arr[j].ObservationCount
	})

	return arr
}
