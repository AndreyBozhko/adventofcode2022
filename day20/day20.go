package main

import (
	"adventofcode2022/utils"
	"strconv"
)

type Element struct {
	value      int
	Prev, Next *Element
}

type LinkedList struct {
	data []*Element
}

func parseInput(input []string, factor int) LinkedList {
	ll := LinkedList{}

	root := (*Element)(nil)
	tail := (*Element)(nil)

	for _, line := range input {
		if line == "" {
			continue
		}

		value := utils.Must(strconv.Atoi(line))
		cur := &Element{value: value * factor}

		if root == nil {
			root = cur
			tail = cur
		} else {
			tail.Next = cur
			cur.Prev = tail
			tail = cur
		}
		ll.data = append(ll.data, cur)
	}
	tail.Next = root
	root.Prev = tail

	return ll
}

func (lst LinkedList) Mix() {
	length := len(lst.data)

	for _, el := range lst.data {
		v := el.value
		if v == 0 {
			continue
		}

		v %= (length - 1)
		if v < 0 {
			v += length - 1
		}

		cur := el
		el.Prev.Next = el.Next
		el.Next.Prev = el.Prev

		for jj := 0; jj < v; jj++ {
			cur = cur.Next
		}

		nei := cur.Next
		cur.Next = el
		el.Prev = cur
		el.Next = nei
		nei.Prev = el
	}
}

func solve(lst LinkedList, attempts int) int {

	for attempt := 0; attempt < attempts; attempt++ {
		lst.Mix()
	}

	cur := lst.data[0]
	for cur.value != 0 {
		cur = cur.Next
	}

	total := 0
	for ii := 1; ii <= 3000; ii++ {
		cur = cur.Next
		if ii%1000 == 0 {
			total += cur.value
		}
	}

	return total
}

func solveA(input []string) int {
	lst := parseInput(input, 1)
	return solve(lst, 1)
}

func solveB(input []string) int {
	lst := parseInput(input, 811589153)
	return solve(lst, 10)
}

func main() {
	problem := utils.Must(utils.NewProblem(2022, 20))
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
