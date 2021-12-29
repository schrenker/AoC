package y2015

import (
	"fmt"

	"github.com/schrenker/go_aoc/tools"
)

type DayOne struct{}

func (d DayOne) PartOne() {
	acc := 0
	for _, i := range tools.ReadFileBytes("input/2015/01.txt") {
		if i == '(' {
			acc++
		} else if i == ')' {
			acc--
		}
	}
	fmt.Println(acc)
}

func (d DayOne) PartTwo() {
	fmt.Println("testualnie2")
}
