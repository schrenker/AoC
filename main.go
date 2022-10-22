package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/schrenker/AoC/tools"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		panic(fmt.Sprintf("3 Arguments are required [year day part], but %v were provided", len(args)))
	}

	if i, _ := strconv.Atoi(args[2]); i == 1 {
		fmt.Println(challenges[args[0]+"/"+args[1]].PartOne(tools.GetDefaultInputPath()))
	} else if i == 2 {
		fmt.Println(challenges[args[0]+"/"+args[1]].PartTwo(tools.GetDefaultInputPath()))
	} else {
		panic(fmt.Sprintf("unrecognized part value: %v. Expected [01, 1, 02, 2]", args[2]))
	}
}
