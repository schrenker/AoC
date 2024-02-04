package y2015_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/schrenker/AoC/input"
	"github.com/schrenker/AoC/y2015"
)

func TestDay03PartOne(t *testing.T) {
	d := y2015.Day03{}
	os.Args = []string{"cmd", "2015", "03", "01"}
	expected := map[string]int{
		"1": 2,
		"2": 4,
		"3": 2,
	}
	for i, v := range expected {
		if r := d.PartOne(input.ReadFileBytes(fmt.Sprintf("../testdata/y2015/03/1.%v.txt", i))); r != v {
			t.Fatalf("Error with test data file 1.%v.txt. Expected %v, got %v\n", i, v, r)
		}
	}
}

func TestDay03PartTwo(t *testing.T) {
	d := y2015.Day03{}
	os.Args = []string{"cmd", "2015", "03", "02"}
	expected := map[string]int{
		"1": 3,
		"2": 3,
		"3": 11,
	}
	for i, v := range expected {
		if r := d.PartTwo(input.ReadFileBytes(fmt.Sprintf("../testdata/y2015/03/2.%v.txt", i))); r != v {
			t.Fatalf("Error with test data file 2.%v.txt. Expected %v, got %v\n", i, v, r)
		}
	}
}
