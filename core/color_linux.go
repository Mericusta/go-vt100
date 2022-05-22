package core

import (
	"fmt"
)

func SetForegroundColor(c Color) {
	if 30 <= c && c <= 39 {
		fmt.Printf("\033[%vm", c)
	}
}

func ClearForegroundColor() {
	fmt.Printf("\033[%vm", Default)
}

func SetBackgroundColor(c Color) {
	if 40 <= c+10 && c+10 <= 49 {
		fmt.Printf("\033[%vm", c+10)
	}
}

func ClearBackgroundColor() {
	fmt.Printf("\033[%vm", Default+10)
}
