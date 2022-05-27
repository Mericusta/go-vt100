package main

import (
	"github.com/Mericusta/go-vt100/container"
	"github.com/Mericusta/go-vt100/core"
	"github.com/Mericusta/go-vt100/shape"
)

func main() {
	defer core.Destruct()
	core.ClearScreen()
	core.CursorInvisible()
	// 3 point expand to 2x2 grid, each grid size is 2x1
	p := shape.NewPoint('❤')
	c := container.NewCanvas(core.Size{Width: p.Width(), Height: p.Height()})
	c.AppendObjects(core.Object{D: p})
	g := container.NewGrids(map[uint]map[uint]container.Canvas{
		1: {1: c, 2: c},
		2: {1: c},
	})
	g.Draw(core.Context(), core.Coordinate{X: 0, Y: 0})
	<-core.ControlSignal
	g.Clear()
	<-core.ControlSignal
	// 3 length 3 line expand to 2x2 grid, each grid size is
	s1 := shape.NewPoint('❤')
	c1 := container.NewCanvas(core.Size{Width: s1.Width(), Height: s1.Height()})
	c1.AppendObjects(core.Object{D: s1})

	s2 := shape.NewLine(s1, 4, core.Horizontal)
	c2 := container.NewCanvas(core.Size{Width: s2.Width(), Height: s2.Height()})
	c2.AppendObjects(core.Object{D: s2})

	s3 := shape.NewLine(s1, 4, core.Vertical)
	c3 := container.NewCanvas(core.Size{Width: s3.Width(), Height: s3.Height()})
	c3.AppendObjects(core.Object{D: s3})

	s4 := shape.NewRectangle(shape.NewLine(s1, 5, core.Horizontal), 5)
	c4 := container.NewCanvas(core.Size{Width: s4.Width(), Height: s4.Height()})
	c4.AppendObjects(core.Object{D: s4})

	g1 := container.NewGrids(map[uint]map[uint]container.Canvas{
		1: {1: c1, 2: c2},
		2: {1: c3},
	})
	g1.Draw(core.Context(), core.Coordinate{X: int(core.Stdout().Width()-g1.Width()) / 2, Y: 1})
	<-core.ControlSignal
	g1.Clear()
	<-core.ControlSignal

	g1.SetCanvas(map[uint]map[uint]container.Canvas{
		2: {2: c4},
	})
	g1.Draw(core.Context(), core.Coordinate{X: int(core.Stdout().Width()-g1.Width()) / 2, Y: 1})
	<-core.ControlSignal
	g1.Clear()
	<-core.ControlSignal
}
