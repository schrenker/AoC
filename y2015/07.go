package y2015

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/schrenker/go_aoc/tools"
)

type DaySeven struct{}

type wireMap map[string]uint16

func (w *wireMap) decode(op string) error {
	operands := strings.Split(op, " -> ")
	source := strings.Split(operands[0], " ")

	if _, ok := (*w)[operands[1]]; ok {
		return nil
	}

	switch len(source) {

	case 1:
		conv, err := strconv.Atoi(source[0])
		if err != nil {
			if _, ok := (*w)[source[0]]; ok {
				(*w)[operands[1]] = (*w)[source[0]]
				return nil
			}
		} else {
			(*w)[operands[1]] = uint16(conv)
			return nil
		}

	case 2:
		conv, err := strconv.Atoi(source[1])
		if err != nil {
			if _, ok := (*w)[source[1]]; ok {
				(*w)[operands[1]] = not((*w)[source[1]])
				return nil
			}
		} else {
			(*w)[operands[1]] = not(uint16(conv))
			return nil
		}

	case 3:
		left, errl := strconv.Atoi(source[0])
		right, errr := strconv.Atoi(source[2])
		if errl == nil && errr == nil {
			(*w)[operands[1]] = opMap[source[1]](uint16(left), uint16(right))
			return nil
		} else if errl == nil && errr != nil {
			if _, ok := (*w)[source[2]]; ok {
				(*w)[operands[1]] = opMap[source[1]](uint16(left), uint16((*w)[source[2]]))
				return nil
			}
		} else if errl != nil && errr == nil {
			if _, ok := (*w)[source[0]]; ok {
				(*w)[operands[1]] = opMap[source[1]](uint16((*w)[source[0]]), uint16(right))
				return nil
			}
		} else {
			_, okl := (*w)[source[0]]
			_, okr := (*w)[source[2]]
			if okl && okr {
				(*w)[operands[1]] = opMap[source[1]](uint16((*w)[source[0]]), uint16((*w)[source[2]]))
				return nil
			}
		}

	}
	return fmt.Errorf("")
}

func not(operand uint16) uint16 {
	return ^operand
}

func and(left, right uint16) uint16 {
	return left & right
}

func or(left, right uint16) uint16 {
	return left | right
}

func rshift(left, right uint16) uint16 {
	return left >> right
}

func lshift(left, right uint16) uint16 {
	return left << right
}

var opMap = map[string]func(uint16, uint16) uint16{
	"AND":    and,
	"OR":     or,
	"RSHIFT": rshift,
	"LSHIFT": lshift,
}

func (d DaySeven) PartOne() {
	operations := tools.ReadFileStringSlice("input/2015/07.txt")
	wires := make(wireMap)
	for {
		for i := 0; i < len(operations); i++ {
			err := wires.decode(operations[i])
			if err == nil {
				operations = tools.RemoveElement(operations, i)
				break
			}
		}
		if len(operations) == 1 && operations[0] == "" {
			break
		}
	}
	fmt.Println(wires["a"])
}

func (d DaySeven) PartTwo() {
	operations := tools.ReadFileStringSlice("input/2015/07.txt")
	wires := make(wireMap)
	for {
		for i := 0; i < len(operations); i++ {
			err := wires.decode(operations[i])
			if err == nil {
				operations = tools.RemoveElement(operations, i)
				break
			}
		}
		if len(operations) == 1 && operations[0] == "" {
			break
		}
	}
	signalA := wires["a"]
	operations = tools.ReadFileStringSlice("input/2015/07.txt")
	wires = make(wireMap)
	wires["b"] = signalA
	for {
		for i := 0; i < len(operations); i++ {
			err := wires.decode(operations[i])
			if err == nil {
				operations = tools.RemoveElement(operations, i)
				break
			}
		}
		if len(operations) == 1 && operations[0] == "" {
			break
		}
	}
	fmt.Println(wires["a"])
}
