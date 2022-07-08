package y2015

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/schrenker/AoC/tools"
)

type Day15 struct{}

type ingredient struct {
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func newIngredient(s string) ingredient {
	s = strings.ReplaceAll(s, ",", "")
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

	return ingredient{
		capacity:   nums[0],
		durability: nums[1],
		flavor:     nums[2],
		texture:    nums[3],
		calories:   nums[4],
	}

}

func (d Day15) PartOne() interface{} {
	data := tools.ReadFileStringSlice()
	base := &ingredient{}
	ingredients := []*ingredient{}
	for i := range data {
		ing := newIngredient(data[i])
		ingredients = append(ingredients, &ing)
	}
	fmt.Println(base)
	return 0
}

func (d Day15) PartTwo() interface{} {
	return 0
}
