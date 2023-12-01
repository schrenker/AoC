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
		for i := 0; i < len(line); i++ {
			if num, err := strconv.Atoi(string(line[i])); err == nil {
				acc += num * 10
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if num, err := strconv.Atoi(string(line[i])); err == nil {
				acc += num
				break
			}
		}
	}
	return acc
}

func (d Day01) PartTwo(data []byte) interface{} {
	// tmp := input.ByteToStringSlice(data)
	// literals := map[string]int{
	// 	"one": 1,
	// 	"two": 2,
	// 	"three": 3,
	// 	"four": 4,
	// 	"five": 5,
	// 	"six": 6,
	// 	"seven": 7,
	// 	"eight": 8,
	// 	"nine": 9,
	// }
	acc := 0
	// for _, line := range tmp {
	// 	for i := 0; i < len(line); i++ {
	// 		if num, err := strconv.Atoi(string(line[i])); err == nil {
	// 			acc += num * 10
	// 			break
	// 		}
	// 	}
	// 	for i := len(line) - 1; i >= 0; i-- {
	// 		if num, err := strconv.Atoi(string(line[i])); err == nil {
	// 			acc += num
	// 			break
	// 		}
	// 	}
	// }

	return acc
}
