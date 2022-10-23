package y2015_test

import (
	"os"
	"testing"

	"github.com/schrenker/AoC/tools"
	"github.com/schrenker/AoC/y2015"
)

func TestDay08PartOne(t *testing.T) {
	d := y2015.Day08{}
	os.Args = []string{"cmd", "2015", "08", "01"}
	tools.GetInput("2015", "08")
    t.Errorf("No tests defined. Placeholder: %v", d)
}

func TestDay08PartTwo(t *testing.T) {
	d := y2015.Day08{}
	os.Args = []string{"cmd", "2015", "08", "02"}
	tools.GetInput("2015", "08")
    t.Errorf("No tests defined. Placeholder: %v", d)
}
