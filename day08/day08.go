package main

import (
	"adventofcode2022/utils"
	"strconv"
)

func solveA(input []string) int {

	if input[len(input)-1] == "" {
		input = input[:len(input)-1]
	}

	visible := make(map[int]bool)

	for row := 1; row < len(input)-1; row++ {
		line := input[row]
		max := line[0]
		for col := 1; col < len(line)-1; col++ {
			if line[col] > max {
				max = line[col]
				visible[row*1000+col] = true
			}
		}

		max = line[len(line)-1]
		for col := len(line) - 2; col > 0; col-- {
			if line[col] > max {
				max = line[col]
				visible[row*1000+col] = true
			}
		}
	}

	for col := 1; col < len(input[0])-1; col++ {
		max := input[0][col]
		for row := 1; row < len(input)-1; row++ {
			if input[row][col] > max {
				max = input[row][col]
				visible[row*1000+col] = true
			}

		}

		max = input[len(input)-1][col]
		for row := len(input) - 2; row > 0; row-- {
			if input[row][col] > max {
				max = input[row][col]
				visible[row*1000+col] = true
			}

		}
	}

	return len(visible) + 2*len(input) + 2*len(input[0]) - 4
}

func treeScore(input []string, i, j int) (score int) {
	score = 1

	var ii, jj int

	for ii = i - 1; ; ii-- {
		if ii <= 0 || input[ii][j] >= input[i][j] {
			score *= i - ii
			break
		}
	}

	for ii = i + 1; ; ii++ {
		if ii >= len(input)-1 || input[ii][j] >= input[i][j] {
			score *= ii - i
			break
		}
	}

	for jj = j - 1; ; jj-- {
		if jj <= 0 || input[i][jj] >= input[i][j] {
			score *= j - jj
			break
		}
	}

	for jj = j + 1; ; jj++ {
		if jj >= len(input[0])-1 || input[i][jj] >= input[i][j] {
			score *= jj - j
			break
		}
	}

	return score
}

func solveB(input []string) int {
	max := 0

	if input[len(input)-1] == "" {
		input = input[:len(input)-1]
	}

	for row := 1; row < len(input)-1; row++ {
		for col := 1; col < len(input[0])-1; col++ {
			score := treeScore(input, row, col)
			if score > max {
				max = score
			}
		}
	}

	return max
}

func main() {
	problem := utils.Must(utils.NewProblem(2022, 8))
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
