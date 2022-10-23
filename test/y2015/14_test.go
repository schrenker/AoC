package y2015_test

import (
	"os"
	"testing"

	"github.com/schrenker/AoC/tools"
	"github.com/schrenker/AoC/y2015"
)

func TestDay14PartOne(t *testing.T) {
	d := y2015.Day14{}
	os.Args = []string{"cmd", "2015", "14", "01"}
	tools.GetInput("2015", "14")
    t.Errorf("No tests defined. Placeholder: %v", d)
}

func TestDay14PartTwo(t *testing.T) {
	d := y2015.Day14{}
	os.Args = []string{"cmd", "2015", "14", "02"}
	tools.GetInput("2015", "14")
    t.Errorf("No tests defined. Placeholder: %v", d)
}
