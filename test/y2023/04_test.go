package y2023_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/schrenker/AoC/input"
	"github.com/schrenker/AoC/y2023"
)

func TestDay04PartOne(t *testing.T) {
	d := y2023.Day04{}
	os.Args = []string{"cmd", "2023", "04", "01"}
	expected := map[string]int{
		"1": 13,
	}
	for i, v := range expected {
		if r := d.PartOne(input.ReadFileBytes(fmt.Sprintf("../testdata/y2023/04/1.%v.txt", i))); r != v {
			t.Fatalf("Error with test data file 1.%v.txt. Expected %v, got %v\n", i, v, r)
		}
	}
}

func TestDay04PartTwo(t *testing.T) {
	d := y2023.Day04{}
	os.Args = []string{"cmd", "2023", "04", "02"}
	expected := map[string]int{
		"1": 30,
	}
	for i, v := range expected {
		if r := d.PartTwo(input.ReadFileBytes(fmt.Sprintf("../testdata/y2023/04/2.%v.txt", i))); r != v {
			t.Fatalf("Error with test data file 2.%v.txt. Expected %v, got %v\n", i, v, r)
		}
	}
}
