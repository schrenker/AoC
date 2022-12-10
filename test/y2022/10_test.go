package y2022_test

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/schrenker/AoC/tools"
	"github.com/schrenker/AoC/y2022"
)

func TestDay10PartOne(t *testing.T) {
	d := y2022.Day10{}
	os.Args = []string{"cmd", "2022", "10", "01"}
	expected := map[string]int{
		"1": 13140,
	}
	for i, v := range expected {
		if r := d.PartOne(tools.ReadFileBytes(fmt.Sprintf("../testdata/y2022/10/1.%v.txt", i))); r != v {
			t.Fatalf("Error with test data file 1.%v.txt. Expected %v, got %v\n", i, v, r)
		}
	}
}

func TestDay10PartTwo(t *testing.T) {
	d := y2022.Day10{}
	os.Args = []string{"cmd", "2022", "10", "02"}
	expected := map[string][]string{
		"1": {
			"##..##..##..##..##..##..##..##..##..##..",
			"###...###...###...###...###...###...###.",
			"####....####....####....####....####....",
			"#####.....#####.....#####.....#####.....",
			"######......######......######......####",
			"#######.......#######.......#######.....",
		},
	}
	for i, v := range expected {
		if r := d.PartTwo(tools.ReadFileBytes(fmt.Sprintf("../testdata/y2022/10/2.%v.txt", i))); !reflect.DeepEqual(r, v) {
			t.Fatalf("Error with test data file 2.%v.txt. Expected %v, got %v\n", i, v, r)
		}
	}
}
