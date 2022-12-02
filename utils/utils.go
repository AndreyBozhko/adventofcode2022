package utils

import (
	"fmt"
	"os"
	"strings"
)

// Part is a problem part (A or B).
type Part string

const (
	A Part = "A"
	B Part = "B"
)

func (p Part) IsA() bool {
	return p == A
}

func (p Part) IsB() bool {
	return p == B
}

func ProblemPart() Part {
	args := os.Args
	if len(args) < 2 {
		panic("Problem part is not provided")
	}
	part := Part(args[1])
	if part == A || part == B {
		return part
	}

	panic("Invalid part: " + part)
}

// Arbitrary is a generic type to use with the Must function.
type Arbitrary interface {
	~int | ~string | ~[]int | ~[]string
}

// Must returns the value if error is nil,
// and panics if error is not nil.
func Must[T Arbitrary](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

// Input loads the problem input for a specified day
// from file as a string slice.
func Input(year, day int) (lines []string, err error) {
	path := fmt.Sprintf("inputs/%d/input%02d.txt", year, day)

	bts, err := os.ReadFile(path)
	if err != nil {
		return
	}

	lines = strings.Split(string(bts), "\n")

	return
}

// Output prints the result to stdout and to '.answer' file.
func Output(part Part, result string) {
	fmt.Printf("Part %s: %s\n", part, result)

	content := string(part) + " " + result
	if err := os.WriteFile(".answer", []byte(content), 0644); err != nil {
		panic(err)
	}
}
