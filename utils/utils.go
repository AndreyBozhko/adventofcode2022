package utils

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// Arbitrary is a generic type to use with the Must function.
type Arbitrary interface {
	~int | ~string | ~[]int | ~[]string | ~rune | *Problem
}

// Must returns the value if error is nil,
// and panics if error is not nil.
func Must[T Arbitrary](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

// Problem represents AoC problem part
// for a given year and day.
type Problem struct {
	year, day int
	part      string
}

func NewProblem(year, day int) (*Problem, error) {
	args := os.Args
	if len(args) < 2 {
		return nil, errors.New("problem part is not provided via os.Args")
	}
	part := args[1]
	if part == "A" || part == "B" {
		return &Problem{year, day, part}, nil
	}

	return nil, fmt.Errorf("invalid part: %s", part)
}

func (p *Problem) IsA() bool {
	return p.part == "A"
}

func (p *Problem) IsB() bool {
	return p.part == "B"
}

// LoadInput loads the problem input from file into a string slice.
func (p *Problem) LoadInput() (lines []string, err error) {
	path := fmt.Sprintf("inputs/%d/input%02d.txt", p.year, p.day)

	bts, err := os.ReadFile(path)
	if err != nil {
		return
	}

	lines = strings.Split(string(bts), "\n")

	return
}

// WriteOutput prints the result to stdout and to '.answer' file.
func (p *Problem) WriteOutput(result string) error {
	fmt.Printf("Part %s: %s\n", p.part, result)

	content := p.part + " " + result
	return os.WriteFile(".answer", []byte(content), 0644)
}
