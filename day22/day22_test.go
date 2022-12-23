package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var rawInput = `
        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

10R5L5R10L4R5L5
`

var input = strings.Split(rawInput, "\n")[1:]

func TestA(t *testing.T) {
	assert.Equal(t, 6032, solveA(input))
}

func TestB(t *testing.T) {
	sz := 4
	assert.Equal(t, 5031, solve(input, func(p Point, dir direction, m Map) (next Point, d direction) {
		faces := [][]int{{0, 0, 1, 0}, {2, 3, 4, 0}, {0, 0, 5, 6}}
		f := faces[p.Y/sz][p.X/sz]

		switch f {
		case 1:
			switch dir {
			case R: // -> 6
				next, d = Point{4*sz - 1, 3*sz - p.Y - 1}, L
			case L: // -> 3
				next, d = Point{p.Y + sz, sz}, D
			case U: // -> 2
				next, d = Point{3*sz - p.X - 1, sz}, D
			default:
				panic("should not happen")
			}
		case 2:
			switch dir {
			case D: // -> 5
				next, d = Point{3*sz - p.X - 1, 3*sz - 1}, U
			case L: // -> 6
				next, d = Point{5*sz - p.Y - 1, 3*sz - 1}, U
			case U: // -> 1
				next, d = Point{3*sz - p.X - 1, 0}, D
			default:
				panic("should not happen")
			}
		case 3:
			switch dir {
			case D: // -> 5
				next, d = Point{2*sz - 1, 4*sz - p.X - 1}, R
			case U: // -> 1
				next, d = Point{2 * sz, p.X - sz}, R
			default:
				panic("should not happen")
			}
		case 4:
			switch dir {
			case R: // -> 6
				next, d = Point{5*sz - p.Y - 1, 2 * sz}, D
			default:
				panic("should not happen")
			}
		case 5:
			switch dir {
			case D: // -> 2
				next, d = Point{3*sz - p.X - 1, 2*sz - 1}, U
			case L: // -> 3
				next, d = Point{4*sz - p.Y - 1, 2*sz - 1}, U
			default:
				panic("should not happen")
			}
		case 6:
			switch dir {
			case R: // -> 1
				next, d = Point{3*sz - 1, 3*sz - p.Y - 1}, L
			case D: // -> 2
				next, d = Point{0, 5*sz - p.X - 1}, R
			case U: // -> 4
				next, d = Point{3*sz - 1, 5*sz - p.X - 1}, L
			default:
				panic("should not happen")
			}
		default:
			panic("unreachable")
		}

		return
	}))
}
