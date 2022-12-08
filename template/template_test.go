package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var rawInput = ``

var input = strings.Split(rawInput, "\n")[1:]

func TestA(t *testing.T) {
	assert.Equal(t, 0, solveA(input))
}

func TestB(t *testing.T) {
	assert.Equal(t, 0, solveB(input))
}
