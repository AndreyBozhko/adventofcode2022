package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var rawInput = `
[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]
`

var input = strings.Split(rawInput, "\n")[1:]

func TestA(t *testing.T) {
	assert.Equal(t, 13, solveA(input))
}

func TestB(t *testing.T) {
	assert.Equal(t, 140, solveB(input))
}
