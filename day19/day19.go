package main

import (
	"adventofcode2022/utils"
	"fmt"
	"math"
	"strconv"
)

const (
	ore      = 0
	clay     = 1
	obsidian = 2
	geode    = 3
)

type cost [4]int

type blueprint struct {
	robotCosts [4]cost
}

func (resourcesNow *cost) NextAvail(robotCost cost, numRobots [4]int) (resourcesAfter cost, waitTime int) {
	resourcesAfter, waitTime = *resourcesNow, 0

	for material := 0; material < 3; material++ {
		if resourcesAfter[material] < robotCost[material] && numRobots[material] == 0 {
			return cost{}, math.MaxInt32
		}
	}

	for {
		if resourcesAfter[ore] >= robotCost[ore] &&
			resourcesAfter[clay] >= robotCost[clay] &&
			resourcesAfter[obsidian] >= robotCost[obsidian] {
			break
		}
		for material := 0; material < 4; material++ {
			resourcesAfter[material] += numRobots[material]
		}
		waitTime++
	}

	for material := 0; material < 4; material++ {
		resourcesAfter[material] += numRobots[material] - robotCost[material]
	}
	return
}

func countGeodes(b blueprint, timeLimit int, resources cost, numRobots [4]int, best []int) int {

	if numRobots[geode] > best[timeLimit] {
		best[timeLimit] = numRobots[geode]
	}

	if timeLimit <= 0 {
		return resources[geode]
	}

	g := 0

	for _, material := range []int{geode, obsidian, clay, ore} {
		if newResources, delta := resources.NextAvail(b.robotCosts[material], numRobots); delta < timeLimit {
			numRobots[material]++
			if numRobots[geode] >= best[timeLimit-delta] {
				if res := countGeodes(b, timeLimit-delta-1, newResources, numRobots, best); res > g {
					g = res
				}
			}
			numRobots[material]--
		}
	}

	return g
}

func parseBlueprints(input []string) []blueprint {
	blueprints := make([]blueprint, 0)

	for _, line := range input {
		if line == "" {
			continue
		}

		var ignored int
		b := blueprint{}

		_ = utils.Must(fmt.Sscanf(
			line,
			"Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
			&ignored,
			&b.robotCosts[ore][ore],
			&b.robotCosts[clay][ore],
			&b.robotCosts[obsidian][ore], &b.robotCosts[obsidian][clay],
			&b.robotCosts[geode][ore], &b.robotCosts[geode][obsidian],
		))

		blueprints = append(blueprints, b)
	}

	return blueprints
}

func solveA(input []string) int {
	total := 0
	limit := 24

	blueprints := parseBlueprints(input)

	for idx, b := range blueprints {
		total += (idx + 1) * countGeodes(b, limit, cost{}, [4]int{1, 0, 0, 0}, make([]int, limit+1))
	}

	return total
}

func solveB(input []string) int {
	total := 1
	limit := 32

	blueprints := parseBlueprints(input)
	if len(blueprints) > 3 {
		blueprints = blueprints[:3]
	}

	for _, b := range blueprints {
		total *= countGeodes(b, limit, cost{}, [4]int{1, 0, 0, 0}, make([]int, limit+1))
	}

	return total
}

func main() {
	problem := utils.Must(utils.NewProblem(2022, 19))
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
