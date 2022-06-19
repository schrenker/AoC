package y2015

import (
	"encoding/json"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/schrenker/AoC/tools"
)

type Day12 struct{}

func procSlice(slc []interface{}) float64 {
	var acc float64 = 0

	for i := range slc {

		switch slc[i].(type) {
		case []interface{}:
			acc += procSlice(slc[i].([]interface{}))
		case map[string]interface{}:
			acc += procMap(slc[i].(map[string]interface{}))
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

func procMap(m map[string]interface{}) float64 {
	var acc float64 = 0

	for _, v := range m {

		switch v.(type) {
		case []interface{}:
			acc += procSlice(v.([]interface{}))
		case map[string]interface{}:
			acc += procMap(v.(map[string]interface{}))
		case string:
			if v.(string) == "red" {
				return 0
			}
		default:
			acc += procNumber(v)
		}

	}
	return acc
}

func (d Day12) PartOne() interface{} {
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

func (d Day12) PartTwo() interface{} {
	data := tools.ReadFileBytes("input/2015/12.txt")
	var jsonDump []interface{}
	err := json.Unmarshal(data, &jsonDump)
	if err != nil {
		log.Fatalln(err)
	}
	return procSlice(jsonDump)
}
