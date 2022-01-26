package vt100

import (
	"fmt"
	"go-vt100/color"
)

func SetForegroundColor(c color.Color) {
	if 30 <= c && c <= 39 {
		fmt.Printf("\x1b[%vm", c)
	}
}

func ClearForegroundColor() {
	fmt.Printf("\x1b[%vm", color.Default)
}

func SetBackgroundColor(c color.Color) {
	if 40 <= c+10 && c+10 <= 49 {
		fmt.Printf("\x1b[%vm", c+10)
	}
}

func ClearBackgroundColor() {
	fmt.Printf("\x1b[%vm", color.Default+10)
}
