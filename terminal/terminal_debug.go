//go:build unittest

package terminal

type debugTerminal struct{}

func init() {
	terminal = debugTerminal{}
}

func (t debugTerminal) Width() uint {
	return 128
}

func (t debugTerminal) Height() uint {
	return 64
}
