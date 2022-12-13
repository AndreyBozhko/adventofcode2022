package main

import (
	"adventofcode2022/utils"
	"sort"
	"strconv"
	"strings"
)

type ordering int8

const (
	Lt ordering = -1
	Eq ordering = 0
	Gt ordering = 1
)

type Entry struct {
	Int int
	Lst []*Entry
}

func ParseEntry(s string) *Entry {
	if val, err := strconv.Atoi(s); err == nil {
		return &Entry{Int: val}
	}

	tmp := s[1 : len(s)-1]
	if len(tmp) == 0 {
		return &Entry{Lst: make([]*Entry, 0)}
	}

	e := &Entry{}

	var idx int

	for len(tmp) > 0 {
		if tmp[0] >= '0' && tmp[0] <= '9' {
			// number
			idx = strings.Index(tmp, ",")
		} else {
			// list
			cnt := 1
			for j := 1; j < len(tmp); j++ {
				if tmp[j] == '[' {
					cnt++
				}
				if tmp[j] == ']' {
					cnt--
				}
				if cnt == 0 {
					idx = j + 1
					break
				}
			}
		}

		if idx == -1 || idx == len(tmp) {
			e.Lst = append(e.Lst, ParseEntry(tmp))
			tmp = tmp[:0]
		} else {
			e.Lst = append(e.Lst, ParseEntry(tmp[:idx]))
			tmp = tmp[idx+1:]
		}
	}

	return e
}

func Compare(left, right *Entry) ordering {
	// both integers
	if left.Lst == nil && right.Lst == nil {
		if left.Int < right.Int {
			return Lt
		}
		if left.Int > right.Int {
			return Gt
		}
		return Eq
	}

	// one integer, one list
	if left.Lst != nil && right.Lst == nil {
		r := &Entry{Lst: []*Entry{{Int: right.Int}}}
		return Compare(left, r)
	}

	if left.Lst == nil && right.Lst != nil {
		l := &Entry{Lst: []*Entry{{Int: left.Int}}}
		return Compare(l, right)
	}

	// both lists
	for idx := 0; idx < len(left.Lst) && idx < len(right.Lst); idx++ {
		if res := Compare(left.Lst[idx], right.Lst[idx]); res != Eq {
			return res
		}
	}

	if len(left.Lst) < len(right.Lst) {
		return Lt
	}
	if len(left.Lst) > len(right.Lst) {
		return Gt
	}
	return Eq
}

func solveA(input []string) int {
	total := 0

	for i := 0; i < len(input); i += 3 {
		left, right := ParseEntry(input[i]), ParseEntry(input[i+1])

		if Compare(left, right) == Lt {
			total += i/3 + 1
		}
	}

	return total
}

func solveB(input []string) int {
	total := 1

	marker1, marker2 := ParseEntry("[[2]]"), ParseEntry("[[6]]")
	arr := []*Entry{marker1, marker2}

	for _, line := range input {
		if line == "" {
			continue
		}

		arr = append(arr, ParseEntry(line))
	}

	sort.Slice(arr, func(i, j int) bool {
		return Compare(arr[i], arr[j]) == Lt
	})

	for idx, e := range arr {
		if e == marker1 || e == marker2 {
			total *= idx + 1
		}
	}

	return total
}

func main() {
	problem := utils.Must(utils.NewProblem(2022, 13))
	input := utils.Must(problem.LoadInput())

	var result int
	if problem.IsA() {
		result = solveA(input)
	}

	if problem.IsB() {
		result = solveB(input)
	}

	_ = problem.WriteOutput(strconv.Itoa(result))
}
