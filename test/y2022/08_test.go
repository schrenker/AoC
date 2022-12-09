package y2022_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/schrenker/AoC/tools"
	"github.com/schrenker/AoC/y2022"
)

func TestDay08PartOne(t *testing.T) {
	d := y2022.Day08{}
	os.Args = []string{"cmd", "2022", "08", "01"}
	expected := map[string]int{
		"1": 21,
	}
	for i, v := range expected {
		if r := d.PartOne(tools.ReadFileBytes(fmt.Sprintf("../testdata/y2022/08/1.%v.txt", i))); r != v {
			t.Fatalf("Error with test data file 1.%v.txt. Expected %v, got %v\n", i, v, r)
		}
	}
}

func TestDay08PartTwo(t *testing.T) {
	d := y2022.Day08{}
	os.Args = []string{"cmd", "2022", "08", "02"}
	expected := map[string]int{
		"1": 8,
	}
	for i, v := range expected {
		if r := d.PartTwo(tools.ReadFileBytes(fmt.Sprintf("../testdata/y2022/08/2.%v.txt", i))); r != v {
			t.Fatalf("Error with test data file 2.%v.txt. Expected %v, got %v\n", i, v, r)
		}
	}
}
