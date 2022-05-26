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
	d = shape.NewLine(shape.NewPoint('❤'), 5, core.Horizontal)
	d.Draw(core.Context(), core.Coordinate{X: 1, Y: 1})
	d = shape.NewLine(shape.NewPoint(border.CT()), 5, core.Horizontal)
	d.Draw(core.Context(), core.Coordinate{X: 1, Y: 2})
	d = shape.NewLine(shape.NewPoint('*'), 5, core.Horizontal)
	d.Draw(core.Context(), core.Coordinate{X: 1, Y: 3})
	d = shape.NewLine(shape.NewPoint('❤'), 5, core.Vertical)
	d.Draw(core.Context(), core.Coordinate{X: 1, Y: 4})
	d = shape.NewLine(shape.NewPoint(border.CT()), 5, core.Vertical)
	d.Draw(core.Context(), core.Coordinate{X: 3, Y: 4})
	d = shape.NewLine(shape.NewPoint('*'), 5, core.Vertical)
	d.Draw(core.Context(), core.Coordinate{X: 4, Y: 4})
	<-core.ControlSignal
	// half outer container
	d = shape.NewLine(shape.NewPoint(border.CT()), 5, core.Horizontal)
	d.Draw(core.Context(), core.Coordinate{X: -2, Y: 15})
	<-core.ControlSignal
	d = shape.NewLine(shape.NewPoint(border.CT()), 5, core.Vertical)
	d.Draw(core.Context(), core.Coordinate{X: 15, Y: -2})
	<-core.ControlSignal
	// total outer container
	d = shape.NewLine(shape.NewPoint(border.CT()), 5, core.Vertical)
	d.Draw(core.Context(), core.Coordinate{X: -1, Y: -1})
	<-core.ControlSignal
}
