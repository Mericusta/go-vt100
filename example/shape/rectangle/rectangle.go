package main

import (
	"github.com/Mericusta/go-vt100/core"
	"github.com/Mericusta/go-vt100/shape"
	"github.com/Mericusta/go-vt100/terminal"
)

func main() {
	defer terminal.Destruct()
	core.ClearScreen()
	core.CursorInvisible()
	// rune not in ASCII
	r := shape.NewRectangle(shape.NewLine(
		shape.NewPoint('❤'),
		5, core.Horizontal,
	), 5)
	r.Draw(2, 2, core.Size{
		Width:  terminal.Stdout().Width(),
		Height: terminal.Stdout().Height(),
	})
	r = shape.NewRectangle(shape.NewLine(
		shape.NewPoint('❤'),
		5, core.Vertical,
	), 5)
	r.Draw(22, 2, core.Size{
		Width:  terminal.Stdout().Width(),
		Height: terminal.Stdout().Height(),
	})
	// // rune in ASCII
	r = shape.NewRectangle(shape.NewLine(
		shape.NewPoint('*'),
		5, core.Horizontal,
	), 5)
	r.Draw(3, 7, core.Size{
		Width:  terminal.Stdout().Width(),
		Height: terminal.Stdout().Height(),
	})
	r = shape.NewRectangle(shape.NewLine(
		shape.NewPoint('*'),
		5, core.Vertical,
	), 5)
	r.Draw(23, 7, core.Size{
		Width:  terminal.Stdout().Width(),
		Height: terminal.Stdout().Height(),
	})
	<-terminal.ControlSignal
}
