//go:build !unittest

package core

import (
	"syscall"

	"golang.org/x/sys/unix"
)

type unixTerminal unix.Winsize

func init() {
	unixWinSize, err := unix.IoctlGetWinsize(syscall.Stdout, unix.TIOCGWINSZ)
	if err != nil {
		panic(err.Error())
	}
	terminal = unixTerminal(*unixWinSize)
}

func (t unixTerminal) Width() uint {
	return uint(t.Col)
}

func (t unixTerminal) Height() uint {
	return uint(t.Row)
}
