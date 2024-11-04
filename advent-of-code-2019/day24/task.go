package day24

import (
	"fmt"
	"strings"
)

func part1(input string) int {
	layout := strings.ReplaceAll(input, "\n", "")

	seen := map[string]bool{layout: true}
	for {
		newLayout := iterate(layout)
		if seen[newLayout] {
			return calc(newLayout)
		}
		layout = newLayout
		seen[layout] = true

	}
}

func calc(newLayout string) int {
	sum := 0

	for i, v := range newLayout {
		if v == '#' {
			bio := 1 << i
			sum += bio
		}
	}

	return sum
}

func iterate(n string) string {
	newLayout := []rune(strings.ReplaceAll(n, "#", "."))

	for i := 0; i < len(n); i++ {
		ns := 0
		for _, d := range []int{1, -1, 5, -5} {
			di := i + d
			if di >= 0 && di < 25 {
				if (d == 1 && i%5 == 4) || (d == -1 && i%5 == 0) {
					continue
				}
				if n[di] == '#' {
					ns += 1
				}
			}
		}

		if n[i] == '#' {
			if ns == 1 {
				newLayout[i] = '#'
			}
		} else {
			if ns == 1 || ns == 2 {
				newLayout[i] = '#'
			}
		}
	}

	return string(newLayout)
}

func part2(input string) int {
	for _, line := range strings.Split(input, "\n") {
		fmt.Println(line)
	}
	return 0
}
