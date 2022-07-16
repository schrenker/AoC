package y2015

import (
	"strconv"
	"strings"

	"github.com/schrenker/AoC/tools"
)

type Day16 struct{}

var base = map[string]int{
	"children":    3,
	"cats":        7,
	"samoyeds":    2,
	"pomeranians": 3,
	"akitas":      0,
	"vizslas":     0,
	"goldfish":    5,
	"trees":       3,
	"cars":        2,
	"perfumes":    1,
}

func unmarshalAunt(line string) map[string]int {
	comp := make(map[string]int)
	tmp := strings.Split(line, " ")
	for i := 2; i < len(tmp); i += 2 {
		num, _ := strconv.Atoi(strings.ReplaceAll(tmp[i+1], ",", ""))
		comp[strings.ReplaceAll(tmp[i], ":", "")] = num
	}
	return comp
}

func matchAunt(base map[string]int, line string) bool {
	comp := unmarshalAunt(line)
	for k, v := range base {
		val, ok := comp[k]
		if ok && val != v {
			return false
		}
	}
	return true
}

func matchAuntRanges(base map[string]int, line string) bool {
	comp := unmarshalAunt(line)
	for k, v := range base {
		val, ok := comp[k]
		if ok {
			switch k {
			case "cats", "trees":
				if val <= v {
					return false
				}
			case "pomeranians", "goldfish":
				if val >= v {
					return false
				}
			default:
				if val != v {
					return false
				}
			}
		}
	}
	return true
}

func (d Day16) PartOne() interface{} {
	data := tools.ReadFileStringSlice()
	matchingAunts := []int{}
	for i := range data {
		if matchAunt(base, data[i]) {
			matchingAunts = append(matchingAunts, i+1)
		}
	}

	if len(matchingAunts) == 1 {
		return matchingAunts[0]
	} else {
		return matchingAunts
	}
}

func (d Day16) PartTwo() interface{} {
	data := tools.ReadFileStringSlice()
	matchingAunts := []int{}
	for i := range data {
		if matchAuntRanges(base, data[i]) {
			matchingAunts = append(matchingAunts, i+1)
		}
	}

	if len(matchingAunts) == 1 {
		return matchingAunts[0]
	} else {
		return matchingAunts
	}
}
