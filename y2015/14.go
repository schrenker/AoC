package y2015

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/schrenker/AoC/tools"
)

type Day14 struct{}

type reindeer struct {
	distance        int
	travelTime      int
	restingTime     int
	currentScore    int
	currentDistance int
	currentTime     int
	isFlying        bool
}

func spawnReindeer(s string) reindeer {
	nums := tools.Map(
		tools.Filter(
			strings.Split(s, " "),
			func(s string) bool {
				if _, err := strconv.Atoi(s); err != nil {
					return false
				}
				return true
			}),
		func(s string) int {
			num, _ := strconv.Atoi(s) // no need to check for errors here, since I did a function ago
			return num
		})

	return reindeer{
		distance:        nums[0],
		travelTime:      nums[1],
		restingTime:     nums[2],
		currentScore:    0,
		currentDistance: 0,
		currentTime:     0,
		isFlying:        true,
	}
}

func (r reindeer) getDistance(seconds int) int {
	distance := 0
	for i, elapsed := 0, 0; elapsed <= seconds; i++ {
		if i%2 == 0 {
			if r.travelTime+elapsed > seconds {
				fmt.Printf("%v, %v, %v\n", r.travelTime, elapsed, (r.travelTime + elapsed))
				distance = distance + r.distance*(seconds-elapsed)
				return distance
			}
			distance = distance + r.distance*r.travelTime
			elapsed += r.travelTime
		} else {
			elapsed += r.restingTime
		}
	}
	return distance
}

func (r *reindeer) second() {
	r.currentTime++

	if r.isFlying {
		r.currentDistance += r.distance
		if r.currentTime == r.travelTime {
			r.currentTime = 0
			r.isFlying = !r.isFlying
		}
	}

	if !r.isFlying && r.currentTime != 0 {
		if r.currentTime == r.restingTime {
			r.currentTime = 0
			r.isFlying = !r.isFlying
		}
	}
}

func awardPoint(r []*reindeer) {
	max := tools.Reduce(r, 0, func(s int, rein *reindeer) int {
		if s < rein.currentDistance {
			s = rein.currentDistance
		}
		return s
	})

	for i := range r {
		fmt.Printf("%v %v %v %v\n", i, r[i].currentScore, r[i].currentDistance, max)
	}

	for i := range r {
		if r[i].currentDistance == max {
			r[i].currentScore++
		}
	}
}

func getWinner(r []*reindeer) int {
	win := tools.Reduce(r, 0, func(s int, rein *reindeer) int {
		if s < rein.currentScore {
			s = rein.currentScore
		}
		return s
	})
	return win
}

func (d Day14) PartOne() interface{} {
	data := tools.ReadFileStringSlice()
	seconds := 2503
	distances := []int{}
	for _, v := range data {
		distances = append(distances, spawnReindeer(v).getDistance(seconds))
	}

	return tools.GetMax(distances...)
}

func (d Day14) PartTwo() interface{} {
	data := tools.ReadFileStringSlice()
	seconds := 2503
	reindeers := []*reindeer{}
	for _, v := range data {
		rein := spawnReindeer(v)
		reindeers = append(reindeers, &rein)
	}

	for i := 0; i < seconds; i++ {
		for j := range reindeers {
			reindeers[j].second()
		}

		awardPoint(reindeers)
	}

	return tools.Reduce(reindeers, 0, func(s int, rein *reindeer) int {
		if s < rein.currentScore {
			s = rein.currentScore
		}
		return s
	})
}
