package y2015

import (
	"fmt"
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
		fmt.Println(tmp)
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

func (d Day13) PartOne() interface{} {
	guests := parseInput(tools.ReadFileStringSlice("input/2015/13.txt"))
	return guests
}

func (d Day13) PartTwo() interface{} {
	return nil
}
