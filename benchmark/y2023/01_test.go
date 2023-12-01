package y2023_bench

import (
	"os"
	"testing"

    "github.com/schrenker/AoC/input"
	"github.com/schrenker/AoC/y2023"
)

func BenchmarkDay01PartOne(b *testing.B) {
	d := y2023.Day01{}
	os.Args = []string{"cmd", "2023", "01", "01"}
	for i := 0; i < b.N; i++ {
		d.PartOne(input.ReadFileBytes(input.GetDefaultInputPath()))
	}
}

func BenchmarkDay01PartTwo(b *testing.B) {
	d := y2023.Day01{}
	os.Args = []string{"cmd", "2023", "01", "02"}
	for i := 0; i < b.N; i++ {
		d.PartTwo(input.ReadFileBytes(input.GetDefaultInputPath()))
	}
}
