package y2015_bench

import (
	"os"
	"testing"

    "github.com/schrenker/AoC/tools"
	"github.com/schrenker/AoC/y2015"
)

func BenchmarkDay08PartOne(b *testing.B) {
	d := y2015.Day08{}
	os.Args = []string{"cmd", "2015", "08", "01"}
	tools.GetInput("2015", "08")
	for i := 0; i < b.N; i++ {
		d.PartOne(tools.GetDefaultInputPath())
	}
}

func BenchmarkDay08PartTwo(b *testing.B) {
	d := y2015.Day08{}
	os.Args = []string{"cmd", "2015", "08", "02"}
	tools.GetInput("2015", "08")
	for i := 0; i < b.N; i++ {
		d.PartTwo(tools.GetDefaultInputPath())
	}
}