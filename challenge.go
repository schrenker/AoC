package main

import "github.com/schrenker/AoC/y2015"

type challenge interface {
	PartOne() interface{}
	PartTwo() interface{}
}

var challenges = map[string]challenge{
	"2015/01": y2015.DayOne{},
	"2015/02": y2015.DayTwo{},
	"2015/03": y2015.DayThree{},
	"2015/04": y2015.DayFour{},
	"2015/05": y2015.DayFive{},
	"2015/06": y2015.DaySix{},
	"2015/07": y2015.DaySeven{},
	"2015/08": y2015.DayEight{},
	"2015/09": y2015.DayNine{},
	"2015/10": y2015.DayTen{},
	"2015/11": y2015.DayEleven{},
	"2015/12": y2015.DayTwelve{},
}
