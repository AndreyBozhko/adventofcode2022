package main

import (
	"adventofcode2022/utils"
	"math"
	"strconv"
)

type move int

const (
	N move = 0
	S move = 1
	W move = 2
	E move = 3
)

var directions = map[move][]Point{
	N: {{-1, -1}, {0, -1}, {1, -1}},
	S: {{-1, 1}, {0, 1}, {1, 1}},
	W: {{-1, -1}, {-1, 0}, {-1, 1}},
	E: {{1, -1}, {1, 0}, {1, 1}},
}

type Point struct {
	X, Y int
}

func parseInput(input []string) map[Point]bool {
	m := make(map[Point]bool)

	for y, line := range input {
		if line == "" {
			continue
		}
		for x, c := range line {
			if c == '#' {
				m[Point{x, y}] = true
			}
		}
	}
	return m
}

func shouldMove(p Point, grid map[Point]bool) bool {
	for _, dir := range directions {
		for _, d := range dir {
			if _, ok := grid[Point{p.X + d.X, p.Y + d.Y}]; ok {
				return true
			}
		}
	}
	return false
}

func moveAround(grid map[Point]bool, prev move) (newGrid map[Point]bool, totalMoved int) {
	part1, part2 := make(map[Point]int), make(map[Point]bool)
	totalMoved = 0

	for pos := range grid {
		if !shouldMove(pos, grid) {
			continue
		}

		var canMove bool
		var choice move

		for i := 0; i < 4; i++ {
			choice = (prev + move(i)) % 4
			canMove = true
			for _, delta := range directions[choice] {
				if grid[Point{pos.X + delta.X, pos.Y + delta.Y}] {
					canMove = false
					break
				}
			}
			if canMove {
				dir := directions[choice][1]
				dest := Point{pos.X + dir.X, pos.Y + dir.Y}
				v := part1[dest]
				part1[dest] = v + 1
				break
			}
		}
	}

	for pos := range grid {
		if !shouldMove(pos, grid) {
			part2[pos] = true
			continue
		}

		var canMove bool
		var choice move

		for i := 0; i < 4; i++ {
			choice = (prev + move(i)) % 4
			canMove = true
			for _, delta := range directions[choice] {
				if grid[Point{pos.X + delta.X, pos.Y + delta.Y}] {
					canMove = false
					break
				}
			}
			if canMove {
				break
			}
		}
		if !canMove {
			part2[pos] = true
			continue
		}

		dir := directions[choice][1]
		dest := Point{pos.X + dir.X, pos.Y + dir.Y}
		if v := part1[dest]; v == 1 {
			part2[dest] = true
			totalMoved++
		} else {
			part2[pos] = true
		}
	}

	return part2, totalMoved
}

func area(grid map[Point]bool) int {
	xmin, xmax := math.MaxInt32, math.MinInt32
	ymin, ymax := math.MaxInt32, math.MinInt32

	occupied := 0
	for k := range grid {
		occupied++
		if k.X < xmin {
			xmin = k.X
		}
		if k.X > xmax {
			xmax = k.X
		}
		if k.Y < ymin {
			ymin = k.Y
		}
		if k.Y > ymax {
			ymax = k.Y
		}
	}
	return (xmax - xmin + 1) * (ymax - ymin + 1)
}

func solveA(input []string) int {
	grid := parseInput(input)

	for i := 0; i < 10; i++ {
		choice := move(i % 4)
		grid, _ = moveAround(grid, choice)
	}

	return area(grid) - len(grid)
}

func solveB(input []string) int {
	grid := parseInput(input)
	var moved int

	for i := 0; ; i++ {
		choice := move(i % 4)
		grid, moved = moveAround(grid, choice)
		if moved == 0 {
			return i + 1
		}
	}
}

func main() {
	problem := utils.Must(utils.NewProblem(2022, 23))
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
