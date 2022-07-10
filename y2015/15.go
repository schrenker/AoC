package y2015

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/schrenker/AoC/tools"
)

type Day15 struct{}

func newIngredientLight(s string) (string, map[string]int) {
	s = strings.ReplaceAll(s, ",", "")
	s = strings.ReplaceAll(s, ":", "")
	name := strings.Split(s, " ")[0]
	nums := tools.Map(
		tools.Filter(
			strings.Split(s, " "),
			func(s string) bool {
				if _, err := strconv.Atoi(s); err != nil {
					return false
				}
				return true
			}),
		func(s string) int {
			num, _ := strconv.Atoi(s) // no need to check for errors here, since I did a function ago
			return num
		})

	return name, map[string]int{
		"capacity":   nums[0],
		"durability": nums[1],
		"flavor":     nums[2],
		"texture":    nums[3],
	}
}

func calculateCurrentScore(cookie map[string]int, ingredients map[string]map[string]int) int {
	result := 1
	properties := make([]string, 0)
	for k := range ingredients["Candy"] {
		properties = append(properties, k)
	}
	for _, v := range properties {
		tmp := 0
		for j := range cookie {
			tmp += cookie[j] * ingredients[j][v]
		}
		fmt.Println(tmp)
		if tmp <= 0 {
			continue
		}
		result *= tmp
	}
	return result
}

func (d Day15) PartOne() interface{} {
	data := tools.ReadFileStringSlice()
	ingredients := make(map[string]map[string]int)
	cookie := make(map[string]int)
	for i := range data {
		name, ing := newIngredientLight(data[i])
		ingredients[name] = ing
	}

	for k := range ingredients {
		cookie[k] = 0
	}
	cookie["Butterscotch"] = 44
	cookie["Candy"] = 56
	fmt.Println(cookie)
	fmt.Println(ingredients)
	fmt.Println(calculateCurrentScore(cookie, ingredients))
	return 0
}

func (d Day15) PartTwo() interface{} {
	return 0
}
