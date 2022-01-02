package y2015

import (
	"fmt"

	"github.com/schrenker/go_aoc/tools"
)

type DayThree struct{}

type coords struct {
	x int
	y int
}

func (d DayThree) PartOne() {
	visits := make(map[coords]int)
	x := 0
	y := 0
	visits[coords{x: x, y: y}]++
	for _, v := range tools.ReadFileBytes("input/2015/03.txt") {
		switch v {
		case '>':
			x++
		case '<':
			x--
		case '^':
			y++
		case 'v':
			y--
		}
		visits[coords{x: x, y: y}]++
	}
	fmt.Println(len(visits))
}

func (d DayThree) PartTwo() {
	visits := make(map[coords]int)
	x := []int{0, 0}
	y := []int{0, 0}
	visits[coords{x: x[0], y: y[0]}]++
	for i, v := range tools.ReadFileBytes("input/2015/03.txt") {
		turn := i % 2
		switch v {
		case '>':
			x[turn]++
		case '<':
			x[turn]--
		case '^':
			y[turn]++
		case 'v':
			y[turn]--
		}
		visits[coords{x: x[turn], y: y[turn]}]++
	}
	fmt.Println(len(visits))
}
