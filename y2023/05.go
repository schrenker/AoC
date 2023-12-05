package y2023

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/schrenker/AoC/input"
)

type Day05 struct{}

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
	for _, v := range mapOrder {
		tmp, ok := m[v][loc]
		if ok {
			loc = tmp
		}
	}
	return loc
}

func (d Day05) PartOne(data []byte) interface{} {
	inp := input.ByteToStringSlice(data)
	spacePattern = regexp.MustCompile(`\s`)
	seeds := parseSeeds(inp[0])
	maps := make(map[string]map[int]int)
	for _, v := range mapOrder {
		maps[v] = make(map[int]int)
	}



	processed := 0
	for i := 3; i < len(inp); i++ {
		if inp[i] == "" {
			i++
			processed++
			continue
		}
		ranges := spacePattern.Split(strings.TrimSpace(inp[i]), -1)
		destination, _ := strconv.Atoi(ranges[0])
		source, _ := strconv.Atoi(ranges[1])
		length, _ := strconv.Atoi(ranges[2])
		for j := 0; j < length; j++ {
			maps[mapOrder[processed]][source+j] = destination + j
		}

	}

	fmt.Println("maps built")

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
