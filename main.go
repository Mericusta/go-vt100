package main

import (
	"fmt"
	"go-vt100/terminal"
)

func main() {
	fmt.Printf("width %v\n", terminal.Stdout().Width())
}
