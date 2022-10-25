package y2015

import (
	"github.com/schrenker/AoC/tools"
)

type Day01 struct{}

func (d Day01) PartOne(path string) interface{} {
	acc := 0
	for _, v := range tools.ReadFileBytes(path) {
		if v == '(' {
			acc++
		} else if v == ')' {
			acc--
		}
	}
	return acc
}

func (d Day01) PartTwo(path string) interface{} {
	acc := 0
	for i, v := range tools.ReadFileBytes(path) {
		if v == '(' {
			acc++
		} else if v == ')' {
			acc--
		}
		if acc < 0 {
			return i + 1
		}
	}
	return acc
}
