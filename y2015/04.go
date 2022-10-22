package y2015

import (
	"strconv"
	"strings"

	"github.com/schrenker/AoC/tools"
)

type Day04 struct{}

func mineAdventCoin(comparator string, path string) int {
	input := tools.ReadFileString(path)
	for i := 0; ; i++ {
		hash := tools.GetMD5(strings.TrimSpace(input) + strconv.Itoa(i))
		if hash[:len(comparator)] == comparator {
			return i
		}
	}
}

func (d Day04) PartOne(path string) interface{} {
	return mineAdventCoin("00000", path)
}

func (d Day04) PartTwo(path string) interface{} {
	return mineAdventCoin("000000", path)
}
