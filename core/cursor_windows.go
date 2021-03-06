package core

import (
	"fmt"
)

// vt100 cursor origin is {1,1}

// MoveCursorToAndPrint
// @param x column index in vt100, absolute coordinate of terminal
// @param y row index in vt100, absolute coordinate of terminal
func MoveCursorToAndPrint(x, y uint, c string) {
	// DebugOutput(func() {
	// 	fmt.Printf("c = %v, x = %v, y = %v\n", c, x, y)
	// }, nil)
	fmt.Printf("\033[%d;%dH%v", y+uint(Origin().X), x+uint(Origin().Y), c)
}

func MoveCursorToHome() {
	fmt.Printf("\033[H")
}

func MoveCursorToLine(y uint) {
	fmt.Printf("\033[%df\n", y)
}

func CursorInvisible() {
	fmt.Printf("\033[?25l")
}

func CursorVisible() {
	fmt.Printf("\033[?25h")
}
