package main

import (
	"adventofcode2022/utils"
	"fmt"
	"strconv"
)

type Point struct {
	x, y int
}

func (t *Point) MoveTowards(h Point) {
	dx, dy := h.x-t.x, h.y-t.y
	if dx*dx > 1 || dy*dy > 1 {
		t.x += utils.Signum(dx)
		t.y += utils.Signum(dy)
	}
}

func solve(input []string, sz int) int {
	visited := make(map[Point]bool)

	var d string
	var l int

	knots := make([]Point, sz)
	head, tail := &knots[0], &knots[sz-1]

	visited[*tail] = true

	for _, line := range input {
		if line == "" {
			continue
		}

		_ = utils.Must(fmt.Sscanf(line, "%s %d", &d, &l))

		for i := 0; i < l; i++ {
			switch d {
			case "U":
				head.y += 1
			case "D":
				head.y -= 1
			case "R":
				head.x += 1
			case "L":
				head.x -= 1
			}

			for j := 1; j < sz; j++ {
				knots[j].MoveTowards(knots[j-1])
			}

			visited[*tail] = true
		}
	}

	return len(visited)
}

func solveA(input []string) int {
	return solve(input, 2)
}

func solveB(input []string) int {
	return solve(input, 10)
}

func main() {
	problem := utils.Must(utils.NewProblem(2022, 9))
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
