package main

import (
	"adventofcode2022/utils"
	"strconv"
	"strings"
)

func solve(input []string) (int, []byte) {
	vals := make([]byte, 0, 240)

	cycle, total := 1, 0
	val, next := 1, 20
	delta := 1

	var addx int
	for _, line := range input {
		if line == "" {
			continue
		}

		if line == "noop" {
			addx = 0
			delta = 1
		} else {
			addx = utils.Must(strconv.Atoi(line[5:]))
			delta = 2
		}

		for i := 0; i < delta; i++ {
			if cycle == next {
				total += val * next
				next += 40
			}

			pos := (cycle - 1) % 40
			if (pos-val)*(pos-val) <= 1 {
				vals = append(vals, '#')
			} else {
				vals = append(vals, '.')
			}

			cycle++
		}

		val += addx
	}

	return total, vals
}

func solveA(input []string) string {
	total, _ := solve(input)
	return strconv.Itoa(total)
}

func solveB(input []string) string {
	ans := make([]string, 0, 6)
	_, vals := solve(input)

	for i := 0; i+40 <= len(vals); i += 40 {
		ans = append(ans, string(vals[i:i+40]))
	}

	return strings.Join(ans, "\n")
}

func main() {
	problem := utils.Must(utils.NewProblem(2022, 10))
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
