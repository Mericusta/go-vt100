package main

import (
	"fmt"

	"github.com/Mericusta/go-vt100/container"
	"github.com/Mericusta/go-vt100/core"
)

var (
	header  = []string{"header1", "header2", "header3"}
	content = map[uint]map[uint]string{
		1: {1: "1,1", 2: "1,2", 3: "1,3"},
		2: {2: "2,2"},
		3: {1: "3,1", 3: "3,3"},
	}
)

func main() {
	core.Init()
	core.ClearScreen()
	core.CursorInvisible()
	headerDrawableSlice := make([]core.Drawable, len(header))
	for i, s := range header {
		t := container.NewTextarea(s, core.Horizontal)
		headerDrawableSlice[i] = &t
	}
	cellDrawableMap := make(map[uint]map[uint]core.Drawable, len(content))
	for row, colMap := range content {
		cellDrawableMap[row] = make(map[uint]core.Drawable, len(colMap))
		for col, s := range colMap {
			t := container.NewTextarea(s, core.Horizontal)
			cellDrawableMap[row][col] = &t
		}
	}
	core.ClearScreen()
	core.CursorInvisible()
	fmt.Println()
	table := container.NewTable(headerDrawableSlice, cellDrawableMap)
	table.Draw(core.Context(), core.Coordinate{X: 0, Y: 0})
	core.MoveCursorToLine(int(table.Height()))
	core.ResetAttribute()
	core.CursorVisible()
}
