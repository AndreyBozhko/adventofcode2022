package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var rawInput = `
498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9
`

var input = strings.Split(rawInput, "\n")[1:]

func TestA(t *testing.T) {
	assert.Equal(t, 24, solveA(input))
}

func TestB(t *testing.T) {
	assert.Equal(t, 93, solveB(input))
}
