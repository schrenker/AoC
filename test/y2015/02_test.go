package y2015_test

import (
	"os"
	"testing"

	"github.com/schrenker/AoC/tools"
	"github.com/schrenker/AoC/y2015"
)

func TestDay02PartOne(t *testing.T) {
	d := y2015.Day02{}
	os.Args = []string{"cmd", "2015", "02", "01"}
	tools.GetInput("2015", "02")
    t.Errorf("No tests defined. Placeholder: %v", d)
}

func TestDay02PartTwo(t *testing.T) {
	d := y2015.Day02{}
	os.Args = []string{"cmd", "2015", "02", "02"}
	tools.GetInput("2015", "02")
    t.Errorf("No tests defined. Placeholder: %v", d)
}
