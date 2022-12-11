package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/matgd/advent2022/utils"
)

type Item struct {
	worryLevel int
}

type Monkey struct {
	id                  int
	items               []*Item
	operationOperator   string
	operationArg        int
	operationArgIsOld   bool
	testDivisibleArg    int
	testSuccessMonkeyId int
	testFailMonkeyId    int
	inspectedCount      int
	allMonkeys          *[]*Monkey
}

func (m *Monkey) Play() {
	for len(m.items) > 0 {
		m.inspectedCount++
		m.IncreaseWorryLevel(m.items[0])
		m.GetBored(m.items[0])
		m.PassToOtherMonkey(m.items[0])
		m.items = m.items[1:]
	}
}

func (m Monkey) IncreaseWorryLevel(item *Item) {
	var arg int = 0
	if m.operationArgIsOld {
		arg = item.worryLevel
	} else {
		arg = m.operationArg
	}
	switch m.operationOperator {
	case "+":
		item.worryLevel += arg
	case "*":
		item.worryLevel *= arg
	case "-":
		item.worryLevel -= arg
	}
}

func (m Monkey) GetBored(item *Item) {
	rounded := math.Floor(float64(item.worryLevel) / 3.0)

	item.worryLevel = int(rounded)
}

func (m Monkey) PerformTest(item *Item) bool {
	return item.worryLevel%m.testDivisibleArg == 0
}

func (m Monkey) PassToOtherMonkey(item *Item) {
	if m.PerformTest(item) {
		(*m.allMonkeys)[m.testSuccessMonkeyId].items = append((*m.allMonkeys)[m.testSuccessMonkeyId].items, item)
	} else {
		(*m.allMonkeys)[m.testFailMonkeyId].items = append((*m.allMonkeys)[m.testFailMonkeyId].items, item)
	}
}

func (m Monkey) String() string {
	return fmt.Sprintf("Monkey %d: %d items, %d inspected", m.id, len(m.items), m.inspectedCount)
}

func parseMonkeyId(s string) int {
	split := strings.Split(s, "Monkey ")
	rawId := strings.Split(split[1], ":")[0]
	id, _ := strconv.Atoi(rawId)
	return id
}

func parseStartingItems(s string) []*Item {
	split := strings.Split(s, "Starting items: ")
	rawItems := strings.Split(split[1], ", ")

	items := make([]*Item, len(rawItems))
	for i, rawItem := range rawItems {
		worryLevel, _ := strconv.Atoi(rawItem)
		items[i] = &Item{worryLevel}
	}
	return items
}

func parseOperation(s string) (string, int, bool) {
	split := strings.Split(s, " = ")
	rawOperation := split[1]

	arrOperation := strings.Split(rawOperation, " ")
	operator := arrOperation[1]
	operatorArg := arrOperation[2]

	operatorArgInt, err := strconv.Atoi(operatorArg)
	argIsOld := err != nil
	return operator, operatorArgInt, argIsOld
}

func parseTest(s string) int {
	split := strings.Split(s, "by ")
	rawArg := split[1]
	arg, _ := strconv.Atoi(rawArg)
	return arg
}

func parseTestResult(s string) int {
	split := strings.Split(s, "monkey ")
	rawId := split[1]
	id, _ := strconv.Atoi(rawId)
	return id
}

func Day11(filename string) int {
	lines := utils.ReadFileLines(filename)

	monkeys := make([]*Monkey, 0, 10)

	for i := 0; i < (len(lines)+1)/7; i++ {
		id := parseMonkeyId(lines[i*7])
		startingItems := parseStartingItems(lines[i*7+1])
		operationOperator, operationArg, operationArgIsOld := parseOperation(lines[i*7+2])
		testDivisibleArg := parseTest(lines[i*7+3])
		testSuccessMonkeyId := parseTestResult(lines[i*7+4])
		testFailMonkeyId := parseTestResult(lines[i*7+5])
		// Empty line.

		monkey := &Monkey{
			id:                  id,
			items:               startingItems,
			operationOperator:   operationOperator,
			operationArg:        operationArg,
			operationArgIsOld:   operationArgIsOld,
			testDivisibleArg:    testDivisibleArg,
			testSuccessMonkeyId: testSuccessMonkeyId,
			testFailMonkeyId:    testFailMonkeyId,
			inspectedCount:      0,
			allMonkeys:          &monkeys,
		}
		monkeys = append(monkeys, monkey)
	}

	for i := 0; i < 20; i++ {
		for _, monkey := range monkeys {
			monkey.Play()
		}
	}

	inspections := []int{}
	for _, monkey := range monkeys {
		inspections = append(inspections, monkey.inspectedCount)
	}
	sort.Ints(inspections)
	return inspections[len(inspections)-1] * inspections[len(inspections)-2]
}

func main() {
	fmt.Println("Day 11, Part 1:", Day11("input.txt"))
	fmt.Println("Day 11, Part 2:", Day11("input_example.txt"))
}
