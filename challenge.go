package main

import "github.com/schrenker/AoC/y2015"

type challenge interface {
	PartOne([]byte) interface{}
	PartTwo([]byte) interface{}
}

var challenges = map[string]challenge{
	"2015/01": y2015.Day01{},
	"2015/02": y2015.Day02{},
	"2015/03": y2015.Day03{},
	"2015/04": y2015.Day04{},
	"2015/05": y2015.Day05{},
}
