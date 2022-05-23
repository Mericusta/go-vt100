package shape

import "github.com/Mericusta/go-vt100/core"

// Point is the unit shape, that is 1x1
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
