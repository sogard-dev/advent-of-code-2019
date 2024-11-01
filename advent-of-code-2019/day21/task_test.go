package day21

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2019/utils"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 19357335, part1(utils.GetInput(t, "input.txt")))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 1140147758, part2(utils.GetInput(t, "input.txt")))
}
