package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var rawInput = `
Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi
`

var input = strings.Split(strings.Trim(rawInput, "\n"), "\n")

func TestA(t *testing.T) {
	assert.Equal(t, 31, solveA(input))
}

func TestB(t *testing.T) {
	assert.Equal(t, 29, solveB(input))
}
