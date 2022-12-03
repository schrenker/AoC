package y2022

import (
	"github.com/schrenker/AoC/tools"
)

type Day03 struct{}

func getItemScore(item byte) int {
	swap := int(item) ^ 32
	if swap <= 90 {
		return swap - 64
	} else {
		return swap - 70
	}
}

func (d Day03) PartOne(data []byte) interface{} {
	items := tools.ByteToStringSlice(data)
	acc := 0

	for _, v := range items {
		occurences := make(map[byte]bool)

		for i := 0; i < len(v)/2; i++ {
			occurences[v[i]] = true
		}

		for i := len(v) / 2; i < len(v); i++ {
			if occurences[v[i]] {
				acc += getItemScore(v[i])
				occurences[v[i]] = false
			}
		}
	}

	return acc
}

func (d Day03) PartTwo(data []byte) interface{} {
	items := tools.ByteToStringSlice(data)
	acc := 0

	for i := 0; i < len(items); i += 3 {
		occurences := make(map[byte]int)

		for j := 0; j < 3; j++ {
			uniq := make(map[byte]bool)

			for k := range items[i+j] {
				uniq[items[i+j][k]] = true
			}

			for k := range uniq {
				occurences[k]++
			}
		}

		for k, v := range occurences {
			if v == 3 {
				acc += getItemScore(k)
			}
		}

	}

	return acc
}
