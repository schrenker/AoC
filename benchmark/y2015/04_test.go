package y2015_bench

import (
	"os"
	"testing"

    "github.com/schrenker/AoC/input"
	"github.com/schrenker/AoC/y2015"
)

func BenchmarkDay04PartOne(b *testing.B) {
	d := y2015.Day04{}
	os.Args = []string{"cmd", "2015", "04", "01"}
	for i := 0; i < b.N; i++ {
		d.PartOne(input.ReadFileBytes(input.GetDefaultInputPath()))
	}
}

func BenchmarkDay04PartTwo(b *testing.B) {
	d := y2015.Day04{}
	os.Args = []string{"cmd", "2015", "04", "02"}
	for i := 0; i < b.N; i++ {
		d.PartTwo(input.ReadFileBytes(input.GetDefaultInputPath()))
	}
}
