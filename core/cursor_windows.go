package core

import (
	"fmt"
)

// vt100 cursor origin is {1,1}

// MoveCursorToAndPrint
// @param x column index in vt100
// @param y row index in vt100
func MoveCursorToAndPrint(x, y int, c string) {
	fmt.Printf("\033[%d;%dH%v", y+1, x+1, c)
}

func MoveCursorToHome() {
	fmt.Printf("\033[H")
}

func MoveCursorToLine(y int) {
	fmt.Printf("\033[%df\n", y)
}

func CursorInvisible() {
	fmt.Printf("\033[?25l")
}

func CursorVisible() {
	fmt.Printf("\033[?25h")
}
