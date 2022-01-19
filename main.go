package main

import (
	"go-vt100/table"
)

func main() {
	s := "example"

	t := table.Table{
		Col:        2,
		Row:        2,
		CellWidth:  len(s),
		CellHeight: 1,
		Content:    []byte(s),
	}
	t.Draw()
}
