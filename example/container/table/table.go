package main

import (
	"github.com/Mericusta/go-vt100/core"
	"github.com/Mericusta/go-vt100/terminal"
)

func main() {
	defer terminal.Destruct()
	core.ClearScreen()
	core.CursorInvisible()

	// example 1
	// header1Slice := []string{"header 1", "header 2", "header 3", "operation"}
	// container.NewTable(header1Slice, )
}
