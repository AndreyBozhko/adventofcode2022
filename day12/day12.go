package main

import (
	"adventofcode2022/utils"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

func solve(lines []string, isStart func(c rune) bool) int {
	input := make([]string, 0, len(lines))
	input = append(input, lines...)

	queue1, queue2 := make([]Point, 0), make([]Point, 0)
	visited := make(map[Point]bool)

	var destX, destY int

	steps := make([][]int, len(input))
	for x, line := range input {
		steps[x] = make([]int, len(line))

		for y, c := range line {
			steps[x][y] = math.MaxInt32
			if isStart(c) {
				steps[x][y] = 0
				queue1 = append(queue1, Point{X: x, Y: y})
			}
		}

		input[x] = strings.ReplaceAll(input[x], "S", "a")

		if y := strings.Index(line, "E"); y != -1 {
			destX, destY = x, y
			input[x] = strings.ReplaceAll(input[x], "E", "z")
		}
	}

	maxX, maxY := len(input), len(input[0])

	neighbors := func(point Point) (result []Point) {
		result = make([]Point, 0, 4)
		if point.X > 0 {
			result = append(result, Point{X: point.X - 1, Y: point.Y})
		}
		if point.Y > 0 {
			result = append(result, Point{X: point.X, Y: point.Y - 1})
		}
		if point.X < maxX-1 {
			result = append(result, Point{X: point.X + 1, Y: point.Y})
		}
		if point.Y < maxY-1 {
			result = append(result, Point{X: point.X, Y: point.Y + 1})
		}

		return
	}

	for len(queue1) > 0 {

		for _, p := range queue1 {

			if visited[p] {
				continue
			}

			visited[p] = true

			if p.X == destX && p.Y == destY {
				return steps[p.X][p.Y]
			}

			for _, n := range neighbors(p) {
				stepCondition := steps[n.X][n.Y] > steps[p.X][p.Y]+1

				to, from := input[n.X][n.Y], input[p.X][p.Y]
				heightCondition := to <= from+1

				if stepCondition && heightCondition && !visited[n] {
					steps[n.X][n.Y] = steps[p.X][p.Y] + 1
					queue2 = append(queue2, n)
				}
			}
		}

		queue1, queue2 = queue2, make([]Point, 0)
	}

	panic("unreachable")
}

func solveA(input []string) int {
	return solve(input, func(c rune) bool {
		return c == 'S'
	})
}

func solveB(input []string) int {
	return solve(input, func(c rune) bool {
		return c == 'S' || c == 'a'
	})
}

func main() {
	problem := utils.Must(utils.NewProblem(2022, 12))
	input := utils.Must(problem.LoadInput())

	if input[len(input)-1] == "" {
		input = input[:len(input)-1]
	}

	var result int
	if problem.IsA() {
		result = solveA(input)
	}

	if problem.IsB() {
		result = solveB(input)
	}

	_ = problem.WriteOutput(strconv.Itoa(result))
}
