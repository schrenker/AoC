package y2015_test

import (
	"os"
	"testing"

	"github.com/schrenker/AoC/tools"
	"github.com/schrenker/AoC/y2015"
)

func TestDay09PartOne(t *testing.T) {
	d := y2015.Day09{}
	os.Args = []string{"cmd", "2015", "09", "01"}
	tools.GetInput("2015", "09")
    t.Errorf("No tests defined. Placeholder: %v", d)
}

func TestDay09PartTwo(t *testing.T) {
	d := y2015.Day09{}
	os.Args = []string{"cmd", "2015", "09", "02"}
	tools.GetInput("2015", "09")
    t.Errorf("No tests defined. Placeholder: %v", d)
}
