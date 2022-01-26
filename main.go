package main

import (
	"go-vt100/canvas"
	"go-vt100/color"
	"go-vt100/shape/matrix"
	"go-vt100/table"
	"go-vt100/vt100"
	"os"
	"os/signal"
)

func main() {
	controlSignal := make(chan os.Signal)
	signal.Notify(controlSignal, os.Interrupt)

	// c := canvas.NewVSCodeTerminalCanvas(true)
	c := canvas.NewCanvasWithBoundary(120, 30)
	// for y := 1; y <= 5; y++ {
	// 	for x := 1; x <= 10; x++ {
	// 		if y == x {
	// 			c.AddLayerObject(x, y, point.NewPoint('*'))
	// 		}
	// 	}
	// }
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

	// t = table.NewDecoratedTable(head, value, &table.TableDecoration{
	// 	WidthPadding:  1,
	// 	HeightPadding: 1,
	// })
	// table.Draw(t)

	// n := tree.NewFactorioTree()
	// tree.Draw(n)

	// fmt.Println("Ctrl + C to exit")
	<-controlSignal
	vt100.MoveCursorToLine(0)
	vt100.ClearScreen()
}
