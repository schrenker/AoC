package y2015_bench

import (
	"os"
	"testing"

    "github.com/schrenker/AoC/tools"
	"github.com/schrenker/AoC/y2015"
)

func BenchmarkDay12PartOne(b *testing.B) {
	d := y2015.Day12{}
	os.Args = []string{"cmd", "2015", "12", "01"}
	tools.GetInput("2015", "12")
	for i := 0; i < b.N; i++ {
		d.PartOne(tools.GetDefaultInputPath())
	}
}

func BenchmarkDay12PartTwo(b *testing.B) {
	d := y2015.Day12{}
	os.Args = []string{"cmd", "2015", "12", "02"}
	tools.GetInput("2015", "12")
	for i := 0; i < b.N; i++ {
		d.PartTwo(tools.GetDefaultInputPath())
	}
}