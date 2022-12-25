package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var rawInput = `
1=-0-2
12111
2=0=
21
2=01
111
20012
112
1=-1=
1-12
12
1=
122
`

var input = strings.Split(rawInput, "\n")[1:]

func TestA(t *testing.T) {
	assert.Equal(t, "2=-1=0", solveA(input))
}
