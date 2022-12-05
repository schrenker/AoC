package y2022

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/schrenker/AoC/tools"
)

type Day05 struct{}

func getCrates(data []byte) ([][]string, int) {
	stringData := tools.ByteToStringSlice(data)
	crates := [][]string{}
	index := 0
	for _, v := range stringData {
		index++
		if v == "" {
			break
		}
		crates = append(crates, strings.Split(v, ""))
	}

	return crates, index
}

func stackify(crates [][]string) []*tools.Stack[string] {
	stacks := make([]*tools.Stack[string], 0)
	for i := 0; i < len(crates)-1; i++ {
		stacknum := 0
		for j := 1; j < len(crates[i]); j += 4 {

			if stacknum >= len(stacks) {
				stacks = append(stacks, tools.NewStack[string]())
			}

			if crates[i][j] != " " {
				stacks[stacknum].Push(crates[i][j])
			}

			stacknum++
		}
	}
	for i := range stacks {
		stacks[i].Reverse()
	}
	return stacks
}

func parseInstructions(index int, data []byte) [][]int {
	instructions := make([][]int, 0)
	stringData := tools.ByteToStringSlice(data)
	rgx := regexp.MustCompile(`\D`)
	spaceRgx := regexp.MustCompile(`[ ]{1,}`)

	for i := index; i < len(stringData); i++ {
		tmp := rgx.ReplaceAllString(stringData[i], " ")
		tmp = spaceRgx.ReplaceAllString(tmp, " ")
		splitTmp := strings.Split(tmp, " ")

		instruction := make([]int, 0, 3)

		for _, v := range splitTmp {
			num, err := strconv.Atoi(v)
			if err != nil {
				continue
			}
			instruction = append(instruction, num)
		}

		instructions = append(instructions, instruction)
	}

	return instructions
}

func (d Day05) PartOne(data []byte) interface{} {
	crates, index := getCrates(data)
	stacks := stackify(crates)
	instructions := parseInstructions(index, data)

	for i := range instructions {
		for j := 0; j < instructions[i][0]; j++ {
			tmp := stacks[instructions[i][1]-1].Pop()
			stacks[instructions[i][2]-1].Push(tmp)
		}
	}

	result := ""
	for i := range stacks {
		result += stacks[i].Peek()
	}

	return result
}

func (d Day05) PartTwo(data []byte) interface{} {
	crates, index := getCrates(data)
	stacks := stackify(crates)
	instructions := parseInstructions(index, data)

	for i := range instructions {
		orderPreserver := make([]string, 0, instructions[i][0])
		for j := 0; j < instructions[i][0]; j++ {
			orderPreserver = append(orderPreserver, stacks[instructions[i][1]-1].Pop())
		}
		orderPreserver = tools.Reverse(orderPreserver)
		for _, v := range orderPreserver {
			stacks[instructions[i][2]-1].Push(v)
		}
	}

	result := ""
	for i := range stacks {
		result += stacks[i].Peek()
	}

	return result

}
