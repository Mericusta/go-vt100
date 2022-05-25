package main

import (
	"github.com/Mericusta/go-vt100/container"
	"github.com/Mericusta/go-vt100/core"
	"github.com/Mericusta/go-vt100/terminal"
)

func main() {
	defer terminal.Destruct()
	core.ClearScreen()
	core.CursorInvisible()
	// ASCII textarea
	t := container.NewTextarea("abcde", core.Horizontal)
	t.Draw(terminal.Context(), core.Coordinate{X: 0, Y: 0})
	<-terminal.ControlSignal

	// ASCII textarea and Over Int32
	t = container.NewTextarea("ab❤de", core.Horizontal)
	t.Draw(terminal.Context(), core.Coordinate{X: 1, Y: 3})
	<-terminal.ControlSignal

	// ASCII multi-line textarea and Over Int32
	t = container.NewTextarea("ab❤de\nI❤China!", core.Horizontal)
	t.Draw(terminal.Context(), core.Coordinate{X: 1, Y: 6})
	<-terminal.ControlSignal

	// ASCII multi-line textarea and Over Int32
	t = container.NewTextarea("ab❤de\nI❤China!", core.Vertical)
	t.Draw(terminal.Context(), core.Coordinate{X: 12, Y: 6})
	<-terminal.ControlSignal
}
