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
	d = shape.NewLine(shape.NewPoint('❤'), 5, core.Horizontal)
	d.Draw(core.Coordinate{X: 1, Y: 1})
	d = shape.NewLine(shape.NewPoint(border.CT()), 5, core.Horizontal)
	d.Draw(core.Coordinate{X: 1, Y: 2})
	d = shape.NewLine(shape.NewPoint('*'), 5, core.Horizontal)
	d.Draw(core.Coordinate{X: 1, Y: 3})
	d = shape.NewLine(shape.NewPoint('❤'), 5, core.Vertical)
	d.Draw(core.Coordinate{X: 1, Y: 4})
	d = shape.NewLine(shape.NewPoint(border.CT()), 5, core.Vertical)
	d.Draw(core.Coordinate{X: 3, Y: 4})
	d = shape.NewLine(shape.NewPoint('*'), 5, core.Vertical)
	d.Draw(core.Coordinate{X: 4, Y: 4})
	<-terminal.ControlSignal
}
