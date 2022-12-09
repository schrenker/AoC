package y2022_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/schrenker/AoC/tools"
	"github.com/schrenker/AoC/y2022"
)

func TestDay09PartOne(t *testing.T) {
	d := y2022.Day09{}
	os.Args = []string{"cmd", "2022", "09", "01"}
	expected := map[string]int{
		"1": 13,
	}
	for i, v := range expected {
		if r := d.PartOne(tools.ReadFileBytes(fmt.Sprintf("../testdata/y2022/09/1.%v.txt", i))); r != v {
			t.Fatalf("Error with test data file 1.%v.txt. Expected %v, got %v\n", i, v, r)
		}
	}
}

func TestDay09PartTwo(t *testing.T) {
	d := y2022.Day09{}
	os.Args = []string{"cmd", "2022", "09", "02"}
	expected := map[string]int{
		"1": 1,
		"2": 36,
	}
	for i, v := range expected {
		if r := d.PartTwo(tools.ReadFileBytes(fmt.Sprintf("../testdata/y2022/09/2.%v.txt", i))); r != v {
			t.Fatalf("Error with test data file 2.%v.txt. Expected %v, got %v\n", i, v, r)
		}
	}
}
