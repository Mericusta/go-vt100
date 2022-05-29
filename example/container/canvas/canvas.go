package main

import (
	"github.com/Mericusta/go-vt100/character"
	"github.com/Mericusta/go-vt100/container"
	"github.com/Mericusta/go-vt100/core"
	"github.com/Mericusta/go-vt100/shape"
)

func main() {
	defer core.Destruct()
	core.ClearScreen()
	core.CursorInvisible()
	c := container.NewCanvas(core.Size{Width: 64, Height: 17}, true)
	// in canvas
	c.AppendObjects(core.NewObject(
		core.Coordinate{X: 0, Y: 0},
		shape.NewRectangle(shape.NewLine(shape.NewPoint('❤'), 5, core.Horizontal), 5),
	), core.NewObject(
		core.Coordinate{X: 1, Y: 6},
		shape.NewRectangle(shape.NewLine(shape.NewPoint(character.CT()), 5, core.Horizontal), 5),
	), core.NewObject(
		core.Coordinate{X: 2, Y: 12},
		shape.NewRectangle(shape.NewLine(shape.NewPoint('*'), 5, core.Horizontal), 5),
	))
	c.Draw(core.Context(), core.Coordinate{X: int(core.Stdout().Width()-64) / 2, Y: 0})
	<-core.ControlSignal
	c.Clear()
	c.ClearObjects()
	<-core.ControlSignal

	// coincides with the boundary
	c.AppendObjects(core.NewObject(
		core.Coordinate{X: 50, Y: -2},
		shape.NewRectangle(shape.NewLine(shape.NewPoint('❤'), 5, core.Horizontal), 5),
	), core.NewObject(
		core.Coordinate{X: 61, Y: 6},
		shape.NewRectangle(shape.NewLine(shape.NewPoint(character.CT()), 5, core.Horizontal), 5),
	), core.NewObject(
		core.Coordinate{X: 50, Y: 14},
		shape.NewRectangle(shape.NewLine(shape.NewPoint('*'), 5, core.Horizontal), 5),
	))
	c.Draw(core.Context(), core.Coordinate{X: int(core.Stdout().Width()-64) / 2, Y: 0})
	<-core.ControlSignal
	c.Clear()
	c.ClearObjects()
	<-core.ControlSignal

	// out canvas
	c.AppendObjects(core.NewObject(
		core.Coordinate{X: 50, Y: -5},
		shape.NewRectangle(shape.NewLine(shape.NewPoint('❤'), 5, core.Horizontal), 5),
	), core.NewObject(
		core.Coordinate{X: 71, Y: 6},
		shape.NewRectangle(shape.NewLine(shape.NewPoint(character.CT()), 5, core.Horizontal), 5),
	), core.NewObject(
		core.Coordinate{X: 50, Y: 19},
		shape.NewRectangle(shape.NewLine(shape.NewPoint('*'), 5, core.Horizontal), 5),
	))
	c.Draw(core.Context(), core.Coordinate{X: int(core.Stdout().Width()-64) / 2, Y: 0})
	<-core.ControlSignal
	c.Clear()
	c.ClearObjects()
	<-core.ControlSignal

	// canvas out terminal
	c.Draw(core.Context(), core.Coordinate{X: -15, Y: -15})
	<-core.ControlSignal
	c.Clear()
	c.ClearObjects()
	<-core.ControlSignal

	// canvas resize
	c.Resize(core.Size{Width: 10, Height: 5})
	c.Draw(core.Context(), core.Coordinate{X: 0, Y: 0})
	<-core.ControlSignal
	c.Clear()
	c.ClearObjects()
	<-core.ControlSignal

	// canvas without boundry
	c = container.NewCanvas(core.Size{Width: 64, Height: 17}, false)
	c.AppendObjects(core.NewObject(
		core.Coordinate{X: 0, Y: 0},
		shape.NewRectangle(shape.NewLine(shape.NewPoint('❤'), 5, core.Horizontal), 5),
	), core.NewObject(
		core.Coordinate{X: 1, Y: 6},
		shape.NewRectangle(shape.NewLine(shape.NewPoint(character.CT()), 5, core.Horizontal), 5),
	), core.NewObject(
		core.Coordinate{X: 2, Y: 12},
		shape.NewRectangle(shape.NewLine(shape.NewPoint('*'), 5, core.Horizontal), 5),
	))
	c.Draw(core.Context(), core.Coordinate{X: 0, Y: 0})
	c = container.NewCanvas(core.Size{Width: 64, Height: 17}, true)
	c.AppendObjects(core.NewObject(
		core.Coordinate{X: 0, Y: 0},
		shape.NewRectangle(shape.NewLine(shape.NewPoint('❤'), 5, core.Horizontal), 5),
	), core.NewObject(
		core.Coordinate{X: 1, Y: 6},
		shape.NewRectangle(shape.NewLine(shape.NewPoint(character.CT()), 5, core.Horizontal), 5),
	), core.NewObject(
		core.Coordinate{X: 2, Y: 12},
		shape.NewRectangle(shape.NewLine(shape.NewPoint('*'), 5, core.Horizontal), 5),
	))
	c.Draw(core.Context(), core.Coordinate{X: 30, Y: 0})
	<-core.ControlSignal
	c.Clear()
	c.ClearObjects()
	<-core.ControlSignal
}
