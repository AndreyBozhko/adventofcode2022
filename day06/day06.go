package main

import (
	"adventofcode2022/utils"
	"strconv"
)

func solve(input string, sz int) int {

	counter := utils.NewCounter[byte]()

	for i := 0; i < sz-1; i++ {
		counter.Add(input[i])
	}

	for i := sz; i <= len(input); i++ {
		counter.Add(input[i-1])
		if counter.Len() == sz {
			return i
		}
		counter.Remove(input[i-sz])
	}
	return -1
}

func solveA(input string) int {
	return solve(input, 4)
}

func solveB(input string) int {
	return solve(input, 14)
}

func main() {
	problem := utils.Must(utils.NewProblem(2022, 6))
	input := utils.Must(problem.LoadInput())

	var result int
	if problem.IsA() {
		result = solveA(input[0])
	}

	if problem.IsB() {
		result = solveB(input[0])
	}

	_ = problem.WriteOutput(strconv.Itoa(result))
}
