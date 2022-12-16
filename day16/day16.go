package main

import (
	"adventofcode2022/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Valve struct {
	flow      int
	neighbors []string
}

func solveA(input []string) int {
	valves := make(map[string]Valve)

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

		valves[name] = Valve{flow, neighbors}

	}

	dist := make(map[string]map[string]int)
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

	usefulValves := make([]string, 0)
	for nm, v := range valves {
		if v.flow > 0 {
			usefulValves = append(usefulValves, nm)
		}
	}

	seen := make(map[string]bool)
	seen["AA"] = true

	total := 0
	var dfs func(released int, minutes int, prev string)

	dfs = func(released int, minutes int, prev string) {
		if total < released {
			total = released
		}
		for _, v := range usefulValves {
			if seen[v] {
				continue
			}
			seen[v] = true

			newMinutes := minutes + dist[prev][v] + 1
			if newMinutes < 30 {
				dfs(released+valves[v].flow*(30-newMinutes), newMinutes, v)
			}

			delete(seen, v)
		}
	}

	dfs(0, 0, "AA")

	return total
}

func solveB(input []string) int {
	valves := make(map[string]Valve)

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

		valves[name] = Valve{flow, neighbors}

	}

	dist := make(map[string]map[string]int)
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

	usefulValves := make([]string, 0)
	for nm, v := range valves {
		if v.flow > 0 {
			usefulValves = append(usefulValves, nm)
		}
	}

	seen := make(map[string]bool)
	seen["AA"] = true

	total := 0
	var dfs func(released int, minutes int, prev string, minutes2 int, prev2 string)

	dfs = func(released int, minutes int, prev string, minutes2 int, prev2 string) {
		if total < released {
			total = released
		}
		for _, v := range usefulValves {
			if seen[v] {
				continue
			}
			seen[v] = true

			newMinutes := minutes + dist[prev][v] + 1
			if newMinutes < 26 {
				dfs(released+valves[v].flow*(26-newMinutes), minutes2, prev2, newMinutes, v)
			}

			delete(seen, v)
		}
	}

	dfs(0, 0, "AA", 0, "AA")

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
