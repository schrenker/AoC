package y2015

import (
	"fmt"
	"strings"

	"github.com/schrenker/go_aoc/tools"
)

type DayEight struct{}

func getStringLength(str string) int {
	acc := 0
	var escapeMap = map[byte]int{
		'\\': 1,
		'"':  1,
		'x':  3,
	}
	for i := 0; i < len(str); i++ {
		if str[i] == '\\' {
			i += escapeMap[str[i+1]]
		}
		acc++
	}
	return acc - 2
}

func getCodeLength(str string) int {
	acc := 0
	encodeChars := "\\\""
	for i := 0; i < len(str); i++ {
		if strings.Contains(encodeChars, string(str[i])) {
			acc++
		}
		acc++
	}
	return acc + 2
}

func (d DayEight) PartOne() {
	data := tools.ReadFileStringSlice("input/2015/08.txt")
	codeLength := 0
	stringLength := 0
	for _, v := range data {
		codeLength += len(v)
		stringLength += getStringLength(v)
	}
	fmt.Println(codeLength - stringLength)
}

func (d DayEight) PartTwo() {
	data := tools.ReadFileStringSlice("input/2015/08.txt")
	codeLength := 0
	actualCodeLength := 0
	for _, v := range data {
		codeLength += len(v)
		actualCodeLength += getCodeLength(v)
	}
	fmt.Println(actualCodeLength - codeLength)
}
