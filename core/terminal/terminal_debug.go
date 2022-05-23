//go:build debug

package terminal

type debugTerminal struct{}

func init() {
	terminal = debugTerminal{}
}

func (t debugTerminal) Width() int {
	return 128
}

func (t debugTerminal) Height() int {
	return 64
}
