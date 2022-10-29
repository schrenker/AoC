package y2015

import (
	"strconv"
	"strings"

	"github.com/schrenker/AoC/tools"
)

type Day02 struct{}

func (d Day02) PartOne(data []byte) interface{} {
	acc := 0
	for _, line := range tools.ByteToStringSlice(data) {
		tmp := strings.Split(line, "x")
		numtmp := make([]int, 3)
		for i, num := range tmp {
			n, _ := strconv.Atoi(num)
			numtmp[i] = n
		}
		first := numtmp[0] * numtmp[1]
		second := numtmp[0] * numtmp[2]
		third := numtmp[1] * numtmp[2]

		acc += tools.SumAll(2*first, 2*second, 2*third, (tools.GetMin(first, second, third)))
	}
	return acc
}

func (d Day02) PartTwo(data []byte) interface{} {
	acc := 0
	for _, line := range tools.ByteToStringSlice(data) {
		tmp := strings.Split(line, "x")
		numtmp := make([]int, 3)
		for i, num := range tmp {
			n, _ := strconv.Atoi(num)
			numtmp[i] = n
		}
		first := numtmp[0]
		second := numtmp[1]
		third := numtmp[2]

		twoShort := tools.GetMinMultiple(2, first, second, third)
		acc += (2 * tools.SumAll(twoShort...)) + tools.MulAll(first, second, third)
	}
	return acc
}
