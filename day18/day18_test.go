package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var rawInput = `
2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5
`

var input = strings.Split(rawInput, "\n")[1:]

func TestA(t *testing.T) {
	assert.Equal(t, 64, solveA(input))
}

func TestB(t *testing.T) {
	assert.Equal(t, 58, solveB(input))
}
