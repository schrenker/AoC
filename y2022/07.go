package y2022

import (
	"log"
	"strconv"
	"strings"

	"github.com/schrenker/AoC/tools"
)

type Day07 struct{}

func duSh(term []string) map[string]int {
	dirSizes := make(map[string]int)
	dirStack := tools.NewStack[string]()
	for _, v := range term {
		line := strings.Fields(v)
		switch line[0] {
		case "$":
			if line[1] == "cd" {
				if line[2] == ".." {
					if dirStack.Length() > 1 {
						dirStack.Pop()
					}
				} else {
					dirStack.Push(line[2])
				}
			}

		case "dir":

		default:
			size, err := strconv.Atoi(line[0])
			if err != nil {
				log.Fatalln(err)
			}
			path := ""
			for _, d := range dirStack.List() {
				path += d
				dirSizes[path] += size
			}

		}
	}
	return dirSizes
}

func (d Day07) PartOne(data []byte) interface{} {
	term := tools.ByteToStringSlice(data)

	dirSizes := duSh(term)

	size := 0

	for _, v := range dirSizes {
		if v <= 100000 {
			size += v
		}
	}
	return size
}

func (d Day07) PartTwo(data []byte) interface{} {
	term := tools.ByteToStringSlice(data)
	dirSizes := duSh(term)
	target := 30000000 - (70000000 - dirSizes["/"])
	rmCandidates := make([]int, 0)

	for _, v := range dirSizes {
		if v > target {
			rmCandidates = append(rmCandidates, v)
		}
	}

	return tools.GetMin(rmCandidates...)
}
