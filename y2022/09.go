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

type rope struct {
	head ropeEnd
	tail ropeEnd
}

func (r *rope) checkTail() {
	x := r.head.x - r.tail.x
	y := r.head.y - r.tail.y
	if (int(math.Abs(float64(x))) == 1 && int(math.Abs(float64(y))) == 1) || int(math.Abs(float64(x)))+int(math.Abs(float64(y))) < 2 {
		return
	}
	moves := []int{0, 1, -1}
	for _, i := range moves {
		for _, j := range moves {
			if (x != 0 && i == 0) || (y != 0 && j == 0) {
				continue
			}
			xn := r.head.x - (r.tail.x + i)
			yn := r.head.y - (r.tail.y + j)
			if (int(math.Abs(float64(xn))) == 1 && int(math.Abs(float64(yn))) == 1) || int(math.Abs(float64(xn)))+int(math.Abs(float64(yn))) < 2 {
				r.tail.x += i
				r.tail.y += j
				return
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
	rope := rope{
		head: ropeEnd{x: 0, y: 0},
		tail: ropeEnd{x: 0, y: 0},
	}

	for _, m := range moveList {
		tmp := strings.Fields(m)
		direction := tmp[0]
		num, _ := strconv.Atoi(tmp[1])
		for i := 0; i < num; i++ {
			rope.head.x += moves[direction].x
			rope.head.y += moves[direction].y
			rope.checkTail()
			visited[rope.tail] = true
		}
	}

	return len(visited)
}

func (d Day09) PartTwo(data []byte) interface{} {
	return 0
}
