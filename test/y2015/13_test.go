package y2015_test

import (
	"os"
	"testing"

	"github.com/schrenker/AoC/tools"
	"github.com/schrenker/AoC/y2015"
)

func TestDay13PartOne(t *testing.T) {
	d := y2015.Day13{}
	os.Args = []string{"cmd", "2015", "13", "01"}
	tools.GetInput("2015", "13")
    t.Errorf("No tests defined. Placeholder: %v", d)
}

func TestDay13PartTwo(t *testing.T) {
	d := y2015.Day13{}
	os.Args = []string{"cmd", "2015", "13", "02"}
	tools.GetInput("2015", "13")
    t.Errorf("No tests defined. Placeholder: %v", d)
}
