package day9

import (
	intcodecomputer "github.com/sogard-dev/advent-of-code-2019/intcode"
	"github.com/sogard-dev/advent-of-code-2019/utils"
)

func part1(input string) int {
	return solve(input, 1)
}

func part2(input string) int {
	return solve(input, 2)
}

func solve(input string, n int) int {
	icc := intcodecomputer.New(utils.GetAllNumbers(input))
	icc.SetInput([]int{n})
	outputs := []int{}
	icc.SetOutputter(func(n int) {
		outputs = append(outputs, n)
	})
	icc.ExecuteUntilHalt()
	return outputs[len(outputs)-1]
}
