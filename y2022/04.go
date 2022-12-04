package y2022

import (
	"regexp"
	"strconv"

	"github.com/schrenker/AoC/tools"
)

type Day04 struct{}

func contains(start1, end1, start2, end2 int) bool {
	return start2 >= start1 && end2 <= end1
}

func overlaps(start1, end1, start2, end2 int) bool {
	return end1 >= start2 && start1 <= start2
}

func (d Day04) PartOne(data []byte) interface{} {
	pairs := tools.ByteToStringSlice(data)
	acc := 0

	for _, v := range pairs {
		t := regexp.MustCompile(`[,-]`)
		values := t.Split(v, -1)
		start1, _ := strconv.Atoi(values[0])
		end1, _ := strconv.Atoi(values[1])
		start2, _ := strconv.Atoi(values[2])
		end2, _ := strconv.Atoi(values[3])

		if contains(start1, end1, start2, end2) || contains(start2, end2, start1, end1) {
			acc++
		}
	}

	return acc
}

func (d Day04) PartTwo(data []byte) interface{} {
	pairs := tools.ByteToStringSlice(data)
	acc := 0

	for _, v := range pairs {
		t := regexp.MustCompile(`[,-]`)
		values := t.Split(v, -1)
		start1, _ := strconv.Atoi(values[0])
		end1, _ := strconv.Atoi(values[1])
		start2, _ := strconv.Atoi(values[2])
		end2, _ := strconv.Atoi(values[3])

		if overlaps(start1, end1, start2, end2) || overlaps(start2, end2, start1, end1) {
			acc++
		}
	}

	return acc
}
