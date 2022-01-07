package y2015

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/schrenker/go_aoc/tools"
)

type DaySix struct{}

type boolGrid [][]bool

func makeBoolGrid() boolGrid {
	grid := make(boolGrid, 1000, 1000)
	for i := range grid {
		grid[i] = make([]bool, 1000, 1000)
		for j := range grid[i] {
			grid[i][j] = false
		}
	}
	return grid
}

func (b *boolGrid) turnOff(pos []int) {

	for i := pos[0]; i <= pos[2]; i++ {
		for j := pos[1]; j <= pos[3]; j++ {
			(*b)[i][j] = false
		}
	}
}

func (b *boolGrid) turnOn(pos []int) {

	for i := pos[0]; i <= pos[2]; i++ {
		for j := pos[1]; j <= pos[3]; j++ {
			(*b)[i][j] = true
		}
	}
}

func (b *boolGrid) toggle(pos []int) {

	for i := pos[0]; i <= pos[2]; i++ {
		for j := pos[1]; j <= pos[3]; j++ {

			(*b)[i][j] = !(*b)[i][j]
		}
	}
}

func (b boolGrid) countLights() int {
	acc := 0
	for _, v := range b {
		for _, vv := range v {
			if vv {
				acc++
			}
		}
	}
	return acc
}

func (b *boolGrid) parseCommand(command string) {
	tmp := strings.Split(command, " ")
	if tmp[0] == "toggle" {
		b.toggle(parsePos(tmp[1], tmp[3]))
	} else if tmp[0] == "turn" {
		if tmp[1] == "on" {
			b.turnOn(parsePos(tmp[2], tmp[4]))
		} else {
			b.turnOff(parsePos(tmp[2], tmp[4]))
		}
	}
}

type intGrid [][]int

func makeIntGrid() intGrid {
	grid := make(intGrid, 1000, 1000)
	for i := range grid {
		grid[i] = make([]int, 1000, 1000)
		for j := range grid[i] {
			grid[i][j] = 0
		}
	}
	return grid
}

func (g *intGrid) turnOff(pos []int) {

	for i := pos[0]; i <= pos[2]; i++ {
		for j := pos[1]; j <= pos[3]; j++ {
			if (*g)[i][j] > 0 {
				(*g)[i][j]--
			}
		}
	}
}

func (g *intGrid) turnOn(pos []int) {

	for i := pos[0]; i <= pos[2]; i++ {
		for j := pos[1]; j <= pos[3]; j++ {
			(*g)[i][j]++
		}
	}
}

func (g *intGrid) toggle(pos []int) {

	for i := pos[0]; i <= pos[2]; i++ {
		for j := pos[1]; j <= pos[3]; j++ {

			(*g)[i][j] += 2
		}
	}
}

func (g intGrid) countLights() int {
	acc := 0
	for _, v := range g {
		for _, vv := range v {
			acc += vv
		}
	}
	return acc
}

func (g *intGrid) parseCommand(command string) {
	tmp := strings.Split(command, " ")
	if tmp[0] == "toggle" {
		g.toggle(parsePos(tmp[1], tmp[3]))
	} else if tmp[0] == "turn" {
		if tmp[1] == "on" {
			g.turnOn(parsePos(tmp[2], tmp[4]))
		} else {
			g.turnOff(parsePos(tmp[2], tmp[4]))
		}
	}
}

func parsePos(start, finish string) []int {
	tmpSta := strings.Split(start, ",")
	tmpFin := strings.Split(finish, ",")
	xs, _ := strconv.Atoi(tmpSta[0])
	ys, _ := strconv.Atoi(tmpSta[1])
	xf, _ := strconv.Atoi(tmpFin[0])
	yf, _ := strconv.Atoi(tmpFin[1])
	return []int{xs, ys, xf, yf}
}

func (d DaySix) PartOne() {
	data := tools.ReadFileStringSlice("input/2015/06.txt")
	grid := makeBoolGrid()
	for _, v := range data {
		grid.parseCommand(v)
	}
	fmt.Println(grid.countLights())
}

func (d DaySix) PartTwo() {
	data := tools.ReadFileStringSlice("input/2015/06.txt")
	grid := makeIntGrid()
	for _, v := range data {
		grid.parseCommand(v)
	}
	fmt.Println(grid.countLights())
}
