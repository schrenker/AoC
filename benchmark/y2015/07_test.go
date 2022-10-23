package y2015_bench

import (
	"os"
	"testing"

    "github.com/schrenker/AoC/tools"
	"github.com/schrenker/AoC/y2015"
)

func BenchmarkDay07PartOne(b *testing.B) {
	d := y2015.Day07{}
	os.Args = []string{"cmd", "2015", "07", "01"}
	tools.GetInput("2015", "07")
	for i := 0; i < b.N; i++ {
		d.PartOne(tools.GetDefaultInputPath())
	}
}

func BenchmarkDay07PartTwo(b *testing.B) {
	d := y2015.Day07{}
	os.Args = []string{"cmd", "2015", "07", "02"}
	tools.GetInput("2015", "07")
	for i := 0; i < b.N; i++ {
		d.PartTwo(tools.GetDefaultInputPath())
	}
}