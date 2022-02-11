package main

import (
	"go-vt100/canvas"
	"go-vt100/color"
	"go-vt100/shape/matrix"
	"go-vt100/shape/point"
	"go-vt100/table"
	"go-vt100/terminal"
	"go-vt100/tree"
	"go-vt100/vt100"
)

// "go-vt100/tree"

func main() {
	canvasWidth := terminal.Stdout().Width() - 1
	canvasHeight := terminal.Stdout().Height() - 4

	<-terminal.ControlSignal
	c := canvas.NewCanvasWithBoundary(canvasWidth, canvasHeight*2)
	for y := 1; y <= 5; y++ {
		for x := 1; x <= 10; x++ {
			if y == x {
				c.AddLayerObject(x, y, point.NewPoint('*'))
			}
		}
	}
	c.Draw()
	<-terminal.ControlSignal
	c.Clear()

	// c.SetBackgroundColor(color.Black)

	c.AddLayerObject(5, 5, matrix.NewMatrix(6, 3, color.White))
	c.Draw()
	<-terminal.ControlSignal
	c.Clear()

	fct := table.NewFixedCellTable(2, 2, "standard out", color.Default, color.Default)
	c.AddLayerObject(20, 5, fct)
	c.Draw()
	<-terminal.ControlSignal
	c.Clear()

	head := []string{"ID", "Value", "Desc"}
	value := [][]string{
		{"1", "202201201529", "Date"},
		{"2", "Hello World", "Msg"},
		{"3", "Mericustar", "Author"},
	}
	act := table.NewAdaptiveCellTable(head, value, color.Default, color.Default)
	c.AddLayerObject(20, 5, act)
	c.Draw()
	<-terminal.ControlSignal
	c.Clear()

	dt := table.NewDecoratedTable(head, value, &table.TableDecoration{
		CellWidthPadding:  1,
		CellHeightPadding: 1,
	}, color.Default, color.Default)
	c.AddLayerObject(20, 1, dt)
	c.Draw()
	<-terminal.ControlSignal
	c.Clear()

	ft := tree.NewFactorioTree(1)
	c.AddLayerObject(1, 1, ft)
	c.Draw()
	<-terminal.ControlSignal
	c.Clear()

	return

	// canvas height might lower than tree height
	// so you have to stretch window height if you want to see all content
	mt := tree.NewMarkdownTree("./resources/factorio.md", "Artillery shell", 4, 1)
	// c.AddLayerObject(1, 1, mt)
	mtn := mt.RootNode()
	ft = tree.NewFactorioTreeWithRootNode(mtn, 1)
	c.AddLayerObject(1, 1, ft)
	c.Draw()
	<-terminal.ControlSignal
	c.Clear()

	content := "Ctrl + C to exit"
	fct = table.NewFixedCellTable(1, 1, content, color.Default, color.Default)
	c.AddLayerObject(canvasWidth/2-len(content)/2, canvasHeight/2, fct)
	c.Draw()
	<-terminal.ControlSignal
	c.Clear()

	vt100.MoveCursorToLine(0)
	vt100.ClearScreen()
	vt100.CursorVisible()
	vt100.Reset()
}
