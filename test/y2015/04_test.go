package y2015_test

import (
	"os"
	"testing"

	"github.com/schrenker/AoC/tools"
	"github.com/schrenker/AoC/y2015"
)

func TestDay04PartOne(t *testing.T) {
	d := y2015.Day04{}
	os.Args = []string{"cmd", "2015", "04", "01"}
	tools.GetInput("2015", "04")
    t.Errorf("No tests defined. Placeholder: %v", d)
}

func TestDay04PartTwo(t *testing.T) {
	d := y2015.Day04{}
	os.Args = []string{"cmd", "2015", "04", "02"}
	tools.GetInput("2015", "04")
    t.Errorf("No tests defined. Placeholder: %v", d)
}
