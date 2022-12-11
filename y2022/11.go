package y2022

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/schrenker/AoC/tools"
)

type Day11 struct{}

type monkey struct {
	// amount    int
	fail      *monkey
	inspected int
	items     *tools.Queue[int]
	operation string
	success   *monkey
	test      int
}

func (m *monkey) takeTurn(divideByThree bool) {
	r := regexp.MustCompile(`old`)
	length := len(m.items.List())
	for i := 0; i < length; i++ {
		conv := strconv.Itoa(m.items.Dequeue())
		if conv == "0" {
			return
		}
		m.inspected++
		tmp := r.ReplaceAllString(m.operation, conv)
		result := tools.StringMathToInt(tmp)

		if divideByThree {
			result = int(result / 3)
		}

		if result%m.test == 0 {
			m.success.items.Enqueue(result)
		} else {
			m.fail.items.Enqueue(result)
		}
	}
}

func (d Day11) PartOne(data []byte) interface{} {
	ops := tools.ByteToStringSlice(data)
	monkeys := make([]*monkey, 0)

	for _, v := range ops {
		if v == "" {
			continue
		}
		if strings.Fields(v)[0] == "Monkey" {
			monkeys = append(monkeys, &monkey{
				items: tools.NewQueue[int](),
			})
		}
	}

	numRgx := regexp.MustCompile(`[\d]+`)
	opRgx := regexp.MustCompile(`([a-z]{0,3}[0-9]* [+\-*\/] [a-z]{0,3}[0-9]*)`)

	for i, m := 0, 0; i < len(ops); i, m = i+7, m+1 {
		for _, v := range numRgx.FindAllString(ops[i+1], -1) {
			tmp, err := strconv.Atoi(v)
			if err != nil {
				log.Fatalln(err)
			}
			monkeys[m].items.Enqueue(tmp)
		}
		monkeys[m].operation = opRgx.FindString(ops[i+2])

		tmp := numRgx.FindAllString(ops[i+3], -1)
		num, err := strconv.Atoi(tmp[len(tmp)-1])
		if err != nil {
			log.Fatalln(err)
		}

		monkeys[m].test = num

		tmp = numRgx.FindAllString(ops[i+4], -1)
		num, err = strconv.Atoi(tmp[len(tmp)-1])
		if err != nil {
			log.Fatalln(err)
		}

		monkeys[m].success = monkeys[num]

		tmp = numRgx.FindAllString(ops[i+5], -1)
		num, err = strconv.Atoi(tmp[len(tmp)-1])
		if err != nil {
			log.Fatalln(err)
		}

		monkeys[m].fail = monkeys[num]
	}

	for i := 0; i < 20; i++ {
		for _, m := range monkeys {
			m.takeTurn(true)
		}
	}

	inspections := make([]int, len(monkeys))
	for _, m := range monkeys {
		inspections = append(inspections, m.inspected)
	}

	mostInspected := tools.GetMaxMultiple(2, inspections...)

	return tools.MulAll(mostInspected...)
}

func (d Day11) PartTwo(data []byte) interface{} {
	ops := tools.ByteToStringSlice(data)
	monkeys := make([]*monkey, 0)

	for _, v := range ops {
		if v == "" {
			continue
		}
		if strings.Fields(v)[0] == "Monkey" {
			monkeys = append(monkeys, &monkey{
				items: tools.NewQueue[int](),
			})
		}
	}

	numRgx := regexp.MustCompile(`[\d]+`)
	opRgx := regexp.MustCompile(`([a-z]{0,3}[0-9]* [+\-*\/] [a-z]{0,3}[0-9]*)`)

	for i, m := 0, 0; i < len(ops); i, m = i+7, m+1 {
		for _, v := range numRgx.FindAllString(ops[i+1], -1) {
			tmp, err := strconv.Atoi(v)
			if err != nil {
				log.Fatalln(err)
			}
			monkeys[m].items.Enqueue(tmp)
		}
		monkeys[m].operation = opRgx.FindString(ops[i+2])

		tmp := numRgx.FindAllString(ops[i+3], -1)
		num, err := strconv.Atoi(tmp[len(tmp)-1])
		if err != nil {
			log.Fatalln(err)
		}

		monkeys[m].test = num

		tmp = numRgx.FindAllString(ops[i+4], -1)
		num, err = strconv.Atoi(tmp[len(tmp)-1])
		if err != nil {
			log.Fatalln(err)
		}

		monkeys[m].success = monkeys[num]

		tmp = numRgx.FindAllString(ops[i+5], -1)
		num, err = strconv.Atoi(tmp[len(tmp)-1])
		if err != nil {
			log.Fatalln(err)
		}

		monkeys[m].fail = monkeys[num]
	}

	for i := 0; i < 10000; i++ {
		for _, m := range monkeys {
			m.takeTurn(false)
		}
	}

	inspections := make([]int, 0, len(monkeys))
	for _, m := range monkeys {
		inspections = append(inspections, m.inspected)
	}

	fmt.Println(inspections)
	mostInspected := tools.GetMaxMultiple(2, inspections...)

	return tools.MulAll(mostInspected...)
}
