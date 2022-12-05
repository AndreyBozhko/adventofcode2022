package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const rawInput = `
    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`

var input = strings.Split(rawInput, "\n")[1:]

func TestA(t *testing.T) {
	assert.Equal(t, "CMZ", solveA(input))
}

func TestB(t *testing.T) {
	assert.Equal(t, "MCD", solveB(input))
}
