package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var rawInput = `
....#..
..###.#
#...#.#
.#...##
#.###..
##.#.##
.#..#..
`

var input = strings.Split(rawInput, "\n")[1:]

func TestA(t *testing.T) {
	assert.Equal(t, 110, solveA(input))
}

func TestB(t *testing.T) {
	assert.Equal(t, 20, solveB(input))
}
