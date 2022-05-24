package main

import (
	"github.com/Mericusta/go-vt100/border"
	"github.com/Mericusta/go-vt100/container"
	"github.com/Mericusta/go-vt100/core"
	"github.com/Mericusta/go-vt100/shape"
	"github.com/Mericusta/go-vt100/terminal"
)

func main() {
	defer terminal.Destruct()
	core.ClearScreen()
	core.CursorInvisible()
	c := container.NewCanvas(core.Size{Width: 64, Height: 17})
	// in canvas
	c.AppendObject(core.NewObject(
		core.Coordinate{X: 0, Y: 0},
		shape.NewRectangle(shape.NewLine(shape.NewPoint('❤'), 5, core.Horizontal), 5),
	), core.NewObject(
		core.Coordinate{X: 1, Y: 6},
		shape.NewRectangle(shape.NewLine(shape.NewPoint(border.CT()), 5, core.Horizontal), 5),
	), core.NewObject(
		core.Coordinate{X: 2, Y: 12},
		shape.NewRectangle(shape.NewLine(shape.NewPoint('*'), 5, core.Horizontal), 5),
	))
	c.Draw(terminal.Context(), core.Coordinate{X: int(terminal.Stdout().Width()-64) / 2, Y: 1})
	<-terminal.ControlSignal
	c.Clear()
	// coincides with the boundary
	c.AppendObject(core.NewObject(
		core.Coordinate{X: 50, Y: -2},
		shape.NewRectangle(shape.NewLine(shape.NewPoint('❤'), 5, core.Horizontal), 5),
	), core.NewObject(
		core.Coordinate{X: 61, Y: 6},
		shape.NewRectangle(shape.NewLine(shape.NewPoint(border.CT()), 5, core.Horizontal), 5),
	), core.NewObject(
		core.Coordinate{X: 50, Y: 14},
		shape.NewRectangle(shape.NewLine(shape.NewPoint('*'), 5, core.Horizontal), 5),
	))
	c.Draw(terminal.Context(), core.Coordinate{X: int(terminal.Stdout().Width()-64) / 2, Y: 1})
	<-terminal.ControlSignal
	c.Clear()
	// out canvas
	c.AppendObject(core.NewObject(
		core.Coordinate{X: 50, Y: -5},
		shape.NewRectangle(shape.NewLine(shape.NewPoint('❤'), 5, core.Horizontal), 5),
	), core.NewObject(
		core.Coordinate{X: 71, Y: 6},
		shape.NewRectangle(shape.NewLine(shape.NewPoint(border.CT()), 5, core.Horizontal), 5),
	), core.NewObject(
		core.Coordinate{X: 50, Y: 19},
		shape.NewRectangle(shape.NewLine(shape.NewPoint('*'), 5, core.Horizontal), 5),
	))
	c.Draw(terminal.Context(), core.Coordinate{X: int(terminal.Stdout().Width()-64) / 2, Y: 1})
	<-terminal.ControlSignal
	c.Clear()
}
