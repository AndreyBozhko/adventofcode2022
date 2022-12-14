package main

import (
	"adventofcode2022/utils"
	"fmt"
	"strconv"
	"strings"
)

func addSandUnit(grid [][]bool, startX, startY int, abyssY int) (x, y int, rest bool) {
	x, y = startX, startY
	for y <= abyssY {
		if !grid[x][y+1] {
			y++
		} else if !grid[x-1][y+1] {
			x--
			y++
		} else if !grid[x+1][y+1] {
			x++
			y++
		} else {
			grid[x][y] = true
			return x, y, true
		}
	}
	return x, y, false
}

func parseGrid(input []string) (grid [][]bool, abyssY int) {
	grid = make([][]bool, 1000)
	for i := range grid {
		grid[i] = make([]bool, 1000)
	}

	for _, line := range input {
		if line == "" {
			continue
		}

		pts := strings.Split(line, " -> ")
		var x, y int
		_ = utils.Must(fmt.Sscanf(pts[0], "%d,%d", &x, &y))
		if y > abyssY {
			abyssY = y
		}

		grid[x][y] = true

		for i := 1; i < len(pts); i++ {
			var newX, newY int
			_ = utils.Must(fmt.Sscanf(pts[i], "%d,%d", &newX, &newY))
			if newY > abyssY {
				abyssY = newY
			}

			dx, dy := utils.Signum(newX-x), utils.Signum(newY-y)
			for x != newX || y != newY {
				x += dx
				y += dy
				grid[x][y] = true
			}
		}
	}

	return grid, abyssY
}

func solveA(input []string) int {

	grid, abyssY := parseGrid(input)

	for total := 0; ; total++ {
		if _, _, rest := addSandUnit(grid, 500, 0, abyssY); !rest {
			return total
		}
	}
}

func solveB(input []string) int {
	grid, abyssY := parseGrid(input)

	for i := range grid {
		grid[i][abyssY+2] = true
	}

	for total := 1; ; total++ {
		if x, y, rest := addSandUnit(grid, 500, 0, abyssY+2); rest && x == 500 && y == 0 {
			return total
		}
	}
}

func main() {
	problem := utils.Must(utils.NewProblem(2022, 14))
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
