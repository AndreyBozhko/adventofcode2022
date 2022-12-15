package main

import (
	"adventofcode2022/utils"
	"fmt"
	"math"
	"sort"
	"strconv"
)

type Range struct {
	// from - inclusive
	// to   - exclusive
	from, to int
}

func (r *Range) Len() int {
	return r.to - r.from
}

type Point struct {
	X, Y int
}

func (p *Point) Manhattan(other *Point) int {
	return utils.Abs(p.X-other.X) + utils.Abs(p.Y-other.Y)
}

type Sensor struct {
	location Point
	beacon   Point
}

func parseInput(input []string) []Sensor {
	sensors := make([]Sensor, 0, len(input))

	for _, line := range input {
		if line == "" {
			continue
		}

		var s Sensor
		_ = utils.Must(
			fmt.Sscanf(line,
				"Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
				&s.location.X, &s.location.Y, &s.beacon.X, &s.beacon.Y,
			),
		)
		sensors = append(sensors, s)
	}

	return sensors
}

func findRanges(row int, minX, maxX int, sensors []Sensor) []Range {
	ranges := make([]Range, 0)

	for _, s := range sensors {
		dist := s.location.Manhattan(&s.beacon)
		dist = dist - utils.Abs(s.location.Y-row)
		if dist < 0 {
			continue
		}

		x1, x2 := s.location.X-dist, s.location.X+dist
		if x1 < minX {
			x1 = minX
		}
		if x2 > maxX {
			x2 = maxX
		}

		ranges = append(ranges, Range{x1, x2 + 1})
	}

	if len(ranges) == 0 {
		return nil
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].from < ranges[j].from
	})

	merged := []Range{ranges[0]}

	for _, r := range ranges {
		last := &merged[len(merged)-1]
		if r.from <= last.to && r.to > last.to {
			last.to = r.to
		}
		if r.from > last.to {
			merged = append(merged, r)
		}
	}

	return merged
}

func solveA(input []string, maxSize int) int {
	sensors := parseInput(input)
	ranges := findRanges(maxSize/2, math.MinInt32, math.MaxInt32, sensors)

	excludeBeaconsX := make(map[int]bool)
	for _, s := range sensors {
		if s.beacon.Y == maxSize/2 {
			excludeBeaconsX[s.beacon.X] = true
		}
	}

	occupied := 0
	for _, r := range ranges {
		occupied += r.Len()
		for b := range excludeBeaconsX {
			if b >= r.from && b < r.to {
				occupied--
			}
		}
	}

	return occupied
}

func solveB(input []string, maxSize int) int {
	sensors := parseInput(input)

	x, y := 0, 0
	for y = 0; y <= maxSize; y++ {
		ranges := findRanges(y, 0, maxSize, sensors)
		if len(ranges) == 1 && ranges[0].Len() > maxSize {
			continue
		}
		if len(ranges) != 2 || ranges[0].Len()+ranges[1].Len() != maxSize {
			panic("should not happen")
		}
		x = ranges[0].to
		break
	}

	return x*4000000 + y
}

func main() {
	problem := utils.Must(utils.NewProblem(2022, 15))
	input := utils.Must(problem.LoadInput())

	var result int
	if problem.IsA() {
		result = solveA(input, 4000000)
	}

	if problem.IsB() {
		result = solveB(input, 4000000)
	}

	_ = problem.WriteOutput(strconv.Itoa(result))
}
