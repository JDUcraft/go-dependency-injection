package main

import (
	"fmt"
	"main/cmd"
	"os"
)

func main() {
	weeksLeft, err := cmd.ComputeWeeksLeftBeforeNextYear()
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("%s", weeksLeft)
}
