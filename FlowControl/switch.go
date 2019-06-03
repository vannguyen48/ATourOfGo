// A switch statement is a shorter way to write a sequence of if-else statements
// The break statement is provided automatically in Go.
// Go switch cases cannot be constants, and the values involved cannot be integers.
// Switch cases evaluates cases from top to bottom, stopping when a case succeeds.
// Switch without a condition is the same as switch true


package main

import (
	"fmt"
	"runtime"
	"time"
)

func findTime() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}


func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
	findTime()
}
