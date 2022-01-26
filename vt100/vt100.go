package vt100

import "fmt"

func ClearScreen() {
	fmt.Printf("\033[2J\033[H")
}
