package main

import (
	"go-vt100/table"
)

func main() {
	t := table.Table{
		Col:        2,
		Row:        2,
		CellWidth:  4,
		CellHeight: 1,
	}
	t.Draw()
}
