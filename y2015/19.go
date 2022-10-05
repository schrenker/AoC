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

func (d Day19) PartTwo() interface{} {
	return 0
}
