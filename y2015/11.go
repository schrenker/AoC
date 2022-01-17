package y2015

import (
	"fmt"

	"github.com/schrenker/go_aoc/tools"
)

type DayEleven struct{}

func incrementPass(pass []byte, idx int) []byte {
	lower := 97
	upper := 122
	pass[idx]++
	if int(pass[idx]) > upper {
		pass[idx] = byte(lower)
		pass = incrementPass(pass, idx-1)
	}
	return pass
}

func pairLetters(pass []byte) bool {
	pairs := make(map[byte]int)
	for i := 0; i < len(pass)-1; i++ {
		if pass[i] == pass[i+1] {
			pairs[pass[i]]++
			i++
		}
	}
	acc := 0
	for _, v := range pairs {
		acc += v
	}
	if acc >= 2 {
		return true
	}
	return false
}

func forbiddenLetters(pass []byte) bool {
	forbidden := []byte{'i', 'o', 'l'}
	for _, v := range pass {
		for _, b := range forbidden {
			if v == b {
				return false
			}
		}
	}
	return true
}

func increasingLetters(pass []byte) bool {
	for i := 0; i < len(pass)-2; i++ {
		if tools.MultipleEqualInt(int(pass[i+2])-2, int(pass[i+1])-1, int(pass[i])) {
			return true
		}
	}
	return false
}

func validatePass(pass []byte) bool {
	return increasingLetters(pass) && forbiddenLetters(pass) && pairLetters(pass)
}

func generateNextPass(pass []byte) []byte {
	for {
		if validatePass(pass) {
			break
		} else {
			pass = incrementPass(pass, len(pass)-1)
		}
	}
	return pass
}

func (d DayEleven) PartOne() {
	data := tools.ReadFileBytes("input/2015/11.txt")
	fmt.Println(string(generateNextPass(data)))
}

func (d DayEleven) PartTwo() {
	data := tools.ReadFileBytes("input/2015/11.txt")
	data = incrementPass(generateNextPass(data), len(data)-1)
	fmt.Println(string(generateNextPass(data)))
}
