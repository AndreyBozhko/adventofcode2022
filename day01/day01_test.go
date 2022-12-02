package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const rawInput = `
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`

var input = strings.Split(rawInput, "\n")

func TestA(t *testing.T) {
	assert.Equal(t, 24000, solveA(input))
}

func TestB(t *testing.T) {
	assert.Equal(t, 45000, solveB(input))
}
