package y2015_test

import (
	"os"
	"testing"

	"github.com/schrenker/AoC/tools"
	"github.com/schrenker/AoC/y2015"
)

func TestDay19PartOne(t *testing.T) {
	d := y2015.Day19{}
	os.Args = []string{"cmd", "2015", "19", "01"}
	tools.GetInput("2015", "19")
    t.Errorf("No tests defined. Placeholder: %v", d)
}

func TestDay19PartTwo(t *testing.T) {
	d := y2015.Day19{}
	os.Args = []string{"cmd", "2015", "19", "02"}
	tools.GetInput("2015", "19")
    t.Errorf("No tests defined. Placeholder: %v", d)
}
