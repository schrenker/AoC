package y2015

import (
	"fmt"
	"strings"

	"github.com/schrenker/AoC/tools"
)

type Day19 struct{}

func (d Day19) PartOne() interface{} {
	data := tools.ReadFileStringSlice()
	src := data[len(data)-1]
	replacements := make(map[string][]string)
	uniq := make(map[string]bool)

	for i := 0; i < len(data)-2; i++ {
		tmp := strings.Split(data[i], " => ")
		replacements[tmp[0]] = append(replacements[tmp[0]], tmp[1])
	}

	for i := range src {
		if tmp, ok := replacements[string(src[i])]; ok {
			for _, j := range tmp {
				uniq[fmt.Sprintf("%v%v%v", src[:i], j, src[i+1:])] = true
			}
		}
		if i+2 < len(src) {
			if tmp, ok := replacements[src[i:i+2]]; ok {
				for _, j := range tmp {
					uniq[fmt.Sprintf("%v%v%v", src[:i], j, src[i+2:])] = true
				}
			}

		}
	}
	return len(uniq)
}

func transform(molecule string, step int, replacements map[string][]string, processed map[string]bool) int {
	fmt.Printf("Processing molecule: %v\n", molecule)
	if _, ok := replacements[molecule]; ok {
		return step
	}

	processed[molecule] = true

	min := 99999
	for i := 0; i < len(molecule); i++ {
		for j := i + 1; j <= len(molecule); j++ {
			s, err := deconstruct(molecule[i:j], replacements)
			if err == nil {
				if _, ok := processed[molecule[:i]+s+molecule[j:]]; !ok {
					min = tools.GetMin(min, transform(fmt.Sprintf("%v%v%v", molecule[:i], s, molecule[j:]), step+1, replacements, processed))
				}

			}
		}
	}
	return min
}

func deconstruct(substr string, replacements map[string][]string) (string, error) {
	for i, v := range replacements {
		for _, s := range v {
			if substr == s {
				return i, nil
			}
		}
	}
	return "", fmt.Errorf("not found")
}

func (d Day19) PartTwo() interface{} {
	data := tools.ReadFileStringSlice()
	src := data[len(data)-1]
	replacements := make(map[string][]string)
	processed := make(map[string]bool)

	for i := 0; i < len(data)-2; i++ {
		tmp := strings.Split(data[i], " => ")
		replacements[tmp[0]] = append(replacements[tmp[0]], tmp[1])
	}
	return transform(src, 1, replacements, processed)
}
