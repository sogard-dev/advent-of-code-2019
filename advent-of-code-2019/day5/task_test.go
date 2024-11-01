package day5

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2019/utils"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 14155342, part1(utils.GetInput(t, "input.txt")))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 8684145, part2(utils.GetInput(t, "input.txt")))
}
