package main

import (
	"github.com/Mericusta/go-vt100/border"
	"github.com/Mericusta/go-vt100/core"
	"github.com/Mericusta/go-vt100/shape"
	"github.com/Mericusta/go-vt100/terminal"
)

func main() {
	defer terminal.Destruct()
	core.ClearScreen()
	core.CursorInvisible()
	var d core.Drawable
	d = shape.NewPoint('❤')
	d.Draw(core.Coordinate{X: 1, Y: 1})
	d = shape.NewPoint(border.CT())
	d.Draw(core.Coordinate{X: 1, Y: 2})
	d = shape.NewPoint('*')
	d.Draw(core.Coordinate{X: 1, Y: 3})
	<-terminal.ControlSignal
}
