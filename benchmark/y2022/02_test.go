package y2022_bench

import (
	"os"
	"testing"

    "github.com/schrenker/AoC/tools"
	"github.com/schrenker/AoC/y2022"
)

func BenchmarkDay02PartOne(b *testing.B) {
	d := y2022.Day02{}
	os.Args = []string{"cmd", "2022", "02", "01"}
	for i := 0; i < b.N; i++ {
		d.PartOne(tools.ReadFileBytes(tools.GetDefaultInputPath()))
	}
}

func BenchmarkDay02PartTwo(b *testing.B) {
	d := y2022.Day02{}
	os.Args = []string{"cmd", "2022", "02", "02"}
	for i := 0; i < b.N; i++ {
		d.PartTwo(tools.ReadFileBytes(tools.GetDefaultInputPath()))
	}
}