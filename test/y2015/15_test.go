package y2015_test

import (
	"os"
	"testing"

	"github.com/schrenker/AoC/tools"
	"github.com/schrenker/AoC/y2015"
)

func TestDay15PartOne(t *testing.T) {
	d := y2015.Day15{}
	os.Args = []string{"cmd", "2015", "15", "01"}
	tools.GetInput("2015", "15")
    t.Errorf("No tests defined. Placeholder: %v", d)
}

func TestDay15PartTwo(t *testing.T) {
	d := y2015.Day15{}
	os.Args = []string{"cmd", "2015", "15", "02"}
	tools.GetInput("2015", "15")
    t.Errorf("No tests defined. Placeholder: %v", d)
}
