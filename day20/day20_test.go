package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var rawInput = `
1
2
-3
3
-2
0
4
`

var input = strings.Split(rawInput, "\n")[1:]

func TestA(t *testing.T) {
	assert.Equal(t, 3, solveA(input))
}

func TestB(t *testing.T) {
	assert.Equal(t, 1623178306, solveB(input))
}
