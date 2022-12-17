package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `>>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>`

func TestA(t *testing.T) {
	assert.Equal(t, 3068, solveA(input))
}

func TestB(t *testing.T) {
	assert.Equal(t, 1_514_285_714_288, solveB(input))
}
