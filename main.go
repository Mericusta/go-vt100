package main

import (
	"go-vt100/canvas"
	"os"
	"os/signal"
)

func main() {
	c := canvas.NewCanvasWithBoundary(10, 5)
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
