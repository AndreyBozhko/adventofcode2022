package main

import (
	"adventofcode2022/utils"
	"strconv"
)

type direction int

const (
	R direction = 0
	D direction = 1
	L direction = 2
	U direction = 3
)

var directions = [4]Point{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

type cell int

const (
	null cell = 0
	free cell = 1
	wall cell = 2
)

type Point struct {
	X, Y int
}

func (d *direction) TurnRight() {
	*d = (*d + 1) % 4
}

func (d *direction) TurnLeft() {
	*d = (*d + 3) % 4
}

type Map [][]cell

func (pos Point) Next(dir direction, m Map) Point {
	maxX, maxY := len(m[0]), len(m)

	step := directions[int(dir)]

	x, y := pos.X, pos.Y

	if step.Y == 0 {
		x = (pos.X + step.X + maxX) % maxX
		for m[pos.Y][x] == null {
			x = (x + step.X + maxX) % maxX
		}
	}

	if step.X == 0 {
		y = (pos.Y + step.Y + maxY) % maxY
		for m[y][pos.X] == null {
			y = (y + step.Y + maxY) % maxY
		}
	}

	return Point{x, y}
}

func parseInput(input []string) (m Map, start *Point, dir direction, inst string) {
	m = make([][]cell, len(input)-3)
	start = nil

	var i, maxX int
	var line string

	for i, line = range input {
		if line == "" {
			break
		}

		if len(line) > maxX {
			maxX = len(line)
		}
	}

	for i, line = range input {
		if line == "" {
			break
		}

		m[i] = make([]cell, maxX)

		for j, c := range line {
			switch c {
			case '.':
				m[i][j] = free
				if start == nil {
					start = &Point{j, i}
				}
			case '#':
				m[i][j] = wall
			}
		}
	}

	inst = input[i+1]
	return
}

func solve(input []string, wrap func(p Point, dir direction, m Map) Point) int {
	grid, pos, dir, instr := parseInput(input)

	for instr != "" {
		if instr[0] == 'R' {
			dir.TurnRight()
			instr = instr[1:]
		} else if instr[0] == 'L' {
			dir.TurnLeft()
			instr = instr[1:]
		} else {
			var idx int
			for idx = 0; idx < len(instr) && instr[idx] < 'A'; idx++ {
			}
			steps := utils.Must(strconv.Atoi(instr[:idx]))

			dxdy := directions[int(dir)]

			var p Point
			for s := 0; s < steps; s++ {
				p = *pos
				p.X, p.Y = p.X+dxdy.X, p.Y+dxdy.Y
				if grid[p.Y][p.X] == null {
					p = wrap(*pos, dir, grid)
				}
				if grid[p.Y][p.X] == wall {
					break
				}
				*pos = p
			}

			instr = instr[idx:]
		}
	}

	return 1000*(pos.Y+1) + 4*(pos.X+1) + int(dir)
}

func solveA(input []string) int {
	return solve(input, func(pos Point, dir direction, m Map) Point {
		maxX, maxY := len(m[0]), len(m)
		step := directions[int(dir)]
		x, y := pos.X, pos.Y

		if step.Y == 0 {
			x = (x + step.X + maxX) % maxX
			for m[pos.Y][x] == null {
				x = (x + step.X + maxX) % maxX
			}
		}

		if step.X == 0 {
			y = (y + step.Y + maxY) % maxY
			for m[y][pos.X] == null {
				y = (y + step.Y + maxY) % maxY
			}
		}

		return Point{x, y}
	})
}

func solveB(input []string) int {

	return 0
}

func main() {
	problem := utils.Must(utils.NewProblem(2022, 22))
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
