package y2015

import (
	"strconv"
	"strings"

	"github.com/schrenker/AoC/tools"
)

type Day04 struct{}

func mineAdventCoin(comparator string, data []byte) int {
	input := string(data)
	for i := 0; ; i++ {
		hash := tools.GetMD5(strings.TrimSpace(input) + strconv.Itoa(i))
		if hash[:len(comparator)] == comparator {
			return i
		}
	}
}

func (d Day04) PartOne(data []byte) interface{} {
	return mineAdventCoin("00000", data)
}

func (d Day04) PartTwo(data []byte) interface{} {
	return mineAdventCoin("000000", data)
}
