package day21

import (
	"strings"

	intcodecomputer "github.com/sogard-dev/advent-of-code-2019/intcode"
	"github.com/sogard-dev/advent-of-code-2019/utils"
)

func part1(input string) int {
	program := `
	NOT A T
	OR T J

	NOT B T
	AND D T
	OR T J

	NOT C T
	AND D T
	OR T J
	WALK`
	return solve(input, program)
}

func part2(input string) int {
	//D is safe, ABC is missing, E or H is safe
	program := `
OR D J

OR A T
AND B T
AND C T
NOT T T
AND T J

AND H T
OR E T
AND T J
	RUN`
	return solve(input, program)
}

func solve(input string, program string) int {
	icc := intcodecomputer.New(utils.GetAllNumbers(input))
	inputs := []int{}
	for _, line := range strings.Split(program, "\n") {
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			for _, c := range line {
				inputs = append(inputs, int(c))
			}
			inputs = append(inputs, '\n')
		}

	}

	icc.SetInput(inputs)
	icc.ExecuteUntilHalt()
	outputs := icc.GetOutputs()
	for _, i := range outputs {
		print(string(rune(i)))
	}

	return outputs[len(outputs)-1]
}
