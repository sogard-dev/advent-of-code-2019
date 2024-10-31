package day19

import (
	"strings"
)

type pos struct {
	x, y int
}

type portal struct {
	outer bool
	label string
}

func getDirections() []pos {
	return []pos{
		{x: -1},
		{y: -1},
		{x: 1},
		{y: 1},
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func parse(inputSr string) map[portal]map[portal]int {
	maze := map[pos]bool{}

	input := strings.Split(inputSr, "\n")
	for y, line := range input {
		for x, v := range line {
			if v == '.' {
				maze[pos{x: x, y: y}] = true
			}
		}
	}

	width := len(input[0])
	height := len(input)

	directions := getDirections()
	jumps := map[pos]portal{}
	for mp := range maze {
		for _, dp := range directions {
			p := pos{
				x: mp.x + dp.x,
				y: mp.y + dp.y,
			}
			f := input[p.y][p.x]
			if f != ' ' && f != '.' && f != '#' {
				isInner := true
				if p.x < 5 || p.y < 5 || width-p.x < 5 || height-p.y < 5 {
					isInner = false
				}
				s := input[p.y+dp.y][p.x+dp.x]
				if dp.y == -1 || dp.x == -1 {
					jumps[mp] = portal{label: string([]rune{rune(s), rune(f)}), outer: !isInner}
				} else {
					jumps[mp] = portal{label: string([]rune{rune(f), rune(s)}), outer: !isInner}
				}
			}
		}
	}

	travelMap := map[portal]map[portal]int{}
	var traverse func(startPortal portal, p pos, visited map[pos]bool, steps int)
	traverse = func(startPortal portal, p pos, visited map[pos]bool, steps int) {
		visited[p] = true

		for _, dp := range directions {
			np := pos{
				x: p.x + dp.x,
				y: p.y + dp.y,
			}
			if !visited[np] && maze[np] {
				traverse(startPortal, np, visited, steps+1)
			}
		}

		if label, hasLabel := jumps[p]; hasLabel {
			if _, e := travelMap[startPortal]; !e {
				travelMap[startPortal] = map[portal]int{}
			}

			if label != startPortal {
				if dist, hasSeen := travelMap[startPortal][label]; hasSeen {
					travelMap[startPortal][label] = min(dist, steps)
				} else {
					travelMap[startPortal][label] = steps
				}
			}
		}

		visited[p] = false
	}

	for p, label := range jumps {
		traverse(label, p, map[pos]bool{}, 0)
	}
	return travelMap
}

func part1(input string) int {
	travelMap := parse(input)

	jumpMap := map[string]map[string]int{}
	for from, toMap := range travelMap {
		if _, exist := jumpMap[from.label]; !exist {
			jumpMap[from.label] = map[string]int{}
		}
		for to, dist := range toMap {
			jumpMap[from.label][to.label] = dist
		}
	}

	return solvePart1(jumpMap)
}

func part2(input string) int {
	travelMap := parse(input)
	best := 500
	for {
		solution := solvePart2(travelMap, best)
		if solution != best-1 {
			return solution
		}
		best *= 2
	}
}

func solvePart1(jumps map[string]map[string]int) int {
	distances := map[string]int{}

	var traverse func(label string, visited map[string]bool, steps int)
	traverse = func(label string, visited map[string]bool, steps int) {
		if dist, hasSeen := distances[label]; hasSeen {
			distances[label] = min(dist, steps-1)
		} else {
			distances[label] = steps - 1
		}

		visited[label] = true
		for nextLabel, dist := range jumps[label] {
			if !visited[nextLabel] {
				traverse(nextLabel, visited, steps+dist+1)
			}
		}
		visited[label] = false
	}

	traverse("AA", map[string]bool{}, 0)
	return distances["ZZ"]
}

func solvePart2(travelMap map[portal]map[portal]int, bestDistance int) int {
	var traverse func(current portal, steps int, level int)
	traverse = func(current portal, steps int, level int) {
		if steps > bestDistance {
			return
		}

		if current.label == "ZZ" {
			if steps < bestDistance {
				println("New best", steps)
				bestDistance = steps
			}
		}

		for nextPortal, dist := range travelMap[current] {
			isZZ := nextPortal.label == "ZZ"
			isInner := !nextPortal.outer
			isAllowed := (isZZ && level == 0) || (!isZZ && (isInner || level > 0))
			if isAllowed {
				//Check labels
				if isInner {
					traverse(portal{label: nextPortal.label, outer: true}, steps+dist+1, level+1)
				} else {
					traverse(portal{label: nextPortal.label, outer: false}, steps+dist+1, level-1)
				}
			}
		}
	}

	traverse(portal{label: "AA", outer: true}, 0, 0)
	return bestDistance - 1
}
