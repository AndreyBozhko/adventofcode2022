package main

import (
	"adventofcode2022/utils"
	"container/heap"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	X, Y, Steps int
}

type Heap []Point

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].Steps < h[j].Steps }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x any) {
	*h = append(*h, x.(Point))
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func solve(lines []string, queue Heap) int {
	input := make([]string, 0, len(lines))
	input = append(input, lines...)

	visited := make(map[Point]bool)

	var destX, destY int

	steps := make([][]int, len(input))
	for x, line := range input {
		steps[x] = make([]int, len(line))
		for idx := 0; idx < len(line); idx++ {
			steps[x][idx] = math.MaxInt32
		}

		if y := strings.Index(line, "S"); y != -1 {
			steps[x][y] = 0
			input[x] = strings.ReplaceAll(input[x], "S", "a")
		}

		if y := strings.Index(line, "E"); y != -1 {
			destX, destY = x, y
			input[x] = strings.ReplaceAll(input[x], "E", "z")
		}
	}

	maxX, maxY := len(input), len(input[0])

	neighbors := func(point Point) (result []Point) {
		result = make([]Point, 0, 4)
		if point.X > 0 {
			result = append(result, Point{X: point.X - 1, Y: point.Y})
		}
		if point.Y > 0 {
			result = append(result, Point{X: point.X, Y: point.Y - 1})
		}
		if point.X < maxX-1 {
			result = append(result, Point{X: point.X + 1, Y: point.Y})
		}
		if point.Y < maxY-1 {
			result = append(result, Point{X: point.X, Y: point.Y + 1})
		}

		return
	}

	for len(queue) > 0 {

		p := heap.Pop(&queue).(Point)

		if _, ok := visited[Point{X: p.X, Y: p.Y}]; !ok {
			steps[p.X][p.Y] = p.Steps
			visited[Point{X: p.X, Y: p.Y}] = true
		} else {
			continue
		}

		if p.X == destX && p.Y == destY {
			return p.Steps
		}

		for _, n := range neighbors(p) {
			stepCondition := steps[n.X][n.Y] > p.Steps+1

			to, from := input[n.X][n.Y], input[p.X][p.Y]
			heightCondition := to <= from+1

			_, marked := visited[n]

			if stepCondition && heightCondition && !marked {
				heap.Push(&queue, Point{X: n.X, Y: n.Y, Steps: p.Steps + 1})
			}
		}
	}

	panic("unreachable")
}

func solveA(input []string) int {
	queue := Heap(make([]Point, 0))
	for x, line := range input {
		if y := strings.Index(line, "S"); y != -1 {
			heap.Push(&queue, Point{X: x, Y: y, Steps: 0})
		}
	}

	return solve(input, queue)
}

func solveB(input []string) int {
	queue := Heap(make([]Point, 0))
	for x, line := range input {
		for y, c := range line {
			if c == 'a' || c == 'S' {
				heap.Push(&queue, Point{X: x, Y: y, Steps: 0})
			}
		}
	}
	return solve(input, queue)
}

func main() {
	problem := utils.Must(utils.NewProblem(2022, 12))
	input := utils.Must(problem.LoadInput())

	if input[len(input)-1] == "" {
		input = input[:len(input)-1]
	}

	var result int
	if problem.IsA() {
		result = solveA(input)
	}

	if problem.IsB() {
		result = solveB(input)
	}

	_ = problem.WriteOutput(strconv.Itoa(result))
}
