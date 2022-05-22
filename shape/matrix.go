package vt100

import (
	"github.com/Mericusta/go-vt100/core"
)

type Matrix struct {
	s core.Size
	c core.Color
}

func NewMatrix(w, h int, c core.Color) Matrix {
	return Matrix{core.Size{Width: w, Height: h}, c}
}

func (m Matrix) Draw(x, y int, s core.Size) {
	core.SetBackgroundColor(m.c)
	for _y := y; _y < y+m.s.Height && _y < s.Height; _y++ {
		for _x := x; _x < x+m.s.Width && _x < s.Width; _x++ {
			core.MoveCursorToAndPrint(_x, _y, string(core.Space()))
		}
	}
	core.ClearBackgroundColor()
}
