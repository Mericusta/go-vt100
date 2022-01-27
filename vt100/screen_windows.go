package vt100

import "fmt"

func SaveScreen() {
	fmt.Printf("\033[?47h")
}

func RestoreScreen() {
	fmt.Printf("\033[?47l")
}
