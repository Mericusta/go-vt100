package vt100

import (
	"github.com/Mericusta/go-vt100/core"
)

type Point struct {
	r rune
}

func NewPoint(r rune) Point {
	return Point{r}
}

func (p Point) Draw(x, y int, s core.Size) {
	if x > s.Width || y > s.Height {
		return
	}
	core.MoveCursorToAndPrint(x, y, string(p.r))
}
