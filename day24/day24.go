package main

import (
	"adventofcode2022/utils"
	"strconv"
)

type direction int

const (
	R direction = 0
	D direction = 1
	L direction = 2
	U direction = 3
)

type Point struct {
	X, Y int
}

func (v *Valley) Move(p Point, dir direction, steps int) (dest Point) {
	d := directions[dir]
	dest.X = (p.X + steps*d.X) % v.Width
	for dest.X < 0 {
		dest.X += v.Width
	}

	dest.Y = (p.Y + steps*d.Y) % v.Height
	for dest.Y < 0 {
		dest.Y += v.Height
	}
	return
}

var directions = map[direction]Point{
	R: {1, 0}, D: {0, 1}, L: {-1, 0}, U: {0, -1},
}

type Valley struct {
	blizzards     map[Point]direction
	Height, Width int
	safeSpots     []map[Point]bool
}

func parseInput(input []string) Valley {
	blizzards := make(map[Point]direction)

	for y, line := range input {
		for x, c := range line {
			switch c {
			case '>':
				blizzards[Point{x - 1, y - 1}] = R
			case 'v':
				blizzards[Point{x - 1, y - 1}] = D
			case '<':
				blizzards[Point{x - 1, y - 1}] = L
			case '^':
				blizzards[Point{x - 1, y - 1}] = U
			}
		}
	}

	v := Valley{blizzards, len(input) - 3, len(input[0]) - 2, nil}
	v.PredictSafeSpots()

	return v
}

func (v *Valley) PredictSafeSpots() {
	iterations := utils.LCM(v.Height, v.Width)
	v.safeSpots = make([]map[Point]bool, iterations)

	for i := 0; i < len(v.safeSpots); i++ {
		v.safeSpots[i] = make(map[Point]bool)

		for x := 0; x < v.Width; x++ {
			for y := 0; y < v.Height; y++ {
				free := true
				for pos, d := range v.blizzards {
					blizPos := v.Move(pos, d, i)
					if blizPos.X == x && blizPos.Y == y {
						free = false
						break
					}
				}
				if free {
					v.safeSpots[i][Point{x, y}] = true
				}
			}
		}
	}

}

func (v *Valley) Neighbors(p Point) []Point {
	res := make([]Point, 0, 5)
	for _, d := range directions {
		q := Point{p.X + d.X, p.Y + d.Y}
		if q.X >= 0 && q.X < v.Width && q.Y >= 0 && q.Y < v.Height {
			res = append(res, q)
		}
	}
	res = append(res, p)
	return res
}

func bfs(start, end Point, clock int, valley Valley) int {
	queue1, queue2 := make(map[Point]bool), make(map[Point]bool)
	queue1[start] = true

	for minute := clock; ; minute++ {
		for cur := range queue1 {
			if cur.X == end.X && (cur.Y+1 == end.Y || cur.Y-1 == end.Y) {
				return minute
			}

			for _, move := range valley.Neighbors(cur) {
				if move.X == start.X && move.Y == start.Y {
					queue2[move] = true
				}
				if move.X == end.X && move.Y == end.Y {
					queue2[move] = true
				}

				cycle := minute % len(valley.safeSpots)
				safeSpots := valley.safeSpots[cycle]
				if safeSpots[move] {
					queue2[move] = true
				}
			}
		}

		queue1, queue2 = queue2, make(map[Point]bool)
	}
}

func solveA(input []string) int {
	valley := parseInput(input)

	start, finish := Point{0, -1}, Point{valley.Width - 1, valley.Height}

	ans := bfs(start, finish, 1, valley)

	return ans
}

func solveB(input []string) int {
	valley := parseInput(input)

	start, finish := Point{0, -1}, Point{valley.Width - 1, valley.Height}

	ans := 1
	ans = bfs(start, finish, ans, valley)
	ans = bfs(finish, start, ans, valley)
	ans = bfs(start, finish, ans, valley)

	return ans
}

func main() {
	problem := utils.Must(utils.NewProblem(2022, 24))
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
