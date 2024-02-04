package y2015

import (
	"slices"
	"strconv"
	"strings"

	"github.com/schrenker/AoC/input"
)

type Day02 struct{}

func calcRibbon(dim []string) int {
	l, _ := strconv.Atoi(dim[0])
	w, _ := strconv.Atoi(dim[1])
	h, _ := strconv.Atoi(dim[2])
	slc := []int{l, w, h}
	slices.Sort(slc)
    return 2*slc[0] + 2*slc[1] + l*w*h
}

func findSmallestSide(l, w, h int) int {
	slc := []int{l, w, h}
	slices.Sort(slc)
	return slc[0] * slc[1]
}

func calcPaper(dim []string) int {
	l, _ := strconv.Atoi(dim[0])
	w, _ := strconv.Atoi(dim[1])
	h, _ := strconv.Atoi(dim[2])
	return 2*l*w + 2*l*h + 2*w*h + findSmallestSide(l,w,h)
}

func (d Day02) PartOne(data []byte) interface{} {
	acc := 0
	for _, v := range input.ByteToStringSlice(data) {
		dim := strings.Split(v, "x")
		acc += calcPaper(dim)
	}
	return acc
}

func (d Day02) PartTwo(data []byte) interface{} {
	acc := 0
	for _, v := range input.ByteToStringSlice(data) {
		dim := strings.Split(v, "x")
		acc += calcRibbon(dim)
	}
	return acc
}
