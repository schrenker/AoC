package y2015

import (
	"encoding/json"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/schrenker/AoC/tools"
)

type DayTwelve struct{}

func (d DayTwelve) PartOne() interface{} {
	data := tools.ReadFileString("input/2015/12.txt")
	re := regexp.MustCompile(`[^\d\-,]`)
	spl := strings.Split(re.ReplaceAllString(data, ""), ",")
	acc := 0
	for _, v := range spl {
		conv, _ := strconv.Atoi(v)
		acc += conv
	}
	return acc
}

func (d DayTwelve) PartTwo() interface{} {
	data := tools.ReadFileBytes("input/2015/12.txt")
	var jsonDump []interface{}
	err := json.Unmarshal(data, &jsonDump)
	if err != nil {
		log.Fatalln(err)
	}
	return jsonDump[0]
}
