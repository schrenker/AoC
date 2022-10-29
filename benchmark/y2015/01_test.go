package y2015_bench

import (
	"os"
	"testing"

    "github.com/schrenker/AoC/tools"
	"github.com/schrenker/AoC/y2015"
)

func BenchmarkDay01PartOne(b *testing.B) {
	d := y2015.Day01{}
	os.Args = []string{"cmd", "2015", "01", "01"}
	for i := 0; i < b.N; i++ {
		d.PartOne(tools.ReadFileBytes(tools.GetDefaultInputPath()))
	}
}

func BenchmarkDay01PartTwo(b *testing.B) {
	d := y2015.Day01{}
	os.Args = []string{"cmd", "2015", "01", "02"}
	for i := 0; i < b.N; i++ {
		d.PartTwo(tools.ReadFileBytes(tools.GetDefaultInputPath()))
	}
}