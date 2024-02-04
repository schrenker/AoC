package y2015_bench

import (
	"os"
	"testing"

    "github.com/schrenker/AoC/input"
	"github.com/schrenker/AoC/y2015"
)

func BenchmarkDay05PartOne(b *testing.B) {
	d := y2015.Day05{}
	os.Args = []string{"cmd", "2015", "05", "01"}
	for i := 0; i < b.N; i++ {
		d.PartOne(input.ReadFileBytes(input.GetDefaultInputPath()))
	}
}

func BenchmarkDay05PartTwo(b *testing.B) {
	d := y2015.Day05{}
	os.Args = []string{"cmd", "2015", "05", "02"}
	for i := 0; i < b.N; i++ {
		d.PartTwo(input.ReadFileBytes(input.GetDefaultInputPath()))
	}
}
