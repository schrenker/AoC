package y2015_test

import (
	"os"
	"testing"

	"github.com/schrenker/AoC/tools"
	"github.com/schrenker/AoC/y2015"
)

func TestDay07PartOne(t *testing.T) {
	d := y2015.Day07{}
	os.Args = []string{"cmd", "2015", "07", "01"}
	tools.GetInput("2015", "07")
    t.Errorf("No tests defined. Placeholder: %v", d)
}

func TestDay07PartTwo(t *testing.T) {
	d := y2015.Day07{}
	os.Args = []string{"cmd", "2015", "07", "02"}
	tools.GetInput("2015", "07")
    t.Errorf("No tests defined. Placeholder: %v", d)
}
