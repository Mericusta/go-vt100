package main

import (
	"github.com/Mericusta/go-vt100/border"
	"github.com/Mericusta/go-vt100/container"
	"github.com/Mericusta/go-vt100/core"
)

func main() {
	defer core.Destruct()
	core.ClearScreen()
	core.CursorInvisible()

	// example 1
	header1StrSlice := []string{"header 1", "header 2", "header 3", "operation"}
	header1DrawableSlice := make([]core.Drawable, len(header1StrSlice))
	for i, s := range header1StrSlice {
		t := container.NewTextarea(s, core.Horizontal)
		header1DrawableSlice[i] = &t
	}
	value1StrMap := map[uint]map[uint]string{
		1: {1: "A", 2: "B", 3: "C", 4: "OP1"},
		2: {1: string(border.VL()), 2: string(border.HL()), 3: string(border.CT())},
		4: {1: "❤", 3: "❤"},
	}
	value1DrawableMap := make(map[uint]map[uint]core.Drawable, len(value1StrMap))
	for row, colMap := range value1StrMap {
		value1DrawableMap[row] = make(map[uint]core.Drawable, len(colMap))
		for col, s := range colMap {
			t := container.NewTextarea(s, core.Horizontal)
			value1DrawableMap[row][col] = &t
		}
	}
	table1 := container.NewTable(header1DrawableSlice, value1DrawableMap)
	table1.Draw(core.Context(), core.Coordinate{X: 0, Y: 0})
	<-core.ControlSignal
}
