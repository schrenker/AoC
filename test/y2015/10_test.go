package y2015_test

import (
	"os"
	"testing"

	"github.com/schrenker/AoC/tools"
	"github.com/schrenker/AoC/y2015"
)

func TestDay10PartOne(t *testing.T) {
	d := y2015.Day10{}
	os.Args = []string{"cmd", "2015", "10", "01"}
	tools.GetInput("2015", "10")
    t.Errorf("No tests defined. Placeholder: %v", d)
}

func TestDay10PartTwo(t *testing.T) {
	d := y2015.Day10{}
	os.Args = []string{"cmd", "2015", "10", "02"}
	tools.GetInput("2015", "10")
    t.Errorf("No tests defined. Placeholder: %v", d)
}
