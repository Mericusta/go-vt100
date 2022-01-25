package vt100

import "fmt"

func ClearScreen() {
	fmt.Printf("\033[2J")
}

func MoveCursorToAndPrint(x, y int, c string) {
	fmt.Printf("\033[%d;%dH%v", y, x, c)
}

func MoveCursorToLine(y int) {
	fmt.Printf("\033[%df\n", y)
}
