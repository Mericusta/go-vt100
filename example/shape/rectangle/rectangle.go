package main

import (
	"github.com/Mericusta/go-vt100/border"
	"github.com/Mericusta/go-vt100/core"
	"github.com/Mericusta/go-vt100/shape"
)

func main() {
	defer core.Destruct()
	core.ClearScreen()
	core.CursorInvisible()
	var d core.Drawable
	d = shape.NewRectangle(shape.NewLine(shape.NewPoint('❤'), 5, core.Horizontal), 5)
	d.Draw(core.Context(), core.Coordinate{X: 1, Y: 1})
	d = shape.NewRectangle(shape.NewLine(shape.NewPoint(border.CT()), 5, core.Horizontal), 5)
	d.Draw(core.Context(), core.Coordinate{X: 1, Y: 6})
	d = shape.NewRectangle(shape.NewLine(shape.NewPoint('*'), 5, core.Horizontal), 5)
	d.Draw(core.Context(), core.Coordinate{X: 1, Y: 11})
	<-core.ControlSignal
	// half outer container
	d = shape.NewRectangle(shape.NewLine(shape.NewPoint(border.CT()), 5, core.Horizontal), 5)
	d.Draw(core.Context(), core.Coordinate{X: -2, Y: 16})
	<-core.ControlSignal
	d = shape.NewRectangle(shape.NewLine(shape.NewPoint(border.CT()), 5, core.Vertical), 5)
	d.Draw(core.Context(), core.Coordinate{X: 11, Y: -2})
	<-core.ControlSignal
	// total outer container
	d = shape.NewRectangle(shape.NewLine(shape.NewPoint(border.CT()), 5, core.Horizontal), 5)
	d.Draw(core.Context(), core.Coordinate{X: -5, Y: -5})
	<-core.ControlSignal
}
