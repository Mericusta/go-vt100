package utility

import (
	"fmt"
	"go-vt100/terminal"
	"go-vt100/vt100"
)

func DebugPrintf(y int, format string, content ...interface{}) {
	<-terminal.ControlSignal
	formatContent := fmt.Sprintf(format, content...)
	vt100.MoveCursorToAndPrint(2, y, formatContent)
	<-terminal.ControlSignal
	vt100.ClearLine()
}
