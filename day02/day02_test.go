package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const rawInput = `
A Y
B X
C Z
`

var input = strings.Split(rawInput, "\n")

func TestA(t *testing.T) {
	assert.Equal(t, 15, solveA(input))
}

func TestB(t *testing.T) {
	assert.Equal(t, 12, solveB(input))
}
