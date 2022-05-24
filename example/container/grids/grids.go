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
	g := container.NewGrids(map[uint]map[uint]core.Drawable{
		1: {1: shape.NewPoint('❤')},
	})
	g.Draw(terminal.Context(), core.Coordinate{X: int(terminal.Stdout().Width()-64) / 2, Y: 1})
	<-terminal.ControlSignal
}
