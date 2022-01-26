package vt100

import (
	"fmt"
)

func MoveCursorToAndPrint(x, y int, c string) {
	fmt.Printf("\x1b[%d;%dH%v", y, x, c)
}

func MoveCursorToHome() {
	fmt.Printf("\x1b[H")
}

func MoveCursorToLine(y int) {
	fmt.Printf("\x1b[%df\n", y)
}

func CursorInvisible() {
	fmt.Printf("\x1b[?25l")
}

func CursorVisible() {
	fmt.Printf("\x1b[?25h")
}
