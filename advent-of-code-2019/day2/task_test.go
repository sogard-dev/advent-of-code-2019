package day2

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2019/utils"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 9801, part1([]int{2, 4, 4, 0, 99, 0}))
	require.Equal(t, 30, part1([]int{1, 1, 1, 4, 99, 5, 6, 0, 99}))

	program := utils.GetAllNumbers(utils.GetInput(t, "input.txt"))
	program[1] = 12
	program[2] = 2
	require.Equal(t, 5098658, part1(program))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 5064, part2(utils.GetAllNumbers(utils.GetInput(t, "input.txt"))))

}
