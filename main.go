package main

import (
	"go-vt100/canvas"
	"go-vt100/shape/point"
	"os"
	"os/signal"
)

func main() {
	c := canvas.NewCanvasWithBoundary(10, 5)
	for y := 1; y <= 5; y++ {
		for x := 1; x <= 10; x++ {
			if y == x {
				c.AddLayerObject(x, y, point.NewPoint('*'))
			}
		}
	}

	c.Draw()

	// var t table.Table
	// t = table.NewFixedCellTable(2, 2, "standard out")
	// table.Draw(t)

	// head := []string{"ID", "Value", "Desc"}
	// value := [][]string{
	// 	{"1", "202201201529", "Date"},
	// 	{"2", "Hello World", "Msg"},
	// 	{"3", "Mericustar", "Author"},
	// }
	// t = table.NewAdaptiveCellTable(head, value)
	// table.Draw(t)

	// t = table.NewDecoratedTable(head, value, &table.TableDecoration{
	// 	WidthPadding:  1,
	// 	HeightPadding: 1,
	// })
	// table.Draw(t)

	// n := tree.NewFactorioTree()
	// tree.Draw(n)

	// fmt.Println("Ctrl + C to exit")
	e := make(chan os.Signal)
	signal.Notify(e, os.Interrupt)
	<-e
}
