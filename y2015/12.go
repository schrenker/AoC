package y2015

import (
	"encoding/json"
	"log"
	"regexp"
	"strconv"
	"strings"
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

	switch num := num.(type) {
	case int:
		return float64(num)
	case float64:
		return num
	}
	return 0
}

func procMap(m map[string]interface{}) float64 {
	var acc float64 = 0

	for _, v := range m {

		switch v := v.(type) {
		case []interface{}:
			acc += procSlice(v)
		case map[string]interface{}:
			acc += procMap(v)
		case string:
			if v == "red" {
				return 0
			}
		default:
			acc += procNumber(v)
		}

	}
	return acc
}

func (d Day12) PartOne(data []byte) interface{} {
	re := regexp.MustCompile(`[^\d\-,]`)
	spl := strings.Split(re.ReplaceAllString(string(data), ""), ",")
	acc := 0
	for _, v := range spl {
		conv, _ := strconv.Atoi(v)
		acc += conv
	}
	return acc
}

func (d Day12) PartTwo(data []byte) interface{} {
	var jsonDump []interface{}
	err := json.Unmarshal(data, &jsonDump)
	if err != nil {
		log.Fatalln(err)
	}
	return procSlice(jsonDump)
}
