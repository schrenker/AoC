package y2022

import (
	"math"
	"strconv"
	"strings"

	"github.com/schrenker/AoC/tools"
)

type Day09 struct{}

type ropeEnd struct {
	x int
	y int
}

type rope []ropeEnd

func (r rope) checkTail() {
nodeLoop:
	for t := 1; t < len(r); t++ {
		x := r[t-1].x - r[t].x
		y := r[t-1].y - r[t].y
		if (int(math.Abs(float64(x))) == 1 && int(math.Abs(float64(y))) == 1) || int(math.Abs(float64(x)))+int(math.Abs(float64(y))) < 2 {
			return
		}
		moves := []int{0, 1, -1}
		for _, i := range moves {
			for _, j := range moves {
				if (x != 0 && i == 0) || (y != 0 && j == 0) {
					continue
				}
				xn := r[t-1].x - (r[t].x + i)
				yn := r[t-1].y - (r[t].y + j)
				if (int(math.Abs(float64(xn))) == 1 && int(math.Abs(float64(yn))) == 1) || int(math.Abs(float64(xn)))+int(math.Abs(float64(yn))) < 2 {
					r[t].x += i
					r[t].y += j
					continue nodeLoop
				}
			}
		}
	}
}

func (d Day09) PartOne(data []byte) interface{} {
	moveList := tools.ByteToStringSlice(data)
	moves := map[string]ropeEnd{
		"R": {x: 1, y: 0},
		"L": {x: -1, y: 0},
		"U": {x: 0, y: 1},
		"D": {x: 0, y: -1},
	}
	visited := map[ropeEnd]bool{
		{x: 0, y: 0}: true,
	}

	rope := make(rope, 0, 2)

	for i := 0; i < 2; i++ {
		rope = append(rope, ropeEnd{x: 0, y: 0})
	}

	for _, m := range moveList {
		tmp := strings.Fields(m)
		direction := tmp[0]
		num, _ := strconv.Atoi(tmp[1])
		for i := 0; i < num; i++ {
			rope[0].x += moves[direction].x
			rope[0].y += moves[direction].y
			rope.checkTail()
			visited[rope[len(rope)-1]] = true
		}
	}

	return len(visited)
}

func (d Day09) PartTwo(data []byte) interface{} {
	moveList := tools.ByteToStringSlice(data)
	moves := map[string]ropeEnd{
		"R": {x: 1, y: 0},
		"L": {x: -1, y: 0},
		"U": {x: 0, y: 1},
		"D": {x: 0, y: -1},
	}
	visited := map[ropeEnd]bool{
		{x: 0, y: 0}: true,
	}

	rope := make(rope, 0, 10)

	for i := 0; i < 10; i++ {
		rope = append(rope, ropeEnd{x: 0, y: 0})
	}

	for _, m := range moveList {
		tmp := strings.Fields(m)
		direction := tmp[0]
		num, _ := strconv.Atoi(tmp[1])
		for i := 0; i < num; i++ {
			rope[0].x += moves[direction].x
			rope[0].y += moves[direction].y
			rope.checkTail()
			visited[rope[len(rope)-1]] = true
		}
	}

	return len(visited)
}
