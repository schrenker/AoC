package y2015

import (
	"strconv"

	"github.com/schrenker/AoC/tools"
)

type Day10 struct{}

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

func (d Day10) PartOne() interface{} {
	data := tools.ReadFileBytes("input/2015/10.txt")
	for i := 0; i < 40; i++ {
		data = looknsay(data)
	}
	return len(data)
}

func (d Day10) PartTwo() interface{} {
	data := tools.ReadFileBytes("input/2015/10.txt")
	for i := 0; i < 50; i++ {
		data = looknsay(data)
	}
	return len(data)
}
