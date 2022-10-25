package y2015_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/schrenker/AoC/tools"
	"github.com/schrenker/AoC/y2015"
)

func TestDay03PartOne(t *testing.T) {
	d := y2015.Day03{}
	os.Args = []string{"cmd", "2015", "03", "01"}
	tools.GetInput("2015", "03")
	expected := map[string]int{
		"1": 0,
	}
	for i, v := range expected {
		if r := d.PartOne(fmt.Sprintf("../testdata/y2015/03/1.%v.txt", i)); r != v {
			t.Fatalf("Error with test data file 1.%v.txt. Expected %v, got %v\n", i, v, r)
		}
	}
}

func TestDay03PartTwo(t *testing.T) {
	d := y2015.Day03{}
	os.Args = []string{"cmd", "2015", "03", "02"}
	tools.GetInput("2015", "03")
	expected := map[string]int{
		"1": 0,
	}
	for i, v := range expected {
		if r := d.PartTwo(fmt.Sprintf("../testdata/y2015/03/2.%v.txt", i)); r != v {
			t.Fatalf("Error with test data file 2.%v.txt. Expected %v, got %v\n", i, v, r)
		}
	}
}
