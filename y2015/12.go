package y2015

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/schrenker/go_aoc/tools"
)

type DayTwelve struct{}

func (d DayTwelve) PartOne() {
	data := tools.ReadFileString("input/2015/12.txt")
	re := regexp.MustCompile(`[^\d\-,]`)
	spl := strings.Split(re.ReplaceAllString(data, ""), ",")
	acc := 0
	for _, v := range spl {
		conv, _ := strconv.Atoi(v)
		acc += conv
	}
	fmt.Println(acc)
}

func (d DayTwelve) PartTwo() {
}