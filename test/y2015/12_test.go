package y2015_test

import (
	"os"
	"testing"

	"github.com/schrenker/AoC/tools"
	"github.com/schrenker/AoC/y2015"
)

func TestDay12PartOne(t *testing.T) {
	d := y2015.Day12{}
	os.Args = []string{"cmd", "2015", "12", "01"}
	tools.GetInput("2015", "12")
    t.Errorf("No tests defined. Placeholder: %v", d)
}

func TestDay12PartTwo(t *testing.T) {
	d := y2015.Day12{}
	os.Args = []string{"cmd", "2015", "12", "02"}
	tools.GetInput("2015", "12")
    t.Errorf("No tests defined. Placeholder: %v", d)
}
