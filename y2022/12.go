package y2022

import (
	"bytes"
	"fmt"

	"github.com/schrenker/AoC/tools"
)

type Day12 struct{}

type coord struct {
	x int
	y int
}

var visited map[coord]int

var grid [][]byte

func init() {
	visited = make(map[coord]int)
}

func findStart() coord {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 'S' {
				return coord{x: i, y: j}
			}
		}
	}
	return coord{x: 0, y: 0}
}

func move(cur coord) int {
	visited[cur]++

	poss := []int{0, -1, 1}
	results := []int{99999999999}
	fmt.Printf("Current value: %c\n", grid[cur.x][cur.y])

	for _, i := range poss {
		for _, j := range poss {
			if (i != 0 && j == 0) || (j != 0 && i == 0) {
				new := coord{x: cur.x + i, y: cur.y + j}
				if v, ok := visited[new]; ok && v > 3 {
					// fmt.Printf("Visited coord: %v. Skipping\n", new)
					continue
				}
				if new.x < 0 || new.x >= len(grid) || new.y < 0 || new.y >= len(grid[0]) {
					// fmt.Println("Coord not possible")
					continue
				}
				if grid[new.x][new.y] == 'E' {
					// fmt.Println("FOUND E")
					if grid[cur.x][cur.y] >= 'z' || grid[cur.x][cur.y]+1 == 'z' {
						// fmt.Println("Can I climb E?")
						return 1
					}
				}
				if grid[cur.x][cur.y]+1 == grid[new.x][new.y] || grid[cur.x][cur.y] >= grid[new.x][new.y] && grid[new.x][new.y] != 'E' {
					// fmt.Printf("E not found yet. Cur: %v, New: %v\n", cur, new)
					results = append(results, 1+move(new))
				}
			}
		}
	}
	return tools.GetMin(results...)
}

func (d Day12) PartOne(data []byte) interface{} {

	grid = bytes.Fields(data)
	start := findStart()
	grid[start.x][start.y] = 'a'
	return move(start) - 1
}

func (d Day12) PartTwo(data []byte) interface{} {
	return 0
}
