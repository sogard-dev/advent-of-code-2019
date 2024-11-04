package day24

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2019/utils"
	"github.com/stretchr/testify/require"
)

const testInput = `....#
#..#.
#..##
..#..
#....`

func TestPart1(t *testing.T) {
	require.Equal(t, 2129920, part1(testInput))
	require.Equal(t, 17863741, part1(utils.GetInput(t, "input.txt")))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 99, part2(testInput, 10))
	require.Equal(t, 2029, part2(utils.GetInput(t, "input.txt"), 200))
}
