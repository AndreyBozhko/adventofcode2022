package main

import (
	"adventofcode2022/utils"
	"sort"
	"strconv"
)

func solveA(input []string) int {
	max := -1
	cur := 0

	for _, el := range input {
		if calories, err := strconv.Atoi(el); err == nil {
			cur += calories
		} else {
			if cur >= max {
				max = cur
			}
			cur = 0
		}
	}

	return max
}

func solveB(input []string) int {
	calories := make([]int, 0)
	cur := 0

	for _, el := range input {
		if cal, err := strconv.Atoi(el); err == nil {
			cur -= cal
		} else {
			calories = append(calories, cur)
			cur = 0
		}
	}

	sort.Ints(calories)
	ans := -(calories[0] + calories[1] + calories[2])

	return ans
}

func main() {
	problem := utils.Must(utils.NewProblem(2022, 1))
	input := utils.Must(problem.LoadInput())

	input = append(input, "")

	var result int
	if problem.IsA() {
		result = solveA(input)
	}

	if problem.IsB() {
		result = solveB(input)
	}

	_ = problem.WriteOutput(strconv.Itoa(result))
}
