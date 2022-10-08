package y2015_test

import (
	"os"
	"testing"

	"github.com/schrenker/AoC/y2015"
)

func BenchmarkDayOnePartOne(b *testing.B) {
	t := y2015.Day01{}
	os.Args = []string{"cmd", "2015", "01", "01"}
	for i := 0; i < b.N; i++ {
		t.PartOne()
	}
}

func BenchmarkDayOnePartTwo(b *testing.B) {
	t := y2015.Day01{}
	os.Args = []string{"cmd", "2015", "01", "02"}
	for i := 0; i < b.N; i++ {
		t.PartTwo()
	}
}
