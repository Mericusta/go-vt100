package main

import (
	"github.com/Mericusta/go-vt100/core"
	"github.com/Mericusta/go-vt100/core/terminal"
)

func main() {
	c := core.NewCanvasWithBoundary(128, 64)
	defer c.Destruct()
	c.Draw()
	<-terminal.ControlSignal
}
