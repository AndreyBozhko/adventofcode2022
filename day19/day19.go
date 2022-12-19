package main

import (
	"adventofcode2022/utils"
	"errors"
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

func (c *cost) Allocate(r cost) (cost, error) {
	res := cost{}
	if c[clay] < r[clay] || c[ore] < r[ore] || c[obsidian] < r[obsidian] {
		return cost{}, errors.New("not enough")
	}

	res[clay] = c[clay] - r[clay]
	res[ore] = c[ore] - r[ore]
	res[obsidian] = c[obsidian] - r[obsidian]
	return res, nil
}

func (c *cost) NextAvailable(r cost, robots [4]int) (cost, int) {
	res := *c
	minute := 0
	for res[ore] < r[ore] || res[clay] < r[clay] || res[obsidian] < r[obsidian] {
		if res[ore] < r[ore] && robots[ore] == 0 {
			return cost{}, math.MaxInt32
		}
		if res[clay] < r[clay] && robots[clay] == 0 {
			return cost{}, math.MaxInt32
		}
		if res[obsidian] < r[obsidian] && robots[obsidian] == 0 {
			return cost{}, math.MaxInt32
		}
		res[clay] += robots[clay]
		res[ore] += robots[ore]
		res[obsidian] += robots[obsidian]
		res[geode] += robots[geode]
		minute++
	}
	res[clay] += robots[clay]
	res[ore] += robots[ore]
	res[obsidian] += robots[obsidian]
	res[geode] += robots[geode]
	res[clay] -= r[clay]
	res[ore] -= r[ore]
	res[obsidian] -= r[obsidian]
	return res, minute
}

type blueprint struct {
	OreRobot, ClayRobot, ObsidianRobot, GeodeRobot cost
}

func countGeodes(b blueprint, timeLimit int, resources cost, robotsTotal [4]int) int {

	if timeLimit <= 0 {
		return resources[geode]
	}

	g := 0

	if newResources, delta := resources.NextAvailable(b.GeodeRobot, robotsTotal); delta < timeLimit {
		robotsTotal[geode]++
		if res := countGeodes(b, timeLimit-delta-1, newResources, robotsTotal); res > g {
			g = res
		}
		robotsTotal[geode]--
	}
	if newResources, delta := resources.NextAvailable(b.ObsidianRobot, robotsTotal); delta < timeLimit {
		robotsTotal[obsidian]++
		if res := countGeodes(b, timeLimit-delta-1, newResources, robotsTotal); res > g {
			g = res
		}
		robotsTotal[obsidian]--
	}
	if newResources, delta := resources.NextAvailable(b.ClayRobot, robotsTotal); delta < timeLimit {
		robotsTotal[clay]++
		if res := countGeodes(b, timeLimit-delta-1, newResources, robotsTotal); res > g {
			g = res
		}
		robotsTotal[clay]--
	}
	if newResources, delta := resources.NextAvailable(b.OreRobot, robotsTotal); delta < timeLimit {
		robotsTotal[ore]++
		if res := countGeodes(b, timeLimit-delta-1, newResources, robotsTotal); res > g {
			g = res
		}
		robotsTotal[ore]--
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
			&ignored, &b.OreRobot[ore], &b.ClayRobot[ore], &b.ObsidianRobot[ore], &b.ObsidianRobot[clay], &b.GeodeRobot[ore], &b.GeodeRobot[obsidian],
		))

		blueprints = append(blueprints, b)
	}

	return blueprints
}

func solveA(input []string) int {
	total := 0

	blueprints := parseBlueprints(input)

	for idx, b := range blueprints {
		total += (idx + 1) * countGeodes(b, 24, cost{}, [4]int{1, 0, 0, 0})
		println("Done:", idx+1)
	}

	return total
}

func solveB(input []string) int {
	total := 1

	blueprints := parseBlueprints(input)
	if len(blueprints) > 3 {
		blueprints = blueprints[:3]
	}

	for idx, b := range blueprints {
		total *= countGeodes(b, 32, cost{}, [4]int{1, 0, 0, 0})
		println("Done:", idx+1)
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
