package y2015_test

import (
	"os"
	"testing"

	"github.com/schrenker/AoC/tools"
	"github.com/schrenker/AoC/y2015"
)

func TestDay17PartOne(t *testing.T) {
	d := y2015.Day17{}
	os.Args = []string{"cmd", "2015", "17", "01"}
	tools.GetInput("2015", "17")
    t.Errorf("No tests defined. Placeholder: %v", d)
}

func TestDay17PartTwo(t *testing.T) {
	d := y2015.Day17{}
	os.Args = []string{"cmd", "2015", "17", "02"}
	tools.GetInput("2015", "17")
    t.Errorf("No tests defined. Placeholder: %v", d)
}
