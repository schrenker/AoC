package y2023_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/schrenker/AoC/input"
	"github.com/schrenker/AoC/y2023"
)

func TestDay01PartOne(t *testing.T) {
	d := y2023.Day01{}
	os.Args = []string{"cmd", "2023", "01", "01"}
	expected := map[string]int{
		"1": 142,
	}
	for i, v := range expected {
		if r := d.PartOne(input.ReadFileBytes(fmt.Sprintf("../testdata/y2023/01/1.%v.txt", i))); r != v {
			t.Fatalf("Error with test data file 1.%v.txt. Expected %v, got %v\n", i, v, r)
		}
	}
}

func TestDay01PartTwo(t *testing.T) {
	d := y2023.Day01{}
	os.Args = []string{"cmd", "2023", "01", "02"}
	expected := map[string]int{
		"1": 281,
	}
	for i, v := range expected {
		if r := d.PartTwo(input.ReadFileBytes(fmt.Sprintf("../testdata/y2023/01/2.%v.txt", i))); r != v {
			t.Fatalf("Error with test data file 2.%v.txt. Expected %v, got %v\n", i, v, r)
		}
	}
}
