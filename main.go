package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) < 3 {
		panic(fmt.Sprintf("3 Arguments are required [year day part], but %v were provided", len(args)))
	}

	getInput(args[0], args[1])

	if i, _ := strconv.Atoi(args[2]); i == 1 {
		challenges[args[0]+"/"+args[1]].PartOne()
	} else if i == 2 {
		challenges[args[0]+"/"+args[1]].PartTwo()
	} else {
		panic(fmt.Sprintf("unrecognized part value: %v. Expected [01, 1, 02, 2]", args[2]))
	}
}
