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
	p := shape.NewLine(
		shape.NewPoint('❤'),
		5, core.Horizontal,
	)
	p.Draw(2, 2, core.Size{
		Width:  terminal.Stdout().Width(),
		Height: terminal.Stdout().Height(),
	})
	p = shape.NewLine(
		shape.NewPoint('❤'),
		5, core.Vertical,
	)
	p.Draw(2, 2, core.Size{
		Width:  terminal.Stdout().Width(),
		Height: terminal.Stdout().Height(),
	})
	p = shape.NewLine(
		shape.NewPoint(core.CT()),
		5, core.Horizontal,
	)
	p.Draw(3, 7, core.Size{
		Width:  terminal.Stdout().Width(),
		Height: terminal.Stdout().Height(),
	})
	p = shape.NewLine(
		shape.NewPoint(core.CT()),
		5, core.Vertical,
	)
	p.Draw(3, 7, core.Size{
		Width:  terminal.Stdout().Width(),
		Height: terminal.Stdout().Height(),
	})
	p = shape.NewLine(
		shape.NewPoint('*'),
		5, core.Horizontal,
	)
	p.Draw(3, 12, core.Size{
		Width:  terminal.Stdout().Width(),
		Height: terminal.Stdout().Height(),
	})
	p = shape.NewLine(
		shape.NewPoint('*'),
		5, core.Vertical,
	)
	p.Draw(3, 12, core.Size{
		Width:  terminal.Stdout().Width(),
		Height: terminal.Stdout().Height(),
	})
	<-terminal.ControlSignal
}
