package y2015_test

import (
	"os"
	"testing"

	"github.com/schrenker/AoC/tools"
	"github.com/schrenker/AoC/y2015"
)

func TestDay16PartOne(t *testing.T) {
	d := y2015.Day16{}
	os.Args = []string{"cmd", "2015", "16", "01"}
	tools.GetInput("2015", "16")
    t.Errorf("No tests defined. Placeholder: %v", d)
}

func TestDay16PartTwo(t *testing.T) {
	d := y2015.Day16{}
	os.Args = []string{"cmd", "2015", "16", "02"}
	tools.GetInput("2015", "16")
    t.Errorf("No tests defined. Placeholder: %v", d)
}
