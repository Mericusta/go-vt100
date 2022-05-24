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
	d = shape.NewRectangle(shape.NewLine(shape.NewPoint('❤'), 5, core.Horizontal), 5)
	d.Draw(core.Coordinate{X: 1, Y: 1})
	d = shape.NewRectangle(shape.NewLine(shape.NewPoint(border.CT()), 5, core.Horizontal), 5)
	d.Draw(core.Coordinate{X: 1, Y: 6})
	d = shape.NewRectangle(shape.NewLine(shape.NewPoint('*'), 5, core.Horizontal), 5)
	d.Draw(core.Coordinate{X: 1, Y: 11})
	<-terminal.ControlSignal
}
