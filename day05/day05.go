package main

import (
	"adventofcode2022/utils"
	"fmt"
)

func solve(input []string, reverse bool) string {
	sz := (len(input[0]) + 1) / 4
	stacks := make([][]byte, sz)

	i := 0
	for input[i] != "" {
		i++
	}

	for j := i - 2; j >= 0; j-- {
		for st := 0; st < sz; st++ {
			c := input[j][4*st+1]
			if c != ' ' {
				stacks[st] = append(stacks[st], c)
			}
		}
	}

	var cnt, from, to int

	for j := i + 1; j < len(input); j++ {
		line := input[j]
		if line == "" {
			continue
		}

		_ = utils.Must(fmt.Sscanf(line, "move %d from %d to %d", &cnt, &from, &to))

		from--
		to--

		l := len(stacks[from])
		stacks[to] = append(stacks[to], stacks[from][l-cnt:]...)
		stacks[from] = stacks[from][:l-cnt]

		if reverse {
			utils.ReverseSlice(stacks[to][len(stacks[to])-cnt:])
		}
	}

	result := make([]byte, 0, sz)

	for _, stack := range stacks {
		result = append(result, stack[len(stack)-1])
	}

	return string(result)
}

func solveA(input []string) string {
	return solve(input, true)
}

func solveB(input []string) string {
	return solve(input, false)
}

func main() {
	problem := utils.Must(utils.NewProblem(2022, 5))
	input := utils.Must(problem.LoadInput())

	var result string
	if problem.IsA() {
		result = solveA(input)
	}

	if problem.IsB() {
		result = solveB(input)
	}

	_ = problem.WriteOutput(result)
}
