package y2015_bench

import (
	"os"
	"testing"

    "github.com/schrenker/AoC/tools"
	"github.com/schrenker/AoC/y2015"
)

func BenchmarkDay10PartOne(b *testing.B) {
	d := y2015.Day10{}
	os.Args = []string{"cmd", "2015", "10", "01"}
	for i := 0; i < b.N; i++ {
		d.PartOne(tools.ReadFileBytes(tools.GetDefaultInputPath()))
	}
}

func BenchmarkDay10PartTwo(b *testing.B) {
	d := y2015.Day10{}
	os.Args = []string{"cmd", "2015", "10", "02"}
	for i := 0; i < b.N; i++ {
		d.PartTwo(tools.ReadFileBytes(tools.GetDefaultInputPath()))
	}
}