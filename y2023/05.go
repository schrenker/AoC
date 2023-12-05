package y2023

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/schrenker/AoC/input"
)

type Day05 struct{}

type pair struct {
	srcStart int
	srcEnd int
	dstStart int
	dstEnd int
	rg int
}

var spacePattern *regexp.Regexp
var mapOrder = []string{"soil", "fertilizer", "water", "light", "temperature", "humidity", "location"}

func parseSeeds(s string) []int {
	result := make([]int, 0)
	for _, v := range spacePattern.Split(strings.TrimLeft(s, "seeds: "), -1) {
		tmp, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		result = append(result, tmp)
	}
	return result
}

func findLocation(seed int, m map[string]map[int]int) int {
	loc := seed
	for i := range mapOrder {
		for _, p := range m[mapOrder[i]] {
			if loc >= p.srcStart && loc <= p.srcEnd {
				tmp := loc - p.srcStart
				fmt.Printf("v: %v, loc: %v, p.srcStart: %v, tmp: %v, newloc: %v\n", mapOrder[i], loc, p.srcStart, tmp, p.dstStart + tmp)
				loc = p.dstStart + tmp
				break
			}
		}
	}
	return loc
}

func (d Day05) PartOne(data []byte) interface{} {
	inp := input.ByteToStringSlice(data)
	spacePattern = regexp.MustCompile(`\s`)
	seeds := parseSeeds(inp[0])
	maps := make(map[string][]*pair)
	for _, v := range mapOrder {
		maps[v] = make([]*pair, 0)
	}

	processed := 0
	for i := 3; i < len(inp); i++ {
		if inp[i] == "" {
			i++
			processed++
			continue
		}
		ranges := spacePattern.Split(strings.TrimSpace(inp[i]), -1)
		rg, _ := strconv.Atoi(ranges[2])
		srcStart, _ := strconv.Atoi(ranges[1])
		srcEnd := srcStart + rg - 1
		dstStart, _ := strconv.Atoi(ranges[0])
		dstEnd := dstStart + rg - 1
		fmt.Printf("rg: %v, srcStart: %v, srcEnd: %v, dstStart: %v, dstEnd: %v\n", rg, srcStart, srcEnd, dstStart, dstEnd)
		maps[mapOrder[processed]] = append(maps[mapOrder[processed]], &pair{
			srcStart: srcStart,
			srcEnd: srcEnd,
			dstStart: dstStart,
			dstEnd: dstEnd,
			rg: rg,
		})

	}

	min := 999999999999999
	for _, v := range seeds {
		tmp := findLocation(v, maps)
		if tmp < min {
			min = tmp
		}
	}

	return min
}

func (d Day05) PartTwo(data []byte) interface{} {
	return 0
}
