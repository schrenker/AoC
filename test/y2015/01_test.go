package y2015_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/schrenker/AoC/tools"
	"github.com/schrenker/AoC/y2015"
)

func TestDay01PartOne(t *testing.T) {
	d := y2015.Day01{}
	os.Args = []string{"cmd", "2015", "01", "01"}
	expected := map[string]int{
		"1": 0,
		"2": 3,
		"3": 3,
		"4": -1,
		"5": -3,
	}
	for i, v := range expected {
		if r := d.PartOne(tools.ReadFileBytes(fmt.Sprintf("../testdata/y2015/01/1.%v.txt", i))); r != v {
			t.Fatalf("Error with test data file 1.%v.txt. Expected %v, got %v\n", i, v, r)
		}
	}
}

func TestDay01PartTwo(t *testing.T) {
	d := y2015.Day01{}
	os.Args = []string{"cmd", "2015", "01", "02"}
	expected := map[string]int{
		"1": 1,
		"2": 5,
	}
	for i, v := range expected {
		if r := d.PartTwo(tools.ReadFileBytes(fmt.Sprintf("../testdata/y2015/01/2.%v.txt", i))); r != v {
			t.Fatalf("Error with test data file 2.%v.txt. Expected %v, got %v\n", i, v, r)
		}
	}
}
