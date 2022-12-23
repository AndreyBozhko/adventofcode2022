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

type wrappingFunc func(p Point, dir direction, m Map) (Point, direction)

func (m Map) At(p Point) cell {
	if p.Y < 0 || p.Y >= len(m) {
		return null
	}
	if p.X < 0 || p.X >= len(m[p.Y]) {
		return null
	}
	return m[p.Y][p.X]
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

func solve(input []string, wrap wrappingFunc) int {
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

			var p Point
			var d direction

			for s := 0; s < steps; s++ {
				p, d = *pos, dir
				dxdy := directions[int(dir)]

				p.X, p.Y = p.X+dxdy.X, p.Y+dxdy.Y
				if grid.At(p) == null {
					p, d = wrap(*pos, d, grid)
				}
				if grid.At(p) == wall {
					break
				}
				*pos = p
				dir = d
			}

			instr = instr[idx:]
		}
	}

	return 1000*(pos.Y+1) + 4*(pos.X+1) + int(dir)
}

func solveA(input []string) int {
	return solve(input, func(pos Point, dir direction, m Map) (Point, direction) {
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

		return Point{x, y}, dir
	})
}

func solveB(input []string) int {
	sz := 50
	return solve(input, func(p Point, dir direction, m Map) (next Point, d direction) {
		faces := [][]int{{0, 1, 2}, {0, 3, 0}, {4, 5, 0}, {6, 0, 0}}
		f := faces[p.Y/sz][p.X/sz]

		switch f {
		case 1:
			switch dir {
			case L: // -> 4
				next, d = Point{0, 3*sz - 1 - p.Y}, R
			case U: // -> 6
				next, d = Point{0, p.X + 2*sz}, R
			default:
				panic("should not happen")
			}
		case 2:
			switch dir {
			case R: // -> 5
				next, d = Point{2*sz - 1, 3*sz - 1 - p.Y}, L
			case D: // -> 3
				next, d = Point{2*sz - 1, p.X - sz}, L
			case U: // -> 6
				next, d = Point{p.X - 2*sz, 4*sz - 1}, U
			default:
				panic("should not happen")
			}
		case 3:
			switch dir {
			case R: // -> 2
				next, d = Point{p.Y + sz, sz - 1}, U
			case L: // -> 4
				next, d = Point{p.Y - sz, 2 * sz}, D
			default:
				panic("should not happen")
			}
		case 4:
			switch dir {
			case L: // -> 1
				next, d = Point{sz, 3*sz - 1 - p.Y}, R
			case U: // -> 3
				next, d = Point{sz, p.X + sz}, R
			default:
				panic("should not happen")
			}
		case 5:
			switch dir {
			case R: // -> 2
				next, d = Point{3*sz - 1, 3*sz - 1 - p.Y}, L
			case D: // -> 6
				next, d = Point{sz - 1, p.X + 2*sz}, L
			default:
				panic("should not happen")
			}
		case 6:
			switch dir {
			case R: // -> 5
				next, d = Point{p.Y - 2*sz, 3*sz - 1}, U
			case D: // -> 2
				next, d = Point{p.X + 2*sz, 0}, D
			case L: // -> 1
				next, d = Point{p.Y - 2*sz, 0}, D
			default:
				panic("should not happen")
			}
		default:
			panic("unreachable")
		}

		return
	})
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
