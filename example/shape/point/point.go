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
	p := shape.NewPoint('❤')
	p.Draw(1, 1, core.Size{
		Width:  terminal.Stdout().Width(),
		Height: terminal.Stdout().Height(),
	})
	p = shape.NewPoint(core.CT())
	p.Draw(1, 2, core.Size{
		Width:  terminal.Stdout().Width(),
		Height: terminal.Stdout().Height(),
	})
	p = shape.NewPoint('*')
	p.Draw(1, 3, core.Size{
		Width:  terminal.Stdout().Width(),
		Height: terminal.Stdout().Height(),
	})
	<-terminal.ControlSignal
}
