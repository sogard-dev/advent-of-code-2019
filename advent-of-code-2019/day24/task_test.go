package day24

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2019/utils"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 2129920, part1(`....#
#..#.
#..##
..#..
#....`))
	require.Equal(t, 17863741, part1(utils.GetInput(t, "input.txt")))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 0, part2(``))
	require.Equal(t, 0, part2(utils.GetInput(t, "input.txt")))
}
