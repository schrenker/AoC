package y2015_bench

import (
	"os"
	"testing"

    "github.com/schrenker/AoC/tools"
	"github.com/schrenker/AoC/y2015"
)

func BenchmarkDay13PartOne(b *testing.B) {
	d := y2015.Day13{}
	os.Args = []string{"cmd", "2015", "13", "01"}
	tools.GetInput("2015", "13")
	for i := 0; i < b.N; i++ {
		d.PartOne(tools.GetDefaultInputPath())
	}
}

func BenchmarkDay13PartTwo(b *testing.B) {
	d := y2015.Day13{}
	os.Args = []string{"cmd", "2015", "13", "02"}
	tools.GetInput("2015", "13")
	for i := 0; i < b.N; i++ {
		d.PartTwo(tools.GetDefaultInputPath())
	}
}