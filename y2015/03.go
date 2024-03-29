package y2015

type Day03 struct{}

type coords struct {
	x int
	y int
}

func (d Day03) PartOne(data []byte) interface{} {
	visits := make(map[coords]int)
	x := 0
	y := 0
	visits[coords{x: x, y: y}]++
	for _, v := range data {
		switch v {
		case '>':
			x++
		case '<':
			x--
		case '^':
			y++
		case 'v':
			y--
		}
		visits[coords{x: x, y: y}]++
	}
	return len(visits)
}

func (d Day03) PartTwo(data []byte) interface{} {
	visits := make(map[coords]int)
	x := []int{0, 0}
	y := []int{0, 0}
	visits[coords{x: x[0], y: y[0]}]++
	for i, v := range data {
		turn := i % 2
		switch v {
		case '>':
			x[turn]++
		case '<':
			x[turn]--
		case '^':
			y[turn]++
		case 'v':
			y[turn]--
		}
		visits[coords{x: x[turn], y: y[turn]}]++
	}
	return len(visits)
}
