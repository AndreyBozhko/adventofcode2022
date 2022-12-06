package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testcases = []struct {
	data string
	ansA int
	ansB int
}{
	{`mjqjpqmgbljsphdztnvjfqwrcgsmlb`, 7, 19},
	{`bvwbjplbgvbhsrlpgdmjqwftvncz`, 5, 23},
	{`nppdvjthqldpwncqszvftbrmjlhg`, 6, 23},
	{`nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg`, 10, 29},
	{`zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`, 11, 26},
}

func TestA(t *testing.T) {
	for _, test := range testcases {
		assert.Equal(t, test.ansA, solveA(test.data))
	}
}

func TestB(t *testing.T) {
	for _, test := range testcases {
		assert.Equal(t, test.ansB, solveB(test.data))
	}
}
