package y2015

import (
	"strconv"

	"github.com/schrenker/AoC/tools"
)

type Day17 struct{}

func (d Day17) PartOne(data []byte) interface{} {
	target := 150
	acc := 0
	for _, d := range tools.Combinations(tools.ByteToStringSlice(data)) {
		tmp := make([]int, len(d))
		for i, v := range d {
			tmp[i], _ = strconv.Atoi(v)
		}
		if target == tools.SumAll(tmp...) {
			acc++
		}
	}
	return acc
}

func (d Day17) PartTwo(data []byte) interface{} {
	target := 150
	min := make(map[int]int)
	for _, d := range tools.Combinations(tools.ByteToStringSlice(data)) {
		tmp := make([]int, len(d))
		for i, v := range d {
			tmp[i], _ = strconv.Atoi(v)
		}
		if target == tools.SumAll(tmp...) {
			min[len(tmp)]++
		}
	}

	length := target
	for k := range min {
		if k < length {
			length = k
		}
	}

	return min[length]
}
