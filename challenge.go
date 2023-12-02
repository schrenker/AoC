package main

import "github.com/schrenker/AoC/y2023"

type challenge interface {
	PartOne([]byte) interface{}
	PartTwo([]byte) interface{}
}

var challenges = map[string]challenge{
	"2023/01": y2023.Day01{},
	"2023/02": y2023.Day02{},
}
