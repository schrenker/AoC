package y2022_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/schrenker/AoC/tools"
	"github.com/schrenker/AoC/y2022"
)

func TestDay05PartOne(t *testing.T) {
	d := y2022.Day05{}
	os.Args = []string{"cmd", "2022", "05", "01"}
	expected := map[string]string{
		"1": "CMZ",
	}
	for i, v := range expected {
		if r := d.PartOne(tools.ReadFileBytes(fmt.Sprintf("../testdata/y2022/05/1.%v.txt", i))); r != v {
			t.Fatalf("Error with test data file 1.%v.txt. Expected %v, got %v\n", i, v, r)
		}
	}
}

func TestDay05PartTwo(t *testing.T) {
	d := y2022.Day05{}
	os.Args = []string{"cmd", "2022", "05", "02"}
	expected := map[string]string{
		"1": "MCD",
	}
	for i, v := range expected {
		if r := d.PartTwo(tools.ReadFileBytes(fmt.Sprintf("../testdata/y2022/05/2.%v.txt", i))); r != v {
			t.Fatalf("Error with test data file 2.%v.txt. Expected %v, got %v\n", i, v, r)
		}
	}
}