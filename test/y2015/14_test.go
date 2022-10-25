package y2015_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/schrenker/AoC/tools"
	"github.com/schrenker/AoC/y2015"
)

func TestDay14PartOne(t *testing.T) {
	d := y2015.Day14{}
	os.Args = []string{"cmd", "2015", "14", "01"}
	tools.GetInput("2015", "14")
	expected := map[string]int{
		"1": 0,
	}
	for i, v := range expected {
		if r := d.PartOne(fmt.Sprintf("../testdata/y2015/14/1.%v.txt", i)); r != v {
			t.Fatalf("Error with test data file 1.%v.txt. Expected %v, got %v\n", i, v, r)
		}
	}
}

func TestDay14PartTwo(t *testing.T) {
	d := y2015.Day14{}
	os.Args = []string{"cmd", "2015", "14", "02"}
	tools.GetInput("2015", "14")
	expected := map[string]int{
		"1": 0,
	}
	for i, v := range expected {
		if r := d.PartTwo(fmt.Sprintf("../testdata/y2015/14/2.%v.txt", i)); r != v {
			t.Fatalf("Error with test data file 2.%v.txt. Expected %v, got %v\n", i, v, r)
		}
	}
}
