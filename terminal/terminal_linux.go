package terminal

import (
	"syscall"

	"golang.org/x/sys/unix"
)

type unixTerminal unix.Winsize

func init() {
	unixWinSize, err := unix.IoctlGetWinsize(syscall.Stdout, unix.TIOCGWINSZ)
	if err != nil {
		panic(err)
	}
	terminal = unixTerminal(&unixWinSize)
}

func (t unixTerminal) Width() int {
	return int(t.Col)
}

func (t unixTerminal) Height() int {
	return int(t.Row)
}
