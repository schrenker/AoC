package main

type challenge interface {
	PartOne([]byte) interface{}
	PartTwo([]byte) interface{}
}

var challenges = map[string]challenge{
	// "2022/12": y2022.Day12{},
}
