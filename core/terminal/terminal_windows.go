//go:build !debug

package terminal

import (
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

func (t windowsTerminal) Width() int {
	return int(t.Size.X)
}

func (t windowsTerminal) Height() int {
	return int(t.Size.Y)
}
