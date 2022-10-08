package y2015_test

import (
	"os"
	"testing"

	"github.com/schrenker/AoC/y2015"
)

func BenchmarkDay02PartOne(b *testing.B) {
	t := y2015.Day02{}
	os.Args = []string{"cmd", "2015", "02", "01"}
	for i := 0; i < b.N; i++ {
		t.PartOne()
	}
}

func BenchmarkDay02PartTwo(b *testing.B) {
	t := y2015.Day02{}
	os.Args = []string{"cmd", "2015", "02", "02"}
	for i := 0; i < b.N; i++ {
		t.PartTwo()
	}
}