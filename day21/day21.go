package main

import (
	"adventofcode2022/utils"
	"errors"
	"strconv"
	"strings"
)

type Monkey struct {
	Id    string
	Text  string
	Num   *int
	Op    string
	Left  *Monkey
	Right *Monkey
}

func (m *Monkey) Evaluate() error {
	if m.Num != nil {
		return nil
	}

	if m.Left == nil || m.Right == nil {
		return errors.New("cannot evaluate")
	}

	if m.Left.Num == nil || m.Right.Num == nil {
		return errors.New("cannot evaluate")
	}

	m1, m2 := m.Left, m.Right

	var res int
	switch m.Op {
	case "+":
		res = *m1.Num + *m2.Num
	case "-":
		res = *m1.Num - *m2.Num
	case "*":
		res = *m1.Num * *m2.Num
	case "/":
		res = *m1.Num / *m2.Num
	}

	m.Num = &res
	return nil
}

func (m *Monkey) YellBack() {
	if m.Left == nil || m.Right == nil {
		return
	}

	var res int

	if m.Right.Num == nil {
		switch m.Op {
		case "+":
			res = *m.Num - *m.Left.Num
		case "-":
			res = *m.Left.Num - *m.Num
		case "*":
			res = *m.Num / *m.Left.Num
		case "/":
			res = *m.Left.Num / *m.Num
		}
		m.Right.Num = &res
		m.Right.YellBack()
	} else {
		switch m.Op {
		case "+":
			res = *m.Num - *m.Right.Num
		case "-":
			res = *m.Num + *m.Right.Num
		case "*":
			res = *m.Num / *m.Right.Num
		case "/":
			res = *m.Num * *m.Right.Num
		}
		m.Left.Num = &res
		m.Left.YellBack()
	}
}

func parseMonkeys(input []string) map[string]*Monkey {
	monkeys := make(map[string]*Monkey)

	for _, line := range input {
		if line == "" {
			continue
		}

		parts := strings.Split(line, ": ")
		m := &Monkey{Id: parts[0], Text: parts[1]}
		monkeys[m.Id] = m
	}

	for _, m := range monkeys {
		if val, err := strconv.Atoi(m.Text); err == nil {
			m.Num = &val
			continue
		}

		parts := strings.Split(m.Text, " ")
		m.Left, m.Op, m.Right = monkeys[parts[0]], parts[1], monkeys[parts[2]]
	}

	return monkeys
}

func topoSort(cur *Monkey, visited map[string]bool, stack []*Monkey) []*Monkey {
	visited[cur.Id] = true

	for _, m := range []*Monkey{cur.Left, cur.Right} {
		if m == nil {
			continue
		}
		if visited[m.Id] {
			continue
		}
		stack = topoSort(m, visited, stack)
	}
	stack = append(stack, cur)
	return stack
}

func solveA(input []string) int {
	monkeys := parseMonkeys(input)

	root := monkeys["root"]

	sorted := topoSort(root, make(map[string]bool), nil)

	for _, m := range sorted {
		_ = m.Evaluate()
	}

	return *root.Num
}

func solveB(input []string) int {
	monkeys := parseMonkeys(input)

	root := monkeys["root"]
	root.Op = "="

	monkeys["humn"].Num = nil

	sorted := topoSort(root, make(map[string]bool), nil)

	for _, m := range sorted {
		_ = m.Evaluate()
	}

	if root.Left.Num == nil {
		root.Left.Num = root.Right.Num
		root.Left.YellBack()
	} else {
		root.Right.Num = root.Left.Num
		root.Right.YellBack()
	}

	return *monkeys["humn"].Num
}

func main() {
	problem := utils.Must(utils.NewProblem(2022, 21))
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
