package main

import (
	"adventofcode2022/utils"
	"fmt"
	"strconv"
)

type Cube struct {
	X, Y, Z int
}

func parseCubes(input []string) map[Cube]bool {
	cubes := make(map[Cube]bool)

	for _, line := range input {
		if line == "" {
			continue
		}

		var c Cube
		_ = utils.Must(fmt.Sscanf(line, "%d,%d,%d", &c.X, &c.Y, &c.Z))
		cubes[c] = true
	}

	return cubes
}

func surfaceArea(cubes map[Cube]bool) int {
	total := 0
	for c := range cubes {
		for _, delta := range []Cube{
			{-1, 0, 0},
			{+1, 0, 0},
			{0, -1, 0},
			{0, +1, 0},
			{0, 0, -1},
			{0, 0, +1},
		} {
			if !cubes[Cube{c.X + delta.X, c.Y + delta.Y, c.Z + delta.Z}] {
				total++
			}
		}
	}

	return total
}

func solveA(input []string) int {
	lava := parseCubes(input)
	return surfaceArea(lava)
}

func solveB(input []string) int {
	lava := parseCubes(input)

	max := Cube{}
	for c := range lava {
		if c.X > max.X {
			max.X = c.X
		}
		if c.Y > max.Y {
			max.Y = c.Y
		}
		if c.Z > max.Z {
			max.Z = c.Z
		}
	}

	air := make(map[Cube]bool)
	for x := -1; x <= max.X+1; x++ {
		for y := -1; y <= max.Y+1; y++ {
			for z := -1; z <= max.Z+1; z++ {
				if !lava[Cube{x, y, z}] {
					air[Cube{x, y, z}] = true
				}
			}
		}
	}

	queue1, queue2 := []Cube{max}, make([]Cube, 0)
	delete(air, max)

	for len(queue1) > 0 {
		for _, c := range queue1 {
			for _, delta := range []Cube{
				{-1, 0, 0},
				{+1, 0, 0},
				{0, -1, 0},
				{0, +1, 0},
				{0, 0, -1},
				{0, 0, +1},
			} {
				n := Cube{c.X + delta.X, c.Y + delta.Y, c.Z + delta.Z}
				if air[n] {
					delete(air, n)
					queue2 = append(queue2, n)
				}
			}
		}

		queue1, queue2 = queue2, make([]Cube, 0)
	}

	return surfaceArea(lava) - surfaceArea(air)
}

func main() {
	problem := utils.Must(utils.NewProblem(2022, 18))
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
