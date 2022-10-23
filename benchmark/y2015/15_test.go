package y2015_bench

import (
	"os"
	"testing"

    "github.com/schrenker/AoC/tools"
	"github.com/schrenker/AoC/y2015"
)

func BenchmarkDay15PartOne(b *testing.B) {
	d := y2015.Day15{}
	os.Args = []string{"cmd", "2015", "15", "01"}
	tools.GetInput("2015", "15")
	for i := 0; i < b.N; i++ {
		d.PartOne(tools.GetDefaultInputPath())
	}
}

func BenchmarkDay15PartTwo(b *testing.B) {
	d := y2015.Day15{}
	os.Args = []string{"cmd", "2015", "15", "02"}
	tools.GetInput("2015", "15")
	for i := 0; i < b.N; i++ {
		d.PartTwo(tools.GetDefaultInputPath())
	}
}