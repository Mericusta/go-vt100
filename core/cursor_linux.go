package core

import (
	"fmt"
)

func MoveCursorToAndPrint(x, y int, c string) {
	fmt.Printf("\033[%d;%dH%v", y+Origin().X, x+Origin().Y, c)
}

func MoveCursorTo(x, y int) {
	fmt.Printf("\033[%d;%d", y, x)
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
