package y2015

import (
	"bytes"
	"strconv"
	"strings"

	"github.com/schrenker/AoC/tools"
)

type Day15 struct{}

func parseRecipe(s string) (string, []int) {
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
	return name, nums
}

func newIngredientLight(s string) (string, map[string]int) {
	name, nums := parseRecipe(s)
	return name, map[string]int{
		"capacity":   nums[0],
		"durability": nums[1],
		"flavor":     nums[2],
		"texture":    nums[3],
	}
}

func newIngredientNutritious(s string) (string, map[string]int) {
	name, nums := parseRecipe(s)
	return name, map[string]int{
		"capacity":   nums[0],
		"durability": nums[1],
		"flavor":     nums[2],
		"texture":    nums[3],
		"calories":   nums[4],
	}
}

func makeCookie(cookie map[string]int, ingredients map[string]map[string]int, properties []string) int {
	result := 1
	for _, v := range properties {
		tmp := 0
		for j := range cookie {
			tmp += cookie[j] * ingredients[j][v]
		}
		if tmp <= 0 {
			continue
		}
		result *= tmp
	}
	return result
}

func makeNutritiousCookie(cookie map[string]int, ingredients map[string]map[string]int, properties []string) int {
	result := 1
	cal := 0
	for _, v := range properties {
		tmp := 0
		for j := range cookie {
			if v == "calories" {
				cal += cookie[j] * ingredients[j][v]
			} else {
				tmp += cookie[j] * ingredients[j][v]
			}
		}
		if tmp <= 0 {
			continue
		}
		result *= tmp
	}
	if cal != 500 {
		return 0
	}
	return result
}

func (d Day15) PartOne(data []byte) interface{} {
	str := tools.ByteToStringSlice(bytes.TrimSpace(data))
	ingredients := make(map[string]map[string]int)
	ingredientList := []string{}
	properties := make([]string, 0)

	for i := range str {
		name, ing := newIngredientLight(str[i])
		ingredients[name] = ing
		ingredientList = append(ingredientList, name)
	}

	for i := range ingredients[ingredientList[0]] {
		properties = append(properties, i)
	}

	result := 0
	for i := 0; i <= 100; i++ {
		for j := 0; j <= 100; j++ {
			for k := 0; k <= 100; k++ {
				for l := 0; l <= 100; l++ {
					if i+j+k+l != 100 {
						continue
					} else {
						result = tools.GetMax(result, makeCookie(map[string]int{
							ingredientList[0]: i,
							ingredientList[1]: j,
							ingredientList[2]: k,
							ingredientList[3]: l,
						}, ingredients, properties))
					}
				}
			}
		}
	}

	return result
}

func (d Day15) PartTwo(data []byte) interface{} {
	str := tools.ByteToStringSlice(bytes.TrimSpace(data))
	ingredients := make(map[string]map[string]int)
	ingredientList := []string{}
	properties := make([]string, 0)

	for i := range str {
		name, ing := newIngredientNutritious(str[i])
		ingredients[name] = ing
		ingredientList = append(ingredientList, name)
	}

	for i := range ingredients[ingredientList[0]] {
		properties = append(properties, i)
	}

	result := 0
	for i := 0; i <= 100; i++ {
		for j := 0; j <= 100; j++ {
			for k := 0; k <= 100; k++ {
				for l := 0; l <= 100; l++ {
					if i+j+k+l != 100 {
						continue
					} else {
						result = tools.GetMax(result, makeNutritiousCookie(map[string]int{
							ingredientList[0]: i,
							ingredientList[1]: j,
							ingredientList[2]: k,
							ingredientList[3]: l,
						}, ingredients, properties))
					}
				}
			}
		}
	}

	return result
}
