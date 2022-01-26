package main

import (
	"go-vt100/canvas"
	"go-vt100/color"
	"go-vt100/shape/matrix"
	"go-vt100/shape/point"
	"go-vt100/table"
	"go-vt100/vt100"
	"os"
	"os/signal"
)

func main() {
	controlSignal := make(chan os.Signal)
	signal.Notify(controlSignal, os.Interrupt)

	c := canvas.NewStdoutCanvas(true)
	for y := 1; y <= 5; y++ {
		for x := 1; x <= 10; x++ {
			if y == x {
				c.AddLayerObject(x, y, point.NewPoint('*'))
			}
		}
	}
	c.Draw()
	<-controlSignal
	c.Clear()

	c.SetBackgroundColor(color.Black)

	c.AddLayerObject(5, 5, matrix.NewMatrix(6, 3, color.White))
	c.Draw()
	<-controlSignal
	c.Clear()

	fct := table.NewFixedCellTable(2, 2, "standard out", color.Red, color.Yellow)
	c.AddLayerObject(20, 5, fct)
	c.Draw()
	<-controlSignal
	c.Clear()

	head := []string{"ID", "Value", "Desc"}
	value := [][]string{
		{"1", "202201201529", "Date"},
		{"2", "Hello World", "Msg"},
		{"3", "Mericustar", "Author"},
	}
	act := table.NewAdaptiveCellTable(head, value, color.Red, color.Yellow)
	c.AddLayerObject(20, 5, act)
	c.Draw()
	<-controlSignal
	c.Clear()

	dt := table.NewDecoratedTable(head, value, &table.TableDecoration{
		WidthPadding:  1,
		HeightPadding: 0,
	}, color.Red, color.Yellow)
	c.AddLayerObject(20, 1, dt)
	c.Draw()
	<-controlSignal
	c.Clear()

	// n := tree.NewFactorioTree()
	// tree.Draw(n)

	// fmt.Println("Ctrl + C to exit")
	<-controlSignal
	vt100.MoveCursorToLine(0)
	vt100.ClearScreen()
	vt100.CursorVisible()
	vt100.Reset()
}
