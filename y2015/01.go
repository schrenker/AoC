package y2015

type Day01 struct{}

func (d Day01) PartOne(data []byte) interface{} {
	acc := 0

	for _, v := range data {
		if v == '(' {
			acc++
		} else if v == ')' {
			acc--
		}
	}

	return acc
}

func (d Day01) PartTwo(data []byte) interface{} {
	acc := 0

	for i, v := range data {
		if v == '(' {
			acc++
		} else if v == ')' {
			acc--
		}

		if acc < 0 {
			return i + 1
		}
	}

	return acc
}
