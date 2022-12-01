package y2015_bench

import (
	"os"
	"testing"

    "github.com/schrenker/AoC/tools"
	"github.com/schrenker/AoC/y2015"
)

func BenchmarkDay18PartOne(b *testing.B) {
	d := y2015.Day18{}
	os.Args = []string{"cmd", "2015", "18", "01"}
	for i := 0; i < b.N; i++ {
		d.PartOne(tools.ReadFileBytes(tools.GetDefaultInputPath()))
	}
}

func BenchmarkDay18PartTwo(b *testing.B) {
	d := y2015.Day18{}
	os.Args = []string{"cmd", "2015", "18", "02"}
	for i := 0; i < b.N; i++ {
		d.PartTwo(tools.ReadFileBytes(tools.GetDefaultInputPath()))
	}
}