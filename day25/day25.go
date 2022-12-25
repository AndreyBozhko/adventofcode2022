package main

import (
	"adventofcode2022/utils"
)

func solveA(input []string) string {
	total := 0
	for _, line := range input {
		if line == "" {
			continue
		}

		pow := 1
		for idx := len(line) - 1; idx >= 0; idx-- {
			c := line[idx]
			var num int
			switch c {
			case '2':
				num = 2
			case '1':
				num = 1
			case '0':
				num = 0
			case '-':
				num = -1
			case '=':
				num = -2
			}

			total += pow * num
			pow *= 5
		}
	}

	rslt := make([]int, 0)
	overflow := 0

	for total != 0 {
		cur := total%5 + overflow
		overflow, cur = cur/5, cur%5
		if cur > 2 {
			cur -= 5
			overflow++
		}
		rslt = append(rslt, cur)
		total /= 5
	}
	if overflow > 0 {
		rslt = append(rslt, overflow)
	}

	s := make([]byte, len(rslt))
	for idx := range s {
		c := rslt[len(rslt)-idx-1]
		switch c {
		case 2:
			s[idx] = '2'
		case 1:
			s[idx] = '1'
		case 0:
			s[idx] = '0'
		case -1:
			s[idx] = '-'
		case -2:
			s[idx] = '='
		}
	}

	return string(s)
}

func main() {
	problem := utils.Must(utils.NewProblem(2022, 25))
	input := utils.Must(problem.LoadInput())

	var result string
	if problem.IsA() {
		result = solveA(input)
	}

	_ = problem.WriteOutput(result)
}
