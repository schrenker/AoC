package y2023

import (
	"strconv"

	"github.com/schrenker/AoC/input"
)

type Day03 struct{}

func isSymbolNoDot(b byte) bool {
	return b != 46 &&
		((b >= 33 && b <= 47) || (b >= 58 && b <= 64))
}

func isNumber(b byte) bool {
	return b >= 48 && b <= 57
}

func checkSurroundingsForSymbol(inp []string, x, y int) bool {
	coords := [][]int{{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1}}
	for i := range coords {
		tmpx := x - coords[i][0]
		tmpy := y - coords[i][1]
		if tmpx < len(inp) && tmpx >= 0 && tmpy < len(inp[tmpx]) && tmpy >= 0 && isSymbolNoDot(inp[tmpx][tmpy]) {
			return true
		}
	}
	return false
}

type gearCoords struct {
	x int
	y int
}

func checkSurroundingsForGear(inp []string, x, y int) []gearCoords {
	coords := [][]int{{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1}}
	gears := make([]gearCoords, 0)
	for i := range coords {
		tmpx := x - coords[i][0]
		tmpy := y - coords[i][1]
		if tmpx < len(inp) && tmpx >= 0 && tmpy < len(inp[tmpx]) && tmpy >= 0 && inp[tmpx][tmpy] == '*' {
			gears = append(gears, gearCoords{x: tmpx, y: tmpy})
		}
	}
	return gears
}

func containsGear(gs []gearCoords, g gearCoords) bool {
	for _,i := range gs {
		if g.x == i.x && g.y == i.y {
			return true
		}
	}
	return false
}

func (d Day03) PartOne(data []byte) interface{} {
	acc := 0
	inp := input.ByteToStringSlice(data)
	var currNum []byte
	isPartNumber := false
	for i := range inp {
		for j := range inp[i] {
			if isNumber(inp[i][j]) {
				currNum = append(currNum, inp[i][j])
				if !isPartNumber {
					isPartNumber = checkSurroundingsForSymbol(inp, i, j)
				}
				continue
			}
			if currNum != nil {
				num, _ := strconv.Atoi(string(currNum))
				if isPartNumber {
					acc += num
				}
				isPartNumber = false
				currNum = nil
			}
		}
	}

	return acc
}

func (d Day03) PartTwo(data []byte) interface{} {
	gears := make(map[gearCoords][]int)
	acc := 0
	inp := input.ByteToStringSlice(data)
	var currNum []byte
	var currGear []gearCoords

	for i := range inp {
		for j := range inp[i] {

			if isNumber(inp[i][j]) {
				currNum = append(currNum, inp[i][j])
				for _,v := range checkSurroundingsForGear(inp, i, j) {
					if !containsGear(currGear, v) {
						currGear = append(currGear, v)
					}
				}
				continue
			}

			if currNum != nil {
				num, _ := strconv.Atoi(string(currNum))
				for _, v := range currGear {
					gears[v] = append(gears[v], num)
				}
				currNum = nil
				currGear = nil
			}
		}
	}

	for _,v := range gears {
		if len(v) == 2 {
			acc += v[0] * v[1]
		}
	}
	return acc
}
