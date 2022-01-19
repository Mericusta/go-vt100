package terminal

import (
	"golang.org/x/sys/windows"
)

type windowsTerminal windows.ConsoleScreenBufferInfo

var csbInfo windows.ConsoleScreenBufferInfo

func init() {
	err := windows.GetConsoleScreenBufferInfo(windows.Stdout, &csbInfo)
	if err != nil {
		panic(err)
	}
	terminal = windowsTerminal(csbInfo)
}

func (t windowsTerminal) Width() int {
	return int(t.Size.X)
}

func (t windowsTerminal) Height() int {
	return int(t.Size.Y)
}
