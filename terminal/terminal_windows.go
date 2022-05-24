//go:build !unittest

package terminal

import (
	"github.com/Mericusta/go-vt100/core"
	"golang.org/x/sys/windows"
)

type windowsTerminal windows.ConsoleScreenBufferInfo

var csbInfo windows.ConsoleScreenBufferInfo

func init() {
	var stdoutMode, stdinMode uint32
	if err := windows.GetConsoleMode(windows.Stdout, &stdoutMode); err != nil {
		panic(err.Error())
	}
	if err := windows.GetConsoleMode(windows.Stdin, &stdinMode); err != nil {
		panic(err.Error())
	}
	newStdoutMode := windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING | windows.DISABLE_NEWLINE_AUTO_RETURN | stdoutMode
	newStdinMode := windows.ENABLE_VIRTUAL_TERMINAL_INPUT | stdinMode

	if err := windows.SetConsoleMode(windows.Stdout, newStdoutMode); err != nil {
		panic(err.Error())
	}
	if err := windows.SetConsoleMode(windows.Stdin, newStdinMode); err != nil {
		panic(err.Error())
	}

	if err := windows.GetConsoleScreenBufferInfo(windows.Stdout, &csbInfo); err != nil {
		panic(err.Error())
	}
	terminal = windowsTerminal(csbInfo)
}

func (t windowsTerminal) Width() uint {
	return uint(t.Size.X)
}

func (t windowsTerminal) Height() uint {
	return uint(t.Size.Y)
}

func (t windowsTerminal) Coordinate() core.Coordinate {
	return core.Coordinate{}
}
