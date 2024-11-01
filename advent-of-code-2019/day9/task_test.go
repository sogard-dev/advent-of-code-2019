package day9

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2019/utils"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 3780860499, part1(utils.GetInput(t, "input.txt")))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 33343, part2(utils.GetInput(t, "input.txt")))
}
