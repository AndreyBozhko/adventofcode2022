package main

import (
	"adventofcode2022/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	items     []int
	operation func(val int) int
	divisor   int
	ifTrue    int
	ifFalse   int
	inspect   int
}

func (m *Monkey) DoInspect(val int) int {
	m.inspect++
	return m.operation(val)
}

func newAddOp(right string) func(x int) int {
	if r, err := strconv.Atoi(right); err == nil {
		return func(old int) int {
			return old + r
		}
	} else {
		return func(old int) int {
			return old + old
		}
	}
}

func newMulOp(right string) func(x int) int {
	if r, err := strconv.Atoi(right); err == nil {
		return func(old int) int {
			return old * r
		}
	} else {
		return func(old int) int {
			return old * old
		}
	}
}

func NewMonkey(lines []string) *Monkey {
	m := &Monkey{}

	itms := strings.Split(lines[1][18:], ", ")
	for _, it := range itms {
		m.items = append(m.items, utils.Must(strconv.Atoi(it)))
	}

	var op, right string
	_ = utils.Must(fmt.Sscanf(lines[2], "  Operation: new = old %s %s", &op, &right))

	if op == "+" {
		m.operation = newAddOp(right)
	} else {
		m.operation = newMulOp(right)
	}

	m.divisor = utils.Must(strconv.Atoi(lines[3][21:]))
	m.ifTrue = utils.Must(strconv.Atoi(lines[4][29:]))
	m.ifFalse = utils.Must(strconv.Atoi(lines[5][30:]))

	return m
}

func solve(input []string, shouldDivide bool, rounds int) int {
	lcm := 1

	monkeys := make([]*Monkey, len(input)/7)
	for i := 0; i < len(monkeys); i++ {
		monkeys[i] = NewMonkey(input[i*7 : (i+1)*7])

		lcm *= monkeys[i].divisor
	}

	for round := 0; round < rounds; round++ {
		for _, m := range monkeys {
			for _, item := range m.items {
				level := m.DoInspect(item)
				if shouldDivide {
					level = level / 3
				}

				next := m.ifFalse
				if level%m.divisor == 0 {
					next = m.ifTrue
				}

				monkeys[next].items = append(monkeys[next].items, level%lcm)
			}
			m.items = make([]int, 0)
		}
	}

	inspections := make([]int, 0, len(monkeys))
	for _, m := range monkeys {
		inspections = append(inspections, -m.inspect)
	}
	sort.Ints(inspections)

	return inspections[0] * inspections[1]
}

func solveA(input []string) int {
	return solve(input, true, 20)
}

func solveB(input []string) int {
	return solve(input, false, 10000)
}

func main() {
	problem := utils.Must(utils.NewProblem(2022, 11))
	input := utils.Must(problem.LoadInput())

	var result int
	if problem.IsA() {
		result = solveA(input)
	}

	if problem.IsB() {
		result = solveB(input)
	}

	_ = problem.WriteOutput(strconv.Itoa(result))
}
