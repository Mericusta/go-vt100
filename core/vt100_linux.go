package core

import "fmt"

func ClearScreen() {
	fmt.Printf("\033[2J")
}

func ClearLine() {
	fmt.Printf("\033[2K")
}

func ResetAttribute() {
	fmt.Printf("\033[0m")
}
