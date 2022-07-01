package y2015

import (
	"github.com/schrenker/AoC/tools"
)

type Day01 struct{}

func (d Day01) PartOne() interface{} {
	acc := 0
	for _, v := range tools.ReadFileBytes() {
		if v == '(' {
			acc++
		} else if v == ')' {
			acc--
		}
	}
	return acc
}

func (d Day01) PartTwo() interface{} {
	acc := 0
	for i, v := range tools.ReadFileBytes() {
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
