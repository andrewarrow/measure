package main

import (
	"fmt"
	"math/rand"
	"measure/args"
	"measure/gallons"
	"os"
	"time"
)

func PrintHelp() {
	fmt.Println("")
	fmt.Println("  measure gallons       # use --x=val1 --y=val2 --z=val3")
	fmt.Println("")
}

var argMap map[string]string

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) == 1 {
		PrintHelp()
		return
	}
	command := os.Args[1]
	argMap = args.ToMap()

	if command == "gallons" {
		gallons.HandleInput(argMap)
	} else if command == "other" {
	} else {
	}

}
