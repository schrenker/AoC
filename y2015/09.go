package y2015

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/schrenker/go_aoc/tools"
)

type DayNine struct{}

type edge struct {
	start string
	stop  string
}

func genUniqueArray(s []string) []string {
	umap := make(map[string]bool)
	for _, v := range s {
		spl := strings.Split(v, " ")
		umap[spl[0]] = true
		umap[spl[2]] = true
	}
	result := make([]string, 0)
	for k := range umap {
		result = append(result, k)
	}
	return result
}

func genDistanceMap(s []string) map[edge]int {
	distMap := make(map[edge]int)
	for _, v := range s {
		spl := strings.Split(v, " ")
		dist, _ := strconv.Atoi(spl[4])
		distMap[edge{start: spl[0], stop: spl[2]}] = dist
		distMap[edge{start: spl[2], stop: spl[0]}] = dist
	}
	return distMap
}

func getPermDistance(s []string, dist map[edge]int) int {
	acc := 0
	for i := 0; i < len(s)-1; i++ {
		acc += dist[edge{start: s[i], stop: s[i+1]}]
	}
	return acc
}

func (d DayNine) PartOne() {
	data := tools.ReadFileStringSlice("input/2015/09.txt")
	perms := tools.StringPermutations(genUniqueArray(data))
	distMap := genDistanceMap(data)
	min := math.MaxInt
	for _, v := range perms {
		if tmp := getPermDistance(v, distMap); tmp < min {
			min = tmp
		}
	}
	fmt.Println(min)
}

func (d DayNine) PartTwo() {
	data := tools.ReadFileStringSlice("input/2015/09.txt")
	perms := tools.StringPermutations(genUniqueArray(data))
	distMap := genDistanceMap(data)
	max := math.MinInt
	for _, v := range perms {
		if tmp := getPermDistance(v, distMap); tmp > max {
			max = tmp
		}
	}
	fmt.Println(max)
}