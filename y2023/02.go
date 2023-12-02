package y2023

import (
	"strconv"
	"strings"

	"github.com/schrenker/AoC/input"
)

type Day02 struct{}

type round struct {
	red   int
	blue  int
	green int
}

type game struct {
	id     int
	rounds []*round
}

func getID(s string) int {
	tmp := strings.Split(s, ":")
	id, err := strconv.Atoi(tmp[0][5:len(tmp[0])])
	if err != nil {
		panic(err)
	}
	return id
}

func getRounds(s string) []*round {
	tmp := strings.Split(strings.Split(s, ":")[1], ";")
	result := make([]*round, 0)

	for _, r := range tmp {
		balls := strings.Split(r, ", ")
		for _, ball := range balls {
			ballType := strings.Split(strings.TrimSpace(ball), " ")
			amount, err := strconv.Atoi(ballType[0])
			if err != nil {
				panic(err)
			}
			rd := &round{red: 0, green: 0, blue: 0}

			switch ballType[1] {
			case "red":
				rd.red = amount
			case "green":
				rd.green = amount
			case "blue":
				rd.blue = amount
			}
			if rd != nil {
				result = append(result, rd)
			}
		}
	}

	return result
}

func parseGame(s string) *game {
	return &game{
		id:     getID(s),
		rounds: getRounds(s),
	}
}

func (d Day02) PartOne(data []byte) interface{} {
	acc := 0
	for _, game := range input.ByteToStringSlice(data) {
		possible := true
		g := parseGame(game)
		for _, r := range g.rounds {
			if r.red > 12 || r.green > 13 || r.blue > 14 {
				possible = false
				break
			}
		}
		if possible {
			acc += g.id
		}
	}
	return acc
}

func (d Day02) PartTwo(data []byte) interface{} {
	acc := 0
	for _, game := range input.ByteToStringSlice(data) {
		g := parseGame(game)
		red := 0
		green := 0
		blue := 0

		for _,r := range g.rounds {
			if r.red > red {
				red = r.red
			}
			if r.green > green {
				green = r.green
			}
			if r.blue > blue {
				blue = r.blue
			}
		}
		tmp := 1
		if red > 0 {
			tmp *= red
		}
		if green > 0 {
			tmp *= green
		}
		if blue > 0 {
			tmp *= blue
		}
		acc += tmp
	}
	return acc
}
