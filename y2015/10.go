package y2015

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/schrenker/go_aoc/tools"
)

type DayTen struct{}

func looknsay(seq []byte) []byte {
	buf := make([]byte, 0)
	count := 0
	cur := seq[0]
	for _, v := range seq {
		if v == cur {
			count++
		} else {
			conv := strconv.Itoa(count)
			buf = append(buf, []byte(conv)...)
			buf = append(buf, cur)
			count = 1
			cur = v
		}
	}
	conv := strconv.Itoa(count)
	buf = append(buf, []byte(conv)...)
	buf = append(buf, cur)
	return buf
}

func (d DayTen) PartOne() {
	data := tools.ReadFileBytes("input/2015/10.txt")
	data = bytes.TrimSpace(data)
	for i := 0; i < 40; i++ {
		data = looknsay(data)
	}
	fmt.Println(len(data))
}

func (d DayTen) PartTwo() {
	data := tools.ReadFileBytes("input/2015/10.txt")
	data = bytes.TrimSpace(data)
	for i := 0; i < 50; i++ {
		data = looknsay(data)
	}
	fmt.Println(len(data))
}
