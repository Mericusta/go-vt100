package main

import (
	"go-vt100/table"
	"go-vt100/tree"
)

func main() {
	var t table.Table
	// t = table.NewFixedCellTable(2, 2, "standard out")
	// table.Draw(t)

	head := []string{"ID", "Value", "Desc"}
	value := [][]string{
		{"1", "202201201529", "Date"},
		{"2", "Hello World", "Msg"},
		{"3", "Mericustar", "Author"},
	}
	// t = table.NewAdaptiveCellTable(head, value)
	// table.Draw(t)

	t = table.NewDecoratedTable(head, value, &table.TableDecoration{
		WidthPadding:  1,
		HeightPadding: 1,
	})
	table.Draw(t)

	n := tree.NewFactorioTree()
	tree.Draw(n)
}
