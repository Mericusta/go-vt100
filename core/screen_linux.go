package core

import "fmt"

func SaveScreen() {
	fmt.Printf("\033[?47h")
}

func RestoreScreen() {
	fmt.Printf("\033[?47l")
}
