package y2015

import (
	"encoding/json"
	"fmt"
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
	return procSlice(jsonDump)
}

func procSlice(slc []interface{}) float64 {
	var acc float64 = 0

	for i := range slc {

		switch slc[i].(type) {
		case []interface{}:
			acc += procSlice(slc[i].([]interface{}))
		case map[string]interface{}:
			fmt.Println("mep")
		default:
			acc += procNumber(slc[i])
		}

	}
	return acc
}

func procNumber(num interface{}) float64 {

	switch num.(type) {
	case int:
		return float64(num.(int))
	case float64:
		return num.(float64)
	}
	return 0
}

// func procMap(mp map[string]interface{}) int {

// }
