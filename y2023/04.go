package y2023

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/schrenker/AoC/input"
)

type Day04 struct{}

func (d Day04) PartOne(data []byte) interface{} {
	acc := 0
	for _, line := range input.ByteToStringSlice(data) {
		numbers := make(map[string]int)
		spacePattern := regexp.MustCompile(`[|\s]+`)
		tmp := spacePattern.Split(strings.Split(strings.TrimSpace(line), ": ")[1], -1)
		points := 0
		for _, v := range tmp {
			numbers[v]++
		}
		for _, v := range numbers {
			if v >= 2 {
				if points == 0 {
					points = 1
				} else {
					points = points << 1
				}
			}
		}
		acc += points
	}

	return acc
}

func (d Day04) PartTwo(data []byte) interface{} {
	acc := 0
	scratchcards := make(map[string]int)
	copies := make(map[string]int)
	for _, line := range input.ByteToStringSlice(data) {
		numbers := make(map[string]int)
		spacePattern := regexp.MustCompile(`[|\s]+`)

		tmp := strings.Split(strings.TrimSpace(line), ": ")
		tmp[0] = spacePattern.ReplaceAllString(tmp[0], " ")

		copies[tmp[0]] = 1
		for _, v := range spacePattern.Split(tmp[1], -1) {
			numbers[v]++
		}
		for _, v := range numbers {
			if v >= 2 {
				scratchcards[tmp[0]]++
			}
		}
	}

	for i := 1; i <= len(copies); i++ {
		currentCard := fmt.Sprintf("Card %v", i)
		matches := scratchcards[currentCard]
		for m := 1; m <= matches; m++ {
			if m+i > len(copies) {
				break
			}
			copies[fmt.Sprintf("Card %v", i+m)] += copies[currentCard]
		}
	}

	for _, v := range copies {
		acc += v
	}

	return acc
}
