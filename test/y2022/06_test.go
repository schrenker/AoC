package y2022_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/schrenker/AoC/tools"
	"github.com/schrenker/AoC/y2022"
)

func TestDay06PartOne(t *testing.T) {
	d := y2022.Day06{}
	os.Args = []string{"cmd", "2022", "06", "01"}
	expected := map[string]int{
		"1": 7,
		"2": 5,
		"3": 6,
		"4": 10,
		"5": 11,
	}
	for i, v := range expected {
		if r := d.PartOne(tools.ReadFileBytes(fmt.Sprintf("../testdata/y2022/06/1.%v.txt", i))); r != v {
			t.Fatalf("Error with test data file 1.%v.txt. Expected %v, got %v\n", i, v, r)
		}
	}
}

func TestDay06PartTwo(t *testing.T) {
	d := y2022.Day06{}
	os.Args = []string{"cmd", "2022", "06", "02"}
	expected := map[string]int{
		"1": 19,
		"2": 23,
		"3": 23,
		"4": 29,
		"5": 26,
	}
	for i, v := range expected {
		if r := d.PartTwo(tools.ReadFileBytes(fmt.Sprintf("../testdata/y2022/06/2.%v.txt", i))); r != v {
			t.Fatalf("Error with test data file 2.%v.txt. Expected %v, got %v\n", i, v, r)
		}
	}
}
