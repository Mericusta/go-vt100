package main

import (
	"go-vt100/table"
)

func main() {
	var t table.Table
	// t = table.NewFixedTable("standard out")
	// t.Draw()

	head := []string{"ID", "Value", "Desc"}
	value := [][]string{
		{"1", "202201201529", "Date"},
		{"2", "Hello World", "Msg"},
		{"3", "Mericustar", "Author"},
	}
	t = table.NewAdaptiveTable(head, value)
	t.Draw()
}
