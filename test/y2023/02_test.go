package y2023_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/schrenker/AoC/input"
	"github.com/schrenker/AoC/y2023"
)

func TestDay02PartOne(t *testing.T) {
	d := y2023.Day02{}
	os.Args = []string{"cmd", "2023", "02", "01"}
	expected := map[string]int{
		"1": 8,
	}
	for i, v := range expected {
		if r := d.PartOne(input.ReadFileBytes(fmt.Sprintf("../testdata/y2023/02/1.%v.txt", i))); r != v {
			t.Fatalf("Error with test data file 1.%v.txt. Expected %v, got %v\n", i, v, r)
		}
	}
}

func TestDay02PartTwo(t *testing.T) {
	d := y2023.Day02{}
	os.Args = []string{"cmd", "2023", "02", "02"}
	expected := map[string]int{
		"1": 2286,
	}
	for i, v := range expected {
		if r := d.PartTwo(input.ReadFileBytes(fmt.Sprintf("../testdata/y2023/02/2.%v.txt", i))); r != v {
			t.Fatalf("Error with test data file 2.%v.txt. Expected %v, got %v\n", i, v, r)
		}
	}
}
