package y2015_test

import (
	"os"
	"testing"

	"github.com/schrenker/AoC/tools"
	"github.com/schrenker/AoC/y2015"
)

func TestDay01PartOne(t *testing.T) {
	d := y2015.Day01{}
	os.Args = []string{"cmd", "2015", "01", "01"}
	tools.GetInput("2015", "01")
    t.Errorf("No tests defined. Placeholder: %v", d)
}

func TestDay01PartTwo(t *testing.T) {
	d := y2015.Day01{}
	os.Args = []string{"cmd", "2015", "01", "02"}
	tools.GetInput("2015", "01")
    t.Errorf("No tests defined. Placeholder: %v", d)
}
