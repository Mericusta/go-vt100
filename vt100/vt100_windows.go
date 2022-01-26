package vt100

import "fmt"

func ClearScreen() {
	fmt.Printf("\x1b[2J")
}

func Reset() {
	fmt.Printf("\x1b[0m")
}
