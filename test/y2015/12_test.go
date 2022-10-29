package y2015_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/schrenker/AoC/tools"
	"github.com/schrenker/AoC/y2015"
)

func TestDay12PartOne(t *testing.T) {
	d := y2015.Day12{}
	os.Args = []string{"cmd", "2015", "12", "01"}
	expected := map[string]int{
		"1": 0,
	}
	for i, v := range expected {
		if r := d.PartOne(tools.ReadFileBytes(fmt.Sprintf("../testdata/y2015/12/1.%v.txt", i))); r != v {
			t.Fatalf("Error with test data file 1.%v.txt. Expected %v, got %v\n", i, v, r)
		}
	}
}

func TestDay12PartTwo(t *testing.T) {
	d := y2015.Day12{}
	os.Args = []string{"cmd", "2015", "12", "02"}
	expected := map[string]int{
		"1": 0,
	}
	for i, v := range expected {
		if r := d.PartTwo(tools.ReadFileBytes(fmt.Sprintf("../testdata/y2015/12/2.%v.txt", i))); r != v {
			t.Fatalf("Error with test data file 2.%v.txt. Expected %v, got %v\n", i, v, r)
		}
	}
}
