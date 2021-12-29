package main

import (
	. "github.com/schrenker/go_aoc/tools"
	. "github.com/schrenker/go_aoc/y2015"
)

var challenges = map[string]Challenge{
	"2015/01": DayOne{},
}

func main() {
	challenges["2015/01"].PartOne()
	challenges["2015/01"].PartTwo()
}
