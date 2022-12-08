package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var rawInput = `
30373
25512
65332
33549
35390
`

var input = strings.Split(rawInput, "\n")[1:]

func TestA(t *testing.T) {
	assert.Equal(t, 21, solveA(input))
}

func TestB(t *testing.T) {
	assert.Equal(t, 8, solveB(input))
}
