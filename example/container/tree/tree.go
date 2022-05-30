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

	// example 1, just root
	rootDrawable := container.NewTextarea("root node", core.Horizontal)
	rootNode := container.NewTree(&rootDrawable, nil)
	rootNode.Draw(core.Context(), core.Coordinate{})
	<-core.ControlSignal
	rootNode.Clear()
	<-core.ControlSignal

	// example 2
	rootDrawable = container.NewTextarea("root node", core.Horizontal)
	subDrawable1 := shape.NewPoint('❤')
	subDrawable2 := shape.NewLine(subDrawable1, 5, core.Horizontal)
	subDrawable3 := shape.NewRectangle(subDrawable2, 5)
	subDrawable4 := container.NewCanvas(subDrawable3.Size(), false)
	subDrawable4.AppendObjects(core.NewObject(core.Coordinate{}, subDrawable3))
	subDrawable5 := container.NewCanvas(subDrawable3.Size(), true)
	subDrawable5.AppendObjects(core.NewObject(core.Coordinate{}, subDrawable3))
	rootNode = container.NewTree(&rootDrawable, []core.Drawable{
		&subDrawable1,
		&subDrawable2,
		&subDrawable3,
		&subDrawable4,
		&subDrawable5,
	})
	rootNode.Draw(core.Context(), core.Coordinate{})
	<-core.ControlSignal
	rootNode.Clear()
	<-core.ControlSignal
}
