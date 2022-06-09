package y2015

import (
	"fmt"

	"github.com/schrenker/AoC/tools"
)

type DayOne struct{}

func (d DayOne) PartOne() {
	acc := 0
	for _, v := range tools.ReadFileBytes("input/2015/01.txt") {
		if v == '(' {
			acc++
		} else if v == ')' {
			acc--
		}
	}
	fmt.Println(acc)
}

func (d DayOne) PartTwo() {
	acc := 0
	for i, v := range tools.ReadFileBytes("input/2015/01.txt") {
		if acc < 0 {
			fmt.Println(i)
			break
		}
		if v == '(' {
			acc++
		} else if v == ')' {
			acc--
		}
	}
}
