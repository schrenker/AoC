package main

import "github.com/schrenker/go_aoc/y2015"

type challenge interface {
	PartOne()
	PartTwo()
}

var challenges = map[string]challenge{
	"2015/01": y2015.DayOne{},
	"2015/02": y2015.DayTwo{},
	"2015/03": y2015.DayThree{},
	"2015/04": y2015.DayFour{},
	"2015/05": y2015.DayFive{},
	"2015/06": y2015.DaySix{},
}
