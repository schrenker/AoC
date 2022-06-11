package y2015

import (
	"strconv"
	"strings"

	"github.com/schrenker/AoC/tools"
)

type DayFour struct{}

func mineAdventCoin(comparator string) int {
	input := tools.ReadFileString("input/2015/04.txt")
	for i := 0; ; i++ {
		hash := tools.GetMD5(strings.TrimSpace(input) + strconv.Itoa(i))
		if hash[:len(comparator)] == comparator {
			return i
		}
	}
}

func (d DayFour) PartOne() interface{} {
	return mineAdventCoin("00000")
}

func (d DayFour) PartTwo() interface{} {
	return mineAdventCoin("000000")
}
