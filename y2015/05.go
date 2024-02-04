package y2015

import (
	"strings"

	"github.com/schrenker/AoC/input"
)

type Day05 struct{}

var vowels = "aeiou"
var forbidden = []string{"ab", "cd", "pq", "xy"}

func hasForbidden(str string) bool {
	for _, v := range forbidden {
		if strings.Contains(str, v) {
			return true
		}
	}
	return false
}

func hasThreeVowels(str string) bool {
	acc := 0
	for _, i := range str {
		for _, j := range vowels {
			if i == j {
				acc++
			}
		}
	}
	return acc >= 3
}

func hasDoubleLetters(str string) bool {
	for i := 0; i < len(str)-1; i++ {
		if str[i] == str[i+1] {
			return true
		}
	}
	return false
}

func (d Day05) PartOne(data []byte) interface{} {
	acc := 0

	for _, v := range input.ByteToStringSlice(data) {
		if !hasForbidden(v) && hasThreeVowels(v) && hasDoubleLetters(v) {
			acc++
		}
	}

	return acc
}

func hasAlternateLetters(str string) bool {
	for i := 0; i < len(str)-2; i++ {
		if str[i] == str[i+2] {
			return true
		}
	}
	return false
}

func hasPairs(str string) bool {
	pairs := make(map[string][]int)
	for i := 0; i < len(str)-1; i++ {
		pairs[str[i:i+2]] = append(pairs[str[i:i+2]], i)
	}
	for _, v := range pairs {
		if len(v) > 1 {
			for i := 0; i < len(v)-1; i++ {
				if v[i+1]-v[i] > 1 {
					return true
				}
			}
		}
	}
	return false
}

func (d Day05) PartTwo(data []byte) interface{} {
	acc := 0

	for _, v := range input.ByteToStringSlice(data) {
		if hasAlternateLetters(v) && hasPairs(v) {
			acc++
		}
	}

	return acc
}
