package y2022

import (
	"github.com/schrenker/AoC/tools"
)

type Day08 struct{}

type tree struct {
	x int
	y int
}

func (d Day08) PartOne(data []byte) interface{} {
	grid := tools.ByteToIntGrid(data)

	visibleTrees := make(map[tree]bool)

	for i := 0; i < len(grid); i++ {
		height := -1
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > height {
				visibleTrees[tree{x: i, y: j}] = true
				height = grid[i][j]
			}
		}
	}

	for i := 0; i < len(grid[0]); i++ {
		height := -1
		for j := 0; j < len(grid); j++ {
			if grid[j][i] > height {
				visibleTrees[tree{x: j, y: i}] = true
				height = grid[j][i]
			}
		}
	}

	for i := len(grid) - 1; i >= 0; i-- {
		height := -1
		for j := len(grid[i]) - 1; j >= 0; j-- {
			if grid[i][j] > height {
				visibleTrees[tree{x: i, y: j}] = true
				height = grid[i][j]
			}
		}
	}

	for i := len(grid[0]) - 1; i >= 0; i-- {
		height := -1
		for j := len(grid) - 1; j >= 0; j-- {
			if grid[j][i] > height {
				visibleTrees[tree{x: j, y: i}] = true
				height = grid[j][i]
			}
		}
	}

	return len(visibleTrees)
}

func lookDown(grid [][]int, i, j int) int {
	trees := 0
	for k := i + 1; k < len(grid); k++ {
		trees++
		if grid[k][j] >= grid[i][j] {
			break
		}
	}
	return trees
}

func lookUp(grid [][]int, i, j int) int {
	trees := 0
	for k := i - 1; k >= 0; k-- {
		trees++
		if grid[k][j] >= grid[i][j] {
			break
		}
	}
	return trees
}

func lookRight(grid [][]int, i, j int) int {
	trees := 0
	for k := j + 1; k < len(grid[i]); k++ {
		trees++
		if grid[i][k] >= grid[i][j] {
			break
		}
	}
	return trees
}

func lookLeft(grid [][]int, i, j int) int {
	trees := 0
	for k := j - 1; k >= 0; k-- {
		trees++
		if grid[i][k] >= grid[i][j] {
			break
		}
	}
	return trees
}

func (d Day08) PartTwo(data []byte) interface{} {
	grid := tools.ByteToIntGrid(data)

	visibleTrees := make(map[tree]int)

	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			visibleTrees[tree{x: i, y: j}] = lookDown(grid, i, j) *
				lookLeft(grid, i, j) *
				lookRight(grid, i, j) *
				lookUp(grid, i, j)
		}
	}

	acc := 0
	for _, v := range visibleTrees {
		if v > acc {
			acc = v
		}
	}

	return acc
}
