package point

import "go-vt100/vt100"

type Point struct {
	r rune
}

func NewPoint(r rune) Point {
	return Point{r}
}

func (p Point) Draw(x, y int) {
	vt100.MoveCursorToAndPrint(x, y, string(p.r))
}
