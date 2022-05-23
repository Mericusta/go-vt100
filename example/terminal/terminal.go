package main

import (
	"github.com/Mericusta/go-vt100/core"
	"github.com/Mericusta/go-vt100/terminal"
)

func main() {
	core.ClearScreen()
	core.MoveCursorToAndPrint(0, 0, "0")
	<-terminal.ControlSignal
	core.ClearScreen()
	core.MoveCursorToAndPrint(1, 1, "0")
	<-terminal.ControlSignal
	core.ClearScreen()
	core.MoveCursorToAndPrint(2, 2, "0")
	<-terminal.ControlSignal
	core.ClearScreen()
}
