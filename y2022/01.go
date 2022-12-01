package y2022

import (
	"sort"
	"strconv"

	"github.com/schrenker/AoC/tools"
)

type Day01 struct{}

func (d Day01) PartOne(data []byte) interface{} {
	inp := tools.ByteToStringSlice(data)
	high := 0
	acc := 0
	for _, v := range inp {
		if v == "" {
			high = tools.GetMax(high, acc)
			acc = 0
		} else {
			tmp, _ := strconv.Atoi(v)
			acc += tmp
		}
	}
	return high
}

func (d Day01) PartTwo(data []byte) interface{} {
	inp := tools.ByteToStringSlice(data)
	high := make([]int, 0)
	acc := 0
	for _, v := range inp {
		if v == "" {
			high = append(high, acc)
			acc = 0
		} else {
			tmp, _ := strconv.Atoi(v)
			acc += tmp
		}
	}
	high = append(high, acc)
	sort.Ints(high)
	return tools.SumAll(
		high[len(high)-1],
		high[len(high)-2],
		high[len(high)-3])
}
