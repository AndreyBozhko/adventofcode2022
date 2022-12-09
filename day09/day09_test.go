package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var rawInput1 = `
R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2
`

var input1 = strings.Split(rawInput1, "\n")

func TestA1(t *testing.T) {
	assert.Equal(t, 13, solveA(input1))
}

func TestB1(t *testing.T) {
	assert.Equal(t, 1, solveB(input1))
}

var rawInput2 = `
R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20
`

var input2 = strings.Split(rawInput2, "\n")

func TestA2(t *testing.T) {
	assert.Equal(t, 88, solveA(input2))
}

func TestB2(t *testing.T) {
	assert.Equal(t, 36, solveB(input2))
}
