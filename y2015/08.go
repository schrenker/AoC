package y2015

import (
	"strings"

	"github.com/schrenker/AoC/tools"
)

type Day08 struct{}

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

func (d Day08) PartOne(data []byte) interface{} {
	codeLength := 0
	stringLength := 0
	for _, v := range tools.ByteToStringSlice(data) {
		codeLength += len(v)
		stringLength += getStringLength(v)
	}
	return codeLength - stringLength
}

func (d Day08) PartTwo(data []byte) interface{} {
	codeLength := 0
	actualCodeLength := 0
	for _, v := range tools.ByteToStringSlice(data) {
		codeLength += len(v)
		actualCodeLength += getCodeLength(v)
	}
	return actualCodeLength - codeLength
}
