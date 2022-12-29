package core

import (
	"fmt"
)

// vt100 cursor origin is {1,1}

// MoveCursorToAndPrint
// @param x column index in vt100, absolute coordinate of terminal
// @param y row index in vt100, absolute coordinate of terminal
func MoveCursorToAndPrint(x, y int, c string) {
	// DebugOutput(func() {
	// 	fmt.Printf("c = %v, x = %v, y = %v\n", c, x, y)
	// }, nil)
	fmt.Printf("\033[%d;%dH%v", y+Origin().X, x+Origin().Y, c)
}

func MoveCursorTo(x, y int) {
	fmt.Printf("\033[%d;%d", y, x)
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
