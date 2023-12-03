package y2023

import (
	"strconv"

	"github.com/schrenker/AoC/input"
)

type Day03 struct{}

type gearCoords struct {
	x int
	y int
}

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

func containsGear(gearBuffer []gearCoords, gear gearCoords) bool {
	for _, g := range gearBuffer {
		if gear.x == g.x && gear.y == g.y {
			return true
		}
	}
	return false
}

func (d Day03) PartOne(data []byte) interface{} {
	acc := 0
	inp := input.ByteToStringSlice(data)
	var buffer []byte
	isPartNumber := false
	for i := range inp {
		for j := range inp[i] {
			if isNumber(inp[i][j]) {
				buffer = append(buffer, inp[i][j])
				if !isPartNumber {
					isPartNumber = checkSurroundingsForSymbol(inp, i, j)
				}
				continue
			}
			if buffer != nil {
				number, _ := strconv.Atoi(string(buffer))
				if isPartNumber {
					acc += number
				}
				isPartNumber = false
				buffer = nil
			}
		}
	}

	return acc
}

func (d Day03) PartTwo(data []byte) interface{} {
	gearMap := make(map[gearCoords][]int)
	acc := 0
	inp := input.ByteToStringSlice(data)
	var numberBuffer []byte
	var gearBuffer []gearCoords

	for i := range inp {
		for j := range inp[i] {

			if isNumber(inp[i][j]) {
				numberBuffer = append(numberBuffer, inp[i][j])
				for _, gear := range checkSurroundingsForGear(inp, i, j) {
					if !containsGear(gearBuffer, gear) {
						gearBuffer = append(gearBuffer, gear)
					}
				}
				continue
			}

			if numberBuffer != nil {
				num, _ := strconv.Atoi(string(numberBuffer))
				for _, gear := range gearBuffer {
					gearMap[gear] = append(gearMap[gear], num)
				}
				numberBuffer = nil
				gearBuffer = nil
			}
		}
	}

	for _, v := range gearMap {
		if len(v) == 2 {
			acc += v[0] * v[1]
		}
	}
	return acc
}
