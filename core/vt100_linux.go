package core

import "fmt"

func ClearScreen() {
	fmt.Printf("\033[2J")
}

func Reset() {
	fmt.Printf("\033[0m")
}
