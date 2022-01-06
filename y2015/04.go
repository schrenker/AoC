package y2015

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/schrenker/go_aoc/tools"
)

type DayFour struct{}

func (d DayFour) PartOne() {
	input := tools.ReadFileString("input/2015/04.txt")
	for i := 0; ; i++ {
		hash := tools.GetMD5(strings.TrimSpace(input) + strconv.Itoa(i))
		if hash[:5] == "00000" {
			fmt.Println(i)
			break
		}
	}
}

func (d DayFour) PartTwo() {

}
