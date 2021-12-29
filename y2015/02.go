package y2015

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/schrenker/go_aoc/tools"
)

type DayTwo struct{}

func (d DayTwo) PartOne() {
	acc := 0
	for _, line := range tools.ReadFileStringSlice("input/2015/02.txt") {
		tmp := strings.Split(line, "x")
		numtmp := make([]int, 3)
		for i, num := range tmp {
			n, _ := strconv.Atoi(num)
			numtmp[i] = n
		}
	}
	fmt.Println(acc)
}

func (d DayTwo) PartTwo() {
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
