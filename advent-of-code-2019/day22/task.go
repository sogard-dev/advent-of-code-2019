package day22

import (
	"math/big"
	"strings"

	"github.com/sogard-dev/advent-of-code-2019/utils"
)

func part1(input string) int {
	size := 10007
	actual := 2019

	//y=ax+b % m
	a := big.NewInt(1)
	b := big.NewInt(0)
	x := big.NewInt(int64(actual))
	m := big.NewInt(int64(size))

	for _, line := range strings.Split(input, "\n") {
		if strings.HasPrefix(line, "deal with increment") {
			n := big.NewInt(int64(utils.GetAllNumbers(line)[0]))
			a.Mul(a, n)
			b.Mul(b, n)
		} else if strings.HasPrefix(line, "cut") {
			n := big.NewInt(int64(utils.GetAllNumbers(line)[0]))
			b.Add(b, n.Neg(n))
		} else if strings.HasPrefix(line, "deal into new stack") {
			a.Neg(a)
			b.Neg(b)
			b.Add(b, big.NewInt(-1))
		} else {
			panic("Unknown line: " + line)
		}
	}

	guessBig := big.NewInt(0)
	guessBig.Mul(a, x)
	guessBig.Add(guessBig, b)
	guessBig.Mod(guessBig, m)
	guess := int(guessBig.Int64())

	return guess
}

func part2(input string) int {
	return len(input)
}
