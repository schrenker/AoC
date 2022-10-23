package y2015_test

import (
	"os"
	"testing"

	"github.com/schrenker/AoC/tools"
	"github.com/schrenker/AoC/y2015"
)

func TestDay11PartOne(t *testing.T) {
	d := y2015.Day11{}
	os.Args = []string{"cmd", "2015", "11", "01"}
	tools.GetInput("2015", "11")
    t.Errorf("No tests defined. Placeholder: %v", d)
}

func TestDay11PartTwo(t *testing.T) {
	d := y2015.Day11{}
	os.Args = []string{"cmd", "2015", "11", "02"}
	tools.GetInput("2015", "11")
    t.Errorf("No tests defined. Placeholder: %v", d)
}
