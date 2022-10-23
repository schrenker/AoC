package y2015_test

import (
	"os"
	"testing"

	"github.com/schrenker/AoC/tools"
	"github.com/schrenker/AoC/y2015"
)

func TestDay05PartOne(t *testing.T) {
	d := y2015.Day05{}
	os.Args = []string{"cmd", "2015", "05", "01"}
	tools.GetInput("2015", "05")
    t.Errorf("No tests defined. Placeholder: %v", d)
}

func TestDay05PartTwo(t *testing.T) {
	d := y2015.Day05{}
	os.Args = []string{"cmd", "2015", "05", "02"}
	tools.GetInput("2015", "05")
    t.Errorf("No tests defined. Placeholder: %v", d)
}
