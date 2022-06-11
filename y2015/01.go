package y2015

import (
	"github.com/schrenker/AoC/tools"
)

type DayOne struct{}

func (d DayOne) PartOne() interface{} {
	acc := 0
	for _, v := range tools.ReadFileBytes("input/2015/01.txt") {
		if v == '(' {
			acc++
		} else if v == ')' {
			acc--
		}
	}
	return acc
}

func (d DayOne) PartTwo() interface{} {
	acc := 0
	for i, v := range tools.ReadFileBytes("input/2015/01.txt") {
		if acc < 0 {
			return i
		}
		if v == '(' {
			acc++
		} else if v == ')' {
			acc--
		}
	}
	return acc
}
