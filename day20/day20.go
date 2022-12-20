package main

import (
	"adventofcode2022/utils"
	"strconv"
)

type El struct {
	value      int
	Prev, Next *El
}

func solveA(input []string) int {
	root := (*El)(nil)
	tail := (*El)(nil)

	elements := make([]*El, 0)

	for _, line := range input {
		if line == "" {
			continue
		}

		value := utils.Must(strconv.Atoi(line))
		cur := &El{value: value}

		if root == nil {
			root = cur
			tail = cur
		} else {
			tail.Next = cur
			cur.Prev = tail
			tail = cur
		}
		elements = append(elements, cur)
	}
	tail.Next = root
	root.Prev = tail

	length := len(elements)

	for _, el := range elements {
		vv := el.value
		if vv == 0 {
			continue
		}
		vv %= (length - 1)

		if el.value > 0 {
			cur := el
			el.Prev.Next = el.Next
			el.Next.Prev = el.Prev
			for jj := 0; jj < el.value; jj++ {
				cur = cur.Next
			}
			nei := cur.Next
			cur.Next = el
			el.Prev = cur
			el.Next = nei
			nei.Prev = el
		} else if el.value < 0 {
			cur := el
			el.Prev.Next = el.Next
			el.Next.Prev = el.Prev
			for jj := 0; jj > el.value; jj-- {
				cur = cur.Prev
			}
			nei := cur.Prev
			cur.Prev = el
			el.Next = cur
			el.Prev = nei
			nei.Next = el
		}
	}

	cur := root
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

func solveB(input []string) int {
	root := (*El)(nil)
	tail := (*El)(nil)

	elements := make([]*El, 0)

	for _, line := range input {
		if line == "" {
			continue
		}

		value := utils.Must(strconv.Atoi(line))
		cur := &El{value: value * 811589153}

		if root == nil {
			root = cur
			tail = cur
		} else {
			tail.Next = cur
			cur.Prev = tail
			tail = cur
		}
		elements = append(elements, cur)
	}
	tail.Next = root
	root.Prev = tail

	length := len(elements)

	for attempt := 0; attempt < 10; attempt++ {
		println("Attempt", attempt)
		for _, el := range elements {
			vv := el.value
			if vv == 0 {
				continue
			}
			vv %= (length - 1)

			if el.value > 0 {
				cur := el
				el.Prev.Next = el.Next
				el.Next.Prev = el.Prev
				for jj := 0; jj < vv; jj++ {
					cur = cur.Next
				}
				nei := cur.Next
				cur.Next = el
				el.Prev = cur
				el.Next = nei
				nei.Prev = el
			} else if el.value < 0 {
				cur := el
				el.Prev.Next = el.Next
				el.Next.Prev = el.Prev
				for jj := 0; jj > vv; jj-- {
					cur = cur.Prev
				}
				nei := cur.Prev
				cur.Prev = el
				el.Next = cur
				el.Prev = nei
				nei.Next = el
			}
		}
	}

	cur := root
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

// 1 1 5 1 1
