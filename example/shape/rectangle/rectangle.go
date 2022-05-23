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
	r = shape.NewRectangle(shape.NewLine(
		shape.NewPoint(core.CT()),
		5, core.Horizontal,
	), 5)
	r.Draw(2, 7, core.Size{
		Width:  terminal.Stdout().Width(),
		Height: terminal.Stdout().Height(),
	})
	r = shape.NewRectangle(shape.NewLine(
		shape.NewPoint(core.CT()),
		5, core.Vertical,
	), 5)
	r.Draw(22, 7, core.Size{
		Width:  terminal.Stdout().Width(),
		Height: terminal.Stdout().Height(),
	})
	r = shape.NewRectangle(shape.NewLine(
		shape.NewPoint('*'),
		5, core.Horizontal,
	), 5)
	r.Draw(2, 12, core.Size{
		Width:  terminal.Stdout().Width(),
		Height: terminal.Stdout().Height(),
	})
	r = shape.NewRectangle(shape.NewLine(
		shape.NewPoint('*'),
		5, core.Vertical,
	), 5)
	r.Draw(22, 12, core.Size{
		Width:  terminal.Stdout().Width(),
		Height: terminal.Stdout().Height(),
	})
	<-terminal.ControlSignal
}
