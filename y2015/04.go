package y2015

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/schrenker/go_aoc/tools"
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

func (d DayFour) PartOne() {
	fmt.Println(mineAdventCoin("00000"))
}

func (d DayFour) PartTwo() {
	fmt.Println(mineAdventCoin("000000"))
}
