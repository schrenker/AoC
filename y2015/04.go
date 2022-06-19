package y2015

import (
	"strconv"
	"strings"

	"github.com/schrenker/AoC/tools"
)

type Day04 struct{}

func mineAdventCoin(comparator string) int {
	input := tools.ReadFileString("input/2015/04.txt")
	for i := 0; ; i++ {
		hash := tools.GetMD5(strings.TrimSpace(input) + strconv.Itoa(i))
		if hash[:len(comparator)] == comparator {
			return i
		}
	}
}

func (d Day04) PartOne() interface{} {
	return mineAdventCoin("00000")
}

func (d Day04) PartTwo() interface{} {
	return mineAdventCoin("000000")
}
