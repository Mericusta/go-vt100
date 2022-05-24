package main

import (
	"github.com/Mericusta/go-vt100/container"
	"github.com/Mericusta/go-vt100/core"
	"github.com/Mericusta/go-vt100/shape"
	"github.com/Mericusta/go-vt100/terminal"
)

func main() {
	defer terminal.Destruct()
	core.ClearScreen()
	core.CursorInvisible()
	p := shape.NewPoint('❤')
	c := container.NewCanvas(core.Size{Width: p.Width(), Height: p.Height()})
	c.AppendObjects(core.Object{D: p})
	g := container.NewGrids(map[uint]map[uint]container.Canvas{
		1: {1: c, 2: c},
		2: {1: c},
	})
	g.Draw(terminal.Context(), core.Coordinate{X: 0, Y: 0})
	<-terminal.ControlSignal
}
