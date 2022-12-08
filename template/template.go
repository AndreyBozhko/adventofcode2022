package main

import (
	"adventofcode2022/utils"
	"strconv"
)

func solveA(input []string) int {

	return 0
}

func solveB(input []string) int {

	return 0
}

func main() {
	problem := utils.Must(utils.NewProblem(2022, -1_234))
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
