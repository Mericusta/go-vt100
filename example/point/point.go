package main

import (
	"github.com/Mericusta/go-vt100/core"
	"github.com/Mericusta/go-vt100/core/terminal"
	"github.com/Mericusta/go-vt100/shape"
)

func main() {
	c := core.NewCanvasWithBoundary(terminal.Stdout().Width(), terminal.Stdout().Height())
	defer c.Destruct()

	c.AddLayerObject(core.Coordinate{X: 0, Y: 0}, shape.NewPoint('❤'))

	c.Draw()
	<-terminal.ControlSignal
}
