package y2015

type Day03 struct{}

type coord struct {
	x, y int
}

var instructionSet = map[byte][]int{
	'^': {0, 1},
	'v': {0, -1},
	'<': {-1, 0},
	'>': {1, 0},
}

func (d Day03) PartOne(data []byte) interface{} {
	x, y := 0, 0
	visited := make(map[coord]int)
	visited[coord{x, y}]++
	for _, v := range data {
		if v == '\n' {
			continue
		}
		x += instructionSet[v][0]
		y += instructionSet[v][1]
		visited[coord{x, y}]++
	}
	return len(visited)
}

func (d Day03) PartTwo(data []byte) interface{} {
	sx, sy := 0, 0
	rx, ry := 0, 0
	visited := make(map[coord]int)
	visited[coord{sx, sy}]++
	i := 0
	for _, v := range data {
		if v == '\n' {
			continue
		}
		if i%2 == 0 {
			sx += instructionSet[v][0]
			sy += instructionSet[v][1]
			visited[coord{sx, sy}]++

		} else {
			rx += instructionSet[v][0]
			ry += instructionSet[v][1]
			visited[coord{rx, ry}]++
		}
		i++
	}
	return len(visited)
}
