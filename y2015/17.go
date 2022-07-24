package y2015

import (
	"strconv"

	"github.com/schrenker/AoC/tools"
)

type Day17 struct{}

func (d Day17) PartOne() interface{} {
	target := 150
	data := tools.Combinations(tools.ReadFileStringSlice())
	acc := 0
	for _, d := range data {
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

func (d Day17) PartTwo() interface{} {
	return 0
}
