package point

import (
	"github.com/Mericusta/go-vt100/size"
	"github.com/Mericusta/go-vt100/vt100"
)

type Point struct {
	r rune
}

func NewPoint(r rune) Point {
	return Point{r}
}

func (p Point) Draw(x, y int, s size.Size) {
	if x > s.Width || y > s.Height {
		return
	}
	vt100.MoveCursorToAndPrint(x, y, string(p.r))
}
