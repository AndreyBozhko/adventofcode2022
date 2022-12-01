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
	input := utils.Must(utils.Input(2022, 1))
	part := utils.ProblemPart()

	input = append(input, "")

	var result int
	if part == utils.A {
		result = solveA(input)
	}

	if part == utils.B {
		result = solveB(input)
	}

	utils.Output(part, strconv.Itoa(result))
}
