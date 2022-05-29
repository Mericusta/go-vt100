package main

import (
	"github.com/Mericusta/go-vt100/border"
	"github.com/Mericusta/go-vt100/container"
	"github.com/Mericusta/go-vt100/core"
	"github.com/Mericusta/go-vt100/shape"
)

func main() {
	defer core.Destruct()
	core.ClearScreen()
	core.CursorInvisible()

	// example 1
	header1StrSlice := []string{"HA", "HB", "HC", "HD"}
	header1DrawableSlice := make([]core.Drawable, len(header1StrSlice))
	for i, s := range header1StrSlice {
		t := container.NewTextarea(s, core.Horizontal)
		header1DrawableSlice[i] = &t
	}
	value1StrMap := map[uint]map[uint]string{
		1: {1: "A", 2: "AB", 3: "ABC", 4: "OPERATION"},
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
	table1.Clear()
	<-core.ControlSignal

	// example 2
	header2StrSlice := []string{"HA", "HB", "HC", "HD"}
	header2DrawableSlice := make([]core.Drawable, len(header2StrSlice))
	for i, s := range header2StrSlice {
		t := container.NewTextarea(s, core.Horizontal)
		header2DrawableSlice[i] = &t
	}
	value2StrMap := map[uint]map[uint]string{
		1: {1: "A", 2: "AB", 3: "ABC", 4: "OPERATION"},
		2: {1: string(border.VL()), 2: string(border.HL()), 3: string(border.CT())},
		4: {1: "❤", 3: "❤"},
	}
	value2DrawableMap := make(map[uint]map[uint]core.Drawable, len(value2StrMap))
	for row, colMap := range value2StrMap {
		value2DrawableMap[row] = make(map[uint]core.Drawable, len(colMap))
		for col, s := range colMap {
			t := container.NewTextarea(s, core.Horizontal)
			value2DrawableMap[row][col] = &t
		}
	}
	c := container.NewCanvas(core.Size{Width: 64, Height: 17}, false)
	c.AppendObjects(core.NewObject(
		core.Coordinate{X: 0, Y: 0},
		shape.NewRectangle(shape.NewLine(shape.NewPoint('❤'), 5, core.Horizontal), 5),
	), core.NewObject(
		core.Coordinate{X: 1, Y: 6},
		shape.NewRectangle(shape.NewLine(shape.NewPoint(border.CT()), 5, core.Horizontal), 5),
	), core.NewObject(
		core.Coordinate{X: 2, Y: 12},
		shape.NewRectangle(shape.NewLine(shape.NewPoint('*'), 5, core.Horizontal), 5),
	))
	// value2DrawableMap[5][4] = &c
	value2DrawableMap[5] = map[uint]core.Drawable{4: &c}
	table2 := container.NewTable(header2DrawableSlice, value2DrawableMap)
	table2.Draw(core.Context(), core.Coordinate{X: 0, Y: 0})
	<-core.ControlSignal
	table2.Clear()
	<-core.ControlSignal
}
