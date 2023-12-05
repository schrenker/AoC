package y2023_bench

import (
	"os"
	"testing"

    "github.com/schrenker/AoC/input"
	"github.com/schrenker/AoC/y2023"
)

func BenchmarkDay05PartOne(b *testing.B) {
	d := y2023.Day05{}
	os.Args = []string{"cmd", "2023", "05", "01"}
	for i := 0; i < b.N; i++ {
		d.PartOne(input.ReadFileBytes(input.GetDefaultInputPath()))
	}
}

func BenchmarkDay05PartTwo(b *testing.B) {
	d := y2023.Day05{}
	os.Args = []string{"cmd", "2023", "05", "02"}
	for i := 0; i < b.N; i++ {
		d.PartTwo(input.ReadFileBytes(input.GetDefaultInputPath()))
	}
}
