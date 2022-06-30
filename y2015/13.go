package y2015

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/schrenker/AoC/tools"
)

type Day13 struct{}

func parseInput(input []string) map[string]map[string]int {
	result := make(map[string]map[string]int)

	for i := range input {
		re, err := regexp.Compile(`[^\w]`)
		if err != nil {
			log.Fatalln(err)
		}

		tmp := strings.Split(re.ReplaceAllString(input[i], " "), " ")
		if result[tmp[0]] == nil {
			result[tmp[0]] = make(map[string]int)
		}

		result[tmp[0]][tmp[len(tmp)-2]], err = strconv.Atoi(tmp[3])
		if err != nil {
			log.Fatalln(err)
		}

		if tmp[2] == "lose" {
			result[tmp[0]][tmp[len(tmp)-2]] *= -1
		}
	}

	return result
}

func insertYourself(guestMap map[string]map[string]int) map[string]map[string]int {
	you := make(map[string]int)
	for k := range guestMap {
		you[k] = 0
		guestMap[k]["you"] = 0
	}
	guestMap["you"] = you
	return guestMap
}

func computeHappiness(order []string, scores map[string]map[string]int) int {
	acc := 0
	for i := range order {
		acc = acc + scores[order[i]][order[tools.PMod(i-1, len(order))]] + scores[order[i]][order[tools.PMod(i+1, len(order))]]
	}
	return acc
}

func (d Day13) PartOne() interface{} {
	guestMap := parseInput(tools.ReadFileStringSlice("input/2015/13.txt"))
	guests := make([]string, 0)
	for k := range guestMap {
		if len(k) > 0 {
			guests = append(guests, k)
		}
	}

	guestPermutations := tools.StringPermutations(guests)
	highest := -9999
	for i := range guestPermutations {
		highest = tools.GetMax(computeHappiness(guestPermutations[i], guestMap), highest)
	}

	return highest
}

func (d Day13) PartTwo() interface{} {
	guestMap := parseInput(tools.ReadFileStringSlice("input/2015/13.txt"))
	guestMap = insertYourself(guestMap)
	guests := make([]string, 0)
	for k := range guestMap {
		if len(k) > 0 {
			guests = append(guests, k)
		}
	}

	guestPermutations := tools.StringPermutations(guests)
	highest := -9999
	for i := range guestPermutations {
		highest = tools.GetMax(computeHappiness(guestPermutations[i], guestMap), highest)
	}

	return highest
}
