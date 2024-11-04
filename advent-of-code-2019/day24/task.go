package day24

import (
	"strings"
)

const SIZE = 5

type bug struct {
	x, y int
}

type grid map[bug]bool

func gridToString(m grid) string {
	ret := make([]rune, SIZE*SIZE)

	for y := range SIZE {
		for x := range SIZE {
			if m[bug{x: x, y: y}] {
				ret[y*SIZE+x] = '#'
			} else {
				ret[y*SIZE+x] = '.'
			}
		}
	}

	return string(ret)
}

func part1(input string) int {
	layout := parseGrid(input)

	seen := map[string]bool{}
	asString := gridToString(layout)
	for !seen[asString] {
		seen[asString] = true
		layout = iterate(map[int]grid{0: layout}, false)[0]
		asString = gridToString(layout)
	}
	return calc(layout)
}

func part2(input string, minutes int) int {
	grid0 := parseGrid(input)
	grids := map[int]grid{0: grid0}

	for range minutes {
		grids = iterate(grids, true)
	}

	sum := 0
	for _, g := range grids {
		sum += len(g)
	}
	return sum
}

func parseGrid(input string) grid {
	grid := grid{}
	for y, line := range strings.Split(input, "\n") {
		for x, v := range strings.Split(line, "") {
			if v == "#" {
				grid[bug{x: x, y: y}] = true
			}
		}
	}
	return grid
}

func calc(m grid) int {
	sum := 0

	for k := range m {
		num := k.y*SIZE + k.x
		bio := 1 << num
		sum += bio
	}

	return sum
}

func iterate(m map[int]grid, part2 bool) map[int]grid {
	neighbours := map[int]map[bug]int{}

	directions := []bug{{x: 1, y: 0}, {x: 0, y: 1}, {x: -1, y: 0}, {x: 0, y: -1}}

	addAdjacent := func(level int, b bug) {
		if _, exists := neighbours[level]; !exists {
			neighbours[level] = map[bug]int{}
		}
		neighbours[level][b]++
	}

	for l := range m {
		for y := range SIZE {
			for x := range SIZE {
				if m[l][bug{x: x, y: y}] {
					for _, d := range directions {
						nx := d.x + x
						ny := d.y + y

						isRecMiddle := nx == 2 && ny == 2 && part2
						if nx >= 0 && nx < SIZE && ny >= 0 && ny < SIZE && !isRecMiddle {
							addAdjacent(l, bug{x: nx, y: ny})
						}

						if part2 {
							if isRecMiddle {
								if d.x == -1 {
									for i := range SIZE {
										addAdjacent(l+1, bug{x: 4, y: i})
									}
								}
								if d.x == 1 {
									for i := range SIZE {
										addAdjacent(l+1, bug{y: i})
									}
								}
								if d.y == -1 {
									for i := range SIZE {
										addAdjacent(l+1, bug{x: i, y: 4})
									}
								}
								if d.y == 1 {
									for i := range SIZE {
										addAdjacent(l+1, bug{x: i})
									}
								}
							}
							if nx == SIZE {
								addAdjacent(l-1, bug{x: 3, y: 2})
							} else if nx == -1 {
								addAdjacent(l-1, bug{x: 1, y: 2})
							}

							if ny == SIZE {
								addAdjacent(l-1, bug{x: 2, y: 3})
							} else if ny == -1 {
								addAdjacent(l-1, bug{x: 2, y: 1})
							}
						}
					}
				}
			}
		}
	}

	newGrid := map[int]grid{}
	for l, g := range neighbours {

		for k, v := range g {
			hasBug := m[l] != nil && m[l][k]
			shouldBeBug := (hasBug && v == 1) || (!hasBug && (v == 1 || v == 2))
			if shouldBeBug {
				if _, exists := newGrid[l]; !exists {
					newGrid[l] = grid{}
				}
				newGrid[l][k] = true
			}
		}
	}

	return newGrid
}
