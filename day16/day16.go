package main

import (
	"adventofcode2022/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Progress struct {
	elapsed  int
	previous string
}

func parseValves(input []string) (usefulValves map[string]int, dist map[string]map[string]int) {
	valves := make(map[string]struct {
		flow      int
		neighbors []string
	})

	for _, line := range input {
		if line == "" {
			continue
		}

		var flow int
		var name string

		line := strings.ReplaceAll(line, "valves", "valve")
		part1, part2 := strings.Split(line, "; ")[0], strings.Split(line, "; ")[1]
		part2 = strings.Split(part2, "valve ")[1]
		neighbors := strings.Split(part2, ", ")

		_ = utils.Must(fmt.Sscanf(part1, "Valve %s has flow rate=%d", &name, &flow))

		valves[name] = struct {
			flow      int
			neighbors []string
		}{flow, neighbors}

	}

	dist = make(map[string]map[string]int)
	for name, v := range valves {
		dist[name] = make(map[string]int)
		for nm := range valves {
			dist[name][nm] = math.MaxInt32
		}

		dist[name][name] = 0

		for _, to := range v.neighbors {
			dist[name][to] = 1
		}
	}

	for k := range valves {
		for i := range valves {
			for j := range valves {
				if dist[i][j] > dist[i][k]+dist[k][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	usefulValves = make(map[string]int)
	for nm, v := range valves {
		if v.flow > 0 {
			usefulValves[nm] = v.flow
		}
	}

	return
}

func solveA(input []string) int {
	totalMinutes := 30
	usefulValves, dist := parseValves(input)

	seen := make(map[string]bool)
	seen["AA"] = true

	total := 0

	var dfs func(released int, progress Progress)

	dfs = func(released int, progress Progress) {
		if total < released {
			total = released
		}
		for v, flow := range usefulValves {
			if seen[v] {
				continue
			}
			seen[v] = true

			newMinutes := progress.elapsed + dist[progress.previous][v] + 1
			if newMinutes < totalMinutes {
				dfs(released+flow*(totalMinutes-newMinutes), Progress{newMinutes, v})
			}

			delete(seen, v)
		}
	}

	dfs(0, Progress{0, "AA"})

	return total
}

func solveB(input []string) int {
	totalMinutes := 26
	usefulValves, dist := parseValves(input)

	seen := make(map[string]bool)
	seen["AA"] = true

	total := 0

	var dfs func(released int, progress1, progress2 Progress)

	dfs = func(released int, progress1, progress2 Progress) {
		if total < released {
			total = released
		}
		for v, flow := range usefulValves {
			if seen[v] {
				continue
			}
			seen[v] = true

			newMinutes := progress1.elapsed + dist[progress1.previous][v] + 1
			if newMinutes < totalMinutes {
				dfs(released+flow*(totalMinutes-newMinutes), progress2, Progress{newMinutes, v})
			}

			delete(seen, v)
		}
	}

	dfs(0, Progress{0, "AA"}, Progress{0, "AA"})

	return total
}

func main() {
	problem := utils.Must(utils.NewProblem(2022, 16))
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
