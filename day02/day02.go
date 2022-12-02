package main

import (
	"adventofcode2022/utils"
	"strconv"
)

func solveA(input []string) int {
	score := 0

	for _, line := range input {
		if line == "" {
			continue
		}

		player1, player2 := int(line[0]-'A'+1), int(line[2]-'X'+1)

		score += player2

		if player2 == player1 {
			// draw
			score += 3
		} else if (player2-player1-1)%3 == 0 {
			// player2 wins
			score += 6
		}
	}

	return score
}

func solveB(input []string) int {
	score := 0

	for _, line := range input {
		if line == "" {
			continue
		}

		player1, outcome := int(line[0]-'A'+1), int(line[2]-'X')*3

		score += outcome

		switch outcome {
		case 0:
			// need to lose
			score += (player1+1)%3 + 1
		case 3:
			// need to end in draw
			score += player1
		case 6:
			// need to win
			score += player1%3 + 1
		}
	}

	return score
}

func main() {
	problem := utils.Must(utils.NewProblem(2022, 2))
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
