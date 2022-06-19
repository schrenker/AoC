package main

import "github.com/schrenker/AoC/y2015"

type challenge interface {
	PartOne() interface{}
	PartTwo() interface{}
}

var challenges = map[string]challenge{
	"2015/01": y2015.Day01{},
	"2015/02": y2015.Day02{},
	"2015/03": y2015.Day03{},
	"2015/04": y2015.Day04{},
	"2015/05": y2015.Day05{},
	"2015/06": y2015.Day06{},
	"2015/07": y2015.Day07{},
	"2015/08": y2015.Day08{},
	"2015/09": y2015.Day09{},
	"2015/10": y2015.Day10{},
	"2015/11": y2015.Day11{},
	"2015/12": y2015.Day12{},
	"2015/13": y2015.Day13{},
}
