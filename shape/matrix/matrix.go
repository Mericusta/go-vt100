package matrix

import (
	"go-vt100/color"
	"go-vt100/size"
	"go-vt100/tab"
	"go-vt100/vt100"
)

type Matrix struct {
	s size.Size
	c color.Color
}

func NewMatrix(w, h int, c color.Color) Matrix {
	return Matrix{size.Size{Width: w, Height: h}, c}
}

func (m Matrix) Draw(x, y int, s size.Size) {
	vt100.SetBackgroundColor(m.c)
	for _y := y; _y < y+m.s.Height && _y < s.Height; _y++ {
		for _x := x; _x < x+m.s.Width && _x < s.Width; _x++ {
			vt100.MoveCursorToAndPrint(_x, _y, string(tab.Space()))
		}
	}
	vt100.ClearBackgroundColor()
}
