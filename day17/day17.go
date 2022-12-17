package main

import (
	"adventofcode2022/utils"
	"strconv"
	"strings"
)

type Jet struct {
	dirs string
	idx  int
}

func NewJet(input string) Jet {
	return Jet{input, -1}
}

func (j *Jet) Next() byte {
	j.idx = (j.idx + 1) % len(j.dirs)
	return j.dirs[j.idx]
}

type Point struct {
	X, Y int
}

type Rock [][]byte

func NewRock(s string) Rock {
	lines := strings.Split(s, "\n")
	r := make([][]byte, 0)
	for i := len(lines) - 1; i >= 0; i-- {
		r = append(r, []byte(lines[i]))
	}
	return Rock(r)
}

func (r Rock) Overlaps(x, y int, grid map[Point]bool) bool {
	for j, row := range r {
		for i, ch := range row {
			if ch == '#' && grid[Point{x + i, y + j}] {
				return true
			}
		}
	}
	return false
}

func Snapshot(grid map[Point]bool, height int, limit int) string {
	res := make([]string, 0, 2)

	buf := make([]byte, 9)
	buf[0] = '|'
	buf[8] = '|'

	for y := height; y >= 0; y-- {
		for x := 0; x < 7; x++ {
			if grid[Point{x, y}] {
				buf[x+1] = '#'
			} else {
				buf[x+1] = '.'
			}
		}

		res = append(res, string(buf))

		limit--
		if limit <= 0 {
			break
		}
	}
	return strings.Join(res, "\n")
}

func (r Rock) Move(x, y int, grid map[Point]bool, dir byte) (newX, newY int, stopped bool) {
	switch dir {
	case 'v':
		y--
		if y < 0 || r.Overlaps(x, y, grid) {
			return x, y + 1, true
		}
		return x, y, false

	case '<':
		x--
		if x < 0 || r.Overlaps(x, y, grid) {
			return x + 1, y, false
		}
		return x, y, false

	case '>':
		x++
		width := len(r[0])
		if x+width > 7 || r.Overlaps(x, y, grid) {
			return x - 1, y, false
		}
		return x, y, false

	default:
		panic("should not happen")
	}
}

var rocks = []Rock{
	NewRock("####"),
	NewRock(".#.\n###\n.#."),
	NewRock("..#\n..#\n###"),
	NewRock("#\n#\n#\n#"),
	NewRock("##\n##"),
}

type Info struct {
	rock    int
	jet     int
	tworows string
}

func solve(input string, target int) int {
	grid := make(map[Point]bool)
	jet := NewJet(input)
	height := 0

	heights := make([]int, 0)

	seen := make(map[Info]int)

	for r := 0; r < target; r++ {
		rock := rocks[r%len(rocks)]
		x, y := 2, height+3
		var stopped bool

		for {
			if x, y, stopped = rock.Move(x, y, grid, jet.Next()); stopped {
				break
			}
			if x, y, stopped = rock.Move(x, y, grid, 'v'); stopped {
				break
			}
		}

		for j, row := range rock {
			for i, ch := range row {
				if ch == '#' {
					grid[Point{x + i, y + j}] = true
					if y+j+1 > height {
						height = y + j + 1
					}
				}
			}
		}

		heights = append(heights, height)

		info := Info{rock: r % len(rocks), jet: jet.idx, tworows: Snapshot(grid, height-1, 2)}
		if val, ok := seen[info]; ok {
			rockDiff := r - val
			heightDiff := height - heights[val]
			extra := target % rockDiff

			return (target-extra)/rockDiff*heightDiff + heights[extra-1]
		}
		seen[info] = r
	}

	return height
}

func solveA(input string) int {
	return solve(input, 2022)
}

func solveB(input string) int {
	return solve(input, 1_000_000_000_000)
}

func main() {
	problem := utils.Must(utils.NewProblem(2022, 17))
	input := utils.Must(problem.LoadInput())[0]

	var result int
	if problem.IsA() {
		result = solveA(input)
	}

	if problem.IsB() {
		result = solveB(input)
	}

	_ = problem.WriteOutput(strconv.Itoa(result))
}
