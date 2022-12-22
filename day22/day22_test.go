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
	assert.Equal(t, 5031, solveB(input))
}
