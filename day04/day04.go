package main

import (
	"adventofcode2022/utils"
	"fmt"
	"strconv"
)

type Range struct {
	left, right int
}

func (r Range) Contains(other Range) bool {
	return r.left <= other.left && r.right >= other.right
}

func (r Range) Overlaps(other Range) bool {
	return r.right >= other.left && other.right >= r.left
}

func solve(input []string, condition func(r1, r2 Range) bool) int {
	total := 0
	var range1, range2 Range

	for _, line := range input {
		if line == "" {
			continue
		}

		_ = utils.Must(fmt.Sscanf(line, "%d-%d,%d-%d", &range1.left, &range1.right, &range2.left, &range2.right))

		if condition(range1, range2) {
			total += 1
		}
	}

	return total
}

func solveA(input []string) int {
	return solve(input, func(r1, r2 Range) bool {
		return r1.Contains(r2) || r2.Contains(r1)
	})
}

func solveB(input []string) int {
	return solve(input, func(r1, r2 Range) bool {
		return r1.Overlaps(r2)
	})
}

func main() {
	problem := utils.Must(utils.NewProblem(2022, 4))
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
