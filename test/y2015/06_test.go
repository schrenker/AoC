package y2015_test

import (
	"os"
	"testing"

	"github.com/schrenker/AoC/tools"
	"github.com/schrenker/AoC/y2015"
)

func TestDay06PartOne(t *testing.T) {
	d := y2015.Day06{}
	os.Args = []string{"cmd", "2015", "06", "01"}
	tools.GetInput("2015", "06")
    t.Errorf("No tests defined. Placeholder: %v", d)
}

func TestDay06PartTwo(t *testing.T) {
	d := y2015.Day06{}
	os.Args = []string{"cmd", "2015", "06", "02"}
	tools.GetInput("2015", "06")
    t.Errorf("No tests defined. Placeholder: %v", d)
}
