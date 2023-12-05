package main

import "github.com/schrenker/AoC/y2023"

type challenge interface {
	PartOne([]byte) interface{}
	PartTwo([]byte) interface{}
}

var challenges = map[string]challenge{
	"2023/01": y2023.Day01{},
	"2023/02": y2023.Day02{},
	"2023/03": y2023.Day03{},
	"2023/04": y2023.Day04{},
	"2023/05": y2023.Day05{},
}
