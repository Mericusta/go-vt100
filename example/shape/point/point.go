package main

import (
	"github.com/Mericusta/go-vt100/character"
	"github.com/Mericusta/go-vt100/core"
	"github.com/Mericusta/go-vt100/shape"
)

func main() {
	defer core.Destruct()
	core.Init()
	core.ClearScreen()
	core.CursorInvisible()
	var d core.Drawable
	d = shape.NewPoint('❤')
	d.Draw(core.Context(), core.Coordinate{X: 0, Y: 0})
	d = shape.NewPoint(character.CT())
	d.Draw(core.Context(), core.Coordinate{X: 0, Y: 1})
	d = shape.NewPoint('*')
	d.Draw(core.Context(), core.Coordinate{X: 0, Y: 2})
	<-core.ControlSignal
	// outer container
	d = shape.NewPoint(character.CT())
	d.Draw(core.Context(), core.Coordinate{X: -1, Y: -1})
	<-core.ControlSignal
}
