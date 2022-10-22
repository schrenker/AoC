package y2015

import (
	"github.com/schrenker/AoC/tools"
)

type Day18 struct{}

var directions = [][]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func turnOnCornerLights(state [][]bool) {
	state[0][0] = true
	state[0][len(state[0])-1] = true
	state[len(state)-1][0] = true
	state[len(state)-1][len(state[0])-1] = true
}

func createState(path string) [][]bool {
	data := tools.ReadFileStringSlice(path)
	result := make([][]bool, len(data))
	for i := range data {
		result[i] = make([]bool, len(data[i]))
		for j := range data[i] {
			switch data[i][j] {
			case '.':
				result[i][j] = false
			case '#':
				result[i][j] = true
			}
		}
	}
	return result
}

func countState(state [][]bool) int {
	acc := 0
	for _, i := range state {
		for _, v := range i {
			if v {
				acc++
			}
		}
	}
	return acc
}

func decide(neigh int, state bool) bool {
	if !state {
		if neigh == 3 {
			return true
		}
	} else {
		if neigh == 2 || neigh == 3 {
			return true
		}
	}
	return false
}

func evolve(state [][]bool) [][]bool {
	reference := make([][]bool, len(state))
	for i := range state {
		reference[i] = make([]bool, len(state[i]))
		copy(reference[i], state[i])
	}

	for i := range state {
		for j := range state[i] {
			neigh := 0
			for k := range directions {
				tmp1 := i + directions[k][0]
				tmp2 := j + directions[k][1]
				if tmp1 >= 0 && tmp1 < len(state) &&
					tmp2 >= 0 && tmp2 < len(state[i]) {
					if state[tmp1][tmp2] {
						neigh++
					}
				}
			}
			reference[i][j] = decide(neigh, state[i][j])
		}
	}

	return reference
}

func (d Day18) PartOne(path string) interface{} {
	state := createState(path)
	for i := 0; i < 100; i++ {
		state = evolve(state)
	}
	return countState(state)
}

func (d Day18) PartTwo(path string) interface{} {
	state := createState(path)
	turnOnCornerLights(state)
	for i := 0; i < 100; i++ {
		state = evolve(state)
		turnOnCornerLights(state)
	}
	return countState(state)
}
