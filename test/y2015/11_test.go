package y2015_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/schrenker/AoC/tools"
	"github.com/schrenker/AoC/y2015"
)

func TestDay11PartOne(t *testing.T) {
	d := y2015.Day11{}
	os.Args = []string{"cmd", "2015", "11", "01"}
	tools.GetInput("2015", "11")
	expected := map[string]int{
		"1": 0,
	}
	for i, v := range expected {
		if r := d.PartOne(fmt.Sprintf("../testdata/y2015/11/1.%v.txt", i)); r != v {
			t.Fatalf("Error with test data file 1.%v.txt. Expected %v, got %v\n", i, v, r)
		}
	}
}

func TestDay11PartTwo(t *testing.T) {
	d := y2015.Day11{}
	os.Args = []string{"cmd", "2015", "11", "02"}
	tools.GetInput("2015", "11")
	expected := map[string]int{
		"1": 0,
	}
	for i, v := range expected {
		if r := d.PartTwo(fmt.Sprintf("../testdata/y2015/11/2.%v.txt", i)); r != v {
			t.Fatalf("Error with test data file 2.%v.txt. Expected %v, got %v\n", i, v, r)
		}
	}
}
