package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const rawInput = `
vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`

var input = strings.Split(rawInput, "\n")[1:]

func TestA(t *testing.T) {
	assert.Equal(t, 157, solveA(input))
}

func TestB(t *testing.T) {
	assert.Equal(t, 70, solveB(input))
}
