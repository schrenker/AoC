package y2015_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/schrenker/AoC/input"
	"github.com/schrenker/AoC/y2015"
)

func TestDay05PartOne(t *testing.T) {
	d := y2015.Day05{}
	os.Args = []string{"cmd", "2015", "05", "01"}
	expected := map[string]int{
		"1": 1,
		"2": 1,
		"3": 0,
		"4": 0,
		"5": 0,
	}
	for i, v := range expected {
		if r := d.PartOne(input.ReadFileBytes(fmt.Sprintf("../testdata/y2015/05/1.%v.txt", i))); r != v {
			t.Fatalf("Error with test data file 1.%v.txt. Expected %v, got %v\n", i, v, r)
		}
	}
}

func TestDay05PartTwo(t *testing.T) {
	d := y2015.Day05{}
	os.Args = []string{"cmd", "2015", "05", "02"}
	expected := map[string]int{
		"1": 1,
		"2": 1,
		"3": 0,
		"4": 0,
	}
	for i, v := range expected {
		if r := d.PartTwo(input.ReadFileBytes(fmt.Sprintf("../testdata/y2015/05/2.%v.txt", i))); r != v {
			t.Fatalf("Error with test data file 2.%v.txt. Expected %v, got %v\n", i, v, r)
		}
	}
}