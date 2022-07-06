package y2015

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/schrenker/AoC/tools"
)

type Day14 struct{}

type reindeer struct {
	distance    int
	travelTime  int
	restingTime int
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
		distance:    nums[0],
		travelTime:  nums[1],
		restingTime: nums[2],
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
	return 0
}
