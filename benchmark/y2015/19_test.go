package y2015_bench

import (
	"os"
	"testing"

    "github.com/schrenker/AoC/tools"
	"github.com/schrenker/AoC/y2015"
)

func BenchmarkDay19PartOne(b *testing.B) {
	d := y2015.Day19{}
	os.Args = []string{"cmd", "2015", "19", "01"}
	tools.GetInput("2015", "19")
	for i := 0; i < b.N; i++ {
		d.PartOne(tools.GetDefaultInputPath())
	}
}

func BenchmarkDay19PartTwo(b *testing.B) {
	d := y2015.Day19{}
	os.Args = []string{"cmd", "2015", "19", "02"}
	tools.GetInput("2015", "19")
	for i := 0; i < b.N; i++ {
		d.PartTwo(tools.GetDefaultInputPath())
	}
}