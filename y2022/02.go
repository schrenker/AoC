package y2022

import (
	"github.com/schrenker/AoC/tools"
)

type Day02 struct{}

const (
	lose = 0
	draw = 3
	win  = 6
)

func (d Day02) PartOne(data []byte) interface{} {
	resolve := func(opponent, you int) int {
		res := opponent - you + 1
		switch res {
		case 0, 3:
			return win
		case 1:
			return draw
		}
		return lose
	}

	databook := tools.ByteToStringSlice(data)
	mapping := map[byte]int{
		'A': 1,
		'B': 2,
		'C': 3,
		'X': 1,
		'Y': 2,
		'Z': 3,
	}
	acc := 0

	for _, v := range databook {
		acc += resolve(mapping[v[0]], mapping[v[2]])
		acc += mapping[v[2]]
	}

	return acc
}

func (d Day02) PartTwo(data []byte) interface{} {
	resolve := func(opponent, result int) int {
		switch result {
		case 0:
			return (opponent+1)%3 + 1
		case 3:
			return opponent
		case 6:
			return opponent%3 + 1
		}
		return 1
	}

	databook := tools.ByteToStringSlice(data)
	mapping := map[byte]int{
		'A': 1,
		'B': 2,
		'C': 3,
		'X': lose,
		'Y': draw,
		'Z': win,
	}
	acc := 0

	for _, v := range databook {
		acc += resolve(mapping[v[0]], mapping[v[2]])
		acc += mapping[v[2]]
	}

	return acc
}
