package day2

import intcodecomputer "github.com/sogard-dev/advent-of-code-2019/intcode"

func part1(input []int) int {
	icc := intcodecomputer.New(input)
	icc.ExecuteUntilHalt()
	return icc.GetMemory0()
}

func part2(input []int) int {
	for noun := range 100 {
		for verb := range 100 {
			input[1] = noun
			input[2] = verb
			icc := intcodecomputer.New(input)
			icc.ExecuteUntilHalt()
			if icc.GetMemory0() == 19690720 {
				return noun*100 + verb
			}
		}
	}

	panic("Damn")
}
