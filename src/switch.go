package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")

	// In golang break statement is automatically provided.
	// Go's switch cases need not be constants, and the values involved need not be integers.
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}
