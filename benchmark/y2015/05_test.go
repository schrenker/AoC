package y2015_bench

import (
	"os"
	"testing"

    "github.com/schrenker/AoC/tools"
	"github.com/schrenker/AoC/y2015"
)

func BenchmarkDay05PartOne(b *testing.B) {
	d := y2015.Day05{}
	os.Args = []string{"cmd", "2015", "05", "01"}
	tools.GetInput("2015", "05")
	for i := 0; i < b.N; i++ {
		d.PartOne(tools.GetDefaultInputPath())
	}
}

func BenchmarkDay05PartTwo(b *testing.B) {
	d := y2015.Day05{}
	os.Args = []string{"cmd", "2015", "05", "02"}
	tools.GetInput("2015", "05")
	for i := 0; i < b.N; i++ {
		d.PartTwo(tools.GetDefaultInputPath())
	}
}