package main

import (
	"adventofcode2022/utils"
	"fmt"
	"strconv"
)

func findCommon(left, right string) (rune, error) {
	m := make(map[rune]bool)
	for _, el := range left {
		m[el] = true
	}

	for _, el := range right {
		if _, ok := m[el]; ok {
			return el, nil
		}
	}

	return rune(0), fmt.Errorf("no common rune in '%s' and '%s'", left, right)
}

func findBadge(buf []string) (rune, error) {
	if len(buf) != 3 {
		return rune(0), fmt.Errorf("unexpected buf length: %d", len(buf))
	}

	m := make(map[rune]int)

	for _, el := range buf[0] {
		m[el] = 1
	}

	for _, el := range buf[1] {
		if _, ok := m[el]; ok {
			m[el] = 2
		}
	}

	for _, el := range buf[2] {
		if v, ok := m[el]; ok && v == 2 {
			return el, nil
		}
	}

	return rune(0), fmt.Errorf("badge not found among '%s', '%s', '%s'", buf[0], buf[1], buf[2])
}

func priority(item rune) int {
	if item < 'a' {
		return int(item - 'A' + 27)
	}

	return int(item - 'a' + 1)
}

func solveA(input []string) int {
	total := 0

	for _, line := range input {
		if line == "" {
			continue
		}

		mid := len(line) / 2
		left, right := line[:mid], line[mid:]
		item := utils.Must(findCommon(left, right))

		total += priority(item)
	}

	return total
}

func solveB(input []string) int {
	total := 0

	if input[0] == "" {
		input = input[1:]
	}

	for i := 0; i+2 < len(input); i += 3 {
		item := utils.Must(findBadge(input[i : i+3]))
		total += priority(item)
	}

	return total
}

func main() {
	problem := utils.Must(utils.NewProblem(2022, 3))
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
