package y2022

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/schrenker/AoC/tools"
)

type Day10 struct{}

func (d Day10) PartOne(data []byte) interface{} {
	calls := tools.ByteToStringSlice(data)
	q := tools.NewQueue[int]()

	for i := 0; i < len(calls); i++ {
		tmp := strings.Fields(calls[i])
		q.Enqueue(0)
		if len(tmp) == 2 {
			num, _ := strconv.Atoi(tmp[1])
			q.Enqueue(num)
		}
	}

	regX := 1
	acc := 0
	check := 20

	for i := 1; ; i++ {
		if i == check {
			acc += regX * check
			check += 40
			if check > 220 {
				return acc
			}
		}
		regX += q.Dequeue()
	}
}

func (d Day10) PartTwo(data []byte) interface{} {
	calls := tools.ByteToStringSlice(data)
	q := tools.NewQueue[int]()

	for i := 0; i < len(calls); i++ {
		tmp := strings.Fields(calls[i])
		q.Enqueue(0)
		if len(tmp) == 2 {
			num, _ := strconv.Atoi(tmp[1])
			q.Enqueue(num)
		}
	}
	regX := 1
	screen := make([]byte, 240)
	shift := 0
	for i := range screen {
		if i%40 == 0 {
			shift = i
		}
		if i == regX+shift || i == regX-1+shift || i == regX+1+shift {
			screen[i] = '#'
		} else {
			screen[i] = '.'
		}

		regX += q.Dequeue()
	}

	result := make([]string, 0, 6)

	for i := 0; i < 240; i += 40 {
		result = append(result, string(screen[i:i+40]))
	}

	for _, v := range result {
		fmt.Println(v)
	}

	return result
}
