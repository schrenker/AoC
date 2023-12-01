package y2023

import (
	"strconv"

	"github.com/schrenker/AoC/input"
)

type Day01 struct{}

func (d Day01) PartOne(data []byte) interface{} {
	tmp := input.ByteToStringSlice(data)
	acc := 0
	for _, line := range tmp {
		first, last := 0, 0
		for i := 0; i < len(line); i++ {
			if num, err := strconv.Atoi(string(line[i])); err == nil {
				last = num
				if first == 0 {
					first = num * 10
				}
			}
		}
		acc += first + last
	}
	return acc
}

func (d Day01) PartTwo(data []byte) interface{} {
	tmp := input.ByteToStringSlice(data)
	literals := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	acc := 0
	for _, line := range tmp {
		first, last := 0, 0
		for i := 0; i < len(line); i++ {
			num, err := strconv.Atoi(string(line[i]))
			if err != nil {
				for lit, n := range literals {
					if len(lit)+i <= len(line) && line[i:len(lit)+i] == lit {
						num = n
						break
					}
				}
			}

			if num != 0 {
				last = num
				if first == 0 {
					first = num * 10
				}
			}
		}
		acc += first + last
	}

	return acc
}
